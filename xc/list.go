package 炫彩基类

import (
	"unsafe"
	
	"github.com/888go/xcgui/common"
	
	"github.com/888go/xcgui/xcc"
)

// 列表_创建, 创建列表元素, 返回元素句柄.
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
func X列表_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xList_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 列表_创建Ex, 创建列表元素, 使用内置项模板, 自动创建数据适配器, 返回元素句柄.
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
func X列表_创建Ex(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄, 列数量 int32) int {
	r, _, _ := xList_CreateEx.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄), uintptr(列数量))
	return int(r)
}

// 列表_增加列, 返回位置索引.
//
// hEle: 元素句柄.
//
// width: 列宽度.
func X列表_增加列(元素句柄 int, 列宽度 int) int {
	r, _, _ := xList_AddColumn.Call(uintptr(元素句柄), uintptr(列宽度))
	return int(r)
}

// 列表_插入列, 返回插入位置索引.
//
// hEle: 元素句柄.
//
// width: 列宽度.
//
// iItem: 插入位置索引.
func X列表_插入列(元素句柄 int, 列宽度 int, 项索引 int) int {
	r, _, _ := xList_InsertColumn.Call(uintptr(元素句柄), uintptr(列宽度), uintptr(项索引))
	return int(r)
}

// 列表_启用多选, 启用或关闭多选功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表_启用多选(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xList_EnableMultiSel.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表_启用拖动更改列宽, 启用拖动改变列宽度.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表_启用拖动更改列宽(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xList_EnableDragChangeColumnWidth.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表_启用垂直滚动条顶部对齐.
//
// hEle: 元素句柄.
//
// bTop: 是否启用.
func X列表_启用垂直滚动条顶部对齐(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xList_EnableVScrollBarTop.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表_启用项背景全行模式, 启用项背景全行填充模式.
//
// hEle: 元素句柄.
//
// bFull: 是否启用.
func X列表_启用项背景全行模式(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xList_EnableItemBkFullRow.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表_启用固定行高.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表_启用固定行高(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xList_EnableFixedRowHeight.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表_启用模板复用.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表_启用模板复用(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xList_EnableTemplateReuse.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表_启用虚表.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X列表_启用虚表(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xList_EnableVirtualTable.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 列表_置虚表行数.
//
// hEle: 元素句柄.
//
// nRowCount: 行数.
func X列表_置虚表行数(元素句柄 int, 行数 int) int {
	r, _, _ := xList_SetVirtualRowCount.Call(uintptr(元素句柄), uintptr(行数))
	return int(r)
}

// 列表_置排序, 设置排序属性.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
//
// iColumnAdapter: 需要排序的数据在数据适配器中的列索引.
//
// bEnable: 是否启用排序功能.
func X列表_置排序(元素句柄 int, 列索引 int, 适配器中列索引 int, 是否启用排序功能 bool) int {
	r, _, _ := xList_SetSort.Call(uintptr(元素句柄), uintptr(列索引), uintptr(适配器中列索引), 炫彩工具类.BoolPtr(是否启用排序功能))
	return int(r)
}

// 列表_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// hEle: 元素句柄.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func X列表_置绘制项背景标志(元素句柄 int, 标志位 炫彩常量类.List_DrawItemBk_Flag_) int {
	r, _, _ := xList_SetDrawItemBkFlags.Call(uintptr(元素句柄), uintptr(标志位))
	return int(r)
}

// 列表_置列宽.
//
// hEle: 元素句柄.
//
// iItem: 列索引.
//
// width: 宽度.
func X列表_置列宽(元素句柄 int, 项索引 int, 宽度 int) int {
	r, _, _ := xList_SetColumnWidth.Call(uintptr(元素句柄), uintptr(项索引), uintptr(宽度))
	return int(r)
}

// 列表_置列最小宽度.
//
// hEle: 元素句柄.
//
// iItem: 列索引.
//
// width: 宽度.
func X列表_置列最小宽度(元素句柄 int, 项索引 int, 宽度 int) int {
	r, _, _ := xList_SetColumnMinWidth.Call(uintptr(元素句柄), uintptr(项索引), uintptr(宽度))
	return int(r)
}

// 列表_置列宽度固定.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
//
// bFixed: 是否固定宽度.
func X列表_置列宽度固定(元素句柄 int, 列索引 int, 是否固定宽度 bool) int {
	r, _, _ := xList_SetColumnWidthFixed.Call(uintptr(元素句柄), uintptr(列索引), 炫彩工具类.BoolPtr(是否固定宽度))
	return int(r)
}

// 列表_取列宽度.
//
// hEle: 元素句柄.
//
// iColumn: 列索引.
func X列表_取列宽度(元素句柄 int, 列索引 int) int {
	r, _, _ := xList_GetColumnWidth.Call(uintptr(元素句柄), uintptr(列索引))
	return int(r)
}

// 列表_取列数量.
//
// hEle: 元素句柄.
func X列表_取列数量(元素句柄 int) int {
	r, _, _ := xList_GetColumnCount.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_置项数据, 设置项用户数据.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// data: 用户数据.
func X列表_置项数据(元素句柄 int, 项索引 int, 子项索引 int, 用户数据 int) bool {
	r, _, _ := xList_SetItemData.Call(uintptr(元素句柄), uintptr(项索引), uintptr(子项索引), uintptr(用户数据))
	return r != 0
}

// 列表_取项数据, 获取项用户数据.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
func X列表_取项数据(元素句柄 int, 项索引 int, 子项索引 int) int {
	r, _, _ := xList_GetItemData.Call(uintptr(元素句柄), uintptr(项索引), uintptr(子项索引))
	return int(r)
}

// 列表_置选择项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表_置选择项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xList_SetSelectItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表_取选择项, 返回项索引.
//
// hEle: 元素句柄.
func X列表_取选择项(元素句柄 int) int {
	r, _, _ := xList_GetSelectItem.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_取选择项数量, 获取选择项数量.
//
// hEle: 元素句柄.
func X列表_取选择项数量(元素句柄 int) int {
	r, _, _ := xList_GetSelectItemCount.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_添加选择项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表_添加选择项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xList_AddSelectItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表_置选择全部, 选择全部行.
//
// hEle: 元素句柄.
func X列表_置选择全部(元素句柄 int) int {
	r, _, _ := xList_SetSelectAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_取全部选择, 获取全部选择的行, 返回行数量.
//
// hEle: 元素句柄.
//
// pArray: 接收行索引数组.
//
// nArraySize: 数组大小.
func X列表_取全部选择(元素句柄 int, 接收行索引数组 *[]int32, 数组大小 int) int {
	if 数组大小 < 1 {
		return 0
	}
	*接收行索引数组 = make([]int32, 数组大小)
	r, _, _ := xList_GetSelectAll.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(&(*接收行索引数组)[0])), uintptr(数组大小))
	return int(r)
}

// 列表_显示指定项, 滚动视图让指定项可见.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表_显示指定项(元素句柄 int, 项索引 int) int {
	r, _, _ := xList_VisibleItem.Call(uintptr(元素句柄), uintptr(项索引))
	return int(r)
}

// 列表_取消选择项, 取消选择指定项(这里的项可以理解为行).
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表_取消选择项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xList_CancelSelectItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表_取消全部选择项, 取消选择所有项(这里的项可以理解为行).
//
// hEle: 元素句柄.
func X列表_取消全部选择项(元素句柄 int) int {
	r, _, _ := xList_CancelSelectAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_取列表头, 获取列表头元素, 返回列表头元素句柄.
//
// hEle: 元素句柄.
func X列表_取列表头(元素句柄 int) int {
	r, _, _ := xList_GetHeaderHELE.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_删除列.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表_删除列(元素句柄 int, 项索引 int) bool {
	r, _, _ := xList_DeleteColumn.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表_删除全部列, 删除所有的列, 仅删除List的, 数据适配器的列不变.
//
// hEle: 元素句柄.
func X列表_删除全部列(元素句柄 int) int {
	r, _, _ := xList_DeleteColumnAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄 XAdTable.
func X列表_绑定数据适配器(元素句柄 int, 数据适配器句柄 int) int {
	r, _, _ := xList_BindAdapter.Call(uintptr(元素句柄), uintptr(数据适配器句柄))
	return int(r)
}

// 列表_列表头绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄 XAdMap.
func X列表_列表头绑定数据适配器(元素句柄 int, 数据适配器句柄 int) int {
	r, _, _ := xList_BindAdapterHeader.Call(uintptr(元素句柄), uintptr(数据适配器句柄))
	return int(r)
}

// 列表_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
//
// colExtend_count:	列延伸-预计列表总列数, 默认值0; 限制最大延伸范围, 避免超出范围, 增加不必要的字段.
func X列表_创建数据适配器(元素句柄 int, 预计列表总列数 int) int {
	r, _, _ := xList_CreateAdapter.Call(uintptr(元素句柄), uintptr(预计列表总列数))
	return int(r)
}

// 列表_列表头创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.
func X列表_列表头创建数据适配器(元素句柄 int) int {
	r, _, _ := xList_CreateAdapterHeader.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_取数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.
func X列表_取数据适配器(元素句柄 int) int {
	r, _, _ := xList_GetAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_列表头获取数据适配器, 获取列表头数据适配器句柄.
//
// hEle: 元素句柄.
func X列表_列表头获取数据适配器(元素句柄 int) int {
	r, _, _ := xList_GetAdapterHeader.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_置项模板文件, 设置项布局模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 文件名.
func X列表_置项模板文件(元素句柄 int, 文件名 string) bool {
	r, _, _ := xList_SetItemTemplateXML.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文件名))
	return r != 0
}

// 列表_置项模板从字符串, 设置项布局模板文件.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.
func X列表_置项模板从字符串(元素句柄 int, 字符串 string) bool {
	r, _, _ := xList_SetItemTemplateXMLFromString.Call(uintptr(元素句柄), W2A(字符串))
	return r != 0
}

// 列表_置项模板, 设置列表项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.
func X列表_置项模板(元素句柄 int, 模板句柄 int) bool {
	r, _, _ := xList_SetItemTemplate.Call(uintptr(元素句柄), uintptr(模板句柄))
	return r != 0
}

// 列表_取项模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// nTempItemID: 模板项itemID.
func X列表_取项模板对象(元素句柄 int, 项索引 int, 子项索引 int, 模板项ID int) int {
	r, _, _ := xList_GetTemplateObject.Call(uintptr(元素句柄), uintptr(项索引), uintptr(子项索引), uintptr(模板项ID))
	return int(r)
}

// 列表_取对象所在行, 获取当前对象所在模板实例, 属于列表中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func X列表_取对象所在行(元素句柄 int, 对象句柄 int) int {
	r, _, _ := xList_GetItemIndexFromHXCGUI.Call(uintptr(元素句柄), uintptr(对象句柄))
	return int(r)
}

// 列表_取列表头模板对象, 列表头, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// iItem: 列表头项ID.
//
// nTempItemID: 模板项ID.
func X列表_取列表头模板对象(元素句柄 int, 项索引 int, 模板项ID int) int {
	r, _, _ := xList_GetHeaderTemplateObject.Call(uintptr(元素句柄), uintptr(项索引), uintptr(模板项ID))
	return int(r)
}

// 列表_取列表头对象所在行, 列表头, 获取当前对象所在模板实例, 属于列表头中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄.
func X列表_取列表头对象所在行(元素句柄 int, 对象句柄 int) int {
	r, _, _ := xList_GetHeaderItemIndexFromHXCGUI.Call(uintptr(元素句柄), uintptr(对象句柄))
	return int(r)
}

// 列表_置列表头高度.
//
// hEle: 元素句柄.
//
// height: 高度.
func X列表_置列表头高度(元素句柄 int, 高度 int) int {
	r, _, _ := xList_SetHeaderHeight.Call(uintptr(元素句柄), uintptr(高度))
	return int(r)
}

// 列表_取列表头高度.
//
// hEle: 元素句柄.
func X列表_取列表头高度(元素句柄 int) int {
	r, _, _ := xList_GetHeaderHeight.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_取可视行范围.
//
// hEle: 元素句柄.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func X列表_取可视行范围(元素句柄 int, 开始行索引 *int32, 结束行索引 *int32) int {
	r, _, _ := xList_GetVisibleRowRange.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(开始行索引)), uintptr(unsafe.Pointer(结束行索引)))
	return int(r)
}

// 列表_置项默认高度.
//
// hEle: 元素句柄.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func X列表_置项默认高度(元素句柄 int, 高度 int32, 选中时高度 int32) int {
	r, _, _ := xList_SetItemHeightDefault.Call(uintptr(元素句柄), uintptr(高度), uintptr(选中时高度))
	return int(r)
}

// 列表_取项默认高度.
//
// hEle: 元素句柄.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func X列表_取项默认高度(元素句柄 int, 高度 *int32, 选中时高度 *int32) int {
	r, _, _ := xList_GetItemHeightDefault.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(高度)), uintptr(unsafe.Pointer(选中时高度)))
	return int(r)
}

// 列表_置行间距.
//
// hEle: 元素句柄.
//
// nSpace: 行间距大小.
func X列表_置行间距(元素句柄 int, 行间距大小 int) int {
	r, _, _ := xList_SetRowSpace.Call(uintptr(元素句柄), uintptr(行间距大小))
	return int(r)
}

// 列表_取行间距.
//
// hEle: 元素句柄.
func X列表_取行间距(元素句柄 int) int {
	r, _, _ := xList_GetRowSpace.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_置锁定列左侧, 锁定列, 设置左侧锁定列分界列索引.
//
// hEle: 元素句柄.
//
// iColumn: 列索引, -1代表不锁定.
func X列表_置锁定列左侧(元素句柄 int, 列索引 int) int {
	r, _, _ := xList_SetLockColumnLeft.Call(uintptr(元素句柄), uintptr(列索引))
	return int(r)
}

// 列表_置锁定列右侧.
//
// hEle: 元素句柄.
//
// iColumn: 列索引, -1代表不锁定. 暂时只支持锁定末尾列.
func X列表_置锁定列右侧(元素句柄 int, 列索引 int) int {
	r, _, _ := xList_SetLockColumnRight.Call(uintptr(元素句柄), uintptr(列索引))
	return int(r)
}

// 列表_置锁定行底部, 设置是否锁定末尾行.
//
// hEle: 元素句柄.
//
// bLock: 是否锁定.
func X列表_置锁定行底部(元素句柄 int, 是否锁定 bool) int {
	r, _, _ := xList_SetLockRowBottom.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否锁定))
	return int(r)
}

// 列表_置锁定行底部重叠.
//
// hEle: 元素句柄.
//
// bOverlap: 是否重叠.
func X列表_置锁定行底部重叠(元素句柄 int, 是否重叠 bool) int {
	r, _, _ := xList_SetLockRowBottomOverlap.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否重叠))
	return int(r)
}

