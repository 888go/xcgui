package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
)

// 静态文本链接按钮.
type TextLink struct {
	Button
}

// 文本链接_创建, 创建静态文本链接元素.
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
func X创建文本链接(x坐标 int, y坐标 int, 宽度 int, 高度 int, 文本内容 string, 父窗口句柄或元素句柄 int) *TextLink {
	p := &TextLink{}
	p.X设置句柄(炫彩基类.X文本链接_创建(x坐标, y坐标, 宽度, 高度, 文本内容, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建文本链接并按句柄(handle int) *TextLink {
	p := &TextLink{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建文本链接并按名称(name string) *TextLink {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &TextLink{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建文本链接并按UID(nUID int) *TextLink {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &TextLink{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建文本链接并按UID名称(name string) *TextLink {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &TextLink{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 文本链接_启用离开状态下划线, 启用下划线, 鼠标离开状态.
//
// bEnable: 是否启用.
func (t *TextLink) X启用离开状态下划线(是否启用 bool) int {
	return 炫彩基类.X文本链接_启用离开状态下划线(t.Handle, 是否启用)
}

// 文本链接_停留状态下划线, 启用下划线, 鼠标停留状态.
//
// bEnable: 是否启用.
func (t *TextLink) X停留状态下划线(是否启用 bool) int {
	return 炫彩基类.X文本链接_停留状态下划线(t.Handle, 是否启用)
}

// 文本链接_置停留状态文本颜色, 设置文本颜色, 鼠标停留状态.
//
// color: ABGR 颜色值.
func (t *TextLink) X置停留状态文本颜色(ABGR颜色值 int) int {
	return 炫彩基类.X文本链接_置停留状态文本颜色(t.Handle, ABGR颜色值)
}

// 文本链接_置离开状态下划线颜色, 设置下划线颜色, 鼠标离开状态.
//
// color: ABGR 颜色值.
func (t *TextLink) X置离开状态下划线颜色(ABGR颜色值 int) int {
	return 炫彩基类.X文本链接_置离开状态下划线颜色(t.Handle, ABGR颜色值)
}

// 文本链接_置停留状态下划线颜色, 置下划线颜色, 鼠标停留状态.
//
// color: ABGR 颜色值.
func (t *TextLink) X置停留状态下划线颜色(ABGR颜色值 int) int {
	return 炫彩基类.X文本链接_置停留状态下划线颜色(t.Handle, ABGR颜色值)
}
