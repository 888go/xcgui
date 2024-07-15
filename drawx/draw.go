package 炫彩绘制类 //bm:炫彩绘制类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Draw 图形绘制.
type Draw struct {
	炫彩对象基类.ObjectBase
}

// New 绘制_创建, 创建图形绘制模块实例, 返回句柄.
//
// hWindow: 窗口句柄.

// ff:创建
// hWindow:窗口句柄
func X创建(窗口句柄 int) *Draw {
	p := &Draw{}
	p.X设置句柄(炫彩基类.X绘制_创建(窗口句柄))
	return p
}

// NewGDI 绘制_创建GDI, 创建图形绘制模块实例, 返回图形绘制模块实例句柄.
//
// hWindow: 窗口句柄.
//
// hdc: hdc句柄.

// ff:创建GDI
// hdc:hdc句柄
// hWindow:窗口句柄
func X创建GDI(窗口句柄 int, hdc句柄 uintptr) *Draw {
	p := &Draw{}
	p.X设置句柄(炫彩基类.X绘制_创建GDI(窗口句柄, hdc句柄))
	return p
}

// NewByHandle 从图形绘制模块实例句柄创建对象.

// ff:创建并按图形绘制模块句柄
// handle:
func X创建并按图形绘制模块句柄(handle int) *Draw {
	p := &Draw{}
	p.X设置句柄(handle)
	return p
}

// 绘制_销毁, 销毁图形绘制模块实例句柄.

// ff:销毁
func (d *Draw) X销毁() int {
	return 炫彩基类.X绘制_销毁(d.Handle)
}

// 绘制_虚线, 绘制水平或垂直虚线.
//
// x1: 起点x坐标.
//
// y1: 起点y坐标.
//
// x2: 结束点x坐标.
//
// y2: 结束点y坐标.

// ff:虚线
// y2:结束点y坐标
// x2:结束点x坐标
// y1:起点y坐标
// x1:起点x坐标
func (d *Draw) X虚线(起点x坐标 int, 起点y坐标 int, 结束点x坐标 int, 结束点y坐标 int) int {
	return 炫彩基类.X绘制_虚线(d.Handle, 起点x坐标, 起点y坐标, 结束点x坐标, 结束点y坐标)
}

// 绘制_虚线F, 绘制水平或垂直虚线.
//
// x1: 起点x坐标.
//
// y1: 起点y坐标.
//
// x2: 结束点x坐标.
//
// y2: 结束点y坐标.

// ff:虚线F
// y2:结束点y坐标
// x2:结束点x坐标
// y1:起点y坐标
// x1:起点x坐标
func (d *Draw) X虚线F(起点x坐标, 起点y坐标, 结束点x坐标, 结束点y坐标 float32) int {
	return 炫彩基类.X绘制_虚线F(d.Handle, 起点x坐标, 起点y坐标, 结束点x坐标, 结束点y坐标)
}

// 绘制_圆弧.
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

// ff:圆弧
// sweepAngle:绘制角度
// startAngle:起始角度
// nHeight:高度
// nWidth:宽度
// y:y坐标
// x:x坐标
func (d *Draw) X圆弧(x坐标, y坐标 int, 宽度 int, 高度 int, 起始角度 float32, 绘制角度 float32) int {
	return 炫彩基类.X绘制_圆弧(d.Handle, x坐标, y坐标, 宽度, 高度, 起始角度, 绘制角度)
}

// 绘制_圆弧F.
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

// ff:圆弧F
// sweepAngle:绘制角度
// startAngle:起始角度
// nHeight:高度
// nWidth:宽度
// y:y坐标
// x:x坐标
func (d *Draw) X圆弧F(x坐标, y坐标, 宽度, 高度, 起始角度, 绘制角度 float32) int {
	return 炫彩基类.X绘制_圆弧F(d.Handle, x坐标, y坐标, 宽度, 高度, 起始角度, 绘制角度)
}

// 绘制_曲线, D2D暂时留空.
//
// points: 坐标点数组.
//
// count: 数组大小.
//
// tension: 大于或等于0.0F的值，指定曲线的张力, D2D 忽略此参数。.

// ff:曲线
// tension:
// count:
// points:坐标点切片
func (d *Draw) X曲线(坐标点切片 []炫彩基类.POINT, count int, tension float32) int {
	return 炫彩基类.X绘制_曲线(d.Handle, 坐标点切片, count, tension)
}

// 绘制_曲线F, D2D暂时留空.
//
// points: 坐标点数组.
//
// count: 数组大小.
//
// tension: 大于或等于0.0F的值，指定曲线的张力, D2D 忽略此参数。.

// ff:曲线F
// tension:
// count:
// points:坐标点切片
func (d *Draw) X曲线F(坐标点切片 []炫彩基类.POINTF, count int, tension float32) int {
	return 炫彩基类.X绘制_曲线F(d.Handle, 坐标点切片, count, tension)
}

// 绘制_线条.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.

// ff:线条
// y2:y2坐标
// x2:x2坐标
// y1:y1坐标
// x1:x1坐标
func (d *Draw) X线条(x1坐标 int, y1坐标 int, x2坐标 int, y2坐标 int) int {
	return 炫彩基类.X绘制_线条(d.Handle, x1坐标, y1坐标, x2坐标, y2坐标)
}

// 绘制_线条F.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.

// ff:线条F
// y2:y2坐标
// x2:x2坐标
// y1:y1坐标
// x1:x1坐标
func (d *Draw) X线条F(x1坐标, y1坐标, x2坐标, y2坐标 float32) int {
	return 炫彩基类.X绘制_线条F(d.Handle, x1坐标, y1坐标, x2坐标, y2坐标)
}

// 绘制_多边形, 绘制多边形.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.

