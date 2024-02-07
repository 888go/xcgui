package 炫彩基类

import (
	"unsafe"
	
	"github.com/888go/xcgui/common"
	
	"github.com/888go/xcgui/xcc"
)

// 列表框_创建, 创建列表框元素, 返回元素句柄.
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
func X列表框_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xListBox_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 列表框_创建Ex. 创建列表框元素, 使用内置项模板, 自动创建数据适配器, 返回元素句柄.
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
//
// col_extend_count: 列数量. 例如: 内置模板是1列, 如果数据有5列, 那么此参数填5.
func X列表框_创建Ex(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄, 列数量 int32) int {
	r, _, _ := xListBox_CreateEx.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄), uintptr(列数量))
	return int(r)
}

// 列表框_启用固定行高.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表框_启用固定行高(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xListBox_EnableFixedRowHeight.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表框_启用模板复用.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表框_启用模板复用(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xListBox_EnableTemplateReuse.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表框_启用虚表.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表框_启用虚表(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xListBox_EnableVirtualTable.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表框_置虚表行数.
//
// hEle: 元素句柄.
//
// nRowCount: 行数.
func X列表框_置虚表行数(元素句柄 int, 行数 int) int {
	r, _, _ := xListBox_SetVirtualRowCount.Call(uintptr(元素句柄), uintptr(行数))
	return int(r)
}

// 列表框_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// hEle: 元素句柄.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func X列表框_置绘制项背景标志(元素句柄 int, 标志位 炫彩常量类.List_DrawItemBk_Flag_) int {
	r, _, _ := xListBox_SetDrawItemBkFlags.Call(uintptr(元素句柄), uintptr(标志位))
	return int(r)
}

// 列表框_置项数据, 设置项用户数据.
//
// hEle: 元素句柄.
//
// iItem: 想索引.
//
// nUserData: 用户数据.
func X列表框_置项数据(元素句柄 int, 项索引 int, 用户数据 int) bool {
	r, _, _ := xListBox_SetItemData.Call(uintptr(元素句柄), uintptr(项索引), uintptr(用户数据))
	return r != 0
}

// 列表框_取项数据, 获取项用户数据.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表框_取项数据(元素句柄 int, 项索引 int) int {
	r, _, _ := xListBox_GetItemData.Call(uintptr(元素句柄), uintptr(项索引))
	return int(r)
}

// 列表框_置项信息.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pItem: 项信息.
func X列表框_置项信息(元素句柄 int, 项索引 int, 项信息 *ListBox_Item_Info_) bool {
	r, _, _ := xListBox_SetItemInfo.Call(uintptr(元素句柄), uintptr(项索引), uintptr(unsafe.Pointer(项信息)))
	return r != 0
}

// 列表框_取项背景信息, 获取项信息.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pItem: 项信息.
func X列表框_取项背景信息(元素句柄 int, 项索引 int, 项信息 *ListBox_Item_Info_) bool {
	r, _, _ := xListBox_GetItemInfo.Call(uintptr(元素句柄), uintptr(项索引), uintptr(unsafe.Pointer(项信息)))
	return r != 0
}

// 列表框_置选择项, 设置选择选.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表框_置选择项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xListBox_SetSelectItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表框_取选择项, 返回项索引.
//
// hEle: 元素句柄.
func X列表框_取选择项(元素句柄 int) int {
	r, _, _ := xListBox_GetSelectItem.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_添加选择项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表框_添加选择项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xListBox_AddSelectItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表框_取消选择项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表框_取消选择项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xListBox_CancelSelectItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表框_取消选择全部, 如果之前有选择状态的项返回TRUE, 此时可以更新UI, 否则返回FALSE.
//
// hEle: 元素句柄.
func X列表框_取消选择全部(元素句柄 int) bool {
	r, _, _ := xListBox_CancelSelectAll.Call(uintptr(元素句柄))
	return r != 0
}

// 列表框_取全部选择, 获取所有选择项, 返回接收数量.
//
// hEle: 元素句柄.
//
// pArray: 数组缓冲区.
//
// nArraySize: 数组大小.
func X列表框_取全部选择(元素句柄 int, 数组缓冲区 *[]int32, 数组大小 int) int {
	if 数组大小 < 1 {
		return 0
	}
	*数组缓冲区 = make([]int32, 数组大小)
	r, _, _ := xListBox_GetSelectAll.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(&(*数组缓冲区)[0])), uintptr(数组大小))
	return int(r)
}

