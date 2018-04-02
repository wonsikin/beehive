package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wonsikin/beehive/src/server/handlers"
)

// NewServer returns a new HTTP server
func NewServer(address string) *http.Server {
	handler := Router()

	return &http.Server{
		Addr:    address,
		Handler: handler,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

// Router returns a new router
func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/bhq/message", handlers.NewEntry().PostMessage)

	return r
}
