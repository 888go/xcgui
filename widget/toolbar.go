package 炫彩组件类

import "github.com/888go/xcgui/xc"

// 工具条.
type ToolBar struct {
	Element
}

// 工具条_创建, 创建工具条元素; 如果指定了父为窗口, 默认调用XWnd_AddToolBar()函数, 将工具条添加到窗口非客户区.
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

// ff:创建工具条
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:y坐标
// x:x坐标
func X创建工具条(x坐标, y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄 int) *ToolBar {
	p := &ToolBar{}
	p.X设置句柄(炫彩基类.X工具条_创建(x坐标, y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.

// ff:创建工具条并按句柄
// handle:
func X创建工具条并按句柄(handle int) *ToolBar {
	p := &ToolBar{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建工具条并按名称
// name:
func X创建工具条并按名称(name string) *ToolBar {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ToolBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建工具条并按UID
// nUID:
func X创建工具条并按UID(nUID int) *ToolBar {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ToolBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建工具条并按UID名称
// name:
func X创建工具条并按UID名称(name string) *ToolBar {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ToolBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 工具条_插入元素, 插入元素到工具条, 返回插入位置索引.
//
// hNewEle: 将要插入的元素.
//
// index: 插入位置索引, (-1)插入末尾.

// ff:插入元素
// index:插入位置索引
// hNewEle:将要插入的元素
func (t *ToolBar) X插入元素(将要插入的元素 int, 插入位置索引 int) int {
	return 炫彩基类.X工具条_插入元素(t.Handle, 将要插入的元素, 插入位置索引)
}

// 工具条_插入分割栏, 插入分隔符到工具条, 返回插入位置索引.
//
// index: 插入位置索引, (-1)插入末尾.
//
// color: ABGR 颜色.

// ff:插入分割栏
// color:ABGR颜色
// index:插入位置索引
func (t *ToolBar) X插入分割栏(插入位置索引 int, ABGR颜色 int) int {
	return 炫彩基类.X工具条_插入分割栏(t.Handle, 插入位置索引, ABGR颜色)
}

// 工具条_启用下拉菜单, 启用下拉菜单, 显示隐藏的项.
//
// bEnable: 是否启用.

// ff:启用下拉菜单
// bEnable:是否启用
func (t *ToolBar) X启用下拉菜单(是否启用 bool) int {
	return 炫彩基类.X工具条_启用下拉菜单(t.Handle, 是否启用)
}

// 工具条_取元素, 获取工具条上指定元素, 返回元素句柄.
//
// index: 索引值.

// ff:取元素
// index:索引值
func (t *ToolBar) X取元素(索引值 int) int {
	return 炫彩基类.X工具条_取元素(t.Handle, 索引值)
}

// 工具条_取左滚动按钮, 获取左滚动按钮句柄.

// ff:取左滚动按钮
func (t *ToolBar) X取左滚动按钮() int {
	return 炫彩基类.X工具条_取左滚动按钮(t.Handle)
}

// 工具条_取右滚动按钮, 获取右滚动按钮句柄.

// ff:取右滚动按钮
func (t *ToolBar) X取右滚动按钮() int {
	return 炫彩基类.X工具条_取右滚动按钮(t.Handle)
}

// 工具条_取菜单按钮, 获取菜单按钮句柄.

// ff:取菜单按钮
func (t *ToolBar) X取菜单按钮() int {
	return 炫彩基类.X工具条_取菜单按钮(t.Handle)
}

// 工具条_置间距, 设置对象之间的间距.
//
// nSize: 间距大小.

// ff:置间距
// nSize:间距大小
func (t *ToolBar) X置间距(间距大小 int) int {
	return 炫彩基类.X工具条_置间距(t.Handle, 间距大小)
}

// 工具条_删除元素, 删除元素, 并且销毁.
//
// index: 索引值.

// ff:删除元素
// index:索引值
func (t *ToolBar) X删除元素(索引值 int) int {
	return 炫彩基类.X工具条_删除元素(t.Handle, 索引值)
}

// 工具条_删除全部, 删除所有元素, 并且销毁.

// ff:删除全部
func (t *ToolBar) X删除全部() int {
	return 炫彩基类.X工具条_删除全部(t.Handle)
}
