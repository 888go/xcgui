// 编辑框
package main

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 340, 280, "", 0, 炫彩常量类.Window_Style_Default)

	// 1.普通编辑框
	edit := 炫彩组件类.X创建编辑框(12, 35, 100, 30, w.Handle)
	edit.X置文本颜色(炫彩基类.ABGR(236, 64, 122, 255))
	edit.X置文本("hello")

	// 2.密码输入框
	editPwd := 炫彩组件类.X创建编辑框(12, 75, 100, 30, w.Handle)
	editPwd.X启用密码(true)
	editPwd.X置文本("pwd")
	// 这个可以改变密码字符
	editPwd.X置密码字符('#')

	// 3.多行编辑框
	editMultiLine := 炫彩组件类.X创建编辑框(12, 115, 300, 100, w.Handle)
	editMultiLine.X启用自动显示滚动条(true)
	editMultiLine.X启用多行(true)
	editMultiLine.X启用自动换行(true)
	editMultiLine.X添加文本("你好, 世界")
	// 添加样式
	style1 := editMultiLine.X添加样式EX("Arial", 14, 炫彩常量类.FontStyle_Bold, 炫彩基类.ABGR(0, 191, 165, 255), true)
	// 添加带样式的文本
	editMultiLine.X添加文本EX("\nhello world", style1)
	// 获取编辑框文本
	fmt.Println("获取编辑框文本:", editMultiLine.X取文本Ex())

	// 设置定时器, 循环获取多行编辑框鼠标选中的文本
	editMultiLine.X置炫彩定时器(111, 1000)
	editMultiLine.X事件_炫彩定时器(func(nTimerID int, pbHandled *bool) int {
		if nTimerID == 111 {
			text := editMultiLine.X取选择文本Ex()
			if text != "" {
				fmt.Println("选中的文本:", text)
			}
		}
		return 0
	})

	// 4.只能输入数字的编辑框
	editOnlyNumber := 炫彩组件类.X创建编辑框(12, 225, 100, 30, w.Handle)
	editOnlyNumber.X事件_CHAR(func(wParam, lParam uint, pbHandled *bool) int {
		fmt.Println(wParam)
		if wParam < 58 && wParam > 47 { // 0-9
			return 0
		}

		switch wParam { // 放行删除,复制,撤销,剪切,全选. 有其它需要的自己添加.
		case 8, 3, 26, 24, 1:
			return 0
		}

		*pbHandled = true // 其它的都拦截
		return 0
	})

	w.X显示(true)
	a.X运行()
	a.X退出()
}
