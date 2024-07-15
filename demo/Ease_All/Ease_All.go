// 全部缓动类型
package main

import (
	"time"

	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/drawx"
	"github.com/888go/xcgui/ease"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	w *炫彩窗口基类.Window

	m_easeFlag           = 炫彩常量类.Ease_Type_Out // 缓动方式
	m_easeType   int32   = 11                // 缓动类型
	m_pos                = 0                 // 当前位置
	m_time               = 60                // 缓动点数量
	m_time_pos           = 0                 // 当前点
	m_rect       炫彩基类.RECT                     // 窗口客户区坐标
	m_windowType = 2                         // 窗口水平或垂直缓动
)

func main() {
	a := 炫彩App类.X创建(true)
	a.X置绘制频率(10)
	w = 炫彩窗口基类.X创建窗口(0, 0, 700, 450, "炫彩缓动测试", 0, 炫彩常量类.Window_Style_Default|炫彩常量类.Window_Style_Drag_Window)

	var left int32 = 30
	var top int32 = 35
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

	btn := 炫彩组件类.X创建按钮(445, top, 100, 24, "快速", w.Handle)
	btn.X置类型EX(炫彩常量类.Button_Type_Check)
	btn.X置选中(true)
	btn.X事件_选中(OnBtnCheckSlow)

	btn = 炫彩组件类.X创建按钮(445, 65, 100, 24, "从左向右", w.Handle)
	btn.X置类型EX(炫彩常量类.Button_Type_Radio)
	btn.X置组ID(3)
	btn.X事件_选中(OnBtnCheck_LeftToRight)

	btn = 炫彩组件类.X创建按钮(445, 92, 100, 24, "从上向下", w.Handle)
	btn.X置类型EX(炫彩常量类.Button_Type_Radio)
	btn.X置组ID(3)
	btn.X置选中(true)
	btn.X事件_选中(OnBtnCheck_TopToBottom)

	btn = 炫彩组件类.X创建按钮(550, 65, 110, 50, "Run - 窗口缓动", w.Handle)
	btn.X事件_被单击(OnBtnStartWindow)

	btn = 炫彩组件类.X创建按钮(550, 120, 110, 50, "Run - 缓动曲线", w.Handle)
	btn.X事件_被单击(OnBtnStart)

	// 窗口绘制事件
	w.X线程_绘制消息(OnDrawWindow)

	w.X显示方式(炫彩常量类.SW_SHOW)

	// 窗口第一次出现时的缓动
	time.AfterFunc(time.Millisecond*3, func() {
		// 获取窗口坐标
		var rect 炫彩基类.RECT
		w.X取坐标(&rect)

		for i := 1; i <= 30; i++ {
			v := 炫彩缓动类.X缓动弹跳(float32(i)/30.0, 炫彩常量类.Ease_Type_Out)
			y := int32(v * float32(rect.Top))

			w.X移动窗口(rect.Left, y)
			time.Sleep(time.Millisecond * 10)
		}
	})

	a.X运行()
	a.X退出()
}

// 创建按钮

// ff:
// title:
// cx:
// y:
// x:
// id:
// nGroup:
func CreateButton(nGroup, id, x, y, cx int32, title string) {
	btn := 炫彩组件类.X创建按钮(x, y, cx, 22, title, w.Handle)
	// 设置为单选按钮
	btn.X置类型EX(炫彩常量类.Button_Type_Radio)
	// 设置按钮组id
	btn.X置组ID(nGroup)
	// 设置元素ID
	btn.X置ID(id)

	if id == 1 || id == 21 {
		btn.X置选中(true)
	}
	// 注册按钮选中事件
	btn.X事件_选中1(OnButtonCheck)
}

// 按钮选中事件

// ff:
// pbHandled:
// bCheck:
// hEle:
func OnButtonCheck(hEle int, bCheck bool, pbHandled *bool) int {
	if !bCheck {
		return 0
	}
	id := 炫彩基类.X窗口组件_取ID(hEle)

	if id <= 2 {
		m_easeFlag = 炫彩常量类.Ease_Type_(id)
	} else {
		m_easeType = id - 10
	}

	w.X重绘(false)
	return 0
}

// 快速

