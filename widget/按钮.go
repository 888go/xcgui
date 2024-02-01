package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// Button 按钮.
type Button struct {
	Element
}

// 按钮_创建.
//
// x: 按钮x坐标.
//
// y: 按钮y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 标题.
//
// hParent: 父为窗口句柄或元素句柄.
func I按钮_创建(x int, y int, cx int, cy int, pName string, hParent int) *Button {
	p := &Button{}
	p.SetHandle(xc.XBtn_Create(x, y, cx, cy, pName, hParent))
	return p
}

// 从句柄创建对象.
func I按钮_从句柄创建(handle int) *Button {
	p := &Button{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func I按钮_从name创建(name string) *Button {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func I按钮_从UID创建(nUID int) *Button {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func I按钮_从UID名称创建(name string) *Button {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 按钮_判断选中, 是否选中状态.
func (b *Button) I判断选中() bool {
	return xc.XBtn_IsCheck(b.I句柄)
}

// 按钮_置选中, 设置选中状态.
//
// bCheck: 是否选中.
func (b *Button) I置选中(bCheck bool) bool {
	return xc.XBtn_SetCheck(b.I句柄, bCheck)
}

// I置状态 按钮_置状态.
//	@param nState 按钮状态: xcc.Common_State3_.
//	@return int
//
func (b *Button) I置状态(nState xcc.Common_State3_) int {
	return xc.XBtn_SetState(b.I句柄, nState)
}

// I取状态 按钮_取状态.
//	@return xcc.Common_State3_
//
func (b *Button) I取状态() xcc.Common_State3_ {
	return xc.XBtn_GetState(b.I句柄)
}

// I取状态扩展 按钮_取状态扩展.
//	@return xcc.I常量_按钮状态_
//
func (b *Button) I取状态扩展() xcc.I常量_按钮状态_ {
	return xc.XBtn_GetStateEx(b.I句柄)
}

// 按钮_置类型扩展, 设置按钮类型并自动修改样式和文本对齐方式.
//
// nType: 按钮类型, Button_Type_ , element_type_ , xc_ex_error.
func (b *Button) I置类型扩展(nType xcc.I常量_对象扩展类型_) int {
	return xc.XBtn_SetTypeEx(b.I句柄, nType)
}

// 按钮_置组ID.
//
// nID: 组ID.
func (b *Button) I置组ID(nID int) int {
	return xc.XBtn_SetGroupID(b.I句柄, nID)
}

// 按钮_取组ID.
func (b *Button) I取组ID() int {
	return xc.XBtn_GetGroupID(b.I句柄)
}

// 按钮_置绑定元素.
//
// hBindEle: 将要绑定的元素.
func (b *Button) I置绑定元素(hBindEle int) int {
	return xc.XBtn_SetBindEle(b.I句柄, hBindEle)
}

// 按钮_取绑定元素, 返回: 绑定的元素句柄.
func (b *Button) I取绑定元素() int {
	return xc.XBtn_GetBindEle(b.I句柄)
}

// 按钮_置文本对齐.
//
// nFlags: 对齐方式, I常量_文本对齐_ , TextAlignFlag_ , TextTrimming_.
func (b *Button) I置文本对齐(nFlags xcc.I常量_文本对齐_) int {
	return xc.XBtn_SetTextAlign(b.I句柄, nFlags)
}

// 按钮_取文本对齐方式, 返回: I常量_文本对齐_ , TextAlignFlag_ , TextTrimming_.
func (b *Button) I取文本对齐方式() xcc.I常量_文本对齐_ {
	return xc.XBtn_GetTextAlign(b.I句柄)
}

// 按钮_置图标对齐.
//
// align: 对齐方式, I常量_按钮图标对齐_.
func (b *Button) I置图标对齐(align xcc.I常量_按钮图标对齐_) int {
	return xc.XBtn_SetIconAlign(b.I句柄, align)
}

// 按钮_置偏移, 设置按钮文本坐标偏移量.
//
// x: 偏移量.
//
// y: 偏移量.
func (b *Button) I置偏移(x int, y int) int {
	return xc.XBtn_SetOffset(b.I句柄, x, y)
}

// 按钮_置图标偏移, 设置按钮图标坐标偏移量.
//
// x: 偏移量.
//
// y: 偏移量.
func (b *Button) I置图标偏移(x int, y int) int {
	return xc.XBtn_SetOffsetIcon(b.I句柄, x, y)
}

// 按钮_置图标间隔, 设置图标与文本间隔大小.
//
// size: 间隔大小.
func (b *Button) I置图标间隔(size int) int {
	return xc.XBtn_SetIconSpace(b.I句柄, size)
}

// 按钮_置文本.
//
// pName: 文本内容.
func (b *Button) I置文本(pName string) int {
	return xc.XBtn_SetText(b.I句柄, pName)
}

// 按钮_取文本.
func (b *Button) I取文本() string {
	return xc.XBtn_GetText(b.I句柄)
}

// 按钮_置图标.
//
// hImage: 图片句柄.
func (b *Button) I置图标(hImage int) int {
	return xc.XBtn_SetIcon(b.I句柄, hImage)
}

// 按钮_置禁用图标.
//
// hImage: 图片句柄.
func (b *Button) I置禁用图标(hImage int) int {
	return xc.XBtn_SetIconDisable(b.I句柄, hImage)
}

// 按钮_取图标, 返回图标句柄.
//
// nType: 图标类型, 0:默认图标, 1:禁用状态图标.
func (b *Button) I取图标(nType int) int {
	return xc.XBtn_GetIcon(b.I句柄, nType)
}

// 按钮_添加动画帧.
//
// hImage: 图片句柄.
//
// uElapse: 图片帧延时, 单位毫秒.
func (b *Button) I添加动画帧(hImage int, uElapse int) int {
	return xc.XBtn_AddAnimationFrame(b.I句柄, hImage, uElapse)
}

// 按钮_启用动画, 开始或关闭图片动画的播放.
//
// bEnable: 开始播放动画TRUE, 关闭播放动画FALSE.
//
// bLoopPlay: 是否循环播放.
func (b *Button) I启用动画(bEnable bool, bLoopPlay bool) int {
	return xc.XBtn_EnableAnimation(b.I句柄, bEnable, bLoopPlay)
}

/*
下面都是事件
*/

type XE_BNCLICK func(pbHandled *bool) int                              // 按钮点击事件.
type XE_BNCLICK1 func(hEle int, pbHandled *bool) int                   // 按钮点击事件.
type XE_BUTTON_CHECK func(bCheck bool, pbHandled *bool) int            // 按钮选中事件.
type XE_BUTTON_CHECK1 func(hEle int, bCheck bool, pbHandled *bool) int // 按钮选中事件.

// 事件_按钮被单击.
func (b *Button) I事件_按钮被单击(pFun XE_BNCLICK) bool {
	return xc.XEle_RegEventC(b.I句柄, xcc.XE_BNCLICK, pFun)
}

// 事件_按钮被单击1.
func (b *Button) I事件_按钮被单击1(pFun XE_BNCLICK1) bool {
	return xc.XEle_RegEventC1(b.I句柄, xcc.XE_BNCLICK, pFun)
}

// 按钮选中事件.
func (b *Button) I事件_按钮选中(pFun XE_BUTTON_CHECK) bool {
	return xc.XEle_RegEventC(b.I句柄, xcc.XE_BUTTON_CHECK, pFun)
}

// 按钮选中事件.
func (b *Button) I事件_按钮选中1(pFun XE_BUTTON_CHECK1) bool {
	return xc.XEle_RegEventC1(b.I句柄, xcc.XE_BUTTON_CHECK, pFun)
}
