package v1

import (
	"github.com/golang/mock/gomock"
	"github.com/wangzhen94/iam/internal/apiserver/store"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := store.NewMockFactory(ctrl)
	type args struct {
		store store.Factory
	}
	tests := []struct {
		name string
		args args
		want *service
	}{
		{
			name: "default",
			args: args{
				store: mockFactory,
			},
			want: &service{store: mockFactory},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Policies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := store.NewMockFactory(ctrl)

	type fields struct {
		store store.Factory
	}
	tests := []struct {
		name   string
		fields fields
		want   PolicySrv
	}{
		{
			name: "default",
			fields: fields{
				store: mockFactory,
			},
			want: &policyService{
				store: mockFactory,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if got := s.Policies(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Policies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Secrets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := store.NewMockFactory(ctrl)
	s := &service{
		store: mockFactory,
	}

	tests := []struct {
		name string
		want SecretSrv
	}{
		{
			name: "default",
			want: newSecrets(s),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Secrets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Secrets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Users(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := store.NewMockFactory(ctrl)
	s := &service{
		store: mockFactory,
	}

	tests := []struct {
		name string
		want UserSrv
	}{
		{
			name: "default",
			want: newUsers(s),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Users(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Users() = %v, want %v", got, tt.want)
			}
		})
	}
}
