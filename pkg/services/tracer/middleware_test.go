package tracer

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer/enums"
)

func TestMiddleware(t *testing.T) {

	handler200 := func(w http.ResponseWriter, r *http.Request) {
		AddOperationName(r.Context(), "handler200")
		w.WriteHeader(http.StatusOK)
	}
	handler400 := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}
	handler500 := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	handlerPanic := func(w http.ResponseWriter, r *http.Request) {
		panic(errors.New("error"))
	}
	router := chi.NewRouter()
	req := httptest.NewRequest(http.MethodGet, "http://www.your-domain.com", nil)
	req = req.WithContext(context.WithValue(req.Context(), "some-key", "123ABC"))

	err := os.Setenv(enums.HorusecJaegerName, "test")
	assert.NoError(t, err)
	err = os.Setenv(enums.HorusecJaegerLogInfo, "true")
	assert.NoError(t, err)
	err = os.Setenv(enums.HorusecJaegerLogError, "true")
	assert.NoError(t, err)

	j, err := NewJaeger()
	assert.NoError(t, err)
	closer, err := j.Config(true)
	assert.NoError(t, err)
	defer func() {
		err := closer.Close()
		assert.NoError(t, err)
	}()
	router.Use(Tracer(opentracing.GlobalTracer()))
	router.Get("/swagger/200", handler200)
	router.Get("/400", handler400)
	router.Get("/500", handler500)
	router.Get("/panic", handlerPanic)

	ts := httptest.NewServer(router)
	defer ts.Close()
	t.Run("should get a status 200 ", func(t *testing.T) {
		resp, _ := testRequest(t, ts, http.MethodGet, "/swagger/200", nil)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run("should get a status 400 ", func(t *testing.T) {
		resp, _ := testRequest(t, ts, http.MethodGet, "/400", nil)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("should get a status 500 ", func(t *testing.T) {
		resp, _ := testRequest(t, ts, http.MethodGet, "/500", nil)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
	//t.Run("should panic", func(t *testing.T) {
	//	assert.Panics(t, func() {
	//		testRequest(t, ts, http.MethodGet, "/panic", nil)
	//	} )
	//})

}
func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}
