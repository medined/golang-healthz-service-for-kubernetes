package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	port := getEnv("PORT", "10248")

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK\n")
	})
	host := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(host, nil))
}
