package 炫彩数据适配器类 //bm:炫彩数据适配器类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
)

// 数据适配器.
type adapter struct {
	炫彩对象基类.ObjectBase
}

// 数据适配器_增加引用计数.

// ff:增加引用计数
func (a *adapter) X增加引用计数() int {
	return 炫彩基类.X数据适配器_增加引用计数(a.Handle)
}

// 数据适配器_释放引用计数.

// ff:释放引用计数
func (a *adapter) X释放引用计数() int {
	return 炫彩基类.X数据适配器_释放引用计数(a.Handle)
}

// 数据适配器_取引用计数.

// ff:取引用计数
func (a *adapter) X取引用计数() int {
	return 炫彩基类.X数据适配器_取引用计数(a.Handle)
}

// 数据适配器_销毁.

// ff:销毁
func (a *adapter) X销毁() int {
	return 炫彩基类.X数据适配器_销毁(a.Handle)
}

// 数据适配器_启用自动销毁.
//
// bEnable: 是否启用.

// ff:启用自动销毁
// bEnable:是否启用
func (a *adapter) X启用自动销毁(是否启用 bool) int {
	return 炫彩基类.X数据适配器_启用自动销毁(a.Handle, 是否启用)
}
