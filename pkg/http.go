package pkg

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"net/http"
)

func (a *App) respondWithError(w http.ResponseWriter) {
	a.respondWithErrorMessage(w, "Internal server error")
}

func (a *App) respondWithErrorMessage(w http.ResponseWriter, message string) {
	a.respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": message})
}

func (a *App) respondOKWithJSON(w http.ResponseWriter, payload interface{}) {
	a.respondWithJSON(w, http.StatusOK, payload)
}

func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func (a *App) initializeCORS() {
	corsAO := handlers.AllowedOrigins([]string{"*"})
	corsAM := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsAH := handlers.AllowedHeaders([]string{"Content-Type"})
	a.CORS = []handlers.CORSOption{corsAO, corsAM, corsAH}
}