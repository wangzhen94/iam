package user

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/wangzhen94/iam/pkg/log"
)

func (u *UserController) Delete(c *gin.Context) {
	name := c.Params.ByName("name")
	log.L(c).Infof("delete user %s", name)

	if err := u.srv.Users().Delete(c, name, metav1.DeleteOptions{Unscoped: true}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
