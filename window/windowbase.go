package 炫彩窗口基类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/wapi/wnd"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// windowBase 窗口基类.
type windowBase struct {
	炫彩对象基类.UI
}

// MessageBox 炫彩_消息框.
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param XCStyle xcc.Window_Style_.
//	@return xcc.MessageBox_Flag_. 返回: xcc.MessageBox_Flag_Ok: 点击确定按钮退出. xcc.MessageBox_Flag_Cancel: 点击取消按钮退出. xcc.MessageBox_Flag_Other: 其他方式退出.
func (w *windowBase) X消息框(标题, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, 样式 炫彩常量类.Window_Style_) 炫彩常量类.MessageBox_Flag_ {
	return 炫彩基类.X消息框(标题, 内容文本, 标识, w.X取HWND(), 样式)
}

// Msg_Create 消息框_创建, 然后请调用 DoModal() 方法显示模态窗口.
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param XCStyle xcc.Window_Style_.
//	@return *ModalWindow 模态窗口对象.
func (w *windowBase) X消息框创建(pTitle, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, 样式 炫彩常量类.Window_Style_) *ModalWindow {
	p := &ModalWindow{}
	p.X设置句柄(炫彩基类.X消息框_创建(pTitle, 内容文本, 标识, w.X取HWND(), 样式))
	return p
}

// Msg_CreateEx 消息框_创建扩展, 然后请调用 DoModal() 方法显示模态窗口.
//
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param XCStyle xcc.Window_Style_.
//	@return *ModalWindow 模态窗口对象.
func (w *windowBase) X消息框创建EX(dwExStyle, 窗口样式 int, 窗口类名, 标题, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, xcc 炫彩常量类.Window_Style_) *ModalWindow {
	p := &ModalWindow{}
	p.X设置句柄(炫彩基类.X消息框_创建EX(dwExStyle, 窗口样式, 窗口类名, 标题, 内容文本, 标识, w.X取HWND(), xcc))
	return p
}

// 炫彩_发送窗口消息.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (w *windowBase) X发送窗口消息(消息值 uint32, 参数1, 参数2 uint) uint {
	return 炫彩基类.X发送窗口消息(w.Handle, 消息值, 参数1, 参数2)
}

// 炫彩_投递窗口消息.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (w *windowBase) X投递窗口消息(消息值 uint32, 参数1 int32, 参数2 int32) bool {
	return 炫彩基类.X投递窗口消息(w.Handle, 消息值, 参数1, 参数2)
}

// 炫彩_判断窗口, 判断是否为窗口句柄.
func (w *windowBase) X判断窗口() bool {
	return 炫彩基类.X判断窗口(w.Handle)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象.
//
// nID: ID值.
func (w *windowBase) X取对象并按ID(ID值 int) int {
	return 炫彩基类.X取对象从ID(w.Handle, ID值)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄.
//
// pName: ID名称.
func (w *windowBase) X取对象并按ID名称(ID名称 string) int {
	return 炫彩基类.X取对象从ID名称(w.Handle, ID名称)
}

// 窗口_显示.
//
// nCmdShow: 显示方式: xcc.SW_.
func (w *windowBase) X显示方式(显示方式 炫彩常量类.SW_) int {
	return 炫彩基类.X窗口_显示方式(w.Handle, 显示方式)
}

// 窗口_置顶.
func (w *windowBase) X置顶() int {
	return 炫彩基类.X窗口_置顶(w.Handle)
}

// 窗口_注册事件C.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) X注册事件C(事件类型 炫彩常量类.WM_, 事件函数 interface{}) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 事件类型, 事件函数)
}

// 窗口_注册事件C1.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) X注册事件C1(事件类型 炫彩常量类.WM_, 事件函数 interface{}) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 事件类型, 事件函数)
}

// 窗口_移除事件C.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) X移除事件C(事件类型 炫彩常量类.WM_, 事件函数 interface{}) bool {
	return 炫彩基类.X窗口_移除事件C(w.Handle, 事件类型, 事件函数)
}

// 窗口_移除事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成..
func (w *windowBase) X移除事件CEx(事件类型 炫彩常量类.WM_, 事件函数指针 uintptr) bool {
	return 炫彩基类.X窗口_移除事件CEx(w.Handle, 事件类型, 事件函数指针)
}

// 窗口_添加子对象.
//
// hChild: 要添加的对象句柄.
func (w *windowBase) X添加子对象(要添加的对象句柄 int) bool {
	return 炫彩基类.X窗口_添加子对象(w.Handle, 要添加的对象句柄)
}

// 窗口_插入子对象.
//
// hChild: 要插入的对象句柄.
//
// index: 插入位置索引.
func (w *windowBase) X插入子对象(要插入的对象句柄 int, 插入位置索引 int) bool {
	return 炫彩基类.X窗口_插入子对象(w.Handle, 要插入的对象句柄, 插入位置索引)
}

// 窗口_取HWND.
func (w *windowBase) X取HWND() uintptr {
	return 炫彩基类.X窗口_取HWND(w.Handle)
}

// 窗口_重绘.
//
// bImmediate: 是否立即重绘, 默认为否.
func (w *windowBase) X重绘(是否立即重绘 bool) int {
	return 炫彩基类.X窗口_重绘(w.Handle, 是否立即重绘)
}

// 窗口_重绘指定区域.
//
// pRect: 需要重绘的区域坐标.
//
// bImmediate: TRUE立即重绘, FALSE放入消息队列延迟重绘.
func (w *windowBase) X重绘指定区域(需要重绘的区域坐标 *炫彩基类.RECT, TRUE立即重绘 bool) int {
	return 炫彩基类.X窗口_重绘指定区域(w.Handle, 需要重绘的区域坐标, TRUE立即重绘)
}

// 窗口_置坐标.
//
// pRect: 坐标.
func (w *windowBase) X置坐标(坐标 *炫彩基类.RECT) bool {
	return 炫彩基类.X窗口_置坐标(w.Handle, 坐标)
}

// 窗口_置焦点.
//
// hFocusEle: 将获得焦点的元素.
func (w *windowBase) X置焦点(元素 int) bool {
	return 炫彩基类.X窗口_置焦点(w.Handle, 元素)
}

// 窗口_取焦点.
func (w *windowBase) X取焦点() int {
	return 炫彩基类.X窗口_取焦点(w.Handle)
}

// 窗口_取鼠标停留元素.
func (w *windowBase) X取鼠标停留元素() int {
	return 炫彩基类.X窗口_取鼠标停留元素(w.Handle)
}

// 窗口_绘制, 在自绘事件函数中,用户手动调用绘制窗口, 以便控制绘制顺序.
//
// hDraw: 图形绘制句柄.
func (w *windowBase) X绘制(图形绘制句柄 int) int {
	return 炫彩基类.X窗口_绘制(w.Handle, 图形绘制句柄)
}

// 窗口_居中.
func (w *windowBase) X居中() int {
	return 炫彩基类.X窗口_居中(w.Handle)
}

// 窗口_居中扩展.
//
// width: 窗口宽度.
//
// height: 窗口高度.
func (w *windowBase) X居中EX(窗口宽度, 窗口高度 int) int {
	return 炫彩基类.X窗口_居中EX(w.Handle, 窗口宽度, 窗口高度)
}

// 窗口_置光标.
//
// hCursor: 鼠标光标句柄.
func (w *windowBase) X置光标(鼠标光标句柄 int) int {
	return 炫彩基类.X窗口_置光标(w.Handle, 鼠标光标句柄)
}

