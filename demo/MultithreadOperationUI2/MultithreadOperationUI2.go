// 多协程操作UI, 方式2.
// 必须在UI线程操作UI, 否则随机发生崩溃.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	a   *炫彩App类.App
	w   *炫彩窗口基类.Window
	btn *炫彩组件类.Button
	ls  *炫彩组件类.List
)

func main() {
	// 置随机数种子
	rand.Seed(time.Now().UnixNano())

	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w = 炫彩窗口基类.X创建窗口(0, 0, 550, 300, "MultithreadOperationUI2", 0, 炫彩常量类.Window_Style_Default)

	btn = 炫彩组件类.X创建按钮(15, 33, 70, 24, "click", w.Handle)
	btn.X事件_被单击(onBnClick)

	ls = 炫彩组件类.X创建列表(10, 60, 530, 230, w.Handle)
	ls.X创建列表头数据适配器() // 创建表头数据适配器
	ls.X创建数据适配器(5)      // 创建数据适配器, 5列

	ls.X添加列文本(100, "name1", "column1")
	ls.X添加列文本(100, "name2", "column2")
	ls.X添加列文本(100, "name3", "column3")
	ls.X添加列文本(100, "name4", "column4")
	ls.X添加列文本(100, "name5", "column5")

	for i := 1; i < 10; i++ {
		id := strconv.Itoa(i)
		index := ls.X添加项文本("item" + id + "-column" + id)
		ls.X置项文本(index, 1, "item"+id+"-column"+id)
		ls.X置项文本(index, 2, "item"+id+"-column"+id)
		ls.X置项文本(index, 3, "item"+id+"-column"+id)
		ls.X置项文本(index, 4, "item"+id+"-column"+id)
	}

	a.X显示窗口并运行(w.Handle)
	a.X退出()
}

type updateList struct {
	Item int    // 项索引
	Col  int    // 列索引
	Text string // 项文本

	rwm sync.RWMutex // 保证同时只有1个在给List置入数据
	wg  sync.WaitGroup
}

// 在这里面写操作UI的代码, 是在ui线程操作ui

// ff:
// data:
func (l *updateList) UiThreadCallBack(data int) int {
	炫彩基类.X列表_置项文本(data, l.Item, l.Col, l.Text)
	return 0
}

// 按钮单击事件
func onBnClick(pbHandled *bool) int {
	if !btn.X判断启用() {
		return 0
	}
	btn.X启用(false)
	btn.X重绘(false)

	go func() {
		u := new(updateList)
		t := time.Now() // 记录开始的时间

		// 多协程操作列表框数据
		for i := 0; i < 2022; i++ {
			u.wg.Add(1)

			go func() {
				u.rwm.RLock()
				u.Item = rand.Intn(ls.X取项数量AD())
				u.Col = rand.Intn(ls.X取列数量())
				u.Text = strconv.Itoa(rand.Intn(1000) + 1000)
				// 这种方式能够传递更多的数据进回调函数
				a.X调用界面线程(u, ls.Handle) // 这样是在UI线程进行UI操作, 就不会崩溃了
				u.rwm.RUnlock()

				u.wg.Done()
			}()
		}
		u.wg.Wait()

		// 如果不需要传参数进回调函数, 也不需要返回值时可以调用a.CallUT(), 回调函数写法能简单些.
		a.X简易调用界面线程(func() {
			ls.X刷新项数据() // 刷新列表项数据
			ls.X重绘(false) // 列表重绘
			btn.X启用(true)
			btn.X重绘(true)
			w.X消息框("提示", fmt.Sprintf("全部执行完毕, 耗时: %v", time.Since(t)), 炫彩常量类.MessageBox_Flag_Ok, 炫彩常量类.Window_Style_Default)
		})
	}()
	return 0
}
