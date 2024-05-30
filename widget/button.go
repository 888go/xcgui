package widget//bm:炫彩组件类

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Button 按钮.
type Button struct {
	Element
}

// 按钮_创建.
//.
//.
//.
//.
//.
//.
// ff:创建按钮
// x:按钮x坐标
// y:按钮y坐标
// cx:宽度
// cy:高度
// pName:标题
// hParent:父窗口句柄
func NewButton(x, y, cx, cy int32, pName string, hParent int) *Button {
	p := &Button{}
	p.SetHandle(xc.XBtn_Create(x, y, cx, cy, pName, hParent))
	return p
}

// 从句柄创建对象.
// ff:创建按钮并按句柄
// handle:句柄
func NewButtonByHandle(handle int) *Button {
	p := &Button{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
// ff:创建按钮并按名称
// name:名称
func NewButtonByName(name string) *Button {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
// ff:创建按钮并按UID
// nUID:
func NewButtonByUID(nUID int) *Button {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
// ff:创建按钮并按UID名称
// name:UID名称
func NewButtonByUIDName(name string) *Button {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 按钮_判断选中, 是否选中状态.
// ff:判断选中
func (b *Button) IsCheck() bool {
	return xc.XBtn_IsCheck(b.Handle)
}

// 按钮_置选中, 设置选中状态.
//.
// ff:置选中
// bCheck:是否选中
func (b *Button) SetCheck(bCheck bool) bool {
	return xc.XBtn_SetCheck(b.Handle, bCheck)
}

// SetState 按钮_置状态.
//
//	@param nState 按钮状态: xcc.Common_State3_.
//	@return int
// ff:置状态
// nState:按钮状态
func (b *Button) SetState(nState xcc.Common_State3_) int {
	return xc.XBtn_SetState(b.Handle, nState)
}

// GetState 按钮_取状态.
//
//	@return xcc.Common_State3_
// ff:取状态
func (b *Button) GetState() xcc.Common_State3_ {
	return xc.XBtn_GetState(b.Handle)
}

// GetStateEx 按钮_取状态扩展.
//
//	@return xcc.Button_State_
// ff:取状态EX
func (b *Button) GetStateEx() xcc.Button_State_ {
	return xc.XBtn_GetStateEx(b.Handle)
}

// 按钮_置类型扩展, 设置按钮类型并自动修改样式和文本对齐方式.
//, Button_Type_ , element_type_ , xc_ex_error.
// ff:置类型EX
// nType:按钮类型
func (b *Button) SetTypeEx(nType xcc.XC_OBJECT_TYPE_EX) int {
	return xc.XBtn_SetTypeEx(b.Handle, nType)
}

// 按钮_置组ID.
//.
// ff:置组ID
// nID:组ID
func (b *Button) SetGroupID(nID int32) int {
	return xc.XBtn_SetGroupID(b.Handle, nID)
}

// 按钮_取组ID.
// ff:取组ID
func (b *Button) GetGroupID() int32 {
	return xc.XBtn_GetGroupID(b.Handle)
}

// 按钮_置绑定元素.
//.
// ff:置绑定元素
// hBindEle:将要绑定的元素
func (b *Button) SetBindEle(hBindEle int) int {
	return xc.XBtn_SetBindEle(b.Handle, hBindEle)
}

// 按钮_取绑定元素, 返回: 绑定的元素句柄.
// ff:取绑定元素
func (b *Button) GetBindEle() int {
	return xc.XBtn_GetBindEle(b.Handle)
}

// 按钮_置文本对齐.
//, TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
// ff:置文本对齐
// nFlags:对齐方式
func (b *Button) SetTextAlign(nFlags xcc.TextFormatFlag_) int {
	return xc.XBtn_SetTextAlign(b.Handle, nFlags)
}

// 按钮_取文本对齐方式, 返回: TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
// ff:取文本对齐方式
func (b *Button) GetTextAlign() xcc.TextFormatFlag_ {
	return xc.XBtn_GetTextAlign(b.Handle)
}

// 按钮_置图标对齐.
//, Button_Icon_Align_.
// ff:置图标对齐
// align:对齐方式
func (b *Button) SetIconAlign(align xcc.Button_Icon_Align_) int {
	return xc.XBtn_SetIconAlign(b.Handle, align)
}

// 按钮_置偏移, 设置按钮文本坐标偏移量.
//.
//.
// ff:置偏移
// x:偏移量x
// y:偏移量y
func (b *Button) SetOffset(x int, y int) int {
	return xc.XBtn_SetOffset(b.Handle, x, y)
}

// 按钮_置图标偏移, 设置按钮图标坐标偏移量.
//.
//.
// ff:置图标偏移
// x:偏移量x
// y:偏移量y
func (b *Button) SetOffsetIcon(x int, y int) int {
	return xc.XBtn_SetOffsetIcon(b.Handle, x, y)
}

// 按钮_置图标间隔, 设置图标与文本间隔大小.
//.
// ff:置图标间隔
// size:间隔大小
func (b *Button) SetIconSpace(size int) int {
	return xc.XBtn_SetIconSpace(b.Handle, size)
}

// 按钮_置文本.
//.
// ff:置文本
// pName:文本内容
func (b *Button) SetText(pName string) int {
	return xc.XBtn_SetText(b.Handle, pName)
}

// 按钮_取文本.
// ff:取文本
func (b *Button) GetText() string {
	return xc.XBtn_GetText(b.Handle)
}

// 按钮_置图标.
//.
// ff:置图标
// hImage:图片句柄
func (b *Button) SetIcon(hImage int) int {
	return xc.XBtn_SetIcon(b.Handle, hImage)
}

// 按钮_置禁用图标.
//.
// ff:置禁用图标
// hImage:图片句柄
func (b *Button) SetIconDisable(hImage int) int {
	return xc.XBtn_SetIconDisable(b.Handle, hImage)
}

// 按钮_取图标, 返回图标句柄.
//, 0:默认图标, 1:禁用状态图标.
// ff:取图标
// nType:图标类型
func (b *Button) GetIcon(nType int) int {
	return xc.XBtn_GetIcon(b.Handle, nType)
}

// 按钮_添加动画帧.
//.
//, 单位毫秒.
// ff:添加动画帧
// hImage:图片句柄
// uElapse:图片帧延时
func (b *Button) AddAnimationFrame(hImage int, uElapse int) int {
	return xc.XBtn_AddAnimationFrame(b.Handle, hImage, uElapse)
}

// 按钮_启用动画, 开始或关闭图片动画的播放.
//, 关闭播放动画FALSE.
//.
// ff:启用动画
// bEnable:启用
// bLoopPlay:是否循环播放
func (b *Button) EnableAnimation(bEnable bool, bLoopPlay bool) int {
	return xc.XBtn_EnableAnimation(b.Handle, bEnable, bLoopPlay)
}

/*
下面都是事件
*/

type XE_BNCLICK func(pbHandled *bool) int                              // 按钮点击事件.
type XE_BNCLICK1 func(hEle int, pbHandled *bool) int                   // 按钮点击事件.
type XE_BUTTON_CHECK func(bCheck bool, pbHandled *bool) int            // 按钮选中事件.
type XE_BUTTON_CHECK1 func(hEle int, bCheck bool, pbHandled *bool) int // 按钮选中事件.

// 事件_按钮被单击.
// ff:事件_被单击
// pFun:
func (b *Button) Event_BnClick(pFun XE_BNCLICK) bool {
	return xc.XEle_RegEventC(b.Handle, xcc.XE_BNCLICK, pFun)
}

// 事件_按钮被单击1.
// ff:事件_被单击1
// pFun:
func (b *Button) Event_BnClick1(pFun XE_BNCLICK1) bool {
	return xc.XEle_RegEventC1(b.Handle, xcc.XE_BNCLICK, pFun)
}

// 按钮选中事件.
// ff:事件_选中
// pFun:
func (b *Button) Event_BUTTON_CHECK(pFun XE_BUTTON_CHECK) bool {
	return xc.XEle_RegEventC(b.Handle, xcc.XE_BUTTON_CHECK, pFun)
}

// 按钮选中事件.
// ff:事件_选中1
// pFun:
func (b *Button) Event_BUTTON_CHECK1(pFun XE_BUTTON_CHECK1) bool {
	return xc.XEle_RegEventC1(b.Handle, xcc.XE_BUTTON_CHECK, pFun)
}
