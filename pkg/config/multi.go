package config

import "fmt"

func NewMerged(base Config, configsParams ...Params) (Config, error) {
	for _, params := range configsParams {
		config, err := New(params)
		if err != nil {
			return Config{}, fmt.Errorf("failed to create config: %w", err)
		}

		if err = base.MergeInConfig(config); err != nil {
			return Config{}, fmt.Errorf("failed to merge config: %w", err)
		}
	}

	return base, nil
}
