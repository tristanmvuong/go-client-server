package main

import (
	"os"
	"net/http"
	"io/ioutil"
	"bytes"
	"fmt"
)

func main() {
	arg := os.Args[1]
	switch arg {
	case "echo":
		_, err := Echo(os.Args[2])
		if err == nil {
			fmt.Println("Got back echo")
		} else {
			fmt.Println("Failed to get back echo")
		}
	}


}

func Echo(msg string) (string, error) {
	resp, err := http.Post("http://localhost:8080/echo", "text/plain", bytes.NewBufferString(msg))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	} else {
		return string(body), nil
	}
}