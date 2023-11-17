package policy

import (
	srvv1 "github.com/wangzhen94/iam/internal/apiserver/service/v1"
	"github.com/wangzhen94/iam/internal/apiserver/store"
)

type PolicyController struct {
	srv srvv1.Service
}

func NewPolicyController(store store.Factory) *PolicyController {
	return &PolicyController{srv: srvv1.NewService(store)}
}