// 窗口_取光标.
func (w *windowBase) X取光标() int {
	return 炫彩基类.X窗口_取光标(w.Handle)
}

// 窗口_启用拖动边框.
//
// bEnable: 是否启用.
func (w *windowBase) X启用拖动边框(是否启用 bool) int {
	return 炫彩基类.X窗口_启用拖动边框(w.Handle, 是否启用)
}

// 窗口_启用拖动窗口.
//
// bEnable: 是否启用.
func (w *windowBase) X启用拖动窗口(是否启用 bool) int {
	return 炫彩基类.X窗口_启用拖动窗口(w.Handle, 是否启用)
}

// 窗口_启用拖动标题栏.
//
// bEnable: 是否启用.
func (w *windowBase) X启用拖动标题栏(是否启用 bool) int {
	return 炫彩基类.X窗口_启用拖动标题栏(w.Handle, 是否启用)
}

// 窗口_启用绘制背景.
//
// bEnable: 是否启用.
func (w *windowBase) X启用绘制背景(是否启用 bool) int {
	return 炫彩基类.X窗口_启用绘制背景(w.Handle, 是否启用)
}

// 窗口_启用自动焦点.
//
// bEnable: 是否启用.
func (w *windowBase) X启用自动焦点(是否启用 bool) int {
	return 炫彩基类.X窗口_启用自动焦点(w.Handle, 是否启用)
}

// 窗口_启用允许最大化.
//
// bEnable: 是否启用.
func (w *windowBase) X启用允许最大化(是否启用 bool) int {
	return 炫彩基类.X窗口_启用允许最大化(w.Handle, 是否启用)
}

// 窗口_启用限制窗口大小.
//
// bEnable: 是否启用.
func (w *windowBase) X启用限制窗口大小(是否启用 bool) int {
	return 炫彩基类.X窗口_启用限制窗口大小(w.Handle, 是否启用)
}

// 窗口_启用布局.
//
// bEnable: 是否启用.
func (w *windowBase) X启用布局(是否启用 bool) int {
	return 炫彩基类.X窗口_启用布局(w.Handle, 是否启用)
}

// 窗口_启用布局覆盖边框.
//
// bEnable: 是否启用.
func (w *windowBase) X启用布局覆盖边框(是否启用 bool) int {
	return 炫彩基类.X窗口_启用布局覆盖边框(w.Handle, 是否启用)
}

// 窗口_显示布局边界.
//
// bEnable: 是否启用.
func (w *windowBase) X显示布局边界(是否启用 bool) int {
	return 炫彩基类.X窗口_显示布局边界(w.Handle, 是否启用)
}

// 窗口_判断启用布局.
func (w *windowBase) X判断启用布局() bool {
	return 炫彩基类.X窗口_判断启用布局(w.Handle)
}

// 窗口_是否最大化.
func (w *windowBase) X是否最大化() bool {
	return 炫彩基类.X窗口_是否最大化(w.Handle)
}

// 窗口_置鼠标捕获元素.
//
// hEle: 元素句柄.
func (w *windowBase) X置鼠标捕获元素(元素句柄 int) int {
	return 炫彩基类.X窗口_置鼠标捕获元素(w.Handle, 元素句柄)
}

// 窗口_取鼠标捕获元素.
func (w *windowBase) X取鼠标捕获元素() int {
	return 炫彩基类.X窗口_取鼠标捕获元素(w.Handle)
}

// 窗口_取绘制矩形.
//
// pRcPaint: 重绘区域坐标.
func (w *windowBase) X取绘制矩形(重绘区域坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X窗口_取绘制矩形(w.Handle, 重绘区域坐标)
}

// 窗口_置系统光标.
//
// hCursor: 光标句柄.
func (w *windowBase) X置系统光标(光标句柄 int) int {
	return 炫彩基类.X窗口_置系统光标(w.Handle, 光标句柄)
}

// 窗口_置字体.
//
// hFontx: 炫彩字体句柄.
func (w *windowBase) X置字体(炫彩字体句柄 int) int {
	return 炫彩基类.X窗口_置字体(w.Handle, 炫彩字体句柄)
}

// 窗口_置文本颜色.
//
// color: ABGR 颜色值.
func (w *windowBase) X置文本颜色(ABGR颜色值 int) int {
	return 炫彩基类.X窗口_置文本颜色(w.Handle, ABGR颜色值)
}

// 窗口_取文本颜色, 返回ABGR 颜色.
func (w *windowBase) X取文本颜色() int {
	return 炫彩基类.X窗口_取文本颜色(w.Handle)
}

// 窗口_取文本颜色扩展, 返回ABGR 颜色.
func (w *windowBase) X取文本颜色EX() int {
	return 炫彩基类.X窗口_取文本颜色EX(w.Handle)
}

// 窗口_置ID.
//
// nID: ID值.
func (w *windowBase) X置ID(ID值 int) int {
	return 炫彩基类.X窗口_置ID(w.Handle, ID值)
}

// 窗口_取ID.
func (w *windowBase) X取ID() int {
	return 炫彩基类.X窗口_取ID(w.Handle)
}

// 窗口_置名称.
//
// pName: name值, 字符串.
func (w *windowBase) X置名称(名称 string) int {
	return 炫彩基类.X窗口_置名称(w.Handle, 名称)
}

// 窗口_取名称.
func (w *windowBase) X取名称() string {
	return 炫彩基类.X窗口_取名称(w.Handle)
}

// 窗口_置边大小.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底部大小.
func (w *windowBase) X置边大小(左边, 上边, 右边, 底部 int) int {
	return 炫彩基类.X窗口_置边大小(w.Handle, 左边, 上边, 右边, 底部)
}

// 窗口_取边大小.
func (w *windowBase) X取边大小(pBorder *炫彩基类.RECT) int {
	return 炫彩基类.X窗口_取边大小(w.Handle, pBorder)
}

// 窗口_置布局内填充大小.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func (w *windowBase) X置布局内填充大小(左边, 上边, 右边, 下边 int) int {
	return 炫彩基类.X窗口_置布局内填充大小(w.Handle, 左边, 上边, 右边, 下边)
}

// 窗口_置拖动边框大小.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底边大小.
func (w *windowBase) X置拖动边框大小(左边, 上边, 右边, 底边 int) int {
	return 炫彩基类.X窗口_置拖动边框大小(w.Handle, 左边, 上边, 右边, 底边)
}

// 窗口_取拖动边框大小.
//
// pSize: 拖动边框大小.
func (w *windowBase) X取拖动边框大小(pBorder *炫彩基类.RECT) int {
	return 炫彩基类.X窗口_取拖动边框大小(w.Handle, pBorder)
}

// 窗口_置最小大小.
//
// width: 最小宽度.
//
// height: 最小高度.
func (w *windowBase) X置最小大小(最小宽度, 最小高度 int) int {
	return 炫彩基类.X窗口_置最小大小(w.Handle, 最小宽度, 最小高度)
}

// 窗口_测试点击子元素.
//
// pPt: 左边点.
func (w *windowBase) X测试点击子元素(左边点 *炫彩基类.POINT) int {
	return 炫彩基类.X窗口_测试点击子元素(w.Handle, 左边点)
}

// 窗口_取子对象数量.
func (w *windowBase) X取子对象数量() int {
	return 炫彩基类.X窗口_取子对象数量(w.Handle)
}

// 窗口_取子对象从索引.
//
// index: 元素索引.
func (w *windowBase) X取子对象并按索引(元素索引 int) int {
	return 炫彩基类.X窗口_取子对象从索引(w.Handle, 元素索引)
}

