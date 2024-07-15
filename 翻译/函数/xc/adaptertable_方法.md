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

[func XAdTable_Create() int {]
ff=数据适配器表_创建

[func XAdTable_Sort(hAdapter int, iColumn int, bAscending bool) int {]
ff=数据适配器表_排序
bAscending=是否按照升序排序
iColumn=列索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemDataType(hAdapter int, iItem int, iColumn int) xcc.Adapter_Date_Type_ {]
ff=数据适配器表_取项数据类型
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemDataTypeEx(hAdapter int, iItem int, pName string) xcc.Adapter_Date_Type_ {]
ff=数据适配器表_取项数据类型EX
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_AddColumn(hAdapter int, pName string) int {]
ff=数据适配器表_添加列
pName=字段称
hAdapter=数据适配器句柄

[func XAdTable_SetColumn(hAdapter int, pColName string) int {]
ff=数据适配器表_置列
pColName=列名
hAdapter=数据适配器句柄

[func XAdTable_AddItemText(hAdapter int, pValue string) int {]
ff=数据适配器表_添加项文本
pValue=值
hAdapter=数据适配器句柄

[func XAdTable_AddItemTextEx(hAdapter int, pName string, pValue string) int {]
ff=数据适配器表_添加项文本EX
pValue=值
pName=字段称
hAdapter=数据适配器句柄

[func XAdTable_AddItemImage(hAdapter int, hImage int) int {]
ff=数据适配器表_添加项图片
hImage=图片句柄
hAdapter=数据适配器句柄

[func XAdTable_AddItemImageEx(hAdapter int, pName string, hImage int) int {]
ff=数据适配器表_添加项图片EX
hImage=图片句柄
pName=字段称
hAdapter=数据适配器句柄

[func XAdTable_InsertItemText(hAdapter int, iItem int, pValue string) int {]
ff=数据适配器表_插入项文本
pValue=值
iItem=插入位置索引
hAdapter=数据适配器句柄

[func XAdTable_InsertItemTextEx(hAdapter int, iItem int, pName string, pValue string) int {]
ff=数据适配器表_插入项文本EX
pValue=值
pName=字段称
iItem=插入位置索引
hAdapter=数据适配器句柄

[func XAdTable_InsertItemImage(hAdapter int, iItem int, hImage int) int {]
ff=数据适配器表_插入项图片
hImage=图片句柄
iItem=插入位置索引
hAdapter=数据适配器句柄

[func XAdTable_InsertItemImageEx(hAdapter int, iItem int, pName string, hImage int) int {]
ff=数据适配器表_插入项图片EX
hImage=图片句柄
pName=字段称
iItem=插入位置索引
hAdapter=数据适配器句柄

[func XAdTable_SetItemText(hAdapter int, iItem int, iColumn int, pValue string) bool {]
ff=数据适配器表_置项文本
pValue=值
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_SetItemTextEx(hAdapter int, iItem int, pName string, pValue string) bool {]
ff=数据适配器表_置项文本EX
pValue=值
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_SetItemInt(hAdapter int, iItem int, iColumn int, nValue int32) bool {]
ff=数据适配器表_置项整数值
nValue=值
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_SetItemIntEx(hAdapter int, iItem int, pName string, nValue int32) bool {]
ff=数据适配器表_置项整数值EX
nValue=值
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_SetItemFloat(hAdapter int, iItem int, iColumn int, fValue float32) bool {]
ff=数据适配器表_置项浮点值
fValue=值
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_SetItemFloatEx(hAdapter int, iItem int, pName string, fValue float32) bool {]
ff=数据适配器表_置项浮点值EX
fValue=值
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_SetItemImage(hAdapter int, iItem int, iColumn int, hImage int) bool {]
ff=数据适配器表_置项图片
hImage=图片句柄
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_SetItemImageEx(hAdapter int, iItem int, pName string, hImage int) bool {]
ff=数据适配器表_置项图片EX
hImage=图片句柄
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_DeleteItem(hAdapter int, iItem int) bool {]
ff=数据适配器表_删除项
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_DeleteItemEx(hAdapter int, iItem int, nCount int) bool {]
ff=数据适配器表_删除项EX
nCount=删除项数量
iItem=项起始索引
hAdapter=数据适配器句柄

[func XAdTable_DeleteItemAll(hAdapter int) int {]
ff=数据适配器表_删除项全部
hAdapter=数据适配器句柄

[func XAdTable_DeleteColumnAll(hAdapter int) int {]
ff=数据适配器表_删除列全部
hAdapter=数据适配器句柄

[func XAdTable_GetCount(hAdapter int) int {]
ff=数据适配器表_取项数量
hAdapter=数据适配器句柄

[func XAdTable_GetCountColumn(hAdapter int) int {]
ff=数据适配器表_取列数量
hAdapter=数据适配器句柄

[func XAdTable_GetItemText(hAdapter int, iItem int, iColumn int) string {]
ff=数据适配器表_取项文本
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemTextEx(hAdapter int, iItem int, pName string) string {]
ff=数据适配器表_取项文本EX
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemImage(hAdapter int, iItem int, iColumn int) int {]
ff=数据适配器表_取项图片
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemImageEx(hAdapter int, iItem int, pName string) int {]
ff=数据适配器表_取项图片EX
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemInt(hAdapter int, iItem int, iColumn int, pOutValue *int32) bool {]
ff=数据适配器表_取项整数值
pOutValue=接收返还值
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemIntEx(hAdapter int, iItem int, pName string, pOutValue *int32) bool {]
ff=数据适配器表_取项整数值EX
pOutValue=接收返还值
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemFloat(hAdapter int, iItem int, iColumn int, pOutValue *float32) bool {]
ff=数据适配器表_取项浮点值
pOutValue=接收返还值
iColumn=列索引
iItem=项索引
hAdapter=数据适配器句柄

[func XAdTable_GetItemFloatEx(hAdapter int, iItem int, pName string, pOutValue *float32) bool {]
ff=数据适配器表_取项浮点值EX
pOutValue=接收返还值
pName=字段称
iItem=项索引
hAdapter=数据适配器句柄
