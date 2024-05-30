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

[func NewList(x int, y int, cx int, cy int, hParent int) *List {]
ff=创建列表
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewListEx(x, y, cx, cy int32, hParent, col_extend_count int32) *List {]
ff=创建列表Ex
col_extend_count=列数量
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewListByHandle(handle int) *List {]
ff=创建列表并按句柄

[func NewListByName(name string) *List {]
ff=创建列表并按名称

[func NewListByUID(nUID int) *List {]
ff=创建列表并按UID

[func NewListByUIDName(name string) *List {]
ff=创建列表并按UID名称

[func (l *List) AddColumn(width int) int {]
ff=增加列
width=列宽度

[func (l *List) InsertColumn(width int, iItem int) int {]
ff=插入列
iItem=插入位置索引
width=列宽度

[func (l *List) EnableMultiSel(bEnable bool) int {]
ff=启用多选
bEnable=是否启用

[func (l *List) EnableDragChangeColumnWidth(bEnable bool) int {]
ff=启用拖动更改列宽
bEnable=是否启用

[func (l *List) EnableVScrollBarTop(bTop bool) int {]
ff=启用垂直滚动条顶部对齐
bTop=是否启用

[func (l *List) EnableItemBkFullRow(bFull bool) int {]
ff=启用项背景全行模式
bFull=是否启用

[func (l *List) EnableFixedRowHeight(bEnable bool) int {]
ff=启用固定行高
bEnable=是否启用

[func (l *List) EnableTemplateReuse(bEnable bool) int {]
ff=启用模板复用
bEnable=是否启用

[func (l *List) EnableVirtualTable(bEnable bool) int {]
ff=启用虚表
bEnable=是否启用

[func (l *List) SetVirtualRowCount(nRowCount int) int {]
ff=置虚表行数
nRowCount=行数

[func (l *List) SetSort(iColumn int, iColumnAdapter int, bEnable bool) int {]
ff=置排序
bEnable=是否启用排序
iColumnAdapter=数据适配器中列索引
iColumn=列索引

[func (l *List) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) int {]
ff=置绘制项背景标志
nFlags=标志位

[func (l *List) SetColumnWidth(iItem int, width int) int {]
ff=置列宽
width=宽度
iItem=列索引

[func (l *List) SetColumnMinWidth(iItem int, width int) int {]
ff=置列最小宽度
width=宽度
iItem=列索引

[func (l *List) SetColumnWidthFixed(iColumn int, bFixed bool) int {]
ff=置列宽度固定
bFixed=是否固定宽度
iColumn=列索引

[func (l *List) GetColumnWidth(iColumn int) int {]
ff=取列宽度
iColumn=列索引

[func (l *List) GetColumnCount() int {]
ff=取列数量

[func (l *List) SetItemData(iItem int, iSubItem int, data int) bool {]
ff=置项数据
data=用户数据
iSubItem=子项索引
iItem=项索引

[func (l *List) GetItemData(iItem int, iSubItem int) int {]
ff=取项数据
iSubItem=子项索引
iItem=项索引

[func (l *List) SetSelectItem(iItem int) bool {]
ff=置选择项
iItem=项索引

[func (l *List) GetSelectItem() int {]
ff=取选择项

[func (l *List) GetSelectItemCount() int {]
ff=取选择项数量

[func (l *List) AddSelectItem(iItem int) bool {]
ff=添加选择项
iItem=项索引

[func (l *List) SetSelectAll() int {]
ff=置选择全部

[func (l *List) GetSelectAll(pArray *#左中括号##右中括号#int32, nArraySize int) int {]
ff=取全部选择
nArraySize=切片大小
pArray=接收行索引切片

[func (l *List) VisibleItem(iItem int) int {]
ff=显示指定项
iItem=项索引

[func (l *List) CancelSelectItem(iItem int) bool {]
ff=取消选择项
iItem=项索引

[func (l *List) CancelSelectAll() int {]
ff=取消全部选择项

[func (l *List) GetHeaderHELE() int {]
ff=取列表头

[func (l *List) DeleteColumn(iItem int) bool {]
ff=删除列
iItem=项索引

[func (l *List) DeleteColumnAll() int {]
ff=删除全部列

[func (l *List) BindAdapter(hAdapter int) int {]
ff=绑定数据适配器
hAdapter=数据适配器句柄

[func (l *List) BindAdapterHeader(hAdapter int) int {]
ff=列表头绑定数据适配器
hAdapter=数据适配器句柄

[func (l *List) CreateAdapter(colExtend_count int) int {]
ff=创建数据适配器

[func (l *List) CreateAdapterHeader() int {]
ff=创建列表头数据适配器

[func (l *List) GetAdapter() int {]
ff=取数据适配器

[func (l *List) GetAdapterHeader() int {]
ff=取列表头数据适配器

[func (l *List) SetItemTemplateXML(pXmlFile string) bool {]
ff=置项模板文件
pXmlFile=文件名

[func (l *List) SetItemTemplateXMLFromString(pStringXML string) bool {]
ff=置项模板从字符串
pStringXML=字符串

[func (l *List) SetItemTemplate(hTemp int) bool {]
ff=置项模板
hTemp=模板句柄

[func (l *List) GetTemplateObject(iItem int, iSubItem int, nTempItemID int) int {]
ff=取项模板对象
nTempItemID=模板项itemID
iSubItem=子项索引
iItem=项索引

[func (l *List) GetItemIndexFromHXCGUI(hXCGUI int) int {]
ff=取对象所在行
hXCGUI=对象句柄

[func (l *List) GetHeaderTemplateObject(iItem int, nTempItemID int) int {]
ff=取列表头模板对象
nTempItemID=模板项ID
iItem=列表头项ID

[func (l *List) GetHeaderItemIndexFromHXCGUI(hXCGUI int) int {]
ff=取列表头对象所在行
hXCGUI=对象句柄

[func (l *List) SetHeaderHeight(height int) int {]
ff=置列表头高度
height=高度

[func (l *List) GetHeaderHeight() int {]
ff=取列表头高度

[func (l *List) GetVisibleRowRange(piStart *int32, piEnd *int32) int {]
ff=取可视行范围
piEnd=结束行索引
piStart=开始行索引

[func (l *List) SetItemHeightDefault(nHeight int32, nSelHeight int32) int {]
ff=置项默认高度
nSelHeight=选中时高度
nHeight=高度

[func (l *List) GetItemHeightDefault(pHeight *int32, pSelHeight *int32) int {]
ff=取项默认高度
pSelHeight=选中时高度
pHeight=高度

[func (l *List) SetRowSpace(nSpace int) int {]
ff=置行间距
nSpace=行间距大小

[func (l *List) GetRowSpace() int {]
ff=取行间距

[func (l *List) SetLockColumnLeft(iColumn int) int {]
ff=置锁定列左侧
iColumn=列索引

[func (l *List) SetLockColumnRight(iColumn int) int {]
ff=置锁定列右侧
iColumn=列索引

[func (l *List) SetLockRowBottom(bLock bool) int {]
ff=置锁定行底部
bLock=是否锁定

[func (l *List) SetLockRowBottomOverlap(bOverlap bool) int {]
ff=置锁定行底部重叠
bOverlap=是否重叠

[func (l *List) HitTest(pPt *xc.POINT, piItem *int32, piSubItem *int32) bool {]
ff=测试点击项
pPt=坐标点

[func (l *List) HitTestOffset(pPt *xc.POINT, piItem *int32, piSubItem *int32) bool {]
ff=测试点击项EX
pPt=坐标点

[func (l *List) RefreshData() int {]
ff=刷新项数据

[func (l *List) RefreshItem(iItem int) int {]
ff=刷新指定项
iItem=项索引

[func (l *List) AddColumnText(nWidth int, pName string, pText string) int {]
ff=添加列文本
pText=文本
pName=名称
nWidth=列宽

[func (l *List) AddColumnImage(nWidth int, pName string, hImage int) int {]
ff=添加列图片
hImage=图片句柄
pName=名称
nWidth=列宽

[func (l *List) AddItemText(pText string) int {]
ff=添加项文本
pText=文本

[func (l *List) AddItemTextEx(pName string, pText string) int {]
ff=添加项文本EX
pText=文本
pName=名称

[func (l *List) AddItemImage(hImage int) int {]
ff=添加项图片
hImage=图片

[func (l *List) AddItemImageEx(pName string, hImage int) int {]
ff=添加项图片EX
hImage=图片
pName=名称

[func (l *List) InsertItemText(iItem int, pValue string) int {]
ff=插入项文本
pValue=文本
iItem=项

[func (l *List) InsertItemTextEx(iItem int, pName string, pValue string) int {]
ff=插入项文本EX
pValue=文本
pName=名称
iItem=项

[func (l *List) InsertItemImage(iItem int, hImage int) int {]
ff=插入项图片
hImage=图片
iItem=项

[func (l *List) InsertItemImageEx(iItem int, pName string, hImage int) int {]
ff=插入项图片EX
hImage=图片
pName=名称
iItem=项

[func (l *List) SetItemText(iItem int, iColumn int, pText string) bool {]
ff=置项文本
pText=文本
iColumn=列
iItem=项

[func (l *List) SetItemTextEx(iItem int, pName string, pText string) bool {]
ff=置项文本EX
pText=文本
pName=名称
iItem=项

[func (l *List) SetItemImage(iItem int, iColumn int, hImage int) bool {]
ff=置项图片
hImage=图片
iColumn=列
iItem=项

[func (l *List) SetItemImageEx(iItem int, pName string, hImage int) bool {]
ff=置项图片EX
hImage=图片
pName=名称
iItem=项

[func (l *List) SetItemInt(iItem int, iColumn int, nValue int) bool {]
ff=置项指数值
nValue=值
iColumn=列
iItem=项

[func (l *List) SetItemIntEx(iItem int, pName string, nValue int) bool {]
ff=置项整数值EX
nValue=值
pName=名称
iItem=项

[func (l *List) SetItemFloat(iItem int, iColumn int, fFloat float32) bool {]
ff=置项浮点值
fFloat=值
iColumn=列
iItem=项

[func (l *List) SetItemFloatEx(iItem int, pName string, fFloat float32) bool {]
ff=置项浮点值EX
fFloat=值
pName=名称
iItem=项

[func (l *List) GetItemText(iItem int, iColumn int) string {]
ff=取项文本
iColumn=列
iItem=项

[func (l *List) GetItemTextEx(iItem int, pName string) string {]
ff=取项文本EX
pName=名称
iItem=项

[func (l *List) GetItemImage(iItem int, iColumn int) int {]
ff=取项图片
iColumn=列
iItem=项

[func (l *List) GetItemImageEx(iItem int, pName string) int {]
ff=取项图片EX
pName=名称
iItem=项

[func (l *List) GetItemInt(iItem int, iColumn int, pOutValue *int32) bool {]
ff=取项整数值
pOutValue=值指针
iColumn=列
iItem=项

[func (l *List) GetItemIntEx(iItem int, pName string, pOutValue *int32) bool {]
ff=取项整数值EX
pOutValue=值指针
pName=文本
iItem=项

[func (l *List) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {]
ff=取项浮点值
pOutValue=值指针
iColumn=列
iItem=项

[func (l *List) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {]
ff=取项浮点值EX
pOutValue=值指针
pName=名称
iItem=项

[func (l *List) DeleteItem(iItem int) bool {]
ff=删除项
iItem=项

[func (l *List) DeleteItemEx(iItem int, nCount int) bool {]
ff=删除项EX
nCount=计数
iItem=项

[func (l *List) DeleteItemAll() int {]
ff=删除项全部

[func (l *List) DeleteColumnAll_AD() int {]
ff=删除列全部AD

[func (l *List) GetCount_AD() int {]
ff=取项数量AD

[func (l *List) GetCountColumn_AD() int {]
ff=取列数量AD

[func (l *List) SetSplitLineColor(color int) int {]
ff=置分割线颜色
color=ABGR颜色值

[func (l *List) SetItemHeight(iRow int, nHeight, nSelHeight int32) int {]
ff=置项高度
nSelHeight=选中时高度
nHeight=高度
iRow=行索引

[func (l *List) GetItemHeight(iRow int, pHeight, pSelHeight *int32) int {]
ff=取项高度
pSelHeight=选中时高度
pHeight=高度
iRow=行索引

[func (l *List) SetDragRectColor(color, width int) int {]
ff=置拖动矩形颜色
width=线宽度
color=ABGR颜色值

[func (l *List) GetItemTemplate() int {]
ff=取项模板

[func (l *List) GetItemTemplateHeader() int {]
ff=取项模板列表头

[func (l *List) RefreshDataHeader() int {]
ff=刷新项数据列表头

[func (l *List) SetItemTemplateXMLFromMem(data #左中括号##右中括号#byte) bool {]
ff=置项模板从内存
data=模板数据

[func (l *List) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule uintptr) bool {]
ff=置项模板并按资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=项模板文件名
id=RC资源ID

[func (l *List) Event_LIST_TEMP_CREATE(pFun XE_LIST_TEMP_CREATE) bool {]
ff=事件_项模板创建

[func (l *List) Event_LIST_TEMP_CREATE1(pFun XE_LIST_TEMP_CREATE1) bool {]
ff=事件_项模板创建1

[func (l *List) Event_LIST_TEMP_CREATE_END(pFun XE_LIST_TEMP_CREATE_END) bool {]
ff=事件_项模板创建完成

[func (l *List) Event_LIST_TEMP_CREATE_END1(pFun XE_LIST_TEMP_CREATE_END1) bool {]
ff=事件_项模板创建完成1

[func (l *List) Event_LIST_TEMP_DESTROY(pFun XE_LIST_TEMP_DESTROY) bool {]
ff=事件_项模板销毁

[func (l *List) Event_LIST_TEMP_DESTROY1(pFun XE_LIST_TEMP_DESTROY1) bool {]
ff=事件_项模板销毁1

[func (l *List) Event_LIST_TEMP_ADJUST_COORDINATE(pFun XE_LIST_TEMP_ADJUST_COORDINATE) bool {]
ff=停用_项模板调整坐标

[func (l *List) Event_LIST_TEMP_ADJUST_COORDINATE1(pFun XE_LIST_TEMP_ADJUST_COORDINATE1) bool {]
ff=停用_项模板调整坐标1

[func (l *List) Event_LIST_DRAWITEM(pFun XE_LIST_DRAWITEM) bool {]
ff=事件_绘制项

[func (l *List) Event_LIST_DRAWITEM1(pFun XE_LIST_DRAWITEM1) bool {]
ff=事件_绘制项1

[func (l *List) Event_LIST_SELECT(pFun XE_LIST_SELECT) bool {]
ff=事件_项选择

[func (l *List) Event_LIST_SELECT1(pFun XE_LIST_SELECT1) bool {]
ff=事件_项选择1

[func (l *List) Event_LIST_HEADER_DRAWITEM(pFun XE_LIST_HEADER_DRAWITEM) bool {]
ff=事件_列表元素绘制列表头项

[func (l *List) Event_LIST_HEADER_DRAWITEM1(pFun XE_LIST_HEADER_DRAWITEM1) bool {]
ff=事件_列表元素绘制列表头项1

[func (l *List) Event_LIST_HEADER_CLICK(pFun XE_LIST_HEADER_CLICK) bool {]
ff=事件_列表头项点击

[func (l *List) Event_LIST_HEADER_CLICK1(pFun XE_LIST_HEADER_CLICK1) bool {]
ff=事件_列表头项点击1

[func (l *List) Event_LIST_HEADER_WIDTH_CHANGE(pFun XE_LIST_HEADER_WIDTH_CHANGE) bool {]
ff=事件_列表头项宽度改变

[func (l *List) Event_LIST_HEADER_WIDTH_CHANGE1(pFun XE_LIST_HEADER_WIDTH_CHANGE1) bool {]
ff=事件_列表头项宽度改变1

[func (l *List) Event_LIST_HEADER_TEMP_CREATE(pFun XE_LIST_HEADER_TEMP_CREATE) bool {]
ff=事件_列表头项模板创建

[func (l *List) Event_LIST_HEADER_TEMP_CREATE1(pFun XE_LIST_HEADER_TEMP_CREATE1) bool {]
ff=事件_列表头项模板创建1

[func (l *List) Event_LIST_HEADER_TEMP_CREATE_END(pFun XE_LIST_HEADER_TEMP_CREATE_END) bool {]
ff=事件_列表头项模板创建完成

[func (l *List) Event_LIST_HEADER_TEMP_CREATE_END1(pFun XE_LIST_HEADER_TEMP_CREATE_END1) bool {]
ff=事件_列表头项模板创建完成1

[func (l *List) Event_LIST_HEADER_TEMP_DESTROY(pFun XE_LIST_HEADER_TEMP_DESTROY) bool {]
ff=事件_列表头项模板销毁

[func (l *List) Event_LIST_HEADER_TEMP_DESTROY1(pFun XE_LIST_HEADER_TEMP_DESTROY1) bool {]
ff=事件_列表头项模板销毁1

[func (l *List) Event_LIST_HEADER_TEMP_ADJUST_COORDINATE(pFun XE_LIST_HEADER_TEMP_ADJUST_COORDINATE) bool {]
ff=停用_列表头项模板调整坐标
[func (l *List) Event_LIST_HEADER_TEMP_ADJUST_COORDINATE1(pFun XE_LIST_HEADER_TEMP_ADJUST_COORDINATE1) bool {]
ff=停用_列表头项模板调整坐标1