// 窗口_取子对象从ID.
//
// nID: 元素ID, ID必须大于0.
func (w *windowBase) X取子对象并按ID(元素ID int) int {
	return 炫彩基类.X窗口_取子对象从ID(w.Handle, 元素ID)
}

// 窗口_取子对象.
//
// nID: 对象ID,ID必须大于0.
func (w *windowBase) X取子对象(对象ID int) int {
	return 炫彩基类.X窗口_取子对象(w.Handle, 对象ID)
}

// 窗口_关闭.
func (w *windowBase) X关闭() int {
	return 炫彩基类.X窗口_关闭(w.Handle)
}

// 窗口_调整布局.
func (w *windowBase) X调整布局() int {
	return 炫彩基类.X窗口_调整布局(w.Handle)
}

// 窗口_调整布局扩展.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
func (w *windowBase) X调整布局EX(调整布局标识位 炫彩常量类.AdjustLayout_) int {
	return 炫彩基类.X窗口_调整布局EX(w.Handle, 调整布局标识位)
}

// 窗口_创建插入符.
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
func (w *windowBase) X创建插入符(元素句柄, x坐标, y坐标, 宽度, 高度 int) int {
	return 炫彩基类.X窗口_创建插入符(w.Handle, 元素句柄, x坐标, y坐标, 宽度, 高度)
}

// 窗口_置插入符位置, 设置插入符位置.
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
func (w *windowBase) X置插入符位置(x坐标, y坐标, 宽度, 高度 int, 是否立即更新UI bool) int {
	return 炫彩基类.X窗口_置插入符位置(w.Handle, x坐标, y坐标, 宽度, 高度, 是否立即更新UI)
}

// 窗口_置插入符颜色.
//
// color: 颜色值, ABGR 颜色.
func (w *windowBase) X置插入符颜色(颜色值 int) int {
	return 炫彩基类.X窗口_置插入符颜色(w.Handle, 颜色值)
}

// 窗口_显示插入符.
//
// bShow: 是否显示.
func (w *windowBase) X显示插入符(是否显示 bool) int {
	return 炫彩基类.X窗口_显示插入符(w.Handle, 是否显示)
}

// 窗口_销毁插入符.
func (w *windowBase) X销毁插入符() int {
	return 炫彩基类.X窗口_销毁插入符(w.Handle)
}

// 窗口_取插入符元素.
func (w *windowBase) X取插入符元素() int {
	return 炫彩基类.X窗口_取插入符元素(w.Handle)
}

// 窗口_取客户区坐标.
//
// pRect: 坐标.
func (w *windowBase) X取客户区坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X窗口_取客户区坐标(w.Handle, 坐标)
}

// 窗口_取Body坐标.
//
// pRect: 坐标.
func (w *windowBase) X取Body坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X窗口_取Body坐标(w.Handle, 坐标)
}

// 窗口_取布局坐标.
//
// pRect: 接收返回坐标.
func (w *windowBase) X取布局坐标(接收返回坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X窗口_取布局坐标(w.Handle, 接收返回坐标)
}

// 窗口_移动窗口.
//
// x: X坐标.
//
// y: Y坐标.
func (w *windowBase) X移动窗口(X坐标, Y坐标 int32) {
	炫彩基类.X窗口_移动窗口(w.Handle, X坐标, Y坐标)
}

// 窗口_取坐标.
//
// pRect: 坐标.
func (w *windowBase) X取坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X窗口_取坐标(w.Handle, 坐标)
}

// 窗口_最大化.
//
// bMaximize: 是否最大化.
func (w *windowBase) X最大化(是否最大化 bool) int {
	return 炫彩基类.X窗口_最大化(w.Handle, 是否最大化)
}

// 窗口_置定时器.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) X置定时器(定时器ID, 间隔值 int) int {
	return 炫彩基类.X窗口_置定时器(w.Handle, 定时器ID, 间隔值)
}

// 窗口_关闭定时器.
//
// nIDEvent: 定时器ID.
func (w *windowBase) X关闭定时器(定时器ID int) int {
	return 炫彩基类.X窗口_关闭定时器(w.Handle, 定时器ID)
}

// 窗口_置炫彩定时器.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) X置炫彩定时器(定时器ID, 间隔值 int) int {
	return 炫彩基类.X窗口_置炫彩定时器(w.Handle, 定时器ID, 间隔值)
}

// 窗口_关闭炫彩定时器.
//
// nIDEvent: 定时器ID.
func (w *windowBase) X关闭炫彩定时器(定时器ID int) int {
	return 炫彩基类.X窗口_关闭炫彩定时器(w.Handle, 定时器ID)
}

// 窗口_取背景管理器.
func (w *windowBase) X取背景管理器() int {
	return 炫彩基类.X窗口_取背景管理器(w.Handle)
}

// 窗口_取背景管理器扩展.
func (w *windowBase) X取背景管理器EX() int {
	return 炫彩基类.X窗口_取背景管理器EX(w.Handle)
}

// 窗口_置背景管理器.
//
// hBkInfoM: 背景管理器.
func (w *windowBase) X置背景管理器(背景管理器 int) int {
	return 炫彩基类.X窗口_置背景管理器(w.Handle, 背景管理器)
}

// 窗口_置透明类型.
//
// nType: 窗口透明类型: xcc.Window_Transparent_.
func (w *windowBase) X置透明类型(窗口透明类型 炫彩常量类.Window_Transparent_) int {
	return 炫彩基类.X窗口_置透明类型(w.Handle, 窗口透明类型)
}

// 窗口_置透明度.
//
// alpha: 窗口透明度, 范围0-255之间, 0透明, 255不透明.
func (w *windowBase) X置透明度(窗口透明度 byte) int {
	return 炫彩基类.X窗口_置透明度(w.Handle, 窗口透明度)
}

// 窗口_置透明色.
//
// color: 窗口透明色, ABGR 颜色.
func (w *windowBase) X置透明色(窗口透明色 int) int {
	return 炫彩基类.X窗口_置透明色(w.Handle, 窗口透明色)
}

// 窗口_置阴影信息.
//
// nSize: 阴影大小.
//
// nDepth: 阴影深度, 0-255.
//
// nAngeleSize: 圆角阴影内收大小.
//
// bRightAngle: 是否强制直角.
//
// color: 阴影颜色, ABGR 颜色.
func (w *windowBase) X置阴影信息(阴影大小 int, 阴影深度 byte, 圆角阴影内收大小 int, 是否强制直角 bool, 阴影颜色 int) int {
	return 炫彩基类.X窗口_置阴影信息(w.Handle, 阴影大小, 阴影深度, 圆角阴影内收大小, 是否强制直角, 阴影颜色)
}

// 窗口_取阴影信息.
//
// pnSize: 阴影大小.
//
// pnDepth: 阴影深度.
//
// pnAngeleSize: 圆角阴影内收大小 .
//
// pbRightAngle: 是否强制直角.
//
// pColor: 阴影颜色, ABGR 颜色.
func (w *windowBase) X取阴影信息(阴影大小, 阴影深度, 圆角阴影内收大小 *int32, 是否强制直角 *bool, 阴影颜色 *int) int {
	return 炫彩基类.X窗口_取阴影信息(w.Handle, 阴影大小, 阴影深度, 圆角阴影内收大小, 是否强制直角, 阴影颜色)
}

// 窗口_取透明类型, 返回: xcc.Window_Transparent_.
func (w *windowBase) X取透明类型() 炫彩常量类.Window_Transparent_ {
	return 炫彩基类.X窗口_取透明类型(w.Handle)
}

