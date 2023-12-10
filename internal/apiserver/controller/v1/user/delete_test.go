package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	v1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserSrv := v1.NewMockUserSrv(ctrl)
	mockService := v1.NewMockService(ctrl)
	mockService.EXPECT().Users().Return(mockUserSrv)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("DELETE", "/v1/users", nil)
	c.Params = []gin.Param{{Key: "name", Value: "admin"}}

	mockUserSrv.EXPECT().Delete(gomock.Any(), gomock.Eq("admin"), gomock.Any()).Return(nil)

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
			u.Delete(tt.args.c)
		})
	}
}
