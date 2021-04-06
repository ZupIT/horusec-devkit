package parser

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser/enums"
)

func ParseBodyToEntity(body io.ReadCloser, entityPointer interface{}) error {
	err := json.NewDecoder(body).Decode(entityPointer)
	_ = body.Close()

	if err != nil {
		return checkParseBodyToEntityError(err)
	}

	return nil
}

func checkParseBodyToEntityError(err error) error {
	if strings.EqualFold(err.Error(), enums.EOF) {
		return enums.ErrorBodyEmpty
	}

	if strings.Contains(err.Error(), enums.InvalidCharacter) {
		return enums.ErrorBodyInvalid
	}

	return err
}

func ParseEntityToIOReadCloser(entity interface{}) (io.ReadCloser, error) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(strings.NewReader(string(bytes))), nil
}

func ParseStringToUUID(id string) uuid.UUID {
	parsedID, _ := uuid.Parse(id)
	return parsedID
}
