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

[func NewSliderBar(x int, y int, cx int, cy int, hParent int) *SliderBar {]
ff=创建滑动条
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewSliderBarByHandle(handle int) *SliderBar {]
ff=创建滑动条并按句柄

[func NewSliderBarByName(name string) *SliderBar {]
ff=创建滑动条并按名称

[func NewSliderBarByUID(nUID int) *SliderBar {]
ff=创建滑动条并按UID

[func NewSliderBarByUIDName(name string) *SliderBar {]
ff=创建滑动条并按UID名称

[func (s *SliderBar) SetRange(range_ int) int {]
ff=置范围
range_=范围

[func (s *SliderBar) GetRange() int {]
ff=取范围

[func (s *SliderBar) SetImageLoad(hImage int) int {]
ff=置进度图片
hImage=图片句柄

[func (s *SliderBar) SetButtonWidth(width int) int {]
ff=置滑块宽度
width=宽度

[func (s *SliderBar) SetButtonHeight(height int) int {]
ff=置滑块高度
height=高度

[func (s *SliderBar) SetPos(pos int) int {]
ff=置当前位置
pos=进度点

[func (s *SliderBar) GetPos() int {]
ff=取当前位置

[func (s *SliderBar) GetButton() int {]
ff=取滑块

[func (s *SliderBar) EnableHorizon(bHorizon bool) int {]
ff=置水平
bHorizon=水平或垂直

[func (s *SliderBar) Event_SLIDERBAR_CHANGE(pFun XE_SLIDERBAR_CHANGE) bool {]
ff=事件_滑块位置改变

[func (s *SliderBar) Event_SLIDERBAR_CHANGE1(pFun XE_SLIDERBAR_CHANGE1) bool {]
ff=事件_滑块位置改变1
