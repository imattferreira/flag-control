package Flag

import (
	"io"

	"github.com/imattferreira/flag-control/src/json"
)

type Flag struct {
	id   int
	name string
}

func NewFlag(id int, name string) *Flag {
	return &Flag{
		id,
		name,
	}
}

func Receive(body io.ReadCloser) (*Flag, error) {
	var parsed struct {
		Name string `json:"name"`
	}

	err := json.Decode(body, &parsed)

	if err != nil {
		return nil, err
	}

	return NewFlag(10, parsed.Name), nil
}

func (flag *Flag) Expel() map[string]interface{} {
	var expelled map[string]interface{} = make(map[string]interface{})

	expelled["id"] = flag.id
	expelled["name"] = flag.name

	return expelled
}

func (flag *Flag) GetName() string {
	return flag.name
}
