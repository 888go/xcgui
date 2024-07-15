package 炫彩组件类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Element 基础元素.
type Element struct {
	炫彩对象基类.Widget
}

// 元素_创建, 创建基础元素.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.

// ff:创建基础元素
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:y坐标
// x:x坐标
func X创建基础元素(x坐标, y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄 int) *Element {
	p := &Element{}
	p.X设置句柄(炫彩基类.X元素_创建(x坐标, y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.

// ff:创建基础元素并按句柄
// handle:
func X创建基础元素并按句柄(handle int) *Element {
	p := &Element{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建基础元素并按名称
// name:
func X创建基础元素并按名称(name string) *Element {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &Element{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建基础元素并按UID
// nUID:
func X创建基础元素并按UID(nUID int) *Element {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Element{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建基础元素并按UID名称
// name:
func X创建基础元素并按UID名称(name string) *Element {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &Element{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 元素_注册事件C, 注册事件C方式, 省略2参数.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.

// ff:注册事件C
// pFun:
// nEvent:事件类型
func (e *Element) X注册事件C(事件类型 炫彩常量类.XE_, pFun interface{}) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 事件类型, pFun)
}

// 元素_注册事件C1, 注册事件C1方式, 省略1参数.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.

// ff:注册事件C1
// pFun:
// nEvent:事件类型
func (e *Element) X注册事件C1(事件类型 炫彩常量类.XE_, pFun interface{}) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 事件类型, pFun)
}

// 元素_移除事件C.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数.

// ff:移除事件C
// pFun:
// nEvent:事件类型
func (e *Element) X移除事件C(事件类型 炫彩常量类.XE_, pFun interface{}) bool {
	return 炫彩基类.X元素_移除事件C(e.Handle, 事件类型, pFun)
}

// 元素_注册事件CEx, 注册事件C方式, 省略2参数, 和非Ex版相比只是最后一个参数不同.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.

// ff:注册事件CEx
// pFun:
// nEvent:事件类型
func (e *Element) X注册事件CEx(事件类型 炫彩常量类.XE_, pFun uintptr) bool {
	return 炫彩基类.X元素_注册事件CEx(e.Handle, 事件类型, pFun)
}

// 元素_注册事件C1Ex, 注册事件C1方式, 省略1参数, 和非Ex版相比只是最后一个参数不同.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.

// ff:注册事件C1Ex
// pFun:
// nEvent:事件类型
func (e *Element) X注册事件C1Ex(事件类型 炫彩常量类.XE_, pFun uintptr) bool {
	return 炫彩基类.X元素_注册事件C1Ex(e.Handle, 事件类型, pFun)
}

// 元素_移除事件CEx, 和非Ex版相比只是最后一个参数不同.
//
// nEvent: 事件类型: xcc.XE_.
//
// pFun: 事件函数指针, 使用 syscall.NewCallback() 生成.

// ff:移除事件CEx
// pFun:
// nEvent:事件类型
func (e *Element) X移除事件CEx(事件类型 炫彩常量类.XE_, pFun uintptr) bool {
	return 炫彩基类.X元素_移除事件CEx(e.Handle, 事件类型, pFun)
}

// 元素_发送事件.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.

// ff:发送事件
// lParam:
// wParam:
// nEvent:事件类型
func (e *Element) X发送事件(事件类型 炫彩常量类.XE_, wParam, lParam uint) int {
	return 炫彩基类.X元素_发送事件(e.Handle, 事件类型, wParam, lParam)
}

// 元素_投递事件.
//
// nEvent: 事件类型: xcc.XE_.
//
// wParam: 参数.
//
// lParam: 参数.

// ff:投递事件
// lParam:
// wParam:
// nEvent:事件类型
func (e *Element) X投递事件(事件类型 炫彩常量类.XE_, wParam, lParam uint) int {
	return 炫彩基类.X元素_投递事件(e.Handle, 事件类型, wParam, lParam)
}

// 元素_取坐标.
//
// pRect: 坐标.

// ff:取坐标
// pRect:坐标
func (e *Element) X取坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_取坐标(e.Handle, 坐标)
}

// 元素_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.

// ff:取逻辑坐标
// pRect:坐标
func (e *Element) X取逻辑坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_取逻辑坐标(e.Handle, 坐标)
}

// 元素_取客户区坐标.
//
// pRect: 坐标.

// ff:取客户区坐标
// pRect:坐标
func (e *Element) X取客户区坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_取客户区坐标(e.Handle, 坐标)
}

// 元素_置宽度.
//
// nWidth: 宽度.

// ff:置宽度
// nWidth:宽度
func (e *Element) X置宽度(宽度 int) int {
	return 炫彩基类.X元素_置宽度(e.Handle, 宽度)
}

// 元素_置高度.
//
// nHeight: 高度.

// ff:置高度
// nHeight:高度
func (e *Element) X置高度(高度 int) int {
	return 炫彩基类.X元素_置高度(e.Handle, 高度)
}

// 元素_取宽度.

// ff:取宽度
func (e *Element) X取宽度() int {
	return 炫彩基类.X元素_取宽度(e.Handle)
}

// 元素_取高度.

// ff:取高度
func (e *Element) X取高度() int {
	return 炫彩基类.X元素_取高度(e.Handle)
}

// 元素_窗口客户区坐标到元素客户区坐标, 窗口客户区坐标转换到元素客户区坐标.
//
// pRect: 坐标.

// ff:窗口客户区坐标到元素客户区坐标
// pRect:坐标
func (e *Element) X窗口客户区坐标到元素客户区坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_窗口客户区坐标到元素客户区坐标(e.Handle, 坐标)
}

// 元素_窗口客户区点到元素客户区, 窗口客户区坐标转换到元素客户区坐标.
//
// pPt: 坐标.

// ff:窗口客户区点到元素客户区
// pPt:坐标
func (e *Element) X窗口客户区点到元素客户区(坐标 *炫彩基类.POINT) int {
	return 炫彩基类.X元素_窗口客户区点到元素客户区(e.Handle, 坐标)
}

// 元素_客户区坐标到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// pRect: 坐标.

// ff:客户区坐标到窗口客户区
// pRect:坐标
func (e *Element) X客户区坐标到窗口客户区(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_客户区坐标到窗口客户区(e.Handle, 坐标)
}

// 元素_客户区点到窗口客户区, 元素客户区坐标转换到窗口客户区坐标.
//
// pPt: 坐标.

// ff:客户区点到窗口客户区
// pPt:坐标
func (e *Element) X客户区点到窗口客户区(坐标 *炫彩基类.POINT) int {
	return 炫彩基类.X元素_客户区点到窗口客户区(e.Handle, 坐标)
}

// 元素_基于窗口客户区坐标.
//
// pRect: 坐标.

// ff:取窗口客户区坐标
// pRect:坐标
func (e *Element) X取窗口客户区坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_基于窗口客户区坐标(e.Handle, 坐标)
}

// 元素_取光标, 获取元素鼠标光标, 返回光标句柄.

// ff:取光标
func (e *Element) X取光标() int {
	return 炫彩基类.X元素_取光标(e.Handle)
}

// 元素_置光标, 设置元素鼠标光标.
//
// hCursor: 光标句柄.

// ff:置光标
// hCursor:光标句柄
func (e *Element) X置光标(光标句柄 int) int {
	return 炫彩基类.X元素_置光标(e.Handle, 光标句柄)
}

// 元素_添加子对象.
//
// hChild: 要添加的子元素句柄或形状对象句柄.

// ff:添加子对象
// hChild:子元素句柄或形状对象句柄
func (e *Element) X添加子对象(子元素句柄或形状对象句柄 int) bool {
	return 炫彩基类.X元素_添加子对象(e.Handle, 子元素句柄或形状对象句柄)
}

// 元素_插入子对象, 插入子对象到指定位置.
//
// hChild: 要插入的元素句柄或形状对象句柄.
//
// index: 插入位置索引.

// ff:插入子对象
// index:插入位置索引
// hChild:元素句柄或形状对象句柄
func (e *Element) X插入子对象(元素句柄或形状对象句柄 int, 插入位置索引 int) bool {
	return 炫彩基类.X元素_插入子对象(e.Handle, 元素句柄或形状对象句柄, 插入位置索引)
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

// ff:置坐标
// nAdjustNo:
// nFlags:
// bRedraw:
// pRect:坐标
func (e *Element) X置坐标(坐标 *炫彩基类.RECT, bRedraw bool, nFlags 炫彩常量类.AdjustLayout_, nAdjustNo uint32) int {
	return 炫彩基类.X元素_置坐标(e.Handle, 坐标, bRedraw, nFlags, nAdjustNo)
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

// ff:置坐标EX
// nAdjustNo:
// nFlags:标识
// bRedraw:是否重绘
// cy:高度
// cx:宽度
// y:Y坐标
// x:坐标
func (e *Element) X置坐标EX(坐标 int, Y坐标 int, 宽度 int, 高度 int, 是否重绘 bool, 标识 炫彩常量类.AdjustLayout_, nAdjustNo uint32) int {
	return 炫彩基类.X元素_置坐标EX(e.Handle, 坐标, Y坐标, 宽度, 高度, 是否重绘, 标识, nAdjustNo)
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

// ff:置逻辑坐标
// nAdjustNo:
// nFlags:
// bRedraw:
// pRect:坐标
func (e *Element) X置逻辑坐标(坐标 *炫彩基类.RECT, bRedraw bool, nFlags 炫彩常量类.AdjustLayout_, nAdjustNo uint32) int {
	return 炫彩基类.X元素_置逻辑坐标(e.Handle, 坐标, bRedraw, nFlags, nAdjustNo)
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

// ff:移动
// nAdjustNo:
// nFlags:标识
// bRedraw:是否重绘
// y:Y坐标
// x:坐标
func (e *Element) X移动(坐标, Y坐标 int32, 是否重绘 bool, 标识 炫彩常量类.AdjustLayout_, nAdjustNo uint32) int {
	return 炫彩基类.X元素_移动(e.Handle, 坐标, Y坐标, 是否重绘, 标识, nAdjustNo)
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

// ff:移动逻辑坐标
// nAdjustNo:
// nFlags:标识
// bRedraw:是否重绘
// y:Y坐标
// x:坐标
func (e *Element) X移动逻辑坐标(坐标, Y坐标 int32, 是否重绘 bool, 标识 炫彩常量类.AdjustLayout_, nAdjustNo uint32) int {
	return 炫彩基类.X元素_移动逻辑坐标(e.Handle, 坐标, Y坐标, 是否重绘, 标识, nAdjustNo)
}

// 元素_判断绘制焦点.

// ff:判断绘制焦点
func (e *Element) X判断绘制焦点() bool {
	return 炫彩基类.X元素_判断绘制焦点(e.Handle)
}

// 元素_判断启用, 元素是否为启用状态.

// ff:判断启用
func (e *Element) X判断启用() bool {
	return 炫彩基类.X元素_判断启用(e.Handle)
}

// 元素_判断启用焦点, 元素是否启用焦点.

// ff:判断启用焦点
func (e *Element) X判断启用焦点() bool {
	return 炫彩基类.X元素_判断启用焦点(e.Handle)
}

// 元素_判断鼠标穿透, 元素是否启用鼠标穿透.

// ff:判断鼠标穿透
func (e *Element) X判断鼠标穿透() bool {
	return 炫彩基类.X元素_判断鼠标穿透(e.Handle)
}

// 元素_测试点击元素, 检测坐标点所在元素, 包含子元素的子元素.
//
// pPt: 坐标点.

// ff:测试点击元素
// pPt:坐标点
func (e *Element) X测试点击元素(坐标点 *炫彩基类.POINT) int {
	return 炫彩基类.X元素_测试点击元素(e.Handle, 坐标点)
}

// 元素_判断背景透明.

// ff:判断背景透明
func (e *Element) X判断背景透明() bool {
	return 炫彩基类.X元素_判断背景透明(e.Handle)
}

// 元素_判断启用事件_XE_PAINT_END, 是否启XE_PAINT_END用事件.

// ff:判断启用事件_XE_PAINT_END
func (e *Element) X判断启用事件_XE_PAINT_END() bool {
	return 炫彩基类.X元素_判断启用事件_XE_PAINT_END(e.Handle)
}

// 元素_判断接受TAB, 是否接受Tab键输入; 例如: XRichEdit, XEdit.

// ff:判断接受TAB
func (e *Element) X判断接受TAB() bool {
	return 炫彩基类.X元素_判断接受TAB(e.Handle)
}

// 元素_判断接受切换焦点, 是否接受通过键盘切换焦点(方向键,TAB键).

// ff:判断接受切换焦点
func (e *Element) X判断接受切换焦点() bool {
	return 炫彩基类.X元素_判断接受切换焦点(e.Handle)
}

// 元素_判断启用_XE_MOUSEWHEEL, 判断是否启用鼠标滚动事件, 如果禁用那么事件会发送给他的父元素.

// ff:判断启用_XE_MOUSEWHEEL
func (e *Element) X判断启用_XE_MOUSEWHEEL() bool {
	return 炫彩基类.X元素_判断启用_XE_MOUSEWHEEL(e.Handle)
}

// 元素_判断为子元素, 判断hChildEle是否为hEle的子元素.
//
// hChildEle: 子元素句柄.

// ff:判断为子元素
// hChildEle:子元素句柄
func (e *Element) X判断为子元素(子元素句柄 int) bool {
	return 炫彩基类.X元素_判断为子元素(e.Handle, 子元素句柄)
}

// 元素_判断启用画布, 判断是否启用画布.

// ff:判断启用画布
func (e *Element) X判断启用画布() bool {
	return 炫彩基类.X元素_判断启用画布(e.Handle)
}

// 元素_判断焦点, 判断是否拥有焦点.

// ff:判断焦点
func (e *Element) X判断焦点() bool {
	return 炫彩基类.X元素_判断焦点(e.Handle)
}

// 元素_判断焦点扩展, 判断该元素或该元素的子元素是否拥有焦点.

// ff:判断焦点EX
func (e *Element) X判断焦点EX() bool {
	return 炫彩基类.X元素_判断焦点EX(e.Handle)
}

// 元素_启用, 启用或禁用元素.
//
// bEnable: 启用或禁用.

// ff:启用
// bEnable:启用或禁用
func (e *Element) X启用(启用或禁用 bool) int {
	return 炫彩基类.X元素_启用(e.Handle, 启用或禁用)
}

// 元素_启用焦点, 启用焦点.
//
// bEnable: 是否启用.

// ff:启用焦点
// bEnable:是否启用
func (e *Element) X启用焦点(是否启用 bool) int {
	return 炫彩基类.X元素_启用焦点(e.Handle, 是否启用)
}

// 元素_启用绘制焦点.
//
// bEnable: 是否启用.

// ff:启用绘制焦点
// bEnable:是否启用
func (e *Element) X启用绘制焦点(是否启用 bool) int {
	return 炫彩基类.X元素_启用绘制焦点(e.Handle, 是否启用)
}

// 元素_启用绘制边框, 启用或禁用绘制默认边框.
//
// bEnable: 是否启用.

// ff:启用绘制边框
// bEnable:是否启用
func (e *Element) X启用绘制边框(是否启用 bool) int {
	return 炫彩基类.X元素_启用绘制边框(e.Handle, 是否启用)
}

// 元素_启用画布, 启用或禁用背景画布; 如果禁用那么将绘制在父的画布之上, 也就是说他没有自己的画布.
//
// bEnable: 是否启用.

// ff:启用画布
// bEnable:是否启用
func (e *Element) X启用画布(是否启用 bool) int {
	return 炫彩基类.X元素_启用画布(e.Handle, 是否启用)
}

// 元素_启用事件_XE_PAINT_END.
//
// bEnable: 是否启用.

// ff:启用事件_XE_PAINT_END
// bEnable:是否启用
func (e *Element) X启用事件_XE_PAINT_END(是否启用 bool) int {
	return 炫彩基类.X元素_启用事件_XE_PAINT_END(e.Handle, 是否启用)
}

// 元素_启用背景透明.
//
// bEnable: 是否启用.

// ff:启用背景透明
// bEnable:是否启用
func (e *Element) X启用背景透明(是否启用 bool) int {
	return 炫彩基类.X元素_启用背景透明(e.Handle, 是否启用)
}

// 元素_启用鼠标穿透. 启用鼠标穿透, 如果启用, 那么该元素不能接收到鼠标事件, 但是他的子元素不受影响, 任然可以接收鼠标事件.
//
// bEnable: 是否启用.

// ff:启用鼠标穿透
// bEnable:是否启用
func (e *Element) X启用鼠标穿透(是否启用 bool) int {
	return 炫彩基类.X元素_启用鼠标穿透(e.Handle, 是否启用)
}

// 元素_启用接收TAB, 启用接收Tab输入.
//
// bEnable: 是否启用.

// ff:启用接收TAB
// bEnable:是否启用
func (e *Element) X启用接收TAB(是否启用 bool) int {
	return 炫彩基类.X元素_启用接收TAB(e.Handle, 是否启用)
}

// 元素_启用切换焦点, 启用接受通过键盘切换焦点.
//
// bEnable: 是否启用.

// ff:启用切换焦点
// bEnable:是否启用
func (e *Element) X启用切换焦点(是否启用 bool) int {
	return 炫彩基类.X元素_启用切换焦点(e.Handle, 是否启用)
}

// 元素_启用事件_XE_MOUSEWHEEL, 启用接收鼠标滚动事件, 如果禁用那么事件会传递给父元素.
//
// bEnable: 是否启用.

// ff:启用事件_XE_MOUSEWHEEL
// bEnable:是否启用
func (e *Element) X启用事件_XE_MOUSEWHEEL(是否启用 bool) int {
	return 炫彩基类.X元素_启用事件_XE_MOUSEWHEEL(e.Handle, 是否启用)
}

// 元素_移除, 移除元素, 但不销毁.

// ff:移除
func (e *Element) X移除() int {
	return 炫彩基类.X元素_移除(e.Handle)
}

// 元素_置Z序, 设置元素Z序.
//
// index: 位置索引.

// ff:置Z序
// index:位置索引
func (e *Element) X置Z序(位置索引 int) bool {
	return 炫彩基类.X元素_置Z序(e.Handle, 位置索引)
}

// 元素_置Z序扩展, 设置元素Z序.
//
// hDestEle: 目标元素.
//
// nType: 类型, Zorder_.

// ff:置Z序EX
// nType:类型
// hDestEle:目标元素
func (e *Element) X置Z序EX(目标元素 int, 类型 炫彩常量类.Zorder_) bool {
	return 炫彩基类.X元素_置Z序EX(e.Handle, 目标元素, 类型)
}

// 元素_取Z序, 获取元素Z序索引, 位置索引.

// ff:取Z序
func (e *Element) X取Z序() int {
	return 炫彩基类.X元素_取Z序(e.Handle)
}

// 元素_启用置顶, 设置元素置顶.
//
// bTopmost: 是否置顶显示.

// ff:启用置顶
// bTopmost:是否置顶显示
func (e *Element) X启用置顶(是否置顶显示 bool) bool {
	return 炫彩基类.X元素_启用置顶(e.Handle, 是否置顶显示)
}

// 元素_重绘.
//
// bImmediate: 是否立即重绘.

// ff:重绘
// bImmediate:是否立即重绘
func (e *Element) X重绘(是否立即重绘 bool) int {
	return 炫彩基类.X元素_重绘(e.Handle, 是否立即重绘)
}

// 元素_重绘指定区域.
//
// pRect: 相对于元素客户区坐标.
//
// bImmediate: 是否立即重绘.

// ff:重绘指定区域
// bImmediate:
// pRect:坐标
func (e *Element) X重绘指定区域(坐标 *炫彩基类.RECT, bImmediate bool) int {
	return 炫彩基类.X元素_重绘指定区域(e.Handle, 坐标, bImmediate)
}

// 元素_取子对象数量, 获取子对象(UI元素和形状对象)数量, 只检测当前层子对象.

// ff:取子对象数量
func (e *Element) X取子对象数量() int {
	return 炫彩基类.X元素_取子对象数量(e.Handle)
}

// 元素_取子对象从索引, 获取子对象通过索引, 只检测当前层子对象.
//
// index: 索引.

// ff:取子对象并按索引
// index:索引
func (e *Element) X取子对象并按索引(索引 int) int {
	return 炫彩基类.X元素_取子对象从索引(e.Handle, 索引)
}

// 元素_取子对象从ID, 获取子对象通过ID, 只检测当前层子对象.
//
// nID: 元素ID.

// ff:取子对象并按ID
// nID:元素ID
func (e *Element) X取子对象并按ID(元素ID int) int {
	return 炫彩基类.X元素_取子对象从ID(e.Handle, 元素ID)
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

// ff:置边框大小
// bottom:下边
// right:右边
// top:上边
// left:左边
func (e *Element) X置边框大小(左边 int, 上边 int, 右边 int, 下边 int) int {
	return 炫彩基类.X元素_置边框大小(e.Handle, 左边, 上边, 右边, 下边)
}

// 元素_取边框大小.
//
// pBorder: 大小.

// ff:取边框大小
// pBorder:大小
func (e *Element) X取边框大小(大小 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_取边框大小(e.Handle, 大小)
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

// ff:置内填充大小
// bottom:下边
// right:右边
// top:上边
// left:左边
func (e *Element) X置内填充大小(左边 int, 上边 int, 右边 int, 下边 int) int {
	return 炫彩基类.X元素_置内填充大小(e.Handle, 左边, 上边, 右边, 下边)
}

// 元素_取内填充大小.
//
// pPadding: 大小.

// ff:取内填充大小
// pPadding:大小
func (e *Element) X取内填充大小(大小 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_取内填充大小(e.Handle, 大小)
}

// 元素_置拖动边框.
//
// nFlags: 边框位置组合, Element_Position_.

// ff:置拖动边框
// nFlags:边框位置组合
func (e *Element) X置拖动边框(边框位置组合 炫彩常量类.Element_Position_) int {
	return 炫彩基类.X元素_置拖动边框(e.Handle, 边框位置组合)
}

// 元素_置拖动边框绑定元素, 设置拖动边框绑定元素, 当拖动边框时, 自动调整绑定元素的大小.
//
// nFlags: 边框位置标识, Element_Position_.
//
// hBindEle: 绑定元素.
//
// nSpace: 元素间隔大小.

// ff:置拖动边框绑定元素
// nSpace:
// hBindEle:
// nFlags:边框位置标识
func (e *Element) X置拖动边框绑定元素(边框位置标识 炫彩常量类.Element_Position_, hBindEle int, nSpace int) int {
	return 炫彩基类.X元素_置拖动边框绑定元素(e.Handle, 边框位置标识, hBindEle, nSpace)
}

// 元素_置最小大小.
//
// nWidth: 最小宽度.
//
// nHeight: 最小高度.

// ff:置最小大小
// nHeight:最小高度
// nWidth:最小宽度
func (e *Element) X置最小大小(最小宽度 int, 最小高度 int) int {
	return 炫彩基类.X元素_置最小大小(e.Handle, 最小宽度, 最小高度)
}

// 元素_置最大大小.
//
// nWidth: 最大宽度.
//
// nHeight: 最大高度.

// ff:置最大大小
// nHeight:最大高度
// nWidth:最大宽度
func (e *Element) X置最大大小(最大宽度 int, 最大高度 int) int {
	return 炫彩基类.X元素_置最大大小(e.Handle, 最大宽度, 最大高度)
}

// 元素_置锁定滚动, 设置锁定元素在滚动视图中跟随滚动, 如果设置TRUE将不跟随滚动.
//
// bHorizon: 是否锁定水平滚动.
//
// bVertical: 是否锁定垂直滚动.

// ff:置锁定滚动
// bVertical:是否锁定垂直
// bHorizon:是否锁定水平
func (e *Element) X置锁定滚动(是否锁定水平 bool, 是否锁定垂直 bool) int {
	return 炫彩基类.X元素_置锁定滚动(e.Handle, 是否锁定水平, 是否锁定垂直)
}

// 元素_置文本颜色.
//
// color: ABGR 颜色值.

// ff:置文本颜色
// color:ABGR颜色值
func (e *Element) X置文本颜色(ABGR颜色值 int) int {
	return 炫彩基类.X元素_置文本颜色(e.Handle, ABGR颜色值)
}

// 元素_取文本颜色.

// ff:取文本颜色
func (e *Element) X取文本颜色() int {
	return 炫彩基类.X元素_取文本颜色(e.Handle)
}

// 元素_取文本颜色扩展, 获取文本颜色, 优先从资源中获取.

// ff:取文本颜色EX
func (e *Element) X取文本颜色EX() int {
	return 炫彩基类.X元素_取文本颜色EX(e.Handle)
}

// 元素_置焦点边框颜色.
//
// color: ABGR 颜色值.

// ff:置焦点边框颜色
// color:ABGR颜色值
func (e *Element) X置焦点边框颜色(ABGR颜色值 int) int {
	return 炫彩基类.X元素_置焦点边框颜色(e.Handle, ABGR颜色值)
}

// 元素_取焦点边框颜色.

// ff:取焦点边框颜色
func (e *Element) X取焦点边框颜色() int {
	return 炫彩基类.X元素_取焦点边框颜色(e.Handle)
}

// 元素_置字体.
//
// hFontx: 炫彩字体.

// ff:置字体
// hFontx:炫彩字体
func (e *Element) X置字体(炫彩字体 int) int {
	return 炫彩基类.X元素_置字体(e.Handle, 炫彩字体)
}

// 元素_取字体.

// ff:取字体
func (e *Element) X取字体() int {
	return 炫彩基类.X元素_取字体(e.Handle)
}

// 元素_取字体扩展, 获取元素字体, 优先从资源中获取.

// ff:取字体EX
func (e *Element) X取字体EX() int {
	return 炫彩基类.X元素_取字体EX(e.Handle)
}

// 元素_置透明度.

// ff:置透明度
// alpha:透明度
func (e *Element) X置透明度(透明度 uint8) int {
	return 炫彩基类.X元素_置透明度(e.Handle, 透明度)
}

// 元素_销毁.

// ff:销毁
func (e *Element) X销毁() int {
	return 炫彩基类.X元素_销毁(e.Handle)
}

// 元素_添加背景边框, 添加背景内容边框.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.

// ff:添加背景边框
// width:
// color:
// nState:组合状态
func (e *Element) X添加背景边框(组合状态 炫彩常量类.CombinedState, color int, width int) int {
	return 炫彩基类.X元素_添加背景边框(e.Handle, 组合状态, color, width)
}

// 元素_添加背景填充, 添加背景内容填充.
//
// nState: 组合状态.
//
// color: ABGR 颜色.

// ff:添加背景填充
// color:
// nState:组合状态
func (e *Element) X添加背景填充(组合状态 炫彩常量类.CombinedState, color int) int {
	return 炫彩基类.X元素_添加背景填充(e.Handle, 组合状态, color)
}

// 元素_添加背景图片, 添加背景内容图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.

// ff:添加背景图片
// hImage:
// nState:组合状态
func (e *Element) X添加背景图片(组合状态 炫彩常量类.CombinedState, hImage int) int {
	return 炫彩基类.X元素_添加背景图片(e.Handle, 组合状态, hImage)
}

// 元素_取背景对象数量, 获取背景内容数量.

// ff:取背景对象数量
func (e *Element) X取背景对象数量() int {
	return 炫彩基类.X元素_取背景对象数量(e.Handle)
}

// 元素_清空背景对象, 清空背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.

// ff:清空背景对象
func (e *Element) X清空背景对象() int {
	return 炫彩基类.X元素_清空背景对象(e.Handle)
}

// 元素_取背景管理器, 获取元素背景管理器.

// ff:取背景管理器
func (e *Element) X取背景管理器() int {
	return 炫彩基类.X元素_取背景管理器(e.Handle)
}

// 元素_取背景管理器扩展, 获取元素背景管理器, 优先从资源中获取.

// ff:取背景管理器EX
func (e *Element) X取背景管理器EX() int {
	return 炫彩基类.X元素_取背景管理器EX(e.Handle)
}

// 元素_置背景管理器.
//
// hBkInfoM: 背景管理器.

// ff:置背景管理器
// hBkInfoM:背景管理器
func (e *Element) X置背景管理器(背景管理器 int) int {
	return 炫彩基类.X元素_置背景管理器(e.Handle, 背景管理器)
}

// 元素_取状态, 获取组合状态.

// ff:取状态
func (e *Element) X取状态() 炫彩常量类.CombinedState {
	return 炫彩基类.X元素_取状态(e.Handle)
}

// 元素_绘制焦点, 绘制元素焦点.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.

// ff:绘制焦点
// pRect:区域坐标
// hDraw:图形绘制句柄
func (e *Element) X绘制焦点(图形绘制句柄 int, 区域坐标 *炫彩基类.RECT) bool {
	return 炫彩基类.X元素_绘制焦点(e.Handle, 图形绘制句柄, 区域坐标)
}

// 元素_绘制, 在自绘事件函数中, 用户手动调用绘制元素, 以便控制绘制顺序.
//
// hDraw: 图形绘制句柄.

// ff:绘制
// hDraw:图形绘制句柄
func (e *Element) X绘制(图形绘制句柄 int) int {
	return 炫彩基类.X元素_绘制(e.Handle, 图形绘制句柄)
}

// 元素_置用户数据.
//
// nData: 用户数据.

// ff:置用户数据
// nData:用户数据
func (e *Element) X置用户数据(用户数据 int) int {
	return 炫彩基类.X元素_置用户数据(e.Handle, 用户数据)
}

// 元素_取用户数据.

// ff:取用户数据
func (e *Element) X取用户数据() int {
	return 炫彩基类.X元素_取用户数据(e.Handle)
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

// ff:取内容大小
// pSize:返回大小
// cy:高度
// cx:宽度
// bHorizon:水平或垂直
func (e *Element) X取内容大小(水平或垂直 bool, 宽度 int, 高度 int, 返回大小 *炫彩基类.SIZE) int {
	return 炫彩基类.X元素_取内容大小(e.Handle, 水平或垂直, 宽度, 高度, 返回大小)
}

// 元素_置鼠标捕获.
//
// b: TRUE设置.

// ff:置鼠标捕获
// b:开启
func (e *Element) X置鼠标捕获(开启 bool) int {
	return 炫彩基类.X元素_置鼠标捕获(e.Handle, 开启)
}

// 元素_启用透明通道, 启用或关闭元素透明通道, 如果启用, 将强制设置元素背景不透明, 默认为启用, 此功能是为了兼容GDI不支持透明通道问题.
//
// bEnable: 启用或关闭.

// ff:启用透明通道
// bEnable:启用或关闭
func (e *Element) X启用透明通道(启用或关闭 bool) int {
	return 炫彩基类.X元素_启用透明通道(e.Handle, 启用或关闭)
}

// 元素_置炫彩定时器, 设置元素定时器.
//
// nIDEvent: 事件ID.
//
// uElapse: 延时毫秒.

// ff:置炫彩定时器
// uElapse:延时毫秒
// nIDEvent:事件ID
func (e *Element) X置炫彩定时器(事件ID int, 延时毫秒 int) bool {
	return 炫彩基类.X元素_置炫彩定时器(e.Handle, 事件ID, 延时毫秒)
}

// 元素_关闭炫彩定时器, 关闭元素定时器.
//
// nIDEvent: 事件ID.

// ff:关闭炫彩定时器
// nIDEvent:事件ID
func (e *Element) X关闭炫彩定时器(事件ID int) bool {
	return 炫彩基类.X元素_关闭炫彩定时器(e.Handle, 事件ID)
}

// 元素_置工具提示, 设置工具提示内容.
//
// pText: 工具提示内容.

// ff:置工具提示
// pText:内容
func (e *Element) X置工具提示(内容 string) int {
	return 炫彩基类.X元素_置工具提示(e.Handle, 内容)
}

// 元素_置工具提示扩展, 设置工具提示内容.
//
// pText: 工具提示内容.
//
// nTextAlign: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.

// ff:置工具提示EX
// nTextAlign:对齐方式
// pText:内容
func (e *Element) X置工具提示EX(内容 string, 对齐方式 炫彩常量类.TextFormatFlag_) int {
	return 炫彩基类.X元素_置工具提示EX(e.Handle, 内容, 对齐方式)
}

// 元素_取工具提示, 获取工具提示内容.

// ff:取工具提示
func (e *Element) X取工具提示() int {
	return 炫彩基类.X元素_取工具提示(e.Handle)
}

// 元素_弹出工具提示, 弹出工具提示.
//
// x: X坐标.
//
// y: Y坐标.

// ff:弹出工具提示
// y:Y坐标
// x:坐标
func (e *Element) X弹出工具提示(坐标 int, Y坐标 int) int {
	return 炫彩基类.X元素_弹出工具提示(e.Handle, 坐标, Y坐标)
}

// 元素_调整布局.
//
// nAdjustNo: 调整布局流水号, 可填0.

// ff:调整布局
// nAdjustNo:流水号
func (e *Element) X调整布局(流水号 uint32) int {
	return 炫彩基类.X元素_调整布局(e.Handle, 流水号)
}

// 元素_调整布局扩展.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号, 可填0.

// ff:调整布局EX
// nAdjustNo:
// nFlags:标识
func (e *Element) X调整布局EX(标识 炫彩常量类.AdjustLayout_, nAdjustNo uint32) int {
	return 炫彩基类.X元素_调整布局EX(e.Handle, 标识, nAdjustNo)
}

// 元素_取透明度, 返回透明度.

// ff:取透明度
func (e *Element) X取透明度() byte {
	return 炫彩基类.X元素_取透明度(e.Handle)
}

// 元素_取位置.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.

// ff:取位置
// pOutY:返回Y坐标
// pOutX:返回X坐标
func (e *Element) X取位置(返回X坐标, 返回Y坐标 *int32) int {
	return 炫彩基类.X元素_取位置(e.Handle, 返回X坐标, 返回Y坐标)
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

// ff:置大小
// nAdjustNo:
// nFlags:标识
// bRedraw:是否重绘
// nHeight:高度
// nWidth:宽度
func (e *Element) X置大小(宽度, 高度 int32, 是否重绘 bool, 标识 炫彩常量类.AdjustLayout_, nAdjustNo uint32) int {
	return 炫彩基类.X元素_置大小(e.Handle, 宽度, 高度, 是否重绘, 标识, nAdjustNo)
}

// 元素_取大小.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.

// ff:取大小
// pOutHeight:返回高度
// pOutWidth:返回宽度
func (e *Element) X取大小(返回宽度, 返回高度 *int32) int {
	return 炫彩基类.X元素_取大小(e.Handle, 返回宽度, 返回高度)
}

// 元素_置背景, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.

// ff:置背景
// pText:字符串
func (e *Element) X置背景(字符串 string) int {
	return 炫彩基类.X元素_置背景(e.Handle, 字符串)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// pRect: 接收返回坐标.

// ff:取窗口客户区坐标DPI
// pRect:接收返回坐标
func (e *Element) X取窗口客户区坐标DPI(接收返回坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_取窗口客户区坐标DPI(e.Handle, 接收返回坐标)
}

// 元素_取窗口客户区坐标DPI. 基于DPI缩放后的坐标.
//
// pPt: 接收返回坐标点.

// ff:取窗口客户区坐标点DPI
// pPt:接收返回坐标点
func (e *Element) X取窗口客户区坐标点DPI(接收返回坐标点 *炫彩基类.POINT) int {
	return 炫彩基类.X元素_取窗口客户区坐标点DPI(e.Handle, 接收返回坐标点)
}

// 元素_客户区坐标到窗口客户区DPI. 基于DPI缩放后的坐标.
//
// pRect: 接收返回坐标.

// ff:客户区坐标到窗口客户区DPI
// pRect:接收返回坐标
func (e *Element) X客户区坐标到窗口客户区DPI(接收返回坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X元素_客户区坐标到窗口客户区DPI(e.Handle, 接收返回坐标)
}

// SetFocus 元素_置焦点.

// ff:置焦点
func (e *Element) X置焦点() bool {
	hParent := 0
	hEle := e.Handle
	for {
		hParent = 炫彩基类.X窗口组件_取父对象(hEle)
		if 炫彩基类.X判断窗口(hParent) {
			break
		}

		if hParent == 0 {
			return false
		}

		hEle = hParent
	}

	炫彩基类.X窗口_置焦点(hParent, e.Handle)
	return true
}

// GetLeft 元素_取左边.

// ff:取左边
func (e *Element) X取左边() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X元素_取坐标(e.Handle, &rc)
	return rc.Left
}

// GetTop 元素_取顶边.

// ff:取顶边
func (e *Element) X取顶边() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X元素_取坐标(e.Handle, &rc)
	return rc.Top
}

// GetRight 元素_取右边.

// ff:取右边
func (e *Element) X取右边() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X元素_取坐标(e.Handle, &rc)
	return rc.Right
}

// GetBottom 元素_取底边.

// ff:取底边
func (e *Element) X取底边() int32 {
	var rc 炫彩基类.RECT
	炫彩基类.X元素_取坐标(e.Handle, &rc)
	return rc.Bottom
}

// SetLeft 元素_置左边.
//
// x: 左边x坐标.
//
// bRedraw: 是否重绘.

// ff:置左边
// bRedraw:是否重绘
// x:左边x坐标
func (e *Element) X置左边(左边x坐标 int32, 是否重绘 bool) bool {
	return 炫彩基类.X元素_移动(e.Handle, 左边x坐标, e.X取顶边(), 是否重绘, 炫彩常量类.AdjustLayout_All, 0) != 0
}

// SetLeft 元素_置顶边.
//
// y: 顶边y坐标.
//
// bRedraw: 是否重绘.

// ff:置顶边
// bRedraw:是否重绘
// y:顶边y坐标
func (e *Element) X置顶边(顶边y坐标 int32, 是否重绘 bool) bool {
	return 炫彩基类.X元素_移动(e.Handle, e.X取左边(), 顶边y坐标, 是否重绘, 炫彩常量类.AdjustLayout_All, 0) != 0
}

/*
下面都是事件
*/

type XE_ELEPROCE func(nEvent int, wParam, lParam uint, pbHandled *bool) int               // 元素处理过程事件.
type XE_ELEPROCE1 func(hEle int, nEvent int, wParam, lParam uint, pbHandled *bool) int    // 元素处理过程事件.
type XE_PAINT func(hDraw int, pbHandled *bool) int                                        // 元素绘制事件.
type XE_PAINT1 func(hEle int, hDraw int, pbHandled *bool) int                             // 元素绘制事件.
type XE_PAINT_END func(hDraw int, pbHandled *bool) int                                    // 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
type XE_PAINT_END1 func(hEle int, hDraw int, pbHandled *bool) int                         // 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().
type XE_PAINT_SCROLLVIEW func(hDraw int, pbHandled *bool) int                             // 滚动视图绘制事件.
type XE_PAINT_SCROLLVIEW1 func(hEle int, hDraw int, pbHandled *bool) int                  // 滚动视图绘制事件.
type XE_MOUSEMOVE func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int                  // 元素鼠标移动事件.
type XE_MOUSEMOVE1 func(hEle int, nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int       // 元素鼠标移动事件.
type XE_MOUSESTAY func(pbHandled *bool) int                                               // 元素鼠标进入事件.
type XE_MOUSESTAY1 func(hEle int, pbHandled *bool) int                                    // 元素鼠标进入事件.
type XE_MOUSEHOVER func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int                 // 元素鼠标悬停事件.
type XE_MOUSEHOVER1 func(hEle int, nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int      // 元素鼠标悬停事件.
type XE_MOUSELEAVE func(hEleStay int, pbHandled *bool) int                                // 元素鼠标离开事件.
type XE_MOUSELEAVE1 func(hEle int, hEleStay int, pbHandled *bool) int                     // 元素鼠标离开事件.
type XE_MOUSEWHEEL func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int                 // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL(). flags: 见MSDN中WM_MOUSEWHEEL消息wParam参数说明.
type XE_MOUSEWHEEL1 func(hEle int, nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int      // 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL(). flags: 见MSDN中WM_MOUSEWHEEL消息wParam参数说明.
type XE_LBUTTONDOWN func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int                // 鼠标左键按下事件.
type XE_LBUTTONDOWN1 func(hEle int, nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int     // 鼠标左键按下事件.
type XE_LBUTTONUP func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int                  // 鼠标左键弹起事件.
type XE_LBUTTONUP1 func(hEle int, nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int       // 鼠标左键弹起事件.
type XE_RBUTTONDOWN func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int                // 鼠标右键按下事件.
type XE_RBUTTONDOWN1 func(hEle int, nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int     // 鼠标右键按下事件.
type XE_RBUTTONUP func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int                  // 鼠标右键弹起事件.
type XE_RBUTTONUP1 func(hEle int, nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int       // 鼠标右键弹起事件.
type XE_LBUTTONDBCLICK func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int             // 鼠标左键双击事件.
type XE_LBUTTONDBCLICK1 func(hEle int, nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int  // 鼠标左键双击事件.
type XE_XC_TIMER func(nTimerID int, pbHandled *bool) int                                  // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
type XE_XC_TIMER1 func(hEle int, nTimerID int, pbHandled *bool) int                       // 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.
type XE_ADJUSTLAYOUT func(nFlags int32, nAdjustNo uint32, pbHandled *bool) int            // 调整布局事件. 暂停使用.
type XE_ADJUSTLAYOUT1 func(hEle int, nFlags int32, nAdjustNo uint32, pbHandled *bool) int // 调整布局事件. 暂停使用.
type XE_TOOLTIP_POPUP func(hWindow int, pText uintptr, pbHandled *bool) int               // 元素工具提示弹出事件.
type XE_TOOLTIP_POPUP1 func(hEle int, hWindow int, pText uintptr, pbHandled *bool) int    // 元素工具提示弹出事件1.

// 调整布局完成事件.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_ADJUSTLAYOUT_END func(nFlags 炫彩常量类.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

// 调整布局完成事件.
//
// hEle: 元素句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_ADJUSTLAYOUT_END1 func(hEle int, nFlags 炫彩常量类.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

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
type XE_SIZE func(nFlags 炫彩常量类.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

// 元素大小改变事件1.
//
// hEle: 元素句柄.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
//
// nAdjustNo: 调整布局流水号.
type XE_SIZE1 func(hEle int, nFlags 炫彩常量类.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int

type XE_SHOW func(bShow bool, pbHandled *bool) int                          // 元素显示隐藏事件.
type XE_SHOW1 func(hEle int, bShow bool, pbHandled *bool) int               // 元素显示隐藏事件.
type XE_SETFONT func(pbHandled *bool) int                                   // 元素设置字体事件.
type XE_SETFONT1 func(hEle int, pbHandled *bool) int                        // 元素设置字体事件.
type XE_KEYDOWN func(wParam, lParam uint, pbHandled *bool) int              // 元素按键事件.
type XE_KEYDOWN1 func(hEle int, wParam, lParam uint, pbHandled *bool) int   // 元素按键事件.
type XE_KEYUP func(wParam, lParam uint, pbHandled *bool) int                // 元素按键事件.
type XE_KEYUP1 func(hEle int, wParam, lParam uint, pbHandled *bool) int     // 元素按键事件.
type XE_CHAR func(wParam, lParam uint, pbHandled *bool) int                 // 通过TranslateMessage函数翻译的字符事件.
type XE_CHAR1 func(hEle int, wParam, lParam uint, pbHandled *bool) int      // 通过TranslateMessage函数翻译的字符事件.
type XE_SETCAPTURE func(pbHandled *bool) int                                // 元素设置鼠标捕获.
type XE_SETCAPTURE1 func(hEle int, pbHandled *bool) int                     // 元素设置鼠标捕获.
type XE_KILLCAPTURE func(pbHandled *bool) int                               // 元素失去鼠标捕获.
type XE_KILLCAPTURE1 func(hEle int, pbHandled *bool) int                    // 元素失去鼠标捕获.
type XE_SETCURSOR func(wParam, lParam uint, pbHandled *bool) int            // 设置鼠标光标.
type XE_SETCURSOR1 func(hEle int, wParam, lParam uint, pbHandled *bool) int // 设置鼠标光标.
type XE_DROPFILES func(hDropInfo uintptr, pbHandled *bool) int              // 文件拖放事件, 需先启用:XWnd_EnableDragFiles().
type XE_DROPFILES1 func(hEle int, hDropInfo uintptr, pbHandled *bool) int   // 文件拖放事件, 需先启用:XWnd_EnableDragFiles().

// 元素处理过程事件.

// ff:事件_处理过程
// pFun:
func (e *Element) X事件_处理过程(pFun XE_ELEPROCE) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_ELEPROCE, pFun)
}

// 元素处理过程事件.

// ff:事件_处理过程1
// pFun:
func (e *Element) X事件_处理过程1(pFun XE_ELEPROCE1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_ELEPROCE, pFun)
}

// 元素绘制事件.

// ff:事件_绘制事件
// pFun:
func (e *Element) X事件_绘制事件(pFun XE_PAINT) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_PAINT, pFun)
}

// 元素绘制事件.

// ff:事件_绘制事件1
// pFun:
func (e *Element) X事件_绘制事件1(pFun XE_PAINT1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_PAINT, pFun)
}

// 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().

// ff:事件_绘制完成
// pFun:
func (e *Element) X事件_绘制完成(pFun XE_PAINT_END) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_PAINT_END, pFun)
}

// 该元素及子元素绘制完成事件.启用该功能需要调用XEle_EnableEvent_XE_PAINT_END().

// ff:事件_绘制完成1
// pFun:
func (e *Element) X事件_绘制完成1(pFun XE_PAINT_END1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_PAINT_END, pFun)
}

// 滚动视图绘制事件.

// ff:事件_滚动视图
// pFun:
func (e *Element) X事件_滚动视图(pFun XE_PAINT_SCROLLVIEW) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_PAINT_SCROLLVIEW, pFun)
}

// 滚动视图绘制事件.

// ff:事件_滚动视图1
// pFun:
func (e *Element) X事件_滚动视图1(pFun XE_PAINT_SCROLLVIEW1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_PAINT_SCROLLVIEW, pFun)
}

// 元素鼠标移动事件.

// ff:事件_鼠标移动
// pFun:
func (e *Element) X事件_鼠标移动(pFun XE_MOUSEMOVE) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MOUSEMOVE, pFun)
}

// 元素鼠标移动事件.

// ff:事件_鼠标移动1
// pFun:
func (e *Element) X事件_鼠标移动1(pFun XE_MOUSEMOVE1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MOUSEMOVE, pFun)
}

// 元素鼠标进入事件.

// ff:事件_鼠标进入
// pFun:
func (e *Element) X事件_鼠标进入(pFun XE_MOUSESTAY) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MOUSESTAY, pFun)
}

// 元素鼠标进入事件.

// ff:事件_鼠标进入1
// pFun:
func (e *Element) X事件_鼠标进入1(pFun XE_MOUSESTAY1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MOUSESTAY, pFun)
}

// 元素鼠标悬停事件.

// ff:事件_鼠标悬停
// pFun:
func (e *Element) X事件_鼠标悬停(pFun XE_MOUSEHOVER) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MOUSEHOVER, pFun)
}

