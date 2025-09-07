package logger

import (
	"github.com/rs/zerolog/log"
)

func Error(err error, message string) {
	log.Error().Err(err).Msg(message)
}

func Errorf(err error, message string, args ...interface{}) {
	log.Error().Err(err).Msgf(message, args...)
}

func Info(message string) {
	log.Info().Msg(message)
}

func Infof(message string, args ...interface{}) {
	log.Info().Msgf(message, args...)
}

func Debug(message string) {
	log.Debug().Msg(message)
}

func Debugf(message string, args ...interface{}) {
	log.Debug().Msgf(message, args...)
}

func Warn(message string) {
	log.Warn().Msg(message)
}

func Warnf(message string, args ...interface{}) {
	log.Warn().Msgf(message, args...)
}

func Fatal(err error, message string) {
	log.Fatal().Err(err).Msg(message)
}

func Fatalf(err error, message string, args ...interface{}) {
	log.Fatal().Err(err).Msgf(message, args...)
}
