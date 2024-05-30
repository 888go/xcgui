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

[func NewTree(x, y, cx, cy int32, hParent int) *Tree {]
ff=创建列表树
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewTreeEx(x, y, cx, cy int32, hParent, col_extend_count int32) *Tree {]
ff=创建列表树Ex
col_extend_count=列数量
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewTreeByHandle(handle int) *Tree {]
ff=创建列表树并按句柄

[func NewTreeByName(name string) *Tree {]
ff=创建列表树并按名称

[func NewTreeByUID(nUID int) *Tree {]
ff=创建列表树并按UID

[func NewTreeByUIDName(name string) *Tree {]
ff=创建列表树并按UID名称

[func (t *Tree) EnableDragItem(bEnable bool) {]
ff=启用拖动项
bEnable=是否启用

[func (t *Tree) EnableConnectLine(bEnable bool, bSolid bool) {]
ff=启用连接线
bSolid=实线或虚线
bEnable=是否启用

[func (t *Tree) EnableExpand(bEnable bool) {]
ff=启用展开
bEnable=是否启用

[func (t *Tree) EnableTemplateReuse(bEnable bool) {]
ff=启用模板复用
bEnable=是否启用

[func (t *Tree) SetConnectLineColor(color int) {]
ff=置连接线颜色
color=ABGR颜色

[func (t *Tree) SetExpandButtonSize(nWidth, nHeight int32) {]
ff=置展开按钮大小
nHeight=高度
nWidth=宽度

[func (t *Tree) SetConnectLineLength(nLength int32) {]
ff=置连接线长度
nLength=连线绘制长度

[func (t *Tree) SetDragInsertPositionColor(color int) {]
ff=置拖动项插入位置颜色
color=ABGR颜色

[func (t *Tree) SetItemTemplateXML(pXmlFile string) bool {]
ff=置项模板文件
pXmlFile=文件名

[func (t *Tree) SetItemTemplateXMLSel(pXmlFile string) bool {]
ff=置选择项模板文件
pXmlFile=文件名

[func (t *Tree) SetItemTemplate(hTemp int) bool {]
ff=置项模板
hTemp=模板句柄

[func (t *Tree) SetItemTemplateSel(hTemp int) bool {]
ff=置选择项模板
hTemp=模板句柄

[func (t *Tree) SetItemTemplateXMLFromString(pStringXML string) bool {]
ff=置项模板并按字符串
pStringXML=字符串

[func (t *Tree) SetItemTemplateXMLSelFromString(pStringXML string) bool {]
ff=置选择项模板并按字符串
pStringXML=字符串

[func (t *Tree) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) {]
ff=置项背景绘制标志
nFlags=标志位

[func (t *Tree) SetItemData(nID int32, nUserData int) bool {]
ff=置项数据
nUserData=用户数据
nID=项ID

[func (t *Tree) GetItemData(nID int32) int {]
ff=取项数据
nID=项ID

[func (t *Tree) SetSelectItem(nID int32) bool {]
ff=置选择项
nID=项ID

[func (t *Tree) GetSelectItem() int32 {]
ff=取选择项

[func (t *Tree) VisibleItem(nID int32) {]
ff=可视指定项
nID=项索引

[func (t *Tree) IsExpand(nID int32) bool {]
ff=判断展开
nID=项ID

[func (t *Tree) ExpandItem(nID int32, bExpand bool) bool {]
ff=展开项
bExpand=是否展开
nID=项ID

[func (t *Tree) ExpandAllChildItem(nID int32, bExpand bool) bool {]
ff=展开全部子项
bExpand=是否展开
nID=项ID

[func (t *Tree) HitTest(pPt *xc.POINT) int32 {]
ff=测试点击项
pPt=坐标点

[func (t *Tree) HitTestOffset(pPt *xc.POINT) int32 {]
ff=测试点击项EX
pPt=坐标点

[func (t *Tree) GetFirstChildItem(nID int32) int32 {]
ff=取第一个子项
nID=项ID

[func (t *Tree) GetEndChildItem(nID int32) int32 {]
ff=取末尾子项
nID=项ID

[func (t *Tree) GetPrevSiblingItem(nID int32) int32 {]
ff=取上一个兄弟项
nID=项ID

[func (t *Tree) GetNextSiblingItem(nID int32) int32 {]
ff=取下一个兄弟项
nID=项ID

[func (t *Tree) GetParentItem(nID int32) int32 {]
ff=取父项
nID=项ID

[func (t *Tree) CreateAdapter() int {]
ff=创建数据适配器

[func (t *Tree) BindAdapter(hAdapter int) {]
ff=绑定数据适配器
hAdapter=数据适配器句柄

[func (t *Tree) GetAdapter() int {]
ff=取数据适配器

[func (t *Tree) RefreshData() {]
ff=刷新数据

[func (t *Tree) RefreshItem(nID int32) {]
ff=刷新指定项
nID=项ID

[func (t *Tree) SetIndentation(nWidth int32) {]
ff=置缩进
nWidth=缩进宽度

[func (t *Tree) GetIndentation() int32 {]
ff=取缩进

[func (t *Tree) SetItemHeightDefault(nHeight, nSelHeight int32) {]
ff=置项默认高度
nSelHeight=选中时高度
nHeight=高度

[func (t *Tree) GetItemHeightDefault(pHeight, pSelHeight *int32) {]
ff=取项默认高度
pSelHeight=接收返回值
pHeight=接收返回高度

[func (t *Tree) SetItemHeight(nID, nHeight, nSelHeight int32) {]
ff=置项高度
nSelHeight=选中时高度
nHeight=高度
nID=项ID

[func (t *Tree) GetItemHeight(nID int32, pHeight, pSelHeight *int32) {]
ff=取项高度
pSelHeight=接收返回值
pHeight=接收返回高度
nID=项ID

[func (t *Tree) SetRowSpace(nSpace int32) {]
ff=置行间距
nSpace=行间隔大小

[func (t *Tree) GetRowSpace() int32 {]
ff=取行间距

[func (t *Tree) MoveItem(nMoveItem, nDestItem, nFlag int32) bool {]
ff=移动项
nFlag=位置
nDestItem=目标项ID
nMoveItem=要移动的项ID

[func (t *Tree) GetTemplateObject(nID, nTempItemID int32) int {]
ff=取模板对象
nTempItemID=模板项ID
nID=项ID

[func (t *Tree) GetItemIDFromHXCGUI(hXCGUI int) int32 {]
ff=取对象所在项
hXCGUI=对象句柄

[func (t *Tree) InsertItemText(pValue string, nParentID, insertID int32) int32 {]
ff=插入项文本
insertID=插入id
nParentID=父id
pValue=值

[func (t *Tree) InsertItemTextEx(pName string, pValue string, nParentID, insertID int32) int32 {]
ff=插入项文本EX
insertID=插入id
nParentID=父id
pValue=值
pName=名称

[func (t *Tree) InsertItemImage(hImage int, nParentID, insertID int32) int32 {]
ff=插入项图片
insertID=插入id
nParentID=父id
hImage=图片

[func (t *Tree) InsertItemImageEx(pName string, hImage int, nParentID, insertID int32) int32 {]
ff=插入项图片EX
insertID=插入id
nParentID=父id
hImage=图片
pName=名称

[func (t *Tree) GetCount() int32 {]
ff=取项数量

[func (t *Tree) GetCountColumn() int32 {]
ff=取列数量

[func (t *Tree) SetItemText(nID, iColumn int32, pValue string) bool {]
ff=置项文本
pValue=值
iColumn=列
nID=项ID

[func (t *Tree) SetItemTextEx(nID int32, pName string, pValue string) bool {]
ff=置项文本EX
pValue=值
pName=名称
nID=项ID

[func (t *Tree) SetItemImage(nID, iColumn int32, hImage int) bool {]
ff=置项图片
hImage=图片
iColumn=列
nID=项ID

[func (t *Tree) SetItemImageEx(nID int32, pName string, hImage int) bool {]
ff=置项图片EX
hImage=图片
pName=名称
nID=项ID

[func (t *Tree) GetItemText(nID, iColumn int32) string {]
ff=取项文本
iColumn=列
nID=项ID

[func (t *Tree) GetItemTextEx(nID int32, pName string) string {]
ff=取项文本EX
pName=名称
nID=项ID

[func (t *Tree) GetItemImage(nID, iColumn int32) int {]
ff=取项图片
iColumn=列
nID=项ID

[func (t *Tree) GetItemImageEx(nID int32, pName string) int {]
ff=取项图片EX
pName=名称
nID=项ID

[func (t *Tree) DeleteItem(nID int32) bool {]
ff=删除项

[func (t *Tree) DeleteItemAll() {]
ff=删除全部项

[func (t *Tree) DeleteColumnAll() {]
ff=删除列全部

[func (t *Tree) SetSplitLineColor(color int) {]
ff=置分割线颜色
color=ABGR颜色值

[func (t *Tree) SetItemTemplateXMLFromMem(data #左中括号##右中括号#byte) bool {]
ff=置项模板从内存
data=模板数据

[func (t *Tree) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule uintptr) bool {]
ff=置项模板并按资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=文件名
id=RC资源ID

[func (t *Tree) GetItemTemplate() int {]
ff=取项模板

[func (t *Tree) Event_TREE_TEMP_CREATE(pFun XE_TREE_TEMP_CREATE) bool {]
ff=事件_项模板创建

[func (t *Tree) Event_TREE_TEMP_CREATE1(pFun XE_TREE_TEMP_CREATE1) bool {]
ff=事件_项模板创建1

[func (t *Tree) Event_TREE_TEMP_CREATE_END(pFun XE_TREE_TEMP_CREATE_END) bool {]
ff=事件_项模板创建完成

[func (t *Tree) Event_TREE_TEMP_CREATE_END1(pFun XE_TREE_TEMP_CREATE_END1) bool {]
ff=事件_项模板创建完成1

[func (t *Tree) Event_TREE_TEMP_DESTROY(pFun XE_TREE_TEMP_DESTROY) bool {]
ff=事件_项模板销毁

[func (t *Tree) Event_TREE_TEMP_DESTROY1(pFun XE_TREE_TEMP_DESTROY1) bool {]
ff=事件_项模板销毁1

[func (t *Tree) Event_TREE_TEMP_ADJUST_COORDINATE(pFun XE_TREE_TEMP_ADJUST_COORDINATE) bool {]
ff=停用_项模板

[func (t *Tree) Event_TREE_TEMP_ADJUST_COORDINATE1(pFun XE_TREE_TEMP_ADJUST_COORDINATE1) bool {]
ff=停用_项模板1

[func (t *Tree) Event_TREE_DRAWITEM(pFun XE_TREE_DRAWITEM) bool {]
ff=事件_绘制项

[func (t *Tree) Event_TREE_DRAWITEM1(pFun XE_TREE_DRAWITEM1) bool {]
ff=事件_绘制项1

[func (t *Tree) Event_TREE_SELECT(pFun XE_TREE_SELECT) bool {]
ff=事件_项选择

[func (t *Tree) Event_TREE_SELECT1(pFun XE_TREE_SELECT1) bool {]
ff=事件_项选择1

[func (t *Tree) Event_TREE_EXPAND(pFun XE_TREE_EXPAND) bool {]
ff=事件_项展开收缩

[func (t *Tree) Event_TREE_EXPAND1(pFun XE_TREE_EXPAND1) bool {]
ff=事件_项展开收缩1

[func (t *Tree) Event_TREE_DRAG_ITEM_ING(pFun XE_TREE_DRAG_ITEM_ING) bool {]
ff=事件_用户正在拖动项

[func (t *Tree) Event_TREE_DRAG_ITEM_ING1(pFun XE_TREE_DRAG_ITEM_ING1) bool {]
ff=事件_用户正在拖动项1

[func (t *Tree) Event_TREE_DRAG_ITEM(pFun XE_TREE_DRAG_ITEM) bool {]
ff=事件_拖动项

[func (t *Tree) Event_TREE_DRAG_ITEM1(pFun XE_TREE_DRAG_ITEM1) bool {]
ff=事件_拖动项1
