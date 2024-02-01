package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 树形框元素.
type Tree struct {
	ScrollView
}

// I树形框_创建, 创建树元素.
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
func I树形框_创建(x int, y int, cx int, cy int, hParent int) *Tree {
	p := &Tree{}
	p.SetHandle(xc.XTree_Create(x, y, cx, cy, hParent))
	return p
}

// I列表树_从句柄创建对象.
func I树形框_从句柄创建(handle int) *Tree {
	p := &Tree{}
	p.SetHandle(handle)
	return p
}

// I列表树_从name创建对象, 失败返回nil.
func I树形框_从name创建(name string) *Tree {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Tree{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I列表树_从UID创建对象, 失败返回nil.
func I树形框_从UID创建(nUID int) *Tree {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Tree{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I列表树_从UID名称创建对象, 失败返回nil.
func I树形框_从UID名称创建(name string) *Tree {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Tree{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 列表树_启用拖动项, 启用拖动项功能.
//
// bEnable: 是否启用.
func (t *Tree) I启用拖动项(bEnable bool) int {
	return xc.XTree_EnableDragItem(t.I句柄, bEnable)
}

// 列表树_启用连接线, 启用或禁用显示项的连接线.
//
// bEnable: 是否启用.
//
// bSolid: 实线或虚线; TRUE: 实线, FALSE: 虚线.
func (t *Tree) I启用连接线(bEnable bool, bSolid bool) int {
	return xc.XTree_EnableConnectLine(t.I句柄, bEnable, bSolid)
}

// 列表树_启用展开, 启动或关闭默认展开功能, 如果开启新插入的项将自动展开.
//
// bEnable: 是否启用.
func (t *Tree) I启用展开(bEnable bool) int {
	return xc.XTree_EnableExpand(t.I句柄, bEnable)
}

// 列表树_启用模板复用.
//
// bEnable: 是否启用.
func (t *Tree) I启用模板复用(bEnable bool) int {
	return xc.XTree_EnablemTemplateReuse(t.I句柄, bEnable)
}

// 列表树_置连接线颜色.
//
// color: ABGR 颜色.
func (t *Tree) I置连接线颜色(color int) int {
	return xc.XTree_SetConnectLineColor(t.I句柄, color)
}

// 列表树_置展开按钮大小, 设置展开按钮占用空间大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (t *Tree) I置展开按钮大小(nWidth int, nHeight int) int {
	return xc.XTree_SetExpandButtonSize(t.I句柄, nWidth, nHeight)
}

// 列表树_置连接线长度, 设置连线绘制长度, 展开按钮与项内容之间的连线.
//
// nLength: 连线绘制长度.
func (t *Tree) I置连接线长度(nLength int) int {
	return xc.XTree_SetConnectLineLength(t.I句柄, nLength)
}

// 列表树_置拖动项插入位置颜色, 设置拖动项插入位置颜色提示.
//
// color: ABGR 颜色.
func (t *Tree) I置拖动项插入位置颜色(color int) int {
	return xc.XTree_SetDragInsertPositionColor(t.I句柄, color)
}

// 列表树_置项模板文件.
//
// pXmlFile: 文件名.
func (t *Tree) I置项模板文件(pXmlFile string) bool {
	return xc.XTree_SetItemTemplateXML(t.I句柄, pXmlFile)
}

// 列表树_置选择项模板文件, 设置项模板文件, 项选中状态.
//
// pXmlFile: 文件名.
func (t *Tree) I置选择项模板文件(pXmlFile string) bool {
	return xc.XTree_SetItemTemplateXMLSel(t.I句柄, pXmlFile)
}

// 列表树_置项模板.
//
// hTemp: 模板句柄.
func (t *Tree) I置项模板(hTemp int) bool {
	return xc.XTree_SetItemTemplate(t.I句柄, hTemp)
}

// 列表树_置选择项模板, 设置列表项模板, 项选中状态.
//
// hTemp: 模板句柄.
func (t *Tree) I置选择项模板(hTemp int) bool {
	return xc.XTree_SetItemTemplateSel(t.I句柄, hTemp)
}

// 列表树_置项模板从字符串, 设置项模板文件.
//
// pStringXML: 字符串.
func (t *Tree) I置项模板从字符串(pStringXML string) bool {
	return xc.XTree_SetItemTemplateXMLFromString(t.I句柄, pStringXML)
}

// 列表树_置选择项模板从字符串, 设置项模板文件, 项选中状态.
//
// pStringXML: 字符串.
func (t *Tree) I置选择项模板从字符串(pStringXML string) bool {
	return xc.XTree_SetItemTemplateXMLSelFromString(t.I句柄, pStringXML)
}

// 列表树_置项背景绘制标志, 设置是否绘制指定状态下项的背景.
//
// nFlags: 标志位: List_DrawItemBk_Flag_.
func (t *Tree) I置项背景绘制标志(nFlags xcc.List_DrawItemBk_Flag_) int {
	return xc.XTree_SetDrawItemBkFlags(t.I句柄, nFlags)
}

// 列表树_置项数据, 设置项用户数据.
//
// nID: 项ID.
//
// nUserData: 用户数据.
func (t *Tree) I置项数据(nID int, nUserData int) bool {
	return xc.XTree_SetItemData(t.I句柄, nID, nUserData)
}

// 列表树_取项数据, 获取项用户数据.
//
// nID: 项ID.
func (t *Tree) I取项数据(nID int) int {
	return xc.XTree_GetItemData(t.I句柄, nID)
}

// 列表树_置选择项.
//
// nID: 项ID.
func (t *Tree) I置选择项(nID int) bool {
	return xc.XTree_SetSelectItem(t.I句柄, nID)
}

// 列表树_取选择项, 返回项ID.
func (t *Tree) I取选择项() int {
	return xc.XTree_GetSelectItem(t.I句柄)
}

// 列表树_可视指定项, 滚动视图让指定项可见.
//
// nID: 项索引.
func (t *Tree) I可视指定项(nID int) int {
	return xc.XTree_VisibleItem(t.I句柄, nID)
}

// 列表树_判断展开.
//
// nID: 项ID.
func (t *Tree) I判断展开(nID int) bool {
	return xc.XTree_IsExpand(t.I句柄, nID)
}

// 列表树_展开项, 判断项是否展开.
//
// nID: 项ID.
//
// bExpand: 是否展开.
func (t *Tree) I展开项(nID int, bExpand bool) bool {
	return xc.XTree_ExpandItem(t.I句柄, nID, bExpand)
}

// 列表树_展开全部子项, 展开所有的子项.
//
// nID: 项ID.
//
// bExpand: 是否展开.
func (t *Tree) I展开全部子项(nID int, bExpand bool) bool {
	return xc.XTree_ExpandAllChildItem(t.I句柄, nID, bExpand)
}

// 列表树_测试点击项, 检测坐标点所在项, 返回项ID.
//
// pPt: 坐标点.
func (t *Tree) I测试点击项(pPt *xc.POINT) int {
	return xc.XTree_HitTest(t.I句柄, pPt)
}

// 列表树_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移坐标, 返回项ID.
//
// pPt: 坐标点.
func (t *Tree) I测试点击项扩展(pPt *xc.POINT) int {
	return xc.XTree_HitTestOffset(t.I句柄, pPt)
}

// 列表树_取第一个子项, 获取第一个子项. 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) I取第一个子项(nID int) int {
	return xc.XTree_GetFirstChildItem(t.I句柄, nID)
}

// 列表树_取末尾子项, 获取末尾子项. 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) I取末尾子项(nID int) int {
	return xc.XTree_GetEndChildItem(t.I句柄, nID)
}

// 列表树_取上一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) I取上一个兄弟项(nID int) int {
	return xc.XTree_GetPrevSiblingItem(t.I句柄, nID)
}

// 列表树_取下一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) I取下一个兄弟项(nID int) int {
	return xc.XTree_GetNextSiblingItem(t.I句柄, nID)
}

// 列表树_取父项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) I取父项(nID int) int {
	return xc.XTree_GetParentItem(t.I句柄, nID)
}

