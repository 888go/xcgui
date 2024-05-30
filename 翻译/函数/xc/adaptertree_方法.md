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

[func XAdTree_Create() int {]
ff=数据适配器树_创建

[func XAdTree_AddColumn(hAdapter int, pName string) int {]
ff=数据适配器树_添加列
pName=字段称
hAdapter=数据适配器句柄

[func XAdTree_SetColumn(hAdapter int, pColName string) int {]
ff=数据适配器树_置列
pColName=列名
hAdapter=数据适配器句柄

[func XAdTree_InsertItemText(hAdapter int, pValue string, nParentID int, insertID int) int {]
ff=数据适配器树_插入项文本
insertID=插入位置ID
nParentID=父ID
pValue=值
hAdapter=数据适配器句柄

[func XAdTree_InsertItemTextEx(hAdapter int, pName string, pValue string, nParentID int, insertID int) int {]
ff=数据适配器树_插入项文本EX
insertID=插入位置ID
nParentID=父ID
pValue=值
pName=字段称
hAdapter=数据适配器句柄

[func XAdTree_InsertItemImage(hAdapter int, hImage int, nParentID int, insertID int) int {]
ff=数据适配器树_插入项图片
insertID=插入位置ID
nParentID=父ID
hImage=图片句柄
hAdapter=数据适配器句柄

[func XAdTree_InsertItemImageEx(hAdapter int, pName string, hImage int, nParentID int, insertID int) int {]
ff=数据适配器树_插入项图片EX
insertID=插入位置ID
nParentID=父ID
hImage=图片句柄
pName=字段称
hAdapter=数据适配器句柄

[func XAdTree_GetCount(hAdapter int) int {]
ff=数据适配器树_取项数量
hAdapter=数据适配器句柄

[func XAdTree_GetCountColumn(hAdapter int) int {]
ff=数据适配器树_取列数量
hAdapter=数据适配器句柄

[func XAdTree_SetItemText(hAdapter int, nID int, iColumn int, pValue string) bool {]
ff=数据适配器树_置项文本
pValue=值
iColumn=列索引
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_SetItemTextEx(hAdapter int, nID int, pName string, pValue string) bool {]
ff=数据适配器树_置项文本EX
pValue=值
pName=字段称
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_SetItemImage(hAdapter int, nID int, iColumn int, hImage int) bool {]
ff=数据适配器树_置项图片
hImage=图片句柄
iColumn=列索引
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_SetItemImageEx(hAdapter int, nID int, pName string, hImage int) bool {]
ff=数据适配器树_置项图片EX
hImage=图片句柄
pName=字段称
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_GetItemText(hAdapter int, nID int, iColumn int) string {]
ff=数据适配器树_取项文本
iColumn=列索引
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_GetItemTextEx(hAdapter int, nID int, pName string) string {]
ff=数据适配器树_取项文本EX
pName=字段称
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_GetItemImage(hAdapter int, nID int, iColumn int) int {]
ff=数据适配器树_取项图片
iColumn=列索引
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_GetItemImageEx(hAdapter int, nID int, pName string) int {]
ff=数据适配器树_取项图片EX
pName=字段称
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_DeleteItem(hAdapter int, nID int) bool {]
ff=数据适配器树_删除项
nID=项ID
hAdapter=数据适配器句柄

[func XAdTree_DeleteItemAll(hAdapter int) int {]
ff=数据适配器树_删除项全部
hAdapter=数据适配器句柄

[func XAdTree_DeleteColumnAll(hAdapter int) int {]
ff=数据适配器树_删除列全部
hAdapter=数据适配器句柄
