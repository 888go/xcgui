package bkmanager//bm:炫彩背景管理器类

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

// New 背景_创建, 创建背景管理器.
// ff:创建
func New() *BkManager {
	p := &BkManager{}
	p.SetHandle(xc.XBkM_Create())
	return p
}

// NewByHandle 从句柄创建背景管理器对象.
// ff:创建并按句柄
// handle:
func NewByHandle(handle int) *BkManager {
	p := &BkManager{}
	p.SetHandle(handle)
	return p
}

// NewByName 从name创建背景管理器对象, 失败返回nil.
// ff:创建并按名称
// name:
func NewByName(name string) *BkManager {
	handle := res.GetBkM(name)
	if handle > 0 {
		p := &BkManager{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// Deprecated
//
// !废弃函数, 保留为了兼容旧版; 请使用 SetInfo().
//
// 背景_置内容, 设置背景内容, 返回设置的背景对象数量.
//.
// ff:作废SetBkInfo
// pText:背景内容
func (b *BkManager) SetBkInfo(pText string) int {
	return xc.XBkM_SetBkInfo(b.Handle, pText)
}

// 背景_销毁.
// ff:销毁
func (b *BkManager) Destroy() int {
	return xc.XBkM_Destroy(b.Handle)
}

// 背景_添加内容, 添加背景内容, 返回添加的背景对象数量.
//.
// ff:添加内容
// pText:背景内容
func (b *BkManager) AddInfo(pText string) int {
	return xc.XBkM_AddInfo(b.Handle, pText)
}

// 背景_添加边框, 添加背景内容边框.
//.
//.
//.
//, 可忽略(填0).
// ff:添加边框
// nState:组合状态
// color:
// width:
// id:
func (b *BkManager) AddBorder(nState xcc.CombinedState, color, width, id int) int {
	return xc.XBkM_AddBorder(b.Handle, nState, color, width, id)
}

// 背景_添加填充, 添加背景内容填充.
//.
//.
//, 可忽略(填0).
// ff:添加填充
// nState:组合状态
// color:
// id:
func (b *BkManager) AddFill(nState xcc.CombinedState, color, id int) int {
	return xc.XBkM_AddFill(b.Handle, nState, color, id)
}

// 背景_添加图片, 添加背景内容图片.
//.
//.
//, 可忽略(填0).
// ff:添加图片
// nState:组合状态
// hImage:
// id:
func (b *BkManager) AddImage(nState xcc.CombinedState, hImage, id int) int {
	return xc.XBkM_AddImage(b.Handle, nState, hImage, id)
}

// 背景_取数量, 获取背景内容数量.
// ff:取数量
func (b *BkManager) GetCount() int {
	return xc.XBkM_GetCount(b.Handle)
}

// 背景_清空, 清空背景内容.
// ff:清空
func (b *BkManager) Clear() int {
	return xc.XBkM_Clear(b.Handle)
}

// 背景_绘制, 绘制背景内容.
//.
//.
//.
// ff:绘制
// nState:组合状态
// hDraw:
// pRect:
func (b *BkManager) Draw(nState xcc.CombinedState, hDraw int, pRect *xc.RECT) bool {
	return xc.XBkM_Draw(b.Handle, nState, hDraw, pRect)
}

// 背景_绘制扩展, 绘制背景内容, 设置条件.
//.
//.
//.
//(nState)中包含(nStateEx)中的一个或多个状态时有效.
//
// 注解: 例如用来绘制列表项时, nState中包含项的状态(nStateEx)才会绘制, 避免列表项与元素背景叠加.
// ff:绘制EX
// nState:组合状态
// hDraw:
// pRect:
// nStateEx:
func (b *BkManager) DrawEx(nState xcc.CombinedState, hDraw int, pRect *xc.RECT, nStateEx xcc.CombinedState) bool {
	return xc.XBkM_DrawEx(b.Handle, nState, hDraw, pRect, nStateEx)
}

// 背景_启用自动销毁, 是否自动销毁.
//.
// ff:启用自动销毁
// bEnable:是否启用
func (b *BkManager) EnableAutoDestroy(bEnable bool) int {
	return xc.XBkM_EnableAutoDestroy(b.Handle, bEnable)
}

// 背景_增加引用计数.
// ff:增加引用计数
func (b *BkManager) AddRef() int {
	return xc.XBkM_AddRef(b.Handle)
}

// 背景_释放引用计数.
// ff:释放引用计数
func (b *BkManager) Release() int {
	return xc.XBkM_Release(b.Handle)
}

// 背景_取引用计数.
// ff:取引用计数
func (b *BkManager) GetRefCount() int {
	return xc.XBkM_GetRefCount(b.Handle)
}

// 背景_取引用计数, 设置背景内容, 返回设置的背景对象数量.
//.
// ff:设置背景内容
// pText:背景内容
func (b *BkManager) SetInfo(pText string) int {
	return xc.XBkM_SetInfo(b.Handle, pText)
}

// 背景_取指定状态文本颜色.
//.
//.
// ff:取指定状态文本颜色
// nState:组合状态
// color:
func (b *BkManager) GetStateTextColor(nState xcc.CombinedState, color *int) bool {
	return xc.XBkM_GetStateTextColor(b.Handle, nState, color)
}

// 背景_取背景对象, 返回BkObj对象句柄.
//.
// ff:取背景对象
// id:背景对象ID
func (b *BkManager) GetObject(id int) int {
	return xc.XBkM_GetObject(b.Handle, id)
}
