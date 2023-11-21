package operations

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lucasresone/spycoin/domain/entities"
)

type OperationSqliteRepository struct {
	db *sql.DB
}

func (o *OperationSqliteRepository) Insert(ctx context.Context, op *entities.Operation) error {
	args := []interface{}{op.Price, op.Symbol, op.Type}

	_, err := o.db.ExecContext(
		ctx,
		`
			INSERT INTO operations(price, symbol, type) VALUES (?, ?, ?)
		`,
		args...,
	)

	if err != nil {
		return fmt.Errorf("error when insert operation in database: %w", err)
	}

	log.Println("inserted ok")

	return nil
}

func (o *OperationSqliteRepository) LastOperation(ctx context.Context, symbol, typeOp string) (*entities.Operation, error) {
	args := []interface{}{symbol, typeOp}

	entity := &entities.Operation{}

	err := o.db.QueryRowContext(
		ctx,
		`
			SELECT
				price,
				symbol,
				type
			FROM
				operations
			WHERE
				symbol = ?
				AND type = ?
			ORDER BY id DESC`,
		args...,
	).Scan(&entity.Price, &entity.Symbol, &entity.Type)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("erro when select last operation: %w", err)
	}

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("erro when query last operation: %w", err)
	}

	return entity, nil
}

func NewOperationSqliteRepository(db *sql.DB) *OperationSqliteRepository {
	return &OperationSqliteRepository{
		db: db,
	}
}
