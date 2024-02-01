package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 列表视图.
type ListView struct {
	ScrollView
}

// I列表视图_创建.
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
func I列表视图_创建(x int, y int, cx int, cy int, hParent int) *ListView {
	p := &ListView{}
	p.SetHandle(xc.XListView_Create(x, y, cx, cy, hParent))
	return p
}

// I列表视图_从句柄创建对象.
func I列表视图_从句柄创建(handle int) *ListView {
	p := &ListView{}
	p.SetHandle(handle)
	return p
}

// I列表视图_从name创建对象, 失败返回nil.
func I列表视图_从name创建(name string) *ListView {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ListView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I列表视图_从UID创建对象, 失败返回nil.
func I列表视图_从UID创建(nUID int) *ListView {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ListView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I列表视图_从UID名称创建对象, 失败返回nil.
func I列表视图_从UID名称创建(name string) *ListView {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ListView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 列表视_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *ListView) I创建数据适配器() int {
	return xc.XListView_CreateAdapter(l.I句柄)
}

// 列表视_绑定数据适配器.
//
// hAdapter: 数据适配器XAdListView.
func (l *ListView) I绑定数据适配器(hAdapter int) int {
	return xc.XListView_BindAdapter(l.I句柄, hAdapter)
}

// 列表视_取数据适配器, 返回数据适配器句柄.
func (l *ListView) I取数据适配器() int {
	return xc.XListView_GetAdapter(l.I句柄)
}

// 列表视_置项模板文件.
//
// pXmlFile: 文件名.
func (l *ListView) I置项模板文件(pXmlFile string) bool {
	return xc.XListView_SetItemTemplateXML(l.I句柄, pXmlFile)
}

// 列表视_置项模板从字符串.
//
// pStringXML: 字符串.
func (l *ListView) I置项模板从字符串(pStringXML string) bool {
	return xc.XListView_SetItemTemplateXMLFromString(l.I句柄, pStringXML)
}

// 列表视_置项模板, 置列表项模板.
//
// hTemp: 模板句柄.
func (l *ListView) I置项模板(hTemp int) bool {
	return xc.XListView_SetItemTemplate(l.I句柄, hTemp)
}

// 列表视_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func (l *ListView) I取模板对象(iGroup int, iItem int, nTempItemID int) int {
	return xc.XListView_GetTemplateObject(l.I句柄, iGroup, iItem, nTempItemID)
}

// 列表视_取模板对象组, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iGroup: 组索引.
//
// nTempItemID: 模板项ID.
func (l *ListView) I取模板对象组(iGroup int, nTempItemID int) int {
	return xc.XListView_GetTemplateObjectGroup(l.I句柄, iGroup, nTempItemID)
}

// 列表视_取对象所在项, 获取当前对象所在模板实例, 属于列表视中哪一个项.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func (l *ListView) I取对象所在项(hXCGUI int, piGroup *int, piItem *int) bool {
	return xc.XListView_GetItemIDFromHXCGUI(l.I句柄, hXCGUI, piGroup, piItem)
}

// 列表视_测试点击项, 检查坐标点所在项.
//
// pPt: 坐标点.
//
// pOutGroup: 接收组索引.
//
// pOutItem: 接收项索引.
func (l *ListView) I测试点击项(pPt *xc.POINT, pOutGroup *int, pOutItem *int) bool {
	return xc.XListView_HitTest(l.I句柄, pPt, pOutGroup, pOutItem)
}

// 列表视_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// pPt: 坐标点.
//
// pOutGroup: 接收做索引.
//
// pOutItem: 接收项索引.
func (l *ListView) I测试点击项扩展(pPt *xc.POINT, pOutGroup *int, pOutItem *int) bool {
	return xc.XListView_HitTestOffset(l.I句柄, pPt, pOutGroup, pOutItem)
}

// 列表视_启用多选.
//
// bEnable: 是否启用.
func (l *ListView) I启用多选(bEnable bool) int {
	return xc.XListView_EnableMultiSel(l.I句柄, bEnable)
}

