package config

import "github.com/spf13/viper"

type Option func(v *viper.Viper)

func WithDefault(data map[string]any) Option {
	return func(v *viper.Viper) {
		for key, val := range data {
			v.SetDefault(key, val)
		}
	}
}
