/*
*
TODO:
  - improve error handling
  - improve route handling
*/
package server

import (
	"fmt"
	"net/http"

	"github.com/imattferreira/flag-control/src/entities"
	"github.com/imattferreira/flag-control/src/tools"
)

var flags []*entities.Flag

func getFlags() []*entities.Flag {
	for i := 0; i < 5; i++ {

		flags = append(flags, entities.NewFlag(i, fmt.Sprintf("Flag: %d", i)))
	}

	return flags
}

func createFlag(r *http.Request) (*entities.Flag, error) {
	flag, err := entities.Receive(r.Body)

	if err != nil {
		return nil, err
	}

	flags = append(flags, flag)

	return flag, nil
}

func notFound(w http.ResponseWriter) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}

func internalErr(w http.ResponseWriter) {
	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
}

func Router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	if path == "/flags" && method == "GET" {
		flags := getFlags()
		var expelled []map[string]interface{}

		for _, flag := range flags {
			expelled = append(expelled, flag.Expel())
		}

		body, err := tools.Encode(expelled)

		if err != nil {
			internalErr(w)
			return
		}

		tools.JsonResponse(w, body)
		return
	}

	if path == "/flags" && method == "POST" {
		body, _ := createFlag(r)

		// if err != nil {
		// 	internalErr(w)
		// 	return
		// }

		encoded, _ := tools.Encode(body.Expel())

		tools.JsonResponse(w, encoded)
		return
	}

	notFound(w)
}