// 列表视_启用模板复用.
//
// bEnable: 是否启用.
func (l *ListView) I启用模板复用(bEnable bool) int {
	return xc.XListView_EnablemTemplateReuse(l.I句柄, bEnable)
}

// 列表视_启用虚表.
//
// bEnable: 是否启用.
func (l *ListView) I启用虚表(bEnable bool) int {
	return xc.XListView_EnableVirtualTable(l.I句柄, bEnable)
}

// 列表视_置虚表项数量.
//
// iGroup: 组索引.
//
// nCount: 项数量.
func (l *ListView) I置虚表项数量(iGroup int, nCount int) bool {
	return xc.XListView_SetVirtualItemCount(l.I句柄, iGroup, nCount)
}

// 列表视_置项背景绘制标志, 置是否绘制指定状态下项的背景.
//
// nFlags: 标志位: List_DrawItemBk_Flag_.
func (l *ListView) I置项背景绘制标志(nFlags xcc.List_DrawItemBk_Flag_) int {
	return xc.XListView_SetDrawItemBkFlags(l.I句柄, nFlags)
}

// 列表视_置选择项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) I置选择项(iGroup int, iItem int) bool {
	return xc.XListView_SetSelectItem(l.I句柄, iGroup, iItem)
}

// 列表视_取选择项.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func (l *ListView) I取选择项(piGroup *int, piItem *int) bool {
	return xc.XListView_GetSelectItem(l.I句柄, piGroup, piItem)
}

// 列表视_添加选择项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) I添加选择项(iGroup int, iItem int) bool {
	return xc.XListView_AddSelectItem(l.I句柄, iGroup, iItem)
}

// 列表视_显示指定项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) I显示指定项(iGroup int, iItem int) int {
	return xc.XListView_VisibleItem(l.I句柄, iGroup, iItem)
}

// 列表视_取可视项范围, 获取当前可见项范围.
//
// piGroup1: 可视开始组.
//
// piGroup2: 可视结束组.
//
// piStartGroup: 可视开始组.
//
// piStartItem: 可视开始项.
//
// piEndGroup: 可视结束组.
//
// piEndItem: 可视结束项.
func (l *ListView) I取可视项范围(piGroup1 *int, piGroup2 *int, piStartGroup *int, piStartItem *int, piEndGroup *int, piEndItem *int) int {
	return xc.XListView_GetVisibleItemRange(l.I句柄, piGroup1, piGroup2, piStartGroup, piStartItem, piEndGroup, piEndItem)
}

// 列表视_取选择项数量.
func (l *ListView) I取选择项数量() int {
	return xc.XListView_GetSelectItemCount(l.I句柄)
}

// 列表视_取选择项全部, 获取选择的项ID, 返回接收项数量.
//
// pArray: 数组.
//
// nArraySize: 数组大小.
func (l *ListView) I取选择项全部(pArray int, nArraySize int) int {
	return xc.XListView_GetSelectAll(l.I句柄, pArray, nArraySize)
}

// 列表视_置选择项全部, 选择所有的项.
func (l *ListView) I置选择项全部() int {
	return xc.XListView_SetSelectAll(l.I句柄)
}

// 列表视_取消选择项全部, 取消选择所有项.
func (l *ListView) I取消选择项全部() int {
	return xc.XListView_CancelSelectAll(l.I句柄)
}

// 列表视_置列间隔, 置列间隔大小.
//
// space: 间隔大小.
func (l *ListView) I置列间隔(space int) int {
	return xc.XListView_SetColumnSpace(l.I句柄, space)
}

// 列表视_置行间隔, 置行间隔大小.
//
// space: 间隔大小.
func (l *ListView) I置行间隔(space int) int {
	return xc.XListView_SetRowSpace(l.I句柄, space)
}