// 列表树_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (t *Tree) I创建数据适配器() int {
	return xc.XTree_CreateAdapter(t.I句柄)
}

// 列表树_绑定数据适配器.
//
// hAdapter: 数据适配器句柄, XAdTree.
func (t *Tree) I绑定数据适配器(hAdapter int) int {
	return xc.XTree_BindAdapter(t.I句柄, hAdapter)
}

// 列表树_取数据适配器, 返回数据适配器句柄.
func (t *Tree) I取数据适配器() int {
	return xc.XTree_GetAdapter(t.I句柄)
}

// 列表树_刷新数据, 刷新所有项模板, 以便更新UI.
func (t *Tree) I刷新数据() int {
	return xc.XTree_RefreshData(t.I句柄)
}

// 列表树_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// nID: 项ID.
func (t *Tree) I刷新指定项(nID int) int {
	return xc.XTree_RefreshItem(t.I句柄, nID)
}

// 列表树_置缩进, 设置缩进大小.
//
// nWidth: 缩进宽度.
func (t *Tree) I置缩进(nWidth int) int {
	return xc.XTree_SetIndentation(t.I句柄, nWidth)
}

// 列表树_取缩进, 返回缩进值大小.
func (t *Tree) I取缩进() int {
	return xc.XTree_GetIndentation(t.I句柄)
}

