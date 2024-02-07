package 炫彩SVG类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
)

// SVG矢量图形.
type Svg struct {
	炫彩对象基类.ObjectBase
}

// SVG_加载从文件, 返回Svg对象.
//
// pFileName: 文件名.
func X创建并按文件(文件名 string) *Svg {
	p := &Svg{}
	p.X设置句柄(炫彩基类.XSVG_加载从文件(文件名))
	return p
}

// SVG_加载从字符串, 返回Svg对象.
//
// pString: 字符串.
func X创建并按字符串(字符串 string) *Svg {
	p := &Svg{}
	p.X设置句柄(炫彩基类.XSVG_加载从字符串(字符串))
	return p
}

// SVG_加载从字符串W.
//
// pString: 字符串.
func X创建并按字符串W(字符串 string) *Svg {
	p := &Svg{}
	p.X设置句柄(炫彩基类.XSVG_加载从字符串W(字符串))
	return p
}

// SVG_加载从字符串UTF8.
//
// pString: 字符串.
func X创建并按字符串UTF8(字符串 string) *Svg {
	p := &Svg{}
	p.X设置句柄(炫彩基类.XSVG_加载从字符串UTF8(字符串))
	return p
}

// SVG_加载从ZIP, 返回Svg对象.
//
// pZipFileName: zip文件名.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func X创建并按ZIP(zip文件名, svg文件名, zip密码 string) *Svg {
	p := &Svg{}
	p.X设置句柄(炫彩基类.XSVG_加载从ZIP(zip文件名, svg文件名, zip密码))
	return p
}

// SVG_加载从资源ZIP, 返回SVG对象.
//
// id: 资源ID.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func X创建并按资源ZIP(资源ID int32, svg文件名, zip密码 string, 模块句柄 uintptr) *Svg {
	p := &Svg{}
	p.X设置句柄(炫彩基类.XSVG_加载从资源ZIP(资源ID, svg文件名, zip密码, 模块句柄))
	return p
}

// SVG_加载从内存ZIP, 返回Svg对象.
//
// data: zip数据.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func X创建并按内存ZIP(zip数据 []byte, svg文件名, zip密码 string) *Svg {
	p := &Svg{}
	p.X设置句柄(炫彩基类.XSVG_加载从内存ZIP(zip数据, svg文件名, zip密码))
	return p
}

// SVG_加载从资源, 返回Svg对象.
//
// id: 资源ID.
//
// pType: 资源类型.在rc资源文件中.
//
// hModule: 从指定模块加载.
func X创建并按资源(资源ID int32, 资源类型 string, 从指定模块加载 uintptr) *Svg {
	p := &Svg{}
	p.X设置句柄(炫彩基类.XSVG_加载从资源(资源ID, 资源类型, 从指定模块加载))
	return p
}

// SVG_置大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *Svg) X置大小(宽度, 高度 int32) {
	炫彩基类.XSVG_置大小(s.Handle, 宽度, 高度)
}

// SVG_取大小.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func (s *Svg) X取大小(接收返回宽度, 接收返回高度 *int32) {
	炫彩基类.XSVG_取大小(s.Handle, 接收返回宽度, 接收返回高度)
}

// SVG_取宽度.
func (s *Svg) X取宽度() int32 {
	return 炫彩基类.XSVG_取宽度(s.Handle)
}

// SVG_取高度.
func (s *Svg) X取高度() int32 {
	return 炫彩基类.XSVG_取高度(s.Handle)
}

// SVG_置偏移.
//
// x: x轴偏移.
//
// y: y轴偏移.
func (s *Svg) X置偏移(x轴偏移, y轴偏移 int32) {
	炫彩基类.XSVG_置偏移(s.Handle, x轴偏移, y轴偏移)
}

// SVG_取偏移.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func (s *Svg) X取偏移(x轴偏移, y轴偏移 *int32) {
	炫彩基类.XSVG_取偏移(s.Handle, x轴偏移, y轴偏移)
}

// SVG_置偏移F.
//
// x: x轴偏移.
//
// y: y轴偏移.
func (s *Svg) X置偏移F(x轴偏移, y轴偏移 float32) {
	炫彩基类.XSVG_置偏移F(s.Handle, x轴偏移, y轴偏移)
}

// SVG_取偏移F.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func (s *Svg) X取偏移F(x轴偏移, y轴偏移 *float32) {
	炫彩基类.XSVG_取偏移F(s.Handle, x轴偏移, y轴偏移)
}

