// 全部缓动类型
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/drawx"
	"e.coding.net/gogit/go/xcgui/ease"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
	"time"
)

var (
	w *window.Window

	m_easeFlag   = xcc.I常量_缓动类型_快到慢 // 缓动方式
	m_easeType   = 11               // 缓动类型
	m_pos        = 0                // 当前位置
	m_time       = 60               // 缓动点数量
	m_time_pos   = 0                // 当前点
	m_rect       xc.RECT            // 窗口客户区坐标
	m_windowType = 2                // 窗口水平或垂直缓动
)

func main() {
	a := app.I初始化(true)
	// a.I置绘制频率(10)
	w = window.I窗口_创建(0, 0, 700, 450, "炫彩缓动测试", 0, xcc.I常量_窗口样式_默认)

	left := 30
	top := 35
	CreateButton(2, 11, left, top, 100, "Linear")
	left += 105
	CreateButton(2, 12, left, top, 100, "Quadratic")
	left += 105
	CreateButton(2, 13, left, top, 100, "Cubic")
	left += 105
	CreateButton(2, 14, left, top, 100, "Quartic")
	left += 105
	CreateButton(2, 15, left, top, 100, "Quintic")
	left += 105

	left = 30
	top += 30
	CreateButton(2, 16, left, top, 100, "Sinusoidal")
	left += 105
	CreateButton(2, 17, left, top, 100, "Exponential")
	left += 105
	CreateButton(2, 18, left, top, 100, "Circular")
	left += 105

	left = 30
	top += 30
	CreateButton(2, 19, left, top, 100, "Elastic")
	left += 105
	CreateButton(2, 20, left, top, 100, "Back")
	left += 105
	CreateButton(2, 21, left, top, 100, "Bounce")
	left += 105

	left = 30
	top += 40
	CreateButton(1, 0, left, top, 100, "easeIn")
	left += 105
	CreateButton(1, 1, left, top, 100, "easeOut")
	left += 105
	CreateButton(1, 2, left, top, 100, "easeInOut")
	left += 105

	btn := widget.I按钮_创建(445, top, 100, 24, "快速", w.I句柄)
	btn.I置类型扩展(xcc.I常量_对象扩展类型_复选按钮)
	btn.I置选中(true)
	btn.I事件_按钮选中(OnBtnCheckSlow)

	btn = widget.I按钮_创建(445, 65, 100, 24, "从左向右", w.I句柄)
	btn.I置类型扩展(xcc.I常量_对象扩展类型_单选按钮)
	btn.I置组ID(3)
	btn.I事件_按钮选中(OnBtnCheck_LeftToRight)

	btn = widget.I按钮_创建(445, 92, 100, 24, "从上向下", w.I句柄)
	btn.I置类型扩展(xcc.I常量_对象扩展类型_单选按钮)
	btn.I置组ID(3)
	btn.I置选中(true)
	btn.I事件_按钮选中(OnBtnCheck_TopToBottom)

	btn = widget.I按钮_创建(550, 65, 110, 50, "I运行 - 窗口缓动", w.I句柄)
	btn.I事件_按钮被单击(OnBtnStartWindow)

	btn = widget.I按钮_创建(550, 120, 110, 50, "I运行 - 缓动曲线", w.I句柄)
	btn.I事件_按钮被单击(OnBtnStart)

	// 窗口绘制事件
	w.I事件_窗口绘制(OnDrawWindow)
	// 窗口调整布局
	w.I调整布局()
	w.I显示(xcc.I常量_窗口形式_显示并激活)

	// 窗口第一次出现时的缓动
	time.AfterFunc(time.Millisecond*3, func() {
		// 获取窗口坐标
		var rect xc.RECT
		w.I取坐标(&rect)

		for i := 1; i <= 30; i++ {
			v := ease.Bounce(float32(i)/30.0, xcc.I常量_缓动类型_快到慢)
			y := int(v * float32(rect.Top))

			w.I移动窗口(int(rect.Left), y)
			time.Sleep(time.Millisecond * 10)
		}
	})

	a.I运行()
	a.I退出()
}

