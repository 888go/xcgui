// 滚动条, 设置背景, 获取滚动条上的三个按钮并加以改变
package main

import (
	"fmt"
	"github.com/888go/xcgui/xc"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 430, 300, "ScrollBar", 0, 炫彩常量类.Window_Style_Default)

	// 创建滚动条
	bar1 := 炫彩组件类.NewScrollBar(12, 33, 300, 20, w.Handle)
	bar2 := 炫彩组件类.NewScrollBar(330, 33, 20, 240, w.Handle)

	// 设置为垂直滚动条
	bar2.EnableHorizon(false)
	// 添加背景
	bar2.AddBkFill(炫彩常量类.Element_State_Flag_Leave, xc.RGBA(247, 248, 250, 255))

	// 获取滑块按钮
	btnSlider := 炫彩组件类.NewButtonByHandle(bar2.GetButtonSlider())
	btnSlider.AddBkFill(炫彩常量类.Button_State_Flag_Leave, xc.RGBA(221, 221, 223, 255))
	btnSlider.AddBkFill(炫彩常量类.Button_State_Flag_Stay, xc.RGBA(202, 202, 204, 255))
	btnSlider.AddBkFill(炫彩常量类.Button_State_Flag_Down, xc.RGBA(202, 202, 204, 255))
	// 获取滚动条上按钮
	btnUp := 炫彩组件类.NewButtonByHandle(bar2.GetButtonUp())
	btnUp.AddBkFill(炫彩常量类.Button_State_Flag_Leave, xc.RGBA(137, 140, 151, 255))
	btnUp.AddBkFill(炫彩常量类.Button_State_Flag_Stay, xc.RGBA(255, 135, 250, 255))
	btnUp.AddBkFill(炫彩常量类.Button_State_Flag_Down, xc.RGBA(255, 75, 250, 255))
	// 获取滚动条下按钮
	btnDown := 炫彩组件类.NewButtonByHandle(bar2.GetButtonDown())
	btnDown.AddBkFill(炫彩常量类.Button_State_Flag_Leave, xc.RGBA(137, 140, 151, 255))
	btnDown.AddBkFill(炫彩常量类.Button_State_Flag_Stay, xc.RGBA(255, 135, 250, 255))
	btnDown.AddBkFill(炫彩常量类.Button_State_Flag_Down, xc.RGBA(255, 75, 250, 255))
	// 因为上下按钮背景改变了, 你可以自己准备图片设置到按钮上去
	// btnDown.SetIcon()  或  btnDown.AddBkImage()

	// 注册滚动条元素滚动事件
	bar1.Event_SBAR_SCROLL1(SBAR_SCROLL1)
	bar2.Event_SBAR_SCROLL1(SBAR_SCROLL1)

	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}

// 滚动条元素滚动事件
func SBAR_SCROLL1(hEle int, pos int32, pbHandled *bool) int {
	fmt.Println(pos)
	// 为了鼠标滚轮滚动和点击两端按钮实时显示效果而刷新
	xc.XEle_Redraw(hEle, false)
	return 0
}
