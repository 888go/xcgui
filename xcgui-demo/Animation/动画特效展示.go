// 动画特效展示
package main

import (
	_ "embed"

	"e.coding.net/gogit/go/xcgui/ani"
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var (
	a *app.App
	w *window.Window

	top = 35

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
	a = app.I初始化(true)
	// a.I显示边界(true)
	a.I置绘制频率(16)
	// 创建窗口
	w = window.I窗口_创建(0, 0, 970, 650, "炫彩界面库-动画特效-SVG特效", 0, xcc.I常量_窗口样式_默认)

	// 创建按钮, 注册按钮单击事件
	创建按钮("1.下落 缩放 缓动").I事件_按钮被单击(OnBtnClick1)
	创建按钮("2.下落 呼吸SVG").I事件_按钮被单击(OnBtnClick2)
	创建按钮("3.呼吸SVG").I事件_按钮被单击(OnBtnClick3)
	创建按钮("4.不透明度SVG").I事件_按钮被单击(OnBtnClick4)
	创建按钮("5.移动SVG").I事件_按钮被单击(OnBtnClick5)
	创建按钮("6.形状文本").I事件_按钮被单击(OnBtnClick6)
	创建按钮("7.按钮").I事件_按钮被单击(OnBtnClick7)
	创建按钮("8.布局焦点展开").I事件_按钮被单击(OnBtnClick8)
	创建按钮("9.图片切换").I事件_按钮被单击(OnBtnClick9)
	创建按钮("10.图片切换2").I事件_按钮被单击(OnBtnClick10)
	创建按钮("11.进度 等待").I事件_按钮被单击(OnBtnClick11)
	创建按钮("12.旋转 移动").I事件_按钮被单击(OnBtnClick12)
	创建按钮("13.旋转 摇摆").I事件_按钮被单击(OnBtnClick13)
	创建按钮("14.旋转 移动 缩放").I事件_按钮被单击(OnBtnClick14)
	创建按钮("15.旋转 开合效果").I事件_按钮被单击(OnBtnClick15)
	创建按钮("16.颜色渐变").I事件_按钮被单击(OnBtnClick16)
	创建按钮("17.缩放 位置").I事件_按钮被单击(OnBtnClick17)
	创建按钮("18.按钮 宽度").I事件_按钮被单击(OnBtnClick18)

	w.I事件_窗口绘制(OnWndDrawWindow)

	w.I显示(xcc.I常量_窗口形式_显示并激活)

	a.I运行()
	a.I退出()
}

// 创建按钮
func 创建按钮(name string) *widget.Button {
	btn := widget.I按钮_创建(10, top, 110, 30, name, w.I句柄)
	btn.I置文本对齐(xcc.I常量_文本对齐_左 | xcc.I常量_文本对齐_垂直居中)
	btn.I置类型扩展(xcc.I常量_对象扩展类型_单选按钮)
	btn.I置组ID(1)
	top += 31
	return btn
}

func ReleaseAnimation() {
	for _, v := range list_animation {
		xc.XAnima_Release(v, true)
	}

	for _, v := range list_svg {
		xc.XSvg_Release(v)
	}

	for _, v := range list_xcgui {
		if xc.XC_IsHELE(v) {
			xc.XEle_Destroy(v)
		} else if xc.XC_IsShape(v) {
			xc.XShape_Destroy(v)
		}
	}

	list_animation = list_animation[:0]
	list_svg = list_svg[:0]
	list_xcgui = list_xcgui[:0]
}

func OnWndDrawWindow(hDraw int, pbHandled *bool) int {
	*pbHandled = true
	w.I绘制(hDraw)

	if hSvg != 0 {
		xc.XDraw_DrawSvgSrc(hDraw, hSvg)
	}

	for _, v := range list_svg {
		xc.XDraw_DrawSvgSrc(hDraw, v)
	}

	return 0
}

func OnBtnClick1(pbHandled *bool) int {
	left := 130
	top = 22
	ReleaseAnimation()

	// 加载svg图片
	list_svg = append(list_svg,
		xc.XSvg_LoadStringW(svg1),
		xc.XSvg_LoadStringW(svg2),
		xc.XSvg_LoadStringW(svg3),
		xc.XSvg_LoadStringW(svg4),
		xc.XSvg_LoadStringW(svg5),
		xc.XSvg_LoadStringW(svg6),
	)

	// 创建动画组
	hGroup := xc.XAnimaGroup_Create(0)
	list_animation = append(list_animation, hGroup)
	xc.XAnima_Run(hGroup, w.I句柄)

	for k, v := range list_svg {
		// 设置svg图片大小和位置
		xc.XSvg_SetSize(v, 100, 100)
		xc.XSvg_SetPosition(v, left, top)

		// 创建动画序列
		hAnimation := xc.XAnima_Create(v, 0)
		// 将动画序列添加到动画组中
		xc.XAnimaGroup_AddItem(hGroup, hAnimation)

		xc.XAnima_Move(hAnimation, 500, float32(left), 22, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)
		xc.XAnima_Delay(hAnimation, 500)

		xc.XAnima_Delay(hAnimation, 100*float32(k))
		xc.XAnima_Alpha(hAnimation, 500, 0, 1, 0, false)

		xc.XAnima_Delay(hAnimation, 500)

		xc.XAnima_Alpha(hAnimation, 500, 255, 1, 0, false)
		xc.XAnima_Delay(hAnimation, 1000)

		xc.XAnima_Move(hAnimation, 2000, float32(left), 500, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)
		xc.XAnima_Delay(hAnimation, 1000)
		left += 130

		hAnimation = xc.XAnima_Create(v, 0)
		xc.XAnima_Delay(hAnimation, 6000+float32(k)*200)
		xc.XAnima_Scale(hAnimation, 1200, 2, 2, 1, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快, true)

		xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	}

	return 0
}

func OnBtnClick2(pbHandled *bool) int {
	left := 450
	top = 22
	ReleaseAnimation()

	// 加载svg图片
	list_svg = append(list_svg, xc.XSvg_LoadStringW(svg1))
	// 设置svg图片大小和位置
	xc.XSvg_SetSize(list_svg[0], 100, 100)
	xc.XSvg_SetPosition(list_svg[0], left, top)

	// 创建动画组
	group := ani.NewAnimaGroup(0)
	list_animation = append(list_animation, group.I句柄)
	group.Run(w.I句柄)

	// 下落
	ani1 := ani.NewAnima(list_svg[0], 0)
	group.AddItem(ani1.I句柄)
	ani1.Move(2000, float32(left), 500, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)

	// 停留
	ani1.Delay(2000)

	// 返回顶部
	ani1.Move(500, float32(left), 22, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)

	// 缩放
	ani2 := ani.NewAnima(list_svg[0], 1)
	group.AddItem(ani2.I句柄)

	ani2.Delay(2000)
	ani2.Scale(1000, 2, 2, 0, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快, true)

	/* 以下是纯函数方式实现
		// 创建动画组
	   	hGroup := xc.XAnimaGroup_Create(0)
	   	list_animation = append(list_animation, hGroup)
	   	xc.XAnima_Run(hGroup, w.I句柄)

	   	// 下落
	   	hAnimation := xc.XAnima_Create(list_svg[0], 0)
	   	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	   	xc.XAnima_Move(hAnimation, 2000, float32(left), 500, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)

	   	// 停留
	   	xc.XAnima_Delay(hAnimation, 2000)

	   	// 返回顶部
	   	xc.XAnima_Move(hAnimation, 500, float32(left), 22, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)

	   	// 缩放
	   	hAnimation = xc.XAnima_Create(list_svg[0], 1)
	   	xc.XAnimaGroup_AddItem(hGroup, hAnimation)

	   	xc.XAnima_Delay(hAnimation, 2000)
	   	xc.XAnima_Scale(hAnimation, 1000, 2, 2, 0, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快, true)
	*/

	return 0
}

func OnBtnClick3(pbHandled *bool) int {
	ReleaseAnimation()
	left := 300
	top = 150

	// 加载svg图片
	list_svg = append(list_svg, xc.XSvg_LoadStringW(svg1))
	// 设置svg图片大小和位置
	xc.XSvg_SetSize(list_svg[0], 300, 300)
	xc.XSvg_SetPosition(list_svg[0], left, top)

	// 创建动画序列
	ani1 := ani.NewAnima(list_svg[0], 1)
	list_animation = append(list_animation, ani1.I句柄)

	// 缩放
	ani1.Scale(1500, 2, 2, 0, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_慢到快, true)
	ani1.Run(w.I句柄)

	/* 以下是纯函数方式实现
		// 创建动画序列
	   	hAnimation := xc.XAnima_Create(list_svg[0], 1)
	   	list_animation = append(list_animation, hAnimation)

	   	// 缩放
	   	xc.XAnima_Scale(hAnimation, 1500, 2, 2, 0, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_慢到快, true)
	   	xc.XAnima_Run(hAnimation, w.I句柄)
	*/

	return 0
}

func OnBtnClick4(pbHandled *bool) int {
	ReleaseAnimation()
	left := 200
	top = 30

	// 加载svg图片
	list_svg = append(list_svg,
		xc.XSvg_LoadStringW(svg1),
		xc.XSvg_LoadStringW(svg1),
		xc.XSvg_LoadStringW(svg1),
	)

	// 设置svg图片大小和位置
	for k, v := range list_svg {
		xc.XSvg_SetSize(v, 100, 100)
		xc.XSvg_SetPosition(v, left+k*100, top)
	}

	// 创建动画序列
	hAnimation := xc.XAnima_Create(list_svg[0], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_AlphaEx(hAnimation, 3000, 0, 255, 1, 0, false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	hAnimation = xc.XAnima_Create(list_svg[1], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Alpha(hAnimation, 3000, 0, 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	hAnimation = xc.XAnima_Create(list_svg[2], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Alpha(hAnimation, 3000, 0, 0, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = 100
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)

	xc.XAnima_Alpha(hAnimation, 3000, 0, 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top += 150
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)

	xc.XAnima_AlphaEx(hAnimation, 3000, 255, 50, 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top += 150
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)

	xc.XAnima_AlphaEx(hAnimation, 3000, 255, 50, 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	return 0
}

func OnBtnClick5(pbHandled *bool) int {
	ReleaseAnimation()
	left := 150
	top = 30

	// 加载svg图片
	list_svg = append(list_svg,
		xc.XSvg_LoadStringW(svg1),
		xc.XSvg_LoadStringW(svg2),
		xc.XSvg_LoadStringW(svg3),
	)

	// 设置svg图片大小和位置
	for k, v := range list_svg {
		xc.XSvg_SetSize(v, 100, 100)
		xc.XSvg_SetPosition(v, left, top+k*100)
	}
	top = 22

	// 循环
	ani1 := ani.NewAnima(list_svg[0], 1)
	list_animation = append(list_animation, ani1.I句柄)
	ani1.Run(w.I句柄)
	ani1.Move(2000, 750, float32(top), 10, 0, true)
	top += 100

	// 一次, 往返
	ani2 := ani.NewAnima(list_svg[1], 1)
	list_animation = append(list_animation, ani2.I句柄)
	ani2.Run(w.I句柄)
	ani2.Move(2000, 750, float32(top), 1, 0, true)
	top += 100

	// 一次, 不往返
	ani3 := ani.NewAnima(list_svg[2], 1)
	list_animation = append(list_animation, ani3.I句柄)
	ani3.Run(w.I句柄)
	ani3.Move(2000, 750, float32(top), 1, 0, false)
	top += 100

	/* 以下是纯函数方式实现
	// 循环
	hAnimation := xc.XAnima_Create(list_svg[0], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.I句柄)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 10, 0, true)
	top += 100

	// 一次, 往返
	hAnimation = xc.XAnima_Create(list_svg[1], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.I句柄)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 1, 0, true)
	top += 100

	// 一次, 不往返
	hAnimation = xc.XAnima_Create(list_svg[2], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.I句柄)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 1, 0, false)
	*/

	return 0
}

func OnBtnClick6(pbHandled *bool) int {
	ReleaseAnimation()
	left := 140
	top = 100

	// 创建形状文本
	hShapeText1 := xc.XShapeText_Create(left, top, 100, 30, "循环滚动", w.I句柄)
	top += 50
	hShapeText2 := xc.XShapeText_Create(left, top, 100, 30, "往返滚动", w.I句柄)
	top += 50
	hShapeText3 := xc.XShapeText_Create(left, top, 100, 30, "移动到末尾", w.I句柄)
	top += 50
	list_xcgui = append(list_xcgui,
		hShapeText1,
		hShapeText2,
		hShapeText3,
	)
	top = 100

	ani1 := ani.NewAnima(hShapeText1, 0)
	list_animation = append(list_animation, ani1.I句柄)
	ani1.Run(w.I句柄)
	ani1.Move(3000, 750, float32(top), 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)

	ani2 := ani.NewAnima(hShapeText2, 1)
	list_animation = append(list_animation, ani2.I句柄)
	ani2.Run(w.I句柄)
	ani2.Move(3000, 750, float32(top+50), 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)

	ani3 := ani.NewAnima(hShapeText3, 1)
	list_animation = append(list_animation, ani3.I句柄)
	ani3.Run(w.I句柄)
	ani3.Move(1500, 750, float32(top+100), 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)

	/* 	以下是纯函数方式实现
	hAnimation := xc.XAnima_Create(hShapeText1, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.I句柄)
	xc.XAnima_Move(hAnimation, 3000, 750, float32(top), 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)

	hAnimation = xc.XAnima_Create(hShapeText2, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.I句柄)
	xc.XAnima_Move(hAnimation, 3000, 750, float32(top+50), 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)

	hAnimation = xc.XAnima_Create(hShapeText3, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.I句柄)
	xc.XAnima_Move(hAnimation, 1500, 750, float32(top+100), 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)
	*/

	return 0
}

func OnBtnClick7(pbHandled *bool) int {
	ReleaseAnimation()
	left := 125
	top = 50

	group1 := ani.NewAnimaGroup(0)
	list_animation = append(list_animation, group1.I句柄)
	group1.Run(w.I句柄)
	for i := 0; i < 13; i++ {
		hButton := xc.XBtn_Create(left, top, 60, 30, "透明度", w.I句柄)
		list_xcgui = append(list_xcgui, hButton)

		hAnimation := xc.XAnima_Create(hButton, 0)
		group1.AddItem(hAnimation)

		xc.XAnima_Delay(hAnimation, 500)

		xc.XAnima_Delay(hAnimation, 100*float32(i))
		xc.XAnima_AlphaEx(hAnimation, 1200, 255, 20, 1, 0, true)
		left += 61
	}

	left = 125
	top = 100
	group2 := ani.NewAnimaGroup(0)
	list_animation = append(list_animation, group2.I句柄)
	group2.Run(w.I句柄)
	for i := 0; i < 7; i++ {
		hButton := xc.XBtn_Create(left, top, 80, 30, "循环滚动", w.I句柄)
		list_xcgui = append(list_xcgui, hButton)

		hAnimation := xc.XAnima_Create(hButton, 0)
		group2.AddItem(hAnimation)

		xc.XAnima_Move(hAnimation, 500, float32(left), float32(top), 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)
		xc.XAnima_Delay(hAnimation, 500)

		xc.XAnima_Delay(hAnimation, 100*float32(i))
		xc.XAnima_AlphaEx(hAnimation, 500, 255, 0, 1, 0, false)

		xc.XAnima_Delay(hAnimation, 500)

		xc.XAnima_AlphaEx(hAnimation, 500, 0, 255, 1, 0, false)
		xc.XAnima_Delay(hAnimation, 1000)

		xc.XAnima_Move(hAnimation, 2000, float32(left), 500, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, false)
		xc.XAnima_Delay(hAnimation, 1000)

		hAnimation = xc.XAnima_Create(hButton, 1)
		xc.XAnimaGroup_AddItem(group2.I句柄, hAnimation)
		xc.XAnima_Delay(hAnimation, 6000+float32(i)*200)
		xc.XAnima_Scale(hAnimation, 1200, 1.5, 2, 1, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快, true)

		left += 110
	}

	return 0
}

func OnBtnClick8(pbHandled *bool) int {
	ReleaseAnimation()
	hLayout := xc.XLayout_Create(140, 100, 750, 100, w.I句柄)
	xc.XLayoutBox_SetSpace(hLayout, 20)

	for i := 0; i < 3; i++ {
		hLayout_ := xc.XLayout_Create(0, 0, 100, 100, hLayout)
		xc.XEle_SetPadding(hLayout_, 10, 0, 10, 0)

		hShapeText := xc.XShapeText_Create(0, 0, 100, 100, "炫彩界面库-www.xcgui.com-鼠标移动到上面查看", hLayout_)
		xc.XShapeText_SetTextColor(hShapeText, xc.ABGR(255, 255, 255, 255))
		xc.XWidget_LayoutItem_SetWidth(hShapeText, xcc.I常量_布局大小_填充父, 0)

		list_xcgui = append(list_xcgui, hLayout_)
		xc.XEle_EnableMouseThrough(hLayout_, false)
		xc.XWidget_LayoutItem_SetWidth(hLayout_, xcc.I常量_布局大小_按照比例, 100)

		xc.XBkM_SetBkInfo(xc.XEle_GetBkManager(hLayout_), "{99:1.9.9;98:16(0);5:2(15)20(1)21(3)26(1)22(-7839744)23(255)9(5,5,5,5);}")
		xc.XEle_RegEventC1(hLayout_, xcc.XE_MOUSESTAY, OnMouseStay8)
		xc.XEle_RegEventC1(hLayout_, xcc.XE_MOUSELEAVE, OnMouseLeave8)

		switch i {
		case 0:
			hLayout1 = hLayout_
		case 1:
			hLayout2 = hLayout_
		case 2:
			hLayout3 = hLayout_
		}
	}

	xc.XWnd_AdjustLayout(w.I句柄)
	w.I重绘(false)
	return 0
}

// 鼠标进入事件8
func OnMouseStay8(hLayout int, pbHandled *bool) int {
	if hLayout1 != hLayout {
		xc.XEle_SetAlpha(hLayout1, 200)
	}

	if hLayout2 != hLayout {
		xc.XEle_SetAlpha(hLayout2, 200)
	}

	if hLayout3 != hLayout {
		xc.XEle_SetAlpha(hLayout3, 200)
	}

	hAnimation := xc.XAnima_Create(hLayout, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.I句柄)
	xc.XAnima_LayoutWidth(hAnimation, 300, xcc.I常量_布局大小_按照比例, 200, 1, 0, false)

	return 0
}

// 鼠标离开事件8
func OnMouseLeave8(hLayout, hEleStay int, pbHandled *bool) int {
	hAnimation := xc.XAnima_Create(hLayout, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.I句柄)
	xc.XAnima_LayoutWidth(hAnimation, 300, xcc.I常量_布局大小_按照比例, 100, 1, 0, false)

	xc.XEle_SetAlpha(hLayout1, 255)
	xc.XEle_SetAlpha(hLayout2, 255)
	xc.XEle_SetAlpha(hLayout3, 255)

	return 0
}

func OnBtnClick9(pbHandled *bool) int {
	ReleaseAnimation()
	left := 150
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
		hImage := xc.XImage_LoadMemory(imgMap[i*2+1])
		xc.XImage_SetDrawType(hImage, xcc.Image_Draw_Type_Fixed_Ratio)

		hEle := xc.XEle_Create(left, top, 212, 271, w.I句柄)
		xc.XEle_AddBkImage(hEle, xcc.Element_State_Flag_Leave, hImage)
		list_xcgui = append(list_xcgui, hEle)

		hImage2 := xc.XImage_LoadMemory(imgMap[i*2+2])
		xc.XImage_SetDrawType(hImage2, xcc.Image_Draw_Type_Fixed_Ratio)

		hEle2 := xc.XEle_Create(left, top, 212, 271, w.I句柄)
		xc.XEle_AddBkImage(hEle2, xcc.Element_State_Flag_Leave, hImage2)
		list_xcgui = append(list_xcgui, hEle2)

		hText := xc.XShapeText_Create(left, top+280, 200, 30, "炫彩界面库-图片切换\r\n$66.66", w.I句柄)
		xc.XShapeText_SetTextColor(hText, xc.ABGR(80, 80, 80, 255))
		list_xcgui = append(list_xcgui, hText)

		xc.XEle_SetUserData(hEle, hEle2)
		xc.XEle_SetUserData(hEle2, hEle)
		xc.XWidget_Show(hEle2, false)

		xc.XEle_RegEventC1(hEle, xcc.XE_MOUSESTAY, OnMouseStay9)
		xc.XEle_RegEventC1(hEle2, xcc.XE_MOUSELEAVE, OnMouseLeave9)

		left += 212 + 10
	}

	w.I重绘(false)
	return 0
}

// 鼠标进入事件9
func OnMouseStay9(hEle int, pbHandled *bool) int {
	hEle2 := xc.XEle_GetUserData(hEle)
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			hObjectUI := xc.XAnima_GetObjectUI(list_animation[i])
			if hEle == hObjectUI || hEle2 == hObjectUI {
				xc.XAnima_Release(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hAnimation := xc.XAnima_Create(hEle, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_AlphaEx(hAnimation, 1000, 255, 0, 1, 0, false)
	xc.XAnima_Show(hAnimation, 0, false)

	xc.XEle_SetAlpha(hEle2, 0)
	xc.XWidget_Show(hEle2, true)

	hAnimation = xc.XAnima_Create(hEle2, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle2)
	xc.XAnima_Delay(hAnimation, 500)
	xc.XAnima_AlphaEx(hAnimation, 1000, 0, 255, 1, 0, false)

	return 0
}

// 鼠标离开事件9
func OnMouseLeave9(hEle2, hEleStay int, pbHandled *bool) int {
	hEle := xc.XEle_GetUserData(hEle2)
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			hObjectUI := xc.XAnima_GetObjectUI(list_animation[i])
			if hEle == hObjectUI || hEle2 == hObjectUI {
				xc.XAnima_Release(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hAnimation := xc.XAnima_Create(hEle2, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle2)
	xc.XAnima_AlphaEx(hAnimation, 1000, 255, 0, 1, 0, false)
	xc.XAnima_Show(hAnimation, 0, false)

	xc.XEle_SetAlpha(hEle, 0)
	xc.XWidget_Show(hEle, true)

	hAnimation = xc.XAnima_Create(hEle, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Delay(hAnimation, 500)
	xc.XAnima_AlphaEx(hAnimation, 1000, 0, 255, 1, 0, false)

	return 0
}

func OnBtnClick10(pbHandled *bool) int {
	ReleaseAnimation()
	left := 150
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
		hEle := xc.XEle_Create(left, top, 212, 271, w.I句柄)
		xc.XEle_EnableDrawBorder(hEle, false)
		list_xcgui = append(list_xcgui, hEle)

		hImage := xc.XImage_LoadMemory(imgMap[i*2+1])
		xc.XImage_SetDrawType(hImage, xcc.Image_Draw_Type_Fixed_Ratio)

		hImage2 := xc.XImage_LoadMemory(imgMap[i*2+2])
		xc.XImage_SetDrawType(hImage2, xcc.Image_Draw_Type_Fixed_Ratio)

		hShapePic := xc.XShapePic_Create(0, 0, 212, 271, hEle)
		xc.XShapePic_SetImage(hShapePic, hImage)

		hShapePic2 := xc.XShapePic_Create(212+10, 0, 212, 271, hEle)
		xc.XShapePic_SetImage(hShapePic2, hImage2)

		hText := xc.XShapeText_Create(left, top+280, 200, 30, "炫彩界面库-图片切换\r\n$66.66", w.I句柄)
		xc.XShapeText_SetTextColor(hText, xc.ABGR(80, 80, 80, 255))
		list_xcgui = append(list_xcgui, hText)

		xc.XEle_RegEventC1(hEle, xcc.XE_MOUSESTAY, OnMouseStay10)
		xc.XEle_RegEventC1(hEle, xcc.XE_MOUSELEAVE, OnMouseLeave10)

		left += 212 + 10
	}

	w.I重绘(false)
	return 0
}

// 鼠标进入事件10
func OnMouseStay10(hEle int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			if hEle == xc.XAnima_GetObjectUI(list_animation[i]) {
				xc.XAnima_Release(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hPic := xc.XEle_GetChildByIndex(hEle, 0)

	hAnimation := xc.XAnima_Create(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Move(hAnimation, 500, -(212 + 10), 0, 1, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快, false)

	hPic = xc.XEle_GetChildByIndex(hEle, 1)

	hAnimation = xc.XAnima_Create(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Move(hAnimation, 500, 0, 0, 1, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快, false)

	return 0
}

// 鼠标离开事件10
func OnMouseLeave10(hEle, hEleStay int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			if hEle == xc.XAnima_GetObjectUI(list_animation[i]) {
				xc.XAnima_Release(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hPic := xc.XEle_GetChildByIndex(hEle, 0)

	hAnimation := xc.XAnima_Create(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Move(hAnimation, 500, 0, 0, 1, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快, false)

	hPic = xc.XEle_GetChildByIndex(hEle, 1)

	hAnimation = xc.XAnima_Create(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Move(hAnimation, 500, 212+10, 0, 1, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快, false)

	return 0
}

func OnBtnClick11(pbHandled *bool) int {
	ReleaseAnimation()
	left := 160
	top = 80

	// 两个球型交替移动
	hSvg = xc.XSvg_LoadStringW(svg11)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)

	hGroup := xc.XAnimaGroup_Create(0)
	list_animation = append(list_animation, hGroup)
	xc.XAnima_Run(hGroup, w.I句柄)

	hAnimation := xc.XAnima_Create(hSvg, 1)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	xc.XAnima_Move(hAnimation, 1000, float32(left)+50, float32(top), 1, xcc.I常量_缓动标识_正弦|xcc.I常量_缓动标识_慢到快再到慢, false)
	xc.XAnima_Move(hAnimation, 1000, float32(left), float32(top), 1, xcc.I常量_缓动标识_正弦|xcc.I常量_缓动标识_慢到快再到慢, false)

	hSvg = xc.XSvg_LoadStringW(svg12)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left+50, top)

	hGroup = xc.XAnimaGroup_Create(0)
	list_animation = append(list_animation, hGroup)
	xc.XAnima_Run(hGroup, w.I句柄)

	hAnimation = xc.XAnima_Create(hSvg, 1)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	xc.XAnima_Move(hAnimation, 1000, float32(left), float32(top), 1, xcc.I常量_缓动标识_正弦|xcc.I常量_缓动标识_慢到快再到慢, false)
	xc.XAnima_Move(hAnimation, 1000, float32(left)+50, float32(top), 1, xcc.I常量_缓动标识_正弦|xcc.I常量_缓动标识_慢到快再到慢, false)

	// 一排小球 缩放
	left = 350
	hGroup = xc.XAnimaGroup_Create(0)
	list_animation = append(list_animation, hGroup)
	xc.XAnima_Run(hGroup, w.I句柄)

	for i := 0; i < 10; i++ {
		hSvg = xc.XSvg_LoadStringW(svg13)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left+i*50, top)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		xc.XAnimaGroup_AddItem(hGroup, hAnimation)

		xc.XAnima_Delay(hAnimation, float32(i)*200)
		xc.XAnima_Scale(hAnimation, 1000, 2, 2, 1, 0, true)
	}

	// 一排小球 垂直缩放
	top = 150
	hGroup = xc.XAnimaGroup_Create(0)
	list_animation = append(list_animation, hGroup)
	xc.XAnima_Run(hGroup, w.I句柄)

	for i := 0; i < 10; i++ {
		hSvg = xc.XSvg_LoadStringW(svg13)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left+i*50, top)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		xc.XAnimaGroup_AddItem(hGroup, hAnimation)

		xc.XAnima_Delay(hAnimation, float32(i)*200)
		xc.XAnima_Scale(hAnimation, 1000, 1, 2, 1, 0, true)
	}

	// 一排小球 上下波浪
	left = 150
	top = 200
	for i := 0; i < 10; i++ {
		hSvg = xc.XSvg_LoadStringW(svg13)
		list_svg = append(list_svg, hSvg)
		x := left + i*35
		xc.XSvg_SetPosition(hSvg, x, top)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Run(hAnimation, w.I句柄)

		xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*100), true)
		xc.XAnima_Move(hAnimation, 1200, float32(x), float32(top)+100, 1, 0, true)
	}

	left = 550
	for i := 0; i < 10; i++ {
		hSvg = xc.XSvg_LoadStringW(svg13)
		list_svg = append(list_svg, hSvg)
		x := left + i*35
		xc.XSvg_SetPosition(hSvg, x, top)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Run(hAnimation, w.I句柄)

		xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*150), true)
		xc.XAnima_Move(hAnimation, 1000, float32(x), float32(top)+50, 1, xcc.I常量_缓动标识_正弦|xcc.I常量_缓动标识_慢到快再到慢, true)
	}

	// 一排小球 跳动
	left = 150
	top = 350
	for i := 0; i < 10; i++ {
		hSvg = xc.XSvg_LoadStringW(svg13)
		list_svg = append(list_svg, hSvg)
		x := left + i*35
		xc.XSvg_SetPosition(hSvg, x, top)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Run(hAnimation, w.I句柄)

		xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*200), true)
		xc.XAnima_Move(hAnimation, 500, float32(x), float32(top)+50, 1, xcc.I常量_缓动标识_五次方曲线|xcc.I常量_缓动标识_快到慢, true)
		xc.XAnima_Delay(hAnimation, 1700)
	}

	// 一排小球 移动
	left = 220
	top = 600
	for i := 0; i < 10; i++ {
		hSvg = xc.XSvg_LoadStringW(svg14)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, 100-i*25, top)
		xc.XSvg_SetAlpha(hSvg, 0)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Run(hAnimation, w.I句柄)

		xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*100), true)
		xc.XAnima_Move(hAnimation, 2000, 550-(float32(i)+1)*25, float32(top), 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_快到慢, false)
		xc.XAnima_Move(hAnimation, 2000, 900-(float32(i)+1)*25, float32(top), 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_慢到快, false)
		xc.XAnima_Move(hAnimation, 0, 100-float32(i)*25, float32(top), 1, 0, false)
		xc.XAnima_Delay(hAnimation, 500)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Run(hAnimation, w.I句柄)

		xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*100), true)
		xc.XAnima_AlphaEx(hAnimation, 2000, 0, 255, 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_快到慢, false)
		xc.XAnima_AlphaEx(hAnimation, 2000, 255, 0, 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_慢到快, false)
		xc.XAnima_Delay(hAnimation, 500)
	}

	w.I重绘(false)
	return 0
}

func OnBtnClick12(pbHandled *bool) int {
	ReleaseAnimation()
	left := 120
	top = 100

	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, 0)

	hAnimation := xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Rotate(hAnimation, 1700, 360, 1, 0, false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Move(hAnimation, 3000, float32(left)+500, float32(top), 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	// 移动 往返旋转
	top = 350
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, -45)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 0, 0, 255), true)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Rotate(hAnimation, 1500, 45, 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_慢到快, false)
	xc.XAnima_Rotate(hAnimation, 1500, -45, 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_慢到快, false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Move(hAnimation, 3000, float32(left)+500, float32(top), 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	return 0
}

func OnBtnClick13(pbHandled *bool) int {
	ReleaseAnimation()
	left := 120
	top = 80

	// 自身 摇摆 往返
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, -45)

	hAnimation := xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Rotate(hAnimation, 1000, 45, 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	// 自身 旋转
	left = 500
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Rotate(hAnimation, 1000, 360, 1, xcc.I常量_缓动标识_突击|xcc.I常量_缓动标识_慢到快, false)
	xc.XAnima_Rotate(hAnimation, 0, 0, 1, xcc.I常量_缓动标识_线性, false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	// 两个叠加 悬挂摆动
	left = 300
	top = 250
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, 45)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate := xc.XAnima_Rotate(hAnimation, 3000, 100, 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_慢到快再到慢, true)
	xc.XAnimaRotate_SetCenter(hRotate, float32(left)+10, float32(top)+50, false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, 45)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 0, 0, 255), true)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = xc.XAnima_Rotate(hAnimation, 3000, 100, 1, xcc.I常量_缓动标识_三次方曲线|xcc.I常量_缓动标识_慢到快再到慢, true)
	xc.XAnimaRotate_SetCenter(hRotate, float32(left)+10, float32(top)+50, false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	// 砍东西效果
	left = 500
	top = 400
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, -45)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 0, 0, 255), true)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = xc.XAnima_Rotate(hAnimation, 1500, 0, 1, xcc.I常量_缓动标识_突击|xcc.I常量_缓动标识_慢到快, true)
	xc.XAnimaRotate_SetCenter(hRotate, float32(left), float32(top), false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	return 0
}

func OnBtnClick14(pbHandled *bool) int {
	ReleaseAnimation()
	left := 130
	top = 50

	// 加载svg, 设置大小和填充颜色
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetSize(hSvg, 50, 50)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 0, 0, 255), true)

	// 移动 360度旋转
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, 0)

	// 创建动画组
	hGroup := xc.XAnimaGroup_Create(0)
	list_animation = append(list_animation, hGroup)

	// 旋转
	hAnimation := xc.XAnima_Create(hSvg, 0)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	xc.XAnima_Rotate(hAnimation, 600, 360, 4, 0, false)

	// 缩放
	hAnimation = xc.XAnima_Create(hSvg, 0)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	xc.XAnima_Scale(hAnimation, 2400, 7, 7, 1, 0, false)
	xc.XAnima_Delay(hAnimation, 1000)
	xc.XAnima_Scale(hAnimation, 1000, 1.0/7.0, 1.0/7.0, 1, 0, false)

	// 移动
	hAnimation = xc.XAnima_Create(hSvg, 0)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	xc.XAnima_Move(hAnimation, 2400, float32(left)+500, float32(top)+300, 1, 0, false)
	xc.XAnima_Delay(hAnimation, 1000)
	xc.XAnima_Move(hAnimation, 1000, float32(left), float32(top), 1, 0, false)
	xc.XAnima_Run(hGroup, w.I句柄)

	return 0
}

