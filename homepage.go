package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"dalle2/dalle2"

	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the dalle2 text to img!")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	apiKey, ok := os.LookupEnv("OPENAI_API_KEY")

	if !ok {
		log.Fatal("Environment variable OPENAI_API_KEY is not set")
	}
	client, err := dalle2.MakeNewClientV1(apiKey)

	if err != nil {
		log.Fatalf("Error initializing client: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")

	resp, err := client.Create(
		context.Background(),
		"draw a portrait of dhaka in dawn",
		dalle2.WithNumImages(1),
		dalle2.WithSize(dalle2.LARGE),
		dalle2.WithFormat(dalle2.URL),
	)

	json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created: %d\n", resp.Created)

	fmt.Println("Images:")
	for _, img := range resp.Data {
		fmt.Printf("\t%s\n", img.Url)
	}
	fmt.Println("Endpoint Hit: homePage")
}
