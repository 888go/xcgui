package 炫彩背景对象类

import (
	"github.com/888go/xcgui/bkmanager"
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// BkObj 背景对象.
type BkObj struct {
	炫彩对象基类.ObjectBase
}

// NewByHandle 从BkObj句柄创建BkObj对象.
func X创建并按句柄(handle int) *BkObj {
	p := &BkObj{}
	p.X设置句柄(handle)
	return p
}

// NewByBkm 从BkManager对象创建BkObj对象, 失败返回nil.
//
// bkm: 背景管理器对象.
//
// id: 背景对象ID.
func X创建并按背景管理器对象(背景管理器对象 *炫彩背景管理器类.BkManager, 背景对象ID int) *BkObj {
	handle := 背景管理器对象.X取背景对象(背景对象ID)
	if handle == 0 {
		return nil
	}
	p := &BkObj{}
	p.X设置句柄(handle)
	return p
}

// NewByBkmHandle 从BkManager句柄创建BkObj对象, 失败返回nil.
//
// hBkm: 背景管理器句柄.
//
// id: 背景对象ID.
func X创建并按背景管理器对象句柄(背景管理器句柄, 背景对象ID int) *BkObj {
	handle := 炫彩基类.X背景_取背景对象(背景管理器句柄, 背景对象ID)
	if handle == 0 {
		return nil
	}
	p := &BkObj{}
	p.X设置句柄(handle)
	return p
}

// 背景对象_置外间距, 外间距与对齐标识(BkObject_Align_Flag_)互相依赖.
//
// left: 左边间距.
//
// top: 顶部间距.
//
// right: 右边间距.
//
// bottom: 底部间距.
func (b *BkObj) X置外间距(左边间距 int, 顶部间距 int, 右边间距 int, 底部间距 int) int {
	return 炫彩基类.X背景对象_置外间距(b.Handle, 左边间距, 顶部间距, 右边间距, 底部间距)
}

// 背景对象_置对齐.
//
// nFlags: 对齐方式: xcc.BkObject_Align_Flag_.
func (b *BkObj) X置对齐(对齐方式 炫彩常量类.BkObject_Align_Flag_) int {
	return 炫彩基类.X背景对象_置对齐(b.Handle, 对齐方式)
}

// 背景对象_置图片.
//
// hImage: 图片句柄.
func (b *BkObj) X置图片(图片句柄 int) int {
	return 炫彩基类.X背景对象_置图片(b.Handle, 图片句柄)
}

// 背景对象_置旋转.
//
// angle: 旋转角度.
func (b *BkObj) X置旋转(旋转角度 float32) int {
	return 炫彩基类.X背景对象_置旋转(b.Handle, 旋转角度)
}

// 背景对象_置填充颜色.
//
// color: ABGR 颜色值.
func (b *BkObj) X置填充颜色(ABGR颜色值 int) int {
	return 炫彩基类.X背景对象_置填充颜色(b.Handle, ABGR颜色值)
}

// 背景对象_置边框宽度.
//
// width: 宽度.
func (b *BkObj) X置边框宽度(宽度 int) int {
	return 炫彩基类.X背景对象_置边框宽度(b.Handle, 宽度)
}

// 背景对象_置边框颜色.
//
// color: ABGR 颜色值.
func (b *BkObj) X置边框颜色(ABGR颜色值 int) int {
	return 炫彩基类.X背景对象_置边框颜色(b.Handle, ABGR颜色值)
}

// 背景对象_置矩形圆角.
//
// leftTop: 左上角.
//
// leftBottom: 左下角.
//
// rightTop: 右上角.
//
// rightBottom: 右下角.
func (b *BkObj) X置矩形圆角(左上角 int, 左下角 int, 右上角 int, 右下角 int) int {
	return 炫彩基类.X背景对象_置矩形圆角(b.Handle, 左上角, 左下角, 右上角, 右下角)
}

// 背景对象_启用填充, 启用绘制填充.
//
// bEnable: 是否启用.
func (b *BkObj) X启用填充(是否启用 bool) int {
	return 炫彩基类.X背景对象_启用填充(b.Handle, 是否启用)
}

// 背景对象_启用边框, 启用绘制边框.
//
// bEnable: 是否启用.
func (b *BkObj) X启用边框(是否启用 bool) int {
	return 炫彩基类.X背景对象_启用边框(b.Handle, 是否启用)
}

// 背景对象_置文本.
//
// pText: 文本字符串.
func (b *BkObj) X置文本(文本字符串 string) int {
	return 炫彩基类.X背景对象_置文本(b.Handle, 文本字符串)
}

// 背景对象_置字体.
//
// hFont: 字体句柄.
func (b *BkObj) X置字体(字体句柄 int) int {
	return 炫彩基类.X背景对象_置字体(b.Handle, 字体句柄)
}

// 背景对象_置文本对齐.
//
// nAlign: 文本对齐方式: xcc.TextFormatFlag_, xcc.TextAlignFlag_, xcc.TextTrimming_.
func (b *BkObj) X置文本对齐(对齐方式 炫彩常量类.TextFormatFlag_) int {
	return 炫彩基类.X背景对象_置文本对齐(b.Handle, 对齐方式)
}

// 背景对象_取外间距.
//
// pMargin: 接收返回外间距.
func (b *BkObj) X取外间距(接收返回 *炫彩基类.RECT) int {
	return 炫彩基类.X背景对象_取外间距(b.Handle, 接收返回)
}

// 背景对象_取对齐, 返回对齐标识: xcc.BkObject_Align_Flag_.
func (b *BkObj) X取对齐() 炫彩常量类.BkObject_Align_Flag_ {
	return 炫彩基类.X背景对象_取对齐(b.Handle)
}

// 背景对象_取图片, 返回图片句柄.
func (b *BkObj) X取图片() int {
	return 炫彩基类.X背景对象_取图片(b.Handle)
}

// 背景对象_取旋转角度, 返回旋转角度.
func (b *BkObj) X取旋转角度() int {
	return 炫彩基类.X背景对象_取旋转角度(b.Handle)
}

// 背景对象_取填充色, 返回ABGR填充色.
func (b *BkObj) X取填充色() int {
	return 炫彩基类.X背景对象_取填充色(b.Handle)
}

// 背景对象_取边框色, 返回ABGR边框色.
func (b *BkObj) X取边框色() int {
	return 炫彩基类.X背景对象_取边框色(b.Handle)
}

// 背景对象_取边框宽度, 返回边框宽度.
func (b *BkObj) X取边框宽度() int {
	return 炫彩基类.X背景对象_取边框宽度(b.Handle)
}

// 背景对象_取矩形圆角.
//
// pRect: 接收返回圆角大小.
func (b *BkObj) X取矩形圆角(接收返回圆角大小 *炫彩基类.RECT) int {
	return 炫彩基类.X背景对象_取矩形圆角(b.Handle, 接收返回圆角大小)
}

// 背景对象_是否填充.
func (b *BkObj) X是否填充() bool {
	return 炫彩基类.X背景对象_是否填充(b.Handle)
}

// 背景对象_是否边框.
func (b *BkObj) X是否边框() bool {
	return 炫彩基类.X背景对象_是否边框(b.Handle)
}

// 背景对象_取文本.
func (b *BkObj) X取文本() string {
	return 炫彩基类.X背景对象_取文本(b.Handle)
}

// 背景对象_取字体, 返回字体句柄.
func (b *BkObj) X取字体() int {
	return 炫彩基类.X背景对象_取字体(b.Handle)
}

// 背景对象_取文本对齐, 返回文本对齐方式: xcc.TextFormatFlag_.
func (b *BkObj) X取文本对齐() 炫彩常量类.TextFormatFlag_ {
	return 炫彩基类.X背景对象_取文本对齐(b.Handle)
}
