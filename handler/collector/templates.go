package collector

import _ "embed"

//go:embed template/daily-forecast.md.template
var dailyStonksTemplateData string

var templates = []string{
	dailyStonksTemplateData,
}
