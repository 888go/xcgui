package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 列表.
type List struct {
	ScrollView
}

// I列表_创建, 创建列表元素.
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
func I列表_创建(x int, y int, cx int, cy int, hParent int) *List {
	p := &List{}
	p.SetHandle(xc.XList_Create(x, y, cx, cy, hParent))
	return p
}

// I列表_从句柄创建对象.
func I列表_从句柄创建(handle int) *List {
	p := &List{}
	p.SetHandle(handle)
	return p
}

// I列表_从name创建对象, 失败返回nil.
func I列表_从name创建(name string) *List {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &List{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I列表_从UID创建对象, 失败返回nil.
func I列表_从UID创建(nUID int) *List {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &List{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I列表_从UID名称创建对象, 失败返回nil.
func I列表_从UID名称创建(name string) *List {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &List{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 列表_增加列, 返回位置索引.
//
// width: 列宽度.
func (l *List) I增加列(width int) int {
	return xc.XList_AddColumn(l.I句柄, width)
}

// 列表_插入列, 返回插入位置索引.
//
// width: 列宽度.
//
// iItem: 插入位置索引.
func (l *List) I插入列(width int, iItem int) int {
	return xc.XList_InsertColumn(l.I句柄, width, iItem)
}

// 列表_启用多选, 启用或关闭多选功能.
//
// bEnable: 是否启用.
func (l *List) I启用多选(bEnable bool) int {
	return xc.XList_EnableMultiSel(l.I句柄, bEnable)
}

// 列表_启用拖动更改列宽, 启用拖动改变列宽度.
//
// bEnable: 是否启用.
func (l *List) I启用拖动更改列宽(bEnable bool) int {
	return xc.XList_EnableDragChangeColumnWidth(l.I句柄, bEnable)
}

// 列表_启用垂直滚动条顶部对齐.
//
// bTop: 是否启用.
func (l *List) I启用垂直滚动条顶部对齐(bTop bool) int {
	return xc.XList_EnableVScrollBarTop(l.I句柄, bTop)
}

// 列表_启用项背景全行模式, 启用项背景全行填充模式.
//
// bFull: 是否启用.
func (l *List) I启用项背景全行模式(bFull bool) int {
	return xc.XList_EnableItemBkFullRow(l.I句柄, bFull)
}

// 列表_启用固定行高.
//
// bEnable: 是否启用.
func (l *List) I启用固定行高(bEnable bool) int {
	return xc.XList_EnableFixedRowHeight(l.I句柄, bEnable)
}

// 列表_启用模板复用.
//
// bEnable: 是否启用.
func (l *List) I启用模板复用(bEnable bool) int {
	return xc.XList_EnablemTemplateReuse(l.I句柄, bEnable)
}

// 列表_启用虚表.
//
// bEnable: 是否启用.
func (l *List) I启用虚表(bEnable bool) int {
	return xc.XList_EnableVirtualTable(l.I句柄, bEnable)
}

// 列表_置虚表行数.
//
// nRowCount: 行数.
func (l *List) I置虚表行数(nRowCount int) int {
	return xc.XList_SetVirtualRowCount(l.I句柄, nRowCount)
}

// 列表_置排序, 设置排序属性.
//
// iColumn: 列索引.
//
// iColumnAdapter: 需要排序的数据在数据适配器中的列索引.
//
// bEnable: 是否启用排序功能.
func (l *List) I置排序(iColumn int, iColumnAdapter int, bEnable bool) int {
	return xc.XList_SetSort(l.I句柄, iColumn, iColumnAdapter, bEnable)
}

// 列表_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func (l *List) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) int {
	return xc.XList_SetDrawItemBkFlags(l.I句柄, nFlags)
}

// 列表_置列宽.
//
// iItem: 列索引.
//
// width: 宽度.
func (l *List) I置列宽(iItem int, width int) int {
	return xc.XList_SetColumnWidth(l.I句柄, iItem, width)
}

// 列表_置列最小宽度.
//
// iItem: 列索引.
//
// width: 宽度.
func (l *List) I置列最小宽度(iItem int, width int) int {
	return xc.XList_SetColumnMinWidth(l.I句柄, iItem, width)
}

// 列表_置列宽度固定.
//
// iColumn: 列索引.
//
// bFixed: 是否固定宽度.
func (l *List) I置列宽度固定(iColumn int, bFixed bool) int {
	return xc.XList_SetColumnWidthFixed(l.I句柄, iColumn, bFixed)
}

// 列表_取列宽度.
//
// iColumn: 列索引.
func (l *List) I取列宽度(iColumn int) int {
	return xc.XList_GetColumnWidth(l.I句柄, iColumn)
}

// 列表_取列数量.
func (l *List) I取列数量() int {
	return xc.XList_GetColumnCount(l.I句柄)
}

// 列表_置项数据, 设置项用户数据.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// data: 用户数据.
func (l *List) I置项数据(iItem int, iSubItem int, data int) bool {
	return xc.XList_SetItemData(l.I句柄, iItem, iSubItem, data)
}

// 列表_取项数据, 获取项用户数据.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
func (l *List) I取项数据(iItem int, iSubItem int) int {
	return xc.XList_GetItemData(l.I句柄, iItem, iSubItem)
}

// 列表_置选择项.
//
// iItem: 项索引.
func (l *List) I置选择项(iItem int) bool {
	return xc.XList_SetSelectItem(l.I句柄, iItem)
}

// 列表_取选择项, 返回项索引.
func (l *List) I取选择项() int {
	return xc.XList_GetSelectItem(l.I句柄)
}

// 列表_取选择项数量, 获取选择项数量.
func (l *List) I取选择项数量() int {
	return xc.XList_GetSelectItemCount(l.I句柄)
}

// 列表_添加选择项.
//
// iItem: 项索引.
func (l *List) I添加选择项(iItem int) bool {
	return xc.XList_AddSelectItem(l.I句柄, iItem)
}

// 列表_置选择全部, 选择全部行.
func (l *List) I置选择全部() int {
	return xc.XList_SetSelectAll(l.I句柄)
}

// 列表_取全部选择, 获取全部选择的行, 返回行数量.
//
// pArray: 接收行索引数组.
//
// nArraySize: 数组大小.
func (l *List) I取全部选择(pArray *[]int32, nArraySize int) int {
	return xc.XList_GetSelectAll(l.I句柄, pArray, nArraySize)
}

// 列表_显示指定项, 滚动视图让指定项可见.
//
// iItem: 项索引.
func (l *List) I显示指定项(iItem int) int {
	return xc.XList_VisibleItem(l.I句柄, iItem)
}

// 列表_取消选择项, 取消选择指定项(这里的项可以理解为行).
//
// iItem: 项索引.
func (l *List) I取消选择项(iItem int) bool {
	return xc.XList_CancelSelectItem(l.I句柄, iItem)
}

// 列表_取消全部选择项, 取消选择所有项(这里的项可以理解为行).
func (l *List) I取消全部选择项() int {
	return xc.XList_CancelSelectAll(l.I句柄)
}

// 列表_取列表头, 获取列表头元素, 返回列表头元素句柄.
func (l *List) I取列表头() int {
	return xc.XList_GetHeaderHELE(l.I句柄)
}

// 列表_删除列.
//
// iItem: 项索引.
func (l *List) I删除列(iItem int) bool {
	return xc.XList_DeleteColumn(l.I句柄, iItem)
}

// 列表_删除全部列, 删除所有的列, 仅删除List的, 数据适配器的列不变.
func (l *List) I删除全部列() int {
	return xc.XList_DeleteColumnAll(l.I句柄)
}

// 列表_绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdTable.
func (l *List) I绑定数据适配器(hAdapter int) int {
	return xc.XList_BindAdapter(l.I句柄, hAdapter)
}

// 列表_列表头绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdMap.
func (l *List) I列表头绑定数据适配器(hAdapter int) int {
	return xc.XList_BindAdapterHeader(l.I句柄, hAdapter)
}

// 列表_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// colExtend_count:	列延伸-预计列表总列数, 默认值0; 限制最大延伸范围, 避免超出范围, 增加不必要的字段.
func (l *List) I创建数据适配器(colExtend_count int) int {
	return xc.XList_CreateAdapter(l.I句柄, colExtend_count)
}

// 列表_列表头创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *List) I列表头创建数据适配器() int {
	return xc.XList_CreateAdapterHeader(l.I句柄)
}

// 列表_取数据适配器, 返回数据适配器句柄.
func (l *List) I取数据适配器() int {
	return xc.XList_GetAdapter(l.I句柄)
}

// 列表_列表头获取数据适配器, 获取列表头数据适配器句柄.
func (l *List) I列表头获取数据适配器() int {
	return xc.XList_GetAdapterHeader(l.I句柄)
}

// 列表_置项模板文件, 设置项布局模板文件.
//
// pXmlFile: 文件名.
func (l *List) I置项模板文件(pXmlFile string) bool {
	return xc.XList_SetItemTemplateXML(l.I句柄, pXmlFile)
}

// 列表_置项模板从字符串, 设置项布局模板文件.
//
// pStringXML: 字符串.
func (l *List) I置项模板从字符串(pStringXML string) bool {
	return xc.XList_SetItemTemplateXMLFromString(l.I句柄, pStringXML)
}

// 列表_置项模板, 设置列表项模板.
//
// hTemp: 模板句柄.
func (l *List) I置项模板(hTemp int) bool {
	return xc.XList_SetItemTemplate(l.I句柄, hTemp)
}

// 列表_取项模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// nTempItemID: 模板项itemID.
func (l *List) I取项模板对象(iItem int, iSubItem int, nTempItemID int) int {
	return xc.XList_GetTemplateObject(l.I句柄, iItem, iSubItem, nTempItemID)
}

// 列表_取对象所在行, 获取当前对象所在模板实例, 属于列表中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func (l *List) I取对象所在行(hXCGUI int) int {
	return xc.XList_GetItemIndexFromHXCGUI(l.I句柄, hXCGUI)
}

// 列表_取列表头模板对象, 列表头, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iItem: 列表头项ID.
//
// nTempItemID: 模板项ID.
func (l *List) I取列表头模板对象(iItem int, nTempItemID int) int {
	return xc.XList_GetHeaderTemplateObject(l.I句柄, iItem, nTempItemID)
}

// 列表_取列表头对象所在行, 列表头, 获取当前对象所在模板实例, 属于列表头中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄.
func (l *List) I取列表头对象所在行(hXCGUI int) int {
	return xc.XList_GetHeaderItemIndexFromHXCGUI(l.I句柄, hXCGUI)
}

// 列表_置列表头高度.
//
// height: 高度.
func (l *List) I置列表头高度(height int) int {
	return xc.XList_SetHeaderHeight(l.I句柄, height)
}

// 列表_取列表头高度.
func (l *List) I取列表头高度() int {
	return xc.XList_GetHeaderHeight(l.I句柄)
}

// 列表_取可视行范围.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func (l *List) I取可视行范围(piStart *int, piEnd *int) int {
	return xc.XList_GetVisibleRowRange(l.I句柄, piStart, piEnd)
}

// 列表_置项默认高度.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (l *List) I置项默认高度(nHeight int, nSelHeight int) int {
	return xc.XList_SetItemHeightDefault(l.I句柄, nHeight, nSelHeight)
}

// 列表_取项默认高度.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *List) I取项默认高度(pHeight *int, pSelHeight *int) int {
	return xc.XList_GetItemHeightDefault(l.I句柄, pHeight, pSelHeight)
}

