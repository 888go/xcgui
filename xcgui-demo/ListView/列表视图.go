// 列表视图
package main

import (
	_ "embed"
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/imagex"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

//go:embed res/1.png
var img1 []byte

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 465, 400, "ListView", 0, xcc.I常量_窗口样式_默认)

	// 创建ListView
	lv := widget.I列表视图_创建(10, 32, 445, 357, w.I句柄)
	// 创建数据适配器
	lv.I创建数据适配器()

	// 添加分组
	group1 := lv.I组添加项文本("group1", -1)
	group2 := lv.I组添加项文本("group2", -1)
	// 图片加载从内存
	img := imagex.NewImage_LoadMemory(img1)

	// 循环把图片加到分组里
	var index int
	for i := 0; i < 3; i++ {
		index = lv.I项添加图片(group1, img.I句柄, -1)
		lv.I项置文本(group1, index, 1, fmt.Sprintf("group1-item%d", i))

		index = lv.I项添加图片(group2, img.I句柄, -1)
		lv.I项置文本(group2, index, 1, fmt.Sprintf("group2-item%d", i))
	}

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
