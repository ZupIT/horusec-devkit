package parser

import (
	"encoding/json"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser/enums"
	"io"
	"io/ioutil"
	"strings"
)

func ParseBodyToEntity(body io.ReadCloser, entityPointer interface{}) error {
	err := json.NewDecoder(body).Decode(entityPointer)
	_ = body.Close()

	if err != nil {
		if strings.EqualFold(err.Error(), "eof") {
			return enums.ErrorBodyEmpty
		}
		if strings.Contains(err.Error(), "invalid character") {
			return enums.ErrorBodyInvalid
		}
	}

	return nil
}

func ParseEntityToIOReadCloser(entity interface{}) (io.ReadCloser, error) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(strings.NewReader(string(bytes))), nil
}
