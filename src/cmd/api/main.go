package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	ollama "github.com/ollama/ollama/api"
	"github.com/radia78/orpheus/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	// Step 1: Initialize Milvus Database
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	milvusAddr := "localhost:19530" // Will be given via config (LATER)

	milvusClient, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address: milvusAddr,
	})

	if err != nil {
		log.Error(err)
	}

	// Step 2: Initialize Ollama Client
	// For the time being we're just using the default Ollama config, we'll change this later
	ollamaClient, err := ollama.ClientFromEnvironment()
	if err != nil {
		log.Error(err)
	}

	// Step 3: Now the API is live
	var r *chi.Mux = chi.NewRouter()

	handlers.Handler(r, cli)

	fmt.Println("Starting Orpheus...")

	fmt.Println(`
  ____  ____  ____  _   _ _____ _   _ ____  
 / __ \|  _ \|  _ \| | | | ____| | | / ___| 
| |  | | |_) | |_) | |_| |  _| | | | \___ \ 
| |__| |  _ <|  __/|  _  | |___| |_| |___) |
 \____/|_| \_\_|   |_| |_|_____|\___/|____/
 `)
	fmt.Println("Version 0.1.0")

	// Serve the API service
	err = http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

	// Shutting down the milvus and ollama clients
	err = milvusClient.Close(ctx)
	if err != nil {
		log.Error(err)
	}
	err = ollamaClient.Signout(ctx)
	if err != nil {
		log.Error(err)
	}
}
