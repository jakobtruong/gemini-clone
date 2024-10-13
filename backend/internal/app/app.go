package app

import (
	"fmt"
	"log"
	"net/http"
)

const PortNumber string = "8080"

func Run() {
	fmt.Println("Starting http server.")
	log.Printf("http server starting here: http://localhost:%s", PortNumber)
	fmt.Println("To close connection, press 'CTRL+C'.")

	err := http.ListenAndServe(":"+PortNumber, nil)
	if err != nil {
		log.Fatal(err)
	}
}
