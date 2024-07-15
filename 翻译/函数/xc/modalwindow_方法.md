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

[func XModalWnd_Create(nWidth, nHeight int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {]
ff=模态窗口_创建
XCStyle=炫彩窗口样式
hWndParent=父窗口句柄
pTitle=标题内容
nHeight=高度
nWidth=宽度

[func XModalWnd_CreateEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {]
ff=模态窗口_创建EX
XCStyle=GUI库窗口样式
hWndParent=父窗口
pTitle=窗口名
cy=高度
cx=宽度
y=左上角y坐标
x=左上角x坐标
lpClassName=类名
dwStyle=样式
dwExStyle=扩展样式

[func XModalWnd_EnableAutoClose(hWindow int, bEnable bool) int {]
ff=模态窗口_启用自动关闭
bEnable=开启开关
hWindow=模态窗口句柄

[func XModalWnd_EnableEscClose(hWindow int, bEnable bool) int {]
ff=模态窗口_启用ESC关闭
bEnable=是否启用
hWindow=模态窗口句柄

[func XModalWnd_DoModal(hWindow int) xcc.MessageBox_Flag_ {]
ff=模态窗口_启动
hWindow=模态窗口句柄

[func XModalWnd_EndModal(hWindow int, nResult xcc.MessageBox_Flag_) int {]
ff=模态窗口_结束
nResult=用作XModalWnd_DoModal
hWindow=窗口句柄

[func XModalWnd_Attach(hWnd uintptr, XCStyle xcc.Window_Style_) int {]
ff=模态窗口_附加窗口
XCStyle=炫彩窗口样式
hWnd=外部窗口句柄
