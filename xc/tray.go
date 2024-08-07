package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// 托盘图标_重置. 重置数据, 当未添加到系统托盘状态才可调用.

// ff:托盘图标_重置
func X托盘图标_重置() {
	xTrayIcon_Reset.Call()
}

// 托盘图标_删除. 从系统托盘删除.

// ff:托盘图标_删除
func X托盘图标_删除() bool {
	r, _, _ := xTrayIcon_Del.Call()
	return r != 0
}

// 托盘图标_修改. 修改托盘图标.

// ff:托盘图标_修改
func X托盘图标_修改() bool {
	r, _, _ := xTrayIcon_Modify.Call()
	return r != 0
}

// 托盘图标_置焦点. 将焦点设置回系统托盘.

// ff:托盘图标_置焦点
func X托盘图标_置焦点() bool {
	r, _, _ := xTrayIcon_SetFocus.Call()
	return r != 0
}

// 托盘图标_添加. 将图标添加到系统托盘.
//
// hWindow: 关联窗口句柄.
//
// id: 托盘图标唯一标识符.

// ff:托盘图标_添加
// id:托盘图标唯一标识符
// hWindow:关联窗口句柄
func X托盘图标_添加(关联窗口句柄 int, 托盘图标唯一标识符 int32) bool {
	r, _, _ := xTrayIcon_Add.Call(uintptr(关联窗口句柄), uintptr(托盘图标唯一标识符))
	return r != 0
}

// 托盘图标_置图标. 设置图标.
//
// hIcon: 图标句柄.

// ff:托盘图标_置图标
// hIcon:图标句柄
func X托盘图标_置图标(图标句柄 uintptr) {
	xTrayIcon_SetIcon.Call(图标句柄)
}

// 托盘图标_置提示文本. 设置工具提示内容.
//
// pTips: 提示文本内容, 长度不能超过127个字符.

// ff:托盘图标_置提示文本
// pTips:提示文本内容
func X托盘图标_置提示文本(提示文本内容 string) {
	xTrayIcon_SetTips.Call(炫彩工具类.StrPtr(提示文本内容))
}

// 托盘图标_置回调消息. 设置用户自定义的回调消息类型, 触发事件后, 系统会发送到此消息.
//
// user_message: 用户自定义消息, 界面库默认定义消息为: xcc.XWM_TRAYICON.

// ff:托盘图标_置回调消息
// user_message:用户自定义消息
func X托盘图标_置回调消息(用户自定义消息 uint32) {
	xTrayIcon_SetCallbackMessage.Call(uintptr(用户自定义消息))
}

// 托盘图标_置弹出气泡. 设置弹出气泡信息.
//
// pTitle: 弹出气泡标题.
//
// pText: 弹出气泡内容.
//
// hBalloonIcon: 气泡图标. 可填0.
//
// flags: 标识, 可设置默认图标类型, 禁用声音等: xcc.TrayIcon_Flag_

// ff:托盘图标_置弹出气泡
// flags:标识
// hBalloonIcon:气泡图标
// pText:弹出气泡内容
// pTitle:弹出气泡标题
func X托盘图标_置弹出气泡(弹出气泡标题, 弹出气泡内容 string, 气泡图标 uintptr, 标识 炫彩常量类.TrayIcon_Flag_) {
	xTrayIcon_SetPopupBalloon.Call(炫彩工具类.StrPtr(弹出气泡标题), 炫彩工具类.StrPtr(弹出气泡内容), 气泡图标, uintptr(标识))
}
