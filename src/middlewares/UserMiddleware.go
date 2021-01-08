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
	fmt.Println("这是用户中间件")
	fmt.Println(context.Query("name"))
	return nil
}
