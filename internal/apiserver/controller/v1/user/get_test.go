package user

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserController_Get(t *testing.T) {
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
			u.Get(tt.args.c)
		})
	}
}
