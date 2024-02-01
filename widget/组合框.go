package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 组合框.
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
func I组合框_创建(x int, y int, cx int, cy int, hParent int) *ComboBox {
	p := &ComboBox{}
	p.SetHandle(xc.XComboBox_Create(x, y, cx, cy, hParent))
	return p
}

// I组合框_从句柄创建对象.
func I组合框_从句柄创建(handle int) *ComboBox {
	p := &ComboBox{}
	p.SetHandle(handle)
	return p
}

// I组合框_从name创建对象, 失败返回nil.
func I组合框_从name创建(name string) *ComboBox {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ComboBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I组合框_从UID创建对象, 失败返回nil.
func I组合框_从UID创建(nUID int) *ComboBox {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ComboBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I组合框_从UID名称创建对象, 失败返回nil.
func I组合框_从UID名称创建(name string) *ComboBox {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ComboBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 组合框_置选择项.
//
// iIndex: 项索引.
func (c *ComboBox) I置选择项(iIndex int) bool {
	return xc.XComboBox_SetSelItem(c.I句柄, iIndex)
}

// 组合框_创建数据适配器, 返回数据适配器句柄.
func (c *ComboBox) I创建数据适配器() int {
	return xc.XComboBox_CreateAdapter(c.I句柄)
}

// 组合框_绑定数据适配器.
//
// hAdapter: 适配器句柄.
func (c *ComboBox) I绑定数据适配器(hAdapter int) int {
	return xc.XComboBox_BindAdapter(c.I句柄, hAdapter)
}

// 组合框_取数据适配器, 获取绑定的数据适配器.
func (c *ComboBox) I取数据适配器() int {
	return xc.XComboBox_GetAdapter(c.I句柄)
}

// 组合框_置绑定名称.
//
// pName: 字段名.
func (c *ComboBox) I置绑定名称(pName string) int {
	return xc.XComboBox_SetBindName(c.I句柄, pName)
}

// 组合框_取下拉按钮坐标.
//
// pRect: 坐标.
func (c *ComboBox) I取下拉按钮坐标(pRect *xc.RECT) int {
	return xc.XComboBox_GetButtonRect(c.I句柄, pRect)
}

// 组合框_置下拉按钮大小.
//
// size: 大小.
func (c *ComboBox) I置下拉按钮大小(size int) int {
	return xc.XComboBox_SetButtonSize(c.I句柄, size)
}

// 组合框_置下拉列表高度.
//
// height: 高度, -1自动计算高度.
func (c *ComboBox) I置下拉列表高度(height int) int {
	return xc.XComboBox_SetDropHeight(c.I句柄, height)
}

// 组合框_取下拉列表高度.
func (c *ComboBox) I取下拉列表高度() int {
	return xc.XComboBox_GetDropHeight(c.I句柄)
}

// 组合框_置项模板, 设置下拉列表项模板文件.
//
// pXmlFile: 项模板文件.
func (c *ComboBox) I置项模板(pXmlFile string) int {
	return xc.XComboBox_SetItemTemplateXML(c.I句柄, pXmlFile)
}

// 组合框_置项模板从字符串, 设置下拉列表项模板.
//
// pStringXML: 字符串.
func (c *ComboBox) I置项模板从字符串(pStringXML string) int {
	return xc.XComboBox_SetItemTemplateXMLFromString(c.I句柄, pStringXML)
}

// 组合框_启用绘制下拉按钮, 是否绘制下拉按钮.
//
// bEnable: 是否绘制.
func (c *ComboBox) I启用绘制下拉按钮(bEnable bool) int {
	return xc.XComboBox_EnableDrawButton(c.I句柄, bEnable)
}

// 组合框_启用编辑, 启用可编辑显示的文本内容.
//
// bEdit: TRUE可编辑.
func (c *ComboBox) I启用编辑(bEdit bool) int {
	return xc.XComboBox_EnableEdit(c.I句柄, bEdit)
}

// 组合框_启用下拉列表高度固定大小, 启用/关闭下拉列表高度固定大小.
//
// bEnable: 是否启用.
func (c *ComboBox) I启用下拉列表高度固定大小(bEnable bool) int {
	return xc.XComboBox_EnableDropHeightFixed(c.I句柄, bEnable)
}

// 组合框_取选择项, 获取组合框下拉列表中选择项索引.
func (c *ComboBox) I取选择项() int {
	return xc.XComboBox_GetSelItem(c.I句柄)
}

// 组合框_取状态, 返回: I常量_组合框状态_.
func (c *ComboBox) I取状态() xcc.I常量_组合框状态_ {
	return xc.XComboBox_GetState(c.I句柄)
}

// 组合框_添加项文本, 返回项索引.
//
// pText:.
func (c *ComboBox) I添加项文本(pText string) int {
	return xc.XComboBox_AddItemText(c.I句柄, pText)
}

// 组合框_添加项文本扩展, 返回项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) I添加项文本扩展(pName string, pText string) int {
	return xc.XComboBox_AddItemTextEx(c.I句柄, pName, pText)
}

// 组合框_添加项图片, 返回项索引.
//
// hImage: 图片句柄.
func (c *ComboBox) I添加项图片(hImage int) int {
	return xc.XComboBox_AddItemImage(c.I句柄, hImage)
}

// 组合框_添加项图片扩展, 返回项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) I添加项图片扩展(pName string, hImage int) int {
	return xc.XComboBox_AddItemImageEx(c.I句柄, pName, hImage)
}

// 组合框_插入项文本, 返回项索引.
//
// iItem: 项索引.
//
// pText: 文本.
func (c *ComboBox) I插入项文本(iItem int, pText string) int {
	return xc.XComboBox_InsertItemText(c.I句柄, iItem, pText)
}

// 组合框_插入项文本扩展, 返回项索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) I插入项文本扩展(iItem int, pName string, pText string) int {
	return xc.XComboBox_InsertItemTextEx(c.I句柄, iItem, pName, pText)
}

