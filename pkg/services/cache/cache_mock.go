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

package cache

import (
	"time"

	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Get(_ string) interface{} {
	args := m.MethodCalled("Get")
	return args.Get(0).(interface{})
}

func (m *Mock) GetAndParse(_ string, _ interface{}) error {
	args := m.MethodCalled("GetAndParse")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) GetString(_ string) (result string, err error) {
	args := m.MethodCalled("GetString")
	return args.Get(0).(string), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Delete(_ string) {
	_ = m.MethodCalled("Delete")
}

func (m *Mock) Set(_ string, _ interface{}, _ time.Duration) {
	_ = m.MethodCalled("Set")
}
