package 炫彩窗口基类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// ModalWindow 模态窗口.
type ModalWindow struct {
	windowBase
}

// 模态窗口_创建, 创建模态窗口, 然后你需要调用DoModal()来显示窗口; 当模态窗口关闭时, 会自动销毁模态窗口资源句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// pTitle: 窗口标题内容.
//
// hWndParent: 父窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.

// ff:创建模态窗口
// XCStyle:炫彩窗口样式
// hWndParent:父窗口句柄
// pTitle:标题
// nHeight:高度
// nWidth:宽度
func X创建模态窗口(宽度, 高度 int32, 标题 string, 父窗口句柄 uintptr, 炫彩窗口样式 炫彩常量类.Window_Style_) *ModalWindow {
	p := &ModalWindow{}
	p.X设置句柄(炫彩基类.X模态窗口_创建(宽度, 高度, 标题, 父窗口句柄, 炫彩窗口样式))
	return p
}

// 模态窗口_创建扩展, 创建模态窗口, 增强功能.
//
// dwExStyle: 窗口扩展样式.
//
// dwStyle: 窗口样式.
//
// lpClassName: 窗口类名.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口名.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式: Window_Style_.

// ff:创建模态窗口EX
// XCStyle:GUI库窗口样式
// hWndParent:父窗口
// pTitle:窗口名
// cy:高度
// cx:宽度
// y:y坐标
// x:x坐标
// lpClassName:类名
// dwStyle:样式
// dwExStyle:扩展样式
func X创建模态窗口EX(扩展样式 int, 样式 int, 类名 string, x坐标, y坐标, 宽度, 高度 int32, 窗口名 string, 父窗口 uintptr, GUI库窗口样式 炫彩常量类.Window_Style_) *ModalWindow {
	p := &ModalWindow{}
	p.X设置句柄(炫彩基类.X模态窗口_创建EX(扩展样式, 样式, 窗口名, x坐标, y坐标, 宽度, 高度, 类名, 父窗口, GUI库窗口样式))
	return p
}

// NewModalWindowByLayout 从布局文件创建对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *ModalWindow

