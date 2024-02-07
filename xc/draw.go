package 炫彩基类

import (
	"unsafe"

	"github.com/888go/xcgui/common"

	"github.com/888go/xcgui/xcc"
)

// 绘制_创建, 创建图形绘制模块实例, 返回句柄.
//
// hWindow: 窗口句柄.
func X绘制_创建(窗口句柄 int) int {
	r, _, _ := xDraw_Create.Call(uintptr(窗口句柄))
	return int(r)
}

// 绘制_创建GDI, 创建图形绘制模块实例, 返回图形绘制模块实例句柄.
//
// hWindow: 窗口句柄.
//
// hdc: hdc句柄.
func X绘制_创建GDI(窗口句柄 int, hdc句柄 uintptr) int {
	r, _, _ := xDraw_CreateGDI.Call(uintptr(窗口句柄), hdc句柄)
	return int(r)
}

// 绘制_销毁, 销毁图形绘制模块实例句柄.
//
// hDraw: 图形绘制句柄.
func X绘制_销毁(图形绘制句柄 int) int {
	r, _, _ := xDraw_Destroy.Call(uintptr(图形绘制句柄))
	return int(r)
}

// 绘制_虚线, 绘制水平或垂直虚线.
//
// hDraw: 图形绘制句柄.
//
// x1: 起点x坐标.
//
// y1: 起点y坐标.
//
// x2: 结束点x坐标.
//
// y2: 结束点y坐标.
func X绘制_虚线(图形绘制句柄 int, 起点x坐标 int, 起点y坐标 int, 结束点x坐标 int, 结束点y坐标 int) int {
	r, _, _ := xDraw_Dottedline.Call(uintptr(图形绘制句柄), uintptr(起点x坐标), uintptr(起点y坐标), uintptr(结束点x坐标), uintptr(结束点y坐标))
	return int(r)
}

// 绘制_虚线F, 绘制水平或垂直虚线.
//
// hDraw: 图形绘制句柄.
//
// x1: 起点x坐标.
//
// y1: 起点y坐标.
//
// x2: 结束点x坐标.
//
// y2: 结束点y坐标.
func X绘制_虚线F(图形绘制句柄 int, 起点x坐标, 起点y坐标, 结束点x坐标, 结束点y坐标 float32) int {
	r, _, _ := xDraw_DottedlineF.Call(uintptr(图形绘制句柄), 炫彩工具类.Float32Ptr(起点x坐标), 炫彩工具类.Float32Ptr(起点y坐标), 炫彩工具类.Float32Ptr(结束点x坐标), 炫彩工具类.Float32Ptr(结束点y坐标))
	return int(r)
}

// 绘制_圆弧.
//
// hDraw: 图形绘制句柄.
//
// x: 坐标.
//
// y: 坐标.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// startAngle: 起始角度.
//
// sweepAngle: 绘制角度, 从起始角度开始计算.
func X绘制_圆弧(图形绘制句柄 int, 坐标x, 坐标y int, 宽度 int, 高度 int, 起始角度 float32, 绘制角度 float32) int {
	r, _, _ := xDraw_DrawArc.Call(uintptr(图形绘制句柄), uintptr(坐标x), uintptr(坐标y), uintptr(宽度), uintptr(高度), 炫彩工具类.Float32Ptr(起始角度), 炫彩工具类.Float32Ptr(绘制角度))
	return int(r)
}

// 绘制_圆弧F.
//
// hDraw: 图形绘制句柄.
//
// x: 坐标.
//
// y: 坐标.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// startAngle: 起始角度.
//
// sweepAngle: 绘制角度, 从起始角度开始计算.
func X绘制_圆弧F(图形绘制句柄 int, 坐标x, 坐标y, 宽度, 高度, 起始角度, 绘制角度 float32) int {
	r, _, _ := xDraw_DrawArcF.Call(uintptr(图形绘制句柄), 炫彩工具类.Float32Ptr(坐标x), 炫彩工具类.Float32Ptr(坐标y), 炫彩工具类.Float32Ptr(宽度), 炫彩工具类.Float32Ptr(高度), 炫彩工具类.Float32Ptr(起始角度), 炫彩工具类.Float32Ptr(绘制角度))
	return int(r)
}

// 绘制_曲线, D2D暂时留空.
//
// hDraw: 图形绘制句柄.
//
// points: 坐标点数组.
//
// count: 数组大小.
//
// tension: 大于或等于0.0F的值，指定曲线的张力, D2D 忽略此参数.
func X绘制_曲线(图形绘制句柄 int, 坐标点数组 []POINT, 数组大小 int, 曲线张力 float32) int {
	r, _, _ := xDraw_DrawCurve.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(&坐标点数组[0])), uintptr(数组大小), 炫彩工具类.Float32Ptr(曲线张力))
	return int(r)
}

// 绘制_曲线F, D2D暂时留空.
//
// hDraw: 图形绘制句柄.
//
// points: 坐标点数组.
//
// count: 数组大小.
//
// tension: 大于或等于0.0F的值，指定曲线的张力, D2D 忽略此参数.
func X绘制_曲线F(图形绘制句柄 int, 坐标点数组 []POINTF, 数组大小 int, 曲线的张力 float32) int {
	r, _, _ := xDraw_DrawCurveF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(&坐标点数组[0])), uintptr(数组大小), 炫彩工具类.Float32Ptr(曲线的张力))
	return int(r)
}

// 绘制_线条.
//
// hDraw: 图形绘制句柄.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func X绘制_线条(图形绘制句柄 int, x1 int, y1 int, x2 int, y2 int) int {
	r, _, _ := xDraw_DrawLine.Call(uintptr(图形绘制句柄), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return int(r)
}

// 绘制_线条F.
//
// hDraw: 图形绘制句柄.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func X绘制_线条F(图形绘制句柄 int, x1, y1, x2, y2 float32) int {
	r, _, _ := xDraw_DrawLineF.Call(uintptr(图形绘制句柄), 炫彩工具类.Float32Ptr(x1), 炫彩工具类.Float32Ptr(y1), 炫彩工具类.Float32Ptr(x2), 炫彩工具类.Float32Ptr(y2))
	return int(r)
}

