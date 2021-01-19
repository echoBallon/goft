package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g           *gin.RouterGroup
	beanFactory *BeanFactory
	exprData    map[string]interface{}
}

func Ignite() *Goft {
	g := &Goft{Engine: gin.Default(), beanFactory: NewBeanFactory(), exprData: map[string]interface{}{}}
	g.Use(ErrorHandler()) //must use error middleware
	config := InitConfig()
	g.beanFactory.setBean(config) //整个配置加载进bean中
	if config.Server.Html != "" {
		g.LoadHTMLGlob("src/" + config.Server.Html)
	}
	return g
}
func (this *Goft) Launch() {
	var port int32 = 8080
	if config := this.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	getCronTask().Start()
	this.Run(fmt.Sprintf(":%d", port))
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
func (this *Goft) Beans(beans ...Bean) *Goft {
	// 取出bean的名称 ，然后 加入到exprData里面
	for _, bean := range beans {
		this.exprData[bean.Name()] = bean
	}
	this.beanFactory.setBean(beans...)
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
		this.beanFactory.inject(class)
		this.Beans(class) //控制器 也作为bean加入到bean容器
	}
	return this
}

/**
cron task
*/
func (this *Goft) Task(cron string, expr interface{}) *Goft {
	var err error
	//func
	if f, ok := expr.(func()); ok {
		_, err = getCronTask().AddFunc(cron, f)
	}
	// string
	if exp, ok := expr.(Expr); ok {
		_, err = getCronTask().AddFunc(cron, func() {
			_, expErr := ExecExpr(exp, this.exprData)
			Error(expErr)
		})
	}
	Error(err)
	return this
}
