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

package database

import (
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) IsAvailable() bool {
	args := m.MethodCalled("IsAvailable")
	return mockUtils.ReturnBool(args, 0)
}

func (m *Mock) Find(_ interface{}, _ map[string]interface{}, _ string) response.IResponse {
	args := m.MethodCalled("Find")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) First(_ interface{}, _ map[string]interface{}, _ string) response.IResponse {
	args := m.MethodCalled("First")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) Raw(_ string, _ interface{}, _ ...interface{}) response.IResponse {
	args := m.MethodCalled("Raw")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) StartTransaction() IDatabaseWrite {
	args := m.MethodCalled("StartTransaction")
	return args.Get(0).(IDatabaseWrite)
}

func (m *Mock) RollbackTransaction() response.IResponse {
	args := m.MethodCalled("RollbackTransaction")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) CommitTransaction() response.IResponse {
	args := m.MethodCalled("CommitTransaction")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) Create(_ interface{}, _ string) response.IResponse {
	args := m.MethodCalled("Create")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) CreateOrUpdate(_ interface{}, _ map[string]interface{}, _ string) response.IResponse {
	args := m.MethodCalled("CreateOrUpdate")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) Update(_ interface{}, _ map[string]interface{}, _ string) response.IResponse {
	args := m.MethodCalled("Update")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) Delete(_ map[string]interface{}, _ string) response.IResponse {
	args := m.MethodCalled("Delete")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) FindPreload(_ interface{}, _ map[string]interface{}, _ map[string][]interface{}, _ string) response.IResponse {
	args := m.MethodCalled("FindPreload")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) FindPreloadWitLimitAndPage(_ interface{}, _ map[string]interface{},
	_ map[string][]interface{}, _ string, _, _ int) response.IResponse {
	args := m.MethodCalled("FindPreloadWitLimitAndPage")
	return args.Get(0).(response.IResponse)
}
