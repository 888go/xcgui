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

[func NewShapeRect(x int, y int, cx int, cy int, hParent int) *ShapeRect {]
ff=创建形状矩形
hParent=父对象句柄
cy=高度
cx=宽度
y=Y坐标
x=坐标

[func NewShapeRectByHandle(handle int) *ShapeRect {]
ff=创建形状矩形并按句柄

[func NewShapeRectByName(name string) *ShapeRect {]
ff=创建形状矩形并按名称

[func NewShapeRectByUID(nUID int) *ShapeRect {]
ff=创建形状矩形并按UID

[func NewShapeRectByUIDName(name string) *ShapeRect {]
ff=创建形状矩形并按UID名称

[func (s *ShapeRect) SetBorderColor(color int) int {]
ff=置边框色
color=ABGR颜色值

[func (s *ShapeRect) SetFillColor(color int) int {]
ff=置填充色
color=ABGR颜色值

[func (s *ShapeRect) SetRoundAngle(nWidth int, nHeight int32) int {]
ff=置圆角大小
nHeight=圆角高度
nWidth=圆角宽度

[func (s *ShapeRect) GetRoundAngle(pWidth *int, pHeight *int32) int {]
ff=取圆角大小
pHeight=圆角高度
pWidth=圆角宽度

[func (s *ShapeRect) EnableBorder(bEnable bool) int {]
ff=启用边框
bEnable=是否启用

[func (s *ShapeRect) EnableFill(bEnable bool) int {]
ff=启用填充
bEnable=是否启用

[func (s *ShapeRect) EnableRoundAngle(bEnable bool) int {]
ff=启用圆角
bEnable=是否启用
