package 炫彩基类

import (
	"unsafe"
	
	"github.com/888go/xcgui/common"
	
	"github.com/888go/xcgui/xcc"
)

// 列表视_创建, 返回元素句柄.
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
func X列表视_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xListView_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 列表视_创建Ex. 创建列表视图元素, 使用内置项模板, 自动创建数据适配器, 返回元素句柄.
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
func X列表视_创建Ex(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄, 列数量 int32) int {
	r, _, _ := xListView_CreateEx.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄), uintptr(列数量))
	return int(r)
}

// 列表视_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
func X列表视_创建数据适配器(元素句柄 int) int {
	r, _, _ := xListView_CreateAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器XAdListView.
func X列表视_绑定数据适配器(元素句柄 int, 数据适配器 int) int {
	r, _, _ := xListView_BindAdapter.Call(uintptr(元素句柄), uintptr(数据适配器))
	return int(r)
}

// 列表视_取数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.
func X列表视_取数据适配器(元素句柄 int) int {
	r, _, _ := xListView_GetAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_置项模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 文件名.
func X列表视_置项模板文件(元素句柄 int, 文件名 string) bool {
	r, _, _ := xListView_SetItemTemplateXML.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文件名))
	return r != 0
}

// 列表视_置项模板从字符串.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.
func X列表视_置项模板从字符串(元素句柄 int, 字符串 string) bool {
	r, _, _ := xListView_SetItemTemplateXMLFromString.Call(uintptr(元素句柄), W2A(字符串))
	return r != 0
}

// 列表视_置项模板, 置列表项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.
func X列表视_置项模板(元素句柄 int, 模板句柄 int) bool {
	r, _, _ := xListView_SetItemTemplate.Call(uintptr(元素句柄), uintptr(模板句柄))
	return r != 0
}

// 列表视_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func X列表视_取模板对象(元素句柄 int, 组索引 int, 项索引 int, 模板项ID int) int {
	r, _, _ := xListView_GetTemplateObject.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), uintptr(模板项ID))
	return int(r)
}

// 列表视_取模板对象组, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// nTempItemID: 模板项ID.
func X列表视_取模板对象组(元素句柄 int, 组索引 int, 模板项ID int) int {
	r, _, _ := xListView_GetTemplateObjectGroup.Call(uintptr(元素句柄), uintptr(组索引), uintptr(模板项ID))
	return int(r)
}

// 列表视_取对象所在项, 获取当前对象所在模板实例, 属于列表视中哪一个项.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func X列表视_取对象所在项(元素句柄 int, 对象句柄 int, 接收组索引 *int32, 接收项索引 *int32) bool {
	r, _, _ := xListView_GetItemIDFromHXCGUI.Call(uintptr(元素句柄), uintptr(对象句柄), uintptr(unsafe.Pointer(接收组索引)), uintptr(unsafe.Pointer(接收项索引)))
	return r != 0
}

// 列表视_测试点击项, 检查坐标点所在项.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
//
// pOutGroup: 接收组索引.
//
// pOutItem: 接收项索引.
func X列表视_测试点击项(元素句柄 int, 坐标点 *POINT, 接收组索引 *int32, 接收项索引 *int32) bool {
	r, _, _ := xListView_HitTest.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)), uintptr(unsafe.Pointer(接收组索引)), uintptr(unsafe.Pointer(接收项索引)))
	return r != 0
}

// 列表视_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
//
// pOutGroup: 接收做索引.
//
// pOutItem: 接收项索引.
func X列表视_测试点击项EX(元素句柄 int, 坐标点 *POINT, 接收做索引 *int32, 接收项索引 *int32) bool {
	r, _, _ := xListView_HitTestOffset.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)), uintptr(unsafe.Pointer(接收做索引)), uintptr(unsafe.Pointer(接收项索引)))
	return r != 0
}

// 列表视_启用多选.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表视_启用多选(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xListView_EnableMultiSel.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表视_启用模板复用.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表视_启用模板复用(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xListView_EnableTemplateReuse.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表视_启用虚表.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表视_启用虚表(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xListView_EnableVirtualTable.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表视_置虚表项数量.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// nCount: 项数量.
func X列表视_置虚表项数量(元素句柄 int, 组索引 int, 项数量 int) bool {
	r, _, _ := xListView_SetVirtualItemCount.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项数量))
	return r != 0
}

