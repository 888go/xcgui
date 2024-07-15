package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"

	"github.com/888go/xcgui/xcc"
)

// 背景对象_置外间距, 外间距与对齐标识(BkObject_Align_Flag_)互相依赖.
//
// hObj: 背景对象句柄.
//
// left: 左边间距.
//
// top: 顶部间距.
//
// right: 右边间距.
//
// bottom: 底部间距.

// ff:背景对象_置外间距
// bottom:底部间距
// right:右边间距
// top:顶部间距
// left:左边间距
// hObj:背景对象句柄
func X背景对象_置外间距(背景对象句柄 int, 左边间距 int, 顶部间距 int, 右边间距 int, 底部间距 int) int {
	r, _, _ := xBkObj_SetMargin.Call(uintptr(背景对象句柄), uintptr(左边间距), uintptr(顶部间距), uintptr(右边间距), uintptr(底部间距))
	return int(r)
}

// 背景对象_置对齐.
//
// hObj: 背景对象句柄.
//
// nFlags: 对齐方式: xcc.BkObject_Align_Flag_.

// ff:背景对象_置对齐
// nFlags:对齐方式
// hObj:背景对象句柄
func X背景对象_置对齐(背景对象句柄 int, 对齐方式 炫彩常量类.BkObject_Align_Flag_) int {
	r, _, _ := xBkObj_SetAlign.Call(uintptr(背景对象句柄), uintptr(对齐方式))
	return int(r)
}

// 背景对象_置图片.
//
// hObj: 背景对象句柄.
//
// hImage: 图片句柄.

// ff:背景对象_置图片
// hImage:图片句柄
// hObj:背景对象句柄
func X背景对象_置图片(背景对象句柄 int, 图片句柄 int) int {
	r, _, _ := xBkObj_SetImage.Call(uintptr(背景对象句柄), uintptr(图片句柄))
	return int(r)
}

// 背景对象_置旋转.
//
// hObj: 背景对象句柄.
//
// angle: 旋转角度.

// ff:背景对象_置旋转
// angle:旋转角度
// hObj:背景对象句柄
func X背景对象_置旋转(背景对象句柄 int, 旋转角度 float32) int {
	r, _, _ := xBkObj_SetRotate.Call(uintptr(背景对象句柄), 炫彩工具类.Float32Ptr(旋转角度))
	return int(r)
}

// 背景对象_置填充颜色.
//
// hObj: 背景对象句柄.
//
// color: ABGR 颜色值.

