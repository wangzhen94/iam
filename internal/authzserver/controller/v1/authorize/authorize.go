package authorize

import (
	"github.com/gin-gonic/gin"
	"github.com/wangzhen94/iam/internal/authzserver/authorization/authorizer"
)

type AuthzController struct {
	store authorizer.PolicyGetter
}

func NewAuthzController(store authorizer.PolicyGetter) AuthzController {
	return AuthzController{store}
}

func (a *AuthzController) Authorize(c *gin.Context) {

}