// 元素鼠标悬停事件.

// ff:事件_鼠标悬停1
// pFun:
func (e *Element) X事件_鼠标悬停1(pFun XE_MOUSEHOVER1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MOUSEHOVER, pFun)
}

// 元素鼠标离开事件.

// ff:事件_鼠标离开
// pFun:
func (e *Element) X事件_鼠标离开(pFun XE_MOUSELEAVE) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MOUSELEAVE, pFun)
}

// 元素鼠标离开事件.

// ff:事件_鼠标离开1
// pFun:
func (e *Element) X事件_鼠标离开1(pFun XE_MOUSELEAVE1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MOUSELEAVE, pFun)
}

// 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL().

// ff:事件_鼠标滚轮滚动
// pFun:
func (e *Element) X事件_鼠标滚轮滚动(pFun XE_MOUSEWHEEL) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MOUSEWHEEL, pFun)
}

// 元素鼠标滚轮滚动事件. 如果非滚动视图需要调用 XEle_EnableEvent_XE_MOUSEWHEEL().

// ff:事件_鼠标滚轮滚动1
// pFun:
func (e *Element) X事件_鼠标滚轮滚动1(pFun XE_MOUSEWHEEL1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MOUSEWHEEL, pFun)
}

// 鼠标左键按下事件.

