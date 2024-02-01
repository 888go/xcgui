package objectbase

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 炫彩对象基类.
type ObjectBase struct {
	I句柄 int // 句柄.
}

// 给本类的Handle赋值.
func (o *ObjectBase) SetHandle(handle int) {
	o.I句柄 = handle
}

// 炫彩对象_取类型, 获取对象最终类型, 返回对象类型: XC_.
func (o *ObjectBase) GetType() xcc.I常量_对象句柄类型_ {
	return xc.XObj_GetType(o.I句柄)
}

// 炫彩对象_取基础类型, 获取对象的基础类型, 返回对象类型, 以下类型之一: I常量_对象句柄类型_错误类型, I常量_对象句柄类型_窗口, I常量_对象句柄类型_基础元素, I常量_对象句柄类型_形状对象, XC_ADAPTER.
func (o *ObjectBase) GetTypeBase() int {
	return xc.XObj_GetTypeBase(o.I句柄)
}

// 炫彩对象_取类型扩展, 获取对象扩展类型, 返回对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *ObjectBase) GetTypeEx() xcc.I常量_对象扩展类型_ {
	return xc.XObj_GetTypeEx(o.I句柄)
}

// 炫彩对象_置类型扩展, 如果是按钮, 请使用按钮的增强接口 XBtn_SetTypeEx().
//
// nType: 对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func (o *ObjectBase) SetTypeEx(nType xcc.I常量_对象扩展类型_) int {
	return xc.XObj_SetTypeEx(o.I句柄, nType)
}
