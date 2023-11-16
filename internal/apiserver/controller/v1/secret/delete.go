package secret

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/wangzhen94/iam/internal/pkg/middleware"
	"github.com/wangzhen94/iam/pkg/log"
)

func (s *SecretController) Delete(c *gin.Context) {
	log.L(c).Info("delete secret called.")
	opts := metav1.DeleteOptions{Unscoped: true}
	err := s.srv.Secrets().Delete(c, c.GetString(middleware.UsernameKey), c.Param("name"), opts)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
