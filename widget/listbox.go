package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 列表框.
type ListBox struct {
	ScrollView
}

// 列表框_创建, 创建列表框元素.
//.
//.
//.
//.
//. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
// ff:创建列表框
// x:元素x坐标
// y:元素y坐标
// cx:宽度
// cy:高度
// hParent:父窗口句柄或元素句柄
func NewListBox(x int, y int, cx int, cy int, hParent int) *ListBox {
	p := &ListBox{}
	p.SetHandle(xc.XListBox_Create(x, y, cx, cy, hParent))
	return p
}

// 列表框_创建Ex, 创建列表框元素, 使用内置项模板, 自动创建数据适配器.
//.
//.
//.
//.
//. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
//. 例如: 内置模板是1列, 如果数据有5列, 那么此参数填5.
// ff:创建列表框Ex
// x:元素x坐标
// y:元素y坐标
// cx:宽度
// cy:高度
// hParent:父窗口句柄或元素句柄
// col_extend_count:列数量
func NewListBoxEx(x, y, cx, cy int32, hParent, col_extend_count int32) *ListBox {
	p := &ListBox{}
	p.SetHandle(xc.XListBox_CreateEx(x, y, cx, cy, hParent, col_extend_count))
	return p
}

// 从句柄创建对象.
// ff:创建列表框并按句柄
// handle:
func NewListBoxByHandle(handle int) *ListBox {
	p := &ListBox{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
// ff:创建列表框并按名称
// name:
func NewListBoxByName(name string) *ListBox {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
// ff:创建列表框并按UID
// nUID:
func NewListBoxByUID(nUID int) *ListBox {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
// ff:创建列表框并按UID名称
// name:
func NewListBoxByUIDName(name string) *ListBox {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ListBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 列表框_启用固定行高.
//.
// ff:启用固定行高
// bEnable:是否启用
func (l *ListBox) EnableFixedRowHeight(bEnable bool) int {
	return xc.XListBox_EnableFixedRowHeight(l.Handle, bEnable)
}

// 列表框_启用模板复用.
//.
// ff:启用模板复用
// bEnable:是否启用
func (l *ListBox) EnableTemplateReuse(bEnable bool) int {
	return xc.XListBox_EnableTemplateReuse(l.Handle, bEnable)
}

// 列表框_启用虚表.
//.
// ff:启用虚表
// bEnable:是否启用
func (l *ListBox) EnableVirtualTable(bEnable bool) int {
	return xc.XListBox_EnableVirtualTable(l.Handle, bEnable)
}

// 列表框_置虚表行数.
//.
// ff:置虚表行数
// nRowCount:行数
func (l *ListBox) SetVirtualRowCount(nRowCount int) int {
	return xc.XListBox_SetVirtualRowCount(l.Handle, nRowCount)
}

// 列表框_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//, List_DrawItemBk_Flag_.
// ff:置绘制项背景标志
// nFlags:标志位
func (l *ListBox) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) int {
	return xc.XListBox_SetDrawItemBkFlags(l.Handle, nFlags)
}

// 列表框_置项数据, 设置项用户数据.
//.
//.
// ff:置项数据
// iItem:想索引
// nUserData:用户数据
func (l *ListBox) SetItemData(iItem int, nUserData int) bool {
	return xc.XListBox_SetItemData(l.Handle, iItem, nUserData)
}

// 列表框_取项数据, 获取项用户数据.
//.
// ff:取项数据
// iItem:项索引
func (l *ListBox) GetItemData(iItem int) int {
	return xc.XListBox_GetItemData(l.Handle, iItem)
}

// 列表框_置项信息.
//.
//.
// ff:置项信息
// iItem:项索引
// pItem:项信息
func (l *ListBox) SetItemInfo(iItem int, pItem *xc.ListBox_Item_Info_) bool {
	return xc.XListBox_SetItemInfo(l.Handle, iItem, pItem)
}

// 列表框_取项背景信息, 获取项信息.
//.
//.
// ff:取项背景信息
// iItem:项索引
// pItem:项信息
func (l *ListBox) GetItemInfo(iItem int, pItem *xc.ListBox_Item_Info_) bool {
	return xc.XListBox_GetItemInfo(l.Handle, iItem, pItem)
}

// 列表框_置选择项, 设置选择选.
//.
// ff:置选择项
// iItem:项索引
func (l *ListBox) SetSelectItem(iItem int) bool {
	return xc.XListBox_SetSelectItem(l.Handle, iItem)
}

// 列表框_取选择项, 返回项索引.
// ff:取选择项
func (l *ListBox) GetSelectItem() int {
	return xc.XListBox_GetSelectItem(l.Handle)
}

// 列表框_添加选择项.
//.
// ff:添加选择项
// iItem:项索引
func (l *ListBox) AddSelectItem(iItem int) bool {
	return xc.XListBox_AddSelectItem(l.Handle, iItem)
}

// 列表框_取消选择项.
//.
// ff:取消选择项
// iItem:项索引
func (l *ListBox) CancelSelectItem(iItem int) bool {
	return xc.XListBox_CancelSelectItem(l.Handle, iItem)
}

// 列表框_取消选择全部, 如果之前有选择状态的项返回TRUE, 此时可以更新UI, 否则返回FALSE.
// ff:取消选择全部
func (l *ListBox) CancelSelectAll() bool {
	return xc.XListBox_CancelSelectAll(l.Handle)
}

// 列表框_取全部选择, 获取所有选择项, 返回接收数量.
//.
//.
// ff:取全部选择
// pArray:切片缓冲区
// nArraySize:切片大小
func (l *ListBox) GetSelectAll(pArray *[]int32, nArraySize int) int {
	return xc.XListBox_GetSelectAll(l.Handle, pArray, nArraySize)
}

// 列表框_取选择项数量, 获取选择项数量.
// ff:取选择项数量
func (l *ListBox) GetSelectCount() int {
	return xc.XListBox_GetSelectCount(l.Handle)
}

// 列表框_取鼠标停留项, 返回鼠标所在项.
// ff:取鼠标停留项
func (l *ListBox) GetItemMouseStay() int {
	return xc.XListBox_GetItemMouseStay(l.Handle)
}

// 列表框_选择全部项.
// ff:选择全部项
func (l *ListBox) SelectAll() bool {
	return xc.XListBox_SelectAll(l.Handle)
}

// 列表框_显示指定项, 滚动视图让指定项可见.
//.
// ff:显示指定项
// iItem:项索引
func (l *ListBox) VisibleItem(iItem int) int {
	return xc.XListBox_VisibleItem(l.Handle, iItem)
}

// 列表框_取可视行范围, 获取当前可见行范围.
//.
//.
// ff:取可视行范围
// piStart:开始行索引
// piEnd:结束行索引
func (l *ListBox) GetVisibleRowRange(piStart *int32, piEnd *int32) int {
	return xc.XListBox_GetVisibleRowRange(l.Handle, piStart, piEnd)
}

// 列表框_置项默认高度.
//.
//.
// ff:置项默认高度
// nHeight:项高度
// nSelHeight:选中项高度
func (l *ListBox) SetItemHeightDefault(nHeight int32, nSelHeight int32) int {
	return xc.XListBox_SetItemHeightDefault(l.Handle, nHeight, nSelHeight)
}

// 列表框_取项默认高度.
//.
//.
// ff:取项默认高度
// pHeight:高度
// pSelHeight:选中时高度
func (l *ListBox) GetItemHeightDefault(pHeight *int32, pSelHeight *int32) int {
	return xc.XListBox_GetItemHeightDefault(l.Handle, pHeight, pSelHeight)
}

// 列表框_取所在行索引, 获取当前对象所在模板实例, 属于列表中哪一个项(行). 成功返回项索引, 否则返回XC_ID_ERROR.
//, UI元素句柄或形状对象句柄.
// ff:取所在行索引
// hXCGUI:对象句柄
func (l *ListBox) GetItemIndexFromHXCGUI(hXCGUI int) int {
	return xc.XListBox_GetItemIndexFromHXCGUI(l.Handle, hXCGUI)
}

// 列表框_置行间距.
//.
// ff:置行间距
// nSpace:间距大小
func (l *ListBox) SetRowSpace(nSpace int) int {
	return xc.XListBox_SetRowSpace(l.Handle, nSpace)
}

// 列表框_取行间距.
// ff:取行间距
func (l *ListBox) GetRowSpace() int {
	return xc.XListBox_GetRowSpace(l.Handle)
}

// 列表框_测试点击项, 检测坐标点所在项, 返回项索引.
//.
// ff:测试点击项
// pPt:坐标点
func (l *ListBox) HitTest(pPt *xc.POINT) int {
	return xc.XListBox_HitTest(l.Handle, pPt)
}

// 列表框_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移量, 返回项索引.
//.
// ff:测试点击项EX
// pPt:坐标点
func (l *ListBox) HitTestOffset(pPt *xc.POINT) int {
	return xc.XListBox_HitTestOffset(l.Handle, pPt)
}

// 列表框_置项模板文件, 设置列表项模板文件.
//.
// ff:置项模板文件
// pXmlFile:文件名
func (l *ListBox) SetItemTemplateXML(pXmlFile string) bool {
	return xc.XListBox_SetItemTemplateXML(l.Handle, pXmlFile)
}

// 列表框_置项模板, 设置列表项模板.
//.
// ff:置项模板
// hTemp:模板句柄
func (l *ListBox) SetItemTemplate(hTemp int) bool {
	return xc.XListBox_SetItemTemplate(l.Handle, hTemp)
}

// 列表框_置项模板从字符串, 设置项布局模板文件.
//.
// ff:置项模板并按字符串
// pStringXML:字符串
func (l *ListBox) SetItemTemplateXMLFromString(pStringXML string) bool {
	return xc.XListBox_SetItemTemplateXMLFromString(l.Handle, pStringXML)
}

// 列表框_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄, 成功返回对象句柄, 否则返回NULL.
//.
//.
// ff:取模板对象
// iItem:项索引
// nTempItemID:模板项ID
func (l *ListBox) GetTemplateObject(iItem int, nTempItemID int) int {
	return xc.XListBox_GetTemplateObject(l.Handle, iItem, nTempItemID)
}

// 列表框_启用多选, 是否启用多行选择功能.
//.
// ff:启用多选
// bEnable:是否启用
func (l *ListBox) EnableMultiSel(bEnable bool) int {
	return xc.XListBox_EnableMultiSel(l.Handle, bEnable)
}

// 列表框_创建数据适配器, 创建数据适配器并绑定, 根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
// ff:创建数据适配器
func (l *ListBox) CreateAdapter() int {
	return xc.XListBox_CreateAdapter(l.Handle)
}

// 列表框_绑定数据适配器, 绑定数据适配器.
//.
// ff:绑定数据适配器
// hAdapter:数据适配器句柄
func (l *ListBox) BindAdapter(hAdapter int) int {
	return xc.XListBox_BindAdapter(l.Handle, hAdapter)
}

// 列表框_取数据适配器, 获取绑定的数据适配器, 返回数据适配器句柄.
// ff:取数据适配器
func (l *ListBox) GetAdapter() int {
	return xc.XListBox_GetAdapter(l.Handle)
}

// 列表框_排序.
//.
//(TRUE)或降序(FALSE).
// ff:排序
// iColumnAdapter:列索引
// bAscending:升序
func (l *ListBox) Sort(iColumnAdapter int, bAscending bool) int {
	return xc.XListBox_Sort(l.Handle, iColumnAdapter, bAscending)
}

// 列表框_刷新数据.
// ff:刷新数据
func (l *ListBox) RefreshData() int {
	return xc.XListBox_RefreshData(l.Handle)
}

// 列表框_刷新指定项, 刷新指定项模板, 以便更新UI.
//.
// ff:刷新指定项
// iItem:项索引
func (l *ListBox) RefreshItem(iItem int) int {
	return xc.XListBox_RefreshItem(l.Handle, iItem)
}

// 列表框_添加项文本, XAdTable_AddItemText, 返回项索引.
//.
// ff:添加项文本
// pText:文本
func (l *ListBox) AddItemText(pText string) int {
	return xc.XListBox_AddItemText(l.Handle, pText)
}

// 列表框_添加项文本扩展, XAdTable_AddItemTextEx.
//.
//.
// ff:添加项文本EX
// pName:名称
// pText:文本
func (l *ListBox) AddItemTextEx(pName string, pText string) int {
	return xc.XListBox_AddItemTextEx(l.Handle, pName, pText)
}

// 列表框_添加项图片, XAdTable_AddItemImage.
//.
// ff:添加项图片
// hImage:图片
func (l *ListBox) AddItemImage(hImage int) int {
	return xc.XListBox_AddItemImage(l.Handle, hImage)
}

// 列表框_添加项图片扩展, XAdTable_AddItemImageEx.
//.
//.
// ff:添加项图片EX
// pName:名称
// hImage:图片
func (l *ListBox) AddItemImageEx(pName string, hImage int) int {
	return xc.XListBox_AddItemImageEx(l.Handle, pName, hImage)
}

// 列表框_插入项文本.
//.
//.
// ff:插入项文本
// iItem:项
// pValue:文本
func (l *ListBox) InsertItemText(iItem int, pValue string) int {
	return xc.XListBox_InsertItemText(l.Handle, iItem, pValue)
}

// 列表框_插入项文本扩展.
//.
//.
//.
// ff:插入项文本EX
// iItem:项
// pName:名称
// pValue:文本
func (l *ListBox) InsertItemTextEx(iItem int, pName string, pValue string) int {
	return xc.XListBox_InsertItemTextEx(l.Handle, iItem, pName, pValue)
}

// 列表框_插入项图片.
//.
//.
// ff:插入项图片
// iItem:项
// hImage:图片
func (l *ListBox) InsertItemImage(iItem int, hImage int) int {
	return xc.XListBox_InsertItemImage(l.Handle, iItem, hImage)
}

// 列表框_插入项图片扩展.
//.
//.
//.
// ff:插入项图片EX
// iItem:项
// pName:名称
// hImage:图片
func (l *ListBox) InsertItemImageEx(iItem int, pName string, hImage int) int {
	return xc.XListBox_InsertItemImageEx(l.Handle, iItem, pName, hImage)
}

// 列表框_置项文本.
//.
//.
//.
// ff:置项文本
// iItem:项
// iColumn:列
// pText:文本
func (l *ListBox) SetItemText(iItem int, iColumn int, pText string) bool {
	return xc.XListBox_SetItemText(l.Handle, iItem, iColumn, pText)
}

// 列表框_置项文本扩展.
//.
//.
//.
// ff:置项文本EX
// iItem:项
// pName:名称
// pText:文本
func (l *ListBox) SetItemTextEx(iItem int, pName string, pText string) bool {
	return xc.XListBox_SetItemTextEx(l.Handle, iItem, pName, pText)
}

// 列表框_置项图片.
//.
//.
//.
// ff:置项图片
// iItem:项
// iColumn:列
// hImage:图片
func (l *ListBox) SetItemImage(iItem int, iColumn int, hImage int) bool {
	return xc.XListBox_SetItemImage(l.Handle, iItem, iColumn, hImage)
}

// 列表框_置项图片扩展.
//.
//.
//.
// ff:置项图片EX
// iItem:项
// pName:名称
// hImage:图片
func (l *ListBox) SetItemImageEx(iItem int, pName string, hImage int) bool {
	return xc.XListBox_SetItemImageEx(l.Handle, iItem, pName, hImage)
}

// 列表框_置项整数值.
//.
//.
//.
// ff:置项整数值
// iItem:项
// iColumn:列
// nValue:文本
func (l *ListBox) SetItemInt(iItem int, iColumn int, nValue int) bool {
	return xc.XListBox_SetItemInt(l.Handle, iItem, iColumn, nValue)
}

// 列表框_置项整数值扩展.
//.
//.
//.
// ff:置项整数值EX
// iItem:项
// pName:名称
// nValue:文本
func (l *ListBox) SetItemIntEx(iItem int, pName string, nValue int) bool {
	return xc.XListBox_SetItemIntEx(l.Handle, iItem, pName, nValue)
}

// 列表框_置项浮点值.
//.
//.
//.
// ff:置项浮点值
// iItem:项
// iColumn:列
// fFloat:小数值
func (l *ListBox) SetItemFloat(iItem int, iColumn int, fFloat float32) bool {
	return xc.XListBox_SetItemFloat(l.Handle, iItem, iColumn, fFloat)
}

// 列表框_置项浮点值扩展.
//.
//.
//.
// ff:置项浮点值EX
// iItem:项
// pName:名称
// fFloat:小数值
func (l *ListBox) SetItemFloatEx(iItem int, pName string, fFloat float32) bool {
	return xc.XListBox_SetItemFloatEx(l.Handle, iItem, pName, fFloat)
}

// 列表框_取项文本.
//.
//.
// ff:取项文本
// iItem:项
// iColumn:列
func (l *ListBox) GetItemText(iItem int, iColumn int) string {
	return xc.XListBox_GetItemText(l.Handle, iItem, iColumn)
}

// 列表框_取项文本扩展.
//.
//.
// ff:取项文本EX
// iItem:项
// pName:名称
func (l *ListBox) GetItemTextEx(iItem int, pName string) string {
	return xc.XListBox_GetItemTextEx(l.Handle, iItem, pName)
}

// 列表框_取项图片.
//.
//.
// ff:取项图片
// iItem:项
// iColumn:列
func (l *ListBox) GetItemImage(iItem int, iColumn int) int {
	return xc.XListBox_GetItemImage(l.Handle, iItem, iColumn)
}

// 列表框_取项图片扩展.
//.
//.
// ff:取项图片EX
// iItem:项
// pName:名称
func (l *ListBox) GetItemImageEx(iItem int, pName string) int {
	return xc.XListBox_GetItemImageEx(l.Handle, iItem, pName)
}

// 列表框_取项整数值.
//.
//.
//.
// ff:取项整数值
// iItem:项
// iColumn:列
// pOutValue:值指针
func (l *ListBox) GetItemInt(iItem int, iColumn int, pOutValue *int32) bool {
	return xc.XListBox_GetItemInt(l.Handle, iItem, iColumn, pOutValue)
}

// 列表框_取项整数值扩展.
//.
//.
//.
// ff:取项整数值EX
// iItem:项
// pName:名称
// pOutValue:值指针
func (l *ListBox) GetItemIntEx(iItem int, pName string, pOutValue *int32) bool {
	return xc.XListBox_GetItemIntEx(l.Handle, iItem, pName, pOutValue)
}

// 列表框_取项浮点值.
//.
//.
//.
// ff:取项浮点值
// iItem:项
// iColumn:列
// pOutValue:值指针
func (l *ListBox) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XListBox_GetItemFloat(l.Handle, iItem, iColumn, pOutValue)
}

// 列表框_取项浮点值扩展.
//.
//.
//.
// ff:取项浮点值EX
// iItem:项
// pName:名称
// pOutValue:值指针
func (l *ListBox) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {
	return xc.XListBox_GetItemFloatEx(l.Handle, iItem, pName, pOutValue)
}

// 列表框_删除项.
//.
// ff:删除项
// iItem:项
func (l *ListBox) DeleteItem(iItem int) bool {
	return xc.XListBox_DeleteItem(l.Handle, iItem)
}

// 列表框_删除项扩展.
//.
//.
// ff:删除项EX
// iItem:项
// nCount:计数
func (l *ListBox) DeleteItemEx(iItem int, nCount int) bool {
	return xc.XListBox_DeleteItemEx(l.Handle, iItem, nCount)
}

// 列表框_删除项全部.
// ff:删除项全部
func (l *ListBox) DeleteItemAll() int {
	return xc.XListBox_DeleteItemAll(l.Handle)
}

// 列表框_删除列全部.
// ff:删除列全部
func (l *ListBox) DeleteColumnAll() int {
	return xc.XListBox_DeleteColumnAll(l.Handle)
}

// 列表框_取项数量AD.
// ff:取项数量AD
func (l *ListBox) GetCount_AD() int {
	return xc.XListBox_GetCount_AD(l.Handle)
}

// 列表框_取列数量AD.
// ff:取列数量AD
func (l *ListBox) GetCountColumn_AD() int {
	return xc.XListBox_GetCountColumn_AD(l.Handle)
}

// 列表框_置分割线颜色.
//.
// ff:置分割线颜色
// color:ABGR颜色值
func (l *ListBox) SetSplitLineColor(color int) int {
	return xc.XListBox_SetSplitLineColor(l.Handle, color)
}

// 列表框_置拖动矩形颜色.
//.
//.
// ff:置拖动矩形颜色
// color:ABGR颜色值
// width:线宽度
func (l *ListBox) SetDragRectColor(color, width int) int {
	return xc.XListBox_SetDragRectColor(l.Handle, color, width)
}

// 列表框_置项模板从内存. 设置项模板文件.
//.
// ff:置项模板从内存
// data:模板数据
func (l *ListBox) SetItemTemplateXMLFromMem(data []byte) bool {
	return xc.XListBox_SetItemTemplateXMLFromMem(l.Handle, data)
}

// 列表框_置项模板从资源ZIP. 设置项模板文件.
//.
//.
//.
//, 可填0.
// ff:置项模板并按资源ZIP
// id:RC资源ID
// pFileName:项模板文件名
// pPassword:zip密码
// hModule:模块句柄
func (l *ListBox) SetItemTemplateXMLFromZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {
	return xc.XListBox_SetItemTemplateXMLFromZipRes(l.Handle, id, pFileName, pPassword, hModule)
}

// 列表框_取项模板. 获取列表项模板, 返回项模板句柄.
// ff:取项模板
func (l *ListBox) GetItemTemplate() int {
	return xc.XListBox_GetItemTemplate(l.Handle)
}

/*
以下都是事件
*/

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE func(pItem *xc.ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE1 func(hEle int, pItem *xc.ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE_END func(pItem *xc.ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE_END1 func(hEle int, pItem *xc.ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LISTBOX_TEMP_DESTROY func(pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LISTBOX_TEMP_DESTROY1 func(hEle int, pItem *xc.ListBox_Item_, nFlag int, pbHandled *bool) int
type XE_LISTBOX_TEMP_ADJUST_COORDINATE func(pItem *xc.ListBox_Item_, pbHandled *bool) int            // 列表框元素,项模板调整坐标. 已停用.
type XE_LISTBOX_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *xc.ListBox_Item_, pbHandled *bool) int // 列表框元素,项模板调整坐标. 已停用.
type XE_LISTBOX_DRAWITEM func(hDraw int, pItem *xc.ListBox_Item_, pbHandled *bool) int               // 列表框元素,项绘制事件.
type XE_LISTBOX_DRAWITEM1 func(hEle int, hDraw int, pItem *xc.ListBox_Item_, pbHandled *bool) int    // 列表框元素,项绘制事件.
type XE_LISTBOX_SELECT func(iItem int32, pbHandled *bool) int                                        // 列表框元素,项选择事件.
type XE_LISTBOX_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int                             // 列表框元素,项选择事件.

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
// ff:事件_项模板创建
// pFun:
func (l *ListBox) Event_LISTBOX_TEMP_CREATE(pFun XE_LISTBOX_TEMP_CREATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE, pFun)
}

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
// ff:事件_项模板创建1
// pFun:
func (l *ListBox) Event_LISTBOX_TEMP_CREATE1(pFun XE_LISTBOX_TEMP_CREATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE, pFun)
}

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
// ff:事件_项模板创建完成
// pFun:
func (l *ListBox) Event_LISTBOX_TEMP_CREATE_END(pFun XE_LISTBOX_TEMP_CREATE_END) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE_END, pFun)
}

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
// ff:事件_项模板创建完成1
// pFun:
func (l *ListBox) Event_LISTBOX_TEMP_CREATE_END1(pFun XE_LISTBOX_TEMP_CREATE_END1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_TEMP_CREATE_END, pFun)
}

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
// ff:事件_项模板销毁
// pFun:
func (l *ListBox) Event_LISTBOX_TEMP_DESTROY(pFun XE_LISTBOX_TEMP_DESTROY) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_TEMP_DESTROY, pFun)
}

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
// ff:事件_项模板销毁1
// pFun:
func (l *ListBox) Event_LISTBOX_TEMP_DESTROY1(pFun XE_LISTBOX_TEMP_DESTROY1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_TEMP_DESTROY, pFun)
}