// ff:多边形
// nCount:
// points:顶点坐标切片
func (d *Draw) X多边形(顶点坐标切片 []炫彩基类.POINT, nCount int) int {
	return 炫彩基类.X绘制_多边形(d.Handle, 顶点坐标切片, nCount)
}

// 绘制_多边形F, 绘制多边形.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.

// ff:多边形F
// nCount:
// points:顶点坐标切片
func (d *Draw) X多边形F(顶点坐标切片 []炫彩基类.POINTF, nCount int) int {
	return 炫彩基类.X绘制_多边形F(d.Handle, 顶点坐标切片, nCount)
}

// 绘制_矩形, 绘制矩形边框.
//
// pRect: 矩形坐标 .

// ff:矩形边框
// pRect:矩形坐标
func (d *Draw) X矩形边框(矩形坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_矩形边框(d.Handle, 矩形坐标)
}

// 绘制_矩形F, 绘制矩形边框.
//
// pRect: 矩形坐标 .

// ff:矩形边框F
// pRect:矩形坐标
func (d *Draw) X矩形边框F(矩形坐标 *炫彩基类.RECTF) int {
	return 炫彩基类.X绘制_矩形边框F(d.Handle, 矩形坐标)
}

// 绘制_置偏移, 设置坐标偏移量, X向左偏移为负数, 向右偏移为正数.
//
// x: X轴偏移量.
//
// y: Y轴偏移量.

// ff:置偏移
// y:Y轴偏移量
// x:轴偏移量
func (d *Draw) X置偏移(轴偏移量, Y轴偏移量 int32) int {
	return 炫彩基类.X绘制_置偏移(d.Handle, 轴偏移量, Y轴偏移量)
}

// 绘制_取偏移, 获取坐标偏移量, X向左偏移为负数, 向右偏移为正数.
//
// pX: 接收X轴偏移量.
//
// pY: 接收Y轴偏移量.

// ff:取偏移
// pY:接收Y轴偏移量
// pX:接收X轴偏移量
func (d *Draw) X取偏移(接收X轴偏移量, 接收Y轴偏移量 *int32) int {
	return 炫彩基类.X绘制_取偏移(d.Handle, 接收X轴偏移量, 接收Y轴偏移量)
}

// 绘制_还原状态, 还原状态, 释放用户绑定的GDI对象, 例如画刷, 画笔.

// ff:还原状态
func (d *Draw) X还原状态() int {
	return 炫彩基类.X绘制_还原状态(d.Handle)
}

// 绘制_取HDC, 获取绑定的设备上下文HDC, 返回HDC句柄.

// ff:取HDC
func (d *Draw) X取HDC() uintptr {
	return 炫彩基类.X绘制_取HDC(d.Handle)
}

// 绘制_置画刷颜色, 设置画刷颜色.
//
// color: ABGR 颜色值.

// ff:置画刷颜色
// color:ABGR
func (d *Draw) X置画刷颜色(ABGR int) int {
	return 炫彩基类.X绘制_置画刷颜色(d.Handle, ABGR)
}

// 绘制_置文本垂直, 设置文本垂直显示.
//
// bVertical: 是否垂直显示文本.

// ff:置文本垂直
// bVertical:是否垂直显示文本
func (d *Draw) X置文本垂直(是否垂直显示文本 bool) int {
	return 炫彩基类.X绘制_置文本垂直(d.Handle, 是否垂直显示文本)
}

// 绘制_置文本对齐, 设置文本对齐.
//
// nFlags: 对齐标识, TextFormatFlag_, TextAlignFlag_, TextTrimming_.

// ff:置文本对齐
// nFlags:对齐标识
func (d *Draw) X置文本对齐(对齐标识 炫彩常量类.TextFormatFlag_) int {
	return 炫彩基类.X绘制_置文本对齐(d.Handle, 对齐标识)
}

// 绘制_置字体.
//
// hFontx: 炫彩字体.

// ff:置字体
// hFontx:炫彩字体
func (d *Draw) X置字体(炫彩字体 int) int {
	return 炫彩基类.X绘制_置字体(d.Handle, 炫彩字体)
}

// 绘制_置线宽.
//
// nWidth: 宽度.

// ff:置线宽
// nWidth:宽度
func (d *Draw) X置线宽(宽度 int) int {
	return 炫彩基类.X绘制_置线宽(d.Handle, 宽度)
}

// 绘制_置线宽F.
//
// nWidth: 宽度.

// ff:置线宽F
// nWidth:宽度
func (d *Draw) X置线宽F(宽度 float32) int {
	return 炫彩基类.X绘制_置线宽F(d.Handle, 宽度)
}

// 绘制_置背景模式, SetBkMode() 参见MSDN.
//
// bTransparent: 参见MSDN.

// ff:置背景模式
// bTransparent:参见MSDN
func (d *Draw) X置背景模式(参见MSDN bool) int {
	return 炫彩基类.X绘制_置背景模式(d.Handle, 参见MSDN)
}

// 绘制_置裁剪区域, 设置裁剪区域.
//
// pRect: 区域坐标.

// ff:置裁剪区域
// pRect:区域坐标
func (d *Draw) X置裁剪区域(区域坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_置裁剪区域(d.Handle, 区域坐标)
}

// 绘制_置D2D文本渲染模式.
//
// mode	渲染模式 XC_DWRITE_RENDERING_MODE_.

// ff:置D2D文本渲染模式
// mode:渲染模式
func (d *Draw) X置D2D文本渲染模式(渲染模式 炫彩常量类.XC_DWRITE_RENDERING_MODE_) int {
	return 炫彩基类.X绘制_置D2D文本渲染模式(d.Handle, 渲染模式)
}

// 绘制_清除裁剪区域.

