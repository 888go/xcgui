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

[func XTemp_Load(nType xcc.ListItemTemp_Type_, pFileName string) int {]
ff=模板_加载从文件
nType=模板类型

[func XTemp_LoadZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) int {]
ff=模板_加载从ZIP
nType=模板类型

[func XTemp_LoadZipMem(nType xcc.ListItemTemp_Type_, data #左中括号##右中括号#byte, pFileName string, pPassword string) int {]
ff=模板_加载从内存ZIP
nType=模板类型

[func XTemp_LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=模板_加载从文件EX
nType=模板类型

[func XTemp_LoadFromMem(nType xcc.ListItemTemp_Type_, data #左中括号##右中括号#byte) int {]
ff=项模板_加载从内存
nType=模板类型

[func XTemp_LoadFromMemEx(nType xcc.ListItemTemp_Type_, data #左中括号##右中括号#byte, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=项模板_加载从内存EX
nType=模板类型

[func XTemp_LoadZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) int {]
ff=项模板_加载从资源ZIP
nType=模板类型

[func XTemp_LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int, hModule uintptr) int {]
ff=项模板_加载从资源ZIPEX
nType=模板类型

[func XTemp_LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=模板_加载从ZIPEX
nType=模板类型

[func XTemp_LoadZipMemEx(nType xcc.ListItemTemp_Type_, data #左中括号##右中括号#byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=模板_加载从内存ZIPEX
nType=模板类型

[func XTemp_LoadFromString(nType xcc.ListItemTemp_Type_, pStringXML string) int {]
ff=模板_加载从字符串
nType=模板类型

[func XTemp_LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=模板_加载从字符串EX
nType=模板类型

[func XTemp_GetType(hTemp int) xcc.ListItemTemp_Type_ {]
ff=模板_取类型
hTemp=列表项模板句柄

[func XTemp_Destroy(hTemp int) bool {]
ff=模板_销毁
hTemp=项模板句柄

[func XTemp_Create(nType xcc.ListItemTemp_Type_) int {]
ff=模板_创建
nType=模板类型

[func XTemp_AddNodeRoot(hTemp int, pNode int) bool {]
ff=模板_添加根节点
pNode=节点指针
hTemp=项模板句柄

[func XTemp_AddNode(pParentNode int, pNode int) bool {]
ff=模板_添加子节点
pNode=节点指针
pParentNode=父节点指针

[func XTemp_CreateNode(nType xcc.XC_OBJECT_TYPE) int {]
ff=模板_创建节点
nType=对象类型

[func XTemp_SetNodeAttribute(pNode int, pName string, pAttr string) bool {]
ff=模板_置节点属性
pAttr=属性值
pName=属性名
pNode=节点指针

[func XTemp_SetNodeAttributeEx(pNode int, itemID int32, pName string, pAttr string) bool {]
ff=模板_置节点属性EX
pAttr=属性值
pName=属性名
itemID=模板项ID
pNode=节点指针

[func XTemp_List_GetNode(hTemp int, index int32) int {]
ff=模板_取列表中的节点
index=节点位置索引
hTemp=模板句柄

[func XTemp_GetNode(pNode int, itemID int32) int {]
ff=模板_取节点
itemID=ID
pNode=节点指针

[func XTemp_CloneNode(pNode int) int {]
ff=模板_克隆节点
pNode=节点指针

[func XTemp_Clone(hTemp int) int {]
ff=项模板_克隆
hTemp=列表项模板句柄

[func XTemp_List_InsertNode(hTemp int, index int32, pNode int) bool {]
ff=项模板_列表_插入节点
pNode=节点指针
index=插入位置索引
hTemp=列表项模板句柄

[func XTemp_List_DeleteNode(hTemp int, index int32) bool {]
ff=项模板_列表_删除节点
index=位置索引
hTemp=列表项模板句柄

[func XTemp_List_GetCount(hTemp int) int32 {]
ff=项模板_列表_取数量
hTemp=列表项模板句柄

[func XTemp_List_MoveColumn(hTemp int, iColSrc, iColDest int32) bool {]
ff=项模板_列表_移动列
iColDest=目标列索引
iColSrc=源列索引
hTemp=列表项模板句柄
