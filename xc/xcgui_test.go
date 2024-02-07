package 炫彩基类_test

import (
	"testing"
	
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func TestXInitXCGUI(t *testing.T) {
	// 可自定义xcgui.dll的路径, 如无需自定义, 则删掉这句代码.
	/*	err := xc.SetXcguiPath(`C:\Users\Administrator\Desktop\XCGUI.dll`)
		if err != nil {
			panic(err)
		}*/
	炫彩基类.XLoadXCGUI()
	炫彩基类.X初始化(true)
	hWindow := 炫彩基类.X窗口_创建(0, 0, 500, 500, "", 0, 炫彩常量类.Window_Style_Default)
	炫彩基类.X窗口_显示方式(hWindow, 炫彩常量类.SW_SHOW)
	炫彩基类.X运行()
	炫彩基类.X退出()
}

func TestWriteDll(t *testing.T) {
	err := 炫彩基类.X写出炫彩dll到临时目录([]byte("0"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(炫彩基类.X取炫彩DLL路径())
}

func TestABGR(t *testing.T) {
	t.Log(炫彩基类.ABGR(255, 201, 100, 255))
}

func TestRGBA(t *testing.T) {
	t.Log(炫彩基类.RGBA(255, 201, 100, 255))
}

func TestABGR2(t *testing.T) {
	t.Log(炫彩基类.ABGR2(炫彩基类.BGR(255, 201, 100), 255))
}

func TestRGB(t *testing.T) {
	t.Log(炫彩基类.RGB(255, 201, 100))
}

func TestGetXcgui(t *testing.T) {
	炫彩基类.XLoadXCGUI()
	dll := 炫彩基类.X取已加载炫彩DLL()
	t.Log(dll.Name, dll.Handle())
}

func TestGetXcgui1(t *testing.T) {
	t.Log(炫彩基类.X取炫彩DLL路径())
}

func TestSetXcguiPath(t *testing.T) {
	err := 炫彩基类.X设置炫彩DLL路径("C:\\xcgui.dll")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(炫彩基类.X取炫彩DLL路径())
}

func TestHexRGB2RGB(t *testing.T) {
	t.Log(炫彩基类.HexRGB2RGB("#ffc964"))
	t.Log(炫彩基类.HexRGB2RGB("ffc964"))
}

func TestHexRGB2ABGR(t *testing.T) {
	t.Log(炫彩基类.HexRGB2ABGR("#ffc964", 255))
	t.Log(炫彩基类.HexRGB2ABGR("ffc964", 255))
}

func TestHexRGB2BGR(t *testing.T) {
	t.Log(炫彩基类.HexRGB2BGR("#ffc964"))
	t.Log(炫彩基类.HexRGB2BGR("ffc964"))
}

func TestRGB2ABGR(t *testing.T) {
	t.Log(炫彩基类.RGB2ABGR(炫彩基类.HexRGB2RGB("ffc964"), 255))
}

func TestRGB2BGR(t *testing.T) {
	t.Log(炫彩基类.RGB2BGR(炫彩基类.HexRGB2RGB("ffc964")))
}

func TestBGR(t *testing.T) {
	t.Log(炫彩基类.BGR(255, 201, 100))
}
