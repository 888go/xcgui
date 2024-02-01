package ani

import (
	"e.coding.net/gogit/go/xcgui/objectbase"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 动画缩放项.
type AnimaScale struct {
	objectbase.ObjectBase
}

// 动画缩放_置延伸位置, 设置缩放起点, 确定延伸方向.
//
// position: 位置, I常量_位置标识_.
func (a *AnimaScale) SetPosition(position xcc.I常量_位置标识_) bool {
	return xc.XAnimaScale_SetPosition(a.I句柄, position)
}
