package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 进度条.
type ProgressBar struct {
	Element
}

// I进度条_创建, 创建进度条元素.
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
func I进度条_创建(x int, y int, cx int, cy int, hParent int) *ProgressBar {
	p := &ProgressBar{}
	p.SetHandle(xc.XProgBar_Create(x, y, cx, cy, hParent))
	return p
}

// I进度条_从句柄创建对象.
func I进度条_从句柄创建(handle int) *ProgressBar {
	p := &ProgressBar{}
	p.SetHandle(handle)
	return p
}

// I进度条_从name创建对象, 失败返回nil.
func I进度条_从name创建(name string) *ProgressBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I进度条_从UID创建对象, 失败返回nil.
func I进度条_从UID创建(nUID int) *ProgressBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I进度条_从UID名称创建对象, 失败返回nil.
func I进度条_从UID名称创建(name string) *ProgressBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 进度条_置范围, 设置范围.
//
// range_: 范围.
func (p *ProgressBar) I置范围(range_ int) int {
	return xc.XProgBar_SetRange(p.I句柄, range_)
}

// 进度条_取范围.
func (p *ProgressBar) I取范围() int {
	return xc.XProgBar_GetRange(p.I句柄)
}

// 进度条_置进度图片.
//
// hImage: 图片句柄.
func (p *ProgressBar) I置进度图片(hImage int) int {
	return xc.XProgBar_SetImageLoad(p.I句柄, hImage)
}

// 进度条_置进度, 设置位置点.
//
// pos: 位置点.
func (p *ProgressBar) I置进度(pos int) int {
	return xc.XProgBar_SetPos(p.I句柄, pos)
}

// 进度条_取进度, 获取当前位置点.
func (p *ProgressBar) I取进度() int {
	return xc.XProgBar_GetPos(p.I句柄)
}

// 进度条_置水平, 设置水平或垂直.
//
// bHorizon: 水平或垂直.
func (p *ProgressBar) I置水平(bHorizon bool) int {
	return xc.XProgBar_EnableHorizon(p.I句柄, bHorizon)
}

// 进度条_启用缩放, 缩放进度贴图为当前进度区域(当前进度所显示区域), 否则为整体100进度区域.
//
// bStretch: 缩放.
func (p *ProgressBar) I启用缩放(bStretch bool) bool {
	return xc.XProgBar_EnableStretch(p.I句柄, bStretch)
}

// 进度条_启用进度文本 显示进度值文本.
//
// bShow: 是否启用.
func (p *ProgressBar) I启用进度文本(bShow bool) bool {
	return xc.XProgBar_EnableShowText(p.I句柄, bShow)
}

/*
以下都是事件
*/

type XE_PROGRESSBAR_CHANGE func(pos int, pbHandled *bool) int            // 进度条元素,进度改变事件.
type XE_PROGRESSBAR_CHANGE1 func(hEle int, pos int, pbHandled *bool) int // 进度条元素,进度改变事件.

// 进度条元素, I事件_进度改变 事件.
func (p *ProgressBar) I事件_进度改变(pFun XE_PROGRESSBAR_CHANGE) bool {
	return xc.XEle_RegEventC(p.I句柄, xcc.XE_PROGRESSBAR_CHANGE, pFun)
}

// 进度条元素, I事件_进度改变 事件.
func (p *ProgressBar) I事件_进度改变1(pFun XE_PROGRESSBAR_CHANGE1) bool {
	return xc.XEle_RegEventC1(p.I句柄, xcc.XE_PROGRESSBAR_CHANGE, pFun)
}
