package model

type Stonks struct {
	Currency                   string  `json:"currency"`                   // Moeda de referência
	ShortName                  string  `json:"shortName"`                  // Nome da empresa
	LongName                   string  `json:"longName"`                   // Nome longo da empresa
	RegularMarketPrice         float64 `json:"regularMarketPrice"`         // Preço atual
	RegularMarketChangePercent float64 `json:"regularMarketChangePercent"` // Variação percentual do preço diário
	Symbol                     string  `json:"symbol"`                     // Ticker na bolsa
	RegularMarketPreviousClose float64 `json:"regularMarketPreviousClose"` // Preço de fechamento do dia anterior
	Logourl                    string  `json:"logourl"`                    // URL do logo da empresa
	Icon                       string  `json:"icon"`                       // Ícone da situação em relação ao dia anterior (subiu ou desceu)
}
