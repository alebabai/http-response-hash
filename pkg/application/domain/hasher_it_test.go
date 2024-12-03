package domain_test

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/alebabai/http-response-hash/pkg/application/domain"
	"github.com/alebabai/http-response-hash/pkg/application/domain/hasher"
)

func TestService_Process_IntegrationTest(t *testing.T) {
	ctx := context.Background()

	const testResponseData = "test"

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprint(w, testResponseData)
		}),
	)

	defer srv.Close()

	hash := md5.New()
	svc := domain.NewHasherService(hash)

	got, err := svc.HashURLContent(ctx, hasher.HashURLContentInput{URL: srv.URL})
	if err != nil {
		t.Errorf("h.Process() error = %v, want nil", err)
		return
	}

	want := &hasher.HashURLContentOutput{
		Sum:  hash.Sum([]byte(testResponseData)),
		Size: hash.Size(),
		URL:  srv.URL,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Process() got = %v, want %v", got, want)
		return
	}
}
