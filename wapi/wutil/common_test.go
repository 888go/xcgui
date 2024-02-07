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
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		w.EnableDragFiles(true)
		w.Event_DROPFILES(func(hDropInfo uintptr, pbHandled *bool) int {
			fmt.Println(炫彩WinApi工具类.GetDropFiles(hDropInfo))
			return 0
		})
	})
}

func TestOpenDir(t *testing.T) {
	fmt.Println(炫彩WinApi工具类.OpenDir(0))
}

func TestOpenFile(t *testing.T) {
	fmt.Println(炫彩WinApi工具类.OpenFile(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
}

func TestOpenFiles(t *testing.T) {
	for _, s := range 炫彩WinApi工具类.OpenFiles(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "") {
		fmt.Println(s)
	}
}

func TestSaveFile(t *testing.T) {
	fmt.Println(炫彩WinApi工具类.SaveFile(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "", "默认文件名.txt"))
}

func TestChooseColor(t *testing.T) {
	rgb := 炫彩WinApi工具类.ChooseColor(0)
	fmt.Println("rgb颜色", rgb)
	fmt.Println("abgr颜色", xc.RGB2ABGR(rgb, 255))
}
