package goft

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Ignite() *Goft {
	return &Goft{Engine: gin.New()}
}
func (this *Goft) Launch() {
	this.Run(":8080")
}

/**
add middleware
*/
func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		err := f.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.Next()
		}
	})
	return this
}

/**
overwrite
*/
func (this *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Goft {
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}

//mount classes
func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}
