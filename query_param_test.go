package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s ", name)
	}
}

func TestQuery(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func MultiQuery(writer http.ResponseWriter, request *http.Request) {
	firstname := request.URL.Query().Get("firstname")
	lastname := request.URL.Query().Get("lastname")
	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestMultiQuery(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?firstname=Eko&lastname=Wayne", nil)
	recorder := httptest.NewRecorder()

	MultiQuery(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultiParamValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprintf(writer, strings.Join(names, " "))
}

func TestMutliParamValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Eko&name=Gamal&name=thomas", nil)
	recorder := httptest.NewRecorder()

	MultiParamValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