// 列表视_置项大小.
//
// width: 宽度.
//
// height: 高度.
func (l *ListView) I置项大小(width int, height int) int {
	return xc.XListView_SetItemSize(l.I句柄, width, height)
}

// 列表视_取项大小.
//
// pSize: 接收返回大小.
func (l *ListView) I取项大小(pSize *xc.SIZE) int {
	return xc.XListView_GetItemSize(l.I句柄, pSize)
}

// 列表视_置组高度.
//
// height: 高度.
func (l *ListView) I置组高度(height int) int {
	return xc.XListView_SetGroupHeight(l.I句柄, height)
}

// 列表视_取组高度.
func (l *ListView) I取组高度() int {
	return xc.XListView_GetGroupHeight(l.I句柄)
}

// 列表视_置组用户数据.
//
// iGroup: 组索引.
//
// nData: 数据.
func (l *ListView) I置组用户数据(iGroup int, nData int) int {
	return xc.XListView_SetGroupUserData(l.I句柄, iGroup, nData)
}

// 列表视_置项用户数据.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nData: 数据.
func (l *ListView) I置项用户数据(iGroup int, iItem int, nData int) int {
	return xc.XListView_SetItemUserData(l.I句柄, iGroup, iItem, nData)
}

// 列表视_取组用户数据.
//
// iGroup: 组索引.
func (l *ListView) I取组用户数据(iGroup int) int {
	return xc.XListView_GetGroupUserData(l.I句柄, iGroup)
}

// 列表视_取项用户数据.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) I取项用户数据(iGroup int, iItem int) int {
	return xc.XListView_GetItemUserData(l.I句柄, iGroup, iItem)
}

// 列表视_刷新项数据.
func (l *ListView) I刷新项数据() int {
	return xc.XListView_RefreshData(l.I句柄)
}

// 列表视_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// iGroup: 组索引.
//
// iItem: 项索引, 如果为-1, 代表为组.
func (l *ListView) I刷新指定项(iGroup int, iItem int) int {
	return xc.XListView_RefreshItem(l.I句柄, iGroup, iItem)
}

// 列表视_展开组, 成功返回TRUE否则返回FALSE, 如果状态没有改变返回FALSE.
//
// iGroup: 组索引.
//
// bExpand: 是否展开.
func (l *ListView) I展开组(iGroup int, bExpand bool) bool {
	return xc.XListView_ExpandGroup(l.I句柄, iGroup, bExpand)
}

// 列表视_组添加列, 返回列索引.
//
// pName: 字段称.
func (l *ListView) I组添加列(pName string) int {
	return xc.XListView_Group_AddColumn(l.I句柄, pName)
}

// 列表视_组添加项文本, 返回组索引.
//
// pValue: 值.
//
// iPos: 插入位置.
func (l *ListView) I组添加项文本(pValue string, iPos int) int {
	return xc.XListView_Group_AddItemText(l.I句柄, pValue, iPos)
}

// 列表视_组添加项文本扩展, 返回组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置.
func (l *ListView) I组添加项文本扩展(pName string, pValue string, iPos int) int {
	return xc.XListView_Group_AddItemTextEx(l.I句柄, pName, pValue, iPos)
}

// 列表视_组添加项图片, 返回组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func (l *ListView) I组添加项图片(hImage int, iPos int) int {
	return xc.XListView_Group_AddItemImage(l.I句柄, hImage, iPos)
}

// 列表视_组添加项图片扩展, 返回组索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func (l *ListView) I组添加项图片扩展(pName string, hImage int, iPos int) int {
	return xc.XListView_Group_AddItemImageEx(l.I句柄, pName, hImage, iPos)
}

// 列表视_组置文本.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (l *ListView) I组置文本(iGroup int, iColumn int, pValue string) bool {
	return xc.XListView_Group_SetText(l.I句柄, iGroup, iColumn, pValue)
}

// 列表视_组置文本扩展.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func (l *ListView) I组置文本扩展(iGroup int, pName string, pValue string) bool {
	return xc.XListView_Group_SetTextEx(l.I句柄, iGroup, pName, pValue)
}

