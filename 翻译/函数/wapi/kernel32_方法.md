# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func GetModuleHandleW(lpModuleName string) uintptr {]
ff=检索指定模块句柄
lpModuleName=已加载模块名称

[func Sleep(ms uint32) {]
ff=延时
ms=毫秒

[func SleepEx(dwMilliseconds uint32, bAlertable bool) uint32 {]
ff=延时EX
dwMilliseconds=毫秒

[func CloseHandle(handle uintptr) bool {]
ff=XCloseHandle
handle=对象句柄

[func GlobalLock(hMem uintptr) uintptr {]
ff=全局内存对象锁定
hMem=内存对象句柄

[func GlobalAlloc(uFlags GMEM_, dwBytes uint) uintptr {]
ff=全局内存分配
dwBytes=字节数
uFlags=内存分配属性

[func GlobalUnlock(hMem uintptr) bool {]
ff=全局内存对象解锁
hMem=内存对象句柄

[func GlobalSize(hMem uintptr) uint {]
ff=全局内存对象取大小
hMem=内存对象句柄

[func LstrcpyW(lpString1, lpString2 uintptr) uintptr {]
ff=复制到缓冲区
lpString2=文本指针2
lpString1=文本指针1

[func GlobalFree(hMem uintptr) uintptr {]
ff=全局内存对象关闭
hMem=内存对象句柄

[func GetLastError() int32 {]
ff=取最后一个错误
