package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/bye", handleBye)

	addr := "localhost:8000"
	log.Printf("Listening on %s ... ", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {
	if (request.Method) != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(writer, "Hello-world!")
}

func handleBye(writer http.ResponseWriter, request *http.Request) {
	if (request.Method) != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(writer, "Bye!")
}

func writeResponse(writer http.ResponseWriter, responseString string) {
	response := []byte(responseString)
	fmt.Println(response)
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
