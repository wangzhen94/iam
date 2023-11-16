package user

import (
	"reflect"
	"testing"

	"github.com/wangzhen94/iam/internal/apiserver/store"
)

func TestNewUserController(t *testing.T) {
	type args struct {
		store store.Factory
	}
	tests := []struct {
		name string
		args args
		want *UserController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserController() = %v, want %v", got, tt.want)
			}
		})
	}
}
