package 炫彩资源类//bm:炫彩资源类

import "github.com/888go/xcgui/xc"

// 资源_启用延迟加载, 启用延迟加载; 图片文件, 列表项模板文件.
//
// bEnable: 是否启用.

// ff:启用延迟加载
// bEnable:是否启用
func X启用延迟加载(是否启用 bool) int {
	return 炫彩基类.X资源_启用延迟加载(是否启用)
}

// 资源_置文件加载回调, 设置文件加载回调函数.
//
// pFun: 回调函数.

// ff:置文件加载回调
// pFun:回调函数
func X置文件加载回调(回调函数 interface{}) int {
	return 炫彩基类.X资源_置文件加载回调(回调函数)
}

// 资源_取ID值, 获取资源ID整型值.
//
// pName: 资源ID名称.

// ff:取ID值
// pName:资源ID名称
func X取ID值(资源ID名称 string) int {
	return 炫彩基类.X资源_取ID值(资源ID名称)
}

// 资源_取图片, 查找资源图片.
//
// pName: 资源名称.

// ff:取图片
// pName:资源名称
func X取图片(资源名称 string) int {
	return 炫彩基类.X资源_取图片(资源名称)
}

// 资源_取图片扩展, 从指定的资源文件中查找图片.
//
// pFileName: 资源文件名.
//
// pName: 资源名称.

// ff:取图片EX
// pName:资源名称
// pFileName:资源文件名
func X取图片EX(资源文件名 string, 资源名称 string) int {
	return 炫彩基类.X资源_取图片EX(资源文件名, 资源名称)
}

// 资源_取颜色, 从资源中查找颜色.
//
// pName: 资源名称.

// ff:取颜色
// pName:资源名称
func X取颜色(资源名称 string) int {
	return 炫彩基类.X资源_取颜色(资源名称)
}

// 资源_取字体, 从资源中查找字体.
//
// pName: 资源名称.

// ff:取字体
// pName:资源名称
func X取字体(资源名称 string) int {
	return 炫彩基类.X资源_取字体(资源名称)
}

// 资源_取背景管理器, 从资源中查找背景.
//
// pName: 资源名称.

// ff:取背景管理器
// pName:资源名称
func X取背景管理器(资源名称 string) int {
	return 炫彩基类.X资源_取背景管理器(资源名称)
}