// 列表树_置项默认高度.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (t *Tree) I置项默认高度(nHeight int, nSelHeight int) int {
	return xc.XTree_SetItemHeightDefault(t.I句柄, nHeight, nSelHeight)
}

// 列表树_取项默认高度.
//
// pHeight: 接收返回高度.
//
// pSelHeight: 接收返回值, 当项选中时的高度.
func (t *Tree) I取项默认高度(pHeight *int, pSelHeight *int) int {
	return xc.XTree_GetItemHeightDefault(t.I句柄, pHeight, pSelHeight)
}

// 列表树_置项高度.
//
// nID: 项ID.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (t *Tree) I置项高度(nID int, nHeight int, nSelHeight int) int {
	return xc.XTree_SetItemHeight(t.I句柄, nID, nHeight, nSelHeight)
}

// 列表树_取项高度.
//
// nID: 项ID.
//
// pHeight: 接收返回高度.
//
// pSelHeight: 接收返回值, 当项选中时的高度.
func (t *Tree) I取项高度(nID int, pHeight *int, pSelHeight *int) int {
	return xc.XTree_GetItemHeight(t.I句柄, nID, pHeight, pSelHeight)
}

// 列表树_置行间距.
//
// nSpace: 行间隔大小.
func (t *Tree) I置行间距(nSpace int) int {
	return xc.XTree_SetRowSpace(t.I句柄, nSpace)
}

// 列表树_取行间距.
func (t *Tree) I取行间距() int {
	return xc.XTree_GetRowSpace(t.I句柄)
}

// 列表树_移动项, 移动项的位置.
//
// nMoveItem: 要移动的项ID.
//
// nDestItem: 目标项ID, 参照位置.
//
// nFlag: 0:目标前面, 1:目标后面, 2:目标子项首, 3:目标子项尾.
func (t *Tree) I移动项(nMoveItem int, nDestItem int, nFlag int) bool {
	return xc.XTree_MoveItem(t.I句柄, nMoveItem, nDestItem, nFlag)
}

// 列表树_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// nID: 树项ID.
//
// nTempItemID: 模板项ID.
func (t *Tree) I取模板对象(nID int, nTempItemID int) int {
	return xc.XTree_GetTemplateObject(t.I句柄, nID, nTempItemID)
}

// 列表树_取对象所在项, 获取当前对象所在模板实例, 属于列表树中哪一个项. 成功返回项ID, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄.
func (t *Tree) I取对象所在项(hXCGUI int) int {
	return xc.XTree_GetItemIDFromHXCGUI(t.I句柄, hXCGUI)
}

// 列表树_插入项文本.
//
// pValue:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) I插入项文本(pValue string, nParentID int, insertID int) int {
	return xc.XTree_InsertItemText(t.I句柄, pValue, nParentID, insertID)
}

// 列表树_插入项文本扩展.
//
// pName:.
//
// pValue:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) I插入项文本扩展(pName string, pValue string, nParentID int, insertID int) int {
	return xc.XTree_InsertItemTextEx(t.I句柄, pName, pValue, nParentID, insertID)
}

// 列表树_插入项图片.
//
// hImage:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) I插入项图片(hImage int, nParentID int, insertID int) int {
	return xc.XTree_InsertItemImage(t.I句柄, hImage, nParentID, insertID)
}

// 列表树_插入项图片扩展.
//
// pName:.
//
// hImage:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) I插入项图片扩展(pName string, hImage int, nParentID int, insertID int) int {
	return xc.XTree_InsertItemImageEx(t.I句柄, pName, hImage, nParentID, insertID)
}

// 列表树_取项数量.
func (t *Tree) I取项数量() int {
	return xc.XTree_GetCount(t.I句柄)
}

