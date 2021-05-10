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

package request

import (
	"crypto/tls"
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/request/entities"

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

func (m *Mock) DoRequest(_ *http.Request, _ *tls.Config) (*entities.HTTPResponse, error) {
	args := m.MethodCalled("DoRequest")
	return args.Get(0).(*entities.HTTPResponse), mockUtils.ReturnNilOrError(args, 1)
}
