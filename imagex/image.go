package 炫彩图片类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/res"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Image 图片操作.
type Image struct {
	炫彩对象基类.ObjectBase
}

// 图片_加载从图片源.
//
// hImageSrc: 图片源句柄.
func X创建并按图片源句柄(图片源句柄 int) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从图片源(图片源句柄))
	return p
}

// 图片_加载从文件.
//
// pFileName: 图片文件.
func X创建并按文件(图片文件 string) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从文件(图片文件))
	return p
}

// 图片_加载从文件自适应, 加载图片从文件, 自适应图片.
//
// pFileName: 图片文件.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func X创建并按文件且自适应(图片文件 string, 左, 上, 右, 下 int32) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从文件自适应(图片文件, 左, 上, 右, 下))
	return p
}

// 图片_加载从文件指定区域, 加载图片, 指定区位置及大小.
//
// pFileName: 图片文件.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func X创建并按文件且指定区域(图片文件 string, 坐标x, 坐标y, 宽度, 高度 int32) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从文件指定区域(图片文件, 坐标x, 坐标y, 宽度, 高度))
	return p
}

// 图片_加载从资源自适应, 加载图片从资源, 自适应图片.
//
// id: 资源ID.
//
// pType: 资源类型.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func X创建并按资源且自适应(资源ID int32, 资源类型 string, 左, 上, 右, 下 int32, 从指定模块加载 uintptr) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从资源自适应(资源ID, 资源类型, 左, 上, 右, 下, 从指定模块加载))
	return p
}

// 图片_加载从资源.
//
// id: 资源ID.
//
// pType: 资源类型.
//
// bStretch: 是否拉伸图片.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func X创建并按资源(资源ID int32, 资源类型 string, 是否拉伸图片 bool, 从指定模块加载 uintptr) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从资源(资源ID, 资源类型, 是否拉伸图片, 从指定模块加载))
	return p
}

// 图片_加载从ZIP, 加载图片从ZIP压缩包.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
func X创建并按ZIP(ZIP压缩包文件名 string, 图片文件名 string, ZIP压缩包密码 string) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从ZIP(ZIP压缩包文件名, 图片文件名, ZIP压缩包密码))
	return p
}

// 图片_加载从资源ZIP.
//
// id: RC资源ID.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
//
// hModule: 模块句柄, 可填0.
func X创建并按资源ZIP(RC资源ID int32, 图片文件名 string, ZIP压缩包密码 string, 模块句柄 uintptr) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从资源ZIP(RC资源ID, 图片文件名, ZIP压缩包密码, 模块句柄))
	return p
}

// 图片_加载从ZIP自适应, 加载图片从ZIP压缩包, 自适应图片.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
//
// x1: 坐标.
//
// x2: 坐标.
//
// y1: 坐标.
//
// y2: 坐标.
func X创建并按ZIP且自适应(ZIP压缩包文件名 string, 图片文件名 string, ZIP压缩包密码 string, 坐标x1, 坐标x2, 坐标y1, 坐标y2 int32) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从ZIP自适应(ZIP压缩包文件名, 图片文件名, ZIP压缩包密码, 坐标x1, 坐标x2, 坐标y1, 坐标y2))
	return p
}

// 图片_加载从ZIP指定区域, 加载ZIP图片, 指定区位置及大小.
//
// pZipFileName: ZIP文件.
//
// pFileName: 图片名称.
//
// pPassword: 密码.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func X创建并按ZIP且指定区域(ZIP文件 string, 图片名称 string, 密码 string, 坐标x, 坐标y, 宽度, 高度 int32) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从ZIP指定区域(ZIP文件, 图片名称, 密码, 坐标x, 坐标y, 宽度, 高度))
	return p
}

// 图片_加载从内存ZIP.
//
// data: 图片数据.
//
// pFileName: 图片名称.
//
// pPassword: zip压缩包密码.
func X创建并按内存ZIP(图片数据 []byte, 图片名称 string, zip压缩包密码 string) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从内存ZIP(图片数据, 图片名称, zip压缩包密码))
	return p
}

// 图片_加载从内存, 加载流图片.
//
// pBuffer: 图片数据.
func X创建并按内存(图片数据 []byte) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从内存(图片数据))
	return p
}