// 窗口_启用拖放文件.
//
// bEnable: 是否启用.
func (w *windowBase) X启用拖放文件(是否启用 bool) bool {
	return 炫彩基类.X窗口_启用拖放文件(w.Handle, 是否启用)
}

// 窗口_显示 显示隐藏窗口.
//
// bShow: 是否显示.
func (w *windowBase) X显示(是否显示 bool) int {
	return 炫彩基类.X窗口_显示(w.Handle, 是否显示)
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
func (w *windowBase) X取插入符信息(接收返回x坐标, 接收返回y坐标, 接收返回宽度, 接收返回高度 *int32) int {
	return 炫彩基类.X窗口_取插入符信息(w.Handle, 接收返回x坐标, 接收返回y坐标, 接收返回宽度, 接收返回高度)
}

// 窗口_置图标.
//
// hImage: 图标句柄.
func (w *windowBase) X置图标(图标句柄 int) int {
	return 炫彩基类.X窗口_置图标(w.Handle, 图标句柄)
}

// 窗口_置标题.
//
// pTitle: 标题文本.
func (w *windowBase) X置标题(标题文本 string) int {
	return 炫彩基类.X窗口_置标题(w.Handle, 标题文本)
}

// 窗口_置标题颜色.
//
// color: ABGR 颜色.
func (w *windowBase) X置标题颜色(ABGR颜色 int) int {
	return 炫彩基类.X窗口_置标题颜色(w.Handle, ABGR颜色)
}

// 窗口_取控制按钮, 返回按钮句柄.
//
// nFlag: xcc.Window_Style_ . 可用值: xcc.Window_Style_Btn_Min , xcc.Window_Style_Btn_Max , xcc.Window_Style_Btn_Close .
func (w *windowBase) X取控制按钮(xcc 炫彩常量类.Window_Style_) int {
	return 炫彩基类.X窗口_取控制按钮(w.Handle, xcc)
}

// 窗口_取图标, 返回图标句柄.
func (w *windowBase) X取图标() int {
	return 炫彩基类.X窗口_取图标(w.Handle)
}

// 窗口_取标题, 返回标题文本.
func (w *windowBase) X取标题() string {
	return 炫彩基类.X窗口_取标题(w.Handle)
}

// 窗口_取标题颜色, 返回ABGR 颜色.
func (w *windowBase) X取标题颜色() int {
	return 炫彩基类.X窗口_取标题颜色(w.Handle)
}

// 窗口_添加背景边框.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
func (w *windowBase) X添加背景边框(组合状态 炫彩常量类.Window_State_Flag_, ABGR颜色 int, 线宽 int) int {
	return 炫彩基类.X窗口_添加背景边框(w.Handle, 组合状态, ABGR颜色, 线宽)
}

// 窗口_添加背景填充.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
func (w *windowBase) X添加背景填充(组合状态 炫彩常量类.Window_State_Flag_, ABGR颜色 int) int {
	return 炫彩基类.X窗口_添加背景填充(w.Handle, 组合状态, ABGR颜色)
}

// 窗口_添加背景图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func (w *windowBase) X添加背景图片(组合状态 炫彩常量类.Window_State_Flag_, 图片句柄 int) int {
	return 炫彩基类.X窗口_添加背景图片(w.Handle, 组合状态, 图片句柄)
}

// 窗口_取背景对象数量.
func (w *windowBase) X取背景对象数量() int {
	return 炫彩基类.X窗口_取背景对象数量(w.Handle)
}

// 窗口_清空背景对象.
func (w *windowBase) X清空背景对象() int {
	return 炫彩基类.X窗口_清空背景对象(w.Handle)
}

// 通知消息_窗口中弹出, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
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
func (w *windowBase) X通知消息_窗口中弹出(位置 炫彩常量类.Position_Flag_, 标题, 内容 string, 图标 int, 外观类型 炫彩常量类.NotifyMsg_Skin_) int {
	return 炫彩基类.X通知消息_窗口中弹出(w.Handle, 位置, 标题, 内容, 图标, 外观类型)
}

// 通知消息_窗口中弹出扩展, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
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
func (w *windowBase) X通知消息_窗口中弹出EX(位置 炫彩常量类.Position_Flag_, 标题, 内容 string, 图标 int, 外观类型 炫彩常量类.NotifyMsg_Skin_, 是否启用关闭按钮, 是否自动关闭 bool, 自定义宽度, 自定义高度 int) int {
	return 炫彩基类.X通知消息_窗口中弹出EX(w.Handle, 位置, 标题, 内容, 图标, 外观类型, 是否启用关闭按钮, 是否自动关闭, 自定义宽度, 自定义高度)
}

// 通知消息_置持续时间.
//
// duration: 持续时间.
func (w *windowBase) X通知消息_置持续时间(持续时间 int) int {
	return 炫彩基类.X通知消息_置持续时间(w.Handle, 持续时间)
}

// 通知消息_置父边距 设置通知消息与父对象的四边间隔.
//
// left: 左侧间隔, 未实现, 预留功能.
//
// top: 顶部间隔.
//
// right: 右侧间隔.
//
// bottom: 底部间隔, 未实现, 预留功能.
func (w *windowBase) X通知消息_置父边距(左侧间隔, 顶部间隔, 右侧间隔, 底部间隔 int) int {
	return 炫彩基类.X通知消息_置父边距(w.Handle, 左侧间隔, 顶部间隔, 右侧间隔, 底部间隔)
}

// 通知消息_置标题高度.
//
// nHeight: 高度.
func (w *windowBase) X通知消息_置标题高度(高度 int) int {
	return 炫彩基类.X通知消息_置标题高度(w.Handle, 高度)
}

// 通知消息_置宽度, 设置默认宽度.
//
// nWidth: 宽度.
func (w *windowBase) X通知消息_置宽度(宽度 int) int {
	return 炫彩基类.X通知消息_置宽度(w.Handle, 宽度)
}

// 通知消息_置间隔.
//
// nSpace: 间隔大小.
func (w *windowBase) X通知消息_置间隔(间隔大小 int) int {
	return 炫彩基类.X通知消息_置间隔(w.Handle, 间隔大小)
}

// 通知消息_置边大小, 设置通知消息面板边大小.
//
// left: 左边.
//
// top: 顶边.
//
// right: 右边.
//
// bottom: 底边.
func (w *windowBase) X通知消息_置边大小(左边, 顶边, 右边, 底边 int) int {
	return 炫彩基类.X通知消息_置边大小(w.Handle, 左边, 顶边, 右边, 底边)
}

// 窗口_置背景, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.
func (w *windowBase) X置背景文本(背景内容字符串 string) int {
	return 炫彩基类.X窗口_置背景(w.Handle, 背景内容字符串)
}

// 窗口_是否可拖动标题栏.
func (w *windowBase) X是否可拖动标题栏() bool {
	return 炫彩基类.X窗口_是否可拖动标题栏(w.Handle)
}

// 窗口_是否可拖动窗口.
func (w *windowBase) X是否可拖动窗口() bool {
	return 炫彩基类.X窗口_是否可拖动窗口(w.Handle)
}

// 窗口_是否可拖动边框.
func (w *windowBase) X是否可拖动边框() bool {
	return 炫彩基类.X窗口_是否可拖动边框(w.Handle)
}