func OnBtnClick15(pbHandled *bool) int {
	ReleaseAnimation()
	left := 150
	top = 200
	height := 0
	width := 0

	// 砍东西效果

	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	height = xc.XSvg_GetHeight(hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, -45)

	hAnimation := xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate := xc.XAnima_Rotate(hAnimation, 2000, 0, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)
	xc.XAnimaRotate_SetCenter(hRotate, float32(left), float32(top+height/2.0), false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	// 砍东西效果
	top = 300
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	height = xc.XSvg_GetHeight(hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, 45)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = xc.XAnima_Rotate(hAnimation, 2000, 0, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)
	xc.XAnimaRotate_SetCenter(hRotate, float32(left), float32(top+height/2.0), false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	// 砍东西效果
	left = 500
	top = 200
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	width = xc.XSvg_GetWidth(hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, 45)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 0, 0, 255), true)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = xc.XAnima_Rotate(hAnimation, 2000, 0, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)
	xc.XAnimaRotate_SetCenter(hRotate, float32(left+width), float32(top+height/2.0), false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	// 砍东西效果
	top = 300
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	width = xc.XSvg_GetWidth(hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, -45)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 0, 0, 255), true)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hRotate = xc.XAnima_Rotate(hAnimation, 2000, 0, 1, xcc.I常量_缓动标识_弹跳|xcc.I常量_缓动标识_快到慢, true)
	xc.XAnimaRotate_SetCenter(hRotate, float32(left+width), float32(top+height/2.0), false)
	xc.XAnima_Run(hAnimation, w.I句柄)

	return 0
}

