package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raiashpanda007/go-api-project/pkg/config"
)

func main() {
	configData := config.MustLoad()
	fmt.Println(*configData)
	router := http.NewServeMux()

	router.HandleFunc("GET /health-status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	server := http.Server{
		Addr:    configData.Addr,
		Handler: router,
	}
	fmt.Println("Server running at the server :: ", configData.Addr)

	// We are doing this inorder to gracefully close the server.

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()

		if err != nil {
			log.Fatalf("Unable to start the server :: %s", err.Error())
		}
	}()

	//Blocking the thread

	<-done

	slog.Info("Shutting down the server ")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("Unable to shutdown the server :: ", slog.String("error", err.Error()))
	}

}