// 列表_置行间距.
//
// nSpace: 行间距大小.
func (l *List) I置行间距(nSpace int) int {
	return xc.XList_SetRowSpace(l.I句柄, nSpace)
}

// 列表_取行间距.
func (l *List) I取行间距() int {
	return xc.XList_GetRowSpace(l.I句柄)
}

// 列表_置锁定列左侧, 锁定列, 设置左侧锁定列分界列索引.
//
// iColumn: 列索引, -1代表不锁定.
func (l *List) I置锁定列左侧(iColumn int) int {
	return xc.XList_SetLockColumnLeft(l.I句柄, iColumn)
}

// 列表_置锁定列右侧.
//
// iColumn: 列索引, -1代表不锁定. 暂时只支持锁定末尾列.
func (l *List) I置锁定列右侧(iColumn int) int {
	return xc.XList_SetLockColumnRight(l.I句柄, iColumn)
}

// 列表_置锁定行底部, 设置是否锁定末尾行.
//
// bLock: 是否锁定.
func (l *List) I置锁定行底部(bLock bool) int {
	return xc.XList_SetLockRowBottom(l.I句柄, bLock)
}

// 列表_置锁定行底部重叠.
//
// bOverlap: 是否重叠.
func (l *List) I置锁定行底部重叠(bOverlap bool) int {
	return xc.XList_SetLockRowBottomOverlap(l.I句柄, bOverlap)
}

