package main

import (
	"net/http"
	"time"
)

func main()  {


	sm := http.NewServeMux()

	sv := &http.Server{
		Addr: ":8080",
		Handler: sm,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	sv.ListenAndServe()


}