package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kevin-Bian/BianPhotography2.0/src/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}