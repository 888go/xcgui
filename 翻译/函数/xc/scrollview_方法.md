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

[func XSView_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=滚动视_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XSView_SetTotalSize(hEle int, cx int, cy int) bool {]
ff=滚动视_置视图大小
cy=高度
cx=宽度
hEle=元素句柄

[func XSView_GetTotalSize(hEle int, pSize *SIZE) int {]
ff=滚动视_取视图大小
pSize=大小
hEle=元素句柄

[func XSView_SetLineSize(hEle int, nWidth int, nHeight int) bool {]
ff=滚动视_置滚动单位大小
nHeight=高度
nWidth=宽度
hEle=元素句柄

[func XSView_GetLineSize(hEle int, pSize *SIZE) int {]
ff=滚动视_取滚动单位大小
pSize=返回大小
hEle=元素句柄

[func XSView_SetScrollBarSize(hEle int, size int) int {]
ff=滚动视_置滚动条大小
size=滚动条大小
hEle=元素句柄

[func XSView_GetViewPosH(hEle int) int {]
ff=滚动视_取视口原点X
hEle=元素句柄

[func XSView_GetViewPosV(hEle int) int {]
ff=滚动视_取视口原点Y
hEle=元素句柄

[func XSView_GetViewWidth(hEle int) int {]
ff=滚动视_取视口宽度
hEle=元素句柄

[func XSView_GetViewHeight(hEle int) int {]
ff=滚动视_取视口高度
hEle=元素句柄

[func XSView_GetViewRect(hEle int, pRect *RECT) int {]
ff=滚动视_取视口坐标
pRect=坐标
hEle=元素句柄

[func XSView_GetScrollBarH(hEle int) int {]
ff=滚动视_取水平滚动条
hEle=元素句柄

[func XSView_GetScrollBarV(hEle int) int {]
ff=滚动视_取垂直滚动条
hEle=元素句柄

[func XSView_ScrollPosH(hEle int, pos int) bool {]
ff=滚动视_水平滚动
pos=位置点
hEle=元素句柄

[func XSView_ScrollPosV(hEle int, pos int) bool {]
ff=滚动视_垂直滚动
pos=位置点
hEle=元素句柄

[func XSView_ScrollPosXH(hEle int, posX int) bool {]
ff=滚动视_水平滚动到X
posX=坐标
hEle=元素句柄

[func XSView_ScrollPosYV(hEle int, posY int) bool {]
ff=滚动视_垂直滚动到Y
posY=Y坐标
hEle=元素句柄

[func XSView_ShowSBarH(hEle int, bShow bool) int {]
ff=滚动视_显示水平滚动条
bShow=是否显示
hEle=元素句柄

[func XSView_ShowSBarV(hEle int, bShow bool) int {]
ff=滚动视_显示垂直滚动条
bShow=是否显示
hEle=元素句柄

[func XSView_EnableAutoShowScrollBar(hEle int, bEnable bool) int {]
ff=滚动视_启用自动显示滚动条
bEnable=是否启用
hEle=元素句柄

[func XSView_ScrollLeftLine(hEle int) bool {]
ff=滚动视_向左滚动
hEle=元素句柄

[func XSView_ScrollRightLine(hEle int) bool {]
ff=滚动视_向右滚动
hEle=元素句柄

[func XSView_ScrollTopLine(hEle int) bool {]
ff=滚动视_向上滚动
hEle=元素句柄

[func XSView_ScrollBottomLine(hEle int) bool {]
ff=滚动视_向下滚动
hEle=元素句柄

[func XSView_ScrollLeft(hEle int) bool {]
ff=滚动视_滚动到左侧
hEle=元素句柄

[func XSView_ScrollRight(hEle int) bool {]
ff=滚动视_滚动到右侧
hEle=元素句柄

[func XSView_ScrollTop(hEle int) bool {]
ff=滚动视_滚动到顶部
hEle=元素句柄

[func XSView_ScrollBottom(hEle int) bool {]
ff=滚动视_滚动到底部
hEle=元素句柄
