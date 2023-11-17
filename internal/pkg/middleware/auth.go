package middleware

import "github.com/gin-gonic/gin"

type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

type AuthOperator struct {
	strategy AuthStrategy
}

func (o *AuthOperator) SetStrategy(strategy AuthStrategy) {
	o.strategy = strategy
}

func (o *AuthOperator) AuthFunc() gin.HandlerFunc {
	return o.strategy.AuthFunc()
}
