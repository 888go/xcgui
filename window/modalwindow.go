package window

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ModalWindow 模态窗口.
type ModalWindow struct {
	windowBase
}

// 模态窗口_创建, 创建模态窗口, 然后你需要调用DoModal()来显示窗口; 当模态窗口关闭时, 会自动销毁模态窗口资源句柄.
//.
//.
//.
//.
//: Window_Style_.
// ff:创建模态窗口
// nWidth:宽度
// nHeight:高度
// pTitle:标题
// hWndParent:父窗口句柄
// XCStyle:炫彩窗口样式
func NewModalWindow(nWidth, nHeight int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(xc.XModalWnd_Create(nWidth, nHeight, pTitle, hWndParent, XCStyle))
	return p
}

// 模态窗口_创建扩展, 创建模态窗口, 增强功能.
//.
//.
//.
//.
//.
//.
//.
//.
//.
//: Window_Style_.
// ff:创建模态窗口EX
// dwExStyle:扩展样式
// dwStyle:样式
// lpClassName:类名
// x:x坐标
// y:y坐标
// cx:宽度
// cy:高度
// pTitle:窗口名
// hWndParent:父窗口
// XCStyle:GUI库窗口样式
func NewModalWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(xc.XModalWnd_CreateEx(dwExStyle, dwStyle, pTitle, x, y, cx, cy, lpClassName, hWndParent, XCStyle))
	return p
}

// NewModalWindowByLayout 从布局文件创建对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *ModalWindow
// ff:创建模态窗口并按布局文件
// pFileName:布局文件名
// hParent:父对象句柄
// hAttachWnd:附加窗口句柄
func NewModalWindowByLayout(pFileName string, hParent int, hAttachWnd uintptr) *ModalWindow {
	handle := xc.XC_LoadLayout(pFileName, hParent, hAttachWnd)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewModalWindowByLayoutZip 从压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *ModalWindow
// ff:创建模态窗口并按压缩包布局文件
// pZipFileName:zip文件名
// pFileName:布局文件名
// pPassword:zip密码
// hParent:父对象句柄
// hAttachWnd:附加窗口句柄
func NewModalWindowByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *ModalWindow {
	handle := xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewModalWindowByLayoutZipMem 从内存压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *ModalWindow
// ff:创建模态窗口并按内存压缩包布局文件
// data:布局文件数据
// pFileName:布局文件名
// pPassword:zip密码
// hParent:父对象句柄
// hAttachWnd:附加窗口句柄
func NewModalWindowByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *ModalWindow {
	handle := xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewModalWindowByLayoutStringW 从布局文件字符串W创建对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param hParent 父对象.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
// ff:创建模态窗口并按布局文件字符串W
// pStringXML:字符串
// hParent:父对象
// hAttachWnd:附加窗口句柄
func NewModalWindowByLayoutStringW(pStringXML string, hParent int, hAttachWnd uintptr) *ModalWindow {
	handle := xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewModalWindowByLayoutEx 从布局文件创建对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *ModalWindow
// ff:创建模态窗口并按布局文件EX
// pFileName:布局文件名
// pPrefixName:名称前缀
// hParent:父对象句柄
// hParentWnd:父窗口句柄
// hAttachWnd:附加窗口句柄
func NewModalWindowByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {
	handle := xc.XC_LoadLayoutEx(pFileName, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewModalWindowByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param id RC资源ID.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@param hModule 模块句柄, 可填0.
//	@return *ModalWindow
// ff:创建模态窗口并按RC资源zip压缩包布局文件EX
// id:RC资源ID
// pFileName:布局文件名
// pPassword:zip密码
// pPrefixName:名称前缀
// hParent:父对象句柄
// hParentWnd:父窗口句柄
// hAttachWnd:附加窗口句柄
// hModule:模块句柄
func NewModalWindowByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *ModalWindow {
	handle := xc.XC_LoadLayoutZipResEx(id, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewModalWindowByLayoutZipEx 从压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *ModalWindow
// ff:创建模态窗口并按压缩包布局文件EX
// pZipFileName:zip文件名
// pFileName:布局文件名
// pPassword:zip密码
// pPrefixName:名称前缀
// hParent:父对象句柄
// hParentWnd:父窗口句柄HWND
// hAttachWnd:附加窗口句柄
func NewModalWindowByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {
	handle := xc.XC_LoadLayoutZipEx(pZipFileName, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewModalWindowByLayoutZipMemEx 从内存压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *ModalWindow
// ff:创建模态窗口并按内存压缩包布局文件EX
// data:布局文件数据
// pFileName:布局文件名
// pPassword:zip密码
// pPrefixName:名称前缀
// hParent:父对象句柄
// hParentWnd:父窗口句柄HWND
// hAttachWnd:附加窗口句柄
func NewModalWindowByLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {
	handle := xc.XC_LoadLayoutZipMemEx(data, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewModalWindowByLayoutStringWEx 从布局文件字符串W创建对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
// ff:创建模态窗口并按布局文件字符串WEX
// pStringXML:字符串
// pPrefixName:名称前缀
// hParent:父对象
// hParentWnd:父窗口句柄HWND
// hAttachWnd:附加窗口句柄
func NewModalWindowByLayoutStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {
	handle := xc.XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 模态窗口_附加窗口, 返回窗口对象.
//.
//: xcc.Window_Style_.
// ff:模态窗口附加窗口
// hWnd:外部窗口句柄
// XCStyle:炫彩窗口样式
func ModalWnd_Attach(hWnd uintptr, XCStyle xcc.Window_Style_) *Window {
	p := &Window{}
	p.SetHandle(xc.XModalWnd_Attach(hWnd, XCStyle))
	return p
}

// 从句柄创建对象.
// ff:创建模态窗口并按句柄
// handle:
func NewModalWindowByHandle(handle int) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
// ff:创建模态窗口并按名称
// name:
func NewModalWindowByName(name string) *ModalWindow {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
// ff:创建模态窗口并按UID
// nUID:
func NewModalWindowByUID(nUID int) *ModalWindow {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
// ff:创建模态窗口并按UID名称
// name:
func NewModalWindowByUIDName(name string) *ModalWindow {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ModalWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 模态窗口_启用自动关闭, 是否自动关闭窗口, 当窗口失去焦点时.
//.
// ff:启用自动关闭
// bEnable:开启开关
func (m *ModalWindow) EnableAutoClose(bEnable bool) int {
	return xc.XModalWnd_EnableAutoClose(m.Handle, bEnable)
}

// 模态窗口_启用ESC关闭, 当用户按ESC键时自动关闭模态窗口.
//.
// ff:启用ESC关闭
// bEnable:是否启用
func (m *ModalWindow) EnableEscClose(bEnable bool) int {
	return xc.XModalWnd_EnableEscClose(m.Handle, bEnable)
}

// 模态窗口_启动, 启动显示模态窗口, 当窗口关闭时返回: MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出.
// ff:启动
func (m *ModalWindow) DoModal() xcc.MessageBox_Flag_ {
	return xc.XModalWnd_DoModal(m.Handle)
}

// 模态窗口_结束, 结束模态窗口.
//()的返回值. MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出.
// ff:结束
// nResult:结果
func (m *ModalWindow) EndModal(nResult xcc.MessageBox_Flag_) int {
	return xc.XModalWnd_EndModal(m.Handle, nResult)
}
