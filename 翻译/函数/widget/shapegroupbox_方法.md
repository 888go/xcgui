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

[func NewShapeGroupBox(x int, y int, cx int, cy int, pName string, hParent int) *ShapeGroupBox {]
ff=创建形状组框
hParent=父对象句柄
pName=名称
cy=高度
cx=宽度
y=Y坐标
x=坐标

[func NewShapeGroupBoxByHandle(handle int) *ShapeGroupBox {]
ff=创建形状组框并按句柄

[func NewShapeGroupBoxByName(name string) *ShapeGroupBox {]
ff=创建形状组框并按名称

[func NewShapeGroupBoxByUID(nUID int) *ShapeGroupBox {]
ff=创建形状组框并按UID

[func NewShapeGroupBoxByUIDName(name string) *ShapeGroupBox {]
ff=创建形状组框并按UID名称

[func (s *ShapeGroupBox) SetBorderColor(color int) int {]
ff=置边框颜色
color=ABGR颜色值

[func (s *ShapeGroupBox) SetTextColor(color int) int {]
ff=置文本颜色
color=ABGR颜色值

[func (s *ShapeGroupBox) SetFontX(hFontX int) int {]
ff=置字体
hFontX=炫彩字体

[func (s *ShapeGroupBox) SetTextOffset(offsetX int32, offsetY int32) int {]
ff=置文本偏移
offsetY=垂直偏移
offsetX=水平偏移

[func (s *ShapeGroupBox) SetRoundAngle(nWidth int32, nHeight int32) int {]
ff=置圆角大小
nHeight=圆角高度
nWidth=圆角宽度

[func (s *ShapeGroupBox) SetText(pText string) int {]
ff=置文本
pText=文本内容

[func (s *ShapeGroupBox) GetTextOffset(pOffsetX *int32, pOffsetY *int32) int {]
ff=取文本偏移
pOffsetY=Y坐标偏移量
pOffsetX=坐标偏移量

[func (s *ShapeGroupBox) GetRoundAngle(pWidth *int32, pHeight *int32) int {]
ff=取圆角大小
pHeight=返回圆角高度
pWidth=返回圆角宽度

[func (s *ShapeGroupBox) EnableRoundAngle(bEnable bool) int {]
ff=启用圆角
bEnable=是否启用