// 列表_测试点击项, 检测坐标点所在项.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func X列表_测试点击项(元素句柄 int, 坐标点 *POINT, 项索引 *int32, 子项索引 *int32) bool {
	r, _, _ := xList_HitTest.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)), uintptr(unsafe.Pointer(项索引)), uintptr(unsafe.Pointer(子项索引)))
	return r != 0
}

// 列表_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func X列表_测试点击项EX(元素句柄 int, 坐标点 *POINT, 项索引 *int32, 子项索引 *int32) bool {
	r, _, _ := xList_HitTestOffset.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)), uintptr(unsafe.Pointer(项索引)), uintptr(unsafe.Pointer(子项索引)))
	return r != 0
}

// 列表_刷新项数据.
//
// hEle: 元素句柄.
func X列表_刷新项数据(元素句柄 int) int {
	r, _, _ := xList_RefreshData.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X列表_刷新指定项(元素句柄 int, 项索引 int) int {
	r, _, _ := xList_RefreshItem.Call(uintptr(元素句柄), uintptr(项索引))
	return int(r)
}

// 列表_添加列文本.
//
// hEle: 元素句柄.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen. 你可以理解为创建表头数据适配器后, 内部有一个Map来存储每一列的表头名(列名), 这个name名就是Map的Key, 这个函数就相当于给每一列的Key赋值, 然后List会根据这个name名从Map读取Value来显示表头到界面.
//
// pText: 文本.
func X列表_添加列文本(元素句柄 int, 列宽 int, 名称 string, 文本 string) int {
	r, _, _ := xList_AddColumnText.Call(uintptr(元素句柄), uintptr(列宽), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 列表_添加列图片.
//
// hEle: 元素句柄.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen. 你可以理解为创建表头数据适配器后, 内部有一个Map来存储每一列的表头名(列名), 这个name名就是Map的Key, 这个函数就相当于给每一列的Key赋值, 然后List会根据这个name名从Map读取Value来显示表头到界面.
//
// hImage: 图片句柄.
func X列表_添加列图片(元素句柄 int, 列宽 int, 名称 string, 图片 int) int {
	r, _, _ := xList_AddColumnImage.Call(uintptr(元素句柄), uintptr(列宽), 炫彩工具类.StrPtr(名称), uintptr(图片))
	return int(r)
}

// 列表_添加项文本.
//
// hEle:.
//
// pText:.
func X列表_添加项文本(元素句柄 int, 文本 string) int {
	r, _, _ := xList_AddItemText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 列表_添加项文本扩展.
//
// hEle:.
//
// pName:.
//
// pText:.
func X列表_添加项文本EX(元素句柄 int, 名称 string, 文本 string) int {
	r, _, _ := xList_AddItemTextEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 列表_添加项图片.
//
// hEle:.
//
// hImage:.
func X列表_添加项图片(元素句柄 int, 图片 int) int {
	r, _, _ := xList_AddItemImage.Call(uintptr(元素句柄), uintptr(图片))
	return int(r)
}

// 列表_添加项图片扩展.
//
// hEle:.
//
// pName:.
//
// hImage:.
func X列表_添加项图片EX(元素句柄 int, 名称 string, 图片 int) int {
	r, _, _ := xList_AddItemImageEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(名称), uintptr(图片))
	return int(r)
}

// 列表_插入项文本.
//
// hEle:.
//
// iItem:.
//
// pValue:.
func X列表_插入项文本(元素句柄 int, 项索引 int, 值 string) int {
	r, _, _ := xList_InsertItemText.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(值))
	return int(r)
}

// 列表_插入项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pValue:.
func X列表_插入项文本EX(元素句柄 int, 项索引 int, 名称 string, 值 string) int {
	r, _, _ := xList_InsertItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(值))
	return int(r)
}

// 列表_插入项图片.
//
// hEle:.
//
// iItem:.
//
// hImage:.
func X列表_插入项图片(元素句柄 int, 项索引 int, 图片 int) int {
	r, _, _ := xList_InsertItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(图片))
	return int(r)
}

// 列表_插入项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// hImage:.
func X列表_插入项图片EX(元素句柄 int, 项索引 int, 名称 string, 图片 int) int {
	r, _, _ := xList_InsertItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(图片))
	return int(r)
}

