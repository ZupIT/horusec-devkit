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

package entities

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/request/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type HTTPResponse struct {
	*http.Response
}

func (h *HTTPResponse) GetBodyBytes() ([]byte, error) {
	if h.Body == nil {
		return []byte{}, nil
	}

	return ioutil.ReadAll(h.Body)
}

func (h *HTTPResponse) GetStatusCodeString() string {
	return http.StatusText(h.StatusCode)
}

func (h *HTTPResponse) GetContentType() string {
	return h.Header.Get(enums.ContentType)
}

func (h *HTTPResponse) GetStatusCode() int {
	return h.StatusCode
}

func (h *HTTPResponse) CloseBody() {
	if h.Body != nil {
		logger.LogError(enums.MessageFailedCloseRequestBody, h.Body.Close())
	}
}

func (h *HTTPResponse) ErrorByStatusCode() error {
	body, _ := h.GetBodyBytes()
	switch {
	case h.StatusCode >= http.StatusInternalServerError:
		logger.LogError(enums.MessageHTTPResponseErrorStatusCode, errors.New(string(body)), h.mapResponse())
		return enums.ErrorRequestServerError
	case h.StatusCode >= http.StatusBadRequest && h.StatusCode < http.StatusInternalServerError:
		logger.LogError(enums.MessageHTTPResponseErrorStatusCode, errors.New(string(body)), h.mapResponse())
		return enums.ErrorRequestClientError
	default:
		return nil
	}
}

func (h *HTTPResponse) mapResponse() map[string]interface{} {
	return map[string]interface{}{
		"statusCode": h.GetStatusCode(),
		"status":     h.GetStatusCodeString(),
	}
}
