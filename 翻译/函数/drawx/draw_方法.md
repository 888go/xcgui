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

[func New(hWindow int) *Draw {]
ff=创建
hWindow=窗口句柄

[func NewGDI(hWindow int, hdc uintptr) *Draw {]
ff=创建GDI
hdc=hdc句柄
hWindow=窗口句柄

[func NewByHandle(handle int) *Draw {]
ff=创建并按图形绘制模块句柄

[func (d *Draw) Destroy() int {]
ff=销毁

[func (d *Draw) Dottedline(x1 int, y1 int, x2 int, y2 int) int {]
ff=虚线
y2=结束点y坐标
x2=结束点x坐标
y1=起点y坐标
x1=起点x坐标

[func (d *Draw) DottedlineF(x1, y1, x2, y2 float32) int {]
ff=虚线F
y2=结束点y坐标
x2=结束点x坐标
y1=起点y坐标
x1=起点x坐标

[func (d *Draw) DrawArc(x, y int, nWidth int, nHeight int, startAngle float32, sweepAngle float32) int {]
ff=圆弧
sweepAngle=绘制角度
startAngle=起始角度
nHeight=高度
nWidth=宽度
y=y坐标
x=x坐标

[func (d *Draw) DrawArcF(x, y, nWidth, nHeight, startAngle, sweepAngle float32) int {]
ff=圆弧F
sweepAngle=绘制角度
startAngle=起始角度
nHeight=高度
nWidth=宽度
y=y坐标
x=x坐标

[func (d *Draw) DrawCurve(points #左中括号##右中括号#xc.POINT, count int, tension float32) int {]
ff=曲线
points=坐标点切片

[func (d *Draw) DrawCurveF(points #左中括号##右中括号#xc.POINTF, count int, tension float32) int {]
ff=曲线F
points=坐标点切片

[func (d *Draw) DrawLine(x1 int, y1 int, x2 int, y2 int) int {]
ff=线条
y2=y2坐标
x2=x2坐标
y1=y1坐标
x1=x1坐标

[func (d *Draw) DrawLineF(x1, y1, x2, y2 float32) int {]
ff=线条F
y2=y2坐标
x2=x2坐标
y1=y1坐标
x1=x1坐标

[func (d *Draw) DrawPolygon(points #左中括号##右中括号#xc.POINT, nCount int) int {]
ff=多边形
points=顶点坐标切片

[func (d *Draw) DrawPolygonF(points #左中括号##右中括号#xc.POINTF, nCount int) int {]
ff=多边形F
points=顶点坐标切片

[func (d *Draw) DrawRect(pRect *xc.RECT) int {]
ff=矩形边框
pRect=矩形坐标

[func (d *Draw) DrawRectF(pRect *xc.RECTF) int {]
ff=矩形边框F
pRect=矩形坐标

[func (d *Draw) SetOffset(x, y int32) int {]
ff=置偏移
y=Y轴偏移量
x=轴偏移量

[func (d *Draw) GetOffset(pX, pY *int32) int {]
ff=取偏移
pY=接收Y轴偏移量
pX=接收X轴偏移量

[func (d *Draw) GDI_RestoreGDIOBJ() int {]
ff=还原状态

[func (d *Draw) GetHDC() uintptr {]
ff=取HDC

[func (d *Draw) SetBrushColor(color int) int {]
ff=置画刷颜色
color=ABGR

[func (d *Draw) SetTextVertical(bVertical bool) int {]
ff=置文本垂直
bVertical=是否垂直显示文本

[func (d *Draw) SetTextAlign(nFlags xcc.TextFormatFlag_) int {]
ff=置文本对齐
nFlags=对齐标识

[func (d *Draw) SetFont(hFontx int) int {]
ff=置字体
hFontx=炫彩字体

[func (d *Draw) SetLineWidth(nWidth int) int {]
ff=置线宽
nWidth=宽度

[func (d *Draw) SetLineWidthF(nWidth float32) int {]
ff=置线宽F
nWidth=宽度

[func (d *Draw) GDI_SetBkMode(bTransparent bool) int {]
ff=置背景模式
bTransparent=参见MSDN

[func (d *Draw) SetClipRect(pRect *xc.RECT) int {]
ff=置裁剪区域
pRect=区域坐标

[func (d *Draw) SetD2dTextRenderingMode(mode xcc.XC_DWRITE_RENDERING_MODE_) int {]
ff=置D2D文本渲染模式
mode=渲染模式

[func (d *Draw) ClearClip() int {]
ff=清除裁剪区域

[func (d *Draw) EnableSmoothingMode(bEnable bool) int {]
ff=启用平滑模式
bEnable=是否启用

[func (d *Draw) EnableWndTransparent(bTransparent bool) int {]
ff=启用窗口透明判断
bTransparent=是否启用

[func (d *Draw) GDI_CreateSolidBrush(crColor int) int {]
ff=创建实心画刷
crColor=画刷颜色

[func (d *Draw) GDI_CreatePen(fnPenStyle int, nWidth int, crColor int) int {]
ff=创建画笔
crColor=ABGR颜色
nWidth=画笔宽度
fnPenStyle=画笔样式

[func (d *Draw) GDI_CreateRectRgn(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) int {]
ff=创建矩形区域
nBottomRect=右下角Y坐标
nRightRect=右下角X坐标
nTopRect=左上角Y坐标
nLeftRect=左上角X坐标

[func (d *Draw) GDI_CreateRoundRectRgn(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int, nWidthEllipse int, nHeightEllipse int) int {]
ff=创建圆角矩形区域
nHeightEllipse=椭圆的高度
nWidthEllipse=椭圆的宽度
nBottomRect=Y坐标右下角
nRightRect=坐标右下角
nTopRect=Y坐标左上角
nLeftRect=坐标左上角

[func (d *Draw) GDI_CreatePolygonRgn(pPt #左中括号##右中括号#xc.POINT, cPoints, fnPolyFillMode int) int {]
ff=创建多边形区域
pPt=POINT切片

[func (d *Draw) GDI_Ellipse(pRect *xc.RECT) bool {]
ff=GDI椭圆
pRect=矩形区域

[func (d *Draw) GDI_SelectClipRgn(hRgn int) int {]
ff=选择裁剪区域
hRgn=区域句柄

[func (d *Draw) FillRect(pRect *xc.RECT) int {]
ff=填充矩形
pRect=矩形区域

[func (d *Draw) FillRectF(pRect *xc.RECTF) int {]
ff=填充矩形F
pRect=矩形区域

[func (d *Draw) FillRectColor(pRect *xc.RECT, color int) int {]
ff=填充矩形指定颜色
pRect=矩形区域

[func (d *Draw) FillRectColorF(pRect *xc.RECTF, color int) int {]
ff=填充矩形指定颜色F
pRect=矩形区域

[func (d *Draw) GDI_FillRgn(hrgn int, hbr int) bool {]
ff=填充区域
hbr=画刷句柄
hrgn=区域句柄

[func (d *Draw) FillEllipse(pRect *xc.RECT) int {]
ff=填充圆形
pRect=矩形区域

[func (d *Draw) FillEllipseF(pRect *xc.RECTF) int {]
ff=填充圆形F
pRect=矩形区域

[func (d *Draw) DrawEllipse(pRect *xc.RECT) int {]
ff=圆形
pRect=矩形区域

[func (d *Draw) FillRoundRect(pRect *xc.RECT, nWidth, nHeight int) int {]
ff=填充圆角矩形
pRect=矩形坐标

[func (d *Draw) FillRoundRectF(pRect *xc.RECTF, nWidth, nHeight float32) int {]
ff=填充圆角矩形F
pRect=矩形坐标

[func (d *Draw) DrawRoundRect(pRect *xc.RECT, nWidth int, nHeight int) int {]
ff=圆角矩形
pRect=矩形坐标

[func (d *Draw) DrawRoundRectF(pRect *xc.RECT, nWidth, nHeight float32) int {]
ff=圆角矩形F
pRect=矩形坐标

[func (d *Draw) FillRoundRectEx(pRect *xc.RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom int) int {]
ff=填充圆角矩形EX
pRect=坐标

[func (d *Draw) FillRoundRectExF(pRect *xc.RECTF, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {]
ff=填充圆角矩形EXF
pRect=坐标

[func (d *Draw) DrawRoundRectEx(pRect *xc.RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {]
ff=圆角矩形EX
pRect=坐标

[func (d *Draw) DrawRoundRectExF(pRect *xc.RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {]
ff=圆角矩形EXF
pRect=坐标

[func (d *Draw) GDI_Rectangle(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) bool {]
ff=矩形
nBottomRect=右下角Y坐标
nRightRect=右下角X坐标
nTopRect=左上角Y坐标
nLeftRect=左上角X坐标

[func (d *Draw) GradientFill2(pRect *xc.RECT, color1 int, color2 int, mode xcc.GRADIENT_FILL_) int {]
ff=渐变填充2
pRect=矩形坐标

[func (d *Draw) GradientFill2F(pRect *xc.RECTF, color1 int, color2 int, mode xcc.GRADIENT_FILL_) int {]
ff=渐变填充2F
pRect=矩形坐标

[func (d *Draw) GradientFill4(pRect *xc.RECT, color1 int, color2 int, color3 int, color4 int, mode xcc.GRADIENT_FILL_) bool {]
ff=渐变填充4
pRect=矩形坐标

[func (d *Draw) GradientFill4F(pRect *xc.RECTF, color1 int, color2 int, color3 int, color4 int, mode xcc.GRADIENT_FILL_) bool {]
ff=渐变填充4F
pRect=矩形坐标

[func (d *Draw) GDI_FrameRgn(hrgn int, hbr int, nWidth int, nHeight int) bool {]
ff=边框区域
nHeight=边框高度
nWidth=边框宽度
hbr=画刷句柄
hrgn=区域句柄

[func (d *Draw) FocusRect(pRect *xc.RECT) int {]
ff=焦点矩形
pRect=矩形坐标

[func (d *Draw) FocusRectF(pRect *xc.RECTF) int {]
ff=焦点矩形F
pRect=矩形坐标

[func (d *Draw) GDI_MoveToEx(X int, Y int, pPoint *xc.POINT) bool {]
ff=移动到起点
pPoint=接收原位置指针
Y=坐标y
X=坐标x

[func (d *Draw) GDI_LineTo(nXEnd int, nYEnd int) bool {]
ff=线终点
nYEnd=Y坐标
nXEnd=坐标

[func (d *Draw) GDI_Polyline(pArrayPt #左中括号##右中括号#xc.POINT, arrayPtSize int) bool {]
ff=折线

[func (d *Draw) GDI_SetPixel(X int, Y int, crColor int) int {]
ff=置像素颜色
crColor=RGB颜色值
Y=坐标y
X=坐标x

[func (d *Draw) XDraw_GetD2dRenderTarget() int {]
ff=取D2D渲染目标

[func (d *Draw) GDI_DrawIconEx(xLeft int, yTop int, hIcon uintptr, cxWidth int, cyWidth int, istepIfAniCur int, hbrFlickerFreeDraw int, diFlags int) bool {]
ff=图标

[func (d *Draw) GDI_BitBlt(nXDest, nYDest, nWidth, nHeight int32, hdcSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {]
ff=复制

[func (d *Draw) GDI_BitBlt2(nXDest, nYDest, nWidth, nHeight int32, hDrawSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {]
ff=复制2

[func (d *Draw) GDI_AlphaBlend(nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int32, hdcSrc uintptr, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha int32) bool {]
ff=带透明复制

[func (d *Draw) FillPolygon(points #左中括号##右中括号#xc.POINT, nCount int) int {]
ff=填充多边形
points=顶点坐标切片

[func (d *Draw) FillPolygonF(points #左中括号##右中括号#xc.POINTF, nCount int) int {]
ff=填充多边形F
points=顶点坐标切片

[func (d *Draw) Image(hImageFrame int, x, y int32) {]
ff=图片
y=y坐标
x=x坐标
hImageFrame=图片句柄

[func (d *Draw) ImageF(hImageFrame int, x, y float32) int {]
ff=图片F
y=y坐标
x=x坐标
hImageFrame=图片句柄

[func (d *Draw) ImageAdaptive(hImageFrame int, pRect *xc.RECT, bOnlyBorder bool) int {]
ff=图片自适应
pRect=坐标
hImageFrame=图片句柄

[func (d *Draw) ImageAdaptiveF(hImageFrame int, pRect *xc.RECTF, bOnlyBorder bool) int {]
ff=图片自适应F
pRect=坐标
hImageFrame=图片句柄

[func (d *Draw) XDraw_ImageEx(hImageFrame int, x, y, width, height int) int {]
ff=图片EX
height=高度
width=宽度
y=y坐标
x=x坐标
hImageFrame=图片句柄

[func (d *Draw) XDraw_ImageExF(hImageFrame int, x, y, width, height float32) int {]
ff=图片EXF
height=高度
width=宽度
y=y坐标
x=x坐标
hImageFrame=图片句柄

[func (d *Draw) ImageSuper(hImageFrame int, pRect *xc.RECT, bClip bool) int {]
ff=图片增强
pRect=坐标
hImageFrame=图片句柄

[func (d *Draw) ImageSuperF(hImageFrame int, pRect *xc.RECTF, bClip bool) int {]
ff=图片增强F
pRect=坐标
hImageFrame=图片句柄

[func (d *Draw) ImageSuperEx(hImageFrame int, prcDest *xc.RECT, prcSrc *xc.RECT) int {]
ff=图片增强扩展
prcDest=目标坐标
hImageFrame=图片句柄

[func (d *Draw) ImageSuperExF(hImageFrame int, prcDest *xc.RECTF, prcSrc *xc.RECT) int {]
ff=图片增强EXF
prcDest=目标坐标
hImageFrame=图片句柄

[func (d *Draw) ImageSuperMask(hImageFrame int, hImageFrameMask int, pRect *xc.RECT, pRectMask *xc.RECT, bClip bool) int {]
ff=图片增强遮盖
pRect=坐标
hImageFrameMask=图片遮盖句柄
hImageFrame=图片句柄

[func (d *Draw) ImageTile(hImageFrame int, hImageFrameMask int, pRect *xc.RECT, flag int) int {]
ff=图片平铺
pRect=坐标
hImageFrame=图片句柄

[func (d *Draw) ImageTileF(hImageFrame int, hImageFrameMask int, pRect *xc.RECTF, flag int) int {]
ff=图片平铺F
pRect=坐标
hImageFrame=图片句柄

[func (d *Draw) ImageMask(hImageFrame int, hImageFrameMask int, x int, y int, x2 int, y2 int) int {]
ff=图片遮盖
y2=y2坐标
x2=x2坐标
y=y1坐标
x=x1坐标
hImageFrameMask=图片遮盖句柄
hImageFrame=图片句柄

[func (d *Draw) DrawText(lpString string, lpRect *xc.RECT) int {]
ff=文本指定矩形
lpRect=坐标
lpString=字符串

[func (d *Draw) DrawTextF(lpString string, lpRect *xc.RECTF) int {]
ff=文本指定矩形F
lpRect=坐标
lpString=字符串

[func (d *Draw) DrawTextUnderline(lpString string, lpRect *xc.RECT, colorLine int) int {]
ff=文本下划线
lpRect=坐标
lpString=字符串

[func (d *Draw) DrawTextUnderlineF(lpString string, lpRect *xc.RECTF, colorLine int) int {]
ff=文本下划线F
lpRect=坐标
lpString=字符串

[func (d *Draw) TextOut(nXStart int, nYStart int, lpString string, cbString string) int {]
ff=文本输出

[func (d *Draw) TextOutF(nXStart, nYStart float32, lpString string, cbString string) int {]
ff=文本输出F

[func (d *Draw) TextOutEx(nXStart int, nYStart int, lpString string) int {]
ff=文本输出EX

[func (d *Draw) TextOutExF(nXStart, nYStart float32, lpString string) int {]
ff=文本输出EXF

[func (d *Draw) TextOutA(nXStart int, nYStart int, lpString string) int {]
ff=文本输出A

[func (d *Draw) TextOutAF(nXStart, nYStart float32, lpString string) int {]
ff=文本输出AF

[func (d *Draw) SetTextRenderingHint(nType int) int {]
ff=设置文本渲染提示

[func (d *Draw) DrawSvgSrc(hSvg int) int {]
ff=SVG源
hSvg=SVG句柄

[func (d *Draw) DrawSvg(hSvg int, x int, y int) int {]
ff=SVG
y=y坐标
x=x坐标
hSvg=SVG句柄

[func (d *Draw) DrawSvgEx(hSvg int, x int, y int, nWidth int, nHeight int) int {]
ff=SVGEX
nHeight=高度
nWidth=宽度
y=y坐标
x=x坐标
hSvg=SVG句柄

[func (d *Draw) DrawSvgSize(hSvg int, nWidth int, nHeight int) int {]
ff=SVG大小
nHeight=高度
nWidth=宽度
hSvg=SVG句柄

[func (d *Draw) D2D_Clear(color int) int {]
ff=D2D清理
color=ABGR

[func (d *Draw) ImageMaskRect(hImageFrame int, pRect *xc.RECT, pRcMask *xc.RECT, pRcRoundAngle *xc.RECT) int {]
ff=图片遮盖矩形
pRect=矩形坐标
hImageFrame=图片句柄

[func (d *Draw) ImageMaskEllipse(hImageFrame int, pRect *xc.RECT, pRcMask *xc.RECT) int {]
ff=图片遮盖圆型
pRect=矩形坐标
hImageFrame=图片句柄

[func (d *Draw) GetFont() int {]
ff=取字体
