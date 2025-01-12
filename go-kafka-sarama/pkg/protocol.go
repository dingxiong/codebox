package pkg

import (
	"github.com/rs/zerolog/log"
)

func RunProtocol(api KafkaApiKey) {
	switch api {
	case MetaDataApi:
		GetMetaData()
	default:
		log.Fatal().Msgf("Unsupported api key %v", api)
	}
}
