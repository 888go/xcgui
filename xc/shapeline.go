package 炫彩基类

// 形状线_创建.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
//
// hParent: 父对象句柄.
func X形状线_创建(x1 int, y1 int, x2 int, y2 int, 父对象句柄 int) int {
	r, _, _ := xShapeLine_Create.Call(uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(父对象句柄))
	return int(r)
}

// 形状线_置位置.
//
// hShape: 形状对象句柄.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func X形状线_置位置(形状对象句柄 int, x1 int, y1 int, x2 int, y2 int) int {
	r, _, _ := xShapeLine_SetPosition.Call(uintptr(形状对象句柄), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return int(r)
}

// 形状线_置颜色, 设置直线颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func X形状线_置颜色(形状对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xShapeLine_SetColor.Call(uintptr(形状对象句柄), uintptr(ABGR颜色值))
	return int(r)
}
