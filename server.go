package main

import (
	"log"
	"net/http"
	"time"

	"github.com/pedrazzani/golang/backend-api/config"
	"github.com/pedrazzani/golang/backend-api/router"
)

func main() {

	s := &http.Server{
		Addr:           config.ServerPort,
		Handler:        router.RoutersHandler,
		ReadTimeout:    time.Duration(config.ServerReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.ServerWriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