// 列表树_取列数量.
func (t *Tree) I取列数量() int {
	return xc.XTree_GetCountColumn(t.I句柄)
}

// 列表树_置项文本.
//
// nID:.
//
// iColumn:.
//
// pValue:.
func (t *Tree) I置项文本(nID int, iColumn int, pValue string) bool {
	return xc.XTree_SetItemText(t.I句柄, nID, iColumn, pValue)
}

// 列表树_置项文本扩展.
//
// nID:.
//
// pName:.
//
// pValue:.
func (t *Tree) I置项文本扩展(nID int, pName string, pValue string) bool {
	return xc.XTree_SetItemTextEx(t.I句柄, nID, pName, pValue)
}

// 列表树_置项图片.
//
// nID:.
//
// iColumn:.
//
// hImage:.
func (t *Tree) I置项图片(nID int, iColumn int, hImage int) bool {
	return xc.XTree_SetItemImage(t.I句柄, nID, iColumn, hImage)
}

// 列表树_置项图片扩展.
//
// nID:.
//
// pName:.
//
// hImage:.
func (t *Tree) I置项图片扩展(nID int, pName string, hImage int) bool {
	return xc.XTree_SetItemImageEx(t.I句柄, nID, pName, hImage)
}

// 列表树_取项文本.
//
// nID:.
//
// iColumn:.
func (t *Tree) I取项文本(nID int, iColumn int) string {
	return xc.XTree_GetItemText(t.I句柄, nID, iColumn)
}

// 列表树_取项文本扩展.
//
// nID:.
//
// pName:.
func (t *Tree) I取项文本扩展(nID int, pName string) string {
	return xc.XTree_GetItemTextEx(t.I句柄, nID, pName)
}

// 列表树_取项图片.
//
// nID:.
//
// iColumn:.
func (t *Tree) I取项图片(nID int, iColumn int) int {
	return xc.XTree_GetItemImage(t.I句柄, nID, iColumn)
}

// 列表树_取项图片扩展.
//
// nID:.
//
// pName:.
func (t *Tree) I取项图片扩展(nID int, pName string) int {
	return xc.XTree_GetItemImageEx(t.I句柄, nID, pName)
}

// 列表树_删除项.
//
// nID:.
func (t *Tree) I删除项(nID int) bool {
	return xc.XTree_DeleteItem(t.I句柄, nID)
}

// 列表树_删除全部项.
func (t *Tree) I删除全部项() int {
	return xc.XTree_DeleteItemAll(t.I句柄)
}

// 列表树_删除列全部.
func (t *Tree) I删除列全部() int {
	return xc.XTree_DeleteColumnAll(t.I句柄)
}

// 列表树_置分割线颜色.
//
// color: ABGR 颜色值.
func (t *Tree) I置分割线颜色(color int) int {
	return xc.XTree_SetSplitLineColor(t.I句柄, color)
}

/*
以下都是事件
*/

type XE_TREE_TEMP_CREATE func(pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int                // 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_TREE_TEMP_CREATE1 func(hEle int, pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int     // 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
type XE_TREE_TEMP_CREATE_END func(pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int            // 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_TREE_TEMP_CREATE_END1 func(hEle int, pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int // 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
type XE_TREE_TEMP_DESTROY func(pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int               // 列表树元素-项模板销毁,模板复用机制需先启用;.
type XE_TREE_TEMP_DESTROY1 func(hEle int, pItem *xc.Tree_Item_, nFlag int, pbHandled *bool) int    // 列表树元素-项模板销毁,模板复用机制需先启用;.
type XE_TREE_TEMP_ADJUST_COORDINATE func(pItem *xc.Tree_Item_, pbHandled *bool) int                // 树元素,项模板,调整项坐标. 已停用.
type XE_TREE_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.Tree_Item_, pbHandled *bool) int     // 树元素,项模板,调整项坐标. 已停用.
type XE_TREE_DRAWITEM func(hDraw int, pItem *xc.Tree_Item_, pbHandled *bool) int                   // 树元素,绘制项.
type XE_TREE_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.Tree_Item_, pbHandled *bool) int        // 树元素,绘制项.
type XE_TREE_SELECT func(nItemID int, pbHandled *bool) int                                         // 树元素,项选择事件.
type XE_TREE_SELECT1 func(hEle int, nItemID int, pbHandled *bool) int                              // 树元素,项选择事件.
type XE_TREE_EXPAND func(id int, bExpand bool, pbHandled *bool) int                                // 树元素,项展开收缩事件.
type XE_TREE_EXPAND1 func(hEle int, id int, bExpand bool, pbHandled *bool) int                     // 树元素,项展开收缩事件.
type XE_TREE_DRAG_ITEM_ING func(pInfo *xc.Tree_Drag_Item_, pbHandled *bool) int                    // 树元素,用户正在拖动项, 可对参数值修改.
type XE_TREE_DRAG_ITEM_ING1 func(hEle int, pInfo *xc.Tree_Drag_Item_, pbHandled *bool) int         // 树元素,用户正在拖动项, 可对参数值修改.
type XE_TREE_DRAG_ITEM func(pInfo *xc.Tree_Drag_Item_, pbHandled *bool) int                        // 树元素,拖动项事件.
type XE_TREE_DRAG_ITEM1 func(hEle int, pInfo *xc.Tree_Drag_Item_, pbHandled *bool) int             // 树元素,拖动项事件.

// 列表树元素- I事件_项模板创建 ,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (t *Tree) I事件_项模板创建(pFun XE_TREE_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_TEMP_CREATE, pFun)
}

