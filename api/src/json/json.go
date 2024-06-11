package json

import (
	"encoding/json"
	"io"
)

func Encode(data any) ([]byte, error) {
	return json.Marshal(data)
}

func Decode(b io.ReadCloser, v interface{}) error {
	return json.NewDecoder(b).Decode(v)
}