// 列表视_组置图片.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (l *ListView) I组置图片(iGroup int, iColumn int, hImage int) bool {
	return xc.XListView_Group_SetImage(l.I句柄, iGroup, iColumn, hImage)
}

// 列表视_组置图片扩展.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (l *ListView) I组置图片扩展(iGroup int, pName string, hImage int) bool {
	return xc.XListView_Group_SetImageEx(l.I句柄, iGroup, pName, hImage)
}

// 列表视_组获取数量, 返回组数量.
func (l *ListView) I组获取数量() int {
	return xc.XListView_Group_GetCount(l.I句柄)
}

// 列表视_项获取数量, 成功返回项数量, 否则返回 I常量_错误.
//
// iGroup: 组索引.
func (l *ListView) I项获取数量(iGroup int) int {
	return xc.XListView_Item_GetCount(l.I句柄, iGroup)
}

// 列表视_项添加列, 返回列索引.
//
// pName: 字段名.
func (l *ListView) I项添加列(pName string) int {
	return xc.XListView_Item_AddColumn(l.I句柄, pName)
}

// 列表视_项添加文本, 返回项索引.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) I项添加文本(iGroup int, pValue string, iPos int) int {
	return xc.XListView_Item_AddItemText(l.I句柄, iGroup, pValue, iPos)
}

// 列表视_项添加文本扩展, 返回项索引.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) I项添加文本扩展(iGroup int, pName string, pValue string, iPos int) int {
	return xc.XListView_Item_AddItemTextEx(l.I句柄, iGroup, pName, pValue, iPos)
}

// 列表视_项添加图片, 返回项索引.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) I项添加图片(iGroup int, hImage int, iPos int) int {
	return xc.XListView_Item_AddItemImage(l.I句柄, iGroup, hImage, iPos)
}

// 列表视_项添加图片扩展, 返回项索引.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) I项添加图片扩展(iGroup int, pName string, hImage int, iPos int) int {
	return xc.XListView_Item_AddItemImageEx(l.I句柄, iGroup, pName, hImage, iPos)
}

// 列表视_项置文本.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (l *ListView) I项置文本(iGroup int, iItem int, iColumn int, pValue string) bool {
	return xc.XListView_Item_SetText(l.I句柄, iGroup, iItem, iColumn, pValue)
}

// 列表视_项置文本扩展.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pValue: 值.
func (l *ListView) I项置文本扩展(iGroup int, iItem int, pName string, pValue string) bool {
	return xc.XListView_Item_SetTextEx(l.I句柄, iGroup, iItem, pName, pValue)
}

// 列表视_项置图片.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (l *ListView) I项置图片(iGroup int, iItem int, iColumn int, hImage int) bool {
	return xc.XListView_Item_SetImage(l.I句柄, iGroup, iItem, iColumn, hImage)
}

// 列表视_项置图片扩展.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 列名称.
//
// hImage: 图片句柄.
func (l *ListView) I项置图片扩展(iGroup int, iItem int, pName string, hImage int) bool {
	return xc.XListView_Item_SetImageEx(l.I句柄, iGroup, iItem, pName, hImage)
}

// 列表视_组删除项.
//
// iGroup: 组索引.
func (l *ListView) I组删除项(iGroup int) bool {
	return xc.XListView_Group_DeleteItem(l.I句柄, iGroup)
}

// 列表视_组删除全部子项.
//
// iGroup: 组索引.
func (l *ListView) I组删除全部子项(iGroup int) int {
	return xc.XListView_Group_DeleteAllChildItem(l.I句柄, iGroup)
}

// 列表视_项删除.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) I项删除(iGroup int, iItem int) bool {
	return xc.XListView_Item_DeleteItem(l.I句柄, iGroup, iItem)
}

// 列表视_删除全部.
func (l *ListView) I删除全部() int {
	return xc.XListView_DeleteAll(l.I句柄)
}

