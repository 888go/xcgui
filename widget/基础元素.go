package widget

import (
	"e.coding.net/gogit/go/xcgui/objectbase"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// Element 基础元素.
type Element struct {
	objectbase.Widget
}

// I元素_创建, 创建基础元素.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func I元素_创建(x int, y int, cx int, cy int, hParent int) *Element {
	p := &Element{}
	p.SetHandle(xc.XEle_Create(x, y, cx, cy, hParent))
	return p
}

// I元素_从句柄创建对象.
func I元素_从句柄创建(handle int) *Element {
	p := &Element{}
	p.SetHandle(handle)
	return p
}

// I元素_从name创建对象, 失败返回nil.
func I元素_从name创建(name string) *Element {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Element{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I元素_从UID创建对象, 失败返回nil.
func I元素_从UID创建(nUID int) *Element {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Element{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I元素_从UID名称创建对象, 失败返回nil.
func I元素_从UID名称创建(name string) *Element {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Element{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 元素_注册事件C, 注册事件C方式, 省略2参数.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针.
func (e *Element) I注册事件C(nEvent xcc.XE_, pFun interface{}) bool {
	return xc.XEle_RegEventC(e.I句柄, nEvent, pFun)
}

// 元素_注册事件C1, 注册事件C1方式, 省略1参数.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针.
func (e *Element) I注册事件C1(nEvent xcc.XE_, pFun interface{}) bool {
	return xc.XEle_RegEventC1(e.I句柄, nEvent, pFun)
}

// 元素_移除事件C.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针.
func (e *Element) I移除事件C(nEvent xcc.XE_, pFun interface{}) bool {
	return xc.XEle_RemoveEventC(e.I句柄, nEvent, pFun)
}

// 元素_发送事件.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func (e *Element) I发送事件(nEvent xcc.XE_, wParam int, lParam int) int {
	return xc.XEle_SendEvent(e.I句柄, nEvent, wParam, lParam)
}

// 元素_投递事件.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.
func (e *Element) I投递事件(nEvent xcc.XE_, wParam int, lParam int) int {
	return xc.XEle_PostEvent(e.I句柄, nEvent, wParam, lParam)
}

// 元素_取坐标.
//
// pRect: 坐标.
func (e *Element) I取坐标(pRect *xc.RECT) int {
	return xc.XEle_GetRect(e.I句柄, pRect)
}

// 元素_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.
func (e *Element) I取逻辑坐标(pRect *xc.RECT) int {
	return xc.XEle_GetRectLogic(e.I句柄, pRect)
}

// 元素_取客户区坐标.
//
// pRect: 坐标.
func (e *Element) I取客户区坐标(pRect *xc.RECT) int {
	return xc.XEle_GetClientRect(e.I句柄, pRect)
}

// 元素_置宽度.
//
// nWidth: 宽度.
func (e *Element) SetWidth(nWidth int) int {
	return xc.XEle_SetWidth(e.I句柄, nWidth)
}

// 元素_置高度.
//
// nHeight: 高度.
func (e *Element) I置高度(nHeight int) int {
	return xc.XEle_SetHeight(e.I句柄, nHeight)
}

// 元素_取宽度.
func (e *Element) I取宽度() int {
	return xc.XEle_GetWidth(e.I句柄)
}

// 元素_取高度.
func (e *Element) I取高度() int {
	return xc.XEle_GetHeight(e.I句柄)
}

// 元素_窗口客户区坐标到元素客户区坐标, 窗口客户区坐标转换到元素客户区坐标.
//
// pRect: 坐标.
func (e *Element) RectWndClientToEleClient(pRect *xc.RECT) int {
	return xc.XEle_RectWndClientToEleClient(e.I句柄, pRect)
}

// 元素_窗口客户区点到元素客户区, 窗口客户区坐标转换到元素客户区坐标.
//
// pPt: 坐标.
func (e *Element) PointWndClientToEleClient(pPt *xc.POINT) int {
	return xc.XEle_PointWndClientToEleClient(e.I句柄, pPt)
}

// 元素_客户区坐标到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// pRect: 坐标.
func (e *Element) RectClientToWndClient(pRect *xc.RECT) int {
	return xc.XEle_RectClientToWndClient(e.I句柄, pRect)
}

// 元素_客户区点到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// pPt: 坐标.
func (e *Element) PointClientToWndClient(pPt *xc.POINT) int {
	return xc.XEle_PointClientToWndClient(e.I句柄, pPt)
}

// 元素_基于窗口客户区坐标.
//
// pRect: 坐标.
func (e *Element) I基于窗口客户区坐标(pRect *xc.RECT) int {
	return xc.XEle_GetWndClientRect(e.I句柄, pRect)
}

// 元素_取光标, 获取元素鼠标光标, 返回光标句柄.
func (e *Element) I取光标() int {
	return xc.XEle_GetCursor(e.I句柄)
}

// 元素_置光标, 设置元素鼠标光标.
//
// hCursor: 光标句柄.
func (e *Element) I置光标(hCursor int) int {
	return xc.XEle_SetCursor(e.I句柄, hCursor)
}

// 元素_添加子对象.
//
// hChild: 要添加的子元素句柄或形状对象句柄.
func (e *Element) I添加子对象(hChild int) bool {
	return xc.XEle_AddChild(e.I句柄, hChild)
}

// 元素_插入子对象, 插入子对象到指定位置.
//
// hChild: 要插入的元素句柄或形状对象句柄.
//
// index: 插入位置索引.
func (e *Element) I插入子对象(hChild int, index int) bool {
	return xc.XEle_InsertChild(e.I句柄, hChild, index)
}

// 元素_置坐标, 如果返回0坐标没有改变, 如果大小改变返回2(触发XE_SIZE), 否则返回1(仅改变left,top,没有改变大小).
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) I置坐标(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_SetRect(e.I句柄, pRect, bRedraw, nFlags, nAdjustNo)
}

// 元素_置坐标扩展, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
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
func (e *Element) I置坐标扩展(x int, y int, cx int, cy int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_SetRectEx(e.I句柄, x, y, cx, cy, bRedraw, nFlags, nAdjustNo)
}

// 元素_置逻辑坐标, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_. 此参数将会传入XE_SIZE ,XE_ADJUSTLAYOUT 事件回调.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) I置逻辑坐标(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_SetRectLogic(e.I句柄, pRect, bRedraw, nFlags, nAdjustNo)
}

// 元素_移动, 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
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
func (e *Element) I移动(x int, y int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_SetPosition(e.I句柄, x, y, bRedraw, nFlags, nAdjustNo)
}

// 元素_移动逻辑坐标, 移动元素坐标, 逻辑坐标, 包含滚动视图偏移. 如果坐标未改变返回0, 如果大小改变返回2(触发XE_SIZE), 否则返回1.
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
func (e *Element) I移动逻辑坐标(x int, y int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_SetPositionLogic(e.I句柄, x, y, bRedraw, nFlags, nAdjustNo)
}

// 元素_判断绘制焦点.
func (e *Element) I判断绘制焦点() bool {
	return xc.XEle_IsDrawFocus(e.I句柄)
}

// 元素_判断启用, 元素是否为启用状态.
func (e *Element) I判断启用() bool {
	return xc.XEle_IsEnable(e.I句柄)
}

// 元素_判断启用焦点, 元素是否启用焦点.
func (e *Element) I判断启用焦点() bool {
	return xc.XEle_IsEnableFocus(e.I句柄)
}

// 元素_判断鼠标穿透, 元素是否启用鼠标穿透.
func (e *Element) I判断鼠标穿透() bool {
	return xc.XEle_IsMouseThrough(e.I句柄)
}

// 元素_测试点击元素, 检测坐标点所在元素, 包含子元素的子元素.
//
// pPt: 坐标点.
func (e *Element) I测试点击元素(pPt *xc.POINT) int {
	return xc.XEle_HitChildEle(e.I句柄, pPt)
}

// 元素_判断背景透明.
func (e *Element) I判断背景透明() bool {
	return xc.XEle_IsBkTransparent(e.I句柄)
}

// 元素_判断启用事件_XE_PAINT_END, 是否启XE_PAINT_END用事件.
func (e *Element) I判断启用事件_XE_PAINT_END() bool {
	return xc.XEle_IsEnableEvent_XE_PAINT_END(e.I句柄)
}

// 元素_判断接受TAB, 是否接受Tab键输入; 例如: XRichEdit, XEdit.
func (e *Element) I判断接受TAB() bool {
	return xc.XEle_IsKeyTab(e.I句柄)
}

// 元素_判断接受切换焦点, 是否接受通过键盘切换焦点(方向键,TAB键).
func (e *Element) IsSwitchFocus() bool {
	return xc.XEle_IsSwitchFocus(e.I句柄)
}

// 元素_判断启用_XE_MOUSEWHEEL, 判断是否启用鼠标滚动事件, 如果禁用那么事件会发送给他的父元素.
func (e *Element) I判断启用_XE_MOUSEWHEEL() bool {
	return xc.XEle_IsEnable_XE_MOUSEWHEEL(e.I句柄)
}

// 元素_判断为子元素, 判断hChildEle是否为hEle的子元素.
//
// hChildEle: 子元素句柄.
func (e *Element) I判断为子元素(hChildEle int) bool {
	return xc.XEle_IsChildEle(e.I句柄, hChildEle)
}

// 元素_判断启用画布, 判断是否启用画布.
func (e *Element) I判断启用画布() bool {
	return xc.XEle_IsEnableCanvas(e.I句柄)
}

// 元素_判断焦点, 判断是否拥有焦点.
func (e *Element) I判断焦点() bool {
	return xc.XEle_IsFocus(e.I句柄)
}

// 元素_判断焦点扩展, 判断该元素或该元素的子元素是否拥有焦点.
func (e *Element) I判断焦点扩展() bool {
	return xc.XEle_IsFocusEx(e.I句柄)
}

// 元素_启用, 启用或禁用元素.
//
// bEnable: 启用或禁用.
func (e *Element) I启用(bEnable bool) int {
	return xc.XEle_Enable(e.I句柄, bEnable)
}

// 元素_启用焦点, 启用焦点.
//
// bEnable: 是否启用.
func (e *Element) I启用焦点(bEnable bool) int {
	return xc.XEle_EnableFocus(e.I句柄, bEnable)
}

// 元素_启用绘制焦点.
//
// bEnable: 是否启用.
func (e *Element) I启用绘制焦点(bEnable bool) int {
	return xc.XEle_EnableDrawFocus(e.I句柄, bEnable)
}

// 元素_启用绘制边框, 启用或禁用绘制默认边框.
//
// bEnable: 是否启用.
func (e *Element) I启用绘制边框(bEnable bool) int {
	return xc.XEle_EnableDrawBorder(e.I句柄, bEnable)
}

// 元素_启用画布, 启用或禁用背景画布; 如果禁用那么将绘制在父的画布之上, 也就是说他没有自己的画布.
//
// bEnable: 是否启用.
func (e *Element) I启用画布(bEnable bool) int {
	return xc.XEle_EnableCanvas(e.I句柄, bEnable)
}

// 元素_启用事件_XE_PAINT_END.
//
// bEnable: 是否启用.
func (e *Element) I启用事件_XE_PAINT_END(bEnable bool) int {
	return xc.XEle_EnableEvent_XE_PAINT_END(e.I句柄, bEnable)
}

// 元素_启用背景透明.
//
// bEnable: 是否启用.
func (e *Element) I启用背景透明(bEnable bool) int {
	return xc.XEle_EnableBkTransparent(e.I句柄, bEnable)
}

// 元素_启用鼠标穿透. 启用鼠标穿透, 如果启用, 那么该元素不能接收到鼠标事件, 但是他的子元素不受影响, 任然可以接收鼠标事件.
//
// bEnable: 是否启用.
func (e *Element) I启用鼠标穿透(bEnable bool) int {
	return xc.XEle_EnableMouseThrough(e.I句柄, bEnable)
}

// 元素_启用接收TAB, 启用接收Tab输入.
//
// bEnable: 是否启用.
func (e *Element) I启用接收TAB(bEnable bool) int {
	return xc.XEle_EnableKeyTab(e.I句柄, bEnable)
}

// 元素_启用切换焦点, 启用接受通过键盘切换焦点.
//
// bEnable: 是否启用.
func (e *Element) I启用切换焦点(bEnable bool) int {
	return xc.XEle_EnableSwitchFocus(e.I句柄, bEnable)
}

// 元素_启用事件_XE_MOUSEWHEEL, 启用接收鼠标滚动事件, 如果禁用那么事件会传递给父元素.
//
// bEnable: 是否启用.
func (e *Element) I启用事件_XE_MOUSEWHEEL(bEnable bool) int {
	return xc.XEle_EnableEvent_XE_MOUSEWHEEL(e.I句柄, bEnable)
}

// 元素_移除, 移除元素, 但不销毁.
func (e *Element) I移除() int {
	return xc.XEle_Remove(e.I句柄)
}

// 元素_置Z序, 设置元素Z序.
//
// index: 位置索引.
func (e *Element) I置Z序(index int) bool {
	return xc.XEle_SetZOrder(e.I句柄, index)
}

// 元素_置Z序扩展, 设置元素Z序.
//
// hDestEle: 目标元素.
//
// nType: 类型, Zorder_.
func (e *Element) I置Z序扩展(hDestEle int, nType xcc.Zorder_) bool {
	return xc.XEle_SetZOrderEx(e.I句柄, hDestEle, nType)
}

// 元素_取Z序, 获取元素Z序索引, 位置索引.
func (e *Element) I取Z序() int {
	return xc.XEle_GetZOrder(e.I句柄)
}

// 元素_启用置顶, 设置元素置顶.
//
// bTopmost: 是否置顶显示.
func (e *Element) I启用置顶(bTopmost bool) bool {
	return xc.XEle_EnableTopmost(e.I句柄, bTopmost)
}

// 元素_重绘.
//
// bImmediate: 是否立即重绘.
func (e *Element) I重绘(bImmediate bool) int {
	return xc.XEle_Redraw(e.I句柄, bImmediate)
}

// 元素_重绘指定区域.
//
// pRect: 相对于元素客户区坐标.
//
// bImmediate: 是否立即重绘.
func (e *Element) I重绘指定区域(pRect *xc.RECT, bImmediate bool) int {
	return xc.XEle_RedrawRect(e.I句柄, pRect, bImmediate)
}

// 元素_取子对象数量, 获取子对象(UI元素和形状对象)数量, 只检测当前层子对象.
func (e *Element) I取子对象数量() int {
	return xc.XEle_GetChildCount(e.I句柄)
}

// 元素_取子对象从索引, 获取子对象通过索引, 只检测当前层子对象.
//
// index: 索引.
func (e *Element) I取子对象从索引(index int) int {
	return xc.XEle_GetChildByIndex(e.I句柄, index)
}

// 元素_取子对象从ID, 获取子对象通过ID, 只检测当前层子对象.
//
// nID: 元素ID.
func (e *Element) I取子对象从ID(nID int) int {
	return xc.XEle_GetChildByID(e.I句柄, nID)
}

// 元素_置边框大小.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func (e *Element) I置边框大小(left int, top int, right int, bottom int) int {
	return xc.XEle_SetBorderSize(e.I句柄, left, top, right, bottom)
}

// 元素_取边框大小.
//
// pBorder: 大小.
func (e *Element) I取边框大小(pBorder *xc.RECT) int {
	return xc.XEle_GetBorderSize(e.I句柄, pBorder)
}

// 元素_置内填充大小.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func (e *Element) I置内填充大小(left int, top int, right int, bottom int) int {
	return xc.XEle_SetPadding(e.I句柄, left, top, right, bottom)
}

// 元素_取内填充大小.
//
// pPadding: 大小.
func (e *Element) I取内填充大小(pPadding *xc.RECT) int {
	return xc.XEle_GetPadding(e.I句柄, pPadding)
}

// 元素_置拖动边框.
//
// nFlags: 边框位置组合, Element_Position_.
func (e *Element) I置拖动边框(nFlags xcc.Element_Position_) int {
	return xc.XEle_SetDragBorder(e.I句柄, nFlags)
}

// 元素_置拖动边框绑定元素, 设置拖动边框绑定元素, 当拖动边框时, 自动调整绑定元素的大小.
//
// nFlags: 边框位置标识, Element_Position_.
//
// hBindEle: 绑定元素.
//
// nSpace: 元素间隔大小.
func (e *Element) I置拖动边框绑定元素(nFlags xcc.Element_Position_, hBindEle int, nSpace int) int {
	return xc.XEle_SetDragBorderBindEle(e.I句柄, nFlags, hBindEle, nSpace)
}

// 元素_置最小大小.
//
// nWidth: 最小宽度.
//
// nHeight: 最小高度.
func (e *Element) I置最小大小(nWidth int, nHeight int) int {
	return xc.XEle_SetMinSize(e.I句柄, nWidth, nHeight)
}

// 元素_置最大大小.
//
// nWidth: 最大宽度.
//
// nHeight: 最大高度.
func (e *Element) I置最大大小(nWidth int, nHeight int) int {
	return xc.XEle_SetMaxSize(e.I句柄, nWidth, nHeight)
}

// 元素_置锁定滚动, 设置锁定元素在滚动视图中跟随滚动, 如果设置TRUE将不跟随滚动.
//
// bHorizon: 是否锁定水平滚动.
//
// bVertical: 是否锁定垂直滚动.
func (e *Element) I置锁定滚动(bHorizon bool, bVertical bool) int {
	return xc.XEle_SetLockScroll(e.I句柄, bHorizon, bVertical)
}

// 元素_置文本颜色.
//
// color: ABGR 颜色值.
func (e *Element) I置文本颜色(color int) int {
	return xc.XEle_SetTextColor(e.I句柄, color)
}

// 元素_取文本颜色.
func (e *Element) I取文本颜色() int {
	return xc.XEle_GetTextColor(e.I句柄)
}

// 元素_取文本颜色扩展, 获取文本颜色, 优先从资源中获取.
func (e *Element) I取文本颜色扩展() int {
	return xc.XEle_GetTextColorEx(e.I句柄)
}

// 元素_置焦点边框颜色.
//
// color: ABGR 颜色值.
func (e *Element) I置焦点边框颜色(color int) int {
	return xc.XEle_SetFocusBorderColor(e.I句柄, color)
}

// 元素_取焦点边框颜色.
func (e *Element) I取焦点边框颜色() int {
	return xc.XEle_GetFocusBorderColor(e.I句柄)
}

// 元素_置字体.
//
// hFontx: 炫彩字体.
func (e *Element) I置字体(hFontx int) int {
	return xc.XEle_SetFont(e.I句柄, hFontx)
}

// 元素_取字体.
func (e *Element) I取字体() int {
	return xc.XEle_GetFont(e.I句柄)
}

// 元素_取字体扩展, 获取元素字体, 优先从资源中获取.
func (e *Element) I取字体扩展() int {
	return xc.XEle_GetFontEx(e.I句柄)
}

// 元素_置透明度.
func (e *Element) I置透明度(alpha uint8) int {
	return xc.XEle_SetAlpha(e.I句柄, alpha)
}

// 元素_销毁.
func (e *Element) I销毁() int {
	return xc.XEle_Destroy(e.I句柄)
}

// 元素_添加背景边框, 添加背景内容边框.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
func (e *Element) I添加背景边框(nState xcc.CombinedState, color int, width int) int {
	return xc.XEle_AddBkBorder(e.I句柄, nState, color, width)
}

// 元素_添加背景填充, 添加背景内容填充.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
func (e *Element) I添加背景填充(nState xcc.CombinedState, color int) int {
	return xc.XEle_AddBkFill(e.I句柄, nState, color)
}

// 元素_添加背景图片, 添加背景内容图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func (e *Element) I添加背景图片(nState xcc.CombinedState, hImage int) int {
	return xc.XEle_AddBkImage(e.I句柄, nState, hImage)
}

