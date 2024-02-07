package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

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
func X按钮_创建(x坐标, y坐标, 宽度, 高度 int32, 标题 string, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xBtn_Create.Call(uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(标题), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 按钮_判断选中, 是否选中状态.
//
// hEle: 元素句柄.
func X按钮_判断选中(元素句柄 int) bool {
	r, _, _ := xBtn_IsCheck.Call(uintptr(元素句柄))
	return r != 0
}

// 按钮_置选中, 设置选中状态.
//
// hEle: 元素句柄.
//
// bCheck: 是否选中.
func X按钮_置选中(元素句柄 int, 是否选中 bool) bool {
	r, _, _ := xBtn_SetCheck.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否选中))
	return r != 0
}

// XBtn_SetState 按钮_置状态.
//
//	@param hEle 元素句柄.
//	@param nState 按钮状态: xcc.Common_State3_
//	@return int
func X按钮_置状态(元素句柄 int, 按钮状态 炫彩常量类.Common_State3_) int {
	r, _, _ := xBtn_SetState.Call(uintptr(元素句柄), uintptr(按钮状态))
	return int(r)
}

// XBtn_GetState 按钮_取状态.
//
//	@param hEle 元素句柄.
//	@return xcc.Common_State3_
func X按钮_取状态(元素句柄 int) 炫彩常量类.Common_State3_ {
	r, _, _ := xBtn_GetState.Call(uintptr(元素句柄))
	return 炫彩常量类.Common_State3_(r)
}

// XBtn_GetStateEx 按钮_取状态扩展.
//
//	@param hEle 元素句柄.
//	@return xcc.Button_State_
func X按钮_取状态EX(元素句柄 int) 炫彩常量类.Button_State_ {
	r, _, _ := xBtn_GetStateEx.Call(uintptr(元素句柄))
	return 炫彩常量类.Button_State_(r)
}

// 按钮_置类型扩展, 设置按钮类型并自动修改样式和文本对齐方式.
//
// hEle: 元素句柄.
//
// nType: 按钮类型, Button_Type_ , element_type_ , xc_ex_error.
func X按钮_置类型EX(元素句柄 int, 按钮类型 炫彩常量类.XC_OBJECT_TYPE_EX) int {
	r, _, _ := xBtn_SetTypeEx.Call(uintptr(元素句柄), uintptr(按钮类型))
	return int(r)
}

// 按钮_置组ID.
//
// hEle: 元素句柄.
//
// nID: 组ID.
func X按钮_置组ID(元素句柄 int, 组ID int32) int {
	r, _, _ := xBtn_SetGroupID.Call(uintptr(元素句柄), uintptr(组ID))
	return int(r)
}

// 按钮_取组ID.
//
// hEle: 元素句柄.
func X按钮_取组ID(元素句柄 int) int32 {
	r, _, _ := xBtn_GetGroupID.Call(uintptr(元素句柄))
	return int32(r)
}

// 按钮_置绑定元素.
//
// hEle: 元素句柄.
//
// hBindEle: 将要绑定的元素.
func X按钮_置绑定元素(元素句柄 int, 将要绑定的元素 int) int {
	r, _, _ := xBtn_SetBindEle.Call(uintptr(元素句柄), uintptr(将要绑定的元素))
	return int(r)
}

// 按钮_取绑定元素, 返回: 绑定的元素句柄.
//
// hEle: 元素句柄.
func X按钮_取绑定元素(元素句柄 int) int {
	r, _, _ := xBtn_GetBindEle.Call(uintptr(元素句柄))
	return int(r)
}

// 按钮_置文本对齐.
//
// hEle: 元素句柄.
//
// nFlags: 对齐方式, TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
func X按钮_置文本对齐(元素句柄 int, 对齐方式 炫彩常量类.TextFormatFlag_) int {
	r, _, _ := xBtn_SetTextAlign.Call(uintptr(元素句柄), uintptr(对齐方式))
	return int(r)
}

// 按钮_取文本对齐方式, 返回: TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
//
// hEle: 元素句柄.
func X按钮_取文本对齐方式(元素句柄 int) 炫彩常量类.TextFormatFlag_ {
	r, _, _ := xBtn_GetTextAlign.Call(uintptr(元素句柄))
	return 炫彩常量类.TextFormatFlag_(r)
}

