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

[func XAdListView_Create() int {]
ff=数据适配器列表视_创建

[func XAdListView_Group_AddColumn(hAdapter int, pName string) int {]
ff=数据适配器列表视_组添加列
pName=字段称
hAdapter=数据适配器句柄

[func XAdListView_Group_AddItemText(hAdapter int, pValue string, iPos int) int {]
ff=数据适配器列表视_添加组文本
iPos=插入位置
pValue=值
hAdapter=数据适配器句柄

[func XAdListView_Group_AddItemTextEx(hAdapter int, pName string, pValue string, iPos int) int {]
ff=数据适配器列表视_添加组文本EX
iPos=插入位置
pValue=值
pName=字段称
hAdapter=数据适配器句柄

[func XAdListView_Group_AddItemImage(hAdapter int, hImage int, iPos int) int {]
ff=数据适配器列表视_添加组图片
iPos=插入位置
hImage=图片句柄
hAdapter=数据适配器句柄

[func XAdListView_Group_AddItemImageEx(hAdapter int, pName string, hImage int, iPos int) int {]
ff=数据适配器列表视_添加组图片EX
iPos=插入位置
hImage=图片句柄
pName=字段称
hAdapter=数据适配器句柄

[func XAdListView_Group_SetText(hAdapter int, iGroup int, iColumn int, pValue string) bool {]
ff=数据适配器列表视_组设置文本
pValue=值
iColumn=列索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_SetTextEx(hAdapter int, iGroup int, pName string, pValue string) bool {]
ff=数据适配器列表视_组设置文本EX
pValue=值
pName=字段名
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_SetImage(hAdapter int, iGroup int, iColumn int, hImage int) bool {]
ff=数据适配器列表视_组设置图片
hImage=图片句柄
iColumn=列索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_SetImageEx(hAdapter int, iGroup int, pName string, hImage int) bool {]
ff=数据适配器列表视_组设置图片EX
hImage=图片句柄
pName=字段名
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_AddColumn(hAdapter int, pName string) int {]
ff=数据适配器列表视_项添加列
pName=字段称
hAdapter=数据适配器句柄

[func XAdListView_Group_GetCount(hAdapter int) int {]
ff=数据适配器列表视_取组数量
hAdapter=数据适配器句柄

[func XAdListView_Item_GetCount(hAdapter int, iGroup int) int {]
ff=数据适配器列表视_取组中项数量
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_AddItemText(hAdapter int, iGroup int, pValue string, iPos int) int {]
ff=数据适配器列表视_添加项文本
iPos=插入位置
pValue=值
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_AddItemTextEx(hAdapter int, iGroup int, pName string, pValue string, iPos int) int {]
ff=数据适配器列表视_添加项文本EX
iPos=插入位置
pValue=值
pName=字段称
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_AddItemImage(hAdapter int, iGroup int, hImage int, iPos int) int {]
ff=数据适配器列表视_添加项图片
iPos=插入位置
hImage=图片句柄
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_AddItemImageEx(hAdapter int, iGroup int, pName string, hImage int, iPos int) int {]
ff=数据适配器列表视_添加项图片EX
iPos=插入位置
hImage=图片句柄
pName=字段称
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_DeleteItem(hAdapter int, iGroup int) bool {]
ff=数据适配器列表视_组删除项
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_DeleteAllChildItem(hAdapter int, iGroup int) int {]
ff=数据适配器列表视_删除全部子项
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_DeleteItem(hAdapter int, iGroup int, iItem int) bool {]
ff=数据适配器列表视_删除项
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_DeleteAll(hAdapter int) int {]
ff=数据适配器列表视_删除全部
hAdapter=数据适配器句柄

[func XAdListView_DeleteAllGroup(hAdapter int) int {]
ff=数据适配器列表视_删除全部组
hAdapter=数据适配器句柄

[func XAdListView_DeleteAllItem(hAdapter int) int {]
ff=数据适配器列表视_删除全部项
hAdapter=数据适配器句柄

[func XAdListView_DeleteColumnGroup(hAdapter int, iColumn int) int {]
ff=数据适配器列表视_组删除列
iColumn=列索引
hAdapter=数据适配器句柄

[func XAdListView_DeleteColumnItem(hAdapter int, iColumn int) int {]
ff=数据适配器列表视_项删除列
iColumn=列索引
hAdapter=数据适配器句柄

[func XAdListView_Item_GetTextEx(hAdapter int, iGroup int, iItem int, pName string) string {]
ff=数据适配器列表视_项获取文本EX
pName=字段称
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_GetImageEx(hAdapter int, iGroup int, iItem int, pName string) int {]
ff=数据适配器列表视_项获取图片EX
pName=字段称
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_SetText(hAdapter int, iGroup int, iItem int, iColumn int, pValue string) bool {]
ff=数据适配器列表视_项置文本
pValue=值
iColumn=列索引
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_SetTextEx(hAdapter int, iGroup int, iItem int, pName string, pValue string) bool {]
ff=数据适配器列表视_项置文本EX
pValue=值
pName=字段称
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_SetImage(hAdapter int, iGroup int, iItem int, iColumn int, hImage int) bool {]
ff=数据适配器列表视_项置图片
hImage=图片句柄
iColumn=列索引
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_SetImageEx(hAdapter int, iGroup int, iItem int, pName string, hImage int) bool {]
ff=数据适配器列表视_项置图片EX
hImage=图片句柄
pName=字段称
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_GetText(hAdapter int, iGroup int, iColumn int) string {]
ff=数据适配器列表视_组取文本
iColumn=列索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_GetTextEx(hAdapter int, iGroup int, pName string) string {]
ff=数据适配器列表视_组取文本EX
pName=字段名称
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_GetImage(hAdapter int, iGroup int, iColumn int) int {]
ff=数据适配器列表视_组取图片
iColumn=列索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Group_GetImageEx(hAdapter int, iGroup int, pName string) int {]
ff=数据适配器列表视_组取图片EX
pName=字段名称
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_GetText(hAdapter int, iGroup int, iItem int, iColumn int) string {]
ff=数据适配器列表视_项取文本
iColumn=列索引
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄

[func XAdListView_Item_GetImage(hAdapter int, iGroup int, iItem int, iColumn int) int {]
ff=数据适配器列表视_项取图片
iColumn=列索引
iItem=项索引
iGroup=组索引
hAdapter=数据适配器句柄
