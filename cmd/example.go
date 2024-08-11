package main

import (
	"context"
	"log"
	"os"
)

func main() {
	f, err := os.Create("example.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = Example().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
