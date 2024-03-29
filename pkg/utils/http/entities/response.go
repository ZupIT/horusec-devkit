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

package entities

import "encoding/json"

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Content interface{} `json:"content,omitempty"`
}

func (r *Response) ToBytes() []byte {
	data, _ := json.Marshal(r)

	return data
}

func (r *Response) ToString() string {
	return string(r.ToBytes())
}

func (r *Response) SetResponseData(code int, status string, content interface{}) {
	r.Code = code
	r.Status = status
	r.Content = content
}

func (r *Response) GetStatusCode() int {
	return r.Code
}

func (r *Response) ContentToBytes() []byte {
	data, _ := json.Marshal(r.Content)

	return data
}
