package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerIntegration_findTaskByIDHandler_Success(t *testing.T) {
	s := httptest.NewServer(server.NewServer().Handler)
	defer s.Close()

	resp, err := http.Get(s.URL + "/tasks/1")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var task server.Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	assert.NoError(t, err)
	assert.Equal(t, task, server.Tasks[1])
}

func TestServerIntegration_findTaskByIDHandler_NotFound(t *testing.T) {
	s := httptest.NewServer(server.NewServer().Handler)
	defer s.Close()

	resp, err := http.Get(s.URL + "/tasks/9")
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestServerIntegration_findTaskByIDHandler_BadRequest(t *testing.T) {
	s := httptest.NewServer(server.NewServer().Handler)
	defer s.Close()

	resp, err := http.Get(s.URL + "/tasks/invalid_id")
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestServerIntegration_findTaskByIDHandler_MethodNotAllowed(t *testing.T) {
	s := httptest.NewServer(server.NewServer().Handler)
	defer s.Close()

	resp, err := http.Post(s.URL+"/tasks/1", "application/json", nil)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}
