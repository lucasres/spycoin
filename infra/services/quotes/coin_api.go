package quotes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lucasresone/spycoin/domain/entities"
)

const urlBase = "https://rest.coinapi.io/v1"

type CoinApiQuotesService struct {
	cli *http.Client
}

func (c *CoinApiQuotesService) Quote(ctx context.Context, symbol string) (*entities.Quote, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/quotes/current", urlBase), nil)
	if err != nil {
		return nil, fmt.Errorf("cant create request: %w", err)
	}

	q := url.Values{}
	q.Add("filter_symbol_id", "BINANCE_SPOT_BTC_BRL")
	req.Header.Add("X-CoinAPI-Key", "")
	req.URL.RawQuery = q.Encode()

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cant do request at coinapi: %w", err)
	}
	defer resp.Body.Close()

	var (
		decoded []quoteResponse
		entity  *entities.Quote
	)

	err = json.NewDecoder(resp.Body).Decode(&decoded)
	if err != nil {
		return nil, fmt.Errorf("cant decode response of quote coinapi: %w", err)
	}

	for _, k := range decoded {
		if k.SymbolID == symbol {
			entity = &entities.Quote{
				Symbol: k.SymbolID,
				Price:  int64(k.BidPrice) * 100,
			}
		}
	}

	return entity, nil
}

func NewCoinApiQuotesService() *CoinApiQuotesService {
	return &CoinApiQuotesService{
		cli: &http.Client{
			Timeout: time.Second * 60,
		},
	}
}
