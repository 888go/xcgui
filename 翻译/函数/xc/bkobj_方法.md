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

[func XBkObj_SetMargin(hObj int, left int, top int, right int, bottom int) int {]
ff=背景对象_置外间距
bottom=底部间距
right=右边间距
top=顶部间距
left=左边间距
hObj=背景对象句柄

[func XBkObj_SetAlign(hObj int, nFlags xcc.BkObject_Align_Flag_) int {]
ff=背景对象_置对齐
nFlags=对齐方式
hObj=背景对象句柄

[func XBkObj_SetImage(hObj int, hImage int) int {]
ff=背景对象_置图片
hImage=图片句柄
hObj=背景对象句柄

[func XBkObj_SetRotate(hObj int, angle float32) int {]
ff=背景对象_置旋转
angle=旋转角度
hObj=背景对象句柄

[func XBkObj_SetFillColor(hObj int, color int) int {]
ff=背景对象_置填充颜色
color=ABGR颜色值
hObj=背景对象句柄

[func XBkObj_SetBorderWidth(hObj int, width int) int {]
ff=背景对象_置边框宽度
width=宽度
hObj=背景对象句柄

[func XBkObj_SetBorderColor(hObj int, color int) int {]
ff=背景对象_置边框颜色
color=ABGR颜色值
hObj=背景对象句柄

[func XBkObj_SetRectRoundAngle(hObj int, leftTop int, leftBottom int, rightTop int, rightBottom int) int {]
ff=背景对象_置矩形圆角
rightBottom=右下角
rightTop=右上角
leftBottom=左下角
leftTop=左上角
hObj=背景对象句柄

[func XBkObj_EnableFill(hObj int, bEnable bool) int {]
ff=背景对象_启用填充
bEnable=是否启用
hObj=背景对象句柄

[func XBkObj_EnableBorder(hObj int, bEnable bool) int {]
ff=背景对象_启用边框
bEnable=是否启用
hObj=背景对象句柄

[func XBkObj_SetText(hObj int, pText string) int {]
ff=背景对象_置文本
pText=文本字符串
hObj=背景对象句柄

[func XBkObj_SetFont(hObj int, hFont int) int {]
ff=背景对象_置字体
hFont=字体句柄
hObj=背景对象句柄

[func XBkObj_SetTextAlign(hObj int, nAlign xcc.TextFormatFlag_) int {]
ff=背景对象_置文本对齐
nAlign=文本对齐方式
hObj=背景对象句柄

[func XBkObj_GetMargin(hObj int, pMargin *RECT) int {]
ff=背景对象_取外间距
pMargin=接收返回外间距
hObj=背景对象句柄

[func XBkObj_GetAlign(hObj int) xcc.BkObject_Align_Flag_ {]
ff=背景对象_取对齐
hObj=背景对象句柄

[func XBkObj_GetImage(hObj int) int {]
ff=背景对象_取图片
hObj=背景对象句柄

[func XBkObj_GetRotate(hObj int) int {]
ff=背景对象_取旋转角度
hObj=背景对象句柄

[func XBkObj_GetFillColor(hObj int) int {]
ff=背景对象_取填充色
hObj=背景对象句柄

[func XBkObj_GetBorderColor(hObj int) int {]
ff=背景对象_取边框色
hObj=背景对象句柄

[func XBkObj_GetBorderWidth(hObj int) int {]
ff=背景对象_取边框宽度
hObj=背景对象句柄

[func XBkObj_GetRectRoundAngle(hObj int, pRect *RECT) int {]
ff=背景对象_取矩形圆角
pRect=接收返回圆角大小
hObj=背景对象句柄

[func XBkObj_IsFill(hObj int) bool {]
ff=背景对象_是否填充
hObj=背景对象句柄

[func XBkObj_IsBorder(hObj int) bool {]
ff=背景对象_是否边框
hObj=背景对象句柄

[func XBkObj_GetText(hObj int) string {]
ff=背景对象_取文本
hObj=背景对象句柄

[func XBkObj_GetFont(hObj int) int {]
ff=背景对象_取字体
hObj=背景对象句柄

[func XBkObj_GetTextAlign(hObj int) xcc.TextFormatFlag_ {]
ff=背景对象_取文本对齐
hObj=背景对象句柄