// 元素_取背景对象数量, 获取背景内容数量.
func (e *Element) I取背景对象数量() int {
	return xc.XEle_GetBkInfoCount(e.I句柄)
}

// 元素_清空背景对象, 清空背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.
func (e *Element) I清空背景对象() int {
	return xc.XEle_ClearBkInfo(e.I句柄)
}

// 元素_取背景管理器, 获取元素背景管理器.
func (e *Element) I取背景管理器() int {
	return xc.XEle_GetBkManager(e.I句柄)
}

// 元素_取背景管理器扩展, 获取元素背景管理器, 优先从资源中获取.
func (e *Element) I取背景管理器扩展() int {
	return xc.XEle_GetBkManagerEx(e.I句柄)
}

// 元素_置背景管理器.
//
// hBkInfoM: 背景管理器.
func (e *Element) I置背景管理器(hBkInfoM int) int {
	return xc.XEle_SetBkManager(e.I句柄, hBkInfoM)
}

// 元素_取状态, 获取组合状态.
func (e *Element) I取状态() xcc.CombinedState {
	return xc.XEle_GetStateFlags(e.I句柄)
}

// 元素_绘制焦点, 绘制元素焦点.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func (e *Element) I绘制焦点(hDraw int, pRect *xc.RECT) bool {
	return xc.XEle_DrawFocus(e.I句柄, hDraw, pRect)
}

