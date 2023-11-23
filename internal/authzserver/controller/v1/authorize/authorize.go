package authorize

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/ory/ladon"
	"github.com/wangzhen94/iam/internal/authzserver/authorization"
	"github.com/wangzhen94/iam/internal/authzserver/authorization/authorizer"
	"github.com/wangzhen94/iam/internal/pkg/code"
)

type AuthzController struct {
	store authorizer.PolicyGetter
}

func NewAuthzController(store authorizer.PolicyGetter) AuthzController {
	return AuthzController{store}
}

func (a *AuthzController) Authorize(c *gin.Context) {
	var r ladon.Request
	if err := c.ShouldBind(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	auth := authorization.NewAuthorizer(authorizer.NewAuthorization(a.store))
	if r.Context == nil {
		r.Context = ladon.Context{}
	}

	r.Context["username"] = c.GetString("username")
	rsp := auth.Authorize(&r)

	core.WriteResponse(c, nil, rsp)
}
