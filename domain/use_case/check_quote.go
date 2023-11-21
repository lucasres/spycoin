package use_case

import (
	"context"
	"fmt"
	"log"

	"github.com/lucasresone/spycoin/domain/repositories"
	"github.com/lucasresone/spycoin/domain/service"
	"github.com/martinlindhe/notify"
)

type CheckQuote struct {
	QuoteService        service.QuoteCriptoService
	OperationRepository repositories.OperationRepository
}

func (c *CheckQuote) Execute(ctx context.Context, symbol string) error {
	q, err := c.QuoteService.Quote(ctx, "BINANCE_SPOT_BTC_BRL")
	if err != nil {
		log.Fatalf("erro when quote: %v", err)
	}

	o, err := c.OperationRepository.LastOperation(ctx, symbol, "BUY")
	if err != nil {
		return err
	}

	if o == nil {
		notify.Notify(
			"spycoin",
			"Cotação BTC_BRL",
			fmt.Sprintf(
				"ultima de %s cotação foi: R$%.2f",
				q.Symbol,
				float64(q.Price/100),
			),
			"",
		)
	} else {
		if q.Price > o.Price {
			notify.Notify(
				"spycoin",
				"Cotação BTC_BRL",
				fmt.Sprintf(
					"ultima de %s cotação foi: R$%.2f, ultima compra foi de R$%.2f",
					q.Symbol,
					float64(q.Price/100),
					float64(o.Price/100),
				),
				"",
			)
		}
	}

	return nil
}

func NewCheckQuote(q service.QuoteCriptoService, r repositories.OperationRepository) *CheckQuote {
	return &CheckQuote{
		QuoteService:        q,
		OperationRepository: r,
	}
}
