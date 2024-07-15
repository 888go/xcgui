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

[func XList_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=列表_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XList_CreateEx(x, y, cx, cy int32, hParent, col_extend_count int32) int {]
ff=列表_创建Ex
col_extend_count=列数量
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XList_AddColumn(hEle int, width int) int {]
ff=列表_增加列
width=列宽度
hEle=元素句柄

[func XList_InsertColumn(hEle int, width int, iItem int) int {]
ff=列表_插入列
iItem=项索引
width=列宽度
hEle=元素句柄

[func XList_EnableMultiSel(hEle int, bEnable bool) int {]
ff=列表_启用多选
bEnable=是否启用
hEle=元素句柄

[func XList_EnableDragChangeColumnWidth(hEle int, bEnable bool) int {]
ff=列表_启用拖动更改列宽
bEnable=是否启用
hEle=元素句柄

[func XList_EnableVScrollBarTop(hEle int, bTop bool) int {]
ff=列表_启用垂直滚动条顶部对齐
bTop=是否启用
hEle=元素句柄

[func XList_EnableItemBkFullRow(hEle int, bFull bool) int {]
ff=列表_启用项背景全行模式
bFull=是否启用
hEle=元素句柄

[func XList_EnableFixedRowHeight(hEle int, bEnable bool) int {]
ff=列表_启用固定行高
bEnable=是否启用
hEle=元素句柄

[func XList_EnableTemplateReuse(hEle int, bEnable bool) int {]
ff=列表_启用模板复用
bEnable=是否启用
hEle=元素句柄

[func XList_EnableVirtualTable(hEle int, bEnable bool) int {]
ff=列表_启用虚表
bEnable=是否启用
hEle=元素句柄

[func XList_SetVirtualRowCount(hEle int, nRowCount int) int {]
ff=列表_置虚表行数
nRowCount=行数
hEle=元素句柄

[func XList_SetSort(hEle int, iColumn int, iColumnAdapter int, bEnable bool) int {]
ff=列表_置排序
bEnable=是否启用排序功能
iColumnAdapter=适配器中列索引
iColumn=列索引
hEle=元素句柄

[func XList_SetDrawItemBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) int {]
ff=列表_置绘制项背景标志
nFlags=标志位
hEle=元素句柄

[func XList_SetColumnWidth(hEle int, iItem int, width int) int {]
ff=列表_置列宽
width=宽度
iItem=项索引
hEle=元素句柄

[func XList_SetColumnMinWidth(hEle int, iItem int, width int) int {]
ff=列表_置列最小宽度
width=宽度
iItem=项索引
hEle=元素句柄

[func XList_SetColumnWidthFixed(hEle int, iColumn int, bFixed bool) int {]
ff=列表_置列宽度固定
bFixed=是否固定宽度
iColumn=列索引
hEle=元素句柄

[func XList_GetColumnWidth(hEle int, iColumn int) int {]
ff=列表_取列宽度
iColumn=列索引
hEle=元素句柄

[func XList_GetColumnCount(hEle int) int {]
ff=列表_取列数量
hEle=元素句柄

[func XList_SetItemData(hEle int, iItem int, iSubItem int, data int) bool {]
ff=列表_置项数据
data=用户数据
iSubItem=子项索引
iItem=项索引
hEle=元素句柄

[func XList_GetItemData(hEle int, iItem int, iSubItem int) int {]
ff=列表_取项数据
iSubItem=子项索引
iItem=项索引
hEle=元素句柄

[func XList_SetSelectItem(hEle int, iItem int) bool {]
ff=列表_置选择项
iItem=项索引
hEle=元素句柄

[func XList_GetSelectItem(hEle int) int {]
ff=列表_取选择项
hEle=元素句柄

[func XList_GetSelectItemCount(hEle int) int {]
ff=列表_取选择项数量
hEle=元素句柄

[func XList_AddSelectItem(hEle int, iItem int) bool {]
ff=列表_添加选择项
iItem=项索引
hEle=元素句柄

[func XList_SetSelectAll(hEle int) int {]
ff=列表_置选择全部
hEle=元素句柄

[func XList_GetSelectAll(hEle int, pArray *#左中括号##右中括号#int32, nArraySize int) int {]
ff=列表_取全部选择
nArraySize=切片大小
pArray=接收行索引切片
hEle=元素句柄

[func XList_VisibleItem(hEle int, iItem int) int {]
ff=列表_显示指定项
iItem=项索引
hEle=元素句柄

[func XList_CancelSelectItem(hEle int, iItem int) bool {]
ff=列表_取消选择项
iItem=项索引
hEle=元素句柄

[func XList_CancelSelectAll(hEle int) int {]
ff=列表_取消全部选择项
hEle=元素句柄

[func XList_GetHeaderHELE(hEle int) int {]
ff=列表_取列表头
hEle=元素句柄

[func XList_DeleteColumn(hEle int, iItem int) bool {]
ff=列表_删除列
iItem=项索引
hEle=元素句柄

[func XList_DeleteColumnAll(hEle int) int {]
ff=列表_删除全部列
hEle=元素句柄

[func XList_BindAdapter(hEle int, hAdapter int) int {]
ff=列表_绑定数据适配器
hAdapter=数据适配器句柄
hEle=元素句柄

[func XList_BindAdapterHeader(hEle int, hAdapter int) int {]
ff=列表_列表头绑定数据适配器
hAdapter=数据适配器句柄
hEle=元素句柄

[func XList_CreateAdapter(hEle int, colExtend_count int) int {]
ff=列表_创建数据适配器
colExtend_count=预计列表总列数
hEle=元素句柄

[func XList_CreateAdapterHeader(hEle int) int {]
ff=列表_列表头创建数据适配器
hEle=元素句柄

[func XList_GetAdapter(hEle int) int {]
ff=列表_取数据适配器
hEle=元素句柄

[func XList_GetAdapterHeader(hEle int) int {]
ff=列表_列表头获取数据适配器
hEle=元素句柄

[func XList_SetItemTemplateXML(hEle int, pXmlFile string) bool {]
ff=列表_置项模板文件
pXmlFile=文件名
hEle=元素句柄

[func XList_SetItemTemplateXMLFromString(hEle int, pStringXML string) bool {]
ff=列表_置项模板从字符串
pStringXML=字符串
hEle=元素句柄

[func XList_SetItemTemplate(hEle int, hTemp int) bool {]
ff=列表_置项模板
hTemp=模板句柄
hEle=元素句柄

[func XList_GetTemplateObject(hEle int, iItem int, iSubItem int, nTempItemID int) int {]
ff=列表_取项模板对象
nTempItemID=模板项ID
iSubItem=子项索引
iItem=项索引
hEle=元素句柄

[func XList_GetItemIndexFromHXCGUI(hEle int, hXCGUI int) int {]
ff=列表_取对象所在行
hXCGUI=对象句柄
hEle=元素句柄

[func XList_GetHeaderTemplateObject(hEle int, iItem int, nTempItemID int) int {]
ff=列表_取列表头模板对象
nTempItemID=模板项ID
iItem=项索引
hEle=元素句柄

[func XList_GetHeaderItemIndexFromHXCGUI(hEle int, hXCGUI int) int {]
ff=列表_取列表头对象所在行
hXCGUI=对象句柄
hEle=元素句柄

[func XList_SetHeaderHeight(hEle int, height int) int {]
ff=列表_置列表头高度
height=高度
hEle=元素句柄

[func XList_GetHeaderHeight(hEle int) int {]
ff=列表_取列表头高度
hEle=元素句柄

[func XList_GetVisibleRowRange(hEle int, piStart *int32, piEnd *int32) int {]
ff=列表_取可视行范围
piEnd=结束行索引
piStart=开始行索引
hEle=元素句柄

[func XList_SetItemHeightDefault(hEle int, nHeight int32, nSelHeight int32) int {]
ff=列表_置项默认高度
nSelHeight=选中时高度
nHeight=高度
hEle=元素句柄

[func XList_GetItemHeightDefault(hEle int, pHeight *int32, pSelHeight *int32) int {]
ff=列表_取项默认高度
pSelHeight=选中时高度
pHeight=高度
hEle=元素句柄

[func XList_SetRowSpace(hEle int, nSpace int) int {]
ff=列表_置行间距
nSpace=行间距大小
hEle=元素句柄

[func XList_GetRowSpace(hEle int) int {]
ff=列表_取行间距
hEle=元素句柄

[func XList_SetLockColumnLeft(hEle int, iColumn int) int {]
ff=列表_置锁定列左侧
iColumn=列索引
hEle=元素句柄

[func XList_SetLockColumnRight(hEle int, iColumn int) int {]
ff=列表_置锁定列右侧
iColumn=列索引
hEle=元素句柄

[func XList_SetLockRowBottom(hEle int, bLock bool) int {]
ff=列表_置锁定行底部
bLock=是否锁定
hEle=元素句柄

[func XList_SetLockRowBottomOverlap(hEle int, bOverlap bool) int {]
ff=列表_置锁定行底部重叠
bOverlap=是否重叠
hEle=元素句柄

[func XList_HitTest(hEle int, pPt *POINT, piItem *int32, piSubItem *int32) bool {]
ff=列表_测试点击项
piSubItem=子项索引
piItem=项索引
pPt=坐标点
hEle=元素句柄

[func XList_HitTestOffset(hEle int, pPt *POINT, piItem *int32, piSubItem *int32) bool {]
ff=列表_测试点击项EX
piSubItem=子项索引
piItem=项索引
pPt=坐标点
hEle=元素句柄

[func XList_RefreshData(hEle int) int {]
ff=列表_刷新项数据
hEle=元素句柄

[func XList_RefreshItem(hEle int, iItem int) int {]
ff=列表_刷新指定项
iItem=项索引
hEle=元素句柄

[func XList_AddColumnText(hEle int, nWidth int, pName string, pText string) int {]
ff=列表_添加列文本
pText=文本
pName=名称
nWidth=列宽
hEle=元素句柄

[func XList_AddColumnImage(hEle int, nWidth int, pName string, hImage int) int {]
ff=列表_添加列图片
hImage=图片
pName=名称
nWidth=列宽
hEle=元素句柄

[func XList_AddItemText(hEle int, pText string) int {]
ff=列表_添加项文本
pText=文本
hEle=元素句柄

[func XList_AddItemTextEx(hEle int, pName string, pText string) int {]
ff=列表_添加项文本EX
pText=文本
pName=名称
hEle=元素句柄

[func XList_AddItemImage(hEle int, hImage int) int {]
ff=列表_添加项图片
hImage=图片
hEle=元素句柄

[func XList_AddItemImageEx(hEle int, pName string, hImage int) int {]
ff=列表_添加项图片EX
hImage=图片
pName=名称
hEle=元素句柄

[func XList_InsertItemText(hEle int, iItem int, pValue string) int {]
ff=列表_插入项文本
pValue=值
iItem=项索引
hEle=元素句柄

[func XList_InsertItemTextEx(hEle int, iItem int, pName string, pValue string) int {]
ff=列表_插入项文本EX
pValue=值
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_InsertItemImage(hEle int, iItem int, hImage int) int {]
ff=列表_插入项图片
hImage=图片
iItem=项索引
hEle=元素句柄

[func XList_InsertItemImageEx(hEle int, iItem int, pName string, hImage int) int {]
ff=列表_插入项图片EX
hImage=图片
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_SetItemText(hEle int, iItem int, iColumn int, pText string) bool {]
ff=列表_置项文本
pText=文本
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XList_SetItemTextEx(hEle int, iItem int, pName string, pText string) bool {]
ff=列表_置项文本EX
pText=文本
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_SetItemImage(hEle int, iItem int, iColumn int, hImage int) bool {]
ff=列表_置项图片
hImage=图片
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XList_SetItemImageEx(hEle int, iItem int, pName string, hImage int) bool {]
ff=列表_置项图片EX
hImage=图片
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_SetItemInt(hEle int, iItem int, iColumn int, nValue int) bool {]
ff=列表_置项指数值
nValue=值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XList_SetItemIntEx(hEle int, iItem int, pName string, nValue int) bool {]
ff=列表_置项整数值EX
nValue=值
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_SetItemFloat(hEle int, iItem int, iColumn int, fFloat float32) bool {]
ff=列表_置项浮点值
fFloat=值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XList_SetItemFloatEx(hEle int, iItem int, pName string, fFloat float32) bool {]
ff=列表_置项浮点值EX
fFloat=值
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_GetItemText(hEle int, iItem int, iColumn int) string {]
ff=列表_取项文本
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XList_GetItemTextEx(hEle int, iItem int, pName string) string {]
ff=列表_取项文本EX
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_GetItemImage(hEle int, iItem int, iColumn int) int {]
ff=列表_取项图片
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XList_GetItemImageEx(hEle int, iItem int, pName string) int {]
ff=列表_取项图片EX
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_GetItemInt(hEle int, iItem int, iColumn int, pOutValue *int32) bool {]
ff=列表_取项整数值
pOutValue=返回值指针
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XList_GetItemIntEx(hEle int, iItem int, pName string, pOutValue *int32) bool {]
ff=列表_取项整数值EX
pOutValue=返回值指针
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_GetItemFloat(hEle int, iItem int, iColumn int, pOutValue *float32) bool {]
ff=列表_取项浮点值
pOutValue=返回值指针
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XList_GetItemFloatEx(hEle int, iItem int, pName string, pOutValue *float32) bool {]
ff=列表_取项浮点值EX
pOutValue=返回值指针
pName=名称
iItem=项索引
hEle=元素句柄

[func XList_DeleteItem(hEle int, iItem int) bool {]
ff=列表_删除项
iItem=项索引
hEle=元素句柄

[func XList_DeleteItemEx(hEle int, iItem int, nCount int) bool {]
ff=列表_删除项EX
nCount=数量
iItem=项索引
hEle=元素句柄

[func XList_DeleteItemAll(hEle int) int {]
ff=列表_删除项全部
hEle=元素句柄

[func XList_DeleteColumnAll_AD(hEle int) int {]
ff=列表_删除列全部AD
hEle=元素句柄

[func XList_GetCount_AD(hEle int) int {]
ff=列表_取项数量AD
hEle=元素句柄

[func XList_GetCountColumn_AD(hEle int) int {]
ff=列表_取列数量AD
hEle=元素句柄

[func XList_SetSplitLineColor(hEle int, color int) int {]
ff=列表_置分割线颜色
color=ABGR颜色值
hEle=元素句柄

[func XList_SetItemHeight(hEle int, iRow int, nHeight, nSelHeight int32) int {]
ff=列表_置项高度
nSelHeight=选中时高度
nHeight=高度
iRow=行索引
hEle=元素句柄

[func XList_GetItemHeight(hEle int, iRow int, pHeight, pSelHeight *int32) int {]
ff=列表_取项高度
pSelHeight=选中时高度
pHeight=高度
iRow=行索引
hEle=元素句柄

[func XList_SetDragRectColor(hEle int, color, width int) int {]
ff=列表_置拖动矩形颜色
width=线宽度
color=ABGR颜色值
hEle=元素句柄

[func XList_GetItemTemplate(hEle int) int {]
ff=列表_取项模板
hEle=元素句柄

[func XList_GetItemTemplateHeader(hEle int) int {]
ff=列表_取项模板列表头
hEle=元素句柄

[func XList_RefreshDataHeader(hEle int) int {]
ff=列表_刷新项数据列表头
hEle=元素句柄

[func XList_SetItemTemplateXMLFromMem(hEle int, data #左中括号##右中括号#byte) bool {]
ff=列表_置项模板从内存
data=模板数据
hEle=元素句柄

[func XList_SetItemTemplateXMLFromZipRes(hEle int, id int32, pFileName string, pPassword string, hModule uintptr) bool {]
ff=列表_置项模板从资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=项模板文件名
id=RC资源ID
hEle=元素句柄
