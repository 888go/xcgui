// 列表视图
package main

import (
	_ "embed"
	"fmt"
	"github.com/888go/xcgui/xc"

	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/imagex"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

//go:embed res/1.png
var img1 []byte

func main() {
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 465, 400, "ListView", 0, 炫彩常量类.Window_Style_Default)

	// 创建ListView
	lv := 炫彩组件类.X创建列表视(10, 32, 445, 357, w.Handle)
	// 创建数据适配器
	lv.X创建数据适配器()

	// 添加分组
	group1 := lv.X组添加项文本("group1", -1)
	group2 := lv.X组添加项文本("group2", -1)
	// 图片加载从内存
	img := 炫彩图片类.X创建并按内存(img1)

	// 循环把图片加到分组里
	var index int
	for i := 0; i < 3; i++ {
		index = lv.X项添加图片(group1, img.Handle, -1)
		lv.X项置文本(group1, index, 1, fmt.Sprintf("group1-item%d", i))

		index = lv.X项添加图片(group2, img.Handle, -1)
		lv.X项置文本(group2, index, 1, fmt.Sprintf("group2-item%d", i))
	}

	炫彩组件类.X创建按钮(150, 0, 70, 30, "取选中项", w.Handle).X事件_被单击(func(pbHandled *bool) int {
		n := lv.X取选择项数量()
		fmt.Println("个数:", n)
		if n == 0 {
			return 0
		}

		var slice []炫彩基类.ListView_Item_Id_
		lv.X取选择项全部(&slice, n)
		for _, item := range slice {
			fmt.Println(item)
		}

		var a, b int32
		lv.X取选择项(&a, &b)
		fmt.Println("GetSelectItem:", a, b)
		return 0
	})

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
