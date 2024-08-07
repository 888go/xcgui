package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"

	"github.com/888go/xcgui/xcc"
)

// 背景_创建, 创建背景管理器, 返回背景管理器句柄.

// ff:背景_创建
func X背景_创建() int {
	r, _, _ := xBkM_Create.Call()
	return int(r)
}

// 背景_销毁.
//
// hBkInfoM: 背景管理器句柄.

// ff:背景_销毁
// hBkInfoM:背景管理器句柄
func X背景_销毁(背景管理器句柄 int) int {
	r, _, _ := xBkM_Destroy.Call(uintptr(背景管理器句柄))
	return int(r)
}

// !废弃函数, 保留为了兼容旧版; 请使用 XBkM_SetInfo().
//
// 背景_置内容, 设置背景内容, 返回设置的背景对象数量.
//
// hBkInfoM: 背景管理器句柄.
//
// pText: 背景内容字符串.

// ff:废弃_XBkM_SetBkInfo
// pText:背景内容字符串
// hBkInfoM:背景管理器句柄
func X废弃_XBkM_SetBkInfo(背景管理器句柄 int, 背景内容字符串 string) int {
	r, _, _ := xBkM_SetBkInfo.Call(uintptr(背景管理器句柄), 炫彩工具类.StrPtr(背景内容字符串))
	return int(r)
}

// 背景_添加内容, 添加背景内容, 返回添加的背景对象数量.
//
// hBkInfoM: 背景管理器句柄.
//
// pText: 背景内容字符串.

// ff:背景_添加内容
// pText:背景内容字符串
// hBkInfoM:背景管理器句柄
func X背景_添加内容(背景管理器句柄 int, 背景内容字符串 string) int {
	r, _, _ := xBkM_AddInfo.Call(uintptr(背景管理器句柄), 炫彩工具类.StrPtr(背景内容字符串))
	return int(r)
}

// 背景_添加边框, 添加背景内容边框.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
//
// id: 背景对象ID, 可忽略(填0).

// ff:背景_添加边框
// id:
// width:
// color:
// nState:组合状态
// hBkInfoM:背景管理器句柄
func X背景_添加边框(背景管理器句柄 int, 组合状态 炫彩常量类.CombinedState, color, width, id int) int {
	r, _, _ := xBkM_AddBorder.Call(uintptr(背景管理器句柄), uintptr(组合状态), uintptr(color), uintptr(width), uintptr(id))
	return int(r)
}

// 背景_添加填充, 添加背景内容填充.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// id: 背景对象ID, 可忽略(填0).

// ff:背景_添加填充
// id:
// color:
// nState:组合状态
// hBkInfoM:背景管理器句柄
func X背景_添加填充(背景管理器句柄 int, 组合状态 炫彩常量类.CombinedState, color, id int) int {
	r, _, _ := xBkM_AddFill.Call(uintptr(背景管理器句柄), uintptr(组合状态), uintptr(color), uintptr(id))
	return int(r)
}

// 背景_添加图片, 添加背景内容图片.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
//
// id: 背景对象ID, 可忽略(填0).

// ff:背景_添加图片
// id:
// hImage:
// nState:组合状态
// hBkInfoM:背景管理器句柄
func X背景_添加图片(背景管理器句柄 int, 组合状态 炫彩常量类.CombinedState, hImage, id int) int {
	r, _, _ := xBkM_AddImage.Call(uintptr(背景管理器句柄), uintptr(组合状态), uintptr(hImage), uintptr(id))
	return int(r)
}

// 背景_取数量, 获取背景内容数量.
//
// hBkInfoM: 背景管理器句柄.

// ff:背景_取数量
// hBkInfoM:背景管理器句柄
func X背景_取数量(背景管理器句柄 int) int {
	r, _, _ := xBkM_GetCount.Call(uintptr(背景管理器句柄))
	return int(r)
}

// 背景_清空, 清空背景内容.
//
// hBkInfoM: 背景管理器句柄.

// ff:背景_清空
// hBkInfoM:背景管理器句柄
func X背景_清空(背景管理器句柄 int) int {
	r, _, _ := xBkM_Clear.Call(uintptr(背景管理器句柄))
	return int(r)
}

// 背景_绘制, 绘制背景内容.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.

