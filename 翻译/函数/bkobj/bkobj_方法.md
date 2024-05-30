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

[func NewByHandle(handle int) *BkObj {]
ff=创建并按句柄

[func NewByBkm(bkm *bkmanager.BkManager, id int) *BkObj {]
ff=创建并按背景管理器对象
bkm=背景管理器对象

[func NewByBkmHandle(hBkm, id int) *BkObj {]
ff=创建并按背景管理器对象句柄
id=背景对象ID
hBkm=背景管理器句柄

[func (b *BkObj) SetMargin(left int, top int, right int, bottom int) int {]
ff=置外间距
bottom=底部间距
right=右边间距
top=顶部间距
left=左边间距

[func (b *BkObj) SetAlign(nFlags xcc.BkObject_Align_Flag_) int {]
ff=置对齐
nFlags=对齐方式

[func (b *BkObj) SetImage(hImage int) int {]
ff=置图片
hImage=图片句柄

[func (b *BkObj) SetRotate(angle float32) int {]
ff=置旋转
angle=旋转角度

[func (b *BkObj) SetFillColor(color int) int {]
ff=置填充颜色
color=ABGR颜色值

[func (b *BkObj) SetBorderWidth(width int) int {]
ff=置边框宽度
width=宽度

[func (b *BkObj) SetBorderColor(color int) int {]
ff=置边框颜色
color=ABGR颜色值

[func (b *BkObj) SetRectRoundAngle(leftTop int, leftBottom int, rightTop int, rightBottom int) int {]
ff=置矩形圆角
rightBottom=右下角
rightTop=右上角
leftBottom=左下角
leftTop=左上角

[func (b *BkObj) EnableFill(bEnable bool) int {]
ff=启用填充
bEnable=是否启用

[func (b *BkObj) EnableBorder(bEnable bool) int {]
ff=启用边框
bEnable=是否启用

[func (b *BkObj) SetText(pText string) int {]
ff=置文本
pText=文本字符串

[func (b *BkObj) SetFont(hFont int) int {]
ff=置字体
hFont=字体句柄

[func (b *BkObj) SetTextAlign(nAlign xcc.TextFormatFlag_) int {]
ff=置文本对齐
nAlign=对齐方式

[func (b *BkObj) GetMargin(pMargin *xc.RECT) int {]
ff=取外间距
pMargin=接收返回

[func (b *BkObj) GetAlign() xcc.BkObject_Align_Flag_ {]
ff=取对齐

[func (b *BkObj) GetImage() int {]
ff=取图片

[func (b *BkObj) GetRotate() int {]
ff=取旋转角度

[func (b *BkObj) GetFillColor() int {]
ff=取填充色

[func (b *BkObj) GetBorderColor() int {]
ff=取边框色

[func (b *BkObj) GetBorderWidth() int {]
ff=取边框宽度

[func (b *BkObj) GetRectRoundAngle(pRect *xc.RECT) int {]
ff=取矩形圆角
pRect=接收返回圆角大小

[func (b *BkObj) IsFill() bool {]
ff=是否填充

[func (b *BkObj) IsBorder() bool {]
ff=是否边框

[func (b *BkObj) GetText() string {]
ff=取文本

[func (b *BkObj) GetFont() int {]
ff=取字体
[func (b *BkObj) GetTextAlign() xcc.TextFormatFlag_ {]
ff=取文本对齐
