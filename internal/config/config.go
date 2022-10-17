package config

import (
	"fmt"
	"strings"

	"github.com/ra9dev/go-template/pkg/config"
	"github.com/ra9dev/go-template/pkg/sre"
	"github.com/ra9dev/go-template/pkg/sre/log"
)

const defaultConfigPath = "./config"

var defaultEnvKeyReplacer = strings.NewReplacer(".", "_")

const (
	defaultHTTPPort      = 80
	defaultGRPCPort      = 82
	defaultHTTPAdminPort = 84
	defaultEnv           = "local"
)

const (
	ServiceName    = "go-template"
	ServiceVersion = "1.0.0"
)

type Config struct {
	Env       sre.Env         `mapstructure:"env"`
	LogLevel  log.Level       `mapstructure:"log_level"`
	Ports     PortsConfig     `mapstructure:"ports"`
	DataStore DataStoreConfig `mapstructure:"data_store"`
	Tracing   TracingConfig   `mapstructure:"tracing"`
}

type PortsConfig struct {
	HTTP      uint `mapstructure:"http"`
	GRPC      uint `mapstructure:"grpc"`
	AdminHTTP uint `mapstructure:"admin_http"`
}

type DataStoreConfig struct {
	URL string `mapstructure:"url"`
}

type TracingConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	Endpoint string `mapstructure:"endpoint"`
}

func NewConfig() (Config, error) {
	paths := []string{defaultConfigPath}

	defaultConfig := map[string]any{
		"env":              defaultEnv,
		"ports.http":       defaultHTTPPort,
		"ports.grpc":       defaultGRPCPort,
		"ports.admin_http": defaultHTTPAdminPort,
	}

	defaultOpts := []config.Option{
		config.WithEnvKeyReplacer(defaultEnvKeyReplacer),
	}

	rawCfg, err := config.New(
		config.Params{
			Paths:   paths,
			Options: append([]config.Option{config.WithDefault(defaultConfig)}, defaultOpts...),
		},
	)
	if err != nil {
		return Config{}, fmt.Errorf("could not create config: %w", err)
	}

	baseConfig := Config{}
	if err = rawCfg.Unmarshal(&baseConfig); err != nil {
		return Config{}, fmt.Errorf("could not unmarshal config: %w", err)
	}

	if additionalConfigsParams := getAdditionalConfigs(baseConfig, paths); len(additionalConfigsParams) > 0 {
		for i := range additionalConfigsParams {
			additionalConfigsParams[i].Options = append(additionalConfigsParams[i].Options, defaultOpts...)
		}

		rawCfg, err = config.NewMerged(rawCfg, additionalConfigsParams...)
		if err != nil {
			return Config{}, fmt.Errorf("could not merge configs: %w", err)
		}
	}

	if err = rawCfg.Unmarshal(&baseConfig); err != nil {
		return Config{}, fmt.Errorf("could not unmarshal config: %w", err)
	}

	return baseConfig, nil
}

func getAdditionalConfigs(baseConfig Config, basePaths []string) []config.Params {
	additionalConfigsParams := make([]config.Params, 0)

	if baseConfig.Env != "" {
		envPaths := make([]string, 0, len(basePaths))

		for _, path := range basePaths {
			envPaths = append(envPaths, fmt.Sprintf("%s/%s", path, baseConfig.Env))
		}

		additionalConfigsParams = append(
			additionalConfigsParams,
			config.Params{
				Paths: envPaths,
			},
		)
	}

	return additionalConfigsParams
}
