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

[func NewMenu() *Menu {]
ff=创建菜单

[func NewMenuByHandle(handle int) *Menu {]
ff=创建菜单并按句柄

[func NewMenuByName(name string) *Menu {]
ff=创建菜单并按名称

[func NewMenuByUID(nUID int) *Menu {]
ff=创建菜单并按UID

[func NewMenuByUIDName(name string) *Menu {]
ff=创建菜单并按UID名称

[func (m *Menu) AddItem(nID int32, pText string, nParentID int32, nFlags xcc.Menu_Item_Flag_) {]
ff=添加项
nFlags=标识
nParentID=父项ID
pText=文本内容
nID=项ID

[func (m *Menu) AddItemIcon(nID int32, pText string, nParentID int32, hIcon int, nFlags xcc.Menu_Item_Flag_) {]
ff=添加项图标
nFlags=标识
hIcon=图标句柄
nParentID=父项ID
pText=文本内容
nID=项ID

[func (m *Menu) InsertItem(nID int32, pText string, nFlags xcc.Menu_Item_Flag_, insertID int32) {]
ff=插入项
nFlags=标识
pText=文本内容
nID=项ID

[func (m *Menu) InsertItemIcon(nID int32, pText string, hIcon int, nFlags xcc.Menu_Item_Flag_, insertID int32) {]
ff=插入项图标
nFlags=标识
hIcon=图标句柄
pText=文本内容
nID=项ID

[func (m *Menu) GetFirstChildItem(nID int32) int32 {]
ff=取第一个子项
nID=项ID

[func (m *Menu) GetEndChildItem(nID int32) int32 {]
ff=取末尾子项
nID=项ID

[func (m *Menu) GetPrevSiblingItem(nID int32) int32 {]
ff=取上一个兄弟项
nID=项ID

[func (m *Menu) GetNextSiblingItem(nID int32) int32 {]
ff=取下一个兄弟项
nID=项ID

[func (m *Menu) GetParentItem(nID int32) int32 {]
ff=取父项
nID=项ID

[func (m *Menu) SetAutoDestroy(bAuto bool) {]
ff=置自动销毁
bAuto=是否自动销毁

[func (m *Menu) EnableDrawBackground(bEnable bool) {]
ff=启用用户绘制背景
bEnable=是否启用

[func (m *Menu) EnableDrawItem(bEnable bool) {]
ff=启用用户绘制项
bEnable=是否启用

[func (m *Menu) Popup(hParentWnd uintptr, x, y int32, hParentEle int, nPosition xcc.Menu_Popup_Position_) bool {]
ff=弹出
nPosition=弹出位置
hParentEle=父元素句柄
y=y坐标
x=x坐标
hParentWnd=父窗口句柄

[func (m *Menu) DestroyMenu() {]
ff=销毁

[func (m *Menu) CloseMenu() {]
ff=关闭

[func (m *Menu) SetBkImage(hImage int) {]
ff=置背景图片
hImage=图片句柄

[func (m *Menu) SetItemText(nID int32, pText string) bool {]
ff=置项文本
pText=文本内容
nID=项ID

[func (m *Menu) GetItemText(nID int32) string {]
ff=取项文本
nID=项ID

[func (m *Menu) GetItemTextLength(nID int32) int32 {]
ff=取项文本长度
nID=项ID

[func (m *Menu) SetItemIcon(nID int32, hIcon int) bool {]
ff=置项图标
hIcon=菜单项图标句柄
nID=项ID

[func (m *Menu) SetItemFlags(nID int32, uFlags xcc.Menu_Item_Flag_) bool {]
ff=置项标志
uFlags=标识
nID=项ID

[func (m *Menu) SetItemHeight(height int32) {]
ff=置项高度
height=高度

[func (m *Menu) GetItemHeight() int32 {]
ff=取项高度

[func (m *Menu) SetBorderColor(crColor int) {]
ff=置边框颜色
crColor=ABGR颜色

[func (m *Menu) SetBorderSize(nLeft, nTop, nRight, nBottom int32) {]
ff=置边框大小
nBottom=下
nRight=右
nTop=上
nLeft=左

[func (m *Menu) GetLeftWidth() int32 {]
ff=取左侧宽度

[func (m *Menu) GetLeftSpaceText() int32 {]
ff=取左侧文本间隔

[func (m *Menu) GetItemCount() int32 {]
ff=取项数量

[func (m *Menu) SetItemCheck(nID int32, bCheck bool) bool {]
ff=置项勾选
bCheck=勾选TRUE
nID=菜单项ID

[func (m *Menu) IsItemCheck(nID int32) bool {]
ff=判断项勾选
nID=菜单项ID

[func (m *Menu) SetItemWidth(nID, nWidth int32) bool {]
ff=置项宽度
nWidth=宽度
nID=项ID
[func (m *Menu) GetMenuBar() int {]
ff=取菜单条
