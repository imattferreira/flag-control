package tools

import (
	"encoding/json"
	"io"
	"net/http"
)

func Encode(data any) ([]byte, error) {
	return json.Marshal(data)
}

func Decode(b io.ReadCloser, d *map[string]string) error {
	return json.NewDecoder(b).Decode(&d)
}

func JsonResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
