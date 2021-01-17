package classes

import (
	"github.com/gin-gonic/gin"
	"goft/src/goft"
	"goft/src/models"
)

type ArticleClass struct {
	*goft.GormAdapter
}

func NewArticleClass() *ArticleClass {
	return &ArticleClass{}
}

func (this *ArticleClass) ArticleDetail(ctx *gin.Context) goft.Model {
	article := models.NewArticleModel()
	err := ctx.BindUri(article)
	goft.Error(err)
	goft.Error(this.Table("articles").Where("id=?", article.NewsId).Find(&article).Error)
	return article
}


/**
router register
*/
func (this *ArticleClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/article/:id", this.ArticleDetail)
}
