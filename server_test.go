package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Server")
	}

	server := http.Server{
		Addr:    "localhost:8001",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8001",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
