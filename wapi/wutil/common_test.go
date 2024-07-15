package 炫彩WinApi工具类_test

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/tf"
	"github.com/888go/xcgui/wapi/wutil"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"testing"
)

func TestGetDropFiles(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		w.X启用拖放文件(true)
		w.X线程_拖动文件到窗口(func(hDropInfo uintptr, pbHandled *bool) int {
			fmt.Println(炫彩WinApi工具类.X拖放文件取路径(hDropInfo))
			return 0
		})
	})
}

func TestOpenDir(t *testing.T) {
	fmt.Println(炫彩WinApi工具类.X对话框打开文件夹(0))
}

func TestOpenFile(t *testing.T) {
	fmt.Println(炫彩WinApi工具类.X对话框打开单个文件(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
}

func TestOpenFiles(t *testing.T) {
	for _, s := range 炫彩WinApi工具类.X对话框打开多个文件(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "") {
		fmt.Println(s)
	}
}

func TestSaveFile(t *testing.T) {
	fmt.Println(炫彩WinApi工具类.X对话框保存文件(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "", "默认文件名.txt"))
}

func TestChooseColor(t *testing.T) {
	rgb := 炫彩WinApi工具类.X对话框选择颜色(0)
	fmt.Println("rgb颜色", rgb)
	fmt.Println("abgr颜色", 炫彩基类.RGB2ABGR(rgb, 255))
}
