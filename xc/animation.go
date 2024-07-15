package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"syscall"

	"github.com/888go/xcgui/xcc"
)

// 动画_运行, 并且增加引用计数.
//
// hAnimation: 动画序列或动画组句柄.
//
// hRedrawObjectUI: 当更新UI时重绘的UI层. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.

// ff:动画_运行
// hRedrawObjectUI:重绘的UI层
// hAnimation:动画序列或动画组句柄
func X动画_运行(动画序列或动画组句柄 int, 重绘的UI层 int) int {
	r, _, _ := xAnima_Run.Call(uintptr(动画序列或动画组句柄), uintptr(重绘的UI层))
	return int(r)
}

// 动画_释放, 停止动画, 并释放引用, 当引用计数为0时自动销毁.
//
// hAnimation: 动画序列或动画组句柄.
//
// bEnd: 是否立即执行到终点.

// ff:动画_释放
// bEnd:是否立即执行到终点
// hAnimation:动画序列或动画组句柄
func X动画_释放(动画序列或动画组句柄 int, 是否立即执行到终点 bool) bool {
	r, _, _ := xAnima_Release.Call(uintptr(动画序列或动画组句柄), 炫彩工具类.BoolPtr(是否立即执行到终点))
	return r != 0
}

// 动画_释放扩展, 释放与指定UI对象关联的所有动画, 返回释放动画数量.
//
// hObjectUI: 指定UI对象句柄.
//
// bEnd: 是否立即执行到终点.

// ff:动画_释放EX
// bEnd:是否立即执行到终点
// hObjectUI:指定UI对象句柄
func X动画_释放EX(指定UI对象句柄 int, 是否立即执行到终点 bool) int {
	r, _, _ := xAnima_ReleaseEx.Call(uintptr(指定UI对象句柄), 炫彩工具类.BoolPtr(是否立即执行到终点))
	return int(r)
}

// 动画_创建动画序列, 按顺序执行的动画列表, 返回动画序列句柄.
//
// hObjectUI: 绑定的UI对象. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.
//
// nLoopCount: 动画循环次数, 0: 无限循环.

// ff:动画_创建动画序列
// nLoopCount:动画循环次数
// hObjectUI:绑定的UI对象
func X动画_创建动画序列(绑定的UI对象 int, 动画循环次数 int) int {
	r, _, _ := xAnima_Create.Call(uintptr(绑定的UI对象), uintptr(动画循环次数))
	return int(r)
}

