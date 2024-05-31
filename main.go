package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/imattferreira/flag-control/src/server"
)

func main() {
	server.Register()
	http.HandleFunc("/", server.Handler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("🚀 Server running at: http://localhost:8080/")

	if err != nil {
		log.Fatal(err)
	}
}
