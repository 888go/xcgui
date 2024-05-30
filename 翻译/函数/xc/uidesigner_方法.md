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

[func XC_LoadLayout(pFileName string, hParent int, hAttachWnd uintptr) int {]
ff=炫彩_加载布局文件
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pFileName=布局文件名

[func XC_LoadLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {]
ff=炫彩_加载布局文件Ex
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pFileName=布局文件名

[func XC_LoadLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) int {]
ff=炫彩_加载布局文件ZIP
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pPassword=zip密码
pFileName=布局文件名
pZipFileName=zip文件名

[func XC_LoadLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {]
ff=炫彩_加载布局文件ZIPEx
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pPassword=zip密码
pFileName=布局文件名
pZipFileName=zip文件名

[func XC_LoadLayoutZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) int {]
ff=炫彩_加载布局文件内存ZIP
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pPassword=zip密码
pFileName=布局文件名
data=布局文件数据

[func XC_LoadLayoutZipMemEx(data #左中括号##右中括号#byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {]
ff=炫彩_加载布局文件内存ZIPEx
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pPassword=zip密码
pFileName=布局文件名
data=布局文件数据

[func XC_LoadLayoutFromStringW(pStringXML string, hParent int, hAttachWnd uintptr) int {]
ff=炫彩_加载布局文件从字符串W
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pStringXML=字符串

[func XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {]
ff=炫彩_加载布局文件从字符串WEx
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pStringXML=字符串

[func XC_LoadStyle(pFileName string) bool {]
ff=炫彩_加载样式文件
pFileName=样式文件名称

[func XC_LoadStyleZip(pZipFile string, pFileName string, pPassword string) bool {]
ff=炫彩_加载样式文件ZIP
pPassword=密码
pFileName=文件名
pZipFile=ZIP文件名

[func XC_LoadStyleZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string) bool {]
ff=炫彩_加载样式文件从内存ZIP
pPassword=密码
pFileName=文件名
data=样式文件数据

[func XC_LoadResource(pFileName string) bool {]
ff=炫彩_加载资源文件
pFileName=资源文件名

[func XC_LoadResourceZip(pZipFileName string, pFileName string, pPassword string) bool {]
ff=炫彩_加载资源文件ZIP
pPassword=zip压缩包密码
pFileName=资源文件名
pZipFileName=zip文件名

[func XC_LoadResourceZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string) bool {]
ff=炫彩_加载资源文件内存ZIP
pPassword=zip压缩包密码
pFileName=资源文件名
data=资源文件数据

[func XC_LoadResourceFromStringW(pStringXML string, pFileName string) bool {]
ff=炫彩_加载资源文件从字符串W
pFileName=资源文件名
pStringXML=字符串

[func XC_LoadStyleFromStringW(pFileName string, pString string) bool {]
ff=炫彩_加载样式文件从字符串W
pString=字符串
pFileName=样式文件名

[func XC_LoadLayoutZipResEx(id int32, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) int {]
ff=炫彩_加载布局文件资源ZIPEX
hModule=模块句柄
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pPassword=zip密码
pFileName=布局文件名
id=RC资源ID

[func XC_LoadResourceZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {]
ff=炫彩_加载资源文件资源ZIP
hModule=模块句柄
pPassword=zip压缩包密码
pFileName=资源文件名
id=RC资源ID

[func XC_LoadStyleZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {]
ff=炫彩_加载样式文件从资源ZIP
hModule=模块句柄
pPassword=密码
pFileName=文件名
id=RC资源ID
