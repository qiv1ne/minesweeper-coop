package main

import (
	"net/http"
	"os"

	"github.com/Sinoverg/minesweeper-svc/components"
	"github.com/joho/godotenv"

	"github.com/a-h/templ"
	"github.com/rs/zerolog/log"
)

type MinesweeperService interface {

}

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	mux := http.NewServeMux()
	mux.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	mux.Handle("/",templ.Handler(components.Index()))
	mux.HandleFunc("/play",PlayHandler)
	
	log.Info().Msg("Starting server on port :"+os.Getenv("PORT")+"...")
	err = http.ListenAndServe(":"+os.Getenv("PORT"),mux)
	log.Err(err).Msg("Error ListenAndServer")
}

func PlayHandler(w http.ResponseWriter, r *http.Request){

}