// 绘制_多边形, 绘制多边形.
//
// hDraw: 图形绘制句柄.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func X绘制_多边形(图形绘制句柄 int, 顶点坐标数组 []POINT, 顶点数量 int) int {
	r, _, _ := xDraw_DrawPolygon.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(&顶点坐标数组[0])), uintptr(顶点数量))
	return int(r)
}

// 绘制_多边形F, 绘制多边形.
//
// hDraw: 图形绘制句柄.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func X绘制_多边形F(图形绘制句柄 int, 顶点坐标数组 []POINTF, 顶点数量 int) int {
	r, _, _ := xDraw_DrawPolygonF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(&顶点坐标数组[0])), uintptr(顶点数量))
	return int(r)
}

// 绘制_矩形, 绘制矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标 .
func X绘制_矩形边框(图形绘制句柄 int, 矩形坐标 *RECT) int {
	r, _, _ := xDraw_DrawRect.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)))
	return int(r)
}

// 绘制_矩形F, 绘制矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标 .
func X绘制_矩形边框F(图形绘制句柄 int, 矩形坐标 *RECTF) int {
	r, _, _ := xDraw_DrawRectF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)))
	return int(r)
}

// 绘制_置偏移, 设置坐标偏移量, X向左偏移为负数, 向右偏移为正数.
//
// hDraw: 图形绘制句柄.
//
// x: X轴偏移量.
//
// y: Y轴偏移量.
func X绘制_置偏移(图形绘制句柄 int, X轴偏移量, Y轴偏移量 int32) int {
	r, _, _ := xDraw_SetOffset.Call(uintptr(图形绘制句柄), uintptr(X轴偏移量), uintptr(Y轴偏移量))
	return int(r)
}

// 绘制_取偏移, 获取坐标偏移量, X向左偏移为负数, 向右偏移为正数.
//
// hDraw: 图形绘制句柄.
//
// pX: 接收X轴偏移量.
//
// pY: 接收Y轴偏移量.
func X绘制_取偏移(图形绘制句柄 int, 接收X轴偏移量, 接收Y轴偏移量 *int32) int {
	r, _, _ := xDraw_GetOffset.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(接收X轴偏移量)), uintptr(unsafe.Pointer(接收Y轴偏移量)))
	return int(r)
}

// 绘制_还原状态, 还原状态, 释放用户绑定的GDI对象, 例如画刷, 画笔.
//
// hDraw: 图形绘制句柄.
func X绘制_还原状态(图形绘制句柄 int) int {
	r, _, _ := xDraw_GDI_RestoreGDIOBJ.Call(uintptr(图形绘制句柄))
	return int(r)
}

// 绘制_取HDC, 获取绑定的设备上下文HDC, 返回HDC句柄.
//
// hDraw: 图形绘制句柄.
func X绘制_取HDC(图形绘制句柄 int) uintptr {
	r, _, _ := xDraw_GetHDC.Call(uintptr(图形绘制句柄))
	return r
}

// 绘制_置画刷颜色, 设置画刷颜色.
//
// hDraw: 图形绘制句柄.
//
// color: ABGR 颜色值.
func X绘制_置画刷颜色(图形绘制句柄 int, ABGR颜色值 int) int {
	r, _, _ := xDraw_SetBrushColor.Call(uintptr(图形绘制句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 绘制_置文本垂直, 设置文本垂直显示.
//
// hDraw: 图形绘制句柄.
//
// bVertical: 是否垂直显示文本.
func X绘制_置文本垂直(图形绘制句柄 int, 是否垂直显示文本 bool) int {
	r, _, _ := xDraw_SetTextVertical.Call(uintptr(图形绘制句柄), 炫彩工具类.BoolPtr(是否垂直显示文本))
	return int(r)
}

// 绘制_置文本对齐, 设置文本对齐.
//
// hDraw: 图形绘制句柄.
//
// nFlags: 对齐标识, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func X绘制_置文本对齐(图形绘制句柄 int, 对齐标识 炫彩常量类.TextFormatFlag_) int {
	r, _, _ := xDraw_SetTextAlign.Call(uintptr(图形绘制句柄), uintptr(对齐标识))
	return int(r)
}

// 绘制_置字体.
//
// hDraw: 图形绘制句柄.
//
// hFontx: 炫彩字体.
func X绘制_置字体(图形绘制句柄 int, 炫彩字体 int) int {
	r, _, _ := xDraw_SetFont.Call(uintptr(图形绘制句柄), uintptr(炫彩字体))
	return int(r)
}

// 绘制_置线宽.
//
// hDraw: 图形绘制句柄.
//
// nWidth: 宽度.
func X绘制_置线宽(图形绘制句柄 int, 宽度 int) int {
	r, _, _ := xDraw_SetLineWidth.Call(uintptr(图形绘制句柄), uintptr(宽度))
	return int(r)
}

// 绘制_置线宽F.
//
// hDraw: 图形绘制句柄.
//
// nWidth: 宽度.
func X绘制_置线宽F(图形绘制句柄 int, 宽度 float32) int {
	r, _, _ := xDraw_SetLineWidthF.Call(uintptr(图形绘制句柄), 炫彩工具类.Float32Ptr(宽度))
	return int(r)
}

// 绘制_置背景模式, SetBkMode() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// bTransparent: 参见MSDN.
func X绘制_置背景模式(图形绘制句柄 int, bTransparent bool) int {
	r, _, _ := xDraw_GDI_SetBkMode.Call(uintptr(图形绘制句柄), 炫彩工具类.BoolPtr(bTransparent))
	return int(r)
}

// 绘制_置裁剪区域, 设置裁剪区域.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func X绘制_置裁剪区域(图形绘制句柄 int, 区域坐标 *RECT) int {
	r, _, _ := xDraw_SetClipRect.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(区域坐标)))
	return int(r)
}

// 绘制_置D2D文本渲染模式.
//
// hDraw: 图形绘制句柄.
//
// mode	渲染模式 XC_DWRITE_RENDERING_MODE_.
func X绘制_置D2D文本渲染模式(图形绘制句柄 int, mode 炫彩常量类.XC_DWRITE_RENDERING_MODE_) int {
	r, _, _ := xDraw_SetD2dTextRenderingMode.Call(uintptr(图形绘制句柄), uintptr(mode))
	return int(r)
}

// 绘制_清除裁剪区域.
//
// hDraw: 图形绘制句柄.
func X绘制_清除裁剪区域(图形绘制句柄 int) int {
	r, _, _ := xDraw_ClearClip.Call(uintptr(图形绘制句柄))
	return int(r)
}