// 元素_绘制, 在自绘事件函数中, 用户手动调用绘制元素, 以便控制绘制顺序.
//
// hDraw: 图形绘制句柄.
func (e *Element) I绘制(hDraw int) int {
	return xc.XEle_DrawEle(e.I句柄, hDraw)
}

// 元素_置用户数据.
//
// nData: 用户数据.
func (e *Element) I置用户数据(nData int) int {
	return xc.XEle_SetUserData(e.I句柄, nData)
}

// 元素_取用户数据.
func (e *Element) I取用户数据() int {
	return xc.XEle_GetUserData(e.I句柄)
}

// 元素_取内容大小.
//
// bHorizon: 水平或垂直, 布局属性交换依赖.
//
// cx: 宽度.
//
// cy: 高度.
//
// pSize: 返回大小.
func (e *Element) I取内容大小(bHorizon bool, cx int, cy int, pSize *xc.SIZE) int {
	return xc.XEle_GetContentSize(e.I句柄, bHorizon, cx, cy, pSize)
}

// 元素_置鼠标捕获.
//
// b: TRUE设置.
func (e *Element) I置鼠标捕获(b bool) int {
	return xc.XEle_SetCapture(e.I句柄, b)
}

// 元素_启用透明通道, 启用或关闭元素透明通道, 如果启用, 将强制设置元素背景不透明, 默认为启用, 此功能是为了兼容GDI不支持透明通道问题.
//
// bEnable: 启用或关闭.
func (e *Element) I启用透明通道(bEnable bool) int {
	return xc.XEle_EnableTransparentChannel(e.I句柄, bEnable)
}

