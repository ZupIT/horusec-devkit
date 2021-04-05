package parser

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

func ParseBodyToEntity(body io.ReadCloser, entityPointer interface{}) error {
	err := json.NewDecoder(body).Decode(entityPointer)

	_ = body.Close()

	return err
}

func ParseEntityToIOReadCloser(entity interface{}) (io.ReadCloser, error) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(strings.NewReader(string(bytes))), nil
}
