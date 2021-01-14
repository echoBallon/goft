package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Goft struct {
	*gin.Engine
	g     *gin.RouterGroup
	props []interface{}
}

func Ignite() *Goft {
	g := &Goft{Engine: gin.Default(), props: make([]interface{}, 0)}
	g.Use(ErrorHandler()) //must use error middleware
	return g
}
func (this *Goft) Launch() {
	config := InitConfig()
	this.Run(fmt.Sprintf(":%d" ,config.Server.Port))
}

/**
add middleware
*/
func (this *Goft) Attach(fs...Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		for _,f :=range fs {
			err := f.OnRequest(context)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
		context.Next()
	})
	return this
}

/**
init db
*/
func (this *Goft) DB(beans ...interface{}) *Goft {
	this.props = append(this.props, beans...)
	return this
}

/**
overwrite
*/
func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := Convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}

// 加一个group 参数
func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
		this.setProp(class)
	}
	return this
}

//获取属性
func (this *Goft) getProp(t reflect.Type) interface{} {
	for _, p := range this.props {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}
func (this *Goft) setProp(class IClass) {
	vClass := reflect.ValueOf(class).Elem()
	for i := 0; i < vClass.NumField(); i++ {
		f := vClass.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}
		if p := this.getProp(f.Type()); p != nil {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}
