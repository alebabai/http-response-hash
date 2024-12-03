package hasher

import (
	"crypto/md5"
	"fmt"
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
			res := HashURLContentOutput{
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
