package policy

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/wangzhen94/iam/internal/pkg/middleware"
	"github.com/wangzhen94/iam/pkg/log"
)

func (p *PolicyController) Delete(c *gin.Context) {
	log.L(c).Info("delete policy function called.")

	if err := p.srv.Policies().Delete(c, c.GetString(middleware.UsernameKey), c.Param("name"),
		metav1.DeleteOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
