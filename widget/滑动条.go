package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 滑动条元素.
type SliderBar struct {
	Element
}

// I滑动条_创建, 创建滑动条元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func I滑动条_创建(x int, y int, cx int, cy int, hParent int) *SliderBar {
	p := &SliderBar{}
	p.SetHandle(xc.XSliderBar_Create(x, y, cx, cy, hParent))
	return p
}

// I滑动条_从句柄创建对象.
func I滑动条_从句柄创建(handle int) *SliderBar {
	p := &SliderBar{}
	p.SetHandle(handle)
	return p
}

// I滑动条_从name创建对象, 失败返回nil.
func I滑动条_从name创建(name string) *SliderBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &SliderBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I滑动条_从UID创建对象, 失败返回nil.
func I滑动条_从UID创建(nUID int) *SliderBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &SliderBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I滑动条_从UID名称创建对象, 失败返回nil.
func I滑动条_从UID名称创建(name string) *SliderBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &SliderBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 滑动条_置范围, 设置滑动范围.
//
// range_: 范围.
func (s *SliderBar) I置范围(range_ int) int {
	return xc.XSliderBar_SetRange(s.I句柄, range_)
}

// 滑动条_取范围, 获取滚动范围.
func (s *SliderBar) I取范围() int {
	return xc.XSliderBar_GetRange(s.I句柄)
}

// 滑动条_置进度图片, 设置进度贴图.
//
// hImage: 图片句柄.
func (s *SliderBar) I置进度图片(hImage int) int {
	return xc.XSliderBar_SetImageLoad(s.I句柄, hImage)
}

// 滑动条_置滑块宽度, 设置滑块按钮宽度.
//
// width: 宽度.
func (s *SliderBar) I置滑块宽度(width int) int {
	return xc.XSliderBar_SetButtonWidth(s.I句柄, width)
}

// 滑动条_置滑块高度, 设置滑块按钮高度.
//
// height: 高度.
func (s *SliderBar) I置滑块高度(height int) int {
	return xc.XSliderBar_SetButtonHeight(s.I句柄, height)
}

// 滑动条_置当前位置, 设置当前进度点.
//
// pos: 进度点.
func (s *SliderBar) I置当前位置(pos int) int {
	return xc.XSliderBar_SetPos(s.I句柄, pos)
}

// 滑动条_取当前位置, 获取当前进度点.
func (s *SliderBar) I取当前位置() int {
	return xc.XSliderBar_GetPos(s.I句柄)
}

// 滑动条_取滑块, 返回滑块按钮句柄.
func (s *SliderBar) I取滑块() int {
	return xc.XSliderBar_GetButton(s.I句柄)
}

// 滑动条_置水平, 设置水平或垂直.
//
// bHorizon: 水平或垂直.
func (s *SliderBar) I置水平(bHorizon bool) int {
	return xc.XSliderBar_EnableHorizon(s.I句柄, bHorizon)
}

/*
以下都是事件
*/

type XE_SLIDERBAR_CHANGE func(pos int, pbHandled *bool) int            // 滑动条元素,滑块位置改变事件.
type XE_SLIDERBAR_CHANGE1 func(hEle int, pos int, pbHandled *bool) int // 滑动条元素,滑块位置改变事件.

// 滑动条元素, I事件_滑块位置改变 事件.
func (s *SliderBar) I事件_滑块位置改变(pFun XE_SLIDERBAR_CHANGE) bool {
	return xc.XEle_RegEventC(s.I句柄, xcc.XE_SLIDERBAR_CHANGE, pFun)
}

// 滑动条元素,滑块位置改变事件.
func (s *SliderBar) I事件_滑块位置改变1(pFun XE_SLIDERBAR_CHANGE1) bool {
	return xc.XEle_RegEventC1(s.I句柄, xcc.XE_SLIDERBAR_CHANGE, pFun)
}
