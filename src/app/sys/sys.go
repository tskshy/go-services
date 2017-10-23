package sys

import (
	"app/routes"
	"libs/log"
	"net/http"
	"time"
)

func Boot() {
	var srv = &http.Server{
		Handler:      routes.Init(),
		Addr:         ":4243",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("server start")
	var _ = srv.ListenAndServe()
}
