package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/lucasresone/spycoin/infra/cli"
	"github.com/lucasresone/spycoin/infra/env"
	"github.com/lucasresone/spycoin/infra/repositories/sqlite"
	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatalf("erro when open database: %v", err)
	}

	ctx = context.WithValue(ctx, env.DBCtx, db)

	sqlite.BuildDataBase(ctx, db)

	rootCmd := &cobra.Command{
		Use:   "spycoin",
		Short: "programa usado para me ajudar a fazer trade de cripto moedas",
	}

	rootCmd.AddCommand(cli.NewCheckQuoteCommand())
	rootCmd.AddCommand(cli.NewInsertOperationCommand())

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Fatalf("error when execute root command: %v", err)
	}
}
