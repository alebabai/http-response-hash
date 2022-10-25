package hasher

import "fmt"

func (h Hasher) Validate() error {
	if h.client == nil {
		return fmt.Errorf("'client' is required")
	}

	if h.hash == nil {
		return fmt.Errorf("'hash' is required")
	}

	return nil
}