// 列表框_取选择项数量, 获取选择项数量.
//
// hEle: 元素句柄.
func X列表框_取选择项数量(元素句柄 int) int {
	r, _, _ := xListBox_GetSelectCount.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_取鼠标停留项, 返回鼠标所在项.
//
// hEle: 元素句柄.
func X列表框_取鼠标停留项(元素句柄 int) int {
	r, _, _ := xListBox_GetItemMouseStay.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_选择全部项.
//
// hEle: 元素句柄.
func X列表框_选择全部项(元素句柄 int) bool {
	r, _, _ := xListBox_SelectAll.Call(uintptr(元素句柄))
	return r != 0
}

// 列表框_显示指定项, 滚动视图让指定项可见.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表框_显示指定项(元素句柄 int, 项索引 int) int {
	r, _, _ := xListBox_VisibleItem.Call(uintptr(元素句柄), uintptr(项索引))
	return int(r)
}

// 列表框_取可视行范围, 获取当前可见行范围.
//
// hEle: 元素句柄.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func X列表框_取可视行范围(元素句柄 int, 开始行索引 *int32, 结束行索引 *int32) int {
	r, _, _ := xListBox_GetVisibleRowRange.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(开始行索引)), uintptr(unsafe.Pointer(结束行索引)))
	return int(r)
}

// 列表框_置项默认高度.
//
// hEle: 元素句柄.
//
// nHeight: 项高度.
//
// nSelHeight: 选中项高度.
func X列表框_置项默认高度(元素句柄 int, 项高度, 选中项高度 int32) int {
	r, _, _ := xListBox_SetItemHeightDefault.Call(uintptr(元素句柄), uintptr(项高度), uintptr(选中项高度))
	return int(r)
}

// 列表框_取项默认高度.
//
// hEle: 元素句柄.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func X列表框_取项默认高度(元素句柄 int, 高度, 选中时高度 *int32) int {
	r, _, _ := xListBox_GetItemHeightDefault.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(高度)), uintptr(unsafe.Pointer(选中时高度)))
	return int(r)
}

// 列表框_取所在行索引, 获取当前对象所在模板实例, 属于列表中哪一个项(行). 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func X列表框_取所在行索引(元素句柄 int, 对象句柄 int) int {
	r, _, _ := xListBox_GetItemIndexFromHXCGUI.Call(uintptr(元素句柄), uintptr(对象句柄))
	return int(r)
}

// 列表框_置行间距.
//
// hEle: 元素句柄.
//
// nSpace: 间距大小.
func X列表框_置行间距(元素句柄 int, 间距大小 int) int {
	r, _, _ := xListBox_SetRowSpace.Call(uintptr(元素句柄), uintptr(间距大小))
	return int(r)
}

// 列表框_取行间距.
//
// hEle: 元素句柄.
func X列表框_取行间距(元素句柄 int) int {
	r, _, _ := xListBox_GetRowSpace.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_测试点击项, 检测坐标点所在项, 返回项索引.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
func X列表框_测试点击项(元素句柄 int, 坐标点 *POINT) int {
	r, _, _ := xListBox_HitTest.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)))
	return int(r)
}

// 列表框_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移量, 返回项索引.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
func X列表框_测试点击项EX(元素句柄 int, 坐标点 *POINT) int {
	r, _, _ := xListBox_HitTestOffset.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)))
	return int(r)
}

