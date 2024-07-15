// 代码编辑框
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/font"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func main() {
	panic("由于代码编辑框的API正在升级, 所以[代码编辑框的部分函数]会用不了, 等待xcgui原作者更新后将会开放大量接口, 比以前更好用")
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 1000, 600, "Editor", 0, 炫彩常量类.Window_Style_Default)

	// 创建Editor
	Editor := 炫彩组件类.X创建代码编辑框(12, 35, 975, 555, w.Handle)
	// 启用接收Tab输入
	Editor.X启用接收TAB(true)
	// 启用自动换行
	Editor.X启用自动换行(true)

	// 创建字体
	font1 := 炫彩字体类.X创建EX("Arial", 12, 炫彩常量类.FontStyle_Regular)
	// 设置Editor的字体
	Editor.X置字体(font1.Handle)
	// 设置默认颜色
	Editor.X置文本颜色(炫彩基类.ABGR(100, 100, 100, 255))

	// 添加样式
	iStyle_fun := Editor.X添加样式(0, 炫彩基类.ABGR(255, 128, 0, 255), true)     // 函数
	iStyle_str := Editor.X添加样式(0, 炫彩基类.ABGR(206, 145, 120, 255), true)   // 字符串
	iStyle_comment := Editor.X添加样式(0, 炫彩基类.ABGR(67, 166, 74, 255), true) // 注释
	iStyle_key1 := Editor.X添加样式(0, 炫彩基类.ABGR(86, 156, 214, 255), true)   // key1
	iStyle_key2 := Editor.X添加样式(0, 炫彩基类.ABGR(200, 0, 0, 255), true)      // key2

	// 设置样式
	Editor.X置函数样式(iStyle_fun)
	Editor.X置字符串样式(iStyle_str)
	Editor.X置注释样式(iStyle_comment)
	// 添加关键字
	Editor.X添加关键字("if", iStyle_key1)
	Editor.X添加关键字("else", iStyle_key1)
	Editor.X添加关键字("switch", iStyle_key1)
	Editor.X添加关键字("case", iStyle_key1)
	Editor.X添加关键字("break", iStyle_key1)
	Editor.X添加关键字("int", iStyle_key1)

	Editor.X添加关键字("function", iStyle_key2)
	Editor.X添加关键字("return", iStyle_key2)

	// 添加自动匹配常量
	Editor.X添加自动匹配字符串(`XE_BNCLICK //按钮点击事件`)
	Editor.X添加自动匹配字符串(`XE_PAINT //元素绘制事件`)
	// 添加自动匹配函数
	Editor.X添加自动匹配函数(`function Tmp1(pFileName string, hParent int) //我是 Tmp1`)
	Editor.X添加自动匹配函数(`function Tmp2(pFileName string, hParent int) //我是 Tmp2`)
	Editor.X添加自动匹配函数(`function Tmp3(pFileName string, hParent int) //我是 Tmp3`)

	// 设置断点
	Editor.X置断点(1, true)
	Editor.X置断点(2, true)
	Editor.X置断点(3, false)

	/* 	//获取设置的断点
	   	var BreakPoints []int32
	   	Editor.GetBreakpoints(&BreakPoints, Editor.GetBreakpointCount())
	   	for _, v := range BreakPoints {
	   		fmt.Println(v)
	   	} */

	// 设置当前运行行
	Editor.X置当前运行(0)

	code := `// 123456
function foo(a int,b int) int{
	Tmp1("layout.xml",0);
	Tmp2("layout.xml",0);
	Tmp3("layout.xml",0);
	XE_BNCLICK;
	XE_PAINT;
	if(a == 1){

	}else{

	}

	switch(a){
	case 0:
		break;
	case 1:
		break;
	}

	return 0
}`
	Editor.X置文本(code)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
