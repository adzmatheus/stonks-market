package collector

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"time"

	"github.com/adzmatheus/stonks-market/model"
	"github.com/adzmatheus/stonks-market/pkg/errs"
)

type Collector struct {
	stonksService StonksService
}

func NewCollector(stonksService StonksService) *Collector {
	return &Collector{stonksService}
}

func (c *Collector) Collect(ctx context.Context, crypto string, days int, templateFilePath string, outFilePath string) error {
	slog.Info(fmt.Sprintf("Collecting stonks for %s for %d days - Template file: %s", crypto, days, templateFilePath))
	stonkses, err := c.stonksService.Market(ctx, crypto, days)
	if err != nil {
		return errs.Joinf(err, "[stonksService.Market]")
	}
	readmeTemplate, err := os.ReadFile(templateFilePath)
	if err != nil {
		return errs.Joinf(err, "[os.ReadFile] "+templateFilePath)
	}
	readme, err := generateOutput(stonkses, string(readmeTemplate), templates...)
	if err != nil {
		return errs.Joinf(err, "[generateOutput]")
	}

	return os.WriteFile(outFilePath, []byte(*readme), 0644)
}

func generateOutput(stonkses []model.Stonks, readmeTemplate string, templates ...string) (*string, error) {
	if len(stonkses) == 0 {
		return nil, errors.New("stonkses must be not empty")
	}
	tmpl, err := template.
		New("readme").
		Funcs(template.FuncMap{
			"formatDate": formatDate,
			"formatHour": formatHour,
			"formatTime": formatTime,
		}).
		Parse(readmeTemplate)
	if err != nil {
		return nil, err
	}

	for _, t := range templates {
		tmpl, err = tmpl.Parse(t)
		if err != nil {
			return nil, err
		}
	}

	var result bytes.Buffer
	err = tmpl.ExecuteTemplate(&result, "readme", map[string]any{
		"Stonkses":    stonkses,
		"UpdatedAt":   time.Now(),
		"TodayStonks": stonkses[0],
	})
	if err != nil {
		return nil, err
	}
	stringResult := result.String()
	return &stringResult, nil
}
