package repositories

import (
	"context"

	"github.com/lucasresone/spycoin/domain/entities"
)

type OperationRepository interface {
	Insert(ctx context.Context, op *entities.Operation) error
	LastOperation(ctx context.Context, symbol, typeOp string) (*entities.Operation, error)
}
