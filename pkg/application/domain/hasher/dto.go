package hasher

import (
	"fmt"
)

type HashURLContentInput struct {
	URL string
}

type HashURLContentOutput struct {
	URL  string
	Sum  []byte
	Size int
}

func (dto HashURLContentOutput) String() string {
	return fmt.Sprintf("%s %x", dto.URL, dto.Sum[:dto.Size])
}
