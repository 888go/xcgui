package 炫彩基类

import (
	"unsafe"

	"github.com/888go/xcgui/common"

	"github.com/888go/xcgui/xcc"
)

// XTree_Create 列表树_创建, 创建树元素, 返回元素句柄.
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

// ff:列表树_创建
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X列表树_创建(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xTree_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// XTree_CreateEx 列表树_创建Ex, 创建树元素, 使用内置项模板, 自动创建数据适配器, 返回元素句柄.
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

// ff:列表树_创建Ex
// col_extend_count:列数量
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X列表树_创建Ex(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄, 列数量 int32) int {
	r, _, _ := xTree_CreateEx.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄), uintptr(列数量))
	return int(r)
}

// 列表树_启用拖动项, 启用拖动项功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.

// ff:列表树_启用拖动项
// bEnable:是否启用
// hEle:元素句柄
func X列表树_启用拖动项(元素句柄 int, 是否启用 bool) {
	xTree_EnableDragItem.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
}

// 列表树_启用连接线, 启用或禁用显示项的连接线.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
//
// bSolid: 实线或虚线; TRUE: 实线, FALSE: 虚线.

// ff:列表树_启用连接线
// bSolid:实线或虚线
// bEnable:是否启用
// hEle:元素句柄
func X列表树_启用连接线(元素句柄 int, 是否启用 bool, 实线或虚线 bool) {
	xTree_EnableConnectLine.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用), 炫彩工具类.BoolPtr(实线或虚线))
}

// 列表树_启用展开, 启动或关闭默认展开功能, 如果开启新插入的项将自动展开.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.

// ff:列表树_启用展开
// bEnable:是否启用
// hEle:元素句柄
func X列表树_启用展开(元素句柄 int, 是否启用 bool) {
	xTree_EnableExpand.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
}

// 列表树_启用模板复用.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.

// ff:列表树_启用模板复用
// bEnable:是否启用
// hEle:元素句柄
func X列表树_启用模板复用(元素句柄 int, 是否启用 bool) {
	xTree_EnableTemplateReuse.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
}

// 列表树_置连接线颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色.

// ff:列表树_置连接线颜色
// color:ABGR颜色
// hEle:元素句柄
func X列表树_置连接线颜色(元素句柄 int, ABGR颜色 int) {
	xTree_SetConnectLineColor.Call(uintptr(元素句柄), uintptr(ABGR颜色))
}

// 列表树_置展开按钮大小, 设置展开按钮占用空间大小.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.

// ff:列表树_置展开按钮大小
// nHeight:高度
// nWidth:宽度
// hEle:元素句柄
func X列表树_置展开按钮大小(元素句柄 int, 宽度, 高度 int32) {
	xTree_SetExpandButtonSize.Call(uintptr(元素句柄), uintptr(宽度), uintptr(高度))
}

// 列表树_置连接线长度, 设置连线绘制长度, 展开按钮与项内容之间的连线.
//
// hEle: 元素句柄.
//
// nLength: 连线绘制长度.

// ff:列表树_置连接线长度
// nLength:连线绘制长度
// hEle:元素句柄
func X列表树_置连接线长度(元素句柄 int, 连线绘制长度 int32) {
	xTree_SetConnectLineLength.Call(uintptr(元素句柄), uintptr(连线绘制长度))
}

// 列表树_置拖动项插入位置颜色, 设置拖动项插入位置颜色提示.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色.

// ff:列表树_置拖动项插入位置颜色
// color:ABGR颜色
// hEle:元素句柄
func X列表树_置拖动项插入位置颜色(元素句柄 int, ABGR颜色 int) {
	xTree_SetDragInsertPositionColor.Call(uintptr(元素句柄), uintptr(ABGR颜色))
}

// 列表树_置项模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 文件名.

// ff:列表树_置项模板文件
// pXmlFile:文件名
// hEle:元素句柄
func X列表树_置项模板文件(元素句柄 int, 文件名 string) bool {
	r, _, _ := xTree_SetItemTemplateXML.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文件名))
	return r != 0
}

