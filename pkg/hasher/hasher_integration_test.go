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

func TestService_Process_IntegrationTest(t *testing.T) {
	const testResponseData = "test"

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprint(w, testResponseData)
		}),
	)
	defer srv.Close()

	h, err := hasher.New(
		http.DefaultClient,
		md5.New(),
	)
	if err != nil {
		t.Errorf("hasher.New() error = %v, want nil", err)
	}

	got, err := h.Process(srv.URL)
	if err != nil {
		t.Errorf("h.Process() error = %v, want nil", err)
	}

	want := &hasher.Result{
		Sum:  md5.New().Sum([]byte(testResponseData)),
		Size: md5.Size,
		URL:  srv.URL,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Process() got = %v, want %v", got, want)
	}

}