// 元素_置炫彩定时器, 设置元素定时器.
//
// nIDEvent: 事件ID.
//
// uElapse: 延时毫秒.
func (e *Element) I置定时器(nIDEvent int, uElapse int) bool {
	return xc.XEle_SetXCTimer(e.I句柄, nIDEvent, uElapse)
}

// 元素_关闭炫彩定时器, 关闭元素定时器.
//
// nIDEvent: 事件ID.
func (e *Element) I关闭定时器(nIDEvent int) bool {
	return xc.XEle_KillXCTimer(e.I句柄, nIDEvent)
}

// 元素_置工具提示, 设置工具提示内容.
//
// pText: 工具提示内容.
func (e *Element) I置工具提示(pText string) int {
	return xc.XEle_SetToolTip(e.I句柄, pText)
}

// 元素_置工具提示扩展, 设置工具提示内容.
//
// pText: 工具提示内容.
//
// nTextAlign: 文本对齐方式, I常量_文本对齐_, TextAlignFlag_, TextTrimming_.
func (e *Element) I置工具提示扩展(pText string, nTextAlign xcc.I常量_文本对齐_) int {
	return xc.XEle_SetToolTipEx(e.I句柄, pText, nTextAlign)
}

// 元素_取工具提示, 获取工具提示内容.
func (e *Element) I取工具提示() int {
	return xc.XEle_GetToolTip(e.I句柄)
}

