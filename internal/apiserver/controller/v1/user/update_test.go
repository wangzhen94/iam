package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/marmotedu/api/apiserver/v1"
	srvv1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	"testing"
)

func TestUserController_Update(t *testing.T) {
	type fields struct {
		srv srvv1.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				srv: tt.fields.srv,
			}
			u.Update(tt.args.c)
		})
	}
}

func Test_checkOldPassword(t *testing.T) {
	type args struct {
		user        *v1.User
		oldPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkOldPassword(tt.args.user, tt.args.oldPassword); (err != nil) != tt.wantErr {
				t.Errorf("checkOldPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
