package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 列表框.
type ListBox struct {
	ScrollView
}

// I列表框_创建, 创建列表框元素.
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
func I列表框_创建(x int, y int, cx int, cy int, hParent int) *ListBox {
	p := &ListBox{}
	p.SetHandle(xc.XListBox_Create(x, y, cx, cy, hParent))
	return p
}

// I列表框_从句柄创建 对象.
func I列表框_从句柄创建(handle int) *ListBox {
	p := &ListBox{}
	p.SetHandle(handle)
	return p
}

// I列表框_从name创建对象, 失败返回nil.
func I列表框_从name创建(name string) *ListBox {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I列表框_从UID创建对象, 失败返回nil.
func I列表框_从UID创建(nUID int) *ListBox {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I列表框_从UID名称创建对象, 失败返回nil.
func I列表框_从UID名称创建(name string) *ListBox {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 列表框_启用固定行高.
//
// bEnable: 是否启用.
func (l *ListBox) I启用固定行高(bEnable bool) int {
	return xc.XListBox_EnableFixedRowHeight(l.I句柄, bEnable)
}

// 列表框_启用模板复用.
//
// bEnable: 是否启用.
func (l *ListBox) I启用模板复用(bEnable bool) int {
	return xc.XListBox_EnablemTemplateReuse(l.I句柄, bEnable)
}

// 列表框_启用虚表.
//
// bEnable: 是否启用.
func (l *ListBox) I启用虚表(bEnable bool) int {
	return xc.XListBox_EnableVirtualTable(l.I句柄, bEnable)
}

// 列表框_置虚表行数.
//
// nRowCount: 行数.
func (l *ListBox) I置虚表行数(nRowCount int) int {
	return xc.XListBox_SetVirtualRowCount(l.I句柄, nRowCount)
}

// 列表框_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func (l *ListBox) I置绘制项背景标志(nFlags xcc.List_DrawItemBk_Flag_) int {
	return xc.XListBox_SetDrawItemBkFlags(l.I句柄, nFlags)
}

// 列表框_置项数据, 设置项用户数据.
//
// iItem: 想索引.
//
// nUserData: 用户数据.
func (l *ListBox) I置项数据(iItem int, nUserData int) bool {
	return xc.XListBox_SetItemData(l.I句柄, iItem, nUserData)
}

// 列表框_取项数据, 获取项用户数据.
//
// iItem: 项索引.
func (l *ListBox) I取项数据(iItem int) int {
	return xc.XListBox_GetItemData(l.I句柄, iItem)
}

// 列表框_置项信息.
//
// iItem: 项索引.
//
// pItem: 项信息.
func (l *ListBox) I置项信息(iItem int, pItem *xc.ListBox_Item_Info_) bool {
	return xc.XListBox_SetItemInfo(l.I句柄, iItem, pItem)
}

// 列表框_取项背景信息, 获取项信息.
//
// iItem: 项索引.
//
// pItem: 项信息.
func (l *ListBox) I取项背景信息(iItem int, pItem *xc.ListBox_Item_Info_) bool {
	return xc.XListBox_GetItemInfo(l.I句柄, iItem, pItem)
}

// 列表框_置选择项, 设置选择选.
//
// iItem: 项索引.
func (l *ListBox) I置选择项(iItem int) bool {
	return xc.XListBox_SetSelectItem(l.I句柄, iItem)
}

// 列表框_取选择项, 返回项索引.
func (l *ListBox) I取选择项() int {
	return xc.XListBox_GetSelectItem(l.I句柄)
}

// 列表框_添加选择项.
//
// iItem: 项索引.
func (l *ListBox) I添加选择项(iItem int) bool {
	return xc.XListBox_AddSelectItem(l.I句柄, iItem)
}

// 列表框_取消选择项.
//
// iItem: 项索引.
func (l *ListBox) I取消选择项(iItem int) bool {
	return xc.XListBox_CancelSelectItem(l.I句柄, iItem)
}

// 列表框_取消选择全部, 如果之前有选择状态的项返回TRUE, 此时可以更新UI, 否则返回FALSE.
func (l *ListBox) I取消选择全部() bool {
	return xc.XListBox_CancelSelectAll(l.I句柄)
}

// 列表框_取全部选择, 获取所有选择项, 返回接收数量.
//
// pArray: 数组缓冲区.
//
// nArraySize: 数组大小.
func (l *ListBox) I取全部选择(pArray *[]int32, nArraySize int) int {
	return xc.XListBox_GetSelectAll(l.I句柄, pArray, nArraySize)
}

// 列表框_取选择项数量, 获取选择项数量.
func (l *ListBox) I取选择项数量() int {
	return xc.XListBox_GetSelectCount(l.I句柄)
}

// 列表框_取鼠标停留项, 返回鼠标所在项.
func (l *ListBox) I取鼠标停留项() int {
	return xc.XListBox_GetItemMouseStay(l.I句柄)
}

// 列表框_选择全部项.
func (l *ListBox) I选择全部项() bool {
	return xc.XListBox_SelectAll(l.I句柄)
}

// 列表框_显示指定项, 滚动视图让指定项可见.
//
// iItem: 项索引.
func (l *ListBox) I显示指定项(iItem int) int {
	return xc.XListBox_VisibleItem(l.I句柄, iItem)
}

// 列表框_取可视行范围, 获取当前可见行范围.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func (l *ListBox) I取可视行范围(piStart *int, piEnd *int) int {
	return xc.XListBox_GetVisibleRowRange(l.I句柄, piStart, piEnd)
}

// 列表框_置项默认高度.
//
// nHeight: 项高度.
//
// nSelHeight: 选中项高度.
func (l *ListBox) I置项默认高度(nHeight int, nSelHeight int) int {
	return xc.XListBox_SetItemHeightDefault(l.I句柄, nHeight, nSelHeight)
}

// 列表框_取项默认高度.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *ListBox) I取项默认高度(pHeight *int, pSelHeight *int) int {
	return xc.XListBox_GetItemHeightDefault(l.I句柄, pHeight, pSelHeight)
}

