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

[func XTree_Create(x, y, cx, cy int32, hParent int) int {]
ff=列表树_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XTree_CreateEx(x, y, cx, cy int32, hParent, col_extend_count int32) int {]
ff=列表树_创建Ex
col_extend_count=列数量
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XTree_EnableDragItem(hEle int, bEnable bool) {]
ff=列表树_启用拖动项
bEnable=是否启用
hEle=元素句柄

[func XTree_EnableConnectLine(hEle int, bEnable bool, bSolid bool) {]
ff=列表树_启用连接线
bSolid=实线或虚线
bEnable=是否启用
hEle=元素句柄

[func XTree_EnableExpand(hEle int, bEnable bool) {]
ff=列表树_启用展开
bEnable=是否启用
hEle=元素句柄

[func XTree_EnableTemplateReuse(hEle int, bEnable bool) {]
ff=列表树_启用模板复用
bEnable=是否启用
hEle=元素句柄

[func XTree_SetConnectLineColor(hEle int, color int) {]
ff=列表树_置连接线颜色
color=ABGR颜色
hEle=元素句柄

[func XTree_SetExpandButtonSize(hEle int, nWidth, nHeight int32) {]
ff=列表树_置展开按钮大小
nHeight=高度
nWidth=宽度
hEle=元素句柄

[func XTree_SetConnectLineLength(hEle int, nLength int32) {]
ff=列表树_置连接线长度
nLength=连线绘制长度
hEle=元素句柄

[func XTree_SetDragInsertPositionColor(hEle int, color int) {]
ff=列表树_置拖动项插入位置颜色
color=ABGR颜色
hEle=元素句柄

[func XTree_SetItemTemplateXML(hEle int, pXmlFile string) bool {]
ff=列表树_置项模板文件
pXmlFile=文件名
hEle=元素句柄

[func XTree_SetItemTemplateXMLSel(hEle int, pXmlFile string) bool {]
ff=列表树_置选择项模板文件
pXmlFile=文件名
hEle=元素句柄

[func XTree_SetItemTemplate(hEle int, hTemp int) bool {]
ff=列表树_置项模板
hTemp=模板句柄
hEle=元素句柄

[func XTree_SetItemTemplateSel(hEle int, hTemp int) bool {]
ff=列表树_置选择项模板
hTemp=模板句柄
hEle=元素句柄

[func XTree_SetItemTemplateXMLFromString(hEle int, pStringXML string) bool {]
ff=列表树_置项模板从字符串
pStringXML=字符串
hEle=元素句柄

[func XTree_SetItemTemplateXMLSelFromString(hEle int, pStringXML string) bool {]
ff=列表树_置选择项模板从字符串
pStringXML=字符串
hEle=元素句柄

[func XTree_SetDrawItemBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) {]
ff=列表树_置项背景绘制标志
nFlags=标志位
hEle=元素句柄

[func XTree_SetItemData(hEle int, nID int32, nUserData int) bool {]
ff=列表树_置项数据
nUserData=用户数据
nID=项ID
hEle=元素句柄

[func XTree_GetItemData(hEle int, nID int32) int {]
ff=列表树_取项数据
nID=项ID
hEle=元素句柄

[func XTree_SetSelectItem(hEle int, nID int32) bool {]
ff=列表树_置选择项
nID=项ID
hEle=元素句柄

[func XTree_GetSelectItem(hEle int) int32 {]
ff=列表树_取选择项
hEle=元素句柄

[func XTree_VisibleItem(hEle int, nID int32) {]
ff=列表树_可视指定项
nID=项ID
hEle=元素句柄

[func XTree_IsExpand(hEle int, nID int32) bool {]
ff=列表树_判断展开
nID=项ID
hEle=元素句柄

[func XTree_ExpandItem(hEle int, nID int32, bExpand bool) bool {]
ff=列表树_展开项
bExpand=是否展开
nID=项ID
hEle=元素句柄

[func XTree_ExpandAllChildItem(hEle int, nID int32, bExpand bool) bool {]
ff=列表树_展开全部子项
bExpand=是否展开
nID=项ID
hEle=元素句柄

[func XTree_HitTest(hEle int, pPt *POINT) int32 {]
ff=列表树_测试点击项
pPt=坐标点
hEle=元素句柄

[func XTree_HitTestOffset(hEle int, pPt *POINT) int32 {]
ff=列表树_测试点击项EX
pPt=坐标点
hEle=元素句柄

[func XTree_GetFirstChildItem(hEle int, nID int32) int32 {]
ff=列表树_取第一个子项
nID=项ID
hEle=元素句柄

[func XTree_GetEndChildItem(hEle int, nID int32) int32 {]
ff=列表树_取末尾子项
nID=项ID
hEle=元素句柄

[func XTree_GetPrevSiblingItem(hEle int, nID int32) int32 {]
ff=列表树_取上一个兄弟项
nID=项ID
hEle=元素句柄

[func XTree_GetNextSiblingItem(hEle int, nID int32) int32 {]
ff=列表树_取下一个兄弟项
nID=项ID
hEle=元素句柄

[func XTree_GetParentItem(hEle int, nID int32) int32 {]
ff=列表树_取父项
nID=项ID
hEle=元素句柄

[func XTree_CreateAdapter(hEle int) int {]
ff=列表树_创建数据适配器
hEle=元素句柄

[func XTree_BindAdapter(hEle int, hAdapter int) {]
ff=列表树_绑定数据适配器
hAdapter=数据适配器句柄
hEle=元素句柄

[func XTree_GetAdapter(hEle int) int {]
ff=列表树_取数据适配器
hEle=元素句柄

[func XTree_RefreshData(hEle int) {]
ff=列表树_刷新数据
hEle=元素句柄

[func XTree_RefreshItem(hEle int, nID int32) {]
ff=列表树_刷新指定项
nID=项ID
hEle=元素句柄

[func XTree_SetIndentation(hEle int, nWidth int32) {]
ff=列表树_置缩进
nWidth=缩进宽度
hEle=元素句柄

[func XTree_GetIndentation(hEle int) int32 {]
ff=列表树_取缩进
hEle=元素句柄

[func XTree_SetItemHeightDefault(hEle int, nHeight, nSelHeight int32) {]
ff=列表树_置项默认高度
nSelHeight=选中时高度
nHeight=高度
hEle=元素句柄

[func XTree_GetItemHeightDefault(hEle int, pHeight, pSelHeight *int32) {]
ff=列表树_取项默认高度
pSelHeight=接收返回值
pHeight=接收返回高度
hEle=元素句柄

[func XTree_SetItemHeight(hEle int, nID, nHeight, nSelHeight int32) {]
ff=列表树_置项高度
nSelHeight=选中时高度
nHeight=高度
nID=项ID
hEle=元素句柄

[func XTree_GetItemHeight(hEle int, nID int32, pHeight, pSelHeight *int32) {]
ff=列表树_取项高度
pSelHeight=接收返回值
pHeight=接收返回高度
nID=项ID
hEle=元素句柄

[func XTree_SetRowSpace(hEle int, nSpace int32) {]
ff=列表树_置行间距
nSpace=行间隔大小
hEle=元素句柄

[func XTree_GetRowSpace(hEle int) int32 {]
ff=列表树_取行间距
hEle=元素句柄

[func XTree_MoveItem(hEle int, nMoveItem, nDestItem, nFlag int32) bool {]
ff=列表树_移动项
nFlag=移动方式
nDestItem=目标项ID
nMoveItem=要移动的项ID
hEle=元素句柄

[func XTree_GetTemplateObject(hEle int, nID, nTempItemID int32) int {]
ff=列表树_取模板对象
nTempItemID=模板项ID
nID=项ID
hEle=元素句柄

[func XTree_GetItemIDFromHXCGUI(hEle int, hXCGUI int) int32 {]
ff=列表树_取对象所在项
hXCGUI=对象句柄
hEle=元素句柄

[func XTree_InsertItemText(hEle int, pValue string, nParentID, insertID int32) int32 {]
ff=列表树_插入项文本
insertID=插入位置ID
nParentID=父ID
pValue=值
hEle=元素句柄

[func XTree_InsertItemTextEx(hEle int, pName string, pValue string, nParentID, insertID int32) int32 {]
ff=列表树_插入项文本EX
insertID=插入位置ID
nParentID=父ID
pValue=值
pName=名称
hEle=元素句柄

[func XTree_InsertItemImage(hEle int, hImage int, nParentID, insertID int32) int32 {]
ff=列表树_插入项图片
insertID=插入位置ID
nParentID=父ID
hImage=图片句柄
hEle=元素句柄

[func XTree_InsertItemImageEx(hEle int, pName string, hImage int, nParentID, insertID int32) int32 {]
ff=列表树_插入项图片EX
insertID=插入位置ID
nParentID=父ID
hImage=图片句柄
pName=名称
hEle=元素句柄

[func XTree_GetCount(hEle int) int32 {]
ff=列表树_取项数量
hEle=元素句柄

[func XTree_GetCountColumn(hEle int) int32 {]
ff=列表树_取列数量
hEle=元素句柄

[func XTree_SetItemText(hEle int, nID, iColumn int32, pValue string) bool {]
ff=列表树_置项文本
pValue=值
iColumn=列索引
nID=项ID
hEle=元素句柄

[func XTree_SetItemTextEx(hEle int, nID int32, pName string, pValue string) bool {]
ff=列表树_置项文本EX
pValue=值
pName=名称
nID=项ID
hEle=元素句柄

[func XTree_SetItemImage(hEle int, nID, iColumn int32, hImage int) bool {]
ff=列表树_置项图片
hImage=图片句柄
iColumn=列索引
nID=项ID
hEle=元素句柄

[func XTree_SetItemImageEx(hEle int, nID int32, pName string, hImage int) bool {]
ff=列表树_置项图片EX
hImage=图片句柄
pName=名称
nID=项ID
hEle=元素句柄

[func XTree_GetItemText(hEle int, nID, iColumn int32) string {]
ff=列表树_取项文本
iColumn=列索引
nID=项ID
hEle=元素句柄

[func XTree_GetItemTextEx(hEle int, nID int32, pName string) string {]
ff=列表树_取项文本EX
pName=名称
nID=项ID
hEle=元素句柄

[func XTree_GetItemImage(hEle int, nID, iColumn int32) int {]
ff=列表树_取项图片
iColumn=列索引
nID=项ID
hEle=元素句柄

[func XTree_GetItemImageEx(hEle int, nID int32, pName string) int {]
ff=列表树_取项图片EX
pName=名称
nID=项ID
hEle=元素句柄

[func XTree_DeleteItem(hEle int, nID int32) bool {]
ff=列表树_删除项
nID=项ID
hEle=元素句柄

[func XTree_DeleteItemAll(hEle int) {]
ff=列表树_删除全部项
hEle=元素句柄

[func XTree_DeleteColumnAll(hEle int) {]
ff=列表树_删除列全部
hEle=元素句柄

[func XTree_SetSplitLineColor(hEle int, color int) {]
ff=列表树_置分割线颜色
color=ABGR颜色值
hEle=元素句柄

[func XTree_SetItemTemplateXMLFromMem(hEle int, data #左中括号##右中括号#byte) bool {]
ff=列表树_置项模板从内存
data=模板数据
hEle=元素句柄

[func XTree_SetItemTemplateXMLFromZipRes(hEle int, id int32, pFileName string, pPassword string, hModule uintptr) bool {]
ff=列表树_置项模板从资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=文件名
id=RC资源ID
hEle=元素句柄

[func XTree_GetItemTemplate(hEle int) int {]
ff=列表树_取项模板
hEle=元素句柄
