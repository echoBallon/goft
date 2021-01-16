package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	this.Run(fmt.Sprintf(":%d", config.Server.Port))
}

/**
add middleware
*/
func (this *Goft) Attach(fs ...Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		for _, f := range fs {
			err := f.OnRequest(context)
			Error(err)
		}
		context.Next()
	})
	return this
}

/**
init db
*/
func (this *Goft) Beans(beans ...interface{}) *Goft {
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
	//设置当前的路由需要设置的group
	this.g = this.Group(group)
	for _, class := range classes {
		//执行传入函数的build 注册路由
		class.Build(this)
		//注入初始化db
		this.setProp(class)
	}
	return this
}

//获取属性
func (this *Goft) getProp(t reflect.Type) interface{} {
	//寻找当前db注入的是否能包含我们需要使用的Orm
	for _, p := range this.props {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}
func (this *Goft) setProp(class IClass) {
	//获struct值
	vClass := reflect.ValueOf(class).Elem()
	//循环获取struct参数
	for i := 0; i < vClass.NumField(); i++ {
		f := vClass.Field(i)
		//如果参数不存在 或者不是指针 就退出
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}
		if p := this.getProp(f.Type()); p != nil {
			//初始化
			f.Set(reflect.New(f.Type().Elem()))
			//赋值
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}
