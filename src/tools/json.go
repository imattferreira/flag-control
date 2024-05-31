package tools

import (
	"encoding/json"
	"io"
	"net/http"
)

func Encode(data any) ([]byte, error) {
	return json.Marshal(data)
}

func Decode(b io.ReadCloser, v interface{}) error {
	return json.NewDecoder(b).Decode(v)
}

func JsonResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
