package models

type ArticleModel struct {
	NewsId int `json:"id" gorm:"column:id" uri:"id" binding:"required"`
	NewsTitle  string `json:"title" gorm:"column:title"`
	NewsView int `json:"view" gorm:"column:view"`
}

func NewArticleModel() *ArticleModel {
	return &ArticleModel{}
}
func (this *ArticleModel) String() string {
	return ""
}