// 列表视_删除全部组.
func (l *ListView) I删除全部组() int {
	return xc.XListView_DeleteAllGroup(l.I句柄)
}

// 列表视_删除全部项.
func (l *ListView) I删除全部项() int {
	return xc.XListView_DeleteAllItem(l.I句柄)
}

// 列表视_组删除列.
//
// iColumn: 列索引.
func (l *ListView) I组删除列(iColumn int) int {
	return xc.XListView_DeleteColumnGroup(l.I句柄, iColumn)
}

// 列表视_项删除列.
//
// iColumn: 列索引.
func (l *ListView) I项删除列(iColumn int) int {
	return xc.XListView_DeleteColumnItem(l.I句柄, iColumn)
}

// 列表视_项获取文本.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (l *ListView) I项获取文本(iGroup int, iItem int, pName string) string {
	return xc.XListView_Item_GetTextEx(l.I句柄, iGroup, iItem, pName)
}

// 列表视_项获取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (l *ListView) I项获取图片扩展(iGroup int, iItem int, pName string) int {
	return xc.XListView_Item_GetImageEx(l.I句柄, iGroup, iItem, pName)
}

// 列表视_组取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (l *ListView) I组取文本(iGroup int, iColumn int) string {
	return xc.XListView_Group_GetText(l.I句柄, iGroup, iColumn)
}

// 列表视_组取文本扩展, 返回文本内容.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (l *ListView) I组取文本扩展(iGroup int, pName string) string {
	return xc.XListView_Group_GetTextEx(l.I句柄, iGroup, pName)
}

// 列表视_组取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (l *ListView) I组取图片(iGroup int, iColumn int) int {
	return xc.XListView_Group_GetImage(l.I句柄, iGroup, iColumn)
}

// 列表视_组取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (l *ListView) I组取图片扩展(iGroup int, pName string) int {
	return xc.XListView_Group_GetImageEx(l.I句柄, iGroup, pName)
}

// 列表视_项取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (l *ListView) I项取文本(iGroup int, iItem int, iColumn int) string {
	return xc.XListView_Item_GetText(l.I句柄, iGroup, iItem, iColumn)
}

// 列表视_项取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (l *ListView) I项取图片(iGroup int, iItem int, iColumn int) int {
	return xc.XListView_Item_GetImage(l.I句柄, iGroup, iItem, iColumn)
}

// 列表视_置拖动矩形颜色.
//
// color: ABGR 颜色.
//
// width: 线宽度.
func (l *ListView) I置拖动矩形颜色(color int, width int) int {
	return xc.XListView_SetDragRectColor(l.I句柄, color, width)
}

/*
以下都是事件
*/

type XE_LISTVIEW_TEMP_CREATE func(pItem *xc.ListView_Item_, nFlag int, pbHandled *bool) int                // 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_LISTVIEW_TEMP_CREATE1 func(hEle int, pItem *xc.ListView_Item_, nFlag int, pbHandled *bool) int     // 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_LISTVIEW_TEMP_CREATE_END func(pItem *xc.ListView_Item_, nFlag int, pbHandled *bool) int            // 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_LISTVIEW_TEMP_CREATE_END1 func(hEle int, pItem *xc.ListView_Item_, nFlag int, pbHandled *bool) int // 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_LISTVIEW_TEMP_DESTROY func(pItem *xc.ListView_Item_, nFlag int, pbHandled *bool) int               // 列表视元素-项模板销毁, 模板复用机制需先启用;.
type XE_LISTVIEW_TEMP_DESTROY1 func(hEle int, pItem *xc.ListView_Item_, nFlag int, pbHandled *bool) int    // 列表视元素-项模板销毁, 模板复用机制需先启用;.
type XE_LISTVIEW_TEMP_ADJUST_COORDINATE func(pItem *xc.ListView_Item_, pbHandled *bool) int                // 列表视元素,项模板调整坐标.已停用.
type XE_LISTVIEW_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.ListView_Item_, pbHandled *bool) int     // 列表视元素,项模板调整坐标.已停用.
type XE_LISTVIEW_DRAWITEM func(hDraw int, pItem *xc.ListView_Item_, pbHandled *bool) int                   // 列表视元素,自绘项.
type XE_LISTVIEW_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.ListView_Item_, pbHandled *bool) int        // 列表视元素,自绘项.
type XE_LISTVIEW_SELECT func(iGroup int, iItem int, pbHandled *bool) int                                   // 列表视元素,项选择事件.
type XE_LISTVIEW_SELECT1 func(hEle int, iGroup int, iItem int, pbHandled *bool) int                        // 列表视元素,项选择事件.
type XE_LISTVIEW_EXPAND func(iGroup int, bExpand bool, pbHandled *bool) int                                // 列表视元素,组展开收缩事件.
type XE_LISTVIEW_EXPAND1 func(hEle int, iGroup int, bExpand bool, pbHandled *bool) int                     // 列表视元素,组展开收缩事件.

