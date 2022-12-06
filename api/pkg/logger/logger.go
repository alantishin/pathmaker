package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(level string) *zerolog.Logger {
	skipFrameCount := 3

	pLvl, err := zerolog.ParseLevel(level)

	if err != nil {
		panic(err)
	}

	logger := zerolog.New(os.Stdout).
		Level(pLvl).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).
		Logger()

	return &logger
}
