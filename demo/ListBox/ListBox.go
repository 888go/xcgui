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
	a := 炫彩App类.New(true)
	w := window.New(0, 0, 430, 500, "ListBox", 0, 炫彩常量类.Window_Style_Default)

	// 创建ListBox
	lb := 炫彩组件类.NewListBox(12, 33, 400, 450, w.Handle)

	// 创建数据适配器, 这个必须创建, 存储数据的
	lb.CreateAdapter()

	for i := 0; i < 15; i++ {
		// 添加行
		lb.AddItemText(fmt.Sprintf("item-%d", i))
	}

	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}
