package config

import (
	"fmt"

	"github.com/ra9dev/go-template/pkg/config"
	"github.com/ra9dev/go-template/pkg/log"
)

const relativeConfigPath = "./config"

const (
	defaultHTTPPort      = 80
	defaultGRPCPort      = 82
	defaultHTTPAdminPort = 84

	ServiceName    = "go-template"
	ServiceVersion = "1.0.0"
)

type Config struct {
	LogLevel    log.Level       `mapstructure:"log_level"`
	Ports       PortsConfig     `mapstructure:"ports"`
	DataStore   DataStoreConfig `mapstructure:"data_store"`
	Tracing     TracingConfig   `mapstructure:"tracing"`
	Environment string          `mapstructure:"environment"`
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
	paths := []string{relativeConfigPath}

	defaultConfig := map[string]any{
		"ports.http":       defaultHTTPPort,
		"ports.grpc":       defaultGRPCPort,
		"ports.admin_http": defaultHTTPAdminPort,
	}

	rawCfg, err := config.NewConfig(
		config.DefaultConfigName,
		config.DefaultConfigExtension,
		paths,
		config.WithDefault(defaultConfig),
	)
	if err != nil {
		return Config{}, fmt.Errorf("could not create config: %w", err)
	}

	cfg := Config{}
	if err = rawCfg.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("could not unmarshal config: %w", err)
	}

	return cfg, nil
}
