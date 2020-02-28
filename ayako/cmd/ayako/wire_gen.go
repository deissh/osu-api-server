// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/deissh/osu-lazer/ayako/app"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// Injectors from main.go:

func Injector() *app.App {
	storeStore := store.Init()
	appApp := app.NewApp(storeStore)
	return appApp
}

// main.go:

var Version string

var Commit string

var Branch string

var BuildTimestamp string

func main() {
	config.WithOptions(config.ParseEnv, config.Readonly)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yaml")
	if err != nil {
		panic(err)
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if config.Bool("debug", false) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
		).With().Caller().Logger()
	}
	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Msg("Starting API")
	log.Debug().
		Msg("Loaded configuration and logger")
	log.Debug().Msg("Start initialize dependencies")
	app2 := Injector()
	log.Debug().Msg("Initialize dependencies successful done")

	if err := app2.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