// ff:背景_绘制
// pRect:
// hDraw:
// nState:组合状态
// hBkInfoM:背景管理器句柄
func X背景_绘制(背景管理器句柄 int, 组合状态 炫彩常量类.CombinedState, hDraw int, pRect *RECT) bool {
	r, _, _ := xBkM_Draw.Call(uintptr(背景管理器句柄), uintptr(组合状态), uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return r != 0
}

// 背景_绘制扩展, 绘制背景内容, 设置条件.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
//
// nStateEx: 当(nState)中包含(nStateEx)中的一个或多个状态时有效.
//
// 注解: 例如用来绘制列表项时, nState中包含项的状态(nStateEx)才会绘制, 避免列表项与元素背景叠加.

// ff:背景_绘制EX
// nStateEx:
// pRect:
// hDraw:
// nState:组合状态
// hBkInfoM:背景管理器句柄
func X背景_绘制EX(背景管理器句柄 int, 组合状态 炫彩常量类.CombinedState, hDraw int, pRect *RECT, nStateEx 炫彩常量类.CombinedState) bool {
	r, _, _ := xBkM_DrawEx.Call(uintptr(背景管理器句柄), uintptr(组合状态), uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nStateEx))
	return r != 0
}

// 背景_启用自动销毁, 是否自动销毁.
//
// hBkInfoM: 背景管理器句柄.
//
// bEnable: 是否启用.

// ff:背景_启用自动销毁
// bEnable:是否启用
// hBkInfoM:背景管理器句柄
func X背景_启用自动销毁(背景管理器句柄 int, 是否启用 bool) int {
	r, _, _ := xBkM_EnableAutoDestroy.Call(uintptr(背景管理器句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 背景_增加引用计数.
//
// hBkInfoM: 背景管理器句柄.

// ff:背景_增加引用计数
// hBkInfoM:背景管理器句柄
func X背景_增加引用计数(背景管理器句柄 int) int {
	r, _, _ := xBkM_AddRef.Call(uintptr(背景管理器句柄))
	return int(r)
}

// 背景_释放引用计数.
//
// hBkInfoM: 背景管理器句柄.

// ff:背景_释放引用计数
// hBkInfoM:背景管理器句柄
func X背景_释放引用计数(背景管理器句柄 int) int {
	r, _, _ := xBkM_Release.Call(uintptr(背景管理器句柄))
	return int(r)
}

// 背景_取引用计数.
//
// hBkInfoM: 背景管理器句柄.

// ff:背景_取引用计数
// hBkInfoM:背景管理器句柄
func X背景_取引用计数(背景管理器句柄 int) int {
	r, _, _ := xBkM_GetRefCount.Call(uintptr(背景管理器句柄))
	return int(r)
}

// 背景_取引用计数, 设置背景内容, 返回设置的背景对象数量.
//
// hBkInfoM: 背景管理器句柄.
//
// pText: 背景内容字符串.

// ff:背景_设置内容
// pText:背景内容字符串
// hBkInfoM:背景管理器句柄
func X背景_设置内容(背景管理器句柄 int, 背景内容字符串 string) int {
	r, _, _ := xBkM_SetInfo.Call(uintptr(背景管理器句柄), 炫彩工具类.StrPtr(背景内容字符串))
	return int(r)
}

// 背景_取指定状态文本颜色.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// color: 接收返回的ABGR 颜色.

// ff:背景_取指定状态文本颜色
// color:
// nState:组合状态
// hBkInfoM:背景管理器句柄
func X背景_取指定状态文本颜色(背景管理器句柄 int, 组合状态 炫彩常量类.CombinedState, color *int) bool {
	r, _, _ := xBkM_GetStateTextColor.Call(uintptr(背景管理器句柄), uintptr(组合状态), uintptr(unsafe.Pointer(color)))
	return r != 0
}

// 背景_取背景对象, 返回BkObj对象句柄.
//
// hBkInfoM: 背景管理器句柄.
//
// id: 背景对象ID.

// ff:背景_取背景对象
// id:背景对象ID
// hBkInfoM:背景管理器句柄
func X背景_取背景对象(背景管理器句柄 int, 背景对象ID int) int {
	r, _, _ := xBkM_GetObject.Call(uintptr(背景管理器句柄), uintptr(背景对象ID))
	return int(r)
}
