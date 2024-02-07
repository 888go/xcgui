package 炫彩对象基类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// ObjectBase 炫彩对象基类.
type ObjectBase struct {
	Handle int // 句柄.
}

// 给本类的Handle赋值.
func (o *ObjectBase) X设置句柄(句柄 int) {
	o.Handle = 句柄
}

// 炫彩对象_取类型, 获取对象最终类型, 返回对象类型: XC_.
func (o *ObjectBase) X取类型() 炫彩常量类.XC_OBJECT_TYPE {
	return 炫彩基类.X炫彩对象_取类型(o.Handle)
}

// 炫彩对象_取基础类型, 获取对象的基础类型, 返回对象类型, 以下类型之一: XC_ERROR, XC_WINDOW, XC_ELE, XC_SHAPE, XC_ADAPTER.
func (o *ObjectBase) X取基础类型() int {
	return 炫彩基类.X炫彩对象_取基础类型(o.Handle)
}

// 炫彩对象_取类型扩展, 获取对象扩展类型, 返回对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *ObjectBase) X取类型EX() 炫彩常量类.XC_OBJECT_TYPE_EX {
	return 炫彩基类.X炫彩对象_取类型EX(o.Handle)
}

// 炫彩对象_置类型扩展, 如果是按钮, 请使用按钮的增强接口 XBtn_SetTypeEx().
//
// nType: 对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *ObjectBase) X置类型EX(对象类型EX 炫彩常量类.XC_OBJECT_TYPE_EX) int {
	return 炫彩基类.X炫彩对象_置类型EX(o.Handle, 对象类型EX)
}

// 炫彩_置属性, 设置对象属性.
//
// pName: 属性名.
//
// pValue: 属性值.
func (o *ObjectBase) X置属性(属性名 string, 属性值 string) bool {
	return 炫彩基类.X置属性(o.Handle, 属性名, 属性值)
}

// 炫彩_取属性, 获取对象属性, 返回属性值.
//
// pName: 属性名.
func (o *ObjectBase) X取属性(属性名 string) string {
	return 炫彩基类.X取属性(o.Handle, 属性名)
}