// ff:背景对象_置填充颜色
// color:ABGR颜色值
// hObj:背景对象句柄
func X背景对象_置填充颜色(背景对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xBkObj_SetFillColor.Call(uintptr(背景对象句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 背景对象_置边框宽度.
//
// hObj: 背景对象句柄.
//
// width: 宽度.

// ff:背景对象_置边框宽度
// width:宽度
// hObj:背景对象句柄
func X背景对象_置边框宽度(背景对象句柄 int, 宽度 int) int {
	r, _, _ := xBkObj_SetBorderWidth.Call(uintptr(背景对象句柄), uintptr(宽度))
	return int(r)
}

// 背景对象_置边框颜色.
//
// hObj: 背景对象句柄.
//
// color: ABGR 颜色值.

// ff:背景对象_置边框颜色
// color:ABGR颜色值
// hObj:背景对象句柄
func X背景对象_置边框颜色(背景对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xBkObj_SetBorderColor.Call(uintptr(背景对象句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 背景对象_置矩形圆角.
//
// hObj: 背景对象句柄.
//
// leftTop: 左上角.
//
// leftBottom: 左下角.
//
// rightTop: 右上角.
//
// rightBottom: 右下角.

// ff:背景对象_置矩形圆角
// rightBottom:右下角
// rightTop:右上角
// leftBottom:左下角
// leftTop:左上角
// hObj:背景对象句柄
func X背景对象_置矩形圆角(背景对象句柄 int, 左上角 int, 左下角 int, 右上角 int, 右下角 int) int {
	r, _, _ := xBkObj_SetRectRoundAngle.Call(uintptr(背景对象句柄), uintptr(左上角), uintptr(左下角), uintptr(右上角), uintptr(右下角))
	return int(r)
}

// 背景对象_启用填充, 启用绘制填充.
//
// hObj: 背景对象句柄.
//
// bEnable: 是否启用.

// ff:背景对象_启用填充
// bEnable:是否启用
// hObj:背景对象句柄
func X背景对象_启用填充(背景对象句柄 int, 是否启用 bool) int {
	r, _, _ := xBkObj_EnableFill.Call(uintptr(背景对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 背景对象_启用边框, 启用绘制边框.
//
// hObj: 背景对象句柄.
//
// bEnable: 是否启用.

// ff:背景对象_启用边框
// bEnable:是否启用
// hObj:背景对象句柄
func X背景对象_启用边框(背景对象句柄 int, 是否启用 bool) int {
	r, _, _ := xBkObj_EnableBorder.Call(uintptr(背景对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 背景对象_置文本.
//
// hObj: 背景对象句柄.
//
// pText: 文本字符串.

// ff:背景对象_置文本
// pText:文本字符串
// hObj:背景对象句柄
func X背景对象_置文本(背景对象句柄 int, 文本字符串 string) int {
	r, _, _ := xBkObj_SetText.Call(uintptr(背景对象句柄), 炫彩工具类.StrPtr(文本字符串))
	return int(r)
}

// 背景对象_置字体.
//
// hObj: 背景对象句柄.
//
// hFont: 字体句柄.

// ff:背景对象_置字体
// hFont:字体句柄
// hObj:背景对象句柄
func X背景对象_置字体(背景对象句柄 int, 字体句柄 int) int {
	r, _, _ := xBkObj_SetFont.Call(uintptr(背景对象句柄), uintptr(字体句柄))
	return int(r)
}

// 背景对象_置文本对齐.
//
// hObj: 背景对象句柄.
//
// nAlign: 文本对齐方式: xcc.TextFormatFlag_, xcc.TextAlignFlag_, xcc.TextTrimming_.

// ff:背景对象_置文本对齐
// nAlign:文本对齐方式
// hObj:背景对象句柄
func X背景对象_置文本对齐(背景对象句柄 int, 文本对齐方式 炫彩常量类.TextFormatFlag_) int {
	r, _, _ := xBkObj_SetTextAlign.Call(uintptr(背景对象句柄), uintptr(文本对齐方式))
	return int(r)
}

// 背景对象_取外间距.
//
// hObj: 背景对象句柄.
//
// pMargin: 接收返回外间距.

// ff:背景对象_取外间距
// pMargin:接收返回外间距
// hObj:背景对象句柄
func X背景对象_取外间距(背景对象句柄 int, 接收返回外间距 *RECT) int {
	r, _, _ := xBkObj_GetMargin.Call(uintptr(背景对象句柄), uintptr(unsafe.Pointer(接收返回外间距)))
	return int(r)
}

// 背景对象_取对齐, 返回对齐标识: xcc.BkObject_Align_Flag_.
//
// hObj: 背景对象句柄.

// ff:背景对象_取对齐
// hObj:背景对象句柄
func X背景对象_取对齐(背景对象句柄 int) 炫彩常量类.BkObject_Align_Flag_ {
	r, _, _ := xBkObj_GetAlign.Call(uintptr(背景对象句柄))
	return 炫彩常量类.BkObject_Align_Flag_(r)
}

// 背景对象_取图片, 返回图片句柄.
//
// hObj: 背景对象句柄.

// ff:背景对象_取图片
// hObj:背景对象句柄
func X背景对象_取图片(背景对象句柄 int) int {
	r, _, _ := xBkObj_GetImage.Call(uintptr(背景对象句柄))
	return int(r)
}

// 背景对象_取旋转角度, 返回旋转角度.
//
// hObj: 背景对象句柄.

// ff:背景对象_取旋转角度
// hObj:背景对象句柄
func X背景对象_取旋转角度(背景对象句柄 int) int {
	r, _, _ := xBkObj_GetRotate.Call(uintptr(背景对象句柄))
	return int(r)
}

// 背景对象_取填充色, 返回ABGR填充色.
//
// hObj: 背景对象句柄.

// ff:背景对象_取填充色
// hObj:背景对象句柄
func X背景对象_取填充色(背景对象句柄 int) int {
	r, _, _ := xBkObj_GetFillColor.Call(uintptr(背景对象句柄))
	return int(r)
}

// 背景对象_取边框色, 返回ABGR边框色.
//
// hObj: 背景对象句柄.

// ff:背景对象_取边框色
// hObj:背景对象句柄
func X背景对象_取边框色(背景对象句柄 int) int {
	r, _, _ := xBkObj_GetBorderColor.Call(uintptr(背景对象句柄))
	return int(r)
}

// 背景对象_取边框宽度, 返回边框宽度.
//
// hObj: 背景对象句柄.

// ff:背景对象_取边框宽度
// hObj:背景对象句柄
func X背景对象_取边框宽度(背景对象句柄 int) int {
	r, _, _ := xBkObj_GetBorderWidth.Call(uintptr(背景对象句柄))
	return int(r)
}

// 背景对象_取矩形圆角.
//
// hObj: 背景对象句柄.
//
// pRect: 接收返回圆角大小.

// ff:背景对象_取矩形圆角
// pRect:接收返回圆角大小
// hObj:背景对象句柄
func X背景对象_取矩形圆角(背景对象句柄 int, 接收返回圆角大小 *RECT) int {
	r, _, _ := xBkObj_GetRectRoundAngle.Call(uintptr(背景对象句柄), uintptr(unsafe.Pointer(接收返回圆角大小)))
	return int(r)
}

// 背景对象_是否填充.
//
// hObj: 背景对象句柄.

// ff:背景对象_是否填充
// hObj:背景对象句柄
func X背景对象_是否填充(背景对象句柄 int) bool {
	r, _, _ := xBkObj_IsFill.Call(uintptr(背景对象句柄))
	return r != 0
}

// 背景对象_是否边框.
//
// hObj: 背景对象句柄.

// ff:背景对象_是否边框
// hObj:背景对象句柄
func X背景对象_是否边框(背景对象句柄 int) bool {
	r, _, _ := xBkObj_IsBorder.Call(uintptr(背景对象句柄))
	return r != 0
}

// 背景对象_取文本.
//
// hObj: 背景对象句柄.

// ff:背景对象_取文本
// hObj:背景对象句柄
func X背景对象_取文本(背景对象句柄 int) string {
	r, _, _ := xBkObj_GetText.Call(uintptr(背景对象句柄))
	return 炫彩工具类.UintPtrToString(r)
}

// 背景对象_取字体, 返回字体句柄.
//
// hObj: 背景对象句柄.

// ff:背景对象_取字体
// hObj:背景对象句柄
func X背景对象_取字体(背景对象句柄 int) int {
	r, _, _ := xBkObj_GetFont.Call(uintptr(背景对象句柄))
	return int(r)
}

// 背景对象_取文本对齐, 返回文本对齐方式: xcc.TextFormatFlag_.
//
// hObj: 背景对象句柄.

// ff:背景对象_取文本对齐
// hObj:背景对象句柄
func X背景对象_取文本对齐(背景对象句柄 int) 炫彩常量类.TextFormatFlag_ {
	r, _, _ := xBkObj_GetTextAlign.Call(uintptr(背景对象句柄))
	return 炫彩常量类.TextFormatFlag_(r)
}
