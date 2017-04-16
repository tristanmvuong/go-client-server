package main

import (
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	http.Handle("/echo", http.HandlerFunc(EchoHandler))
	fmt.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func EchoHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err == nil {
		fmt.Fprintf(w, string(body))
		fmt.Println("Responding with echo", string(body))
	} else {
		http.Error(w, "Failed to send echo", http.StatusInternalServerError)
	}
}
