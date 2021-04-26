package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go_workshop_1/internal/api"
	"github.com/go_workshop_1/internal/api/mocks"
)

func TestHandler_Hello(t *testing.T) {
	tests := []struct {
		name     string
		weather  *api.WeatherResponse // init field for json
		err      error
		codeWant int
		bodyWant string
	}{
		{
			name:     "simple test",
			weather:  &api.WeatherResponse{Weather: "test weather"},
			err:      nil,
			codeWant: 200,
			bodyWant: "test weather",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := &mocks.Client{}
			// when call GetWeather what must return
			apiMock.On("GetWeather").Return(tt.weather, tt.err)

			h := NewHandler(apiMock)

			req, _ := http.NewRequest("GET", "/hello", nil)
			rr := httptest.NewRecorder()

			h.Hello(rr, req)

			// same as above line
			// handler := http.HandlerFunc(h.Hello)
			// handler.ServeHTTP(rr, req)

			gotRaw, _ := ioutil.ReadAll(rr.Body)
			got := string(gotRaw)

			if got != tt.bodyWant {
				t.Errorf("wrong response body %s want %s", got, tt.bodyWant)
			}

			if status := rr.Result().StatusCode; status != tt.codeWant {
				t.Errorf("wrong response body %d want %d", status, tt.codeWant)
			}
		})
	}
}
