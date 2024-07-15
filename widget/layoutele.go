package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// LayoutEle 布局元素.
type LayoutEle struct {
	Element
}

// NewLayoutEle 布局_创建, 创建布局元素.
//
//	@param x 元素x坐标.
//	@param y 元素y坐标.
//	@param cx 宽度.
//	@param cy 高度.
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutEle

// ff:创建布局
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X创建布局(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *LayoutEle {
	p := &LayoutEle{}
	p.X设置句柄(炫彩基类.X布局_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// NewLayoutEleEx 布局_创建扩展, 创建布局元素.
//
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutEle

// ff:创建布局EX
// hParent:父窗口句柄或元素句柄
func X创建布局EX(父窗口句柄或元素句柄 int) *LayoutEle {
	p := &LayoutEle{}
	p.X设置句柄(炫彩基类.X布局_创建EX(父窗口句柄或元素句柄))
	return p
}

// NewLayoutEleByHandle 从句柄创建对象.
//
//	@param handle
//	@return *LayoutEle

// ff:创建布局并按句柄
// handle:
func X创建布局并按句柄(handle int) *LayoutEle {
	p := &LayoutEle{}
	p.X设置句柄(handle)
	return p
}

// NewLayoutEleByLayout 从布局文件创建对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle

// ff:创建布局并按布局文件
// hAttachWnd:附加窗口句柄
// hParent:父对象句柄
// pFileName:布局文件名
func X创建布局并按布局文件(布局文件名 string, 父对象句柄 int, 附加窗口句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件(布局文件名, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZip 从压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle

// ff:创建布局并按压缩包中布局文件
// hAttachWnd:附加窗口句柄
// hParent:父对象句柄
// pPassword:zip密码
// pFileName:布局文件名
// pZipFileName:zip文件名
func X创建布局并按压缩包中布局文件(zip文件名 string, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件ZIP(zip文件名, 布局文件名, zip密码, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param id RC资源ID.
//	@param pFileName zip文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@param hModule 模块句柄, 可填0.
//	@return *LayoutEle

// ff:创建布局并按RC资源压缩包中布局文件
// hModule:模块句柄
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄HWND
// hParent:父对象句柄
// pPrefixName:名称
// pPassword:zip密码
// pFileName:布局文件名
// id:RC资源ID
func X创建布局并按RC资源压缩包中布局文件(RC资源ID int32, 布局文件名, zip密码, 名称 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄, 模块句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件资源ZIPEX(RC资源ID, 布局文件名, zip密码, 名称, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄, 模块句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZipMem 从内存压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle

// ff:创建布局并按内存压缩包中布局文件
// hAttachWnd:附加窗口句柄
// hParent:父对象句柄
// pPassword:zip密码
// pFileName:布局文件名
// data:布局文件数据
func X创建布局并按内存压缩包中布局文件(布局文件数据 []byte, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件内存ZIP(布局文件数据, 布局文件名, zip密码, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByStringW 从布局文件字符串W创建对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param hParent 父对象.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle

// ff:创建布局并按布局文件字符串W
// hAttachWnd:附加窗口句柄
// hParent:父对象
// pStringXML:xml字符串
func X创建布局并按布局文件字符串W(xml字符串 string, 父对象 int, 附加窗口句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件从字符串W(xml字符串, 父对象, 附加窗口句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutEx 从布局文件创建对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle

// ff:创建布局并按布局文件EX
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄HWND
// hParent:父对象句柄
// pPrefixName:名称前缀
// pFileName:布局文件名
func X创建布局并按布局文件EX(布局文件名, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件Ex(布局文件名, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZipEx 从压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle

// ff:创建布局并按压缩包中的布局文件EX
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄HWND
// hParent:父对象句柄
// pPrefixName:名称前缀
// pPassword:zip密码
// pFileName:布局文件名
// pZipFileName:zip文件名
func X创建布局并按压缩包中的布局文件EX(zip文件名 string, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件ZIPEx(zip文件名, 布局文件名, zip密码, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZipMemEx 从内存压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle

// ff:创建布局并按内存压缩包中布局文件EX
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄HWND
// hParent:父对象句柄
// pPrefixName:名称
// pPassword:zip密码
// pFileName:布局文件名
// data:布局文件数据
func X创建布局并按内存压缩包中布局文件EX(布局文件数据 []byte, 布局文件名 string, zip密码, 名称 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件内存ZIPEx(布局文件数据, 布局文件名, zip密码, 名称, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByStringWEx 从布局文件字符串W创建对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle

// ff:创建布局并按布局文件字符串WEX
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄HWND
// hParent:父对象
// pPrefixName:名称前缀
// pStringXML:字符串
func X创建布局并按布局文件字符串WEX(字符串, 名称前缀 string, 父对象 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *LayoutEle {
	handle := 炫彩基类.X炫彩_加载布局文件从字符串WEx(字符串, 名称前缀, 父对象, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByName 从name创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutEle

// ff:创建布局并按名称
// name:
func X创建布局并按名称(name string) *LayoutEle {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByID 从ID创建对象, 失败返回nil.
//
//	@param hWindow 父窗口句柄
//	@param nID
//	@return *LayoutEle

// ff:创建布局并按ID
// nID:
// hWindow:父窗口句柄
func X创建布局并按ID(父窗口句柄, nID int) *LayoutEle {
	handle := 炫彩基类.X取对象从ID(父窗口句柄, nID)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByIDName 从ID名称创建对象, 失败返回nil.
//
//	@param hWindow 父窗口句柄
//	@param name
//	@return *LayoutEle

// ff:创建布局并按ID名称
// name:名称
// hWindow:父窗口句柄
func X创建布局并按ID名称(父窗口句柄 int, 名称 string) *LayoutEle {
	handle := 炫彩基类.X取对象从ID名称(父窗口句柄, 名称)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByUID 从UID创建对象, 失败返回nil.
//
//	@param nUID
//	@return *LayoutEle

// ff:创建布局并按UID
// nUID:
func X创建布局并按UID(nUID int) *LayoutEle {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutEleByUIDName 从UID名称创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutEle

// ff:创建布局并按UID名称
// name:名称
func X创建布局并按UID名称(名称 string) *LayoutEle {
	handle := 炫彩基类.X取对象从UID名称(名称)
	if handle > 0 {
		p := &LayoutEle{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// IsEnableLayout 布局_判断启用, 是否已经启用布局功能.
//
//	@return bool

// ff:判断启用
func (l *LayoutEle) X判断启用() bool {
	return 炫彩基类.X布局_判断启用(l.Handle)
}

// EnableLayout 布局_启用, 启用布局功能.
//
//	@param bEnable 是否启用.
//	@return int

// ff:启用
// bEnable:是否启用
func (l *LayoutEle) X启用(是否启用 bool) int {
	return 炫彩基类.X布局_启用(l.Handle, 是否启用)
}

// ShowLayoutFrame 布局_显示布局边界, 显示布局边界.
//
//	@param bEnable 是否显示.
//	@return int

// ff:显示布局边界
// bEnable:是否显示
func (l *LayoutEle) X显示布局边界(是否显示 bool) int {
	return 炫彩基类.X布局_显示布局边界(l.Handle, 是否显示)
}

// GetWidthIn 布局_取内宽度, 获取宽度,不包含内边距大小.
//
//	@return int

// ff:取内宽度
func (l *LayoutEle) X取内宽度() int {
	return 炫彩基类.X布局_取内宽度(l.Handle)
}

// GetHeightIn 布局_取内高度, 获取高度,不包含内边距大小.
//
//	@return int

// ff:取内高度
func (l *LayoutEle) X取内高度() int {
	return 炫彩基类.X布局_取内高度(l.Handle)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
//	@param bEnable 是否启用.
//	@return int

// ff:启用水平
// bEnable:是否启用
func (l *LayoutEle) X启用水平(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用水平(l.Handle, 是否启用)
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
//	@param bEnable 是否启用.
//	@return int

// ff:启用自动换行
// bEnable:是否启用
func (l *LayoutEle) X启用自动换行(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用自动换行(l.Handle, 是否启用)
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
//	@param bEnable 是否启用.
//	@return int

// ff:启用溢出隐藏
// bEnable:是否启用
func (l *LayoutEle) X启用溢出隐藏(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用溢出隐藏(l.Handle, 是否启用)
}

// SetAlignH 布局盒子_置水平对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int

// ff:置水平对齐
// nAlign:对齐方式
func (l *LayoutEle) X置水平对齐(对齐方式 炫彩常量类.Layout_Align_) int {
	return 炫彩基类.X布局盒子_置水平对齐(l.Handle, 对齐方式)
}

// SetAlignV 布局盒子_置垂直对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int

// ff:置垂直对齐
// nAlign:对齐方式
func (l *LayoutEle) X置垂直对齐(对齐方式 炫彩常量类.Layout_Align_) int {
	return 炫彩基类.X布局盒子_置垂直对齐(l.Handle, 对齐方式)
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_Axis_.
//	@return int

// ff:置对齐基线
// nAlign:对齐方式
func (l *LayoutEle) X置对齐基线(对齐方式 炫彩常量类.Layout_Align_Axis_) int {
	return 炫彩基类.X布局盒子_置对齐基线(l.Handle, 对齐方式)
}

// SetSpace 布局盒子_置间距.
//
//	@param nSpace 项间距大小.
//	@return int

// ff:置间距
// nSpace:项间距大小
func (l *LayoutEle) X置间距(项间距大小 int) int {
	return 炫彩基类.X布局盒子_置间距(l.Handle, 项间距大小)
}

// SetSpaceRow 布局盒子_置行距.
//
//	@param nSpace 行间距大小.
//	@return int

// ff:置行距
// nSpace:行间距大小
func (l *LayoutEle) X置行距(行间距大小 int) int {
	return 炫彩基类.X布局盒子_置行距(l.Handle, 行间距大小)
}