// 图片_加载从内存指定区域, 加载流图片, 指定区位置及大小.
//
// pBuffer: 图片数据.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func X创建并按内存且指定区域(图片数据 []byte, 坐标x, 坐标y, 宽度, 高度 int32) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从内存指定区域(图片数据, 坐标x, 坐标y, 宽度, 高度))
	return p
}

// 图片_加载从内存自适应, 加载流图片压缩包, 自适应图片(九宫格).
//
// pBuffer: 图片数据.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func X创建并按内存且自适应(图片数据 []byte, 坐标左, 坐标上, 坐标右, 坐标下 int32) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从内存自适应(图片数据, 坐标左, 坐标上, 坐标右, 坐标下))
	return p
}

// 图片_加载从Image, 加载图片从GDI+的Image对象.
//
// pImage: GDI图片对象指针Image*.
func X创建并按GDI图片对象指针(GDI图片对象指针 uintptr) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从Image(GDI图片对象指针))
	return p
}

// 图片_加载文件图标, 加载文件图标, 从一个EXE文件或DLL文件或图标文件; 例如:*.exe文件的图标.
//
// pFileName: 文件名.
func X创建并按图标文件(文件名 string) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载文件图标(文件名))
	return p
}

// 图片_加载从HICON, 创建一个炫彩图片句柄, 从一个现有的图标句柄HICON.
//
// hIcon: 图标句柄.
func X创建并按图标句柄(图标句柄 uintptr) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从HICON(图标句柄))
	return p
}

// 图片_加载从HBITMAP, 创建一个炫彩图片句柄, 从一个现有的位图句柄HBITMAP.
//
// hBitmap: 位图句柄.
func X创建并按位图句柄(位图句柄 uintptr) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从HBITMAP(位图句柄))
	return p
}

// 图片_加载从SVG.
//
// hSvg: SVG句柄.
func X创建并按SVG(SVG句柄 int) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从SVG(SVG句柄))
	return p
}

// 图片_加载从SVG文件.
//
// pFileName: 文件名.
func X创建并按SVG文件(文件名 string) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从SVG文件(文件名))
	return p
}

// 图片_加载从SVG字符串.
//
// pString: 字符串.
func X创建并按SVG字符串(字符串 string) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从SVG字符串(字符串))
	return p
}

// 图片_加载从SVG字符串W.
//
// pString: 字符串.
func X创建并按SVG字符串W(字符串 string) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从SVG字符串W(字符串))
	return p
}

// 图片_加载从SVG字符串UTF8, 更推荐使用 imagex.NewBySvgStringW.
//
// pString: 字符串.
func X创建并按SVG字符串UTF8(字符串 string) *Image {
	p := &Image{}
	p.X设置句柄(炫彩基类.X图片_加载从SVG字符串UTF8(字符串))
	return p
}

// 从句柄创建对象.
func X创建并按句柄(句柄 int) *Image {
	p := &Image{}
	p.X设置句柄(句柄)
	return p
}

