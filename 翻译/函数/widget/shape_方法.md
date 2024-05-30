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

[func NewShapeByHandle(handle int) *Shape {]
ff=创建形状对象基类并按句柄

[func NewShapeByName(name string) *Shape {]
ff=创建形状对象基类并按名称

[func NewShapeByUID(nUID int) *Shape {]
ff=创建形状对象基类并按UID

[func NewShapeByUIDName(name string) *Shape {]
ff=创建形状对象基类并按UID名称

[func (s *Shape) RemoveShape() int {]
ff=移除

[func (s *Shape) GetZOrder() int {]
ff=取Z序

[func (s *Shape) Redraw() int {]
ff=重绘

[func (s *Shape) GetWidth() int32 {]
ff=取宽度

[func (s *Shape) GetHeight() int32 {]
ff=取高度

[func (s *Shape) SetPosition(x, y int32) int {]
ff=移动位置
y=y坐标
x=x坐标

[func (s *Shape) GetRect(pRect *xc.RECT) int {]
ff=取坐标
pRect=接收返回坐标

[func (s *Shape) SetRect(pRect *xc.RECT) int {]
ff=置坐标
pRect=坐标

[func (s *Shape) SetRectLogic(pRect *xc.RECT, bRedraw bool) bool {]
ff=置逻辑坐标
pRect=坐标

[func (s *Shape) GetRectLogic(pRect *xc.RECT) int {]
ff=取逻辑坐标
pRect=坐标

[func (s *Shape) GetWndClientRect(pRect *xc.RECT) int {]
ff=取基于窗口客户区坐标
pRect=坐标

[func (s *Shape) GetContentSize(pSize *xc.SIZE) int {]
ff=取内容大小
pSize=接收返回大小值

[func (s *Shape) ShowLayout(bShow bool) int {]
ff=显示布局边界
bShow=是否显示

[func (s *Shape) AdjustLayout() int {]
ff=调整布局

[func (s *Shape) Destroy() int {]
ff=销毁

[func (s *Shape) GetPosition(pOutX, pOutY *int32) int {]
ff=取位置
pOutY=返回Y坐标
pOutX=返回X坐标

[func (s *Shape) SetSize(nWidth, nHeight int32) int {]
ff=置大小
nHeight=高度
nWidth=宽度

[func (s *Shape) GetSize(pOutWidth, pOutHeight *int32) int {]
ff=取大小
pOutHeight=返回高度
pOutWidth=返回宽度

[func (s *Shape) SetAlpha(alpha uint8) int {]
ff=置透明度
alpha=透明度
[func (s *Shape) GetAlpha() int {]
ff=取透明度