// 列表树_置选择项模板文件, 设置项模板文件, 项选中状态.
//
// hEle: 元素句柄.
//
// pXmlFile: 文件名.

// ff:列表树_置选择项模板文件
// pXmlFile:文件名
// hEle:元素句柄
func X列表树_置选择项模板文件(元素句柄 int, 文件名 string) bool {
	r, _, _ := xTree_SetItemTemplateXMLSel.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文件名))
	return r != 0
}

// 列表树_置项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.

// ff:列表树_置项模板
// hTemp:模板句柄
// hEle:元素句柄
func X列表树_置项模板(元素句柄 int, 模板句柄 int) bool {
	r, _, _ := xTree_SetItemTemplate.Call(uintptr(元素句柄), uintptr(模板句柄))
	return r != 0
}

// 列表树_置选择项模板, 设置列表项模板, 项选中状态.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.

// ff:列表树_置选择项模板
// hTemp:模板句柄
// hEle:元素句柄
func X列表树_置选择项模板(元素句柄 int, 模板句柄 int) bool {
	r, _, _ := xTree_SetItemTemplateSel.Call(uintptr(元素句柄), uintptr(模板句柄))
	return r != 0
}

// 列表树_置项模板从字符串, 设置项模板文件.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.

// ff:列表树_置项模板从字符串
// pStringXML:字符串
// hEle:元素句柄
func X列表树_置项模板从字符串(元素句柄 int, 字符串 string) bool {
	r, _, _ := xTree_SetItemTemplateXMLFromString.Call(uintptr(元素句柄), W2A(字符串))
	return r != 0
}

// 列表树_置选择项模板从字符串, 设置项模板文件, 项选中状态.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.

// ff:列表树_置选择项模板从字符串
// pStringXML:字符串
// hEle:元素句柄
func X列表树_置选择项模板从字符串(元素句柄 int, 字符串 string) bool {
	r, _, _ := xTree_SetItemTemplateXMLSelFromString.Call(uintptr(元素句柄), W2A(字符串))
	return r != 0
}

// 列表树_置项背景绘制标志, 设置是否绘制指定状态下项的背景.
//
// hEle: 元素句柄.
//
// nFlags: 标志位: xcc.List_DrawItemBk_Flag_.

// ff:列表树_置项背景绘制标志
// nFlags:标志位
// hEle:元素句柄
func X列表树_置项背景绘制标志(元素句柄 int, 标志位 炫彩常量类.List_DrawItemBk_Flag_) {
	xTree_SetDrawItemBkFlags.Call(uintptr(元素句柄), uintptr(标志位))
}

// 列表树_置项数据, 设置项用户数据.
//
// hEle: 元素句柄.
//
// nID: 项ID.
//
// nUserData: 用户数据.

// ff:列表树_置项数据
// nUserData:用户数据
// nID:项ID
// hEle:元素句柄
func X列表树_置项数据(元素句柄 int, 项ID int32, 用户数据 int) bool {
	r, _, _ := xTree_SetItemData.Call(uintptr(元素句柄), uintptr(项ID), uintptr(用户数据))
	return r != 0
}

// 列表树_取项数据, 获取项用户数据.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_取项数据
// nID:项ID
// hEle:元素句柄
func X列表树_取项数据(元素句柄 int, 项ID int32) int {
	r, _, _ := xTree_GetItemData.Call(uintptr(元素句柄), uintptr(项ID))
	return int(r)
}

// 列表树_置选择项.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_置选择项
// nID:项ID
// hEle:元素句柄
func X列表树_置选择项(元素句柄 int, 项ID int32) bool {
	r, _, _ := xTree_SetSelectItem.Call(uintptr(元素句柄), uintptr(项ID))
	return r != 0
}

// 列表树_取选择项, 返回项ID.
//
// hEle: 元素句柄.

// ff:列表树_取选择项
// hEle:元素句柄
func X列表树_取选择项(元素句柄 int) int32 {
	r, _, _ := xTree_GetSelectItem.Call(uintptr(元素句柄))
	return int32(r)
}