// ff:事件_左键按下
// pFun:
func (e *Element) X事件_左键按下(pFun XE_LBUTTONDOWN) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_LBUTTONDOWN, pFun)
}

// 鼠标左键按下事件.

// ff:事件_左键按下1
// pFun:
func (e *Element) X事件_左键按下1(pFun XE_LBUTTONDOWN1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_LBUTTONDOWN, pFun)
}

// 鼠标左键弹起事件.

// ff:事件_左键弹起
// pFun:
func (e *Element) X事件_左键弹起(pFun XE_LBUTTONUP) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_LBUTTONUP, pFun)
}

// 鼠标左键弹起事件.

// ff:事件_左键弹起1
// pFun:
func (e *Element) X事件_左键弹起1(pFun XE_LBUTTONUP1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_LBUTTONUP, pFun)
}

// 鼠标右键按下事件.

// ff:事件_右键按下
// pFun:
func (e *Element) X事件_右键按下(pFun XE_RBUTTONDOWN) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_RBUTTONDOWN, pFun)
}

// 鼠标右键按下事件.

// ff:事件_右键按下1
// pFun:
func (e *Element) X事件_右键按下1(pFun XE_RBUTTONDOWN1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_RBUTTONDOWN, pFun)
}

// 鼠标右键弹起事件.

