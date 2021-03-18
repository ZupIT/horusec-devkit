package http

import (
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	text string
}

func TestNewHTTPClientUtil(t *testing.T) {
	t.Run("should success create a new http util", func(t *testing.T) {
		assert.NotNil(t, NewHTTPRequestUtil(0))
	})
}

func TestDoRequest(t *testing.T) {
	t.Run("should success make a http request with no error and status code 200", func(t *testing.T) {
		requestUtil := NewHTTPRequestUtil(20)

		request, err := requestUtil.NewHTTPRequest(http.MethodGet, "https://httpbin.org/get",
			nil, map[string]string{"accept": "application/json"})
		assert.NoError(t, err)

		response, err := requestUtil.DoRequest(request, &tls.Config{})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("should return error while making request", func(t *testing.T) {
		requestUtil := NewHTTPRequestUtil(0)

		request, err := requestUtil.NewHTTPRequest(http.MethodGet, "test", nil, nil)
		assert.NoError(t, err)

		response, err := requestUtil.DoRequest(request, &tls.Config{})
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestNewHTTPRequest(t *testing.T) {
	t.Run("should success create a new http request with headers and body", func(t *testing.T) {
		requestUtil := NewHTTPRequestUtil(0)

		request, err := requestUtil.NewHTTPRequest(http.MethodGet, "http://localhost:9999", &test{text: "test"},
			map[string]string{"test": "test"})

		assert.NoError(t, err)
		assert.NotNil(t, request)
	})

	t.Run("should success create a new http request with body but without headers", func(t *testing.T) {
		requestUtil := NewHTTPRequestUtil(0)

		request, err := requestUtil.NewHTTPRequest(http.MethodGet, "http://localhost:9999", &test{text: "test"},
			nil)

		assert.NoError(t, err)
		assert.NotNil(t, request)
	})

	t.Run("should success create a new http request with headers but without body", func(t *testing.T) {
		requestUtil := NewHTTPRequestUtil(0)

		request, err := requestUtil.NewHTTPRequest(http.MethodGet, "http://localhost:9999", nil,
			map[string]string{"test": "test"})

		assert.NoError(t, err)
		assert.NotNil(t, request)
	})

	t.Run("should return error when invalid body", func(t *testing.T) {
		requestUtil := NewHTTPRequestUtil(0)

		request, err := requestUtil.NewHTTPRequest(http.MethodGet, "http://localhost:9999", make(chan string),
			nil)

		assert.Error(t, err)
		assert.Nil(t, request)
	})
}
