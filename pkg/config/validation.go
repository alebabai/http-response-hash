package config

import (
	"fmt"
)

func (cfg Config) Validate() error {
	if cfg.Parallel <= 0 {
		return fmt.Errorf("'parallel' has invalid value: should be positive number")
	}

	return nil
}