// 列表_置项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func X列表_置项文本(元素句柄 int, 项索引 int, 列索引 int, 文本 string) bool {
	r, _, _ := xList_SetItemText.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), 炫彩工具类.StrPtr(文本))
	return r != 0
}

// 列表_置项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pText:.
func X列表_置项文本EX(元素句柄 int, 项索引 int, 名称 string, 文本 string) bool {
	r, _, _ := xList_SetItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(文本))
	return r != 0
}

// 列表_置项图片.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func X列表_置项图片(元素句柄 int, 项索引 int, 列索引 int, 图片 int) bool {
	r, _, _ := xList_SetItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(图片))
	return r != 0
}

// 列表_置项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// hImage:.
func X列表_置项图片EX(元素句柄 int, 项索引 int, 名称 string, 图片 int) bool {
	r, _, _ := xList_SetItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(图片))
	return r != 0
}

// 列表_置项指数值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func X列表_置项指数值(元素句柄 int, 项索引 int, 列索引 int, 值 int) bool {
	r, _, _ := xList_SetItemInt.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(值))
	return r != 0
}

// 列表_置项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// nValue:.
func X列表_置项整数值EX(元素句柄 int, 项索引 int, 名称 string, 值 int) bool {
	r, _, _ := xList_SetItemIntEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(值))
	return r != 0
}