// ff:事件_右键弹起
// pFun:
func (e *Element) X事件_右键弹起(pFun XE_RBUTTONUP) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_RBUTTONUP, pFun)
}

// 鼠标右键弹起事件.

// ff:事件_右键弹起1
// pFun:
func (e *Element) X事件_右键弹起1(pFun XE_RBUTTONUP1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_RBUTTONUP, pFun)
}

// 鼠标左键双击事件.

// ff:事件_左键双击
// pFun:
func (e *Element) X事件_左键双击(pFun XE_LBUTTONDBCLICK) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_LBUTTONDBCLICK, pFun)
}

// 鼠标左键双击事件.

// ff:事件_左键双击1
// pFun:
func (e *Element) X事件_左键双击1(pFun XE_LBUTTONDBCLICK1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_LBUTTONDBCLICK, pFun)
}

// 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.

// ff:事件_炫彩定时器
// pFun:
func (e *Element) X事件_炫彩定时器(pFun XE_XC_TIMER) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_XC_TIMER, pFun)
}

// 炫彩定时器,非系统定时器,定时器消息 XM_TIMER.

// ff:事件_炫彩定时器1
// pFun:
func (e *Element) X事件_炫彩定时器1(pFun XE_XC_TIMER1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_XC_TIMER, pFun)
}

