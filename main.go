package main

import (
	"fmt"
	"goft/src/goft"
)

func main() {
	fmt.Println(goft.InitConfig().Config)
//	goft.Ignite().
//		DB(goft.NewGormAdapter(),goft.NewXOrmAdapter()).
//		Attach(middlewares.NewUserMiddleware()).
//		Mount("v1", classes.NewIndexClass(), classes.NewUserClass()).
//		Launch()
}