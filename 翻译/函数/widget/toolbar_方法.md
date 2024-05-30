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

[func NewToolBar(x, y, cx, cy int32, hParent int) *ToolBar {]
ff=创建工具条
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func NewToolBarByHandle(handle int) *ToolBar {]
ff=创建工具条并按句柄

[func NewToolBarByName(name string) *ToolBar {]
ff=创建工具条并按名称

[func NewToolBarByUID(nUID int) *ToolBar {]
ff=创建工具条并按UID

[func NewToolBarByUIDName(name string) *ToolBar {]
ff=创建工具条并按UID名称

[func (t *ToolBar) InsertEle(hNewEle int, index int) int {]
ff=插入元素
index=插入位置索引
hNewEle=将要插入的元素

[func (t *ToolBar) InsertSeparator(index int, color int) int {]
ff=插入分割栏
color=ABGR颜色
index=插入位置索引

[func (t *ToolBar) EnableButtonMenu(bEnable bool) int {]
ff=启用下拉菜单
bEnable=是否启用

[func (t *ToolBar) GetEle(index int) int {]
ff=取元素
index=索引值

[func (t *ToolBar) GetButtonLeft() int {]
ff=取左滚动按钮

[func (t *ToolBar) GetButtonRight() int {]
ff=取右滚动按钮

[func (t *ToolBar) GetButtonMenu() int {]
ff=取菜单按钮

[func (t *ToolBar) SetSpace(nSize int) int {]
ff=置间距
nSize=间距大小

[func (t *ToolBar) DeleteEle(index int) int {]
ff=删除元素
index=索引值
[func (t *ToolBar) DeleteAllEle() int {]
ff=删除全部
