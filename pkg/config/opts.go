package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Option func(v *viper.Viper)

func WithEnvPrefix(prefix string) Option {
	return func(v *viper.Viper) {
		v.SetEnvPrefix(prefix)
	}
}

func WithEnvKeyReplacer(r *strings.Replacer) Option {
	return func(v *viper.Viper) {
		v.SetEnvKeyReplacer(r)
	}
}

func WithDefault(data map[string]any) Option {
	return func(v *viper.Viper) {
		for key, val := range data {
			v.SetDefault(key, val)
		}
	}
}
