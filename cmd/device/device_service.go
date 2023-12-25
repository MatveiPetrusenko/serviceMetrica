package device

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"serviceMetrica/internal/config"
	"time"

	"github.com/gorilla/mux"
)

func RunDeviceService() {
	deviceAPIHandler := NewDeviceHandler()
	deviceAPIMiddleware := NewMiddleware()

	router := mux.NewRouter()

	router.HandleFunc("/tokens/jwt", deviceAPIHandler.GetJWTTokenHandler).Methods("POST")
	router.Handle("/events", deviceAPIMiddleware.AuthMiddleware(http.HandlerFunc(deviceAPIHandler.SendEventHandler))).Methods("POST")

	server := &http.Server{
		Addr:    config.New().Device.Host + ":" + config.New().Device.Port,
		Handler: router,
	}

	// Запуск HTTP-сервера в отдельной горутине
	go func() {
		log.Printf("Device Service listening on: %v\n", config.New().Device.Host+":"+config.New().Device.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Ожидание сигнала завершения работы приложения
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Создание контекста для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Остановка HTTP-сервера с graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server gracefully stopped")
}