// 列表框_取所在行索引, 获取当前对象所在模板实例, 属于列表中哪一个项(行). 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func (l *ListBox) I取所在行索引(hXCGUI int) int {
	return xc.XListBox_GetItemIndexFromHXCGUI(l.I句柄, hXCGUI)
}

// 列表框_置行间距.
//
// nSpace: 间距大小.
func (l *ListBox) I置行间距(nSpace int) int {
	return xc.XListBox_SetRowSpace(l.I句柄, nSpace)
}

// 列表框_取行间距.
func (l *ListBox) I取行间距() int {
	return xc.XListBox_GetRowSpace(l.I句柄)
}

// 列表框_测试点击项, 检测坐标点所在项, 返回项索引.
//
// pPt: 坐标点.
func (l *ListBox) I测试点击项(pPt *xc.POINT) int {
	return xc.XListBox_HitTest(l.I句柄, pPt)
}

// 列表框_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移量, 返回项索引.
//
// pPt: 坐标点.
func (l *ListBox) I测试点击项扩展(pPt *xc.POINT) int {
	return xc.XListBox_HitTestOffset(l.I句柄, pPt)
}

// 列表框_置项模板文件, 设置列表项模板文件.
//
// pXmlFile: 文件名.
func (l *ListBox) I置项模板文件(pXmlFile string) bool {
	return xc.XListBox_SetItemTemplateXML(l.I句柄, pXmlFile)
}

// 列表框_置项模板, 设置列表项模板.
//
// hTemp: 模板句柄.
func (l *ListBox) I置项模板(hTemp int) bool {
	return xc.XListBox_SetItemTemplate(l.I句柄, hTemp)
}

// 列表框_置项模板从字符串, 设置项布局模板文件.
//
// pStringXML: 字符串.
func (l *ListBox) I置项模板从字符串(pStringXML string) bool {
	return xc.XListBox_SetItemTemplateXMLFromString(l.I句柄, pStringXML)
}

// 列表框_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄, 成功返回对象句柄, 否则返回NULL.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func (l *ListBox) I取模板对象(iItem int, nTempItemID int) int {
	return xc.XListBox_GetTemplateObject(l.I句柄, iItem, nTempItemID)
}

