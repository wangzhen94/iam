package secret

import (
	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	v1 "github.com/marmotedu/api/apiserver/v1"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/component-base/pkg/util/idutil"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
)

const maxSecretCount = 10

func (s *SecretController) Create(c *gin.Context) {
	var req v1.Secret

	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := req.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	//username := c.GetString(middleware.UsernameKey)
	username := c.Query("username")

	log.Infof("user: %s create secret %s", username, req)
	secrets, err := s.srv.Secrets().List(c, username, metav1.ListOptions{
		Offset: pointer.ToInt64(0),
		Limit:  pointer.ToInt64(-1),
	})

	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if secrets.TotalCount >= maxSecretCount {
		core.WriteResponse(c, errors.WithCode(code.ErrReachMaxCount, "secret count: %d", secrets.TotalCount), nil)

		return
	}

	req.Username = username

	req.SecretID = idutil.NewSecretID()
	req.SecretKey = idutil.NewSecretKey()

	if err = s.srv.Secrets().Create(c, &req, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, req)
}