// 调整布局事件. 暂停使用.

// ff:事件_调整布局
// pFun:
func (e *Element) X事件_调整布局(pFun XE_ADJUSTLAYOUT) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_ADJUSTLAYOUT, pFun)
}

// 调整布局事件. 暂停使用.

// ff:事件_调整布局1
// pFun:
func (e *Element) X事件_调整布局1(pFun XE_ADJUSTLAYOUT1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_ADJUSTLAYOUT, pFun)
}

// 调整布局完成事件.

// ff:事件_调整布局完成
// pFun:
func (e *Element) X事件_调整布局完成(pFun XE_ADJUSTLAYOUT_END) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_ADJUSTLAYOUT_END, pFun)
}

// 调整布局完成事件.

// ff:事件_调整布局完成1
// pFun:
func (e *Element) X事件_调整布局完成1(pFun XE_ADJUSTLAYOUT_END1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_ADJUSTLAYOUT_END, pFun)
}

// 元素获得焦点事件.

// ff:事件_获得焦点
// pFun:
func (e *Element) X事件_获得焦点(pFun XE_SETFOCUS) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_SETFOCUS, pFun)
}

// 元素获得焦点事件.

// ff:事件_获得焦点1
// pFun:
func (e *Element) X事件_获得焦点1(pFun XE_SETFOCUS1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_SETFOCUS, pFun)
}