// 列表_测试点击项, 检测坐标点所在项.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func (l *List) I测试点击项(pPt *xc.POINT, piItem *int, piSubItem *int) bool {
	return xc.XList_HitTest(l.I句柄, pPt, piItem, piSubItem)
}

// 列表_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func (l *List) I测试点击项扩展(pPt *xc.POINT, piItem *int, piSubItem *int) bool {
	return xc.XList_HitTestOffset(l.I句柄, pPt, piItem, piSubItem)
}

// 列表_刷新项数据.
func (l *List) I刷新项数据() int {
	return xc.XList_RefreshData(l.I句柄)
}

// 列表_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// iItem: 项索引.
func (l *List) I刷新指定项(iItem int) int {
	return xc.XList_RefreshItem(l.I句柄, iItem)
}

// 列表_添加列文本.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen. 你可以理解为创建数据适配器后, 内部有一个表, 这个name是每一列的字段名, 这个函数就相当于给这一行的这个字段赋值, 然后List会根据这个name名读取数据来显示到界面.
//
// pText: 文本.
func (l *List) I添加列文本(nWidth int, pName string, pText string) int {
	return xc.XList_AddColumnText(l.I句柄, nWidth, pName, pText)
}

// 列表_添加列图片.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen. 你可以理解为创建数据适配器后, 内部有一个表, 这个name是每一列的字段名, 这个函数就相当于给这一行的这个字段赋值, 然后List会根据这个name名读取数据来显示到界面.
//
// hImage: 图片句柄.
func (l *List) I添加列图片(nWidth int, pName string, hImage int) int {
	return xc.XList_AddColumnImage(l.I句柄, nWidth, pName, hImage)
}

