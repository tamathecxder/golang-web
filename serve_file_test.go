package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("auth_key") != "" {
		http.ServeFile(writer, request, "./resources/dashboard.html")
	} else {
		http.ServeFile(writer, request, "./resources/401.html")
	}
}

//go:embed resources/dashboard.html
var resourceDashboard string

//go:embed resources/401.html
var resourceUnauthorized string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("auth_key") != "" {
		fmt.Fprint(writer, resourceDashboard)
	} else {
		fmt.Fprint(writer, resourceUnauthorized)
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
func TestServeFileEmbedServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
