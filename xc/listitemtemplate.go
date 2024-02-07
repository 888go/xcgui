package 炫彩基类

import (
	"unsafe"
	
	"github.com/888go/xcgui/common"
	
	"github.com/888go/xcgui/xcc"
)

// 模板_加载从文件, 列表项模板文件载入, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
func X模板_加载从文件(模板类型 炫彩常量类.ListItemTemp_Type_, 文件名 string) int {
	r, _, _ := xTemp_Load.Call(uintptr(模板类型), 炫彩工具类.StrPtr(文件名))
	return int(r)
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func X模板_加载从ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, zip文件 string, 文件名 string, zip密码 string) int {
	r, _, _ := xTemp_LoadZip.Call(uintptr(模板类型), 炫彩工具类.StrPtr(zip文件), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(zip密码))
	return int(r)
}

// 模板_加载从内存ZIP, 加载列表项模板从内存zip压缩包中, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func X模板_加载从内存ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, 模板文件数据 []byte, 文件名 string, zip密码 string) int {
	r, _, _ := xTemp_LoadZipMem.Call(uintptr(模板类型), 炫彩工具类.ByteSliceDataPtr(&模板文件数据), uintptr(len(模板文件数据)), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(zip密码))
	return int(r)
}

// 模板_加载从文件扩展, 加载列表项模板从文件.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func X模板_加载从文件EX(模板类型 炫彩常量类.ListItemTemp_Type_, 文件名 string, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	r, _, _ := xTemp_LoadEx.Call(uintptr(模板类型), 炫彩工具类.StrPtr(文件名), uintptr(unsafe.Pointer(&返回模板句柄1)), uintptr(unsafe.Pointer(&返回模板句柄2)))
	return r != 0
}

// 项模板_加载从内存. 加载列表项模板文件从内存, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
func X项模板_加载从内存(模板类型 炫彩常量类.ListItemTemp_Type_, 模板文件数据 []byte) int {
	r, _, _ := xTemp_LoadFromMem.Call(uintptr(模板类型), 炫彩工具类.ByteSliceDataPtr(&模板文件数据), uintptr(len(模板文件数据)))
	return int(r)
}

// 项模板_加载从内存扩展. 加载列表项模板文件从内存.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func X项模板_加载从内存EX(模板类型 炫彩常量类.ListItemTemp_Type_, 模板文件数据 []byte, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	r, _, _ := xTemp_LoadFromMemEx.Call(uintptr(模板类型), 炫彩工具类.ByteSliceDataPtr(&模板文件数据), uintptr(len(模板文件数据)), uintptr(unsafe.Pointer(&返回模板句柄1)), uintptr(unsafe.Pointer(&返回模板句柄2)))
	return r != 0
}

// 项模板_加载从资源ZIP. 加载列表项模板文件从RC资源ZIP, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// id: RC资源ID.
//
// pFileName: 模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func X项模板_加载从资源ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, RC资源ID int32, 模板文件名 string, zip密码 string, 模块句柄 uintptr) int {
	r, _, _ := xTemp_LoadZipRes.Call(uintptr(模板类型), uintptr(RC资源ID), 炫彩工具类.StrPtr(模板文件名), 炫彩工具类.StrPtr(zip密码), 模块句柄)
	return int(r)
}

// 项模板_加载从资源ZIP扩展. 加载列表项模板文件从RC资源ZIP, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// id: RC资源ID.
//
// pFileName: 模板文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
//
// hModule: 模块句柄, 可填0.
func X项模板_加载从资源ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, RC资源ID int32, 模板文件名 string, zip密码 string, 返回模板句柄1 *int, 返回模板句柄2 *int, 模块句柄 uintptr) int {
	r, _, _ := xTemp_LoadZipResEx.Call(uintptr(模板类型), uintptr(RC资源ID), 炫彩工具类.StrPtr(模板文件名), 炫彩工具类.StrPtr(zip密码), uintptr(unsafe.Pointer(&返回模板句柄1)), uintptr(unsafe.Pointer(&返回模板句柄2)), 模块句柄)
	return int(r)
}

// 模板_加载从ZIP扩展, 加载列表项模板从zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func X模板_加载从ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, zip文件 string, 文件名 string, zip密码 string, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	r, _, _ := xTemp_LoadZipEx.Call(uintptr(模板类型), 炫彩工具类.StrPtr(zip文件), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(zip密码), uintptr(unsafe.Pointer(&返回模板句柄1)), uintptr(unsafe.Pointer(&返回模板句柄2)))
	return r != 0
}

// 模板_加载从内存ZIP扩展, 加载列表项模板从内存zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func X模板_加载从内存ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, 模板文件数据 []byte, 文件名 string, zip密码 string, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	r, _, _ := xTemp_LoadZipMemEx.Call(uintptr(模板类型), 炫彩工具类.ByteSliceDataPtr(&模板文件数据), uintptr(len(模板文件数据)), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(zip密码), uintptr(unsafe.Pointer(&返回模板句柄1)), uintptr(unsafe.Pointer(&返回模板句柄2)))
	return r != 0
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串.
func X模板_加载从字符串(模板类型 炫彩常量类.ListItemTemp_Type_, 字符串 string) int {
	r, _, _ := xTemp_LoadFromString.Call(uintptr(模板类型), W2A(字符串))
	return int(r)
}

// 模板_加载从字符串扩展, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串内容.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func X模板_加载从字符串EX(模板类型 炫彩常量类.ListItemTemp_Type_, 字符串内容 string, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	r, _, _ := xTemp_LoadFromStringEx.Call(uintptr(模板类型), W2A(字符串内容), uintptr(unsafe.Pointer(&返回模板句柄1)), uintptr(unsafe.Pointer(&返回模板句柄2)))
	return r != 0
}

