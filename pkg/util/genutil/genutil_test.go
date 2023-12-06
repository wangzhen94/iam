package genutil

import (
	"testing"
)

func TestOutDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "cur dir",
			args: args{
				path: "./",
			},
			want:    "/Users/wangzhen/go/src/github.com/wangzhen94/iam/pkg/util/genutil/",
			wantErr: false,
		},

		{
			name: "no exist dir err",
			args: args{
				path: "./noexistdir",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "if is file err",
			args: args{
				path: "./genutil_test.go",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OutDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("OutDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OutDir() got = %v, want %v", got, tt.want)
			}
		})
	}

}