// 元素_弹出工具提示, 弹出工具提示.
//
// x: X坐标.
//
// y: Y坐标.
func (e *Element) I弹出工具提示(x int, y int) int {
	return xc.XEle_PopupToolTip(e.I句柄, x, y)
}

// 元素_调整布局.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) I调整布局(nAdjustNo uint32) int {
	return xc.XEle_AdjustLayout(e.I句柄, nAdjustNo)
}

// 元素_调整布局扩展.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.
func (e *Element) I调整布局扩展(nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_AdjustLayoutEx(e.I句柄, nFlags, nAdjustNo)
}

// 元素_取透明度, 返回透明度.
func (e *Element) I取透明度() byte {
	return xc.XEle_GetAlpha(e.I句柄)
}

// 元素_取位置.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.
func (e *Element) I取位置(pOutX *int, pOutY *int) int {
	return xc.XEle_GetPosition(e.I句柄, pOutX, pOutY)
}

// 元素_置大小.
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
func (e *Element) I置大小(nWidth int, nHeight int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {
	return xc.XEle_SetSize(e.I句柄, nWidth, nHeight, bRedraw, nFlags, nAdjustNo)
}

// 元素_取大小.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.
func (e *Element) I取大小(pOutWidth *int, pOutHeight *int) int {
	return xc.XEle_GetSize(e.I句柄, pOutWidth, pOutHeight)
}

// 元素_置背景, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.
func (e *Element) I置背景(pText string) int {
	return xc.XEle_SetBkInfo(e.I句柄, pText)
}

/*
下面都是事件
*/

type XE_ELEPROCE func(nEvent int, wParam int, lParam int, pbHandled *bool) int            // 元素处理过程事件.
type XE_ELEPROCE1 func(hEle int, nEvent int, wParam int, lParam int, pbHandled *bool) int // 元素处理过程事件.
type XE_PAINT func(hDraw int, pbHandled *bool) int                                        // 元素绘制事件.
type XE_PAINT1 func(hEle int, hDraw int, pbHandled *bool) int                             // 元素绘制事件.
type XE_PAINT_END func(hDraw int, pbHandled *bool) int                                    // 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
type XE_PAINT_END1 func(hEle int, hDraw int, pbHandled *bool) int                         // 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
type XE_PAINT_SCROLLVIEW func(hDraw int, pbHandled *bool) int                             // 滚动视图绘制事件.
type XE_PAINT_SCROLLVIEW1 func(hEle int, hDraw int, pbHandled *bool) int                  // 滚动视图绘制事件.
type XE_MOUSEMOVE func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                    // 元素鼠标移动事件.
type XE_MOUSEMOVE1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int         // 元素鼠标移动事件.
type XE_MOUSESTAY func(pbHandled *bool) int                                               // 元素鼠标进入事件.
type XE_MOUSESTAY1 func(hEle int, pbHandled *bool) int                                    // 元素鼠标进入事件.
type XE_MOUSEHOVER func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                   // 元素鼠标悬停事件.
type XE_MOUSEHOVER1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int        // 元素鼠标悬停事件.
type XE_MOUSELEAVE func(hEleStay int, pbHandled *bool) int                                // 元素鼠标离开事件.
type XE_MOUSELEAVE1 func(hEle int, hEleStay int, pbHandled *bool) int                     // 元素鼠标离开事件.
type XE_MOUSEWHEEL func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                   // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL(). flags: 见MSDN中WM_MOUSEWHEEL消息wParam参数说明.
type XE_MOUSEWHEEL1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int        // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL(). flags: 见MSDN中WM_MOUSEWHEEL消息wParam参数说明.
type XE_LBUTTONDOWN func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                  // 鼠标左键按下事件.
type XE_LBUTTONDOWN1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int       // 鼠标左键按下事件.
type XE_LBUTTONUP func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                    // 鼠标左键弹起事件.
type XE_LBUTTONUP1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int         // 鼠标左键弹起事件.
type XE_RBUTTONDOWN func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                  // 鼠标右键按下事件.
type XE_RBUTTONDOWN1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int       // 鼠标右键按下事件.
type XE_RBUTTONUP func(nFlags int, pPt *xc.POINT, pbHandled *bool) int                    // 鼠标右键弹起事件.
type XE_RBUTTONUP1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int         // 鼠标右键弹起事件.
type XE_LBUTTONDBCLICK func(nFlags int, pPt *xc.POINT, pbHandled *bool) int               // 鼠标左键双击事件.
type XE_LBUTTONDBCLICK1 func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int    // 鼠标左键双击事件.
type XE_XC_TIMER func(nTimerID int, pbHandled *bool) int                                  // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
type XE_XC_TIMER1 func(hEle int, nTimerID int, pbHandled *bool) int                       // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
type XE_ADJUSTLAYOUT func(nFlags int, nAdjustNo uint32, pbHandled *bool) int              // 调整布局事件. 暂停使用.
type XE_ADJUSTLAYOUT1 func(hEle int, nFlags int, nAdjustNo uint32, pbHandled *bool) int   // 调整布局事件. 暂停使用.

// 调整布局完成事件.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_ADJUSTLAYOUT_END func(nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

// 调整布局完成事件.
//
// hEle: 元素句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_ADJUSTLAYOUT_END1 func(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

type XE_SETFOCUS func(pbHandled *bool) int               // 元素获得焦点事件.
type XE_SETFOCUS1 func(hEle int, pbHandled *bool) int    // 元素获得焦点事件.
type XE_KILLFOCUS func(pbHandled *bool) int              // 元素失去焦点事件.
type XE_KILLFOCUS1 func(hEle int, pbHandled *bool) int   // 元素失去焦点事件.
type XE_DESTROY func(pbHandled *bool) int                // 元素即将销毁事件. 在销毁子对象之前触发.
type XE_DESTROY1 func(hEle int, pbHandled *bool) int     // 元素即将销毁事件. 在销毁子对象之前触发.
type XE_DESTROY_END func(pbHandled *bool) int            // 元素销毁完成事件. 在销毁子对象之后触发.
type XE_DESTROY_END1 func(hEle int, pbHandled *bool) int // 元素销毁完成事件. 在销毁子对象之后触发.

// 元素大小改变事件.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_SIZE func(nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

// 元素大小改变事件1.
//
// hEle: 元素句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_SIZE1 func(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

type XE_SHOW func(bShow bool, pbHandled *bool) int                             // 元素显示隐藏事件.
type XE_SHOW1 func(hEle int, bShow bool, pbHandled *bool) int                  // 元素显示隐藏事件.
type XE_SETFONT func(pbHandled *bool) int                                      // 元素设置字体事件.
type XE_SETFONT1 func(hEle int, pbHandled *bool) int                           // 元素设置字体事件.
type XE_KEYDOWN func(wParam int, lParam int, pbHandled *bool) int              // 元素按键事件.
type XE_KEYDOWN1 func(hEle int, wParam int, lParam int, pbHandled *bool) int   // 元素按键事件.
type XE_KEYUP func(wParam int, lParam int, pbHandled *bool) int                // 元素按键事件.
type XE_KEYUP1 func(hEle int, wParam int, lParam int, pbHandled *bool) int     // 元素按键事件.
type XE_CHAR func(wParam int, lParam int, pbHandled *bool) int                 // 通过TranslateMessage函数翻译的字符事件.
type XE_CHAR1 func(hEle int, wParam int, lParam int, pbHandled *bool) int      // 通过TranslateMessage函数翻译的字符事件.
type XE_SETCAPTURE func(pbHandled *bool) int                                   // 元素设置鼠标捕获.
type XE_SETCAPTURE1 func(hEle int, pbHandled *bool) int                        // 元素设置鼠标捕获.
type XE_KILLCAPTURE func(pbHandled *bool) int                                  // 元素失去鼠标捕获.
type XE_KILLCAPTURE1 func(hEle int, pbHandled *bool) int                       // 元素失去鼠标捕获.
type XE_SETCURSOR func(wParam int, lParam int, pbHandled *bool) int            // 设置鼠标光标.
type XE_SETCURSOR1 func(hEle int, wParam int, lParam int, pbHandled *bool) int // 设置鼠标光标.
type XE_DROPFILES func(hDropInfo int, pbHandled *bool) int                     // 文件拖放事件, 需先启用:XWnd_EnableDragFiles().
type XE_DROPFILES1 func(hEle int, hDropInfo int, pbHandled *bool) int          // 文件拖放事件, 需先启用:XWnd_EnableDragFiles().

// I事件_元素处理过程 事件.
func (e *Element) I事件_元素处理过程(pFun XE_ELEPROCE) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_ELEPROCE, pFun)
}

