// 动画特效展示
package main

import (
	_ "embed"

	"github.com/888go/xcgui/ani"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	a *炫彩App类.App
	w *炫彩窗口基类.Window

	top int32 = 35

	hSvg           int
	list_svg       []int
	list_animation []int
	list_xcgui     []int

	hLayout1 int
	hLayout2 int
	hLayout3 int
)

var (
	//go:embed svg/公益.svg
	svg1 string
	//go:embed svg/时间戳.svg
	svg2 string
	//go:embed svg/技术服务.svg
	svg3 string
	//go:embed svg/底层架构.svg
	svg4 string
	//go:embed svg/查验.svg
	svg5 string
	//go:embed svg/接口配置.svg
	svg6 string
	//go:embed svg/淘公仔文字.svg
	svg7 string

	//go:embed image/img-1.jpg
	img1 []byte
	//go:embed image/img-2.jpg
	img2 []byte
	//go:embed image/img-3.jpg
	img3 []byte
	//go:embed image/img-4.jpg
	img4 []byte
	//go:embed image/img-5.jpg
	img5 []byte
	//go:embed image/img-6.jpg
	img6 []byte

	svg11 = `<svg x="0" y="0" width="25" height="25" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#ee6362" /></svg>`
	svg12 = `<svg x="0" y="0" width="25" height="25" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#2cb0b2" /></svg>`
	svg13 = `<svg x="0" y="0" width="20" height="20" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#f00" /></svg>`
	svg14 = `<svg x="0" y="0" width="15" height="15" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#f00" /></svg>`
	svg15 = `<svg viewBox="0 0 200 200"><circle cx="100" cy="100" r="100" fill="#ff0" /></svg>`
)

