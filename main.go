package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofoody/consumer-service/pkg/config"
	"github.com/gofoody/consumer-service/pkg/ctrl"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.New()

	initLogger(config.GetLogLevel())

	router := mountEndpoints()
	startService(config.GetHttpPort(), router)
}

func initLogger(logLevel string) {
	level, _ := log.ParseLevel(logLevel)
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}

func mountEndpoints() *mux.Router {
	r := mux.NewRouter()

	statusCtrl := ctrl.NewStatusCtrl()
	r.HandleFunc("/api/status", statusCtrl.Show)

	consumerCtrl := ctrl.NewConsumerCtrl()
	r.HandleFunc("/api/consumers/{consumerId}", consumerCtrl.Show)
	r.HandleFunc("/api/consumers", consumerCtrl.Create)

	return r
}

func startService(port int, router *mux.Router) {
	addr := fmt.Sprintf("localhost:%d", port)
	log.Infof("consumer service running at:%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("failed to start consumer service, error:%v", err)
	}
}
