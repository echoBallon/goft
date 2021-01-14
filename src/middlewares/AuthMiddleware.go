package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
}
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}
func (this *AuthMiddleware) OnRequest(context *gin.Context) error {
	fmt.Println("auth-mid")
	return nil
}
