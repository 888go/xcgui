// 列表框
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
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 500, "ListBox", 0, 炫彩常量类.Window_Style_Default)

	// 创建ListBox
	lb := 炫彩组件类.X创建列表框(12, 33, 400, 450, w.Handle)

	// 创建数据适配器, 这个必须创建, 存储数据的
	lb.X创建数据适配器()

	for i := 0; i < 15; i++ {
		// 添加行
		lb.X添加项文本(fmt.Sprintf("item-%d", i))
	}

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
