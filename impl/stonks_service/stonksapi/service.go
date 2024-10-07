package stonksapi

import (
	"context"
	"fmt"

	"github.com/adzmatheus/stonks-market/model"
	"github.com/adzmatheus/stonks-market/pkg/stonksapi_com"
)

const expense string = "https://github.com/adzmatheus/stonks-market/blob/main/assets/expense.svg"
const income string = "https://github.com/adzmatheus/stonks-market/blob/main/assets/income.svg"

type StonksService struct {
	service *stonksapi_com.Service
}

func NewStonksService(service *stonksapi_com.Service) *StonksService {
	return &StonksService{service}
}

func (s *StonksService) Market(ctx context.Context, ticker string, days int) ([]model.Stonks, error) {

	// Verifica se s.service é nil antes de tentar usá-lo
	if s.service == nil {
		print("service is not initialized")
		return nil, fmt.Errorf("service is not initialized")
	}

	market, err := s.service.Market(ctx, ticker, days)
	if s.service == nil || err != nil {
		print("service is not initialized")
		return nil, err
	}

	stonkses, err := toStonkses(*market)
	if err != nil {
		return nil, err
	}

	return stonkses, nil
}

func toStonkses(market stonksapi_com.Market) ([]model.Stonks, error) {

	var stonkses []model.Stonks

	for _, result := range market.Results {
		stonks := marketResultToStonks(result)
		stonkses = append(stonkses, stonks)
	}

	return stonkses, nil
}

func marketResultToStonks(result stonksapi_com.Result) model.Stonks {

	return model.Stonks{
		Currency:                   result.Currency,
		ShortName:                  result.ShortName,
		LongName:                   result.LongName,
		RegularMarketPrice:         result.RegularMarketPrice,
		RegularMarketChangePercent: result.RegularMarketChangePercent,
		Symbol:                     result.Symbol,
		RegularMarketPreviousClose: result.RegularMarketPreviousClose,
		Logourl:                    result.LogoURL,
		Icon:                       findIcon(result.RegularMarketPreviousClose, result.RegularMarketPrice),
	}
}

func findIcon(lastPrice float64, updatedPrice float64) string {
	iconAddress := expense
	if lastPrice < updatedPrice {
		iconAddress = income
	}
	return iconAddress
}