// 元素失去焦点事件.

// ff:事件_失去焦点
// pFun:
func (e *Element) X事件_失去焦点(pFun XE_KILLFOCUS) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_KILLFOCUS, pFun)
}

// 元素失去焦点事件.

// ff:事件_失去焦点1
// pFun:
func (e *Element) X事件_失去焦点1(pFun XE_KILLFOCUS1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_KILLFOCUS, pFun)
}

// 元素即将销毁事件. 在销毁子对象之前触发.

// ff:事件_即将销毁
// pFun:
func (e *Element) X事件_即将销毁(pFun XE_DESTROY) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_DESTROY, pFun)
}

// 元素即将销毁事件. 在销毁子对象之前触发.

// ff:事件_即将销毁1
// pFun:
func (e *Element) X事件_即将销毁1(pFun XE_DESTROY1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_DESTROY, pFun)
}

// 元素销毁完成事件. 在销毁子对象之后触发.

// ff:事件_销毁完成
// pFun:
func (e *Element) X事件_销毁完成(pFun XE_DESTROY_END) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_DESTROY_END, pFun)
}

// 元素销毁完成事件. 在销毁子对象之后触发.

// ff:事件_销毁完成1
// pFun:
func (e *Element) X事件_销毁完成1(pFun XE_DESTROY_END1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_DESTROY_END, pFun)
}

