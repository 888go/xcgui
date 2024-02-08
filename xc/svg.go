package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
)

// SVG_加载从文件, 返回SVG句柄.
//
// pFileName: 文件名.
func SVG_加载从文件(文件名 string) int {
	r, _, _ := xSvg_LoadFile.Call(炫彩工具类.StrPtr(文件名))
	return int(r)
}

// SVG_加载从字符串, 返回SVG句柄.
//
// pString: 字符串.
func SVG_加载从字符串(字符串 string) int {
	r, _, _ := xSvg_LoadString.Call(W2A(字符串))
	return int(r)
}

// SVG_加载从ZIP, 返回SVG句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func SVG_加载从ZIP(zip文件名, svg文件名, zip密码 string) int {
	r, _, _ := xSvg_LoadZip.Call(炫彩工具类.StrPtr(zip文件名), 炫彩工具类.StrPtr(svg文件名), 炫彩工具类.StrPtr(zip密码))
	return int(r)
}

// SVG_加载从资源ZIP, 返回SVG句柄.
//
// id: 资源ID.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func SVG_加载从资源ZIP(资源ID int32, svg文件名, zip密码 string, 模块句柄 uintptr) int {
	r, _, _ := xSvg_LoadZipRes.Call(uintptr(资源ID), 炫彩工具类.StrPtr(svg文件名), 炫彩工具类.StrPtr(zip密码), 模块句柄)
	return int(r)
}

// SVG_加载从资源, 返回SVG句柄.
//
// id: 资源ID.
//
// pType: 资源类型.在rc资源文件中.
//
// hModule: 从指定模块加载.
func SVG_加载从资源(资源ID int32, 资源类型 string, 从指定模块加载 uintptr) int {
	r, _, _ := xSvg_LoadRes.Call(uintptr(资源ID), 炫彩工具类.StrPtr(资源类型), 从指定模块加载)
	return int(r)
}

// SVG_置大小.
//
// hSvg: SVG句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func SVG_置大小(SVG句柄 int, 宽度, 高度 int32) {
	xSvg_SetSize.Call(uintptr(SVG句柄), uintptr(宽度), uintptr(高度))
}

// SVG_取大小.
//
// hSvg: SVG句柄.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func SVG_取大小(SVG句柄 int, 接收返回宽度, 接收返回高度 *int32) {
	xSvg_GetSize.Call(uintptr(SVG句柄), uintptr(unsafe.Pointer(接收返回宽度)), uintptr(unsafe.Pointer(接收返回高度)))
}

// SVG_取宽度.
//
// hSvg: SVG句柄.
func SVG_取宽度(SVG句柄 int) int32 {
	r, _, _ := xSvg_GetWidth.Call(uintptr(SVG句柄))
	return int32(r)
}

// SVG_取高度.
//
// hSvg: SVG句柄.
func SVG_取高度(SVG句柄 int) int32 {
	r, _, _ := xSvg_GetHeight.Call(uintptr(SVG句柄))
	return int32(r)
}

// SVG_置偏移.
//
// hSvg: SVG句柄.
//
// x: x轴偏移.
//
// y: y轴偏移.
func SVG_置偏移(SVG句柄 int, x轴偏移, y轴偏移 int32) {
	xSvg_SetPosition.Call(uintptr(SVG句柄), uintptr(x轴偏移), uintptr(y轴偏移))
}

// SVG_取偏移.
//
// hSvg: SVG句柄.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func SVG_取偏移(SVG句柄 int, x轴偏移, y轴偏移 *int32) {
	xSvg_GetPosition.Call(uintptr(SVG句柄), uintptr(unsafe.Pointer(x轴偏移)), uintptr(unsafe.Pointer(y轴偏移)))
}

// SVG_置偏移F.
//
// hSvg: SVG句柄.
//
// x: x轴偏移.
//
// y: y轴偏移.
func SVG_置偏移F(SVG句柄 int, x轴偏移, y轴偏移 float32) {
	xSvg_SetPositionF.Call(uintptr(SVG句柄), 炫彩工具类.Float32Ptr(x轴偏移), 炫彩工具类.Float32Ptr(y轴偏移))
}

