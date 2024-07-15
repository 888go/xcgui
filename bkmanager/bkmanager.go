package 炫彩背景管理器类 //bm:炫彩背景管理器类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/res"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// BkManager 背景管理器.
type BkManager struct {
	炫彩对象基类.ObjectBase
}

// New 背景_创建, 创建背景管理器.

// ff:创建
func X创建() *BkManager {
	p := &BkManager{}
	p.X设置句柄(炫彩基类.X背景_创建())
	return p
}

// NewByHandle 从句柄创建背景管理器对象.

// ff:创建并按句柄
// handle:
func X创建并按句柄(handle int) *BkManager {
	p := &BkManager{}
	p.X设置句柄(handle)
	return p
}

// NewByName 从name创建背景管理器对象, 失败返回nil.

// ff:创建并按名称
// name:
func X创建并按名称(name string) *BkManager {
	handle := 炫彩资源类.X取背景管理器(name)
	if handle > 0 {
		p := &BkManager{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// Deprecated
//
// !废弃函数, 保留为了兼容旧版; 请使用 SetInfo().
//
// 背景_置内容, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.

// ff:作废SetBkInfo
// pText:背景内容
func (b *BkManager) X作废SetBkInfo(背景内容 string) int {
	return 炫彩基类.X废弃_XBkM_SetBkInfo(b.Handle, 背景内容)
}

// 背景_销毁.

// ff:销毁
func (b *BkManager) X销毁() int {
	return 炫彩基类.X背景_销毁(b.Handle)
}

// 背景_添加内容, 添加背景内容, 返回添加的背景对象数量.
//
// pText: 背景内容字符串.

// ff:添加内容
// pText:背景内容
func (b *BkManager) X添加内容(背景内容 string) int {
	return 炫彩基类.X背景_添加内容(b.Handle, 背景内容)
}

// 背景_添加边框, 添加背景内容边框.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
//
// id: 背景对象ID, 可忽略(填0).

// ff:添加边框
// id:
// width:
// color:
// nState:组合状态
func (b *BkManager) X添加边框(组合状态 炫彩常量类.CombinedState, color, width, id int) int {
	return 炫彩基类.X背景_添加边框(b.Handle, 组合状态, color, width, id)
}

// 背景_添加填充, 添加背景内容填充.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// id: 背景对象ID, 可忽略(填0).

// ff:添加填充
// id:
// color:
// nState:组合状态
func (b *BkManager) X添加填充(组合状态 炫彩常量类.CombinedState, color, id int) int {
	return 炫彩基类.X背景_添加填充(b.Handle, 组合状态, color, id)
}

// 背景_添加图片, 添加背景内容图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
//
// id: 背景对象ID, 可忽略(填0).

// ff:添加图片
// id:
// hImage:
// nState:组合状态
func (b *BkManager) X添加图片(组合状态 炫彩常量类.CombinedState, hImage, id int) int {
	return 炫彩基类.X背景_添加图片(b.Handle, 组合状态, hImage, id)
}

// 背景_取数量, 获取背景内容数量.

// ff:取数量
func (b *BkManager) X取数量() int {
	return 炫彩基类.X背景_取数量(b.Handle)
}

// 背景_清空, 清空背景内容.

// ff:清空
func (b *BkManager) X清空() int {
	return 炫彩基类.X背景_清空(b.Handle)
}

// 背景_绘制, 绘制背景内容.
//
// nState: 组合状态.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.

// ff:绘制
// pRect:
// hDraw:
// nState:组合状态
func (b *BkManager) X绘制(组合状态 炫彩常量类.CombinedState, hDraw int, pRect *炫彩基类.RECT) bool {
	return 炫彩基类.X背景_绘制(b.Handle, 组合状态, hDraw, pRect)
}

// 背景_绘制扩展, 绘制背景内容, 设置条件.
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

// ff:绘制EX
// nStateEx:
// pRect:
// hDraw:
// nState:组合状态
func (b *BkManager) X绘制EX(组合状态 炫彩常量类.CombinedState, hDraw int, pRect *炫彩基类.RECT, nStateEx 炫彩常量类.CombinedState) bool {
	return 炫彩基类.X背景_绘制EX(b.Handle, 组合状态, hDraw, pRect, nStateEx)
}

// 背景_启用自动销毁, 是否自动销毁.
//
// bEnable: 是否启用.

// ff:启用自动销毁
// bEnable:是否启用
func (b *BkManager) X启用自动销毁(是否启用 bool) int {
	return 炫彩基类.X背景_启用自动销毁(b.Handle, 是否启用)
}

// 背景_增加引用计数.

// ff:增加引用计数
func (b *BkManager) X增加引用计数() int {
	return 炫彩基类.X背景_增加引用计数(b.Handle)
}

// 背景_释放引用计数.

// ff:释放引用计数
func (b *BkManager) X释放引用计数() int {
	return 炫彩基类.X背景_释放引用计数(b.Handle)
}

// 背景_取引用计数.

// ff:取引用计数
func (b *BkManager) X取引用计数() int {
	return 炫彩基类.X背景_取引用计数(b.Handle)
}

// 背景_取引用计数, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.

// ff:设置背景内容
// pText:背景内容
func (b *BkManager) X设置背景内容(背景内容 string) int {
	return 炫彩基类.X背景_设置内容(b.Handle, 背景内容)
}

// 背景_取指定状态文本颜色.
//
// nState: 组合状态.
//
// color: 接收返回的ABGR 颜色.

// ff:取指定状态文本颜色
// color:
// nState:组合状态
func (b *BkManager) X取指定状态文本颜色(组合状态 炫彩常量类.CombinedState, color *int) bool {
	return 炫彩基类.X背景_取指定状态文本颜色(b.Handle, 组合状态, color)
}

// 背景_取背景对象, 返回BkObj对象句柄.
//
// id: 背景对象ID.

// ff:取背景对象
// id:背景对象ID
func (b *BkManager) X取背景对象(背景对象ID int) int {
	return 炫彩基类.X背景_取背景对象(b.Handle, 背景对象ID)
}