// 元素大小改变事件.

// ff:事件_大小改变事件
// pFun:
func (e *Element) X事件_大小改变事件(pFun XE_SIZE) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_SIZE, pFun)
}

// 元素大小改变事件.

// ff:事件_大小改变事件1
// pFun:
func (e *Element) X事件_大小改变事件1(pFun XE_SIZE1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_SIZE, pFun)
}

// 元素显示隐藏事件.

// ff:事件_显示隐藏事件
// pFun:
func (e *Element) X事件_显示隐藏事件(pFun XE_SHOW) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_SHOW, pFun)
}

// 元素显示隐藏事件.

// ff:事件_显示隐藏事件1
// pFun:
func (e *Element) X事件_显示隐藏事件1(pFun XE_SHOW1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_SHOW, pFun)
}

// 元素设置字体事件.

// ff:事件_设置字体
// pFun:
func (e *Element) X事件_设置字体(pFun XE_SETFONT) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_SETFONT, pFun)
}

// 元素设置字体事件.

// ff:事件_设置字体1
// pFun:
func (e *Element) X事件_设置字体1(pFun XE_SETFONT1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_SETFONT, pFun)
}

// 元素按键事件.

// ff:事件_按键按下
// pFun:
func (e *Element) X事件_按键按下(pFun XE_KEYDOWN) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_KEYDOWN, pFun)
}

