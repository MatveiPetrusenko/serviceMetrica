package main

import (
	"log"
	main2 "serviceMetrica/cmd/Internal"
	main1 "serviceMetrica/cmd/device"
	"serviceMetrica/internal/config"
)

func main() {
	err := config.New().Load()
	if err != nil {
		log.Fatalf("Config error: %v\n", err)
	}

	go main1.RunDeviceService()
	go main2.RunInternalService()
	for {
	}
}
