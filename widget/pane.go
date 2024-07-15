package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Pane元素.
type Pane struct {
	Element
}

// 窗格_创建, 创建窗格元素, 返回元素句柄.
//
// pName: 窗格标题.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// hFrameWnd: 框架窗口.

// ff:创建窗格
// hFrameWnd:框架窗口
// nHeight:高度
// nWidth:宽度
// pName:窗格标题
func X创建窗格(窗格标题 string, 宽度 int, 高度 int, 框架窗口 int) *Pane {
	p := &Pane{}
	p.X设置句柄(炫彩基类.X窗格_创建(窗格标题, 宽度, 高度, 框架窗口))
	return p
}

// 从句柄创建对象.

// ff:创建窗格并按句柄
// handle:
func X创建窗格并按句柄(handle int) *Pane {
	p := &Pane{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建窗格并按名称
// name:
func X创建窗格并按名称(name string) *Pane {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &Pane{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建窗格并按UID
// nUID:
func X创建窗格并按UID(nUID int) *Pane {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Pane{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建窗格并按UID名称
// name:
func X创建窗格并按UID名称(name string) *Pane {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &Pane{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 窗格_置视图, 设置窗格视图元素.
//
// hView: 绑定视图元素.

// ff:置视图
// hView:绑定视图元素
func (p *Pane) X置视图(绑定视图元素 int) int {
	return 炫彩基类.X窗格_置视图(p.Handle, 绑定视图元素)
}

// 窗格_置标题, 设置标题文本.
//
// pTitle: 文本内容.

// ff:置标题
// pTitle:文本内容
func (p *Pane) X置标题(文本内容 string) int {
	return 炫彩基类.X窗格_置标题(p.Handle, 文本内容)
}

// 窗格_取标题, 获取标题文本.

// ff:取标题
func (p *Pane) X取标题() string {
	return 炫彩基类.X窗格_取标题(p.Handle)
}

// 窗格_置标题栏高度, 设置标题栏高度.
//
// nHeight: 高度.

// ff:置标题栏高度
// nHeight:高度
func (p *Pane) X置标题栏高度(高度 int) int {
	return 炫彩基类.X窗格_置标题栏高度(p.Handle, 高度)
}

// 窗格_取标题栏高度, 获取标题栏高度.

// ff:取标题栏高度
func (p *Pane) X取标题栏高度() int {
	return 炫彩基类.X窗格_取标题栏高度(p.Handle)
}

// 窗格_判断显示, 判断窗格是否显示.

// ff:判断显示
func (p *Pane) X判断显示() bool {
	return 炫彩基类.X窗格_判断显示(p.Handle)
}

// 窗格_置大小, 设置窗格大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.

// ff:置大小
// nHeight:高度
// nWidth:宽度
func (p *Pane) X置大小(宽度 int, 高度 int) int {
	return 炫彩基类.X窗格_置大小(p.Handle, 宽度, 高度)
}

// 窗格_取状态, 获取窗格停靠状态, 返回: Pane_State_.

// ff:取状态
func (p *Pane) X取状态() 炫彩常量类.Pane_State_ {
	return 炫彩基类.X窗格_取状态(p.Handle)
}

// 窗格_取视图坐标, 获取窗格视图坐标.
//
// pRect: 接收返回的坐标值.

// ff:取视图坐标
// pRect:接收返回坐标值
func (p *Pane) X取视图坐标(接收返回坐标值 *炫彩基类.RECT) int {
	return 炫彩基类.X窗格_取视图坐标(p.Handle, 接收返回坐标值)
}

// HidePane 窗格_隐藏.
//
//	@param bGroupActivate 当为窗格组成员时, 延迟处理窗格组成员激活的切换.
//	@return int

// ff:隐藏
// bGroupActivate:延迟组成员激活
func (p *Pane) X隐藏(延迟组成员激活 bool) int {
	return 炫彩基类.X窗格_隐藏(p.Handle, 延迟组成员激活)
}

// ShowPane 窗格_显示.
//
//	@param bGroupActivate 如果是窗格组成员, 那么窗格组切换当前窗格为显示状态.
//	@return int

// ff:显示
// bGroupActivate:延迟组成员激活
func (p *Pane) X显示(延迟组成员激活 bool) int {
	return 炫彩基类.X窗格_显示(p.Handle, 延迟组成员激活)
}

// 窗格_停靠, 窗格停靠到码头.

// ff:停靠
func (p *Pane) X停靠() int {
	return 炫彩基类.X窗格_停靠(p.Handle)
}

// 窗格_锁定, 锁定窗格.

// ff:锁定
func (p *Pane) X锁定() int {
	return 炫彩基类.X窗格_锁定(p.Handle)
}

// 窗格_浮动.

// ff:浮动
func (p *Pane) X浮动() int {
	return 炫彩基类.X窗格_浮动(p.Handle)
}

// 窗格_绘制, 手动调用该函数绘制窗格, 以便控制绘制顺序.
//
// hDraw: 图形绘制句柄.

// ff:绘制
// hDraw:图形绘制句柄
func (p *Pane) X绘制(图形绘制句柄 int) int {
	return 炫彩基类.X窗格_绘制(p.Handle, 图形绘制句柄)
}

// 窗口_置选中, 如果窗格是组成员, 设置选中当前窗格可见.

// ff:置选中
func (p *Pane) X置选中() bool {
	return 炫彩基类.X窗口_置选中(p.Handle)
}

// 窗格_是否激活. 判断窗格是否激活, 当为组成员时有效.

// ff:是否激活
func (p *Pane) X是否激活() bool {
	return 炫彩基类.X窗格_是否激活(p.Handle)
}