// 动画_移动, 移动到目标位置, 默认以UI对象中心点为操作方式, 避免出现坐标错位, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_移动
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// y:终点位置Y
// x:终点位置X
// duration:持续时间
// hSequence:动画序列句柄
func X动画_移动(动画序列句柄 int, 持续时间 int, 终点位置X float32, 终点位置Y float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Move.Call(uintptr(动画序列句柄), uintptr(持续时间), 炫彩工具类.Float32Ptr(终点位置X), 炫彩工具类.Float32Ptr(终点位置Y), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_移动扩展, 从指定位置移动到目标位置, 默认以UI对象中心点为操作方式, 避免出现坐标错位, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_移动EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// to_y:终点位置Y
// to_x:终点位置X
// from_y:起点位置Y
// from_x:起点位置X
// duration:持续时间
// hSequence:动画序列句柄
func X动画_移动EX(动画序列句柄 int, 持续时间 int, 起点位置X float32, 起点位置Y float32, 终点位置X float32, 终点位置Y float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_MoveEx.Call(uintptr(动画序列句柄), uintptr(持续时间), 炫彩工具类.Float32Ptr(起点位置X), 炫彩工具类.Float32Ptr(起点位置Y), 炫彩工具类.Float32Ptr(终点位置X), 炫彩工具类.Float32Ptr(终点位置Y), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_旋转, 旋转角度支持负数值, 因为负数可以控制反向旋转, 返回动画旋转项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_旋转
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// angle:角度
// duration:持续时间
// hSequence:动画序列句柄
func X动画_旋转(动画序列句柄 int, 持续时间 int, 角度 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Rotate.Call(uintptr(动画序列句柄), uintptr(持续时间), 炫彩工具类.Float32Ptr(角度), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_旋转扩展, 指定起点和终点, 返回动画旋转项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_旋转EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// to:终点角度
// from:起点角度
// duration:持续时间
// hSequence:动画序列句柄
func X动画_旋转EX(动画序列句柄 int, 持续时间 int, 起点角度 float32, 终点角度 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_RotateEx.Call(uintptr(动画序列句柄), uintptr(持续时间), 炫彩工具类.Float32Ptr(起点角度), 炫彩工具类.Float32Ptr(终点角度), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_缩放, 缩放对象, 默认以自身为中心缩放, 返回动画缩放项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_缩放
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// scaleY:Y轴缩放比例
// scaleX:轴缩放比例
// duration:持续时间
// hSequence:动画序列句柄
func X动画_缩放(动画序列句柄 int, 持续时间 int, 轴缩放比例 float32, Y轴缩放比例 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Scale.Call(uintptr(动画序列句柄), uintptr(持续时间), 炫彩工具类.Float32Ptr(轴缩放比例), 炫彩工具类.Float32Ptr(Y轴缩放比例), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_缩放大小, 修改UI对象大小, 默认向右延伸, 返回动画缩放项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_缩放大小
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// height:高度
// width:宽度
// duration:持续时间
// hSequence:动画序列句柄
func X动画_缩放大小(动画序列句柄 int, 持续时间 int, 宽度 float32, 高度 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_ScaleSize.Call(uintptr(动画序列句柄), uintptr(持续时间), 炫彩工具类.Float32Ptr(宽度), 炫彩工具类.Float32Ptr(高度), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_透明度, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_透明度
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// alpha:透明度
// duration:持续时间
// hSequence:动画序列句柄
func X动画_透明度(动画序列句柄 int, 持续时间 int, 透明度 uint8, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Alpha.Call(uintptr(动画序列句柄), uintptr(持续时间), uintptr(透明度), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_透明度扩展, 从指定透明度到目标透明度, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_透明度EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// to_alpha:终止透明度
// from_alpha:起始透明度
// duration:持续时间
// hSequence:动画序列句柄
func X动画_透明度EX(动画序列句柄 int, 持续时间 int, 起始透明度 uint8, 终止透明度 uint8, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_AlphaEx.Call(uintptr(动画序列句柄), uintptr(持续时间), uintptr(起始透明度), uintptr(终止透明度), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_颜色, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_颜色
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// color:ABGR
// duration:持续时间
// hSequence:动画序列句柄
func X动画_颜色(动画序列句柄 int, 持续时间 int, ABGR int, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Color.Call(uintptr(动画序列句柄), uintptr(持续时间), uintptr(ABGR), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_颜色扩展, 从指定颜色到目标颜色, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_颜色EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// to:终点颜色
// from:起点颜色
// duration:持续时间
// hSequence:动画序列句柄
func X动画_颜色EX(动画序列句柄 int, 持续时间 int, 起点颜色 int, 终点颜色 int, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_ColorEx.Call(uintptr(动画序列句柄), uintptr(持续时间), uintptr(起点颜色), uintptr(终点颜色), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_布局宽度, 修改布局宽度属性, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_布局宽度
// bGoBack:
// ease_flag:
// nLoopCount:
// width:
// nType:布局宽度类型
// duration:持续时间
// hSequence:动画序列句柄
func X动画_布局宽度(动画序列句柄 int, 持续时间 int, 布局宽度类型 炫彩常量类.Layout_Size_, width float32, nLoopCount int, ease_flag 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_LayoutWidth.Call(uintptr(动画序列句柄), uintptr(持续时间), uintptr(布局宽度类型), 炫彩工具类.Float32Ptr(width), uintptr(nLoopCount), uintptr(ease_flag), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_布局高度, 修改布局高度属性, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_布局高度
// bGoBack:
// ease_flag:
// nLoopCount:
// height:
// nType:布局高度类型
// duration:持续时间
// hSequence:动画序列句柄
func X动画_布局高度(动画序列句柄 int, 持续时间 int, 布局高度类型 炫彩常量类.Layout_Size_, height float32, nLoopCount int, ease_flag 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_LayoutHeight.Call(uintptr(动画序列句柄), uintptr(持续时间), uintptr(布局高度类型), 炫彩工具类.Float32Ptr(height), uintptr(nLoopCount), uintptr(ease_flag), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_布局大小, 修改布局宽度和高度, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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

// ff:动画_布局大小
// bGoBack:
// ease_flag:
// nLoopCount:
// height:
// nHeightType:
// width:
// nWidthType:布局宽度类型
// duration:持续时间
// hSequence:动画序列句柄
func X动画_布局大小(动画序列句柄 int, 持续时间 int, 布局宽度类型 炫彩常量类.Layout_Size_, width float32, nHeightType 炫彩常量类.Layout_Size_, height float32, nLoopCount int, ease_flag 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_LayoutSize.Call(uintptr(动画序列句柄), uintptr(持续时间), uintptr(布局宽度类型), 炫彩工具类.Float32Ptr(width), uintptr(nHeightType), 炫彩工具类.Float32Ptr(height), uintptr(nLoopCount), uintptr(ease_flag), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画_延迟, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
//
// duration: 持续时间.

// ff:动画_延迟
// duration:持续时间
// hSequence:动画序列句柄
func X动画_延迟(动画序列句柄 int, 持续时间 float32) int {
	r, _, _ := xAnima_Delay.Call(uintptr(动画序列句柄), 炫彩工具类.Float32Ptr(持续时间))
	return int(r)
}

// 动画_显示, 显示或隐藏UI对象, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
//
// duration: 持续时间.
//
// bShow: 显示或隐藏.

// ff:动画_显示
// bShow:显示或隐藏
// duration:持续时间
// hSequence:动画序列句柄
func X动画_显示(动画序列句柄 int, 持续时间 float32, 显示或隐藏 bool) int {
	r, _, _ := xAnima_Show.Call(uintptr(动画序列句柄), 炫彩工具类.Float32Ptr(持续时间), 炫彩工具类.BoolPtr(显示或隐藏))
	return int(r)
}

// 动画组_创建, 动画同步组, 当组中动画序列全部完成后, 重新开始.
//
// 当遇到无限循环项时, 直至其他序列完成后终止循环, 避免出现无法到达终点, 无法返回头部进行同步, 返回动画组句柄.
//
// nLoopCount: 动画循环次数, 0: 无限循环.

// ff:动画组_创建
// nLoopCount:循环次数
func X动画组_创建(循环次数 int) int {
	r, _, _ := xAnimaGroup_Create.Call(uintptr(循环次数))
	return int(r)
}

// 动画组_添加项, 将动画序列添加到组中.
//
// hGroup: 动画组句柄.
//
// hSequence: 动画序列句柄.

// ff:动画组_添加项
// hSequence:动画序列句柄
// hGroup:动画组句柄
func X动画组_添加项(动画组句柄 int, 动画序列句柄 int) int {
	r, _, _ := xAnimaGroup_AddItem.Call(uintptr(动画组句柄), uintptr(动画序列句柄))
	return int(r)
}

// 动画旋转_置中心, 设置旋转中心点坐标.
//
// hAnimationRotate: 动画旋转项句柄.
//
// x: 坐标X.
//
// y: 坐标Y.
//
// bOffset: TRUE: 相对于自身中心点偏移, FALSE: 绝对坐标.

// ff:动画旋转_置中心
// bOffset:TRUE
// y:坐标Y
// x:坐标X
// hAnimationRotate:动画旋转项句柄
func X动画旋转_置中心(动画旋转项句柄 int, 坐标X float32, 坐标Y float32, TRUE bool) bool {
	r, _, _ := xAnimaRotate_SetCenter.Call(uintptr(动画旋转项句柄), 炫彩工具类.Float32Ptr(坐标X), 炫彩工具类.Float32Ptr(坐标Y), 炫彩工具类.BoolPtr(TRUE))
	return r != 0
}

// 动画缩放_置延伸位置, 设置缩放起点, 确定延伸方向.
//
// hAnimationScale: 动画缩放项句柄.
//
// position: 位置, Position_Flag_.

// ff:动画缩放_置延伸位置
// position:位置
// hAnimationScale:动画缩放项句柄
func X动画缩放_置延伸位置(动画缩放项句柄 int, 位置 炫彩常量类.Position_Flag_) bool {
	r, _, _ := xAnimaScale_SetPosition.Call(uintptr(动画缩放项句柄), uintptr(位置))
	return r != 0
}

// 动画_取UI对象, 获取动画关联的UI对象, 返回UI对象句柄.
//
// hAnimation: 动画序列或动画组句柄.

// ff:动画_取UI对象
// hAnimation:动画序列或动画组句柄
func X动画_取UI对象(动画序列或动画组句柄 int) int {
	r, _, _ := xAnima_GetObjectUI.Call(uintptr(动画序列或动画组句柄))
	return int(r)
}

// 动画项_启用完成释放, 当动画项完成后自动释放.
//
// 例如对多个动画序列进行渐近式延迟, 在动画序列头标添加延时项(时间差), 当延时项完成时自动释放, 后续动画循环就形成一种时间差(因为对齐的时间差销毁了, 他们永远无法对齐时间).
//
// hAnimationItem: 动画项句柄.
//
// bEnable: 是否启用.

// ff:动画项_启用完成释放
// bEnable:是否启用
// hAnimationItem:动画项句柄
func X动画项_启用完成释放(动画项句柄 int, 是否启用 bool) int {
	r, _, _ := xAnimaItem_EnableCompleteRelease.Call(uintptr(动画项句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 动画_启用自动销毁, TRUE: 当引用计数为0时自动销毁, FALSE: 手动销毁.
//
// hAnimation: 动画项或动画序列或动画组句柄.
//
// bEnable: 是否启用.

// ff:动画_启用自动销毁
// bEnable:是否启用
// hAnimation:动画项或动画序列或动画组句柄
func X动画_启用自动销毁(动画项或动画序列或动画组句柄 int, 是否启用 bool) int {
	r, _, _ := xAnima_EnableAutoDestroy.Call(uintptr(动画项或动画序列或动画组句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 动画_销毁UI对象, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
//
// duration: 持续时间.

// ff:动画_销毁UI对象
// duration:持续时间
// hSequence:动画序列句柄
func X动画_销毁UI对象(动画序列句柄 int, 持续时间 float32) int {
	r, _, _ := xAnima_DestroyObjectUI.Call(uintptr(动画序列句柄), 炫彩工具类.Float32Ptr(持续时间))
	return int(r)
}

// 动画_置回调.
//
// hAnimationEx: 动画序列或动画组句柄.
//
// callback: 回调函数.

// ff:动画_置回调
// callback:回调函数
// hAnimationEx:动画序列或动画组句柄
func X动画_置回调(动画序列或动画组句柄 int, 回调函数 interface{}) int {
	r, _, _ := xAnima_SetCallBack.Call(uintptr(动画序列或动画组句柄), syscall.NewCallback(回调函数))
	return int(r)
}

// 动画_置用户数据.
//
// hAnimationEx: 动画序列或动画组句柄.
//
// nUserData: 用户数据.

// ff:动画_置用户数据
// nUserData:用户数据
// hAnimationEx:动画序列或动画组句柄
func X动画_置用户数据(动画序列或动画组句柄, 用户数据 int) int {
	r, _, _ := xAnima_SetUserData.Call(uintptr(动画序列或动画组句柄), uintptr(用户数据))
	return int(r)
}

// 动画_取用户数据, 返回用户数据.
//
// hAnimationEx: 动画序列或动画组句柄.

// ff:动画_取用户数据
// hAnimationEx:动画序列或动画组句柄
func X动画_取用户数据(动画序列或动画组句柄 int) int {
	r, _, _ := xAnima_GetUserData.Call(uintptr(动画序列或动画组句柄))
	return int(r)
}

// 动画_停止.
//
// hAnimationEx: 动画序列或动画组句柄.

// ff:动画_停止
// hAnimationEx:动画序列或动画组句柄
func X动画_停止(动画序列或动画组句柄 int) bool {
	r, _, _ := xAnima_Stop.Call(uintptr(动画序列或动画组句柄))
	return r != 0
}

// 动画_启动.
//
// hAnimationEx: 动画序列或动画组句柄.

// ff:动画_启动
// hAnimationEx:动画序列或动画组句柄
func X动画_启动(动画序列或动画组句柄 int) bool {
	r, _, _ := xAnima_Start.Call(uintptr(动画序列或动画组句柄))
	return r != 0
}

// 动画_暂停.
//
// hAnimationEx: 动画序列或动画组句柄.

// ff:动画_暂停
// hAnimationEx:动画序列或动画组句柄
func X动画_暂停(动画序列或动画组句柄 int) bool {
	r, _, _ := xAnima_Pause.Call(uintptr(动画序列或动画组句柄))
	return r != 0
}

// 动画项_置回调.
//
// hAnimationItem: 动画项句柄.
//
// callback: 回调函数.

// ff:动画项_置回调
// callback:回调函数
// hAnimationItem:动画项句柄
func X动画项_置回调(动画项句柄 int, 回调函数 interface{}) int {
	r, _, _ := xAnimaItem_SetCallback.Call(uintptr(动画项句柄), syscall.NewCallback(回调函数))
	return int(r)
}

// 动画项_置用户数据.
//
// hAnimationItem: 动画项句柄.
//
// nUserData: 用户数据.

// ff:动画项_置用户数据
// nUserData:用户数据
// hAnimationItem:动画项句柄
func X动画项_置用户数据(动画项句柄, 用户数据 int) int {
	r, _, _ := xAnimaItem_SetUserData.Call(uintptr(动画项句柄), uintptr(用户数据))
	return int(r)
}

// 动画项_取用户数据, 返回用户数据.
//
// hAnimationItem: 动画项句柄.

// ff:动画项_取用户数据
// hAnimationItem:动画项句柄
func X动画项_取用户数据(动画项句柄 int) int {
	r, _, _ := xAnimaItem_GetUserData.Call(uintptr(动画项句柄))
	return int(r)
}

// 动画项_启用自动销毁.
//
// hAnimationItem: 动画项句柄.
//
// bEnable: 是否启用.

// ff:动画项_启用自动销毁
// bEnable:是否启用
// hAnimationItem:动画项句柄
func X动画项_启用自动销毁(动画项句柄 int, 是否启用 bool) int {
	r, _, _ := xAnimaItem_EnableAutoDestroy.Call(uintptr(动画项句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 动画_延迟扩展, 可以作为一个空动画, 然后在回调里处理自己的算法, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
//
// duration: 持续时间.
//
// nLoopCount: 动画循环次数, 0:无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.

// ff:动画_延迟EX
// bGoBack:
// ease_flag:缓动标识
// nLoopCount:动画循环次数
// duration:持续时间
// hSequence:动画序列句柄
func X动画_延迟EX(动画序列句柄 int, 持续时间 float32, 动画循环次数 int, 缓动标识 炫彩常量类.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_DelayEx.Call(uintptr(动画序列句柄), 炫彩工具类.Float32Ptr(持续时间), uintptr(动画循环次数), uintptr(缓动标识), 炫彩工具类.BoolPtr(bGoBack))
	return int(r)
}

// 动画移动_置标识, 此接口可独立设置x轴移动或y轴移动.
//
// hAnimationMove: 动画移动项句柄.
//
// flags: 动画移动标识, 可组合使用, Animation_Move_.
//
// TODO: 此函数尚未封装到类中.

// ff:动画移动_置标识
// flags:动画移动标识
// hAnimationMove:动画移动项句柄
func X动画移动_置标识(动画移动项句柄 int, 动画移动标识 炫彩常量类.Animation_Move_) int {
	r, _, _ := xAnimaMove_SetFlag.Call(uintptr(动画移动项句柄), uintptr(动画移动标识))
	return int(r)
}
