package user

import (
	srvv1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	"github.com/wangzhen94/iam/internal/apiserver/store"
)

type UserController struct {
	srv srvv1.Service
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{srv: srvv1.NewService(store)}
}
