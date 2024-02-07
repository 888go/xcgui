package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
	
	"github.com/888go/xcgui/xcc"
)

// 组合框_创建, 返回元素句柄.
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
func X组合框_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xComboBox_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 组合框_置选择项.
//
// hEle: 元素句柄.
//
// iIndex: 项索引.
func X组合框_置选择项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xComboBox_SetSelItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 组合框_创建数据适配器, 返回数据适配器句柄.
//
// hEle: 元素句柄.
func X组合框_创建数据适配器(元素句柄 int) int {
	r, _, _ := xComboBox_CreateAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 组合框_绑定数据适配器.
//
// hEle: 元素句柄.
//
// hAdapter: 适配器句柄.
func X组合框_绑定数据适配器(元素句柄 int, 适配器句柄 int) int {
	r, _, _ := xComboBox_BindAdapter.Call(uintptr(元素句柄), uintptr(适配器句柄))
	return int(r)
}

// 组合框_取数据适配器, 获取绑定的数据适配器.
//
// hEle: 元素句柄.
func X组合框_取数据适配器(元素句柄 int) int {
	r, _, _ := xComboBox_GetAdapter.Call(uintptr(元素句柄))
	return int(r)
}

// 组合框_置绑定名称.
//
// hEle: 元素句柄.
//
// pName: 字段名.
func X组合框_置绑定名称(元素句柄 int, 字段名 string) int {
	r, _, _ := xComboBox_SetBindName.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字段名))
	return int(r)
}

// 组合框_取下拉按钮坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.
func X组合框_取下拉按钮坐标(元素句柄 int, 坐标 *RECT) int {
	r, _, _ := xComboBox_GetButtonRect.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 组合框_置下拉按钮大小.
//
// hEle: 元素句柄.
//
// size: 大小.
func X组合框_置下拉按钮大小(元素句柄 int, 大小 int) int {
	r, _, _ := xComboBox_SetButtonSize.Call(uintptr(元素句柄), uintptr(大小))
	return int(r)
}

// 组合框_置下拉列表高度.
//
// hEle: 元素句柄.
//
// height: 高度, -1自动计算高度.
func X组合框_置下拉列表高度(元素句柄 int, 高度 int) int {
	r, _, _ := xComboBox_SetDropHeight.Call(uintptr(元素句柄), uintptr(高度))
	return int(r)
}

// 组合框_取下拉列表高度.
//
// hEle: 元素句柄.
func X组合框_取下拉列表高度(元素句柄 int) int {
	r, _, _ := xComboBox_GetDropHeight.Call(uintptr(元素句柄))
	return int(r)
}

// 组合框_置项模板, 设置下拉列表项模板文件.
//
// hEle: 元素句柄.
//
// pXmlFile: 项模板文件.
func X组合框_置项模板(元素句柄 int, 项模板文件 string) int {
	r, _, _ := xComboBox_SetItemTemplateXML.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(项模板文件))
	return int(r)
}

// 组合框_置项模板从字符串, 设置下拉列表项模板.
//
// hEle: 元素句柄.
//
// pStringXML: 字符串.
func X组合框_置项模板从字符串(元素句柄 int, 字符串 string) int {
	r, _, _ := xComboBox_SetItemTemplateXMLFromString.Call(uintptr(元素句柄), W2A(字符串))
	return int(r)
}

// 组合框_启用绘制下拉按钮, 是否绘制下拉按钮.
//
// hEle: 元素句柄.
//
// bEnable: 是否绘制.
func X组合框_启用绘制下拉按钮(元素句柄 int, 是否绘制 bool) int {
	r, _, _ := xComboBox_EnableDrawButton.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否绘制))
	return int(r)
}

// 组合框_启用编辑, 启用可编辑显示的文本内容.
//
// hEle: 元素句柄.
//
// bEdit: TRUE可编辑.
func X组合框_启用编辑(元素句柄 int, TRUE可编辑 bool) int {
	r, _, _ := xComboBox_EnableEdit.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(TRUE可编辑))
	return int(r)
}

// 组合框_启用下拉列表高度固定大小, 启用/关闭下拉列表高度固定大小.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X组合框_启用下拉列表高度固定大小(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xComboBox_EnableDropHeightFixed.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 组合框_取选择项, 获取组合框下拉列表中选择项索引.
//
// hEle: 元素句柄.
func X组合框_取选择项(元素句柄 int) int {
	r, _, _ := xComboBox_GetSelItem.Call(uintptr(元素句柄))
	return int(r)
}

// 组合框_取状态, 返回: ComboBox_State_.
//
// hEle: 元素句柄.
func X组合框_取状态(元素句柄 int) 炫彩常量类.ComboBox_State_ {
	r, _, _ := xComboBox_GetState.Call(uintptr(元素句柄))
	return 炫彩常量类.ComboBox_State_(r)
}

