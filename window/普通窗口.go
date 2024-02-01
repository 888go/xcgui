package window

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// Window 普通窗口.
type Window struct {
	windowBase
}

// I窗口_创建 窗口_创建.
//	@param x 窗口左上角x坐标.
//	@param y 窗口左上角y坐标.
//	@param cx 窗口宽度.
//	@param cy 窗口高度.
//	@param pTitle 窗口标题.
//	@param hWndParent 父窗口.
//	@param XCStyle 窗口样式: xcc.I常量_窗口样式_.
//	@return *Window
//
func I窗口_创建(x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle xcc.I常量_窗口样式_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// I窗口_创建扩展 窗口_创建扩展.
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param x 窗口左上角x坐标.
//	@param y 窗口左上角y坐标.
//	@param cx 窗口宽度.
//	@param cy 窗口高度.
//	@param pTitle 窗口名.
//	@param hWndParent 父窗口.
//	@param XCStyle 窗口样式, xcc.I常量_窗口样式_.
//	@return *Window
//
func I窗口_创建扩展(dwExStyle int, dwStyle int, lpClassName string, x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle xcc.I常量_窗口样式_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_CreateEx(dwExStyle, dwStyle, lpClassName, x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// I窗口_附加窗口 窗口_附加窗口, 返回窗口对象.
//	@param hWnd 要附加的外部窗口句柄.
//	@param XCStyle 窗口样式: xcc.I常量_窗口样式_.
//	@return *Window
//
func I窗口_附加窗口(hWnd int, XCStyle xcc.I常量_窗口样式_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_Attach(hWnd, XCStyle))
	return p
}

// I窗口_从句柄创建 从句柄创建对象.
//	@param hWindow 窗口资源句柄.
//	@return *Window
//
func I窗口_从句柄创建(hWindow int) *Window {
	p := &Window{}
	p.SetHandle(hWindow)
	return p
}

// I窗口_布局文件创建 从布局文件创建对象, 失败返回nil.
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
//
func I窗口_布局文件创建(pFileName string, hParent, hAttachWnd int) *Window {
	handle := xc.XC_LoadLayout(pFileName, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I窗口_从zip布局文件创建 从压缩包中的布局文件创建对象, 失败返回nil.
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
//
func I窗口_从zip布局文件创建(pZipFileName string, pFileName string, pPassword string, hParent, hAttachWnd int) *Window {
	handle := xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I窗口_从内存zip布局文件创建 从内存压缩包中的布局文件创建对象, 失败返回nil.
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
//
func I窗口_从内存zip布局文件创建(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd int) *Window {
	handle := xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I窗口_从布局文件字符串W创建 从布局文件字符串W创建对象, 失败返回nil.
//	@param pStringXML 字符串.
//	@param hParent 父对象.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
//
func I窗口_从布局文件字符串W创建(pStringXML string, hParent int, hAttachWnd int) *Window {
	handle := xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I窗口_从name创建 从name创建对象, 失败返回nil.
//	@param name
//	@return *Window
//
func I窗口_从name创建(name string) *Window {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I窗口_从UID创建 从UID创建对象, 失败返回nil.
//	@param nUID
//	@return *Window
//
func I窗口_从UID创建(nUID int) *Window {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I窗口_从UID名称创建 从UID名称创建对象, 失败返回nil.
//	@param name
//	@return *Window
//
func I窗口_从UID名称创建(name string) *Window {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

/*
LayoutBox-布局盒子
*/

// I布局盒子_启用水平 布局盒子_启用水平.
//	@param bEnable 是否启用.
//	@return int
//
func (w *Window) I布局盒子_启用水平(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(w.I句柄, bEnable)
}

// I布局盒子_启用自动换行 布局盒子_启用自动换行.
//	@param bEnable 是否启用.
//	@return int
//
func (w *Window) I布局盒子_启用自动换行(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(w.I句柄, bEnable)
}

// I布局盒子_启用溢出隐藏 布局盒子_启用溢出隐藏.
//	@param bEnable 是否启用.
//	@return int
//
func (w *Window) I布局盒子_启用溢出隐藏(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(w.I句柄, bEnable)
}

// I布局盒子_置水平对齐 布局盒子_置水平对齐.
//	@param nAlign 对齐方式: xcc.I常量_布局对齐_.
//	@return int
//
func (w *Window) I布局盒子_置水平对齐(nAlign xcc.I常量_布局对齐_) int {
	return xc.XLayoutBox_SetAlignH(w.I句柄, nAlign)
}

// I布局盒子_置垂直对齐 布局盒子_置垂直对齐.
//	@param nAlign 对齐方式: xcc.I常量_布局对齐_.
//	@return int
//
func (w *Window) I布局盒子_置垂直对齐(nAlign xcc.I常量_布局对齐_) int {
	return xc.XLayoutBox_SetAlignV(w.I句柄, nAlign)
}

// I布局盒子_置对齐基线 布局盒子_置对齐基线.
//	@param nAlign 对齐方式: xcc.I常量_布局轴对齐_.
//	@return int
//
func (w *Window) I布局盒子_置对齐基线(nAlign xcc.I常量_布局轴对齐_) int {
	return xc.XLayoutBox_SetAlignBaseline(w.I句柄, nAlign)
}

// I布局盒子_置间距 布局盒子_置间距.
//	@param nSpace 项间距大小.
//	@return int
//
func (w *Window) I布局盒子_置间距(nSpace int) int {
	return xc.XLayoutBox_SetSpace(w.I句柄, nSpace)
}

// I布局盒子_置行距 布局盒子_置行距.
//	@param nSpace 行间距大小.
//	@return int
//
func (w *Window) I布局盒子_置行距(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(w.I句柄, nSpace)
}
