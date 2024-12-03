package hasher

import (
	"context"
)

type Service interface {
	HashURLContent(ctx context.Context, in HashURLContentInput) (*HashURLContentOutput, error)
}
