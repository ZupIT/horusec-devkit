package parser

import (
	"encoding/json"
	"io"
)

func ParseBodyToEntity(body io.ReadCloser, entityPointer interface{}) error {
	err := json.NewDecoder(body).Decode(entityPointer)

	_ = body.Close()

	return err
}
