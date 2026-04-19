package http_test

import (
	"encoding/json"
	handler "http"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUserHandler(t *testing.T) {
	tests := []struct {
		name       string
		userID     string
		method     string
		wantStatus int
		wantUser   *handler.User
	}{
		{
			name:       "valid id 1",
			userID:     "1",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
			wantUser:   handler.Users[1],
		},
		{
			name:       "valid id 2",
			userID:     "2",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
			wantUser:   handler.Users[2],
		},
		{
			name:       "not found",
			userID:     "8",
			method:     http.MethodGet,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "invalid id",
			userID:     "id",
			method:     http.MethodGet,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "method not allowed",
			userID:     "1",
			method:     http.MethodPut,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.method, "/users?id="+test.userID, nil)
			w := httptest.NewRecorder()

			handler.FindUserHandler(w, req)

			assert.Equal(t, w.Code, test.wantStatus)

			if test.wantUser != nil {
				var user handler.User
				err := json.NewDecoder(w.Body).Decode(&user)
				assert.NoError(t, err)
				assert.Equal(t, test.wantUser, &user)
			}
		})
	}
}
