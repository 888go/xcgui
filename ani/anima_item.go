package 炫彩动画类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
)

// AnimaItem 动画项.
type AnimaItem struct {
	炫彩对象基类.ObjectBase
}

// 动画项_启用完成释放, 当动画项完成后自动释放.
//
// 例如对多个动画序列进行渐近式延迟, 在动画序列头标添加延时项(时间差), 当延时项完成时自动释放, 后续动画循环就形成一种时间差(因为对齐的时间差销毁了, 他们永远无法对齐时间).
//
// bEnable: 是否启用.
func (a *AnimaItem) X启用完成释放(是否启用 bool) int {
	return 炫彩基类.X动画项_启用完成释放(a.Handle, 是否启用)
}

// 动画项_置回调.
//
// callback: 回调函数.
func (a *AnimaItem) X置回调(回调函数 interface{}) int {
	return 炫彩基类.X动画项_置回调(a.Handle, 回调函数)
}

// 动画项_置用户数据.
//
// nUserData: 用户数据.
func (a *AnimaItem) X置用户数据(用户数据 int) int {
	return 炫彩基类.X动画项_置用户数据(a.Handle, 用户数据)
}

// 动画项_取用户数据, 返回用户数据.
func (a *AnimaItem) X取用户数据() int {
	return 炫彩基类.X动画项_取用户数据(a.Handle)
}

// 动画项_启用自动销毁.
//
// bEnable: 是否启用.
func (a *AnimaItem) X启用自动销毁(是否启用 bool) int {
	return 炫彩基类.X动画项_启用自动销毁(a.Handle, 是否启用)
}
