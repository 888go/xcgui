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

[func XListView_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=列表视_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XListView_CreateEx(x, y, cx, cy int32, hParent, col_extend_count int32) int {]
ff=列表视_创建Ex
col_extend_count=列数量
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XListView_CreateAdapter(hEle int) int {]
ff=列表视_创建数据适配器
hEle=元素句柄

[func XListView_BindAdapter(hEle int, hAdapter int) int {]
ff=列表视_绑定数据适配器
hAdapter=数据适配器
hEle=元素句柄

[func XListView_GetAdapter(hEle int) int {]
ff=列表视_取数据适配器
hEle=元素句柄

[func XListView_SetItemTemplateXML(hEle int, pXmlFile string) bool {]
ff=列表视_置项模板文件
pXmlFile=文件名
hEle=元素句柄

[func XListView_SetItemTemplateXMLFromString(hEle int, pStringXML string) bool {]
ff=列表视_置项模板从字符串
pStringXML=字符串
hEle=元素句柄

[func XListView_SetItemTemplate(hEle int, hTemp int) bool {]
ff=列表视_置项模板
hTemp=模板句柄
hEle=元素句柄

[func XListView_GetTemplateObject(hEle int, iGroup int, iItem int, nTempItemID int) int {]
ff=列表视_取模板对象
nTempItemID=模板项ID
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_GetTemplateObjectGroup(hEle int, iGroup int, nTempItemID int) int {]
ff=列表视_取模板对象组
nTempItemID=模板项ID
iGroup=组索引
hEle=元素句柄

[func XListView_GetItemIDFromHXCGUI(hEle int, hXCGUI int, piGroup *int32, piItem *int32) bool {]
ff=列表视_取对象所在项
piItem=接收项索引
piGroup=接收组索引
hXCGUI=对象句柄
hEle=元素句柄

[func XListView_HitTest(hEle int, pPt *POINT, pOutGroup *int32, pOutItem *int32) bool {]
ff=列表视_测试点击项
pOutItem=接收项索引
pOutGroup=接收组索引
pPt=坐标点
hEle=元素句柄

[func XListView_HitTestOffset(hEle int, pPt *POINT, pOutGroup *int32, pOutItem *int32) bool {]
ff=列表视_测试点击项EX
pOutItem=接收项索引
pOutGroup=接收做索引
pPt=坐标点
hEle=元素句柄

[func XListView_EnableMultiSel(hEle int, bEnable bool) int {]
ff=列表视_启用多选
bEnable=是否启用
hEle=元素句柄

[func XListView_EnableTemplateReuse(hEle int, bEnable bool) int {]
ff=列表视_启用模板复用
bEnable=是否启用
hEle=元素句柄

[func XListView_EnableVirtualTable(hEle int, bEnable bool) int {]
ff=列表视_启用虚表
bEnable=是否启用
hEle=元素句柄

[func XListView_SetVirtualItemCount(hEle int, iGroup int, nCount int) bool {]
ff=列表视_置虚表项数量
nCount=项数量
iGroup=组索引
hEle=元素句柄

[func XListView_SetDrawItemBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) int {]
ff=列表视_置项背景绘制标志
nFlags=标志位
hEle=元素句柄

[func XListView_SetSelectItem(hEle int, iGroup int, iItem int) bool {]
ff=列表视_置选择项
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_GetSelectItem(hEle int, piGroup *int32, piItem *int32) bool {]
ff=列表视_取选择项
piItem=接收项索引
piGroup=接收组索引
hEle=元素句柄

[func XListView_AddSelectItem(hEle int, iGroup int, iItem int) bool {]
ff=列表视_添加选择项
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_VisibleItem(hEle int, iGroup int, iItem int) int {]
ff=列表视_显示指定项
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_GetVisibleItemRange(hEle int, piGroup1 *int32, piGroup2 *int32, piStartGroup *int32, piStartItem *int32, piEndGroup *int32, piEndItem *int32) int {]
ff=列表视_取可视项范围
piEndItem=可视结束项
piEndGroup=可视结束组
piStartItem=可视开始项
piStartGroup=可视开始组
piGroup2=组2
piGroup1=组1
hEle=元素句柄

[func XListView_GetSelectItemCount(hEle int) int {]
ff=列表视_取选择项数量
hEle=元素句柄

[func XListView_GetSelectAll(hEle int, pArray *#左中括号##右中括号#ListView_Item_Id_, nArraySize int) int {]
ff=列表视_取选择项全部
nArraySize=切片大小
pArray=切片
hEle=元素句柄

[func XListView_SetSelectAll(hEle int) int {]
ff=列表视_置选择项全部
hEle=元素句柄

[func XListView_CancelSelectAll(hEle int) int {]
ff=列表视_取消选择项全部
hEle=元素句柄

[func XListView_SetColumnSpace(hEle int, space int) int {]
ff=列表视_置列间隔
space=间隔大小
hEle=元素句柄

[func XListView_SetRowSpace(hEle int, space int) int {]
ff=列表视_置行间隔
space=间隔大小
hEle=元素句柄

[func XListView_SetItemSize(hEle int, width int, height int) int {]
ff=列表视_置项大小
height=高度
width=宽度
hEle=元素句柄

[func XListView_GetItemSize(hEle int, pSize *SIZE) int {]
ff=列表视_取项大小
pSize=接收返回大小
hEle=元素句柄

[func XListView_SetGroupHeight(hEle int, height int) int {]
ff=列表视_置组高度
height=高度
hEle=元素句柄

[func XListView_GetGroupHeight(hEle int) int {]
ff=列表视_取组高度
hEle=元素句柄

[func XListView_SetGroupUserData(hEle int, iGroup int, nData int) int {]
ff=列表视_置组用户数据
nData=数据
iGroup=组索引
hEle=元素句柄

[func XListView_SetItemUserData(hEle int, iGroup int, iItem int, nData int) int {]
ff=列表视_置项用户数据
nData=数据
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_GetGroupUserData(hEle int, iGroup int) int {]
ff=列表视_取组用户数据
iGroup=组索引
hEle=元素句柄

[func XListView_GetItemUserData(hEle int, iGroup int, iItem int) int {]
ff=列表视_取项用户数据
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_RefreshData(hEle int) int {]
ff=列表视_刷新项数据
hEle=元素句柄

[func XListView_RefreshItem(hEle int, iGroup int, iItem int) int {]
ff=列表视_刷新指定项
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_ExpandGroup(hEle int, iGroup int, bExpand bool) bool {]
ff=列表视_展开组
bExpand=是否展开
iGroup=组索引
hEle=元素句柄

[func XListView_Group_AddColumn(hEle int, pName string) int {]
ff=列表视_组添加列
pName=字段称
hEle=元素句柄

[func XListView_Group_AddItemText(hEle int, pValue string, iPos int) int {]
ff=列表视_组添加项文本
iPos=插入位置
pValue=值
hEle=元素句柄

[func XListView_Group_AddItemTextEx(hEle int, pName string, pValue string, iPos int) int {]
ff=列表视_组添加项文本EX
iPos=插入位置
pValue=值
pName=字段称
hEle=元素句柄

[func XListView_Group_AddItemImage(hEle int, hImage int, iPos int) int {]
ff=列表视_组添加项图片
iPos=插入位置
hImage=图片句柄
hEle=元素句柄

[func XListView_Group_AddItemImageEx(hEle int, pName string, hImage int, iPos int) int {]
ff=列表视_组添加项图片EX
iPos=插入位置
hImage=图片句柄
pName=字段称
hEle=元素句柄

[func XListView_Group_SetText(hEle int, iGroup int, iColumn int, pValue string) bool {]
ff=列表视_组置文本
pValue=值
iColumn=列索引
iGroup=组索引
hEle=元素句柄

[func XListView_Group_SetTextEx(hEle int, iGroup int, pName string, pValue string) bool {]
ff=列表视_组置文本EX
pValue=值
pName=字段名
iGroup=组索引
hEle=元素句柄

[func XListView_Group_SetImage(hEle int, iGroup int, iColumn int, hImage int) bool {]
ff=列表视_组置图片
hImage=图片句柄
iColumn=列索引
iGroup=组索引
hEle=元素句柄

[func XListView_Group_SetImageEx(hEle int, iGroup int, pName string, hImage int) bool {]
ff=列表视_组置图片EX
hImage=图片句柄
pName=字段名
iGroup=组索引
hEle=元素句柄

[func XListView_Group_GetCount(hEle int) int {]
ff=列表视_组获取数量
hEle=元素句柄

[func XListView_Item_GetCount(hEle int, iGroup int) int {]
ff=列表视_项获取数量
iGroup=组索引
hEle=元素句柄

[func XListView_Item_AddColumn(hEle int, pName string) int {]
ff=列表视_项添加列
pName=字段名
hEle=元素句柄

[func XListView_Item_AddItemText(hEle int, iGroup int, pValue string, iPos int) int {]
ff=列表视_项添加文本
iPos=插入位置
pValue=值
iGroup=组索引
hEle=元素句柄

[func XListView_Item_AddItemTextEx(hEle int, iGroup int, pName string, pValue string, iPos int) int {]
ff=列表视_项添加文本EX
iPos=插入位置
pValue=值
pName=字段名
iGroup=组索引
hEle=元素句柄

[func XListView_Item_AddItemImage(hEle int, iGroup int, hImage int, iPos int) int {]
ff=列表视_项添加图片
iPos=插入位置
hImage=图片句柄
iGroup=组索引
hEle=元素句柄

[func XListView_Item_AddItemImageEx(hEle int, iGroup int, pName string, hImage int, iPos int) int {]
ff=列表视_项添加图片EX
iPos=插入位置
hImage=图片句柄
pName=字段名
iGroup=组索引
hEle=元素句柄

[func XListView_Item_SetText(hEle int, iGroup int, iItem int, iColumn int, pValue string) bool {]
ff=列表视_项置文本
pValue=值
iColumn=列索引
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_Item_SetTextEx(hEle int, iGroup int, iItem int, pName string, pValue string) bool {]
ff=列表视_项置文本EX
pValue=值
pName=字段名
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_Item_SetImage(hEle int, iGroup int, iItem int, iColumn int, hImage int) bool {]
ff=列表视_项置图片
hImage=图片句柄
iColumn=列索引
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_Item_SetImageEx(hEle int, iGroup int, iItem int, pName string, hImage int) bool {]
ff=列表视_项置图片EX
hImage=图片句柄
pName=列名称
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_Group_DeleteItem(hEle int, iGroup int) bool {]
ff=列表视_组删除项
iGroup=组索引
hEle=元素句柄

[func XListView_Group_DeleteAllChildItem(hEle int, iGroup int) int {]
ff=列表视_组删除全部子项
iGroup=组索引
hEle=元素句柄

[func XListView_Item_DeleteItem(hEle int, iGroup int, iItem int) bool {]
ff=列表视_项删除
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_DeleteAll(hEle int) int {]
ff=列表视_删除全部
hEle=元素句柄

[func XListView_DeleteAllGroup(hEle int) int {]
ff=列表视_删除全部组
hEle=元素句柄

[func XListView_DeleteAllItem(hEle int) int {]
ff=列表视_删除全部项
hEle=元素句柄

[func XListView_DeleteColumnGroup(hEle int, iColumn int) int {]
ff=列表视_组删除列
iColumn=列索引
hEle=元素句柄

[func XListView_DeleteColumnItem(hEle int, iColumn int) int {]
ff=列表视_项删除列
iColumn=列索引
hEle=元素句柄

[func XListView_Item_GetTextEx(hEle int, iGroup int, iItem int, pName string) string {]
ff=列表视_项获取文本
pName=字段称
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_Item_GetImageEx(hEle int, iGroup int, iItem int, pName string) int {]
ff=列表视_项获取图片EX
pName=字段称
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_Group_GetText(hEle int, iGroup int, iColumn int) string {]
ff=列表视_组取文本
iColumn=列索引
iGroup=组索引
hEle=元素句柄

[func XListView_Group_GetTextEx(hEle int, iGroup int, pName string) string {]
ff=列表视_组取文本EX
pName=字段名称
iGroup=组索引
hEle=元素句柄

[func XListView_Group_GetImage(hEle int, iGroup int, iColumn int) int {]
ff=列表视_组取图片
iColumn=列索引
iGroup=组索引
hEle=元素句柄

[func XListView_Group_GetImageEx(hEle int, iGroup int, pName string) int {]
ff=列表视_组取图片EX
pName=字段名称
iGroup=组索引
hEle=元素句柄

[func XListView_Item_GetText(hEle int, iGroup int, iItem int, iColumn int) string {]
ff=列表视_项取文本
iColumn=列索引
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_Item_GetImage(hEle int, iGroup int, iItem int, iColumn int) int {]
ff=列表视_项取图片
iColumn=列索引
iItem=项索引
iGroup=组索引
hEle=元素句柄

[func XListView_SetDragRectColor(hEle int, color int, width int) int {]
ff=列表视_置拖动矩形颜色
width=线宽度
color=ABGR颜色
hEle=元素句柄

[func XListView_SetItemTemplateXMLFromMem(hEle int, data #左中括号##右中括号#byte) bool {]
ff=列表视_置项模板从内存
data=模板数据
hEle=元素句柄

[func XListView_SetItemTemplateXMLFromZipRes(hEle int, id int32, pFileName string, pPassword string, hModule uintptr) bool {]
ff=列表视_置项模板从资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=文件名
id=RC资源ID
hEle=元素句柄

[func XListView_GetItemTemplate(hEle int) int {]
ff=列表视_取项模板
hEle=元素句柄

[func XListView_GetItemTemplateGroup(hEle int) int {]
ff=列表视_取项模板组
hEle=元素句柄
