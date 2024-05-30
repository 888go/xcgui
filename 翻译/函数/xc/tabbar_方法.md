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

[func XTabBar_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=TAB条_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XTabBar_AddLabel(hEle int, pName string) int {]
ff=TAB条_添加标签
pName=标签文本内容
hEle=元素句柄

[func XTabBar_InsertLabel(hEle int, index int, pName string) int {]
ff=TAB条插入_标签
pName=标签文本内容
index=插入位置
hEle=元素句柄

[func XTabBar_MoveLabel(hEle int, iSrc int, iDest int) bool {]
ff=TAB条_移动标签
iDest=目标位置索引
iSrc=源位置索引
hEle=元素句柄

[func XTabBar_DeleteLabel(hEle int, index int) bool {]
ff=TAB条_删除标签
index=位置索引
hEle=元素句柄

[func XTabBar_DeleteLabelAll(hEle int) int {]
ff=TAB条_删除全部
hEle=元素句柄

[func XTabBar_GetLabel(hEle int, index int) int {]
ff=TAB条_取标签
index=位置索引
hEle=元素句柄

[func XTabBar_GetLabelClose(hEle int, index int) int {]
ff=TAB条_取标签上的关闭按钮
index=位置索引
hEle=元素句柄

[func XTabBar_GetButtonLeft(hEle int) int {]
ff=TAB条_取左滚动按钮
hEle=元素句柄

[func XTabBar_GetButtonRight(hEle int) int {]
ff=TAB条_取右滚动按钮
hEle=元素句柄

[func XTabBar_GetButtonDropMenu(hEle int) int {]
ff=TAB条_取下拉菜单按钮句柄
hEle=元素句柄

[func XTabBar_GetSelect(hEle int) int {]
ff=TAB条_取当前选择
hEle=元素句柄

[func XTabBar_GetLabelSpacing(hEle int) int {]
ff=TAB条_取间隔
hEle=元素句柄

[func XTabBar_GetLabelCount(hEle int) int {]
ff=TAB条_取标签数量
hEle=元素句柄

[func XTabBar_GetindexByEle(hEle int, hLabel int) int {]
ff=TAB条_取标签位置索引
hLabel=标签按钮句柄
hEle=元素句柄

[func XTabBar_SetLabelSpacing(hEle int, spacing int) int {]
ff=TAB条_置间隔
spacing=标签间隔大小
hEle=元素句柄

[func XTabBar_SetPadding(hEle int, left int, top int, right int, bottom int) int {]
ff=TAB条_置边距
bottom=下
right=右
top=上
left=左
hEle=元素句柄

[func XTabBar_SetSelect(hEle int, index int) int {]
ff=TAB条_置选择
index=标签位置索引
hEle=元素句柄

[func XTabBar_SetUp(hEle int) int {]
ff=TAB条_左滚动
hEle=元素句柄

[func XTabBar_SetDown(hEle int) int {]
ff=TAB条_右滚动
hEle=元素句柄

[func XTabBar_EnableTile(hEle int, bTile bool) int {]
ff=TAB条_启用平铺
bTile=是否启用
hEle=元素句柄

[func XTabBar_EnableDropMenu(hEle int, bEnable bool) int {]
ff=TAB条_启用下拉菜单按钮
hEle=元素句柄

[func XTabBar_EnableClose(hEle int, bEnable bool) int {]
ff=TAB条_启用标签带关闭按钮
bEnable=是否启用
hEle=元素句柄

[func XTabBar_SetCloseSize(hEle int, pSize *SIZE) int {]
ff=TAB条_置关闭按钮大小
pSize=大小值
hEle=元素句柄

[func XTabBar_SetTurnButtonSize(hEle int, pSize *SIZE) int {]
ff=TAB条_置滚动按钮大小
pSize=大小值
hEle=元素句柄

[func XTabBar_SetLabelWidth(hEle int, index int, nWidth int) int {]
ff=TAB条_置指定标签固定宽度
nWidth=宽度
index=索引
hEle=元素句柄

[func XTabBar_ShowLabel(hEle int, index int, bShow bool) bool {]
ff=TAB条_显示标签
bShow=是否显示
index=标签索引
hEle=元素句柄