// 根据资源文件中的name创建对象, 失败返回nil.
//
// pName: 资源名称.
func X创建并按资源名称(资源名称 string) *Image {
	handle := 炫彩资源类.X取图片(资源名称)
	if handle > 0 {
		p := &Image{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从指定的资源文件中, 根据name创建对象, 失败返回nil.
//
// pFileName: 资源文件名.
//
// pName: 资源名称.
func X创建并按资源文件名(资源文件名, 资源名称 string) *Image {
	handle := 炫彩资源类.X取图片EX(资源文件名, 资源名称)
	if handle > 0 {
		p := &Image{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 图片_取SVG, 返回SVG句柄.
//
// hImage: 图片句柄.
func (i *Image) X取SVG() int {
	return 炫彩基类.X图片_取SVG(i.Handle)
}

// 图片_判断缩放, 是否为拉伸图片句柄.
func (i *Image) X判断缩放() bool {
	return 炫彩基类.X图片_判断缩放(i.Handle)
}

// 图片_判断自适应, 是否为自适应图片句柄.
func (i *Image) X判断自适应() bool {
	return 炫彩基类.X图片_判断自适应(i.Handle)
}

// 图片_判断平铺, 是否为平铺图片.
func (i *Image) X判断平铺() bool {
	return 炫彩基类.X图片_判断平铺(i.Handle)
}

// 图片_置绘制类型, 设置图片绘制类型.
//
// nType: 图片绘制类型, Image_Draw_Type_.
func (i *Image) X置绘制类型(图片绘制类型 炫彩常量类.Image_Draw_Type_) bool {
	return 炫彩基类.X图片_置绘制类型(i.Handle, 图片绘制类型)
}

// 图片_置绘制类型自适应, 设置图片自适应(九宫格).
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func (i *Image) X置绘制类型自适应(坐标左, 坐标上, 坐标右, 坐标下 int32) bool {
	return 炫彩基类.X图片_置绘制类型自适应(i.Handle, 坐标左, 坐标上, 坐标右, 坐标下)
}

// 图片_置透明色, 指定图片透明颜色.
//
// color: ABGR 颜色.
func (i *Image) X置透明色(ABGR颜色 int) {
	炫彩基类.X图片_置透明色(i.Handle, ABGR颜色)
}

// 图片_置透明色扩展, 指定图片透明颜色及透明度.
//
// color: ABGR 颜色.
//
// tranColor: 透明色的透明度.
func (i *Image) X置透明色EX(ABGR颜色 int, 透明度 byte) {
	炫彩基类.X图片_置透明色EX(i.Handle, ABGR颜色, 透明度)
}

// 图片_置旋转角度, 设置旋转角度, 返回先前角度.
//
// fAngle: 选择角度.
func (i *Image) X置旋转角度(选择角度 float32) float32 {
	return 炫彩基类.X图片_置旋转角度(i.Handle, 选择角度)
}

// 图片_置等分.
//
// nCount: 等分数量.
//
// iIndex: 索引.
func (i *Image) X置等分(等分数量, 索引 int32) {
	炫彩基类.X图片_置等分(i.Handle, 等分数量, 索引)
}

// 图片_启用透明色, 启用或关闭图片透明色.
//
// bEnable: 启用TRUE.
func (i *Image) X启用透明色(启用TRUE bool) {
	炫彩基类.X图片_启用透明色(i.Handle, 启用TRUE)
}

// 图片_启用自动销毁, 启用或关闭自动销毁, 当与UI元素关联时有效.
//
// bEnable: 启用自动销毁TRUE.
func (i *Image) X启用自动销毁(启用自动销毁TRUE bool) {
	炫彩基类.X图片_启用自动销毁(i.Handle, 启用自动销毁TRUE)
}

// 图片_启用居中, 启用或关闭图片居中显示，默认属性图片有效.
//
// bCenter: 是否居中显示.
func (i *Image) X启用居中(是否居中显示 bool) {
	炫彩基类.X图片_启用居中(i.Handle, 是否居中显示)
}

// 图片_判断居中, 判断图片是否居中显示.
func (i *Image) X判断居中() bool {
	return 炫彩基类.X图片_判断居中(i.Handle)
}

// 图片_取绘制类型, 获取图片绘制类型, 返回: xcc.Image_Draw_Type_.
func (i *Image) X取绘制类型() 炫彩常量类.Image_Draw_Type_ {
	return 炫彩基类.X图片_取绘制类型(i.Handle)
}

// 图片_取宽度.
func (i *Image) X取宽度() int32 {
	return 炫彩基类.X图片_取宽度(i.Handle)
}

// 图片_取高度.
func (i *Image) X取高度() int32 {
	return 炫彩基类.X图片_取高度(i.Handle)
}

// 图片_取图片源.
func (i *Image) X取图片源() int {
	return 炫彩基类.X图片_取图片源(i.Handle)
}

// 图片_增加引用计数.
func (i *Image) X增加引用计数() {
	炫彩基类.X图片_增加引用计数(i.Handle)
}

// 图片_释放引用计数, 释放引用计数, 当引用计数为0时, 自动销毁.
func (i *Image) X释放引用计数() {
	炫彩基类.X图片_释放引用计数(i.Handle)
}

// 图片_取引用计数.
func (i *Image) X取引用计数() int32 {
	return 炫彩基类.X图片_取引用计数(i.Handle)
}

// 图片_销毁, 强制销毁图片, 谨慎使用, 建议使用 XImage_Release() 释放.
func (i *Image) X销毁() {
	炫彩基类.X图片_销毁(i.Handle)
}

// 图片_置缩放大小, 启用缩放属性后有效, 值大于0有效.
//
// width: 宽度.
//
// height: 高度.
func (i *Image) X置缩放大小(宽度, 高度 int32) {
	炫彩基类.X图片_置缩放大小(i.Handle, 宽度, 高度)
}
