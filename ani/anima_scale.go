package 炫彩动画类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// AnimaScale 动画缩放项.
type AnimaScale struct {
	炫彩对象基类.ObjectBase
}

// 动画缩放_置延伸位置, 设置缩放起点, 确定延伸方向.
//
// position: 位置, Position_Flag_.
func (a *AnimaScale) SetPosition(position 炫彩常量类.Position_Flag_) bool {
	return xc.XAnimaScale_SetPosition(a.Handle, position)
}
