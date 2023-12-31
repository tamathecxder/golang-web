package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func IfTemplateAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(writer, "if.gohtml", map[string]interface{}{
		"Name": "Testing",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	IfTemplateAction(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ComparisonTemplateAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparison.gohtml"))

	t.ExecuteTemplate(writer, "comparison.gohtml", map[string]interface{}{
		"Title": "Comparison Operator",
		"Age":   "50",
	})
}

func TestTemplateActionComparison(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	ComparisonTemplateAction(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func RangeTemplateAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Range Operator",
		"Bands": []string{
			"I Prevail", "Alpha Wolf", "The Beatles",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	RangeTemplateAction(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func WithTemplateAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "With Operator",
		"Name":  "Mamat",
		"Address": map[string]interface{}{
			"Street": "Jalan Seksama",
			"City":   "Cianjur",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	WithTemplateAction(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
