package tmpl//bm:炫彩模板类

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ListItemTemplate 列表项模板.
type ListItemTemplate struct {
	objectbase.ObjectBase
}

// 模板_创建, 创建项模板.
//, xcc.ListItemTemp_Type_.
// ff:创建
// nType:模板类型
func New(nType xcc.ListItemTemp_Type_) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_Create(nType))
	return p
}

// NewByHandle 从句柄创建对象.
// ff:创建并按句柄
// handle:句柄
func NewByHandle(handle int) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(handle)
	return p
}

// 模板_加载从文件, 列表项模板文件载入.
//, xcc.ListItemTemp_Type_.
//.
// ff:创建并按文件
// nType:模板类型
// pFileName:
func NewByXML(nType xcc.ListItemTemp_Type_, pFileName string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_Load(nType, pFileName))
	return p
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中.
//, xcc.ListItemTemp_Type_.
//.
//.
//.
// ff:创建并按ZIP
// nType:模板类型
// pZipFile:
// pFileName:
// pPassword:
func NewByZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_LoadZip(nType, pZipFile, pFileName, pPassword))
	return p
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中.
//, xcc.ListItemTemp_Type_.
//.
//.
//.
//, 可填0.
// ff:创建并按RC资源ZIP
// nType:模板类型
// id:
// pFileName:
// pPassword:
// hModule:
func NewByZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_LoadZipRes(nType, id, pFileName, pPassword, hModule))
	return p
}

// 模板_加载从内存ZIP, 加载列表项模板从内存zip压缩包中.
//, xcc.ListItemTemp_Type_.
//.
//.
//.
// ff:创建并按内存ZIP
// nType:模板类型
// data:
// pFileName:
// pPassword:
func NewByZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_LoadZipMem(nType, data, pFileName, pPassword))
	return p
}

// 项模板_加载从内存. 加载列表项模板文件从内存.
//, xcc.ListItemTemp_Type_.
//.
// ff:创建并按内存
// nType:模板类型
// data:
func NewByXmlMem(nType xcc.ListItemTemp_Type_, data []byte) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_LoadFromMem(nType, data))
	return p
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串.
//, xcc.ListItemTemp_Type_.
//.
// ff:创建并按字符串
// nType:模板类型
// pStringXML:
func NewByString(nType xcc.ListItemTemp_Type_, pStringXML string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_LoadFromString(nType, pStringXML))
	return p
}

// 模板_克隆, 复制一份新的项模板, 返回模板对象.
// ff:取副本
func (l *ListItemTemplate) Clone() *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_Clone(l.Handle))
	return p
}

// 模板_取类型, 获取列表项模板类型, 返回: xcc.ListItemTemp_Type_.
// ff:取类型
func (l *ListItemTemplate) GetType() xcc.ListItemTemp_Type_ {
	return xc.XTemp_GetType(l.Handle)
}

// 模板_销毁, 项模板销毁.
// ff:销毁
func (l *ListItemTemplate) Destroy() bool {
	return xc.XTemp_Destroy(l.Handle)
}

// 模板_添加根节点.
//.
// ff:添加根节点
// pNode:节点指针
func (l *ListItemTemplate) AddNodeRoot(pNode int) bool {
	return xc.XTemp_AddNodeRoot(l.Handle, pNode)
}

// 模板_取列表中的节点.
//.
// ff:取列表中节点
// index:节点位置索引
func (l *ListItemTemplate) List_GetNode(index int32) int {
	return xc.XTemp_List_GetNode(l.Handle, index)
}

// 项模板_列表_插入节点.
//.
//.
// ff:列表插入节点
// index:插入位置索引
// pNode:节点指针
func (l *ListItemTemplate) List_InsertNode(index int32, pNode int) bool {
	return xc.XTemp_List_InsertNode(l.Handle, index, pNode)
}

// 项模板_列表_删除节点.
//.
// ff:列表删除节点
// index:位置索引
func (l *ListItemTemplate) List_DeleteNode(index int32) bool {
	return xc.XTemp_List_DeleteNode(l.Handle, index)
}

// 项模板_列表_取数量, 取子节点数量, 只当前层子节点.
// ff:列表取数量
func (l *ListItemTemplate) List_GetCount() int32 {
	return xc.XTemp_List_GetCount(l.Handle)
}

// 项模板_列表_移动列, 将指定列移动到目标位置.
//.
//.
// ff:列表移动列
// iColSrc:源列索引
// iColDest:目标列索引
func (l *ListItemTemplate) List_MoveColumn(iColSrc, iColDest int32) bool {
	return xc.XTemp_List_MoveColumn(l.Handle, iColSrc, iColDest)
}

// 模板_加载从文件扩展, 加载列表项模板从文件.
//, xcc.ListItemTemp_Type_.
//.
//, 项模板.
//, 列表头模板或列表视组模板.
// ff:加载并按文件EX
// nType:模板类型
// pFileName:
// pOutTemp1:
// pOutTemp2:
func LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadEx(nType, pFileName, pOutTemp1, pOutTemp2)
}

// 模板_加载从ZIP扩展, 加载列表项模板从zip压缩包中.
//, xcc.ListItemTemp_Type_.
//.
//.
//.
//, 项模板.
//, 列表头模板或列表视组模板.
// ff:加载并按ZIPEX
// nType:模板类型
// pZipFile:
// pFileName:
// pPassword:
// pOutTemp1:
// pOutTemp2:
func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadZipEx(nType, pZipFile, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从内存ZIP扩展, 加载列表项模板从内存zip压缩包中.
//, xcc.ListItemTemp_Type_.
//.
//.
//.
//, 项模板.
//, 列表头模板或列表视组模板.
// ff:加载并按内存ZIPEX
// nType:模板类型
// data:
// pFileName:
// pPassword:
// pOutTemp1:
// pOutTemp2:
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadZipMemEx(nType, data, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从字符串扩展, 加载列表项模板文件从内存字符串.
//, xcc.ListItemTemp_Type_.
//内容.
//, 项模板.
//, 列表头模板或列表视组模板.
// ff:加载并按字符串EX
// nType:模板类型
// pStringXML:
// pOutTemp1:
// pOutTemp2:
func LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadFromStringEx(nType, pStringXML, pOutTemp1, pOutTemp2)
}

// 项模板_加载从内存扩展. 加载列表项模板文件从内存.
//, xcc.ListItemTemp_Type_.
//.
//, 项模板.
//, 列表头模板或列表视组模板.
// ff:加载并按内存EX
// nType:模板类型
// data:
// pOutTemp1:
// pOutTemp2:
func LoadFromMemEx(nType xcc.ListItemTemp_Type_, data []byte, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadFromMemEx(nType, data, pOutTemp1, pOutTemp2)
}

// 项模板_加载从资源ZIP扩展. 加载列表项模板文件从RC资源ZIP, 返回模板句柄.
//, xcc.ListItemTemp_Type_.
//.
//.
//.
//, 项模板.
//, 列表头模板或列表视组模板.
//, 可填0.
// ff:加载并按资源ZIPEX
// nType:模板类型
// id:
// pFileName:
// pPassword:
// pOutTemp1:
// pOutTemp2:
// hModule:
func LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int, hModule uintptr) int {
	return xc.XTemp_LoadZipResEx(nType, id, pFileName, pPassword, pOutTemp1, pOutTemp2, hModule)
}
