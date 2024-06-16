package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/symonk/toodoo/internal/server"
)

func TestHealtchCheckStatus(t *testing.T) {
	r := server.NewRouter()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/healthcheck", nil)
	r.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Uptime")

}
