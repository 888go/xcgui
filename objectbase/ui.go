package objectbase

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// UI 可视对象.
type UI struct {
	ObjectBase
}

// I置样式 可视对象_ I置样式, 设置UI对象样式.
//
//	@param nStyle xcc.XC_OBJECT_STYLE, 样式值: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
//	@return int
func (u *UI) I置样式(nStyle xcc.XC_OBJECT_STYLE) int {
	return xc.XUI_SetStyle(u.I句柄, nStyle)
}

// I取样式 可视对象_ I取样式, 获取UI对象样式.
//
//	@return xcc.XC_OBJECT_STYLE 返回: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
func (u *UI) I取样式() xcc.XC_OBJECT_STYLE {
	return xc.XUI_GetStyle(u.I句柄)
}

// I启用CSS 可视对象_ I启用CSS, 启用或禁用样式.
//
//	@param bEnable 是否启用.
//	@return int
func (u *UI) I启用CSS(bEnable bool) int {
	return xc.XUI_EnableCSS(u.I句柄, bEnable)
}

// I置CSS名称 可视对象_ I置CSS名称, 设置CSS[套用样式]名称.
//
//	@param pName 套用样式名称.
//	@return int
func (u *UI) I置CSS名称(pName string) int {
	return xc.XUI_SetCssName(u.I句柄, pName)
}

// I取CSS名称 可视对象_ I取CSS名称, 获取CSS样式名称.
//
//	@return string
func (u *UI) I取CSS名称() string {
	return xc.XUI_GetCssName(u.I句柄)
}
