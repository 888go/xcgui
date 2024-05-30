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
# //zj:
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func NewFrameWindow(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *FrameWindow {]
ff=创建框架窗口
XCStyle=GUI库窗口样式
hWndParent=父窗口句柄
pTitle=标题
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *FrameWindow {]
ff=创建框架窗口EX
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

[func NewFrameWindowByHandle(handle int) *FrameWindow {]
ff=创建框架窗口并按句柄

[func NewFrameWindowByName(name string) *FrameWindow {]
ff=创建框架窗口并按名称

[func NewFrameWindowByUID(nUID int) *FrameWindow {]
ff=创建框架窗口并按UID

[func NewFrameWindowByUIDName(name string) *FrameWindow {]
ff=创建框架窗口并按UID名称

[func (fw *FrameWindow) GetLayoutAreaRect(pRect *xc.RECT) int {]
ff=取布局区域坐标
pRect=返回坐标

[func (fw *FrameWindow) SetView(hEle int) int {]
ff=置视图
hEle=元素句柄

[func (fw *FrameWindow) SetPaneSplitBarColor(color int) int {]
ff=置窗格分隔条颜色
color=ABGR颜色值

[func (fw *FrameWindow) SetTabBarHeight(nHeight int32) int {]
ff=置TabBar条高度
nHeight=高度

[func (fw *FrameWindow) SaveLayoutToFile(pFileName string) bool {]
ff=保存布局到文件
pFileName=文件名

[func (fw *FrameWindow) LoadLayoutFile(aPaneList #左中括号##右中括号#int, nPaneCount int32, pFileName string) bool {]
ff=加载布局信息文件
pFileName=文件名
nPaneCount=窗格数量
aPaneList=窗格句柄切片

[func (fw *FrameWindow) AddPane(hPaneDest int, hPaneNew int, align xcc.Pane_Align_) bool {]
ff=添加窗格
align=对齐方式
hPaneNew=当前窗格
hPaneDest=目标窗格

[func (fw *FrameWindow) MergePane(hPaneDest int, hPaneNew int) bool {]
ff=合并窗格
hPaneNew=当前窗格
hPaneDest=目标窗格

[func (fw *FrameWindow) GetDragFloatWndTopFlag() xcc.FrameWnd_Cell_Type_ {]
ff=取拖动浮动窗格停留位置标识

[func FrameWnd_Attach(hWnd uintptr, XCStyle int) *Window {]
ff=附加窗口
XCStyle=炫彩窗口样式
hWnd=外部窗口句柄

[func (fw *FrameWindow) GetViewRect(pRect *xc.RECT) {]
ff=取主视图坐标
pRect=返回坐标

[func (fw *FrameWindow) SetPaneSplitBarWidth(nWidth int32) {]
ff=置窗格分隔条宽度
nWidth=宽度
[func (fw *FrameWindow) GetPaneSplitBarWidth() int32 {]
ff=取窗格分隔条宽度