// 模板_取类型, 获取列表项模板类型, 返回: xcc.ListItemTemp_Type_.
//
// hTemp: 列表项模板句柄.
func X模板_取类型(列表项模板句柄 int) 炫彩常量类.ListItemTemp_Type_ {
	r, _, _ := xTemp_GetType.Call(uintptr(列表项模板句柄))
	return 炫彩常量类.ListItemTemp_Type_(r)
}

// 模板_销毁, 项模板销毁.
//
// hTemp: 项模板句柄.
func X模板_销毁(项模板句柄 int) bool {
	r, _, _ := xTemp_Destroy.Call(uintptr(项模板句柄))
	return r != 0
}

// 模板_创建, 创建项模板, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
func X模板_创建(模板类型 炫彩常量类.ListItemTemp_Type_) int {
	r, _, _ := xTemp_Create.Call(uintptr(模板类型))
	return int(r)
}

// 模板_添加根节点.
//
// hTemp: 项模板句柄.
//
// pNode: 节点指针.
func X模板_添加根节点(项模板句柄 int, 节点指针 int) bool {
	r, _, _ := xTemp_AddNodeRoot.Call(uintptr(项模板句柄), uintptr(节点指针))
	return r != 0
}

// 模板_添加子节点.
//
// pParentNode: 父节点指针.
//
// pNode: 节点指针.
func X模板_添加子节点(父节点指针 int, 节点指针 int) bool {
	r, _, _ := xTemp_AddNode.Call(uintptr(父节点指针), uintptr(节点指针))
	return r != 0
}

// 模板_创建节点.
//
// nType: 对象类型: xcc.XC_OBJECT_TYPE.
func X模板_创建节点(对象类型 炫彩常量类.XC_OBJECT_TYPE) int {
	r, _, _ := xTemp_CreateNode.Call(uintptr(对象类型))
	return int(r)
}

// 模板_置节点属性.
//
// pNode: 节点指针.
//
// pName: 属性名.
//
// pAttr: 属性值.
func X模板_置节点属性(节点指针 int, 属性名 string, 属性值 string) bool {
	r, _, _ := xTemp_SetNodeAttribute.Call(uintptr(节点指针), 炫彩工具类.StrPtr(属性名), 炫彩工具类.StrPtr(属性值))
	return r != 0
}

// 模板_置节点属性扩展.
//
// pNode: 节点指针.
//
// itemID: 模板项ID.
//
// pName: 属性名.
//
// pAttr: 属性值.
func X模板_置节点属性EX(节点指针 int, 模板项ID int32, 属性名 string, 属性值 string) bool {
	r, _, _ := xTemp_SetNodeAttributeEx.Call(uintptr(节点指针), uintptr(模板项ID), 炫彩工具类.StrPtr(属性名), 炫彩工具类.StrPtr(属性值))
	return r != 0
}

// 模板_取列表中的节点.
//
// hTemp: 模板句柄.
//
// index: 节点位置索引.
func X模板_取列表中的节点(模板句柄 int, 节点位置索引 int32) int {
	r, _, _ := xTemp_List_GetNode.Call(uintptr(模板句柄), uintptr(节点位置索引))
	return int(r)
}

// 模板_取节点, 获取节点, 根据itemID. 返回itemID对应的节点指针.
//
// pNode: 节点指针.
//
// itemID: ID.
func X模板_取节点(节点指针 int, ID int32) int {
	r, _, _ := xTemp_GetNode.Call(uintptr(节点指针), uintptr(ID))
	return int(r)
}

// 模板_克隆节点, 克隆一个节点, 返回克隆的节点.
//
// pNode: 节点指针.
func X模板_克隆节点(节点指针 int) int {
	r, _, _ := xTemp_CloneNode.Call(uintptr(节点指针))
	return int(r)
}

// 项模板_克隆, 复制一份新的项模板, 返回模板句柄.
//
// hTemp: 列表项模板句柄.
func X项模板_克隆(列表项模板句柄 int) int {
	r, _, _ := xTemp_Clone.Call(uintptr(列表项模板句柄))
	return int(r)
}

// 项模板_列表_插入节点.
//
// hTemp: 列表项模板句柄.
//
// index: 插入位置索引.
//
// pNode: 节点指针.
func X项模板_列表_插入节点(列表项模板句柄 int, 插入位置索引 int32, 节点指针 int) bool {
	r, _, _ := xTemp_List_InsertNode.Call(uintptr(列表项模板句柄), uintptr(插入位置索引), uintptr(节点指针))
	return r != 0
}

// 项模板_列表_删除节点.
//
// hTemp: 列表项模板句柄.
//
// index: 位置索引.
func X项模板_列表_删除节点(列表项模板句柄 int, 位置索引 int32) bool {
	r, _, _ := xTemp_List_DeleteNode.Call(uintptr(列表项模板句柄), uintptr(位置索引))
	return r != 0
}

// 项模板_列表_取数量, 取子节点数量, 只当前层子节点.
//
// hTemp: 列表项模板句柄.
func X项模板_列表_取数量(列表项模板句柄 int) int32 {
	r, _, _ := xTemp_List_GetCount.Call(uintptr(列表项模板句柄))
	return int32(r)
}

// 项模板_列表_移动列, 将指定列移动到目标位置.
//
// hTemp: 列表项模板句柄.
//
// iColSrc: 源列索引.
//
// iColDest: 目标列索引.
func X项模板_列表_移动列(列表项模板句柄 int, 源列索引, 目标列索引 int32) bool {
	r, _, _ := xTemp_List_MoveColumn.Call(uintptr(列表项模板句柄), uintptr(源列索引), uintptr(目标列索引))
	return r != 0
}
