// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/google/uuid"

	brokerPacket "github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
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

	if strings.Contains(err.Error(), enums.InvalidJSONInput) {
		return enums.ErrorBodyInvalid
	}

	return err
}

func ParseEntityToIOReadCloser(entity interface{}) (io.ReadCloser, error) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(strings.NewReader(string(bytes))), nil
}

func ParseStringToUUID(id string) uuid.UUID {
	parsedID, _ := uuid.Parse(id)
	return parsedID
}

func ParsePacketToEntity(packet brokerPacket.IPacket, entityPointer interface{}) error {
	if err := json.Unmarshal(packet.GetBody(), entityPointer); err != nil {
		return checkParseBodyToEntityError(err)
	}

	return nil
}
