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
