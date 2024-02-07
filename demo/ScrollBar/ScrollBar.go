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
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "ScrollBar", 0, 炫彩常量类.Window_Style_Default)

	// 创建滚动条
	bar1 := 炫彩组件类.X创建滚动条(12, 33, 300, 20, w.Handle)
	bar2 := 炫彩组件类.X创建滚动条(330, 33, 20, 240, w.Handle)

	// 设置为垂直滚动条
	bar2.X置水平(false)
	// 添加背景
	bar2.X添加背景填充(炫彩常量类.Element_State_Flag_Leave, 炫彩基类.RGBA(247, 248, 250, 255))

	// 获取滑块按钮
	btnSlider := 炫彩组件类.X创建按钮并按句柄(bar2.X取滑块())
	btnSlider.X添加背景填充(炫彩常量类.Button_State_Flag_Leave, 炫彩基类.RGBA(221, 221, 223, 255))
	btnSlider.X添加背景填充(炫彩常量类.Button_State_Flag_Stay, 炫彩基类.RGBA(202, 202, 204, 255))
	btnSlider.X添加背景填充(炫彩常量类.Button_State_Flag_Down, 炫彩基类.RGBA(202, 202, 204, 255))
	// 获取滚动条上按钮
	btnUp := 炫彩组件类.X创建按钮并按句柄(bar2.X取上按钮())
	btnUp.X添加背景填充(炫彩常量类.Button_State_Flag_Leave, 炫彩基类.RGBA(137, 140, 151, 255))
	btnUp.X添加背景填充(炫彩常量类.Button_State_Flag_Stay, 炫彩基类.RGBA(255, 135, 250, 255))
	btnUp.X添加背景填充(炫彩常量类.Button_State_Flag_Down, 炫彩基类.RGBA(255, 75, 250, 255))
	// 获取滚动条下按钮
	btnDown := 炫彩组件类.X创建按钮并按句柄(bar2.X取下按钮())
	btnDown.X添加背景填充(炫彩常量类.Button_State_Flag_Leave, 炫彩基类.RGBA(137, 140, 151, 255))
	btnDown.X添加背景填充(炫彩常量类.Button_State_Flag_Stay, 炫彩基类.RGBA(255, 135, 250, 255))
	btnDown.X添加背景填充(炫彩常量类.Button_State_Flag_Down, 炫彩基类.RGBA(255, 75, 250, 255))
	// 因为上下按钮背景改变了, 你可以自己准备图片设置到按钮上去
	// btnDown.SetIcon()  或  btnDown.AddBkImage()

	// 注册滚动条元素滚动事件
	bar1.X事件_滚动事件1(SBAR_SCROLL1)
	bar2.X事件_滚动事件1(SBAR_SCROLL1)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}

// 滚动条元素滚动事件
func SBAR_SCROLL1(hEle int, pos int32, pbHandled *bool) int {
	fmt.Println(pos)
	// 为了鼠标滚轮滚动和点击两端按钮实时显示效果而刷新
	炫彩基类.X元素_重绘(hEle, false)
	return 0
}