// 列表树_可视指定项, 滚动视图让指定项可见.
//
// hEle: 元素句柄.
//
// nID: 项索引.

// ff:列表树_可视指定项
// nID:项ID
// hEle:元素句柄
func X列表树_可视指定项(元素句柄 int, 项ID int32) {
	xTree_VisibleItem.Call(uintptr(元素句柄), uintptr(项ID))
}

// 列表树_判断展开.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_判断展开
// nID:项ID
// hEle:元素句柄
func X列表树_判断展开(元素句柄 int, 项ID int32) bool {
	r, _, _ := xTree_IsExpand.Call(uintptr(元素句柄), uintptr(项ID))
	return r != 0
}

// 列表树_展开项, 判断项是否展开.
//
// hEle: 元素句柄.
//
// nID: 项ID.
//
// bExpand: 是否展开.

// ff:列表树_展开项
// bExpand:是否展开
// nID:项ID
// hEle:元素句柄
func X列表树_展开项(元素句柄 int, 项ID int32, 是否展开 bool) bool {
	r, _, _ := xTree_ExpandItem.Call(uintptr(元素句柄), uintptr(项ID), 炫彩工具类.BoolPtr(是否展开))
	return r != 0
}

// 列表树_展开全部子项, 展开所有的子项.
//
// hEle: 元素句柄.
//
// nID: 项ID.
//
// bExpand: 是否展开.

// ff:列表树_展开全部子项
// bExpand:是否展开
// nID:项ID
// hEle:元素句柄
func X列表树_展开全部子项(元素句柄 int, 项ID int32, 是否展开 bool) bool {
	r, _, _ := xTree_ExpandAllChildItem.Call(uintptr(元素句柄), uintptr(项ID), 炫彩工具类.BoolPtr(是否展开))
	return r != 0
}

// 列表树_测试点击项, 检测坐标点所在项, 返回项ID.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.

// ff:列表树_测试点击项
// pPt:坐标点
// hEle:元素句柄
func X列表树_测试点击项(元素句柄 int, 坐标点 *POINT) int32 {
	r, _, _ := xTree_HitTest.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)))
	return int32(r)
}

// 列表树_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移坐标, 返回项ID.
//
// hEle: 元素句柄.
//
// pPt: 坐标点.

// ff:列表树_测试点击项EX
// pPt:坐标点
// hEle:元素句柄
func X列表树_测试点击项EX(元素句柄 int, 坐标点 *POINT) int32 {
	r, _, _ := xTree_HitTestOffset.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标点)))
	return int32(r)
}

// 列表树_取第一个子项, 获取第一个子项. 成功返回项ID, 失败返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_取第一个子项
// nID:项ID
// hEle:元素句柄
func X列表树_取第一个子项(元素句柄 int, 项ID int32) int32 {
	r, _, _ := xTree_GetFirstChildItem.Call(uintptr(元素句柄), uintptr(项ID))
	return int32(r)
}

// 列表树_取末尾子项, 获取末尾子项. 成功返回项ID, 失败返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_取末尾子项
// nID:项ID
// hEle:元素句柄
func X列表树_取末尾子项(元素句柄 int, 项ID int32) int32 {
	r, _, _ := xTree_GetEndChildItem.Call(uintptr(元素句柄), uintptr(项ID))
	return int32(r)
}

// 列表树_取上一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_取上一个兄弟项
// nID:项ID
// hEle:元素句柄
func X列表树_取上一个兄弟项(元素句柄 int, 项ID int32) int32 {
	r, _, _ := xTree_GetPrevSiblingItem.Call(uintptr(元素句柄), uintptr(项ID))
	return int32(r)
}

// 列表树_取下一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_取下一个兄弟项
// nID:项ID
// hEle:元素句柄
func X列表树_取下一个兄弟项(元素句柄 int, 项ID int32) int32 {
	r, _, _ := xTree_GetNextSiblingItem.Call(uintptr(元素句柄), uintptr(项ID))
	return int32(r)
}

// 列表树_取父项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_取父项
// nID:项ID
// hEle:元素句柄
func X列表树_取父项(元素句柄 int, 项ID int32) int32 {
	r, _, _ := xTree_GetParentItem.Call(uintptr(元素句柄), uintptr(项ID))
	return int32(r)
}

