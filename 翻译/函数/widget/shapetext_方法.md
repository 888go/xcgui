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

[func NewShapeText(x, y, cx, cy int32, pName string, hParent int) *ShapeText {]
ff=创建形状文本
hParent=父对象句柄
pName=文本内容
cy=高度
cx=宽度
y=Y坐标
x=坐标

[func NewShapeTextByHandle(handle int) *ShapeText {]
ff=创建形状文本并按句柄

[func NewShapeTextByName(name string) *ShapeText {]
ff=创建形状文本并按名称

[func NewShapeTextByUID(nUID int) *ShapeText {]
ff=创建形状文本并按UID

[func NewShapeTextByUIDName(name string) *ShapeText {]
ff=创建形状文本并按UID名称

[func (s *ShapeText) SetText(pName string) int {]
ff=置文本
pName=文本内容

[func (s *ShapeText) GetText() string {]
ff=取文本

[func (s *ShapeText) GetTextLength() int {]
ff=取文本长度

[func (s *ShapeText) SetFont(hFontx int) int {]
ff=置字体
hFontx=字体句柄

[func (s *ShapeText) GetFont() int {]
ff=取字体

[func (s *ShapeText) SetTextColor(color int) int {]
ff=置文本颜色
color=ABGR颜色值

[func (s *ShapeText) GetTextColor() int {]
ff=取文本颜色

[func (s *ShapeText) SetTextAlign(align xcc.TextFormatFlag_) int {]
ff=置文本对齐
align=文本对齐方式

[func (s *ShapeText) SetOffset(x int, y int) int {]
ff=置偏移
y=Y坐标
x=坐标
