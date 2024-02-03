package bkmanager

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/res"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// BkManager 背景管理器.
type BkManager struct {
	objectbase.ObjectBase
}

// 创建背景管理器, 创建背景管理器.
func New() *BkManager {
	p := &BkManager{}
	p.SetHandle(xc.XBkM_Create())
	return p
}

// 创建背景管理器并按句柄
func NewByHandle(handle int) *BkManager {
	p := &BkManager{}
	p.SetHandle(handle)
	return p
}

// 创建背景管理器并按名称, 失败返回nil.
func NewByName(name string) *BkManager {
	handle := res.GetBkM(name)
	if handle > 0 {
		p := &BkManager{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 作废SetBkInfo
//
// !废弃函数, 保留为了兼容旧版; 请使用 SetInfo().
//
// 置内容, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容
func (b *BkManager) SetBkInfo(pText string) int {
	return xc.XBkM_SetBkInfo(b.Handle, pText)
}

// 销毁.
func (b *BkManager) Destroy() int {
	return xc.XBkM_Destroy(b.Handle)
}

// 添加内容, 添加背景内容, 返回添加的背景对象数量.
//
// pText: 背景内容
func (b *BkManager) AddInfo(pText string) int {
	return xc.XBkM_AddInfo(b.Handle, pText)
}

// 添加边框, 添加背景内容边框.
//
// nState: 组合状态.
//
// color: ABGR颜色.
//
// width: 线宽.
//
// id: 背景对象ID, 可忽略(填0).
func (b *BkManager) AddBorder(nState xcc.CombinedState, color, width, id int) int {
	return xc.XBkM_AddBorder(b.Handle, nState, color, width, id)
}

// 添加填充, 添加背景内容填充.
//
// nState: 组合状态.
//
// color: ABGR颜色.
//
// id: 背景对象ID, 可忽略(填0).
func (b *BkManager) AddFill(nState xcc.CombinedState, color, id int) int {
	return xc.XBkM_AddFill(b.Handle, nState, color, id)
}

// 添加图片, 添加背景内容图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
//
// id: 背景对象ID, 可忽略(填0).
func (b *BkManager) AddImage(nState xcc.CombinedState, hImage, id int) int {
	return xc.XBkM_AddImage(b.Handle, nState, hImage, id)
}

// 取数量, 获取背景内容数量.
func (b *BkManager) GetCount() int {
	return xc.XBkM_GetCount(b.Handle)
}

// 清空, 清空背景内容.
func (b *BkManager) Clear() int {
	return xc.XBkM_Clear(b.Handle)
}

// 绘制, 绘制背景内容.
//
// nState: 组合状态.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func (b *BkManager) Draw(nState xcc.CombinedState, hDraw int, pRect *xc.RECT) bool {
	return xc.XBkM_Draw(b.Handle, nState, hDraw, pRect)
}

// 绘制扩展, 绘制背景内容, 设置条件.
//
// nState: 组合状态.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
//
// nStateEx: 状态Ex
//
// 注解: 例如用来绘制列表项时, nState中包含项的状态(nStateEx)才会绘制, 避免列表项与元素背景叠加.
func (b *BkManager) DrawEx(nState xcc.CombinedState, hDraw int, pRect *xc.RECT, nStateEx xcc.CombinedState) bool {
	return xc.XBkM_DrawEx(b.Handle, nState, hDraw, pRect, nStateEx)
}

// 启用自动销毁, 是否自动销毁.
//
// bEnable: 是否启用.
func (b *BkManager) EnableAutoDestroy(bEnable bool) int {
	return xc.XBkM_EnableAutoDestroy(b.Handle, bEnable)
}

// 增加引用计数.
func (b *BkManager) AddRef() int {
	return xc.XBkM_AddRef(b.Handle)
}

// 释放引用计数.
func (b *BkManager) Release() int {
	return xc.XBkM_Release(b.Handle)
}

// 取引用计数.
func (b *BkManager) GetRefCount() int {
	return xc.XBkM_GetRefCount(b.Handle)
}

// 取引用计数, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容
func (b *BkManager) SetInfo(pText string) int {
	return xc.XBkM_SetInfo(b.Handle, pText)
}

// 取指定状态文本颜色.
//
// nState: 组合状态.
//
// color: 返回ABGR颜色.
func (b *BkManager) GetStateTextColor(nState xcc.CombinedState, color *int) bool {
	return xc.XBkM_GetStateTextColor(b.Handle, nState, color)
}

// 取背景对象, 返回BkObj对象句柄.
//
// id: 背景对象ID.
func (b *BkManager) GetObject(id int) int {
	return xc.XBkM_GetObject(b.Handle, id)
}
