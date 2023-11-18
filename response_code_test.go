package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if len(name) != 0 {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello, %s", name)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is required")
	}
}

func TestBadRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)

	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}

func TestOKRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?name=Jenadammaya", nil)

	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}
