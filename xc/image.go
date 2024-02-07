package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// 图片_加载从图片源.
//
// hImageSrc: 图片源句柄.
func X图片_加载从图片源(图片源句柄 int) int {
	r, _, _ := xImage_LoadSrc.Call(uintptr(图片源句柄))
	return int(r)
}

// 图片_加载从文件.
//
// pFileName: 图片文件.
func X图片_加载从文件(图片文件 string) int {
	r, _, _ := xImage_LoadFile.Call(炫彩工具类.StrPtr(图片文件))
	return int(r)
}

// 图片_加载从文件自适应, 加载图片从文件, 自适应图片.
//
// pFileName: 图片文件.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func X图片_加载从文件自适应(图片文件 string, 左, 上, 右, 下 int32) int {
	r, _, _ := xImage_LoadFileAdaptive.Call(炫彩工具类.StrPtr(图片文件), uintptr(左), uintptr(上), uintptr(右), uintptr(下))
	return int(r)
}

// 图片_加载从文件指定区域, 加载图片, 指定区位置及大小.
//
// pFileName: 图片文件.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func X图片_加载从文件指定区域(图片文件 string, 左, 右, 宽度, 高度 int32) int {
	r, _, _ := xImage_LoadFileRect.Call(炫彩工具类.StrPtr(图片文件), uintptr(左), uintptr(右), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 图片_加载从资源自适应, 加载图片从资源, 自适应图片.
//
// id: 资源ID.
//
// pType: 资源类型.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func X图片_加载从资源自适应(资源ID int32, 资源类型 string, 左, 上, 右, 下 int32, 从指定模块加载 uintptr) int {
	r, _, _ := xImage_LoadResAdaptive.Call(uintptr(资源ID), 炫彩工具类.StrPtr(资源类型), uintptr(左), uintptr(上), uintptr(右), uintptr(下), 从指定模块加载)
	return int(r)
}

// 图片_加载从资源.
//
// id: 资源ID.
//
// pType: 资源类型.
//
// bStretch: 是否拉伸图片.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func X图片_加载从资源(资源ID int32, 资源类型 string, 是否拉伸图片 bool, 从指定模块加载 uintptr) int {
	r, _, _ := xImage_LoadRes.Call(uintptr(资源ID), 炫彩工具类.StrPtr(资源类型), 炫彩工具类.BoolPtr(是否拉伸图片), 从指定模块加载)
	return int(r)
}

// 图片_加载从ZIP, 加载图片从ZIP压缩包.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
func X图片_加载从ZIP(ZIP压缩包文件名 string, 图片文件名 string, ZIP压缩包密码 string) int {
	r, _, _ := xImage_LoadZip.Call(炫彩工具类.StrPtr(ZIP压缩包文件名), 炫彩工具类.StrPtr(图片文件名), 炫彩工具类.StrPtr(ZIP压缩包密码))
	return int(r)
}

// 图片_加载从资源ZIP, 返回图片句柄.
//
// id: RC资源ID.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
//
// hModule: 模块句柄, 可填0.
func X图片_加载从资源ZIP(RC资源ID int32, 图片文件名 string, ZIP压缩包密码 string, 模块句柄 uintptr) int {
	r, _, _ := xImage_LoadZipRes.Call(uintptr(RC资源ID), 炫彩工具类.StrPtr(图片文件名), 炫彩工具类.StrPtr(ZIP压缩包密码), 模块句柄)
	return int(r)
}

// 图片_加载从ZIP自适应, 加载图片从ZIP压缩包, 自适应图片.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
//
// x1: 坐标.
//
// x2: 坐标.
//
// y1: 坐标.
//
// y2: 坐标.
func X图片_加载从ZIP自适应(ZIP压缩包文件名 string, 图片文件名 string, ZIP压缩包密码 string, x1, x2, y1, y2 int32) int {
	r, _, _ := xImage_LoadZipAdaptive.Call(炫彩工具类.StrPtr(ZIP压缩包文件名), 炫彩工具类.StrPtr(图片文件名), 炫彩工具类.StrPtr(ZIP压缩包密码), uintptr(x1), uintptr(x2), uintptr(y1), uintptr(y2))
	return int(r)
}

