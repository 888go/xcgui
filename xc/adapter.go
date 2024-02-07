package 炫彩基类

import (
	"github.com/888go/xcgui/common"
)

// 数据适配器_增加引用计数.
//
// hAdapter: 数据适配器句柄.
func X数据适配器_增加引用计数(数据适配器句柄 int) int {
	r, _, _ := xAd_AddRef.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器_释放引用计数.
//
// hAdapter: 数据适配器句柄.
func X数据适配器_释放引用计数(数据适配器句柄 int) int {
	r, _, _ := xAd_Release.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器_取引用计数.
//
// hAdapter: 数据适配器句柄.
func X数据适配器_取引用计数(数据适配器句柄 int) int {
	r, _, _ := xAd_GetRefCount.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器_销毁.
//
// hAdapter: 数据适配器句柄.
func X数据适配器_销毁(数据适配器句柄 int) int {
	r, _, _ := xAd_Destroy.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器_启用自动销毁.
//
// hAdapter: 数据适配器句柄.
//
// bEnable: 是否启用.
func X数据适配器_启用自动销毁(数据适配器句柄 int, 是否启用 bool) int {
	r, _, _ := xAd_EnableAutoDestroy.Call(uintptr(数据适配器句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}
