package server

import "net/http"

func notFound(w http.ResponseWriter) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}

func internalErr(w http.ResponseWriter) {
	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
}
