package cmd

import (
	"kafka-sarama-example/pkg"

	"github.com/spf13/cobra"
)

var (
	apiKey string
)

var protocolCmd = &cobra.Command{
	Use:   "protocol",
	Short: "Test kafka protocol",
	Long:  "Test kafka protocol",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.RunProtocol(pkg.KafkaApiKey(apiKey))
	},
}

func init() {
	rootCmd.AddCommand(protocolCmd)

	protocolCmd.PersistentFlags().
		StringVarP(&apiKey, "kafka-api-key", "a", string(pkg.MetaDataApi), "Kafka protocol api key")
}
