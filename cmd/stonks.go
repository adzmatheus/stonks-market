package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/adzmatheus/stonks-market/handler/collector"
	"github.com/adzmatheus/stonks-market/impl/stonks_service/stonksapi"
	"github.com/adzmatheus/stonks-market/pkg/stonksapi_com"
	"github.com/spf13/cobra"
)

func UpdateStonks(use string) *cobra.Command {
	var stonksApiComKey string        // Chave/Token da API
	var ticker string                 // Ticker que será buscado pela consulta
	var days int                      // Quantos dias retroativos (lista de valores possíveis)
	var stonksTemplateFilePath string // Caminho do template
	var outputFilePath string         // Saída (README)

	command := &cobra.Command{
		Use: use,
		Run: func(cmd *cobra.Command, args []string) {
			stonksApiService := stonksapi.NewStonksService(stonksapi_com.NewService(stonksApiComKey))
			handler := collector.NewCollector(stonksApiService)
			err := handler.Collect(context.Background(), ticker, days, stonksTemplateFilePath, outputFilePath)
			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			slog.Info("Updated stonks")
		},
	}

	command.Flags().StringVarP(&stonksApiComKey, "stonks-api-key", "k", "", "stonksapi.com API key")
	command.Flags().StringVarP(&stonksTemplateFilePath, "template-file", "f", "", "Readme template file path")
	command.Flags().StringVarP(&outputFilePath, "out-file", "o", "", "Output file path")
	command.Flags().StringVar(&ticker, "ticker", "", "Ticker")
	command.Flags().IntVar(&days, "days", 5, "Days of market")

	err := command.MarkFlagRequired("stonks-api-key")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("template-file")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("ticker")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("out-file")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return command
}
