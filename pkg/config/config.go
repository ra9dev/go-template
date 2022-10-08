package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	DefaultName      = "config"
	DefaultExtension = "yaml"
	DefaultPath      = "."
)

type (
	Config struct {
		driver *viper.Viper
	}

	Params struct {
		Name      string
		Extension string
		Paths     []string
		Options   []Option
	}
)

func New(params Params) (Config, error) {
	if params.Name == "" {
		params.Name = DefaultName
	}

	if params.Extension == "" {
		params.Extension = DefaultExtension
	}

	if len(params.Paths) == 0 {
		params.Paths = append(params.Paths, DefaultPath)
	}

	driver := viper.New()
	driver.AutomaticEnv()
	driver.AllowEmptyEnv(true)
	driver.SetConfigName(params.Name)      // name of config file (without extension)
	driver.SetConfigType(params.Extension) // REQUIRED if the config file does not have the extension in the name

	for _, path := range params.Paths {
		driver.AddConfigPath(path)
	}

	for _, opt := range params.Options {
		opt(driver)
	}

	if err := driver.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("could not read config from file: %w", err)
	}

	return Config{driver: driver}, nil
}

func (c Config) Unmarshal(value any, opts ...viper.DecoderConfigOption) error {
	if err := c.driver.Unmarshal(value, opts...); err != nil {
		return fmt.Errorf("failed to unmarshal config of type %T: %w", value, err)
	}

	return nil
}

func (c Config) GetSettings() map[string]any {
	return c.driver.AllSettings()
}

func (c Config) MergeInConfig(other Config) error {
	settings := other.GetSettings()

	if err := c.driver.MergeConfigMap(settings); err != nil {
		return fmt.Errorf("failed to merge config: %w", err)
	}

	return nil
}
