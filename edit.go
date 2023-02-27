package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Andrew-peng/go-dalle2/dalle2"
	"github.com/joho/godotenv"
)

func edit(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the dalle2 image edit with commands!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		log.Fatal("Environment variable OPENAI_API_KEY is not set")
	}

	w.Header().Set("Content-Type", "application/json")

	client, err := dalle2.MakeNewClientV1(apiKey)
	if err != nil {
		log.Fatalf("Error initializing client: %s", err)
	}

	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Read image files
	imgPath := filepath.Join(curDir, "portrait.png")
	img, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	imgBytes, err := io.ReadAll(img)
	if err != nil {
		log.Fatal(err)
	}
	maskPath := filepath.Join(curDir, "mask.png")
	mask, err := os.Open(maskPath)
	if err != nil {
		log.Fatal(err)
	}
	maskBytes, err := io.ReadAll(mask)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Edit(
		context.Background(),
		imgBytes,
		maskBytes,
		"A rainy view of Dhaka during the sunset, watercolor",
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

	fmt.Println("Endpoint Hit: edit of image")
}
