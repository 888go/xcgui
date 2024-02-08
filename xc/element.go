package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"syscall"
	"unsafe"
	
	"github.com/888go/xcgui/xcc"
)

// 元素_创建, 创建基础元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func X元素_创建(x坐标, y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xEle_Create.Call(uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 元素_注册事件C, 注册事件C方式, 省略2参数.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func X元素_注册事件C(元素句柄 int, 事件类型 炫彩常量类.XE_, 事件函数 interface{}) bool {
	r, _, _ := xEle_RegEventC.Call(uintptr(元素句柄), uintptr(事件类型), syscall.NewCallback(事件函数))
	return r != 0
}

// 元素_注册事件C1, 注册事件C1方式, 省略1参数.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func X元素_注册事件C1(元素句柄 int, 事件类型 炫彩常量类.XE_, 事件函数 interface{}) bool {
	r, _, _ := xEle_RegEventC1.Call(uintptr(元素句柄), uintptr(事件类型), syscall.NewCallback(事件函数))
	return r != 0
}

// 元素_移除事件C.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.
func X元素_移除事件C(元素句柄 int, 事件类型 炫彩常量类.XE_, 事件函数 interface{}) bool {
	r, _, _ := xEle_RemoveEventC.Call(uintptr(元素句柄), uintptr(事件类型), syscall.NewCallback(事件函数))
	return r != 0
}

// 元素_注册事件CEx, 注册事件C方式, 省略2参数, 和非Ex版相比只是最后一个参数不同.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func X元素_注册事件CEx(元素句柄 int, 事件类型 炫彩常量类.XE_, 事件函数指针 uintptr) bool {
	r, _, _ := xEle_RegEventC.Call(uintptr(元素句柄), uintptr(事件类型), 事件函数指针)
	return r != 0
}

// 元素_注册事件C1Ex, 注册事件C1方式, 省略1参数, 和非Ex版相比只是最后一个参数不同.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func X元素_注册事件C1Ex(元素句柄 int, 事件类型 炫彩常量类.XE_, 事件函数指针 uintptr) bool {
	r, _, _ := xEle_RegEventC1.Call(uintptr(元素句柄), uintptr(事件类型), 事件函数指针)
	return r != 0
}

// 元素_移除事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.
func X元素_移除事件CEx(元素句柄 int, 事件类型 炫彩常量类.XE_, 事件函数指针 uintptr) bool {
	r, _, _ := xEle_RemoveEventC.Call(uintptr(元素句柄), uintptr(事件类型), 事件函数指针)
	return r != 0
}

// 元素_发送事件.
//
// hEle: 目标元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func X元素_发送事件(元素句柄 int, 事件类型 炫彩常量类.XE_, 参数1, 参数2 uint) int {
	r, _, _ := xEle_SendEvent.Call(uintptr(元素句柄), uintptr(事件类型), uintptr(参数1), uintptr(参数2))
	return int(r)
}

// 元素_投递事件.
//
// hEle: 元素句柄.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func X元素_投递事件(元素句柄 int, 事件类型 炫彩常量类.XE_, 参数1, 参数2 uint) int {
	r, _, _ := xEle_PostEvent.Call(uintptr(元素句柄), uintptr(事件类型), uintptr(参数1), uintptr(参数2))
	return int(r)
}