// 窗口_置标题外间距, 设置标题内容(图标, 标题, 控制按钮)外间距.
//
// left: 左边间距.
//
// top: 上边间距.
//
// right: 右边间距.
//
// bottom: 下边间距.
func (w *windowBase) X置标题外间距(左 int, 上 int, 右 int, 下 int) int {
	return 炫彩基类.X窗口_置标题外间距(w.Handle, 左, 上, 右, 下)
}

// SetTopEx 窗口_置顶Ex.
//
//	@param b 是否置顶.
//	@return bool
func (w *windowBase) XSetTopEx(是否置顶 bool) bool {
	return 炫彩WinApi窗口类.X窗口置顶(w.X取HWND(), 是否置顶)
}

// 窗口_置窗口位置. 封装系统API SetWindowPos(), 内部做了DPI适配.
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
func (w *windowBase) X置窗口位置(置顶方式 炫彩常量类.HWND_, X坐标, Y坐标, 宽度, 高度 int32, 标志 炫彩常量类.SWP_) int {
	return 炫彩基类.X窗口_置窗口位置(w.Handle, 置顶方式, X坐标, Y坐标, 宽度, 高度, 标志)
}

// 窗口_取DPI. 获取当前窗口所在显示器DPI, 返回窗口DPI.
func (w *windowBase) X取DPI() int {
	return 炫彩基类.X窗口_取DPI(w.Handle)
}

// 窗口_坐标转换DPI. 窗口客户区坐标转换到缩放后DPI坐标.
//
// pRect: 接收返回坐标.
func (w *windowBase) X坐标转换DPI(接收返回坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X窗口_坐标转换DPI(w.Handle, 接收返回坐标)
}

// 窗口_坐标点转换DPI. 窗口客户区坐标点转换到缩放后.
//
// pPt: 接收返回坐标点.
func (w *windowBase) X坐标点转换DPI(接收返回坐标点 *炫彩基类.POINT) int {
	return 炫彩基类.X窗口_坐标点转换DPI(w.Handle, 接收返回坐标点)
}

// 窗口_取光标位置. 封装的系统API: GetCursorPos(), 内部做了DPI适配.
//
// pPt: 接收返回坐标点.
func (w *windowBase) X取光标位置(接收返回坐标点 *炫彩基类.POINT) bool {
	return 炫彩基类.X窗口_取光标位置(w.Handle, 接收返回坐标点)
}

// 窗口_客户区坐标点到屏幕坐标点. 封装的系统API: ClientToScreen, 内部做了DPI适配.
//
// pPt: 接收返回坐标点.
func (w *windowBase) X客户区坐标点到屏幕坐标点(接收返回坐标点 *炫彩基类.POINT) bool {
	return 炫彩基类.X窗口_客户区坐标点到屏幕坐标点(w.Handle, 接收返回坐标点)
}

// 窗口_屏幕坐标点到客户区坐标点. 封装的系统API: ScreenToClient(), 内部做了DPI适配.
//
// pPt: 接收返回坐标点.
func (w *windowBase) X屏幕坐标点到客户区坐标点(接收返回坐标点 *炫彩基类.POINT) bool {
	return 炫彩基类.X窗口_屏幕坐标点到客户区坐标点(w.Handle, 接收返回坐标点)
}

// 窗口_置DPI. 设置当前窗口DPI, 默认DPI为96.
//
// nDPI: DPI值.
func (w *windowBase) X置DPI(DPI值 int) int {
	return 炫彩基类.X窗口_置DPI(w.Handle, DPI值)
}

// 窗口_置大小. 设置窗口宽高.
//
// 宽: width.
//
// 高: height.
func (w *windowBase) X置大小(宽, 高 int32) bool {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	rc.Right = rc.Left + 宽
	rc.Bottom = rc.Top + 高
	return 炫彩基类.X窗口_置坐标(w.Handle, &rc)
}

// 窗口_置宽度.
//
// 宽: width.
func (w *windowBase) X置宽度(宽 int32) bool {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	rc.Right = rc.Left + 宽
	return 炫彩基类.X窗口_置坐标(w.Handle, &rc)
}

// 窗口_置高度.
//
// 高: height.
func (w *windowBase) X置高度(高 int32) bool {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	rc.Bottom = rc.Top + 高
	return 炫彩基类.X窗口_置坐标(w.Handle, &rc)
}

// 窗口_取宽度.
func (w *windowBase) X取宽度() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	return rc.Right - rc.Left
}

// 窗口_取高度.
func (w *windowBase) X取高度() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	return rc.Bottom - rc.Top
}

// 窗口_取左边.
func (w *windowBase) X取左边() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	return rc.Left
}

// 窗口_取顶边.
func (w *windowBase) X取顶边() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	return rc.Top
}

// 窗口_取右边.
func (w *windowBase) X取右边() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	return rc.Right
}

// 窗口_取底边.
func (w *windowBase) X取底边() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X窗口_取坐标(w.Handle, &rc)
	return rc.Bottom
}

// 窗口_置左边.
//
// x: 左边x坐标.
func (w *windowBase) X置左边(x坐标 int32) {
	炫彩基类.X窗口_移动窗口(w.Handle, x坐标, w.X取顶边())
}

