package main

import (
	"github.com/rs/zerolog/log"

	"github.com/joeshaw/envdecode"
	"github.com/nayan9229/ad_prox_dsp/server"
)

var appname = "ad-prox-dsp-service"

var release = "0.0.1"

func main() {
	var cfg server.Config

	err := envdecode.StrictDecode(&cfg)
	if err != nil {
		log.Fatal().Err(err).
			Msg("failed to process environment variables")
	}

	cfg.AppName = appname
	cfg.Release = release
	serv := server.NewServer(&cfg)

	log.Info().
		Str("app", appname).
		Str("release", release).
		Msg("starting server")
	serv.Serve()
}