// 创建按钮
func CreateButton(nGroup, id, x, y, cx int, title string) {
	btn := widget.I按钮_创建(x, y, cx, 22, title, w.I句柄)
	// 设置为单选按钮
	btn.I置类型扩展(xcc.I常量_对象扩展类型_单选按钮)
	// 设置按钮组id
	btn.I置组ID(nGroup)
	// 设置元素ID
	btn.I置ID(id)

	if id == 1 || id == 21 {
		btn.I置选中(true)
	}
	// 注册按钮选中事件
	btn.I事件_按钮选中1(OnButtonCheck)
}

// 按钮选中事件
func OnButtonCheck(hEle int, bCheck bool, pbHandled *bool) int {
	if !bCheck {
		return 0
	}
	id := xc.XWidget_GetID(hEle)

	if id <= 2 {
		m_easeFlag = xcc.I常量_缓动类型_(id)
	} else {
		m_easeType = id - 10
	}

	w.I重绘(false)
	return 0
}

// 快速
func OnBtnCheckSlow(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_time = 60
	} else {
		m_time = 120
	}
	return 0
}

// 从左向右
func OnBtnCheck_LeftToRight(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_windowType = 1
	}
	return 0
}

// 从上向下
func OnBtnCheck_TopToBottom(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_windowType = 2
	}
	return 0
}

// 窗口缓动
func OnBtnStartWindow(pbHandled *bool) int {
	var rect xc.RECT
	w.I取坐标(&rect)

	time2 := m_time / 2
	for i := 0; i <= time2; i++ {
		var v float32
		switch m_easeType {
		case 1:
			v = ease.Linear(float32(i) / float32(time2))
		case 2:
			v = ease.Quad(float32(i)/float32(time2), m_easeFlag)
		case 3:
			v = ease.Cubic(float32(i)/float32(time2), m_easeFlag)
		case 4:
			v = ease.Quart(float32(i)/float32(time2), m_easeFlag)
		case 5:
			v = ease.Quint(float32(i)/float32(time2), m_easeFlag)
		case 6:
			v = ease.Sine(float32(i)/float32(time2), m_easeFlag)
		case 7:
			v = ease.Expo(float32(i)/float32(time2), m_easeFlag)
		case 8:
			v = ease.Circ(float32(i)/float32(time2), m_easeFlag)
		case 9:
			v = ease.Elastic(float32(i)/float32(time2), m_easeFlag)
		case 10:
			v = ease.Back(float32(i)/float32(time2), m_easeFlag)
		case 11:
			v = ease.Bounce(float32(i)/float32(time2), m_easeFlag)
		}

		if m_windowType == 1 {
			x := int(v * float32(rect.Left))
			w.I移动窗口(x, int(rect.Top))
		} else {
			y := int(v * float32(rect.Top))
			w.I移动窗口(int(rect.Left), y)
		}
		time.Sleep(20 * time.Millisecond)
	}

	return 0
}

// 缓动曲线
func OnBtnStart(pbHandled *bool) int {
	var width float32 = 400.0
	for i := 0; i <= m_time; i++ {
		var v float32
		switch m_easeType {
		case 1:
			v = ease.Linear(float32(i) / float32(m_time))
		case 2:
			v = ease.Quad(float32(i)/float32(m_time), m_easeFlag)
		case 3:
			v = ease.Cubic(float32(i)/float32(m_time), m_easeFlag)
		case 4:
			v = ease.Quart(float32(i)/float32(m_time), m_easeFlag)
		case 5:
			v = ease.Quint(float32(i)/float32(m_time), m_easeFlag)
		case 6:
			v = ease.Sine(float32(i)/float32(m_time), m_easeFlag)
		case 7:
			v = ease.Expo(float32(i)/float32(m_time), m_easeFlag)
		case 8:
			v = ease.Circ(float32(i)/float32(m_time), m_easeFlag)
		case 9:
			v = ease.Elastic(float32(i)/float32(m_time), m_easeFlag)
		case 10:
			v = ease.Back(float32(i)/float32(m_time), m_easeFlag)
		case 11:
			v = ease.Bounce(float32(i)/float32(m_time), m_easeFlag)
		}

		m_pos = int(v * width)
		m_time_pos = i
		time.Sleep(20 * time.Millisecond)

		rc := m_rect
		rc.Top = 170
		w.I重绘指定区域(&rc, true)
	}
	return 0
}

