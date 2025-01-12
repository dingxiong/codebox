package cmd

import (
	"kafka-sarama-example/pkg"

	"github.com/spf13/cobra"
)

var (
	kafkaTopic string
	MsgToSend  string
)

var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Test kafka producer",
	Long:  "Test kafka producer",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.RunProducer(kafkaTopic, MsgToSend)
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)

	producerCmd.PersistentFlags().
		StringVarP(&kafkaTopic, "kafka-topic", "t", "filebeat", "Kafka topic")
	producerCmd.PersistentFlags().
		StringVarP(&MsgToSend, "kafka-message", "m", "xiong-test-kafka-producer", "Message sent to kafka")
}
