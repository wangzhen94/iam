package user

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	v1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := v1.NewMockService(ctrl)
	mockUserService := v1.NewMockUserSrv(ctrl)
	mockUserService.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockService.EXPECT().Users().Return(mockUserService)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	body := bytes.NewBufferString(
		`{"metadata":{"name":"admin"},"nickname":"admin","email":"aaa@qq.com","password":"Admin@2020","phone":"1812884xxx"}`,
	)
	c.Request, _ = http.NewRequest("POST", "/v1/users", body)
	c.Request.Header.Set("Content-Type", "application/json")

	type fields struct {
		srv v1.Service
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
			u.Create(tt.args.c)
		})
	}
}