func OnBtnClick16(pbHandled *bool) int {
	ReleaseAnimation()
	left := 150
	top = 50

	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 0, 0, 255), true)

	hAnimation := xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Color(hAnimation, 1500, xc.ABGR(0, 0, 255, 255), 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = 225
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(0, 255, 0, 255), true)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Color(hAnimation, 1500, xc.ABGR(255, 0, 0, 255), 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = 400
	hSvg = xc.XSvg_LoadStringW(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 255, 0, 255), true)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Color(hAnimation, 1500, xc.ABGR(0, 0, 255, 255), 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	hSvg = xc.XSvg_LoadString(svg15)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, 500, 300)
	xc.XSvg_SetUserFillColor(hSvg, xc.ABGR(255, 255, 0, 255), true)

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Color(hAnimation, 1500, xc.ABGR(0, 255, 255, 255), 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	hFontx := xc.XFont_CreateEx("微软雅黑", 36, xcc.I常量_字体样式_粗体)
	hShapeText := xc.XShapeText_Create(500, 100, 200, 50, "炫彩界面库", w.I句柄)
	list_xcgui = append(list_xcgui, hShapeText)
	xc.XShapeText_SetFont(hShapeText, hFontx)
	xc.XShapeText_SetTextColor(hShapeText, xc.ABGR(255, 0, 0, 255))

	hAnimation = xc.XAnima_Create(hShapeText, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Color(hAnimation, 1500, xc.ABGR(0, 0, 255, 255), 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	hShapeText = xc.XShapeText_Create(500, 200, 100, 20, "炫彩界面库", w.I句柄)
	list_xcgui = append(list_xcgui, hShapeText)

	hAnimation = xc.XAnima_Create(hShapeText, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Color(hAnimation, 1500, xc.ABGR(0, 255, 0, 255), 1, 0, true)
	xc.XAnima_Run(hAnimation, w.I句柄)

	return 0
}

func OnBtnClick17(pbHandled *bool) int {
	ReleaseAnimation()
	left := 150
	top = 50

	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_leftTop", w.I句柄))

	hAnimation := xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale := xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_左上角)
	xc.XAnima_Run(hAnimation, w.I句柄)
	top = top + 150
	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_left", w.I句柄))

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_左)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = top + 150
	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_leftBottom", w.I句柄))

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_左下角)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = 50
	left = left + 150
	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_top", w.I句柄))

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_上)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = top + 150
	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_center", w.I句柄))

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_中心)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = top + 150
	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_bottom", w.I句柄))

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_下)
	xc.XAnima_Run(hAnimation, w.I句柄)

	left = left + 150
	top = 50
	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_rightTop", w.I句柄))

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_右上角)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = top + 150
	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_right", w.I句柄))

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_右)
	xc.XAnima_Run(hAnimation, w.I句柄)

	top = top + 150
	hSvg = xc.XSvg_LoadStringW(svg5)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetPosition(hSvg, left, top)
	list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_rightBottom", w.I句柄))

	hAnimation = xc.XAnima_Create(hSvg, 0)
	list_animation = append(list_animation, hAnimation)
	hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_右下角)
	xc.XAnima_Run(hAnimation, w.I句柄)

	return 0
}

