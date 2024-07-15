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
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
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
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func XFont_Create(size int32) int {]
ff=字体_创建
size=字体大小

[func XFont_CreateEx(pName string, size int32, style xcc.FontStyle_) int {]
ff=字体_创建EX
style=字体样式
size=字体大小
pName=字体名称

[func XFont_CreateLOGFONTW(pFontInfo *LOGFONTW) int {]
ff=字体_创建从LOGFONT
pFontInfo=字体信息

[func XFont_CreateFromHFONT(hFont uintptr) int {]
ff=字体_创建从HFONT
hFont=字体句柄

[func XFont_CreateFromFont(pFont uintptr) int {]
ff=字体_创建从Font
pFont=GDI字体指针

[func XFont_CreateFromFile(pFontFile string, size int32, style xcc.FontStyle_) int {]
ff=字体_创建从文件
style=字体样式
size=字体大小
pFontFile=字体文件名

[func XFont_CreateFromZip(pZipFileName, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) int {]
ff=字体_创建从ZIP
style=字体样式
fontSize=字体大小
pPassword=zip密码
pFileName=字体文件名
pZipFileName=zip文件名

[func XFont_CreateFromZipMem(data #左中括号##右中括号#byte, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) int {]
ff=字体_创建从内存ZIP
style=字体样式
fontSize=字体大小
pPassword=zip密码
pFileName=字体文件名
data=zip数据

[func XFont_CreateFromMem(data #左中括号##右中括号#byte, fontSize int32, style xcc.FontStyle_) int {]
ff=字体_创建从内存
style=字体样式
fontSize=字体大小
data=字体文件数据

[func XFont_CreateFromRes(id int32, pType string, fontSize int32, style xcc.FontStyle_, hModule uintptr) int {]
ff=字体_创建从资源
style=字体样式
fontSize=字体大小
pType=类型

[func XFont_EnableAutoDestroy(hFontX int, bEnable bool) int {]
ff=字体_启用自动销毁
bEnable=是否启用
hFontX=字体句柄

[func XFont_GetFont(hFontX int) int {]
ff=字体_取Font
hFontX=字体句柄

[func XFont_GetFontInfo(hFontX int, pInfo *Font_Info_) int {]
ff=字体_取信息
pInfo=接收返回的字体信息
hFontX=字体句柄

[func XFont_GetLOGFONTW(hFontX int, hdc uintptr, pOut *LOGFONTW) bool {]
ff=字体_取LOGFONTW
pOut=接收返回信息
hdc=hdc句柄
hFontX=字体句柄

[func XFont_Destroy(hFontX int) {]
ff=字体_销毁
hFontX=字体句柄

[func XFont_AddRef(hFontX int) {]
ff=字体_增加引用计数
hFontX=字体句柄

[func XFont_GetRefCount(hFontX int) int32 {]
ff=字体_取引用计数
hFontX=字体句柄

[func XFont_Release(hFontX int) {]
ff=字体_释放引用计数
hFontX=字体句柄