// 绘制_启用平滑模式.
//
// hDraw: 图形绘制句柄.
//
// bEnable: 是否启用.
func X绘制_启用平滑模式(图形绘制句柄 int, 是否启用 bool) int {
	r, _, _ := xDraw_EnableSmoothingMode.Call(uintptr(图形绘制句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 绘制_启用窗口透明判断, 当启用之后, 调用GDI+函数时, 如果参数alpha=255, 将自动修改为254, 应对GDI+的bug, 否则透明通道异常.
//
// hDraw: 图形绘制句柄.
//
// bTransparent: 是否启用.
func X绘制_启用窗口透明判断(图形绘制句柄 int, 是否启用 bool) int {
	r, _, _ := xDraw_EnableWndTransparent.Call(uintptr(图形绘制句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 绘制_创建实心画刷, GDI创建具有指定的纯色逻辑刷.
//
// hDraw: 图形绘制句柄.
//
// crColor: 画刷颜色, RGB颜色.
func X绘制_创建实心画刷(图形绘制句柄 int, 画刷颜色 int) int {
	r, _, _ := xDraw_GDI_CreateSolidBrush.Call(uintptr(图形绘制句柄), uintptr(画刷颜色))
	return int(r)
}

// 绘制_创建画笔, GDI创建一个逻辑笔, 指定的样式, 宽度和颜色, 随后的笔可以选择到设备上下文, 用于绘制线条和曲线.
//
// hDraw: 图形绘制句柄.
//
// fnPenStyle: 画笔样式, PS_SOLID:实线, PS_DASH:段线, PS_DOT:点线, PS_DASHDOT:段线_点线, PS_DASHDOTDOT:段线_点_点, PS_NULL:空, PS_INSIDEFRAME:实线_笔宽是向里扩展.
//
// nWidth: 画笔宽度.
//
// crColor: RGB颜色.
func X绘制_创建画笔(图形绘制句柄 int, 画笔样式 int, 画笔宽度 int, RGB颜色 int) int {
	r, _, _ := xDraw_GDI_CreatePen.Call(uintptr(图形绘制句柄), uintptr(画笔样式), uintptr(画笔宽度), uintptr(RGB颜色))
	return int(r)
}

// 绘制_创建矩形区域, GDI创建矩形区域, 成功返回区域句柄, 失败返回NULL.
//
// hDraw: 图形绘制句柄.
//
// nLeftRect: 左上角X坐标.
//
// nTopRect: 左上角Y坐标.
//
// nRightRect: 右下角X坐标.
//
// nBottomRect: 右下角Y坐标.
func X绘制_创建矩形区域(图形绘制句柄 int, 左上角X坐标 int, 左上角Y坐标 int, 右下角X坐标 int, 右下角Y坐标 int) int {
	r, _, _ := xDraw_GDI_CreateRectRgn.Call(uintptr(图形绘制句柄), uintptr(左上角X坐标), uintptr(左上角Y坐标), uintptr(右下角X坐标), uintptr(右下角Y坐标))
	return int(r)
}

// 绘制_创建圆角矩形区域, GDI创建一个圆角的矩形区域, 成功返回区域句柄, 失败返回NULL.
//
// hDraw: 图形绘制句柄.
//
// nLeftRect: X-坐标的左上角.
//
// nTopRect: Y-坐标左上角坐标.
//
// nRightRect: X-坐标右下角.
//
// nBottomRect: Y-坐标右下角.
//
// nWidthEllipse: 椭圆的宽度.
//
// nHeightEllipse: 椭圆的高度.
func X绘制_创建圆角矩形区域(图形绘制句柄 int, 左上角X坐标 int, 左上角Y坐标 int, 右下角X坐标 int, 右下角Y坐标 int, 椭圆的宽度 int, 椭圆的高度 int) int {
	r, _, _ := xDraw_GDI_CreateRoundRectRgn.Call(uintptr(图形绘制句柄), uintptr(左上角X坐标), uintptr(左上角Y坐标), uintptr(右下角X坐标), uintptr(右下角Y坐标), uintptr(椭圆的宽度), uintptr(椭圆的高度))
	return int(r)
}

// 绘制_创建多边形区域, GDI创建一个多边形区域, 成功返回区域句柄, 失败返回NULL.
//
// hDraw: 图形绘制句柄.
//
// pPt: POINT数组.
//
// cPoints: 数组大小.
//
// fnPolyFillMode: 多边形填充模式, 指定用于确定在该地区的像素填充模式,这个参数可以是下列值之一.
//
// ALTERNATE Selects alternate mode (fills area between odd-numbered and even-numbered polygon sides on each scan line).
//
// WINDING Selects winding mode (fills any region with a nonzero winding value).
func X绘制_创建多边形区域(图形绘制句柄 int, POINT数组 []POINT, 数组大小 int, 多边形填充模式 int) int {
	r, _, _ := xDraw_GDI_CreatePolygonRgn.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(&POINT数组[0])), uintptr(数组大小), uintptr(多边形填充模式))
	return int(r)
}

// 绘制_选择裁剪区域, 选择一个区域作为当前裁剪区域, 注意: 该函数只对GDI有效.
//
// hDraw: 图形绘制句柄.
//
// hRgn: 区域句柄.
//
// 返回: 返回值指定地区的复杂性，可以是下列值之一.
//
// NULLREGION Region is empty.
//
// SIMPLEREGION Region is a single rectangle.
//
// COMPLEXREGION Region is more than one rectangle.
//
// ERROR An error occurred. (The previous clipping region is unaffected).
func X绘制_选择裁剪区域(图形绘制句柄 int, 区域句柄 int) int {
	r, _, _ := xDraw_GDI_SelectClipRgn.Call(uintptr(图形绘制句柄), uintptr(区域句柄))
	return int(r)
}

// 绘制_填充矩形, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func X绘制_填充矩形(图形绘制句柄 int, 矩形区域 *RECT) int {
	r, _, _ := xDraw_FillRect.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形区域)))
	return int(r)
}

// 绘制_填充矩形F, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func X绘制_填充矩形F(图形绘制句柄 int, 矩形区域 *RECTF) int {
	r, _, _ := xDraw_FillRectF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形区域)))
	return int(r)
}