// 列表_添加项文本.
//
// pText:.
func (l *List) I添加项文本(pText string) int {
	return xc.XList_AddItemText(l.I句柄, pText)
}

// 列表_添加项文本扩展.
//
// pName:.
//
// pText:.
func (l *List) I添加项文本扩展(pName string, pText string) int {
	return xc.XList_AddItemTextEx(l.I句柄, pName, pText)
}

// 列表_添加项图片.
//
// hImage:.
func (l *List) I添加项图片(hImage int) int {
	return xc.XList_AddItemImage(l.I句柄, hImage)
}

// 列表_添加项图片扩展.
//
// pName:.
//
// hImage:.
func (l *List) I添加项图片扩展(pName string, hImage int) int {
	return xc.XList_AddItemImageEx(l.I句柄, pName, hImage)
}

// 列表_插入项文本.
//
// iItem:.
//
// pValue:.
func (l *List) I插入项文本(iItem int, pValue string) int {
	return xc.XList_InsertItemText(l.I句柄, iItem, pValue)
}

// 列表_插入项文本扩展.
//
// iItem:.
//
// pName:.
//
// pValue:.
func (l *List) I插入项文本扩展(iItem int, pName string, pValue string) int {
	return xc.XList_InsertItemTextEx(l.I句柄, iItem, pName, pValue)
}

// 列表_插入项图片.
//
// iItem:.
//
// hImage:.
func (l *List) I插入项图片(iItem int, hImage int) int {
	return xc.XList_InsertItemImage(l.I句柄, iItem, hImage)
}

// 列表_插入项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *List) I插入项图片扩展(iItem int, pName string, hImage int) int {
	return xc.XList_InsertItemImageEx(l.I句柄, iItem, pName, hImage)
}

// 列表_置项文本.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func (l *List) I置项文本(iItem int, iColumn int, pText string) bool {
	return xc.XList_SetItemText(l.I句柄, iItem, iColumn, pText)
}

// 列表_置项文本扩展.
//
// iItem:.
//
// pName:.
//
// pText:.
func (l *List) I置项文本扩展(iItem int, pName string, pText string) bool {
	return xc.XList_SetItemTextEx(l.I句柄, iItem, pName, pText)
}

// 列表_置项图片.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func (l *List) I置项图片(iItem int, iColumn int, hImage int) bool {
	return xc.XList_SetItemImage(l.I句柄, iItem, iColumn, hImage)
}

// 列表_置项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *List) I置项图片扩展(iItem int, pName string, hImage int) bool {
	return xc.XList_SetItemImageEx(l.I句柄, iItem, pName, hImage)
}

// 列表_置项指数值.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func (l *List) I置项指数值(iItem int, iColumn int, nValue int) bool {
	return xc.XList_SetItemInt(l.I句柄, iItem, iColumn, nValue)
}

