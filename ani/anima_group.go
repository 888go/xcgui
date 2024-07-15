package 炫彩动画类

import (
	"github.com/888go/xcgui/xc"
)

// AnimaGroup 动画组.
type AnimaGroup struct {
	animaBase
}

// 动画组_创建, 动画同步组, 当组中动画序列全部完成后, 重新开始.
//
// 当遇到无限循环项时, 直至其他序列完成后终止循环, 避免出现无法到达终点, 无法返回头部进行同步, 返回动画组句柄.
//
// nLoopCount: 动画循环次数, 0: 无限循环.

// ff:创建动画组
// nLoopCount:动画循环次数
func X创建动画组(动画循环次数 int) *AnimaGroup {
	p := &AnimaGroup{}
	p.X设置句柄(炫彩基类.X动画组_创建(动画循环次数))
	return p
}

// 动画组_添加项, 将动画序列添加到组中.
//
// hSequence: 动画序列句柄.

// ff:添加项
// hSequence:动画序列句柄
func (a *AnimaGroup) X添加项(动画序列句柄 int) int {
	return 炫彩基类.X动画组_添加项(a.Handle, 动画序列句柄)
}
