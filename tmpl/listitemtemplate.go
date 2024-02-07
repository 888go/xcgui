package 炫彩模板类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// ListItemTemplate 列表项模板.
type ListItemTemplate struct {
	炫彩对象基类.ObjectBase
}

// 模板_创建, 创建项模板.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
func X创建(模板类型 炫彩常量类.ListItemTemp_Type_) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(炫彩基类.X模板_创建(模板类型))
	return p
}

// NewByHandle 从句柄创建对象.
func X创建并按句柄(句柄 int) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(句柄)
	return p
}

// 模板_加载从文件, 列表项模板文件载入.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
func X创建并按文件(模板类型 炫彩常量类.ListItemTemp_Type_, 文件名 string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(炫彩基类.X模板_加载从文件(模板类型, 文件名))
	return p
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func X创建并按ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, zip文件 string, 文件名 string, zip密码 string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(炫彩基类.X模板_加载从ZIP(模板类型, zip文件, 文件名, zip密码))
	return p
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中.
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
func X创建并按RC资源ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, RC资源ID int32, 模板文件名 string, zip密码 string, 模块句柄 uintptr) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(炫彩基类.X项模板_加载从资源ZIP(模板类型, RC资源ID, 模板文件名, zip密码, 模块句柄))
	return p
}

// 模板_加载从内存ZIP, 加载列表项模板从内存zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func X创建并按内存ZIP(模板类型 炫彩常量类.ListItemTemp_Type_, 模板文件数据 []byte, 文件名 string, zip密码 string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(炫彩基类.X模板_加载从内存ZIP(模板类型, 模板文件数据, 文件名, zip密码))
	return p
}

// 项模板_加载从内存. 加载列表项模板文件从内存.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
func X创建并按内存(模板类型 炫彩常量类.ListItemTemp_Type_, 模板文件数据 []byte) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(炫彩基类.X项模板_加载从内存(模板类型, 模板文件数据))
	return p
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串.
func X创建并按字符串(模板类型 炫彩常量类.ListItemTemp_Type_, 字符串 string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(炫彩基类.X模板_加载从字符串(模板类型, 字符串))
	return p
}

// 模板_克隆, 复制一份新的项模板, 返回模板对象.
func (l *ListItemTemplate) X取副本() *ListItemTemplate {
	p := &ListItemTemplate{}
	p.X设置句柄(炫彩基类.X项模板_克隆(l.Handle))
	return p
}

// 模板_取类型, 获取列表项模板类型, 返回: xcc.ListItemTemp_Type_.
func (l *ListItemTemplate) X取类型() 炫彩常量类.ListItemTemp_Type_ {
	return 炫彩基类.X模板_取类型(l.Handle)
}

// 模板_销毁, 项模板销毁.
func (l *ListItemTemplate) X销毁() bool {
	return 炫彩基类.X模板_销毁(l.Handle)
}

// 模板_添加根节点.
//
// pNode: 节点指针.
func (l *ListItemTemplate) X添加根节点(节点指针 int) bool {
	return 炫彩基类.X模板_添加根节点(l.Handle, 节点指针)
}

// 模板_取列表中的节点.
//
// index: 节点位置索引.
func (l *ListItemTemplate) X取列表中节点(节点位置索引 int32) int {
	return 炫彩基类.X模板_取列表中的节点(l.Handle, 节点位置索引)
}

// 项模板_列表_插入节点.
//
// index: 插入位置索引.
//
// pNode: 节点指针.
func (l *ListItemTemplate) X列表插入节点(插入位置索引 int32, 节点指针 int) bool {
	return 炫彩基类.X项模板_列表_插入节点(l.Handle, 插入位置索引, 节点指针)
}

// 项模板_列表_删除节点.
//
// index: 位置索引.
func (l *ListItemTemplate) X列表删除节点(位置索引 int32) bool {
	return 炫彩基类.X项模板_列表_删除节点(l.Handle, 位置索引)
}

// 项模板_列表_取数量, 取子节点数量, 只当前层子节点.
func (l *ListItemTemplate) X列表取数量() int32 {
	return 炫彩基类.X项模板_列表_取数量(l.Handle)
}

// 项模板_列表_移动列, 将指定列移动到目标位置.
//
// iColSrc: 源列索引.
//
// iColDest: 目标列索引.
func (l *ListItemTemplate) X列表移动列(源列索引, 目标列索引 int32) bool {
	return 炫彩基类.X项模板_列表_移动列(l.Handle, 源列索引, 目标列索引)
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
func X加载并按文件EX(模板类型 炫彩常量类.ListItemTemp_Type_, 文件名 string, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	return 炫彩基类.X模板_加载从文件EX(模板类型, 文件名, 返回模板句柄1, 返回模板句柄2)
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
func X加载并按ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, zip文件 string, 文件名 string, zip密码 string, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	return 炫彩基类.X模板_加载从ZIPEX(模板类型, zip文件, 文件名, zip密码, 返回模板句柄1, 返回模板句柄2)
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
func X加载并按内存ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, 模板文件数据 []byte, 文件名 string, zip密码 string, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	return 炫彩基类.X模板_加载从内存ZIPEX(模板类型, 模板文件数据, 文件名, zip密码, 返回模板句柄1, 返回模板句柄2)
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
func X加载并按字符串EX(模板类型 炫彩常量类.ListItemTemp_Type_, 字符串内容 string, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	return 炫彩基类.X模板_加载从字符串EX(模板类型, 字符串内容, 返回模板句柄1, 返回模板句柄2)
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
func X加载并按内存EX(模板类型 炫彩常量类.ListItemTemp_Type_, 模板文件数据 []byte, 返回模板句柄1 *int, 返回模板句柄2 *int) bool {
	return 炫彩基类.X项模板_加载从内存EX(模板类型, 模板文件数据, 返回模板句柄1, 返回模板句柄2)
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
func X加载并按资源ZIPEX(模板类型 炫彩常量类.ListItemTemp_Type_, RC资源ID int32, 模板文件名 string, zip密码 string, 返回模板句柄1 *int, 返回模板句柄2 *int, 模块句柄 uintptr) int {
	return 炫彩基类.X项模板_加载从资源ZIPEX(模板类型, RC资源ID, 模板文件名, zip密码, 返回模板句柄1, 返回模板句柄2, 模块句柄)
}
