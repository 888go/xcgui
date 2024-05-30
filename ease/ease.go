package ease//bm:炫彩缓动类

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 缓动_Linear, 线性.
//, 0.0f - 1.0f.
// ff:缓动线性
// p:位置
func Linear(p float32) float32 {
	return xc.XEase_Linear(p)
}

// 缓动_Quad, 二次方曲线.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动二次方曲线
// p:位置
// flag:缓动类型
func Quad(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Quad(p, flag)
}

// 缓动_Cubic, 三次方曲线, 圆弧.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动三次方曲线
// p:位置
// flag:缓动类型
func Cubic(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Cubic(p, flag)
}

// 缓动_Quart, 四次方曲线.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动四次方曲线
// p:位置
// flag:缓动类型
func Quart(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Quart(p, flag)
}

// 缓动_Quint, 五次方曲线.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动五次方曲线
// p:位置
// flag:缓动类型
func Quint(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Quint(p, flag)
}

// 缓动_Sine, 正弦曲线, 在末端变化.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动正弦曲线
// p:位置
// flag:缓动类型
func Sine(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Sine(p, flag)
}

// 缓动_Expo, 突击曲线, 突然一下.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动突击曲线
// p:位置
// flag:缓动类型
func Expo(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Expo(p, flag)
}

// 缓动_Circ, 圆环, 好比绕过一个圆环.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动圆环
// p:位置
// flag:缓动类型
func Circ(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Circ(p, flag)
}

// 缓动_Elastic, 强力回弹.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动强力回弹
// p:位置
// flag:缓动类型
func Elastic(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Elastic(p, flag)
}

// 缓动_Back, 回弹, 比较缓慢.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动回弹
// p:位置
// flag:缓动类型
func Back(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Back(p, flag)
}

// 缓动_Bounce.
//, 0.0f - 1.0f.
//, Ease_Type_.
// ff:缓动弹跳
// p:位置
// flag:缓动类型
func Bounce(p float32, flag xcc.Ease_Type_) float32 {
	return xc.XEase_Bounce(p, flag)
}

// 缓动_扩展, 全部缓动类型.
//, 0.0f - 1.0f.
//, Ease_Flag_.
// ff:缓动EX
// pos:位置
// flag:缓动标识
func Ex(pos float32, flag xcc.Ease_Flag_) float32 {
	return xc.XEase_Ex(pos, flag)
}
