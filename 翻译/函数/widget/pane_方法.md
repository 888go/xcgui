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

[func NewPane(pName string, nWidth int, nHeight int, hFrameWnd int) *Pane {]
ff=创建窗格
hFrameWnd=框架窗口
nHeight=高度
nWidth=宽度
pName=窗格标题

[func NewPaneByHandle(handle int) *Pane {]
ff=创建窗格并按句柄

[func NewPaneByName(name string) *Pane {]
ff=创建窗格并按名称

[func NewPaneByUID(nUID int) *Pane {]
ff=创建窗格并按UID

[func NewPaneByUIDName(name string) *Pane {]
ff=创建窗格并按UID名称

[func (p *Pane) SetView(hView int) int {]
ff=置视图
hView=绑定视图元素

[func (p *Pane) SetTitle(pTitle string) int {]
ff=置标题
pTitle=文本内容

[func (p *Pane) GetTitle() string {]
ff=取标题

[func (p *Pane) SetCaptionHeight(nHeight int) int {]
ff=置标题栏高度
nHeight=高度

[func (p *Pane) GetCaptionHeight() int {]
ff=取标题栏高度

[func (p *Pane) IsShowPane() bool {]
ff=判断显示

[func (p *Pane) SetSize(nWidth int, nHeight int) int {]
ff=置大小
nHeight=高度
nWidth=宽度

[func (p *Pane) GetState() xcc.Pane_State_ {]
ff=取状态

[func (p *Pane) GetViewRect(pRect *xc.RECT) int {]
ff=取视图坐标
pRect=接收返回坐标值

[func (p *Pane) HidePane(bGroupActivate bool) int {]
ff=隐藏
bGroupActivate=延迟组成员激活

[func (p *Pane) ShowPane(bGroupActivate bool) int {]
ff=显示
bGroupActivate=延迟组成员激活

[func (p *Pane) DockPane() int {]
ff=停靠

[func (p *Pane) LockPane() int {]
ff=锁定

[func (p *Pane) FloatPane() int {]
ff=浮动

[func (p *Pane) DrawPane(hDraw int) int {]
ff=绘制
hDraw=图形绘制句柄

[func (p *Pane) SetSelect() bool {]
ff=置选中
[func (p *Pane) IsGroupActivate() bool {]
ff=是否激活
