// 调用UI设计器设计好的布局文件和资源文件
package main

import (
	_ "embed"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
)

//go:embed res/qqmusic.zip
var qqmusic []byte

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	// 从内存zip中加载资源文件
	a.X加载资源文件内存ZIP(qqmusic, "resource.res", "")
	// 从内存zip中加载布局文件, 创建窗口对象
	w := 炫彩窗口基类.X创建窗口并按内存压缩包布局文件(qqmusic, "main.xml", "", 0, 0)

	// songTitle是在main.xml中给歌曲名(shapeText组件)设置的name属性的值.
	// 通过 GetObjectByName 可以获取布局文件中设置了name属性的组件的句柄.
	// 可简化为: widget.NewShapeTextByName("songTitle").
	song := 炫彩组件类.X创建形状文本并按句柄(a.X取对象从名称("songTitle"))
	println(song.X取文本()) // 输出: 两只老虎爱跳舞

	// 调整布局
	w.X调整布局()
	// 显示窗口
	w.X显示(true)
	a.X运行()
	a.X退出()
}