// 窗口_置顶边.
//
// y: 顶边y坐标.
func (w *windowBase) X置顶边(y坐标 int32) {
	炫彩基类.X窗口_移动窗口(w.Handle, w.X取左边(), y坐标)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
//	@param bEnable 是否启用.
//	@return int
func (w *windowBase) X启用水平(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用水平(w.Handle, 是否启用)
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
//	@param bEnable 是否启用.
//	@return int
func (w *windowBase) X启用自动换行(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用自动换行(w.Handle, 是否启用)
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
//	@param bEnable 是否启用.
//	@return int
func (w *windowBase) X启用溢出隐藏(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用溢出隐藏(w.Handle, 是否启用)
}

// SetAlignH 布局盒子_置水平对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (w *windowBase) X置水平对齐(对齐方式 炫彩常量类.Layout_Align_) int {
	return 炫彩基类.X布局盒子_置水平对齐(w.Handle, 对齐方式)
}

// SetAlignV 布局盒子_置垂直对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (w *windowBase) X置垂直对齐(对齐方式 炫彩常量类.Layout_Align_) int {
	return 炫彩基类.X布局盒子_置垂直对齐(w.Handle, 对齐方式)
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_Axis_.
//	@return int
func (w *windowBase) X置对齐基线(对齐方式 炫彩常量类.Layout_Align_Axis_) int {
	return 炫彩基类.X布局盒子_置对齐基线(w.Handle, 对齐方式)
}

// SetSpace 布局盒子_置间距.
//
//	@param nSpace 项间距大小.
//	@return int
func (w *windowBase) X置间距(项间距大小 int) int {
	return 炫彩基类.X布局盒子_置间距(w.Handle, 项间距大小)
}

// SetSpaceRow 布局盒子_置行距.
//
//	@param nSpace 行间距大小.
//	@return int
func (w *windowBase) X置行距(行间距大小 int) int {
	return 炫彩基类.X布局盒子_置行距(w.Handle, 行间距大小)
}

/*
下面都是事件
*/

type XWM_TRAYICON func(wParam, lParam uint, pbHandled *bool) int // 托盘图标事件.

type XWM_WINDPROC func(message uint32, wParam, lParam uint, pbHandled *bool) int               // 窗口消息过程.
type XWM_WINDPROC1 func(hWindow int, message uint32, wParam, lParam uint, pbHandled *bool) int // 窗口消息过程.
type XWM_XC_TIMER func(nTimerID uint, pbHandled *bool) int                                     // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_XC_TIMER1 func(hWindow int, nTimerID uint, pbHandled *bool) int                       // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_SETFOCUS_ELE func(hEle int, pbHandled *bool) int                                      // 窗口事件_置焦点元素. 指定元素获得焦点
type XWM_SETFOCUS_ELE1 func(hWindow int, hEle int, pbHandled *bool) int                        // 窗口事件_置焦点元素. 指定元素获得焦点
type XWM_FLOAT_PANE func(hFloatWnd int, hPane int, pbHandled *bool) int                        // 浮动窗格.
type XWM_FLOAT_PANE1 func(hWindow int, hFloatWnd int, hPane int, pbHandled *bool) int          // 浮动窗格.
type XWM_PAINT_END func(hDraw int, pbHandled *bool) int                                        // 窗口绘制完成消息.
type XWM_PAINT_END1 func(hWindow int, hDraw int, pbHandled *bool) int                          // 窗口绘制完成消息.
type XWM_PAINT_DISPLAY func(pbHandled *bool) int                                               // 窗口绘制完成并且已经显示到屏幕.
type XWM_PAINT_DISPLAY1 func(hWindow int, pbHandled *bool) int                                 // 窗口绘制完成并且已经显示到屏幕.
type XWM_DOCK_POPUP func(hWindowDock, hPane int, pbHandled *bool) int                          // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_DOCK_POPUP1 func(hWindow int, hWindowDock, hPane int, pbHandled *bool) int            // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_BODYVIEW_RECT func(width, height int32, pbHandled *bool) int                          // 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件
type XWM_BODYVIEW_RECT1 func(hWindow int, width, height int32, pbHandled *bool) int            // 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件
// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
type XWM_FLOATWND_DRAG func(hFloatWnd int, hArray *[6]int, pbHandled *bool) int

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hWindow: 传入的窗口资源句柄.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
type XWM_FLOATWND_DRAG1 func(hWindow int, hFloatWnd int, hArray *[6]int, pbHandled *bool) int

type WM_PAINT func(hDraw int, pbHandled *bool) int                                        // 窗口绘制消息.
type WM_PAINT1 func(hWindow int, hDraw int, pbHandled *bool) int                          // 窗口绘制消息.
type WM_CLOSE func(pbHandled *bool) int                                                   // 窗口关闭消息.
type WM_CLOSE1 func(hWindow int, pbHandled *bool) int                                     // 窗口关闭消息.
type WM_DESTROY func(pbHandled *bool) int                                                 // 窗口销毁消息.
type WM_DESTROY1 func(hWindow int, pbHandled *bool) int                                   // 窗口销毁消息.
type WM_NCDESTROY func(pbHandled *bool) int                                               // 窗口非客户区销毁消息.
type WM_NCDESTROY1 func(hWindow int, pbHandled *bool) int                                 // 窗口非客户区销毁消息.
type WM_MOUSEMOVE func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int                   // 窗口鼠标移动消息.
type WM_MOUSEMOVE1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int     // 窗口鼠标移动消息.
type WM_LBUTTONDOWN func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int                 // 窗口鼠标左键按下消息.
type WM_LBUTTONDOWN1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int   // 窗口鼠标左键按下消息.
type WM_LBUTTONUP func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int                   // 窗口鼠标左键弹起消息.
type WM_LBUTTONUP1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int     // 窗口鼠标左键弹起消息.
type WM_RBUTTONDOWN func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int                 // 窗口鼠标右键按下消息.
type WM_RBUTTONDOWN1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int   // 窗口鼠标右键按下消息.
type WM_RBUTTONUP func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int                   // 窗口鼠标右键弹起消息.
type WM_RBUTTONUP1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int     // 窗口鼠标右键弹起消息.
type WM_LBUTTONDBLCLK func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int               // 窗口鼠标左键双击消息.
type WM_LBUTTONDBLCLK1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int // 窗口鼠标左键双击消息.
type WM_RBUTTONDBLCLK func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int               // 窗口鼠标右键双击消息.
type WM_RBUTTONDBLCLK1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int // 窗口鼠标右键双击消息.
type WM_MOUSEWHEEL func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int                  // 窗口鼠标滚轮滚动消息.
type WM_MOUSEWHEEL1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int    // 窗口鼠标滚轮滚动消息.
type WM_EXITSIZEMOVE func(pbHandled *bool) int                                            // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
type WM_EXITSIZEMOVE1 func(hWindow int, pbHandled *bool) int                              // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
type WM_MOUSEHOVER func(nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int                  // 窗口鼠标进入消息.
type WM_MOUSEHOVER1 func(hWindow int, nFlags uint, pPt *炫彩基类.POINT, pbHandled *bool) int    // 窗口鼠标进入消息.
type WM_MOUSELEAVE func(pbHandled *bool) int                                              // 窗口鼠标离开消息.
type WM_MOUSELEAVE1 func(hWindow int, pbHandled *bool) int                                // 窗口鼠标离开消息.
type WM_SIZE func(nFlags uint, pPt *炫彩基类.SIZE, pbHandled *bool) int                         // 窗口大小改变消息.
type WM_SIZE1 func(hWindow int, nFlags uint, pPt *炫彩基类.SIZE, pbHandled *bool) int           // 窗口大小改变消息.
type WM_TIMER func(nIDEvent uint, pbHandled *bool) int                                    // 窗口定时器消息.
type WM_TIMER1 func(hWindow int, nIDEvent uint, pbHandled *bool) int                      // 窗口定时器消息.
type WM_SETFOCUS func(pbHandled *bool) int                                                // 窗口获得焦点.
type WM_SETFOCUS1 func(hWindow int, pbHandled *bool) int                                  // 窗口获得焦点.
type WM_KILLFOCUS func(pbHandled *bool) int                                               // 窗口失去焦点.
type WM_KILLFOCUS1 func(hWindow int, pbHandled *bool) int                                 // 窗口失去焦点.
type WM_KEYDOWN func(wParam, lParam uint, pbHandled *bool) int                            // 窗口键盘按键消息.
type WM_KEYDOWN1 func(hWindow int, wParam, lParam uint, pbHandled *bool) int              // 窗口键盘按键消息.
type WM_CAPTURECHANGED func(hWnd uintptr, pbHandled *bool) int                            // 窗口鼠标捕获改变消息.
type WM_CAPTURECHANGED1 func(hWindow int, hWnd uintptr, pbHandled *bool) int              // 窗口鼠标捕获改变消息.
type WM_SETCURSOR func(wParam, lParam uint, pbHandled *bool) int                          // 窗口设置鼠标光标.
type WM_SETCURSOR1 func(hWindow int, wParam, lParam uint, pbHandled *bool) int            // 窗口设置鼠标光标.
type WM_CHAR func(wParam, lParam uint, pbHandled *bool) int                               // 窗口字符消息.
type WM_CHAR1 func(hWindow int, wParam, lParam uint, pbHandled *bool) int                 // 窗口字符消息.
type WM_DROPFILES func(hDropInfo uintptr, pbHandled *bool) int                            // 拖动文件到窗口.
type WM_DROPFILES1 func(hWindow int, hDropInfo uintptr, pbHandled *bool) int              // 拖动文件到窗口.

// 托盘图标事件.
func (w *windowBase) X线程_托盘图标事件(pFun XWM_TRAYICON) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_TRAYICON, pFun)
}

// 窗口消息过程.
func (w *windowBase) X线程_消息过程(pFun XWM_WINDPROC) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_WINDPROC, pFun)
}

// 窗口消息过程.
func (w *windowBase) X线程_消息过程1(pFun XWM_WINDPROC1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_WINDPROC, pFun)
}

