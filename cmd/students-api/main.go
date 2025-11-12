package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raiashpanda007/go-api-project/pkg/config"
)

func main() {
	configData := config.MustLoad()
	fmt.Println(*configData);
	router := http.NewServeMux()

	router.HandleFunc("GET /health-status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	server := http.Server{
		Addr:    configData.Addr,
		Handler: router,
	}
	fmt.Println("Server running at the server :: ", configData.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Unable to start the server :: %s", err.Error())
	}

}
