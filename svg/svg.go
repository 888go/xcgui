package svg

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// SVG矢量图形.
type Svg struct {
	objectbase.ObjectBase
}

// 创建SVG并按文件, 返回Svg对象.
//
// pFileName: 文件名.
func NewByFile(pFileName string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadFile(pFileName))
	return p
}

// 创建SVG并按字符串, 返回Svg对象.
//
// pString: 字符串.
func NewByString(pString string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadString(pString))
	return p
}

// 创建SVG并按字符串W.
//
// pString: 字符串.
func NewByStringW(pString string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadStringW(pString))
	return p
}

// 创建SVG并按字符串UTF8.
//
// pString: 字符串.
func NewByStringUtf8(pString string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadStringUtf8(pString))
	return p
}

// 创建SVG并按ZIP, 返回Svg对象.
//
// pZipFileName: zip文件名.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func NewByZip(pZipFileName, pFileName, pPassword string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadZip(pZipFileName, pFileName, pPassword))
	return p
}

// 创建SVG并按资源ZIP, 返回SVG对象.
//
// id: 资源ID.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func NewByZipRes(id int32, pFileName, pPassword string, hModule uintptr) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadZipRes(id, pFileName, pPassword, hModule))
	return p
}

// 创建SVG并按内存ZIP, 返回Svg对象.
//
// data: zip数据.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func NewByZipMem(data []byte, pFileName, pPassword string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadZipMem(data, pFileName, pPassword))
	return p
}

// 创建SVG并按资源, 返回Svg对象.
//
// id: 资源ID.
//
// pType: 资源类型  .在rc资源文件中.
//
// hModule: 从指定模块加载.
func NewByRes(id int32, pType string, hModule uintptr) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadRes(id, pType, hModule))
	return p
}

// 置大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *Svg) SetSize(nWidth, nHeight int32) {
	xc.XSvg_SetSize(s.Handle, nWidth, nHeight)
}

// 取大小.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func (s *Svg) GetSize(pWidth, pHeight *int32) {
	xc.XSvg_GetSize(s.Handle, pWidth, pHeight)
}

// 取宽度.
func (s *Svg) GetWidth() int32 {
	return xc.XSvg_GetWidth(s.Handle)
}

// 取高度.
func (s *Svg) GetHeight() int32 {
	return xc.XSvg_GetHeight(s.Handle)
}

// 置偏移.
//
// x: x轴偏移.
//
// y: y轴偏移.
func (s *Svg) SetPosition(x, y int32) {
	xc.XSvg_SetPosition(s.Handle, x, y)
}

// 取偏移.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func (s *Svg) GetPosition(pX, pY *int32) {
	xc.XSvg_GetPosition(s.Handle, pX, pY)
}

// 置偏移F.
//
// x: x轴偏移.
//
// y: y轴偏移.
func (s *Svg) SetPositionF(x, y float32) {
	xc.XSvg_SetPositionF(s.Handle, x, y)
}

// 取偏移F.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func (s *Svg) GetPositionF(pX, pY *float32) {
	xc.XSvg_GetPositionF(s.Handle, pX, pY)
}

// 取视图框.
//
// pViewBox: 接收返回视图框.
func (s *Svg) GetViewBox(pViewBox *xc.RECT) {
	xc.XSvg_GetViewBox(s.Handle, pViewBox)
}

// 启用自动销毁.
//
// bEnable: 是否自动销毁.
func (s *Svg) EnableAutoDestroy(bEnable bool) {
	xc.XSvg_EnableAutoDestroy(s.Handle, bEnable)
}

// 增加引用计数.
func (s *Svg) AddRef() {
	xc.XSvg_AddRef(s.Handle)
}

// 释放引用计数.
func (s *Svg) Release() {
	xc.XSvg_Release(s.Handle)
}

// 取引用计数.
func (s *Svg) GetRefCount() int32 {
	return xc.XSvg_GetRefCount(s.Handle)
}

// 销毁.
func (s *Svg) Destroy() {
	xc.XSvg_Destroy(s.Handle)
}

// 置透明度.
//
// alpha: 透明度.
func (s *Svg) SetAlpha(alpha byte) {
	xc.XSvg_SetAlpha(s.Handle, alpha)
}

// 取透明度, 返回透明度.
func (s *Svg) GetAlpha() byte {
	return xc.XSvg_GetAlpha(s.Handle)
}

// 置用户填充颜色, 用户颜色将覆盖默认样式.
//
// color: 颜色, AGBR颜色.
//
// bEnable: 是否有效.
func (s *Svg) SetUserFillColor(color int, bEnable bool) {
	xc.XSvg_SetUserFillColor(s.Handle, color, bEnable)
}

// 置用户笔触颜色, 用户颜色将覆盖默认样式.
//
// color: 颜色, AGBR颜色.
//
// strokeWidth: 笔触宽度.
//
// bEnable: 是否有效.
func (s *Svg) SetUserStrokeColor(color int, strokeWidth float32, bEnable bool) {
	xc.XSvg_SetUserStrokeColor(s.Handle, color, strokeWidth, bEnable)
}

// 取用户填充颜色.
//
// pColor: 返回颜色值, AGBR颜色.
func (s *Svg) GetUserFillColor(pColor *int) bool {
	return xc.XSvg_GetUserFillColor(s.Handle, pColor)
}

// 取用户笔触颜色.
//
// pColor: 返回颜色值, AGBR颜色.
//
// pStrokeWidth: .
func (s *Svg) GetUserStrokeColor(pColor *int, pStrokeWidth *float32) bool {
	return xc.XSvg_GetUserStrokeColor(s.Handle, pColor, pStrokeWidth)
}

// 置旋转角度, 默认以自身中心点旋转.
//
// angle: 转角度.
func (s *Svg) SetRotateAngle(angle float32) {
	xc.XSvg_SetRotateAngle(s.Handle, angle)
}

// 取旋转角度, 返回旋转角度.
func (s *Svg) GetRotateAngle() float32 {
	return xc.XSvg_GetRotateAngle(s.Handle)
}

// 置旋转.
//
// angle: 角度.
//
// x: 旋转中心点X.
//
// y: 旋转中心点Y.
//
// bOffset: TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func (s *Svg) SetRotate(angle float32, x float32, y float32, bOffset bool) {
	xc.XSvg_SetRotate(s.Handle, angle, x, y, bOffset)
}

// 取旋转.
//
// pAngle: 返回角度.
//
// pX: 返回 旋转中心点X.
//
// pY: 返回 旋转中心点Y.
//
// pbOffset: 返回TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func (s *Svg) GetRotate(pAngle *float32, pX *float32, pY *float32, pbOffset *bool) {
	xc.XSvg_GetRotate(s.Handle, pAngle, pX, pY, pbOffset)
}

// 显示, 显示或隐藏.
//
// bShow: 是否显示.
func (s *Svg) Show(bShow bool) {
	xc.XSvg_Show(s.Handle, bShow)
}