// 图片_加载从ZIP指定区域, 加载ZIP图片, 指定区位置及大小.
//
// pZipFileName: ZIP文件.
//
// pFileName: 图片名称.
//
// pPassword: 密码.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func X图片_加载从ZIP指定区域(ZIP文件 string, 图片名称 string, 密码 string, 左, 右, 宽度, 高度 int32) int {
	r, _, _ := xImage_LoadZipRect.Call(炫彩工具类.StrPtr(ZIP文件), 炫彩工具类.StrPtr(图片名称), 炫彩工具类.StrPtr(密码), uintptr(左), uintptr(右), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 图片_加载从内存ZIP.
//
// data: 图片数据.
//
// pFileName: 图片名称.
//
// pPassword: zip压缩包密码.
func X图片_加载从内存ZIP(图片数据 []byte, 图片名称 string, zip压缩包密码 string) int {
	r, _, _ := xImage_LoadZipMem.Call(炫彩工具类.ByteSliceDataPtr(&图片数据), uintptr(len(图片数据)), 炫彩工具类.StrPtr(图片名称), 炫彩工具类.StrPtr(zip压缩包密码))
	return int(r)
}

// 图片_加载从内存, 加载流图片.
//
// pBuffer: 图片数据.
func X图片_加载从内存(图片数据 []byte) int {
	r, _, _ := xImage_LoadMemory.Call(炫彩工具类.ByteSliceDataPtr(&图片数据), uintptr(len(图片数据)))
	return int(r)
}

// 图片_加载从内存指定区域, 加载流图片, 指定区位置及大小.
//
// pBuffer: 图片数据.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func X图片_加载从内存指定区域(图片数据 []byte, 左, 右, 宽度, 高度 int32) int {
	r, _, _ := xImage_LoadMemoryRect.Call(炫彩工具类.ByteSliceDataPtr(&图片数据), uintptr(len(图片数据)), uintptr(左), uintptr(右), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 图片_加载从内存自适应, 加载流图片压缩包, 自适应图片(九宫格).
//
// pBuffer: 图片数据.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func X图片_加载从内存自适应(图片数据 []byte, 左, 上, 右, 下 int32) int {
	r, _, _ := xImage_LoadMemoryAdaptive.Call(炫彩工具类.ByteSliceDataPtr(&图片数据), uintptr(len(图片数据)), uintptr(左), uintptr(上), uintptr(右), uintptr(下))
	return int(r)
}

// 图片_加载从Image, 加载图片从GDI+的Image对象.
//
// pImage: GDI图片对象指针Image*.
func X图片_加载从Image(GDI图片对象指针 uintptr) int {
	r, _, _ := xImage_LoadFromImage.Call(GDI图片对象指针)
	return int(r)
}

// 图片_加载文件图标, 加载文件图标, 从一个EXE文件或DLL文件或图标文件; 例如:*.exe文件的图标.
//
// pFileName: 文件名.
func X图片_加载文件图标(文件名 string) int {
	r, _, _ := xImage_LoadFromExtractIcon.Call(炫彩工具类.StrPtr(文件名))
	return int(r)
}

// 图片_加载从HICON, 创建一个炫彩图片句柄, 从一个现有的图标句柄HICON.
//
// hIcon: 图标句柄.
func X图片_加载从HICON(图标句柄 uintptr) int {
	r, _, _ := xImage_LoadFromHICON.Call(图标句柄)
	return int(r)
}

// 图片_加载从HBITMAP, 创建一个炫彩图片句柄, 从一个现有的位图句柄HBITMAP.
//
// hBitmap: 位图句柄.
func X图片_加载从HBITMAP(位图句柄 uintptr) int {
	r, _, _ := xImage_LoadFromHBITMAP.Call(位图句柄)
	return int(r)
}

// 图片_判断缩放, 是否为拉伸图片句柄.
//
// hImage: 图片句柄.
func X图片_判断缩放(图片句柄 int) bool {
	r, _, _ := xImage_IsStretch.Call(uintptr(图片句柄))
	return r != 0
}

// 图片_判断自适应, 是否为自适应图片句柄.
//
// hImage: 图片句柄.
func X图片_判断自适应(图片句柄 int) bool {
	r, _, _ := xImage_IsAdaptive.Call(uintptr(图片句柄))
	return r != 0
}

// 图片_判断平铺, 是否为平铺图片.
//
// hImage: 图片句柄.
func X图片_判断平铺(图片句柄 int) bool {
	r, _, _ := xImage_IsTile.Call(uintptr(图片句柄))
	return r != 0
}

// 图片_置绘制类型, 设置图片绘制类型.
//
// hImage: 图片句柄.
//
// nType: 图片绘制类型, Image_Draw_Type_.
func X图片_置绘制类型(图片句柄 int, 图片绘制类型 炫彩常量类.Image_Draw_Type_) bool {
	r, _, _ := xImage_SetDrawType.Call(uintptr(图片句柄), uintptr(图片绘制类型))
	return r != 0
}

// 图片_置绘制类型自适应, 设置图片自适应(九宫格).
//
// hImage: 图片句柄.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func X图片_置绘制类型自适应(图片句柄 int, 左, 上, 右, 下 int32) bool {
	r, _, _ := xImage_SetDrawTypeAdaptive.Call(uintptr(图片句柄), uintptr(左), uintptr(上), uintptr(右), uintptr(下))
	return r != 0
}

