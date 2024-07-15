package 炫彩WinApi类_test

import (
	"fmt"
	"github.com/888go/xcgui/wapi"
	"testing"
)

func TestSleep(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		炫彩WinApi类.X延时(1000)
	}
}
