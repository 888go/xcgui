package 炫彩对象基类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// UI 可视对象.
type UI struct {
	ObjectBase
}

// SetStyle 可视对象_置样式, 设置UI对象样式.
//
//	@param nStyle xcc.XC_OBJECT_STYLE, 样式值: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
//	@return int
func (u *UI) X置样式(样式值 炫彩常量类.XC_OBJECT_STYLE) int {
	return 炫彩基类.X可视对象_置样式(u.Handle, 样式值)
}

// GetStyle 可视对象_取样式, 获取UI对象样式.
//
//	@return xcc.XC_OBJECT_STYLE 返回: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
func (u *UI) X取样式() 炫彩常量类.XC_OBJECT_STYLE {
	return 炫彩基类.X可视对象_取样式(u.Handle)
}

// EnableCSS 可视对象_启用CSS, 启用或禁用样式.
//
//	@param bEnable 是否启用.
//	@return int
func (u *UI) X启用CSS(是否启用 bool) int {
	return 炫彩基类.X可视对象_启用CSS(u.Handle, 是否启用)
}

// SetCssName 可视对象_置CSS名称, 设置CSS[套用样式]名称.
//
//	@param pName 套用样式名称.
//	@return int
func (u *UI) X置CSS名称(套用样式名称 string) int {
	return 炫彩基类.X可视对象_置CSS名称(u.Handle, 套用样式名称)
}

// GetCssName 可视对象_取CSS名称, 获取CSS样式名称.
//
//	@return string
func (u *UI) X取CSS名称() string {
	return 炫彩基类.X可视对象_取CSS名称(u.Handle)
}
