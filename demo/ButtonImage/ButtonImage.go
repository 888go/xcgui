// 1. 给窗口添加背景色
// 2. 给按钮加上三种状态下的图片
// 纯代码的要记好多api, 还是建议用设计器来做
package main

import (
	_ "embed"
	"fmt"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/imagex"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	//go:embed res/button_min.png
	img1 []byte
	//go:embed res/button_close.png
	img2 []byte
)

func main() {
	// 1.初始化UI库
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	// 2.创建窗口
	w := 炫彩窗口基类.X创建窗口(0, 0, 465, 300, "", 0, 炫彩常量类.Window_Style_Simple|炫彩常量类.Window_Style_Title|炫彩常量类.Window_Style_Drag_Window)
	// 设置窗口透明类型
	w.X置透明类型(炫彩常量类.Window_Transparent_Shadow)
	// 设置窗口阴影
	w.X置阴影信息(8, 255, 10, false, 0)
	// 给整个窗口添加背景色
	w.X添加背景填充(炫彩常量类.Window_State_Flag_Leave, 炫彩基类.ABGR(51, 57, 60, 254))

	// 创建最小化按钮
	btnMin := 炫彩组件类.X创建按钮(397, 8, 30, 30, "", w.Handle)
	btnMin.X置类型EX(炫彩常量类.Button_Type_Min)
	// 创建结束按钮
	btnClose := 炫彩组件类.X创建按钮(427, 8, 30, 30, "", w.Handle)
	btnClose.X置类型EX(炫彩常量类.Button_Type_Close)

	// 给按钮加上三种状态下的图片
	setBtnImg(btnMin, img1)
	setBtnImg(btnClose, img2)

	// 3.显示窗口
	w.X显示方式(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.X运行()
	// 5.释放UI库
	a.X退出()
}

// 给按钮加上三态图片
func setBtnImg(btn *炫彩组件类.Button, file []byte) {
	for i := 0; i < 3; i++ {
		x := int32(i * 31)
		// 图片_加载从内存, 指定区域位置及大小
		img := 炫彩图片类.X创建并按内存且指定区域(file, x, 0, 30, 30)

		if img.Handle == 0 {
			fmt.Println("Error: hImg=", img.Handle)
			continue
		}

		// 启用图片透明色
		img.X启用透明色(true)
		// 添加背景图片
		switch i {
		case 0:
			btn.X添加背景图片(炫彩常量类.Button_State_Flag_Leave, img.Handle)
		case 1:
			btn.X添加背景图片(炫彩常量类.Button_State_Flag_Stay, img.Handle)
		case 2:
			btn.X添加背景图片(炫彩常量类.Button_State_Flag_Down, img.Handle)
		}
		// 启用按钮背景透明
		btn.X启用背景透明(true)
	}
}
