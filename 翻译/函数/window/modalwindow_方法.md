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

[func NewModalWindow(nWidth, nHeight int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *ModalWindow {]
ff=创建模态窗口
XCStyle=炫彩窗口样式
hWndParent=父窗口句柄
pTitle=标题
nHeight=高度
nWidth=宽度

[func NewModalWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *ModalWindow {]
ff=创建模态窗口EX
XCStyle=GUI库窗口样式
hWndParent=父窗口
pTitle=窗口名
cy=高度
cx=宽度
y=y坐标
x=x坐标
lpClassName=类名
dwStyle=样式
dwExStyle=扩展样式

[func NewModalWindowByLayout(pFileName string, hParent int, hAttachWnd uintptr) *ModalWindow {]
ff=创建模态窗口并按布局文件
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pFileName=布局文件名

[func NewModalWindowByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *ModalWindow {]
ff=创建模态窗口并按压缩包布局文件
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pPassword=zip密码
pFileName=布局文件名
pZipFileName=zip文件名

[func NewModalWindowByLayoutZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *ModalWindow {]
ff=创建模态窗口并按内存压缩包布局文件
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pPassword=zip密码
pFileName=布局文件名
data=布局文件数据

[func NewModalWindowByLayoutStringW(pStringXML string, hParent int, hAttachWnd uintptr) *ModalWindow {]
ff=创建模态窗口并按布局文件字符串W
hAttachWnd=附加窗口句柄
hParent=父对象
pStringXML=字符串

[func NewModalWindowByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {]
ff=创建模态窗口并按布局文件EX
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄
hParent=父对象句柄
pPrefixName=名称前缀
pFileName=布局文件名

[func NewModalWindowByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *ModalWindow {]
ff=创建模态窗口并按RC资源zip压缩包布局文件EX
hModule=模块句柄
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄
hParent=父对象句柄
pPrefixName=名称前缀
pPassword=zip密码
pFileName=布局文件名
id=RC资源ID

[func NewModalWindowByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {]
ff=创建模态窗口并按压缩包布局文件EX
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pPassword=zip密码
pFileName=布局文件名
pZipFileName=zip文件名

[func NewModalWindowByLayoutZipMemEx(data #左中括号##右中括号#byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {]
ff=创建模态窗口并按内存压缩包布局文件EX
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pPassword=zip密码
pFileName=布局文件名
data=布局文件数据

[func NewModalWindowByLayoutStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {]
ff=创建模态窗口并按布局文件字符串WEX
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象
pPrefixName=名称前缀
pStringXML=字符串

[func ModalWnd_Attach(hWnd uintptr, XCStyle xcc.Window_Style_) *Window {]
ff=模态窗口附加窗口
XCStyle=炫彩窗口样式
hWnd=外部窗口句柄

[func NewModalWindowByHandle(handle int) *ModalWindow {]
ff=创建模态窗口并按句柄

[func NewModalWindowByName(name string) *ModalWindow {]
ff=创建模态窗口并按名称

[func NewModalWindowByUID(nUID int) *ModalWindow {]
ff=创建模态窗口并按UID

[func NewModalWindowByUIDName(name string) *ModalWindow {]
ff=创建模态窗口并按UID名称

[func (m *ModalWindow) EnableAutoClose(bEnable bool) int {]
ff=启用自动关闭
bEnable=开启开关

[func (m *ModalWindow) EnableEscClose(bEnable bool) int {]
ff=启用ESC关闭
bEnable=是否启用

[func (m *ModalWindow) DoModal() xcc.MessageBox_Flag_ {]
ff=启动

[func (m *ModalWindow) EndModal(nResult xcc.MessageBox_Flag_) int {]
ff=结束
nResult=结果
