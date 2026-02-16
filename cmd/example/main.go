package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	addr := ":8080"
	router := http.NewServeMux()
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	fmt.Fprintf(os.Stderr, "Listening on %s\n", addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "Shutting down: %s\n", err.Error())
		os.Exit(1)
	}
}
