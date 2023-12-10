package user

import (
	"github.com/golang/mock/gomock"
	srvv1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	"github.com/wangzhen94/iam/internal/apiserver/store"
	"reflect"
	"testing"
)

func TestNewUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := store.NewMockFactory(ctrl)

	type args struct {
		store store.Factory
	}
	tests := []struct {
		name string
		args args
		want *UserController
	}{
		{
			name: "default",
			args: args{store: mockFactory},
			want: &UserController{srv: srvv1.NewService(mockFactory)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserController() = %v, want %v", got, tt.want)
			}
		})
	}
}
