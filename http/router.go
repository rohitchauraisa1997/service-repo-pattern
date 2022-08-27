package router

import "net/http"

// seperated Router from Mux-Router to enable ease of use for different httprouters
// and prevent cluttered code.
type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
