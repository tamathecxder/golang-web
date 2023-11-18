package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	fullName := request.PostForm.Get("full_name")
	age := request.PostForm.Get("age")

	fmt.Fprintf(writer, "Hi, i am %s, my age is %s", fullName, age)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("full_name=Agus&age=19")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
