// 绘制圆角按钮
// 自己绘制要记一些api, 还是建议使用设计器来实现
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/drawx"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	// 2.创建窗口
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "绘制圆角按钮", 0, 炫彩常量类.Window_Style_Default)

	// 创建一个按钮
	btn := 炫彩组件类.X创建按钮((w.X取宽度()-100)/2, 100, 100, 30, "圆角按钮", w.Handle)
	// 设置按钮字体颜色, 白色
	btn.X置文本颜色(炫彩基类.ABGR(255, 255, 255, 255))
	// 设置按钮圆角
	setBtnRound(btn, 14)

	// 3.显示窗口
	w.X显示方式(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.X运行()
	// 5.释放UI库
	a.X退出()
}

// 设置按钮圆角
func setBtnRound(btn *炫彩组件类.Button, round int) {
	// 启用按钮背景透明
	btn.X启用背景透明(true)
	// 注册按钮绘制事件
	btn.X事件_绘制事件1(func(hEle int, hDraw int, pbHandled *bool) int {
		// 创建Draw对象
		draw := 炫彩绘制类.X创建并按图形绘制模块句柄(hDraw)
		// 启用平滑模式
		draw.X启用平滑模式(true)

		// 设置不同状态下的按钮背景色
		nState := 炫彩基类.X按钮_取状态EX(hEle)
		bgcolor := 炫彩基类.ABGR(1, 162, 232, 255) // 默认
		switch nState {
		case 炫彩常量类.Button_State_Stay:
			bgcolor = 炫彩基类.ABGR(1, 182, 252, 255)
		case 炫彩常量类.Button_State_Down:
			bgcolor = 炫彩基类.ABGR(1, 122, 192, 255)
		case 炫彩常量类.Button_State_Disable:
			bgcolor = 炫彩基类.ABGR(211, 215, 212, 255)
		}
		// 设置画刷颜色
		draw.X置画刷颜色(bgcolor)

		// 绘制填充圆角矩形
		var rc 炫彩基类.RECT
		炫彩基类.X元素_取客户区坐标(hEle, &rc)
		draw.X填充圆角矩形(&rc, round, round)
		return 0
	})
}
