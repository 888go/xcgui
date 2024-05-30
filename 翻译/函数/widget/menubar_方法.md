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

[func NewMenuBar(x int32, y int32, cx int32, cy int32, hParent int) *MenuBar {]
ff=创建菜单条
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewMenuBarByHandle(handle int) *MenuBar {]
ff=创建菜单条并按句柄

[func NewMenuBarByName(name string) *MenuBar {]
ff=创建菜单条并按名称

[func NewMenuBarByUID(nUID int) *MenuBar {]
ff=创建菜单条并按UID

[func NewMenuBarByUIDName(name string) *MenuBar {]
ff=创建菜单条并按UID名称

[func (m *MenuBar) AddButton(pText string) int32 {]
ff=添加按钮
pText=文本内容

[func (m *MenuBar) SetButtonHeight(height int32) {]
ff=置按钮高度
height=高度

[func (m *MenuBar) GetMenu(nIndex int32) int {]
ff=取菜单
nIndex=按钮索引

[func (m *MenuBar) DeleteButton(nIndex int32) bool {]
ff=删除按钮
nIndex=按钮索引

[func (m *MenuBar) EnableAutoWidth(bEnable bool) {]
ff=启用自动宽度
bEnable=是否启用

[func (m *MenuBar) GetButton(nIndex int32) bool {]
ff=取菜单按钮
nIndex=按钮索引

[func (m *MenuBar) GetSelect() int32 {]
ff=取选择项