// 元素按键事件.

// ff:事件_按键按下1
// pFun:
func (e *Element) X事件_按键按下1(pFun XE_KEYDOWN1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_KEYDOWN, pFun)
}

// 元素按键事件.

// ff:事件_按键弹起
// pFun:
func (e *Element) X事件_按键弹起(pFun XE_KEYUP) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_KEYUP, pFun)
}

// 元素按键事件.

// ff:事件_按键弹起1
// pFun:
func (e *Element) X事件_按键弹起1(pFun XE_KEYUP1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_KEYUP, pFun)
}

// 通过TranslateMessage函数翻译的字符事件.

// ff:事件_CHAR
// pFun:
func (e *Element) X事件_CHAR(pFun XE_CHAR) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_CHAR, pFun)
}

// 通过TranslateMessage函数翻译的字符事件.

// ff:事件_CHAR1
// pFun:
func (e *Element) X事件_CHAR1(pFun XE_CHAR1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_CHAR, pFun)
}

// 元素设置鼠标捕获.

// ff:事件_设置鼠标捕获
// pFun:
func (e *Element) X事件_设置鼠标捕获(pFun XE_SETCAPTURE) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_SETCAPTURE, pFun)
}

// 元素设置鼠标捕获.

// ff:事件_设置鼠标捕获1
// pFun:
func (e *Element) X事件_设置鼠标捕获1(pFun XE_SETCAPTURE1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_SETCAPTURE, pFun)
}

// 元素失去鼠标捕获.

// ff:事件_失去鼠标捕获
// pFun:
func (e *Element) X事件_失去鼠标捕获(pFun XE_KILLCAPTURE) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_KILLCAPTURE, pFun)
}

// 元素失去鼠标捕获.

// ff:事件_失去鼠标捕获1
// pFun:
func (e *Element) X事件_失去鼠标捕获1(pFun XE_KILLCAPTURE1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_KILLCAPTURE, pFun)
}

// 设置鼠标光标.

// ff:事件_设置鼠标光标
// pFun:
func (e *Element) X事件_设置鼠标光标(pFun XE_SETCURSOR) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_SETCURSOR, pFun)
}

// 设置鼠标光标.

// ff:事件_设置鼠标光标1
// pFun:
func (e *Element) X事件_设置鼠标光标1(pFun XE_SETCURSOR1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_SETCURSOR, pFun)
}

// 文件拖放事件, 需先启用:XWnd_EnableDragFiles().

// ff:事件_文件拖放
// pFun:
func (e *Element) X事件_文件拖放(pFun XE_DROPFILES) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_DROPFILES, pFun)
}

// 文件拖放事件, 需先启用:XWnd_EnableDragFiles().

// ff:事件_文件拖放事件1
// pFun:
func (e *Element) X事件_文件拖放事件1(pFun XE_DROPFILES1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_DROPFILES, pFun)
}

// 事件_元素工具提示弹出, 可使用 common.UintPtrToString 把uintptr转换到文本.

// ff:事件_工具提示弹出
// pFun:
func (e *Element) X事件_工具提示弹出(pFun XE_TOOLTIP_POPUP) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_TOOLTIP_POPUP, pFun)
}

// 事件_元素工具提示弹出1, 可使用 common.UintPtrToString 把uintptr转换到文本.

// ff:事件_工具提示弹出1
// pFun:
func (e *Element) X事件_工具提示弹出1(pFun XE_TOOLTIP_POPUP1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_TOOLTIP_POPUP, pFun)
}

type XE_MENU_SELECT func(nID int32, pbHandled *bool) int                                            // 弹出菜单项选择事件.
type XE_MENU_POPUP func(HMENUX int, pbHandled *bool) int                                            // 菜单弹出.
type XE_MENU_EXIT func(pbHandled *bool) int                                                         // 弹出菜单退出事件.
type XE_MENU_POPUP_WND func(hMenu int, pInfo *炫彩基类.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XE_MENU_DRAW_BACKGROUND func(hDraw int, pInfo *炫彩基类.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XE_MENU_DRAWITEM func(hDraw int, pInfo *炫彩基类.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// 事件_弹出菜单项被选择.

// ff:事件_弹出菜单项被选择
// pFun:
func (e *Element) X事件_弹出菜单项被选择(pFun XE_MENU_SELECT) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MENU_SELECT, pFun)
}

// 事件_菜单弹出.

// ff:事件_菜单弹出
// pFun:
func (e *Element) X事件_菜单弹出(pFun XE_MENU_POPUP) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MENU_POPUP, pFun)
}

// 事件_菜单退出.

// ff:事件_菜单退出
// pFun:
func (e *Element) X事件_菜单退出(pFun XE_MENU_EXIT) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MENU_EXIT, pFun)
}

// 菜单弹出窗口.

// ff:事件_菜单弹出窗口
// pFun:
func (e *Element) X事件_菜单弹出窗口(pFun XE_MENU_POPUP_WND) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MENU_POPUP_WND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().

// ff:事件_绘制菜单背景
// pFun:
func (e *Element) X事件_绘制菜单背景(pFun XE_MENU_DRAW_BACKGROUND) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// ff:事件_绘制菜单项事件
// pFun:
func (e *Element) X事件_绘制菜单项事件(pFun XE_MENU_DRAWITEM) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_MENU_DRAWITEM, pFun)
}

type XE_MENU_SELECT1 func(hEle int, nID int, pbHandled *bool) int                                              // 弹出菜单项选择事件.
type XE_MENU_POPUP1 func(hEle int, HMENUX int, pbHandled *bool) int                                            // 菜单弹出.
type XE_MENU_EXIT1 func(hEle int, pbHandled *bool) int                                                         // 弹出菜单退出事件.
type XE_MENU_POPUP_WND1 func(hEle int, hMenu int, pInfo *炫彩基类.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XE_MENU_DRAW_BACKGROUND1 func(hEle int, hDraw int, pInfo *炫彩基类.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XE_MENU_DRAWITEM1 func(hEle int, hDraw int, pInfo *炫彩基类.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// 事件_弹出菜单项被选择.

// ff:事件_弹出菜单项被选择1
// pFun:
func (e *Element) X事件_弹出菜单项被选择1(pFun XE_MENU_SELECT1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MENU_SELECT, pFun)
}

// 事件_菜单弹出.

// ff:事件_菜单弹出1
// pFun:
func (e *Element) X事件_菜单弹出1(pFun XE_MENU_POPUP1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MENU_POPUP, pFun)
}

// 事件_菜单退出.

// ff:事件_菜单退出1
// pFun:
func (e *Element) X事件_菜单退出1(pFun XE_MENU_EXIT1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MENU_EXIT, pFun)
}

// 菜单弹出窗口.

// ff:菜单弹出窗口1
// pFun:
func (e *Element) X菜单弹出窗口1(pFun XE_MENU_POPUP_WND1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MENU_POPUP_WND, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().

// ff:事件_绘制菜单背景1
// pFun:
func (e *Element) X事件_绘制菜单背景1(pFun XE_MENU_DRAW_BACKGROUND1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// ff:事件_绘制菜单项1
// pFun:
func (e *Element) X事件_绘制菜单项1(pFun XE_MENU_DRAWITEM1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_MENU_DRAWITEM, pFun)
}
