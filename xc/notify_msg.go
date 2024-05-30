package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// 通知消息_弹出, 未实现, 预留接口.
//, Position_Flag_.
//.
//.
//.
//, NotifyMsg_Skin_.
// ff:通知消息_弹出
// position:位置
// pTitle:
// pText:
// hIcon:
// skin:
func XNotifyMsg_Popup(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {
	r, _, _ := xNotifyMsg_Popup.Call(uintptr(position), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(hIcon), uintptr(skin))
	return int(r)
}

// 通知消息_弹出扩展, 未实现, 预留接口.
//, Position_Flag_.
//.
//.
//.
//, NotifyMsg_Skin_.
//.
//.
//, -1(使用默认值).
//, -1(使用默认值).
// ff:通知消息_弹出EX
// position:位置
// pTitle:
// pText:
// hIcon:
// skin:
// bBtnClose:
// bAutoClose:
// nWidth:
// nHeight:
func XNotifyMsg_PopupEx(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	r, _, _ := xNotifyMsg_PopupEx.Call(uintptr(position), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(hIcon), uintptr(skin), common.BoolPtr(bBtnClose), common.BoolPtr(bAutoClose), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 通知消息_窗口中弹出, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//.
//, Position_Flag_.
//.
//.
//.
//, NotifyMsg_Skin_.
// ff:通知消息_窗口中弹出
// hWindow:窗口句柄
// position:位置
// pTitle:
// pText:
// hIcon:
// skin:
func XNotifyMsg_WindowPopup(hWindow int, position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {
	r, _, _ := xNotifyMsg_WindowPopup.Call(uintptr(hWindow), uintptr(position), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(hIcon), uintptr(skin))
	return int(r)
}

// 通知消息_窗口中弹出扩展, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//.
//, Position_Flag_.
//.
//.
//.
//, NotifyMsg_Skin_.
//.
//.
//, -1(使用默认值).
//, -1(使用默认值).
// ff:通知消息_窗口中弹出EX
// hWindow:窗口句柄
// position:位置
// pTitle:
// pText:
// hIcon:
// skin:
// bBtnClose:
// bAutoClose:
// nWidth:
// nHeight:
func XNotifyMsg_WindowPopupEx(hWindow int, position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	r, _, _ := xNotifyMsg_WindowPopupEx.Call(uintptr(hWindow), uintptr(position), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(hIcon), uintptr(skin), common.BoolPtr(bBtnClose), common.BoolPtr(bAutoClose), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 通知消息_置持续时间.
//, 如果未指定那么认为是桌面通知消息.
//.
// ff:通知消息_置持续时间
// hWindow:窗口句柄
// duration:持续时间
func XNotifyMsg_SetDuration(hWindow, duration int) int {
	r, _, _ := xNotifyMsg_SetDuration.Call(uintptr(hWindow), uintptr(duration))
	return int(r)
}

// 通知消息_置父边距 设置通知消息与父对象的四边间隔.
//, 如果未指定那么认为是桌面通知消息.
//, 未实现, 预留功能.
//.
//.
//, 未实现, 预留功能.
// ff:通知消息_置父边距
// hWindow:窗口句柄
// left:左侧间隔
// top:顶部间隔
// right:右侧间隔
// bottom:底部间隔
func XNotifyMsg_SetParentMargin(hWindow, left, top, right, bottom int) int {
	r, _, _ := xNotifyMsg_SetParentMargin.Call(uintptr(hWindow), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 通知消息_置标题高度.
//, 如果未指定那么认为是桌面通知消息.
//.
// ff:通知消息_置标题高度
// hWindow:窗口句柄
// nHeight:高度
func XNotifyMsg_SetCaptionHeight(hWindow, nHeight int) int {
	r, _, _ := xNotifyMsg_SetCaptionHeight.Call(uintptr(hWindow), uintptr(nHeight))
	return int(r)
}

// 通知消息_置宽度, 设置默认宽度.
//, 如果未指定那么认为是桌面通知消息.
//.
// ff:通知消息_置宽度
// hWindow:窗口句柄
// nWidth:宽度
func XNotifyMsg_SetWidth(hWindow, nWidth int) int {
	r, _, _ := xNotifyMsg_SetWidth.Call(uintptr(hWindow), uintptr(nWidth))
	return int(r)
}

// 通知消息_置间隔.
//, 如果未指定那么认为是桌面通知消息.
//.
// ff:通知消息_置间隔
// hWindow:窗口句柄
// nSpace:间隔大小
func XNotifyMsg_SetSpace(hWindow, nSpace int) int {
	r, _, _ := xNotifyMsg_SetSpace.Call(uintptr(hWindow), uintptr(nSpace))
	return int(r)
}

// 通知消息_置边大小, 设置通知消息面板边大小.
//, 如果未指定那么认为是桌面通知消息.
//.
//.
//.
//.
// ff:通知消息_置边大小
// hWindow:窗口句柄
// left:左边
// top:顶边
// right:右边
// bottom:底边
func XNotifyMsg_SetBorderSize(hWindow, left, top, right, bottom int) int {
	r, _, _ := xNotifyMsg_SetBorderSize.Call(uintptr(hWindow), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}
