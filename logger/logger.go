package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Log zerolog.Logger
	Lig zerolog.Logger
)

func InitZeroLog(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
	Log = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	Lig = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
