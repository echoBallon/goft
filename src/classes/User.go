package classes

import (
	"github.com/gin-gonic/gin"
	"goft/src/goft"
	"net/http"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}


func (this *UserClass) UserList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
		"message": "user success",
		})
	}
}
/**
router register
 */
func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET","/user", this.UserList()) //index
}