func main() {
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	// a.ShowSvgFrame(true)
	a.X置绘制频率(10)
	// 创建窗口
	w = 炫彩窗口基类.X创建窗口(0, 0, 970, 650, "炫彩界面库-动画特效-SVG特效", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮, 注册按钮单击事件
	CreateButton("1.下落 缩放 缓动").X事件_被单击(OnBtnClick1)
	CreateButton("2.下落 呼吸SVG").X事件_被单击(OnBtnClick2)
	CreateButton("3.呼吸SVG").X事件_被单击(OnBtnClick3)
	CreateButton("4.不透明度SVG").X事件_被单击(OnBtnClick4)
	CreateButton("5.移动SVG").X事件_被单击(OnBtnClick5)
	CreateButton("6.形状文本").X事件_被单击(OnBtnClick6)
	CreateButton("7.按钮").X事件_被单击(OnBtnClick7)
	CreateButton("8.布局焦点展开").X事件_被单击(OnBtnClick8)
	CreateButton("9.图片切换").X事件_被单击(OnBtnClick9)
	CreateButton("10.图片切换2").X事件_被单击(OnBtnClick10)
	CreateButton("11.进度 等待").X事件_被单击(OnBtnClick11)
	CreateButton("12.旋转 移动").X事件_被单击(OnBtnClick12)
	CreateButton("13.旋转 摇摆").X事件_被单击(OnBtnClick13)
	CreateButton("14.旋转 移动 缩放").X事件_被单击(OnBtnClick14)
	CreateButton("15.旋转 开合效果").X事件_被单击(OnBtnClick15)
	CreateButton("16.颜色渐变").X事件_被单击(OnBtnClick16)
	CreateButton("17.缩放 位置").X事件_被单击(OnBtnClick17)
	CreateButton("18.按钮 宽度").X事件_被单击(OnBtnClick18)

	w.X线程_绘制消息(OnWndDrawWindow)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}

// 创建按钮

// ff:
// name:
func CreateButton(name string) *炫彩组件类.Button {
	btn := 炫彩组件类.X创建按钮(10, top, 110, 30, name, w.Handle)
	btn.X置文本对齐(炫彩常量类.TextAlignFlag_Left | 炫彩常量类.TextAlignFlag_Vcenter)
	btn.X置类型EX(炫彩常量类.Button_Type_Radio)
	btn.X置组ID(1)
	top += 31
	return btn
}

// ff:
func ReleaseAnimation() {
	for _, v := range list_animation {
		炫彩基类.X动画_释放(v, true)
	}

	for _, v := range list_svg {
		炫彩基类.SVG_释放引用计数(v)
	}

	for _, v := range list_xcgui {
		if 炫彩基类.X判断元素(v) {
			炫彩基类.X元素_销毁(v)
		} else if 炫彩基类.X判断形状对象(v) {
			炫彩基类.X形状_销毁(v)
		}
	}

	list_animation = list_animation[:0]
	list_svg = list_svg[:0]
	list_xcgui = list_xcgui[:0]
}

// ff:
// pbHandled:
// hDraw:
func OnWndDrawWindow(hDraw int, pbHandled *bool) int {
	*pbHandled = true
	w.X绘制(hDraw)

	if hSvg != 0 {
		炫彩基类.X绘制_SVG源(hDraw, hSvg)
	}

	for _, v := range list_svg {
		炫彩基类.X绘制_SVG源(hDraw, v)
	}

	return 0
}

// ff:
// pbHandled:
func OnBtnClick1(pbHandled *bool) int {
	var left int32 = 130
	top = 22
	ReleaseAnimation()

	// 加载svg图片
	list_svg = append(list_svg,
		炫彩基类.SVG_加载从字符串W(svg1),
		炫彩基类.SVG_加载从字符串W(svg2),
		炫彩基类.SVG_加载从字符串W(svg3),
		炫彩基类.SVG_加载从字符串W(svg4),
		炫彩基类.SVG_加载从字符串W(svg5),
		炫彩基类.SVG_加载从字符串W(svg6),
	)

	// 创建动画组
	hGroup := 炫彩基类.X动画组_创建(0)
	list_animation = append(list_animation, hGroup)
	炫彩基类.X动画_运行(hGroup, w.Handle)

	for k, v := range list_svg {
		// 设置svg图片大小和位置
		炫彩基类.SVG_置大小(v, 100, 100)
		炫彩基类.SVG_置偏移(v, left, top)

		// 创建动画序列
		hAnimation := 炫彩基类.X动画_创建动画序列(v, 0)
		// 将动画序列添加到动画组中
		炫彩基类.X动画组_添加项(hGroup, hAnimation)

		炫彩基类.X动画_移动(hAnimation, 500, float32(left), 22, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, false)
		炫彩基类.X动画_延迟(hAnimation, 500)

		炫彩基类.X动画_延迟(hAnimation, 100*float32(k))
		炫彩基类.X动画_透明度(hAnimation, 500, 0, 1, 0, false)

		炫彩基类.X动画_延迟(hAnimation, 500)

		炫彩基类.X动画_透明度(hAnimation, 500, 255, 1, 0, false)
		炫彩基类.X动画_延迟(hAnimation, 1000)

		炫彩基类.X动画_移动(hAnimation, 2000, float32(left), 500, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, false)
		炫彩基类.X动画_延迟(hAnimation, 1000)
		left += 130

		hAnimation = 炫彩基类.X动画_创建动画序列(v, 0)
		炫彩基类.X动画_延迟(hAnimation, 6000+float32(k)*200)
		炫彩基类.X动画_缩放(hAnimation, 1200, 2, 2, 1, 炫彩常量类.Ease_Flag_Cubic|炫彩常量类.Ease_Flag_In, true)

		炫彩基类.X动画组_添加项(hGroup, hAnimation)
	}

	return 0
}

// ff:
// pbHandled:
func OnBtnClick2(pbHandled *bool) int {
	var left int32 = 450
	top = 22
	ReleaseAnimation()

	// 加载svg图片
	list_svg = append(list_svg, 炫彩基类.SVG_加载从字符串W(svg1))
	// 设置svg图片大小和位置
	炫彩基类.SVG_置大小(list_svg[0], 100, 100)
	炫彩基类.SVG_置偏移(list_svg[0], left, top)

	// 创建动画组
	group := 炫彩动画类.X创建动画组(0)
	list_animation = append(list_animation, group.Handle)
	group.X运行(w.Handle)

	// 下落
	ani1 := 炫彩动画类.X创建动画序列(list_svg[0], 0)
	group.X添加项(ani1.Handle)
	ani1.X移动(2000, float32(left), 500, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, false)

	// 停留
	ani1.X延迟(2000)

	// 返回顶部
	ani1.X移动(500, float32(left), 22, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, false)

	// 缩放
	ani2 := 炫彩动画类.X创建动画序列(list_svg[0], 1)
	group.X添加项(ani2.Handle)

	ani2.X延迟(2000)
	ani2.X缩放(1000, 2, 2, 0, 炫彩常量类.Ease_Flag_Cubic|炫彩常量类.Ease_Flag_In, true)

	/* 以下是纯函数方式实现
		// 创建动画组
	   	hGroup := xc.XAnimaGroup_Create(0)
	   	list_animation = append(list_animation, hGroup)
	   	xc.XAnima_Run(hGroup, w.Handle)

	   	// 下落
	   	hAnimation := xc.XAnima_Create(list_svg[0], 0)
	   	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	   	xc.XAnima_Move(hAnimation, 2000, float32(left), 500, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)

	   	// 停留
	   	xc.XAnima_Delay(hAnimation, 2000)

	   	// 返回顶部
	   	xc.XAnima_Move(hAnimation, 500, float32(left), 22, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)

	   	// 缩放
	   	hAnimation = xc.XAnima_Create(list_svg[0], 1)
	   	xc.XAnimaGroup_AddItem(hGroup, hAnimation)

	   	xc.XAnima_Delay(hAnimation, 2000)
	   	xc.XAnima_Scale(hAnimation, 1000, 2, 2, 0, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, true)
	*/

	return 0
}

// ff:
// pbHandled:
func OnBtnClick3(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 300
	top = 150

	// 加载svg图片
	list_svg = append(list_svg, 炫彩基类.SVG_加载从字符串W(svg1))
	// 设置svg图片大小和位置
	炫彩基类.SVG_置大小(list_svg[0], 300, 300)
	炫彩基类.SVG_置偏移(list_svg[0], left, top)

	// 创建动画序列
	ani1 := 炫彩动画类.X创建动画序列(list_svg[0], 1)
	list_animation = append(list_animation, ani1.Handle)

	// 缩放
	ani1.X缩放(1500, 2, 2, 0, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_In, true)
	ani1.X运行(w.Handle)

	/* 以下是纯函数方式实现
		// 创建动画序列
	   	hAnimation := xc.XAnima_Create(list_svg[0], 1)
	   	list_animation = append(list_animation, hAnimation)

	   	// 缩放
	   	xc.XAnima_Scale(hAnimation, 1500, 2, 2, 0, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, true)
	   	xc.XAnima_Run(hAnimation, w.Handle)
	*/

	return 0
}

// ff:
// pbHandled:
func OnBtnClick4(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 200
	top = 30

	// 加载svg图片
	list_svg = append(list_svg,
		炫彩基类.SVG_加载从字符串W(svg1),
		炫彩基类.SVG_加载从字符串W(svg1),
		炫彩基类.SVG_加载从字符串W(svg1),
	)

	// 设置svg图片大小和位置
	for k, v := range list_svg {
		炫彩基类.SVG_置大小(v, 100, 100)
		炫彩基类.SVG_置偏移(v, left+int32(k)*100, top)
	}

	// 创建动画序列
	hAnimation := 炫彩基类.X动画_创建动画序列(list_svg[0], 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_透明度EX(hAnimation, 3000, 0, 255, 1, 0, false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	hAnimation = 炫彩基类.X动画_创建动画序列(list_svg[1], 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_透明度(hAnimation, 3000, 0, 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	hAnimation = 炫彩基类.X动画_创建动画序列(list_svg[2], 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_透明度(hAnimation, 3000, 0, 0, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = 100
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)

	炫彩基类.X动画_透明度(hAnimation, 3000, 0, 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top += 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)

	炫彩基类.X动画_透明度EX(hAnimation, 3000, 255, 50, 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top += 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)

	炫彩基类.X动画_透明度EX(hAnimation, 3000, 255, 50, 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick5(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	top = 30

	// 加载svg图片
	list_svg = append(list_svg,
		炫彩基类.SVG_加载从字符串W(svg1),
		炫彩基类.SVG_加载从字符串W(svg2),
		炫彩基类.SVG_加载从字符串W(svg3),
	)

	// 设置svg图片大小和位置
	for k, v := range list_svg {
		炫彩基类.SVG_置大小(v, 100, 100)
		炫彩基类.SVG_置偏移(v, left, top+int32(k)*100)
	}
	top = 22

	// 循环
	ani1 := 炫彩动画类.X创建动画序列(list_svg[0], 1)
	list_animation = append(list_animation, ani1.Handle)
	ani1.X运行(w.Handle)
	ani1.X移动(2000, 750, float32(top), 10, 0, true)
	top += 100

	// 一次, 往返
	ani2 := 炫彩动画类.X创建动画序列(list_svg[1], 1)
	list_animation = append(list_animation, ani2.Handle)
	ani2.X运行(w.Handle)
	ani2.X移动(2000, 750, float32(top), 1, 0, true)
	top += 100

	// 一次, 不往返
	ani3 := 炫彩动画类.X创建动画序列(list_svg[2], 1)
	list_animation = append(list_animation, ani3.Handle)
	ani3.X运行(w.Handle)
	ani3.X移动(2000, 750, float32(top), 1, 0, false)
	top += 100

	/* 以下是纯函数方式实现
	// 循环
	hAnimation := xc.XAnima_Create(list_svg[0], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 10, 0, true)
	top += 100

	// 一次, 往返
	hAnimation = xc.XAnima_Create(list_svg[1], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 1, 0, true)
	top += 100

	// 一次, 不往返
	hAnimation = xc.XAnima_Create(list_svg[2], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 1, 0, false)
	*/

	return 0
}

// ff:
// pbHandled:
func OnBtnClick6(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 140
	top = 100

	// 创建形状文本
	hShapeText1 := 炫彩基类.X形状文本_创建(left, top, 100, 30, "循环滚动", w.Handle)
	top += 50
	hShapeText2 := 炫彩基类.X形状文本_创建(left, top, 100, 30, "往返滚动", w.Handle)
	top += 50
	hShapeText3 := 炫彩基类.X形状文本_创建(left, top, 100, 30, "移动到末尾", w.Handle)
	top += 50
	list_xcgui = append(list_xcgui,
		hShapeText1,
		hShapeText2,
		hShapeText3,
	)
	top = 100

	ani1 := 炫彩动画类.X创建动画序列(hShapeText1, 0)
	list_animation = append(list_animation, ani1.Handle)
	ani1.X运行(w.Handle)
	ani1.X移动(3000, 750, float32(top), 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, true)

	ani2 := 炫彩动画类.X创建动画序列(hShapeText2, 1)
	list_animation = append(list_animation, ani2.Handle)
	ani2.X运行(w.Handle)
	ani2.X移动(3000, 750, float32(top+50), 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, true)

	ani3 := 炫彩动画类.X创建动画序列(hShapeText3, 1)
	list_animation = append(list_animation, ani3.Handle)
	ani3.X运行(w.Handle)
	ani3.X移动(1500, 750, float32(top+100), 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, true)

	/* 	以下是纯函数方式实现
	hAnimation := xc.XAnima_Create(hShapeText1, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 3000, 750, float32(top), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)

	hAnimation = xc.XAnima_Create(hShapeText2, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 3000, 750, float32(top+50), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)

	hAnimation = xc.XAnima_Create(hShapeText3, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 1500, 750, float32(top+100), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
	*/

	return 0
}

// ff:
// pbHandled:
func OnBtnClick7(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 125
	top = 50

	group1 := 炫彩动画类.X创建动画组(0)
	list_animation = append(list_animation, group1.Handle)
	group1.X运行(w.Handle)
	for i := 0; i < 13; i++ {
		hButton := 炫彩基类.X按钮_创建(left, top, 60, 30, "透明度", w.Handle)
		list_xcgui = append(list_xcgui, hButton)

		hAnimation := 炫彩基类.X动画_创建动画序列(hButton, 0)
		group1.X添加项(hAnimation)

		炫彩基类.X动画_延迟(hAnimation, 500)

		炫彩基类.X动画_延迟(hAnimation, 100*float32(i))
		炫彩基类.X动画_透明度EX(hAnimation, 1200, 255, 20, 1, 0, true)
		left += 61
	}

	left = 125
	top = 100
	group2 := 炫彩动画类.X创建动画组(0)
	list_animation = append(list_animation, group2.Handle)
	group2.X运行(w.Handle)
	for i := 0; i < 7; i++ {
		hButton := 炫彩基类.X按钮_创建(left, top, 80, 30, "循环滚动", w.Handle)
		list_xcgui = append(list_xcgui, hButton)

		hAnimation := 炫彩基类.X动画_创建动画序列(hButton, 0)
		group2.X添加项(hAnimation)

		炫彩基类.X动画_移动(hAnimation, 500, float32(left), float32(top), 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, false)
		炫彩基类.X动画_延迟(hAnimation, 500)

		炫彩基类.X动画_延迟(hAnimation, 100*float32(i))
		炫彩基类.X动画_透明度EX(hAnimation, 500, 255, 0, 1, 0, false)

		炫彩基类.X动画_延迟(hAnimation, 500)

		炫彩基类.X动画_透明度EX(hAnimation, 500, 0, 255, 1, 0, false)
		炫彩基类.X动画_延迟(hAnimation, 1000)

		炫彩基类.X动画_移动(hAnimation, 2000, float32(left), 500, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, false)
		炫彩基类.X动画_延迟(hAnimation, 1000)

		hAnimation = 炫彩基类.X动画_创建动画序列(hButton, 1)
		炫彩基类.X动画组_添加项(group2.Handle, hAnimation)
		炫彩基类.X动画_延迟(hAnimation, 6000+float32(i)*200)
		炫彩基类.X动画_缩放(hAnimation, 1200, 1.5, 2, 1, 炫彩常量类.Ease_Flag_Cubic|炫彩常量类.Ease_Flag_In, true)

		left += 110
	}

	return 0
}

// ff:
// pbHandled:
func OnBtnClick8(pbHandled *bool) int {
	ReleaseAnimation()
	hLayout := 炫彩基类.X布局_创建(140, 100, 750, 100, w.Handle)
	炫彩基类.X布局盒子_置间距(hLayout, 20)

	for i := 0; i < 3; i++ {
		hLayout_ := 炫彩基类.X布局_创建(0, 0, 100, 100, hLayout)
		炫彩基类.X元素_置内填充大小(hLayout_, 10, 0, 10, 0)

		hShapeText := 炫彩基类.X形状文本_创建(0, 0, 100, 100, "炫彩界面库-www.xcgui.com-鼠标移动到上面查看", hLayout_)
		炫彩基类.X形状文本_置文本颜色(hShapeText, 炫彩基类.ABGR(255, 255, 255, 255))
		炫彩基类.X窗口组件_布局项_置宽度(hShapeText, 炫彩常量类.Layout_Size_Fill, 0)

		list_xcgui = append(list_xcgui, hLayout_)
		炫彩基类.X元素_启用鼠标穿透(hLayout_, false)
		炫彩基类.X窗口组件_布局项_置宽度(hLayout_, 炫彩常量类.Layout_Size_Weight, 100)

		炫彩基类.X废弃_XBkM_SetBkInfo(炫彩基类.X元素_取背景管理器(hLayout_), "{99:1.9.9;98:16(0);5:2(15)20(1)21(3)26(1)22(-7839744)23(255)9(5,5,5,5);}")
		炫彩基类.X元素_注册事件C1(hLayout_, 炫彩常量类.XE_MOUSESTAY, OnMouseStay8)
		炫彩基类.X元素_注册事件C1(hLayout_, 炫彩常量类.XE_MOUSELEAVE, OnMouseLeave8)

		switch i {
		case 0:
			hLayout1 = hLayout_
		case 1:
			hLayout2 = hLayout_
		case 2:
			hLayout3 = hLayout_
		}
	}

	炫彩基类.X窗口_调整布局(w.Handle)
	w.X重绘(false)
	return 0
}

// 鼠标进入事件8

// ff:
// pbHandled:
// hLayout:
func OnMouseStay8(hLayout int, pbHandled *bool) int {
	if hLayout1 != hLayout {
		炫彩基类.X元素_置透明度(hLayout1, 200)
	}

	if hLayout2 != hLayout {
		炫彩基类.X元素_置透明度(hLayout2, 200)
	}

	if hLayout3 != hLayout {
		炫彩基类.X元素_置透明度(hLayout3, 200)
	}

	hAnimation := 炫彩基类.X动画_创建动画序列(hLayout, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, w.Handle)
	炫彩基类.X动画_布局宽度(hAnimation, 300, 炫彩常量类.Layout_Size_Weight, 200, 1, 0, false)

	return 0
}

// 鼠标离开事件8

// ff:
// pbHandled:
// hEleStay:
// hLayout:
func OnMouseLeave8(hLayout, hEleStay int, pbHandled *bool) int {
	hAnimation := 炫彩基类.X动画_创建动画序列(hLayout, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, w.Handle)
	炫彩基类.X动画_布局宽度(hAnimation, 300, 炫彩常量类.Layout_Size_Weight, 100, 1, 0, false)

	炫彩基类.X元素_置透明度(hLayout1, 255)
	炫彩基类.X元素_置透明度(hLayout2, 255)
	炫彩基类.X元素_置透明度(hLayout3, 255)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick9(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	top = 50

	imgMap := map[int][]byte{
		1: img1,
		2: img2,
		3: img3,
		4: img4,
		5: img5,
		6: img6,
	}

	for i := 0; i < 3; i++ {
		hImage := 炫彩基类.X图片_加载从内存(imgMap[i*2+1])
		炫彩基类.X图片_置绘制类型(hImage, 炫彩常量类.Image_Draw_Type_Fixed_Ratio)

		hEle := 炫彩基类.X元素_创建(left, top, 212, 271, w.Handle)
		炫彩基类.X元素_添加背景图片(hEle, 炫彩常量类.Element_State_Flag_Leave, hImage)
		list_xcgui = append(list_xcgui, hEle)

		hImage2 := 炫彩基类.X图片_加载从内存(imgMap[i*2+2])
		炫彩基类.X图片_置绘制类型(hImage2, 炫彩常量类.Image_Draw_Type_Fixed_Ratio)

		hEle2 := 炫彩基类.X元素_创建(left, top, 212, 271, w.Handle)
		炫彩基类.X元素_添加背景图片(hEle2, 炫彩常量类.Element_State_Flag_Leave, hImage2)
		list_xcgui = append(list_xcgui, hEle2)

		hText := 炫彩基类.X形状文本_创建(left, top+280, 200, 30, "炫彩界面库-图片切换\r\n$66.66", w.Handle)
		炫彩基类.X形状文本_置文本颜色(hText, 炫彩基类.ABGR(80, 80, 80, 255))
		list_xcgui = append(list_xcgui, hText)

		炫彩基类.X元素_置用户数据(hEle, hEle2)
		炫彩基类.X元素_置用户数据(hEle2, hEle)
		炫彩基类.X窗口组件_显示(hEle2, false)

		炫彩基类.X元素_注册事件C1(hEle, 炫彩常量类.XE_MOUSESTAY, OnMouseStay9)
		炫彩基类.X元素_注册事件C1(hEle2, 炫彩常量类.XE_MOUSELEAVE, OnMouseLeave9)

		left += 212 + 10
	}

	w.X重绘(false)
	return 0
}

// 鼠标进入事件9

// ff:
// pbHandled:
// hEle:
func OnMouseStay9(hEle int, pbHandled *bool) int {
	hEle2 := 炫彩基类.X元素_取用户数据(hEle)
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			hObjectUI := 炫彩基类.X动画_取UI对象(list_animation[i])
			if hEle == hObjectUI || hEle2 == hObjectUI {
				炫彩基类.X动画_释放(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hAnimation := 炫彩基类.X动画_创建动画序列(hEle, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, hEle)
	炫彩基类.X动画_透明度EX(hAnimation, 1000, 255, 0, 1, 0, false)
	炫彩基类.X动画_显示(hAnimation, 0, false)

	炫彩基类.X元素_置透明度(hEle2, 0)
	炫彩基类.X窗口组件_显示(hEle2, true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hEle2, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, hEle2)
	炫彩基类.X动画_延迟(hAnimation, 500)
	炫彩基类.X动画_透明度EX(hAnimation, 1000, 0, 255, 1, 0, false)

	return 0
}

// 鼠标离开事件9

// ff:
// pbHandled:
// hEleStay:
// hEle2:
func OnMouseLeave9(hEle2, hEleStay int, pbHandled *bool) int {
	hEle := 炫彩基类.X元素_取用户数据(hEle2)
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			hObjectUI := 炫彩基类.X动画_取UI对象(list_animation[i])
			if hEle == hObjectUI || hEle2 == hObjectUI {
				炫彩基类.X动画_释放(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hAnimation := 炫彩基类.X动画_创建动画序列(hEle2, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, hEle2)
	炫彩基类.X动画_透明度EX(hAnimation, 1000, 255, 0, 1, 0, false)
	炫彩基类.X动画_显示(hAnimation, 0, false)

	炫彩基类.X元素_置透明度(hEle, 0)
	炫彩基类.X窗口组件_显示(hEle, true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hEle, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, hEle)
	炫彩基类.X动画_延迟(hAnimation, 500)
	炫彩基类.X动画_透明度EX(hAnimation, 1000, 0, 255, 1, 0, false)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick10(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	top = 50

	imgMap := map[int][]byte{
		1: img1,
		2: img2,
		3: img3,
		4: img4,
		5: img5,
		6: img6,
	}

	for i := 0; i < 3; i++ {
		hEle := 炫彩基类.X元素_创建(left, top, 212, 271, w.Handle)
		炫彩基类.X元素_启用绘制边框(hEle, false)
		list_xcgui = append(list_xcgui, hEle)

		hImage := 炫彩基类.X图片_加载从内存(imgMap[i*2+1])
		炫彩基类.X图片_置绘制类型(hImage, 炫彩常量类.Image_Draw_Type_Fixed_Ratio)

		hImage2 := 炫彩基类.X图片_加载从内存(imgMap[i*2+2])
		炫彩基类.X图片_置绘制类型(hImage2, 炫彩常量类.Image_Draw_Type_Fixed_Ratio)

		hShapePic := 炫彩基类.X形状图片_创建(0, 0, 212, 271, hEle)
		炫彩基类.X形状图片_置图片(hShapePic, hImage)

		hShapePic2 := 炫彩基类.X形状图片_创建(212+10, 0, 212, 271, hEle)
		炫彩基类.X形状图片_置图片(hShapePic2, hImage2)

		hText := 炫彩基类.X形状文本_创建(left, top+280, 200, 30, "炫彩界面库-图片切换\r\n$66.66", w.Handle)
		炫彩基类.X形状文本_置文本颜色(hText, 炫彩基类.ABGR(80, 80, 80, 255))
		list_xcgui = append(list_xcgui, hText)

		炫彩基类.X元素_注册事件C1(hEle, 炫彩常量类.XE_MOUSESTAY, OnMouseStay10)
		炫彩基类.X元素_注册事件C1(hEle, 炫彩常量类.XE_MOUSELEAVE, OnMouseLeave10)

		left += 212 + 10
	}

	w.X重绘(false)
	return 0
}

// 鼠标进入事件10

// ff:
// pbHandled:
// hEle:
func OnMouseStay10(hEle int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			if hEle == 炫彩基类.X动画_取UI对象(list_animation[i]) {
				炫彩基类.X动画_释放(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hPic := 炫彩基类.X元素_取子对象从索引(hEle, 0)

	hAnimation := 炫彩基类.X动画_创建动画序列(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, hEle)
	炫彩基类.X动画_移动(hAnimation, 500, -(212 + 10), 0, 1, 炫彩常量类.Ease_Flag_Cubic|炫彩常量类.Ease_Flag_In, false)

	hPic = 炫彩基类.X元素_取子对象从索引(hEle, 1)

	hAnimation = 炫彩基类.X动画_创建动画序列(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, hEle)
	炫彩基类.X动画_移动(hAnimation, 500, 0, 0, 1, 炫彩常量类.Ease_Flag_Cubic|炫彩常量类.Ease_Flag_In, false)

	return 0
}

// 鼠标离开事件10

// ff:
// pbHandled:
// hEleStay:
// hEle:
func OnMouseLeave10(hEle, hEleStay int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			if hEle == 炫彩基类.X动画_取UI对象(list_animation[i]) {
				炫彩基类.X动画_释放(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hPic := 炫彩基类.X元素_取子对象从索引(hEle, 0)

	hAnimation := 炫彩基类.X动画_创建动画序列(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, hEle)
	炫彩基类.X动画_移动(hAnimation, 500, 0, 0, 1, 炫彩常量类.Ease_Flag_Cubic|炫彩常量类.Ease_Flag_In, false)

	hPic = 炫彩基类.X元素_取子对象从索引(hEle, 1)

	hAnimation = 炫彩基类.X动画_创建动画序列(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_运行(hAnimation, hEle)
	炫彩基类.X动画_移动(hAnimation, 500, 212+10, 0, 1, 炫彩常量类.Ease_Flag_Cubic|炫彩常量类.Ease_Flag_In, false)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick11(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 160
	top = 80

	// 两个球型交替移动
	hSvg = 炫彩基类.SVG_加载从字符串W(svg11)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)

	hGroup := 炫彩基类.X动画组_创建(0)
	list_animation = append(list_animation, hGroup)
	炫彩基类.X动画_运行(hGroup, w.Handle)

	hAnimation := 炫彩基类.X动画_创建动画序列(hSvg, 1)
	炫彩基类.X动画组_添加项(hGroup, hAnimation)
	炫彩基类.X动画_移动(hAnimation, 1000, float32(left)+50, float32(top), 1, 炫彩常量类.Ease_Flag_Sine|炫彩常量类.Ease_Flag_InOut, false)
	炫彩基类.X动画_移动(hAnimation, 1000, float32(left), float32(top), 1, 炫彩常量类.Ease_Flag_Sine|炫彩常量类.Ease_Flag_InOut, false)

	hSvg = 炫彩基类.SVG_加载从字符串W(svg12)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left+50, top)

	hGroup = 炫彩基类.X动画组_创建(0)
	list_animation = append(list_animation, hGroup)
	炫彩基类.X动画_运行(hGroup, w.Handle)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 1)
	炫彩基类.X动画组_添加项(hGroup, hAnimation)
	炫彩基类.X动画_移动(hAnimation, 1000, float32(left), float32(top), 1, 炫彩常量类.Ease_Flag_Sine|炫彩常量类.Ease_Flag_InOut, false)
	炫彩基类.X动画_移动(hAnimation, 1000, float32(left)+50, float32(top), 1, 炫彩常量类.Ease_Flag_Sine|炫彩常量类.Ease_Flag_InOut, false)

	// 一排小球 缩放
	left = 350
	hGroup = 炫彩基类.X动画组_创建(0)
	list_animation = append(list_animation, hGroup)
	炫彩基类.X动画_运行(hGroup, w.Handle)

	for i := 0; i < 10; i++ {
		hSvg = 炫彩基类.SVG_加载从字符串W(svg13)
		list_svg = append(list_svg, hSvg)
		炫彩基类.SVG_置偏移(hSvg, left+int32(i)*50, top)

		hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
		炫彩基类.X动画组_添加项(hGroup, hAnimation)

		炫彩基类.X动画_延迟(hAnimation, float32(i)*200)
		炫彩基类.X动画_缩放(hAnimation, 1000, 2, 2, 1, 0, true)
	}

	// 一排小球 垂直缩放
	top = 150
	hGroup = 炫彩基类.X动画组_创建(0)
	list_animation = append(list_animation, hGroup)
	炫彩基类.X动画_运行(hGroup, w.Handle)

	for i := 0; i < 10; i++ {
		hSvg = 炫彩基类.SVG_加载从字符串W(svg13)
		list_svg = append(list_svg, hSvg)
		炫彩基类.SVG_置偏移(hSvg, left+int32(i)*50, top)

		hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
		炫彩基类.X动画组_添加项(hGroup, hAnimation)

		炫彩基类.X动画_延迟(hAnimation, float32(i)*200)
		炫彩基类.X动画_缩放(hAnimation, 1000, 1, 2, 1, 0, true)
	}

	// 一排小球 上下波浪
	left = 150
	top = 200
	for i := 0; i < 10; i++ {
		hSvg = 炫彩基类.SVG_加载从字符串W(svg13)
		list_svg = append(list_svg, hSvg)
		x := left + int32(i)*35
		炫彩基类.SVG_置偏移(hSvg, x, top)

		hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		炫彩基类.X动画_运行(hAnimation, w.Handle)

		炫彩基类.X动画项_启用完成释放(炫彩基类.X动画_延迟(hAnimation, float32(i)*100), true)
		炫彩基类.X动画_移动(hAnimation, 1200, float32(x), float32(top)+100, 1, 0, true)
	}

	left = 550
	for i := 0; i < 10; i++ {
		hSvg = 炫彩基类.SVG_加载从字符串W(svg13)
		list_svg = append(list_svg, hSvg)
		x := left + int32(i)*35
		炫彩基类.SVG_置偏移(hSvg, x, top)

		hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		炫彩基类.X动画_运行(hAnimation, w.Handle)

		炫彩基类.X动画项_启用完成释放(炫彩基类.X动画_延迟(hAnimation, float32(i)*150), true)
		炫彩基类.X动画_移动(hAnimation, 1000, float32(x), float32(top)+50, 1, 炫彩常量类.Ease_Flag_Sine|炫彩常量类.Ease_Flag_InOut, true)
	}

	// 一排小球 跳动
	left = 150
	top = 350
	for i := 0; i < 10; i++ {
		hSvg = 炫彩基类.SVG_加载从字符串W(svg13)
		list_svg = append(list_svg, hSvg)
		x := left + int32(i)*35
		炫彩基类.SVG_置偏移(hSvg, x, top)

		hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		炫彩基类.X动画_运行(hAnimation, w.Handle)

		炫彩基类.X动画项_启用完成释放(炫彩基类.X动画_延迟(hAnimation, float32(i)*200), true)
		炫彩基类.X动画_移动(hAnimation, 500, float32(x), float32(top)+50, 1, 炫彩常量类.Ease_Flag_Quint|炫彩常量类.Ease_Flag_Out, true)
		炫彩基类.X动画_延迟(hAnimation, 1700)
	}

	// 一排小球 移动
	left = 220
	top = 600
	for i := 0; i < 10; i++ {
		hSvg = 炫彩基类.SVG_加载从字符串W(svg14)
		list_svg = append(list_svg, hSvg)
		炫彩基类.SVG_置偏移(hSvg, 100-int32(i)*25, top)
		炫彩基类.SVG_置透明度(hSvg, 0)

		hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		炫彩基类.X动画_运行(hAnimation, w.Handle)

		炫彩基类.X动画项_启用完成释放(炫彩基类.X动画_延迟(hAnimation, float32(i)*100), true)
		炫彩基类.X动画_移动(hAnimation, 2000, 550-(float32(i)+1)*25, float32(top), 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_Out, false)
		炫彩基类.X动画_移动(hAnimation, 2000, 900-(float32(i)+1)*25, float32(top), 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_In, false)
		炫彩基类.X动画_移动(hAnimation, 0, 100-float32(i)*25, float32(top), 1, 0, false)
		炫彩基类.X动画_延迟(hAnimation, 500)

		hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		炫彩基类.X动画_运行(hAnimation, w.Handle)

		炫彩基类.X动画项_启用完成释放(炫彩基类.X动画_延迟(hAnimation, float32(i)*100), true)
		炫彩基类.X动画_透明度EX(hAnimation, 2000, 0, 255, 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_Out, false)
		炫彩基类.X动画_透明度EX(hAnimation, 2000, 255, 0, 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_In, false)
		炫彩基类.X动画_延迟(hAnimation, 500)
	}

	w.X重绘(false)
	return 0
}

// ff:
// pbHandled:
func OnBtnClick12(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 120
	top = 100

	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, 0)

	hAnimation := 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_旋转(hAnimation, 1700, 360, 1, 0, false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_移动(hAnimation, 3000, float32(left)+500, float32(top), 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	// 移动 往返旋转
	top = 350
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, -45)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 0, 0, 255), true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_旋转(hAnimation, 1500, 45, 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_In, false)
	炫彩基类.X动画_旋转(hAnimation, 1500, -45, 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_In, false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_移动(hAnimation, 3000, float32(left)+500, float32(top), 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick13(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 120
	top = 80

	// 自身 摇摆 往返
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, -45)

	hAnimation := 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_旋转(hAnimation, 1000, 45, 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	// 自身 旋转
	left = 500
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_旋转(hAnimation, 1000, 360, 1, 炫彩常量类.Ease_Flag_Expo|炫彩常量类.Ease_Flag_In, false)
	炫彩基类.X动画_旋转(hAnimation, 0, 0, 1, 炫彩常量类.Ease_Flag_Linear, false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	// 两个叠加 悬挂摆动
	left = 300
	top = 250
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, 45)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate := 炫彩基类.X动画_旋转(hAnimation, 3000, 100, 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_InOut, true)
	炫彩基类.X动画旋转_置中心(hRotate, float32(left)+10, float32(top)+50, false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, 45)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 0, 0, 255), true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = 炫彩基类.X动画_旋转(hAnimation, 3000, 100, 1, 炫彩常量类.Ease_Flag_Cubic|炫彩常量类.Ease_Flag_InOut, true)
	炫彩基类.X动画旋转_置中心(hRotate, float32(left)+10, float32(top)+50, false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	// 砍东西效果
	left = 500
	top = 400
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, -45)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 0, 0, 255), true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = 炫彩基类.X动画_旋转(hAnimation, 1500, 0, 1, 炫彩常量类.Ease_Flag_Expo|炫彩常量类.Ease_Flag_In, true)
	炫彩基类.X动画旋转_置中心(hRotate, float32(left), float32(top), false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick14(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 130
	top = 50

	// 加载svg, 设置大小和填充颜色
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置大小(hSvg, 50, 50)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 0, 0, 255), true)

	// 移动 360度旋转
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, 0)

	// 创建动画组
	hGroup := 炫彩基类.X动画组_创建(0)
	list_animation = append(list_animation, hGroup)

	// 旋转
	hAnimation := 炫彩基类.X动画_创建动画序列(hSvg, 0)
	炫彩基类.X动画组_添加项(hGroup, hAnimation)
	炫彩基类.X动画_旋转(hAnimation, 600, 360, 4, 0, false)

	// 缩放
	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	炫彩基类.X动画组_添加项(hGroup, hAnimation)
	炫彩基类.X动画_缩放(hAnimation, 2400, 7, 7, 1, 0, false)
	炫彩基类.X动画_延迟(hAnimation, 1000)
	炫彩基类.X动画_缩放(hAnimation, 1000, 1.0/7.0, 1.0/7.0, 1, 0, false)

	// 移动
	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	炫彩基类.X动画组_添加项(hGroup, hAnimation)
	炫彩基类.X动画_移动(hAnimation, 2400, float32(left)+500, float32(top)+300, 1, 0, false)
	炫彩基类.X动画_延迟(hAnimation, 1000)
	炫彩基类.X动画_移动(hAnimation, 1000, float32(left), float32(top), 1, 0, false)
	炫彩基类.X动画_运行(hGroup, w.Handle)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick15(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	top = 200
	var height int32 = 0
	var width int32 = 0

	// 砍东西效果

	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	height = 炫彩基类.SVG_取高度(hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, -45)

	hAnimation := 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate := 炫彩基类.X动画_旋转(hAnimation, 2000, 0, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, true)
	炫彩基类.X动画旋转_置中心(hRotate, float32(left), float32(top+height/2.0), false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	// 砍东西效果
	top = 300
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	height = 炫彩基类.SVG_取高度(hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, 45)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = 炫彩基类.X动画_旋转(hAnimation, 2000, 0, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, true)
	炫彩基类.X动画旋转_置中心(hRotate, float32(left), float32(top+height/2.0), false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	// 砍东西效果
	left = 500
	top = 200
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	width = 炫彩基类.SVG_取宽度(hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, 45)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 0, 0, 255), true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = 炫彩基类.X动画_旋转(hAnimation, 2000, 0, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, true)
	炫彩基类.X动画旋转_置中心(hRotate, float32(left+width), float32(top+height/2.0), false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	// 砍东西效果
	top = 300
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	width = 炫彩基类.SVG_取宽度(hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置旋转角度(hSvg, -45)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 0, 0, 255), true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = 炫彩基类.X动画_旋转(hAnimation, 2000, 0, 1, 炫彩常量类.Ease_Flag_Bounce|炫彩常量类.Ease_Flag_Out, true)
	炫彩基类.X动画旋转_置中心(hRotate, float32(left+width), float32(top+height/2.0), false)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick16(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	top = 50

	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 0, 0, 255), true)

	hAnimation := 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_颜色(hAnimation, 1500, 炫彩基类.ABGR(0, 0, 255, 255), 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = 225
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(0, 255, 0, 255), true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_颜色(hAnimation, 1500, 炫彩基类.ABGR(255, 0, 0, 255), 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = 400
	hSvg = 炫彩基类.SVG_加载从字符串W(svg7)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 255, 0, 255), true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_颜色(hAnimation, 1500, 炫彩基类.ABGR(0, 0, 255, 255), 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	hSvg = 炫彩基类.SVG_加载从字符串(svg15)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, 500, 300)
	炫彩基类.SVG_置用户填充颜色(hSvg, 炫彩基类.ABGR(255, 255, 0, 255), true)

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_颜色(hAnimation, 1500, 炫彩基类.ABGR(0, 255, 255, 255), 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	hFontx := 炫彩基类.X字体_创建EX("微软雅黑", 36, 炫彩常量类.FontStyle_Bold)
	hShapeText := 炫彩基类.X形状文本_创建(500, 100, 300, 50, "炫彩界面库", w.Handle)
	list_xcgui = append(list_xcgui, hShapeText)
	炫彩基类.X形状文本_置字体(hShapeText, hFontx)
	炫彩基类.X形状文本_置文本颜色(hShapeText, 炫彩基类.ABGR(255, 0, 0, 255))

	hAnimation = 炫彩基类.X动画_创建动画序列(hShapeText, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_颜色(hAnimation, 1500, 炫彩基类.ABGR(0, 0, 255, 255), 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	hShapeText = 炫彩基类.X形状文本_创建(500, 200, 100, 20, "炫彩界面库", w.Handle)
	list_xcgui = append(list_xcgui, hShapeText)

	hAnimation = 炫彩基类.X动画_创建动画序列(hShapeText, 0)
	list_animation = append(list_animation, hAnimation)
	炫彩基类.X动画_颜色(hAnimation, 1500, 炫彩基类.ABGR(0, 255, 0, 255), 1, 0, true)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick17(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	top = 50

	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_leftTop", w.Handle))

	hAnimation := 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale := 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_LeftTop)
	炫彩基类.X动画_运行(hAnimation, w.Handle)
	top = top + 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_left", w.Handle))

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_Left)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = top + 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_leftBottom", w.Handle))

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_LeftBottom)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = 50
	left = left + 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_top", w.Handle))

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_Top)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = top + 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_center", w.Handle))

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_Center)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = top + 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_bottom", w.Handle))

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_Bottom)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	left = left + 150
	top = 50
	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_rightTop", w.Handle))

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_RightTop)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = top + 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_right", w.Handle))

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_Right)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	top = top + 150
	hSvg = 炫彩基类.SVG_加载从字符串W(svg5)
	list_svg = append(list_svg, hSvg)
	炫彩基类.SVG_置偏移(hSvg, left, top)
	list_xcgui = append(list_xcgui, 炫彩基类.X形状文本_创建(left, top+65, 150, 20, "position_flag_rightBottom", w.Handle))

	hAnimation = 炫彩基类.X动画_创建动画序列(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = 炫彩基类.X动画_缩放(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_RightBottom)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	return 0
}

// ff:
// pbHandled:
func OnBtnClick18(pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	top = 50

	for i := 0; i < 5; i++ {
		hButton := 炫彩基类.X按钮_创建(left, top, 100, 50, "鼠标 停留 离开", w.Handle)
		list_xcgui = append(list_xcgui, hButton)
		炫彩基类.X元素_注册事件C1(hButton, 炫彩常量类.XE_MOUSESTAY, OnMouseStay18)
		炫彩基类.X元素_注册事件C1(hButton, 炫彩常量类.XE_MOUSELEAVE, OnMouseLeave18)
		top = top + 60
	}
	w.X重绘(false)

	return 0
}

// 鼠标进入事件18

// ff:
// pbHandled:
// hButton:
func OnMouseStay18(hButton int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			if hButton == 炫彩基类.X动画_取UI对象(list_animation[i]) {
				炫彩基类.X动画_释放(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hAnimation := 炫彩基类.X动画_创建动画序列(hButton, 1)
	list_animation = append(list_animation, hAnimation)
	hScale := 炫彩基类.X动画_缩放大小(hAnimation, 400, 200, 50, 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_Out, false)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_Left)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	return 0
}

// 鼠标离开事件18

// ff:
// pbHandled:
// hEleStay:
// hButton:
func OnMouseLeave18(hButton, hEleStay int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			if hButton == 炫彩基类.X动画_取UI对象(list_animation[i]) {
				炫彩基类.X动画_释放(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hAnimation := 炫彩基类.X动画_创建动画序列(hButton, 1)
	list_animation = append(list_animation, hAnimation)
	hScale := 炫彩基类.X动画_缩放大小(hAnimation, 400, 100, 50, 1, 炫彩常量类.Ease_Flag_Quad|炫彩常量类.Ease_Flag_In, false)
	炫彩基类.X动画缩放_置延伸位置(hScale, 炫彩常量类.Position_Flag_Left)
	炫彩基类.X动画_运行(hAnimation, w.Handle)

	return 0
}
