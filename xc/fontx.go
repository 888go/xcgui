package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
	
	"github.com/888go/xcgui/xcc"
)

// XFont_Create 字体_创建, 创建炫彩字体. 当字体句柄与元素关联后, 会自动释放.
//
//	@param size 字体大小,单位(pt,磅).
//	@return int 返回字体句柄.
func X字体_创建(字体大小 int32) int {
	r, _, _ := xFont_Create.Call(uintptr(字体大小))
	return int(r)
}

// XFont_CreateEx 字体_创建扩展. 创建炫彩字体.
//
//	@param pName 字体名称.
//	@param size 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return int 返回字体句柄.
func X字体_创建EX(字体名称 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) int {
	r, _, _ := xFont_CreateEx.Call(炫彩工具类.StrPtr(字体名称), uintptr(字体大小), uintptr(字体样式))
	return int(r)
}

// XFont_CreateLOGFONTW 字体_创建从LOGFONT. 创建炫彩字体.
//
//	@param pFontInfo 字体信息.
//	@return int 返回字体句柄.
func X字体_创建从LOGFONT(字体信息 *LOGFONTW) int {
	r, _, _ := xFont_CreateLOGFONTW.Call(uintptr(unsafe.Pointer(字体信息)))
	return int(r)
}

// XFont_CreateFromHFONT 字体_创建从HFONT. 创建炫彩字体从现有HFONT字体.
//
//	@param hFont 字体句柄.
//	@return int 返回字体句柄.
func X字体_创建从HFONT(字体句柄 uintptr) int {
	r, _, _ := xFont_CreateFromHFONT.Call(字体句柄)
	return int(r)
}

// XFont_CreateFromFont 字体_创建从Font. 创建炫彩字体从GDI+字体.
//
//	@param pFont GDI+字体指针.
//	@return int 返回字体句柄.
func X字体_创建从Font(GDI字体指针 uintptr) int {
	r, _, _ := xFont_CreateFromFont.Call(GDI字体指针)
	return int(r)
}

// XFont_CreateFromFile 字体_创建从文件. 创建字体从文件.
//
//	@param pFontFile 字体文件名.
//	@param size 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return int 返回炫彩字体句柄.
func X字体_创建从文件(字体文件名 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) int {
	r, _, _ := xFont_CreateFromFile.Call(炫彩工具类.StrPtr(字体文件名), uintptr(字体大小), uintptr(字体样式))
	return int(r)
}

// XFont_CreateFromZip 字体_创建从ZIP.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 字体文件名.
//	@param pPassword zip密码.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式: xcc.FontStyle_.
//	@return int 返回炫彩字体句柄.
func X字体_创建从ZIP(zip文件名, 字体文件名, zip密码 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) int {
	r, _, _ := xFont_CreateFromZip.Call(炫彩工具类.StrPtr(zip文件名), 炫彩工具类.StrPtr(字体文件名), 炫彩工具类.StrPtr(zip密码), uintptr(字体大小), uintptr(字体样式))
	return int(r)
}

// XFont_CreateFromZipMem 字体_创建从内存ZIP.
//
//	@param data zip数据.
//	@param pFileName 字体文件名.
//	@param pPassword zip密码.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式: xcc.FontStyle_.
//	@return int 返回炫彩字体句柄.
func X字体_创建从内存ZIP(zip数据 []byte, 字体文件名, zip密码 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) int {
	r, _, _ := xFont_CreateFromZipMem.Call(炫彩工具类.ByteSliceDataPtr(&zip数据), 炫彩工具类.StrPtr(字体文件名), 炫彩工具类.StrPtr(zip密码), uintptr(字体大小), uintptr(字体样式))
	return int(r)
}

// XFont_CreateFromMem 字体_创建从内存. 创建炫彩字体从内存.
//
//	@param data 字体文件数据.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return int 返回字体句柄.
func X字体_创建从内存(字体文件数据 []byte, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) int {
	r, _, _ := xFont_CreateFromMem.Call(炫彩工具类.ByteSliceDataPtr(&字体文件数据), uintptr(len(字体文件数据)), uintptr(字体大小), uintptr(字体样式))
	return int(r)
}

// XFont_CreateFromRes 字体_创建从资源. 创建字体从资源.
//
//	@param id xx.
//	@param pType xx.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@param hModule xx.
//	@return int 返回炫彩字体句柄.
func X字体_创建从资源(id int32, 类型 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_, hModule uintptr) int {
	r, _, _ := xFont_CreateFromRes.Call(uintptr(id), 炫彩工具类.StrPtr(类型), uintptr(字体大小), uintptr(字体样式), hModule)
	return int(r)
}

// XFont_EnableAutoDestroy 字体_启用自动销毁. 是否自动销毁.
//
//	@param hFontX 字体句柄.
//	@param bEnable 是否启用.
//	@return int
func X字体_启用自动销毁(字体句柄 int, 是否启用 bool) int {
	r, _, _ := xFont_EnableAutoDestroy.Call(uintptr(字体句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// XFont_GetFont 字体_取Font. 获取字体.
//
//	@param hFontX 字体句柄.
//	@return int 返回GDI+ Font指针.
func X字体_取Font(字体句柄 int) int {
	r, _, _ := xFont_GetFont.Call(uintptr(字体句柄))
	return int(r)
}

// XFont_GetFontInfo 字体_取信息. 获取字体信息.
//
//	@param hFontX 字体句柄.
//	@param pInfo 接收返回的字体信息.
//	@return int
func X字体_取信息(字体句柄 int, 接收返回的字体信息 *Font_Info_) int {
	r, _, _ := xFont_GetFontInfo.Call(uintptr(字体句柄), uintptr(unsafe.Pointer(接收返回的字体信息)))
	return int(r)
}

// XFont_GetLOGFONTW 字体_取LOGFONTW. 获取字体LOGFONTW.
//
//	@param hFontX 字体句柄.
//	@param hdc hdc句柄.
//	@param pOut 接收返回信息.
//	@return bool
func X字体_取LOGFONTW(字体句柄 int, hdc句柄 uintptr, 接收返回信息 *LOGFONTW) bool {
	r, _, _ := xFont_GetLOGFONTW.Call(uintptr(字体句柄), hdc句柄, uintptr(unsafe.Pointer(接收返回信息)))
	return r != 0
}

// XFont_Destroy 字体_销毁. 强制销毁炫彩字体, 谨慎使用, 建议使用 XFont_Release() 释放.
//
//	@param hFontX 字体句柄.
//	@return int
func X字体_销毁(字体句柄 int) {
	xFont_Destroy.Call(uintptr(字体句柄))
}

// XFont_AddRef 字体_增加引用计数.
//
//	@param hFontX 字体句柄.
//	@return int
func X字体_增加引用计数(字体句柄 int) {
	xFont_AddRef.Call(uintptr(字体句柄))
}

// XFont_GetRefCount 字体_取引用计数.
//
//	@param hFontX 字体句柄.
//	@return int
func X字体_取引用计数(字体句柄 int) int32 {
	r, _, _ := xFont_GetRefCount.Call(uintptr(字体句柄))
	return int32(r)
}

// XFont_Release 字体_释放引用计数. 释放引用计数, 当引用计数为0时自动销毁.
//
//	@param hFontX 字体句柄.
//	@return int
func X字体_释放引用计数(字体句柄 int) {
	xFont_Release.Call(uintptr(字体句柄))
}
