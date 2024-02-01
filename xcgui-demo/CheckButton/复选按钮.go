// 复选按钮
package main

import (
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 430, 300, "", 0, xcc.I常量_窗口样式_默认)

	// 创建按钮
	Check1 := widget.I按钮_创建(10, 35, 70, 30, "Check1", w.I句柄)
	Check2 := widget.I按钮_创建(10, 75, 70, 30, "Check2", w.I句柄)
	Check3 := widget.I按钮_创建(10, 115, 70, 30, "Check3", w.I句柄)
	// 设置按钮类型
	Check1.I置类型扩展(xcc.I常量_对象扩展类型_复选按钮)
	Check2.I置类型扩展(xcc.I常量_对象扩展类型_复选按钮)
	Check3.I置类型扩展(xcc.I常量_对象扩展类型_复选按钮)

	// 设置选中
	Check1.I置选中(true)

	// 注册事件_按钮被选中
	Check1.I事件_按钮选中1(btn_check)
	Check2.I事件_按钮选中1(btn_check)
	Check3.I事件_按钮选中1(btn_check)

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
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
