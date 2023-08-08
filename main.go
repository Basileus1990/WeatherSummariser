package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := setRouting()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Printf("Server is running at port %s\n", server.Addr)
	server.ListenAndServe()
}

func setRouting() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", nil)

	return mux
}
