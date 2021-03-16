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

func (m *Mock) Raw(_ string, _ interface{}) response.IResponse {
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
