// 自定义窗口不同区域的背景颜色.
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/imagex"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	// 2.创建窗口
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "xcgui window", 0, 炫彩常量类.Window_Style_Default|炫彩常量类.Window_Style_Drag_Window)
	// 设置窗口边框大小
	w.X置边大小(1, 30, 1, 1)
	// 设置窗口图标
	a.X置窗口图标(炫彩图片类.X创建并按SVG字符串W(svgIcon).Handle)

	// 设置窗口主体颜色
	w.X添加背景填充(炫彩常量类.Window_State_Flag_Body_Leave, 炫彩基类.RGBA(248, 249, 251, 255))
	// 设置窗口边框颜色
	w.X添加背景填充(炫彩常量类.Window_State_Flag_Top_Leave, 炫彩基类.RGBA(39, 40, 46, 255))
	w.X添加背景填充(炫彩常量类.Window_State_Flag_Left_Leave, 炫彩基类.RGBA(177, 177, 177, 255))
	w.X添加背景填充(炫彩常量类.Window_State_Flag_Right_Leave, 炫彩基类.RGBA(177, 177, 177, 255))
	w.X添加背景填充(炫彩常量类.Window_State_Flag_Bottom_Leave, 炫彩基类.RGBA(177, 177, 177, 255))

	// 3.显示窗口
	w.X显示(true)
	// 4.运行程序
	a.X运行()
	// 5.释放UI库
	a.X退出()
}

const svgIcon = `<svg t="1690463587383" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6388" width="22" height="22"><path d="M511.97 318.12m-224.46 0a224.46 224.46 0 1 0 448.92 0 224.46 224.46 0 1 0-448.92 0Z" fill="#386BF3" p-id="6389"></path><path d="M735.51 705.31m-224.46 0a224.46 224.46 0 1 0 448.92 0 224.46 224.46 0 1 0-448.92 0Z" fill="#3D4265" p-id="6390"></path><path d="M288.42 705.31m-224.46 0a224.46 224.46 0 1 0 448.92 0 224.46 224.46 0 1 0-448.92 0Z" fill="#386BF3" p-id="6391"></path></svg>`