// 列表树_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// hEle: 元素句柄.

// ff:列表树_创建数据适配器
// hEle:元素句柄
func X列表树_创建数据适配器(元素句柄 int) int {
	r, _, _ := xTree_CreateAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 列表树_绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 数据适配器句柄, XAdTree.

// ff:列表树_绑定数据适配器
// hAdapter:数据适配器句柄
// hEle:元素句柄
func X列表树_绑定数据适配器(元素句柄 int, 数据适配器句柄 int) {
	xTree_BindAdapter.Call(uintptr(元素句柄), uintptr(数据适配器句柄))
}

// 列表树_取数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.

// ff:列表树_取数据适配器
// hEle:元素句柄
func X列表树_取数据适配器(元素句柄 int) int {
	r, _, _ := xTree_GetAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 列表树_刷新数据, 刷新所有项模板, 以便更新UI.
//
// hEle: 元素句柄.

// ff:列表树_刷新数据
// hEle:元素句柄
func X列表树_刷新数据(元素句柄 int) {
	xTree_RefreshData.Call(uintptr(元素句柄))
}

// 列表树_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// hEle: 元素句柄.
//
// nID: 项ID.

// ff:列表树_刷新指定项
// nID:项ID
// hEle:元素句柄
func X列表树_刷新指定项(元素句柄 int, 项ID int32) {
	xTree_RefreshItem.Call(uintptr(元素句柄), uintptr(项ID))
}

// 列表树_置缩进, 设置缩进大小.
//
// hEle: 元素句柄.
//
// nWidth: 缩进宽度.

// ff:列表树_置缩进
// nWidth:缩进宽度
// hEle:元素句柄
func X列表树_置缩进(元素句柄 int, 缩进宽度 int32) {
	xTree_SetIndentation.Call(uintptr(元素句柄), uintptr(缩进宽度))
}

// 列表树_取缩进, 返回缩进值大小.
//
// hEle: 元素句柄.

// ff:列表树_取缩进
// hEle:元素句柄
func X列表树_取缩进(元素句柄 int) int32 {
	r, _, _ := xTree_GetIndentation.Call(uintptr(元素句柄))
	return int32(r)
}

// 列表树_置项默认高度.
//
// hEle: 元素句柄.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.

// ff:列表树_置项默认高度
// nSelHeight:选中时高度
// nHeight:高度
// hEle:元素句柄
func X列表树_置项默认高度(元素句柄 int, 高度, 选中时高度 int32) {
	xTree_SetItemHeightDefault.Call(uintptr(元素句柄), uintptr(高度), uintptr(选中时高度))
}

// 列表树_取项默认高度.
//
// hEle: 元素句柄.
//
// pHeight: 接收返回高度.
//
// pSelHeight: 接收返回值, 当项选中时的高度.

// ff:列表树_取项默认高度
// pSelHeight:接收返回值
// pHeight:接收返回高度
// hEle:元素句柄
func X列表树_取项默认高度(元素句柄 int, 接收返回高度, 接收返回值 *int32) {
	xTree_GetItemHeightDefault.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(接收返回高度)), uintptr(unsafe.Pointer(接收返回值)))
}

// 列表树_置项高度.
//
// hEle: 元素句柄.
//
// nID: 项ID.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.

// ff:列表树_置项高度
// nSelHeight:选中时高度
// nHeight:高度
// nID:项ID
// hEle:元素句柄
func X列表树_置项高度(元素句柄 int, 项ID, 高度, 选中时高度 int32) {
	xTree_SetItemHeight.Call(uintptr(元素句柄), uintptr(项ID), uintptr(高度), uintptr(选中时高度))
}

// 列表树_取项高度.
//
// hEle: 元素句柄.
//
// nID: 项ID.
//
// pHeight: 接收返回高度.
//
// pSelHeight: 接收返回值, 当项选中时的高度.

