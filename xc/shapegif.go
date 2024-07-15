package 炫彩基类

// XShapeGif_Create 形状GIF_创建, 创建形状对象GIF, 成功返回形状对象GIF句柄, 否则返回NULL.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.

// ff:形状GIF_创建
// hParent:父对象句柄
// cy:高度
// cx:宽度
// y:Y坐标
// x:坐标
func X形状GIF_创建(坐标 int, Y坐标 int, 宽度 int, 高度 int, 父对象句柄 int) int {
	r, _, _ := xShapeGif_Create.Call(uintptr(坐标), uintptr(Y坐标), uintptr(宽度), uintptr(高度), uintptr(父对象句柄))
	return int(r)
}

// XShapeGif_SetImage 形状GIF_置图片, 设置GIF图片.
//
// hShape: 形状对象句柄.
//
// hImage: 图片句柄.

// ff:形状GIF_置图片
// hImage:图片句柄
// hShape:形状对象句柄
func X形状GIF_置图片(形状对象句柄 int, 图片句柄 int) int {
	r, _, _ := xShapeGif_SetImage.Call(uintptr(形状对象句柄), uintptr(图片句柄))
	return int(r)
}

// XShapeGif_GetImage 形状GIF_取图片, 获取图片句柄.
//
// hShape: 形状对象句柄.

// ff:形状GIF_取图片
// hShape:形状对象句柄
func X形状GIF_取图片(形状对象句柄 int) int {
	r, _, _ := xShapeGif_GetImage.Call(uintptr(形状对象句柄))
	return int(r)
}
