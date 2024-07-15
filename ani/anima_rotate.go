package 炫彩动画类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
)

// AnimaRotate 动画旋转项.
type AnimaRotate struct {
	炫彩对象基类.ObjectBase
}

// 动画旋转_置中心, 设置旋转中心点坐标.
//
// x: 坐标X.
//
// y: 坐标Y.
//
// bOffset: TRUE: 相对于自身中心点偏移, FALSE: 绝对坐标.

// ff:置中心
// bOffset:偏移方式
// y:坐标Y
// x:坐标X
func (a *AnimaRotate) X置中心(坐标X float32, 坐标Y float32, 偏移方式 bool) bool {
	return 炫彩基类.X动画旋转_置中心(a.Handle, 坐标X, 坐标Y, 偏移方式)
}
