package collector

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/adzmatheus/stonks-market/model"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/stonkses.json
var stonksesData []byte

func TestGenerateReadme(t *testing.T) {
	var stonkses []model.Stonks
	err := json.Unmarshal(stonksesData, &stonkses)
	if err != nil {
		panic(err)
	}

	// Construct the path to data/test.txt relative to the test file
	readme, err := generateOutput(stonkses, "../../template/README.md.template")
	assert.NoError(t, err)
	assert.NotNil(t, readme)
	assert.NotEmpty(t, *readme)
	os.WriteFile(".README.md", []byte(*readme), 0644)
}
