package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	srvv1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := srvv1.NewMockService(ctrl)
	mockUserSrv := srvv1.NewMockUserSrv(ctrl)
	mockUserSrv.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, nil)
	mockService.EXPECT().Users().Return(mockUserSrv)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/v1/users", nil)
	c.Params = []gin.Param{{Key: "name", Value: "colin"}}

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
		{
			name: "default",
			fields: fields{
				srv: mockService,
			},
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				srv: tt.fields.srv,
			}
			u.List(tt.args.c)
		})
	}
}