// ff:
// pbHandled:
// bCheck:
func OnBtnCheckSlow(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_time = 60
	} else {
		m_time = 120
	}
	return 0
}

// 从左向右

// ff:
// pbHandled:
// bCheck:
func OnBtnCheck_LeftToRight(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_windowType = 1
	}
	return 0
}

// 从上向下

// ff:
// pbHandled:
// bCheck:
func OnBtnCheck_TopToBottom(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_windowType = 2
	}
	return 0
}

// 窗口缓动

// ff:
// pbHandled:
func OnBtnStartWindow(pbHandled *bool) int {
	var rect 炫彩基类.RECT
	w.X取坐标(&rect)

	time2 := m_time / 2
	for i := 0; i <= time2; i++ {
		var v float32
		switch m_easeType {
		case 1:
			v = 炫彩缓动类.X缓动线性(float32(i) / float32(time2))
		case 2:
			v = 炫彩缓动类.X缓动二次方曲线(float32(i)/float32(time2), m_easeFlag)
		case 3:
			v = 炫彩缓动类.X缓动三次方曲线(float32(i)/float32(time2), m_easeFlag)
		case 4:
			v = 炫彩缓动类.X缓动四次方曲线(float32(i)/float32(time2), m_easeFlag)
		case 5:
			v = 炫彩缓动类.X缓动五次方曲线(float32(i)/float32(time2), m_easeFlag)
		case 6:
			v = 炫彩缓动类.X缓动正弦曲线(float32(i)/float32(time2), m_easeFlag)
		case 7:
			v = 炫彩缓动类.X缓动突击曲线(float32(i)/float32(time2), m_easeFlag)
		case 8:
			v = 炫彩缓动类.X缓动圆环(float32(i)/float32(time2), m_easeFlag)
		case 9:
			v = 炫彩缓动类.X缓动强力回弹(float32(i)/float32(time2), m_easeFlag)
		case 10:
			v = 炫彩缓动类.X缓动回弹(float32(i)/float32(time2), m_easeFlag)
		case 11:
			v = 炫彩缓动类.X缓动弹跳(float32(i)/float32(time2), m_easeFlag)
		}

		if m_windowType == 1 {
			x := int32(v * float32(rect.Left))
			w.X移动窗口(x, rect.Top)
		} else {
			y := int32(v * float32(rect.Top))
			w.X移动窗口(rect.Left, y)
		}
		time.Sleep(20 * time.Millisecond)
	}
	return 0
}

// 缓动曲线

// ff:
// pbHandled:
func OnBtnStart(pbHandled *bool) int {
	var width float32 = 400.0
	for i := 0; i <= m_time; i++ {
		var v float32
		switch m_easeType {
		case 1:
			v = 炫彩缓动类.X缓动线性(float32(i) / float32(m_time))
		case 2:
			v = 炫彩缓动类.X缓动二次方曲线(float32(i)/float32(m_time), m_easeFlag)
		case 3:
			v = 炫彩缓动类.X缓动三次方曲线(float32(i)/float32(m_time), m_easeFlag)
		case 4:
			v = 炫彩缓动类.X缓动四次方曲线(float32(i)/float32(m_time), m_easeFlag)
		case 5:
			v = 炫彩缓动类.X缓动五次方曲线(float32(i)/float32(m_time), m_easeFlag)
		case 6:
			v = 炫彩缓动类.X缓动正弦曲线(float32(i)/float32(m_time), m_easeFlag)
		case 7:
			v = 炫彩缓动类.X缓动突击曲线(float32(i)/float32(m_time), m_easeFlag)
		case 8:
			v = 炫彩缓动类.X缓动圆环(float32(i)/float32(m_time), m_easeFlag)
		case 9:
			v = 炫彩缓动类.X缓动强力回弹(float32(i)/float32(m_time), m_easeFlag)
		case 10:
			v = 炫彩缓动类.X缓动回弹(float32(i)/float32(m_time), m_easeFlag)
		case 11:
			v = 炫彩缓动类.X缓动弹跳(float32(i)/float32(m_time), m_easeFlag)
		}

		m_pos = int(v * width)
		m_time_pos = i
		time.Sleep(20 * time.Millisecond)

		rc := m_rect
		rc.Top = 170
		w.X重绘指定区域(&rc, true)
	}
	return 0
}

