// 编辑框
package main

import (
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 430, 300, "", 0, xcc.I常量_窗口样式_默认)

	// 1.普通编辑框
	edit := widget.I编辑框_创建(12, 35, 100, 30, w.I句柄)
	edit.I置文本颜色(xc.ABGR(236, 64, 122, 255))
	edit.I置文本("hello")

	// 2.密码输入框
	edit_pwd := widget.I编辑框_创建(12, 75, 100, 30, w.I句柄)
	edit_pwd.I启用密码(true)
	edit_pwd.I置文本("pwd")
	// 这个可以改变密码字符
	edit_pwd.I置密码字符('#')

	// 3.多行编辑框
	edit_MultiLine := widget.I编辑框_创建(12, 115, 300, 100, w.I句柄)
	edit_MultiLine.I启用多行(true)
	edit_MultiLine.I添加文本("你好, 世界")

	// 添加样式
	style1 := edit_MultiLine.I添加样式扩展("Arial", 12, xcc.I常量_字体样式_粗体, xc.ABGR(0, 191, 165, 255), true)
	// 添加带样式的文本
	edit_MultiLine.I添加文本扩展("\nhello world", style1)

	// 获取编辑框文本
	var s string
	edit_MultiLine.I取文本(&s, edit_MultiLine.I取内容长度()+1) // 长度必须+1
	fmt.Printf("s: %s\n", s)

	w.I显示(xcc.I常量_窗口形式_显示并激活)

	a.I运行()
	a.I退出()
}
