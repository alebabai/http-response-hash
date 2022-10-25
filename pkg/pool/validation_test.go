package pool

import (
	"testing"
)

func TestPool_Validate(t *testing.T) {
	type fields struct {
		action   Action[any, any]
		consumer Consumer[any]
		size     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				action: func(in any) any {
					return 1
				},
				consumer: func(in any) {},
				size:     1,
			},
		},
		{
			name: "err  no action",
			fields: fields{
				action:   nil,
				consumer: func(in any) {},
				size:     1,
			},
			wantErr: true,
		},
		{
			name: "err  no consumer",
			fields: fields{
				action: func(in any) any {
					return 1
				},
				consumer: nil,
				size:     1,
			},
			wantErr: true,
		},
		{
			name: "err  invalid size",
			fields: fields{
				action: func(in any) any {
					return 1
				},
				consumer: func(in any) {},
				size:     0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Pool[any, any]{
				action:   tt.fields.action,
				consumer: tt.fields.consumer,
				size:     tt.fields.size,
			}
			if err := h.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Pool.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
