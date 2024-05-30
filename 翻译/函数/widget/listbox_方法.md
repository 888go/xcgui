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

[func NewListBox(x int, y int, cx int, cy int, hParent int) *ListBox {]
ff=创建列表框
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewListBoxEx(x, y, cx, cy int32, hParent, col_extend_count int32) *ListBox {]
ff=创建列表框Ex
col_extend_count=列数量
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewListBoxByHandle(handle int) *ListBox {]
ff=创建列表框并按句柄

[func NewListBoxByName(name string) *ListBox {]
ff=创建列表框并按名称

[func NewListBoxByUID(nUID int) *ListBox {]
ff=创建列表框并按UID

[func NewListBoxByUIDName(name string) *ListBox {]
ff=创建列表框并按UID名称

[func (l *ListBox) EnableFixedRowHeight(bEnable bool) int {]
ff=启用固定行高
bEnable=是否启用

[func (l *ListBox) EnableTemplateReuse(bEnable bool) int {]
ff=启用模板复用
bEnable=是否启用

[func (l *ListBox) EnableVirtualTable(bEnable bool) int {]
ff=启用虚表
bEnable=是否启用

[func (l *ListBox) SetVirtualRowCount(nRowCount int) int {]
ff=置虚表行数
nRowCount=行数

[func (l *ListBox) SetDrawItemBkFlags(nFlags xcc.List_DrawItemBk_Flag_) int {]
ff=置绘制项背景标志
nFlags=标志位

[func (l *ListBox) SetItemData(iItem int, nUserData int) bool {]
ff=置项数据
nUserData=用户数据
iItem=想索引

[func (l *ListBox) GetItemData(iItem int) int {]
ff=取项数据
iItem=项索引

[func (l *ListBox) SetItemInfo(iItem int, pItem *xc.ListBox_Item_Info_) bool {]
ff=置项信息
pItem=项信息
iItem=项索引

[func (l *ListBox) GetItemInfo(iItem int, pItem *xc.ListBox_Item_Info_) bool {]
ff=取项背景信息
pItem=项信息
iItem=项索引

[func (l *ListBox) SetSelectItem(iItem int) bool {]
ff=置选择项
iItem=项索引

[func (l *ListBox) GetSelectItem() int {]
ff=取选择项

[func (l *ListBox) AddSelectItem(iItem int) bool {]
ff=添加选择项
iItem=项索引

[func (l *ListBox) CancelSelectItem(iItem int) bool {]
ff=取消选择项
iItem=项索引

[func (l *ListBox) CancelSelectAll() bool {]
ff=取消选择全部

[func (l *ListBox) GetSelectAll(pArray *#左中括号##右中括号#int32, nArraySize int) int {]
ff=取全部选择
nArraySize=切片大小
pArray=切片缓冲区

[func (l *ListBox) GetSelectCount() int {]
ff=取选择项数量

[func (l *ListBox) GetItemMouseStay() int {]
ff=取鼠标停留项

[func (l *ListBox) SelectAll() bool {]
ff=选择全部项

[func (l *ListBox) VisibleItem(iItem int) int {]
ff=显示指定项
iItem=项索引

[func (l *ListBox) GetVisibleRowRange(piStart *int32, piEnd *int32) int {]
ff=取可视行范围
piEnd=结束行索引
piStart=开始行索引

[func (l *ListBox) SetItemHeightDefault(nHeight int32, nSelHeight int32) int {]
ff=置项默认高度
nSelHeight=选中项高度
nHeight=项高度

[func (l *ListBox) GetItemHeightDefault(pHeight *int32, pSelHeight *int32) int {]
ff=取项默认高度
pSelHeight=选中时高度
pHeight=高度

[func (l *ListBox) GetItemIndexFromHXCGUI(hXCGUI int) int {]
ff=取所在行索引
hXCGUI=对象句柄

[func (l *ListBox) SetRowSpace(nSpace int) int {]
ff=置行间距
nSpace=间距大小

[func (l *ListBox) GetRowSpace() int {]
ff=取行间距

[func (l *ListBox) HitTest(pPt *xc.POINT) int {]
ff=测试点击项
pPt=坐标点

[func (l *ListBox) HitTestOffset(pPt *xc.POINT) int {]
ff=测试点击项EX
pPt=坐标点

[func (l *ListBox) SetItemTemplateXML(pXmlFile string) bool {]
ff=置项模板文件
pXmlFile=文件名

[func (l *ListBox) SetItemTemplate(hTemp int) bool {]
ff=置项模板
hTemp=模板句柄

[func (l *ListBox) SetItemTemplateXMLFromString(pStringXML string) bool {]
ff=置项模板并按字符串
pStringXML=字符串

[func (l *ListBox) GetTemplateObject(iItem int, nTempItemID int) int {]
ff=取模板对象
nTempItemID=模板项ID
iItem=项索引

[func (l *ListBox) EnableMultiSel(bEnable bool) int {]
ff=启用多选
bEnable=是否启用

[func (l *ListBox) CreateAdapter() int {]
ff=创建数据适配器

[func (l *ListBox) BindAdapter(hAdapter int) int {]
ff=绑定数据适配器
hAdapter=数据适配器句柄

[func (l *ListBox) GetAdapter() int {]
ff=取数据适配器

[func (l *ListBox) Sort(iColumnAdapter int, bAscending bool) int {]
ff=排序
bAscending=升序
iColumnAdapter=列索引

[func (l *ListBox) RefreshData() int {]
ff=刷新数据

[func (l *ListBox) RefreshItem(iItem int) int {]
ff=刷新指定项
iItem=项索引

[func (l *ListBox) AddItemText(pText string) int {]
ff=添加项文本
pText=文本

[func (l *ListBox) AddItemTextEx(pName string, pText string) int {]
ff=添加项文本EX
pText=文本
pName=名称

[func (l *ListBox) AddItemImage(hImage int) int {]
ff=添加项图片
hImage=图片

[func (l *ListBox) AddItemImageEx(pName string, hImage int) int {]
ff=添加项图片EX
hImage=图片
pName=名称

[func (l *ListBox) InsertItemText(iItem int, pValue string) int {]
ff=插入项文本
pValue=文本
iItem=项

[func (l *ListBox) InsertItemTextEx(iItem int, pName string, pValue string) int {]
ff=插入项文本EX
pValue=文本
pName=名称
iItem=项

[func (l *ListBox) InsertItemImage(iItem int, hImage int) int {]
ff=插入项图片
hImage=图片
iItem=项

[func (l *ListBox) InsertItemImageEx(iItem int, pName string, hImage int) int {]
ff=插入项图片EX
hImage=图片
pName=名称
iItem=项

[func (l *ListBox) SetItemText(iItem int, iColumn int, pText string) bool {]
ff=置项文本
pText=文本
iColumn=列
iItem=项

[func (l *ListBox) SetItemTextEx(iItem int, pName string, pText string) bool {]
ff=置项文本EX
pText=文本
pName=名称
iItem=项

[func (l *ListBox) SetItemImage(iItem int, iColumn int, hImage int) bool {]
ff=置项图片
hImage=图片
iColumn=列
iItem=项

[func (l *ListBox) SetItemImageEx(iItem int, pName string, hImage int) bool {]
ff=置项图片EX
hImage=图片
pName=名称
iItem=项

[func (l *ListBox) SetItemInt(iItem int, iColumn int, nValue int) bool {]
ff=置项整数值
nValue=文本
iColumn=列
iItem=项

[func (l *ListBox) SetItemIntEx(iItem int, pName string, nValue int) bool {]
ff=置项整数值EX
nValue=文本
pName=名称
iItem=项

[func (l *ListBox) SetItemFloat(iItem int, iColumn int, fFloat float32) bool {]
ff=置项浮点值
fFloat=小数值
iColumn=列
iItem=项

[func (l *ListBox) SetItemFloatEx(iItem int, pName string, fFloat float32) bool {]
ff=置项浮点值EX
fFloat=小数值
pName=名称
iItem=项

[func (l *ListBox) GetItemText(iItem int, iColumn int) string {]
ff=取项文本
iColumn=列
iItem=项

[func (l *ListBox) GetItemTextEx(iItem int, pName string) string {]
ff=取项文本EX
pName=名称
iItem=项

[func (l *ListBox) GetItemImage(iItem int, iColumn int) int {]
ff=取项图片
iColumn=列
iItem=项

[func (l *ListBox) GetItemImageEx(iItem int, pName string) int {]
ff=取项图片EX
pName=名称
iItem=项

[func (l *ListBox) GetItemInt(iItem int, iColumn int, pOutValue *int32) bool {]
ff=取项整数值
pOutValue=值指针
iColumn=列
iItem=项

[func (l *ListBox) GetItemIntEx(iItem int, pName string, pOutValue *int32) bool {]
ff=取项整数值EX
pOutValue=值指针
pName=名称
iItem=项

[func (l *ListBox) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {]
ff=取项浮点值
pOutValue=值指针
iColumn=列
iItem=项

[func (l *ListBox) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {]
ff=取项浮点值EX
pOutValue=值指针
pName=名称
iItem=项

[func (l *ListBox) DeleteItem(iItem int) bool {]
ff=删除项
iItem=项

[func (l *ListBox) DeleteItemEx(iItem int, nCount int) bool {]
ff=删除项EX
nCount=计数
iItem=项

[func (l *ListBox) DeleteItemAll() int {]
ff=删除项全部

[func (l *ListBox) DeleteColumnAll() int {]
ff=删除列全部

[func (l *ListBox) GetCount_AD() int {]
ff=取项数量AD

[func (l *ListBox) GetCountColumn_AD() int {]
ff=取列数量AD

[func (l *ListBox) SetSplitLineColor(color int) int {]
ff=置分割线颜色
color=ABGR颜色值

[func (l *ListBox) SetDragRectColor(color, width int) int {]
ff=置拖动矩形颜色
width=线宽度
color=ABGR颜色值

[func (l *ListBox) SetItemTemplateXMLFromMem(data #左中括号##右中括号#byte) bool {]
ff=置项模板从内存
data=模板数据

[func (l *ListBox) SetItemTemplateXMLFromZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {]
ff=置项模板并按资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=项模板文件名
id=RC资源ID

[func (l *ListBox) GetItemTemplate() int {]
ff=取项模板

[func (l *ListBox) Event_LISTBOX_TEMP_CREATE(pFun XE_LISTBOX_TEMP_CREATE) bool {]
ff=事件_项模板创建

[func (l *ListBox) Event_LISTBOX_TEMP_CREATE1(pFun XE_LISTBOX_TEMP_CREATE1) bool {]
ff=事件_项模板创建1

[func (l *ListBox) Event_LISTBOX_TEMP_CREATE_END(pFun XE_LISTBOX_TEMP_CREATE_END) bool {]
ff=事件_项模板创建完成

[func (l *ListBox) Event_LISTBOX_TEMP_CREATE_END1(pFun XE_LISTBOX_TEMP_CREATE_END1) bool {]
ff=事件_项模板创建完成1

[func (l *ListBox) Event_LISTBOX_TEMP_DESTROY(pFun XE_LISTBOX_TEMP_DESTROY) bool {]
ff=事件_项模板销毁

[func (l *ListBox) Event_LISTBOX_TEMP_DESTROY1(pFun XE_LISTBOX_TEMP_DESTROY1) bool {]
ff=事件_项模板销毁1

[func (l *ListBox) Event_LISTBOX_TEMP_ADJUST_COORDINATE(pFun XE_LISTBOX_TEMP_ADJUST_COORDINATE) bool {]
ff=停用_项模板调整坐标

[func (l *ListBox) Event_LISTBOX_TEMP_ADJUST_COORDINATE1(pFun XE_LISTBOX_TEMP_ADJUST_COORDINATE1) bool {]
ff=停用_项模板调整坐标1

[func (l *ListBox) Event_LISTBOX_DRAWITEM(pFun XE_LISTBOX_DRAWITEM) bool {]
ff=事件_项绘制事件

[func (l *ListBox) Event_LISTBOX_DRAWITEM1(pFun XE_LISTBOX_DRAWITEM1) bool {]
ff=事件_项绘制事件1

[func (l *ListBox) Event_LISTBOX_SELECT(pFun XE_LISTBOX_SELECT) bool {]
ff=事件_项选择事件

[func (l *ListBox) Event_LISTBOX_SELECT1(pFun XE_LISTBOX_SELECT1) bool {]
ff=事件_项选择事件1
