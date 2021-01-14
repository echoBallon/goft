package main

import (
	"goft/src/classes"
	"goft/src/goft"
	"goft/src/middlewares"
)

func main() {
	//fmt.Println(goft.InitConfig().Config)
	goft.Ignite().
		DB(goft.NewGormAdapter(),goft.NewXOrmAdapter()).
		Attach(middlewares.NewUserMiddleware(),middlewares.NewAuthMiddleware()).
		Mount("v1", classes.NewIndexClass(), classes.NewUserClass()).
		Launch()
}