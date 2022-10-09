package parking_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"parking-lot/cmd/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParking(t *testing.T) {

	type args struct {
		method    string
		command   string
		value     string
		attribute string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Create parking request",
			args: args{
				method:  http.MethodPost,
				command: "/create_parking_lot",
				value:   "/6",
			},
			want:    "Created a parking lot with 6 slot(s)",
			wantErr: nil,
		},
	}
	gw := service.GatewayServer{}
	router := gw.SetupServer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080%s%s%s", tt.args.command, tt.args.value, tt.args.attribute)
			request := httptest.NewRequest(tt.args.method, url, nil)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			response := recorder.Result()
			body, err := io.ReadAll(response.Body)

			assert.Equal(t, tt.wantErr, err)
			assert.Contains(t, string(body), tt.want)
		})
	}
}
