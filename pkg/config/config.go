package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	DefaultConfigName      = "config"
	DefaultConfigExtension = "yaml"
	DefaultConfigPath      = "."
)

type Config struct {
	driver *viper.Viper
}

func NewConfig(configName, configExtension string, configPaths []string, opts ...Option) (Config, error) {
	if configName == "" {
		configName = DefaultConfigName
	}

	if configExtension == "" {
		configExtension = DefaultConfigExtension
	}

	if len(configPaths) == 0 {
		configPaths = append(configPaths, DefaultConfigPath)
	}

	driver := viper.NewWithOptions()
	driver.AutomaticEnv()
	driver.AllowEmptyEnv(true)
	driver.SetConfigName(configName)      // name of config file (without extension)
	driver.SetConfigType(configExtension) // REQUIRED if the config file does not have the extension in the name

	for _, path := range configPaths {
		driver.AddConfigPath(path)
	}

	for _, opt := range opts {
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
