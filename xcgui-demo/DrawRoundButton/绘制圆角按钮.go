// 绘制圆角按钮
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/drawx"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.I初始化(true)
	// 2.创建窗口
	w := window.I窗口_创建(0, 0, 430, 300, "xc", 0, xcc.I常量_窗口样式_简单|xcc.I常量_窗口样式_按钮_关闭)

	// 创建一个按钮
	btn := widget.I按钮_创建(30, 50, 70, 30, "Button", w.I句柄)
	// 设置按钮字体颜色, 白色
	btn.I置文本颜色(xc.ABGR(255, 255, 255, 254))
	// 设置按钮圆角
	setBtnRound(btn, 14)

	// 3.显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	// 4.运行程序
	a.I运行()
	// 5.释放UI库
	a.I退出()
}

// 设置按钮圆角
func setBtnRound(btn *widget.Button, round int) {
	// 启用按钮背景透明
	btn.I启用背景透明(true)
	// 注册按钮绘制事件
	btn.I事件_元素绘制1(func(hEle int, hDraw int, pbHandled *bool) int {
		// 创建Draw对象
		draw := drawx.NewDrawByHandle(hDraw)
		// 启用平滑模式
		draw.EnableSmoothingMode(true)

		// 设置三种状态下的按钮背景色
		nState := xc.XBtn_GetStateEx(hEle)
		var bgcolor int
		switch nState {
		case xcc.I常量_按钮状态_离开:
			bgcolor = xc.ABGR(1, 162, 232, 254)
		case xcc.I常量_按钮状态_停留:
			bgcolor = xc.ABGR(1, 182, 252, 254)
		case xcc.I常量_按钮状态_按下:
			bgcolor = xc.ABGR(1, 122, 192, 254)
		case xcc.I常量_按钮状态_禁用:
			bgcolor = xc.ABGR(211, 215, 212, 255)
		}
		// 设置画刷颜色
		draw.SetBrushColor(bgcolor)

		// 绘制填充圆角矩形
		rc := xc.RECT{}
		rc.Right = int32(xc.XEle_GetWidth(hEle))
		rc.Bottom = int32(xc.XEle_GetHeight(hEle))
		draw.FillRoundRect(&rc, round, round)
		return 0
	})
}
