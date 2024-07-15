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

[func NewLayoutEle(x int, y int, cx int, cy int, hParent int) *LayoutEle {]
ff=创建布局
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewLayoutEleEx(hParent int) *LayoutEle {]
ff=创建布局EX
hParent=父窗口句柄或元素句柄

[func NewLayoutEleByHandle(handle int) *LayoutEle {]
ff=创建布局并按句柄

[func NewLayoutEleByLayout(pFileName string, hParent int, hAttachWnd uintptr) *LayoutEle {]
ff=创建布局并按布局文件
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pFileName=布局文件名

[func NewLayoutEleByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *LayoutEle {]
ff=创建布局并按压缩包中布局文件
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pPassword=zip密码
pFileName=布局文件名
pZipFileName=zip文件名

[func NewLayoutEleByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *LayoutEle {]
ff=创建布局并按RC资源压缩包中布局文件
hModule=模块句柄
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称
pPassword=zip密码
pFileName=布局文件名
id=RC资源ID

[func NewLayoutEleByLayoutZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *LayoutEle {]
ff=创建布局并按内存压缩包中布局文件
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pPassword=zip密码
pFileName=布局文件名
data=布局文件数据

[func NewLayoutEleByStringW(pStringXML string, hParent int, hAttachWnd uintptr) *LayoutEle {]
ff=创建布局并按布局文件字符串W
hAttachWnd=附加窗口句柄
hParent=父对象
pStringXML=xml字符串

[func NewLayoutEleByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {]
ff=创建布局并按布局文件EX
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pFileName=布局文件名

[func NewLayoutEleByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {]
ff=创建布局并按压缩包中的布局文件EX
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pPassword=zip密码
pFileName=布局文件名
pZipFileName=zip文件名

[func NewLayoutEleByLayoutZipMemEx(data #左中括号##右中括号#byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {]
ff=创建布局并按内存压缩包中布局文件EX
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称
pPassword=zip密码
pFileName=布局文件名
data=布局文件数据

[func NewLayoutEleByStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {]
ff=创建布局并按布局文件字符串WEX
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象
pPrefixName=名称前缀
pStringXML=字符串

[func NewLayoutEleByName(name string) *LayoutEle {]
ff=创建布局并按名称

[func NewLayoutEleByID(hWindow, nID int) *LayoutEle {]
ff=创建布局并按ID
hWindow=父窗口句柄

[func NewLayoutEleByIDName(hWindow int, name string) *LayoutEle {]
ff=创建布局并按ID名称
name=名称
hWindow=父窗口句柄

[func NewLayoutEleByUID(nUID int) *LayoutEle {]
ff=创建布局并按UID

[func NewLayoutEleByUIDName(name string) *LayoutEle {]
ff=创建布局并按UID名称
name=名称

[func (l *LayoutEle) IsEnableLayout() bool {]
ff=判断启用

[func (l *LayoutEle) EnableLayout(bEnable bool) int {]
ff=启用
bEnable=是否启用

[func (l *LayoutEle) ShowLayoutFrame(bEnable bool) int {]
ff=显示布局边界
bEnable=是否显示

[func (l *LayoutEle) GetWidthIn() int {]
ff=取内宽度

[func (l *LayoutEle) GetHeightIn() int {]
ff=取内高度

[func (l *LayoutEle) EnableHorizon(bEnable bool) int {]
ff=启用水平
bEnable=是否启用

[func (l *LayoutEle) EnableAutoWrap(bEnable bool) int {]
ff=启用自动换行
bEnable=是否启用

[func (l *LayoutEle) EnableOverflowHide(bEnable bool) int {]
ff=启用溢出隐藏
bEnable=是否启用

[func (l *LayoutEle) SetAlignH(nAlign xcc.Layout_Align_) int {]
ff=置水平对齐
nAlign=对齐方式

[func (l *LayoutEle) SetAlignV(nAlign xcc.Layout_Align_) int {]
ff=置垂直对齐
nAlign=对齐方式

[func (l *LayoutEle) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) int {]
ff=置对齐基线
nAlign=对齐方式

[func (l *LayoutEle) SetSpace(nSpace int) int {]
ff=置间距
nSpace=项间距大小

[func (l *LayoutEle) SetSpaceRow(nSpace int) int {]
ff=置行距
nSpace=行间距大小
