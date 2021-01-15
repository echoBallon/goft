package main

import (
	"goft/src/classes"
	"goft/src/goft"
	"goft/src/middlewares"
)

func main() {
	//fmt.Println(goft.InitConfig().Config)
	goft.Ignite().
		DB(goft.NewGormAdapter(),goft.NewXOrmAdapter()).//初始化db
		Attach(middlewares.NewUserMiddleware(),middlewares.NewAuthMiddleware()).//初始化中间件
		Mount("v1", classes.NewIndexClass(), classes.NewUserClass()).//加载应用
		Launch()//开始执行
}