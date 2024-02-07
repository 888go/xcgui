// 复选按钮
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
	w := window.New(0, 0, 430, 300, "复选按钮", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮
	Check1 := 炫彩组件类.NewButton(10, 35, 70, 30, "Check1", w.Handle)
	Check2 := 炫彩组件类.NewButton(10, 75, 70, 30, "Check2", w.Handle)
	Check3 := 炫彩组件类.NewButton(10, 115, 70, 30, "Check3", w.Handle)
	// 设置按钮类型
	Check1.SetTypeEx(炫彩常量类.Button_Type_Check)
	Check2.SetTypeEx(炫彩常量类.Button_Type_Check)
	Check3.SetTypeEx(炫彩常量类.Button_Type_Check)

	// 设置选中
	Check1.SetCheck(true)

	// 注册事件_按钮被选中
	Check1.Event_BUTTON_CHECK1(btn_check)
	Check2.Event_BUTTON_CHECK1(btn_check)
	Check3.Event_BUTTON_CHECK1(btn_check)

	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}

// 事件_按钮被选中
func btn_check(hEle int, bCheck bool, pbHandled *bool) int {
	if bCheck {
		fmt.Println(xc.XBtn_GetText(hEle), "Selected")
	} else {
		fmt.Println(xc.XBtn_GetText(hEle), "Unselected")
	}
	return 0
}
