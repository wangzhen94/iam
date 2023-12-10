package user

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	v1 "github.com/marmotedu/api/apiserver/v1"
	srvv1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	_ "github.com/wangzhen94/iam/pkg/validator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_ChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wangx := &v1.User{
		Password: "$2a$10$KqZhl5WStpa2K.ddEyzyf.zXllEXP4gIG8xQUgMhU1ZvMUn/Ta5um",
	}

	mockService := srvv1.NewMockService(ctrl)
	mockUserSrv := srvv1.NewMockUserSrv(ctrl)
	mockService.EXPECT().Users().Return(mockUserSrv).Times(2)
	mockUserSrv.EXPECT().ChangePassword(gomock.Any(), gomock.Any()).Return(nil)
	mockUserSrv.EXPECT().Get(gomock.Any(), gomock.Eq("wangx"), gomock.Any()).Return(wangx, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	body := bytes.NewBufferString(`{"oldPassword":"Admin@2020","newPassword":"Wangx@2023"}`)
	c.Request, _ = http.NewRequest("PUT", "/v1/users/wangx/change_password", body)
	c.Params = gin.Params{{Key: "name", Value: "wangx"}}
	c.Request.Header.Set("Content-Type", "application/json")

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
			name:   "default",
			fields: fields{srv: mockService},
			args:   args{c: c},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				srv: tt.fields.srv,
			}
			u.ChangePassword(tt.args.c)
		})
	}
}
