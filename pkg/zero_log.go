package pkg

import (
	"os"

	"github.com/rs/zerolog"
)

func ZeroLog() {
	file, err := os.OpenFile("go.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := zerolog.New(file).With().Timestamp().Logger()

	logger.Info().Msg("Application Started")
}
