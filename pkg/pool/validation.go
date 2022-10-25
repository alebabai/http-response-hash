package pool

import (
	"fmt"
)

func (p Pool[T, V]) Validate() error {
	if p.action == nil {
		return fmt.Errorf("'action' is required")
	}

	if p.consumer == nil {
		return fmt.Errorf("'consumer' is required")
	}

	if p.size <= 0 {
		return fmt.Errorf("'size' has invalid value: should be positive number")
	}

	return nil
}
