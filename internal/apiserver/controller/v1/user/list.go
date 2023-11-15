package user

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
)
import metav1 "github.com/marmotedu/component-base/pkg/meta/v1"

func (u *UserController) List(c *gin.Context) {
	log.L(c).Info("list user function called.")

	var r metav1.ListOptions

	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	users, err := u.srv.Users().List(c, metav1.ListOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, users)
}
