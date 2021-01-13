package models

import "fmt"

type UserModel struct {
	UserId   int `uri:"id" binding:"required,gt=0"`
	UserName string
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this *UserModel) String() string {
	return fmt.Sprintf("userId:%d,username:%s", this.UserId, this.UserName)
}