// 列表视_置项背景绘制标志, 置是否绘制指定状态下项的背景.
//
// hEle: 元素句柄.
//
// nFlags: 标志位: List_DrawItemBk_Flag_.
func X列表视_置项背景绘制标志(元素句柄 int, 标志位 炫彩常量类.List_DrawItemBk_Flag_) int {
	r, _, _ := xListView_SetDrawItemBkFlags.Call(uintptr(元素句柄), uintptr(标志位))
	return int(r)
}

// 列表视_置选择项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func X列表视_置选择项(元素句柄 int, 组索引 int, 项索引 int) bool {
	r, _, _ := xListView_SetSelectItem.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引))
	return r != 0
}

// 列表视_取选择项.
//
// hEle: 元素句柄.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func X列表视_取选择项(元素句柄 int, 接收组索引 *int32, 接收项索引 *int32) bool {
	r, _, _ := xListView_GetSelectItem.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(接收组索引)), uintptr(unsafe.Pointer(接收项索引)))
	return r != 0
}

// 列表视_添加选择项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func X列表视_添加选择项(元素句柄 int, 组索引 int, 项索引 int) bool {
	r, _, _ := xListView_AddSelectItem.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引))
	return r != 0
}

// 列表视_显示指定项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func X列表视_显示指定项(元素句柄 int, 组索引 int, 项索引 int) int {
	r, _, _ := xListView_VisibleItem.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引))
	return int(r)
}

// 列表视_取可视项范围, 获取当前可见项范围.
//
// hEle: 元素句柄.
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
func X列表视_取可视项范围(元素句柄 int, 组1 *int32, 组2 *int32, 可视开始组 *int32, 可视开始项 *int32, 可视结束组 *int32, 可视结束项 *int32) int {
	r, _, _ := xListView_GetVisibleItemRange.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(组1)), uintptr(unsafe.Pointer(组2)), uintptr(unsafe.Pointer(可视开始组)), uintptr(unsafe.Pointer(可视开始项)), uintptr(unsafe.Pointer(可视结束组)), uintptr(unsafe.Pointer(可视结束项)))
	return int(r)
}

// 列表视_取选择项数量.
//
// hEle: 元素句柄.
func X列表视_取选择项数量(元素句柄 int) int {
	r, _, _ := xListView_GetSelectItemCount.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_取选择项全部, 获取选择的项ID, 返回接收项数量.
//
// hEle: 元素句柄.
//
// pArray: 数组.
//
// nArraySize: 数组大小.
func X列表视_取选择项全部(元素句柄 int, 数组 *[]ListView_Item_Id_, 数组大小 int) int {
	if 数组大小 < 1 {
		return 0
	}
	*数组 = make([]ListView_Item_Id_, 数组大小)
	r, _, _ := xListView_GetSelectAll.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(&(*数组)[0])), uintptr(数组大小))
	return int(r)
}

// 列表视_置选择项全部, 选择所有的项.
//
// hEle: 元素句柄.
func X列表视_置选择项全部(元素句柄 int) int {
	r, _, _ := xListView_SetSelectAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_取消选择项全部, 取消选择所有项.
//
// hEle: 元素句柄.
func X列表视_取消选择项全部(元素句柄 int) int {
	r, _, _ := xListView_CancelSelectAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_置列间隔, 置列间隔大小.
//
// hEle: 元素句柄.
//
// space: 间隔大小.
func X列表视_置列间隔(元素句柄 int, 间隔大小 int) int {
	r, _, _ := xListView_SetColumnSpace.Call(uintptr(元素句柄), uintptr(间隔大小))
	return int(r)
}

// 列表视_置行间隔, 置行间隔大小.
//
// hEle: 元素句柄.
//
// space: 间隔大小.
func X列表视_置行间隔(元素句柄 int, 间隔大小 int) int {
	r, _, _ := xListView_SetRowSpace.Call(uintptr(元素句柄), uintptr(间隔大小))
	return int(r)
}

// 列表视_置项大小.
//
// hEle: 元素句柄.
//
// width: 宽度.
//
// height: 高度.
func X列表视_置项大小(元素句柄 int, 宽度 int, 高度 int) int {
	r, _, _ := xListView_SetItemSize.Call(uintptr(元素句柄), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 列表视_取项大小.
//
// hEle: 元素句柄.
//
// pSize: 接收返回大小.
func X列表视_取项大小(元素句柄 int, 接收返回大小 *SIZE) int {
	r, _, _ := xListView_GetItemSize.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(接收返回大小)))
	return int(r)
}

// 列表视_置组高度.
//
// hEle: 元素句柄.
//
// height: 高度.
func X列表视_置组高度(元素句柄 int, 高度 int) int {
	r, _, _ := xListView_SetGroupHeight.Call(uintptr(元素句柄), uintptr(高度))
	return int(r)
}

// 列表视_取组高度.
//
// hEle: 元素句柄.
func X列表视_取组高度(元素句柄 int) int {
	r, _, _ := xListView_GetGroupHeight.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_置组用户数据.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// nData: 数据.
func X列表视_置组用户数据(元素句柄 int, 组索引 int, 数据 int) int {
	r, _, _ := xListView_SetGroupUserData.Call(uintptr(元素句柄), uintptr(组索引), uintptr(数据))
	return int(r)
}

// 列表视_置项用户数据.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nData: 数据.
func X列表视_置项用户数据(元素句柄 int, 组索引 int, 项索引 int, 数据 int) int {
	r, _, _ := xListView_SetItemUserData.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), uintptr(数据))
	return int(r)
}

// 列表视_取组用户数据.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
func X列表视_取组用户数据(元素句柄 int, 组索引 int) int {
	r, _, _ := xListView_GetGroupUserData.Call(uintptr(元素句柄), uintptr(组索引))
	return int(r)
}

// 列表视_取项用户数据.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func X列表视_取项用户数据(元素句柄 int, 组索引 int, 项索引 int) int {
	r, _, _ := xListView_GetItemUserData.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引))
	return int(r)
}

