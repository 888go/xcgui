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

[func NewScrollBar(x int, y int, cx int, cy int, hParent int) *ScrollBar {]
ff=创建滚动条
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewScrollBarByHandle(handle int) *ScrollBar {]
ff=创建滚动条并按句柄

[func NewScrollBarByName(name string) *ScrollBar {]
ff=创建滚动条并按名称

[func NewScrollBarByUID(nUID int) *ScrollBar {]
ff=创建滚动条并按UID

[func NewScrollBarByUIDName(name string) *ScrollBar {]
ff=创建滚动条并按UID名称

[func (s *ScrollBar) SetRange(range_ int) int {]
ff=置范围
range_=范围

[func (s *ScrollBar) GetRange() int {]
ff=取范围

[func (s *ScrollBar) ShowButton(bShow bool) int {]
ff=显示上下按钮
bShow=是否显示

[func (s *ScrollBar) SetSliderLength(length int) int {]
ff=置滑块长度
length=长度

[func (s *ScrollBar) SetSliderMinLength(minLength int) int {]
ff=置滑块最小长度
minLength=长度

[func (s *ScrollBar) SetSliderPadding(nPadding int) int {]
ff=置滑块两边间隔
nPadding=间隔大小

[func (s *ScrollBar) EnableHorizon(bHorizon bool) bool {]
ff=置水平
bHorizon=水平或垂直

[func (s *ScrollBar) GetSliderMaxLength() int {]
ff=取滑块最大长度

[func (s *ScrollBar) ScrollUp() bool {]
ff=向上滚动

[func (s *ScrollBar) ScrollDown() bool {]
ff=向下滚动

[func (s *ScrollBar) ScrollTop() bool {]
ff=滚动到顶部

[func (s *ScrollBar) ScrollBottom() bool {]
ff=滚动到底部

[func (s *ScrollBar) ScrollPos(pos int) bool {]
ff=滚动到指定位置
pos=位置点

[func (s *ScrollBar) GetButtonUp() int {]
ff=取上按钮

[func (s *ScrollBar) GetButtonDown() int {]
ff=取下按钮

[func (s *ScrollBar) GetButtonSlider() int {]
ff=取滑块

[func (s *ScrollBar) Event_SBAR_SCROLL(pFun XE_SBAR_SCROLL) bool {]
ff=事件_滚动事件
[func (s *ScrollBar) Event_SBAR_SCROLL1(pFun XE_SBAR_SCROLL1) bool {]
ff=事件_滚动事件1
