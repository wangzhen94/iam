package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	srvv1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserController_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := srvv1.NewMockService(ctrl)
	mockUserSrv := srvv1.NewMockUserSrv(ctrl)
	mockService.EXPECT().Users().Return(mockUserSrv)
	user := v1.User{
		ObjectMeta:  metav1.ObjectMeta{},
		Status:      0,
		Nickname:    "wangx",
		Password:    "",
		Email:       "123456789@163.com",
		Phone:       "1769687xxxx",
		IsAdmin:     0,
		TotalPolicy: 3,
		LoginedAt:   time.Time{},
	}
	mockUserSrv.EXPECT().Get(gomock.Any(), gomock.Eq("wangx"), gomock.Any()).Return(&user, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/v1/users/wangx", nil)
	c.Params = gin.Params{{Key: "name", Value: "wangx"}}
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
			u.Get(tt.args.c)
		})
	}
}
