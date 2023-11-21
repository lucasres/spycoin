package service

import (
	"context"

	"github.com/lucasresone/spycoin/domain/entities"
)

type QuoteCriptoService interface {
	Quote(ctx context.Context, symbol string) (*entities.Quote, error)
}