// 列表视_刷新项数据.
//
// hEle: 元素句柄.
func X列表视_刷新项数据(元素句柄 int) int {
	r, _, _ := xListView_RefreshData.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引, 如果为-1, 代表为组.
func X列表视_刷新指定项(元素句柄 int, 组索引 int, 项索引 int) int {
	r, _, _ := xListView_RefreshItem.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引))
	return int(r)
}

// 列表视_展开组, 成功返回TRUE否则返回FALSE, 如果状态没有改变返回FALSE.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// bExpand: 是否展开.
func X列表视_展开组(元素句柄 int, 组索引 int, 是否展开 bool) bool {
	r, _, _ := xListView_ExpandGroup.Call(uintptr(元素句柄), uintptr(组索引), 炫彩工具类.BoolPtr(是否展开))
	return r != 0
}

// 列表视_组添加列, 返回列索引.
//
// hEle: 元素句柄.
//
// pName: 字段称.
func X列表视_组添加列(元素句柄 int, 字段称 string) int {
	r, _, _ := xListView_Group_AddColumn.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 列表视_组添加项文本, 返回组索引.
//
// hEle: 元素句柄.
//
// pValue: 值.
//
// iPos: 插入位置.
func X列表视_组添加项文本(元素句柄 int, 值 string, 插入位置 int) int {
	r, _, _ := xListView_Group_AddItemText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(值), uintptr(插入位置))
	return int(r)
}

// 列表视_组添加项文本扩展, 返回组索引.
//
// hEle: 元素句柄.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置.
func X列表视_组添加项文本EX(元素句柄 int, 字段称 string, 值 string, 插入位置 int) int {
	r, _, _ := xListView_Group_AddItemTextEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值), uintptr(插入位置))
	return int(r)
}

// 列表视_组添加项图片, 返回组索引.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func X列表视_组添加项图片(元素句柄 int, 图片句柄 int, 插入位置 int) int {
	r, _, _ := xListView_Group_AddItemImage.Call(uintptr(元素句柄), uintptr(图片句柄), uintptr(插入位置))
	return int(r)
}

// 列表视_组添加项图片扩展, 返回组索引.
//
// hEle: 元素句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func X列表视_组添加项图片EX(元素句柄 int, 字段称 string, 图片句柄 int, 插入位置 int) int {
	r, _, _ := xListView_Group_AddItemImageEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄), uintptr(插入位置))
	return int(r)
}