// 列表_置项整数值扩展.
//
// iItem:.
//
// pName:.
//
// nValue:.
func (l *List) I置项整数值扩展(iItem int, pName string, nValue int) bool {
	return xc.XList_SetItemIntEx(l.I句柄, iItem, pName, nValue)
}

// 列表_置项浮点值.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func (l *List) I置项浮点值(iItem int, iColumn int, fFloat float32) bool {
	return xc.XList_SetItemFloat(l.I句柄, iItem, iColumn, fFloat)
}

// 列表_置项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func (l *List) I置项浮点值扩展(iItem int, pName string, fFloat float32) bool {
	return xc.XList_SetItemFloatEx(l.I句柄, iItem, pName, fFloat)
}

// 列表_取项文本.
//
// iItem:.
//
// iColumn:.
func (l *List) I取项文本(iItem int, iColumn int) string {
	return xc.XList_GetItemText(l.I句柄, iItem, iColumn)
}

// 列表_取项文本扩展.
//
// iItem:.
//
// pName:.
func (l *List) I取项文本扩展(iItem int, pName string) string {
	return xc.XList_GetItemTextEx(l.I句柄, iItem, pName)
}

// 列表_取项图片.
//
// iItem:.
//
// iColumn:.
func (l *List) I取项图片(iItem int, iColumn int) int {
	return xc.XList_GetItemImage(l.I句柄, iItem, iColumn)
}

// 列表_取项图片扩展.
//
// iItem:.
//
// pName:.
func (l *List) I取项图片扩展(iItem int, pName string) int {
	return xc.XList_GetItemImageEx(l.I句柄, iItem, pName)
}

// 列表_取项整数值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *List) I取项整数值(iItem int, iColumn int, pOutValue *int) bool {
	return xc.XList_GetItemInt(l.I句柄, iItem, iColumn, pOutValue)
}

// 列表_取项整数值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *List) I取项整数值扩展(iItem int, pName string, pOutValue *int) bool {
	return xc.XList_GetItemIntEx(l.I句柄, iItem, pName, pOutValue)
}

// 列表_取项浮点值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *List) I取项浮点值(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XList_GetItemFloat(l.I句柄, iItem, iColumn, pOutValue)
}

// 列表_取项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *List) I取项浮点值扩展(iItem int, pName string, pOutValue *float32) bool {
	return xc.XList_GetItemFloatEx(l.I句柄, iItem, pName, pOutValue)
}

// 列表_删除项.
//
// iItem:.
func (l *List) I删除项(iItem int) bool {
	return xc.XList_DeleteItem(l.I句柄, iItem)
}

// 列表_删除项扩展.
//
// iItem:.
//
// nCount:.
func (l *List) I删除项扩展(iItem int, nCount int) bool {
	return xc.XList_DeleteItemEx(l.I句柄, iItem, nCount)
}

// 列表_删除项全部.
func (l *List) I删除项全部() int {
	return xc.XList_DeleteItemAll(l.I句柄)
}

// 列表_删除列全部AD.
func (l *List) I删除列全部AD() int {
	return xc.XList_DeleteColumnAll_AD(l.I句柄)
}

// 列表_取项数量AD.
func (l *List) I取项数量AD() int {
	return xc.XList_GetCount_AD(l.I句柄)
}

// 列表_取列数量AD.
func (l *List) I取列数量AD() int {
	return xc.XList_GetCountColumn_AD(l.I句柄)
}

// 列表_置分割线颜色.
//
// color: ABGR 颜色值.
func (l *List) I置分割线颜色(color int) int {
	return xc.XList_SetSplitLineColor(l.I句柄, color)
}

// 列表_置项高度.
//
// iRow: 行索引.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (l *List) I置项高度(iRow int, nHeight int, nSelHeight int) int {
	return xc.XList_SetItemHeight(l.I句柄, iRow, nHeight, nSelHeight)
}

// 列表_取项高度.
//
// iRow: 行索引.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *List) I取项高度(iRow int, pHeight *int, pSelHeight *int) int {
	return xc.XList_GetItemHeight(l.I句柄, iRow, pHeight, pSelHeight)
}

// 列表_置拖动矩形颜色.
//
// color: ABGR 颜色值.
//
// width: 线宽度.
func (l *List) I置拖动矩形颜色(color, width int) int {
	return xc.XList_SetDragRectColor(l.I句柄, color, width)
}