// 列表框元素,项模板调整坐标. 已停用.
// ff:停用_项模板调整坐标
// pFun:
func (l *ListBox) Event_LISTBOX_TEMP_ADJUST_COORDINATE(pFun XE_LISTBOX_TEMP_ADJUST_COORDINATE) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表框元素,项模板调整坐标. 已停用.
// ff:停用_项模板调整坐标1
// pFun:
func (l *ListBox) Event_LISTBOX_TEMP_ADJUST_COORDINATE1(pFun XE_LISTBOX_TEMP_ADJUST_COORDINATE1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表框元素,项绘制事件.
// ff:事件_项绘制事件
// pFun:
func (l *ListBox) Event_LISTBOX_DRAWITEM(pFun XE_LISTBOX_DRAWITEM) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_DRAWITEM, pFun)
}

// 列表框元素,项绘制事件.
// ff:事件_项绘制事件1
// pFun:
func (l *ListBox) Event_LISTBOX_DRAWITEM1(pFun XE_LISTBOX_DRAWITEM1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_DRAWITEM, pFun)
}

// 列表框元素,项选择事件.
// ff:事件_项选择事件
// pFun:
func (l *ListBox) Event_LISTBOX_SELECT(pFun XE_LISTBOX_SELECT) bool {
	return xc.XEle_RegEventC(l.Handle, xcc.XE_LISTBOX_SELECT, pFun)
}

// 列表框元素,项选择事件.
// ff:事件_项选择事件1
// pFun:
func (l *ListBox) Event_LISTBOX_SELECT1(pFun XE_LISTBOX_SELECT1) bool {
	return xc.XEle_RegEventC1(l.Handle, xcc.XE_LISTBOX_SELECT, pFun)
}
