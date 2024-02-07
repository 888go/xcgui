package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// ProgressBar 进度条.
type ProgressBar struct {
	Element
}

// 进度条_创建, 创建进度条元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口.
func X创建进度条(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *ProgressBar {
	p := &ProgressBar{}
	p.X设置句柄(炫彩基类.X进度条_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建进度条并按句柄(handle int) *ProgressBar {
	p := &ProgressBar{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建进度条并按名称(name string) *ProgressBar {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ProgressBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建进度条并按UID(nUID int) *ProgressBar {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ProgressBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建进度条并按UID名称(name string) *ProgressBar {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ProgressBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 进度条_置范围, 设置范围.
//
// range_: 范围.
func (p *ProgressBar) X置范围(范围 int) int {
	return 炫彩基类.X进度条_置范围(p.Handle, 范围)
}

// 进度条_取范围.
func (p *ProgressBar) X取范围() int {
	return 炫彩基类.X进度条_取范围(p.Handle)
}

// 进度条_置进度图片.
//
// hImage: 图片句柄.
func (p *ProgressBar) X置进度图片(图片句柄 int) int {
	return 炫彩基类.X进度条_置进度图片(p.Handle, 图片句柄)
}

// 进度条_置进度, 设置位置点.
//
// pos: 位置点.
func (p *ProgressBar) X置进度(位置点 int) int {
	return 炫彩基类.X进度条_置进度(p.Handle, 位置点)
}

// 进度条_取进度, 获取当前位置点.
func (p *ProgressBar) X取进度() int {
	return 炫彩基类.X进度条_取进度(p.Handle)
}

// 进度条_置水平, 设置水平或垂直.
//
// bHorizon: 水平或垂直.
func (p *ProgressBar) X置水平(水平或垂直 bool) int {
	return 炫彩基类.X进度条_置水平(p.Handle, 水平或垂直)
}

// 进度条_启用缩放, 缩放进度贴图为当前进度区域(当前进度所显示区域), 否则为整体100进度区域.
//
// bStretch: 缩放.
func (p *ProgressBar) X启用缩放(缩放 bool) bool {
	return 炫彩基类.X进度条_启用缩放(p.Handle, 缩放)
}

// 进度条_启用进度文本 显示进度值文本.
//
// bShow: 是否启用.
func (p *ProgressBar) X启用进度文本(是否启用 bool) bool {
	return 炫彩基类.X进度条_启用进度文本(p.Handle, 是否启用)
}

// 进度条_置进度颜色. 设置进度颜色.
//
// color: ABGR 颜色.
func (p *ProgressBar) X置进度颜色(ABGR颜色 int) bool {
	return 炫彩基类.X进度条_置进度颜色(p.Handle, ABGR颜色)
}

/*
以下都是事件
*/

type XE_PROGRESSBAR_CHANGE func(pos int32, pbHandled *bool) int            // 进度条元素,进度改变事件.
type XE_PROGRESSBAR_CHANGE1 func(hEle int, pos int32, pbHandled *bool) int // 进度条元素,进度改变事件.

// 进度条元素,进度改变事件.
func (p *ProgressBar) X事件_进度改变事件(pFun XE_PROGRESSBAR_CHANGE) bool {
	return 炫彩基类.X元素_注册事件C(p.Handle, 炫彩常量类.XE_PROGRESSBAR_CHANGE, pFun)
}

// 进度条元素,进度改变事件.
func (p *ProgressBar) X事件_进度改变1(pFun XE_PROGRESSBAR_CHANGE1) bool {
	return 炫彩基类.X元素_注册事件C1(p.Handle, 炫彩常量类.XE_PROGRESSBAR_CHANGE, pFun)
}
