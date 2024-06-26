package xc

// 形状线_创建.
//.
//.
//.
//.
//.
// ff:形状线_创建
// x1:
// y1:
// x2:
// y2:
// hParent:父对象句柄
func XShapeLine_Create(x1 int, y1 int, x2 int, y2 int, hParent int) int {
	r, _, _ := xShapeLine_Create.Call(uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(hParent))
	return int(r)
}

// 形状线_置位置.
//.
//.
//.
//.
//.
// ff:形状线_置位置
// hShape:形状对象句柄
// x1:
// y1:
// x2:
// y2:
func XShapeLine_SetPosition(hShape int, x1 int, y1 int, x2 int, y2 int) int {
	r, _, _ := xShapeLine_SetPosition.Call(uintptr(hShape), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return int(r)
}

// 形状线_置颜色, 设置直线颜色.
//.
//.
// ff:形状线_置颜色
// hShape:形状对象句柄
// color:ABGR颜色值
func XShapeLine_SetColor(hShape int, color int) int {
	r, _, _ := xShapeLine_SetColor.Call(uintptr(hShape), uintptr(color))
	return int(r)
}
