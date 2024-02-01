// 调用 wapi 选择颜色
package main

import (
	"fmt"
	"unsafe"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/wapi"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var w *window.Window
var custColors [16]uint32 // 保存自定义颜色的数组

func main() {
	a := app.I初始化(true)
	w = window.I窗口_创建(0, 0, 430, 300, "选择颜色", 0, xcc.I常量_窗口样式_默认)

	widget.I按钮_创建(20, 40, 100, 30, "选择颜色", w.I句柄).I事件_按钮被单击(onBnClick)

	a.I显示窗口并运行(w.I句柄)
	a.I退出()
}

func onBnClick(pbHandled *bool) int {
	ExampleChooseColorW()
	return 0
}

func ExampleChooseColorW() {
	cc := wapi.ChooseColor{
		LStructSize:    36,
		HwndOwner:      w.I取HWND(),
		HInstance:      0,
		RgbResult:      0,
		LpCustColors:   &custColors[0],
		Flags:          wapi.CC_FULLOPEN, // 默认打开自定义颜色
		LCustData:      0,
		LpfnHook:       0,
		LpTemplateName: 0,
	}
	cc.LStructSize = uint32(unsafe.Sizeof(cc))
	ret := wapi.ChooseColorW(&cc)
	fmt.Println(ret)
	fmt.Println(cc.RgbResult) // rgb颜色
	fmt.Println(custColors)   // 如果你添加了自定义颜色, 会保存在这个数组里面, 然后只要这个数组还在, 再次打开选择颜色界面时, 之前添加的自定义颜色还会存在

	// 设置窗口标题颜色
	w.I置标题颜色(xc.ABGR2(int(cc.RgbResult), 255))
	w.I重绘(true)
}
