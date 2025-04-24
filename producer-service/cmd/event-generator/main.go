package main

import (
	"event-generator/iternal/api/message"
	"event-generator/iternal/config"
	messageService "event-generator/iternal/service/order"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	configPath = "config/config.yaml"
)

func main() {
	mainConfig, err := config.InitMainConfig(configPath)
	if err != nil {
		log.Fatalf("loading config error: %s", err)
	}

	brockerList, topikName := mainConfig.KafkaConfigLoad()

	r := chi.NewRouter()

	r.Use()

	messageServ, err := messageService.NewKafkaService(brockerList, topikName)
	if err != nil {
		log.Fatalf("kafka service failed: %s", err)
	}
	messageApi := message.NewMessageImplementation(messageServ)

	r.Post("/data", messageApi.Post)

	go func() {
		time.Sleep(200 * time.Millisecond)
		log.Printf("server listening and serving on %s - succeccful", mainConfig.ServerAdressLoad())
	}()

	if err := http.ListenAndServe(mainConfig.ServerAdressLoad(), r); err != nil {
		log.Fatalf("server failed to listen and serve on localhost:8080 : %s", err)
	}
}
