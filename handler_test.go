package rfc9457_test

import (
	"errors"
	"fmt"
	"github.com/alecthomas/assert/v2"
	"github.com/frederik-jatzkowski/go-rfc9457"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	ErrTest1 = errors.New("test 1")
)

func TestHandler_ServeHTTP(t *testing.T) {
	registry := rfc9457.NewRegistry()
	err := registry.RegisterError(ErrTest1, http.StatusBadRequest)
	require.NoError(t, err)

	handler, err := rfc9457.NewHandler(registry)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", ErrTest1.Error()), nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	response := recorder.Result()

	data, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Contains(t, string(data), fmt.Sprintf("%s</h1>", ErrTest1.Error()))
}