// I事件_元素处理过程 事件.
func (e *Element) I事件_元素处理过程1(pFun XE_ELEPROCE1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_ELEPROCE, pFun)
}

// I事件_元素绘制 事件.
func (e *Element) I事件_元素绘制(pFun XE_PAINT) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_PAINT, pFun)
}

// I事件_元素绘制 事件.
func (e *Element) I事件_元素绘制1(pFun XE_PAINT1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_PAINT, pFun)
}

// I事件_元素及子元素绘制完成  事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
func (e *Element) I事件_元素及子元素绘制完成(pFun XE_PAINT_END) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_PAINT_END, pFun)
}

// I事件_元素及子元素绘制完成  事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
func (e *Element) I事件_元素及子元素绘制完成1(pFun XE_PAINT_END1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_PAINT_END, pFun)
}

// I事件_滚动视图绘制 事件.
func (e *Element) I事件_滚动视图绘制(pFun XE_PAINT_SCROLLVIEW) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_PAINT_SCROLLVIEW, pFun)
}

// I事件_滚动视图绘制  事件.
func (e *Element) I事件_滚动视图绘制1(pFun XE_PAINT_SCROLLVIEW1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_PAINT_SCROLLVIEW, pFun)
}

// I事件_鼠标移动  事件.
func (e *Element) I事件_鼠标移动(pFun XE_MOUSEMOVE) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MOUSEMOVE, pFun)
}

// I事件_鼠标移动  事件.
func (e *Element) I事件_鼠标移动1(pFun XE_MOUSEMOVE1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_MOUSEMOVE, pFun)
}

// I事件_鼠标进入  事件.
func (e *Element) I事件_鼠标进入(pFun XE_MOUSESTAY) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MOUSESTAY, pFun)
}

// I事件_鼠标进入  事件.
func (e *Element) I事件_鼠标进入1(pFun XE_MOUSESTAY1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_MOUSESTAY, pFun)
}

