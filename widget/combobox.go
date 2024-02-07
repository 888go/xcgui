package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 下拉组合框.
type ComboBox struct {
	Edit
}

// 组合框_创建.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口.
func X创建组合框(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *ComboBox {
	p := &ComboBox{}
	p.X设置句柄(炫彩基类.X组合框_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建组合框并按句柄(句柄 int) *ComboBox {
	p := &ComboBox{}
	p.X设置句柄(句柄)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建组合框并按名称(名称 string) *ComboBox {
	handle := 炫彩基类.X取对象从名称(名称)
	if handle > 0 {
		p := &ComboBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建组合框并按UID(nUID int) *ComboBox {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ComboBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建组合框并按UID名称(UID名称 string) *ComboBox {
	handle := 炫彩基类.X取对象从UID名称(UID名称)
	if handle > 0 {
		p := &ComboBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 组合框_置选择项.
//
// iIndex: 项索引.
func (c *ComboBox) X置选择项(项索引 int) bool {
	return 炫彩基类.X组合框_置选择项(c.Handle, 项索引)
}

// 组合框_创建数据适配器, 返回数据适配器句柄.
func (c *ComboBox) X创建数据适配器() int {
	return 炫彩基类.X组合框_创建数据适配器(c.Handle)
}

// 组合框_绑定数据适配器.
//
// hAdapter: 适配器句柄.
func (c *ComboBox) X绑定数据适配器(适配器句柄 int) int {
	return 炫彩基类.X组合框_绑定数据适配器(c.Handle, 适配器句柄)
}

// 组合框_取数据适配器, 获取绑定的数据适配器.
func (c *ComboBox) X取数据适配器() int {
	return 炫彩基类.X组合框_取数据适配器(c.Handle)
}

// 组合框_置绑定名称.
//
// pName: 字段名.
func (c *ComboBox) X置绑定名称(字段名 string) int {
	return 炫彩基类.X组合框_置绑定名称(c.Handle, 字段名)
}

// 组合框_取下拉按钮坐标.
//
// pRect: 坐标.
func (c *ComboBox) X取下拉按钮坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X组合框_取下拉按钮坐标(c.Handle, 坐标)
}

// 组合框_置下拉按钮大小.
//
// size: 大小.
func (c *ComboBox) X置下拉按钮大小(大小 int) int {
	return 炫彩基类.X组合框_置下拉按钮大小(c.Handle, 大小)
}

// 组合框_置下拉列表高度.
//
// height: 高度, -1自动计算高度.
func (c *ComboBox) X置下拉列表高度(高度 int) int {
	return 炫彩基类.X组合框_置下拉列表高度(c.Handle, 高度)
}

// 组合框_取下拉列表高度.
func (c *ComboBox) X取下拉列表高度() int {
	return 炫彩基类.X组合框_取下拉列表高度(c.Handle)
}

// 组合框_置项模板, 设置下拉列表项模板文件.
//
// pXmlFile: 项模板文件.
func (c *ComboBox) X置项模板(项模板文件 string) int {
	return 炫彩基类.X组合框_置项模板(c.Handle, 项模板文件)
}

// 组合框_置项模板从字符串, 设置下拉列表项模板.
//
// pStringXML: 字符串.
func (c *ComboBox) X置项模板并按字符串(字符串 string) int {
	return 炫彩基类.X组合框_置项模板从字符串(c.Handle, 字符串)
}

// 组合框_启用绘制下拉按钮, 是否绘制下拉按钮.
//
// bEnable: 是否绘制.
func (c *ComboBox) X启用绘制下拉按钮(是否绘制 bool) int {
	return 炫彩基类.X组合框_启用绘制下拉按钮(c.Handle, 是否绘制)
}

// 组合框_启用编辑, 启用可编辑显示的文本内容.
//
// bEdit: TRUE可编辑.
func (c *ComboBox) X启用编辑(TRUE可编辑 bool) int {
	return 炫彩基类.X组合框_启用编辑(c.Handle, TRUE可编辑)
}

// 组合框_启用下拉列表高度固定大小, 启用/关闭下拉列表高度固定大小.
//
// bEnable: 是否启用.
func (c *ComboBox) X启用下拉列表高度固定大小(是否启用 bool) int {
	return 炫彩基类.X组合框_启用下拉列表高度固定大小(c.Handle, 是否启用)
}

// 组合框_取选择项, 获取组合框下拉列表中选择项索引.
func (c *ComboBox) X取选择项() int {
	return 炫彩基类.X组合框_取选择项(c.Handle)
}

// 组合框_取状态, 返回: ComboBox_State_.
func (c *ComboBox) X取状态() 炫彩常量类.ComboBox_State_ {
	return 炫彩基类.X组合框_取状态(c.Handle)
}

// 组合框_添加项文本, 返回项索引.
//
// pText:.
func (c *ComboBox) X添加项文本(文本 string) int {
	return 炫彩基类.X组合框_添加项文本(c.Handle, 文本)
}

// 组合框_添加项文本扩展, 返回项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) X添加项文本EX(字段名 string, 文本 string) int {
	return 炫彩基类.X组合框_添加项文本EX(c.Handle, 字段名, 文本)
}

// 组合框_添加项图片, 返回项索引.
//
// hImage: 图片句柄.
func (c *ComboBox) X添加项图片(图片句柄 int) int {
	return 炫彩基类.X组合框_添加项图片(c.Handle, 图片句柄)
}

// 组合框_添加项图片扩展, 返回项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) X添加项图片EX(字段名 string, 图片句柄 int) int {
	return 炫彩基类.X组合框_添加项图片EX(c.Handle, 字段名, 图片句柄)
}

// 组合框_插入项文本, 返回项索引.
//
// iItem: 项索引.
//
// pText: 文本.
func (c *ComboBox) X插入项文本(项索引 int, 文本 string) int {
	return 炫彩基类.X组合框_插入项文本(c.Handle, 项索引, 文本)
}

// 组合框_插入项文本扩展, 返回项索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) X插入项文本EX(项索引 int, 字段名 string, 文本 string) int {
	return 炫彩基类.X组合框_插入项文本EX(c.Handle, 项索引, 字段名, 文本)
}

// 组合框_插入项图片, 返回项索引.
//
// iItem: 项索引.
//
// hImage: 图片句柄.
func (c *ComboBox) X插入项图片(项索引 int, 图片句柄 int) int {
	return 炫彩基类.X组合框_插入项图片(c.Handle, 项索引, 图片句柄)
}

// 组合框_插入项图片扩展, 返回项索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) X插入项图片EX(项索引 int, 字段名 string, 图片句柄 int) int {
	return 炫彩基类.X组合框_插入项图片EX(c.Handle, 项索引, 字段名, 图片句柄)
}

// 组合框_置项文本.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pText: 文本.
func (c *ComboBox) X置项文本(项索引 int, 列索引 int, 文本 string) bool {
	return 炫彩基类.X组合框_置项文本(c.Handle, 项索引, 列索引, 文本)
}

// 组合框_置项文本扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) X置项文本EX(项索引 int, 字段名 string, 文本 string) bool {
	return 炫彩基类.X组合框_置项文本EX(c.Handle, 项索引, 字段名, 文本)
}

// 组合框_置项图片.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (c *ComboBox) X置项图片(项索引 int, 列索引 int, 图片句柄 int) bool {
	return 炫彩基类.X组合框_置项图片(c.Handle, 项索引, 列索引, 图片句柄)
}

// 组合框_置项图片扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) X置项图片EX(项索引 int, 字段名 string, 图片句柄 int) bool {
	return 炫彩基类.X组合框_置项图片EX(c.Handle, 项索引, 字段名, 图片句柄)
}

// 组合框_置项整数值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 整数值.
func (c *ComboBox) X置项整数值(项索引 int, 列索引 int, 整数值 int32) bool {
	return 炫彩基类.X组合框_置项整数值(c.Handle, 项索引, 列索引, 整数值)
}

// 组合框_置项指数值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// nValue: 整数值.
func (c *ComboBox) X置项指数值EX(项索引 int, 字段名 string, 整数值 int32) bool {
	return 炫彩基类.X组合框_置项指数值EX(c.Handle, 项索引, 字段名, 整数值)
}

// 组合框_置项浮点值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fFloat: 浮点数.
func (c *ComboBox) X置项浮点值(项索引 int, 列索引 int, 浮点数 float32) bool {
	return 炫彩基类.X组合框_置项浮点值(c.Handle, 项索引, 列索引, 浮点数)
}

// 组合框_置项浮点值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// fFloat: 浮点数.
func (c *ComboBox) X置项浮点值EX(项索引 int, 字段名 string, 浮点数 float32) bool {
	return 炫彩基类.X组合框_置项浮点值EX(c.Handle, 项索引, 字段名, 浮点数)
}

// 组合框_取项文本.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (c *ComboBox) X取项文本(项索引 int32, 列索引 int32) string {
	return 炫彩基类.X组合框_取项文本(c.Handle, 项索引, 列索引)
}

