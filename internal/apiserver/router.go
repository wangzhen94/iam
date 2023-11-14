package apiserver

import (
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	// Middlewares.
	//jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	//g.POST("/login", jwtStrategy.LoginHandler)
	//g.POST("/logout", jwtStrategy.LogoutHandler)
	//// Refresh time can be longer than token timeout
	//g.POST("/refresh", jwtStrategy.RefreshHandler)

	// v1 handlers, requiring authentication
	//storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	//
	//v1 := g.Group("/v1")
	//{
	//	// user RESTful resource
	//	userv1 := v1.Group("/users")
	//	{
	//		//userController := user.NewUserController(storeIns)
	//
	//		//userv1.POST("", userController.Create)
	//		//userv1.Use(auto.AuthFunc(), middleware.Validation())
	//		//// v1.PUT("/find_password", userController.FindPassword)
	//		//userv1.DELETE("", userController.DeleteCollection) // admin api
	//		//userv1.DELETE(":name", userController.Delete)      // admin api
	//		//userv1.PUT(":name/change-password", userController.ChangePassword)
	//		//userv1.PUT(":name", userController.Update)
	//		//userv1.GET("", userController.List)
	//		//userv1.GET(":name", userController.Get) // admin api
	//	}
	//
	//}

	return g
}
