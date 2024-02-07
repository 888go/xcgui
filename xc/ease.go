package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// 缓动_Linear, 线性.
//
// p: 位置, 0.0f - 1.0f.
func X缓动_线性(位置 float32) float32 {
	_, r, _ := xEase_Linear.Call(炫彩工具类.Float32Ptr(位置))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Quad, 二次方曲线.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_二次方曲线(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Quad.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Cubic, 三次方曲线, 圆弧.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_三次方曲线(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Cubic.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Quart, 四次方曲线.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_四次方曲线(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Quart.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Quint, 五次方曲线.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_五次方曲线(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Quint.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Sine, 正弦曲线, 在末端变化.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_正弦曲线(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Sine.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Expo, 突击曲线, 突然一下.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_突击曲线(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Expo.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Circ, 圆环, 好比绕过一个圆环.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_圆环(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Circ.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Elastic, 强力回弹.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_强力回弹(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Elastic.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Back, 回弹, 比较缓慢.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_回弹(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Back.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_Bounce, 弹跳, 模拟小球落地弹跳.
//
// p: 位置, 0.0f - 1.0f.
//
// flag: 缓动类型, Ease_Type_.
func X缓动_弹跳(位置 float32, 缓动类型 炫彩常量类.Ease_Type_) float32 {
	_, r, _ := xEase_Bounce.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动类型))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 缓动_扩展, 全部缓动类型.
//
// pos: 位置, 0.0f - 1.0f.
//
// flag: 缓动标识, Ease_Flag_.
func X缓动_全部缓动类型(位置 float32, 缓动标识 炫彩常量类.Ease_Flag_) float32 {
	_, r, _ := xEase_Ex.Call(炫彩工具类.Float32Ptr(位置), uintptr(缓动标识))
	return 炫彩工具类.UintPtrToFloat32(r)
}
