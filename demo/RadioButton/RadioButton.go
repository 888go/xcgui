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
	a := 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 430, 300, "单选按钮", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮
	Radio1 := 炫彩组件类.NewButton(10, 35, 70, 30, "Radio1", w.Handle)
	Radio2 := 炫彩组件类.NewButton(10, 75, 70, 30, "Radio2", w.Handle)
	Radio3 := 炫彩组件类.NewButton(10, 115, 70, 30, "Radio3", w.Handle)
	// 设置按钮类型
	Radio1.SetTypeEx(炫彩常量类.Button_Type_Radio)
	Radio2.SetTypeEx(炫彩常量类.Button_Type_Radio)
	Radio3.SetTypeEx(炫彩常量类.Button_Type_Radio)
	// 设置分组id
	Radio1.SetGroupID(1)
	Radio2.SetGroupID(1)
	Radio3.SetGroupID(1)

	// 设置选中
	Radio1.SetCheck(true)

	// 注册事件_按钮被选中
	Radio1.Event_BUTTON_CHECK1(btn_check)
	Radio2.Event_BUTTON_CHECK1(btn_check)
	Radio3.Event_BUTTON_CHECK1(btn_check)

	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}

// 事件_按钮被选中
func btn_check(hEle int, bCheck bool, pbHandled *bool) int {
	if bCheck {
		fmt.Println(xc.XBtn_GetText(hEle), "selected")
	}
	return 0
}
