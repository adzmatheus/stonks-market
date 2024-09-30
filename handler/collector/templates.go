package collector

import _ "embed"

//go:embed template/daily-market.md.template
var dailyStonksTemplateData string

var templates = []string{
	dailyStonksTemplateData,
}
