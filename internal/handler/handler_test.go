package handler

import (
	"net/http"
	"testing"

	"github.com/go_workshop_1/internal/api"
)

func TestHandler_Hello(t *testing.T) {
	type fields struct {
		weatherClient api.Client
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				weatherClient: tt.fields.weatherClient,
			}
			h.Hello(tt.args.w, tt.args.r)
		})
	}
}
