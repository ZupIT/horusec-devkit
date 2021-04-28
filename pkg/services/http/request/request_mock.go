package request

import (
	"crypto/tls"
	"net/http"

	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) NewHTTPRequest(_, _ string, _ interface{}, _ map[string]string) (*http.Request, error) {
	args := m.MethodCalled("NewHTTPRequest")
	return args.Get(0).(*http.Request), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) DoRequest(_ *http.Request, _ *tls.Config) (*http.Response, error) {
	args := m.MethodCalled("DoRequest")
	return args.Get(0).(*http.Response), mockUtils.ReturnNilOrError(args, 1)
}
