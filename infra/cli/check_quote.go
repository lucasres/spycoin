package cli

import (
	"database/sql"
	"errors"

	"github.com/lucasresone/spycoin/domain/use_case"
	"github.com/lucasresone/spycoin/infra/env"
	"github.com/lucasresone/spycoin/infra/repositories/operations"
	"github.com/lucasresone/spycoin/infra/services/quotes"
	"github.com/spf13/cobra"
)

func NewCheckQuoteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "check-quote",
		Short: "Check last cotation and compare with last buy",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("invalid args need 1 must be a symbol of crypto to spy [ex: btc]")
			}

			db, err := env.GetFromCtx[sql.DB](cmd.Context(), env.DBCtx)
			if err != nil {
				return err
			}

			s := quotes.NewCoinApiQuotesService()
			r := operations.NewOperationSqliteRepository(db)

			uc := use_case.NewCheckQuote(s, r)
			return uc.Execute(cmd.Context(), args[0])
		},
	}
}