// 图片_置透明色, 指定图片透明颜色.
//
// hImage: 图片句柄.
//
// color: ABGR 颜色.
func X图片_置透明色(图片句柄 int, ABGR颜色 int) {
	xImage_SetTranColor.Call(uintptr(图片句柄), uintptr(ABGR颜色))
}

// 图片_置透明色扩展, 指定图片透明颜色及透明度.
//
// hImage: 图片句柄.
//
// color: ABGR 颜色.
//
// tranColor: 透明色的透明度.
func X图片_置透明色EX(图片句柄 int, ABGR颜色 int, 透明色的透明度 byte) {
	xImage_SetTranColorEx.Call(uintptr(图片句柄), uintptr(ABGR颜色), uintptr(透明色的透明度))
}

// 图片_置旋转角度, 设置旋转角度, 返回先前角度.
//
// hImage: 图片句柄.
//
// fAngle: 选择角度.
func X图片_置旋转角度(图片句柄 int, 选择角度 float32) float32 {
	_, r, _ := xImage_SetRotateAngle.Call(uintptr(图片句柄), 炫彩工具类.Float32Ptr(选择角度))
	return 炫彩工具类.UintPtrToFloat32(r)
}

// 图片_置等分.
//
// hImage: 图片句柄.
//
// nCount: 等分数量.
//
// iIndex: 索引.
func X图片_置等分(图片句柄 int, 等分数量 int32, 索引 int32) {
	xImage_SetSplitEqual.Call(uintptr(图片句柄), uintptr(等分数量), uintptr(索引))
}

// 图片_启用透明色, 启用或关闭图片透明色.
//
// hImage: 图片句柄.
//
// bEnable: 启用TRUE.
func X图片_启用透明色(图片句柄 int, 启用TRUE bool) {
	xImage_EnableTranColor.Call(uintptr(图片句柄), 炫彩工具类.BoolPtr(启用TRUE))
}

// 图片_启用自动销毁, 启用或关闭自动销毁, 当与UI元素关联时有效.
//
// hImage: 图片句柄.
//
// bEnable: 启用自动销毁TRUE.
func X图片_启用自动销毁(图片句柄 int, 启用自动销毁TRUE bool) {
	xImage_EnableAutoDestroy.Call(uintptr(图片句柄), 炫彩工具类.BoolPtr(启用自动销毁TRUE))
}

// 图片_启用居中, 启用或关闭图片居中显示，默认属性图片有效.
//
// hImage: 图片句柄.
//
// bCenter: 是否居中显示.
func X图片_启用居中(图片句柄 int, 是否居中显示 bool) {
	xImage_EnableCenter.Call(uintptr(图片句柄), 炫彩工具类.BoolPtr(是否居中显示))
}

