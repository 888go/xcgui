package 炫彩字体类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/res"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Font 炫彩字体.
type Font struct {
	炫彩对象基类.ObjectBase
}

// New 字体_创建, 创建炫彩字体. 当字体句柄与元素关联后, 会自动释放.
//
//	@param size 字体大小,单位(pt,磅).
//	@return *Font 返回字体对象.
func X创建(字体大小 int32) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建(字体大小))
	return p
}

// NewEX 字体_创建扩展. 创建炫彩字体.
//
//	@param pName 字体名称.
//	@param size 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return *Font 返回字体对象.
func X创建EX(字体名称 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建EX(字体名称, 字体大小, 字体样式))
	return p
}

// NewLOGFONTW 字体_创建从LOGFONT. 创建炫彩字体.
//
//	@param pFontInfo 字体信息.
//	@return *Font 返回字体对象.
func X创建并按选项(字体选项 *炫彩基类.LOGFONTW) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建从LOGFONT(字体选项))
	return p
}

// NewByHFONT 字体_创建从HFONT. 创建炫彩字体从现有HFONT字体.
//
//	@param hFont 字体句柄.
//	@return *Font 返回字体对象.
func X创建并按字体句柄(字体句柄 uintptr) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建从HFONT(字体句柄))
	return p
}

// NewByFont 字体_创建从Font. 创建炫彩字体从GDI+字体.
//
//	@param pFont GDI+字体指针.
//	@return *Font 返回字体对象.
func X创建并按GDI字体指针(GDI字体指针 uintptr) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建从Font(GDI字体指针))
	return p
}

// NewByFile 字体_创建从文件. 创建字体从文件.
//
//	@param pFontFile 字体文件名.
//	@param size 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return *Font 返回字体对象.
func X创建并按文件(字体文件名 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建从文件(字体文件名, 字体大小, 字体样式))
	return p
}

// NewByZip 字体_创建从ZIP.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 字体文件名.
//	@param pPassword zip密码.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式: xcc.FontStyle_.
//	@return *Font 返回炫彩字体对象.
func X创建并按ZIP(zip文件名, 字体文件名, zip密码 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建从ZIP(zip文件名, 字体文件名, zip密码, 字体大小, 字体样式))
	return p
}

// NewByZipMem 字体_创建从内存ZIP.
//
//	@param data zip数据.
//	@param pFileName 字体文件名.
//	@param pPassword zip密码.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式: xcc.FontStyle_.
//	@return *Font 返回炫彩字体对象.
func X创建并按内存ZIP(zip数据 []byte, 字体文件名, zip密码 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建从内存ZIP(zip数据, 字体文件名, zip密码, 字体大小, 字体样式))
	return p
}

// NewByMem 字体_创建从内存. 创建炫彩字体从内存.
//
//	@param data 字体文件数据.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return *Font 返回字体对象.
func X创建并按内存(字体文件数据 []byte, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建从内存(字体文件数据, 字体大小, 字体样式))
	return p
}

// NewByRes 字体_创建从资源. 创建字体从资源.
//
//	@param id xx.
//	@param pType xx.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@param hModule xx.
//	@return *Font 返回字体对象.
func X创建并按资源(id int32, 类型 string, 字体大小 int32, 字体样式 炫彩常量类.FontStyle_, 模块 uintptr) *Font {
	p := &Font{}
	p.X设置句柄(炫彩基类.X字体_创建从资源(id, 类型, 字体大小, 字体样式, 模块))
	return p
}

// NewByHandle 从句柄创建对象.
//
//	@param handle
//	@return *Font
func X创建并按句柄(句柄 int) *Font {
	p := &Font{}
	p.X设置句柄(句柄)
	return p
}

// NewByName 根据资源文件中的name创建对象, 失败返回nil.
//
//	@param name
//	@return *Font
func X创建并按名称(名称 string) *Font {
	handle := 炫彩资源类.X取字体(名称)
	if handle > 0 {
		p := &Font{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// EnableAutoDestroy 字体_启用自动销毁. 是否自动销毁.
//
//	@param bEnable 是否启用.
//	@return int
func (f *Font) X启用自动销毁(是否启用 bool) int {
	return 炫彩基类.X字体_启用自动销毁(f.Handle, 是否启用)
}

// GetFont 字体_取Font. 获取字体.
//
//	@return int 返回GDI+ Font指针
func (f *Font) X取字体指针() int {
	return 炫彩基类.X字体_取Font(f.Handle)
}

// GetFontInfo 字体_取信息. 获取字体信息.
//
//	@param pInfo 接收返回的字体信息.
//	@return int
func (f *Font) X取信息(返回字体信息 *炫彩基类.Font_Info_) int {
	return 炫彩基类.X字体_取信息(f.Handle, 返回字体信息)
}

// GetLOGFONTW 字体_取LOGFONTW. 获取字体LOGFONTW.
//
//	@param hdc hdc句柄.
//	@param pOut 接收返回信息.
//	@return bool
func (f *Font) X取属性(hdc句柄 uintptr, 返回字体属性 *炫彩基类.LOGFONTW) bool {
	return 炫彩基类.X字体_取LOGFONTW(f.Handle, hdc句柄, 返回字体属性)
}

// Destroy 字体_销毁. 强制销毁炫彩字体, 谨慎使用, 建议使用 Release() 释放.
//
//	@return int
func (f *Font) X销毁() {
	炫彩基类.X字体_销毁(f.Handle)
}

// AddRef 字体_增加引用计数.
//
//	@return int
func (f *Font) X增加引用计数() {
	炫彩基类.X字体_增加引用计数(f.Handle)
}

// GetRefCount 字体_取引用计数.
//
//	@return int
func (f *Font) X取引用计数() int32 {
	return 炫彩基类.X字体_取引用计数(f.Handle)
}

// Release 字体_释放引用计数. 释放引用计数, 当引用计数为0时自动销毁.
//
//	@return int
func (f *Font) X释放引用计数() {
	炫彩基类.X字体_释放引用计数(f.Handle)
}
