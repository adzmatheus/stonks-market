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
		assert.NotEmpty(t, stonks.Country)
		assert.NotEmpty(t, stonks.City)
		assert.NotEmpty(t, stonks.Condition)
		assert.NotEmpty(t, stonks.Timezone)
		assert.NotEmpty(t, stonks.Icon)
		assert.NotEmpty(t, stonks.StartTime)
		assert.NotEmpty(t, stonks.EndTime)
	}
	d, err := json.Marshal(stonkses)
	fmt.Println(string(d))
}