// 列表框_启用多选, 是否启用多行选择功能.
//
// bEnable: 是否启用.
func (l *ListBox) I启用多选(bEnable bool) int {
	return xc.XListBox_EnableMultiSel(l.I句柄, bEnable)
}

// 列表框_创建数据适配器, 创建数据适配器并绑定, 根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *ListBox) I创建数据适配器() int {
	return xc.XListBox_CreateAdapter(l.I句柄)
}

// 列表框_绑定数据适配器, 绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdTable.
func (l *ListBox) I绑定数据适配器(hAdapter int) int {
	return xc.XListBox_BindAdapter(l.I句柄, hAdapter)
}

// 列表框_取数据适配器, 获取绑定的数据适配器, 返回数据适配器句柄.
func (l *ListBox) I取数据适配器() int {
	return xc.XListBox_GetAdapter(l.I句柄)
}

// 列表框_排序.
//
// iColumnAdapter: 需要排序的数据在数据适配器中所属列索引.
//
// bAscending: 升序(TRUE)或降序(FALSE).
func (l *ListBox) I排序(iColumnAdapter int, bAscending bool) int {
	return xc.XListBox_Sort(l.I句柄, iColumnAdapter, bAscending)
}

// 列表框_刷新数据.
func (l *ListBox) I刷新数据() int {
	return xc.XListBox_RefreshData(l.I句柄)
}

// 列表框_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// iItem: 项索引.
func (l *ListBox) I刷新指定项(iItem int) int {
	return xc.XListBox_RefreshItem(l.I句柄, iItem)
}

// 列表框_添加项文本, XAdTable_AddItemText, 返回项索引.
//
// pText:.
func (l *ListBox) I添加项文本(pText string) int {
	return xc.XListBox_AddItemText(l.I句柄, pText)
}

// 列表框_添加项文本扩展, XAdTable_AddItemTextEx.
//
// pName:.
//
// pText:.
func (l *ListBox) I添加项文本扩展(pName string, pText string) int {
	return xc.XListBox_AddItemTextEx(l.I句柄, pName, pText)
}

// 列表框_添加项图片, XAdTable_AddItemImage.
//
// hImage:.
func (l *ListBox) I添加项图片(hImage int) int {
	return xc.XListBox_AddItemImage(l.I句柄, hImage)
}

// 列表框_添加项图片扩展, XAdTable_AddItemImageEx.
//
// pName:.
//
// hImage:.
func (l *ListBox) I添加项图片扩展(pName string, hImage int) int {
	return xc.XListBox_AddItemImageEx(l.I句柄, pName, hImage)
}

// 列表框_插入项文本.
//
// iItem:.
//
// pValue:.
func (l *ListBox) I插入项文本(iItem int, pValue string) int {
	return xc.XListBox_InsertItemText(l.I句柄, iItem, pValue)
}

// 列表框_插入项文本扩展.
//
// iItem:.
//
// pName:.
//
// pValue:.
func (l *ListBox) I插入项文本扩展(iItem int, pName string, pValue string) int {
	return xc.XListBox_InsertItemTextEx(l.I句柄, iItem, pName, pValue)
}

// 列表框_插入项图片.
//
// iItem:.
//
// hImage:.
func (l *ListBox) I插入项图片(iItem int, hImage int) int {
	return xc.XListBox_InsertItemImage(l.I句柄, iItem, hImage)
}

// 列表框_插入项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *ListBox) I插入项图片扩展(iItem int, pName string, hImage int) int {
	return xc.XListBox_InsertItemImageEx(l.I句柄, iItem, pName, hImage)
}

// 列表框_置项文本.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func (l *ListBox) I置项文本(iItem int, iColumn int, pText string) bool {
	return xc.XListBox_SetItemText(l.I句柄, iItem, iColumn, pText)
}

// 列表框_置项文本扩展.
//
// iItem:.
//
// pName:.
//
// pText:.
func (l *ListBox) I置项文本扩展(iItem int, pName string, pText string) bool {
	return xc.XListBox_SetItemTextEx(l.I句柄, iItem, pName, pText)
}

// 列表框_置项图片.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func (l *ListBox) I置项图片(iItem int, iColumn int, hImage int) bool {
	return xc.XListBox_SetItemImage(l.I句柄, iItem, iColumn, hImage)
}

