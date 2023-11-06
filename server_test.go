package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello, ServeMux")
	})

	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "About Page")
	})

	mux.HandleFunc("/profile", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Profile Page")
	})

	mux.HandleFunc("/profile/detail", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Detail Profile Page")
	})

	server := http.Server{
		Addr:    "localhost:8001",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

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
