package goft

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			}
		}()
		context.Next()
	}
}

/**
panic error
*/
func Error(err error, msg ...string) {
	if err != nil {
		errMsg := err.Error()
		if len(msg) > 0 {
			errMsg = msg[0]
		}
		log.Println(errMsg)
		panic(errMsg)
	}
}
