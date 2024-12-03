package hasher

import (
	"fmt"
	"strings"
)

func (dto HashURLContentInput) Validate() error {
	if strings.TrimSpace(dto.URL) == "" {
		return fmt.Errorf("'URL' is required")
	}

	return nil
}
