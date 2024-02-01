package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
)

// 静态文本链接按钮.
type TextLink struct {
	Button
}

// I文本链接_创建, 创建静态文本链接元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 文本内容.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func I文本链接_创建(x int, y int, cx int, cy int, pName string, hParent int) *TextLink {
	p := &TextLink{}
	p.SetHandle(xc.XTextLink_Create(x, y, cx, cy, pName, hParent))
	return p
}

// I文本链接_从句柄创建对象.
func I文本链接_从句柄创建(handle int) *TextLink {
	p := &TextLink{}
	p.SetHandle(handle)
	return p
}

// I文本链接_从name创建对象, 失败返回nil.
func I文本链接_从name创建(name string) *TextLink {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &TextLink{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I文本链接_从UID创建对象, 失败返回nil.
func I文本链接_从UID创建(nUID int) *TextLink {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &TextLink{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I文本链接_从UID名称创建对象, 失败返回nil.
func I文本链接_从UID名称创建(name string) *TextLink {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &TextLink{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 文本链接_启用离开状态下划线, 启用下划线, 鼠标离开状态.
//
// bEnable: 是否启用.
func (t *TextLink) I启用离开状态下划线(bEnable bool) int {
	return xc.XTextLink_EnableUnderlineLeave(t.I句柄, bEnable)
}

// 文本链接_停留状态下划线, 启用下划线, 鼠标停留状态.
//
// bEnable: 是否启用.
func (t *TextLink) I启用停留状态下划线(bEnable bool) int {
	return xc.XTextLink_EnableUnderlineStay(t.I句柄, bEnable)
}

// 文本链接_置停留状态文本颜色, 设置文本颜色, 鼠标停留状态.
//
// color: ABGR 颜色值.
func (t *TextLink) I置停留状态文本颜色(color int) int {
	return xc.XTextLink_SetTextColorStay(t.I句柄, color)
}

// 文本链接_置离开状态下划线颜色, 设置下划线颜色, 鼠标离开状态.
//
// color: ABGR 颜色值.
func (t *TextLink) I置离开状态下划线颜色(color int) int {
	return xc.XTextLink_SetUnderlineColorLeave(t.I句柄, color)
}

// 文本链接_置停留状态下划线颜色, 置下划线颜色, 鼠标停留状态.
//
// color: ABGR 颜色值.
func (t *TextLink) I置停留状态下划线颜色(color int) int {
	return xc.XTextLink_SetUnderlineColorStay(t.I句柄, color)
}
