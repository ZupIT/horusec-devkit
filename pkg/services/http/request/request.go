package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/request/entities"
	"github.com/ZupIT/horusec-devkit/pkg/services/http/request/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type IRequest interface {
	NewHTTPRequest(method, url string, body interface{}, headers map[string]string) (*http.Request, error)
	DoRequest(request *http.Request, tlsConfig *tls.Config) (*entities.HTTPResponse, error)
}

type Request struct {
	timeout int
}

func NewHTTPRequestService(timeout int) IRequest {
	return &Request{
		timeout: timeout,
	}
}

func (r *Request) DoRequest(request *http.Request, tlsConfig *tls.Config) (*entities.HTTPResponse, error) {
	response, err := r.setClient(tlsConfig).Do(request)
	if err != nil {
		logger.LogError(enums.MessageFailedToMakeHTTPRequest, err)
		return r.newHTTPResponse(response), err
	}

	return r.newHTTPResponse(response), nil
}

func (r *Request) setClient(tlsConfig *tls.Config) *http.Client {
	return &http.Client{
		Timeout:   r.getTimeout(),
		Transport: r.getTransport(tlsConfig),
	}
}

func (r *Request) getTimeout() time.Duration {
	return time.Second * time.Duration(r.timeout)
}

func (r *Request) getTransport(tlsConfig *tls.Config) *http.Transport {
	return &http.Transport{
		TLSClientConfig: tlsConfig,
	}
}

func (r *Request) newHTTPResponse(response *http.Response) *entities.HTTPResponse {
	return &entities.HTTPResponse{Response: response}
}

func (r *Request) NewHTTPRequest(method, url string, body interface{},
	headers map[string]string) (*http.Request, error) {
	data, err := r.parseBodyToIOReader(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(strings.ToUpper(method), url, data)
	if err == nil && req != nil {
		req = r.setHTTPRequestHeaders(req, headers)
	}

	return req, err
}

func (r *Request) parseBodyToIOReader(body interface{}) (io.Reader, error) {
	if body == nil || body == "" {
		return nil, nil
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}

func (r *Request) setHTTPRequestHeaders(req *http.Request, headers map[string]string) *http.Request {
	for key, value := range headers {
		if key != "" && value != "" {
			req.Header.Add(key, value)
		}
	}

	return req
}
