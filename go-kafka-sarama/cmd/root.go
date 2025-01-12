package cmd

import (
	"os"

	"kafka-sarama-example/pkg"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Kafka test suite",
	Short: "A tool to test and debug kafka cluster",
	Long:  "A tool to test and debug kafka cluster",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().
		StringSliceVarP(&pkg.AppConfig.BootstrapBrokers, "bootstrap-brokers", "b", []string{"localhost:9092"}, "Bootstrap brokers to Kafka")
}
