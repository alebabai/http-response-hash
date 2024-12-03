package domain

import (
	"context"
	"crypto/md5"
	"errors"
	"hash"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/alebabai/http-response-hash/pkg/application/domain/hasher"
)

type testHTTPClient struct {
	body string
	err  error
}

func (c *testHTTPClient) Do(_ *http.Request) (*http.Response, error) {
	if c.err != nil {
		return nil, c.err
	}

	return &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(
			strings.NewReader(c.body),
		),
	}, nil
}

func TestHasher_Process(t *testing.T) {
	const (
		testRequestURL   = "test.com"
		testResponseData = "test"
	)

	var (
		testHash = md5.New()
	)

	type fields struct {
		client HTTPClient
		hash   hash.Hash
	}
	type args struct {
		in hasher.HashURLContentInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hasher.HashURLContentOutput
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				client: &testHTTPClient{
					body: testResponseData,
				},
				hash: testHash,
			},
			args: args{
				in: hasher.HashURLContentInput{
					URL: testRequestURL,
				},
			},
			want: &hasher.HashURLContentOutput{
				URL:  testRequestURL,
				Sum:  testHash.Sum([]byte(testResponseData)),
				Size: testHash.Size(),
			},
		},
		{
			name: "err  client error",
			fields: fields{
				client: &testHTTPClient{
					err: errors.New("some client error"),
				},
				hash: testHash,
			},
			args: args{
				in: hasher.HashURLContentInput{
					URL: testRequestURL,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HasherService{
				client: tt.fields.client,
				hash:   tt.fields.hash,
			}
			got, err := h.HashURLContent(context.Background(), tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
