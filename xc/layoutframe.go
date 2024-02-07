package 炫彩基类

import (
	"github.com/888go/xcgui/common"
)

// 布局框架_创建.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func X布局框架_创建(左 int, 右 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xLayoutFrame_Create.Call(uintptr(左), uintptr(右), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 布局框架_显示布局边界.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X布局框架_显示布局边界(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xLayoutFrame_ShowLayoutFrame.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}
