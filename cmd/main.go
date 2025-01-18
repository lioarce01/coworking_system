package main

import (
	"cowork_system/internal/interface/http"
	"log"
)

func main() {
	//Initialize router
	r := http.SetupRouter()

	//Initialize server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}