func OnBtnClick18(pbHandled *bool) int {
	ReleaseAnimation()
	left := 150
	top = 50

	for i := 0; i < 5; i++ {
		hButton := xc.XBtn_Create(left, top, 100, 50, "鼠标 停留 离开", w.I句柄)
		list_xcgui = append(list_xcgui, hButton)
		xc.XEle_RegEventC1(hButton, xcc.XE_MOUSESTAY, OnMouseStay18)
		xc.XEle_RegEventC1(hButton, xcc.XE_MOUSELEAVE, OnMouseLeave18)
		top = top + 60
	}
	w.I重绘(false)

	return 0
}

// 鼠标进入事件18
func OnMouseStay18(hButton int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			if hButton == xc.XAnima_GetObjectUI(list_animation[i]) {
				xc.XAnima_Release(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hAnimation := xc.XAnima_Create(hButton, 1)
	list_animation = append(list_animation, hAnimation)
	hScale := xc.XAnima_ScaleSize(hAnimation, 400, 200, 50, 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_快到慢, false)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_左)
	xc.XAnima_Run(hAnimation, w.I句柄)

	return 0
}

// 鼠标离开事件18
func OnMouseLeave18(hButton, hEleStay int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := 0; i < len(list_animation); i++ {
		if len(list_animation)-2 == i {
			if hButton == xc.XAnima_GetObjectUI(list_animation[i]) {
				xc.XAnima_Release(list_animation[i], false)
				list_animation = append(list_animation[:i], list_animation[i+1:]...)
			}
		}
	}

	hAnimation := xc.XAnima_Create(hButton, 1)
	list_animation = append(list_animation, hAnimation)
	hScale := xc.XAnima_ScaleSize(hAnimation, 400, 100, 50, 1, xcc.I常量_缓动标识_二次方曲线|xcc.I常量_缓动标识_慢到快, false)
	xc.XAnimaScale_SetPosition(hScale, xcc.I常量_位置标识_左)
	xc.XAnima_Run(hAnimation, w.I句柄)

	return 0
}
