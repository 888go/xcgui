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

[func XDraw_Create(hWindow int) int {]
ff=绘制_创建
hWindow=窗口句柄

[func XDraw_CreateGDI(hWindow int, hdc uintptr) int {]
ff=绘制_创建GDI
hdc=hdc句柄
hWindow=窗口句柄

[func XDraw_Destroy(hDraw int) int {]
ff=绘制_销毁
hDraw=图形绘制句柄

[func XDraw_Dottedline(hDraw int, x1 int, y1 int, x2 int, y2 int) int {]
ff=绘制_虚线
y2=结束点y坐标
x2=结束点x坐标
y1=起点y坐标
x1=起点x坐标
hDraw=图形绘制句柄

[func XDraw_DottedlineF(hDraw int, x1, y1, x2, y2 float32) int {]
ff=绘制_虚线F
y2=结束点y坐标
x2=结束点x坐标
y1=起点y坐标
x1=起点x坐标
hDraw=图形绘制句柄

[func XDraw_DrawArc(hDraw int, x, y int, nWidth int, nHeight int, startAngle float32, sweepAngle float32) int {]
ff=绘制_圆弧
sweepAngle=绘制角度
startAngle=起始角度
nHeight=高度
nWidth=宽度
y=坐标y
x=坐标x
hDraw=图形绘制句柄

[func XDraw_DrawArcF(hDraw int, x, y, nWidth, nHeight, startAngle, sweepAngle float32) int {]
ff=绘制_圆弧F
sweepAngle=绘制角度
startAngle=起始角度
nHeight=高度
nWidth=宽度
y=坐标y
x=坐标x
hDraw=图形绘制句柄

[func XDraw_DrawCurve(hDraw int, points #左中括号##右中括号#POINT, count int, tension float32) int {]
ff=绘制_曲线
tension=曲线张力
count=切片大小
points=坐标点切片
hDraw=图形绘制句柄

[func XDraw_DrawCurveF(hDraw int, points #左中括号##右中括号#POINTF, count int, tension float32) int {]
ff=绘制_曲线F
tension=曲线的张力
count=切片大小
points=坐标点切片
hDraw=图形绘制句柄

[func XDraw_DrawLine(hDraw int, x1 int, y1 int, x2 int, y2 int) int {]
ff=绘制_线条
hDraw=图形绘制句柄

[func XDraw_DrawLineF(hDraw int, x1, y1, x2, y2 float32) int {]
ff=绘制_线条F
hDraw=图形绘制句柄

[func XDraw_DrawPolygon(hDraw int, points #左中括号##右中括号#POINT, nCount int) int {]
ff=绘制_多边形
nCount=顶点数量
points=顶点坐标切片
hDraw=图形绘制句柄

[func XDraw_DrawPolygonF(hDraw int, points #左中括号##右中括号#POINTF, nCount int) int {]
ff=绘制_多边形F
nCount=顶点数量
points=顶点坐标切片
hDraw=图形绘制句柄

[func XDraw_DrawRect(hDraw int, pRect *RECT) int {]
ff=绘制_矩形边框
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_DrawRectF(hDraw int, pRect *RECTF) int {]
ff=绘制_矩形边框F
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_SetOffset(hDraw int, x, y int32) int {]
ff=绘制_置偏移
y=Y轴偏移量
x=轴偏移量
hDraw=图形绘制句柄

[func XDraw_GetOffset(hDraw int, pX, pY *int32) int {]
ff=绘制_取偏移
pY=接收Y轴偏移量
pX=接收X轴偏移量
hDraw=图形绘制句柄

[func XDraw_GDI_RestoreGDIOBJ(hDraw int) int {]
ff=绘制_还原状态
hDraw=图形绘制句柄

[func XDraw_GetHDC(hDraw int) uintptr {]
ff=绘制_取HDC
hDraw=图形绘制句柄

[func XDraw_SetBrushColor(hDraw int, color int) int {]
ff=绘制_置画刷颜色
color=ABGR颜色值
hDraw=图形绘制句柄

[func XDraw_SetTextVertical(hDraw int, bVertical bool) int {]
ff=绘制_置文本垂直
bVertical=是否垂直显示文本
hDraw=图形绘制句柄

[func XDraw_SetTextAlign(hDraw int, nFlags xcc.TextFormatFlag_) int {]
ff=绘制_置文本对齐
nFlags=对齐标识
hDraw=图形绘制句柄

[func XDraw_SetFont(hDraw int, hFontx int) int {]
ff=绘制_置字体
hFontx=炫彩字体
hDraw=图形绘制句柄

[func XDraw_SetLineWidth(hDraw int, nWidth int) int {]
ff=绘制_置线宽
nWidth=宽度
hDraw=图形绘制句柄

[func XDraw_SetLineWidthF(hDraw int, nWidth float32) int {]
ff=绘制_置线宽F
nWidth=宽度
hDraw=图形绘制句柄

[func XDraw_GDI_SetBkMode(hDraw int, bTransparent bool) int {]
ff=绘制_置背景模式
hDraw=图形绘制句柄

[func XDraw_SetClipRect(hDraw int, pRect *RECT) int {]
ff=绘制_置裁剪区域
pRect=区域坐标
hDraw=图形绘制句柄

[func XDraw_SetD2dTextRenderingMode(hDraw int, mode xcc.XC_DWRITE_RENDERING_MODE_) int {]
ff=绘制_置D2D文本渲染模式
hDraw=图形绘制句柄

[func XDraw_ClearClip(hDraw int) int {]
ff=绘制_清除裁剪区域
hDraw=图形绘制句柄

[func XDraw_EnableSmoothingMode(hDraw int, bEnable bool) int {]
ff=绘制_启用平滑模式
bEnable=是否启用
hDraw=图形绘制句柄

[func XDraw_EnableWndTransparent(hDraw int, bTransparent bool) int {]
ff=绘制_启用窗口透明判断
bTransparent=是否启用
hDraw=图形绘制句柄

[func XDraw_GDI_CreateSolidBrush(hDraw int, crColor int) int {]
ff=绘制_创建实心画刷
crColor=画刷颜色
hDraw=图形绘制句柄

[func XDraw_GDI_CreatePen(hDraw int, fnPenStyle int, nWidth int, crColor int) int {]
ff=绘制_创建画笔
crColor=RGB颜色
nWidth=画笔宽度
fnPenStyle=画笔样式
hDraw=图形绘制句柄

[func XDraw_GDI_CreateRectRgn(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) int {]
ff=绘制_创建矩形区域
nBottomRect=右下角Y坐标
nRightRect=右下角X坐标
nTopRect=左上角Y坐标
nLeftRect=左上角X坐标
hDraw=图形绘制句柄

[func XDraw_GDI_CreateRoundRectRgn(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int, nWidthEllipse int, nHeightEllipse int) int {]
ff=绘制_创建圆角矩形区域
nHeightEllipse=椭圆的高度
nWidthEllipse=椭圆的宽度
nBottomRect=右下角Y坐标
nRightRect=右下角X坐标
nTopRect=左上角Y坐标
nLeftRect=左上角X坐标
hDraw=图形绘制句柄

[func XDraw_GDI_CreatePolygonRgn(hDraw int, pPt #左中括号##右中括号#POINT, cPoints int, fnPolyFillMode int) int {]
ff=绘制_创建多边形区域
fnPolyFillMode=多边形填充模式
cPoints=切片大小
pPt=POINT切片
hDraw=图形绘制句柄

[func XDraw_GDI_SelectClipRgn(hDraw int, hRgn int) int {]
ff=绘制_选择裁剪区域
hRgn=区域句柄
hDraw=图形绘制句柄

[func XDraw_FillRect(hDraw int, pRect *RECT) int {]
ff=绘制_填充矩形
pRect=矩形区域
hDraw=图形绘制句柄

[func XDraw_FillRectF(hDraw int, pRect *RECTF) int {]
ff=绘制_填充矩形F
pRect=矩形区域
hDraw=图形绘制句柄

[func XDraw_FillRectColor(hDraw int, pRect *RECT, color int) int {]
ff=绘制_填充矩形指定颜色
color=ABGR颜色
pRect=矩形区域
hDraw=图形绘制句柄

[func XDraw_FillRectColorF(hDraw int, pRect *RECTF, color int) int {]
ff=绘制_填充矩形指定颜色F
color=ABGR颜色
pRect=矩形区域
hDraw=图形绘制句柄

[func XDraw_GDI_FillRgn(hDraw int, hrgn int, hbr int) bool {]
ff=绘制_填充区域
hbr=画刷句柄
hrgn=区域句柄
hDraw=图形绘制句柄

[func XDraw_FillEllipse(hDraw int, pRect *RECT) int {]
ff=绘制_填充圆形
pRect=矩形区域
hDraw=图形绘制句柄

[func XDraw_FillEllipseF(hDraw int, pRect *RECTF) int {]
ff=绘制_填充圆形F
pRect=矩形区域
hDraw=图形绘制句柄

[func XDraw_DrawEllipse(hDraw int, pRect *RECT) int {]
ff=绘制_圆形
pRect=矩形区域
hDraw=图形绘制句柄

[func XDraw_FillRoundRect(hDraw int, pRect *RECT, nWidth, nHeight int) int {]
ff=绘制_填充圆角矩形
nHeight=圆角高度
nWidth=圆角宽度
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_FillRoundRectF(hDraw int, pRect *RECTF, nWidth, nHeight float32) int {]
ff=绘制_填充圆角矩形F
nHeight=圆角高度
nWidth=圆角宽度
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_DrawRoundRect(hDraw int, pRect *RECT, nWidth int, nHeight int) int {]
ff=绘制_圆角矩形
nHeight=圆角高度
nWidth=圆角宽度
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_DrawRoundRectF(hDraw int, pRect *RECT, nWidth, nHeight float32) int {]
ff=绘制_圆角矩形F
nHeight=圆角高度
nWidth=圆角宽度
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_FillRoundRectEx(hDraw int, pRect *RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom int) int {]
ff=绘制_填充圆角矩形EX
nLeftBottom=左下角
nRightBottom=右下角
nRightTop=右上角
nLeftTop=左上角
pRect=坐标
hDraw=图形绘制句柄

[func XDraw_FillRoundRectExF(hDraw int, pRect *RECTF, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {]
ff=绘制_填充圆角矩形EXF
nLeftBottom=左下角
nRightBottom=右下角
nRightTop=右上角
nLeftTop=左上角
pRect=坐标
hDraw=图形绘制句柄

[func XDraw_DrawRoundRectEx(hDraw int, pRect *RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {]
ff=绘制_圆角矩形EX
nLeftBottom=左下角
nRightBottom=右下角
nRightTop=右上角
nLeftTop=左上角
pRect=坐标
hDraw=图形绘制句柄

[func XDraw_DrawRoundRectExF(hDraw int, pRect *RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {]
ff=绘制_圆角矩形EXF
nLeftBottom=左下角
nRightBottom=右下角
nRightTop=右上角
nLeftTop=左上角
pRect=坐标
hDraw=图形绘制句柄

[func XDraw_GDI_Rectangle(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) bool {]
ff=绘制_矩形
nBottomRect=右下角Y坐标
nRightRect=右下角X坐标
nTopRect=左上角Y坐标
nLeftRect=左上角X坐标
hDraw=图形绘制句柄

[func XDraw_GradientFill2(hDraw int, pRect *RECT, color1 int, color2 int, mode xcc.GRADIENT_FILL_) int {]
ff=绘制_渐变填充2
mode=模式
color2=结束颜色
color1=开始颜色
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_GradientFill2F(hDraw int, pRect *RECTF, color1 int, color2 int, mode xcc.GRADIENT_FILL_) int {]
ff=绘制_渐变填充2F
mode=模式
color2=结束颜色
color1=开始颜色
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_GradientFill4(hDraw int, pRect *RECT, color1 int, color2 int, color3 int, color4 int, mode xcc.GRADIENT_FILL_) bool {]
ff=绘制_渐变填充4
mode=模式
color4=颜色4
color3=颜色3
color2=颜色2
color1=颜色1
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_GradientFill4F(hDraw int, pRect *RECTF, color1 int, color2 int, color3 int, color4 int, mode xcc.GRADIENT_FILL_) bool {]
ff=绘制_渐变填充4F
mode=模式
color4=颜色4
color3=颜色3
color2=颜色2
color1=颜色1
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_GDI_FrameRgn(hDraw int, hrgn int, hbr int, nWidth int, nHeight int) bool {]
ff=绘制_边框区域
nHeight=边框高度
nWidth=边框宽度
hbr=画刷句柄
hrgn=区域句柄
hDraw=图形绘制句柄

[func XDraw_FocusRect(hDraw int, pRect *RECT) int {]
ff=绘制_焦点矩形
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_FocusRectF(hDraw int, pRect *RECTF) int {]
ff=绘制_焦点矩形F
pRect=矩形坐标
hDraw=图形绘制句柄

[func XDraw_GDI_MoveToEx(hDraw int, X int, Y int, pPoint *POINT) bool {]
ff=绘制_移动到起点
pPoint=接收以前的当前位置到一个POINT结构的指针
Y=坐标y
X=坐标x
hDraw=图形绘制句柄

[func XDraw_GDI_LineTo(hDraw int, nXEnd int, nYEnd int) bool {]
ff=绘制_线终点
nYEnd=Y坐标
nXEnd=坐标
hDraw=图形绘制句柄

[func XDraw_GDI_Polyline(hDraw int, pArrayPt #左中括号##右中括号#POINT, arrayPtSize int) bool {]
ff=绘制_折线
hDraw=图形绘制句柄

[func XDraw_GDI_SetPixel(hDraw int, X int, Y int, crColor int) int {]
ff=绘制_置像素颜色
crColor=RGB颜色值
Y=坐标y
X=坐标x
hDraw=图形绘制句柄

[func XDraw_GetD2dRenderTarget(hDraw int) int {]
ff=绘制_取D2D渲染目标
hDraw=图形绘制句柄

[func XDraw_GDI_DrawIconEx(hDraw int, xLeft int, yTop int, hIcon uintptr, cxWidth int, cyWidth int, istepIfAniCur int, hbrFlickerFreeDraw int, diFlags int) bool {]
ff=绘制_图标

[func XDraw_GDI_BitBlt(hDrawDest int, nXDest, nYDest, nWidth, nHeight int32, hdcSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {]
ff=绘制_复制

[func XDraw_GDI_BitBlt2(hDrawDest int, nXDest, nYDest, nWidth, nHeight int32, hDrawSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {]
ff=绘制_复制2

[func XDraw_GDI_AlphaBlend(hDraw int, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int32, hdcSrc uintptr, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha int32) bool {]
ff=绘制_带透明复制

[func XDraw_GDI_Ellipse(hDraw int, pRect *RECT) bool {]
ff=绘制_GDI_椭圆
pRect=矩形区域
hDraw=图形绘制句柄

[func XDraw_FillPolygon(hDraw int, points #左中括号##右中括号#POINT, nCount int) int {]
ff=绘制_填充多边形
nCount=顶点数量
points=顶点坐标切片
hDraw=图形绘制句柄

[func XDraw_FillPolygonF(hDraw int, points #左中括号##右中括号#POINTF, nCount int) int {]
ff=绘制_填充多边形F
nCount=顶点数量
points=顶点坐标切片
hDraw=图形绘制句柄

[func XDraw_Image(hDraw int, hImageFrame int, x, y int32) {]
ff=绘制_图片
y=y坐标
x=x坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageF(hDraw int, hImageFrame int, x, y float32) int {]
ff=绘制_图片F
y=y坐标
x=x坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageAdaptive(hDraw int, hImageFrame int, pRect *RECT, bOnlyBorder bool) int {]
ff=绘制_图片自适应
bOnlyBorder=是否只绘制边缘区域
pRect=坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageAdaptiveF(hDraw int, hImageFrame int, pRect *RECTF, bOnlyBorder bool) int {]
ff=绘制_图片自适应F
bOnlyBorder=是否只绘制边缘区域
pRect=坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageEx(hDraw int, hImageFrame int, x, y, width, height int) int {]
ff=绘制_图片EX
height=高度
width=宽度
y=y坐标
x=x坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageExF(hDraw int, hImageFrame int, x, y, width, height float32) int {]
ff=绘制_图片EXF
height=高度
width=宽度
y=y坐标
x=x坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageSuper(hDraw int, hImageFrame int, pRect *RECT, bClip bool) int {]
ff=绘制_图片增强
bClip=是否裁剪区域
pRect=坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageSuperF(hDraw int, hImageFrame int, pRect *RECTF, bClip bool) int {]
ff=绘制_图片增强F
bClip=是否裁剪区域
pRect=坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageSuperEx(hDraw int, hImageFrame int, prcDest *RECT, prcSrc *RECT) int {]
ff=绘制_图片增强EX
prcSrc=源坐标
prcDest=目标坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageSuperExF(hDraw int, hImageFrame int, prcDest *RECTF, prcSrc *RECT) int {]
ff=绘制_图片增强EXF
prcSrc=源坐标
prcDest=目标坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageSuperMask(hDraw int, hImageFrame int, hImageFrameMask int, pRect *RECT, pRectMask *RECT, bClip bool) int {]
ff=绘制_图片增强遮盖
bClip=是否裁剪区域
pRectMask=遮盖坐标
pRect=坐标
hImageFrameMask=图片遮盖句柄
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageTile(hDraw int, hImageFrame int, hImageFrameMask int, pRect *RECT, flag int) int {]
ff=绘制_图片平铺
flag=标识
pRect=坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageTileF(hDraw int, hImageFrame int, hImageFrameMask int, pRect *RECTF, flag int) int {]
ff=绘制_图片平铺F
flag=标识
pRect=坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageMask(hDraw int, hImageFrame int, hImageFrameMask int, x int, y int, x2 int, y2 int) int {]
ff=绘制_图片遮盖
y2=坐标y2
x2=坐标x2
y=坐标y
x=坐标x
hImageFrameMask=图片遮盖句柄
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageMaskRect(hDraw int, hImageFrame int, pRect *RECT, pRcMask *RECT, pRcRoundAngle *RECT) int {]
ff=绘制_图片遮盖矩形
pRcRoundAngle=遮罩圆角
pRcMask=遮罩坐标
pRect=矩形坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_ImageMaskEllipse(hDraw int, hImageFrame int, pRect *RECT, pRcMask *RECT) int {]
ff=绘制_图片遮盖圆型
pRcMask=遮罩坐标
pRect=矩形坐标
hImageFrame=图片句柄
hDraw=图形绘制句柄

[func XDraw_DrawText(hDraw int, lpString string, lpRect *RECT) int {]
ff=绘制_文本指定矩形
lpRect=坐标
lpString=字符串
hDraw=图形绘制句柄

[func XDraw_DrawTextF(hDraw int, lpString string, lpRect *RECTF) int {]
ff=绘制_文本指定矩形F
lpRect=坐标
lpString=字符串
hDraw=图形绘制句柄

[func XDraw_DrawTextUnderline(hDraw int, lpString string, lpRect *RECT, colorLine int) int {]
ff=绘制_文本下划线
colorLine=下划线颜色
lpRect=坐标
lpString=字符串
hDraw=图形绘制句柄

[func XDraw_DrawTextUnderlineF(hDraw int, lpString string, lpRect *RECTF, colorLine int) int {]
ff=绘制_文本下划线F
colorLine=下划线颜色
lpRect=坐标
lpString=字符串
hDraw=图形绘制句柄

[func XDraw_TextOut(hDraw int, nXStart int, nYStart int, lpString string, cbString string) int {]
ff=绘制_文本
hDraw=图形绘制句柄

[func XDraw_TextOutF(hDraw int, nXStart, nYStart float32, lpString string, cbString string) int {]
ff=绘制_文本F
hDraw=图形绘制句柄

[func XDraw_TextOutEx(hDraw int, nXStart int, nYStart int, lpString string) int {]
ff=绘制_文本EX
hDraw=图形绘制句柄

[func XDraw_TextOutExF(hDraw int, nXStart, nYStart float32, lpString string) int {]
ff=绘制_文本EXF
hDraw=图形绘制句柄

[func XDraw_TextOutA(hDraw int, nXStart int, nYStart int, lpString string) int {]
ff=绘制_文本A
hDraw=图形绘制句柄

[func XDraw_TextOutAF(hDraw int, nXStart, nYStart float32, lpString string) int {]
ff=绘制_文本AF
hDraw=图形绘制句柄

[func XDraw_SetTextRenderingHint(hDraw int, nType int) int {]
ff=绘制_设置文本渲染提示
hDraw=图形绘制句柄

[func XDraw_DrawSvgSrc(hDraw int, hSvg int) int {]
ff=绘制_SVG源
hSvg=SVG句柄
hDraw=图形绘制句柄

[func XDraw_DrawSvg(hDraw int, hSvg int, x int, y int) int {]
ff=绘制_SVG
y=y坐标
x=x坐标
hSvg=SVG句柄
hDraw=图形绘制句柄

[func XDraw_DrawSvgEx(hDraw int, hSvg int, x int, y int, nWidth int, nHeight int) int {]
ff=绘制_SVGEX
nHeight=高度
nWidth=宽度
y=y坐标
x=x坐标
hSvg=SVG句柄
hDraw=图形绘制句柄

[func XDraw_DrawSvgSize(hDraw int, hSvg int, nWidth int, nHeight int) int {]
ff=绘制_SVG大小
nHeight=高度
nWidth=宽度
hSvg=SVG句柄
hDraw=图形绘制句柄

[func XDraw_D2D_Clear(hDraw int, color int) int {]
ff=绘制_D2D_清理
color=ABGR颜色值
hDraw=图形绘制句柄

[func XDraw_GetFont(hDraw int) int {]
ff=绘制_取字体
hDraw=图形绘制句柄
