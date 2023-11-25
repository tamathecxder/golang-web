package golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello from `file` template")
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello from `file` template")
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello from `file` template")
}

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(writer, "SIMPLE", "Hello from template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
