package 炫彩窗口基类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Window 普通窗口.
type Window struct {
	windowBase
}

// New 窗口_创建.
//
//	@param x 窗口左上角x坐标.
//	@param y 窗口左上角y坐标.
//	@param cx 窗口宽度.
//	@param cy 窗口高度.
//	@param pTitle 窗口标题.
//	@param hWndParent 父窗口.
//	@param XCStyle 窗口样式: xcc.Window_Style_.
//	@return *Window
func X创建窗口(x坐标, y坐标, 窗口宽度, 窗口高度 int32, 窗口标题 string, 父窗口 uintptr, 窗口样式 炫彩常量类.Window_Style_) *Window {
	p := &Window{}
	p.X设置句柄(炫彩基类.X窗口_创建(x坐标, y坐标, 窗口宽度, 窗口高度, 窗口标题, 父窗口, 窗口样式))
	return p
}

// NewEx 窗口_创建扩展.
//
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param x 窗口左上角x坐标.
//	@param y 窗口左上角y坐标.
//	@param cx 窗口宽度.
//	@param cy 窗口高度.
//	@param pTitle 窗口名.
//	@param hWndParent 父窗口.
//	@param XCStyle 窗口样式, xcc.Window_Style_.
//	@return *Window
func X创建窗口EX(窗口扩展样式 int, 样式 int, 类名 string, x坐标 int, y坐标 int, 宽度 int, 高度 int, 窗口名 string, 父窗口 uintptr, 窗口样式 炫彩常量类.Window_Style_) *Window {
	p := &Window{}
	p.X设置句柄(炫彩基类.X窗口_创建EX(窗口扩展样式, 样式, 类名, x坐标, y坐标, 宽度, 高度, 窗口名, 父窗口, 窗口样式))
	return p
}

// Attach 窗口_附加窗口, 返回窗口对象.
//
//	@param hWnd 要附加的外部窗口句柄.
//	@param XCStyle 窗口样式: xcc.Window_Style_.
//	@return *Window
func X窗口附加窗口(外部窗口句柄 uintptr, XCStyle 炫彩常量类.Window_Style_) *Window {
	p := &Window{}
	p.X设置句柄(炫彩基类.X窗口_附加窗口(外部窗口句柄, XCStyle))
	return p
}

// NewByHandle 从句柄创建窗口对象.
//
//	@param hWindow 窗口资源句柄.
//	@return *Window
func X创建窗口并按句柄(窗口资源句柄 int) *Window {
	p := &Window{}
	p.X设置句柄(窗口资源句柄)
	return p
}

// NewByLayout 从布局文件创建窗口对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func X创建窗口并按布局文件(布局文件名 string, 父对象句柄 int, 附加窗口句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件(布局文件名, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByLayoutZip 从压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func X创建窗口并按压缩包布局文件(zip文件名 string, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件ZIP(zip文件名, 布局文件名, zip密码, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByLayoutZipMem 从内存压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func X创建窗口并按内存压缩包布局文件(布局文件数据 []byte, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件内存ZIP(布局文件数据, 布局文件名, zip密码, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByLayoutStringW 从布局文件字符串W创建窗口对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param hParent 父对象.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func X创建窗口并按布局文件字符串W(字符串 string, 父对象 int, 附加窗口句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件从字符串W(字符串, 父对象, 附加窗口句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByLayoutEx 从布局文件创建窗口对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func X创建窗口并按布局文件EX(布局文件名, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件Ex(布局文件名, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByLayoutZipEx 从压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func X创建窗口并按压缩包布局文件EX(zip文件名 string, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件ZIPEx(zip文件名, 布局文件名, zip密码, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param id RC资源ID.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@param hModule 模块句柄, 可填0.
//	@return *Window
func X创建窗口并按RC资源zip压缩包布局文件EX(RC资源ID int32, 布局文件名, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄, 模块句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件资源ZIPEX(RC资源ID, 布局文件名, zip密码, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄, 模块句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByLayoutZipMemEx 从内存压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func X创建窗口并按内存压缩包布局文件EX(布局文件数据 []byte, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件内存ZIPEx(布局文件数据, 布局文件名, zip密码, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByLayoutStringWEx 从布局文件字符串W创建窗口对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func X创建窗口并按布局文件字符串WEX(字符串, 名称前缀 string, 父对象 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *Window {
	handle := 炫彩基类.X炫彩_加载布局文件从字符串WEx(字符串, 名称前缀, 父对象, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByName 从name创建对象, 失败返回nil.
//
//	@param name
//	@return *Window
func X创建窗口并按名称(名称 string) *Window {
	handle := 炫彩基类.X取对象从名称(名称)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByUID 从UID创建窗口对象, 失败返回nil.
//
//	@param nUID
//	@return *Window
func X创建窗口并按UID(nUID int) *Window {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewByUIDName 从UID名称创建对象, 失败返回nil.
//
//	@param name
//	@return *Window
func X创建窗口并按UID名称(名称 string) *Window {
	handle := 炫彩基类.X取对象从UID名称(名称)
	if handle > 0 {
		p := &Window{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}
