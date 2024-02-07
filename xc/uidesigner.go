package 炫彩基类

import (
	"github.com/888go/xcgui/common"
)

// 炫彩_加载布局文件, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func X炫彩_加载布局文件(布局文件名 string, 父对象句柄 int, 附加窗口句柄 uintptr) int {
	r, _, _ := xC_LoadLayout.Call(炫彩工具类.StrPtr(布局文件名), uintptr(父对象句柄), 附加窗口句柄)
	return int(r)
}

// 炫彩_加载布局文件Ex, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func X炫彩_加载布局文件Ex(布局文件名, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) int {
	r, _, _ := xC_LoadLayoutEx.Call(炫彩工具类.StrPtr(布局文件名), 炫彩工具类.StrPtr(名称前缀), uintptr(父对象句柄), 父窗口句柄HWND, 附加窗口句柄)
	return int(r)
}

// 炫彩_加载布局文件ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func X炫彩_加载布局文件ZIP(zip文件名 string, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) int {
	r, _, _ := xC_LoadLayoutZip.Call(炫彩工具类.StrPtr(zip文件名), 炫彩工具类.StrPtr(布局文件名), 炫彩工具类.StrPtr(zip密码), uintptr(父对象句柄), 附加窗口句柄)
	return int(r)
}

// 炫彩_加载布局文件ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func X炫彩_加载布局文件ZIPEx(zip文件名 string, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) int {
	r, _, _ := xC_LoadLayoutZipEx.Call(炫彩工具类.StrPtr(zip文件名), 炫彩工具类.StrPtr(布局文件名), 炫彩工具类.StrPtr(zip密码), 炫彩工具类.StrPtr(名称前缀), uintptr(父对象句柄), 父窗口句柄HWND, 附加窗口句柄)
	return int(r)
}

// 炫彩_加载布局文件内存ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 布局文件数据.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func X炫彩_加载布局文件内存ZIP(布局文件数据 []byte, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) int {
	r, _, _ := xC_LoadLayoutZipMem.Call(炫彩工具类.ByteSliceDataPtr(&布局文件数据), uintptr(len(布局文件数据)), 炫彩工具类.StrPtr(布局文件名), 炫彩工具类.StrPtr(zip密码), uintptr(父对象句柄), 附加窗口句柄)
	return int(r)
}

// 炫彩_加载布局文件内存ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 布局文件数据.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func X炫彩_加载布局文件内存ZIPEx(布局文件数据 []byte, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) int {
	r, _, _ := xC_LoadLayoutZipMemEx.Call(炫彩工具类.ByteSliceDataPtr(&布局文件数据), uintptr(len(布局文件数据)), 炫彩工具类.StrPtr(布局文件名), 炫彩工具类.StrPtr(zip密码), 炫彩工具类.StrPtr(名称前缀), uintptr(父对象句柄), 父窗口句柄HWND, 附加窗口句柄)
	return int(r)
}

/* // 炫彩_加载布局文件从字符串, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutFromString(pStringXML string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutFromString.Call(XC_wtoa(pStringXML), uintptr(hParent), hAttachWnd)
	return int(r)
} */

// 炫彩_加载布局文件从字符串W, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func X炫彩_加载布局文件从字符串W(字符串 string, 父对象句柄 int, 附加窗口句柄 uintptr) int {
	r, _, _ := xC_LoadLayoutFromStringUtf8.Call(X文本W到UTF8(字符串), uintptr(父对象句柄), 附加窗口句柄)
	return int(r)
}

// 炫彩_加载布局文件从字符串WEx, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func X炫彩_加载布局文件从字符串WEx(字符串, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) int {
	r, _, _ := xC_LoadLayoutFromStringUtf8Ex.Call(X文本W到UTF8(字符串), 炫彩工具类.StrPtr(名称前缀), uintptr(父对象句柄), 父窗口句柄HWND, 附加窗口句柄)
	return int(r)
}

// 炫彩_加载样式文件.
//
// pFileName: 样式文件名称.
func X炫彩_加载样式文件(样式文件名称 string) bool {
	r, _, _ := xC_LoadStyle.Call(炫彩工具类.StrPtr(样式文件名称))
	return r != 0
}

// 炫彩_加载样式文件ZIP.
//
// pZipFile: ZIP文件名.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func X炫彩_加载样式文件ZIP(ZIP文件名 string, 文件名 string, 密码 string) bool {
	r, _, _ := xC_LoadStyleZip.Call(炫彩工具类.StrPtr(ZIP文件名), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(密码))
	return r != 0
}

// 炫彩_加载样式文件从内存ZIP.
//
// data: 样式文件数据.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func X炫彩_加载样式文件从内存ZIP(样式文件数据 []byte, 文件名 string, 密码 string) bool {
	r, _, _ := xC_LoadStyleZipMem.Call(炫彩工具类.ByteSliceDataPtr(&样式文件数据), uintptr(len(样式文件数据)), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(密码))
	return r != 0
}

