// 代码编辑框
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/font"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 1000, 600, "Editor", 0, xcc.I常量_窗口样式_默认)

	// 创建Editor
	Editor := widget.I代码编辑框_创建(12, 35, 975, 555, w.I句柄)
	// 启用接收Tab输入
	Editor.I启用接收TAB(true)
	// 启用自动换行
	Editor.I启用自动换行(true)

	// 创建字体
	font1 := font.NewFontEX("Arial", 12, xcc.I常量_字体样式_正常)
	// 设置Editor的字体
	Editor.I置字体(font1.I句柄)
	// 设置默认颜色
	Editor.I置文本颜色(xc.ABGR(100, 100, 100, 255))

	// 添加样式
	iStyle_fun := Editor.I添加样式(0, xc.ABGR(255, 128, 0, 255), true)     // 函数
	iStyle_str := Editor.I添加样式(0, xc.ABGR(206, 145, 120, 255), true)   // 字符串
	iStyle_comment := Editor.I添加样式(0, xc.ABGR(67, 166, 74, 255), true) // 注释
	iStyle_key1 := Editor.I添加样式(0, xc.ABGR(86, 156, 214, 255), true)   // key1
	iStyle_key2 := Editor.I添加样式(0, xc.ABGR(200, 0, 0, 255), true)      // key2

	// 设置样式
	Editor.I置函数样式(iStyle_fun)
	Editor.I置字符串样式(iStyle_str)
	Editor.I置注释样式(iStyle_comment)
	// 添加关键字
	Editor.I添加关键字("if", iStyle_key1)
	Editor.I添加关键字("else", iStyle_key1)
	Editor.I添加关键字("switch", iStyle_key1)
	Editor.I添加关键字("case", iStyle_key1)
	Editor.I添加关键字("break", iStyle_key1)
	Editor.I添加关键字("int", iStyle_key1)

	Editor.I添加关键字("function", iStyle_key2)
	Editor.I添加关键字("return", iStyle_key2)

	// 添加自动匹配常量
	Editor.I添加自动匹配字符串(`XE_BNCLICK //按钮点击事件`)
	Editor.I添加自动匹配字符串(`XE_PAINT //元素绘制事件`)
	// 添加自动匹配函数
	Editor.I添加自动匹配函数(`function Tmp1(pFileName string, hParent int) //我是 Tmp1`)
	Editor.I添加自动匹配函数(`function Tmp2(pFileName string, hParent int) //我是 Tmp2`)
	Editor.I添加自动匹配函数(`function Tmp3(pFileName string, hParent int) //我是 Tmp3`)

	// 设置断点
	Editor.I置断点(1, true)
	Editor.I置断点(2, true)
	Editor.I置断点(3, false)

	/* 	//获取设置的断点
	   	var BreakPoints []int32
	   	Editor.I取全部断点(&BreakPoints, Editor.I取断点数量())
	   	for _, v := range BreakPoints {
	   		fmt.Println(v)
	   	} */

	// 设置当前运行行
	Editor.I置当前运行(0)

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
	Editor.I置文本(code)

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
