package 炫彩动画类 //bm:炫彩动画类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Anima 动画序列.
type Anima struct {
	animaBase
}

// 动画_创建动画序列, 按顺序执行的动画列表, 返回动画序列对象.
//
// hObjectUI: 绑定的UI对象. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.
//
// nLoopCount: 动画循环次数, 0: 无限循环.

// ff:创建动画序列
// nLoopCount:动画循环次数
// hObjectUI:绑定的UI对象
func X创建动画序列(绑定的UI对象 int, 动画循环次数 int) *Anima {
	p := &Anima{}
	p.X设置句柄(炫彩基类.X动画_创建动画序列(绑定的UI对象, 动画循环次数))
	return p
}

// 动画_移动, 移动到目标位置, 默认以UI对象中心点为操作方式, 避免出现坐标错位, 返回动画项对象.
//
// duration: 持续时间.
//
// x: 终点位置X(对象左上角坐标).
//
// y: 终点位置Y(对象左上角坐标).
//
// nLoopCount: 动画循环次数, 0:无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:移动
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// y:终点位置Y
// x:终点位置X
// duration:持续时间
func (a *Anima) X移动(持续时间 int, 终点位置X float32, 终点位置Y float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_移动(a.Handle, 持续时间, 终点位置X, 终点位置Y, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_移动扩展, 从指定位置移动到目标位置, 默认以UI对象中心点为操作方式, 避免出现坐标错位, 返回动画项对象.
//
// duration: 持续时间.
//
// from_x: 起点位置X(对象左上角坐标).
//
// from_y: 起点位置Y(对象左上角坐标).
//
// to_x: 终点位置X(对象左上角坐标).
//
// to_y: 终点位置Y(对象左上角坐标).
//
// nLoopCount: 动画循环次数, 0:无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:移动EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// to_y:终点位置Y
// to_x:终点位置X
// from_y:起点位置Y
// from_x:起点位置X
// duration:持续时间
func (a *Anima) X移动EX(持续时间 int, 起点位置X float32, 起点位置Y float32, 终点位置X float32, 终点位置Y float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_移动EX(a.Handle, 持续时间, 起点位置X, 起点位置Y, 终点位置X, 终点位置Y, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_旋转, 旋转角度支持负数值, 因为负数可以控制反向旋转, 返回动画旋转项对象.
//
// duration: 持续时间.
//
// angle: 角度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:旋转
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// angle:角度
// duration:持续时间
func (a *Anima) X旋转(持续时间 int, 角度 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaRotate {
	p := &AnimaRotate{}
	p.X设置句柄(炫彩基类.X动画_旋转(a.Handle, 持续时间, 角度, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_旋转扩展, 指定起点和终点, 返回动画旋转项对象.
//
// duration: 持续时间.
//
// from: 起点角度.
//
// to: 终点角度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:旋转EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// to:终点角度
// from:起点角度
// duration:持续时间
func (a *Anima) X旋转EX(持续时间 int, 起点角度 float32, 终点角度 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaRotate {
	p := &AnimaRotate{}
	p.X设置句柄(炫彩基类.X动画_旋转EX(a.Handle, 持续时间, 起点角度, 终点角度, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_缩放, 缩放对象, 默认以自身为中心缩放, 返回动画缩放项对象.
//
// duration: 持续时间.
//
// scaleX: X轴缩放比例.
//
// scaleY: Y轴缩放比例.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:缩放
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// scaleY:Y轴缩放比例
// scaleX:轴缩放比例
// duration:持续时间
func (a *Anima) X缩放(持续时间 int, 轴缩放比例 float32, Y轴缩放比例 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaScale {
	p := &AnimaScale{}
	p.X设置句柄(炫彩基类.X动画_缩放(a.Handle, 持续时间, 轴缩放比例, Y轴缩放比例, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_缩放大小, 修改UI对象大小, 默认向右延伸, 返回动画缩放项对象.
//
// duration: 持续时间.
//
// width: 宽度.
//
// height: 高度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:缩放大小
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// height:高度
// width:宽度
// duration:持续时间
func (a *Anima) X缩放大小(持续时间 int, 宽度 float32, 高度 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaScale {
	p := &AnimaScale{}
	p.X设置句柄(炫彩基类.X动画_缩放大小(a.Handle, 持续时间, 宽度, 高度, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_透明度, 返回动画项对象.
//
// duration: 持续时间.
//
// alpha: 透明度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:透明度
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// alpha:透明度
// duration:持续时间
func (a *Anima) X透明度(持续时间 int, 透明度 uint8, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_透明度(a.Handle, 持续时间, 透明度, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_透明度扩展, 从指定透明度到目标透明度, 返回动画项对象.
//
// duration: 持续时间.
//
// from_alpha: 起始透明度.
//
// to_alpha: 终止透明度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:透明度EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// to_alpha:终止透明度
// from_alpha:起始透明度
// duration:持续时间
func (a *Anima) X透明度EX(持续时间 int, 起始透明度 uint8, 终止透明度 uint8, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_透明度EX(a.Handle, 持续时间, 起始透明度, 终止透明度, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_颜色, 返回动画项对象.
//
// duration: 持续时间.
//
// color: ABGR 颜色.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:颜色
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// color:ABGR
// duration:持续时间
func (a *Anima) X颜色(持续时间 int, ABGR int, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_颜色(a.Handle, 持续时间, ABGR, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_颜色扩展, 从指定颜色到目标颜色, 返回动画项对象.
//
// duration: 持续时间.
//
// from: 起点颜色, ABGR 颜色.
//
// to: 终点颜色, ABGR 颜色.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:颜色EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// to:终点颜色
// from:起点颜色
// duration:持续时间
func (a *Anima) X颜色EX(持续时间 int, 起点颜色 int, 终点颜色 int, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_颜色EX(a.Handle, 持续时间, 起点颜色, 终点颜色, 动画循环次数, 缓动标识, bGoBack))
	return p
}

// 动画_布局宽度, 修改布局宽度属性, 返回动画项对象.
//
// duration: 持续时间.
//
// nType: 布局宽度类型: xcc.Layout_Size_.
//
// width: 布局宽度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:布局宽度
// bGoBack:
// ease_flag:
// nLoopCount:
// width:
// nType:布局宽度类型
// duration:持续时间
func (a *Anima) X布局宽度(持续时间 int, 布局宽度类型 炫彩常量类.Layout_Size_, width float32, nLoopCount int, ease_flag 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_布局宽度(a.Handle, 持续时间, 布局宽度类型, width, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_布局高度, 修改布局高度属性, 返回动画项对象.
//
// duration: 持续时间.
//
// nType: 布局高度类型: xcc.Layout_Size_.
//
// height: 布局高度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:布局高度
// bGoBack:
// ease_flag:
// nLoopCount:
// height:
// nType:布局高度类型
// duration:持续时间
func (a *Anima) X布局高度(持续时间 int, 布局高度类型 炫彩常量类.Layout_Size_, height float32, nLoopCount int, ease_flag 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_布局高度(a.Handle, 持续时间, 布局高度类型, height, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_布局大小, 修改布局宽度和高度, 返回动画项对象.
//
// duration: 持续时间.
//
// nWidthType: 布局宽度类型: xcc.Layout_Size_.
//
// width: 布局宽度.
//
// nHeightType: 布局高度类型: xcc.Layout_Size_.
//
// height: 布局高度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:布局大小
// bGoBack:
// ease_flag:
// nLoopCount:
// height:
// nHeightType:
// width:
// nWidthType:布局宽度类型
// duration:持续时间
func (a *Anima) X布局大小(持续时间 int, 布局宽度类型 炫彩常量类.Layout_Size_, width float32, nHeightType 炫彩常量类.Layout_Size_, height float32, nLoopCount int, ease_flag 炫彩常量类.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_布局大小(a.Handle, 持续时间, 布局宽度类型, width, nHeightType, height, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_延迟, 返回动画项对象.
//
// duration: 持续时间.

// ff:延迟
// duration:持续时间
func (a *Anima) X延迟(持续时间 float32) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_延迟(a.Handle, 持续时间))
	return p
}

// 动画_显示, 显示或隐藏UI对象, 返回动画项对象.
//
// duration: 持续时间.
//
// bShow: 显示或隐藏.

// ff:显示
// bShow:显示或隐藏
// duration:持续时间
func (a *Anima) X显示(持续时间 float32, 显示或隐藏 bool) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_显示(a.Handle, 持续时间, 显示或隐藏))
	return p
}

// 动画_销毁UI对象, 返回动画项对象.
//
// duration: 持续时间.

// ff:销毁UI对象
// duration:持续时间
func (a *Anima) X销毁UI对象(持续时间 float32) *AnimaItem {
	p := &AnimaItem{}
	p.X设置句柄(炫彩基类.X动画_销毁UI对象(a.Handle, 持续时间))
	return p
}

// 动画_延迟扩展, 可以作为一个空动画, 然后在回调里处理自己的算法, 返回动画项句柄.
//
// duration: 持续时间.
//
// nLoopCount: 动画循环次数, 0:无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:延迟EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// duration:持续时间
func (a *Anima) X延迟EX(持续时间 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	return 炫彩基类.X动画_延迟EX(a.Handle, 持续时间, 动画循环次数, 缓动标识, bGoBack)
}
