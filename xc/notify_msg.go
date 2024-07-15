package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// 通知消息_弹出, 未实现, 预留接口.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.

// ff:通知消息_弹出
// skin:
// hIcon:
// pText:
// pTitle:
// position:位置
func X通知消息_弹出(位置 炫彩常量类.Position_Flag_, pTitle, pText string, hIcon int, skin 炫彩常量类.NotifyMsg_Skin_) int {
	r, _, _ := xNotifyMsg_Popup.Call(uintptr(位置), 炫彩工具类.StrPtr(pTitle), 炫彩工具类.StrPtr(pText), uintptr(hIcon), uintptr(skin))
	return int(r)
}

// 通知消息_弹出扩展, 未实现, 预留接口.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
//
// bBtnClose: 是否启用关闭按钮.
//
// bAutoClose: 是否自动关闭.
//
// nWidth: 自定义宽度, -1(使用默认值).
//
// nHeight: 自定义高度, -1(使用默认值).

// ff:通知消息_弹出EX
// nHeight:
// nWidth:
// bAutoClose:
// bBtnClose:
// skin:
// hIcon:
// pText:
// pTitle:
// position:位置
func X通知消息_弹出EX(位置 炫彩常量类.Position_Flag_, pTitle, pText string, hIcon int, skin 炫彩常量类.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	r, _, _ := xNotifyMsg_PopupEx.Call(uintptr(位置), 炫彩工具类.StrPtr(pTitle), 炫彩工具类.StrPtr(pText), uintptr(hIcon), uintptr(skin), 炫彩工具类.BoolPtr(bBtnClose), 炫彩工具类.BoolPtr(bAutoClose), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 通知消息_窗口中弹出, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//
// hWindow: 窗口句柄.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.

// ff:通知消息_窗口中弹出
// skin:
// hIcon:
// pText:
// pTitle:
// position:位置
// hWindow:窗口句柄
func X通知消息_窗口中弹出(窗口句柄 int, 位置 炫彩常量类.Position_Flag_, pTitle, pText string, hIcon int, skin 炫彩常量类.NotifyMsg_Skin_) int {
	r, _, _ := xNotifyMsg_WindowPopup.Call(uintptr(窗口句柄), uintptr(位置), 炫彩工具类.StrPtr(pTitle), 炫彩工具类.StrPtr(pText), uintptr(hIcon), uintptr(skin))
	return int(r)
}

// 通知消息_窗口中弹出扩展, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//
// hWindow: 窗口句柄.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
//
// bBtnClose: 是否启用关闭按钮.
//
// bAutoClose: 是否自动关闭.
//
// nWidth: 自定义宽度, -1(使用默认值).
//
// nHeight: 自定义高度, -1(使用默认值).

// ff:通知消息_窗口中弹出EX
// nHeight:
// nWidth:
// bAutoClose:
// bBtnClose:
// skin:
// hIcon:
// pText:
// pTitle:
// position:位置
// hWindow:窗口句柄
func X通知消息_窗口中弹出EX(窗口句柄 int, 位置 炫彩常量类.Position_Flag_, pTitle, pText string, hIcon int, skin 炫彩常量类.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	r, _, _ := xNotifyMsg_WindowPopupEx.Call(uintptr(窗口句柄), uintptr(位置), 炫彩工具类.StrPtr(pTitle), 炫彩工具类.StrPtr(pText), uintptr(hIcon), uintptr(skin), 炫彩工具类.BoolPtr(bBtnClose), 炫彩工具类.BoolPtr(bAutoClose), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 通知消息_置持续时间.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// duration: 持续时间.

// ff:通知消息_置持续时间
// duration:持续时间
// hWindow:窗口句柄
func X通知消息_置持续时间(窗口句柄, 持续时间 int) int {
	r, _, _ := xNotifyMsg_SetDuration.Call(uintptr(窗口句柄), uintptr(持续时间))
	return int(r)
}

// 通知消息_置父边距 设置通知消息与父对象的四边间隔.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// left: 左侧间隔, 未实现, 预留功能.
//
// top: 顶部间隔.
//
// right: 右侧间隔.
//
// bottom: 底部间隔, 未实现, 预留功能.

// ff:通知消息_置父边距
// bottom:底部间隔
// right:右侧间隔
// top:顶部间隔
// left:左侧间隔
// hWindow:窗口句柄
func X通知消息_置父边距(窗口句柄, 左侧间隔, 顶部间隔, 右侧间隔, 底部间隔 int) int {
	r, _, _ := xNotifyMsg_SetParentMargin.Call(uintptr(窗口句柄), uintptr(左侧间隔), uintptr(顶部间隔), uintptr(右侧间隔), uintptr(底部间隔))
	return int(r)
}

// 通知消息_置标题高度.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nHeight: 高度.

// ff:通知消息_置标题高度
// nHeight:高度
// hWindow:窗口句柄
func X通知消息_置标题高度(窗口句柄, 高度 int) int {
	r, _, _ := xNotifyMsg_SetCaptionHeight.Call(uintptr(窗口句柄), uintptr(高度))
	return int(r)
}

// 通知消息_置宽度, 设置默认宽度.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nWidth: 宽度.

// ff:通知消息_置宽度
// nWidth:宽度
// hWindow:窗口句柄
func X通知消息_置宽度(窗口句柄, 宽度 int) int {
	r, _, _ := xNotifyMsg_SetWidth.Call(uintptr(窗口句柄), uintptr(宽度))
	return int(r)
}

// 通知消息_置间隔.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nSpace: 间隔大小.

// ff:通知消息_置间隔
// nSpace:间隔大小
// hWindow:窗口句柄
func X通知消息_置间隔(窗口句柄, 间隔大小 int) int {
	r, _, _ := xNotifyMsg_SetSpace.Call(uintptr(窗口句柄), uintptr(间隔大小))
	return int(r)
}

// 通知消息_置边大小, 设置通知消息面板边大小.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// left: 左边.
//
// top: 顶边.
//
// right: 右边.
//
// bottom: 底边.

// ff:通知消息_置边大小
// bottom:底边
// right:右边
// top:顶边
// left:左边
// hWindow:窗口句柄
func X通知消息_置边大小(窗口句柄, 左边, 顶边, 右边, 底边 int) int {
	r, _, _ := xNotifyMsg_SetBorderSize.Call(uintptr(窗口句柄), uintptr(左边), uintptr(顶边), uintptr(右边), uintptr(底边))
	return int(r)
}