// 按钮_置图标对齐.
//
// hEle: 元素句柄.
//
// align: 对齐方式, Button_Icon_Align_.
func X按钮_置图标对齐(元素句柄 int, 对齐方式 炫彩常量类.Button_Icon_Align_) int {
	r, _, _ := xBtn_SetIconAlign.Call(uintptr(元素句柄), uintptr(对齐方式))
	return int(r)
}

// 按钮_置偏移, 设置按钮文本坐标偏移量.
//
// hEle: 元素句柄.
//
// x: 偏移量.
//
// y: 偏移量.
func X按钮_置偏移(元素句柄 int, 偏移量x int, 偏移量y int) int {
	r, _, _ := xBtn_SetOffset.Call(uintptr(元素句柄), uintptr(偏移量x), uintptr(偏移量y))
	return int(r)
}

// 按钮_置图标偏移, 设置按钮图标坐标偏移量.
//
// hEle: 元素句柄.
//
// x: 偏移量.
//
// y: 偏移量.
func X按钮_置图标偏移(元素句柄 int, 偏移量x int, 偏移量y int) int {
	r, _, _ := xBtn_SetOffsetIcon.Call(uintptr(元素句柄), uintptr(偏移量x), uintptr(偏移量y))
	return int(r)
}

// 按钮_置图标间隔, 设置图标与文本间隔大小.
//
// hEle: 元素句柄.
//
// size: 间隔大小.
func X按钮_置图标间隔(元素句柄 int, 间隔大小 int) int {
	r, _, _ := xBtn_SetIconSpace.Call(uintptr(元素句柄), uintptr(间隔大小))
	return int(r)
}

// 按钮_置文本.
//
// hEle: 元素句柄.
//
// pName: 文本内容.
func X按钮_置文本(元素句柄 int, 文本内容 string) int {
	r, _, _ := xBtn_SetText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文本内容))
	return int(r)
}

// 按钮_取文本.
//
// hEle: 元素句柄.
func X按钮_取文本(元素句柄 int) string {
	r, _, _ := xBtn_GetText.Call(uintptr(元素句柄))
	return 炫彩工具类.UintPtrToString(r)
}

// 按钮_置图标.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func X按钮_置图标(元素句柄 int, 图片句柄 int) int {
	r, _, _ := xBtn_SetIcon.Call(uintptr(元素句柄), uintptr(图片句柄))
	return int(r)
}

// 按钮_置禁用图标.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func X按钮_置禁用图标(元素句柄 int, 图片句柄 int) int {
	r, _, _ := xBtn_SetIconDisable.Call(uintptr(元素句柄), uintptr(图片句柄))
	return int(r)
}

// 按钮_取图标, 返回图标句柄.
//
// hEle: 元素句柄.
//
// nType: 图标类型, 0:默认图标, 1:禁用状态图标.
func X按钮_取图标(元素句柄 int, 图标类型 int) int {
	r, _, _ := xBtn_GetIcon.Call(uintptr(元素句柄), uintptr(图标类型))
	return int(r)
}

// 按钮_添加动画帧.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
//
// uElapse: 图片帧延时, 单位毫秒.
func X按钮_添加动画帧(元素句柄 int, 图片句柄 int, 图片帧延时 int) int {
	r, _, _ := xBtn_AddAnimationFrame.Call(uintptr(元素句柄), uintptr(图片句柄), uintptr(图片帧延时))
	return int(r)
}

// 按钮_启用动画, 开始或关闭图片动画的播放.
//
// hEle: 元素句柄.
//
// bEnable: 开始播放动画TRUE, 关闭播放动画FALSE.
//
// bLoopPlay: 是否循环播放.
func X按钮_启用动画(元素句柄 int, 开始播放动画TRUE bool, 是否循环播放 bool) int {
	r, _, _ := xBtn_EnableAnimation.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(开始播放动画TRUE), 炫彩工具类.BoolPtr(是否循环播放))
	return int(r)
}
