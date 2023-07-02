package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
    assert := assert.New(t)

    res := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/", nil)

    mux := NewHttpHandler()
    mux.ServeHTTP(res, req)

    assert.Equal(http.StatusOK, res.Code)
    data, _ := ioutil.ReadAll(res.Body)
    assert.Equal("Hello World!", string(data))
}

func TestIndexPathHandler_WithoutName(t *testing.T) {
    assert := assert.New(t)

    res := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/bar", nil)

    mux := NewHttpHandler()
    mux.ServeHTTP(res, req)

    assert.Equal(http.StatusOK, res.Code)
    data, _ := ioutil.ReadAll(res.Body)
    assert.Equal("Hello World!", string(data))
}

func TestIndexPathHandler_WithName(t *testing.T) {
    assert := assert.New(t)

    res := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/bar?name=yamkim", nil)

    mux := NewHttpHandler()
    mux.ServeHTTP(res, req)

    assert.Equal(http.StatusOK, res.Code)
    data, _ := ioutil.ReadAll(res.Body)
    assert.Equal("Hello yamkim!", string(data))
}