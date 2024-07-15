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

[func XShapeText_Create(x, y, cx, cy int32, pName string, hParent int) int {]
ff=形状文本_创建
hParent=父对象句柄
pName=文本内容
cy=高度
cx=宽度
y=Y坐标
x=坐标

[func XShapeText_SetText(hTextBlock int, pName string) int {]
ff=形状文本_置文本
pName=文本内容
hTextBlock=形状对象文本句柄

[func XShapeText_GetText(hTextBlock int) string {]
ff=形状文本_取文本
hTextBlock=形状对象文本句柄

[func XShapeText_GetTextLength(hTextBlock int) int {]
ff=形状文本_取文本长度
hTextBlock=形状对象文本句柄

[func XShapeText_SetFont(hTextBlock int, hFontx int) int {]
ff=形状文本_置字体
hFontx=字体句柄
hTextBlock=形状对象文本句柄

[func XShapeText_GetFont(hTextBlock int) int {]
ff=形状文本_取字体
hTextBlock=形状对象文本句柄

[func XShapeText_SetTextColor(hTextBlock int, color int) int {]
ff=形状文本_置文本颜色
color=ABGR颜色值
hTextBlock=形状对象文本句柄

[func XShapeText_GetTextColor(hTextBlock int) int {]
ff=形状文本_取文本颜色
hTextBlock=形状对象文本句柄

[func XShapeText_SetTextAlign(hTextBlock int, align xcc.TextFormatFlag_) int {]
ff=形状文本_置文本对齐
align=文本对齐方式
hTextBlock=形状对象文本句柄

[func XShapeText_SetOffset(hTextBlock int, x int, y int) int {]
ff=形状文本_置偏移
y=Y坐标
x=坐标
hTextBlock=形状对象文本句柄
