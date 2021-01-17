package models

type ArticleModel struct {
	NewsId int `json:"id" gorm:"column:id" uri:"id" binding:"required"`
	NewTitle  string `json:"title" gorm:"column:title"`
}

func NewArticleModel() *ArticleModel {
	return &ArticleModel{}
}
func (this *ArticleModel) String() string {
	return ""
}
