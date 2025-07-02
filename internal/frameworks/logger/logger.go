package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"time"
)

func Init() {
	// Set zerolog global time field format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// Set default level based on the environment variable
	if os.Getenv("SOMBRA_LOGGER_LEVEL") == "DEBUG" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if os.Getenv("SOMBRA_LOGGER_LEVEL") == "INFO" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Configure zerolog to write to console with human-readable output
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
}

func Info(message string) {
	log.Info().Msg(message)
}

func Error(message string, err error) {
	log.Error().Stack().Err(err).Msg(message)
}

func Panic(message string) {
	log.Panic().Stack().Msg(message)
}
