package server

import (
	"fmt"
	"net/http"

	Flag "github.com/imattferreira/flag-control/src/entities/flag"
	"github.com/imattferreira/flag-control/src/json"
)

var flags []*Flag.Flag

func getFlags(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5; i++ {
		flags = append(flags, Flag.NewFlag(i, fmt.Sprintf("Flag: %d", i)))
	}

	var expelled []map[string]interface{}

	for _, flag := range flags {
		expelled = append(expelled, flag.Expel())
	}

	body, err := json.Encode(expelled)

	if err != nil {
		internalErr(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func createFlag(w http.ResponseWriter, r *http.Request) {
	flag, err := Flag.Receive(r.Body)

	if err != nil {
		internalErr(w)
	}

	flags = append(flags, flag)
	encoded, eerr := json.Encode(flag.Expel())

	if eerr != nil {
		internalErr(w)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(encoded)
}
