package collector

import (
	"context"

	"github.com/adzmatheus/stonks-market/model"
)

type StonksService interface {
	Market(ctx context.Context, ticker string, days int) ([]model.Stonks, error)
}