// 绘制_填充矩形指定颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
//
// color: ABGR 颜色.
func X绘制_填充矩形指定颜色(图形绘制句柄 int, 矩形区域 *RECT, ABGR颜色 int) int {
	r, _, _ := xDraw_FillRectColor.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形区域)), uintptr(ABGR颜色))
	return int(r)
}

// 绘制_填充矩形指定颜色F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
//
// color: ABGR 颜色.
func X绘制_填充矩形指定颜色F(图形绘制句柄 int, 矩形区域 *RECTF, ABGR颜色 int) int {
	r, _, _ := xDraw_FillRectColorF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形区域)), uintptr(ABGR颜色))
	return int(r)
}

// 绘制_填充区域, 通过使用指定的画刷填充一个区域.
//
// hDraw: 图形绘制句柄.
//
// hrgn: 区域句柄.
//
// hbr: 画刷句柄.
func X绘制_填充区域(图形绘制句柄 int, 区域句柄 int, 画刷句柄 int) bool {
	r, _, _ := xDraw_GDI_FillRgn.Call(uintptr(图形绘制句柄), uintptr(区域句柄), uintptr(画刷句柄))
	return r != 0
}

// 绘制_填充圆形.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func X绘制_填充圆形(图形绘制句柄 int, 矩形区域 *RECT) int {
	r, _, _ := xDraw_FillEllipse.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形区域)))
	return int(r)
}

// 绘制_填充圆形F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func X绘制_填充圆形F(图形绘制句柄 int, 矩形区域 *RECTF) int {
	r, _, _ := xDraw_FillEllipseF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形区域)))
	return int(r)
}

// 绘制_圆形, 绘制圆边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func X绘制_圆形(图形绘制句柄 int, 矩形区域 *RECT) int {
	r, _, _ := xDraw_DrawEllipse.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形区域)))
	return int(r)
}

// 绘制_填充圆角矩形.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func X绘制_填充圆角矩形(图形绘制句柄 int, 矩形坐标 *RECT, 圆角宽度, 圆角高度 int) int {
	r, _, _ := xDraw_FillRoundRect.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)), uintptr(圆角宽度), uintptr(圆角高度))
	return int(r)
}

// 绘制_填充圆角矩形F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func X绘制_填充圆角矩形F(图形绘制句柄 int, 矩形坐标 *RECTF, 圆角宽度, 圆角高度 float32) int {
	r, _, _ := xDraw_FillRoundRectF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)), 炫彩工具类.Float32Ptr(圆角宽度), 炫彩工具类.Float32Ptr(圆角高度))
	return int(r)
}

// 绘制_圆角矩形, 绘制圆角矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func X绘制_圆角矩形(图形绘制句柄 int, 矩形坐标 *RECT, 圆角宽度 int, 圆角高度 int) int {
	r, _, _ := xDraw_DrawRoundRect.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)), uintptr(圆角宽度), uintptr(圆角高度))
	return int(r)
}

// 绘制_圆角矩形F, 绘制圆角矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func X绘制_圆角矩形F(图形绘制句柄 int, 矩形坐标 *RECT, 圆角宽度, 圆角高度 float32) int {
	r, _, _ := xDraw_DrawRoundRectF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)), 炫彩工具类.Float32Ptr(圆角宽度), 炫彩工具类.Float32Ptr(圆角高度))
	return int(r)
}

// 绘制_填充圆角矩形扩展.
//
// hDraw: 图形绘制句柄.
//
// pRect: 坐标.
//
// nLeftTop: 圆角大小.
//
// nRightTop: 圆角大小.
//
// nRightBottom: 圆角大小.
//
// nLeftBottom: 圆角大小.
func X绘制_填充圆角矩形EX(图形绘制句柄 int, 坐标 *RECT, 左上角, 右上角, 右下角, 左下角 int) int {
	r, _, _ := xDraw_FillRoundRectEx.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(坐标)), uintptr(左上角), uintptr(右上角), uintptr(右下角), uintptr(左下角))
	return int(r)
}

// 绘制_填充圆角矩形扩展F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 坐标.
//
// nLeftTop: 圆角大小.
//
// nRightTop: 圆角大小.
//
// nRightBottom: 圆角大小.
//
// nLeftBottom: 圆角大小.
func X绘制_填充圆角矩形EXF(图形绘制句柄 int, 坐标 *RECTF, 左上角, 右上角, 右下角, 左下角 float32) int {
	r, _, _ := xDraw_FillRoundRectExF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.Float32Ptr(左上角), 炫彩工具类.Float32Ptr(右上角), 炫彩工具类.Float32Ptr(右下角), 炫彩工具类.Float32Ptr(左下角))
	return int(r)
}

// 绘制_圆角矩形扩展, 绘制圆角矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 坐标.
//
// nLeftTop: 圆角大小.
//
// nRightTop: 圆角大小.
//
// nRightBottom: 圆角大小.
//
// nLeftBottom: 圆角大小.
func X绘制_圆角矩形EX(图形绘制句柄 int, 坐标 *RECT, 左上角 int, 右上角 int, 右下角 int, 左下角 int) int {
	r, _, _ := xDraw_DrawRoundRectEx.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(坐标)), uintptr(左上角), uintptr(右上角), uintptr(右下角), uintptr(左下角))
	return int(r)
}

// 绘制_圆角矩形扩展F, 绘制圆角矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 坐标.
//
// nLeftTop: 圆角大小.
//
// nRightTop: 圆角大小.
//
// nRightBottom: 圆角大小.
//
// nLeftBottom: 圆角大小.
func X绘制_圆角矩形EXF(图形绘制句柄 int, 坐标 *RECT, 左上角, 右上角, 右下角, 左下角 float32) int {
	r, _, _ := xDraw_DrawRoundRectExF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.Float32Ptr(左上角), 炫彩工具类.Float32Ptr(右上角), 炫彩工具类.Float32Ptr(右下角), 炫彩工具类.Float32Ptr(左下角))
	return int(r)
}

