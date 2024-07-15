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

// ff:模板_加载从文件
// pFileName:
// nType:模板类型
func X模板_加载从文件(模板类型 炫彩常量类.ListItemTemp_Type_, pFileName string) int {
	r, _, _ := xTemp_Load.Call(uintptr(模板类型), 炫彩工具类.StrPtr(pFileName))
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

// ff:模板_加载从ZIP
// pPassword:
// pFileName:
// pZipFile:
// nType:模板类型
func X模板_加载从ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) int {
	r, _, _ := xTemp_LoadZip.Call(uintptr(模板类型), 炫彩工具类.StrPtr(pZipFile), 炫彩工具类.StrPtr(pFileName), 炫彩工具类.StrPtr(pPassword))
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

// ff:模板_加载从内存ZIP
// pPassword:
// pFileName:
// data:
// nType:模板类型
func X模板_加载从内存ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) int {
	r, _, _ := xTemp_LoadZipMem.Call(uintptr(模板类型), 炫彩工具类.ByteSliceDataPtr(&data), uintptr(len(data)), 炫彩工具类.StrPtr(pFileName), 炫彩工具类.StrPtr(pPassword))
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

// ff:模板_加载从文件EX
// pOutTemp2:
// pOutTemp1:
// pFileName:
// nType:模板类型
func X模板_加载从文件EX(模板类型 炫彩常量类.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadEx.Call(uintptr(模板类型), 炫彩工具类.StrPtr(pFileName), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
	return r != 0
}

// 项模板_加载从内存. 加载列表项模板文件从内存, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.

// ff:项模板_加载从内存
// data:
// nType:模板类型
func X项模板_加载从内存(模板类型 炫彩常量类.ListItemTemp_Type_, data []byte) int {
	r, _, _ := xTemp_LoadFromMem.Call(uintptr(模板类型), 炫彩工具类.ByteSliceDataPtr(&data), uintptr(len(data)))
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

// ff:项模板_加载从内存EX
// pOutTemp2:
// pOutTemp1:
// data:
// nType:模板类型
func X项模板_加载从内存EX(模板类型 炫彩常量类.ListItemTemp_Type_, data []byte, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadFromMemEx.Call(uintptr(模板类型), 炫彩工具类.ByteSliceDataPtr(&data), uintptr(len(data)), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
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

// ff:项模板_加载从资源ZIP
// hModule:
// pPassword:
// pFileName:
// id:
// nType:模板类型
func X项模板_加载从资源ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) int {
	r, _, _ := xTemp_LoadZipRes.Call(uintptr(模板类型), uintptr(id), 炫彩工具类.StrPtr(pFileName), 炫彩工具类.StrPtr(pPassword), hModule)
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

// ff:项模板_加载从资源ZIPEX
// hModule:
// pOutTemp2:
// pOutTemp1:
// pPassword:
// pFileName:
// id:
// nType:模板类型
func X项模板_加载从资源ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int, hModule uintptr) int {
	r, _, _ := xTemp_LoadZipResEx.Call(uintptr(模板类型), uintptr(id), 炫彩工具类.StrPtr(pFileName), 炫彩工具类.StrPtr(pPassword), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)), hModule)
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

// ff:模板_加载从ZIPEX
// pOutTemp2:
// pOutTemp1:
// pPassword:
// pFileName:
// pZipFile:
// nType:模板类型
func X模板_加载从ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadZipEx.Call(uintptr(模板类型), 炫彩工具类.StrPtr(pZipFile), 炫彩工具类.StrPtr(pFileName), 炫彩工具类.StrPtr(pPassword), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
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

// ff:模板_加载从内存ZIPEX
// pOutTemp2:
// pOutTemp1:
// pPassword:
// pFileName:
// data:
// nType:模板类型
func X模板_加载从内存ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadZipMemEx.Call(uintptr(模板类型), 炫彩工具类.ByteSliceDataPtr(&data), uintptr(len(data)), 炫彩工具类.StrPtr(pFileName), 炫彩工具类.StrPtr(pPassword), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
	return r != 0
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串.

// ff:模板_加载从字符串
// pStringXML:
// nType:模板类型
func X模板_加载从字符串(模板类型 炫彩常量类.ListItemTemp_Type_, pStringXML string) int {
	r, _, _ := xTemp_LoadFromString.Call(uintptr(模板类型), W2A(pStringXML))
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

// ff:模板_加载从字符串EX
// pOutTemp2:
// pOutTemp1:
// pStringXML:
// nType:模板类型
func X模板_加载从字符串EX(模板类型 炫彩常量类.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadFromStringEx.Call(uintptr(模板类型), W2A(pStringXML), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
	return r != 0
}

// 模板_取类型, 获取列表项模板类型, 返回: xcc.ListItemTemp_Type_.
//
// hTemp: 列表项模板句柄.

// ff:模板_取类型
// hTemp:列表项模板句柄
func X模板_取类型(列表项模板句柄 int) 炫彩常量类.ListItemTemp_Type_ {
	r, _, _ := xTemp_GetType.Call(uintptr(列表项模板句柄))
	return 炫彩常量类.ListItemTemp_Type_(r)
}

// 模板_销毁, 项模板销毁.
//
// hTemp: 项模板句柄.

// ff:模板_销毁
// hTemp:项模板句柄
func X模板_销毁(项模板句柄 int) bool {
	r, _, _ := xTemp_Destroy.Call(uintptr(项模板句柄))
	return r != 0
}

// 模板_创建, 创建项模板, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.

// ff:模板_创建
// nType:模板类型
func X模板_创建(模板类型 炫彩常量类.ListItemTemp_Type_) int {
	r, _, _ := xTemp_Create.Call(uintptr(模板类型))
	return int(r)
}

// 模板_添加根节点.
//
// hTemp: 项模板句柄.
//
// pNode: 节点指针.

// ff:模板_添加根节点
// pNode:节点指针
// hTemp:项模板句柄
func X模板_添加根节点(项模板句柄 int, 节点指针 int) bool {
	r, _, _ := xTemp_AddNodeRoot.Call(uintptr(项模板句柄), uintptr(节点指针))
	return r != 0
}

// 模板_添加子节点.
//
// pParentNode: 父节点指针.
//
// pNode: 节点指针.

// ff:模板_添加子节点
// pNode:节点指针
// pParentNode:父节点指针
func X模板_添加子节点(父节点指针 int, 节点指针 int) bool {
	r, _, _ := xTemp_AddNode.Call(uintptr(父节点指针), uintptr(节点指针))
	return r != 0
}

// 模板_创建节点.
//
// nType: 对象类型: xcc.XC_OBJECT_TYPE.

// ff:模板_创建节点
// nType:对象类型
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

// ff:模板_置节点属性
// pAttr:属性值
// pName:属性名
// pNode:节点指针
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

// ff:模板_置节点属性EX
// pAttr:属性值
// pName:属性名
// itemID:模板项ID
// pNode:节点指针
func X模板_置节点属性EX(节点指针 int, 模板项ID int32, 属性名 string, 属性值 string) bool {
	r, _, _ := xTemp_SetNodeAttributeEx.Call(uintptr(节点指针), uintptr(模板项ID), 炫彩工具类.StrPtr(属性名), 炫彩工具类.StrPtr(属性值))
	return r != 0
}

// 模板_取列表中的节点.
//
// hTemp: 模板句柄.
//
// index: 节点位置索引.

// ff:模板_取列表中的节点
// index:节点位置索引
// hTemp:模板句柄
func X模板_取列表中的节点(模板句柄 int, 节点位置索引 int32) int {
	r, _, _ := xTemp_List_GetNode.Call(uintptr(模板句柄), uintptr(节点位置索引))
	return int(r)
}

// 模板_取节点, 获取节点, 根据itemID. 返回itemID对应的节点指针.
//
// pNode: 节点指针.
//
// itemID: ID.

// ff:模板_取节点
// itemID:ID
// pNode:节点指针
func X模板_取节点(节点指针 int, ID int32) int {
	r, _, _ := xTemp_GetNode.Call(uintptr(节点指针), uintptr(ID))
	return int(r)
}

// 模板_克隆节点, 克隆一个节点, 返回克隆的节点.
//
// pNode: 节点指针.

// ff:模板_克隆节点
// pNode:节点指针
func X模板_克隆节点(节点指针 int) int {
	r, _, _ := xTemp_CloneNode.Call(uintptr(节点指针))
	return int(r)
}

// 项模板_克隆, 复制一份新的项模板, 返回模板句柄.
//
// hTemp: 列表项模板句柄.

// ff:项模板_克隆
// hTemp:列表项模板句柄
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

// ff:项模板_列表_插入节点
// pNode:节点指针
// index:插入位置索引
// hTemp:列表项模板句柄
func X项模板_列表_插入节点(列表项模板句柄 int, 插入位置索引 int32, 节点指针 int) bool {
	r, _, _ := xTemp_List_InsertNode.Call(uintptr(列表项模板句柄), uintptr(插入位置索引), uintptr(节点指针))
	return r != 0
}

// 项模板_列表_删除节点.
//
// hTemp: 列表项模板句柄.
//
// index: 位置索引.

// ff:项模板_列表_删除节点
// index:位置索引
// hTemp:列表项模板句柄
func X项模板_列表_删除节点(列表项模板句柄 int, 位置索引 int32) bool {
	r, _, _ := xTemp_List_DeleteNode.Call(uintptr(列表项模板句柄), uintptr(位置索引))
	return r != 0
}

// 项模板_列表_取数量, 取子节点数量, 只当前层子节点.
//
// hTemp: 列表项模板句柄.

// ff:项模板_列表_取数量
// hTemp:列表项模板句柄
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

// ff:项模板_列表_移动列
// iColDest:目标列索引
// iColSrc:源列索引
// hTemp:列表项模板句柄
func X项模板_列表_移动列(列表项模板句柄 int, 源列索引, 目标列索引 int32) bool {
	r, _, _ := xTemp_List_MoveColumn.Call(uintptr(列表项模板句柄), uintptr(源列索引), uintptr(目标列索引))
	return r != 0
}
