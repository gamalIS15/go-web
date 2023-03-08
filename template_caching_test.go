package go_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var tmpl embed.FS

var mytemplates = template.Must(template.ParseFS(tmpl, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	mytemplates.ExecuteTemplate(w, "simple.gohtml", "hello template caching")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
