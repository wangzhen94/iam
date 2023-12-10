package authorization

import (
	"github.com/golang/mock/gomock"
	"github.com/ory/ladon"
	"reflect"
	"testing"
)

func TestAuditLogger_LogGrantedAccessRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthz := NewMockAuthorizationInterface(ctrl)
	mockAuthz.EXPECT().LogGrantedAccessRequest(gomock.Any(), gomock.Any(), gomock.Any())

	type fields struct {
		client AuthorizationInterface
	}
	type args struct {
		r *ladon.Request
		p ladon.Policies
		d ladon.Policies
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "default",
			fields: fields{mockAuthz},
			args: args{
				r: &ladon.Request{},
				p: ladon.Policies{},
				d: ladon.Policies{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuditLogger{
				client: tt.fields.client,
			}
			a.LogGrantedAccessRequest(tt.args.r, tt.args.p, tt.args.d)
		})
	}
}

func TestAuditLogger_LogRejectedAccessRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthz := NewMockAuthorizationInterface(ctrl)
	mockAuthz.EXPECT().LogRejectedAccessRequest(gomock.Any(), gomock.Any(), gomock.Any())

	type fields struct {
		client AuthorizationInterface
	}
	type args struct {
		r *ladon.Request
		p ladon.Policies
		d ladon.Policies
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "default",
			fields: fields{mockAuthz},
			args: args{
				r: &ladon.Request{},
				p: ladon.Policies{},
				d: ladon.Policies{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuditLogger{
				client: tt.fields.client,
			}
			a.LogRejectedAccessRequest(tt.args.r, tt.args.p, tt.args.d)
		})
	}
}

func TestNewAuditLogger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthz := NewMockAuthorizationInterface(ctrl)

	type args struct {
		client AuthorizationInterface
	}
	tests := []struct {
		name string
		args args
		want *AuditLogger
	}{
		{
			name: "default",
			args: args{client: mockAuthz},
			want: &AuditLogger{
				client: mockAuthz,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuditLogger(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuditLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
