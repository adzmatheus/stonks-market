package collector

import (
	"context"

	"github.com/adzmatheus/stonks-market/model"
)

type StonksService interface {
	Market(ctx context.Context, crypto string, days int) ([]model.Stonks, error)
}