// SVG_取视图框.
//
// pViewBox: 接收返回视图框.
func (s *Svg) X取视图框(接收返回视图框 *炫彩基类.RECT) {
	炫彩基类.XSVG_取视图框(s.Handle, 接收返回视图框)
}

// SVG_启用自动销毁.
//
// bEnable: 是否自动销毁.
func (s *Svg) X启用自动销毁(是否自动销毁 bool) {
	炫彩基类.XSVG_启用自动销毁(s.Handle, 是否自动销毁)
}

// SVG_增加引用计数.
func (s *Svg) X增加引用计数() {
	炫彩基类.XSVG_增加引用计数(s.Handle)
}

// SVG_释放引用计数.
func (s *Svg) X释放引用计数() {
	炫彩基类.XSVG_释放引用计数(s.Handle)
}

// SVG_取引用计数.
func (s *Svg) X取引用计数() int32 {
	return 炫彩基类.XSVG_取引用计数(s.Handle)
}

// SVG_销毁.
func (s *Svg) X销毁() {
	炫彩基类.XSVG_销毁(s.Handle)
}

// SVG_置透明度.
//
// alpha: 透明度.
func (s *Svg) X置透明度(透明度 byte) {
	炫彩基类.XSVG_置透明度(s.Handle, 透明度)
}

// SVG_取透明度, 返回透明度.
func (s *Svg) X取透明度() byte {
	return 炫彩基类.XSVG_取透明度(s.Handle)
}

// SVG_置用户填充颜色, 用户颜色将覆盖默认样式.
//
// color: 颜色, AGBR颜色.
//
// bEnable: 是否有效.
func (s *Svg) X置用户填充颜色(颜色 int, 是否有效 bool) {
	炫彩基类.XSVG_置用户填充颜色(s.Handle, 颜色, 是否有效)
}

// SVG_置用户笔触颜色, 用户颜色将覆盖默认样式.
//
// color: 颜色, AGBR颜色.
//
// strokeWidth: 笔触宽度.
//
// bEnable: 是否有效.
func (s *Svg) X置用户笔触颜色(颜色 int, 笔触宽度 float32, 是否有效 bool) {
	炫彩基类.XSVG_置用户笔触颜色(s.Handle, 颜色, 笔触宽度, 是否有效)
}

// SVG_取用户填充颜色.
//
// pColor: 返回颜色值, AGBR颜色.
func (s *Svg) X取用户填充颜色(返回颜色值 *int) bool {
	return 炫彩基类.XSVG_取用户填充颜色(s.Handle, 返回颜色值)
}

// SVG_取用户笔触颜色.
//
// pColor: 返回颜色值, AGBR颜色.
//
// pStrokeWidth: .
func (s *Svg) X取用户笔触颜色(返回颜色值 *int, pStrokeWidth *float32) bool {
	return 炫彩基类.XSVG_取用户笔触颜色(s.Handle, 返回颜色值, pStrokeWidth)
}

// SVG_置旋转角度, 默认以自身中心点旋转.
//
// angle: 转角度.
func (s *Svg) X置旋转角度(转角度 float32) {
	炫彩基类.XSVG_置旋转角度(s.Handle, 转角度)
}

// SVG_取旋转角度, 返回旋转角度.
func (s *Svg) X取旋转角度() float32 {
	return 炫彩基类.XSVG_取旋转角度(s.Handle)
}

// SVG_置旋转.
//
// angle: 角度.
//
// x: 旋转中心点X.
//
// y: 旋转中心点Y.
//
// bOffset: TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func (s *Svg) X置旋转(角度 float32, 旋转中心点X float32, 旋转中心点Y float32, 偏移方式 bool) {
	炫彩基类.XSVG_置旋转(s.Handle, 角度, 旋转中心点X, 旋转中心点Y, 偏移方式)
}

// SVG_取旋转.
//
// pAngle: 返回角度.
//
// pX: 返回 旋转中心点X.
//
// pY: 返回 旋转中心点Y.
//
// pbOffset: 返回TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func (s *Svg) X取旋转(返回角度 *float32, 返回x *float32, 返回y *float32, 返回偏移方式 *bool) {
	炫彩基类.XSVG_取旋转(s.Handle, 返回角度, 返回x, 返回y, 返回偏移方式)
}

// SVG_显示, 显示或隐藏.
//
// bShow: 是否显示.
func (s *Svg) X显示(是否显示 bool) {
	炫彩基类.XSVG_显示(s.Handle, 是否显示)
}
