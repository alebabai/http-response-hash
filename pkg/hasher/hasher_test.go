package hasher

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestResult_String(t *testing.T) {
	hash := md5.New()
	testSum := hash.Sum([]byte("test"))
	type fields struct {
		URL  string
		Sum  []byte
		Size int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				URL:  "http://test.com",
				Sum:  testSum,
				Size: hash.Size(),
			},
			want: fmt.Sprintf("%s %x", "http://test.com", testSum[:hash.Size()]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Result{
				URL:  tt.fields.URL,
				Sum:  tt.fields.Sum,
				Size: tt.fields.Size,
			}
			if got := res.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testHTTPClient struct {
	body string
	err  error
}

type testHash struct {
	sum  []byte
	size int
}

func (h *testHash) Sum([]byte) []byte {
	return h.sum
}

func (h *testHash) Size() int {
	return h.size
}

func (c *testHTTPClient) Get(string) (*http.Response, error) {
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

func TestService_Process(t *testing.T) {
	const (
		testRequestURL   = "test.com"
		testResponseData = "test"
	)

	type fields struct {
		client httpClient
		hash   hashSum
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				client: &testHTTPClient{
					body: testResponseData,
				},
				hash: &testHash{
					sum:  md5.New().Sum([]byte(testResponseData)),
					size: md5.Size,
				},
			},
			args: args{
				url: testRequestURL,
			},
			want: &Result{
				URL:  testRequestURL,
				Sum:  md5.New().Sum([]byte(testResponseData)),
				Size: md5.Size,
			},
		},
		{
			name: "err  client error",
			fields: fields{
				client: &testHTTPClient{
					err: errors.New("some client error"),
				},
				hash: &testHash{
					sum:  md5.New().Sum([]byte(testResponseData)),
					size: md5.Size,
				},
			},
			args: args{
				url: testRequestURL,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Hasher{
				client: tt.fields.client,
				hash:   tt.fields.hash,
			}
			got, err := svc.Process(tt.args.url)
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
