package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// 模态窗口_创建, 创建模态窗口; 当模态窗口关闭时, 会自动销毁模态窗口资源句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// pTitle: 窗口标题内容.
//
// hWndParent: 父窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.
func X模态窗口_创建(宽度, 高度 int32, 标题内容 string, 父窗口句柄 uintptr, 炫彩窗口样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xModalWnd_Create.Call(uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(标题内容), 父窗口句柄, uintptr(炫彩窗口样式))
	return int(r)
}

// 模态窗口_创建扩展, 创建模态窗口, 增强功能.
//
// dwExStyle: 窗口扩展样式.
//
// dwStyle: 窗口样式.
//
// lpClassName: 窗口类名.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口名.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式: Window_Style_.
func X模态窗口_创建EX(扩展样式 int, 样式 int, 类名 string, 左上角x坐标, 左上角y坐标, 宽度, 高度 int32, 窗口名 string, 父窗口 uintptr, GUI库窗口样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xModalWnd_CreateEx.Call(uintptr(扩展样式), uintptr(样式), 炫彩工具类.StrPtr(类名), uintptr(左上角x坐标), uintptr(左上角y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(窗口名), 父窗口, uintptr(GUI库窗口样式))
	return int(r)
}

// 模态窗口_启用自动关闭, 是否自动关闭窗口, 当窗口失去焦点时.
//
// hWindow: 模态窗口句柄.
//
// bEnable: 开启开关.
func X模态窗口_启用自动关闭(模态窗口句柄 int, 开启开关 bool) int {
	r, _, _ := xModalWnd_EnableAutoClose.Call(uintptr(模态窗口句柄), 炫彩工具类.BoolPtr(开启开关))
	return int(r)
}

// 模态窗口_启用ESC关闭, 当用户按ESC键时自动关闭模态窗口.
//
// hWindow: 模态窗口句柄.
//
// bEnable: 是否启用.
func X模态窗口_启用ESC关闭(模态窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xModalWnd_EnableEscClose.Call(uintptr(模态窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 模态窗口_启动, 启动显示模态窗口, 当窗口关闭时返回: MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出.
//
// hWindow: 模态窗口句柄.
func X模态窗口_启动(模态窗口句柄 int) 炫彩常量类.MessageBox_Flag_ {
	r, _, _ := xModalWnd_DoModal.Call(uintptr(模态窗口句柄))
	return 炫彩常量类.MessageBox_Flag_(r)
}

// 模态窗口_结束, 结束模态窗口.
//
// hWindow: 窗口句柄.
//
// nResult: 用作XModalWnd_DoModal()的返回值. MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出.
func X模态窗口_结束(窗口句柄 int, 用作XModalWnd_DoModal 炫彩常量类.MessageBox_Flag_) int {
	r, _, _ := xModalWnd_EndModal.Call(uintptr(窗口句柄), uintptr(用作XModalWnd_DoModal))
	return int(r)
}

// 模态窗口_附加窗口, 返回窗口资源句柄.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.
func X模态窗口_附加窗口(外部窗口句柄 uintptr, 炫彩窗口样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xModalWnd_Attach.Call(外部窗口句柄, uintptr(炫彩窗口样式))
	return int(r)
}
