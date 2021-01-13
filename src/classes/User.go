package classes

import (
	"github.com/gin-gonic/gin"
	"goft/src/goft"
	"goft/src/models"
)

type UserClass struct {
	*goft.GormAdapter
	 //*goft.XOrmAdapter
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserName(ctx *gin.Context) string {
	return "echo"
}
func (this *UserClass) UserList(ctx *gin.Context) goft.Models {
	users := []*models.UserModel{
		&models.UserModel{UserId: 101, UserName: "echo"},
		&models.UserModel{UserId: 102, UserName: "zhanshan"},
	}
	return goft.MakeModels(users)
}

func (this *UserClass) UserDetail(ctx *gin.Context) goft.Model {
	user := models.NewUserModel()
	err := ctx.BindUri(user)
	goft.Error(err)
	this.Table("users").Where("user_id=?", user.UserId).Find(&user)
	//has, err := this.Table("users").Where("user_id=?", user.UserId).Get(user)
	//if !has {
	//	goft.Error(fmt.Errorf("数据不存在"))
	//}
	//goft.Error(err)
	return user
}

/**
router register
*/
func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user_name", this.UserName)         //index
	goft.Handle("GET", "/user_list", this.UserList)         //index
	goft.Handle("GET", "/user_detail/:id", this.UserDetail) //index
}
