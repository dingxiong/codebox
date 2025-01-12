package pkg

import (
	default_log "log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

func RunProducer(topic string, msgToSend string) {
	hosts := AppConfig.BootstrapBrokers
	// hosts := []string{"localhost:9092"}
	// hosts := []string{"10.100.173.102:9092"}
	sarama.Logger = default_log.New(os.Stdout, "", default_log.Ltime)
	config := sarama.NewConfig()
	config.ClientID = "test"
	// config.Producer.Retry.Max = 5
	// config.Producer.RequiredAcks = sarama.WaitForAll
	// config.Producer.Return.Successes = true

	config.Producer.RequiredAcks = sarama.RequiredAcks(sarama.WaitForLocal)
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(hosts, config)
	// producer, err := sarama.NewAsyncProducer(hosts, config)
	if err != nil {
		log.Fatal().Msgf("Fail to create a kafka producer: %v", err)
	}
	log.Info().Msgf("Kafka client: %v", producer)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msgToSend),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Error().Msgf("Erros sending message: %v", err)
	} else {
		log.Info().Msgf("Message %s sent: partition=%d, offset=%d", msgToSend, partition, offset)
	}
}