// ff:清除裁剪区域
func (d *Draw) X清除裁剪区域() int {
	return 炫彩基类.X绘制_清除裁剪区域(d.Handle)
}

// 绘制_启用平滑模式.
//
// bEnable: 是否启用.

// ff:启用平滑模式
// bEnable:是否启用
func (d *Draw) X启用平滑模式(是否启用 bool) int {
	return 炫彩基类.X绘制_启用平滑模式(d.Handle, 是否启用)
}

// 绘制_启用窗口透明判断, 当启用之后, 调用GDI+函数时, 如果参数alpha=255, 将自动修改为254, 应对GDI+的bug, 否则透明通道异常.
//
// bTransparent: 是否启用.

// ff:启用窗口透明判断
// bTransparent:是否启用
func (d *Draw) X启用窗口透明判断(是否启用 bool) int {
	return 炫彩基类.X绘制_启用窗口透明判断(d.Handle, 是否启用)
}

// 绘制_创建实心画刷, GDI创建具有指定的纯色逻辑刷.
//
// crColor: 画刷颜色.

// ff:创建实心画刷
// crColor:画刷颜色
func (d *Draw) X创建实心画刷(画刷颜色 int) int {
	return 炫彩基类.X绘制_创建实心画刷(d.Handle, 画刷颜色)
}

// 绘制_创建画笔, GDI创建一个逻辑笔, 指定的样式, 宽度和颜色, 随后的笔可以选择到设备上下文, 用于绘制线条和曲线.
//
// fnPenStyle: 画笔样式, PS_SOLID:实线, PS_DASH:段线, PS_DOT:点线, PS_DASHDOT:段线_点线, PS_DASHDOTDOT:段线_点_点, PS_NULL:空, PS_INSIDEFRAME:实线_笔宽是向里扩展.
//
// nWidth: 画笔宽度.
//
// crColor: ABGR 颜色.

// ff:创建画笔
// crColor:ABGR颜色
// nWidth:画笔宽度
// fnPenStyle:画笔样式
func (d *Draw) X创建画笔(画笔样式 int, 画笔宽度 int, ABGR颜色 int) int {
	return 炫彩基类.X绘制_创建画笔(d.Handle, 画笔样式, 画笔宽度, ABGR颜色)
}

// 绘制_创建矩形区域, GDI创建矩形区域, 成功返回区域句柄, 失败返回NULL.
//
// nLeftRect: 左上角X坐标.
//
// nTopRect: 左上角Y坐标.
//
// nRightRect: 右下角X坐标.
//
// nBottomRect: 右下角Y坐标.

// ff:创建矩形区域
// nBottomRect:右下角Y坐标
// nRightRect:右下角X坐标
// nTopRect:左上角Y坐标
// nLeftRect:左上角X坐标
func (d *Draw) X创建矩形区域(左上角X坐标 int, 左上角Y坐标 int, 右下角X坐标 int, 右下角Y坐标 int) int {
	return 炫彩基类.X绘制_创建矩形区域(d.Handle, 左上角X坐标, 左上角Y坐标, 右下角X坐标, 右下角Y坐标)
}

// 绘制_创建圆角矩形区域, GDI创建一个圆角的矩形区域, 成功返回区域句柄, 失败返回NULL.
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

// ff:创建圆角矩形区域
// nHeightEllipse:椭圆的高度
// nWidthEllipse:椭圆的宽度
// nBottomRect:Y坐标右下角
// nRightRect:坐标右下角
// nTopRect:Y坐标左上角
// nLeftRect:坐标左上角
func (d *Draw) X创建圆角矩形区域(坐标左上角 int, Y坐标左上角 int, 坐标右下角 int, Y坐标右下角 int, 椭圆的宽度 int, 椭圆的高度 int) int {
	return 炫彩基类.X绘制_创建圆角矩形区域(d.Handle, 坐标左上角, Y坐标左上角, 坐标右下角, Y坐标右下角, 椭圆的宽度, 椭圆的高度)
}

// 绘制_创建多边形区域, GDI创建一个多边形区域, 成功返回区域句柄, 失败返回NULL.
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

// ff:创建多边形区域
// fnPolyFillMode:
// cPoints:
// pPt:POINT切片
func (d *Draw) X创建多边形区域(POINT切片 []炫彩基类.POINT, cPoints, fnPolyFillMode int) int {
	return 炫彩基类.X绘制_创建多边形区域(d.Handle, POINT切片, cPoints, fnPolyFillMode)
}

// 绘制_GDI_椭圆.
//
// pRect: 矩形区域.

// ff:GDI椭圆
// pRect:矩形区域
func (d *Draw) GDI椭圆(矩形区域 *炫彩基类.RECT) bool {
	return 炫彩基类.X绘制_GDI_椭圆(d.Handle, 矩形区域)
}

// 绘制_选择裁剪区域, 选择一个区域作为当前裁剪区域, 注意: 该函数只对GDI有效.
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

// ff:选择裁剪区域
// hRgn:区域句柄
func (d *Draw) X选择裁剪区域(区域句柄 int) int {
	return 炫彩基类.X绘制_选择裁剪区域(d.Handle, 区域句柄)
}

// 绘制_填充矩形, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界.
//
// pRect: 矩形区域.

// ff:填充矩形
// pRect:矩形区域
func (d *Draw) X填充矩形(矩形区域 *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_填充矩形(d.Handle, 矩形区域)
}

// 绘制_填充矩形F, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界.
//
// pRect: 矩形区域.

// ff:填充矩形F
// pRect:矩形区域
func (d *Draw) X填充矩形F(矩形区域 *炫彩基类.RECTF) int {
	return 炫彩基类.X绘制_填充矩形F(d.Handle, 矩形区域)
}

// 绘制_填充矩形指定颜色.
//
// pRect: 矩形区域.
//
// color: ABGR 颜色.

