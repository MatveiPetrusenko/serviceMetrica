package main

import (
	"log"
	"serviceMetrica/config"
	"serviceMetrica/device_service"
	internalservice "serviceMetrica/internal_service"
)

func main() {
	err := config.New().Load()
	if err != nil {
		log.Fatalf("Config error: %v\n", err)
	}

	go internalservice.RunInternalService()
	go device_service.RunDeviceService()
	for {
	}
}
