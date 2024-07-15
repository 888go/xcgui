package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"syscall"
	"unsafe"

	"github.com/888go/xcgui/xcc"
)

// 窗口_创建, 返回: GUI库窗口资源句柄.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口标题.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式, Window_Style_.

// ff:窗口_创建
// XCStyle:GUI库窗口样式
// hWndParent:父窗口
// pTitle:窗口标题
// cy:高度
// cx:宽度
// y:左上角y坐标
// x:左上角x坐标
func X窗口_创建(左上角x坐标, 左上角y坐标, 宽度, 高度 int32, 窗口标题 string, 父窗口 uintptr, GUI库窗口样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xWnd_Create.Call(uintptr(左上角x坐标), uintptr(左上角y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(窗口标题), 父窗口, uintptr(GUI库窗口样式))
	return int(r)
}

// 窗口_创建扩展, 返回: GUI库窗口资源句柄.
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
// XCStyle: GUI库窗口样式, Window_Style_.

// ff:窗口_创建EX
// XCStyle:GUI库窗口样式
// hWndParent:父窗口
// pTitle:窗口名
// cy:高度
// cx:宽度
// y:左上角y坐标
// x:左上角x坐标
// lpClassName:类名
// dwStyle:样式
// dwExStyle:扩展样式
func X窗口_创建EX(扩展样式 int, 样式 int, 类名 string, 左上角x坐标 int, 左上角y坐标 int, 宽度 int, 高度 int, 窗口名 string, 父窗口 uintptr, GUI库窗口样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xWnd_CreateEx.Call(uintptr(扩展样式), uintptr(样式), 炫彩工具类.StrPtr(类名), uintptr(左上角x坐标), uintptr(左上角y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(窗口名), 父窗口, uintptr(GUI库窗口样式))
	return int(r)
}

// 窗口_显示.
//
// hWindow: 窗口句柄.
//
// nCmdShow: 显示方式: xcc.SW_.

// ff:窗口_显示方式
// nCmdShow:显示方式
// hWindow:窗口句柄
func X窗口_显示方式(窗口句柄 int, 显示方式 炫彩常量类.SW_) int {
	r, _, _ := xWnd_ShowWindow.Call(uintptr(窗口句柄), uintptr(显示方式))
	return int(r)
}

// 窗口_置顶.
//
// hWindow: 窗口句柄.

// ff:窗口_置顶
// hWindow:窗口句柄
func X窗口_置顶(窗口句柄 int) int {
	r, _, _ := xWnd_SetTop.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_注册事件C.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.

// ff:窗口_注册事件C
// pFun:
// nEvent:事件类型
// hWindow:窗口句柄
func X窗口_注册事件C(窗口句柄 int, 事件类型 炫彩常量类.WM_, pFun interface{}) bool {
	r, _, _ := xWnd_RegEventC.Call(uintptr(窗口句柄), uintptr(事件类型), syscall.NewCallback(pFun))
	return r != 0
}

// 窗口_注册事件C1.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.

// ff:窗口_注册事件C1
// pFun:
// nEvent:事件类型
// hWindow:窗口句柄
func X窗口_注册事件C1(窗口句柄 int, 事件类型 炫彩常量类.WM_, pFun interface{}) bool {
	r, _, _ := xWnd_RegEventC1.Call(uintptr(窗口句柄), uintptr(事件类型), syscall.NewCallback(pFun))
	return r != 0
}

// 窗口_移除事件C.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.

// ff:窗口_移除事件C
// pFun:
// nEvent:事件类型
// hWindow:窗口句柄
func X窗口_移除事件C(窗口句柄 int, 事件类型 炫彩常量类.WM_, pFun interface{}) bool {
	r, _, _ := xWnd_RemoveEventC.Call(uintptr(窗口句柄), uintptr(事件类型), syscall.NewCallback(pFun))
	return r != 0
}

// 窗口_注册事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_ .
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.

// ff:窗口_注册事件CEx
// pFun:
// nEvent:事件类型
// hWindow:窗口句柄
func X窗口_注册事件CEx(窗口句柄 int, 事件类型 炫彩常量类.WM_, pFun uintptr) bool {
	r, _, _ := xWnd_RegEventC.Call(uintptr(窗口句柄), uintptr(事件类型), pFun)
	return r != 0
}

// 窗口_注册事件C1Ex, 和非Ex版相比只是最后一个参数不同.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.

// ff:窗口_注册事件C1Ex
// pFun:
// nEvent:事件类型
// hWindow:窗口句柄
func X窗口_注册事件C1Ex(窗口句柄 int, 事件类型 炫彩常量类.WM_, pFun uintptr) bool {
	r, _, _ := xWnd_RegEventC1.Call(uintptr(窗口句柄), uintptr(事件类型), pFun)
	return r != 0
}

// 窗口_移除事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// hWindow: 窗口句柄.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.

// ff:窗口_移除事件CEx
// pFun:
// nEvent:事件类型
// hWindow:窗口句柄
func X窗口_移除事件CEx(窗口句柄 int, 事件类型 炫彩常量类.WM_, pFun uintptr) bool {
	r, _, _ := xWnd_RemoveEventC.Call(uintptr(窗口句柄), uintptr(事件类型), pFun)
	return r != 0
}

// 窗口_添加子对象.
//
// hWindow: 窗口句柄.
//
// hChild: 要添加的对象句柄.

// ff:窗口_添加子对象
// hChild:要添加的对象句柄
// hWindow:窗口句柄
func X窗口_添加子对象(窗口句柄 int, 要添加的对象句柄 int) bool {
	r, _, _ := xWnd_AddChild.Call(uintptr(窗口句柄), uintptr(要添加的对象句柄))
	return r != 0
}

// 窗口_插入子对象.
//
// hWindow: 窗口句柄.
//
// hChild: 要插入的对象句柄.
//
// index: 插入位置索引.

// ff:窗口_插入子对象
// index:插入位置索引
// hChild:要插入的对象句柄
// hWindow:窗口句柄
func X窗口_插入子对象(窗口句柄 int, 要插入的对象句柄 int, 插入位置索引 int) bool {
	r, _, _ := xWnd_InsertChild.Call(uintptr(窗口句柄), uintptr(要插入的对象句柄), uintptr(插入位置索引))
	return r != 0
}

// 窗口_取HWND.
//
// hWindow: 窗口句柄.

// ff:窗口_取HWND
// hWindow:窗口句柄
func X窗口_取HWND(窗口句柄 int) uintptr {
	r, _, _ := xWnd_GetHWND.Call(uintptr(窗口句柄))
	return r
}

// 窗口_重绘.
//
// hWindow: 窗口资源句柄.
//
// bImmediate: 是否立即重绘, 默认为否.

// ff:窗口_重绘
// bImmediate:是否立即重绘
// hWindow:窗口资源句柄
func X窗口_重绘(窗口资源句柄 int, 是否立即重绘 bool) int {
	r, _, _ := xWnd_Redraw.Call(uintptr(窗口资源句柄), 炫彩工具类.BoolPtr(是否立即重绘))
	return int(r)
}

// 窗口_重绘指定区域.
//
// hWindow: 窗口资源句柄.
//
// pRect: 需要重绘的区域坐标.
//
// bImmediate: TRUE立即重绘, FALSE放入消息队列延迟重绘.

// ff:窗口_重绘指定区域
// bImmediate:TRUE立即重绘
// pRect:区域坐标
// hWindow:窗口资源句柄
func X窗口_重绘指定区域(窗口资源句柄 int, 区域坐标 *RECT, TRUE立即重绘 bool) int {
	r, _, _ := xWnd_RedrawRect.Call(uintptr(窗口资源句柄), uintptr(unsafe.Pointer(区域坐标)), 炫彩工具类.BoolPtr(TRUE立即重绘))
	return int(r)
}

// 窗口_置坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 坐标.

// ff:窗口_置坐标
// pRect:坐标
// hWindow:窗口句柄
func X窗口_置坐标(窗口句柄 int, 坐标 *RECT) bool {
	r, _, _ := xWnd_SetRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(坐标)))
	return r != 0
}

// 窗口_置焦点.
//
// hWindow: 窗口资源句柄.
//
// hFocusEle: 将获得焦点的元素.

// ff:窗口_置焦点
// hFocusEle:将获得焦点的元素
// hWindow:窗口资源句柄
func X窗口_置焦点(窗口资源句柄 int, 将获得焦点的元素 int) bool {
	r, _, _ := xWnd_SetFocusEle.Call(uintptr(窗口资源句柄), uintptr(将获得焦点的元素))
	return r != 0
}

// 窗口_取焦点.
//
// hWindow: 窗口资源句柄.

// ff:窗口_取焦点
// hWindow:窗口资源句柄
func X窗口_取焦点(窗口资源句柄 int) int {
	r, _, _ := xWnd_GetFocusEle.Call(uintptr(窗口资源句柄))
	return int(r)
}

// 窗口_取鼠标停留元素.
//
// hWindow: 窗口资源句柄.

// ff:窗口_取鼠标停留元素
// hWindow:窗口资源句柄
func X窗口_取鼠标停留元素(窗口资源句柄 int) int {
	r, _, _ := xWnd_GetStayEle.Call(uintptr(窗口资源句柄))
	return int(r)
}

// 窗口_绘制, 在自绘事件函数中,用户手动调用绘制窗口, 以便控制绘制顺序.
//
// hWindow: 窗口资源句柄.
//
// hDraw: 图形绘制句柄.

// ff:窗口_绘制
// hDraw:图形绘制句柄
// hWindow:窗口资源句柄
func X窗口_绘制(窗口资源句柄 int, 图形绘制句柄 int) int {
	r, _, _ := xWnd_DrawWindow.Call(uintptr(窗口资源句柄), uintptr(图形绘制句柄))
	return int(r)
}

// 窗口_居中.
//
// hWindow: 窗口资源句柄.

// ff:窗口_居中
// hWindow:窗口资源句柄
func X窗口_居中(窗口资源句柄 int) int {
	r, _, _ := xWnd_Center.Call(uintptr(窗口资源句柄))
	return int(r)
}

// 窗口_居中扩展.
//
// hWindow: 窗口资源句柄.
//
// width: 窗口宽度.
//
// height: 窗口高度.

// ff:窗口_居中EX
// height:窗口高度
// width:窗口宽度
// hWindow:窗口资源句柄
func X窗口_居中EX(窗口资源句柄, 窗口宽度, 窗口高度 int) int {
	r, _, _ := xWnd_CenterEx.Call(uintptr(窗口资源句柄), uintptr(窗口宽度), uintptr(窗口高度))
	return int(r)
}

// 窗口_置光标.
//
// hWindow: 窗口句柄.
//
// hCursor: 鼠标光标句柄.

// ff:窗口_置光标
// hCursor:鼠标光标句柄
// hWindow:窗口句柄
func X窗口_置光标(窗口句柄 int, 鼠标光标句柄 int) int {
	r, _, _ := xWnd_SetCursor.Call(uintptr(窗口句柄), uintptr(鼠标光标句柄))
	return int(r)
}

// 窗口_取光标.
//
// hWindow: 窗口句柄.

// ff:窗口_取光标
// hWindow:窗口句柄
func X窗口_取光标(窗口句柄 int) int {
	r, _, _ := xWnd_GetCursor.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_启用拖动边框.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用拖动边框
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用拖动边框(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableDragBorder.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_启用拖动窗口.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用拖动窗口
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用拖动窗口(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableDragWindow.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_启用拖动标题栏.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用拖动标题栏
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用拖动标题栏(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableDragCaption.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_启用绘制背景.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用绘制背景
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用绘制背景(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableDrawBk.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_启用自动焦点.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用自动焦点
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用自动焦点(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableAutoFocus.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_启用允许最大化.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用允许最大化
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用允许最大化(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableMaxWindow.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_启用限制窗口大小.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用限制窗口大小
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用限制窗口大小(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableLimitWindowSize.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_启用布局.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用布局
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用布局(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableLayout.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_启用布局覆盖边框.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用布局覆盖边框
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用布局覆盖边框(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_EnableLayoutOverlayBorder.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_显示布局边界.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_显示布局边界
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_显示布局边界(窗口句柄 int, 是否启用 bool) int {
	r, _, _ := xWnd_ShowLayoutFrame.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口_判断启用布局.
//
// hWindow: 窗口句柄.

// ff:窗口_判断启用布局
// hWindow:窗口句柄
func X窗口_判断启用布局(窗口句柄 int) bool {
	r, _, _ := xWnd_IsEnableLayout.Call(uintptr(窗口句柄))
	return r != 0
}

// 窗口_是否最大化.
//
// hWindow: 窗口句柄.

// ff:窗口_是否最大化
// hWindow:窗口句柄
func X窗口_是否最大化(窗口句柄 int) bool {
	r, _, _ := xWnd_IsMaxWindow.Call(uintptr(窗口句柄))
	return r != 0
}

// 窗口_置鼠标捕获元素.
//
// hWindow: 窗口句柄.
//
// hEle: 元素句柄.

// ff:窗口_置鼠标捕获元素
// hEle:元素句柄
// hWindow:窗口句柄
func X窗口_置鼠标捕获元素(窗口句柄, 元素句柄 int) int {
	r, _, _ := xWnd_SetCaptureEle.Call(uintptr(窗口句柄), uintptr(元素句柄))
	return int(r)
}

// 窗口_取鼠标捕获元素.
//
// hWindow: 窗口句柄.

// ff:窗口_取鼠标捕获元素
// hWindow:窗口句柄
func X窗口_取鼠标捕获元素(窗口句柄 int) int {
	r, _, _ := xWnd_GetCaptureEle.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_取绘制矩形.
//
// hWindow: 窗口句柄.
//
// pRcPaint: 重绘区域坐标.

// ff:窗口_取绘制矩形
// pRcPaint:重绘区域坐标
// hWindow:窗口句柄
func X窗口_取绘制矩形(窗口句柄 int, 重绘区域坐标 *RECT) int {
	r, _, _ := xWnd_GetDrawRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(重绘区域坐标)))
	return int(r)
}

// 窗口_置系统光标.
//
// hWindow: 窗口句柄.
//
// hCursor: 光标句柄.

// ff:窗口_置系统光标
// hCursor:光标句柄
// hWindow:窗口句柄
func X窗口_置系统光标(窗口句柄, 光标句柄 int) int {
	r, _, _ := xWnd_SetCursorSys.Call(uintptr(窗口句柄), uintptr(光标句柄))
	return int(r)
}

// 窗口_置字体.
//
// hWindow: 窗口句柄.
//
// hFontx: 炫彩字体句柄.

// ff:窗口_置字体
// hFontx:炫彩字体句柄
// hWindow:窗口句柄
func X窗口_置字体(窗口句柄, 炫彩字体句柄 int) int {
	r, _, _ := xWnd_SetFont.Call(uintptr(窗口句柄), uintptr(炫彩字体句柄))
	return int(r)
}

// 窗口_置文本颜色.
//
// hWindow: 窗口句柄.
//
// color: ABGR 颜色值.

// ff:窗口_置文本颜色
// color:ABGR颜色值
// hWindow:窗口句柄
func X窗口_置文本颜色(窗口句柄, ABGR颜色值 int) int {
	r, _, _ := xWnd_SetTextColor.Call(uintptr(窗口句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 窗口_取文本颜色.
//
// hWindow: 窗口句柄.

// ff:窗口_取文本颜色
// hWindow:窗口句柄
func X窗口_取文本颜色(窗口句柄 int) int {
	r, _, _ := xWnd_GetTextColor.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_取文本颜色扩展.
//
// hWindow: 窗口句柄.

// ff:窗口_取文本颜色EX
// hWindow:窗口句柄
func X窗口_取文本颜色EX(窗口句柄 int) int {
	r, _, _ := xWnd_GetTextColorEx.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_置ID.
//
// hWindow: 窗口句柄.
//
// nID: ID值.

// ff:窗口_置ID
// nID:ID值
// hWindow:窗口句柄
func X窗口_置ID(窗口句柄, ID值 int) int {
	r, _, _ := xWnd_SetID.Call(uintptr(窗口句柄), uintptr(ID值))
	return int(r)
}

// 窗口_取ID.
//
// hWindow: 窗口句柄.

// ff:窗口_取ID
// hWindow:窗口句柄
func X窗口_取ID(窗口句柄 int) int {
	r, _, _ := xWnd_GetID.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_置名称.
//
// hWindow: 窗口句柄.
//
// pName: name值, 字符串.

// ff:窗口_置名称
// pName:名称
// hWindow:窗口句柄
func X窗口_置名称(窗口句柄 int, 名称 string) int {
	r, _, _ := xWnd_SetName.Call(uintptr(窗口句柄), 炫彩工具类.StrPtr(名称))
	return int(r)
}

// 窗口_取名称.
//
// hWindow: 窗口句柄.

// ff:窗口_取名称
// hWindow:窗口句柄
func X窗口_取名称(窗口句柄 int) string {
	r, _, _ := xWnd_GetName.Call(uintptr(窗口句柄))
	return 炫彩工具类.UintPtrToString(r)
}

// 窗口_置边大小.
//
// hWindow: 窗口句柄.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底部大小.

// ff:窗口_置边大小
// bottom:底部
// right:右边
// top:上边
// left:左边
// hWindow:窗口句柄
func X窗口_置边大小(窗口句柄, 左边, 上边, 右边, 底部 int) int {
	r, _, _ := xWnd_SetBorderSize.Call(uintptr(窗口句柄), uintptr(左边), uintptr(上边), uintptr(右边), uintptr(底部))
	return int(r)
}

// 窗口_取边大小.
//
// hWindow: 窗口句柄.

// ff:窗口_取边大小
// pBorder:
// hWindow:窗口句柄
func X窗口_取边大小(窗口句柄 int, pBorder *RECT) int {
	r, _, _ := xWnd_GetBorderSize.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(pBorder)))
	return int(r)
}

// 窗口_置布局内填充大小.
//
// hWindow: 窗口句柄.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.

// ff:窗口_置布局内填充大小
// bottom:底部
// right:右边
// top:上边
// left:左边
// hWindow:窗口句柄
func X窗口_置布局内填充大小(窗口句柄, 左边, 上边, 右边, 底部 int) int {
	r, _, _ := xWnd_SetPadding.Call(uintptr(窗口句柄), uintptr(左边), uintptr(上边), uintptr(右边), uintptr(底部))
	return int(r)
}

// 窗口_置拖动边框大小.
//
// hWindow: 窗口句柄.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底边大小.

// ff:窗口_置拖动边框大小
// bottom:底部
// right:右边
// top:上边
// left:左边
// hWindow:窗口句柄
func X窗口_置拖动边框大小(窗口句柄, 左边, 上边, 右边, 底部 int) int {
	r, _, _ := xWnd_SetDragBorderSize.Call(uintptr(窗口句柄), uintptr(左边), uintptr(上边), uintptr(右边), uintptr(底部))
	return int(r)
}

// 窗口_取拖动边框大小.
//
// hWindow: 窗口句柄.
//
// pSize: 拖动边框大小.

// ff:窗口_取拖动边框大小
// pBorder:
// hWindow:窗口句柄
func X窗口_取拖动边框大小(窗口句柄 int, pBorder *RECT) int {
	r, _, _ := xWnd_GetDragBorderSize.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(pBorder)))
	return int(r)
}

// 窗口_置最小大小.
//
// hWindow: 窗口句柄.
//
// width: 最小宽度.
//
// height: 最小高度.

// ff:窗口_置最小大小
// height:最小高度
// width:最小宽度
// hWindow:窗口句柄
func X窗口_置最小大小(窗口句柄, 最小宽度, 最小高度 int) int {
	r, _, _ := xWnd_SetMinimumSize.Call(uintptr(窗口句柄), uintptr(最小宽度), uintptr(最小高度))
	return int(r)
}

// 窗口_测试点击子元素.
//
// hWindow: 窗口句柄.
//
// pPt: 左边点.

// ff:窗口_测试点击子元素
// pPt:左边点
// hWindow:窗口句柄
func X窗口_测试点击子元素(窗口句柄 int, 左边点 *POINT) int {
	r, _, _ := xWnd_HitChildEle.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(左边点)))
	return int(r)
}

// 窗口_取子对象数量.
//
// hWindow: 窗口句柄.

// ff:窗口_取子对象数量
// hWindow:窗口句柄
func X窗口_取子对象数量(窗口句柄 int) int {
	r, _, _ := xWnd_GetChildCount.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_取子对象从索引.
//
// hWindow: 窗口句柄.
//
// index: 元素索引.

// ff:窗口_取子对象从索引
// index:元素索引
// hWindow:窗口句柄
func X窗口_取子对象从索引(窗口句柄, 元素索引 int) int {
	r, _, _ := xWnd_GetChildByIndex.Call(uintptr(窗口句柄), uintptr(元素索引))
	return int(r)
}

// 窗口_取子对象从ID.
//
// hWindow: 窗口句柄.
//
// nID: 元素ID, ID必须大于0.

// ff:窗口_取子对象从ID
// nID:元素ID
// hWindow:窗口句柄
func X窗口_取子对象从ID(窗口句柄, 元素ID int) int {
	r, _, _ := xWnd_GetChildByID.Call(uintptr(窗口句柄), uintptr(元素ID))
	return int(r)
}

// 窗口_取子对象.
//
// hWindow: 窗口句柄.
//
// nID: 对象ID,ID必须大于0.

// ff:窗口_取子对象
// nID:对象ID
// hWindow:窗口句柄
func X窗口_取子对象(窗口句柄, 对象ID int) int {
	r, _, _ := xWnd_GetChild.Call(uintptr(窗口句柄), uintptr(对象ID))
	return int(r)
}

// 窗口_关闭.
//
// hWindow: 窗口句柄.

// ff:窗口_关闭
// hWindow:窗口句柄
func X窗口_关闭(窗口句柄 int) int {
	r, _, _ := xWnd_CloseWindow.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_调整布局.
//
// hWindow: 窗口句柄.

// ff:窗口_调整布局
// hWindow:窗口句柄
func X窗口_调整布局(窗口句柄 int) int {
	r, _, _ := xWnd_AdjustLayout.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_调整布局扩展.
//
// hWindow: 窗口句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.

// ff:窗口_调整布局EX
// nFlags:调整布局标识位
// hWindow:窗口句柄
func X窗口_调整布局EX(窗口句柄 int, 调整布局标识位 炫彩常量类.AdjustLayout_) int {
	r, _, _ := xWnd_AdjustLayoutEx.Call(uintptr(窗口句柄), uintptr(调整布局标识位))
	return int(r)
}

// 窗口_创建插入符.
//
// hWindow: 窗口句柄.
//
// hEle: 元素句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.

// ff:窗口_创建插入符
// height:高度
// width:宽度
// y:y坐标
// x:x坐标
// hEle:元素句柄
// hWindow:窗口句柄
func X窗口_创建插入符(窗口句柄, 元素句柄, x坐标, y坐标, 宽度, 高度 int) int {
	r, _, _ := xWnd_CreateCaret.Call(uintptr(窗口句柄), uintptr(元素句柄), uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 窗口_置插入符位置, 设置插入符位置.
//
// hWindow: 窗口句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
//
// bUpdate: 是否立即更新UI.

// ff:窗口_置插入符位置
// bUpdate:是否立即更新UI
// height:高度
// width:宽度
// y:y坐标
// x:x坐标
// hWindow:窗口句柄
func X窗口_置插入符位置(窗口句柄, x坐标, y坐标, 宽度, 高度 int, 是否立即更新UI bool) int {
	r, _, _ := xWnd_SetCaretPos.Call(uintptr(窗口句柄), uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.BoolPtr(是否立即更新UI))
	return int(r)
}

// 窗口_置插入符颜色.
//
// hWindow: 窗口句柄.
//
// color: 颜色值.

// ff:窗口_置插入符颜色
// color:颜色值
// hWindow:窗口句柄
func X窗口_置插入符颜色(窗口句柄, 颜色值 int) int {
	r, _, _ := xWnd_SetCaretColor.Call(uintptr(窗口句柄), uintptr(颜色值))
	return int(r)
}

// 窗口_显示插入符.
//
// hWindow: 窗口句柄.
//
// bShow: 是否显示.

// ff:窗口_显示插入符
// bShow:是否显示
// hWindow:窗口句柄
func X窗口_显示插入符(窗口句柄 int, 是否显示 bool) int {
	r, _, _ := xWnd_ShowCaret.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 窗口_销毁插入符.
//
// hWindow: 窗口句柄.

// ff:窗口_销毁插入符
// hWindow:窗口句柄
func X窗口_销毁插入符(窗口句柄 int) int {
	r, _, _ := xWnd_DestroyCaret.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_取插入符元素.
//
// hWindow: 窗口句柄.

// ff:窗口_取插入符元素
// hWindow:窗口句柄
func X窗口_取插入符元素(窗口句柄 int) int {
	r, _, _ := xWnd_GetCaretHELE.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_取客户区坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 坐标.

// ff:窗口_取客户区坐标
// pRect:坐标
// hWindow:窗口句柄
func X窗口_取客户区坐标(窗口句柄 int, 坐标 *RECT) int {
	r, _, _ := xWnd_GetClientRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 窗口_取Body坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 坐标.

// ff:窗口_取Body坐标
// pRect:坐标
// hWindow:窗口句柄
func X窗口_取Body坐标(窗口句柄 int, 坐标 *RECT) int {
	r, _, _ := xWnd_GetBodyRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 窗口_取布局坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 接收返回坐标.

// ff:窗口_取布局坐标
// pRect:接收返回坐标
// hWindow:窗口句柄
func X窗口_取布局坐标(窗口句柄 int, 接收返回坐标 *RECT) int {
	r, _, _ := xWnd_GetLayoutRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(接收返回坐标)))
	return int(r)
}

// 窗口_移动窗口.
//
// hWindow: 窗口句柄.
//
// x: X坐标.
//
// y: Y坐标.

// ff:窗口_移动窗口
// y:Y坐标
// x:坐标
// hWindow:窗口句柄
func X窗口_移动窗口(窗口句柄 int, 坐标, Y坐标 int32) {
	xWnd_SetPosition.Call(uintptr(窗口句柄), uintptr(坐标), uintptr(Y坐标))
}

// 窗口_取坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 坐标.

// ff:窗口_取坐标
// pRect:坐标
// hWindow:窗口句柄
func X窗口_取坐标(窗口句柄 int, 坐标 *RECT) int {
	r, _, _ := xWnd_GetRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 窗口_最大化.
//
// hWindow: 窗口句柄.
//
// bMaximize: 是否最大化.

// ff:窗口_最大化
// bMaximize:是否最大化
// hWindow:窗口句柄
func X窗口_最大化(窗口句柄 int, 是否最大化 bool) int {
	r, _, _ := xWnd_MaxWindow.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否最大化))
	return int(r)
}

// 窗口_置定时器.
//
// hWindow: 窗口句柄.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.

// ff:窗口_置定时器
// uElapse:间隔值
// nIDEvent:定时器ID
// hWindow:窗口句柄
func X窗口_置定时器(窗口句柄, 定时器ID, 间隔值 int) int {
	r, _, _ := xWnd_SetTimer.Call(uintptr(窗口句柄), uintptr(定时器ID), uintptr(间隔值))
	return int(r)
}

// 窗口_关闭定时器.
//
// hWindow: 窗口句柄.
//
// nIDEvent: 定时器ID.

// ff:窗口_关闭定时器
// nIDEvent:定时器ID
// hWindow:窗口句柄
func X窗口_关闭定时器(窗口句柄, 定时器ID int) int {
	r, _, _ := xWnd_KillTimer.Call(uintptr(窗口句柄), uintptr(定时器ID))
	return int(r)
}

// 窗口_置炫彩定时器.
//
// hWindow: 窗口句柄.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.

// ff:窗口_置炫彩定时器
// uElapse:间隔值
// nIDEvent:定时器ID
// hWindow:窗口句柄
func X窗口_置炫彩定时器(窗口句柄, 定时器ID, 间隔值 int) int {
	r, _, _ := xWnd_SetXCTimer.Call(uintptr(窗口句柄), uintptr(定时器ID), uintptr(间隔值))
	return int(r)
}

// 窗口_关闭炫彩定时器.
//
// hWindow: 窗口句柄.
//
// nIDEvent: 定时器ID.

// ff:窗口_关闭炫彩定时器
// nIDEvent:定时器ID
// hWindow:窗口句柄
func X窗口_关闭炫彩定时器(窗口句柄, 定时器ID int) int {
	r, _, _ := xWnd_KillXCTimer.Call(uintptr(窗口句柄), uintptr(定时器ID))
	return int(r)
}

// 窗口_取背景管理器.
//
// hWindow: 窗口句柄.

// ff:窗口_取背景管理器
// hWindow:窗口句柄
func X窗口_取背景管理器(窗口句柄 int) int {
	r, _, _ := xWnd_GetBkManager.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_取背景管理器扩展.
//
// hWindow: 窗口句柄.

// ff:窗口_取背景管理器EX
// hWindow:窗口句柄
func X窗口_取背景管理器EX(窗口句柄 int) int {
	r, _, _ := xWnd_GetBkManagerEx.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_置背景管理器.
//
// hWindow: 窗口句柄.
//
// hBkInfoM: 背景管理器.

// ff:窗口_置背景管理器
// hBkInfoM:背景管理器
// hWindow:窗口句柄
func X窗口_置背景管理器(窗口句柄, 背景管理器 int) int {
	r, _, _ := xWnd_SetBkMagager.Call(uintptr(窗口句柄), uintptr(背景管理器))
	return int(r)
}

// 窗口_置透明类型.
//
// hWindow: 窗口句柄.
//
// nType: 窗口透明类型: xcc.Window_Transparent_.

// ff:窗口_置透明类型
// nType:透明类型
// hWindow:窗口句柄
func X窗口_置透明类型(窗口句柄 int, 透明类型 炫彩常量类.Window_Transparent_) int {
	r, _, _ := xWnd_SetTransparentType.Call(uintptr(窗口句柄), uintptr(透明类型))
	return int(r)
}

// 窗口_置透明度.
//
// hWindow: 窗口句柄.
//
// alpha: 窗口透明度, 范围0-255之间, 0透明, 255不透明.

// ff:窗口_置透明度
// alpha:透明度
// hWindow:窗口句柄
func X窗口_置透明度(窗口句柄 int, 透明度 byte) int {
	r, _, _ := xWnd_SetTransparentAlpha.Call(uintptr(窗口句柄), uintptr(透明度))
	return int(r)
}

// 窗口_置透明色.
//
// hWindow: 窗口句柄.
//
// color: 窗口透明色.

// ff:窗口_置透明色
// color:透明色
// hWindow:窗口句柄
func X窗口_置透明色(窗口句柄, 透明色 int) int {
	r, _, _ := xWnd_SetTransparentColor.Call(uintptr(窗口句柄), uintptr(透明色))
	return int(r)
}

// 窗口_置阴影信息.
//
// hWindow: 窗口句柄.
//
// nSize: 阴影大小.
//
// nDepth: 阴影深度, 0-255.
//
// nAngeleSize: 圆角阴影内收大小.
//
// bRightAngle: 是否强制直角.
//
// color: 阴影颜色.

// ff:窗口_置阴影信息
// color:阴影颜色
// bRightAngle:是否强制直角
// nAngeleSize:圆角阴影内收大小
// nDepth:阴影深度
// nSize:阴影大小
// hWindow:窗口句柄
func X窗口_置阴影信息(窗口句柄, 阴影大小 int, 阴影深度 byte, 圆角阴影内收大小 int, 是否强制直角 bool, 阴影颜色 int) int {
	r, _, _ := xWnd_SetShadowInfo.Call(uintptr(窗口句柄), uintptr(阴影大小), uintptr(阴影深度), uintptr(圆角阴影内收大小), 炫彩工具类.BoolPtr(是否强制直角), uintptr(阴影颜色))
	return int(r)
}

// 窗口_取阴影信息.
//
// hWindow: 窗口句柄.
//
// pnSize: 阴影大小.
//
// pnDepth: 阴影深度.
//
// pnAngeleSize: 圆角阴影内收大小 .
//
// pbRightAngle: 是否强制直角.
//
// pColor: 阴影颜色.

// ff:窗口_取阴影信息
// pColor:阴影颜色
// pbRightAngle:是否强制直角
// pnAngeleSize:圆角阴影内收大小
// pnDepth:阴影深度
// pnSize:阴影大小
// hWindow:窗口句柄
func X窗口_取阴影信息(窗口句柄 int, 阴影大小, 阴影深度, 圆角阴影内收大小 *int32, 是否强制直角 *bool, 阴影颜色 *int) int {
	r, _, _ := xWnd_GetShadowInfo.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(阴影大小)), uintptr(unsafe.Pointer(阴影深度)), uintptr(unsafe.Pointer(圆角阴影内收大小)), uintptr(unsafe.Pointer(是否强制直角)), uintptr(unsafe.Pointer(阴影颜色)))
	return int(r)
}

// 窗口_取透明类型, 返回: xcc.Window_Transparent_.
//
// hWindow: 窗口句柄.

// ff:窗口_取透明类型
// hWindow:窗口句柄
func X窗口_取透明类型(窗口句柄 int) 炫彩常量类.Window_Transparent_ {
	r, _, _ := xWnd_GetTransparentType.Call(uintptr(窗口句柄))
	return 炫彩常量类.Window_Transparent_(r)
}

// 窗口_附加窗口, 返回窗口资源句柄.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.

// ff:窗口_附加窗口
// XCStyle:窗口样式
// hWnd:外部窗口句柄
func X窗口_附加窗口(外部窗口句柄 uintptr, 窗口样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xWnd_Attach.Call(外部窗口句柄, uintptr(窗口样式))
	return int(r)
}

// 窗口_启用拖放文件.
//
// hWindow: 窗口句柄.
//
// bEnable: 是否启用.

// ff:窗口_启用拖放文件
// bEnable:是否启用
// hWindow:窗口句柄
func X窗口_启用拖放文件(窗口句柄 int, 是否启用 bool) bool {
	r, _, _ := xWnd_EnableDragFiles.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否启用))
	return r != 0
}

// 窗口_显示 显示隐藏窗口.
//
// hWindow: 窗口句柄.
//
// bShow: 是否显示.

// ff:窗口_显示
// bShow:是否显示
// hWindow:窗口句柄
func X窗口_显示(窗口句柄 int, 是否显示 bool) int {
	r, _, _ := xWnd_Show.Call(uintptr(窗口句柄), 炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 窗口_取插入符信息, 获取插入符信息, 返回: 插入符元素句柄.
//
// hWindow: 窗口句柄.
//
// pX: 接收返回x坐标.
//
// pY: 接收返回y坐标.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.

// ff:窗口_取插入符信息
// pHeight:接收返回高度
// pWidth:接收返回宽度
// pY:接收返回y坐标
// pX:接收返回x坐标
// hWindow:窗口句柄
func X窗口_取插入符信息(窗口句柄 int, 接收返回x坐标, 接收返回y坐标, 接收返回宽度, 接收返回高度 *int32) int {
	r, _, _ := xWnd_GetCaretInfo.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(接收返回x坐标)), uintptr(unsafe.Pointer(接收返回y坐标)), uintptr(unsafe.Pointer(接收返回宽度)), uintptr(unsafe.Pointer(接收返回高度)))
	return int(r)
}

// 窗口_置图标.
//
// hWindow: 窗口句柄.
//
// hImage: 图标句柄.

// ff:窗口_置图标
// hImage:图片句柄
// hWindow:窗口句柄
func X窗口_置图标(窗口句柄, 图片句柄 int) int {
	r, _, _ := xWnd_SetIcon.Call(uintptr(窗口句柄), uintptr(图片句柄))
	return int(r)
}

// 窗口_置标题.
//
// hWindow: 窗口句柄.
//
// pTitle: 标题文本.

// ff:窗口_置标题
// pTitle:标题文本
// hWindow:窗口句柄
func X窗口_置标题(窗口句柄 int, 标题文本 string) int {
	r, _, _ := xWnd_SetTitle.Call(uintptr(窗口句柄), 炫彩工具类.StrPtr(标题文本))
	return int(r)
}

// 窗口_置标题颜色.
//
// hWindow: 窗口句柄.
//
// color: ABGR 颜色.

// ff:窗口_置标题颜色
// color:ABGR颜色
// hWindow:窗口句柄
func X窗口_置标题颜色(窗口句柄, ABGR颜色 int) int {
	r, _, _ := xWnd_SetTitleColor.Call(uintptr(窗口句柄), uintptr(ABGR颜色))
	return int(r)
}

// 窗口_取控制按钮, 返回按钮句柄.
//
// hWindow: 窗口句柄.
//
// nFlag: xcc.Window_Style_ . 可用值: xcc.Window_Style_Btn_Min , xcc.Window_Style_Btn_Max , xcc.Window_Style_Btn_Close .

// ff:窗口_取控制按钮
// nFlag:标志
// hWindow:窗口句柄
func X窗口_取控制按钮(窗口句柄 int, 标志 炫彩常量类.Window_Style_) int {
	r, _, _ := xWnd_GetButton.Call(uintptr(窗口句柄), uintptr(标志))
	return int(r)
}

// 窗口_取图标, 返回图标句柄.
//
// hWindow: 窗口句柄.

// ff:窗口_取图标
// hWindow:窗口句柄
func X窗口_取图标(窗口句柄 int) int {
	r, _, _ := xWnd_GetIcon.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_取标题, 返回标题文本.
//
// hWindow: 窗口句柄.

// ff:窗口_取标题
// hWindow:窗口句柄
func X窗口_取标题(窗口句柄 int) string {
	r, _, _ := xWnd_GetTitle.Call(uintptr(窗口句柄))
	return 炫彩工具类.UintPtrToString(r)
}

// 窗口_取标题颜色, 返回ABGR 颜色.
//
// hWindow: 窗口句柄.

// ff:窗口_取标题颜色
// hWindow:窗口句柄
func X窗口_取标题颜色(窗口句柄 int) int {
	r, _, _ := xWnd_GetTitleColor.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_添加背景边框.
//
// hWindow: 窗口句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.

// ff:窗口_添加背景边框
// width:
// color:
// nState:组合状态
// hWindow:窗口句柄
func X窗口_添加背景边框(窗口句柄 int, 组合状态 炫彩常量类.Window_State_Flag_, color int, width int) int {
	r, _, _ := xWnd_AddBkBorder.Call(uintptr(窗口句柄), uintptr(组合状态), uintptr(color), uintptr(width))
	return int(r)
}

// 窗口_添加背景填充.
//
// hWindow: 窗口句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.

// ff:窗口_添加背景填充
// color:
// nState:组合状态
// hWindow:窗口句柄
func X窗口_添加背景填充(窗口句柄 int, 组合状态 炫彩常量类.Window_State_Flag_, color int) int {
	r, _, _ := xWnd_AddBkFill.Call(uintptr(窗口句柄), uintptr(组合状态), uintptr(color))
	return int(r)
}

// 窗口_添加背景图片.
//
// hWindow: 窗口句柄.
//
// nState: 组合状态.
//
// hImage: 图片句柄.

// ff:窗口_添加背景图片
// hImage:
// nState:组合状态
// hWindow:窗口句柄
func X窗口_添加背景图片(窗口句柄 int, 组合状态 炫彩常量类.Window_State_Flag_, hImage int) int {
	r, _, _ := xWnd_AddBkImage.Call(uintptr(窗口句柄), uintptr(组合状态), uintptr(hImage))
	return int(r)
}

// 窗口_取背景对象数量.
//
// hWindow: 窗口句柄.

// ff:窗口_取背景对象数量
// hWindow:窗口句柄
func X窗口_取背景对象数量(窗口句柄 int) int {
	r, _, _ := xWnd_GetBkInfoCount.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_清空背景对象.
//
// hWindow: 窗口句柄.

// ff:窗口_清空背景对象
// hWindow:窗口句柄
func X窗口_清空背景对象(窗口句柄 int) int {
	r, _, _ := xWnd_ClearBkInfo.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_置背景, 返回设置的背景对象数量.
//
// hWindow: 窗口句柄.
//
// pText: 背景内容字符串.

// ff:窗口_置背景
// pText:背景内容字符串
// hWindow:窗口句柄
func X窗口_置背景(窗口句柄 int, 背景内容字符串 string) int {
	r, _, _ := xWnd_SetBkInfo.Call(uintptr(窗口句柄), 炫彩工具类.StrPtr(背景内容字符串))
	return int(r)
}

// 窗口_是否可拖动标题栏.
//
// hWindow: 窗口句柄.

// ff:窗口_是否可拖动标题栏
// hWindow:窗口句柄
func X窗口_是否可拖动标题栏(窗口句柄 int) bool {
	r, _, _ := xWnd_IsDragCaption.Call(uintptr(窗口句柄))
	return r != 0
}

// 窗口_是否可拖动窗口.
//
// hWindow: 窗口句柄.

// ff:窗口_是否可拖动窗口
// hWindow:窗口句柄
func X窗口_是否可拖动窗口(窗口句柄 int) bool {
	r, _, _ := xWnd_IsDragWindow.Call(uintptr(窗口句柄))
	return r != 0
}

// 窗口_是否可拖动边框.
//
// hWindow: 窗口句柄.

// ff:窗口_是否可拖动边框
// hWindow:窗口句柄
func X窗口_是否可拖动边框(窗口句柄 int) bool {
	r, _, _ := xWnd_IsDragBorder.Call(uintptr(窗口句柄))
	return r != 0
}

// 窗口_置标题外间距, 设置标题内容(图标, 标题, 控制按钮)外间距.
//
// hWindow: 窗口句柄.
//
// left: 左边间距.
//
// top: 上边间距.
//
// right: 右边间距.
//
// bottom: 下边间距.

// ff:窗口_置标题外间距
// bottom:下边间距
// right:右边间距
// top:上边间距
// left:左边间距
// hWindow:窗口句柄
func X窗口_置标题外间距(窗口句柄 int, 左边间距 int, 上边间距 int, 右边间距 int, 下边间距 int) int {
	r, _, _ := xWnd_SetCaptionMargin.Call(uintptr(窗口句柄), uintptr(左边间距), uintptr(上边间距), uintptr(右边间距), uintptr(下边间距))
	return int(r)
}

// 窗口_置窗口位置. 封装系统API SetWindowPos(), 内部做了DPI适配.
//
// hWindow: 窗口句柄.
//
// hWndInsertAfter: 在Z序中位于定位窗口之前的窗口句柄. 此参数必须是窗口HWND或以下值之一: xcc.HWND_.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// uFlags: 窗口大小调整和定位标志. 可以是以下值的组合: xcc.SWP_.

// ff:窗口_置窗口位置
// uFlags:
// cy:
// cx:
// y:
// x:
// hWndInsertAfter:置顶方式
// hWindow:窗口句柄
func X窗口_置窗口位置(窗口句柄 int, 置顶方式 炫彩常量类.HWND_, x, y, cx, cy int32, uFlags 炫彩常量类.SWP_) int {
	r, _, _ := xWnd_SetWindowPos.Call(uintptr(窗口句柄), uintptr(置顶方式), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return int(r)
}

// 窗口_取DPI. 获取当前窗口所在显示器DPI, 返回窗口DPI.
//
// hWindow: 窗口句柄.

// ff:窗口_取DPI
// hWindow:窗口句柄
func X窗口_取DPI(窗口句柄 int) int {
	r, _, _ := xWnd_GetDPI.Call(uintptr(窗口句柄))
	return int(r)
}

// 窗口_坐标转换DPI. 窗口客户区坐标转换到缩放后DPI坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 接收返回坐标.

// ff:窗口_坐标转换DPI
// pRect:接收返回坐标
// hWindow:窗口句柄
func X窗口_坐标转换DPI(窗口句柄 int, 接收返回坐标 *RECT) int {
	r, _, _ := xWnd_RectToDPI.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(接收返回坐标)))
	return int(r)
}

// 窗口_坐标点转换DPI. 窗口客户区坐标点转换到缩放后.
//
// hWindow: 窗口句柄.
//
// pPt: 接收返回坐标点.

// ff:窗口_坐标点转换DPI
// pPt:接收返回坐标点
// hWindow:窗口句柄
func X窗口_坐标点转换DPI(窗口句柄 int, 接收返回坐标点 *POINT) int {
	r, _, _ := xWnd_PointToDPI.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(接收返回坐标点)))
	return int(r)
}

// 窗口_取光标位置. 封装的系统API: GetCursorPos, 内部做了DPI适配.
//
// hWindow: 窗口句柄.
//
// pPt: 接收返回坐标点.

// ff:窗口_取光标位置
// pPt:接收返回坐标点
// hWindow:窗口句柄
func X窗口_取光标位置(窗口句柄 int, 接收返回坐标点 *POINT) bool {
	r, _, _ := xWnd_GetCursorPos.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(接收返回坐标点)))
	return r != 0
}

// 窗口_客户区坐标点到屏幕坐标点. 封装的系统API: ClientToScreen, 内部做了DPI适配.
//
// hWindow: 窗口句柄.
//
// pPt: 接收返回坐标点.

// ff:窗口_客户区坐标点到屏幕坐标点
// pPt:接收返回坐标点
// hWindow:窗口句柄
func X窗口_客户区坐标点到屏幕坐标点(窗口句柄 int, 接收返回坐标点 *POINT) bool {
	r, _, _ := xWnd_ClientToScreen.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(接收返回坐标点)))
	return r != 0
}

// 窗口_屏幕坐标点到客户区坐标点. 封装的系统API: ScreenToClient(), 内部做了DPI适配.
//
// hWindow: 窗口句柄.
//
// pPt: 接收返回坐标点.

// ff:窗口_屏幕坐标点到客户区坐标点
// pPt:接收返回坐标点
// hWindow:窗口句柄
func X窗口_屏幕坐标点到客户区坐标点(窗口句柄 int, 接收返回坐标点 *POINT) bool {
	r, _, _ := xWnd_ScreenToClient.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(接收返回坐标点)))
	return r != 0
}

// 窗口_置DPI. 设置当前窗口DPI, 默认DPI为96.
//
// hWindow: 窗口句柄.
//
// nDPI: DPI值.

// ff:窗口_置DPI
// nDPI:DPI值
// hWindow:窗口句柄
func X窗口_置DPI(窗口句柄 int, DPI值 int) int {
	r, _, _ := xWnd_SetDPI.Call(uintptr(窗口句柄), uintptr(DPI值))
	return int(r)
}
