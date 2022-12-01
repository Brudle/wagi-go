package main

import (
	"fmt"
	"net/http"

	spin "github.com/fermyon/spin/sdk/go/http"
)

func init() {
	spin.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        fmt.Fprintf(w, "\n")
		fmt.Fprintln(w, "Hello, World!")
	})
}

func main() {
}