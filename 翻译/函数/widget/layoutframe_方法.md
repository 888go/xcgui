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

[func NewLayoutFrame(x int, y int, cx int, cy int, hParent int) *LayoutFrame {]
ff=创建布局框架
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewLayoutFrameByHandle(handle int) *LayoutFrame {]
ff=创建布局框架并按句柄
handle=句柄

[func NewLayoutFrameByName(name string) *LayoutFrame {]
ff=创建布局框架并按名称
name=名称

[func NewLayoutFrameByUID(nUID int) *LayoutFrame {]
ff=创建布局框架并按UID

[func NewLayoutFrameByUIDName(name string) *LayoutFrame {]
ff=创建布局框架并按UID名称
name=名称

[func (l *LayoutFrame) ShowLayoutFrame(bEnable bool) int {]
ff=显示布局边界
bEnable=是否启用

[func (l *LayoutFrame) EnableHorizon(bEnable bool) int {]
ff=启用水平
bEnable=是否启用

[func (l *LayoutFrame) EnableAutoWrap(bEnable bool) int {]
ff=启用自动换行
bEnable=是否启用

[func (l *LayoutFrame) EnableOverflowHide(bEnable bool) int {]
ff=启用溢出隐藏
bEnable=是否启用

[func (l *LayoutFrame) SetAlignH(nAlign xcc.Layout_Align_) int {]
ff=置水平对齐
nAlign=对齐方式

[func (l *LayoutFrame) SetAlignV(nAlign xcc.Layout_Align_) int {]
ff=置垂直对齐
nAlign=对齐方式

[func (l *LayoutFrame) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) int {]
ff=置对齐基线
nAlign=对齐方式

[func (l *LayoutFrame) SetSpace(nSpace int) int {]
ff=置间距
nSpace=项间距大小

[func (l *LayoutFrame) SetSpaceRow(nSpace int) int {]
ff=置行距
nSpace=行间距大小
