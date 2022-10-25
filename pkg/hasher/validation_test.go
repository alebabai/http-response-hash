package hasher

import (
	"crypto/md5"
	"net/http"
	"testing"
)

func TestHasher_Validate(t *testing.T) {
	type fields struct {
		client httpClient
		hash   hash
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				client: http.DefaultClient,
				hash:   md5.New(),
			},
		},
		{
			name: "err  no client",
			fields: fields{
				client: nil,
				hash:   md5.New(),
			},
			wantErr: true,
		},
		{
			name: "err  no hash",
			fields: fields{
				client: http.DefaultClient,
				hash:   nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hasher{
				client: tt.fields.client,
				hash:   tt.fields.hash,
			}
			if err := h.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Hasher.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
