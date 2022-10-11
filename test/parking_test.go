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
	statusResult := `Slot No. Registration No Colour
1 B-1234-RFS Black
2 B-1999-RFD Green
3 B-1000-RFS Black
5 B-1701-RFL Blue
6 B-1141-RFS Black`

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
		{
			name: "Allocate parking 1",
			args: args{
				method:    http.MethodPost,
				command:   "/park",
				value:     "/B-1234-RFS",
				attribute: "/Black",
			},
			want:    "Allocated slot number: 1",
			wantErr: nil,
		},
		{
			name: "Allocate parking 2",
			args: args{
				method:    http.MethodPost,
				command:   "/park",
				value:     "/B-1999-RFD",
				attribute: "/Green",
			},
			want:    "Allocated slot number: 2",
			wantErr: nil,
		},
		{
			name: "Allocate parking 3",
			args: args{
				method:    http.MethodPost,
				command:   "/park",
				value:     "/B-1000-RFS",
				attribute: "/Black",
			},
			want:    "Allocated slot number: 3",
			wantErr: nil,
		},
		{
			name: "Allocate parking 4",
			args: args{
				method:    http.MethodPost,
				command:   "/park",
				value:     "/B-1777-RFU",
				attribute: "/BlueSky",
			},
			want:    "Allocated slot number: 4",
			wantErr: nil,
		},
		{
			name: "Allocate parking 5",
			args: args{
				method:    http.MethodPost,
				command:   "/park",
				value:     "/B-1701-RFL",
				attribute: "/Blue",
			},
			want:    "Allocated slot number: 5",
			wantErr: nil,
		},
		{
			name: "Allocate parking 6",
			args: args{
				method:    http.MethodPost,
				command:   "/park",
				value:     "/B-1141-RFS",
				attribute: "/Black",
			},
			want:    "Allocated slot number: 6",
			wantErr: nil,
		},
		{
			name: "Leave park no 4",
			args: args{
				method:  http.MethodPost,
				command: "/leave",
				value:   "/4",
			},
			want:    "Slot number 4 is free",
			wantErr: nil,
		},
		{
			name: "Status check",
			args: args{
				method:  http.MethodGet,
				command: "/status",
			},
			want:    statusResult,
			wantErr: nil,
		},
		{
			name: "Allocate parking 4",
			args: args{
				method:    http.MethodPost,
				command:   "/park",
				value:     "/B-1333-RFS",
				attribute: "/Black",
			},
			want:    "Allocated slot number: 4",
			wantErr: nil,
		},
		{
			name: "Allocate parking full",
			args: args{
				method:    http.MethodPost,
				command:   "/park",
				value:     "/B-1989-RFU",
				attribute: "/White",
			},
			want:    "Sorry, parking lot is full",
			wantErr: nil,
		},
		{
			name: "get cars registration numbers",
			args: args{
				method:    http.MethodGet,
				command:   "/cars_registration_numbers",
				value:     "/colour",
				attribute: "/Black",
			},
			want:    "B-1234-RFS, B-1000-RFS, B-1333-RFS, B-1141-RFS",
			wantErr: nil,
		},
		{
			name: "get cars slot numbers",
			args: args{
				method:    http.MethodGet,
				command:   "/cars_slot",
				value:     "/colour",
				attribute: "/Black",
			},
			want:    "1, 3, 4, 6",
			wantErr: nil,
		},
		{
			name: "get slot number by exist plate number",
			args: args{
				method:    http.MethodGet,
				command:   "/slot_number",
				value:     "/car_registration_number",
				attribute: "/B-1701-RFL",
			},
			want:    "5",
			wantErr: nil,
		},
		{
			name: "get slot number by not exist plate number",
			args: args{
				method:    http.MethodGet,
				command:   "/slot_number",
				value:     "/car_registration_number",
				attribute: "/RI-1",
			},
			want:    "Not found",
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

			responseMsg := string(body)
			fmt.Println(responseMsg)
			assert.Equal(t, tt.wantErr, err)
			assert.Contains(t, responseMsg, tt.want)
		})
	}
}