// 列表_置项浮点值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func X列表_置项浮点值(元素句柄 int, 项索引 int, 列索引 int, 值 float32) bool {
	r, _, _ := xList_SetItemFloat.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), 炫彩工具类.Float32Ptr(值))
	return r != 0
}

// 列表_置项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func X列表_置项浮点值EX(元素句柄 int, 项索引 int, 名称 string, 值 float32) bool {
	r, _, _ := xList_SetItemFloatEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), 炫彩工具类.Float32Ptr(值))
	return r != 0
}

// 列表_取项文本.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func X列表_取项文本(元素句柄 int, 项索引 int, 列索引 int) string {
	r, _, _ := xList_GetItemText.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表_取项文本扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
func X列表_取项文本EX(元素句柄 int, 项索引 int, 名称 string) string {
	r, _, _ := xList_GetItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表_取项图片.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
func X列表_取项图片(元素句柄 int, 项索引 int, 列索引 int) int {
	r, _, _ := xList_GetItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引))
	return int(r)
}

// 列表_取项图片扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
func X列表_取项图片EX(元素句柄 int, 项索引 int, 名称 string) int {
	r, _, _ := xList_GetItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称))
	return int(r)
}

// 列表_取项整数值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func X列表_取项整数值(元素句柄 int, 项索引 int, 列索引 int, 返回值指针 *int32) bool {
	r, _, _ := xList_GetItemInt.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(unsafe.Pointer(返回值指针)))
	return r != 0
}

