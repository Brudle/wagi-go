package main

import (
	"fmt"
	"net/http"

	spin "github.com/fermyon/spin/sdk/go/http"
)

func init() {
	spin.Handle(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Content-Type: text/plain")
		fmt.Fprintln(w, "Hello, World!")
	})
}

func main() {
}