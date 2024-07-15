# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func DestroyIcon(hIcon uintptr) bool {]
ff=图标销毁
hIcon=图标句柄

[func CreateIconFromResource(presbits uintptr, dwResSize uint32, fIcon bool, dwVer uint32) (uintptr, error) {]
ff=图标创建并按资源

[func LoadImageW(hInst uintptr, name string, Type IMAGE_, cx, cy int32, fuLoad LR_) uintptr {]
ff=加载图像W
cx=宽度
Type=类型
name=名称
hInst=模块句柄

[func FindWindowW(lpClassName, lpWindowName string) uintptr {]
ff=窗口取顶级句柄
lpWindowName=窗口标题
lpClassName=窗口类名

[func IsWindow(hWnd uintptr) bool {]
ff=窗口句柄是否有效
hWnd=窗口的句柄

[func SetWindowPos(hWnd uintptr, hWndInsertAfter HWND_, x, y, cx, cy int32, wFlags SWP_) bool {]
ff=窗口设置位置
wFlags=大小和定位标志
cy=新高度
cx=新宽度
y=新y坐标
x=新x坐标
hWndInsertAfter=置顶方式
hWnd=窗口句柄

[func GetDesktopWindow() uintptr {]
ff=窗口取桌面句柄

[func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType MB_) ID_ {]
ff=窗口消息框W
uType=类型
lpCaption=标题
lpText=显示消息
hWnd=父窗口句柄

[func OpenClipboard(hWnd uintptr) bool {]
ff=剪辑版打开
hWnd=关联窗口句柄

[func CloseClipboard() bool {]
ff=剪辑版关闭

[func EmptyClipboard() bool {]
ff=剪辑版清空

[func IsClipboardFormatAvailable(uFormat CF_) bool {]
ff=剪贴板内容格式判断
uFormat=格式

[func GetClipboardData(uFormat CF_) uintptr {]
ff=剪贴板取指定格式内容
uFormat=格式

[func SetClipboardData(uFormat CF_, hMem uintptr) uintptr {]
ff=剪贴板设置数据
hMem=指定格式数据句柄
uFormat=格式

[func SetForegroundWindow(hWnd uintptr) bool {]
ff=窗口激活
hWnd=窗口句柄

[func FindWindowExW(hWndParent, hWndChildAfter uintptr, lpszClass, lpszWindow string) uintptr {]
ff=窗口模糊搜索子窗口
lpszWindow=窗口标题
lpszClass=类名
hWndChildAfter=子窗口句柄
hWndParent=窗口句柄

[func GetWindowTextLengthW(hWnd uintptr) int {]
ff=窗口取标题长度
hWnd=窗口或控件句柄

[func GetWindowTextW(hWnd uintptr, lpString *string, nMaxCount int) int {]
ff=窗口取标题
nMaxCount=最大字符数
lpString=接收文本
hWnd=窗口或控件句柄

[func ClientToScreen(hWnd uintptr, lpPoint *POINT) bool {]
ff=窗口取屏幕坐标
lpPoint=坐标指针
hWnd=窗口句柄

[func GetCursorPos(lpPoint *POINT) bool {]
ff=鼠标取光标坐标
lpPoint=坐标指针

[func RegisterHotKey(hWnd uintptr, id int32, fsModifiers, vk uint32) bool {]
ff=键盘热键注册
vk=热键代码
fsModifiers=组合键
id=热键标识符
hWnd=窗口句柄

[func UnregisterHotKey(hWnd uintptr, id int32) bool {]
ff=键盘热键释放
id=热键标识符
hWnd=窗口句柄

[func PostQuitMessage(nExitCode int32) {]
ff=结束
nExitCode=结束代码

[func SendMessageW(hWnd uintptr, Msg int32, wParam, lParam uint) int {]
ff=窗口发送消息
lParam=参数2
wParam=参数1
Msg=消息值
hWnd=窗口句柄

[func PostMessageW(hWnd uintptr, Msg int32, wParam, lParam uint) bool {]
ff=窗口投递消息
lParam=参数2
wParam=参数1
Msg=消息值
hWnd=窗口句柄