// ff:列表树_取项高度
// pSelHeight:接收返回值
// pHeight:接收返回高度
// nID:项ID
// hEle:元素句柄
func X列表树_取项高度(元素句柄 int, 项ID int32, 接收返回高度, 接收返回值 *int32) {
	xTree_GetItemHeight.Call(uintptr(元素句柄), uintptr(项ID), uintptr(unsafe.Pointer(接收返回高度)), uintptr(unsafe.Pointer(接收返回值)))
}

// 列表树_置行间距.
//
// hEle: 元素句柄.
//
// nSpace: 行间隔大小.

// ff:列表树_置行间距
// nSpace:行间隔大小
// hEle:元素句柄
func X列表树_置行间距(元素句柄 int, 行间隔大小 int32) {
	xTree_SetRowSpace.Call(uintptr(元素句柄), uintptr(行间隔大小))
}

// 列表树_取行间距.
//
// hEle: 元素句柄.

// ff:列表树_取行间距
// hEle:元素句柄
func X列表树_取行间距(元素句柄 int) int32 {
	r, _, _ := xTree_GetRowSpace.Call(uintptr(元素句柄))
	return int32(r)
}

// 列表树_移动项, 移动项的位置.
//
// hEle: 元素句柄.
//
// nMoveItem: 要移动的项ID.
//
// nDestItem: 目标项ID, 参照位置.
//
// nFlag: 0:目标前面, 1:目标后面, 2:目标子项首, 3:目标子项尾.

// ff:列表树_移动项
// nFlag:移动方式
// nDestItem:目标项ID
// nMoveItem:要移动的项ID
// hEle:元素句柄
func X列表树_移动项(元素句柄 int, 要移动的项ID, 目标项ID, 移动方式 int32) bool {
	r, _, _ := xTree_MoveItem.Call(uintptr(元素句柄), uintptr(要移动的项ID), uintptr(目标项ID), uintptr(移动方式))
	return r != 0
}

// 列表树_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// hEle: 元素句柄.
//
// nID: 树项ID.
//
// nTempItemID: 模板项ID.

// ff:列表树_取模板对象
// nTempItemID:模板项ID
// nID:项ID
// hEle:元素句柄
func X列表树_取模板对象(元素句柄 int, 项ID, 模板项ID int32) int {
	r, _, _ := xTree_GetTemplateObject.Call(uintptr(元素句柄), uintptr(项ID), uintptr(模板项ID))
	return int(r)
}

// 列表树_取对象所在项, 获取当前对象所在模板实例, 属于列表树中哪一个项. 成功返回项ID, 否则返回XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hXCGUI: 对象句柄.

// ff:列表树_取对象所在项
// hXCGUI:对象句柄
// hEle:元素句柄
func X列表树_取对象所在项(元素句柄 int, 对象句柄 int) int32 {
	r, _, _ := xTree_GetItemIDFromHXCGUI.Call(uintptr(元素句柄), uintptr(对象句柄))
	return int32(r)
}

// 列表树_插入项文本.
//
// hEle:.
//
// pValue:.
//
// nParentID:.
//
// insertID:.

// ff:列表树_插入项文本
// insertID:插入位置ID
// nParentID:父ID
// pValue:值
// hEle:元素句柄
func X列表树_插入项文本(元素句柄 int, 值 string, 父ID, 插入位置ID int32) int32 {
	r, _, _ := xTree_InsertItemText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(值), uintptr(父ID), uintptr(插入位置ID))
	return int32(r)
}

// 列表树_插入项文本扩展.
//
// hEle:.
//
// pName:.
//
// pValue:.
//
// nParentID:.
//
// insertID:.

// ff:列表树_插入项文本EX
// insertID:插入位置ID
// nParentID:父ID
// pValue:值
// pName:名称
// hEle:元素句柄
func X列表树_插入项文本EX(元素句柄 int, 名称 string, 值 string, 父ID, 插入位置ID int32) int32 {
	r, _, _ := xTree_InsertItemTextEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(值), uintptr(父ID), uintptr(插入位置ID))
	return int32(r)
}

// 列表树_插入项图片.
//
// hEle:.
//
// hImage:.
//
// nParentID:.
//
// insertID:.