// 列表视_组置文本.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func X列表视_组置文本(元素句柄 int, 组索引 int, 列索引 int, 值 string) bool {
	r, _, _ := xListView_Group_SetText.Call(uintptr(元素句柄), uintptr(组索引), uintptr(列索引), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 列表视_组置文本扩展.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func X列表视_组置文本EX(元素句柄 int, 组索引 int, 字段名 string, 值 string) bool {
	r, _, _ := xListView_Group_SetTextEx.Call(uintptr(元素句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 列表视_组置图片.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func X列表视_组置图片(元素句柄 int, 组索引 int, 列索引 int, 图片句柄 int) bool {
	r, _, _ := xListView_Group_SetImage.Call(uintptr(元素句柄), uintptr(组索引), uintptr(列索引), uintptr(图片句柄))
	return r != 0
}

// 列表视_组置图片扩展.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func X列表视_组置图片EX(元素句柄 int, 组索引 int, 字段名 string, 图片句柄 int) bool {
	r, _, _ := xListView_Group_SetImageEx.Call(uintptr(元素句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名), uintptr(图片句柄))
	return r != 0
}

// 列表视_组获取数量, 返回组数量.
//
// hEle: 元素句柄.
func X列表视_组获取数量(元素句柄 int) int {
	r, _, _ := xListView_Group_GetCount.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_项获取数量, 成功返回项数量, 否则返回 XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
func X列表视_项获取数量(元素句柄 int, 组索引 int) int {
	r, _, _ := xListView_Item_GetCount.Call(uintptr(元素句柄), uintptr(组索引))
	return int(r)
}

// 列表视_项添加列, 返回列索引.
//
// hEle: 元素句柄.
//
// pName: 字段名.
func X列表视_项添加列(元素句柄 int, 字段名 string) int {
	r, _, _ := xListView_Item_AddColumn.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字段名))
	return int(r)
}

// 列表视_项添加文本, 返回项索引.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func X列表视_项添加文本(元素句柄 int, 组索引 int, 值 string, 插入位置 int) int {
	r, _, _ := xListView_Item_AddItemText.Call(uintptr(元素句柄), uintptr(组索引), 炫彩工具类.StrPtr(值), uintptr(插入位置))
	return int(r)
}

// 列表视_项添加文本扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func X列表视_项添加文本EX(元素句柄 int, 组索引 int, 字段名 string, 值 string, 插入位置 int) int {
	r, _, _ := xListView_Item_AddItemTextEx.Call(uintptr(元素句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名), 炫彩工具类.StrPtr(值), uintptr(插入位置))
	return int(r)
}

// 列表视_项添加图片, 返回项索引.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func X列表视_项添加图片(元素句柄 int, 组索引 int, 图片句柄 int, 插入位置 int) int {
	r, _, _ := xListView_Item_AddItemImage.Call(uintptr(元素句柄), uintptr(组索引), uintptr(图片句柄), uintptr(插入位置))
	return int(r)
}

// 列表视_项添加图片扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func X列表视_项添加图片EX(元素句柄 int, 组索引 int, 字段名 string, 图片句柄 int, 插入位置 int) int {
	r, _, _ := xListView_Item_AddItemImageEx.Call(uintptr(元素句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名), uintptr(图片句柄), uintptr(插入位置))
	return int(r)
}

// 列表视_项置文本.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func X列表视_项置文本(元素句柄 int, 组索引 int, 项索引 int, 列索引 int, 值 string) bool {
	r, _, _ := xListView_Item_SetText.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), uintptr(列索引), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 列表视_项置文本扩展.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pValue: 值.
func X列表视_项置文本EX(元素句柄 int, 组索引 int, 项索引 int, 字段名 string, 值 string) bool {
	r, _, _ := xListView_Item_SetTextEx.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), 炫彩工具类.StrPtr(字段名), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 列表视_项置图片.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func X列表视_项置图片(元素句柄 int, 组索引 int, 项索引 int, 列索引 int, 图片句柄 int) bool {
	r, _, _ := xListView_Item_SetImage.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), uintptr(列索引), uintptr(图片句柄))
	return r != 0
}

// 列表视_项置图片扩展.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 列名称.
//
// hImage: 图片句柄.
func X列表视_项置图片EX(元素句柄 int, 组索引 int, 项索引 int, 列名称 string, 图片句柄 int) bool {
	r, _, _ := xListView_Item_SetImageEx.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), 炫彩工具类.StrPtr(列名称), uintptr(图片句柄))
	return r != 0
}

// 列表视_组删除项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
func X列表视_组删除项(元素句柄 int, 组索引 int) bool {
	r, _, _ := xListView_Group_DeleteItem.Call(uintptr(元素句柄), uintptr(组索引))
	return r != 0
}

// 列表视_组删除全部子项.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
func X列表视_组删除全部子项(元素句柄 int, 组索引 int) int {
	r, _, _ := xListView_Group_DeleteAllChildItem.Call(uintptr(元素句柄), uintptr(组索引))
	return int(r)
}

// 列表视_项删除.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func X列表视_项删除(元素句柄 int, 组索引 int, 项索引 int) bool {
	r, _, _ := xListView_Item_DeleteItem.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引))
	return r != 0
}

