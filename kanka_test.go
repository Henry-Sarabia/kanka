package kanka

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const testEndpoint endpoint = "test/"
const testToken string = "not_a_real_token"

//TODO: c.Request has a new parameter; tests should reflect this
func TestClient_request(t *testing.T) {
	type args struct {
		method string
		end    endpoint
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			"Happy path",
			NewClient(testToken, nil),
			args{method: "GET", end: testEndpoint},
			httptest.NewRequest("GET", kankaURL+string(testEndpoint), nil),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.request(tt.args.method, tt.args.end, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Method != tt.want.Method {
				t.Errorf("got: %v, want: %v", got.Method, tt.want.Method)
			}

			if got.URL.String() != tt.want.URL.String() {
				t.Errorf("got: %v, want: %v", got.URL.String(), tt.want.URL.String())
			}

			if got.Header.Get("Authorization") != "Bearer "+tt.c.token {
				t.Errorf("got: %v, want: %v", got.Header.Get("Authorization"), "Bearer "+tt.c.token)
			}

			if got.Header.Get("Accept") != "application/json" {
				t.Errorf("got: %v, want :%v", got.Header.Get("Accept"), "application/json")
			}
		})
	}
}