// I事件_鼠标悬停  事件.
func (e *Element) I事件_鼠标悬停(pFun XE_MOUSEHOVER) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MOUSEHOVER, pFun)
}

// I事件_鼠标悬停  事件.
func (e *Element) I事件_鼠标悬停1(pFun XE_MOUSEHOVER1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_MOUSEHOVER, pFun)
}

// I事件_鼠标离开  事件.
func (e *Element) I事件_鼠标离开(pFun XE_MOUSELEAVE) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MOUSELEAVE, pFun)
}

// I事件_鼠标离开 事件.
func (e *Element) I事件_鼠标离开1(pFun XE_MOUSELEAVE1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_MOUSELEAVE, pFun)
}

// I事件_鼠标滚轮滚动 事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL().
func (e *Element) I事件_鼠标滚轮滚动(pFun XE_MOUSEWHEEL) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MOUSEWHEEL, pFun)
}

// I事件_鼠标滚轮滚动 事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL().
func (e *Element) I事件_鼠标滚轮滚动1(pFun XE_MOUSEWHEEL1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_MOUSEWHEEL, pFun)
}

// I事件_鼠标左键按下 事件.
func (e *Element) I事件_鼠标左键按下(pFun XE_LBUTTONDOWN) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_LBUTTONDOWN, pFun)
}

// I事件_鼠标左键按下 事件.
func (e *Element) I事件_鼠标左键按下1(pFun XE_LBUTTONDOWN1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_LBUTTONDOWN, pFun)
}

// I事件_鼠标左键弹起 事件.
func (e *Element) I事件_鼠标左键弹起(pFun XE_LBUTTONUP) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_LBUTTONUP, pFun)
}

// I事件_鼠标左键弹起 事件.
func (e *Element) I事件_鼠标左键弹起1(pFun XE_LBUTTONUP1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_LBUTTONUP, pFun)
}

// I事件_鼠标右键按下 事件.
func (e *Element) I事件_鼠标右键按下(pFun XE_RBUTTONDOWN) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_RBUTTONDOWN, pFun)
}

// I事件_鼠标右键按下 事件.
func (e *Element) I事件_鼠标右键按下1(pFun XE_RBUTTONDOWN1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_RBUTTONDOWN, pFun)
}

// I事件_鼠标右键弹起 事件.
func (e *Element) I事件_鼠标右键弹起(pFun XE_RBUTTONUP) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_RBUTTONUP, pFun)
}

// I事件_鼠标右键弹起 事件.
func (e *Element) I事件_鼠标右键弹起1(pFun XE_RBUTTONUP1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_RBUTTONUP, pFun)
}

// I事件_鼠标左键双击 事件.
func (e *Element) I事件_鼠标左键双击(pFun XE_LBUTTONDBCLICK) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_LBUTTONDBCLICK, pFun)
}

// I事件_鼠标左键双击 事件.
func (e *Element) I事件_鼠标左键双击1(pFun XE_LBUTTONDBCLICK1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_LBUTTONDBCLICK, pFun)
}

// I事件_定时器,非系统定时器,定时器消息 XM_TIMER.
func (e *Element) I事件_定时器(pFun XE_XC_TIMER) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_XC_TIMER, pFun)
}

// I事件_定时器,非系统定时器,定时器消息 XM_TIMER.
func (e *Element) I事件_定时器1(pFun XE_XC_TIMER1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_XC_TIMER, pFun)
}

// I事件_调整布局 事件. 暂停使用.
func (e *Element) Event_ADJUSTLAYOUT(pFun XE_ADJUSTLAYOUT) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_ADJUSTLAYOUT, pFun)
}

// I事件_调整布局  事件. 暂停使用.
func (e *Element) Event_ADJUSTLAYOUT1(pFun XE_ADJUSTLAYOUT1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_ADJUSTLAYOUT, pFun)
}

// I事件_调整布局完成  事件.
func (e *Element) I事件_调整布局完成(pFun XE_ADJUSTLAYOUT_END) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_ADJUSTLAYOUT_END, pFun)
}

// I事件_调整布局完成  事件.
func (e *Element) I事件_调整布局完成1(pFun XE_ADJUSTLAYOUT_END1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_ADJUSTLAYOUT_END, pFun)
}

// I事件_获得焦点 事件.
func (e *Element) I事件_获得焦点(pFun XE_SETFOCUS) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_SETFOCUS, pFun)
}

// I事件_获得焦点  事件.
func (e *Element) I事件_获得焦点1(pFun XE_SETFOCUS1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_SETFOCUS, pFun)
}

// I事件_失去焦点  事件.
func (e *Element) I事件_失去焦点(pFun XE_KILLFOCUS) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_KILLFOCUS, pFun)
}

// I事件_失去焦点  事件.
func (e *Element) I事件_失去焦点1(pFun XE_KILLFOCUS1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_KILLFOCUS, pFun)
}

// I事件_即将销毁  事件. 在销毁子对象之前触发.
func (e *Element) I事件_即将销毁(pFun XE_DESTROY) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_DESTROY, pFun)
}

// I事件_即将销毁  事件. 在销毁子对象之前触发.
func (e *Element) I事件_即将销毁1(pFun XE_DESTROY1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_DESTROY, pFun)
}

// I事件_销毁完成  事件. 在销毁子对象之后触发.
func (e *Element) I事件_销毁完成(pFun XE_DESTROY_END) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_DESTROY_END, pFun)
}

// I事件_销毁完成  事件. 在销毁子对象之后触发.
func (e *Element) I事件_销毁完成1(pFun XE_DESTROY_END1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_DESTROY_END, pFun)
}

// I事件_大小改变  事件.
func (e *Element) I事件_大小改变(pFun XE_SIZE) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_SIZE, pFun)
}

// I事件_大小改变  事件.
func (e *Element) I事件_大小改变1(pFun XE_SIZE1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_SIZE, pFun)
}

// I事件_显示隐藏  事件.
func (e *Element) I事件_显示隐藏(pFun XE_SHOW) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_SHOW, pFun)
}

