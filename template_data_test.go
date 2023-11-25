package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Bio struct {
	Title, Name string
	Age         int
	Address     map[string]interface{}
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/bio.gohtml"))

	t.ExecuteTemplate(writer, "bio.gohtml", Bio{
		Title: "Title from struct",
		Name:  "Agus",
		Age:   21,
		Address: map[string]interface{}{
			"Street": "Jalan Example",
		},
	})
}

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/bio.gohtml"))

	t.ExecuteTemplate(writer, "bio.gohtml", map[string]interface{}{
		"Title": "My Bio",
		"Name":  "Tama",
		"Age":   19,
		"Address": map[string]interface{}{
			"Street": "Jalan Test",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
