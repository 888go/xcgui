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

[func XListBox_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=列表框_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XListBox_CreateEx(x, y, cx, cy int32, hParent, col_extend_count int32) int {]
ff=列表框_创建Ex
col_extend_count=列数量
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XListBox_EnableFixedRowHeight(hEle int, bEnable bool) int {]
ff=列表框_启用固定行高
bEnable=是否启用
hEle=元素句柄

[func XListBox_EnableTemplateReuse(hEle int, bEnable bool) int {]
ff=列表框_启用模板复用
bEnable=是否启用
hEle=元素句柄

[func XListBox_EnableVirtualTable(hEle int, bEnable bool) int {]
ff=列表框_启用虚表
bEnable=是否启用
hEle=元素句柄

[func XListBox_SetVirtualRowCount(hEle int, nRowCount int) int {]
ff=列表框_置虚表行数
nRowCount=行数
hEle=元素句柄

[func XListBox_SetDrawItemBkFlags(hEle int, nFlags xcc.List_DrawItemBk_Flag_) int {]
ff=列表框_置绘制项背景标志
nFlags=标志位
hEle=元素句柄

[func XListBox_SetItemData(hEle int, iItem int, nUserData int) bool {]
ff=列表框_置项数据
nUserData=用户数据
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemData(hEle int, iItem int) int {]
ff=列表框_取项数据
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemInfo(hEle int, iItem int, pItem *ListBox_Item_Info_) bool {]
ff=列表框_置项信息
pItem=项信息
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemInfo(hEle int, iItem int, pItem *ListBox_Item_Info_) bool {]
ff=列表框_取项背景信息
pItem=项信息
iItem=项索引
hEle=元素句柄

[func XListBox_SetSelectItem(hEle int, iItem int) bool {]
ff=列表框_置选择项
iItem=项索引
hEle=元素句柄

[func XListBox_GetSelectItem(hEle int) int {]
ff=列表框_取选择项
hEle=元素句柄

[func XListBox_AddSelectItem(hEle int, iItem int) bool {]
ff=列表框_添加选择项
iItem=项索引
hEle=元素句柄

[func XListBox_CancelSelectItem(hEle int, iItem int) bool {]
ff=列表框_取消选择项
iItem=项索引
hEle=元素句柄

[func XListBox_CancelSelectAll(hEle int) bool {]
ff=列表框_取消选择全部
hEle=元素句柄

[func XListBox_GetSelectAll(hEle int, pArray *#左中括号##右中括号#int32, nArraySize int) int {]
ff=列表框_取全部选择
nArraySize=切片大小
pArray=切片缓冲区
hEle=元素句柄

[func XListBox_GetSelectCount(hEle int) int {]
ff=列表框_取选择项数量
hEle=元素句柄

[func XListBox_GetItemMouseStay(hEle int) int {]
ff=列表框_取鼠标停留项
hEle=元素句柄

[func XListBox_SelectAll(hEle int) bool {]
ff=列表框_选择全部项
hEle=元素句柄

[func XListBox_VisibleItem(hEle int, iItem int) int {]
ff=列表框_显示指定项
iItem=项索引
hEle=元素句柄

[func XListBox_GetVisibleRowRange(hEle int, piStart *int32, piEnd *int32) int {]
ff=列表框_取可视行范围
piEnd=结束行索引
piStart=开始行索引
hEle=元素句柄

[func XListBox_SetItemHeightDefault(hEle int, nHeight, nSelHeight int32) int {]
ff=列表框_置项默认高度
nSelHeight=选中项高度
nHeight=项高度
hEle=元素句柄

[func XListBox_GetItemHeightDefault(hEle int, pHeight, pSelHeight *int32) int {]
ff=列表框_取项默认高度
pSelHeight=选中时高度
pHeight=高度
hEle=元素句柄

[func XListBox_GetItemIndexFromHXCGUI(hEle int, hXCGUI int) int {]
ff=列表框_取所在行索引
hXCGUI=对象句柄
hEle=元素句柄

[func XListBox_SetRowSpace(hEle int, nSpace int) int {]
ff=列表框_置行间距
nSpace=间距大小
hEle=元素句柄

[func XListBox_GetRowSpace(hEle int) int {]
ff=列表框_取行间距
hEle=元素句柄

[func XListBox_HitTest(hEle int, pPt *POINT) int {]
ff=列表框_测试点击项
pPt=坐标点
hEle=元素句柄

[func XListBox_HitTestOffset(hEle int, pPt *POINT) int {]
ff=列表框_测试点击项EX
pPt=坐标点
hEle=元素句柄

[func XListBox_SetItemTemplateXML(hEle int, pXmlFile string) bool {]
ff=列表框_置项模板文件
pXmlFile=文件名
hEle=元素句柄

[func XListBox_SetItemTemplate(hEle int, hTemp int) bool {]
ff=列表框_置项模板
hTemp=模板句柄
hEle=元素句柄

[func XListBox_SetItemTemplateXMLFromString(hEle int, pStringXML string) bool {]
ff=列表框_置项模板从字符串
pStringXML=字符串
hEle=元素句柄

[func XListBox_GetTemplateObject(hEle int, iItem int, nTempItemID int) int {]
ff=列表框_取模板对象
nTempItemID=模板项ID
iItem=项索引
hEle=元素句柄

[func XListBox_EnableMultiSel(hEle int, bEnable bool) int {]
ff=列表框_启用多选
bEnable=是否启用
hEle=元素句柄

[func XListBox_CreateAdapter(hEle int) int {]
ff=列表框_创建数据适配器
hEle=元素句柄

[func XListBox_BindAdapter(hEle int, hAdapter int) int {]
ff=列表框_绑定数据适配器
hAdapter=数据适配器句柄
hEle=元素句柄

[func XListBox_GetAdapter(hEle int) int {]
ff=列表框_取数据适配器
hEle=元素句柄

[func XListBox_Sort(hEle int, iColumnAdapter int, bAscending bool) int {]
ff=列表框_排序
bAscending=升序
iColumnAdapter=数据适配器列索引
hEle=元素句柄

[func XListBox_RefreshData(hEle int) int {]
ff=列表框_刷新数据
hEle=元素句柄

[func XListBox_RefreshItem(hEle int, iItem int) int {]
ff=列表框_刷新指定项
iItem=项索引
hEle=元素句柄

[func XListBox_AddItemText(hEle int, pText string) int {]
ff=列表框_添加项文本
pText=文本
hEle=元素句柄

[func XListBox_AddItemTextEx(hEle int, pName string, pText string) int {]
ff=列表框_添加项文本EX
pText=文本
pName=名称
hEle=元素句柄

[func XListBox_AddItemImage(hEle int, hImage int) int {]
ff=列表框_添加项图片
hImage=图片
hEle=元素句柄

[func XListBox_AddItemImageEx(hEle int, pName string, hImage int) int {]
ff=列表框_添加项图片EX
hImage=图片
pName=名称
hEle=元素句柄

[func XListBox_InsertItemText(hEle int, iItem int, pValue string) int {]
ff=列表框_插入项文本
pValue=值
iItem=项索引
hEle=元素句柄

[func XListBox_InsertItemTextEx(hEle int, iItem int, pName string, pValue string) int {]
ff=列表框_插入项文本EX
pValue=值
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_InsertItemImage(hEle int, iItem int, hImage int) int {]
ff=列表框_插入项图片
hImage=图片
iItem=项索引
hEle=元素句柄

[func XListBox_InsertItemImageEx(hEle int, iItem int, pName string, hImage int) int {]
ff=列表框_插入项图片EX
hImage=图片
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemText(hEle int, iItem int, iColumn int, pText string) bool {]
ff=列表框_置项文本
pText=文本
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemTextEx(hEle int, iItem int, pName string, pText string) bool {]
ff=列表框_置项文本EX
pText=文本
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemImage(hEle int, iItem int, iColumn int, hImage int) bool {]
ff=列表框_置项图片
hImage=图片
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemImageEx(hEle int, iItem int, pName string, hImage int) bool {]
ff=列表框_置项图片EX
hImage=图片
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemInt(hEle int, iItem int, iColumn int, nValue int) bool {]
ff=列表框_置项整数值
nValue=值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemIntEx(hEle int, iItem int, pName string, nValue int) bool {]
ff=列表框_置项整数值EX
nValue=值
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemFloat(hEle int, iItem int, iColumn int, fFloat float32) bool {]
ff=列表框_置项浮点值
fFloat=值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XListBox_SetItemFloatEx(hEle int, iItem int, pName string, fFloat float32) bool {]
ff=列表框_置项浮点值EX
fFloat=值
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemText(hEle int, iItem int, iColumn int) string {]
ff=列表框_取项文本
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemTextEx(hEle int, iItem int, pName string) string {]
ff=列表框_取项文本EX
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemImage(hEle int, iItem int, iColumn int) int {]
ff=列表框_取项图片
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemImageEx(hEle int, iItem int, pName string) int {]
ff=列表框_取项图片EX
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemInt(hEle int, iItem int, iColumn int, pOutValue *int32) bool {]
ff=列表框_取项整数值
pOutValue=接收返还值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemIntEx(hEle int, iItem int, pName string, pOutValue *int32) bool {]
ff=列表框_取项整数值EX
pOutValue=接收返还值
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemFloat(hEle int, iItem int, iColumn int, pOutValue *float32) bool {]
ff=列表框_取项浮点值
pOutValue=接收返还值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XListBox_GetItemFloatEx(hEle int, iItem int, pName string, pOutValue *float32) bool {]
ff=列表框_取项浮点值EX
pOutValue=接收返还值
pName=名称
iItem=项索引
hEle=元素句柄

[func XListBox_DeleteItem(hEle int, iItem int) bool {]
ff=列表框_删除项
iItem=项索引
hEle=元素句柄

[func XListBox_DeleteItemEx(hEle int, iItem int, nCount int) bool {]
ff=列表框_删除项EX
nCount=数量
iItem=项索引
hEle=元素句柄

[func XListBox_DeleteItemAll(hEle int) int {]
ff=列表框_删除项全部
hEle=元素句柄

[func XListBox_DeleteColumnAll(hEle int) int {]
ff=列表框_删除列全部
hEle=元素句柄

[func XListBox_GetCount_AD(hEle int) int {]
ff=列表框_取项数量AD
hEle=元素句柄

[func XListBox_GetCountColumn_AD(hEle int) int {]
ff=列表框_取列数量AD
hEle=元素句柄

[func XListBox_SetSplitLineColor(hEle int, color int) int {]
ff=列表框_置分割线颜色
color=ABGR颜色值
hEle=元素句柄

[func XListBox_SetDragRectColor(hEle int, color, width int) int {]
ff=列表框_置拖动矩形颜色
width=线宽度
color=ABGR颜色值
hEle=元素句柄

[func XListBox_SetItemTemplateXMLFromMem(hEle int, data #左中括号##右中括号#byte) bool {]
ff=列表框_置项模板从内存
data=模板数据
hEle=元素句柄

[func XListBox_SetItemTemplateXMLFromZipRes(hEle, id int, pFileName string, pPassword string, hModule uintptr) bool {]
ff=列表框_置项模板从资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=项模板文件名
id=RC资源ID
hEle=元素句柄

[func XListBox_GetItemTemplate(hEle int) int {]
ff=列表框_取项模板
hEle=元素句柄
