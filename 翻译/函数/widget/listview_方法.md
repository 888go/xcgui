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

[func NewListView(x int, y int, cx int, cy int, hParent int) *ListView {]
ff=创建列表视
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewListViewEx(x, y, cx, cy int32, hParent, col_extend_count int32) *ListView {]
ff=创建列表视Ex
col_extend_count=列数量
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewListViewByHandle(handle int) *ListView {]
ff=创建列表视并按句柄

[func NewListViewByName(name string) *ListView {]
ff=创建列表视并按名称

[func NewListViewByUID(nUID int) *ListView {]
ff=创建列表视并按UID

[func NewListViewByUIDName(name string) *ListView {]
ff=创建列表视并按UID名称

[func (l *ListView) CreateAdapter() int {]
ff=创建数据适配器

[func (l *ListView) BindAdapter(hAdapter int) int {]
ff=绑定数据适配器
hAdapter=数据适配器

[func (l *ListView) GetAdapter() int {]
ff=取数据适配器

[func (l *ListView) SetItemTemplateXML(pXmlFile string) bool {]
ff=置项模板文件
pXmlFile=文件名

[func (l *ListView) SetItemTemplateXMLFromString(pStringXML string) bool {]
ff=置项模板从字符串
pStringXML=字符串

[func (l *ListView) SetItemTemplate(hTemp int) bool {]
ff=置项模板
hTemp=模板句柄

[func (l *ListView) GetTemplateObject(iGroup int, iItem int, nTempItemID int) int {]
ff=取模板对象
nTempItemID=模板项ID
iItem=项索引
iGroup=组索引

[func (l *ListView) GetTemplateObjectGroup(iGroup int, nTempItemID int) int {]
ff=取模板对象组
nTempItemID=模板项ID
iGroup=组索引

[func (l *ListView) GetItemIDFromHXCGUI(hXCGUI int, piGroup *int32, piItem *int32) bool {]
ff=取对象所在项
piItem=接收项索引
piGroup=接收组索引
hXCGUI=对象句柄

[func (l *ListView) HitTest(pPt *xc.POINT, pOutGroup *int32, pOutItem *int32) bool {]
ff=测试点击项
pPt=坐标点

[func (l *ListView) HitTestOffset(pPt *xc.POINT, pOutGroup *int32, pOutItem *int32) bool {]
ff=测试点击项EX
pPt=坐标点

[func (l *ListView) EnableMultiSel(bEnable bool) int {]
ff=启用多选
bEnable=是否启用

[func (l *ListView) EnableTemplateReuse(bEnable bool) int {]
ff=启用模板复用
bEnable=是否启用

[func (l *ListView) EnableVirtualTable(bEnable bool) int {]
ff=启用虚表
bEnable=是否启用

[func (l *ListView) SetVirtualItemCount(iGroup int, nCount int) bool {]
ff=置虚表项数量
nCount=项数量
iGroup=组索引

[func (l *ListView) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) int {]
ff=置项背景绘制标志
nFlags=标志位

[func (l *ListView) SetSelectItem(iGroup int, iItem int) bool {]
ff=置选择项
iItem=项索引
iGroup=组索引

[func (l *ListView) GetSelectItem(piGroup *int32, piItem *int32) bool {]
ff=取选择项
piItem=接收项索引
piGroup=接收组索引

[func (l *ListView) AddSelectItem(iGroup int, iItem int) bool {]
ff=添加选择项
iItem=项索引
iGroup=组索引

[func (l *ListView) VisibleItem(iGroup int, iItem int) int {]
ff=显示指定项
iItem=项索引
iGroup=组索引

[func (l *ListView) GetVisibleItemRange(piGroup1 *int32, piGroup2 *int32, piStartGroup *int32, piStartItem *int32, piEndGroup *int32, piEndItem *int32) int {]
ff=取可视项范围
piEndItem=可视结束项
piEndGroup=可视结束组
piStartItem=可视开始项
piStartGroup=可视开始组
piGroup2=组2
piGroup1=组1

[func (l *ListView) GetSelectItemCount() int {]
ff=取选择项数量

[func (l *ListView) GetSelectAll(pArray *#左中括号##右中括号#xc.ListView_Item_Id_, nArraySize int) int {]
ff=取选择项全部
pArray=切片

[func (l *ListView) SetSelectAll() int {]
ff=置选择项全部

[func (l *ListView) CancelSelectAll() int {]
ff=取消选择项全部

[func (l *ListView) SetColumnSpace(space int) int {]
ff=置列间隔
space=间隔大小

[func (l *ListView) SetRowSpace(space int) int {]
ff=置行间隔
space=间隔大小

[func (l *ListView) SetItemSize(width int, height int) int {]
ff=置项大小
height=高度
width=宽度

[func (l *ListView) GetItemSize(pSize *xc.SIZE) int {]
ff=取项大小
pSize=接收返回大小

[func (l *ListView) SetGroupHeight(height int) int {]
ff=置组高度
height=高度

[func (l *ListView) GetGroupHeight() int {]
ff=取组高度

[func (l *ListView) SetGroupUserData(iGroup int, nData int) int {]
ff=置组用户数据
nData=数据
iGroup=组索引

[func (l *ListView) SetItemUserData(iGroup int, iItem int, nData int) int {]
ff=置项用户数据
nData=数据
iItem=项索引
iGroup=组索引

[func (l *ListView) GetGroupUserData(iGroup int) int {]
ff=取组用户数据
iGroup=组索引

[func (l *ListView) GetItemUserData(iGroup int, iItem int) int {]
ff=取项用户数据
iItem=项索引
iGroup=组索引

[func (l *ListView) RefreshData() int {]
ff=刷新项数据

[func (l *ListView) RefreshItem(iGroup int, iItem int) int {]
ff=刷新指定项
iItem=项索引
iGroup=组索引

[func (l *ListView) ExpandGroup(iGroup int, bExpand bool) bool {]
ff=展开组
bExpand=是否展开
iGroup=组索引

[func (l *ListView) Group_AddColumn(pName string) int {]
ff=组添加列
pName=字段称

[func (l *ListView) Group_AddItemText(pValue string, iPos int) int {]
ff=组添加项文本
iPos=插入位置
pValue=值

[func (l *ListView) Group_AddItemTextEx(pName string, pValue string, iPos int) int {]
ff=组添加项文本EX
iPos=插入位置
pValue=值
pName=字段称

[func (l *ListView) Group_AddItemImage(hImage int, iPos int) int {]
ff=组添加项图片
iPos=插入位置
hImage=图片句柄

[func (l *ListView) Group_AddItemImageEx(pName string, hImage int, iPos int) int {]
ff=组添加项图片EX
iPos=插入位置
hImage=图片句柄
pName=字段称

[func (l *ListView) Group_SetText(iGroup int, iColumn int, pValue string) bool {]
ff=组置文本
pValue=值
iColumn=列索引
iGroup=组索引

[func (l *ListView) Group_SetTextEx(iGroup int, pName string, pValue string) bool {]
ff=组置文本EX
pValue=值
pName=字段名
iGroup=组索引

[func (l *ListView) Group_SetImage(iGroup int, iColumn int, hImage int) bool {]
ff=组置图片
hImage=图片句柄
iColumn=列索引
iGroup=组索引

[func (l *ListView) Group_SetImageEx(iGroup int, pName string, hImage int) bool {]
ff=组置图片EX
hImage=图片句柄
pName=字段名
iGroup=组索引

[func (l *ListView) Group_GetCount() int {]
ff=组获取数量

[func (l *ListView) Item_GetCount(iGroup int) int {]
ff=项获取数量
iGroup=组索引

[func (l *ListView) Item_AddColumn(pName string) int {]
ff=项添加列
pName=字段名

[func (l *ListView) Item_AddItemText(iGroup int, pValue string, iPos int) int {]
ff=项添加文本
iPos=插入位置
pValue=值
iGroup=组索引

[func (l *ListView) Item_AddItemTextEx(iGroup int, pName string, pValue string, iPos int) int {]
ff=项添加文本EX
iPos=插入位置
pValue=值
pName=字段名
iGroup=组索引

[func (l *ListView) Item_AddItemImage(iGroup int, hImage int, iPos int) int {]
ff=项添加图片
iPos=插入位置
hImage=图片句柄
iGroup=组索引

[func (l *ListView) Item_AddItemImageEx(iGroup int, pName string, hImage int, iPos int) int {]
ff=项添加图片EX
iPos=插入位置
hImage=图片句柄
pName=字段名
iGroup=组索引

[func (l *ListView) Item_SetText(iGroup int, iItem int, iColumn int, pValue string) bool {]
ff=项置文本
pValue=值
iColumn=列索引
iItem=项索引
iGroup=组索引

[func (l *ListView) Item_SetTextEx(iGroup int, iItem int, pName string, pValue string) bool {]
ff=项置文本EX
pValue=值
pName=字段名
iItem=项索引
iGroup=组索引

[func (l *ListView) Item_SetImage(iGroup int, iItem int, iColumn int, hImage int) bool {]
ff=项置图片
hImage=图片句柄
iColumn=列索引
iItem=项索引
iGroup=组索引

[func (l *ListView) Item_SetImageEx(iGroup int, iItem int, pName string, hImage int) bool {]
ff=项置图片EX
hImage=图片句柄
pName=列名称
iItem=项索引
iGroup=组索引

[func (l *ListView) Group_DeleteItem(iGroup int) bool {]
ff=组删除项
iGroup=组索引

[func (l *ListView) Group_DeleteAllChildItem(iGroup int) int {]
ff=组删除全部子项
iGroup=组索引

[func (l *ListView) Item_DeleteItem(iGroup int, iItem int) bool {]
ff=项删除
iItem=项索引
iGroup=组索引

[func (l *ListView) DeleteAll() int {]
ff=删除全部

[func (l *ListView) DeleteAllGroup() int {]
ff=删除全部组

[func (l *ListView) DeleteAllItem() int {]
ff=删除全部项

[func (l *ListView) DeleteColumnGroup(iColumn int) int {]
ff=组删除列
iColumn=列索引

[func (l *ListView) DeleteColumnItem(iColumn int) int {]
ff=项删除列
iColumn=列索引

[func (l *ListView) Item_GetTextEx(iGroup int, iItem int, pName string) string {]
ff=项获取文本
pName=字段称
iItem=项索引
iGroup=组索引

[func (l *ListView) Item_GetImageEx(iGroup int, iItem int, pName string) int {]
ff=项获取图片EX
pName=字段称
iItem=项索引
iGroup=组索引

[func (l *ListView) Group_GetText(iGroup int, iColumn int) string {]
ff=组取文本
iColumn=列索引
iGroup=组索引

[func (l *ListView) Group_GetTextEx(iGroup int, pName string) string {]
ff=组取文本EX
pName=字段名称
iGroup=组索引

[func (l *ListView) Group_GetImage(iGroup int, iColumn int) int {]
ff=组取图片
iColumn=列索引
iGroup=组索引

[func (l *ListView) Group_GetImageEx(iGroup int, pName string) int {]
ff=组取图片EX
pName=字段名称
iGroup=组索引

[func (l *ListView) Item_GetText(iGroup int, iItem int, iColumn int) string {]
ff=项取文本
iColumn=列索引
iItem=项索引
iGroup=组索引

[func (l *ListView) Item_GetImage(iGroup int, iItem int, iColumn int) int {]
ff=项取图片
iColumn=列索引
iItem=项索引
iGroup=组索引

[func (l *ListView) SetDragRectColor(color int, width int) int {]
ff=置拖动矩形颜色
width=线宽度
color=ABGR颜色

[func (l *ListView) SetItemTemplateXMLFromMem(data #左中括号##右中括号#byte) bool {]
ff=置项模板并按内存
data=模板数据

[func (l *ListView) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule uintptr) bool {]
ff=置项模板并按资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=文件名
id=RC资源ID

[func (l *ListView) GetItemTemplate() int {]
ff=取项模板

[func (l *ListView) GetItemTemplateGroup() int {]
ff=取项模板组

[func (l *ListView) Event_LISTVIEW_TEMP_CREATE(pFun XE_LISTVIEW_TEMP_CREATE) bool {]
ff=事件_项模板创建

[func (l *ListView) Event_LISTVIEW_TEMP_CREATE1(pFun XE_LISTVIEW_TEMP_CREATE1) bool {]
ff=事件_项模板创建1

[func (l *ListView) Event_LISTVIEW_TEMP_CREATE_END(pFun XE_LISTVIEW_TEMP_CREATE_END) bool {]
ff=事件_项模板创建完成

[func (l *ListView) Event_LISTVIEW_TEMP_CREATE_END1(pFun XE_LISTVIEW_TEMP_CREATE_END1) bool {]
ff=事件_项模板创建完成1

[func (l *ListView) Event_LISTVIEW_TEMP_DESTROY(pFun XE_LISTVIEW_TEMP_DESTROY) bool {]
ff=事件_项模板销毁

[func (l *ListView) Event_LISTVIEW_TEMP_DESTROY1(pFun XE_LISTVIEW_TEMP_DESTROY1) bool {]
ff=事件_项模板销毁1

[func (l *ListView) Event_LISTVIEW_TEMP_ADJUST_COORDINATE(pFun XE_LISTVIEW_TEMP_ADJUST_COORDINATE) bool {]
ff=事件_项模板调整坐标

[func (l *ListView) Event_LISTVIEW_TEMP_ADJUST_COORDINATE1(pFun XE_LISTVIEW_TEMP_ADJUST_COORDINATE1) bool {]
ff=事件_项模板调整坐标1

[func (l *ListView) Event_LISTVIEW_DRAWITEM(pFun XE_LISTVIEW_DRAWITEM) bool {]
ff=事件_自绘项

[func (l *ListView) Event_LISTVIEW_DRAWITEM1(pFun XE_LISTVIEW_DRAWITEM1) bool {]
ff=事件_自绘项1

[func (l *ListView) Event_LISTVIEW_SELECT(pFun XE_LISTVIEW_SELECT) bool {]
ff=事件_项选择

[func (l *ListView) Event_LISTVIEW_SELECT1(pFun XE_LISTVIEW_SELECT1) bool {]
ff=事件_项选择1

[func (l *ListView) Event_LISTVIEW_EXPAND(pFun XE_LISTVIEW_EXPAND) bool {]
ff=事件_组展开收缩
[func (l *ListView) Event_LISTVIEW_EXPAND1(pFun XE_LISTVIEW_EXPAND1) bool {]
ff=事件_组展开收缩1
