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
func (a *AnimaRotate) SetCenter(x float32, y float32, bOffset bool) bool {
	return xc.XAnimaRotate_SetCenter(a.Handle, x, y, bOffset)
}