// 列表_取项整数值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func X列表_取项整数值EX(元素句柄 int, 项索引 int, 名称 string, 返回值指针 *int32) bool {
	r, _, _ := xList_GetItemIntEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(unsafe.Pointer(返回值指针)))
	return r != 0
}

// 列表_取项浮点值.
//
// hEle:.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func X列表_取项浮点值(元素句柄 int, 项索引 int, 列索引 int, 返回值指针 *float32) bool {
	r, _, _ := xList_GetItemFloat.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(unsafe.Pointer(返回值指针)))
	return r != 0
}

// 列表_取项浮点值扩展.
//
// hEle:.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func X列表_取项浮点值EX(元素句柄 int, 项索引 int, 名称 string, 返回值指针 *float32) bool {
	r, _, _ := xList_GetItemFloatEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(名称), uintptr(unsafe.Pointer(返回值指针)))
	return r != 0
}

// 列表_删除项.
//
// hEle:.
//
// iItem:.
func X列表_删除项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xList_DeleteItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 列表_删除项扩展.
//
// hEle:.
//
// iItem:.
//
// nCount:.
func X列表_删除项EX(元素句柄 int, 项索引 int, 数量 int) bool {
	r, _, _ := xList_DeleteItemEx.Call(uintptr(元素句柄), uintptr(项索引), uintptr(数量))
	return r != 0
}