// 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
func (w *windowBase) X线程_炫彩定时器(pFun XWM_XC_TIMER) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_XC_TIMER, pFun)
}

// 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
func (w *windowBase) X线程_炫彩定时器1(pFun XWM_XC_TIMER1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_XC_TIMER, pFun)
}

// 窗口事件_置焦点元素. 指定元素获得焦点.
func (w *windowBase) X事件_置焦点元素(pFun XWM_SETFOCUS_ELE) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_SETFOCUS_ELE, pFun)
}

// 窗口事件_置焦点元素. 指定元素获得焦点.
func (w *windowBase) X事件_置焦点元素1(pFun XWM_SETFOCUS_ELE1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_SETFOCUS_ELE, pFun)
}

// 浮动窗格.
func (w *windowBase) X线程_浮动窗格(pFun XWM_FLOAT_PANE) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_FLOAT_PANE, pFun)
}

// 浮动窗格.
func (w *windowBase) X线程_浮动窗格1(pFun XWM_FLOAT_PANE1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_FLOAT_PANE, pFun)
}

// 窗口绘制完成消息.
func (w *windowBase) X线程_绘制完成消息(pFun XWM_PAINT_END) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_PAINT_END, pFun)
}

// 窗口绘制完成消息.
func (w *windowBase) X线程_绘制完成消息1(pFun XWM_PAINT_END1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_PAINT_END, pFun)
}

// 窗口绘制完成并且已经显示到屏幕.
func (w *windowBase) X线程_绘制完成并且已经显示到屏幕(pFun XWM_PAINT_DISPLAY) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_PAINT_DISPLAY, pFun)
}

// 窗口绘制完成并且已经显示到屏幕.
func (w *windowBase) X线程_绘制完成并且已经显示到屏幕1(pFun XWM_PAINT_DISPLAY1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_PAINT_DISPLAY, pFun)
}

// 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
func (w *windowBase) X线程_框架窗口码头弹出窗格(pFun XWM_DOCK_POPUP) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_DOCK_POPUP, pFun)
}

// 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
func (w *windowBase) X线程_框架窗口码头弹出窗格1(pFun XWM_DOCK_POPUP1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_DOCK_POPUP, pFun)
}

// 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件.
func (w *windowBase) X线程_框架窗口主视图坐标改变(pFun XWM_BODYVIEW_RECT) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_BODYVIEW_RECT, pFun)
}

// 框架窗口主视图坐标改变, 如果主视图没有绑定元素, 那么当坐标改变时触发此事件.
func (w *windowBase) X线程_框架窗口主视图坐标改变1(pFun XWM_BODYVIEW_RECT) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_BODYVIEW_RECT, pFun)
}

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
func (w *windowBase) X线程_浮动窗口拖动(pFun XWM_FLOATWND_DRAG) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_FLOATWND_DRAG, pFun)
}

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hWindow: 传入的窗口资源句柄.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
func (w *windowBase) X线程_浮动窗口拖动1(pFun XWM_FLOATWND_DRAG1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_FLOATWND_DRAG, pFun)
}

// 窗口绘制消息.
func (w *windowBase) X线程_绘制消息(pFun WM_PAINT) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_PAINT, pFun)
}

// 窗口绘制消息.
func (w *windowBase) X线程_绘制消息1(pFun WM_PAINT1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_PAINT, pFun)
}

// 窗口关闭消息.
func (w *windowBase) X线程_关闭消息(pFun WM_CLOSE) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_CLOSE, pFun)
}

// 窗口关闭消息.
func (w *windowBase) X线程_关闭消息1(pFun WM_CLOSE1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_CLOSE, pFun)
}

// 窗口销毁消息.
func (w *windowBase) X线程_销毁消息(pFun WM_DESTROY) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_DESTROY, pFun)
}

// 窗口销毁消息.
func (w *windowBase) X线程_销毁消息1(pFun WM_DESTROY1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_DESTROY, pFun)
}

// 窗口非客户区销毁消息.
func (w *windowBase) X线程_非客户区销毁消息(pFun WM_NCDESTROY) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_NCDESTROY, pFun)
}

// 窗口非客户区销毁消息.
func (w *windowBase) X线程_非客户区销毁消息1(pFun WM_NCDESTROY1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_NCDESTROY, pFun)
}

// 窗口鼠标移动消息.
func (w *windowBase) X线程_鼠标移动消息(pFun WM_MOUSEMOVE) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_MOUSEMOVE, pFun)
}

// 窗口鼠标移动消息.
func (w *windowBase) X线程_鼠标移动消息1(pFun WM_MOUSEMOVE1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_MOUSEMOVE, pFun)
}

// 窗口鼠标左键按下消息.
func (w *windowBase) X线程_鼠标左键按下消息(pFun WM_LBUTTONDOWN) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_LBUTTONDOWN, pFun)
}

// 窗口鼠标左键按下消息.
func (w *windowBase) X线程_鼠标左键按下消息1(pFun WM_LBUTTONDOWN1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_LBUTTONDOWN, pFun)
}

// 窗口鼠标左键弹起消息.
func (w *windowBase) X线程_鼠标左键弹起消息(pFun WM_LBUTTONUP) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_LBUTTONUP, pFun)
}

// 窗口鼠标左键弹起消息.
func (w *windowBase) X线程_鼠标左键弹起消息1(pFun WM_LBUTTONUP1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_LBUTTONUP, pFun)
}

// 窗口鼠标右键按下消息.
func (w *windowBase) X线程_鼠标右键按下消息(pFun WM_RBUTTONDOWN) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_RBUTTONDOWN, pFun)
}

// 窗口鼠标右键按下消息.
func (w *windowBase) X线程_鼠标右键按下消息1(pFun WM_RBUTTONDOWN1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_RBUTTONDOWN, pFun)
}

// 窗口鼠标右键弹起消息.
func (w *windowBase) X线程_鼠标右键弹起消息(pFun WM_RBUTTONUP) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_RBUTTONUP, pFun)
}

// 窗口鼠标右键弹起消息.
func (w *windowBase) X线程_鼠标右键弹起消息1(pFun WM_RBUTTONUP1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_RBUTTONUP, pFun)
}

// 窗口鼠标左键双击消息.
func (w *windowBase) X线程_鼠标左键双击消息(pFun WM_LBUTTONDBLCLK) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_LBUTTONDBLCLK, pFun)
}

// 窗口鼠标左键双击消息.
func (w *windowBase) X线程_鼠标左键双击消息1(pFun WM_LBUTTONDBLCLK1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_LBUTTONDBLCLK, pFun)
}

// 窗口鼠标右键双击消息.
func (w *windowBase) X线程_鼠标右键双击消息(pFun WM_RBUTTONDBLCLK) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_RBUTTONDBLCLK, pFun)
}

// 窗口鼠标右键双击消息.
func (w *windowBase) X线程_鼠标右键双击消息1(pFun WM_RBUTTONDBLCLK1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_RBUTTONDBLCLK, pFun)
}

// 窗口鼠标滚轮滚动消息.
func (w *windowBase) X线程_鼠标滚轮滚动消息(pFun WM_MOUSEWHEEL) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_MOUSEWHEEL, pFun)
}

// 窗口鼠标滚轮滚动消息.
func (w *windowBase) X线程_鼠标滚轮滚动消息1(pFun WM_MOUSEWHEEL1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_MOUSEWHEEL, pFun)
}

// 窗口退出移动或调整大小模式循环改，详情参见MSDN.
func (w *windowBase) X线程_退出移动或调整大小模式循环改(pFun WM_EXITSIZEMOVE) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_EXITSIZEMOVE, pFun)
}

