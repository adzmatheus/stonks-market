package stonksapi

import (
	"context"
	"strings"
	"time"

	"github.com/adzmatheus/stonks-market/model"
	"github.com/adzmatheus/stonks-market/pkg/stonksapi_com"
	"github.com/adzmatheus/stonks-market/pkg/utils"
)

type StonksService struct {
	service *stonksapi_com.Service
}

func NewStonksService(service *stonksapi_com.Service) *StonksService {
	return &StonksService{service}
}

func (s *StonksService) Market(ctx context.Context, city string, days int) ([]model.Stonks, error) {
	market, err := s.service.Market(ctx, city, days)
	if err != nil {
		return nil, err
	}
	stonkses, err := toStonkses(*market)
	if err != nil {
		return nil, err
	}
	return stonkses, nil
}

func toStonkses(market stonksapi_com.Market) ([]model.Stonks, error) {
	timezoneOffset, err := utils.GetTimezoneOffset(market.Location.TzId)
	if err != nil {
		return nil, err
	}

	var stonkses []model.Stonks
	for _, marketDay := range market.Market.Marketday {
		stonks := marketDayToStonks(marketDay, market.Location.Country, market.Location.Name, market.Location.TzId, int64(timezoneOffset.Seconds()))
		stonkses = append(stonkses, stonks)
	}
	return stonkses, nil
}

func marketDayToStonks(marketDay stonksapi_com.MarketDay, country, city, timezone string, timezoneOffset int64) model.Stonks {
	startTime := time.Unix(marketDay.DateEpoch, 0)
	endTime := startTime.Add(time.Hour)
	return model.Stonks{
		Condition:             marketDay.Day.Condition.Text,
		Icon:                  fillImageSchema(marketDay.Day.Condition.Icon),
		StartTime:             &startTime,
		EndTime:               &endTime,
		Country:               country,
		City:                  city,
		Timezone:              timezone,
		TimezoneOffsetSeconds: timezoneOffset,

		AvgTempC: marketDay.Day.AvgtempC,
		MinTempC: marketDay.Day.MintempC,
		MaxTempC: marketDay.Day.MaxtempC,

		AvgWindKph: marketDay.Day.MaxwindKph,
		MinWindKph: marketDay.Day.MaxwindKph,
		MaxWindKph: marketDay.Day.MaxwindKph,
	}
}

func fillImageSchema(imageUrl string) string {
	if strings.HasPrefix(imageUrl, "//") {
		imageUrl = "https:" + imageUrl
	}
	return imageUrl
}
