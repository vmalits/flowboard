package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/vmalits/taskboard/backend/internal/config"
	"github.com/vmalits/taskboard/backend/internal/http/handler"
	"github.com/vmalits/taskboard/backend/internal/http/router"
	"github.com/vmalits/taskboard/backend/internal/server"
)

func main() {

	cfg := config.Load()

	healthHandler := handler.NewHealthHandler()

	r := router.New(healthHandler)

	srv := server.New(":"+cfg.AppPort, r)

	go func() {
		log.Printf("Server is running on port %s", cfg.AppPort)
		if err := srv.Run(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