// 组合框_取项文本扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
func (c *ComboBox) X取项文本EX(项索引 int, 字段名 string) string {
	return 炫彩基类.X组合框_取项文本EX(c.Handle, 项索引, 字段名)
}

// 组合框_取项图片.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (c *ComboBox) X取项图片(项索引 int, 列索引 int) int {
	return 炫彩基类.X组合框_取项图片(c.Handle, 项索引, 列索引)
}

// 组合框_取项图片扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
func (c *ComboBox) X取项图片EX(项索引 int, 字段名 string) int {
	return 炫彩基类.X组合框_取项图片EX(c.Handle, 项索引, 字段名)
}

// 组合框_取项整数值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回整数值.
func (c *ComboBox) X取项整数值(项索引 int, 列索引 int, 接收返回整数值 *int32) bool {
	return 炫彩基类.X组合框_取项整数值(c.Handle, 项索引, 列索引, 接收返回整数值)
}

// 组合框_取项整数值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回整数值.
func (c *ComboBox) X取项整数值EX(项索引 int, 字段名 string, 接收返回整数值 *int32) bool {
	return 炫彩基类.X组合框_取项整数值EX(c.Handle, 项索引, 字段名, 接收返回整数值)
}

// 组合框_取项浮点值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回浮点值.
func (c *ComboBox) X取项浮点值(项索引 int, 列索引 int, 接收返回浮点值 *float32) bool {
	return 炫彩基类.X组合框_取项浮点值(c.Handle, 项索引, 列索引, 接收返回浮点值)
}

