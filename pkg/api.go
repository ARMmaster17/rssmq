package pkg

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (a *App) registerHTTPRoutes() error {
	a.Router = mux.NewRouter()
	err := a.registerPrometheusEndpoints()
	if err != nil {
		return err
	}
	a.Router.Path("/metrics").Handler(promhttp.Handler())
	a.Router.HandleFunc("/api/feeds", a.HandleGetFeeds).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/api/feed/new", a.HandleCreateFeed).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/api/feed/{id:[0-9]+}/delete", a.HandleDeleteFeed).Methods("POST", "OPTIONS")
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	a.initializeCORS()
	return nil
}

func (a *App) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	var feeds []FeedSource
	result := a.DB.Find(&feeds)
	if result.Error != nil {
		log.Error().Stack().Err(result.Error).Msgf("unable to process request %s", r.RequestURI)
		a.respondWithErrorMessage(w, result.Error.Error())
		return
	}
	a.respondOKWithJSON(w, feeds)
}

func (a *App) HandleCreateFeed(w http.ResponseWriter, r *http.Request) {
	var feed FeedSource
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&feed); err != nil {
		a.respondWithErrorMessage(w, err.Error())
		return
	}
	result := a.DB.Create(&feed)
	if result.Error != nil {
		log.Error().Stack().Err(result.Error).Msgf("unable to process request %s", r.RequestURI)
		a.respondWithErrorMessage(w, result.Error.Error())
		return
	}
	a.respondOKWithJSON(w, nil)
}

func (a *App) HandleDeleteFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fsId, err := strconv.ParseFloat(vars["id"], 64)
	if err != nil {
		a.respondWithErrorMessage(w, err.Error())
		return
	}
	result := a.DB.Delete(&FeedSource{}, fsId)
	if result.Error != nil {
		log.Error().Stack().Err(result.Error).Msgf("unable to process request %s", r.RequestURI)
		a.respondWithErrorMessage(w, result.Error.Error())
		return
	}
}