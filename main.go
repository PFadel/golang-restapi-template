package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PFadel/golang-restapi-template/app"

	log "github.com/sirupsen/logrus"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s := http.Server{Addr: fmt.Sprintf(":%s", port)}
	app.Start()

	log.Infof("Starting server in %s ...", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Errorf("listen: %s\n", err)
	}

	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Errorf("Server Shutdown: %s", err.Error())
	}
	log.Info("Server exiting")
}
