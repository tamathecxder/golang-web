package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func MultipleQueryParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	var hobbies []string = query["hobby"]

	fmt.Fprintln(writer, strings.Join(hobbies, " | "))
}

func TestMultipleQueryParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000?hobby=gaming&hobby=wota", nil)

	recorder := httptest.NewRecorder()

	MultipleQueryParameterValues(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	email := request.URL.Query().Get("email")
	otp := request.URL.Query().Get("otp")

	fmt.Fprintf(writer, "Hi, %s, this is your OTP code: %s", email, otp)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000?email=asep@gmail.com&otp=290192", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func SaySalam(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello, %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000?name=jamal", nil)

	recorder := httptest.NewRecorder()

	SaySalam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Go")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
