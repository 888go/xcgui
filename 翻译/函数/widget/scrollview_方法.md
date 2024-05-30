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

[func NewScrollView(x int, y int, cx int, cy int, hParent int) *ScrollView {]
ff=创建滚动视
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewScrollViewByHandle(handle int) *ScrollView {]
ff=创建滚动视并按句柄

[func NewScrollViewByName(name string) *ScrollView {]
ff=创建滚动视并按名称

[func NewScrollViewByUID(nUID int) *ScrollView {]
ff=创建滚动视并按UID

[func NewScrollViewByUIDName(name string) *ScrollView {]
ff=创建滚动视并按UID名称

[func (s *ScrollView) SetTotalSize(cx int, cy int) bool {]
ff=置视图大小
cy=高度
cx=宽度

[func (s *ScrollView) GetTotalSize(pSize *xc.SIZE) int {]
ff=取视图大小
pSize=大小

[func (s *ScrollView) SetLineSize(nWidth int, nHeight int) bool {]
ff=置滚动单位大小
nHeight=高度
nWidth=宽度

[func (s *ScrollView) GetLineSize(pSize *xc.SIZE) int {]
ff=取滚动单位大小
pSize=返回大小

[func (s *ScrollView) SetScrollBarSize(size int) int {]
ff=置滚动条大小
size=滚动条大小

[func (s *ScrollView) GetViewPosH() int {]
ff=取视口原点X

[func (s *ScrollView) GetViewPosV() int {]
ff=取视口原点Y

[func (s *ScrollView) GetViewWidth() int {]
ff=取视口宽度

[func (s *ScrollView) GetViewHeight() int {]
ff=取视口高度

[func (s *ScrollView) GetViewRect(pRect *xc.RECT) int {]
ff=取视口坐标
pRect=坐标

[func (s *ScrollView) GetScrollBarH() int {]
ff=取水平滚动条

[func (s *ScrollView) GetScrollBarV() int {]
ff=取垂直滚动条

[func (s *ScrollView) ScrollPosH(pos int) bool {]
ff=水平滚动
pos=位置点

[func (s *ScrollView) ScrollPosV(pos int) bool {]
ff=垂直滚动
pos=位置点

[func (s *ScrollView) ScrollPosXH(posX int) bool {]
ff=水平滚动到X
posX=坐标

[func (s *ScrollView) ScrollPosYV(posY int) bool {]
ff=垂直滚动到Y
posY=Y坐标

[func (s *ScrollView) ShowSBarH(bShow bool) int {]
ff=显示水平滚动条
bShow=是否显示

[func (s *ScrollView) ShowSBarV(bShow bool) int {]
ff=显示垂直滚动条
bShow=是否显示

[func (s *ScrollView) EnableAutoShowScrollBar(bEnable bool) int {]
ff=启用自动显示滚动条
bEnable=是否启用

[func (s *ScrollView) ScrollLeftLine() bool {]
ff=向左滚动

[func (s *ScrollView) ScrollRightLine() bool {]
ff=向右滚动

[func (s *ScrollView) ScrollTopLine() bool {]
ff=向上滚动

[func (s *ScrollView) ScrollBottomLine() bool {]
ff=向下滚动

[func (s *ScrollView) ScrollLeft() bool {]
ff=滚动到左侧

[func (s *ScrollView) ScrollRight() bool {]
ff=滚动到右侧

[func (s *ScrollView) ScrollTop() bool {]
ff=滚动到顶部

[func (s *ScrollView) ScrollBottom() bool {]
ff=滚动到底部

[func (s *ScrollView) Event_SCROLLVIEW_SCROLL_H(pFun XE_SCROLLVIEW_SCROLL_H) bool {]
ff=事件_水平滚动

[func (s *ScrollView) Event_SCROLLVIEW_SCROLL_H1(pFun XE_SCROLLVIEW_SCROLL_H1) bool {]
ff=事件_水平滚动1

[func (s *ScrollView) Event_SCROLLVIEW_SCROLL_V(pFun XE_SCROLLVIEW_SCROLL_V) bool {]
ff=事件_垂直滚动
[func (s *ScrollView) Event_SCROLLVIEW_SCROLL_V1(pFun XE_SCROLLVIEW_SCROLL_V1) bool {]
ff=事件_垂直滚动1
