package stonksapi

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adzmatheus/stonks-market/pkg/stonksapi_com"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/market.json
var data []byte

func TestToStonkses(t *testing.T) {
	var market *stonksapi_com.Market
	err := json.Unmarshal(data, &market)
	if err != nil {
		panic(err)
	}

	stonkses, err := toStonkses(*market)
	assert.NoError(t, err)
	assert.NotEmpty(t, stonkses)
	for _, stonks := range stonkses {
		assert.NotEmpty(t, stonks.Currency)
		assert.NotEmpty(t, stonks.RegularMarketPrice)
		assert.NotEmpty(t, stonks.ShortName)
		assert.NotEmpty(t, stonks.RegularMarketPreviousClose)
	}
	d, err := json.Marshal(stonkses)
	fmt.Println(string(d))
}
