package hasher

import (
	"testing"
)

func TestHashURLContentInput_Validate(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				URL: "http://test1.com",
			},
		},
		{
			name: "err  empty url",
			fields: fields{
				URL: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dto := HashURLContentInput{
				URL: tt.fields.URL,
			}
			if err := dto.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("HashURLContentInput.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