// 列表视_删除全部.
//
// hEle: 元素句柄.
func X列表视_删除全部(元素句柄 int) int {
	r, _, _ := xListView_DeleteAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_删除全部组.
//
// hEle: 元素句柄.
func X列表视_删除全部组(元素句柄 int) int {
	r, _, _ := xListView_DeleteAllGroup.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_删除全部项.
//
// hEle: 元素句柄.
func X列表视_删除全部项(元素句柄 int) int {
	r, _, _ := xListView_DeleteAllItem.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_组删除列.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
func X列表视_组删除列(元素句柄 int, 列索引 int) int {
	r, _, _ := xListView_DeleteColumnGroup.Call(uintptr(元素句柄), uintptr(列索引))
	return int(r)
}

// 列表视_项删除列.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
func X列表视_项删除列(元素句柄 int, 列索引 int) int {
	r, _, _ := xListView_DeleteColumnItem.Call(uintptr(元素句柄), uintptr(列索引))
	return int(r)
}

// 列表视_项获取文本.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func X列表视_项获取文本(元素句柄 int, 组索引 int, 项索引 int, 字段称 string) string {
	r, _, _ := xListView_Item_GetTextEx.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), 炫彩工具类.StrPtr(字段称))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表视_项获取图片扩展, 返回图片句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func X列表视_项获取图片EX(元素句柄 int, 组索引 int, 项索引 int, 字段称 string) int {
	r, _, _ := xListView_Item_GetImageEx.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 列表视_组取文本, 返回文本内容.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func X列表视_组取文本(元素句柄 int, 组索引 int, 列索引 int) string {
	r, _, _ := xListView_Group_GetText.Call(uintptr(元素句柄), uintptr(组索引), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表视_组取文本扩展, 返回文本内容.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func X列表视_组取文本EX(元素句柄 int, 组索引 int, 字段名称 string) string {
	r, _, _ := xListView_Group_GetTextEx.Call(uintptr(元素句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名称))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表视_组取图片, 返回图片句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func X列表视_组取图片(元素句柄 int, 组索引 int, 列索引 int) int {
	r, _, _ := xListView_Group_GetImage.Call(uintptr(元素句柄), uintptr(组索引), uintptr(列索引))
	return int(r)
}

// 列表视_组取图片扩展, 返回图片句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func X列表视_组取图片EX(元素句柄 int, 组索引 int, 字段名称 string) int {
	r, _, _ := xListView_Group_GetImageEx.Call(uintptr(元素句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名称))
	return int(r)
}

// 列表视_项取文本, 返回文本内容.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func X列表视_项取文本(元素句柄 int, 组索引 int, 项索引 int, 列索引 int) string {
	r, _, _ := xListView_Item_GetText.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表视_项取图片, 返回图片句柄.
//
// hEle: 元素句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func X列表视_项取图片(元素句柄 int, 组索引 int, 项索引 int, 列索引 int) int {
	r, _, _ := xListView_Item_GetImage.Call(uintptr(元素句柄), uintptr(组索引), uintptr(项索引), uintptr(列索引))
	return int(r)
}

// 列表视_置拖动矩形颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色.
//
// width: 线宽度.
func X列表视_置拖动矩形颜色(元素句柄 int, ABGR颜色 int, 线宽度 int) int {
	r, _, _ := xListView_SetDragRectColor.Call(uintptr(元素句柄), uintptr(ABGR颜色), uintptr(线宽度))
	return int(r)
}

// 列表视_置项模板从内存.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func X列表视_置项模板从内存(元素句柄 int, 模板数据 []byte) bool {
	r, _, _ := xListView_SetItemTemplateXMLFromMem.Call(uintptr(元素句柄), 炫彩工具类.ByteSliceDataPtr(&模板数据), uintptr(len(模板数据)))
	return r != 0
}

// 列表视_置项模板从资源ZIP.
//
// hEle: 元素句柄.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func X列表视_置项模板从资源ZIP(元素句柄 int, RC资源ID int32, 文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	r, _, _ := xListView_SetItemTemplateXMLFromZipRes.Call(uintptr(元素句柄), uintptr(RC资源ID), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(zip密码), 模块句柄)
	return r != 0
}

// 列表视_取项模板, 返回项模板句柄.
//
// hEle: 元素句柄.
func X列表视_取项模板(元素句柄 int) int {
	r, _, _ := xListView_GetItemTemplate.Call(uintptr(元素句柄))
	return int(r)
}

// 列表视_取项模板组, 返回项模板组句柄.
//
// hEle: 元素句柄.
func X列表视_取项模板组(元素句柄 int) int {
	r, _, _ := xListView_GetItemTemplateGroup.Call(uintptr(元素句柄))
	return int(r)
}
