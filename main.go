package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hameedhub/wschat/handlers"
)

func main() {

	sm := mux.NewRouter()

	// Welcome endpoint
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		handlers.SuccessResponse(w, "Welcome to Chat IO", make([]string, 1))
	})

	sv := &http.Server{
		Addr:           ":9090",
		Handler:        sm,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	sv.ListenAndServe()

}
