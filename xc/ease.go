package xc

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xcc"
)

// 缓动_Linear, 线性.
//, 0.0f - 1.0f.
// ff:缓动_线性
// p:位置
func XEase_Linear(p float32) float32 {
	_, r, _ := xEase_Linear.Call(common.Float32Ptr(p))
	return common.UintPtrToFloat32(r)
}

// 缓动_Quad, 二次方曲线.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_二次方曲线
// p:位置
// flag:缓动类型
func XEase_Quad(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Quad.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Cubic, 三次方曲线, 圆弧.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_三次方曲线
// p:位置
// flag:缓动类型
func XEase_Cubic(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Cubic.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Quart, 四次方曲线.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_四次方曲线
// p:位置
// flag:缓动类型
func XEase_Quart(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Quart.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Quint, 五次方曲线.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_五次方曲线
// p:位置
// flag:缓动类型
func XEase_Quint(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Quint.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Sine, 正弦曲线, 在末端变化.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_正弦曲线
// p:位置
// flag:缓动类型
func XEase_Sine(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Sine.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Expo, 突击曲线, 突然一下.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_突击曲线
// p:位置
// flag:缓动类型
func XEase_Expo(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Expo.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Circ, 圆环, 好比绕过一个圆环.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_圆环
// p:位置
// flag:缓动类型
func XEase_Circ(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Circ.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Elastic, 强力回弹.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_强力回弹
// p:位置
// flag:缓动类型
func XEase_Elastic(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Elastic.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Back, 回弹, 比较缓慢.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_回弹
// p:位置
// flag:缓动类型
func XEase_Back(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Back.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_Bounce, 弹跳, 模拟小球落地弹跳.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动_弹跳
// p:位置
// flag:缓动类型
func XEase_Bounce(p float32, flag xcc.Ease_Type_) float32 {
	_, r, _ := xEase_Bounce.Call(common.Float32Ptr(p), uintptr(flag))
	return common.UintPtrToFloat32(r)
}

// 缓动_扩展, 全部缓动类型.
//, 0.0f - 1.0f.
//, Ease_Flag_.
// ff:缓动_全部缓动类型
// pos:位置
// flag:缓动标识
func XEase_Ex(pos float32, flag xcc.Ease_Flag_) float32 {
	_, r, _ := xEase_Ex.Call(common.Float32Ptr(pos), uintptr(flag))
	return common.UintPtrToFloat32(r)
}