// 绘制

// ff:
// pbHandled:
// hDraw:
func OnDrawWindow(hDraw int, pbHandled *bool) int {
	*pbHandled = true

	var rect 炫彩基类.RECT
	w.X取客户区坐标(&rect)
	draw := 炫彩绘制类.X创建并按图形绘制模块句柄(hDraw)

	draw.X置画刷颜色(炫彩基类.ABGR(230, 230, 230, 255))
	draw.X填充矩形(&rect)
	m_rect = rect

	draw.X置画刷颜色(炫彩基类.ABGR(200, 200, 200, 255))
	draw.X矩形边框(&rect)

	draw.X置画刷颜色(炫彩基类.ABGR(0, 0, 200, 255))
	draw.X文本输出EX(260, 10, "炫彩界面库(XCGUI) - 缓动测试")

	var rc 炫彩基类.RECT
	rc.Left = 150
	rc.Top = 190
	rc.Right = rc.Left + 400 + 30
	rc.Bottom = rc.Top + 50

	rcBorder := rc
	rcBorder.Left -= 2
	rcBorder.Top -= 2
	rcBorder.Right += 2
	rcBorder.Bottom += 2
	draw.X置画刷颜色(炫彩基类.ABGR(0, 0, 200, 255))
	draw.X矩形边框(&rcBorder)

	rcFill := rc
	rcFill.Left = rcFill.Left + int32(m_pos)
	rcFill.Right = rcFill.Left + 30
	draw.X置画刷颜色(炫彩基类.ABGR(128, 0, 0, 255))
	draw.X填充矩形(&rcFill)

	var rcBorder_Line 炫彩基类.RECT
	rcBorder_Line.Left = 150
	rcBorder_Line.Right = 150 + 400
	rcBorder_Line.Top = 255
	rcBorder_Line.Bottom = 255 + 180

	rcBorder = rcBorder_Line
	rcBorder.Right++
	rcBorder.Bottom++
	draw.X置画刷颜色(炫彩基类.ABGR(180, 180, 180, 255))
	draw.X矩形边框(&rcBorder)

	pts := make([]炫彩基类.POINT, 121)
	for t := 0; t <= m_time; t++ {
		var v float32
		switch m_easeType {
		case 1:
			v = 炫彩缓动类.X缓动线性(float32(t) / float32(m_time))
		case 2:
			v = 炫彩缓动类.X缓动二次方曲线(float32(t)/float32(m_time), m_easeFlag)
		case 3:
			v = 炫彩缓动类.X缓动三次方曲线(float32(t)/float32(m_time), m_easeFlag)
		case 4:
			v = 炫彩缓动类.X缓动四次方曲线(float32(t)/float32(m_time), m_easeFlag)
		case 5:
			v = 炫彩缓动类.X缓动五次方曲线(float32(t)/float32(m_time), m_easeFlag)
		case 6:
			v = 炫彩缓动类.X缓动正弦曲线(float32(t)/float32(m_time), m_easeFlag)
		case 7:
			v = 炫彩缓动类.X缓动突击曲线(float32(t)/float32(m_time), m_easeFlag)
		case 8:
			v = 炫彩缓动类.X缓动圆环(float32(t)/float32(m_time), m_easeFlag)
		case 9:
			v = 炫彩缓动类.X缓动强力回弹(float32(t)/float32(m_time), m_easeFlag)
		case 10:
			v = 炫彩缓动类.X缓动回弹(float32(t)/float32(m_time), m_easeFlag)
		case 11:
			v = 炫彩缓动类.X缓动弹跳(float32(t)/float32(m_time), m_easeFlag)
		}

		pts[t].X = rc.Left + int32(float32(t)*400.0/float32(m_time))
		pts[t].Y = rcBorder_Line.Bottom - int32(v*180.0)
	}

	draw.X启用平滑模式(true)
	draw.X置画刷颜色(炫彩基类.ABGR(128, 0, 0, 255))

	left := int(rc.Left) + int(float32(m_time_pos)*400.0/float32(m_time))

	draw.X线条(left, int(rcBorder_Line.Top), left, int(rcBorder_Line.Bottom))
	draw.X曲线(pts, m_time+1, 0.5)
	return 0
}
