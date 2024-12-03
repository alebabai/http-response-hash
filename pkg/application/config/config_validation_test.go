package config

import (
	"net/url"
	"testing"
)

func makeURL(rawURL string) url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		return url.URL{}
	}

	return *u
}

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		URLs     []url.URL
		Parallel uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Parallel: 10,
				URLs: []url.URL{
					makeURL("http://test1.com"),
					makeURL("http://test2.com"),
					makeURL("http://test3.com"),
				},
			},
		},
		{
			name: "err  empty urls",
			fields: fields{
				URLs: []url.URL{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Parallel: tt.fields.Parallel,
				URLs:     tt.fields.URLs,
			}
			if err := cfg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
