package sys

import (
	"app/conf"
	"app/routes"
	"fmt"
	"libs/log"
	"net/http"
	"time"
)

func Boot() {
	var srv = &http.Server{
		Handler:      routes.Init(),
		Addr:         fmt.Sprintf("%s:%d", conf.AppConfInfo.HttpHost, conf.AppConfInfo.HttpPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Debug("Run Server:", conf.AppConfInfo.HttpHost, conf.AppConfInfo.HttpPort)
	log.Info("Run Server:", conf.AppConfInfo.HttpHost, conf.AppConfInfo.HttpPort)
	log.Warn("Run Server:", conf.AppConfInfo.HttpHost, conf.AppConfInfo.HttpPort)
	var _ = srv.ListenAndServe()
}
