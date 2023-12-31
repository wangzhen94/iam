package config

import (
	"github.com/wangzhen94/iam/internal/pump/options"
	"reflect"
	"testing"
)

func TestCreateConfigFromOptions(t *testing.T) {
	opts := options.NewOptions()
	type args struct {
		opts *options.Options
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				opts: opts,
			},
			want: &Config{
				opts,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateConfigFromOptions(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConfigFromOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConfigFromOptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
