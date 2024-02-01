// 调用UI设计器设计好的布局文件和资源文件
package main

import (
	_ "embed"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

//go:embed res/qqmusic.zip
var qqmusic []byte

func main() {
	a := app.I初始化(true)
	// 从内存zip中加载资源文件
	a.I加载资源文件内存ZIP(qqmusic, "resource.res", "")
	// 从内存zip中加载布局文件, 创建窗口对象
	w := window.I窗口_从内存zip布局文件创建(qqmusic, "main.xml", "", 0, 0)
	// 调整布局
	w.I调整布局()
	// 显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