// 列表框_置项模板文件, 设置列表项模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 文件名.
func X列表框_置项模板文件(元素句柄 int, 文件名 string) bool {
	r, _, _ := xListBox_SetItemTemplateXML.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文件名))
	return r != 0
}

// 列表框_置项模板, 设置列表项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.
func X列表框_置项模板(元素句柄 int, 模板句柄 int) bool {
	r, _, _ := xListBox_SetItemTemplate.Call(uintptr(元素句柄), uintptr(模板句柄))
	return r != 0
}

// 列表框_置项模板从字符串, 设置项布局模板文件.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.
func X列表框_置项模板从字符串(元素句柄 int, 字符串 string) bool {
	r, _, _ := xListBox_SetItemTemplateXMLFromString.Call(uintptr(元素句柄), W2A(字符串))
	return r != 0
}

// 列表框_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄, 成功返回对象句柄, 否则返回NULL.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func X列表框_取模板对象(元素句柄 int, 项索引 int, 模板项ID int) int {
	r, _, _ := xListBox_GetTemplateObject.Call(uintptr(元素句柄), uintptr(项索引), uintptr(模板项ID))
	return int(r)
}

// 列表框_启用多选, 是否启用多行选择功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表框_启用多选(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xListBox_EnableMultiSel.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表框_创建数据适配器, 创建数据适配器并绑定, 根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
func X列表框_创建数据适配器(元素句柄 int) int {
	r, _, _ := xListBox_CreateAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_绑定数据适配器, 绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄 XAdTable.
func X列表框_绑定数据适配器(元素句柄 int, 数据适配器句柄 int) int {
	r, _, _ := xListBox_BindAdapter.Call(uintptr(元素句柄), uintptr(数据适配器句柄))
	return int(r)
}

// 列表框_取数据适配器, 获取绑定的数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.
func X列表框_取数据适配器(元素句柄 int) int {
	r, _, _ := xListBox_GetAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_排序.
//
// hEle: 元素句柄.
//
// iColumnAdapter: 需要排序的数据在数据适配器中所属列索引.
//
// bAscending: 升序(TRUE)或降序(FALSE).
func X列表框_排序(元素句柄 int, 数据适配器列索引 int, 升序 bool) int {
	r, _, _ := xListBox_Sort.Call(uintptr(元素句柄), uintptr(数据适配器列索引), 炫彩工具类.BoolPtr(升序))
	return int(r)
}

// 列表框_刷新数据.
//
// hEle: 元素句柄.
func X列表框_刷新数据(元素句柄 int) int {
	r, _, _ := xListBox_RefreshData.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表框_刷新指定项(元素句柄 int, 项索引 int) int {
	r, _, _ := xListBox_RefreshItem.Call(uintptr(元素句柄), uintptr(项索引))
	return int(r)
}

// 列表框_添加项文本, XAdTable_AddItemText, 返回项索引.
//
// hEle:.
//
// pText:.
func X列表框_添加项文本(元素句柄 int, 文本 string) int {
	r, _, _ := xListBox_AddItemText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 列表框_添加项文本扩展, XAdTable_AddItemTextEx.
//
// hEle:.
//
// pName:.
//
// pText:.
func X列表框_添加项文本EX(元素句柄 int, 名称 string, 文本 string) int {
	r, _, _ := xListBox_AddItemTextEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 列表框_添加项图片, XAdTable_AddItemImage.
//
// hEle:.
//
// hImage:.
func X列表框_添加项图片(元素句柄 int, 图片 int) int {
	r, _, _ := xListBox_AddItemImage.Call(uintptr(元素句柄), uintptr(图片))
	return int(r)
}

// 列表框_添加项图片扩展, XAdTable_AddItemImageEx.
//
// hEle:.
//
// pName:.
//
// hImage:.
func X列表框_添加项图片EX(元素句柄 int, 名称 string, 图片 int) int {
	r, _, _ := xListBox_AddItemImageEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(名称), uintptr(图片))
	return int(r)
}

// 列表框_插入项文本.
//
// hEle:.
//
// iItem:.
//
// pValue:.
func X列表框_插入项文本(元素句柄 int, 项索引 int, 值 string) int {
	r, _, _ := xListBox_InsertItemText.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(值))
	return int(r)
}

// 列表框_插入项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pValue:.
func X列表框_插入项文本EX(元素句柄 int, 项索引 int, 名称 string, 值 string) int {
	r, _, _ := xListBox_InsertItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(值))
	return int(r)
}

// 列表框_插入项图片.
//
// hEle:.
//
// iItem:.
//
// hImage:.
func X列表框_插入项图片(元素句柄 int, 项索引 int, 图片 int) int {
	r, _, _ := xListBox_InsertItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(图片))
	return int(r)
}

// 列表框_插入项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// hImage:.
func X列表框_插入项图片EX(元素句柄 int, 项索引 int, 名称 string, 图片 int) int {
	r, _, _ := xListBox_InsertItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(图片))
	return int(r)
}

