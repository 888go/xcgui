package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// TabBar Tab条.
type TabBar struct {
	Element
}

// TAB条_创建, 创建tabBar元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.

// ff:创建TAB条
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X创建TAB条(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *TabBar {
	p := &TabBar{}
	p.X设置句柄(炫彩基类.TAB条_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.

// ff:创建TAB条并按句柄
// handle:
func X创建TAB条并按句柄(handle int) *TabBar {
	p := &TabBar{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建TAB条并按名称
// name:
func X创建TAB条并按名称(name string) *TabBar {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &TabBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建TAB条并按UID
// nUID:
func X创建TAB条并按UID(nUID int) *TabBar {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &TabBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建TAB条并按UID名称
// name:
func X创建TAB条并按UID名称(name string) *TabBar {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &TabBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// TAB条_添加标签, 添加一个标签, 返回标签索引.
//
// pName: 标签文本内容.

// ff:添加标签
// pName:标签文本
func (t *TabBar) X添加标签(标签文本 string) int {
	return 炫彩基类.TAB条_添加标签(t.Handle, 标签文本)
}

// TAB条插入_标签, 插入一个标签, 返回标签索引.
//
// index: 插入位置.
//
// pName: 标签文本内容.

// ff:TAB条插入_标签
// pName:标签文本
// index:插入位置
func (t *TabBar) TAB条插入_标签(插入位置 int, 标签文本 string) int {
	return 炫彩基类.TAB条插入_标签(t.Handle, 插入位置, 标签文本)
}

// TAB条_移动标签.
//
// iSrc: 源位置索引.
//
// iDest: 目标位置索引.

// ff:移动标签
// iDest:目标位置索引
// iSrc:源位置索引
func (t *TabBar) X移动标签(源位置索引 int, 目标位置索引 int) bool {
	return 炫彩基类.TAB条_移动标签(t.Handle, 源位置索引, 目标位置索引)
}

// TAB条_删除标签, 删除一个标签.
//
// index: 位置索引.

// ff:删除标签
// index:位置索引
func (t *TabBar) X删除标签(位置索引 int) bool {
	return 炫彩基类.TAB条_删除标签(t.Handle, 位置索引)
}

// TAB条_删除全部, 删除所有标签.

// ff:删除全部
func (t *TabBar) X删除全部() int {
	return 炫彩基类.TAB条_删除全部(t.Handle)
}

// TAB条_取标签, 获取标签按钮句柄.
//
// index: 位置索引.

// ff:取标签
// index:位置索引
func (t *TabBar) X取标签(位置索引 int) int {
	return 炫彩基类.TAB条_取标签(t.Handle, 位置索引)
}

// TAB条_取标签上的关闭按钮, 获取标签上关闭按钮句柄.
//
// index: 位置索引.

// ff:取标签上的关闭按钮
// index:位置索引
func (t *TabBar) X取标签上的关闭按钮(位置索引 int) int {
	return 炫彩基类.TAB条_取标签上的关闭按钮(t.Handle, 位置索引)
}

// TAB条_取左滚动按钮, 获取左滚动按钮句柄.

// ff:取左滚动按钮
func (t *TabBar) X取左滚动按钮() int {
	return 炫彩基类.TAB条_取左滚动按钮(t.Handle)
}

// TAB条_取右滚动按钮, 获取右滚动按钮句柄.

// ff:取右滚动按钮
func (t *TabBar) X取右滚动按钮() int {
	return 炫彩基类.TAB条_取右滚动按钮(t.Handle)
}

// TAB条_取下拉菜单按钮句柄.

// ff:取下拉菜单按钮句柄
func (t *TabBar) X取下拉菜单按钮句柄() int {
	return 炫彩基类.TAB条_取下拉菜单按钮句柄(t.Handle)
}

// TAB条_取当前选择, 获取选择的标签索引.

// ff:取当前选择
func (t *TabBar) X取当前选择() int {
	return 炫彩基类.TAB条_取当前选择(t.Handle)
}

// TAB条_取间隔, 获取标签间距, 0没有间距.

// ff:取间隔
func (t *TabBar) X取间隔() int {
	return 炫彩基类.TAB条_取间隔(t.Handle)
}

// TAB条_取标签数量, 获取标签项数量.

// ff:取标签数量
func (t *TabBar) X取标签数量() int {
	return 炫彩基类.TAB条_取标签数量(t.Handle)
}

// TAB条_取标签位置索引, 获取标签按钮位置索引, 成功返回索引值, 否则返回 XC_ID_ERROR.
//
// hLabel: 标签按钮句柄.

// ff:取标签位置索引
// hLabel:标签按钮句柄
func (t *TabBar) X取标签位置索引(标签按钮句柄 int) int {
	return 炫彩基类.TAB条_取标签位置索引(t.Handle, 标签按钮句柄)
}

// TAB条_置间隔, 设置标签间距, 0没有间距.
//
// spacing: 标签间隔大小.

// ff:置间隔
// spacing:标签间隔大小
func (t *TabBar) X置间隔(标签间隔大小 int) int {
	return 炫彩基类.TAB条_置间隔(t.Handle, 标签间隔大小)
}

// TAB条_置边距, 设置内容与边框的间隔大小.
//
// left: 左边间隔大小.
//
// top: 上边间隔大小.
//
// right: 右边间隔大小.
//
// bottom: 下边间隔大小.

// ff:置边距
// bottom:下
// right:右
// top:上
// left:左
func (t *TabBar) X置边距(左 int, 上 int, 右 int, 下 int) int {
	return 炫彩基类.TAB条_置边距(t.Handle, 左, 上, 右, 下)
}

// TAB条_置选择, 设置选择标签.
//
// index: 标签位置索引.

// ff:置选择
// index:索引
func (t *TabBar) X置选择(索引 int) int {
	return 炫彩基类.TAB条_置选择(t.Handle, 索引)
}

// TAB条_左滚动, 左按钮滚动.

// ff:左滚动
func (t *TabBar) X左滚动() int {
	return 炫彩基类.TAB条_左滚动(t.Handle)
}

// TAB条_右滚动, 右按钮滚动.

// ff:右滚动
func (t *TabBar) X右滚动() int {
	return 炫彩基类.TAB条_右滚动(t.Handle)
}

// TAB条_启用平铺, 平铺标签, 每个标签显示相同大小.
//
// bTile: 是否启用.

// ff:启用平铺
// bTile:是否启用
func (t *TabBar) X启用平铺(是否启用 bool) int {
	return 炫彩基类.TAB条_启用平铺(t.Handle, 是否启用)
}

// TAB条_启用下拉菜单按钮.
//
// bEnable:.

// ff:启用下拉菜单按钮
// bEnable:是否启用
func (t *TabBar) X启用下拉菜单按钮(是否启用 bool) int {
	return 炫彩基类.TAB条_启用下拉菜单按钮(t.Handle, 是否启用)
}

// TAB条_启用标签带关闭按钮, 启用关闭标签功能.
//
// bEnable: 是否启用.

// ff:启用标签带关闭按钮
// bEnable:是否启用
func (t *TabBar) X启用标签带关闭按钮(是否启用 bool) int {
	return 炫彩基类.TAB条_启用标签带关闭按钮(t.Handle, 是否启用)
}

// TAB条_置关闭按钮大小, 设置关闭按钮大小.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.

// ff:置关闭按钮大小
// pSize:大小值
func (t *TabBar) X置关闭按钮大小(大小值 *炫彩基类.SIZE) int {
	return 炫彩基类.TAB条_置关闭按钮大小(t.Handle, 大小值)
}

// TAB条_置滚动按钮大小.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.

// ff:置滚动按钮大小
// pSize:大小值
func (t *TabBar) X置滚动按钮大小(大小值 *炫彩基类.SIZE) int {
	return 炫彩基类.TAB条_置滚动按钮大小(t.Handle, 大小值)
}

// TAB条_置指定标签固定宽度.
//
// index: 索引.
//
// nWidth: 宽度, , 如果值为-1, 那么自动计算宽度.

// ff:置指定标签固定宽度
// nWidth:宽度
// index:索引
func (t *TabBar) X置指定标签固定宽度(索引 int, 宽度 int) int {
	return 炫彩基类.TAB条_置指定标签固定宽度(t.Handle, 索引, 宽度)
}

// TAB条_显示标签, 显示或隐藏指定标签.
//
// index: 标签索引.
//
// bShow: 是否显示.

// ff:显示标签
// bShow:是否显示
// index:标签索引
func (t *TabBar) X显示标签(标签索引 int, 是否显示 bool) bool {
	return 炫彩基类.TAB条_显示标签(t.Handle, 标签索引, 是否显示)
}

/*
以下都是事件
*/

type XE_TABBAR_SELECT func(iItem int32, pbHandled *bool) int            // TabBar标签按钮选择改变事件. iItem: 标签位置索引.
type XE_TABBAR_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int // TabBar标签按钮选择改变事件. iItem: 标签位置索引.
type XE_TABBAR_DELETE func(iItem int32, pbHandled *bool) int            // TabBar标签按钮删除事件. iItem: 标签位置索引.
type XE_TABBAR_DELETE1 func(hEle int, iItem int32, pbHandled *bool) int // TabBar标签按钮删除事件. iItem: 标签位置索引.

// TabBar标签按钮选择改变事件.

// ff:事件_选择改变
// pFun:
func (t *TabBar) X事件_选择改变(pFun XE_TABBAR_SELECT) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TABBAR_SELECT, pFun)
}

// TabBar标签按钮选择改变事件.

// ff:事件_选择改变1
// pFun:
func (t *TabBar) X事件_选择改变1(pFun XE_TABBAR_SELECT1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TABBAR_SELECT, pFun)
}

// TabBar标签按钮删除事件.

// ff:事件_删除
// pFun:
func (t *TabBar) X事件_删除(pFun XE_TABBAR_DELETE) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TABBAR_DELETE, pFun)
}

// TabBar标签按钮删除事件.

// ff:事件_删除1
// pFun:
func (t *TabBar) X事件_删除1(pFun XE_TABBAR_DELETE1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TABBAR_DELETE, pFun)
}
