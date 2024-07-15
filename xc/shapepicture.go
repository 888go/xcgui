package 炫彩基类

// XShapePic_Create 形状图片_创建, 创建形状对象-图片, 成功返回图片对象句柄, 否则返回NULL.
//
// x: x坐标.
//
// y: y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.

// ff:形状图片_创建
// hParent:父对象句柄
// cy:高度
// cx:宽度
// y:y坐标
// x:x坐标
func X形状图片_创建(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父对象句柄 int) int {
	r, _, _ := xShapePic_Create.Call(uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), uintptr(父对象句柄))
	return int(r)
}

// XShapePic_SetImage 形状图片_置图片, 设置图片.
//
// hShape: 形状对象句柄.
//
// hImage: 图片句柄.

// ff:形状图片_置图片
// hImage:图片句柄
// hShape:形状对象句柄
func X形状图片_置图片(形状对象句柄 int, 图片句柄 int) int {
	r, _, _ := xShapePic_SetImage.Call(uintptr(形状对象句柄), uintptr(图片句柄))
	return int(r)
}

// XShapePic_GetImage 形状图片_取图片, 获取图片句柄.
//
// hShape: 形状对象句柄.

// ff:形状图片_取图片
// hShape:形状对象句柄
func X形状图片_取图片(形状对象句柄 int) int {
	r, _, _ := xShapePic_GetImage.Call(uintptr(形状对象句柄))
	return int(r)
}