// 绘制_矩形, 绘制矩形, 使用当前的画刷和画笔. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零.
//
// hDraw: 图形绘制句柄.
//
// nLeftRect: 左上角X坐标.
//
// nTopRect: 左上角Y坐标.
//
// nRightRect: 右下角X坐标.
//
// nBottomRect: 右下角Y坐标.
func X绘制_矩形(图形绘制句柄 int, 左上角X坐标 int, 左上角Y坐标 int, 右下角X坐标 int, 右下角Y坐标 int) bool {
	r, _, _ := xDraw_GDI_Rectangle.Call(uintptr(图形绘制句柄), uintptr(左上角X坐标), uintptr(左上角Y坐标), uintptr(右下角X坐标), uintptr(右下角Y坐标))
	return r != 0
}

// 绘制_渐变填充2, 渐变填充, 从一种颜色过渡到另一种颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func X绘制_渐变填充2(图形绘制句柄 int, 矩形坐标 *RECT, 开始颜色 int, 结束颜色 int, 模式 炫彩常量类.GRADIENT_FILL_) int {
	r, _, _ := xDraw_GradientFill2.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)), uintptr(开始颜色), uintptr(结束颜色), uintptr(模式))
	return int(r)
}

// 绘制_渐变填充2F, 渐变填充, 从一种颜色过渡到另一种颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func X绘制_渐变填充2F(图形绘制句柄 int, 矩形坐标 *RECTF, 开始颜色 int, 结束颜色 int, 模式 炫彩常量类.GRADIENT_FILL_) int {
	r, _, _ := xDraw_GradientFill2F.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)), uintptr(开始颜色), uintptr(结束颜色), uintptr(模式))
	return int(r)
}

// 绘制_渐变填充4, 渐变填充,从一种颜色过渡到另一种颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// color3: 开始颜色, ABGR 颜色.
//
// color4: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func X绘制_渐变填充4(图形绘制句柄 int, 矩形坐标 *RECT, 颜色1 int, 颜色2 int, 颜色3 int, 颜色4 int, 模式 炫彩常量类.GRADIENT_FILL_) bool {
	r, _, _ := xDraw_GradientFill4.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)), uintptr(颜色1), uintptr(颜色2), uintptr(颜色3), uintptr(颜色4), uintptr(模式))
	return r != 0
}

// 绘制_渐变填充4F, 渐变填充,从一种颜色过渡到另一种颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// color3: 开始颜色, ABGR 颜色.
//
// color4: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func X绘制_渐变填充4F(图形绘制句柄 int, 矩形坐标 *RECTF, 颜色1 int, 颜色2 int, 颜色3 int, 颜色4 int, 模式 炫彩常量类.GRADIENT_FILL_) bool {
	r, _, _ := xDraw_GradientFill4F.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)), uintptr(颜色1), uintptr(颜色2), uintptr(颜色3), uintptr(颜色4), uintptr(模式))
	return r != 0
}

// 绘制_边框区域, 绘制边框, 使用指定的画刷绘制指定的区域的边框. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零.
//
// hDraw: 图形绘制句柄.
//
// hrgn: 区域句柄.
//
// hbr: 画刷句柄.
//
// nWidth: 边框宽度, 垂直边.
//
// nHeight: 边框高度, 水平边.
func X绘制_边框区域(图形绘制句柄 int, 区域句柄 int, 画刷句柄 int, 边框宽度 int, 边框高度 int) bool {
	r, _, _ := xDraw_GDI_FrameRgn.Call(uintptr(图形绘制句柄), uintptr(区域句柄), uintptr(画刷句柄), uintptr(边框宽度), uintptr(边框高度))
	return r != 0
}

// 绘制_焦点矩形.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
func X绘制_焦点矩形(图形绘制句柄 int, 矩形坐标 *RECT) int {
	r, _, _ := xDraw_FocusRect.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)))
	return int(r)
}

// 绘制_焦点矩形F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
func X绘制_焦点矩形F(图形绘制句柄 int, 矩形坐标 *RECTF) int {
	r, _, _ := xDraw_FocusRectF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形坐标)))
	return int(r)
}

// 绘制_移动到起点, 更新当前位置到指定点，并返回以前的位置. 如果函数成功, 返回非零值.
//
// hDraw: 图形绘制句柄.
//
// X: 坐标.
//
// Y: 坐标.
//
// pPoint: 接收以前的当前位置到一个POINT结构的指针, 如果这个参数是NULL指针, 没有返回原来的位置.
func X绘制_移动到起点(图形绘制句柄 int, 坐标x int, 坐标y int, 接收以前的当前位置到一个POINT结构的指针 *POINT) bool {
	r, _, _ := xDraw_GDI_MoveToEx.Call(uintptr(图形绘制句柄), uintptr(坐标x), uintptr(坐标y), uintptr(unsafe.Pointer(接收以前的当前位置到一个POINT结构的指针)))
	return r != 0
}

// 绘制_线终点, 函数绘制一条线从当前位置到, 但不包括指定点. 如果函数成功, 返回非零值.
//
// hDraw: 图形绘制句柄.
//
// nXEnd: X坐标, 线结束点.
//
// nYEnd: Y坐标, 线结束点.
func X绘制_线终点(图形绘制句柄 int, X坐标 int, Y坐标 int) bool {
	r, _, _ := xDraw_GDI_LineTo.Call(uintptr(图形绘制句柄), uintptr(X坐标), uintptr(Y坐标))
	return r != 0
}

// 绘制_折线, Polyline() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// pArrayPt: 参见MSDN.
//
// arrayPtSize: 参见MSDN.
func X绘制_折线(图形绘制句柄 int, pArrayPt []POINT, arrayPtSize int) bool {
	r, _, _ := xDraw_GDI_Polyline.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(&pArrayPt[0])), uintptr(arrayPtSize))
	return r != 0
}

// 绘制_置像素颜色, 函数设置在指定的坐标到指定的颜色的像素. 如果函数成功返回RGB值, 如果失败返回-1.
//
// hDraw: 图形绘制句柄.
//
// X: 坐标.
//
// Y: 坐标.
//
// crColor: RGB颜色值.
func X绘制_置像素颜色(图形绘制句柄 int, 坐标x int, 坐标y int, RGB颜色值 int) int {
	r, _, _ := xDraw_GDI_SetPixel.Call(uintptr(图形绘制句柄), uintptr(坐标x), uintptr(坐标y), uintptr(RGB颜色值))
	return int(r)
}

