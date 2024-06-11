/*
*
TODO: improve error handling
*/
package server

import "net/http"

type HttpMethods string

const (
	GET  HttpMethods = "GET"
	POST HttpMethods = "POST"
)

type Route struct {
	Path    string
	Method  HttpMethods
	Handler func(http.ResponseWriter, *http.Request)
}

var routes []*Route

func NewRoute(path string, method HttpMethods, handler func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{
		path,
		method,
		handler,
	}
}

func Register() {
	routes = append(routes, NewRoute("/flags", GET, getFlags))
	routes = append(routes, NewRoute("/flags", POST, createFlag))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	handleCors(&w)

	for _, route := range routes {
		if route.Path == r.URL.Path && route.Method == HttpMethods(r.Method) {
			route.Handler(w, r)
			return
		}
	}

	notFound(w)
}
