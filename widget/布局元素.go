package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// LayoutEle 布局元素.
type LayoutEle struct {
	Element
}

// I布局_创建 I布局_创建, 创建布局元素.
//	@param x 元素x坐标.
//	@param y 元素y坐标.
//	@param cx 宽度.
//	@param cy 高度.
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutEle
//
func I布局_创建(x int, y int, cx int, cy int, hParent int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(xc.XLayout_Create(x, y, cx, cy, hParent))
	return p
}

// I布局_创建扩展 I布局_创建扩展, 创建布局元素.
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutEle
//
func I布局_创建扩展(hParent int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(xc.XLayout_CreateEx(hParent))
	return p
}

// I布局_从句柄创建 I布局_从句柄创建对象.
//	@param handle
//	@return *LayoutEle
//
func I布局_从句柄创建(handle int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(handle)
	return p
}

// I布局_从布局文件创建 I布局_从布局文件创建对象, 失败返回nil.
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
//
func I布局_从布局文件创建(pFileName string, hParent int, hAttachWnd int) *LayoutEle {
	handle := xc.XC_LoadLayout(pFileName, hParent, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I布局_从ZIP布局文件创建 I布局_从压缩包中的布局文件创建对象, 失败返回nil.
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
//
func I布局_从ZIP布局文件创建(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd int) *LayoutEle {
	handle := xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I布局_从内存ZIP布局文件创建 I布局_从内存压缩包中的布局文件创建对象, 失败返回nil.
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
//
func I布局_从内存ZIP布局文件创建(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd int) *LayoutEle {
	handle := xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I布局_从布局文件字符串W创建 I布局_从布局文件字符串W创建对象, 失败返回nil.
//	@param pStringXML 字符串.
//	@param hParent 父对象.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
//
func I布局_从布局文件字符串W创建(pStringXML string, hParent int, hAttachWnd int) *LayoutEle {
	handle := xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I布局_从name创建 I布局_从name创建对象, 失败返回nil.
//	@param name
//	@return *LayoutEle
//
func I布局_从name创建(name string) *LayoutEle {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I布局_从UID创建 I布局_从UID创建对象, 失败返回nil.
//	@param nUID
//	@return *LayoutEle
//
func I布局_从UID创建(nUID int) *LayoutEle {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I布局_从UID名称创建 I布局_从UID名称创建对象, 失败返回nil.
//	@param name
//	@return *LayoutEle
//
func I布局_从UID名称创建(name string) *LayoutEle {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I判断启用 布局_判断启用, 是否已经启用布局功能.
//	@return bool
//
func (l *LayoutEle) I判断启用() bool {
	return xc.XLayout_IsEnableLayout(l.I句柄)
}

// I启用 布局_启用, 启用布局功能.
//	@param bEnable 是否启用.
//	@return int
//
func (l *LayoutEle) I启用(bEnable bool) int {
	return xc.XLayout_EnableLayout(l.I句柄, bEnable)
}

// I显示布局边界 布局_显示布局边界, 显示布局边界.
//	@param bEnable 是否显示.
//	@return int
//
func (l *LayoutEle) I显示布局边界(bEnable bool) int {
	return xc.XLayout_ShowLayoutFrame(l.I句柄, bEnable)
}

// I取内宽度 布局_取内宽度, 获取宽度,不包含内边距大小.
//	@return int
//
func (l *LayoutEle) I取内宽度() int {
	return xc.XLayout_GetWidthIn(l.I句柄)
}

// I取内高度 布局_取内高度, 获取高度,不包含内边距大小.
//	@return int
//
func (l *LayoutEle) I取内高度() int {
	return xc.XLayout_GetHeightIn(l.I句柄)
}

/*
LayoutBox-布局盒子
*/

// I盒子_启用水平 布局盒子_启用水平.
//	@param bEnable 是否启用.
//	@return int
//
func (l *LayoutEle) I盒子_启用水平(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(l.I句柄, bEnable)
}

// I盒子_启用自动换行 布局盒子_启用自动换行.
//	@param bEnable 是否启用.
//	@return int
//
func (l *LayoutEle) I盒子_启用自动换行(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(l.I句柄, bEnable)
}

// I盒子_启用溢出隐藏 布局盒子_启用溢出隐藏.
//	@param bEnable 是否启用.
//	@return int
//
func (l *LayoutEle) I盒子_启用溢出隐藏(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(l.I句柄, bEnable)
}

// SetAlignH 布局盒子_置水平对齐.
//	@param nAlign 对齐方式: xcc.I常量_布局对齐_.
//	@return int
//
func (l *LayoutEle) SetAlignH(nAlign xcc.I常量_布局对齐_) int {
	return xc.XLayoutBox_SetAlignH(l.I句柄, nAlign)
}

// I盒子_置垂直对齐 布局盒子_置垂直对齐.
//	@param nAlign 对齐方式: xcc.I常量_布局对齐_.
//	@return int
//
func (l *LayoutEle) I盒子_置垂直对齐(nAlign xcc.I常量_布局对齐_) int {
	return xc.XLayoutBox_SetAlignV(l.I句柄, nAlign)
}

// I盒子_置对齐基线 布局盒子_置对齐基线.
//	@param nAlign 对齐方式: xcc.I常量_布局轴对齐_.
//	@return int
//
func (l *LayoutEle) I盒子_置对齐基线(nAlign xcc.I常量_布局轴对齐_) int {
	return xc.XLayoutBox_SetAlignBaseline(l.I句柄, nAlign)
}

// I盒子_置间距 布局盒子_置间距.
//	@param nSpace 项间距大小.
//	@return int
//
func (l *LayoutEle) I盒子_置间距(nSpace int) int {
	return xc.XLayoutBox_SetSpace(l.I句柄, nSpace)
}

// I盒子_置行距 布局盒子_置行距.
//	@param nSpace 行间距大小.
//	@return int
//
func (l *LayoutEle) I盒子_置行距(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(l.I句柄, nSpace)
}