// I事件_显示隐藏  事件.
func (e *Element) I事件_显示隐藏1(pFun XE_SHOW1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_SHOW, pFun)
}

// I事件_设置字体  事件.
func (e *Element) I事件_设置字体(pFun XE_SETFONT) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_SETFONT, pFun)
}

// I事件_设置字体  事件.
func (e *Element) I事件_设置字体1(pFun XE_SETFONT1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_SETFONT, pFun)
}

// I事件_按键按下  事件.
func (e *Element) I事件_按键按下(pFun XE_KEYDOWN) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_KEYDOWN, pFun)
}

// I事件_按键按下  事件.
func (e *Element) I事件_按键按下1(pFun XE_KEYDOWN1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_KEYDOWN, pFun)
}

// I事件_按键弹起  事件.
func (e *Element) I事件_按键弹起(pFun XE_KEYUP) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_KEYUP, pFun)
}

//  I事件_按键弹起 元素按键事件.
func (e *Element) I事件_按键弹起1(pFun XE_KEYUP1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_KEYUP, pFun)
}

// 通过TranslateMessage函数翻译的字符事件.
func (e *Element) Event_CHAR(pFun XE_CHAR) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_CHAR, pFun)
}

// 通过TranslateMessage函数翻译的字符事件.
func (e *Element) Event_CHAR1(pFun XE_CHAR1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_CHAR, pFun)
}

// I事件_设置鼠标捕获  .
func (e *Element) I事件_设置鼠标捕获(pFun XE_SETCAPTURE) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_SETCAPTURE, pFun)
}

// I事件_设置鼠标捕获  .
func (e *Element) I事件_设置鼠标捕获1(pFun XE_SETCAPTURE1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_SETCAPTURE, pFun)
}

// I事件_失去鼠标捕获  .
func (e *Element) I事件_失去鼠标捕获(pFun XE_KILLCAPTURE) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_KILLCAPTURE, pFun)
}

// I事件_失去鼠标捕获1.
func (e *Element) I事件_失去鼠标捕获1(pFun XE_KILLCAPTURE1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_KILLCAPTURE, pFun)
}

// I事件_设置鼠标光标  .
func (e *Element) I事件_设置鼠标光标(pFun XE_SETCURSOR) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_SETCURSOR, pFun)
}

// I事件_设置鼠标光标  .
func (e *Element) I事件_设置鼠标光标1(pFun XE_SETCURSOR1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_SETCURSOR, pFun)
}

// I事件_文件拖放  事件, 需先启用:XWnd_EnableDragFiles().
func (e *Element) I事件_文件拖放(pFun XE_DROPFILES) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_DROPFILES, pFun)
}

// I事件_文件拖放  事件, 需先启用:XWnd_EnableDragFiles().
func (e *Element) I事件_文件拖放1(pFun XE_DROPFILES1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_DROPFILES, pFun)
}

type XE_MENU_SELECT func(nID int, pbHandled *bool) int                                            // 弹出菜单项选择事件.
type XE_MENU_POPUP func(HMENUX int, pbHandled *bool) int                                          // 菜单弹出.
type XE_MENU_EXIT func(pbHandled *bool) int                                                       // 弹出菜单退出事件.
type XE_MENU_POPUP_WND func(hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XE_MENU_DRAW_BACKGROUND func(hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XE_MENU_DRAWITEM func(hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// I事件_弹出菜单项被选择.
func (e *Element) I事件_弹出菜单项被选择(pFun XE_MENU_SELECT) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_SELECT, pFun)
}

// I事件_菜单弹出 .
func (e *Element) I事件_菜单弹出(pFun XE_MENU_POPUP) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_POPUP, pFun)
}

// I事件_菜单退出 .
func (e *Element) I事件_菜单退出(pFun XE_MENU_EXIT) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_EXIT, pFun)
}

// I事件_菜单弹出窗口 .
func (e *Element) I事件_菜单弹出窗口(pFun XE_MENU_POPUP_WND) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_POPUP_WND, pFun)
}

// I事件_绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (e *Element) I事件_绘制菜单背景(pFun XE_MENU_DRAW_BACKGROUND) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_DRAW_BACKGROUND, pFun)
}

// I事件_绘制菜单项  事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (e *Element) I事件_绘制菜单项(pFun XE_MENU_DRAWITEM) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_DRAWITEM, pFun)
}

/* type XE_MENU_SELECT1 func(hEle int, nID int, pbHandled *bool) int                                            // 弹出菜单项选择事件.
type XE_MENU_POPUP1 func(hEle int, HMENUX int, pbHandled *bool) int                                          // 菜单弹出.
type XE_MENU_EXIT1 func(hEle int, pbHandled *bool) int                                                       // 弹出菜单退出事件.
type XE_MENU_POPUP_WND1 func(hEle int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XE_MENU_DRAW_BACKGROUND1 func(hEle int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XE_MENU_DRAWITEM1 func(hEle int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// 事件_弹出菜单项被选择.
func (e *Element) Event_MENU_SELECT1(pFun XE_MENU_SELECT1) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_SELECT, pFun)
}

// 事件_菜单弹出.
func (e *Element) Event_MENU_POPUP1(pFun XE_MENU_POPUP1) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_POPUP, pFun)
}

// 事件_菜单退出.
func (e *Element) Event_MENU_EXIT1(pFun XE_MENU_EXIT1) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_MENU_EXIT, pFun)
}

// 菜单弹出窗口.
func (e *Element) Event_MENU_POPUP_WND1(pFun XE_MENU_POPUP_WND1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_MENU_POPUP_WND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (e *Element) Event_MENU_DRAW_BACKGROUND1(pFun XE_MENU_DRAW_BACKGROUND1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (e *Element) Event_MENU_DRAWITEM1(pFun XE_MENU_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_MENU_DRAWITEM, pFun)
} */