// 列表框_置项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func X列表框_置项文本(元素句柄 int, 项索引 int, 列索引 int, 文本 string) bool {
	r, _, _ := xListBox_SetItemText.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), 炫彩工具类.StrPtr(文本))
	return r != 0
}

// 列表框_置项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pText:.
func X列表框_置项文本EX(元素句柄 int, 项索引 int, 名称 string, 文本 string) bool {
	r, _, _ := xListBox_SetItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(文本))
	return r != 0
}

// 列表框_置项图片.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func X列表框_置项图片(元素句柄 int, 项索引 int, 列索引 int, 图片 int) bool {
	r, _, _ := xListBox_SetItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(图片))
	return r != 0
}

// 列表框_置项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// hImage:.
func X列表框_置项图片EX(元素句柄 int, 项索引 int, 名称 string, 图片 int) bool {
	r, _, _ := xListBox_SetItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(图片))
	return r != 0
}

// 列表框_置项整数值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func X列表框_置项整数值(元素句柄 int, 项索引 int, 列索引 int, 值 int) bool {
	r, _, _ := xListBox_SetItemInt.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(值))
	return r != 0
}

// 列表框_置项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// nValue:.
func X列表框_置项整数值EX(元素句柄 int, 项索引 int, 名称 string, 值 int) bool {
	r, _, _ := xListBox_SetItemIntEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(值))
	return r != 0
}

// 列表框_置项浮点值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func X列表框_置项浮点值(元素句柄 int, 项索引 int, 列索引 int, 值 float32) bool {
	r, _, _ := xListBox_SetItemFloat.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), 炫彩工具类.Float32Ptr(值))
	return r != 0
}

// 列表框_置项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func X列表框_置项浮点值EX(元素句柄 int, 项索引 int, 名称 string, 值 float32) bool {
	r, _, _ := xListBox_SetItemFloatEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), 炫彩工具类.Float32Ptr(值))
	return r != 0
}

// 列表框_取项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func X列表框_取项文本(元素句柄 int, 项索引 int, 列索引 int) string {
	r, _, _ := xListBox_GetItemText.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表框_取项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
func X列表框_取项文本EX(元素句柄 int, 项索引 int, 名称 string) string {
	r, _, _ := xListBox_GetItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表框_取项图片.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func X列表框_取项图片(元素句柄 int, 项索引 int, 列索引 int) int {
	r, _, _ := xListBox_GetItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引))
	return int(r)
}

// 列表框_取项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
func X列表框_取项图片EX(元素句柄 int, 项索引 int, 名称 string) int {
	r, _, _ := xListBox_GetItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称))
	return int(r)
}

// 列表框_取项整数值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func X列表框_取项整数值(元素句柄 int, 项索引 int, 列索引 int, 接收返还值 *int32) bool {
	r, _, _ := xListBox_GetItemInt.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(unsafe.Pointer(接收返还值)))
	return r != 0
}

