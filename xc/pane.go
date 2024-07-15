package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"

	"github.com/888go/xcgui/xcc"
)

// 窗格_创建, 创建窗格元素, 返回元素句柄.
//
// pName: 窗格标题.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// hFrameWnd: 框架窗口.

// ff:窗格_创建
// hFrameWnd:框架窗口
// nHeight:高度
// nWidth:宽度
// pName:窗格标题
func X窗格_创建(窗格标题 string, 宽度 int, 高度 int, 框架窗口 int) int {
	r, _, _ := xPane_Create.Call(炫彩工具类.StrPtr(窗格标题), uintptr(宽度), uintptr(高度), uintptr(框架窗口))
	return int(r)
}

// 窗格_置视图, 设置窗格视图元素.
//
// hEle: 元素句柄.
//
// hView: 绑定视图元素.

// ff:窗格_置视图
// hView:绑定视图元素
// hEle:元素句柄
func X窗格_置视图(元素句柄 int, 绑定视图元素 int) int {
	r, _, _ := xPane_SetView.Call(uintptr(元素句柄), uintptr(绑定视图元素))
	return int(r)
}

// 窗格_置标题, 设置标题文本.
//
// hEle: 元素句柄.
//
// pTitle: 文本内容.

// ff:窗格_置标题
// pTitle:文本内容
// hEle:元素句柄
func X窗格_置标题(元素句柄 int, 文本内容 string) int {
	r, _, _ := xPane_SetTitle.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文本内容))
	return int(r)
}

// 窗格_取标题, 获取标题文本.
//
// hEle: 元素句柄.

// ff:窗格_取标题
// hEle:元素句柄
func X窗格_取标题(元素句柄 int) string {
	r, _, _ := xPane_GetTitle.Call(uintptr(元素句柄))
	return 炫彩工具类.UintPtrToString(r)
}

// 窗格_置标题栏高度, 设置标题栏高度.
//
// hEle: 元素句柄.
//
// nHeight: 高度.

// ff:窗格_置标题栏高度
// nHeight:高度
// hEle:元素句柄
func X窗格_置标题栏高度(元素句柄 int, 高度 int) int {
	r, _, _ := xPane_SetCaptionHeight.Call(uintptr(元素句柄), uintptr(高度))
	return int(r)
}

// 窗格_取标题栏高度, 获取标题栏高度.
//
// hEle: 元素句柄.

// ff:窗格_取标题栏高度
// hEle:元素句柄
func X窗格_取标题栏高度(元素句柄 int) int {
	r, _, _ := xPane_GetCaptionHeight.Call(uintptr(元素句柄))
	return int(r)
}

// 窗格_判断显示, 判断窗格是否显示.
//
// hEle: 元素句柄.

// ff:窗格_判断显示
// hEle:元素句柄
func X窗格_判断显示(元素句柄 int) bool {
	r, _, _ := xPane_IsShowPane.Call(uintptr(元素句柄))
	return r != 0
}

// 窗格_置大小, 设置窗格大小.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.

// ff:窗格_置大小
// nHeight:高度
// nWidth:宽度
// hEle:元素句柄
func X窗格_置大小(元素句柄 int, 宽度 int, 高度 int) int {
	r, _, _ := xPane_SetSize.Call(uintptr(元素句柄), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 窗格_取状态, 获取窗格停靠状态, 返回: Pane_State_.
//
// hEle: 元素句柄.

// ff:窗格_取状态
// hEle:元素句柄
func X窗格_取状态(元素句柄 int) 炫彩常量类.Pane_State_ {
	r, _, _ := xPane_GetState.Call(uintptr(元素句柄))
	return 炫彩常量类.Pane_State_(r)
}

// 窗格_取视图坐标, 获取窗格视图坐标.
//
// hEle: 元素句柄.
//
// pRect: 接收返回的坐标值.

// ff:窗格_取视图坐标
// pRect:接收返回坐标
// hEle:元素句柄
func X窗格_取视图坐标(元素句柄 int, 接收返回坐标 *RECT) int {
	r, _, _ := xPane_GetViewRect.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(接收返回坐标)))
	return int(r)
}

// XPane_HidePane 窗格_隐藏.
//
//	@param hEle 元素句柄.
//	@param bGroupActivate 当为窗格组成员时, 延迟处理窗格组成员激活的切换.
//	@return int

// ff:窗格_隐藏
// bGroupActivate:延迟组成员激活
// hEle:元素句柄
func X窗格_隐藏(元素句柄 int, 延迟组成员激活 bool) int {
	r, _, _ := xPane_HidePane.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(延迟组成员激活))
	return int(r)
}

// XPane_ShowPane 窗格_显示.
//
//	@param hEle 元素句柄.
//	@param bGroupActivate 如果是窗格组成员, 那么窗格组切换当前窗格为显示状态.
//	@return int

// ff:窗格_显示
// bGroupActivate:延迟组成员激活
// hEle:元素句柄
func X窗格_显示(元素句柄 int, 延迟组成员激活 bool) int {
	r, _, _ := xPane_ShowPane.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(延迟组成员激活))
	return int(r)
}

// 窗格_停靠, 窗格停靠到码头.
//
// hEle: 元素句柄.

// ff:窗格_停靠
// hEle:元素句柄
func X窗格_停靠(元素句柄 int) int {
	r, _, _ := xPane_DockPane.Call(uintptr(元素句柄))
	return int(r)
}

// 窗格_锁定, 锁定窗格.
//
// hEle: 元素句柄.

// ff:窗格_锁定
// hEle:元素句柄
func X窗格_锁定(元素句柄 int) int {
	r, _, _ := xPane_LockPane.Call(uintptr(元素句柄))
	return int(r)
}

// 窗格_浮动.
//
// hEle: 元素句柄.

// ff:窗格_浮动
// hEle:元素句柄
func X窗格_浮动(元素句柄 int) int {
	r, _, _ := xPane_FloatPane.Call(uintptr(元素句柄))
	return int(r)
}

// 窗格_绘制, 手动调用该函数绘制窗格, 以便控制绘制顺序.
//
// hEle: 元素句柄.
//
// hDraw: 图形绘制句柄.

// ff:窗格_绘制
// hDraw:图形绘制句柄
// hEle:元素句柄
func X窗格_绘制(元素句柄 int, 图形绘制句柄 int) int {
	r, _, _ := xPane_DrawPane.Call(uintptr(元素句柄), uintptr(图形绘制句柄))
	return int(r)
}

// 窗口_置选中, 如果窗格是组成员, 设置选中当前窗格可见.
//
// hEle: 元素句柄.

// ff:窗口_置选中
// hEle:元素句柄
func X窗口_置选中(元素句柄 int) bool {
	r, _, _ := xPane_SetSelect.Call(uintptr(元素句柄))
	return r != 0
}

// 窗格_是否激活. 判断窗格是否激活, 当为组成员时有效.
//
// hEle: 元素句柄.

// ff:窗格_是否激活
// hEle:元素句柄
func X窗格_是否激活(元素句柄 int) bool {
	r, _, _ := xPane_IsGroupActivate.Call(uintptr(元素句柄))
	return r != 0
}