// ff:填充矩形指定颜色
// color:
// pRect:矩形区域
func (d *Draw) X填充矩形指定颜色(矩形区域 *炫彩基类.RECT, color int) int {
	return 炫彩基类.X绘制_填充矩形指定颜色(d.Handle, 矩形区域, color)
}

// 绘制_填充矩形指定颜色F.
//
// pRect: 矩形区域.
//
// color: ABGR 颜色.

// ff:填充矩形指定颜色F
// color:
// pRect:矩形区域
func (d *Draw) X填充矩形指定颜色F(矩形区域 *炫彩基类.RECTF, color int) int {
	return 炫彩基类.X绘制_填充矩形指定颜色F(d.Handle, 矩形区域, color)
}

// 绘制_填充区域, 通过使用指定的画刷填充一个区域.
//
// hrgn: 区域句柄.
//
// hbr: 画刷句柄.

// ff:填充区域
// hbr:画刷句柄
// hrgn:区域句柄
func (d *Draw) X填充区域(区域句柄 int, 画刷句柄 int) bool {
	return 炫彩基类.X绘制_填充区域(d.Handle, 区域句柄, 画刷句柄)
}

// 绘制_填充圆形.
//
// pRect: 矩形区域.

// ff:填充圆形
// pRect:矩形区域
func (d *Draw) X填充圆形(矩形区域 *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_填充圆形(d.Handle, 矩形区域)
}

// 绘制_填充圆形F.
//
// pRect: 矩形区域.

// ff:填充圆形F
// pRect:矩形区域
func (d *Draw) X填充圆形F(矩形区域 *炫彩基类.RECTF) int {
	return 炫彩基类.X绘制_填充圆形F(d.Handle, 矩形区域)
}

// 绘制_圆形, 绘制圆边框.
//
// pRect: 矩形区域.

// ff:圆形
// pRect:矩形区域
func (d *Draw) X圆形(矩形区域 *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_圆形(d.Handle, 矩形区域)
}

// 绘制_填充圆角矩形.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.

// ff:填充圆角矩形
// nHeight:
// nWidth:
// pRect:矩形坐标
func (d *Draw) X填充圆角矩形(矩形坐标 *炫彩基类.RECT, nWidth, nHeight int) int {
	return 炫彩基类.X绘制_填充圆角矩形(d.Handle, 矩形坐标, nWidth, nHeight)
}

// 绘制_填充圆角矩形F.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.

// ff:填充圆角矩形F
// nHeight:
// nWidth:
// pRect:矩形坐标
func (d *Draw) X填充圆角矩形F(矩形坐标 *炫彩基类.RECTF, nWidth, nHeight float32) int {
	return 炫彩基类.X绘制_填充圆角矩形F(d.Handle, 矩形坐标, nWidth, nHeight)
}

// 绘制_圆角矩形, 绘制圆角矩形边框.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.

// ff:圆角矩形
// nHeight:
// nWidth:
// pRect:矩形坐标
func (d *Draw) X圆角矩形(矩形坐标 *炫彩基类.RECT, nWidth int, nHeight int) int {
	return 炫彩基类.X绘制_圆角矩形(d.Handle, 矩形坐标, nWidth, nHeight)
}

// 绘制_圆角矩形F, 绘制圆角矩形边框.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.

// ff:圆角矩形F
// nHeight:
// nWidth:
// pRect:矩形坐标
func (d *Draw) X圆角矩形F(矩形坐标 *炫彩基类.RECT, nWidth, nHeight float32) int {
	return 炫彩基类.X绘制_圆角矩形F(d.Handle, 矩形坐标, nWidth, nHeight)
}

// 绘制_填充圆角矩形扩展.
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

