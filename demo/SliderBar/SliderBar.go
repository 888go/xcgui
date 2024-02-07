// 滑块条.
package main

import (
	"fmt"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "SliderBar", 0, 炫彩常量类.Window_Style_Default)

	// 创建SliderBar
	sb := 炫彩组件类.X创建滑动条(12, 33, 300, 60, w.Handle)
	// 设置滑动范围
	sb.X置范围(10)

	// 设置滑块按钮高度和宽度
	sb.X置滑块高度(27)
	sb.X置滑块宽度(27)

	// 启用背景透明
	sb.X启用背景透明(true)

	// 注册滑块位置改变事件
	sb.X事件_滑块位置改变(func(pos int32, pbHandled *bool) int {
		fmt.Println(pos)
		return 0
	})

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
