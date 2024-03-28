package main

import (
	"net/http"
	"time"

	"github.com/ikeohachidi/live-editor/routes"
	"github.com/ikeohachidi/live-editor/services"

	log "github.com/sirupsen/logrus"
)

func init() {
	go services.PruneStaleSessions(time.Hour * 1)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("POST /session", routes.StartSession)
	router.HandleFunc("PUT /session/{id}", routes.UpdateSession)
	router.HandleFunc("GET /session/{id}", routes.FetchSession)
	router.HandleFunc("OPTIONS /session/{id}", routes.HandleOptionsRequest)

	router.HandleFunc("GET /content/{sessionId}", routes.FetchSessionFiles)

	router.HandleFunc("GET /f/{sessionId}/{file}", routes.GetFile)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	log.Info("Server running")
	log.Fatal(server.ListenAndServe())
}