// 元素_取坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func X元素_取坐标(元素句柄 int, 坐标 *RECT) int {
	r, _, _ := xEle_GetRect.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 元素_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func X元素_取逻辑坐标(元素句柄 int, 坐标 *RECT) int {
	r, _, _ := xEle_GetRectLogic.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 元素_取客户区坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func X元素_取客户区坐标(元素句柄 int, 坐标 *RECT) int {
	r, _, _ := xEle_GetClientRect.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 元素_置宽度.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
func X元素_置宽度(元素句柄 int, 宽度 int) int {
	r, _, _ := xEle_SetWidth.Call(uintptr(元素句柄), uintptr(宽度))
	return int(r)
}

// 元素_置高度.
//
// hEle: 元素句柄.
//
// nHeight: 高度.
func X元素_置高度(元素句柄 int, 高度 int) int {
	r, _, _ := xEle_SetHeight.Call(uintptr(元素句柄), uintptr(高度))
	return int(r)
}

// 元素_取宽度.
//
// hEle: 元素句柄.
func X元素_取宽度(元素句柄 int) int {
	r, _, _ := xEle_GetWidth.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_取高度.
//
// hEle: 元素句柄.
func X元素_取高度(元素句柄 int) int {
	r, _, _ := xEle_GetHeight.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_窗口客户区坐标到元素客户区坐标, 窗口客户区坐标转换到元素客户区坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func X元素_窗口客户区坐标到元素客户区坐标(元素句柄 int, 坐标 *RECT) int {
	r, _, _ := xEle_RectWndClientToEleClient.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 元素_窗口客户区点到元素客户区, 窗口客户区坐标转换到元素客户区坐标.
//
// hEle: 元素句柄.
//
// pPt: 坐标.
func X元素_窗口客户区点到元素客户区(元素句柄 int, 坐标 *POINT) int {
	r, _, _ := xEle_PointWndClientToEleClient.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 元素_客户区坐标到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func X元素_客户区坐标到窗口客户区(元素句柄 int, 坐标 *RECT) int {
	r, _, _ := xEle_RectClientToWndClient.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 元素_客户区点到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// hEle: 元素句柄.
//
// pPt: 坐标.
func X元素_客户区点到窗口客户区(元素句柄 int, 坐标 *POINT) int {
	r, _, _ := xEle_PointClientToWndClient.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 元素_基于窗口客户区坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func X元素_基于窗口客户区坐标(元素句柄 int, 坐标 *RECT) int {
	r, _, _ := xEle_GetWndClientRect.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 元素_取光标, 获取元素鼠标光标, 返回光标句柄.
//
// hEle: 元素句柄.
func X元素_取光标(元素句柄 int) int {
	r, _, _ := xEle_GetCursor.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_置光标, 设置元素鼠标光标.
//
// hEle: 元素句柄.
//
// hCursor: 光标句柄.
func X元素_置光标(元素句柄 int, 光标句柄 int) int {
	r, _, _ := xEle_SetCursor.Call(uintptr(元素句柄), uintptr(光标句柄))
	return int(r)
}

// 元素_添加子对象.
//
// hEle: 元素句柄.
//
// hChild: 要添加的子元素句柄或形状对象句柄.
func X元素_添加子对象(元素句柄 int, 要添加的子元素句柄或形状对象句柄 int) bool {
	r, _, _ := xEle_AddChild.Call(uintptr(元素句柄), uintptr(要添加的子元素句柄或形状对象句柄))
	return r != 0
}

// 元素_插入子对象, 插入子对象到指定位置.
//
// hEle: 元素句柄.
//
// hChild: 要插入的元素句柄或形状对象句柄.
//
// index: 插入位置索引.
func X元素_插入子对象(元素句柄 int, 要插入的元素句柄或形状对象句柄 int, 插入位置索引 int) bool {
	r, _, _ := xEle_InsertChild.Call(uintptr(元素句柄), uintptr(要插入的元素句柄或形状对象句柄), uintptr(插入位置索引))
	return r != 0
}

// 元素_置坐标, 如果返回0坐标没有改变, 如果大小改变返回2(触发XE_SIZE), 否则返回1(仅改变left,top,没有改变大小).
//
// hEle: 元素句柄.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func X元素_置坐标(元素句柄 int, 坐标 *RECT, 是否重绘 bool, 调整布局标识位 炫彩常量类.AdjustLayout_, 调整布局流水号 uint32) int {
	r, _, _ := xEle_SetRect.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.BoolPtr(是否重绘), uintptr(调整布局标识位), uintptr(调整布局流水号))
	return int(r)
}

// 元素_置坐标扩展, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// hEle: 元素句柄.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func X元素_置坐标EX(元素句柄 int, X坐标 int, Y坐标 int, 宽度 int, 高度 int, 是否重绘 bool, 调整布局标识位 炫彩常量类.AdjustLayout_, 调整布局流水号 uint32) int {
	r, _, _ := xEle_SetRectEx.Call(uintptr(元素句柄), uintptr(X坐标), uintptr(Y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.BoolPtr(是否重绘), uintptr(调整布局标识位), uintptr(调整布局流水号))
	return int(r)
}

// 元素_置逻辑坐标, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_ , 此参数将会传入XE_SIZE ,XE_ADJUSTLAYOUT 事件回调.
//
// nAdjustNo: 调整布局流水号, 可填0.
func X元素_置逻辑坐标(元素句柄 int, 坐标 *RECT, 是否重绘 bool, 调整布局标识位 炫彩常量类.AdjustLayout_, 调整布局流水号 uint32) int {
	r, _, _ := xEle_SetRectLogic.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.BoolPtr(是否重绘), uintptr(调整布局标识位), uintptr(调整布局流水号))
	return int(r)
}

// 元素_移动, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// hEle: 元素句柄.
//
// x: X坐标.
//
// y: Y坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func X元素_移动(元素句柄 int, X坐标, Y坐标 int32, 是否重绘 bool, 调整布局标识位 炫彩常量类.AdjustLayout_, 调整布局流水号 uint32) int {
	r, _, _ := xEle_SetPosition.Call(uintptr(元素句柄), uintptr(X坐标), uintptr(Y坐标), 炫彩工具类.BoolPtr(是否重绘), uintptr(调整布局标识位), uintptr(调整布局流水号))
	return int(r)
}

// 元素_移动逻辑坐标, 移动元素坐标, 逻辑坐标, 包含滚动视图偏移. 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// hEle: 元素句柄.
//
// x: X坐标.
//
// y: Y坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func X元素_移动逻辑坐标(元素句柄 int, X坐标, Y坐标 int32, 是否重绘 bool, 调整布局标识位 炫彩常量类.AdjustLayout_, 调整布局流水号 uint32) int {
	r, _, _ := xEle_SetPositionLogic.Call(uintptr(元素句柄), uintptr(X坐标), uintptr(Y坐标), 炫彩工具类.BoolPtr(是否重绘), uintptr(调整布局标识位), uintptr(调整布局流水号))
	return int(r)
}

// 元素_判断绘制焦点.
//
// hEle: 元素句柄.
func X元素_判断绘制焦点(元素句柄 int) bool {
	r, _, _ := xEle_IsDrawFocus.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断启用, 元素是否为启用状态.
//
// hEle: 元素句柄.
func X元素_判断启用(元素句柄 int) bool {
	r, _, _ := xEle_IsEnable.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断启用焦点, 元素是否启用焦点.
//
// hEle: 元素句柄.
func X元素_判断启用焦点(元素句柄 int) bool {
	r, _, _ := xEle_IsEnableFocus.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断鼠标穿透, 元素是否启用鼠标穿透.
//
// hEle: 元素句柄.
func X元素_判断鼠标穿透(元素句柄 int) bool {
	r, _, _ := xEle_IsMouseThrough.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_测试点击元素, 检测坐标点所在元素, 包含子元素的子元素.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
func X元素_测试点击元素(元素句柄 int, 坐标点 *POINT) int {
	r, _, _ := xEle_HitChildEle.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)))
	return int(r)
}

// 元素_判断背景透明.
//
// hEle: 元素句柄.
func X元素_判断背景透明(元素句柄 int) bool {
	r, _, _ := xEle_IsBkTransparent.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断启用事件_XE_PAINT_END, 是否启XE_PAINT_END用事件.
//
// hEle: 元素句柄.
func X元素_判断启用事件_XE_PAINT_END(元素句柄 int) bool {
	r, _, _ := xEle_IsEnableEvent_XE_PAINT_END.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断接受TAB, 是否接受Tab键输入; 例如: XRichEdit, XEdit.
//
// hEle: 元素句柄.
func X元素_判断接受TAB(元素句柄 int) bool {
	r, _, _ := xEle_IsKeyTab.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断接受切换焦点, 是否接受通过键盘切换焦点(方向键,TAB键).
//
// hEle: 元素句柄.
func X元素_判断接受切换焦点(元素句柄 int) bool {
	r, _, _ := xEle_IsSwitchFocus.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断启用_XE_MOUSEWHEEL, 判断是否启用鼠标滚动事件, 如果禁用那么事件会发送给他的父元素.
//
// hEle: 元素句柄.
func X元素_判断启用_XE_MOUSEWHEEL(元素句柄 int) bool {
	r, _, _ := xEle_IsEnable_XE_MOUSEWHEEL.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断为子元素, 判断hChildEle是否为hEle的子元素.
//
// hEle: 元素句柄.
//
// hChildEle: 子元素句柄.
func X元素_判断为子元素(元素句柄 int, 子元素句柄 int) bool {
	r, _, _ := xEle_IsChildEle.Call(uintptr(元素句柄), uintptr(子元素句柄))
	return r != 0
}

// 元素_判断启用画布, 判断是否启用画布.
//
// hEle: 元素句柄.
func X元素_判断启用画布(元素句柄 int) bool {
	r, _, _ := xEle_IsEnableCanvas.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断焦点, 判断是否拥有焦点.
//
// hEle: 元素句柄.
func X元素_判断焦点(元素句柄 int) bool {
	r, _, _ := xEle_IsFocus.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_判断焦点扩展, 判断该元素或该元素的子元素是否拥有焦点.
//
// hEle: 元素句柄.
func X元素_判断焦点EX(元素句柄 int) bool {
	r, _, _ := xEle_IsFocusEx.Call(uintptr(元素句柄))
	return r != 0
}

// 元素_启用, 启用或禁用元素.
//
// hEle: 元素句柄.
//
// bEnable: 启用或禁用.
func X元素_启用(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_Enable.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用焦点, 启用焦点.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用焦点(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableFocus.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用绘制焦点.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用绘制焦点(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableDrawFocus.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用绘制边框, 启用或禁用绘制默认边框.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用绘制边框(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableDrawBorder.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用画布, 启用或禁用背景画布; 如果禁用那么将绘制在父的画布之上, 也就是说他没有自己的画布.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用画布(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableCanvas.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用事件_XE_PAINT_END.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用事件_XE_PAINT_END(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableEvent_XE_PAINT_END.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用背景透明.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用背景透明(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableBkTransparent.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用鼠标穿透. 启用鼠标穿透, 如果启用, 那么该元素不能接收到鼠标事件, 但是他的子元素不受影响, 任然可以接收鼠标事件.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用鼠标穿透(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableMouseThrough.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用接收TAB, 启用接收Tab输入.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用接收TAB(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableKeyTab.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用切换焦点, 启用接受通过键盘切换焦点.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用切换焦点(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableSwitchFocus.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_启用事件_XE_MOUSEWHEEL, 启用接收鼠标滚动事件, 如果禁用那么事件会传递给父元素.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X元素_启用事件_XE_MOUSEWHEEL(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableEvent_XE_MOUSEWHEEL.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_移除, 移除元素, 但不销毁.
//
// hEle: 元素句柄.
func X元素_移除(元素句柄 int) int {
	r, _, _ := xEle_Remove.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_置Z序, 设置元素Z序.
//
// hEle: 元素句柄.
//
// index: 位置索引.
func X元素_置Z序(元素句柄 int, 位置索引 int) bool {
	r, _, _ := xEle_SetZOrder.Call(uintptr(元素句柄), uintptr(位置索引))
	return r != 0
}

// 元素_置Z序扩展, 设置元素Z序.
//
// hEle: 元素句柄.
//
// hDestEle: 目标元素.
//
// nType: 类型, Zorder_.
func X元素_置Z序EX(元素句柄 int, 目标元素 int, 类型 炫彩常量类.Zorder_) bool {
	r, _, _ := xEle_SetZOrderEx.Call(uintptr(元素句柄), uintptr(目标元素), uintptr(类型))
	return r != 0
}

// 元素_取Z序, 获取元素Z序索引, 位置索引.
//
// hEle: 元素句柄.
func X元素_取Z序(元素句柄 int) int {
	r, _, _ := xEle_GetZOrder.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_启用置顶, 设置元素置顶.
//
// hEle: 元素句柄.
//
// bTopmost: 是否置顶显示.
func X元素_启用置顶(元素句柄 int, 是否置顶显示 bool) bool {
	r, _, _ := xEle_EnableTopmost.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否置顶显示))
	return r != 0
}

// 元素_重绘.
//
// hEle: 元素句柄.
//
// bImmediate: 是否立即重绘.
func X元素_重绘(元素句柄 int, 是否立即重绘 bool) int {
	r, _, _ := xEle_Redraw.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否立即重绘))
	return int(r)
}

// 元素_重绘指定区域.
//
// hEle: 元素句柄.
//
// pRect: 相对于元素客户区坐标.
//
// bImmediate: 是否立即重绘.
func X元素_重绘指定区域(元素句柄 int, 相对于元素客户区坐标 *RECT, 是否立即重绘 bool) int {
	r, _, _ := xEle_RedrawRect.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(相对于元素客户区坐标)), 炫彩工具类.BoolPtr(是否立即重绘))
	return int(r)
}

// 元素_取子对象数量, 获取子对象(UI元素和形状对象)数量, 只检测当前层子对象.
//
// hEle: 元素句柄.
func X元素_取子对象数量(元素句柄 int) int {
	r, _, _ := xEle_GetChildCount.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_取子对象从索引, 获取子对象通过索引, 只检测当前层子对象.
//
// hEle: 元素句柄.
//
// index: 索引.
func X元素_取子对象从索引(元素句柄 int, 索引 int) int {
	r, _, _ := xEle_GetChildByIndex.Call(uintptr(元素句柄), uintptr(索引))
	return int(r)
}

// 元素_取子对象从ID, 获取子对象通过ID, 只检测当前层子对象.
//
// hEle: 元素句柄.
//
// nID: 元素ID.
func X元素_取子对象从ID(元素句柄 int, 元素ID int) int {
	r, _, _ := xEle_GetChildByID.Call(uintptr(元素句柄), uintptr(元素ID))
	return int(r)
}

// 元素_置边框大小.
//
// hEle: 元素句柄.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func X元素_置边框大小(元素句柄 int, 左 int, 上 int, 右 int, 下 int) int {
	r, _, _ := xEle_SetBorderSize.Call(uintptr(元素句柄), uintptr(左), uintptr(上), uintptr(右), uintptr(下))
	return int(r)
}

// 元素_取边框大小.
//
// hEle: 元素句柄.
//
// pBorder: 大小.
func X元素_取边框大小(元素句柄 int, 大小 *RECT) int {
	r, _, _ := xEle_GetBorderSize.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(大小)))
	return int(r)
}

// 元素_置内填充大小.
//
// hEle: 元素句柄.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func X元素_置内填充大小(元素句柄 int, 左 int, 上 int, 右 int, 下 int) int {
	r, _, _ := xEle_SetPadding.Call(uintptr(元素句柄), uintptr(左), uintptr(上), uintptr(右), uintptr(下))
	return int(r)
}

// 元素_取内填充大小.
//
// hEle: 元素句柄.
//
// pPadding: 大小.
func X元素_取内填充大小(元素句柄 int, 大小 *RECT) int {
	r, _, _ := xEle_GetPadding.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(大小)))
	return int(r)
}

// 元素_置拖动边框.
//
// hEle: 元素句柄.
//
// nFlags: 边框位置组合, Element_Position_.
func X元素_置拖动边框(元素句柄 int, 边框位置组合 炫彩常量类.Element_Position_) int {
	r, _, _ := xEle_SetDragBorder.Call(uintptr(元素句柄), uintptr(边框位置组合))
	return int(r)
}

// 元素_置拖动边框绑定元素, 设置拖动边框绑定元素, 当拖动边框时, 自动调整绑定元素的大小.
//
// hEle: 元素句柄.
//
// nFlags: 边框位置标识, Element_Position_.
//
// hBindEle: 绑定元素.
//
// nSpace: 元素间隔大小.
func X元素_置拖动边框绑定元素(元素句柄 int, 边框位置标识 炫彩常量类.Element_Position_, 绑定元素 int, 元素间隔大小 int) int {
	r, _, _ := xEle_SetDragBorderBindEle.Call(uintptr(元素句柄), uintptr(边框位置标识), uintptr(绑定元素), uintptr(元素间隔大小))
	return int(r)
}

// 元素_置最小大小.
//
// hEle: 元素句柄.
//
// nWidth: 最小宽度.
//
// nHeight: 最小高度.
func X元素_置最小大小(元素句柄 int, 最小宽度 int, 最小高度 int) int {
	r, _, _ := xEle_SetMinSize.Call(uintptr(元素句柄), uintptr(最小宽度), uintptr(最小高度))
	return int(r)
}

// 元素_置最大大小.
//
// hEle: 元素句柄.
//
// nWidth: 最大宽度.
//
// nHeight: 最大高度.
func X元素_置最大大小(元素句柄 int, 最大宽度 int, 最大高度 int) int {
	r, _, _ := xEle_SetMaxSize.Call(uintptr(元素句柄), uintptr(最大宽度), uintptr(最大高度))
	return int(r)
}

// 元素_置锁定滚动, 设置锁定元素在滚动视图中跟随滚动, 如果设置TRUE将不跟随滚动.
//
// hEle: 元素句柄.
//
// bHorizon: 是否锁定水平滚动.
//
// bVertical: 是否锁定垂直滚动.
func X元素_置锁定滚动(元素句柄 int, 是否锁定水平滚动 bool, 是否锁定垂直滚动 bool) int {
	r, _, _ := xEle_SetLockScroll.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否锁定水平滚动), 炫彩工具类.BoolPtr(是否锁定垂直滚动))
	return int(r)
}

// 元素_置文本颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func X元素_置文本颜色(元素句柄 int, ABGR颜色值 int) int {
	r, _, _ := xEle_SetTextColor.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 元素_取文本颜色.
//
// hEle: 元素句柄.
func X元素_取文本颜色(元素句柄 int) int {
	r, _, _ := xEle_GetTextColor.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_取文本颜色扩展, 获取文本颜色, 优先从资源中获取.
//
// hEle: 元素句柄.
func X元素_取文本颜色EX(元素句柄 int) int {
	r, _, _ := xEle_GetTextColorEx.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_置焦点边框颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func X元素_置焦点边框颜色(元素句柄 int, ABGR颜色值 int) int {
	r, _, _ := xEle_SetFocusBorderColor.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 元素_取焦点边框颜色.
//
// hEle: 元素句柄.
func X元素_取焦点边框颜色(元素句柄 int) int {
	r, _, _ := xEle_GetFocusBorderColor.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_置字体.
//
// hEle: 元素句柄.
//
// hFontx: 炫彩字体.
func X元素_置字体(元素句柄 int, 炫彩字体 int) int {
	r, _, _ := xEle_SetFont.Call(uintptr(元素句柄), uintptr(炫彩字体))
	return int(r)
}

// 元素_取字体.
//
// hEle: 元素句柄.
func X元素_取字体(元素句柄 int) int {
	r, _, _ := xEle_GetFont.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_取字体扩展, 获取元素字体, 优先从资源中获取.
//
// hEle: 元素句柄.
func X元素_取字体EX(元素句柄 int) int {
	r, _, _ := xEle_GetFontEx.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_置透明度.
//
// hEle: 元素句柄.
func X元素_置透明度(元素句柄 int, alpha uint8) int {
	r, _, _ := xEle_SetAlpha.Call(uintptr(元素句柄), uintptr(alpha))
	return int(r)
}

// 元素_销毁.
//
// hEle: 元素句柄.
func X元素_销毁(元素句柄 int) int {
	r, _, _ := xEle_Destroy.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_添加背景边框, 添加背景内容边框.
//
// hEle: 元素句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
func X元素_添加背景边框(元素句柄 int, 组合状态 炫彩常量类.CombinedState, ABGR颜色 int, 线宽 int) int {
	r, _, _ := xEle_AddBkBorder.Call(uintptr(元素句柄), uintptr(组合状态), uintptr(ABGR颜色), uintptr(线宽))
	return int(r)
}

// 元素_添加背景填充, 添加背景内容填充.
//
// hEle: 元素句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
func X元素_添加背景填充(元素句柄 int, 组合状态 炫彩常量类.CombinedState, ABGR颜色 int) int {
	r, _, _ := xEle_AddBkFill.Call(uintptr(元素句柄), uintptr(组合状态), uintptr(ABGR颜色))
	return int(r)
}

// 元素_添加背景图片, 添加背景内容图片.
//
// hEle: 元素句柄.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func X元素_添加背景图片(元素句柄 int, 组合状态 炫彩常量类.CombinedState, 图片句柄 int) int {
	r, _, _ := xEle_AddBkImage.Call(uintptr(元素句柄), uintptr(组合状态), uintptr(图片句柄))
	return int(r)
}

// 元素_取背景对象数量, 获取背景内容数量.
//
// hEle: 元素句柄.
func X元素_取背景对象数量(元素句柄 int) int {
	r, _, _ := xEle_GetBkInfoCount.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_清空背景对象, 清空背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.
//
// hEle: 元素句柄.
func X元素_清空背景对象(元素句柄 int) int {
	r, _, _ := xEle_ClearBkInfo.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_取背景管理器, 获取元素背景管理器.
//
// hEle: 元素句柄.
func X元素_取背景管理器(元素句柄 int) int {
	r, _, _ := xEle_GetBkManager.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_取背景管理器扩展, 获取元素背景管理器, 优先从资源中获取.
//
// hEle: 元素句柄.
func X元素_取背景管理器EX(元素句柄 int) int {
	r, _, _ := xEle_GetBkManagerEx.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_置背景管理器.
//
// hEle: 元素句柄.
//
// hBkInfoM: 背景管理器.
func X元素_置背景管理器(元素句柄 int, 背景管理器 int) int {
	r, _, _ := xEle_SetBkManager.Call(uintptr(元素句柄), uintptr(背景管理器))
	return int(r)
}

// 元素_取状态, 获取组合状态.
//
// hEle: 元素句柄.
func X元素_取状态(元素句柄 int) 炫彩常量类.CombinedState {
	r, _, _ := xEle_GetStateFlags.Call(uintptr(元素句柄))
	return 炫彩常量类.CombinedState(r)
}

// 元素_绘制焦点, 绘制元素焦点.
//
// hEle: 元素句柄.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func X元素_绘制焦点(元素句柄 int, 图形绘制句柄 int, 区域坐标 *RECT) bool {
	r, _, _ := xEle_DrawFocus.Call(uintptr(元素句柄), uintptr(图形绘制句柄), uintptr(unsafe.Pointer(区域坐标)))
	return r != 0
}

// 元素_绘制, 在自绘事件函数中, 用户手动调用绘制元素, 以便控制绘制顺序.
//
// hEle: 元素句柄.
//
// hDraw: 图形绘制句柄.
func X元素_绘制(元素句柄 int, 图形绘制句柄 int) int {
	r, _, _ := xEle_DrawEle.Call(uintptr(元素句柄), uintptr(图形绘制句柄))
	return int(r)
}

// 元素_置用户数据.
//
// hEle: 元素句柄.
//
// nData: 用户数据.
func X元素_置用户数据(元素句柄 int, 用户数据 int) int {
	r, _, _ := xEle_SetUserData.Call(uintptr(元素句柄), uintptr(用户数据))
	return int(r)
}

// 元素_取用户数据.
//
// hEle: 元素句柄.
func X元素_取用户数据(元素句柄 int) int {
	r, _, _ := xEle_GetUserData.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_取内容大小.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直, 布局属性交换依赖.
//
// cx: 宽度.
//
// cy: 高度.
//
// pSize: 返回大小.
func X元素_取内容大小(元素句柄 int, 水平或垂直 bool, 宽度 int, 高度 int, 返回大小 *SIZE) int {
	r, _, _ := xEle_GetContentSize.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(水平或垂直), uintptr(宽度), uintptr(高度), uintptr(unsafe.Pointer(返回大小)))
	return int(r)
}

// 元素_置鼠标捕获.
//
// hEle: 元素句柄.
//
// b: TRUE设置.
func X元素_置鼠标捕获(元素句柄 int, 启用或关闭 bool) int {
	r, _, _ := xEle_SetCapture.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(启用或关闭))
	return int(r)
}

// 元素_启用透明通道, 启用或关闭元素透明通道, 如果启用, 将强制设置元素背景不透明, 默认为启用, 此功能是为了兼容GDI不支持透明通道问题.
//
// hEle: 元素句柄.
//
// bEnable: 启用或关闭.
func X元素_启用透明通道(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEle_EnableTransparentChannel.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 元素_置炫彩定时器, 设置元素定时器.
//
// hEle: 元素句柄.
//
// nIDEvent: 事件ID.
//
// uElapse: 延时毫秒.
func X元素_置炫彩定时器(元素句柄 int, 事件ID int, 延时毫秒 int) bool {
	r, _, _ := xEle_SetXCTimer.Call(uintptr(元素句柄), uintptr(事件ID), uintptr(延时毫秒))
	return r != 0
}

// 元素_关闭炫彩定时器, 关闭元素定时器.
//
// hEle: 元素句柄.
//
// nIDEvent: 事件ID.
func X元素_关闭炫彩定时器(元素句柄 int, 事件ID int) bool {
	r, _, _ := xEle_KillXCTimer.Call(uintptr(元素句柄), uintptr(事件ID))
	return r != 0
}

// 元素_置工具提示, 设置工具提示内容.
//
// hEle: 元素句柄.
//
// pText: 工具提示内容.
func X元素_置工具提示(元素句柄 int, 工具提示内容 string) int {
	r, _, _ := xEle_SetToolTip.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(工具提示内容))
	return int(r)
}

// 元素_置工具提示扩展, 设置工具提示内容.
//
// hEle: 元素句柄.
//
// pText: 工具提示内容.
//
// nTextAlign: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func X元素_置工具提示EX(元素句柄 int, 工具提示内容 string, 文本对齐方式 炫彩常量类.TextFormatFlag_) int {
	r, _, _ := xEle_SetToolTipEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(工具提示内容), uintptr(文本对齐方式))
	return int(r)
}

// 元素_取工具提示, 获取工具提示内容.
//
// hEle: 元素句柄.
func X元素_取工具提示(元素句柄 int) int {
	r, _, _ := xEle_GetToolTip.Call(uintptr(元素句柄))
	return int(r)
}

// 元素_弹出工具提示, 弹出工具提示.
//
// hEle: 元素句柄.
//
// x: X坐标.
//
// y: Y坐标.
func X元素_弹出工具提示(元素句柄 int, X坐标 int, Y坐标 int) int {
	r, _, _ := xEle_PopupToolTip.Call(uintptr(元素句柄), uintptr(X坐标), uintptr(Y坐标))
	return int(r)
}

// 元素_调整布局.
//
// hEle: 元素句柄.
//
// nAdjustNo: 调整布局流水号, 可填0.
func X元素_调整布局(元素句柄 int, 调整布局流水号 uint32) int {
	r, _, _ := xEle_AdjustLayout.Call(uintptr(元素句柄), uintptr(调整布局流水号))
	return int(r)
}

// 元素_调整布局扩展.
//
// hEle: 元素句柄.
//
// nFlags: 调整标识.
//
// nAdjustNo: 调整布局流水号, 可填0.
func X元素_调整布局EX(元素句柄 int, 调整标识 炫彩常量类.AdjustLayout_, 调整布局流水号 uint32) int {
	r, _, _ := xEle_AdjustLayoutEx.Call(uintptr(元素句柄), uintptr(调整标识), uintptr(调整布局流水号))
	return int(r)
}

// 元素_取透明度, 返回透明度.
//
// hEle: 元素句柄.
func X元素_取透明度(元素句柄 int) byte {
	r, _, _ := xEle_GetAlpha.Call(uintptr(元素句柄))
	return byte(r)
}

// 元素_取位置.
//
// hEle: 元素句柄.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.
func X元素_取位置(元素句柄 int, 返回X坐标, 返回Y坐标 *int32) int {
	r, _, _ := xEle_GetPosition.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(返回X坐标)), uintptr(unsafe.Pointer(返回Y坐标)))
	return int(r)
}

// 元素_置大小.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func X元素_置大小(元素句柄 int, 宽度, 高度 int32, 是否重绘 bool, 调整布局标识位 炫彩常量类.AdjustLayout_, 调整布局流水号 uint32) int {
	r, _, _ := xEle_SetSize.Call(uintptr(元素句柄), uintptr(宽度), uintptr(高度), 炫彩工具类.BoolPtr(是否重绘), uintptr(调整布局标识位), uintptr(调整布局流水号))
	return int(r)
}

// 元素_取大小.
//
// hEle: 元素句柄.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.
func X元素_取大小(元素句柄 int, 返回宽度, 返回高度 *int32) int {
	r, _, _ := xEle_GetSize.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(返回宽度)), uintptr(unsafe.Pointer(返回高度)))
	return int(r)
}

// 元素_置背景, 设置背景内容, 返回设置的背景对象数量.
//
// hEle: 元素句柄.
//
// pText: 背景内容字符串.
func X元素_置背景(元素句柄 int, 背景内容字符串 string) int {
	r, _, _ := xEle_SetBkInfo.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(背景内容字符串))
	return int(r)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// hEle: 元素句柄.
//
// pRect: 接收返回坐标.
func X元素_取窗口客户区坐标DPI(元素句柄 int, 接收返回坐标 *RECT) int {
	r, _, _ := xEle_GetWndClientRectDPI.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(接收返回坐标)))
	return int(r)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// hEle: 元素句柄.
//
// pPt: 接收返回坐标点.
func X元素_取窗口客户区坐标点DPI(元素句柄 int, 接收返回坐标点 *POINT) int {
	r, _, _ := xEle_PointClientToWndClientDPI.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(接收返回坐标点)))
	return int(r)
}

// 元素_客户区坐标到窗口客户区DPI. 基于DPI缩放后的坐标.
//
// hEle: 元素句柄.
//
// pRect: 接收返回坐标.
func X元素_客户区坐标到窗口客户区DPI(元素句柄 int, 接收返回坐标 *RECT) int {
	r, _, _ := xEle_RectClientToWndClientDPI.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(接收返回坐标)))
	return int(r)
}
