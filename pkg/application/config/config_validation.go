package config

import (
	"errors"
	"fmt"
)

func (cfg Config) Validate() error {
	var errs []error

	if len(cfg.URLs) <= 0 {
		return fmt.Errorf("'URLs' are required")
	}

	if cfg.Parallel <= 0 {
		return fmt.Errorf("'Parallel' has invalid value: should be positive number")
	}

	return errors.Join(errs...)
}
