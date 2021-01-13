package middlewares

import (
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(context *gin.Context) error {
	return nil
}
