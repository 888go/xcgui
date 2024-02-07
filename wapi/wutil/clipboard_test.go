package 炫彩WinApi工具类_test

import (
	"fmt"
	"github.com/888go/xcgui/wapi/wutil"
	"testing"
)

func TestSetClipboardText(t *testing.T) {
	err := 炫彩WinApi工具类.X剪贴板置文本("SetClipboardText")
	if err != nil {
		fmt.Println(err)
	}
	text, err := 炫彩WinApi工具类.X剪贴板取文本()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
