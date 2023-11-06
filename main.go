package main

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// default
	log.Print("hello world") // {"level":"debug","time":"2023-11-03T18:36:38+08:00","message":"hello world"}

	// level loggin/ only print log above [Info] level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Trace().Msg("trace") // below [Info] level, no print
	log.Debug().Msg("debug") // below [Info] level, no print
	log.Info().Msg("info")   // {"level":"info","time":"2023-11-03T18:36:38+08:00","message":"info"}
	log.Warn().Msg("warn")   // {"level":"warn","time":"2023-11-03T18:36:38+08:00","message":"warn"}
	log.Error().Msg("error") // {"level":"error","time":"2023-11-03T18:36:38+08:00","message":"error"}

	// log.Fatal().Msg("fetal message") // will exit
	// log.Panic().Msg("panic message") // will panic

	// contextual logging (with other fields)
	log.Info().Str("id", "1").Msg("with id") // {"level":"info","id":"1","time":"2023-11-03T18:36:38+08:00","message":"with id"}

	log.Log().Msg("without level") // {"time":"2023-11-03T18:36:38+08:00","message":"without level"}

	// error logging (with "error" field)
	log.Err(errors.New("Error!!")).Msg("error message") // {"level":"error","error":"Error!!","time":"2023-11-03T18:36:38+08:00","message":"error message"}

	// sub logger
	logger := log.With().Str("env", "dev").Logger()
	logger.Info().Str("id", "2").Msg("sub log") // {"level":"info","env":"dev","id":"2","time":"2023-11-03T18:36:38+08:00","message":"sub log"}

	// pretty logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("pretty") // 6:36PM INF pretty

	// stardard logging
	log.Logger = log.Output(os.Stdout)

	// sub dictionary
	log.Info().Str("env", "test").
		Dict("dict", zerolog.Dict().Str("sub", "11").Int("amt", 100)).
		Msg("with dict")
		// {"level":"info","env":"test","dict":{"sub":"11","amt":100},"time":"2023-11-03T18:36:38+08:00","message":"with dict"}

	// caller
	logger = log.With().Caller().Logger()
	logger.Info().Msg("with caller") // {"level":"info","time":"2023-11-03T18:36:38+08:00","caller":"/../go-demo/main.go:50","message":"with caller"}

	// add contextual fields to global logger
	log.Logger = log.With().Str("env", "qat").Logger()
	log.Info().Msg("message") // {"level":"info","env":"qat","time":"2023-11-03T18:36:38+08:00","message":"message"}

}
