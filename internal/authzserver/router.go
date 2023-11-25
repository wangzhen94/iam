package authzserver

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/spf13/viper"
	"github.com/wangzhen94/iam/internal/authzserver/controller/v1/authorize"
	"github.com/wangzhen94/iam/internal/authzserver/load/cache"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/pkg/log"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	auth := newCacheAuth()
	g.NoRoute(auth.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "page not found."), nil)
	})

	cacheIns, _ := cache.GetCacheInsOr(nil)
	if cacheIns == nil {
		log.Panicf("get nil cache instance")
	}

	apiv1 := g.Group("/v1", auth.AuthFunc())
	{
		authzController := authorize.NewAuthzController(cacheIns)

		// Router for authorization
		apiv1.POST("/authz", authzController.Authorize)
	}

	demo := g.Group("")

	demo.GET("/config/:key", func(c *gin.Context) {
		mode := viper.GetString(c.Param("key"))

		core.WriteResponse(c, nil, mode)
	})

	return g
}
