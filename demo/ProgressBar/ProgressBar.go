// 进度条
package main

import (
	_ "embed"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/imagex"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	bar    *炫彩组件类.ProgressBar
	bar2   *炫彩组件类.ProgressBar
	btnAdd *炫彩组件类.Button
	btnSub *炫彩组件类.Button
)

//go:embed jindu.png
var img []byte

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 436, 450, "ProgressBar", 0, 炫彩常量类.Window_Style_Default)

	// 创建一个水平进度条
	bar = 炫彩组件类.X创建进度条(24, 60, 200, 10, w.Handle)
	// 设置进度条边框大小
	bar.X置边框大小(1, 1, 1, 1)
	// 设置进度条不显示进度文字
	bar.X启用进度文本(false)
	// 设置进度条最大值
	bar.X置范围(100)
	// 设置进度条进度
	bar.X置进度(40)
	// 置进度颜色
	bar.X置进度颜色(炫彩基类.RGBA(43, 170, 255, 255))
	// 置进度条背景颜色
	bar.X添加背景填充(炫彩常量类.Element_State_Flag_Leave, 炫彩基类.RGBA(221, 221, 223, 255))

	bar2 = 炫彩组件类.X创建进度条(24, 200, 24, 200, w.Handle)
	// 设置为垂直进度条
	bar2.X置水平(false)
	// 设置进度条边框大小
	bar2.X置边框大小(0, 0, 0, 0)
	// 不显示进度文本
	bar2.X启用进度文本(false)
	// 置进度图片
	bar2.X置进度图片(炫彩图片类.X创建并按内存且自适应(img, 0, 0, 0, 0).Handle)
	// 置进度条背景颜色
	bar2.X添加背景填充(炫彩常量类.Element_State_Flag_Leave, 炫彩基类.RGBA(221, 221, 223, 255))

	// 创建按钮_进度加
	btnAdd = 炫彩组件类.X创建按钮(238, 50, 70, 30, "+", w.Handle)
	btnAdd.X事件_被单击1(onBtnClick)
	// 创建按钮_进度减
	btnSub = 炫彩组件类.X创建按钮(318, 50, 70, 30, "-", w.Handle)
	btnSub.X事件_被单击1(onBtnClick)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}

// 事件_按钮被单击
func onBtnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btnAdd.Handle:
		bar.X置进度(bar.X取进度() + 10)
		bar2.X置进度(bar.X取进度() + 10)
	case btnSub.Handle:
		bar.X置进度(bar.X取进度() - 10)
		bar2.X置进度(bar.X取进度() - 10)
	}
	bar.X重绘(true)
	bar2.X重绘(true)
	return 0
}
