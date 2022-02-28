package hasher_test

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/alebabai/http-response-hash/pkg/hasher"
)

func TestService_Process_Integration(t *testing.T) {
	const (
		testResponseData = "test"
	)

	tests := []struct {
		name    string
		svc     *hasher.Service
		srv     *httptest.Server
		want    *hasher.Output
		prepare func()
		wantErr bool
	}{
		{
			name: "ok",
			svc: hasher.NewService(
				http.DefaultClient,
				md5.New(),
			),
			srv: httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, _ = fmt.Fprintf(w, testResponseData)
				}),
			),
			want: &hasher.Output{
				Sum:  md5.New().Sum([]byte(testResponseData)),
				Size: md5.Size,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.srv.Close()

			// get dynamic url from test http server
			tt.want.URL = tt.srv.URL

			got, err := tt.svc.Process(tt.srv.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() got = %v, want %v", got, tt.want)
			}
		})
	}
}