// 图片_判断居中, 判断图片是否居中显示.
//
// hImage: 图片句柄.
func X图片_判断居中(图片句柄 int) bool {
	r, _, _ := xImage_IsCenter.Call(uintptr(图片句柄))
	return r != 0
}

// 图片_取绘制类型, 获取图片绘制类型, 返回: xcc.Image_Draw_Type_.
//
// hImage: 图片句柄.
func X图片_取绘制类型(图片句柄 int) 炫彩常量类.Image_Draw_Type_ {
	r, _, _ := xImage_GetDrawType.Call(uintptr(图片句柄))
	return 炫彩常量类.Image_Draw_Type_(r)
}

// 图片_取宽度.
//
// hImage: 图片句柄.
func X图片_取宽度(图片句柄 int) int32 {
	r, _, _ := xImage_GetWidth.Call(uintptr(图片句柄))
	return int32(r)
}

// 图片_取高度.
//
// hImage: 图片句柄.
func X图片_取高度(图片句柄 int) int32 {
	r, _, _ := xImage_GetHeight.Call(uintptr(图片句柄))
	return int32(r)
}

// 图片_取图片源.
//
// hImage: 图片句柄.
func X图片_取图片源(图片句柄 int) int {
	r, _, _ := xImage_GetImageSrc.Call(uintptr(图片句柄))
	return int(r)
}

// 图片_增加引用计数.
//
// hImage: 图片句柄.
func X图片_增加引用计数(图片句柄 int) {
	xImage_AddRef.Call(uintptr(图片句柄))
}

// 图片_释放引用计数, 释放引用计数, 当引用计数为0时, 自动销毁.
//
// hImage: 图片句柄.
func X图片_释放引用计数(图片句柄 int) {
	xImage_Release.Call(uintptr(图片句柄))
}

// 图片_取引用计数.
//
// hImage: 图片句柄.
func X图片_取引用计数(图片句柄 int) int32 {
	r, _, _ := xImage_GetRefCount.Call(uintptr(图片句柄))
	return int32(r)
}

// 图片_销毁, 强制销毁图片, 谨慎使用, 建议使用 XImage_Release() 释放.
//
// hImage: 图片句柄.
func X图片_销毁(图片句柄 int) {
	xImage_Destroy.Call(uintptr(图片句柄))
}

// 图片_加载从SVG.
//
// hSvg: SVG句柄.
func X图片_加载从SVG(SVG句柄 int) int {
	r, _, _ := xImage_LoadSvg.Call(uintptr(SVG句柄))
	return int(r)
}

// 图片_加载从SVG文件.
//
// pFileName: 文件名.
func X图片_加载从SVG文件(文件名 string) int {
	r, _, _ := xImage_LoadSvgFile.Call(炫彩工具类.StrPtr(文件名))
	return int(r)
}

// 图片_加载从SVG字符串.
//
// pString: 字符串.
func X图片_加载从SVG字符串(字符串 string) int {
	r, _, _ := xImage_LoadSvgString.Call(W2A(字符串))
	return int(r)
}

// 图片_取SVG, 返回SVG句柄.
//
// hImage: 图片句柄.
func X图片_取SVG(图片句柄 int) int {
	r, _, _ := xImage_GetSvg.Call(uintptr(图片句柄))
	return int(r)
}

// 图片_加载从SVG字符串W.
//
// pString: 字符串.
func X图片_加载从SVG字符串W(字符串 string) int {
	r, _, _ := xImage_LoadSvgStringW.Call(炫彩工具类.StrPtr(字符串))
	return int(r)
}

// 图片_加载从SVG字符串UTF8, 更推荐使用 xc.XImage_LoadSvgStringW.
//
// pString: 字符串.
func X图片_加载从SVG字符串UTF8(字符串 string) int {
	r, _, _ := xImage_LoadSvgStringUtf8.Call(X文本W到UTF8(字符串))
	return int(r)
}

// 图片_置缩放大小, 启用缩放属性后有效, 值大于0有效.
//
// hImage: 图片句柄.
//
// width: 宽度.
//
// height: 高度.
func X图片_置缩放大小(图片句柄 int, 宽度, 高度 int32) {
	xImage_SetScaleSize.Call(uintptr(图片句柄), uintptr(宽度), uintptr(高度))
}
