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

[func XShape_RemoveShape(hShape int) int {]
ff=形状_移除
hShape=形状对象句柄

[func XShape_GetZOrder(hShape int) int {]
ff=形状_取Z序
hShape=形状对象句柄

[func XShape_Redraw(hShape int) int {]
ff=形状_重绘
hShape=形状对象句柄

[func XShape_GetWidth(hShape int) int32 {]
ff=形状_取宽度
hShape=形状对象句柄

[func XShape_GetHeight(hShape int) int32 {]
ff=形状_取高度
hShape=形状对象句柄

[func XShape_SetPosition(hShape int, x, y int32) int {]
ff=形状_移动位置
y=y坐标
x=x坐标
hShape=形状对象句柄

[func XShape_GetRect(hShape int, pRect *RECT) int {]
ff=形状_取坐标
pRect=接收返回坐标
hShape=形状对象句柄

[func XShape_SetRect(hShape int, pRect *RECT) int {]
ff=形状_置坐标
pRect=坐标
hShape=形状对象句柄

[func XShape_SetRectLogic(hShape int, pRect *RECT, bRedraw bool) bool {]
ff=形状_置逻辑坐标
bRedraw=是否重绘
pRect=坐标
hShape=形状对象句柄

[func XShape_GetRectLogic(hShape int, pRect *RECT) int {]
ff=形状_取逻辑坐标
pRect=坐标
hShape=形状对象句柄

[func XShape_GetWndClientRect(hShape int, pRect *RECT) int {]
ff=形状_取基于窗口客户区坐标
pRect=坐标
hShape=形状对象句柄

[func XShape_GetContentSize(hShape int, pSize *SIZE) int {]
ff=形状_取内容大小
pSize=接收返回内容大小值
hShape=形状对象句柄

[func XShape_ShowLayout(hShape int, bShow bool) int {]
ff=形状_显示布局边界
bShow=是否显示
hShape=形状对象句柄

[func XShape_AdjustLayout(hShape int) int {]
ff=形状_调整布局
hShape=形状对象句柄

[func XShape_Destroy(hShape int) int {]
ff=形状_销毁
hShape=形状对象句柄

[func XShape_GetPosition(hShape int, pOutX, pOutY *int32) int {]
ff=形状_取位置
pOutY=返回Y坐标
pOutX=返回X坐标
hShape=形状对象句柄

[func XShape_SetSize(hShape int, nWidth, nHeight int32) int {]
ff=形状_置大小
nHeight=高度
nWidth=宽度
hShape=形状对象句柄

[func XShape_GetSize(hShape int, pOutWidth, pOutHeight *int32) int {]
ff=形状_取大小
pOutHeight=返回高度
pOutWidth=返回宽度
hShape=形状对象句柄

[func XShape_SetAlpha(hShape int, alpha uint8) int {]
ff=形状_置透明度
alpha=透明度
hShape=形状对象句柄

[func XShape_GetAlpha(hShape int) int {]
ff=形状_取透明度
hShape=形状对象句柄