// SVG_取偏移F.
//
// hSvg: SVG句柄.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func SVG_取偏移F(SVG句柄 int, x轴偏移, y轴偏移 *float32) {
	xSvg_GetPositionF.Call(uintptr(SVG句柄), uintptr(unsafe.Pointer(x轴偏移)), uintptr(unsafe.Pointer(y轴偏移)))
}

// SVG_取视图框.
//
// hSvg: SVG句柄.
//
// pViewBox: 接收返回视图框.
func SVG_取视图框(SVG句柄 int, 接收返回视图框 *RECT) {
	xSvg_GetViewBox.Call(uintptr(SVG句柄), uintptr(unsafe.Pointer(接收返回视图框)))
}

// SVG_启用自动销毁.
//
// hSvg: SVG句柄.
//
// bEnable: 是否自动销毁.
func SVG_启用自动销毁(SVG句柄 int, 是否自动销毁 bool) {
	xSvg_EnableAutoDestroy.Call(uintptr(SVG句柄), 炫彩工具类.BoolPtr(是否自动销毁))
}

// SVG_增加引用计数.
//
// hSvg: SVG句柄.
func SVG_增加引用计数(SVG句柄 int) {
	xSvg_AddRef.Call(uintptr(SVG句柄))
}

// SVG_释放引用计数.
//
// hSvg: SVG句柄.
func SVG_释放引用计数(SVG句柄 int) {
	xSvg_Release.Call(uintptr(SVG句柄))
}

// SVG_取引用计数.
//
// hSvg: SVG句柄.
func SVG_取引用计数(SVG句柄 int) int32 {
	r, _, _ := xSvg_GetRefCount.Call(uintptr(SVG句柄))
	return int32(r)
}

// SVG_销毁.
//
// hSvg: SVG句柄.
func SVG_销毁(SVG句柄 int) {
	xSvg_Destroy.Call(uintptr(SVG句柄))
}

// SVG_置透明度.
//
// hSvg: SVG句柄.
//
// alpha: 透明度.
func SVG_置透明度(SVG句柄 int, 透明度 byte) {
	xSvg_SetAlpha.Call(uintptr(SVG句柄), uintptr(透明度))
}

// SVG_取透明度, 返回透明度.
//
// hSvg: SVG句柄.
func SVG_取透明度(SVG句柄 int) byte {
	r, _, _ := xSvg_GetAlpha.Call(uintptr(SVG句柄))
	return byte(r)
}

// SVG_置用户填充颜色, 用户颜色将覆盖默认样式.
//
// hSvg: SVG句柄.
//
// color: 颜色, AGBR颜色.
//
// bEnable: 是否有效.
func SVG_置用户填充颜色(SVG句柄 int, AGBR颜色 int, 是否有效 bool) {
	xSvg_SetUserFillColor.Call(uintptr(SVG句柄), uintptr(AGBR颜色), 炫彩工具类.BoolPtr(是否有效))
}

// SVG_置用户笔触颜色, 用户颜色将覆盖默认样式.
//
// hSvg: SVG句柄.
//
// color: 颜色, AGBR颜色.
//
// strokeWidth: 笔触宽度.
//
// bEnable: 是否有效.
func SVG_置用户笔触颜色(SVG句柄 int, AGBR颜色 int, 笔触宽度 float32, 是否有效 bool) {
	xSvg_SetUserStrokeColor.Call(uintptr(SVG句柄), uintptr(AGBR颜色), 炫彩工具类.Float32Ptr(笔触宽度), 炫彩工具类.BoolPtr(是否有效))
}

// SVG_取用户填充颜色.
//
// hSvg: SVG句柄.
//
// pColor: 返回颜色值, AGBR颜色.
func SVG_取用户填充颜色(SVG句柄 int, 返回颜色值 *int) bool {
	r, _, _ := xSvg_GetUserFillColor.Call(uintptr(SVG句柄), uintptr(unsafe.Pointer(返回颜色值)))
	return r != 0
}