// 列表_删除项全部.
//
// hEle:.
func X列表_删除项全部(元素句柄 int) int {
	r, _, _ := xList_DeleteItemAll.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_删除列全部AD.
//
// hEle:.
func X列表_删除列全部AD(元素句柄 int) int {
	r, _, _ := xList_DeleteColumnAll_AD.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_取项数量AD.
//
// hEle:.
func X列表_取项数量AD(元素句柄 int) int {
	r, _, _ := xList_GetCount_AD.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_取列数量AD.
//
// hEle:.
func X列表_取列数量AD(元素句柄 int) int {
	r, _, _ := xList_GetCountColumn_AD.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_置分割线颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func X列表_置分割线颜色(元素句柄 int, ABGR颜色值 int) int {
	r, _, _ := xList_SetSplitLineColor.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 列表_置项高度.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func X列表_置项高度(元素句柄 int, 行索引 int, 高度, 选中时高度 int32) int {
	r, _, _ := xList_SetItemHeight.Call(uintptr(元素句柄), uintptr(行索引), uintptr(高度), uintptr(选中时高度))
	return int(r)
}

// 列表_取项高度.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func X列表_取项高度(元素句柄 int, 行索引 int, 高度, 选中时高度 *int32) int {
	r, _, _ := xList_GetItemHeight.Call(uintptr(元素句柄), uintptr(行索引), uintptr(unsafe.Pointer(高度)), uintptr(unsafe.Pointer(选中时高度)))
	return int(r)
}

// 列表_置拖动矩形颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
//
// width: 线宽度.
func X列表_置拖动矩形颜色(元素句柄 int, ABGR颜色值, 线宽度 int) int {
	r, _, _ := xList_SetDragRectColor.Call(uintptr(元素句柄), uintptr(ABGR颜色值), uintptr(线宽度))
	return int(r)
}

// 列表_取项模板. 返回列表项模板句柄.
//
// hEle: 元素句柄.
func X列表_取项模板(元素句柄 int) int {
	r, _, _ := xList_GetItemTemplate.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_取项模板列表头. 返回列表头项模板句柄.
//
// hEle: 元素句柄.
func X列表_取项模板列表头(元素句柄 int) int {
	r, _, _ := xList_GetItemTemplateHeader.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_刷新项数据列表头.
//
// hEle: 元素句柄.
func X列表_刷新项数据列表头(元素句柄 int) int {
	r, _, _ := xList_RefreshDataHeader.Call(uintptr(元素句柄))
	return int(r)
}

// 列表_置项模板从内存.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func X列表_置项模板从内存(元素句柄 int, 模板数据 []byte) bool {
	r, _, _ := xList_SetItemTemplateXMLFromMem.Call(uintptr(元素句柄), 炫彩工具类.ByteSliceDataPtr(&模板数据), uintptr(len(模板数据)))
	return r != 0
}

// 列表_置项模板从资源ZIP. 从RC资源ZIP加载.
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
func X列表_置项模板从资源ZIP(元素句柄 int, RC资源ID int32, 项模板文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	r, _, _ := xList_SetItemTemplateXMLFromZipRes.Call(uintptr(元素句柄), uintptr(RC资源ID), 炫彩工具类.StrPtr(项模板文件名), 炫彩工具类.StrPtr(zip密码), 模块句柄)
	return r != 0
}
