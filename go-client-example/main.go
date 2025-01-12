package main

func main() {
	logger.Info().Msg("Program starts...")

	// ListNodes()

	// TODO: Not working
	// ApplyYaml()

	if err := PortForward(); err != nil {
		logger.Fatal().Msgf("%v", err)
	}
}