// 组合框_取项浮点值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回浮点值.
func (c *ComboBox) X取项浮点值EX(项索引 int, 字段名 string, 接收返回浮点值 *float32) bool {
	return 炫彩基类.X组合框_取项浮点值EX(c.Handle, 项索引, 字段名, 接收返回浮点值)
}

// 组合框_删除项.
//
// iItem: 项索引.
func (c *ComboBox) X删除项(项索引 int) bool {
	return 炫彩基类.X组合框_删除项(c.Handle, 项索引)
}

// 组合框_删除项扩展.
//
// iItem: 项索引.
//
// nCount: 删除数量.
func (c *ComboBox) X删除项EX(项索引 int, 删除数量 int) bool {
	return 炫彩基类.X组合框_删除项EX(c.Handle, 项索引, 删除数量)
}

// 组合框_删除项全部.
func (c *ComboBox) X删除项全部() int {
	return 炫彩基类.X组合框_删除项全部(c.Handle)
}

// 组合框_删除列全部.
func (c *ComboBox) X删除列全部() int {
	return 炫彩基类.X组合框_删除列全部(c.Handle)
}

// 组合框_取项数量.
func (c *ComboBox) X取项数量() int {
	return 炫彩基类.X组合框_取项数量(c.Handle)
}

// 组合框_取列数量.
func (c *ComboBox) X取列数量() int {
	return 炫彩基类.X组合框_取列数量(c.Handle)
}

// 组合框_弹出下拉列表.
func (c *ComboBox) X弹出下拉列表() int {
	return 炫彩基类.X组合框_弹出下拉列表(c.Handle)
}

// 组合框_设置项模板.
//
// hTemp: 模板句柄.
func (c *ComboBox) X设置项模板(模板句柄 int) int {
	return 炫彩基类.X组合框_设置项模板(c.Handle, 模板句柄)
}

// 组合框_置项模板从内存.
//
// data: 模板数据.
func (c *ComboBox) X置项模板并按内存(模板数据 []byte) bool {
	return 炫彩基类.X组合框_置项模板从内存(c.Handle, 模板数据)
}