// 列表树元素-I事件_项模板创建 ,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
func (t *Tree) I事件_项模板创建1(pFun XE_TREE_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_TEMP_CREATE, pFun)
}

// 列表树元素-I事件_项模板创建完成 ,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (t *Tree) I事件_项模板创建完成(pFun XE_TREE_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_TEMP_CREATE_END, pFun)
}

// 列表树元素-I事件_项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
func (t *Tree) I事件_项模板创建完成1(pFun XE_TREE_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_TEMP_CREATE_END, pFun)
}

// 列表树元素-I事件_项模板销毁,模板复用机制需先启用;.
func (t *Tree) I事件_项模板销毁(pFun XE_TREE_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_TEMP_DESTROY, pFun)
}

// 列表树元素-I事件_项模板销毁,模板复用机制需先启用;.
func (t *Tree) I事件_项模板销毁1(pFun XE_TREE_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_TEMP_DESTROY, pFun)
}

// 树元素,项模板,I事件_调整项坐标. 已停用.
func (t *Tree) I事件_调整项坐标(pFun XE_TREE_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_TEMP_ADJUST_COORDINATE, pFun)
}

// 树元素,项模板,I事件_调整项坐标. 已停用.
func (t *Tree) I事件_调整项坐标1(pFun XE_TREE_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_TEMP_ADJUST_COORDINATE, pFun)
}

// 树元素,I事件_绘制项.
func (t *Tree) I事件_绘制项(pFun XE_TREE_DRAWITEM) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_DRAWITEM, pFun)
}

// 树元素,I事件_绘制项.
func (t *Tree) I事件_绘制项1(pFun XE_TREE_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_DRAWITEM, pFun)
}

// 树元素,I事件_项选择 事件.
func (t *Tree) I事件_项选择(pFun XE_TREE_SELECT) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_SELECT, pFun)
}

// 树元素,I事件_项选择 事件.
func (t *Tree) I事件_项选择1(pFun XE_TREE_SELECT1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_SELECT, pFun)
}

// 树元素,项展开收缩事件.
func (t *Tree) I事件_项展开收缩(pFun XE_TREE_EXPAND) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_EXPAND, pFun)
}

// 树元素,I事件_项展开收缩 事件.
func (t *Tree) I事件_项展开收缩1(pFun XE_TREE_EXPAND1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_EXPAND, pFun)
}

// 树元素,I事件_用户正在拖动项, 可对参数值修改.
func (t *Tree) I事件_用户正在拖动项(pFun XE_TREE_DRAG_ITEM_ING) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_DRAG_ITEM_ING, pFun)
}

// 树元素,I事件_用户正在拖动项, 可对参数值修改.
func (t *Tree) I事件_用户正在拖动项1(pFun XE_TREE_DRAG_ITEM_ING1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_DRAG_ITEM_ING, pFun)
}

// 树元素,I事件_拖动项 事件.
func (t *Tree) I事件_拖动项(pFun XE_TREE_DRAG_ITEM) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TREE_DRAG_ITEM, pFun)
}

// 树元素,I事件_拖动项 事件.
func (t *Tree) I事件_拖动项1(pFun XE_TREE_DRAG_ITEM1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TREE_DRAG_ITEM, pFun)
}
