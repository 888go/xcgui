package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// 模态窗口_创建, 创建模态窗口; 当模态窗口关闭时, 会自动销毁模态窗口资源句柄.
// nWidth:宽度.
// nHeight:高度.
// pTitle:标题内容.
// hWndParent:父窗口句柄.
// XCStyle:炫彩窗口样式:Window_Style_.
func XModalWnd_Create(nWidth, nHeight int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ : xModalWnd_Create.Call(uintptr(nWidth), uintptr(nHeight), common.StrPtr(pTitle), hWndParent, uintptr(XCStyle))
	return int(r)
}

// 模态窗口_创建EX, 创建模态窗口, 增强功能.
// dwExStyle:EX样式.
// dwStyle:样式.
// lpClassName:类名.
// x:左上角x坐标.
// y:左上角y坐标.
// cx:宽度.
// cy:高度.
// pTitle:窗口名.
// hWndParent:父窗口.
// XCStyle:GUI库窗口样式:Window_Style_.
func XModalWnd_CreateEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ : xModalWnd_CreateEx.Call(uintptr(dwExStyle), uintptr(dwStyle), common.StrPtr(lpClassName), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), common.StrPtr(pTitle), hWndParent, uintptr(XCStyle))
	return int(r)
}

// 模态窗口_启用自动关闭, 是否自动关闭窗口, 当窗口失去焦点时.
// hWindow:模态窗口句柄.
// bEnable:是否启用     .开启开关.
func XModalWnd_EnableAutoClose(hWindow int, bEnable bool) int {
	r, _, _ : xModalWnd_EnableAutoClose.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 模态窗口_启用ESC关闭, 当用户按ESC键时自动关闭模态窗口.
// hWindow:模态窗口句柄.
// bEnable:是否启用     .是否启用.
func XModalWnd_EnableEscClose(hWindow int, bEnable bool) int {
	r, _, _ : xModalWnd_EnableEscClose.Call(uintptr(hWindow), common.BoolPtr(bEnable))
	return int(r)
}

// 模态窗口_启动, 启动显示模态窗口, 当窗口关闭时返回:MessageBox_Flag_Ok:点击确定按钮退出, MessageBox_Flag_Cancel:点击取消按钮退出, MessageBox_Flag_Other:其他方式退出.
// hWindow:模态窗口句柄.
func XModalWnd_DoModal(hWindow int) xcc.MessageBox_Flag_ {
	r, _, _ : xModalWnd_DoModal.Call(uintptr(hWindow))
	return xcc.MessageBox_Flag_(r)
}

// 模态窗口_结束, 结束模态窗口.
// hWindow:窗口句柄.
// nResult:用作XModalWnd_DoModal()的返回值. MessageBox_Flag_Ok:点击确定按钮退出, MessageBox_Flag_Cancel:点击取消按钮退出, MessageBox_Flag_Other:其他方式退出.
func XModalWnd_EndModal(hWindow int, nResult xcc.MessageBox_Flag_) int {
	r, _, _ : xModalWnd_EndModal.Call(uintptr(hWindow), uintptr(nResult))
	return int(r)
}

// 模态窗口_附加窗口, 返回窗口资源句柄.
// hWnd:外部窗口句柄
// XCStyle:炫彩窗口样式:Window_Style_.
func XModalWnd_Attach(hWnd uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ : xModalWnd_Attach.Call(hWnd, uintptr(XCStyle))
	return int(r)
}
