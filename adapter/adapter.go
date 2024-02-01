package adapter

import (
	"e.coding.net/gogit/go/xcgui/objectbase"
	"e.coding.net/gogit/go/xcgui/xc"
)

// 数据适配器.
type adapter struct {
	objectbase.ObjectBase
}

// 数据适配器_增加引用计数.
func (a *adapter) AddRef() int {
	return xc.XAd_AddRef(a.I句柄)
}

// 数据适配器_释放引用计数.
func (a *adapter) Release() int {
	return xc.XAd_Release(a.I句柄)
}

// 数据适配器_取引用计数.
func (a *adapter) GetRefCount() int {
	return xc.XAd_GetRefCount(a.I句柄)
}

// 数据适配器_销毁.
func (a *adapter) Destroy() int {
	return xc.XAd_Destroy(a.I句柄)
}

// 数据适配器_启用自动销毁.
//
// bEnable: 是否启用.
func (a *adapter) EnableAutoDestroy(bEnable bool) int {
	return xc.XAd_EnableAutoDestroy(a.I句柄, bEnable)
}
