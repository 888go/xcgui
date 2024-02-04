package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// SliderBar 滑动条元素.
type SliderBar struct {
	Element
}

// 创建滑动条
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父窗口句柄或元素句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func NewSliderBar(x int, y int, cx int, cy int, hParent int) *SliderBar {
	p := &SliderBar{}
	p.SetHandle(xc.XSliderBar_Create(x, y, cx, cy, hParent))
	return p
}

// 创建滑动条并按句柄
func NewSliderBarByHandle(handle int) *SliderBar {
	p := &SliderBar{}
	p.SetHandle(handle)
	return p
}

// 创建滑动条并按名称, 失败返回nil.
func NewSliderBarByName(name string) *SliderBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &SliderBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建滑动条并按UID
func NewSliderBarByUID(nUID int) *SliderBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &SliderBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建滑动条并按UID名称
func NewSliderBarByUIDName(name string) *SliderBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &SliderBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 置范围, 设置滑动范围.
//
// range_: 范围.
func (s *SliderBar) SetRange(range_ int) int {
	return xc.XSliderBar_SetRange(s.Handle, range_)
}

// 取范围, 获取滚动范围.
func (s *SliderBar) GetRange() int {
	return xc.XSliderBar_GetRange(s.Handle)
}

// 置进度图片, 设置进度贴图.
//
// hImage: 图片句柄.
func (s *SliderBar) SetImageLoad(hImage int) int {
	return xc.XSliderBar_SetImageLoad(s.Handle, hImage)
}

// 置滑块宽度, 设置滑块按钮宽度.
//
// width: 宽度.
func (s *SliderBar) SetButtonWidth(width int) int {
	return xc.XSliderBar_SetButtonWidth(s.Handle, width)
}

// 置滑块高度, 设置滑块按钮高度.
//
// height: 高度.
func (s *SliderBar) SetButtonHeight(height int) int {
	return xc.XSliderBar_SetButtonHeight(s.Handle, height)
}

// 置当前位置, 设置当前进度点.
//
// pos: 进度点.
func (s *SliderBar) SetPos(pos int) int {
	return xc.XSliderBar_SetPos(s.Handle, pos)
}

// 取当前位置, 获取当前进度点.
func (s *SliderBar) GetPos() int {
	return xc.XSliderBar_GetPos(s.Handle)
}

// 取滑块, 返回滑块按钮句柄.
func (s *SliderBar) GetButton() int {
	return xc.XSliderBar_GetButton(s.Handle)
}

// 置水平, 设置水平或垂直.
//
// bHorizon: 水平或垂直.
func (s *SliderBar) EnableHorizon(bHorizon bool) int {
	return xc.XSliderBar_EnableHorizon(s.Handle, bHorizon)
}

/*
以下都是事件
*/

type XE_SLIDERBAR_CHANGE func(pos int32, pbHandled *bool) int            //事件_滑块位置改变事件.
type XE_SLIDERBAR_CHANGE1 func(hEle int, pos int32, pbHandled *bool) int //事件_滑块位置改变事件.

// 事件_滑块位置改变
func (s *SliderBar) Event_SLIDERBAR_CHANGE(pFun XE_SLIDERBAR_CHANGE) bool {
	return xc.XEle_RegEventC(s.Handle, xcc.XE_SLIDERBAR_CHANGE, pFun)
}

// 事件_滑块位置改变1
func (s *SliderBar) Event_SLIDERBAR_CHANGE1(pFun XE_SLIDERBAR_CHANGE1) bool {
	return xc.XEle_RegEventC1(s.Handle, xcc.XE_SLIDERBAR_CHANGE, pFun)
}