// 组合框_置项模板从资源ZIP.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (c *ComboBox) X置项模板并按资源ZIP(RC资源ID int32, 文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	return 炫彩基类.X组合框_置项模板从资源ZIP(c.Handle, RC资源ID, 文件名, zip密码, 模块句柄)
}

// 组合框_取项模板, 返回项模板句柄.
func (c *ComboBox) X取项模板() int {
	return 炫彩基类.X组合框_取项模板(c.Handle)
}

/*
下面都是事件
*/

type XE_COMBOBOX_SELECT_END func(iItem int32, pbHandled *bool) int                          // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
type XE_COMBOBOX_SELECT_END1 func(hEle int, iItem int32, pbHandled *bool) int               // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
type XE_COMBOBOX_SELECT func(iItem int32, pbHandled *bool) int                              // 组合框下拉列表项选择事件.
type XE_COMBOBOX_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int                   // 组合框下拉列表项选择事件.
type XE_COMBOBOX_POPUP_LIST func(hWindow int, hListBox int, pbHandled *bool) int            // 组合框下拉列表弹出事件.
type XE_COMBOBOX_POPUP_LIST1 func(hEle int, hWindow int, hListBox int, pbHandled *bool) int // 组合框下拉列表弹出事件.
type XE_COMBOBOX_EXIT_LIST func(pbHandled *bool) int                                        // 组合框下拉列表退出事件.
type XE_COMBOBOX_EXIT_LIST1 func(hEle int, pbHandled *bool) int                             // 组合框下拉列表退出事件.

// 事件_组合框_下拉列表项选择完成, 编辑框内容已经改变.
func (c *ComboBox) X事件_下拉列表项选择完成(pFun XE_COMBOBOX_SELECT_END) bool {
	return 炫彩基类.X元素_注册事件C(c.Handle, 炫彩常量类.XE_COMBOBOX_SELECT_END, pFun)
}

// 事件_组合框_下拉列表项选择完成, 编辑框内容已经改变.
func (c *ComboBox) X事件_下拉列表项选择完成1(pFun XE_COMBOBOX_SELECT_END1) bool {
	return 炫彩基类.X元素_注册事件C1(c.Handle, 炫彩常量类.XE_COMBOBOX_SELECT_END, pFun)
}

// 组合框下拉列表项选择事件.
func (c *ComboBox) X事件_下拉列表项选择(pFun XE_COMBOBOX_SELECT) bool {
	return 炫彩基类.X元素_注册事件C(c.Handle, 炫彩常量类.XE_COMBOBOX_SELECT, pFun)
}

// 组合框下拉列表项选择事件.
func (c *ComboBox) X事件_下拉列表项选择1(pFun XE_COMBOBOX_SELECT1) bool {
	return 炫彩基类.X元素_注册事件C1(c.Handle, 炫彩常量类.XE_COMBOBOX_SELECT, pFun)
}

// 组合框下拉列表弹出事件.
func (c *ComboBox) X事件_下拉列表弹出(pFun XE_COMBOBOX_POPUP_LIST) bool {
	return 炫彩基类.X元素_注册事件C(c.Handle, 炫彩常量类.XE_COMBOBOX_POPUP_LIST, pFun)
}

// 组合框下拉列表弹出事件.
func (c *ComboBox) X事件_下拉列表弹出1(pFun XE_COMBOBOX_POPUP_LIST1) bool {
	return 炫彩基类.X元素_注册事件C1(c.Handle, 炫彩常量类.XE_COMBOBOX_POPUP_LIST, pFun)
}

// 组合框下拉列表退出事件.
func (c *ComboBox) X事件_下拉列表退出(pFun XE_COMBOBOX_EXIT_LIST) bool {
	return 炫彩基类.X元素_注册事件C(c.Handle, 炫彩常量类.XE_COMBOBOX_EXIT_LIST, pFun)
}

// 组合框下拉列表退出事件.
func (c *ComboBox) X事件_下拉列表退出1(pFun XE_COMBOBOX_EXIT_LIST1) bool {
	return 炫彩基类.X元素_注册事件C1(c.Handle, 炫彩常量类.XE_COMBOBOX_EXIT_LIST, pFun)
}
