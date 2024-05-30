package entities

import (
	"io"

	"github.com/imattferreira/flag-control/src/tools"
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

func Receive(body *io.ReadCloser) (*Flag, error) {
	var data map[string]string

	err := tools.Decode(*body, &data)

	if err != nil {
		return nil, err
	}

	return NewFlag(10, data["name"]), nil
}

func (flag *Flag) Expel() map[string]interface{} {
	var expelled map[string]interface{} = make(map[string]interface{})

	expelled["id"] = flag.id
	expelled["name"] = flag.name

	return expelled
}