// 组合框_添加项文本, 返回项索引.
//
// hEle: 元素句柄.
//
// pText:.
func X组合框_添加项文本(元素句柄 int, 文本 string) int {
	r, _, _ := xComboBox_AddItemText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 组合框_添加项文本扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// pName: 字段名.
//
// pText: 文本.
func X组合框_添加项文本EX(元素句柄 int, 字段名 string, 文本 string) int {
	r, _, _ := xComboBox_AddItemTextEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字段名), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 组合框_添加项图片, 返回项索引.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func X组合框_添加项图片(元素句柄 int, 图片句柄 int) int {
	r, _, _ := xComboBox_AddItemImage.Call(uintptr(元素句柄), uintptr(图片句柄))
	return int(r)
}

// 组合框_添加项图片扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func X组合框_添加项图片EX(元素句柄 int, 字段名 string, 图片句柄 int) int {
	r, _, _ := xComboBox_AddItemImageEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字段名), uintptr(图片句柄))
	return int(r)
}

// 组合框_插入项文本, 返回项索引.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pText: 文本.
func X组合框_插入项文本(元素句柄 int, 项索引 int, 文本 string) int {
	r, _, _ := xComboBox_InsertItemText.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 组合框_插入项文本扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func X组合框_插入项文本EX(元素句柄 int, 项索引 int, 字段名 string, 文本 string) int {
	r, _, _ := xComboBox_InsertItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 组合框_插入项图片, 返回项索引.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// hImage: 图片句柄.
func X组合框_插入项图片(元素句柄 int, 项索引 int, 图片句柄 int) int {
	r, _, _ := xComboBox_InsertItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(图片句柄))
	return int(r)
}

// 组合框_插入项图片扩展, 返回项索引.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func X组合框_插入项图片EX(元素句柄 int, 项索引 int, 字段名 string, 图片句柄 int) int {
	r, _, _ := xComboBox_InsertItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名), uintptr(图片句柄))
	return int(r)
}

// 组合框_置项文本.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pText: 文本.
func X组合框_置项文本(元素句柄 int, 项索引 int, 列索引 int, 文本 string) bool {
	r, _, _ := xComboBox_SetItemText.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), 炫彩工具类.StrPtr(文本))
	return r != 0
}

// 组合框_置项文本扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pText: 文本.
func X组合框_置项文本EX(元素句柄 int, 项索引 int, 字段名 string, 文本 string) bool {
	r, _, _ := xComboBox_SetItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名), 炫彩工具类.StrPtr(文本))
	return r != 0
}

// 组合框_置项图片.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func X组合框_置项图片(元素句柄 int, 项索引 int, 列索引 int, 图片句柄 int) bool {
	r, _, _ := xComboBox_SetItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(图片句柄))
	return r != 0
}

// 组合框_置项图片扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func X组合框_置项图片EX(元素句柄 int, 项索引 int, 字段名 string, 图片句柄 int) bool {
	r, _, _ := xComboBox_SetItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名), uintptr(图片句柄))
	return r != 0
}

// 组合框_置项整数值.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 整数值.
func X组合框_置项整数值(元素句柄 int, 项索引 int, 列索引 int, 整数值 int32) bool {
	r, _, _ := xComboBox_SetItemInt.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(整数值))
	return r != 0
}

// 组合框_置项指数值扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// nValue: 整数值.
func X组合框_置项指数值EX(元素句柄 int, 项索引 int, 字段名 string, 整数值 int32) bool {
	r, _, _ := xComboBox_SetItemIntEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名), uintptr(整数值))
	return r != 0
}

// 组合框_置项浮点值.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fFloat: 浮点数.
func X组合框_置项浮点值(元素句柄 int, 项索引 int, 列索引 int, 浮点数 float32) bool {
	r, _, _ := xComboBox_SetItemFloat.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), 炫彩工具类.Float32Ptr(浮点数))
	return r != 0
}

// 组合框_置项浮点值扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// fFloat: 浮点数.
func X组合框_置项浮点值EX(元素句柄 int, 项索引 int, 字段名 string, 浮点数 float32) bool {
	r, _, _ := xComboBox_SetItemFloatEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名), 炫彩工具类.Float32Ptr(浮点数))
	return r != 0
}

// 组合框_取项文本.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func X组合框_取项文本(元素句柄 int, 项索引 int32, 列索引 int32) string {
	r, _, _ := xComboBox_GetItemText.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 组合框_取项文本扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
func X组合框_取项文本EX(元素句柄 int, 项索引 int, 字段名 string) string {
	r, _, _ := xComboBox_GetItemTextEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名))
	return 炫彩工具类.UintPtrToString(r)
}

// 组合框_取项图片.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func X组合框_取项图片(元素句柄 int, 项索引 int, 列索引 int) int {
	r, _, _ := xComboBox_GetItemImage.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引))
	return int(r)
}

