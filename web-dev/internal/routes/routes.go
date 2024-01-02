package routes

import (
	"ex/server/internal/templates/hello"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/api/data", apiDataHandler)

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	component := hello.Hello("mom")
	handler := templ.Handler(component)

	handler.ServeHTTP(w, r)
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API data page, Dave")
}
