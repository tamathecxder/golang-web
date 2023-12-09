package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (page MyPage) SaySalam(name string) string {
	return "Assalamualaikum, " + name + ". My name is " + page.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SaySalam "Agus" }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Asep",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)

	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	result := recorder.Result()

	body, _ := io.ReadAll(result.Body)

	fmt.Print(string(body))
}
