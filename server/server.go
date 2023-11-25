package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func getGoogle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /google request\n")
	res, err := http.Get("https://localhost:3333")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("empty", string(body))
	if err != nil {
		fmt.Printf("error reading response body: %s\n", err)
		os.Exit(1)
	}
	io.WriteString(w, string(body))
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/google", getGoogle)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
