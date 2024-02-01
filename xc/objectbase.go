package xc

import "e.coding.net/gogit/go/xcgui/xcc"

// 炫彩对象_取类型, 获取对象最终类型, 返回对象类型: XC_.
//
// hXCGUI: 对象句柄.
func XObj_GetType(hXCGUI int) xcc.I常量_对象句柄类型_ {
	r, _, _ := xObj_GetType.Call(uintptr(hXCGUI))
	return xcc.I常量_对象句柄类型_(r)
}

// 炫彩对象_取基础类型, 获取对象的基础类型, 返回对象类型, 以下类型之一: I常量_对象句柄类型_错误类型, I常量_对象句柄类型_窗口, I常量_对象句柄类型_基础元素, I常量_对象句柄类型_形状对象, XC_ADAPTER.
//
// hXCGUI: 对象句柄.
func XObj_GetTypeBase(hXCGUI int) int {
	r, _, _ := xObj_GetTypeBase.Call(uintptr(hXCGUI))
	return int(r)
}

// 炫彩对象_取类型扩展, 获取对象扩展类型, 返回对象扩展类型: button_type_ , element_type_ , xc_ex_error.
//
// hXCGUI: 对象句柄.
func XObj_GetTypeEx(hXCGUI int) xcc.I常量_对象扩展类型_ {
	r, _, _ := xObj_GetTypeEx.Call(uintptr(hXCGUI))
	return xcc.I常量_对象扩展类型_(r)
}

// 炫彩对象_置类型扩展, 如果是按钮, 请使用按钮的增强接口 XBtn_SetTypeEx().
//
// hXCGUI: 对象句柄.
//
// nType: 对象扩展类型: button_type_ , element_type_ , xc_ex_error.
func XObj_SetTypeEx(hXCGUI int, nType xcc.I常量_对象扩展类型_) int {
	r, _, _ := xObj_SetTypeEx.Call(uintptr(hXCGUI), uintptr(nType))
	return int(r)
}
