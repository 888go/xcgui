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

[func XMenu_Create() int {]
ff=菜单_创建

[func XMenu_AddItem(hMenu int, nID int32, pText string, nParentID int32, nFlags xcc.Menu_Item_Flag_) {]
ff=菜单_添加项
nFlags=标识
nParentID=父项ID
pText=文本内容
nID=项ID
hMenu=菜单句柄

[func XMenu_AddItemIcon(hMenu int, nID int32, pText string, nParentID int32, hIcon int, nFlags xcc.Menu_Item_Flag_) {]
ff=菜单_添加项图标
nFlags=标识
hIcon=菜单项图标句柄
nParentID=父项ID
pText=文本内容
nID=项ID
hMenu=菜单句柄

[func XMenu_InsertItem(hMenu int, nID int32, pText string, nFlags xcc.Menu_Item_Flag_, insertID int32) {]
ff=菜单_插入项
nFlags=标识
pText=文本内容
nID=项ID
hMenu=菜单句柄

[func XMenu_InsertItemIcon(hMenu int, nID int32, pText string, hIcon int, nFlags xcc.Menu_Item_Flag_, insertID int32) {]
ff=菜单_插入项图标
nFlags=标识
hIcon=菜单项图标句柄
pText=文本内容
nID=项ID
hMenu=菜单句柄

[func XMenu_GetFirstChildItem(hMenu int, nID int32) int32 {]
ff=菜单_取第一个子项
nID=项ID
hMenu=菜单句柄

[func XMenu_GetEndChildItem(hMenu int, nID int32) int32 {]
ff=菜单_取末尾子项
nID=项ID
hMenu=菜单句柄

[func XMenu_GetPrevSiblingItem(hMenu int, nID int32) int32 {]
ff=菜单_取上一个兄弟项
nID=项ID
hMenu=菜单句柄

[func XMenu_GetNextSiblingItem(hMenu int, nID int32) int32 {]
ff=菜单_取下一个兄弟项
nID=项ID
hMenu=菜单句柄

[func XMenu_GetParentItem(hMenu int, nID int32) int32 {]
ff=菜单_取父项
nID=项ID
hMenu=菜单句柄

[func XMenu_SetAutoDestroy(hMenu int, bAuto bool) {]
ff=菜单_置自动销毁
bAuto=是否自动销毁
hMenu=菜单句柄

[func XMenu_EnableDrawBackground(hMenu int, bEnable bool) {]
ff=菜单_启用用户绘制背景
bEnable=是否启用
hMenu=菜单句柄

[func XMenu_EnableDrawItem(hMenu int, bEnable bool) {]
ff=菜单_启用用户绘制项
bEnable=是否启用
hMenu=菜单句柄

[func XMenu_Popup(hMenu int, hParentWnd uintptr, x, y int32, hParentEle int, nPosition xcc.Menu_Popup_Position_) bool {]
ff=菜单_弹出
nPosition=弹出位置
hParentEle=父元素句柄
y=y坐标
x=x坐标
hParentWnd=父窗口真实句柄
hMenu=菜单句柄

[func XMenu_DestroyMenu(hMenu int) {]
ff=菜单_销毁
hMenu=菜单句柄

[func XMenu_CloseMenu(hMenu int) {]
ff=菜单_关闭
hMenu=菜单句柄

[func XMenu_SetBkImage(hMenu int, hImage int) {]
ff=菜单_置背景图片
hImage=图片句柄
hMenu=菜单句柄

[func XMenu_SetItemText(hMenu int, nID int32, pText string) bool {]
ff=菜单_置项文本
pText=文本内容
nID=项ID
hMenu=菜单句柄

[func XMenu_GetItemText(hMenu int, nID int32) string {]
ff=菜单_取项文本
nID=项ID
hMenu=菜单句柄

[func XMenu_GetItemTextLength(hMenu int, nID int32) int32 {]
ff=菜单_取项文本长度
nID=项ID
hMenu=菜单句柄

[func XMenu_SetItemIcon(hMenu int, nID int32, hIcon int) bool {]
ff=菜单_置项图标
hIcon=菜单项图标句柄
nID=项ID
hMenu=菜单句柄

[func XMenu_SetItemFlags(hMenu int, nID int32, uFlags xcc.Menu_Item_Flag_) bool {]
ff=菜单_置项标志
uFlags=标识
nID=项ID
hMenu=菜单句柄

[func XMenu_SetItemHeight(hMenu int, height int32) {]
ff=菜单_置项高度
height=高度
hMenu=菜单句柄

[func XMenu_GetItemHeight(hMenu int) int32 {]
ff=菜单_取项高度
hMenu=菜单句柄

[func XMenu_SetBorderColor(hMenu int, crColor int) {]
ff=菜单_置边框颜色
crColor=ABGR颜色
hMenu=菜单句柄

[func XMenu_SetBorderSize(hMenu int, nLeft, nTop, nRight, nBottom int32) {]
ff=菜单_置边框大小
nBottom=下
nRight=右
nTop=上
nLeft=左
hMenu=菜单句柄

[func XMenu_GetLeftWidth(hMenu int) int32 {]
ff=菜单_取左侧宽度
hMenu=菜单句柄

[func XMenu_GetLeftSpaceText(hMenu int) int32 {]
ff=菜单_取左侧文本间隔
hMenu=菜单句柄

[func XMenu_GetItemCount(hMenu int) int32 {]
ff=菜单_取项数量
hMenu=菜单句柄

[func XMenu_SetItemCheck(hMenu int, nID int32, bCheck bool) bool {]
ff=菜单_置项勾选
bCheck=勾选TRUE
nID=菜单项ID
hMenu=菜单句柄

[func XMenu_IsItemCheck(hMenu int, nID int32) bool {]
ff=菜单_判断项勾选
nID=菜单项ID
hMenu=菜单句柄

[func XMenu_SetItemWidth(hMenu int, nID, nWidth int32) bool {]
ff=菜单_置项宽度
nWidth=指定文本区域宽度
nID=项ID
hMenu=菜单句柄

[func XMenu_GetMenuBar(hMenu int) int {]
ff=菜单_取菜单条
hMenu=菜单句柄
