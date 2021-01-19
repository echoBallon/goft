package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goft/src/goft"
	"goft/src/models"
	"log"
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
	goft.Error(this.Table("articles").Where("id=?", article.NewsId).First(&article).Error)
	//协程任务
	goft.Task(this.UpdateViews, this.UpdateViewsDone, article.NewsId)
	return article
}

func (this *ArticleClass) UpdateViews(params ...interface{}) {
	this.Table("articles").Where("id=?", params[0]).Update("view", gorm.Expr("view+1"))
}
func (this *ArticleClass) UpdateViewsDone() {
	log.Println("update view done")
}
func (this *ArticleClass) Test() interface{}{
	log.Println("test")
	return nil
}

/**
router register
*/
func (this *ArticleClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/article/:id", this.ArticleDetail)
}

func(this *ArticleClass) Name() string{
	return "ArticleClass"
}