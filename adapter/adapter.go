package adapter

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// 数据适配器.
type adapter struct {
	objectbase.ObjectBase
}

// 增加引用计数.
func (a *adapter) AddRef() int {
	return xc.XAd_AddRef(a.Handle)
}

// 释放引用计数.
func (a *adapter) Release() int {
	return xc.XAd_Release(a.Handle)
}

// 取引用计数.
func (a *adapter) GetRefCount() int {
	return xc.XAd_GetRefCount(a.Handle)
}

// 销毁.
func (a *adapter) Destroy() int {
	return xc.XAd_Destroy(a.Handle)
}

// 启用自动销毁.
// // bEnable: 是否启用.
func (a *adapter) EnableAutoDestroy(bEnable bool) int {
	return xc.XAd_EnableAutoDestroy(a.Handle, bEnable)
}
