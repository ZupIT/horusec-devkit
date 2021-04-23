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
