package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(context *gin.Context) error {
	fmt.Println("user-mid")
	return nil
}
