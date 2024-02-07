package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// 缓动_线性, 线性.
// p:位置, 0.0f - 1.0f.
func XEase_Linear(p float32) float32 {
	_, r, _ : xEase_Linear.Call(common.Float32Ptr(p))
	return common.UintPtrToFloat32(r)
}

// 缓动_二次方曲线, 二次方曲线.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Quad(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Quad.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_三次方曲线, 三次方曲线, 圆弧.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Cubic(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Cubic.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_四次方曲线, 四次方曲线.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Quart(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Quart.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_五次方曲线, 五次方曲线.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Quint(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Quint.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_正弦曲线, 正弦曲线, 在末端变化.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Sine(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Sine.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_突击曲线, 突击曲线, 突然一下.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Expo(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Expo.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_圆环, 圆环, 好比绕过一个圆环.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Circ(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Circ.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_强力回弹, 强力回弹.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Elastic(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Elastic.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_回弹, 回弹, 比较缓慢.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Back(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Back.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_弹跳, 弹跳, 模拟小球落地弹跳.
// p:位置, 0.0f - 1.0f.
// flag:缓动类型, Ease_Type_.
func XEase_Bounce(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ : xEase_Bounce.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_全部缓动类型, 全部缓动类型.
// pos:位置, 0.0f - 1.0f.
// flag:缓动标识, Ease_Flag_.
func XEase_Ex(pos float32, flag xcc.Ease_Flag_) float32 {
	_, r, _ : xEase_Ex.Call(common.Float32Ptr(pos), uintptr(flag))
	return common.UintPtrToFloat32(r)
}
