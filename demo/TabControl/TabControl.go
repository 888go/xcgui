// 纯代码模拟选择夹(选项卡)切换页面
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

var (
	a *炫彩App类.App
	w *炫彩窗口基类.Window

	layoutTab *炫彩组件类.LayoutEle
	tabBtn1   *炫彩组件类.Button
	tabBtn2   *炫彩组件类.Button
	tabBtn3   *炫彩组件类.Button

	layoutBody  *炫彩组件类.LayoutEle
	layoutPage1 *炫彩组件类.LayoutEle
	layoutPage2 *炫彩组件类.LayoutEle
	layoutPage3 *炫彩组件类.LayoutEle
)

func main() {
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w = 炫彩窗口基类.X创建窗口(0, 0, 600, 400, "选择夹切换页面", 0, 炫彩常量类.Window_Style_Default)

	// 我是喜欢创建一个水平布局元素, tab按钮都放在里面
	layoutTab = 炫彩组件类.X创建布局(14, 35, 500, 30, w.Handle)
	// 放在水平布局元素中的组件, x, y绝对坐标是无效的
	tabBtn1 = 炫彩组件类.X创建按钮(0, 0, 100, 30, "页面1", layoutTab.Handle)
	tabBtn2 = 炫彩组件类.X创建按钮(0, 0, 100, 30, "页面2", layoutTab.Handle)
	tabBtn3 = 炫彩组件类.X创建按钮(0, 0, 100, 30, "页面3", layoutTab.Handle)

	// 设为单选类型按钮
	tabBtn1.X置类型EX(炫彩常量类.Button_Type_Radio)
	tabBtn2.X置类型EX(炫彩常量类.Button_Type_Radio)
	tabBtn3.X置类型EX(炫彩常量类.Button_Type_Radio)
	// 设置为默认按钮样式, 就不会是单选按钮样式了
	tabBtn1.X置样式(炫彩常量类.Button_Style_Default)
	tabBtn2.X置样式(炫彩常量类.Button_Style_Default)
	tabBtn3.X置样式(炫彩常量类.Button_Style_Default)
	// 默认选中第一个
	tabBtn1.X置选中(true)

	// 主体部分, 页面都放进这里面, 我是喜欢这样设计, 不是必须
	layoutBody = 炫彩组件类.X创建布局(14, 65, 500, 350, w.Handle)
	// 第一页
	layoutPage1 = 炫彩组件类.X创建布局(0, 0, 500, 350, layoutBody.Handle)
	// 第二页
	layoutPage2 = 炫彩组件类.X创建布局(0, 0, 500, 350, layoutBody.Handle)
	// 第三页
	layoutPage3 = 炫彩组件类.X创建布局(0, 0, 500, 350, layoutBody.Handle)
	// 只让第一页显示, 其他页都设为不显示
	layoutPage2.X显示(false)
	layoutPage3.X显示(false)

	// 给按钮绑定页面, 绑定后切换页面的原理就是: 你点哪个按钮就显示哪个页面
	tabBtn1.X置绑定元素(layoutPage1.Handle)
	tabBtn2.X置绑定元素(layoutPage2.Handle)
	tabBtn3.X置绑定元素(layoutPage3.Handle)

	// 页面中放置内容
	炫彩组件类.X创建形状文本(0, 0, 100, 30, "我是第一页", layoutPage1.Handle)
	炫彩组件类.X创建形状文本(0, 0, 100, 30, "我是第二页", layoutPage2.Handle)
	炫彩组件类.X创建形状文本(0, 0, 100, 30, "我是第三页", layoutPage3.Handle)

	w.X显示(true)
	a.X运行()
	a.X退出()
}