// 绘制_取D2D渲染目标, 返回 *ID2D1RenderTarget.
//
// hDraw: 图形绘制句柄.
func X绘制_取D2D渲染目标(图形绘制句柄 int) int {
	r, _, _ := xDraw_GetD2dRenderTarget.Call(uintptr(图形绘制句柄))
	return int(r)
}

// 绘制_图标, 绘制图标, DrawIconEx()参见MSDN.
//
// hDraw: .
//
// xLeft: .
//
// yTop: .
//
// hIcon: .
//
// cxWidth: .
//
// cyWidth: .
//
// istepIfAniCur: .
//
// hbrFlickerFreeDraw: .
//
// diFlags: .
func X绘制_图标(hDraw int, xLeft int, yTop int, hIcon uintptr, cxWidth int, cyWidth int, istepIfAniCur int, hbrFlickerFreeDraw int, diFlags int) bool {
	r, _, _ := xDraw_GDI_DrawIconEx.Call(uintptr(hDraw), uintptr(xLeft), uintptr(yTop), hIcon, uintptr(cxWidth), uintptr(cyWidth), uintptr(istepIfAniCur), uintptr(hbrFlickerFreeDraw), uintptr(diFlags))
	return r != 0
}

// 绘制_复制, BitBlt() 参见MSDN.
//
// hDrawDest: XX.
//
// nXDest: XX.
//
// nYDest: XX.
//
// nWidth: XX.
//
// nHeight: XX.
//
// hdcSrc: XX.
//
// nXSrc: XX.
//
// nYSrc: XX.
//
// dwRop: XX.
func X绘制_复制(hDrawDest int, nXDest, nYDest, nWidth, nHeight int32, hdcSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {
	r, _, _ := xDraw_GDI_BitBlt.Call(uintptr(hDrawDest), uintptr(nXDest), uintptr(nYDest), uintptr(nWidth), uintptr(nHeight), hdcSrc, uintptr(nXSrc), uintptr(nYSrc), uintptr(dwRop))
	return r != 0
}

// 绘制_复制2, BitBlt() 参见MSDN.
//
// hDrawDest: XX.
//
// nXDest: XX.
//
// nYDest: XX.
//
// nWidth: XX.
//
// nHeight: XX.
//
// hDrawSrc: XX.
//
// nXSrc: XX.
//
// nYSrc: XX.
//
// dwRop: XX.
func X绘制_复制2(hDrawDest int, nXDest, nYDest, nWidth, nHeight int32, hDrawSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {
	r, _, _ := xDraw_GDI_BitBlt2.Call(uintptr(hDrawDest), uintptr(nXDest), uintptr(nYDest), uintptr(nWidth), uintptr(nHeight), hDrawSrc, uintptr(nXSrc), uintptr(nYSrc), uintptr(dwRop))
	return r != 0
}

// 绘制_带透明复制, AlphaBlend() 参见MSDN.
//
// hDraw: XX.
//
// nXOriginDest: XX.
//
// nYOriginDest: XX.
//
// nWidthDest: XX.
//
// nHeightDest: XX.
//
// hdcSrc: XX.
//
// nXOriginSrc: XX.
//
// nYOriginSrc: XX.
//
// nWidthSrc: XX.
//
// nHeightSrc: XX.
//
// alpha: XX.
func X绘制_带透明复制(hDraw int, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int32, hdcSrc uintptr, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha int32) bool {
	r, _, _ := xDraw_GDI_AlphaBlend.Call(uintptr(hDraw), uintptr(nXOriginDest), uintptr(nYOriginDest), uintptr(nWidthDest), uintptr(nHeightDest), hdcSrc, uintptr(nXOriginSrc), uintptr(nYOriginSrc), uintptr(nWidthSrc), uintptr(nHeightSrc), uintptr(alpha))
	return r != 0
}

// 绘制_GDI_椭圆.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func X绘制_GDI_椭圆(图形绘制句柄 int, 矩形区域 *RECT) bool {
	r, _, _ := xDraw_GDI_Ellipse.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(矩形区域)))
	return r != 0
}

// 绘制_填充多边形, 填充多边形.
//
// hDraw: 图形绘制句柄.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func X绘制_填充多边形(图形绘制句柄 int, 顶点坐标数组 []POINT, 顶点数量 int) int {
	r, _, _ := xDraw_FillPolygon.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(&顶点坐标数组[0])), uintptr(顶点数量))
	return int(r)
}

// 绘制_填充多边形F, 填充多边形.
//
// hDraw: 图形绘制句柄.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func X绘制_填充多边形F(图形绘制句柄 int, 顶点坐标数组 []POINTF, 顶点数量 int) int {
	r, _, _ := xDraw_FillPolygonF.Call(uintptr(图形绘制句柄), uintptr(unsafe.Pointer(&顶点坐标数组[0])), uintptr(顶点数量))
	return int(r)
}

// 绘制_图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
func X绘制_图片(图形绘制句柄 int, 图片句柄 int, x坐标, y坐标 int32) {
	xDraw_Image.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(x坐标), uintptr(y坐标))
}

// 绘制_图片F.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
func X绘制_图片F(图形绘制句柄 int, 图片句柄 int, x坐标, y坐标 float32) int {
	r, _, _ := xDraw_ImageF.Call(uintptr(图形绘制句柄), uintptr(图片句柄), 炫彩工具类.Float32Ptr(x坐标), 炫彩工具类.Float32Ptr(y坐标))
	return int(r)
}

// 绘制_图片自适应.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bOnlyBorder: 是否只绘制边缘区域.
func X绘制_图片自适应(图形绘制句柄 int, 图片句柄 int, 坐标 *RECT, 是否只绘制边缘区域 bool) int {
	r, _, _ := xDraw_ImageAdaptive.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.BoolPtr(是否只绘制边缘区域))
	return int(r)
}

// 绘制_图片自适应F.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bOnlyBorder: 是否只绘制边缘区域.
func X绘制_图片自适应F(图形绘制句柄 int, 图片句柄 int, 坐标 *RECTF, 是否只绘制边缘区域 bool) int {
	r, _, _ := xDraw_ImageAdaptiveF.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.BoolPtr(是否只绘制边缘区域))
	return int(r)
}