/*
以下都是事件
*/

type XE_LIST_TEMP_CREATE func(pItem *xc.List_Item_, nFlag int, pbHandled *bool) int                          // 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_LIST_TEMP_CREATE1 func(hEle int, pItem *xc.List_Item_, nFlag int, pbHandled *bool) int               // 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_LIST_TEMP_CREATE_END func(pItem *xc.List_Item_, nFlag int, pbHandled *bool) int                      // 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_LIST_TEMP_CREATE_END1 func(hEle int, pItem *xc.List_Item_, nFlag int, pbHandled *bool) int           // 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_LIST_TEMP_DESTROY func(pItem *xc.List_Item_, nFlag int, pbHandled *bool) int                         // 列表元素,项模板销毁.
type XE_LIST_TEMP_DESTROY1 func(hEle int, pItem *xc.List_Item_, nFlag int, pbHandled *bool) int              // 列表元素,项模板销毁.
type XE_LIST_TEMP_ADJUST_COORDINATE func(pItem *xc.List_Item_, pbHandled *bool) int                          // 列表元素,项模板调整坐标. 已停用.
type XE_LIST_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.List_Item_, pbHandled *bool) int               // 列表元素,项模板调整坐标. 已停用.
type XE_LIST_DRAWITEM func(hDraw int, pItem *xc.List_Item_, pbHandled *bool) int                             // 列表元素,绘制项.
type XE_LIST_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.List_Item_, pbHandled *bool) int                  // 列表元素,绘制项.
type XE_LIST_SELECT func(iItem int, pbHandled *bool) int                                                     // 列表元素,项选择事件.
type XE_LIST_SELECT1 func(hEle int, iItem int, pbHandled *bool) int                                          // 列表元素,项选择事件.
type XE_LIST_HEADER_DRAWITEM func(hDraw int, pItem *xc.List_Header_Item_, pbHandled *bool) int               // 列表元素绘制列表头项.
type XE_LIST_HEADER_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.List_Header_Item_, pbHandled *bool) int    // 列表元素绘制列表头项.
type XE_LIST_HEADER_CLICK func(iItem int, pbHandled *bool) int                                               // 列表元素,列表头项点击事件.
type XE_LIST_HEADER_CLICK1 func(hEle int, iItem int, pbHandled *bool) int                                    // 列表元素,列表头项点击事件.
type XE_LIST_HEADER_WIDTH_CHANGE func(iItem int, nWidth int, pbHandled *bool) int                            // 列表元素,列表头项宽度改变事件.
type XE_LIST_HEADER_WIDTH_CHANGE1 func(hEle int, iItem int, nWidth int, pbHandled *bool) int                 // 列表元素,列表头项宽度改变事件.
type XE_LIST_HEADER_TEMP_CREATE func(pItem *xc.List_Header_Item_, pbHandled *bool) int                       // 列表元素,列表头项模板创建.
type XE_LIST_HEADER_TEMP_CREATE1 func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int            // 列表元素,列表头项模板创建.
type XE_LIST_HEADER_TEMP_CREATE_END func(pItem *xc.List_Header_Item_, pbHandled *bool) int                   // 列表元素,列表头项模板创建完成事件.
type XE_LIST_HEADER_TEMP_CREATE_END1 func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int        // 列表元素,列表头项模板创建完成事件.
type XE_LIST_HEADER_TEMP_DESTROY func(pItem *xc.List_Header_Item_, pbHandled *bool) int                      // 列表元素,列表头项模板销毁.
type XE_LIST_HEADER_TEMP_DESTROY1 func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int           // 列表元素,列表头项模板销毁.
type XE_LIST_HEADER_TEMP_ADJUST_COORDINATE func(pItem *xc.List_Header_Item_, pbHandled *bool) int            // 列表元素,列表头项模板调整坐标. 已停用.
type XE_LIST_HEADER_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int // 列表元素,列表头项模板调整坐标. 已停用.

// I事件_项模板创建 事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (l *List) I事件_项模板创建(pFun XE_LIST_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_TEMP_CREATE, pFun)
}

// I事件_项模板创建  事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (l *List) I事件_项模板创建1(pFun XE_LIST_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_TEMP_CREATE, pFun)
}

// I事件_项模板创建完成  事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (l *List) I事件_项模板创建完成(pFun XE_LIST_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_TEMP_CREATE_END, pFun)
}