// 窗口退出移动或调整大小模式循环改，详情参见MSDN.
func (w *windowBase) X线程_退出移动或调整大小模式循环改1(pFun WM_EXITSIZEMOVE1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_EXITSIZEMOVE, pFun)
}

// 窗口鼠标进入消息.
func (w *windowBase) X线程_鼠标进入消息(pFun WM_MOUSEHOVER) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_MOUSEHOVER, pFun)
}

// 窗口鼠标进入消息.
func (w *windowBase) X线程_鼠标进入消息1(pFun WM_MOUSEHOVER1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_MOUSEHOVER, pFun)
}

// 窗口鼠标离开消息.
func (w *windowBase) X线程_鼠标离开消息(pFun WM_MOUSELEAVE) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_MOUSELEAVE, pFun)
}

// 窗口鼠标离开消息.
func (w *windowBase) X线程_鼠标离开消息1(pFun WM_MOUSELEAVE1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_MOUSELEAVE, pFun)
}

// 窗口大小改变消息.
func (w *windowBase) X线程_大小改变消息(pFun WM_SIZE) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_SIZE, pFun)
}

// 窗口大小改变消息.
func (w *windowBase) X线程_大小改变消息1(pFun WM_SIZE1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_SIZE, pFun)
}

// 窗口定时器消息.
func (w *windowBase) X线程_定时器消息(pFun WM_TIMER) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_TIMER, pFun)
}

// 窗口定时器消息.
func (w *windowBase) X线程_定时器消息1(pFun WM_TIMER1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_TIMER, pFun)
}

// 窗口获得焦点.
func (w *windowBase) X线程_获得焦点(pFun WM_SETFOCUS) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_SETFOCUS, pFun)
}

// 窗口获得焦点.
func (w *windowBase) X线程_获得焦点1(pFun WM_SETFOCUS1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_SETFOCUS, pFun)
}

// 窗口失去焦点.
func (w *windowBase) X线程_失去焦点(pFun WM_KILLFOCUS) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_KILLFOCUS, pFun)
}

// 窗口失去焦点.
func (w *windowBase) X线程_失去焦点1(pFun WM_KILLFOCUS1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_KILLFOCUS, pFun)
}

// 窗口键盘按键消息.
func (w *windowBase) X线程_键盘按键消息(pFun WM_KEYDOWN) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_KEYDOWN, pFun)
}

// 窗口键盘按键消息.
func (w *windowBase) X线程_键盘按键消息1(pFun WM_KEYDOWN1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_KEYDOWN, pFun)
}

// 窗口鼠标捕获改变消息.
func (w *windowBase) X线程_鼠标捕获改变消息(pFun WM_CAPTURECHANGED) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_CAPTURECHANGED, pFun)
}

// 窗口鼠标捕获改变消息.
func (w *windowBase) X线程_鼠标捕获改变消息1(pFun WM_CAPTURECHANGED1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_CAPTURECHANGED, pFun)
}

// 窗口设置鼠标光标.
func (w *windowBase) X线程_设置鼠标光标(pFun WM_SETCURSOR) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_SETCURSOR, pFun)
}

// 窗口设置鼠标光标.
func (w *windowBase) X线程_设置鼠标光标1(pFun WM_SETCURSOR1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_SETCURSOR, pFun)
}

// 窗口字符消息.
func (w *windowBase) X线程_字符消息(pFun WM_CHAR) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_CHAR, pFun)
}

// 窗口字符消息.
func (w *windowBase) X线程_字符消息1(pFun WM_CHAR1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_CHAR, pFun)
}

// 拖动文件到窗口.
func (w *windowBase) X线程_拖动文件到窗口(pFun WM_DROPFILES) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.WM_DROPFILES, pFun)
}

// 拖动文件到窗口.
func (w *windowBase) X线程_拖动文件到窗口1(pFun WM_DROPFILES1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.WM_DROPFILES, pFun)
}

type XWM_MENU_POPUP func(hMenu int, pbHandled *bool) int                                                         // 菜单弹出.
type XWM_MENU_POPUP1 func(hWindow int, hMenu int, pbHandled *bool) int                                           // 菜单弹出.
type XWM_MENU_POPUP_WND func(hMenu int, pInfo *炫彩基类.Menu_PopupWnd_, pbHandled *bool) int                           // 菜单弹出窗口.
type XWM_MENU_POPUP_WND1 func(hWindow int, hMenu int, pInfo *炫彩基类.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XWM_MENU_SELECT func(nID int32, pbHandled *bool) int                                                        // 菜单选择.
type XWM_MENU_SELECT1 func(hWindow int, nID int32, pbHandled *bool) int                                          // 菜单选择.
type XWM_MENU_EXIT func(pbHandled *bool) int                                                                     // 菜单退出.
type XWM_MENU_EXIT1 func(hWindow int, pbHandled *bool) int                                                       // 菜单退出.
type XWM_MENU_DRAW_BACKGROUND func(hDraw int, pInfo *炫彩基类.Menu_DrawBackground_, pbHandled *bool) int               // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XWM_MENU_DRAW_BACKGROUND1 func(hWindow int, hDraw int, pInfo *炫彩基类.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XWM_MENU_DRAWITEM func(hDraw int, pInfo *炫彩基类.Menu_DrawItem_, pbHandled *bool) int                            // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
type XWM_MENU_DRAWITEM1 func(hWindow int, hDraw int, pInfo *炫彩基类.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// 菜单弹出.
func (w *windowBase) X线程_菜单弹出(pFun XWM_MENU_POPUP) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_MENU_POPUP, pFun)
}

// 菜单弹出.
func (w *windowBase) X线程_菜单弹出1(pFun XWM_MENU_POPUP1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_MENU_POPUP, pFun)
}

// 菜单弹出窗口.
func (w *windowBase) X线程_菜单弹出窗口(pFun XWM_MENU_POPUP_WND) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_MENU_POPUP_WND, pFun)
}

// 菜单弹出窗口.
func (w *windowBase) X线程_菜单弹出窗口1(pFun XWM_MENU_POPUP_WND1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_MENU_POPUP_WND, pFun)
}

// 菜单选择.
func (w *windowBase) X线程_菜单选择(pFun XWM_MENU_SELECT) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_MENU_SELECT, pFun)
}

// 菜单选择.
func (w *windowBase) X线程_菜单选择1(pFun XWM_MENU_SELECT1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_MENU_SELECT, pFun)
}

// 菜单退出.
func (w *windowBase) X线程_菜单退出(pFun XWM_MENU_EXIT) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_MENU_EXIT, pFun)
}

// 菜单退出.
func (w *windowBase) X线程_菜单退出1(pFun XWM_MENU_EXIT1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_MENU_EXIT, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (w *windowBase) X线程_绘制菜单背景(pFun XWM_MENU_DRAW_BACKGROUND) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (w *windowBase) X线程_绘制菜单背景1(pFun XWM_MENU_DRAW_BACKGROUND1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (w *windowBase) X线程_绘制菜单项事件(pFun XWM_MENU_DRAWITEM) bool {
	return 炫彩基类.X窗口_注册事件C(w.Handle, 炫彩常量类.XWM_MENU_DRAWITEM, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (w *windowBase) X线程_绘制菜单项事件1(pFun XWM_MENU_DRAWITEM1) bool {
	return 炫彩基类.X窗口_注册事件C1(w.Handle, 炫彩常量类.XWM_MENU_DRAWITEM, pFun)
}
