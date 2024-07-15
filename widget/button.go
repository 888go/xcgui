package 炫彩组件类 //bm:炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
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

// ff:创建按钮
// hParent:父窗口句柄
// pName:标题
// cy:高度
// cx:宽度
// y:按钮y坐标
// x:按钮x坐标
func X创建按钮(按钮x坐标, 按钮y坐标, 宽度, 高度 int32, 标题 string, 父窗口句柄 int) *Button {
	p := &Button{}
	p.X设置句柄(炫彩基类.X按钮_创建(按钮x坐标, 按钮y坐标, 宽度, 高度, 标题, 父窗口句柄))
	return p
}

// 从句柄创建对象.

// ff:创建按钮并按句柄
// handle:句柄
func X创建按钮并按句柄(句柄 int) *Button {
	p := &Button{}
	p.X设置句柄(句柄)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建按钮并按名称
// name:名称
func X创建按钮并按名称(名称 string) *Button {
	handle := 炫彩基类.X取对象从名称(名称)
	if handle > 0 {
		p := &Button{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建按钮并按UID
// nUID:
func X创建按钮并按UID(nUID int) *Button {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Button{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建按钮并按UID名称
// name:UID名称
func X创建按钮并按UID名称(UID名称 string) *Button {
	handle := 炫彩基类.X取对象从UID名称(UID名称)
	if handle > 0 {
		p := &Button{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 按钮_判断选中, 是否选中状态.

// ff:判断选中
func (b *Button) X判断选中() bool {
	return 炫彩基类.X按钮_判断选中(b.Handle)
}

// 按钮_置选中, 设置选中状态.
//
// bCheck: 是否选中.

// ff:置选中
// bCheck:是否选中
func (b *Button) X置选中(是否选中 bool) bool {
	return 炫彩基类.X按钮_置选中(b.Handle, 是否选中)
}

// SetState 按钮_置状态.
//
//	@param nState 按钮状态: xcc.Common_State3_.
//	@return int

// ff:置状态
// nState:按钮状态
func (b *Button) X置状态(按钮状态 炫彩常量类.Common_State3_) int {
	return 炫彩基类.X按钮_置状态(b.Handle, 按钮状态)
}

// GetState 按钮_取状态.
//
//	@return xcc.Common_State3_

// ff:取状态
func (b *Button) X取状态() 炫彩常量类.Common_State3_ {
	return 炫彩基类.X按钮_取状态(b.Handle)
}

// GetStateEx 按钮_取状态扩展.
//
//	@return xcc.Button_State_

// ff:取状态EX
func (b *Button) X取状态EX() 炫彩常量类.Button_State_ {
	return 炫彩基类.X按钮_取状态EX(b.Handle)
}

// 按钮_置类型扩展, 设置按钮类型并自动修改样式和文本对齐方式.
//
// nType: 按钮类型, Button_Type_ , element_type_ , xc_ex_error.

// ff:置类型EX
// nType:按钮类型
func (b *Button) X置类型EX(按钮类型 炫彩常量类.XC_OBJECT_TYPE_EX) int {
	return 炫彩基类.X按钮_置类型EX(b.Handle, 按钮类型)
}

// 按钮_置组ID.
//
// nID: 组ID.

// ff:置组ID
// nID:组ID
func (b *Button) X置组ID(组ID int32) int {
	return 炫彩基类.X按钮_置组ID(b.Handle, 组ID)
}

// 按钮_取组ID.

// ff:取组ID
func (b *Button) X取组ID() int32 {
	return 炫彩基类.X按钮_取组ID(b.Handle)
}

// 按钮_置绑定元素.
//
// hBindEle: 将要绑定的元素.

// ff:置绑定元素
// hBindEle:将要绑定的元素
func (b *Button) X置绑定元素(将要绑定的元素 int) int {
	return 炫彩基类.X按钮_置绑定元素(b.Handle, 将要绑定的元素)
}

// 按钮_取绑定元素, 返回: 绑定的元素句柄.

// ff:取绑定元素
func (b *Button) X取绑定元素() int {
	return 炫彩基类.X按钮_取绑定元素(b.Handle)
}

// 按钮_置文本对齐.
//
// nFlags: 对齐方式, TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.

// ff:置文本对齐
// nFlags:对齐方式
func (b *Button) X置文本对齐(对齐方式 炫彩常量类.TextFormatFlag_) int {
	return 炫彩基类.X按钮_置文本对齐(b.Handle, 对齐方式)
}

// 按钮_取文本对齐方式, 返回: TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.

// ff:取文本对齐方式
func (b *Button) X取文本对齐方式() 炫彩常量类.TextFormatFlag_ {
	return 炫彩基类.X按钮_取文本对齐方式(b.Handle)
}

// 按钮_置图标对齐.
//
// align: 对齐方式, Button_Icon_Align_.

// ff:置图标对齐
// align:对齐方式
func (b *Button) X置图标对齐(对齐方式 炫彩常量类.Button_Icon_Align_) int {
	return 炫彩基类.X按钮_置图标对齐(b.Handle, 对齐方式)
}

// 按钮_置偏移, 设置按钮文本坐标偏移量.
//
// x: 偏移量.
//
// y: 偏移量.

// ff:置偏移
// y:偏移量y
// x:偏移量x
func (b *Button) X置偏移(偏移量x int, 偏移量y int) int {
	return 炫彩基类.X按钮_置偏移(b.Handle, 偏移量x, 偏移量y)
}

// 按钮_置图标偏移, 设置按钮图标坐标偏移量.
//
// x: 偏移量.
//
// y: 偏移量.

// ff:置图标偏移
// y:偏移量y
// x:偏移量x
func (b *Button) X置图标偏移(偏移量x int, 偏移量y int) int {
	return 炫彩基类.X按钮_置图标偏移(b.Handle, 偏移量x, 偏移量y)
}

// 按钮_置图标间隔, 设置图标与文本间隔大小.
//
// size: 间隔大小.

// ff:置图标间隔
// size:间隔大小
func (b *Button) X置图标间隔(间隔大小 int) int {
	return 炫彩基类.X按钮_置图标间隔(b.Handle, 间隔大小)
}

// 按钮_置文本.
//
// pName: 文本内容.

// ff:置文本
// pName:文本内容
func (b *Button) X置文本(文本内容 string) int {
	return 炫彩基类.X按钮_置文本(b.Handle, 文本内容)
}

// 按钮_取文本.

// ff:取文本
func (b *Button) X取文本() string {
	return 炫彩基类.X按钮_取文本(b.Handle)
}

// 按钮_置图标.
//
// hImage: 图片句柄.

// ff:置图标
// hImage:图片句柄
func (b *Button) X置图标(图片句柄 int) int {
	return 炫彩基类.X按钮_置图标(b.Handle, 图片句柄)
}

// 按钮_置禁用图标.
//
// hImage: 图片句柄.

// ff:置禁用图标
// hImage:图片句柄
func (b *Button) X置禁用图标(图片句柄 int) int {
	return 炫彩基类.X按钮_置禁用图标(b.Handle, 图片句柄)
}

// 按钮_取图标, 返回图标句柄.
//
// nType: 图标类型, 0:默认图标, 1:禁用状态图标.

// ff:取图标
// nType:图标类型
func (b *Button) X取图标(图标类型 int) int {
	return 炫彩基类.X按钮_取图标(b.Handle, 图标类型)
}

// 按钮_添加动画帧.
//
// hImage: 图片句柄.
//
// uElapse: 图片帧延时, 单位毫秒.

// ff:添加动画帧
// uElapse:图片帧延时
// hImage:图片句柄
func (b *Button) X添加动画帧(图片句柄 int, 图片帧延时 int) int {
	return 炫彩基类.X按钮_添加动画帧(b.Handle, 图片句柄, 图片帧延时)
}

// 按钮_启用动画, 开始或关闭图片动画的播放.
//
// bEnable: 开始播放动画TRUE, 关闭播放动画FALSE.
//
// bLoopPlay: 是否循环播放.

// ff:启用动画
// bLoopPlay:是否循环播放
// bEnable:启用
func (b *Button) X启用动画(启用 bool, 是否循环播放 bool) int {
	return 炫彩基类.X按钮_启用动画(b.Handle, 启用, 是否循环播放)
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
func (b *Button) X事件_被单击(pFun XE_BNCLICK) bool {
	return 炫彩基类.X元素_注册事件C(b.Handle, 炫彩常量类.XE_BNCLICK, pFun)
}

// 事件_按钮被单击1.

// ff:事件_被单击1
// pFun:
func (b *Button) X事件_被单击1(pFun XE_BNCLICK1) bool {
	return 炫彩基类.X元素_注册事件C1(b.Handle, 炫彩常量类.XE_BNCLICK, pFun)
}

// 按钮选中事件.

// ff:事件_选中
// pFun:
func (b *Button) X事件_选中(pFun XE_BUTTON_CHECK) bool {
	return 炫彩基类.X元素_注册事件C(b.Handle, 炫彩常量类.XE_BUTTON_CHECK, pFun)
}

// 按钮选中事件.

// ff:事件_选中1
// pFun:
func (b *Button) X事件_选中1(pFun XE_BUTTON_CHECK1) bool {
	return 炫彩基类.X元素_注册事件C1(b.Handle, 炫彩常量类.XE_BUTTON_CHECK, pFun)
}
