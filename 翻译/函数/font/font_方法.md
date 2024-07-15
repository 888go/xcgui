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

[func New(size int32) *Font {]
ff=创建
size=字体大小

[func NewEX(pName string, size int32, style xcc.FontStyle_) *Font {]
ff=创建EX
style=字体样式
size=字体大小
pName=字体名称

[func NewLOGFONTW(pFontInfo *xc.LOGFONTW) *Font {]
ff=创建并按选项
pFontInfo=字体选项

[func NewByHFONT(hFont uintptr) *Font {]
ff=创建并按字体句柄
hFont=字体句柄

[func NewByFont(pFont uintptr) *Font {]
ff=创建并按GDI字体指针
pFont=GDI字体指针

[func NewByFile(pFontFile string, size int32, style xcc.FontStyle_) *Font {]
ff=创建并按文件
style=字体样式
size=字体大小
pFontFile=字体文件名

[func NewByZip(pZipFileName, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) *Font {]
ff=创建并按ZIP
style=字体样式
fontSize=字体大小
pPassword=zip密码
pFileName=字体文件名
pZipFileName=zip文件名

[func NewByZipMem(data #左中括号##右中括号#byte, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) *Font {]
ff=创建并按内存ZIP
style=字体样式
fontSize=字体大小
pPassword=zip密码
pFileName=字体文件名
data=zip数据

[func NewByMem(data #左中括号##右中括号#byte, fontSize int32, style xcc.FontStyle_) *Font {]
ff=创建并按内存
style=字体样式
fontSize=字体大小
data=字体文件数据

[func NewByRes(id int32, pType string, fontSize int32, style xcc.FontStyle_, hModule uintptr) *Font {]
ff=创建并按资源
style=字体样式
fontSize=字体大小
pType=类型

[func NewByHandle(handle int) *Font {]
ff=创建并按句柄
handle=句柄

[func NewByName(name string) *Font {]
ff=创建并按名称
name=名称

[func (f *Font) EnableAutoDestroy(bEnable bool) int {]
ff=启用自动销毁
bEnable=是否启用

[func (f *Font) GetFont() int {]
ff=取字体指针

[func (f *Font) GetFontInfo(pInfo *xc.Font_Info_) int {]
ff=取信息
pInfo=返回字体信息

[func (f *Font) GetLOGFONTW(hdc uintptr, pOut *xc.LOGFONTW) bool {]
ff=取属性
pOut=返回字体属性
hdc=hdc句柄

[func (f *Font) Destroy() {]
ff=销毁

[func (f *Font) AddRef() {]
ff=增加引用计数

[func (f *Font) GetRefCount() int32 {]
ff=取引用计数

[func (f *Font) Release() {]
ff=释放引用计数
