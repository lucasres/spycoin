package cli

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/lucasresone/spycoin/domain/entities"
	"github.com/lucasresone/spycoin/infra/env"
	"github.com/lucasresone/spycoin/infra/repositories/operations"
	"github.com/spf13/cobra"
)

func NewInsertOperationCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "insert-operation",
		Short: "insert new transaction that buy or sell one crypto",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := env.GetFromCtx[sql.DB](cmd.Context(), env.DBCtx)
			if err != nil {
				return err
			}

			var (
				symbolIn string
				priceIn  string
				typeIn   string
			)

			fmt.Println("Input Symbol(BTC or USDT): ")
			fmt.Scanf("%s", &symbolIn)

			fmt.Println("Input Price(R$ 100 = 10000): ")
			fmt.Scanf("%s", &priceIn)

			fmt.Println("Input Type(BUY or SELL): ")
			fmt.Scanf("%s", &typeIn)

			price, err := strconv.Atoi(priceIn)
			if err != nil {
				return err
			}

			r := operations.NewOperationSqliteRepository(db)
			return r.Insert(cmd.Context(), &entities.Operation{
				Symbol: symbolIn,
				Price:  int64(price),
				Type:   strings.ToUpper(typeIn),
			})
		},
	}
}
