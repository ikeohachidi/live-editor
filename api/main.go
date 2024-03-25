package main

import (
	"net/http"

	"github.com/ikeohachidi/live-editor/routes"

	log "github.com/sirupsen/logrus"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("POST /session/{id}", routes.StartSession)
	router.HandleFunc("PUT /session/{id}", routes.UpdateSession)
	router.HandleFunc("GET /session/{id}", routes.WriteContent)
	router.HandleFunc("OPTIONS /session/{id}", routes.HandleOptionsRequest)

	router.HandleFunc("GET /content/{sessionId}", routes.GetFileContent)

	router.HandleFunc("GET /f/{sessionId}/{file}", routes.GetFile)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	log.Info("Server running")
	log.Fatal(server.ListenAndServe())
}
