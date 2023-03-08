package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	mytemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escapte",
		"Body":  "<p>Ini body</p>",
	})
}

func TemplateAutoEscapeDis(w http.ResponseWriter, r *http.Request) {
	mytemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escapte",
		"Body":  template.HTML("<p>Ini body</p>"),
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestTemplateEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestTemplateAutoEscapeDis(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDis(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestTemplateEscapeServerDis(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDis),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXss(w http.ResponseWriter, r *http.Request) {
	mytemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto XSS",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXss(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?body=<p>HelloWorld</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXss(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestTemplateEscapeServerXss(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXss),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
