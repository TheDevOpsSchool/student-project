package main

import (
	"fmt"
	"os"
)

func main() {
	serviceAURL := os.Getenv("SERVICEA_URL")
	serviceBURL := os.Getenv("SERVICEB_URL")

	// Use serviceAURL and serviceBURL to connect to ServiceA and ServiceB
	// ...
	
	fmt.Println("Connected to ServiceA at", serviceAURL)
	fmt.Println("Connected to ServiceB at", serviceBURL)
}
