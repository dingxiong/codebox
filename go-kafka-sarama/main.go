package main

import (
	// log2 "log"
	"kafka-sarama-example/cmd"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("Program starts...")

	cmd.Execute()

	/*

	 */
}
