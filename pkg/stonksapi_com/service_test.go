package stonksapi_com

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarket(t *testing.T) {
	key := os.Getenv("STONKS_API_KEY")
	if key == "" {
		t.Skipf("Missing STONKS_API_KEY")
	}
	service := NewService(key)
	market, err := service.Market(context.Background(), "^BVSP", 5)
	assert.NoError(t, err)
	assert.NotNil(t, market)
	assert.NotEmpty(t, market.Results)
}
