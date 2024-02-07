// 单选按钮
package main

import (
	"fmt"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "单选按钮", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮
	Radio1 := 炫彩组件类.X创建按钮(10, 35, 70, 30, "Radio1", w.Handle)
	Radio2 := 炫彩组件类.X创建按钮(10, 75, 70, 30, "Radio2", w.Handle)
	Radio3 := 炫彩组件类.X创建按钮(10, 115, 70, 30, "Radio3", w.Handle)
	// 设置按钮类型
	Radio1.X置类型EX(炫彩常量类.Button_Type_Radio)
	Radio2.X置类型EX(炫彩常量类.Button_Type_Radio)
	Radio3.X置类型EX(炫彩常量类.Button_Type_Radio)
	// 设置分组id
	Radio1.X置组ID(1)
	Radio2.X置组ID(1)
	Radio3.X置组ID(1)

	// 设置选中
	Radio1.X置选中(true)

	// 注册事件_按钮被选中
	Radio1.X事件_选中1(btn_check)
	Radio2.X事件_选中1(btn_check)
	Radio3.X事件_选中1(btn_check)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}

// 事件_按钮被选中
func btn_check(hEle int, bCheck bool, pbHandled *bool) int {
	if bCheck {
		fmt.Println(炫彩基类.X按钮_取文本(hEle), "selected")
	}
	return 0
}
