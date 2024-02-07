// 调用 wapi 选择颜色
package main

import (
	"fmt"
	"github.com/888go/xcgui/xc"
	"unsafe"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

var w *window.Window
var custColors [16]uint32 // 保存自定义颜色的数组

func main() {
	a := 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 430, 300, "选择颜色", 0, 炫彩常量类.Window_Style_Default)

	炫彩组件类.NewButton(20, 40, 100, 30, "选择颜色", w.Handle).Event_BnClick(onBnClick)

	a.ShowAndRun(w.Handle)
	a.Exit()
}

func onBnClick(pbHandled *bool) int {
	// 可直接调用: wutil.ChooseColor(w.Handle)
	ExampleChooseColorW()
	return 0
}

func ExampleChooseColorW() {
	cc := 炫彩WinApi类.ChooseColor{
		LStructSize:    36,
		HwndOwner:      w.GetHWND(),
		HInstance:      0,
		RgbResult:      0,
		LpCustColors:   &custColors[0],
		Flags:          炫彩WinApi类.CC_FULLOPEN, // 默认打开自定义颜色
		LCustData:      0,
		LpfnHook:       0,
		LpTemplateName: 0,
	}
	cc.LStructSize = uint32(unsafe.Sizeof(cc))
	ret := 炫彩WinApi类.ChooseColorW(&cc)
	fmt.Println(ret)
	if !ret {
		return
	}

	rgb := cc.RgbResult
	abgr := xc.RGB2ABGR(int(rgb), 255)
	fmt.Println("rgb颜色:", rgb)
	fmt.Println("abgr颜色:", abgr)
	fmt.Println(custColors) // 如果你添加了自定义颜色, 会保存在这个数组里面, 然后只要这个数组还在, 再次打开选择颜色界面时, 之前添加的自定义颜色还会存在

	// 设置窗口背景颜色
	w.AddBkFill(炫彩常量类.Window_State_Flag_Leave, abgr)
	w.Redraw(true)
}
