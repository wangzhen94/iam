package gormutil

import (
	"github.com/AlekSi/pointer"
	"reflect"
	"testing"
)

func TestUnPointer(t *testing.T) {
	type args struct {
		offset *int64
		limit  *int64
	}
	tests := []struct {
		name string
		args args
		want *LimitAndOffset
	}{
		{
			name: "normal",
			args: args{
				offset: pointer.ToInt64(1),
				limit:  pointer.ToInt64(10),
			},
			want: &LimitAndOffset{
				Offset: 1,
				Limit:  10,
			},
		},
		{
			name: "nil",
			args: args{
				offset: nil,
				limit:  nil,
			},
			want: &LimitAndOffset{
				Offset: 0,
				Limit:  1000,
			},
		},
		{
			name: "nil1",
			args: args{
				offset: pointer.ToInt64(10),
				limit:  nil,
			},
			want: &LimitAndOffset{
				Offset: 10,
				Limit:  1000,
			},
		},
		{
			name: "nil2",
			args: args{
				offset: nil,
				limit:  pointer.ToInt64(10),
			},
			want: &LimitAndOffset{
				Offset: 0,
				Limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnPointer(tt.args.offset, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