// 组合框_插入项图片, 返回项索引.
//
// iItem: 项索引.
//
// hImage: 图片句柄.
func (c *ComboBox) I插入项图片(iItem int, hImage int) int {
	return xc.XComboBox_InsertItemImage(c.I句柄, iItem, hImage)
}

// 组合框_插入项图片扩展, 返回项索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) I插入项图片扩展(iItem int, pName string, hImage int) int {
	return xc.XComboBox_InsertItemImageEx(c.I句柄, iItem, pName, hImage)
}

// 组合框_置项文本.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pText: 文本.
func (c *ComboBox) I置项文本(iItem int, iColumn int, pText string) bool {
	return xc.XComboBox_SetItemText(c.I句柄, iItem, iColumn, pText)
}

// 组合框_置项文本扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func (c *ComboBox) I置项文本扩展(iItem int, pName string, pText string) bool {
	return xc.XComboBox_SetItemTextEx(c.I句柄, iItem, pName, pText)
}

// 组合框_置项图片.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (c *ComboBox) I置项图片(iItem int, iColumn int, hImage int) bool {
	return xc.XComboBox_SetItemImage(c.I句柄, iItem, iColumn, hImage)
}

// 组合框_置项图片扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (c *ComboBox) I置项图片扩展(iItem int, pName string, hImage int) bool {
	return xc.XComboBox_SetItemImageEx(c.I句柄, iItem, pName, hImage)
}

// 组合框_置项整数值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 整数值.
func (c *ComboBox) I置项整数值(iItem int, iColumn int, nValue int) bool {
	return xc.XComboBox_SetItemInt(c.I句柄, iItem, iColumn, nValue)
}

// 组合框_置项指数值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// nValue: 整数值.
func (c *ComboBox) I置项指数值扩展(iItem int, pName string, nValue int) bool {
	return xc.XComboBox_SetItemIntEx(c.I句柄, iItem, pName, nValue)
}

// 组合框_置项浮点值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fFloat: 浮点数.
func (c *ComboBox) I置项浮点值(iItem int, iColumn int, fFloat float32) bool {
	return xc.XComboBox_SetItemFloat(c.I句柄, iItem, iColumn, fFloat)
}

// 组合框_置项浮点值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// fFloat: 浮点数.
func (c *ComboBox) I置项浮点值扩展(iItem int, pName string, fFloat float32) bool {
	return xc.XComboBox_SetItemFloatEx(c.I句柄, iItem, pName, fFloat)
}

// 组合框_取项文本.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (c *ComboBox) I取项文本(iItem int, iColumn int) string {
	return xc.XComboBox_GetItemText(c.I句柄, iItem, iColumn)
}

// 组合框_取项文本扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
func (c *ComboBox) I取项文本扩展(iItem int, pName string) string {
	return xc.XComboBox_GetItemTextEx(c.I句柄, iItem, pName)
}

// 组合框_取项图片.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (c *ComboBox) I取项图片(iItem int, iColumn int) int {
	return xc.XComboBox_GetItemImage(c.I句柄, iItem, iColumn)
}

// 组合框_取项图片扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
func (c *ComboBox) I取项图片扩展(iItem int, pName string) int {
	return xc.XComboBox_GetItemImageEx(c.I句柄, iItem, pName)
}

// 组合框_取项整数值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回整数值.
func (c *ComboBox) I取项整数值(iItem int, iColumn int, pOutValue *int) bool {
	return xc.XComboBox_GetItemInt(c.I句柄, iItem, iColumn, pOutValue)
}

// 组合框_取项整数值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回整数值.
func (c *ComboBox) I取项整数值扩展(iItem int, pName string, pOutValue *int) bool {
	return xc.XComboBox_GetItemIntEx(c.I句柄, iItem, pName, pOutValue)
}

// 组合框_取项浮点值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回浮点值.
func (c *ComboBox) I取项浮点值(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XComboBox_GetItemFloat(c.I句柄, iItem, iColumn, pOutValue)
}

// 组合框_取项浮点值扩展.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回浮点值.
func (c *ComboBox) I取项浮点值扩展(iItem int, pName string, pOutValue *float32) bool {
	return xc.XComboBox_GetItemFloatEx(c.I句柄, iItem, pName, pOutValue)
}

