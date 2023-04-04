package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get environment variables
	service1URL := os.Getenv("SERVICE1_URL")
	service2URL := os.Getenv("SERVICE2_URL")

	// Define HTTP handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the main service!")
	})

	http.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get(service1URL)
		if err != nil {
			fmt.Fprintln(w, "Error calling service1:", err)
			return
		}
		defer response.Body.Close()

		fmt.Fprintln(w, "Response from service1:", response.StatusCode)
	})

	http.HandleFunc("/service2", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get(service2URL)
		if err != nil {
			fmt.Fprintln(w, "Error calling service2:", err)
			return
		}
		defer response.Body.Close()

		fmt.Fprintln(w, "Response from service2:", response.StatusCode)
	})

	// Start HTTP server
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