// 列表框_置项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *ListBox) I置项图片扩展(iItem int, pName string, hImage int) bool {
	return xc.XListBox_SetItemImageEx(l.I句柄, iItem, pName, hImage)
}

// 列表框_置项整数值.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func (l *ListBox) I置项整数值(iItem int, iColumn int, nValue int) bool {
	return xc.XListBox_SetItemInt(l.I句柄, iItem, iColumn, nValue)
}

// 列表框_置项整数值扩展.
//
// iItem:.
//
// pName:.
//
// nValue:.
func (l *ListBox) I置项整数值扩展(iItem int, pName string, nValue int) bool {
	return xc.XListBox_SetItemIntEx(l.I句柄, iItem, pName, nValue)
}

// 列表框_置项浮点值.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func (l *ListBox) I置项浮点值(iItem int, iColumn int, fFloat float32) bool {
	return xc.XListBox_SetItemFloat(l.I句柄, iItem, iColumn, fFloat)
}

// 列表框_置项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func (l *ListBox) I置项浮点值扩展(iItem int, pName string, fFloat float32) bool {
	return xc.XListBox_SetItemFloatEx(l.I句柄, iItem, pName, fFloat)
}

// 列表框_取项文本.
//
// iItem:.
//
// iColumn:.
func (l *ListBox) I取项文本(iItem int, iColumn int) string {
	return xc.XListBox_GetItemText(l.I句柄, iItem, iColumn)
}

// 列表框_取项文本扩展.
//
// iItem:.
//
// pName:.
func (l *ListBox) I取项文本扩展(iItem int, pName string) string {
	return xc.XListBox_GetItemTextEx(l.I句柄, iItem, pName)
}

// 列表框_取项图片.
//
// iItem:.
//
// iColumn:.
func (l *ListBox) I取项图片(iItem int, iColumn int) int {
	return xc.XListBox_GetItemImage(l.I句柄, iItem, iColumn)
}

// 列表框_取项图片扩展.
//
// iItem:.
//
// pName:.
func (l *ListBox) I取项图片扩展(iItem int, pName string) int {
	return xc.XListBox_GetItemImageEx(l.I句柄, iItem, pName)
}

// 列表框_取项整数值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *ListBox) I取项整数值(iItem int, iColumn int, pOutValue *int) bool {
	return xc.XListBox_GetItemInt(l.I句柄, iItem, iColumn, pOutValue)
}

// 列表框_取项整数值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *ListBox) I取项整数值扩展(iItem int, pName string, pOutValue *int) bool {
	return xc.XListBox_GetItemIntEx(l.I句柄, iItem, pName, pOutValue)
}

// 列表框_取项浮点值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *ListBox) I取项浮点值(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XListBox_GetItemFloat(l.I句柄, iItem, iColumn, pOutValue)
}

// 列表框_取项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *ListBox) I取项浮点值扩展(iItem int, pName string, pOutValue *float32) bool {
	return xc.XListBox_GetItemFloatEx(l.I句柄, iItem, pName, pOutValue)
}

// 列表框_删除项.
//
// iItem:.
func (l *ListBox) I删除项(iItem int) bool {
	return xc.XListBox_DeleteItem(l.I句柄, iItem)
}

// 列表框_删除项扩展.
//
// iItem:.
//
// nCount:.
func (l *ListBox) I删除项扩展(iItem int, nCount int) bool {
	return xc.XListBox_DeleteItemEx(l.I句柄, iItem, nCount)
}

// 列表框_删除项全部.
func (l *ListBox) I删除项全部() int {
	return xc.XListBox_DeleteItemAll(l.I句柄)
}

// 列表框_删除列全部.
func (l *ListBox) I删除列全部() int {
	return xc.XListBox_DeleteColumnAll(l.I句柄)
}

// 列表框_取项数量AD.
func (l *ListBox) I取项数量AD() int {
	return xc.XListBox_GetCount_AD(l.I句柄)
}

// 列表框_取列数量AD.
func (l *ListBox) I取列数量AD() int {
	return xc.XListBox_GetCountColumn_AD(l.I句柄)
}