// 绘制
func OnDrawWindow(hDraw int, pbHandled *bool) int {
	*pbHandled = true

	var rect xc.RECT
	w.I取客户区坐标(&rect)
	draw := drawx.NewDrawByHandle(hDraw)

	draw.SetBrushColor(xc.ABGR(230, 230, 230, 255))
	draw.FillRect(&rect)
	m_rect = rect

	draw.SetBrushColor(xc.ABGR(200, 200, 200, 255))
	draw.DrawRect(&rect)

	draw.SetBrushColor(xc.ABGR(0, 0, 200, 255))
	draw.TextOutEx(260, 10, "炫彩界面库(XCGUI) - 缓动测试")

	var rc xc.RECT
	rc.Left = 150
	rc.Top = 190
	rc.Right = rc.Left + 400 + 30
	rc.Bottom = rc.Top + 50

	rcBorder := rc
	rcBorder.Left -= 2
	rcBorder.Top -= 2
	rcBorder.Right += 2
	rcBorder.Bottom += 2
	draw.SetBrushColor(xc.ABGR(0, 0, 200, 255))
	draw.DrawRect(&rcBorder)

	rcFill := rc
	rcFill.Left = rcFill.Left + int32(m_pos)
	rcFill.Right = rcFill.Left + 30
	draw.SetBrushColor(xc.ABGR(128, 0, 0, 255))
	draw.FillRect(&rcFill)

	var rcBorder_Line xc.RECT
	rcBorder_Line.Left = 150
	rcBorder_Line.Right = 150 + 400
	rcBorder_Line.Top = 255
	rcBorder_Line.Bottom = 255 + 180

	rcBorder = rcBorder_Line
	rcBorder.Right++
	rcBorder.Bottom++
	draw.SetBrushColor(xc.ABGR(180, 180, 180, 255))
	draw.DrawRect(&rcBorder)

	pts := make([]xc.POINT, 121)
	for t := 0; t <= m_time; t++ {
		var v float32
		switch m_easeType {
		case 1:
			v = ease.Linear(float32(t) / float32(m_time))
		case 2:
			v = ease.Quad(float32(t)/float32(m_time), m_easeFlag)
		case 3:
			v = ease.Cubic(float32(t)/float32(m_time), m_easeFlag)
		case 4:
			v = ease.Quart(float32(t)/float32(m_time), m_easeFlag)
		case 5:
			v = ease.Quint(float32(t)/float32(m_time), m_easeFlag)
		case 6:
			v = ease.Sine(float32(t)/float32(m_time), m_easeFlag)
		case 7:
			v = ease.Expo(float32(t)/float32(m_time), m_easeFlag)
		case 8:
			v = ease.Circ(float32(t)/float32(m_time), m_easeFlag)
		case 9:
			v = ease.Elastic(float32(t)/float32(m_time), m_easeFlag)
		case 10:
			v = ease.Back(float32(t)/float32(m_time), m_easeFlag)
		case 11:
			v = ease.Bounce(float32(t)/float32(m_time), m_easeFlag)
		}

		pts[t].X = rc.Left + int32(float32(t)*400.0/float32(m_time))
		pts[t].Y = rcBorder_Line.Bottom - int32(v*180.0)
	}

	draw.EnableSmoothingMode(true)
	draw.SetBrushColor(xc.ABGR(128, 0, 0, 255))

	left := int(rc.Left) + int(float32(m_time_pos)*400.0/float32(m_time))

	draw.DrawLine(left, int(rcBorder_Line.Top), left, int(rcBorder_Line.Bottom))
	draw.DrawCurve(pts, m_time+1, 0.5)
	return 0
}
