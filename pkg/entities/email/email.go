// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package email

import (
	"encoding/json"

	"github.com/ZupIT/horusec-devkit/pkg/enums/email"
)

type Message struct {
	To           string         `json:"to"`
	Subject      string         `json:"subject"`
	TemplateName email.Template `json:"templateName"`
	Data         interface{}    `json:"data"`
}

func (m *Message) ToBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
