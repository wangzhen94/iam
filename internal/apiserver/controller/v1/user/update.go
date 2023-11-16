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
	var r *v1.User
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

	user.Nickname = r.Nickname
	user.Email = r.Email
	user.Phone = r.Phone
	user.Extend = r.Extend

	err = changePassword(user, r, oldPassword)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrPasswordIncorrect, err.Error()), nil)

		return
	}

	if errs := user.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	if err := u.srv.Users().Update(c, user, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}

func changePassword(db *v1.User, req *v1.User, oldPassword string) (err error) {
	if oldPassword == "" {
		return nil
	}
	if err := db.Compare(oldPassword); err != nil {
		return errors.WithCode(code.ErrPasswordIncorrect, err.Error())
	}

	db.Password, err = auth.Encrypt(req.Password)

	return err
}
