package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/apiserver/controller/v1/policy"
	"github.com/wangzhen94/iam/internal/apiserver/controller/v1/secret"
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
		{
			userV1 := v1.Group("/user")
			userController := user.NewUserController(storeIns)

			userV1.POST("", userController.Create)
			userV1.DELETE("/:name", userController.Delete)
			userV1.GET("", userController.List)
			userV1.PUT("", userController.Update)
			userV1.GET("/:name", userController.Get)
		}

		{
			secretV1 := v1.Group("/secret")
			secretController := secret.NewSecretController(storeIns)
			secretV1.GET("/:name", secretController.Get)
			secretV1.POST("", secretController.Create)
			secretV1.PUT("/:name", secretController.Get)
			secretV1.GET("", secretController.List)
			secretV1.DELETE("", secretController.Delete)
		}

		{
			policyV1 := v1.Group("/policy")
			policyController := policy.NewPolicyController(storeIns)
			policyV1.POST("", policyController.Create)
			policyV1.DELETE("/:name", policyController.Delete)
			policyV1.PUT("/:name", policyController.Update)
			policyV1.GET("", policyController.List)
			policyV1.GET("/:name", policyController.Get)
		}
	}

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