// SVG_取用户笔触颜色.
//
// hSvg: SVG句柄.
//
// pColor: 返回颜色值, AGBR颜色.
//
// pStrokeWidth: .
func SVG_取用户笔触颜色(SVG句柄 int, 返回颜色值 *int, pStrokeWidth *float32) bool {
	r, _, _ := xSvg_GetUserStrokeColor.Call(uintptr(SVG句柄), uintptr(unsafe.Pointer(返回颜色值)), uintptr(unsafe.Pointer(pStrokeWidth)))
	return r != 0
}

// SVG_置旋转角度, 默认以自身中心点旋转.
//
// hSvg: SVG句柄.
//
// angle: 转角度.
func SVG_置旋转角度(SVG句柄 int, 转角度 float32) {
	xSvg_SetRotateAngle.Call(uintptr(SVG句柄), 炫彩工具类.Float32Ptr(转角度))
}

// SVG_取旋转角度, 返回旋转角度.
//
// hSvg: SVG句柄.
func SVG_取旋转角度(SVG句柄 int) float32 {
	_, r, _ := xSvg_GetRotateAngle.Call(uintptr(SVG句柄))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// SVG_置旋转.
//
// hSvg: SVG句柄.
//
// angle: 角度.
//
// x: 旋转中心点X.
//
// y: 旋转中心点Y.
//
// bOffset: TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func SVG_置旋转(SVG句柄 int, 角度 float32, 旋转中心点X float32, 旋转中心点Y float32, TRUE bool) {
	xSvg_SetRotate.Call(uintptr(SVG句柄), 炫彩工具类.Float32Ptr(角度), 炫彩工具类.Float32Ptr(旋转中心点X), 炫彩工具类.Float32Ptr(旋转中心点Y), 炫彩工具类.BoolPtr(TRUE))
}

// SVG_取旋转.
//
// hSvg: SVG句柄.
//
// pAngle: 返回角度.
//
// pX: 返回 旋转中心点X.
//
// pY: 返回 旋转中心点Y.
//
// pbOffset: 返回TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func SVG_取旋转(SVG句柄 int, 返回角度 *float32, 返回旋转中心点X *float32, 返回旋转中心点Y *float32, 返回TRUE *bool) {
	xSvg_GetRotate.Call(uintptr(SVG句柄), uintptr(unsafe.Pointer(返回角度)), uintptr(unsafe.Pointer(返回旋转中心点X)), uintptr(unsafe.Pointer(返回旋转中心点Y)), uintptr(unsafe.Pointer(返回TRUE)))
}

// SVG_显示, 显示或隐藏.
//
// hSvg: SVG句柄.
//
// bShow: 是否显示.
func SVG_显示(SVG句柄 int, 是否显示 bool) {
	xSvg_Show.Call(uintptr(SVG句柄), 炫彩工具类.BoolPtr(是否显示))
}

// SVG_加载从字符串W.
//
// pString: 字符串.
func SVG_加载从字符串W(字符串 string) int {
	r, _, _ := xSvg_LoadStringW.Call(炫彩工具类.StrPtr(字符串))
	return int(r)
}

// SVG_加载从字符串UTF8.
//
// pString: 字符串.
func SVG_加载从字符串UTF8(字符串 string) int {
	r, _, _ := xSvg_LoadStringUtf8.Call(X文本W到UTF8(字符串))
	return int(r)
}

// SVG_加载从内存ZIP.
//
// data: zip数据.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func SVG_加载从内存ZIP(zip数据 []byte, svg文件名, zip密码 string) int {
	r, _, _ := xSvg_LoadZipMem.Call(炫彩工具类.ByteSliceDataPtr(&zip数据), uintptr(len(zip数据)), 炫彩工具类.StrPtr(svg文件名), 炫彩工具类.StrPtr(zip密码))
	return int(r)
}