// 炫彩_加载资源文件.
//
// pFileName: 资源文件名.
func X炫彩_加载资源文件(资源文件名 string) bool {
	r, _, _ := xC_LoadResource.Call(炫彩工具类.StrPtr(资源文件名))
	return r != 0
}

// 炫彩_加载资源文件ZIP.
//
// pZipFileName: zip文件名.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func X炫彩_加载资源文件ZIP(zip文件名 string, 资源文件名 string, zip压缩包密码 string) bool {
	r, _, _ := xC_LoadResourceZip.Call(炫彩工具类.StrPtr(zip文件名), 炫彩工具类.StrPtr(资源文件名), 炫彩工具类.StrPtr(zip压缩包密码))
	return r != 0
}

// 炫彩_加载资源文件内存ZIP.
//
// data: 资源文件数据.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func X炫彩_加载资源文件内存ZIP(资源文件数据 []byte, 资源文件名 string, zip压缩包密码 string) bool {
	r, _, _ := xC_LoadResourceZipMem.Call(炫彩工具类.ByteSliceDataPtr(&资源文件数据), uintptr(len(资源文件数据)), 炫彩工具类.StrPtr(资源文件名), 炫彩工具类.StrPtr(zip压缩包密码))
	return r != 0
}

// 炫彩_加载资源文件从字符串W.
//
// pStringXML: 字符串.
//
// pFileName: 资源文件名.
func X炫彩_加载资源文件从字符串W(字符串 string, 资源文件名 string) bool {
	r, _, _ := xC_LoadResourceFromStringUtf8.Call(X文本W到UTF8(字符串), 炫彩工具类.StrPtr(资源文件名))
	return r != 0
}

// 炫彩_加载样式文件从字符串W.
//
// pFileName: 样式文件名.
//
// pString: 字符串.
func X炫彩_加载样式文件从字符串W(样式文件名 string, 字符串 string) bool {
	r, _, _ := xC_LoadStyleFromStringW.Call(炫彩工具类.StrPtr(样式文件名), 炫彩工具类.StrPtr(字符串))
	return r != 0
}

// 炫彩_加载布局文件资源ZIP扩展. 加载布局文件从RC资源zip压缩包中, 返回窗口句柄或元素句柄.
//
// id: RC资源ID.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄, 可填0.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用, 可填0.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 可填0.
func X炫彩_加载布局文件资源ZIPEX(RC资源ID int32, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄, 模块句柄 uintptr) int {
	r, _, _ := xC_LoadLayoutZipResEx.Call(uintptr(RC资源ID), 炫彩工具类.StrPtr(布局文件名), 炫彩工具类.StrPtr(zip密码), 炫彩工具类.StrPtr(名称前缀), uintptr(父对象句柄), 父窗口句柄HWND, 附加窗口句柄, 模块句柄)
	return int(r)
}

// 炫彩_加载资源文件资源ZIP. 加载资源文件从RC资源zip压缩包中.
//
// id: RC资源ID.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
//
// hModule: 模块句柄, 可填0.
func X炫彩_加载资源文件资源ZIP(RC资源ID int, 资源文件名 string, zip压缩包密码 string, 模块句柄 uintptr) bool {
	r, _, _ := xC_LoadResourceZipRes.Call(uintptr(RC资源ID), 炫彩工具类.StrPtr(资源文件名), 炫彩工具类.StrPtr(zip压缩包密码), 模块句柄)
	return r != 0
}

// 炫彩_加载样式文件从资源ZIP. 从RC资源中的ZIP包中, 加载样式文件.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: 密码.
//
// hModule: 模块句柄, 可填0.
func X炫彩_加载样式文件从资源ZIP(RC资源ID int, 文件名 string, 密码 string, 模块句柄 uintptr) bool {
	r, _, _ := xC_LoadStyleZipRes.Call(uintptr(RC资源ID), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(密码), 模块句柄)
	return r != 0
}

/*
// 炫彩_加载资源文件从字符串.
//
// pStringXML: 字符串.
//
// pFileName: 资源文件名.
func XC_LoadResourceFromString(pStringXML string, pFileName string) bool {
	r, _, _ := xC_LoadResourceFromString.Call(strPtr(pStringXML), strPtr(pFileName))
	return r!=0
}

// 炫彩_加载样式文件从字符串.
//
// pFileName: 样式文件名, 用于打印错误文件和定位关联资源文件位置.
//
// pString: 字符串.
func XC_LoadStyleFromString(pFileName string, pString string) bool {
	r, _, _ := xC_LoadStyleFromString.Call(strPtr(pFileName), strPtr(pString))
	return r!=0
} */
