
# <翻译开始>
func DestroyIcon(hIcon
图标句柄
# <翻译结束>

# <翻译开始>
func DestroyIcon
X图标销毁
# <翻译结束>


# <翻译开始>
func CreateIconFromResource
X图标创建并按资源
# <翻译结束>


# <翻译开始>
func LoadImageW(hInst uintptr, name string, Type IMAGE_, cx
宽度
# <翻译结束>

# <翻译开始>
func LoadImageW(hInst uintptr, name string, Type
类型
# <翻译结束>

# <翻译开始>
func LoadImageW(hInst uintptr, name
名称
# <翻译结束>

# <翻译开始>
func LoadImageW(hInst
模块句柄
# <翻译结束>

# <翻译开始>
func LoadImageW
X加载图像W
# <翻译结束>


# <翻译开始>
func FindWindowW(lpClassName, lpWindowName
窗口标题
# <翻译结束>

# <翻译开始>
func FindWindowW(lpClassName
窗口类名
# <翻译结束>

# <翻译开始>
func FindWindowW
X窗口取顶级句柄
# <翻译结束>


# <翻译开始>
func IsWindow(hWnd
窗口的句柄
# <翻译结束>

# <翻译开始>
func IsWindow
X窗口句柄是否有效
# <翻译结束>


# <翻译开始>
func SetWindowPos(hWnd uintptr, hWndInsertAfter HWND_, x, y, cx, cy int32, wFlags
大小和定位标志
# <翻译结束>

# <翻译开始>
func SetWindowPos(hWnd uintptr, hWndInsertAfter HWND_, x, y, cx, cy
新高度
# <翻译结束>

# <翻译开始>
func SetWindowPos(hWnd uintptr, hWndInsertAfter HWND_, x, y, cx
新宽度
# <翻译结束>

# <翻译开始>
func SetWindowPos(hWnd uintptr, hWndInsertAfter HWND_, x, y
新y坐标
# <翻译结束>

# <翻译开始>
func SetWindowPos(hWnd uintptr, hWndInsertAfter HWND_, x
新x坐标
# <翻译结束>

# <翻译开始>
func SetWindowPos(hWnd uintptr, hWndInsertAfter
置顶方式
# <翻译结束>

# <翻译开始>
func SetWindowPos(hWnd
窗口句柄
# <翻译结束>

# <翻译开始>
func SetWindowPos
X窗口设置位置
# <翻译结束>


# <翻译开始>
func GetDesktopWindow
X窗口取桌面句柄
# <翻译结束>


# <翻译开始>
func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType
类型
# <翻译结束>

# <翻译开始>
func MessageBoxW(hWnd uintptr, lpText, lpCaption
标题
# <翻译结束>

# <翻译开始>
func MessageBoxW(hWnd uintptr, lpText
显示消息
# <翻译结束>

# <翻译开始>
func MessageBoxW(hWnd
父窗口句柄
# <翻译结束>

# <翻译开始>
func MessageBoxW
X窗口消息框W
# <翻译结束>


# <翻译开始>
func OpenClipboard(hWnd
关联窗口句柄
# <翻译结束>

# <翻译开始>
func OpenClipboard
X剪辑版打开
# <翻译结束>


# <翻译开始>
func CloseClipboard
X剪辑版关闭
# <翻译结束>


# <翻译开始>
func EmptyClipboard
X剪辑版清空
# <翻译结束>


# <翻译开始>
func IsClipboardFormatAvailable(uFormat
格式
# <翻译结束>

# <翻译开始>
func IsClipboardFormatAvailable
X剪贴板内容格式判断
# <翻译结束>


# <翻译开始>
func GetClipboardData(uFormat
格式
# <翻译结束>

# <翻译开始>
func GetClipboardData
X剪贴板取指定格式内容
# <翻译结束>


# <翻译开始>
func SetClipboardData(uFormat CF_, hMem
指定格式数据句柄
# <翻译结束>

# <翻译开始>
func SetClipboardData(uFormat
格式
# <翻译结束>

# <翻译开始>
func SetClipboardData
X剪贴板设置数据
# <翻译结束>


# <翻译开始>
func SetForegroundWindow(hWnd
窗口句柄
# <翻译结束>

# <翻译开始>
func SetForegroundWindow
X窗口激活
# <翻译结束>


# <翻译开始>
func FindWindowExW(hWndParent, hWndChildAfter uintptr, lpszClass, lpszWindow
窗口标题
# <翻译结束>

# <翻译开始>
func FindWindowExW(hWndParent, hWndChildAfter uintptr, lpszClass
类名
# <翻译结束>

# <翻译开始>
func FindWindowExW(hWndParent, hWndChildAfter
子窗口句柄
# <翻译结束>

# <翻译开始>
func FindWindowExW(hWndParent
窗口句柄
# <翻译结束>

# <翻译开始>
func FindWindowExW
X窗口模糊搜索子窗口
# <翻译结束>


# <翻译开始>
func GetWindowTextLengthW(hWnd
窗口或控件句柄
# <翻译结束>

# <翻译开始>
func GetWindowTextLengthW
X窗口取标题长度
# <翻译结束>


# <翻译开始>
func GetWindowTextW(hWnd uintptr, lpString *string, nMaxCount
最大字符数
# <翻译结束>

# <翻译开始>
func GetWindowTextW(hWnd uintptr, lpString
接收文本
# <翻译结束>

# <翻译开始>
func GetWindowTextW(hWnd
窗口或控件句柄
# <翻译结束>

# <翻译开始>
func GetWindowTextW
X窗口取标题
# <翻译结束>


# <翻译开始>
func ClientToScreen(hWnd uintptr, lpPoint
坐标指针
# <翻译结束>

# <翻译开始>
func ClientToScreen(hWnd
窗口句柄
# <翻译结束>


# <翻译开始>
func ClientToScreen
X窗口取屏幕坐标
# <翻译结束>


# <翻译开始>
func GetCursorPos(lpPoint
坐标指针
# <翻译结束>

# <翻译开始>
func GetCursorPos
X鼠标取光标坐标
# <翻译结束>


# <翻译开始>
func RegisterHotKey(hWnd uintptr, id int32, fsModifiers, vk
热键代码
# <翻译结束>

# <翻译开始>
func RegisterHotKey(hWnd uintptr, id int32, fsModifiers
组合键
# <翻译结束>

# <翻译开始>
func RegisterHotKey(hWnd uintptr, id
热键标识符
# <翻译结束>

# <翻译开始>
func RegisterHotKey(hWnd
窗口句柄
# <翻译结束>

# <翻译开始>
func RegisterHotKey
X键盘热键注册
# <翻译结束>


# <翻译开始>
func UnregisterHotKey(hWnd uintptr, id
热键标识符
# <翻译结束>

# <翻译开始>
func UnregisterHotKey(hWnd
窗口句柄
# <翻译结束>

# <翻译开始>
func UnregisterHotKey
X键盘热键释放
# <翻译结束>


# <翻译开始>
func PostQuitMessage(nExitCode
结束代码
# <翻译结束>

# <翻译开始>
func PostQuitMessage
X结束
# <翻译结束>


# <翻译开始>
func SendMessageW(hWnd uintptr, Msg int32, wParam, lParam
参数2
# <翻译结束>

# <翻译开始>
func SendMessageW(hWnd uintptr, Msg int32, wParam
参数1
# <翻译结束>

# <翻译开始>
func SendMessageW(hWnd uintptr, Msg
消息值
# <翻译结束>

# <翻译开始>
func SendMessageW(hWnd
窗口句柄
# <翻译结束>

# <翻译开始>
func SendMessageW
X窗口发送消息
# <翻译结束>


# <翻译开始>
func PostMessageW(hWnd uintptr, Msg int32, wParam, lParam
参数2
# <翻译结束>

# <翻译开始>
func PostMessageW(hWnd uintptr, Msg int32, wParam
参数1
# <翻译结束>

# <翻译开始>
func PostMessageW(hWnd uintptr, Msg
消息值
# <翻译结束>

# <翻译开始>
func PostMessageW(hWnd
窗口句柄
# <翻译结束>

# <翻译开始>
func PostMessageW
X窗口投递消息
# <翻译结束>