// ff:填充圆角矩形EX
// nLeftBottom:
// nRightBottom:
// nRightTop:
// nLeftTop:
// pRect:坐标
func (d *Draw) X填充圆角矩形EX(坐标 *炫彩基类.RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom int) int {
	return 炫彩基类.X绘制_填充圆角矩形EX(d.Handle, 坐标, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_填充圆角矩形扩展F.
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

// ff:填充圆角矩形EXF
// nLeftBottom:
// nRightBottom:
// nRightTop:
// nLeftTop:
// pRect:坐标
func (d *Draw) X填充圆角矩形EXF(坐标 *炫彩基类.RECTF, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {
	return 炫彩基类.X绘制_填充圆角矩形EXF(d.Handle, 坐标, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_圆角矩形扩展, 绘制圆角矩形边框.
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

// ff:圆角矩形EX
// nLeftBottom:
// nRightBottom:
// nRightTop:
// nLeftTop:
// pRect:坐标
func (d *Draw) X圆角矩形EX(坐标 *炫彩基类.RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {
	return 炫彩基类.X绘制_圆角矩形EX(d.Handle, 坐标, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_圆角矩形扩展F, 绘制圆角矩形边框.
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

// ff:圆角矩形EXF
// nLeftBottom:
// nRightBottom:
// nRightTop:
// nLeftTop:
// pRect:坐标
func (d *Draw) X圆角矩形EXF(坐标 *炫彩基类.RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {
	return 炫彩基类.X绘制_圆角矩形EXF(d.Handle, 坐标, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_矩形, 绘制矩形, 使用当前的画刷和画笔. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零.
//
// nLeftRect: 左上角X坐标.
//
// nTopRect: 左上角Y坐标.
//
// nRightRect: 右下角X坐标.
//
// nBottomRect: 右下角Y坐标.

// ff:矩形
// nBottomRect:右下角Y坐标
// nRightRect:右下角X坐标
// nTopRect:左上角Y坐标
// nLeftRect:左上角X坐标
func (d *Draw) X矩形(左上角X坐标 int, 左上角Y坐标 int, 右下角X坐标 int, 右下角Y坐标 int) bool {
	return 炫彩基类.X绘制_矩形(d.Handle, 左上角X坐标, 左上角Y坐标, 右下角X坐标, 右下角Y坐标)
}

// 绘制_渐变填充2, 渐变填充, 从一种颜色过渡到另一种颜色.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.

// ff:渐变填充2
// mode:
// color2:
// color1:
// pRect:矩形坐标
func (d *Draw) X渐变填充2(矩形坐标 *炫彩基类.RECT, color1 int, color2 int, mode 炫彩常量类.GRADIENT_FILL_) int {
	return 炫彩基类.X绘制_渐变填充2(d.Handle, 矩形坐标, color1, color2, mode)
}

// 绘制_渐变填充2F, 渐变填充, 从一种颜色过渡到另一种颜色.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.

// ff:渐变填充2F
// mode:
// color2:
// color1:
// pRect:矩形坐标
func (d *Draw) X渐变填充2F(矩形坐标 *炫彩基类.RECTF, color1 int, color2 int, mode 炫彩常量类.GRADIENT_FILL_) int {
	return 炫彩基类.X绘制_渐变填充2F(d.Handle, 矩形坐标, color1, color2, mode)
}

// 绘制_渐变填充4, 渐变填充,从一种颜色过渡到另一种颜色.
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

// ff:渐变填充4
// mode:
// color4:
// color3:
// color2:
// color1:
// pRect:矩形坐标
func (d *Draw) X渐变填充4(矩形坐标 *炫彩基类.RECT, color1 int, color2 int, color3 int, color4 int, mode 炫彩常量类.GRADIENT_FILL_) bool {
	return 炫彩基类.X绘制_渐变填充4(d.Handle, 矩形坐标, color1, color2, color3, color4, mode)
}

// 绘制_渐变填充4F, 渐变填充,从一种颜色过渡到另一种颜色.
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

// ff:渐变填充4F
// mode:
// color4:
// color3:
// color2:
// color1:
// pRect:矩形坐标
func (d *Draw) X渐变填充4F(矩形坐标 *炫彩基类.RECTF, color1 int, color2 int, color3 int, color4 int, mode 炫彩常量类.GRADIENT_FILL_) bool {
	return 炫彩基类.X绘制_渐变填充4F(d.Handle, 矩形坐标, color1, color2, color3, color4, mode)
}

// 绘制_边框区域, 绘制边框, 使用指定的画刷绘制指定的区域的边框. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零.
//
// hrgn: 区域句柄.
//
// hbr: 画刷句柄.
//
// nWidth: 边框宽度, 垂直边.
//
// nHeight: 边框高度, 水平边.

// ff:边框区域
// nHeight:边框高度
// nWidth:边框宽度
// hbr:画刷句柄
// hrgn:区域句柄
func (d *Draw) X边框区域(区域句柄 int, 画刷句柄 int, 边框宽度 int, 边框高度 int) bool {
	return 炫彩基类.X绘制_边框区域(d.Handle, 区域句柄, 画刷句柄, 边框宽度, 边框高度)
}

// 绘制_焦点矩形.
//
// pRect: 矩形坐标.

// ff:焦点矩形
// pRect:矩形坐标
func (d *Draw) X焦点矩形(矩形坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_焦点矩形(d.Handle, 矩形坐标)
}

// 绘制_焦点矩形F.
//
// pRect: 矩形坐标.

// ff:焦点矩形F
// pRect:矩形坐标
func (d *Draw) X焦点矩形F(矩形坐标 *炫彩基类.RECTF) int {
	return 炫彩基类.X绘制_焦点矩形F(d.Handle, 矩形坐标)
}

// 绘制_移动到起点, 更新当前位置到指定点，并返回以前的位置. 如果函数成功, 返回非零值.
//
// X: 坐标.
//
// Y: 坐标.
//
// pPoint: 接收以前的当前位置到一个POINT结构的指针, 如果这个参数是NULL指针, 没有返回原来的位置.

// ff:移动到起点
// pPoint:接收原位置指针
// Y:坐标y
// X:坐标x
func (d *Draw) X移动到起点(坐标x int, 坐标y int, 接收原位置指针 *炫彩基类.POINT) bool {
	return 炫彩基类.X绘制_移动到起点(d.Handle, 坐标x, 坐标y, 接收原位置指针)
}

// 绘制_线终点, 函数绘制一条线从当前位置到, 但不包括指定点. 如果函数成功, 返回非零值.
//
// nXEnd: X坐标, 线结束点.
//
// nYEnd: Y坐标, 线结束点.

// ff:线终点
// nYEnd:Y坐标
// nXEnd:坐标
func (d *Draw) X线终点(坐标 int, Y坐标 int) bool {
	return 炫彩基类.X绘制_线终点(d.Handle, 坐标, Y坐标)
}

// 绘制_折线, Polyline() 参见MSDN.
//
// pArrayPt: 参见MSDN.
//
// arrayPtSize: 参见MSDN.

// ff:折线
// arrayPtSize:
// pArrayPt:
func (d *Draw) X折线(pArrayPt []炫彩基类.POINT, arrayPtSize int) bool {
	return 炫彩基类.X绘制_折线(d.Handle, pArrayPt, arrayPtSize)
}

// 绘制_置像素颜色, 函数设置在指定的坐标到指定的颜色的像素. 如果函数成功返回RGB值, 如果失败返回-1.
//
// X: 坐标.
//
// Y: 坐标.
//
// crColor: RGB颜色值.

// ff:置像素颜色
// crColor:RGB颜色值
// Y:坐标y
// X:坐标x
func (d *Draw) X置像素颜色(坐标x int, 坐标y int, RGB颜色值 int) int {
	return 炫彩基类.X绘制_置像素颜色(d.Handle, 坐标x, 坐标y, RGB颜色值)
}

// 绘制_取D2D渲染目标, 返回 *ID2D1RenderTarget.

// ff:取D2D渲染目标
func (d *Draw) X取D2D渲染目标() int {
	return 炫彩基类.X绘制_取D2D渲染目标(d.Handle)
}

// 绘制_图标, 绘制图标, DrawIconEx()参见MSDN.
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

// ff:图标
// diFlags:
// hbrFlickerFreeDraw:
// istepIfAniCur:
// cyWidth:
// cxWidth:
// hIcon:
// yTop:
// xLeft:
func (d *Draw) X图标(xLeft int, yTop int, hIcon uintptr, cxWidth int, cyWidth int, istepIfAniCur int, hbrFlickerFreeDraw int, diFlags int) bool {
	return 炫彩基类.X绘制_图标(d.Handle, xLeft, yTop, hIcon, cxWidth, cyWidth, istepIfAniCur, hbrFlickerFreeDraw, diFlags)
}

// 绘制_复制, BitBlt() 参见MSDN.
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

// ff:复制
// dwRop:
// nYSrc:
// nXSrc:
// hdcSrc:
// nHeight:
// nWidth:
// nYDest:
// nXDest:
func (d *Draw) X复制(nXDest, nYDest, nWidth, nHeight int32, hdcSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {
	return 炫彩基类.X绘制_复制(d.Handle, nXDest, nYDest, nWidth, nHeight, hdcSrc, nXSrc, nYSrc, dwRop)
}

// 绘制_复制2, BitBlt() 参见MSDN.
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

// ff:复制2
// dwRop:
// nYSrc:
// nXSrc:
// hDrawSrc:
// nHeight:
// nWidth:
// nYDest:
// nXDest:
func (d *Draw) X复制2(nXDest, nYDest, nWidth, nHeight int32, hDrawSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {
	return 炫彩基类.X绘制_复制2(d.Handle, nXDest, nYDest, nWidth, nHeight, hDrawSrc, nXSrc, nYSrc, dwRop)
}

// 绘制_带透明复制, AlphaBlend() 参见MSDN.
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

// ff:带透明复制
// alpha:
// nHeightSrc:
// nWidthSrc:
// nYOriginSrc:
// nXOriginSrc:
// hdcSrc:
// nHeightDest:
// nWidthDest:
// nYOriginDest:
// nXOriginDest:
func (d *Draw) X带透明复制(nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int32, hdcSrc uintptr, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha int32) bool {
	return 炫彩基类.X绘制_带透明复制(d.Handle, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest, hdcSrc, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha)
}

// 绘制_填充多边形, 填充多边形.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.

// ff:填充多边形
// nCount:
// points:顶点坐标切片
func (d *Draw) X填充多边形(顶点坐标切片 []炫彩基类.POINT, nCount int) int {
	return 炫彩基类.X绘制_填充多边形(d.Handle, 顶点坐标切片, nCount)
}

// 绘制_填充多边形F, 填充多边形.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.

// ff:填充多边形F
// nCount:
// points:顶点坐标切片
func (d *Draw) X填充多边形F(顶点坐标切片 []炫彩基类.POINTF, nCount int) int {
	return 炫彩基类.X绘制_填充多边形F(d.Handle, 顶点坐标切片, nCount)
}

// 绘制_图片.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.

// ff:图片
// y:y坐标
// x:x坐标
// hImageFrame:图片句柄
func (d *Draw) X图片(图片句柄 int, x坐标, y坐标 int32) {
	炫彩基类.X绘制_图片(d.Handle, 图片句柄, x坐标, y坐标)
}

// 绘制_图片F.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.

// ff:图片F
// y:y坐标
// x:x坐标
// hImageFrame:图片句柄
func (d *Draw) X图片F(图片句柄 int, x坐标, y坐标 float32) int {
	return 炫彩基类.X绘制_图片F(d.Handle, 图片句柄, x坐标, y坐标)
}

// 绘制_图片自适应.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bOnlyBorder: 是否只绘制边缘区域.

// ff:图片自适应
// bOnlyBorder:
// pRect:坐标
// hImageFrame:图片句柄
func (d *Draw) X图片自适应(图片句柄 int, 坐标 *炫彩基类.RECT, bOnlyBorder bool) int {
	return 炫彩基类.X绘制_图片自适应(d.Handle, 图片句柄, 坐标, bOnlyBorder)
}

// 绘制_图片自适应F.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bOnlyBorder: 是否只绘制边缘区域.

// ff:图片自适应F
// bOnlyBorder:
// pRect:坐标
// hImageFrame:图片句柄
func (d *Draw) X图片自适应F(图片句柄 int, 坐标 *炫彩基类.RECTF, bOnlyBorder bool) int {
	return 炫彩基类.X绘制_图片自适应F(d.Handle, 图片句柄, 坐标, bOnlyBorder)
}

// 绘制_图片扩展, 绘制图片.
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

// ff:图片EX
// height:高度
// width:宽度
// y:y坐标
// x:x坐标
// hImageFrame:图片句柄
func (d *Draw) X图片EX(图片句柄 int, x坐标, y坐标, 宽度, 高度 int) int {
	return 炫彩基类.X绘制_图片EX(d.Handle, 图片句柄, x坐标, y坐标, 宽度, 高度)
}

// 绘制_图片扩展F, 绘制图片.
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

// ff:图片EXF
// height:高度
// width:宽度
// y:y坐标
// x:x坐标
// hImageFrame:图片句柄
func (d *Draw) X图片EXF(图片句柄 int, x坐标, y坐标, 宽度, 高度 float32) int {
	return 炫彩基类.X绘制_图片EXF(d.Handle, 图片句柄, x坐标, y坐标, 宽度, 高度)
}

// 绘制_图片增强.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bClip: 是否裁剪区域.

// ff:图片增强
// bClip:
// pRect:坐标
// hImageFrame:图片句柄
func (d *Draw) X图片增强(图片句柄 int, 坐标 *炫彩基类.RECT, bClip bool) int {
	return 炫彩基类.X绘制_图片增强(d.Handle, 图片句柄, 坐标, bClip)
}

// 绘制_图片增强F.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bClip: 是否裁剪区域.

// ff:图片增强F
// bClip:
// pRect:坐标
// hImageFrame:图片句柄
func (d *Draw) X图片增强F(图片句柄 int, 坐标 *炫彩基类.RECTF, bClip bool) int {
	return 炫彩基类.X绘制_图片增强F(d.Handle, 图片句柄, 坐标, bClip)
}

// 绘制_图片增强扩展.
//
// hImageFrame: 图片句柄.
//
// prcDest: 目标坐标.
//
// prcSrc: 源坐标.

// ff:图片增强扩展
// prcSrc:
// prcDest:目标坐标
// hImageFrame:图片句柄
func (d *Draw) X图片增强扩展(图片句柄 int, 目标坐标 *炫彩基类.RECT, prcSrc *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_图片增强EX(d.Handle, 图片句柄, 目标坐标, prcSrc)
}

// 绘制_图片增强扩展F.
//
// hImageFrame: 图片句柄.
//
// prcDest: 目标坐标.
//
// prcSrc: 源坐标.

// ff:图片增强EXF
// prcSrc:
// prcDest:目标坐标
// hImageFrame:图片句柄
func (d *Draw) X图片增强EXF(图片句柄 int, 目标坐标 *炫彩基类.RECTF, prcSrc *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_图片增强EXF(d.Handle, 图片句柄, 目标坐标, prcSrc)
}

// 绘制_图片增强遮盖, 绘制带遮盖的图片. D2D留空.
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

// ff:图片增强遮盖
// bClip:
// pRectMask:
// pRect:坐标
// hImageFrameMask:图片遮盖句柄
// hImageFrame:图片句柄
func (d *Draw) X图片增强遮盖(图片句柄 int, 图片遮盖句柄 int, 坐标 *炫彩基类.RECT, pRectMask *炫彩基类.RECT, bClip bool) int {
	return 炫彩基类.X绘制_图片增强遮盖(d.Handle, 图片句柄, 图片遮盖句柄, 坐标, pRectMask, bClip)
}

// 绘制_图片平铺, 绘制图片.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.

// ff:图片平铺
// flag:
// pRect:坐标
// hImageFrameMask:
// hImageFrame:图片句柄
func (d *Draw) X图片平铺(图片句柄 int, hImageFrameMask int, 坐标 *炫彩基类.RECT, flag int) int {
	return 炫彩基类.X绘制_图片平铺(d.Handle, 图片句柄, hImageFrameMask, 坐标, flag)
}

// 绘制_图片平铺F, 绘制图片.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.

// ff:图片平铺F
// flag:
// pRect:坐标
// hImageFrameMask:
// hImageFrame:图片句柄
func (d *Draw) X图片平铺F(图片句柄 int, hImageFrameMask int, 坐标 *炫彩基类.RECTF, flag int) int {
	return 炫彩基类.X绘制_图片平铺F(d.Handle, 图片句柄, hImageFrameMask, 坐标, flag)
}

// 绘制_图片遮盖, 绘制带遮盖的图片, D2D留空.
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

// ff:图片遮盖
// y2:y2坐标
// x2:x2坐标
// y:y1坐标
// x:x1坐标
// hImageFrameMask:图片遮盖句柄
// hImageFrame:图片句柄
func (d *Draw) X图片遮盖(图片句柄 int, 图片遮盖句柄 int, x1坐标 int, y1坐标 int, x2坐标 int, y2坐标 int) int {
	return 炫彩基类.X绘制_图片遮盖(d.Handle, 图片句柄, 图片遮盖句柄, x1坐标, y1坐标, x2坐标, y2坐标)
}

// 绘制_文本指定矩形, DrawText() 参见MSDN.
//
// lpString: 字符串.
//
// lpRect: 坐标.

// ff:文本指定矩形
// lpRect:坐标
// lpString:字符串
func (d *Draw) X文本指定矩形(字符串 string, 坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_文本指定矩形(d.Handle, 字符串, 坐标)
}

// 绘制_文本指定矩形F, DrawText() 参见MSDN.
//
// lpString: 字符串.
//
// lpRect: 坐标.

// ff:文本指定矩形F
// lpRect:坐标
// lpString:字符串
func (d *Draw) X文本指定矩形F(字符串 string, 坐标 *炫彩基类.RECTF) int {
	return 炫彩基类.X绘制_文本指定矩形F(d.Handle, 字符串, 坐标)
}

// 绘制_文本下划线.
//
// lpString: 字符串.
//
// lpRect: 坐标.
//
// colorLine: 下划线颜色, ABGR 颜色.

// ff:文本下划线
// colorLine:
// lpRect:坐标
// lpString:字符串
func (d *Draw) X文本下划线(字符串 string, 坐标 *炫彩基类.RECT, colorLine int) int {
	return 炫彩基类.X绘制_文本下划线(d.Handle, 字符串, 坐标, colorLine)
}

// 绘制_文本下划线F.
//
// lpString: 字符串.
//
// lpRect: 坐标.
//
// colorLine: 下划线颜色, ABGR 颜色.

// ff:文本下划线F
// colorLine:
// lpRect:坐标
// lpString:字符串
func (d *Draw) X文本下划线F(字符串 string, 坐标 *炫彩基类.RECTF, colorLine int) int {
	return 炫彩基类.X绘制_文本下划线F(d.Handle, 字符串, 坐标, colorLine)
}

// 绘制_文本, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
//
// cbString: XX.

// ff:文本输出
// cbString:
// lpString:
// nYStart:
// nXStart:
func (d *Draw) X文本输出(nXStart int, nYStart int, lpString string, cbString string) int {
	return 炫彩基类.X绘制_文本(d.Handle, nXStart, nYStart, lpString, cbString)
}

// 绘制_文本F, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
//
// cbString: XX.

// ff:文本输出F
// cbString:
// lpString:
// nYStart:
// nXStart:
func (d *Draw) X文本输出F(nXStart, nYStart float32, lpString string, cbString string) int {
	return 炫彩基类.X绘制_文本F(d.Handle, nXStart, nYStart, lpString, cbString)
}

// 绘制_文本扩展, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.

// ff:文本输出EX
// lpString:
// nYStart:
// nXStart:
func (d *Draw) X文本输出EX(nXStart int, nYStart int, lpString string) int {
	return 炫彩基类.X绘制_文本EX(d.Handle, nXStart, nYStart, lpString)
}

// 绘制_文本扩展F, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.

// ff:文本输出EXF
// lpString:
// nYStart:
// nXStart:
func (d *Draw) X文本输出EXF(nXStart, nYStart float32, lpString string) int {
	return 炫彩基类.X绘制_文本EXF(d.Handle, nXStart, nYStart, lpString)
}

// 绘制_文本A, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.

// ff:文本输出A
// lpString:
// nYStart:
// nXStart:
func (d *Draw) X文本输出A(nXStart int, nYStart int, lpString string) int {
	return 炫彩基类.X绘制_文本A(d.Handle, nXStart, nYStart, lpString)
}

// 绘制_文本AF, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.

// ff:文本输出AF
// lpString:
// nYStart:
// nXStart:
func (d *Draw) X文本输出AF(nXStart, nYStart float32, lpString string) int {
	return 炫彩基类.X绘制_文本AF(d.Handle, nXStart, nYStart, lpString)
}

// 绘制_设置文本渲染提示.
//
// nType: XX.

// ff:设置文本渲染提示
// nType:
func (d *Draw) X设置文本渲染提示(nType int) int {
	return 炫彩基类.X绘制_设置文本渲染提示(d.Handle, nType)
}

// 绘制_SVG源.
//
// hSvg: SVG句柄.

// ff:SVG源
// hSvg:SVG句柄
func (d *Draw) SVG源(SVG句柄 int) int {
	return 炫彩基类.X绘制_SVG源(d.Handle, SVG句柄)
}

// 绘制_SVG.
//
// hSvg: SVG句柄.
//
// x: x坐标.
//
// y: y坐标.

// ff:SVG
// y:y坐标
// x:x坐标
// hSvg:SVG句柄
func (d *Draw) SVG(SVG句柄 int, x坐标 int, y坐标 int) int {
	return 炫彩基类.X绘制_SVG(d.Handle, SVG句柄, x坐标, y坐标)
}

// 绘制_SVG扩展.
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

// ff:SVGEX
// nHeight:高度
// nWidth:宽度
// y:y坐标
// x:x坐标
// hSvg:SVG句柄
func (d *Draw) SVGEX(SVG句柄 int, x坐标 int, y坐标 int, 宽度 int, 高度 int) int {
	return 炫彩基类.X绘制_SVGEX(d.Handle, SVG句柄, x坐标, y坐标, 宽度, 高度)
}

// 绘制_SVG大小.
//
// hSvg: SVG句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.

// ff:SVG大小
// nHeight:高度
// nWidth:宽度
// hSvg:SVG句柄
func (d *Draw) SVG大小(SVG句柄 int, 宽度 int, 高度 int) int {
	return 炫彩基类.X绘制_SVG大小(d.Handle, SVG句柄, 宽度, 高度)
}

// 绘制_D2D_清理, 使用指定颜色清理画布.
//
// color: ABGR 颜色值.

// ff:D2D清理
// color:ABGR
func (d *Draw) D2D清理(ABGR int) int {
	return 炫彩基类.X绘制_D2D_清理(d.Handle, ABGR)
}

// 绘制_图片遮盖矩形, 使用矩形作为遮罩.
//
// hImageFrame: 图片句柄.
//
// pRect: 矩形坐标.
//
// pRcMask: 遮罩坐标.
//
// pRcRoundAngle: 遮罩圆角.

// ff:图片遮盖矩形
// pRcRoundAngle:
// pRcMask:
// pRect:矩形坐标
// hImageFrame:图片句柄
func (d *Draw) X图片遮盖矩形(图片句柄 int, 矩形坐标 *炫彩基类.RECT, pRcMask *炫彩基类.RECT, pRcRoundAngle *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_图片遮盖矩形(d.Handle, 图片句柄, 矩形坐标, pRcMask, pRcRoundAngle)
}

// 绘制_图片遮盖圆型, 使用圆形作为遮罩.
//
// hImageFrame: 图片句柄.
//
// pRect: 矩形坐标.
//
// pRcMask: 遮罩坐标.

// ff:图片遮盖圆型
// pRcMask:
// pRect:矩形坐标
// hImageFrame:图片句柄
func (d *Draw) X图片遮盖圆型(图片句柄 int, 矩形坐标 *炫彩基类.RECT, pRcMask *炫彩基类.RECT) int {
	return 炫彩基类.X绘制_图片遮盖圆型(d.Handle, 图片句柄, 矩形坐标, pRcMask)
}

// 绘制_取字体, 返回字体句柄.

// ff:取字体
func (d *Draw) X取字体() int {
	return 炫彩基类.X绘制_取字体(d.Handle)
}
