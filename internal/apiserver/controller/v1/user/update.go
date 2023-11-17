package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/marmotedu/api/apiserver/v1"
	"github.com/marmotedu/component-base/pkg/auth"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
)

func (u *UserController) Update(c *gin.Context) {
	var r *v1.Policy
	oldPassword := c.Query("password")
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	log.L(c).Infof("update user %s, function called.", r.Name)

	user, err := u.srv.Users().Get(c, r.Name, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	err = checkOldPassword(user, oldPassword)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrPasswordIncorrect, err.Error()), nil)

		return
	}
	user.Nickname = r.Nickname
	user.Email = r.Email
	user.Phone = r.Phone
	user.Extend = r.Extend
	user.Password = r.Password

	if errs := user.Validate(); len(errs) != 0 {
		log.Errorf("validate failed %s .", errs.ToAggregate().Error())
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	user.Password, _ = auth.Encrypt(user.Password)

	if err := u.srv.Users().Update(c, user, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}

func checkOldPassword(user *v1.Policy, oldPassword string) error {
	if oldPassword == "" {
		return nil
	}

	log.Infof("old password %s, db password %s", oldPassword, user.Password)
	if err := user.Compare(oldPassword); err != nil {
		return errors.WithCode(code.ErrPasswordIncorrect, err.Error())
	}

	return nil
}
