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
	var stonksApiComKey string
	var crypto string
	var days int
	var stonksTemplateFilePath string
	var outputFilePath string

	command := &cobra.Command{
		Use: use,
		Run: func(cmd *cobra.Command, args []string) {
			stonksApiService := stonksapi.NewStonksService(stonksapi_com.NewService(stonksApiComKey))
			handler := collector.NewCollector(stonksApiService)
			err := handler.Collect(context.Background(), crypto, days, stonksTemplateFilePath, outputFilePath)
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
	command.Flags().StringVar(&crypto, "crypto", "", "Crypto")
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
	err = command.MarkFlagRequired("crypto")
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
