package classes

import (
	"github.com/gin-gonic/gin"
	"goft/src/goft"
	"net/http"
)

type IndexClass struct {

}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}


func (this *IndexClass) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
		"message": "index success",
		})
	}
}
/**
router register
 */
func (this *IndexClass) Build(goft *goft.Goft) {
	//goft.Handle("GET","/", this.GetIndex()) //index
}
