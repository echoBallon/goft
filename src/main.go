package main

import (
	. "goft/src/classes"
	"goft/src/goft"
	. "goft/src/middlewares"
)

func main() {
	goft.Ignite().
		Attach(NewUserMiddleware()).
		Mount("v1", NewIndexClass(), NewUserClass()).
		Launch()
}
