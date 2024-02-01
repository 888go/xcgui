// 1. 给窗口添加背景色
// 2. 给按钮加上三种状态下的图片
package main

import (
	_ "embed"
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/imagex"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var (
	//go:embed res/button_min.png
	img1 []byte
	//go:embed res/button_close.png
	img2 []byte
)

func main() {
	// 1.初始化UI库
	a := app.I初始化(true)
	// 2.创建窗口
	w := window.I窗口_创建(0, 0, 465, 300, "", 0, xcc.I常量_窗口样式_简单|xcc.I常量_窗口样式_标题|xcc.I常量_窗口样式_拖动窗口)
	// 设置窗口透明类型
	w.I置透明类型(xcc.I常量_窗口透明_阴影)
	// 设置窗口阴影
	w.I置阴影信息(8, 255, 10, false, 0)
	// 给整个窗口添加背景色
	w.I添加背景填充(xcc.Window_State_Flag_Leave, xc.ABGR(51, 57, 60, 254))

	// 创建最小化按钮
	btn_Min := widget.I按钮_创建(397, 8, 30, 30, "", w.I句柄)
	btn_Min.I置类型扩展(xcc.I常量_对象扩展类型_窗口最小化按钮)
	// 创建结束按钮
	btn_Close := widget.I按钮_创建(427, 8, 30, 30, "", w.I句柄)
	btn_Close.I置类型扩展(xcc.I常量_对象扩展类型_窗口关闭按钮)

	// 给按钮加上三种状态下的图片
	setBtnImg(btn_Min, img1)
	setBtnImg(btn_Close, img2)

	// 3.显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	// 4.运行程序
	a.I运行()
	// 5.释放UI库
	a.I退出()
}

// 给按钮加上三态图片
func setBtnImg(btn *widget.Button, file []byte) {
	for i := 0; i < 3; i++ {
		x := i * 31
		// 图片_加载从内存, 指定区域位置及大小
		img := imagex.NewImage_LoadMemoryRect(file, x, 0, 30, 30)

		if img.I句柄 == 0 {
			fmt.Println("Error: hImg=", img.I句柄)
			continue
		}

		// 启用图片透明色
		img.EnableTranColor(true)
		// 添加背景图片
		switch i {
		case 0:
			btn.I添加背景图片(xcc.Button_State_Flag_Leave, img.I句柄)
		case 1:
			btn.I添加背景图片(xcc.Button_State_Flag_Stay, img.I句柄)
		case 2:
			btn.I添加背景图片(xcc.Button_State_Flag_Down, img.I句柄)
		}
		// 启用按钮背景透明
		btn.I启用背景透明(true)
	}
}
