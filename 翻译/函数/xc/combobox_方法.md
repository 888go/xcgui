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

[func XComboBox_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=组合框_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func XComboBox_SetSelItem(hEle int, iIndex int) bool {]
ff=组合框_置选择项
iIndex=项索引
hEle=元素句柄

[func XComboBox_CreateAdapter(hEle int) int {]
ff=组合框_创建数据适配器
hEle=元素句柄

[func XComboBox_BindAdapter(hEle int, hAdapter int) int {]
ff=组合框_绑定数据适配器
hAdapter=适配器句柄
hEle=元素句柄

[func XComboBox_GetAdapter(hEle int) int {]
ff=组合框_取数据适配器
hEle=元素句柄

[func XComboBox_SetBindName(hEle int, pName string) int {]
ff=组合框_置绑定名称
pName=字段名
hEle=元素句柄

[func XComboBox_GetButtonRect(hEle int, pRect *RECT) int {]
ff=组合框_取下拉按钮坐标
pRect=坐标
hEle=元素句柄

[func XComboBox_SetButtonSize(hEle int, size int) int {]
ff=组合框_置下拉按钮大小
size=大小
hEle=元素句柄

[func XComboBox_SetDropHeight(hEle int, height int) int {]
ff=组合框_置下拉列表高度
height=高度
hEle=元素句柄

[func XComboBox_GetDropHeight(hEle int) int {]
ff=组合框_取下拉列表高度
hEle=元素句柄

[func XComboBox_SetItemTemplateXML(hEle int, pXmlFile string) int {]
ff=组合框_置项模板
pXmlFile=项模板文件
hEle=元素句柄

[func XComboBox_SetItemTemplateXMLFromString(hEle int, pStringXML string) int {]
ff=组合框_置项模板从字符串
pStringXML=字符串
hEle=元素句柄

[func XComboBox_EnableDrawButton(hEle int, bEnable bool) int {]
ff=组合框_启用绘制下拉按钮
bEnable=是否绘制
hEle=元素句柄

[func XComboBox_EnableEdit(hEle int, bEdit bool) int {]
ff=组合框_启用编辑
bEdit=TRUE可编辑
hEle=元素句柄

[func XComboBox_EnableDropHeightFixed(hEle int, bEnable bool) int {]
ff=组合框_启用下拉列表高度固定大小
bEnable=是否启用
hEle=元素句柄

[func XComboBox_GetSelItem(hEle int) int {]
ff=组合框_取选择项
hEle=元素句柄

[func XComboBox_GetState(hEle int) xcc.ComboBox_State_ {]
ff=组合框_取状态
hEle=元素句柄

[func XComboBox_AddItemText(hEle int, pText string) int {]
ff=组合框_添加项文本
pText=文本
hEle=元素句柄

[func XComboBox_AddItemTextEx(hEle int, pName string, pText string) int {]
ff=组合框_添加项文本EX
pText=文本
pName=字段名
hEle=元素句柄

[func XComboBox_AddItemImage(hEle int, hImage int) int {]
ff=组合框_添加项图片
hImage=图片句柄
hEle=元素句柄

[func XComboBox_AddItemImageEx(hEle int, pName string, hImage int) int {]
ff=组合框_添加项图片EX
hImage=图片句柄
pName=字段名
hEle=元素句柄

[func XComboBox_InsertItemText(hEle int, iItem int, pText string) int {]
ff=组合框_插入项文本
pText=文本
iItem=项索引
hEle=元素句柄

[func XComboBox_InsertItemTextEx(hEle int, iItem int, pName string, pText string) int {]
ff=组合框_插入项文本EX
pText=文本
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_InsertItemImage(hEle int, iItem int, hImage int) int {]
ff=组合框_插入项图片
hImage=图片句柄
iItem=项索引
hEle=元素句柄

[func XComboBox_InsertItemImageEx(hEle int, iItem int, pName string, hImage int) int {]
ff=组合框_插入项图片EX
hImage=图片句柄
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_SetItemText(hEle int, iItem int, iColumn int, pText string) bool {]
ff=组合框_置项文本
pText=文本
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XComboBox_SetItemTextEx(hEle int, iItem int, pName string, pText string) bool {]
ff=组合框_置项文本EX
pText=文本
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_SetItemImage(hEle int, iItem int, iColumn int, hImage int) bool {]
ff=组合框_置项图片
hImage=图片句柄
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XComboBox_SetItemImageEx(hEle int, iItem int, pName string, hImage int) bool {]
ff=组合框_置项图片EX
hImage=图片句柄
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_SetItemInt(hEle int, iItem int, iColumn int, nValue int32) bool {]
ff=组合框_置项整数值
nValue=整数值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XComboBox_SetItemIntEx(hEle int, iItem int, pName string, nValue int32) bool {]
ff=组合框_置项指数值EX
nValue=整数值
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_SetItemFloat(hEle int, iItem int, iColumn int, fFloat float32) bool {]
ff=组合框_置项浮点值
fFloat=浮点数
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XComboBox_SetItemFloatEx(hEle int, iItem int, pName string, fFloat float32) bool {]
ff=组合框_置项浮点值EX
fFloat=浮点数
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_GetItemText(hEle int, iItem int32, iColumn int32) string {]
ff=组合框_取项文本
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XComboBox_GetItemTextEx(hEle int, iItem int, pName string) string {]
ff=组合框_取项文本EX
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_GetItemImage(hEle int, iItem int, iColumn int) int {]
ff=组合框_取项图片
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XComboBox_GetItemImageEx(hEle int, iItem int, pName string) int {]
ff=组合框_取项图片EX
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_GetItemInt(hEle int, iItem int, iColumn int, pOutValue *int32) bool {]
ff=组合框_取项整数值
pOutValue=接收返回整数值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XComboBox_GetItemIntEx(hEle int, iItem int, pName string, pOutValue *int32) bool {]
ff=组合框_取项整数值EX
pOutValue=接收返回整数值
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_GetItemFloat(hEle int, iItem int, iColumn int, pOutValue *float32) bool {]
ff=组合框_取项浮点值
pOutValue=接收返回浮点值
iColumn=列索引
iItem=项索引
hEle=元素句柄

[func XComboBox_GetItemFloatEx(hEle int, iItem int, pName string, pOutValue *float32) bool {]
ff=组合框_取项浮点值EX
pOutValue=接收返回浮点值
pName=字段名
iItem=项索引
hEle=元素句柄

[func XComboBox_DeleteItem(hEle int, iItem int) bool {]
ff=组合框_删除项
iItem=项索引
hEle=元素句柄

[func XComboBox_DeleteItemEx(hEle int, iItem int, nCount int) bool {]
ff=组合框_删除项EX
nCount=删除数量
iItem=项索引
hEle=元素句柄

[func XComboBox_DeleteItemAll(hEle int) int {]
ff=组合框_删除项全部
hEle=元素句柄

[func XComboBox_DeleteColumnAll(hEle int) int {]
ff=组合框_删除列全部
hEle=元素句柄

[func XComboBox_GetCount(hEle int) int {]
ff=组合框_取项数量

[func XComboBox_GetCountColumn(hEle int) int {]
ff=组合框_取列数量
hEle=元素句柄

[func XComboBox_PopupDropList(hEle int) int {]
ff=组合框_弹出下拉列表
hEle=元素句柄

[func XComboBox_SetItemTemplate(hEle, hTemp int) int {]
ff=组合框_设置项模板
hTemp=模板句柄
hEle=元素句柄

[func XComboBox_SetItemTemplateXMLFromMem(hEle int, data #左中括号##右中括号#byte) bool {]
ff=组合框_置项模板从内存
data=模板数据
hEle=元素句柄

[func XComboBox_SetItemTemplateXMLFromZipRes(hEle int, id int32, pFileName string, pPassword string, hModule uintptr) bool {]
ff=组合框_置项模板从资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=文件名
id=RC资源ID
hEle=元素句柄

[func XComboBox_GetItemTemplate(hEle int) int {]
ff=组合框_取项模板
hEle=元素句柄
