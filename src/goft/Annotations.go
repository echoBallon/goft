package goft

import (
	"fmt"
	"reflect"
	"strings"
)

//注
type Annotation interface {
	SetTag(tag reflect.StructTag)
}

var AnnotationsList []Annotation

/**
判断当前的注入对象是否是注解
*/
func IsAnnotation(t reflect.Type) bool {
	for _, item := range AnnotationsList {
		if reflect.TypeOf(item) == t {
			return true
		}
	}
	return false
}

/**
init
*/
func init() {
	AnnotationsList = make([]Annotation, 0)
	AnnotationsList = append(AnnotationsList, new(Value))
}

type Value struct {
	tag reflect.StructTag
	Beanfactory *BeanFactory
}

func (this *Value) SetTag(tag reflect.StructTag) {
	this.tag = tag
}
func(this *Value) String() string {
	get_prefix:=this.tag.Get("prefix")
	if get_prefix==""{
		return ""
	}
	prefix:=strings.Split(get_prefix,".")
	if config:=this.Beanfactory.GetBean(new(SysConfig));config!=nil{
		get_value:=GetConfigValue(config.(*SysConfig).Config,prefix,0)
		fmt.Println(get_value)
		if get_value!=nil{
			return fmt.Sprintf("%v",get_value)
		}else{
			return ""
		}
	}else{
		return ""
	}
}

