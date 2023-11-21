package env

import (
	"context"
	"fmt"
)

type ctxKeys string

const DBCtx ctxKeys = "__db"

func GetFromCtx[T any](ctx context.Context, key ctxKeys) (*T, error) {
	val := ctx.Value(key)

	parsed, ok := val.(*T)
	if !ok {
		return nil, fmt.Errorf("erro when retive \"%s\" key from ctx", key)
	}

	return parsed, nil
}