// ff:创建模态窗口并按布局文件
// hAttachWnd:附加窗口句柄
// hParent:父对象句柄
// pFileName:布局文件名
func X创建模态窗口并按布局文件(布局文件名 string, 父对象句柄 int, 附加窗口句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件(布局文件名, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
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
// hAttachWnd:附加窗口句柄
// hParent:父对象句柄
// pPassword:zip密码
// pFileName:布局文件名
// pZipFileName:zip文件名
func X创建模态窗口并按压缩包布局文件(zip文件名 string, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件ZIP(zip文件名, 布局文件名, zip密码, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
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
// hAttachWnd:附加窗口句柄
// hParent:父对象句柄
// pPassword:zip密码
// pFileName:布局文件名
// data:布局文件数据
func X创建模态窗口并按内存压缩包布局文件(布局文件数据 []byte, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件内存ZIP(布局文件数据, 布局文件名, zip密码, 父对象句柄, 附加窗口句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
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
// hAttachWnd:附加窗口句柄
// hParent:父对象
// pStringXML:字符串
func X创建模态窗口并按布局文件字符串W(字符串 string, 父对象 int, 附加窗口句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件从字符串W(字符串, 父对象, 附加窗口句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
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
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄
// hParent:父对象句柄
// pPrefixName:名称前缀
// pFileName:布局文件名
func X创建模态窗口并按布局文件EX(布局文件名, 名称前缀 string, 父对象句柄 int, 父窗口句柄, 附加窗口句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件Ex(布局文件名, 名称前缀, 父对象句柄, 父窗口句柄, 附加窗口句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
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
// hModule:模块句柄
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄
// hParent:父对象句柄
// pPrefixName:名称前缀
// pPassword:zip密码
// pFileName:布局文件名
// id:RC资源ID
func X创建模态窗口并按RC资源zip压缩包布局文件EX(RC资源ID int32, 布局文件名, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄, 附加窗口句柄, 模块句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件资源ZIPEX(RC资源ID, 布局文件名, zip密码, 名称前缀, 父对象句柄, 父窗口句柄, 附加窗口句柄, 模块句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
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
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄HWND
// hParent:父对象句柄
// pPrefixName:名称前缀
// pPassword:zip密码
// pFileName:布局文件名
// pZipFileName:zip文件名
func X创建模态窗口并按压缩包布局文件EX(zip文件名 string, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件ZIPEx(zip文件名, 布局文件名, zip密码, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
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
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄HWND
// hParent:父对象句柄
// pPrefixName:名称前缀
// pPassword:zip密码
// pFileName:布局文件名
// data:布局文件数据
func X创建模态窗口并按内存压缩包布局文件EX(布局文件数据 []byte, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件内存ZIPEx(布局文件数据, 布局文件名, zip密码, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
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
// hAttachWnd:附加窗口句柄
// hParentWnd:父窗口句柄HWND
// hParent:父对象
// pPrefixName:名称前缀
// pStringXML:字符串
func X创建模态窗口并按布局文件字符串WEX(字符串, 名称前缀 string, 父对象 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) *ModalWindow {
	handle := 炫彩基类.X炫彩_加载布局文件从字符串WEx(字符串, 名称前缀, 父对象, 父窗口句柄HWND, 附加窗口句柄)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 模态窗口_附加窗口, 返回窗口对象.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: xcc.Window_Style_.

// ff:模态窗口附加窗口
// XCStyle:炫彩窗口样式
// hWnd:外部窗口句柄
func X模态窗口附加窗口(外部窗口句柄 uintptr, 炫彩窗口样式 炫彩常量类.Window_Style_) *Window {
	p := &Window{}
	p.X设置句柄(炫彩基类.X模态窗口_附加窗口(外部窗口句柄, 炫彩窗口样式))
	return p
}

// 从句柄创建对象.

// ff:创建模态窗口并按句柄
// handle:
func X创建模态窗口并按句柄(handle int) *ModalWindow {
	p := &ModalWindow{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建模态窗口并按名称
// name:
func X创建模态窗口并按名称(name string) *ModalWindow {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建模态窗口并按UID
// nUID:
func X创建模态窗口并按UID(nUID int) *ModalWindow {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建模态窗口并按UID名称
// name:
func X创建模态窗口并按UID名称(name string) *ModalWindow {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ModalWindow{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 模态窗口_启用自动关闭, 是否自动关闭窗口, 当窗口失去焦点时.
//
// bEnable: 开启开关.

// ff:启用自动关闭
// bEnable:开启开关
func (m *ModalWindow) X启用自动关闭(开启开关 bool) int {
	return 炫彩基类.X模态窗口_启用自动关闭(m.Handle, 开启开关)
}

// 模态窗口_启用ESC关闭, 当用户按ESC键时自动关闭模态窗口.
//
// bEnable: 是否启用.

// ff:启用ESC关闭
// bEnable:是否启用
func (m *ModalWindow) X启用ESC关闭(是否启用 bool) int {
	return 炫彩基类.X模态窗口_启用ESC关闭(m.Handle, 是否启用)
}

// 模态窗口_启动, 启动显示模态窗口, 当窗口关闭时返回: MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出.

// ff:启动
func (m *ModalWindow) X启动() 炫彩常量类.MessageBox_Flag_ {
	return 炫彩基类.X模态窗口_启动(m.Handle)
}

// 模态窗口_结束, 结束模态窗口.
//
// nResult: 用作XModalWnd_DoModal()的返回值. MessageBox_Flag_Ok: 点击确定按钮退出, MessageBox_Flag_Cancel: 点击取消按钮退出, MessageBox_Flag_Other: 其他方式退出.

// ff:结束
// nResult:结果
func (m *ModalWindow) X结束(结果 炫彩常量类.MessageBox_Flag_) int {
	return 炫彩基类.X模态窗口_结束(m.Handle, 结果)
}