// 列表框_取项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func X列表框_取项整数值EX(元素句柄 int, 项索引 int, 名称 string, 接收返还值 *int32) bool {
	r, _, _ := xListBox_GetItemIntEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(unsafe.Pointer(接收返还值)))
	return r != 0
}

// 列表框_取项浮点值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func X列表框_取项浮点值(元素句柄 int, 项索引 int, 列索引 int, 接收返还值 *float32) bool {
	r, _, _ := xListBox_GetItemFloat.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(unsafe.Pointer(接收返还值)))
	return r != 0
}

// 列表框_取项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func X列表框_取项浮点值EX(元素句柄 int, 项索引 int, 名称 string, 接收返还值 *float32) bool {
	r, _, _ := xListBox_GetItemFloatEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(unsafe.Pointer(接收返还值)))
	return r != 0
}

// 列表框_删除项.
//
// hEle:.
//
// iItem:.
func X列表框_删除项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xListBox_DeleteItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表框_删除项扩展.
//
// hEle:.
//
// iItem:.
//
// nCount:.
func X列表框_删除项EX(元素句柄 int, 项索引 int, 数量 int) bool {
	r, _, _ := xListBox_DeleteItemEx.Call(uintptr(元素句柄), uintptr(项索引), uintptr(数量))
	return r != 0
}

// 列表框_删除项全部.
//
// hEle:.
func X列表框_删除项全部(元素句柄 int) int {
	r, _, _ := xListBox_DeleteItemAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_删除列全部.
//
// hEle:.
func X列表框_删除列全部(元素句柄 int) int {
	r, _, _ := xListBox_DeleteColumnAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_取项数量AD.
//
// hEle:.
func X列表框_取项数量AD(元素句柄 int) int {
	r, _, _ := xListBox_GetCount_AD.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_取列数量AD.
//
// hEle:.
func X列表框_取列数量AD(元素句柄 int) int {
	r, _, _ := xListBox_GetCountColumn_AD.Call(uintptr(元素句柄))
	return int(r)
}

// 列表框_置分割线颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func X列表框_置分割线颜色(元素句柄 int, ABGR颜色值 int) int {
	r, _, _ := xListBox_SetSplitLineColor.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 列表框_置拖动矩形颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
//
// width: 线宽度.
func X列表框_置拖动矩形颜色(元素句柄 int, ABGR颜色值, 线宽度 int) int {
	r, _, _ := xListBox_SetDragRectColor.Call(uintptr(元素句柄), uintptr(ABGR颜色值), uintptr(线宽度))
	return int(r)
}

// 列表框_置项模板从内存. 设置项模板文件.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func X列表框_置项模板从内存(元素句柄 int, 模板数据 []byte) bool {
	r, _, _ := xListBox_SetItemTemplateXMLFromMem.Call(uintptr(元素句柄), 炫彩工具类.ByteSliceDataPtr(&模板数据), uintptr(len(模板数据)))
	return r != 0
}

// 列表框_置项模板从资源ZIP. 设置项模板文件.
//
// hEle: 元素句柄.
//
// id: RC资源ID.
//
// pFileName: 项模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func X列表框_置项模板从资源ZIP(元素句柄, RC资源ID int, 项模板文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	r, _, _ := xListBox_SetItemTemplateXMLFromZipRes.Call(uintptr(元素句柄), uintptr(RC资源ID), 炫彩工具类.StrPtr(项模板文件名), 炫彩工具类.StrPtr(zip密码), 模块句柄)
	return r != 0
}

// 列表框_取项模板. 获取列表项模板, 返回项模板句柄.
//
// hEle: 元素句柄.
func X列表框_取项模板(元素句柄 int) int {
	r, _, _ := xListBox_GetItemTemplate.Call(uintptr(元素句柄))
	return int(r)
}
