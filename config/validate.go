package config

import (
	"fmt"
)

func (cfg Config) Validate() error {
	if cfg.Parallel <= 0 || cfg.Parallel > len(cfg.Inputs) {
		return fmt.Errorf("parallel should be positive and less than or equal to the number of inputs")
	}

	return nil
}
