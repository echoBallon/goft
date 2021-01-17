package classes

import (
	"github.com/gin-gonic/gin"
	"goft/src/goft"
)

type IndexClass struct {

}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}


func (this *IndexClass) GetIndex(ctx *gin.Context) goft.View {
	return  "index"
}
/**
router register
 */
func (this *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", this.GetIndex) //index
}