// 列表框_置分割线颜色.
//
// color: ABGR 颜色值.
func (l *ListBox) I置分割线颜色(color int) int {
	return xc.XListBox_SetSplitLineColor(l.I句柄, color)
}

// 列表框_置拖动矩形颜色.
//
// color: ABGR 颜色值.
//
// width: 线宽度.
func (l *ListBox) I置拖动矩形颜色(color, width int) int {
	return xc.XListBox_SetDragRectColor(l.I句柄, color, width)
}

/*
以下都是事件
*/

type XE_LISTBOX_TEMP_CREATE func(pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int                // 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_LISTBOX_TEMP_CREATE1 func(hEle int, pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int     // 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_LISTBOX_TEMP_CREATE_END func(pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int            // 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_LISTBOX_TEMP_CREATE_END1 func(hEle int, pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int // 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_LISTBOX_TEMP_DESTROY func(pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int               // 列表框元素,项模板销毁.
type XE_LISTBOX_TEMP_DESTROY1 func(hEle int, pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int    // 列表框元素,项模板销毁.
type XE_LISTBOX_TEMP_ADJUST_COORDINATE func(pItem *xc.ListBox_Item_, pbHandled *bool) int                // 列表框元素,项模板调整坐标. 已停用.
type XE_LISTBOX_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.ListBox_Item_, pbHandled *bool) int     // 列表框元素,项模板调整坐标. 已停用.
type XE_LISTBOX_DRAWITEM func(hDraw int, pItem *xc.ListBox_Item_, pbHandled *bool) int                   // 列表框元素,项绘制事件.
type XE_LISTBOX_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.ListBox_Item_, pbHandled *bool) int        // 列表框元素,项绘制事件.
type XE_LISTBOX_SELECT func(iItem int, pbHandled *bool) int                                              // 列表框元素,项选择事件.
type XE_LISTBOX_SELECT1 func(hEle int, iItem int, pbHandled *bool) int                                   // 列表框元素,项选择事件.

// 列表框元素-  I事件_项模板创建  事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (l *ListBox) I事件_项模板创建(pFun XE_LISTBOX_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTBOX_TEMP_CREATE, pFun)
}

//  I事件_项模板创建  事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (l *ListBox) I事件_项模板创建1(pFun XE_LISTBOX_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTBOX_TEMP_CREATE, pFun)
}

//  I事件_项模板创建完成 事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (l *ListBox) I事件_项模板创建完成(pFun XE_LISTBOX_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTBOX_TEMP_CREATE_END, pFun)
}

// I事件_项模板创建完成  事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (l *ListBox) I事件_项模板创建完成1(pFun XE_LISTBOX_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTBOX_TEMP_CREATE_END, pFun)
}

// I事件_项模板销毁.
func (l *ListBox) I事件_项模板销毁(pFun XE_LISTBOX_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTBOX_TEMP_DESTROY, pFun)
}

// I事件_项模板销毁.
func (l *ListBox) I事件_项模板销毁1(pFun XE_LISTBOX_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTBOX_TEMP_DESTROY, pFun)
}

// I事件_项模板调整坐标 . 已停用.
func (l *ListBox) Event_LISTBOX_TEMP_ADJUST_COORDINATE(pFun XE_LISTBOX_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun)
}

// I事件_项模板调整坐标. 已停用.
func (l *ListBox) Event_LISTBOX_TEMP_ADJUST_COORDINATE1(pFun XE_LISTBOX_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun)
}

// I事件_项绘制事件.
func (l *ListBox) I事件_项绘制事件(pFun XE_LISTBOX_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTBOX_DRAWITEM, pFun)
}

// I事件_项绘制事件.
func (l *ListBox) I事件_项绘制事件1(pFun XE_LISTBOX_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTBOX_DRAWITEM, pFun)
}

// I事件_项选择事件.
func (l *ListBox) I事件_项选择事件(pFun XE_LISTBOX_SELECT) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTBOX_SELECT, pFun)
}

// I事件_项选择事件.
func (l *ListBox) I事件_项选择事件1(pFun XE_LISTBOX_SELECT1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTBOX_SELECT, pFun)
}
