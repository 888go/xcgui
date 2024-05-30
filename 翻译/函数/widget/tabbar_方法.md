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

[func NewTabBar(x int, y int, cx int, cy int, hParent int) *TabBar {]
ff=创建TAB条
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewTabBarByHandle(handle int) *TabBar {]
ff=创建TAB条并按句柄

[func NewTabBarByName(name string) *TabBar {]
ff=创建TAB条并按名称

[func NewTabBarByUID(nUID int) *TabBar {]
ff=创建TAB条并按UID

[func NewTabBarByUIDName(name string) *TabBar {]
ff=创建TAB条并按UID名称

[func (t *TabBar) AddLabel(pName string) int {]
ff=添加标签
pName=标签文本

[func (t *TabBar) InsertLabel(index int, pName string) int {]
ff=TAB条插入_标签
pName=标签文本
index=插入位置

[func (t *TabBar) MoveLabel(iSrc int, iDest int) bool {]
ff=移动标签
iDest=目标位置索引
iSrc=源位置索引

[func (t *TabBar) DeleteLabel(index int) bool {]
ff=删除标签
index=位置索引

[func (t *TabBar) DeleteLabelAll() int {]
ff=删除全部

[func (t *TabBar) GetLabel(index int) int {]
ff=取标签
index=位置索引

[func (t *TabBar) GetLabelClose(index int) int {]
ff=取标签上的关闭按钮
index=位置索引

[func (t *TabBar) GetButtonLeft() int {]
ff=取左滚动按钮

[func (t *TabBar) GetButtonRight() int {]
ff=取右滚动按钮

[func (t *TabBar) GetButtonDropMenu() int {]
ff=取下拉菜单按钮句柄

[func (t *TabBar) GetSelect() int {]
ff=取当前选择

[func (t *TabBar) GetLabelSpacing() int {]
ff=取间隔

[func (t *TabBar) GetLabelCount() int {]
ff=取标签数量

[func (t *TabBar) GetindexByEle(hLabel int) int {]
ff=取标签位置索引
hLabel=标签按钮句柄

[func (t *TabBar) SetLabelSpacing(spacing int) int {]
ff=置间隔
spacing=标签间隔大小

[func (t *TabBar) SetPadding(left int, top int, right int, bottom int) int {]
ff=置边距
bottom=下
right=右
top=上
left=左

[func (t *TabBar) SetSelect(index int) int {]
ff=置选择
index=索引

[func (t *TabBar) SetUp() int {]
ff=左滚动

[func (t *TabBar) SetDown() int {]
ff=右滚动

[func (t *TabBar) EnableTile(bTile bool) int {]
ff=启用平铺
bTile=是否启用

[func (t *TabBar) EnableDropMenu(bEnable bool) int {]
ff=启用下拉菜单按钮
bEnable=是否启用

[func (t *TabBar) EnableClose(bEnable bool) int {]
ff=启用标签带关闭按钮
bEnable=是否启用

[func (t *TabBar) SetCloseSize(pSize *xc.SIZE) int {]
ff=置关闭按钮大小
pSize=大小值

[func (t *TabBar) SetTurnButtonSize(pSize *xc.SIZE) int {]
ff=置滚动按钮大小
pSize=大小值

[func (t *TabBar) SetLabelWidth(index int, nWidth int) int {]
ff=置指定标签固定宽度
nWidth=宽度
index=索引

[func (t *TabBar) ShowLabel(index int, bShow bool) bool {]
ff=显示标签
bShow=是否显示
index=标签索引

[func (t *TabBar) Event_TABBAR_SELECT(pFun XE_TABBAR_SELECT) bool {]
ff=事件_选择改变

[func (t *TabBar) Event_TABBAR_SELECT1(pFun XE_TABBAR_SELECT1) bool {]
ff=事件_选择改变1

[func (t *TabBar) Event_TABBAR_DELETE(pFun XE_TABBAR_DELETE) bool {]
ff=事件_删除

[func (t *TabBar) Event_TABBAR_DELETE1(pFun XE_TABBAR_DELETE1) bool {]
ff=事件_删除1