// 组合框_取项图片扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
func X组合框_取项图片EX(元素句柄 int, 项索引 int, 字段名 string) int {
	r, _, _ := xComboBox_GetItemImageEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名))
	return int(r)
}

// 组合框_取项整数值.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回整数值.
func X组合框_取项整数值(元素句柄 int, 项索引 int, 列索引 int, 接收返回整数值 *int32) bool {
	r, _, _ := xComboBox_GetItemInt.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(unsafe.Pointer(接收返回整数值)))
	return r != 0
}

// 组合框_取项整数值扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回整数值.
func X组合框_取项整数值EX(元素句柄 int, 项索引 int, 字段名 string, 接收返回整数值 *int32) bool {
	r, _, _ := xComboBox_GetItemIntEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名), uintptr(unsafe.Pointer(接收返回整数值)))
	return r != 0
}

// 组合框_取项浮点值.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返回浮点值.
func X组合框_取项浮点值(元素句柄 int, 项索引 int, 列索引 int, 接收返回浮点值 *float32) bool {
	r, _, _ := xComboBox_GetItemFloat.Call(uintptr(元素句柄), uintptr(项索引), uintptr(列索引), uintptr(unsafe.Pointer(接收返回浮点值)))
	return r != 0
}

// 组合框_取项浮点值扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pOutValue: 接收返回浮点值.
func X组合框_取项浮点值EX(元素句柄 int, 项索引 int, 字段名 string, 接收返回浮点值 *float32) bool {
	r, _, _ := xComboBox_GetItemFloatEx.Call(uintptr(元素句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段名), uintptr(unsafe.Pointer(接收返回浮点值)))
	return r != 0
}

// 组合框_删除项.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
func X组合框_删除项(元素句柄 int, 项索引 int) bool {
	r, _, _ := xComboBox_DeleteItem.Call(uintptr(元素句柄), uintptr(项索引))
	return r != 0
}

// 组合框_删除项扩展.
//
// hEle: 元素句柄.
//
// iItem: 项索引.
//
// nCount: 删除数量.
func X组合框_删除项EX(元素句柄 int, 项索引 int, 删除数量 int) bool {
	r, _, _ := xComboBox_DeleteItemEx.Call(uintptr(元素句柄), uintptr(项索引), uintptr(删除数量))
	return r != 0
}

// 组合框_删除项全部.
//
// hEle: 元素句柄.
func X组合框_删除项全部(元素句柄 int) int {
	r, _, _ := xComboBox_DeleteItemAll.Call(uintptr(元素句柄))
	return int(r)
}

// 组合框_删除列全部.
//
// hEle: 元素句柄.
func X组合框_删除列全部(元素句柄 int) int {
	r, _, _ := xComboBox_DeleteColumnAll.Call(uintptr(元素句柄))
	return int(r)
}

// 组合框_取项数量.
//
// hEle:.
func X组合框_取项数量(hEle int) int {
	r, _, _ := xComboBox_GetCount.Call(uintptr(hEle))
	return int(r)
}

// 组合框_取列数量.
//
// hEle: 元素句柄.
func X组合框_取列数量(元素句柄 int) int {
	r, _, _ := xComboBox_GetCountColumn.Call(uintptr(元素句柄))
	return int(r)
}

// 组合框_弹出下拉列表.
//
// hEle: 元素句柄.
func X组合框_弹出下拉列表(元素句柄 int) int {
	r, _, _ := xComboBox_PopupDropList.Call(uintptr(元素句柄))
	return int(r)
}

// 组合框_设置项模板.
//
// hEle: 元素句柄.
//
// hTemp: 模板句柄.
func X组合框_设置项模板(元素句柄, 模板句柄 int) int {
	r, _, _ := xComboBox_SetItemTemplate.Call(uintptr(元素句柄), uintptr(模板句柄))
	return int(r)
}

// 组合框_置项模板从内存.
//
// hEle: 元素句柄.
//
// data: 模板数据.
func X组合框_置项模板从内存(元素句柄 int, 模板数据 []byte) bool {
	r, _, _ := xComboBox_SetItemTemplateXMLFromMem.Call(uintptr(元素句柄), 炫彩工具类.ByteSliceDataPtr(&模板数据), uintptr(len(模板数据)))
	return r != 0
}

// 组合框_置项模板从资源ZIP.
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
func X组合框_置项模板从资源ZIP(元素句柄 int, RC资源ID int32, 文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	r, _, _ := xComboBox_SetItemTemplateXMLFromZipRes.Call(uintptr(元素句柄), uintptr(RC资源ID), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(zip密码), 模块句柄)
	return r != 0
}

// 组合框_取项模板, 返回项模板句柄.
//
// hEle: 元素句柄.
func X组合框_取项模板(元素句柄 int) int {
	r, _, _ := xComboBox_GetItemTemplate.Call(uintptr(元素句柄))
	return int(r)
}
