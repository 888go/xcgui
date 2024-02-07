package 炫彩WinApi工具类_test

import (
	"fmt"
	"github.com/888go/xcgui/wapi/wutil"
	"testing"
)

func TestSetClipboardText(t *testing.T) {
	err := 炫彩WinApi工具类.SetClipboardText("SetClipboardText")
	if err != nil {
		fmt.Println(err)
	}
	text, err := 炫彩WinApi工具类.GetClipboardText()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
