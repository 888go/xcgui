package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/window"
)

func main() {
	tf.TFunc(func(a *app.App, w *window.Window) {
		w.EnableDragFiles(true)
		w.Event_DROPFILES(func(hDropInfo uintptr, pbHandled *bool) int {
			fmt.Println(wutil.GetDropFiles(hDropInfo))
			return 0
		})
	})
}
