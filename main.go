package main

import (
	"flag"

	"github.com/comhttp/our/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	path := flag.String("path", "", "Path")

	loglevel := flag.String("loglevel", "debug", "Logging level (debug, info, warn, error)")
	flag.Parse()

	//j.Log.SetLevel(parseLogLevel(*loglevel))
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Default level for this example is info, unless debug flag is present

	switch *loglevel {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	//log.Debug().Msg("This message appears only when log level set to Debug")
	//log.Info().Msg("This message appears when log level set to Debug or Info")

	o := app.NewOUR(*path)
	o.OurSRV()
	//j.ServicesSRV(*service, *port, *coin)
	log.Info().Msg("Starting Ouroboros...")

	//log.Fatal(j.WWW.ListenAndServe())
}

//
/*
jdbs -bind 192.168.192.99:14477 -dbdir jdbinfo -loglevel info &&
jdbs -bind 192.168.192.99:14488 -dbdir jdbcoins -loglevel info &&
jdbs -bind 192.168.192.99:14489 -dbdir jdbnodes -loglevel info &&
jdbs -bind 192.168.192.99:14499 -dbdir jdbexchanges -loglevel info &&
jdbs -bind 192.168.192.99:15502 -dbdir jdbparallelcoin -loglevel info
*/
//systemctl restart jdbparallelcoin && systemctl restart jdbcoins && systemctl restart jdbnodes && systemctl restart jdbexchanges && systemctl restart jdbinfo
//systemctl restart jorm && systemctl restart enso && systemctl restart comhttp && systemctl restart our && systemctl restart reverseproxy

//systemctl stop jdbparallelcoin && systemctl stop jdbcoins && systemctl stop jdbnodes && systemctl stop jdbexchanges && systemctl stop jdbinfo
//systemctl stop jorm && systemctl stop enso && systemctl stop our && systemctl stop comhttp  && systemctl stop reverseproxy  && systemctl stop parallelcoinexplorer.service  && systemctl stop okno.service
//systemctl restart jorm && systemctl restart enso && systemctl restart our && systemctl restart comhttp  && systemctl restart reverseproxy  && systemctl restart parallelcoinexplorer.service  && systemctl restart okno.service
//    os.Setenv("EDITOR", "emacs")
