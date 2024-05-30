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

[func XToolBar_Create(x, y, cx, cy int32, hParent int) int {]
ff=工具条_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XToolBar_InsertEle(hEle int, hNewEle int, index int) int {]
ff=工具条_插入元素
index=插入位置索引
hNewEle=将要插入的元素
hEle=元素句柄

[func XToolBar_InsertSeparator(hEle int, index int, color int) int {]
ff=工具条_插入分割栏
color=ABGR颜色
index=插入位置索引
hEle=元素句柄

[func XToolBar_EnableButtonMenu(hEle int, bEnable bool) int {]
ff=工具条_启用下拉菜单
bEnable=是否启用
hEle=元素句柄

[func XToolBar_GetEle(hEle int, index int) int {]
ff=工具条_取元素
index=索引值
hEle=元素句柄

[func XToolBar_GetButtonLeft(hEle int) int {]
ff=工具条_取左滚动按钮
hEle=元素句柄

[func XToolBar_GetButtonRight(hEle int) int {]
ff=工具条_取右滚动按钮
hEle=元素句柄

[func XToolBar_GetButtonMenu(hEle int) int {]
ff=工具条_取菜单按钮
hEle=元素句柄

[func XToolBar_SetSpace(hEle int, nSize int) int {]
ff=工具条_置间距
nSize=间距大小
hEle=元素句柄

[func XToolBar_DeleteEle(hEle int, index int) int {]
ff=工具条_删除元素
index=索引值
hEle=元素句柄

[func XToolBar_DeleteAllEle(hEle int) int {]
ff=工具条_删除全部
hEle=元素句柄