// 组合框_删除项.
//
// iItem: 项索引.
func (c *ComboBox) I删除项(iItem int) bool {
	return xc.XComboBox_DeleteItem(c.I句柄, iItem)
}

// 组合框_删除项扩展.
//
// iItem: 项索引.
//
// nCount: 删除数量.
func (c *ComboBox) I删除项扩展(iItem int, nCount int) bool {
	return xc.XComboBox_DeleteItemEx(c.I句柄, iItem, nCount)
}

// 组合框_删除项全部.
func (c *ComboBox) I删除项全部() int {
	return xc.XComboBox_DeleteItemAll(c.I句柄)
}

// 组合框_删除列全部.
func (c *ComboBox) I删除列全部() int {
	return xc.XComboBox_DeleteColumnAll(c.I句柄)
}

// 组合框_取项数量.
func (c *ComboBox) I取项数量() int {
	return xc.XComboBox_GetCount(c.I句柄)
}

// 组合框_取列数量.
func (c *ComboBox) I取列数量() int {
	return xc.XComboBox_GetCountColumn(c.I句柄)
}

// 组合框_弹出下拉列表.
func (c *ComboBox) I弹出下拉列表() int {
	return xc.XComboBox_PopupDropList(c.I句柄)
}

// 组合框_设置项模板.
//
// hTemp: 模板句柄.
func (c *ComboBox) I设置项模板(hTemp int) int {
	return xc.XComboBox_SetItemTemplate(c.I句柄, hTemp)
}

/*
下面都是事件
*/

type XE_COMBOBOX_SELECT_END func(iItem int, pbHandled *bool) int                            // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
type XE_COMBOBOX_SELECT_END1 func(hEle int, iItem int, pbHandled *bool) int                 // 组合框下拉列表项选择完成事件,编辑框内容已经改变.
type XE_COMBOBOX_SELECT func(iItem int, pbHandled *bool) int                                // 组合框下拉列表项选择事件.
type XE_COMBOBOX_SELECT1 func(hEle int, iItem int, pbHandled *bool) int                     // 组合框下拉列表项选择事件.
type XE_COMBOBOX_POPUP_LIST func(hWindow int, hListBox int, pbHandled *bool) int            // 组合框下拉列表弹出事件.
type XE_COMBOBOX_POPUP_LIST1 func(hEle int, hWindow int, hListBox int, pbHandled *bool) int // 组合框下拉列表弹出事件.
type XE_COMBOBOX_EXIT_LIST func(pbHandled *bool) int                                        // 组合框下拉列表退出事件.
type XE_COMBOBOX_EXIT_LIST1 func(hEle int, pbHandled *bool) int                             // 组合框下拉列表退出事件.

// 事件_组合框_下拉列表项 I事件_选择完成, 编辑框内容已经改变.
func (c *ComboBox) I事件_选择完成(pFun XE_COMBOBOX_SELECT_END) bool {
	return xc.XEle_RegEventC(c.I句柄, xcc.XE_COMBOBOX_SELECT_END, pFun)
}

// 事件_组合框_下拉列表项 I事件_选择完成, 编辑框内容已经改变.
func (c *ComboBox) I事件_选择完成1(pFun XE_COMBOBOX_SELECT_END1) bool {
	return xc.XEle_RegEventC1(c.I句柄, xcc.XE_COMBOBOX_SELECT_END, pFun)
}

// 组合框下拉列表 I事件_选择 事件.
func (c *ComboBox) I事件_选择(pFun XE_COMBOBOX_SELECT) bool {
	return xc.XEle_RegEventC(c.I句柄, xcc.XE_COMBOBOX_SELECT, pFun)
}

// 组合框下拉列表 I事件_选择 事件.
func (c *ComboBox) I事件_选择1(pFun XE_COMBOBOX_SELECT1) bool {
	return xc.XEle_RegEventC1(c.I句柄, xcc.XE_COMBOBOX_SELECT, pFun)
}

// 组合框下拉列表 I事件_弹出  事件.
func (c *ComboBox) I事件_弹出(pFun XE_COMBOBOX_POPUP_LIST) bool {
	return xc.XEle_RegEventC(c.I句柄, xcc.XE_COMBOBOX_POPUP_LIST, pFun)
}

// 组合框下拉列表 I事件_弹出  事件.
func (c *ComboBox) I事件_弹出1(pFun XE_COMBOBOX_POPUP_LIST1) bool {
	return xc.XEle_RegEventC1(c.I句柄, xcc.XE_COMBOBOX_POPUP_LIST, pFun)
}

// 组合框下拉列表 I事件_退出 事件.
func (c *ComboBox) I事件_退出(pFun XE_COMBOBOX_EXIT_LIST) bool {
	return xc.XEle_RegEventC(c.I句柄, xcc.XE_COMBOBOX_EXIT_LIST, pFun)
}

// 组合框下拉列表 I事件_退出 事件.
func (c *ComboBox) I事件_退出1(pFun XE_COMBOBOX_EXIT_LIST1) bool {
	return xc.XEle_RegEventC1(c.I句柄, xcc.XE_COMBOBOX_EXIT_LIST, pFun)
}
