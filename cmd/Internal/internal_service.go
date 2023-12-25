package Internal

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"serviceMetrica/internal/config"
	"time"
)

func RunInternalService() {
	internalHandler := NewInternalHandler()
	middleware := NewMiddleware()

	router := http.NewServeMux()

	// Apply middleware to the handlers
	registerUserHandler := middleware.AuthenticationMiddleware(http.HandlerFunc(internalHandler.RegisterUser))
	registerAppHandler := middleware.AuthenticationMiddleware(http.HandlerFunc(internalHandler.RegisterApp))
	registerAppEventsHandler := middleware.AuthenticationMiddleware(http.HandlerFunc(internalHandler.RegisterEvents))
	getAPITokenHandler := middleware.AuthenticationMiddleware(http.HandlerFunc(internalHandler.GetAPIToken))
	getAggregatedDataHandler := middleware.AuthenticationMiddleware(http.HandlerFunc(internalHandler.GetAggregatedDataHandler))

	router.Handle("/register/user", registerUserHandler)
	router.Handle("/register/app", registerAppHandler)
	router.Handle("/register/app/events", registerAppEventsHandler)
	router.Handle("/tokens", getAPITokenHandler)
	router.Handle("/get/events", getAggregatedDataHandler)

	server := &http.Server{
		Addr:    config.New().Service.Host + ":" + config.New().Service.Port,
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		log.Printf("Internal Service listening on: %v\n", config.New().Service.Port+":"+config.New().Service.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan

	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
