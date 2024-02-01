// 简单窗口.
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.I初始化(true)
	// 2.创建窗口
	w := window.I窗口_创建(0, 0, 430, 300, "", 0, xcc.I常量_窗口样式_简单|xcc.I常量_窗口样式_按钮_关闭)
	// 设置窗口边框大小
	w.I置边大小(0, 30, 0, 0)
	// 设置窗口透明类型
	w.I置透明类型(xcc.I常量_窗口透明_阴影)
	// 设置窗口阴影
	w.I置阴影信息(8, 254, 10, false, 0)
	// 3.显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	// 4.运行程序
	a.I运行()
	// 5.释放UI库
	a.I退出()
}