// I事件_项模板创建 事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (l *ListView) I事件_项模板创建(pFun XE_LISTVIEW_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTVIEW_TEMP_CREATE, pFun)
}

// I事件_项模板创建  事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (l *ListView) I事件_项模板创建1(pFun XE_LISTVIEW_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTVIEW_TEMP_CREATE, pFun)
}

// I事件_项模板创建完成  事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (l *ListView) I事件_项模板创建完成(pFun XE_LISTVIEW_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTVIEW_TEMP_CREATE_END, pFun)
}

// I事件_项模板创建完成  事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (l *ListView) I事件_项模板创建完成1(pFun XE_LISTVIEW_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTVIEW_TEMP_CREATE_END, pFun)
}

// I事件_项模板销毁  , 模板复用机制需先启用;.
func (l *ListView) I事件_项模板销毁(pFun XE_LISTVIEW_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTVIEW_TEMP_DESTROY, pFun)
}

// I事件_项模板销毁  , 模板复用机制需先启用;.
func (l *ListView) I事件_项模板销毁1(pFun XE_LISTVIEW_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTVIEW_TEMP_DESTROY, pFun)
}

// I事件_项模板调整坐标.已停用.
func (l *ListView) Event_LISTVIEW_TEMP_ADJUST_COORDINATE(pFun XE_LISTVIEW_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTVIEW_TEMP_ADJUST_COORDINATE, pFun)
}

// I事件_项模板调整坐标.已停用.
func (l *ListView) Event_LISTVIEW_TEMP_ADJUST_COORDINATE1(pFun XE_LISTVIEW_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTVIEW_TEMP_ADJUST_COORDINATE, pFun)
}

// I事件_自绘项  .
func (l *ListView) I事件_自绘项(pFun XE_LISTVIEW_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTVIEW_DRAWITEM, pFun)
}

// I事件_自绘项  .
func (l *ListView) I事件_自绘项1(pFun XE_LISTVIEW_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTVIEW_DRAWITEM, pFun)
}

// I事件_项选择  事件.
func (l *ListView) I事件_项选择(pFun XE_LISTVIEW_SELECT) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTVIEW_SELECT, pFun)
}

// I事件_项选择  事件.
func (l *ListView) I事件_项选择1(pFun XE_LISTVIEW_SELECT1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTVIEW_SELECT, pFun)
}

// I事件_组展开收缩  事件.
func (l *ListView) I事件_组展开收缩(pFun XE_LISTVIEW_EXPAND) bool {
	return xc.XEle_RegEventC(l.I句柄, xcc.XE_LISTVIEW_EXPAND, pFun)
}

// I事件_组展开收缩  事件.
func (l *ListView) I事件_组展开收缩1(pFun XE_LISTVIEW_EXPAND1) bool {
	return xc.XEle_RegEventC1(l.I句柄, xcc.XE_LISTVIEW_EXPAND, pFun)
}
