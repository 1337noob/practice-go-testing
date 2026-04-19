package weather_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"weather"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetWeatherWithTestify(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockClient := new(MockHTTPClient)
		wantResponse := &http.Response{
			StatusCode: http.StatusOK,
			Status:     http.StatusText(http.StatusOK),
			Body:       io.NopCloser(bytes.NewBufferString(`{"main":{"temp":11.86}}`)),
		}
		mockClient.On("Get", mock.Anything).Return(wantResponse, nil)

		temp, err := weather.GetWeather(mockClient, "Moscow", "apiKey")
		assert.NoError(t, err)
		assert.Equal(t, 11.86, temp)
		mockClient.AssertExpectations(t)
	})

	t.Run("api error", func(t *testing.T) {
		mockClient := new(MockHTTPClient)
		errorResponse := &http.Response{
			StatusCode: http.StatusNotFound,
			Status:     http.StatusText(http.StatusNotFound),
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}
		mockClient.On("Get", mock.Anything).Return(errorResponse, nil)

		_, err := weather.GetWeather(mockClient, "UnknownCity", "apiKey")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "API returned status code: 404")
		mockClient.AssertExpectations(t)
	})
}
