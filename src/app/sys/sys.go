package sys

import (
	"app/routes"
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

	var _ = srv.ListenAndServe()
}