// I事件_项模板创建完成  事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (l *List) I事件_项模板创建完成1(pFun XE_LIST_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_TEMP_CREATE_END, pFun)
}

// I事件_项模板销毁  .
func (l *List) I事件_项模板销毁(pFun XE_LIST_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_TEMP_DESTROY, pFun)
}

// I事件_项模板销毁.
func (l *List) I事件_项模板销毁1(pFun XE_LIST_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_TEMP_DESTROY, pFun)
}

// I事件_项模板调整坐标  . 已停用.
func (l *List) Event_LIST_TEMP_ADJUST_COORDINATE(pFun XE_LIST_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_TEMP_ADJUST_COORDINATE, pFun)
}

// I事件_项模板调整坐标  . 已停用.
func (l *List) Event_LIST_TEMP_ADJUST_COORDINATE1(pFun XE_LIST_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_TEMP_ADJUST_COORDINATE, pFun)
}

// I事件_绘制项 .
func (l *List) I事件_绘制项(pFun XE_LIST_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_DRAWITEM, pFun)
}

// I事件_绘制项 .
func (l *List) I事件_绘制项1(pFun XE_LIST_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_DRAWITEM, pFun)
}

// I事件_项选择  事件.
func (l *List) I事件_项选择(pFun XE_LIST_SELECT) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_SELECT, pFun)
}

// I事件_项选择  事件.
func (l *List) I事件_项选择1(pFun XE_LIST_SELECT1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_SELECT, pFun)
}

// I事件_绘制列表头项 .
func (l *List) I事件_绘制列表头项(pFun XE_LIST_HEADER_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_HEADER_DRAWITEM, pFun)
}

// I事件_绘制列表头项 .
func (l *List) I事件_绘制列表头项1(pFun XE_LIST_HEADER_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_HEADER_DRAWITEM, pFun)
}

// I事件_列表头项点击事件.
func (l *List) I事件_列表头项点击事件(pFun XE_LIST_HEADER_CLICK) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_HEADER_CLICK, pFun)
}

// I事件_列表头项点击事件.
func (l *List) I事件_列表头项点击事件1(pFun XE_LIST_HEADER_CLICK1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_HEADER_CLICK, pFun)
}

// I事件_列表头项宽度改变  事件.
func (l *List) I事件_列表头项宽度改变(pFun XE_LIST_HEADER_WIDTH_CHANGE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_HEADER_WIDTH_CHANGE, pFun)
}

// I事件_列表头项宽度改变  事件.
func (l *List) I事件_列表头项宽度改变1(pFun XE_LIST_HEADER_WIDTH_CHANGE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_HEADER_WIDTH_CHANGE, pFun)
}

// I事件_列表头项模板创建.
func (l *List) I事件_列表头项模板创建(pFun XE_LIST_HEADER_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_HEADER_TEMP_CREATE, pFun)
}

// I事件_列表头项模板创建.
func (l *List) I事件_列表头项模板创建1(pFun XE_LIST_HEADER_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_HEADER_TEMP_CREATE, pFun)
}

// I事件_列表头项模板创建完成  事件.
func (l *List) I事件_列表头项模板创建完成(pFun XE_LIST_HEADER_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_HEADER_TEMP_CREATE_END, pFun)
}

// I事件_列表头项模板创建完成  事件.
func (l *List) I事件_列表头项模板创建完成1(pFun XE_LIST_HEADER_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_HEADER_TEMP_CREATE_END, pFun)
}

// I事件_列表头项模板销毁.
func (l *List) I事件_列表头项模板销毁(pFun XE_LIST_HEADER_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_HEADER_TEMP_DESTROY, pFun)
}

// I事件_列表头项模板销毁.
func (l *List) I事件_列表头项模板销毁1(pFun XE_LIST_HEADER_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_HEADER_TEMP_DESTROY, pFun)
}

// I事件_列表头项模板调整坐标. 已停用.
func (l *List) Event_LIST_HEADER_TEMP_ADJUST_COORDINATE(pFun XE_LIST_HEADER_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LIST_HEADER_TEMP_ADJUST_COORDINATE, pFun)
}

// I事件_列表头项模板调整坐标. 已停用.
func (l *List) Event_LIST_HEADER_TEMP_ADJUST_COORDINATE1(pFun XE_LIST_HEADER_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LIST_HEADER_TEMP_ADJUST_COORDINATE, pFun)
}
