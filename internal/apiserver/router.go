// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/spf13/viper"
	"github.com/wangzhen94/iam/internal/apiserver/controller/v1/policy"
	"github.com/wangzhen94/iam/internal/apiserver/controller/v1/secret"
	"github.com/wangzhen94/iam/internal/apiserver/controller/v1/user"
	"github.com/wangzhen94/iam/internal/apiserver/store/mysql"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/internal/pkg/middleware"
	"github.com/wangzhen94/iam/internal/pkg/middleware/auth"
	_ "github.com/wangzhen94/iam/pkg/validator"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	// Middlewares.
	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	// Refresh time can be longer than token timeout
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "Page not found."), nil)
	})

	// v1 handlers, requiring authentication
	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	v1 := g.Group("/v1")
	{
		// user RESTful resource
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)

			userv1.POST("", userController.Create)
			userv1.Use(auto.AuthFunc(), middleware.Validation())
			userv1.DELETE(":name", userController.Delete) // admin api
			userv1.PUT(":name/change_password", userController.ChangePassword)
			userv1.PUT(":name", userController.Update)
			userv1.GET("", userController.List)
			userv1.GET(":name", userController.Get) // admin api
		}

		v1.Use(auto.AuthFunc())

		// policy RESTful resource
		policyv1 := v1.Group("/policies")
		{
			policyController := policy.NewPolicyController(storeIns)

			policyv1.POST("", policyController.Create)
			policyv1.DELETE(":name", policyController.Delete)
			policyv1.PUT(":name", policyController.Update)
			policyv1.GET("", policyController.List)
			policyv1.GET(":name", policyController.Get)
		}

		// secret RESTful resource
		secretv1 := v1.Group("/secrets")
		{
			secretController := secret.NewSecretController(storeIns)

			secretv1.POST("", secretController.Create)
			secretv1.DELETE(":name", secretController.Delete)
			secretv1.PUT(":name", secretController.Update)
			secretv1.GET("", secretController.List)
			secretv1.GET(":name", secretController.Get)
		}
	}

	demo := g.Group("")

	demo.GET("/config/:key", func(c *gin.Context) {
		mode := viper.GetString(c.Param("key"))

		core.WriteResponse(c, nil, mode)
	})

	return g
}
