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

[func XMenuBar_Create(x, y, cx, cy int32, hParent int) int {]
ff=菜单条_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XMenuBar_AddButton(hEle int, pText string) int32 {]
ff=菜单条_添加按钮
pText=文本内容
hEle=元素句柄

[func XMenuBar_SetButtonHeight(hEle int, height int32) {]
ff=菜单条_置按钮高度
height=高度
hEle=元素句柄

[func XMenuBar_GetMenu(hEle int, nIndex int32) int {]
ff=菜单条_取菜单
nIndex=菜单条按钮索引
hEle=元素句柄

[func XMenuBar_DeleteButton(hEle int, nIndex int32) bool {]
ff=菜单条_删除按钮
nIndex=菜单条按钮索引
hEle=元素句柄

[func XMenuBar_EnableAutoWidth(hEle int, bEnable bool) {]
ff=菜单条_启用自动宽度
bEnable=是否启用
hEle=元素句柄

[func XMenuBar_GetButton(hEle int, nIndex int32) bool {]
ff=菜单条_取菜单按钮
nIndex=菜单条按钮索引
hEle=元素句柄

[func XMenuBar_GetSelect(hEle int) int32 {]
ff=菜单条_取选择项
hEle=元素句柄
