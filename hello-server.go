package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Server started. %s = %s \n", "environment", os.Getenv("environment"))

	name, exist := os.LookupEnv("NAME")
	if !exist {
		name = "sunshine"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got a request")
		fmt.Fprintf(w, "Hello, " + name + "!\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
