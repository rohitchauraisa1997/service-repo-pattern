package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

var chiDispatcher = chi.NewRouter()

func (*chiRouter) GET(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Println("Chi HTTP server running on port: ", port)
	http.ListenAndServe(port, chiDispatcher)
}
