// 炫彩_调用界面线程, 在界面线程操作UI
package main

import (
	"fmt"
	"strconv"
	"time"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	a         *炫彩App类.App
	w         *炫彩窗口基类.Window
	btn       *炫彩组件类.Button
	radioBtn1 *炫彩组件类.Button
	radioBtn2 *炫彩组件类.Button

	t = 1 // 方式类型
)

func main() {
	a = 炫彩App类.X创建(true)
	w = 炫彩窗口基类.X创建窗口(0, 0, 550, 300, "ThreadOperationUI", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮
	btn = 炫彩组件类.X创建按钮(20, 35, 100, 30, "click", w.Handle)
	btn.X事件_被单击(onBnClick)

	// 单选按钮
	radioBtn1 = 炫彩组件类.X创建按钮(20, 70, 70, 30, "方式1", w.Handle)
	radioBtn2 = 炫彩组件类.X创建按钮(100, 70, 70, 30, "方式2", w.Handle)
	radioBtn1.X置类型EX(炫彩常量类.Button_Type_Radio)
	radioBtn2.X置类型EX(炫彩常量类.Button_Type_Radio)
	radioBtn1.X置选中(true) // 默认选中radioBtn1
	radioBtn1.X事件_选中1(onBnCheck)
	radioBtn2.X事件_选中1(onBnCheck)

	a.X显示窗口并运行(w.Handle)
	a.X退出()
}

func onBnClick(pbHandled *bool) int {
	// 禁用按钮
	btn.X启用(false)
	btn.X重绘(true)

	// 另起线程是为了不卡界面
	switch t {
	case 1:
		go updateBtn1() // 第一种方式: xc.XC_CallUiThreadEx
	case 2:
		go updateBtn2() // 第二种方式: xc.XC_CallUiThreader
	}
	return 0
}

// 第一种方式
func updateBtn1() {
	fmt.Println("使用方式1: CallUiThreadEx")
	for i := 0; i < 2010; i += 10 {
		// 如果直接在非界面线程内操作UI, 次数多了程序必将崩溃, 而且你不会知道它在什么时候崩溃.
		// 使用 a.CallUiThreadEx() 这样是在界面线程进行UI操作, 就不会崩溃了.
		a.X调用界面线程EX(func(data int) int {
			btn.X置文本(strconv.Itoa(data))
			btn.X置宽度(data / 5)
			w.X重绘(false)
			return 0
		}, i) // 把i传进回调函数了
		time.Sleep(time.Millisecond * 1)
	}

	// 解禁按钮.
	// 如果不需要传参数进回调函数, 也不需要返回值时可以调用 a.CallUT(), 回调函数写法能简单些.
	a.X简易调用界面线程(func() {
		btn.X启用(true)
		btn.X重绘(true)
	})
}

// 第2种方式, 这种方式明显可以传更多的参数, 完成更复杂的操作
type updateButton struct {
	HEle         int
	Text         string
	Width        int
	RedrawWindow bool
}

func (u updateButton) UiThreadCallBack(data int) int {
	炫彩基类.X按钮_置文本(u.HEle, u.Text)
	炫彩基类.X元素_置宽度(u.HEle, u.Width)
	w.X重绘(u.RedrawWindow)
	return 0
}

func updateBtn2() {
	fmt.Println("使用方式2: CallUiThreader")
	u := updateButton{
		HEle:         btn.Handle,
		RedrawWindow: false,
	}

	for i := 0; i < 2010; i += 10 {
		// 如果直接在非界面线程内操作UI, 次数多了程序必将崩溃, 而且你不会知道它在什么时候崩溃.
		// 使用 CallUiThreader 这样是在界面线程进行UI操作, 就不会崩溃了.
		u.Text = strconv.Itoa(i)
		u.Width = i / 5
		a.X调用界面线程(u, 0)
		time.Sleep(time.Millisecond * 1)
	}

	// 解禁按钮.
	// 如果不需要传参数进回调函数, 也不需要返回值时可以调用xc.XC_CallUT(), 回调函数写法能简单些.
	a.X简易调用界面线程(func() {
		btn.X启用(true)
		btn.X重绘(true)
	})
}

// 单选按钮被选择
func onBnCheck(hEle int, bCheck bool, pbHandled *bool) int {
	if bCheck {
		switch hEle {
		case radioBtn1.Handle:
			t = 1
		case radioBtn2.Handle:
			t = 2
		}
	}
	return 0
}