// ff:列表树_插入项图片
// insertID:插入位置ID
// nParentID:父ID
// hImage:图片句柄
// hEle:元素句柄
func X列表树_插入项图片(元素句柄 int, 图片句柄 int, 父ID, 插入位置ID int32) int32 {
	r, _, _ := xTree_InsertItemImage.Call(uintptr(元素句柄), uintptr(图片句柄), uintptr(父ID), uintptr(插入位置ID))
	return int32(r)
}

// 列表树_插入项图片扩展.
//
// hEle:.
//
// pName:.
//
// hImage:.
//
// nParentID:.
//
// insertID:.

// ff:列表树_插入项图片EX
// insertID:插入位置ID
// nParentID:父ID
// hImage:图片句柄
// pName:名称
// hEle:元素句柄
func X列表树_插入项图片EX(元素句柄 int, 名称 string, 图片句柄 int, 父ID, 插入位置ID int32) int32 {
	r, _, _ := xTree_InsertItemImageEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(名称), uintptr(图片句柄), uintptr(父ID), uintptr(插入位置ID))
	return int32(r)
}

// 列表树_取项数量.
//
// hEle:.

// ff:列表树_取项数量
// hEle:元素句柄
func X列表树_取项数量(元素句柄 int) int32 {
	r, _, _ := xTree_GetCount.Call(uintptr(元素句柄))
	return int32(r)
}

// 列表树_取列数量.
//
// hEle:.

// ff:列表树_取列数量
// hEle:元素句柄
func X列表树_取列数量(元素句柄 int) int32 {
	r, _, _ := xTree_GetCountColumn.Call(uintptr(元素句柄))
	return int32(r)
}

// 列表树_置项文本.
//
// hEle:.
//
// nID:.
//
// iColumn:.
//
// pValue:.