// 绘制_图片扩展, 绘制图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
func X绘制_图片EX(图形绘制句柄 int, 图片句柄 int, x坐标, y坐标, 宽度, 高度 int) int {
	r, _, _ := xDraw_ImageEx.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 绘制_图片扩展F, 绘制图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
func X绘制_图片EXF(图形绘制句柄 int, 图片句柄 int, x坐标, y坐标, 宽度, 高度 float32) int {
	r, _, _ := xDraw_ImageExF.Call(uintptr(图形绘制句柄), uintptr(图片句柄), 炫彩工具类.Float32Ptr(x坐标), 炫彩工具类.Float32Ptr(y坐标), 炫彩工具类.Float32Ptr(宽度), 炫彩工具类.Float32Ptr(高度))
	return int(r)
}

// 绘制_图片增强.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bClip: 是否裁剪区域.
func X绘制_图片增强(图形绘制句柄 int, 图片句柄 int, 坐标 *RECT, 是否裁剪区域 bool) int {
	r, _, _ := xDraw_ImageSuper.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.BoolPtr(是否裁剪区域))
	return int(r)
}

// 绘制_图片增强F.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bClip: 是否裁剪区域.
func X绘制_图片增强F(图形绘制句柄 int, 图片句柄 int, 坐标 *RECTF, 是否裁剪区域 bool) int {
	r, _, _ := xDraw_ImageSuperF.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.BoolPtr(是否裁剪区域))
	return int(r)
}

// 绘制_图片增强扩展.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// prcDest: 目标坐标.
//
// prcSrc: 源坐标.
func X绘制_图片增强EX(图形绘制句柄 int, 图片句柄 int, 目标坐标 *RECT, 源坐标 *RECT) int {
	r, _, _ := xDraw_ImageSuperEx.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(unsafe.Pointer(目标坐标)), uintptr(unsafe.Pointer(源坐标)))
	return int(r)
}

// 绘制_图片增强扩展F.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// prcDest: 目标坐标.
//
// prcSrc: 源坐标.
func X绘制_图片增强EXF(图形绘制句柄 int, 图片句柄 int, 目标坐标 *RECTF, 源坐标 *RECT) int {
	r, _, _ := xDraw_ImageSuperExF.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(unsafe.Pointer(目标坐标)), uintptr(unsafe.Pointer(源坐标)))
	return int(r)
}

// 绘制_图片增强遮盖, 绘制带遮盖的图片. D2D留空.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// hImageFrameMask: 图片句柄, 遮盖.
//
// pRect: 坐标.
//
// pRectMask: 坐标, 遮盖.
//
// bClip: 是否裁剪区域.
func X绘制_图片增强遮盖(图形绘制句柄 int, 图片句柄 int, 图片遮盖句柄 int, 遮盖坐标 *RECT, 坐标 *RECT, 是否裁剪区域 bool) int {
	r, _, _ := xDraw_ImageSuperMask.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(图片遮盖句柄), uintptr(unsafe.Pointer(坐标)), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.BoolPtr(是否裁剪区域))
	return int(r)
}

// 绘制_图片平铺, 绘制图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.
func X绘制_图片平铺(图形绘制句柄 int, 图片句柄 int, hImageFrameMask int, 坐标 *RECT, 标识 int) int {
	r, _, _ := xDraw_ImageTile.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(hImageFrameMask), uintptr(unsafe.Pointer(坐标)), uintptr(标识))
	return int(r)
}

// 绘制_图片平铺F, 绘制图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.
func X绘制_图片平铺F(图形绘制句柄 int, 图片句柄 int, hImageFrameMask int, 坐标 *RECTF, 标识 int) int {
	r, _, _ := xDraw_ImageTileF.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(hImageFrameMask), uintptr(unsafe.Pointer(坐标)), uintptr(标识))
	return int(r)
}

// 绘制_图片遮盖, 绘制带遮盖的图片, D2D留空.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// hImageFrameMask: 图片句柄, 遮盖.
//
// x: hImageFrame X坐标.
//
// y: hImageFrame Y坐标.
//
// x2: hImageFrameMask X坐标.
//
// y2: hImageFrameMask Y坐标.
func X绘制_图片遮盖(图形绘制句柄 int, 图片句柄 int, 图片遮盖句柄 int, x int, y int, X坐标 int, Y坐标 int) int {
	r, _, _ := xDraw_ImageMask.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(图片遮盖句柄), uintptr(图片句柄), uintptr(图片句柄), uintptr(图片遮盖句柄), uintptr(图片遮盖句柄))
	return int(r)
}

// 绘制_图片遮盖矩形, 使用矩形作为遮罩.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 矩形坐标.
//
// pRcMask: 遮罩坐标.
//
// pRcRoundAngle: 遮罩圆角.
func X绘制_图片遮盖矩形(图形绘制句柄 int, 图片句柄 int, 矩形坐标 *RECT, 遮罩坐标 *RECT, 遮罩圆角 *RECT) int {
	r, _, _ := xDraw_ImageMaskRect.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(unsafe.Pointer(矩形坐标)), uintptr(unsafe.Pointer(遮罩坐标)), uintptr(unsafe.Pointer(遮罩圆角)))
	return int(r)
}

// 绘制_图片遮盖圆型, 使用圆形作为遮罩.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 矩形坐标.
//
// pRcMask: 遮罩坐标.
func X绘制_图片遮盖圆型(图形绘制句柄 int, 图片句柄 int, 矩形坐标 *RECT, 遮罩坐标 *RECT) int {
	r, _, _ := xDraw_ImageMaskEllipse.Call(uintptr(图形绘制句柄), uintptr(图片句柄), uintptr(unsafe.Pointer(矩形坐标)), uintptr(unsafe.Pointer(遮罩坐标)))
	return int(r)
}

