package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/apiserver/controller/v1/user"
	"github.com/wangzhen94/iam/internal/apiserver/store/mysql"
	"github.com/wangzhen94/iam/internal/pkg/code"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	// Middlewares.

	storeIns, _ := mysql.GetMySQLFactoryOr(nil)

	v1 := g.Group("/v1")

	{
		userV1 := v1.Group("/user")
		userController := user.NewUserController(storeIns)

		userV1.POST("", userController.Create)
		userV1.DELETE("/:name", userController.Delete)

	}

	g.GET("/user/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		tp := c.Query("type")
		if err := getUser(name); err != nil {
			core.WriteResponse(c, err, nil)
			return
		}

		core.WriteResponse(c, nil, map[string]string{"email": name + "@foxmail.com",
			"type": tp})
	})

	return g
}

func getUser(name string) error {
	if err := queryDataBase(name); err != nil {
		return errors.Wrap(err, "get user error")
	}
	return nil
}

func queryDataBase(name string) error {
	if "wang" == name {
		return nil
	} else {
		return errors.WithCode(code.ErrDatabase, "user '%s' not found", name)
	}
}