// ff:列表树_置项文本
// pValue:值
// iColumn:列索引
// nID:项ID
// hEle:元素句柄
func X列表树_置项文本(元素句柄 int, 项ID, 列索引 int32, 值 string) bool {
	r, _, _ := xTree_SetItemText.Call(uintptr(元素句柄), uintptr(项ID), uintptr(列索引), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 列表树_置项文本扩展.
//
// hEle:.
//
// nID:.
//
// pName:.
//
// pValue:.

// ff:列表树_置项文本EX
// pValue:值
// pName:名称
// nID:项ID
// hEle:元素句柄
func X列表树_置项文本EX(元素句柄 int, 项ID int32, 名称 string, 值 string) bool {
	r, _, _ := xTree_SetItemTextEx.Call(uintptr(元素句柄), uintptr(项ID), 炫彩工具类.StrPtr(名称), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 列表树_置项图片.
//
// hEle:.
//
// nID:.
//
// iColumn:.
//
// hImage:.

// ff:列表树_置项图片
// hImage:图片句柄
// iColumn:列索引
// nID:项ID
// hEle:元素句柄
func X列表树_置项图片(元素句柄 int, 项ID, 列索引 int32, 图片句柄 int) bool {
	r, _, _ := xTree_SetItemImage.Call(uintptr(元素句柄), uintptr(项ID), uintptr(列索引), uintptr(图片句柄))
	return r != 0
}

// 列表树_置项图片扩展.
//
// hEle:.
//
// nID:.
//
// pName:.
//
// hImage:.

// ff:列表树_置项图片EX
// hImage:图片句柄
// pName:名称
// nID:项ID
// hEle:元素句柄
func X列表树_置项图片EX(元素句柄 int, 项ID int32, 名称 string, 图片句柄 int) bool {
	r, _, _ := xTree_SetItemImageEx.Call(uintptr(元素句柄), uintptr(项ID), 炫彩工具类.StrPtr(名称), uintptr(图片句柄))
	return r != 0
}

// 列表树_取项文本.
//
// hEle:.
//
// nID:.
//
// iColumn:.

// ff:列表树_取项文本
// iColumn:列索引
// nID:项ID
// hEle:元素句柄
func X列表树_取项文本(元素句柄 int, 项ID, 列索引 int32) string {
	r, _, _ := xTree_GetItemText.Call(uintptr(元素句柄), uintptr(项ID), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表树_取项文本扩展.
//
// hEle:.
//
// nID:.
//
// pName:.

// ff:列表树_取项文本EX
// pName:名称
// nID:项ID
// hEle:元素句柄
func X列表树_取项文本EX(元素句柄 int, 项ID int32, 名称 string) string {
	r, _, _ := xTree_GetItemTextEx.Call(uintptr(元素句柄), uintptr(项ID), 炫彩工具类.StrPtr(名称))
	return 炫彩工具类.UintPtrToString(r)
}

// 列表树_取项图片.
//
// hEle:.
//
// nID:.
//
// iColumn:.

// ff:列表树_取项图片
// iColumn:列索引
// nID:项ID
// hEle:元素句柄
func X列表树_取项图片(元素句柄 int, 项ID, 列索引 int32) int {
	r, _, _ := xTree_GetItemImage.Call(uintptr(元素句柄), uintptr(项ID), uintptr(列索引))
	return int(r)
}

// 列表树_取项图片扩展.
//
// hEle:.
//
// nID:.
//
// pName:.

// ff:列表树_取项图片EX
// pName:名称
// nID:项ID
// hEle:元素句柄
func X列表树_取项图片EX(元素句柄 int, 项ID int32, 名称 string) int {
	r, _, _ := xTree_GetItemImageEx.Call(uintptr(元素句柄), uintptr(项ID), 炫彩工具类.StrPtr(名称))
	return int(r)
}

// 列表树_删除项.
//
// hEle:.
//
// nID:.

// ff:列表树_删除项
// nID:项ID
// hEle:元素句柄
func X列表树_删除项(元素句柄 int, 项ID int32) bool {
	r, _, _ := xTree_DeleteItem.Call(uintptr(元素句柄), uintptr(项ID))
	return r != 0
}

// 列表树_删除全部项.
//
// hEle:.

// ff:列表树_删除全部项
// hEle:元素句柄
func X列表树_删除全部项(元素句柄 int) {
	xTree_DeleteItemAll.Call(uintptr(元素句柄))
}

// 列表树_删除列全部.
//
// hEle:.

// ff:列表树_删除列全部
// hEle:元素句柄
func X列表树_删除列全部(元素句柄 int) {
	xTree_DeleteColumnAll.Call(uintptr(元素句柄))
}

// 列表树_置分割线颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.

// ff:列表树_置分割线颜色
// color:ABGR颜色值
// hEle:元素句柄
func X列表树_置分割线颜色(元素句柄 int, ABGR颜色值 int) {
	xTree_SetSplitLineColor.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
}

// 列表树_置项模板从内存.
//
// hEle: 元素句柄.
//
// data: 模板数据.

// ff:列表树_置项模板从内存
// data:模板数据
// hEle:元素句柄
func X列表树_置项模板从内存(元素句柄 int, 模板数据 []byte) bool {
	r, _, _ := xTree_SetItemTemplateXMLFromMem.Call(uintptr(元素句柄), 炫彩工具类.ByteSliceDataPtr(&模板数据), uintptr(len(模板数据)))
	return r != 0
}

// 列表树_置项模板从资源ZIP.
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

// ff:列表树_置项模板从资源ZIP
// hModule:模块句柄
// pPassword:zip密码
// pFileName:文件名
// id:RC资源ID
// hEle:元素句柄
func X列表树_置项模板从资源ZIP(元素句柄 int, RC资源ID int32, 文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	r, _, _ := xTree_SetItemTemplateXMLFromZipRes.Call(uintptr(元素句柄), uintptr(RC资源ID), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(zip密码), 模块句柄)
	return r != 0
}

// 列表树_取项模板, 返回项模板句柄.
//
// hEle: 元素句柄.

// ff:列表树_取项模板
// hEle:元素句柄
func X列表树_取项模板(元素句柄 int) int {
	r, _, _ := xTree_GetItemTemplate.Call(uintptr(元素句柄))
	return int(r)
}
