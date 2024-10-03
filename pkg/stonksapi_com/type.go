package stonksapi_com

import "time"

const APIEndpoint = "https://brapi.dev/api/quote/"

type Result struct {
	Currency                   string                `json:"currency"`
	ShortName                  string                `json:"shortName"`
	LongName                   string                `json:"longName"`
	RegularMarketChange        float64               `json:"regularMarketChange"`
	RegularMarketChangePercent float64               `json:"regularMarketChangePercent"`
	RegularMarketTime          time.Time             `json:"regularMarketTime"`
	RegularMarketPrice         float64               `json:"regularMarketPrice"`
	RegularMarketDayHigh       float64               `json:"regularMarketDayHigh"`
	RegularMarketDayRange      string                `json:"regularMarketDayRange"`
	RegularMarketDayLow        float64               `json:"regularMarketDayLow"`
	RegularMarketVolume        int64                 `json:"regularMarketVolume"`
	RegularMarketPreviousClose float64               `json:"regularMarketPreviousClose"`
	RegularMarketOpen          float64               `json:"regularMarketOpen"`
	FiftyTwoWeekRange          string                `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekLow            float64               `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh           float64               `json:"fiftyTwoWeekHigh"`
	Symbol                     string                `json:"symbol"`
	UsedInterval               string                `json:"usedInterval"`
	UsedRange                  string                `json:"usedRange"`
	HistoricalDataPrices       []HistoricalDataPrice `json:"historicalDataPrice"`
	ValidRanges                []string              `json:"validRanges"`
	ValidIntervals             []string              `json:"validIntervals"`
	PriceEarnings              *float64              `json:"priceEarnings"`
	EarningsPerShare           *float64              `json:"earningsPerShare"`
	LogoURL                    string                `json:"logourl"`
}

type HistoricalDataPrice struct {
	Date          int64   `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	Volume        int64   `json:"volume"`
	AdjustedClose float64 `json:"adjustedClose"`
}

type Market struct {
	Results     []Result  `json:"results"`
	RequestedAt time.Time `json:"requestedAt"`
	Took        string    `json:"took"`
}