// 绘制_文本指定矩形, DrawText() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// lpString: 字符串.
//
// lpRect: 坐标.
func X绘制_文本指定矩形(图形绘制句柄 int, 字符串 string, 坐标 *RECT) int {
	r, _, _ := xDraw_DrawText.Call(uintptr(图形绘制句柄), 炫彩工具类.StrPtr(字符串), uintptr(len([]rune(字符串))), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 绘制_文本指定矩形F, DrawText() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// lpString: 字符串.
//
// lpRect: 坐标.
func X绘制_文本指定矩形F(图形绘制句柄 int, 字符串 string, 坐标 *RECTF) int {
	r, _, _ := xDraw_DrawTextF.Call(uintptr(图形绘制句柄), 炫彩工具类.StrPtr(字符串), uintptr(len([]rune(字符串))), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 绘制_文本下划线.
//
// hDraw: 图形绘制句柄.
//
// lpString: 字符串.
//
// lpRect: 坐标.
//
// colorLine: 下划线颜色, ABGR 颜色.
func X绘制_文本下划线(图形绘制句柄 int, 字符串 string, 坐标 *RECT, 下划线颜色 int) int {
	r, _, _ := xDraw_DrawTextUnderline.Call(uintptr(图形绘制句柄), 炫彩工具类.StrPtr(字符串), uintptr(len([]rune(字符串))), uintptr(unsafe.Pointer(坐标)), uintptr(下划线颜色))
	return int(r)
}

// 绘制_文本下划线F.
//
// hDraw: 图形绘制句柄.
//
// lpString: 字符串.
//
// lpRect: 坐标.
//
// colorLine: 下划线颜色, ABGR 颜色.
func X绘制_文本下划线F(图形绘制句柄 int, 字符串 string, 坐标 *RECTF, 下划线颜色 int) int {
	r, _, _ := xDraw_DrawTextUnderlineF.Call(uintptr(图形绘制句柄), 炫彩工具类.StrPtr(字符串), uintptr(len([]rune(字符串))), uintptr(unsafe.Pointer(坐标)), uintptr(下划线颜色))
	return int(r)
}

// 绘制_文本, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
//
// cbString: XX.
func X绘制_文本(图形绘制句柄 int, nXStart int, nYStart int, lpString string, cbString string) int {
	r, _, _ := xDraw_TextOut.Call(uintptr(图形绘制句柄), uintptr(nXStart), uintptr(nYStart), 炫彩工具类.StrPtr(lpString), 炫彩工具类.StrPtr(cbString))
	return int(r)
}

// 绘制_文本F, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
//
// cbString: XX.
func X绘制_文本F(图形绘制句柄 int, nXStart, nYStart float32, lpString string, cbString string) int {
	r, _, _ := xDraw_TextOutF.Call(uintptr(图形绘制句柄), 炫彩工具类.Float32Ptr(nXStart), 炫彩工具类.Float32Ptr(nYStart), 炫彩工具类.StrPtr(lpString), 炫彩工具类.StrPtr(cbString))
	return int(r)
}

// 绘制_文本扩展, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func X绘制_文本EX(图形绘制句柄 int, nXStart int, nYStart int, lpString string) int {
	r, _, _ := xDraw_TextOutEx.Call(uintptr(图形绘制句柄), uintptr(nXStart), uintptr(nYStart), 炫彩工具类.StrPtr(lpString))
	return int(r)
}

// 绘制_文本扩展F, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func X绘制_文本EXF(图形绘制句柄 int, nXStart, nYStart float32, lpString string) int {
	r, _, _ := xDraw_TextOutExF.Call(uintptr(图形绘制句柄), 炫彩工具类.Float32Ptr(nXStart), 炫彩工具类.Float32Ptr(nYStart), 炫彩工具类.StrPtr(lpString))
	return int(r)
}

// 绘制_文本A, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func X绘制_文本A(图形绘制句柄 int, nXStart int, nYStart int, lpString string) int {
	r, _, _ := xDraw_TextOutA.Call(uintptr(图形绘制句柄), uintptr(nXStart), uintptr(nYStart), 炫彩工具类.StrPtr(lpString))
	return int(r)
}

// 绘制_文本AF, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func X绘制_文本AF(图形绘制句柄 int, nXStart, nYStart float32, lpString string) int {
	r, _, _ := xDraw_TextOutAF.Call(uintptr(图形绘制句柄), 炫彩工具类.Float32Ptr(nXStart), 炫彩工具类.Float32Ptr(nYStart), 炫彩工具类.StrPtr(lpString))
	return int(r)
}

// 绘制_设置文本渲染提示.
//
// hDraw: 图形绘制句柄.
//
// nType: XX.
func X绘制_设置文本渲染提示(图形绘制句柄 int, nType int) int {
	r, _, _ := xDraw_SetTextRenderingHint.Call(uintptr(图形绘制句柄), uintptr(nType))
	return int(r)
}

// 绘制_SVG源.
//
// hDraw: 图形绘制句柄.
//
// hSvg: SVG句柄.
func X绘制_SVG源(图形绘制句柄 int, SVG句柄 int) int {
	r, _, _ := xDraw_DrawSvgSrc.Call(uintptr(图形绘制句柄), uintptr(SVG句柄))
	return int(r)
}

// 绘制_SVG.
//
// hDraw: 图形绘制句柄.
//
// hSvg: SVG句柄.
//
// x: x坐标.
//
// y: y坐标.
func X绘制_SVG(图形绘制句柄 int, SVG句柄 int, x坐标 int, y坐标 int) int {
	r, _, _ := xDraw_DrawSvg.Call(uintptr(图形绘制句柄), uintptr(SVG句柄), uintptr(x坐标), uintptr(y坐标))
	return int(r)
}

// 绘制_SVG扩展.
//
// hDraw: 图形绘制句柄.
//
// hSvg: SVG句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func X绘制_SVGEX(图形绘制句柄 int, SVG句柄 int, x坐标 int, y坐标 int, 宽度 int, 高度 int) int {
	r, _, _ := xDraw_DrawSvgEx.Call(uintptr(图形绘制句柄), uintptr(SVG句柄), uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 绘制_SVG大小.
//
// hDraw: 图形绘制句柄.
//
// hSvg: SVG句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func X绘制_SVG大小(图形绘制句柄 int, SVG句柄 int, 宽度 int, 高度 int) int {
	r, _, _ := xDraw_DrawSvgSize.Call(uintptr(图形绘制句柄), uintptr(SVG句柄), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 绘制_D2D_清理, 使用指定颜色清理画布.
//
// hDraw: 图形绘制句柄.
//
// color: ABGR 颜色值.
func X绘制_D2D_清理(图形绘制句柄 int, ABGR颜色值 int) int {
	r, _, _ := xDraw_D2D_Clear.Call(uintptr(图形绘制句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 绘制_取字体, 返回字体句柄.
//
// hDraw: 图形绘制句柄.
func X绘制_取字体(图形绘制句柄 int) int {
	r, _, _ := xDraw_GetFont.Call(uintptr(图形绘制句柄))
	return int(r)
}
