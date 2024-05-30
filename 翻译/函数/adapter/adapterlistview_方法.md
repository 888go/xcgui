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

[func NewAdapterListView() *AdapterListView {]
ff=创建适配器列表视

[func NewAdapterListViewByHandle(handle int) *AdapterListView {]
ff=创建适配器列表视并按句柄

[func (a *AdapterListView) Group_AddColumn(pName string) int {]
ff=组添加列
pName=字段称

[func (a *AdapterListView) Group_AddItemText(pValue string, iPos int) int {]
ff=添加组文本
iPos=插入位置
pValue=值

[func (a *AdapterListView) Group_AddItemTextEx(pName string, pValue string, iPos int) int {]
ff=添加组文本EX
iPos=插入位置
pValue=值
pName=字段称

[func (a *AdapterListView) Group_AddItemImage(hImage int, iPos int) int {]
ff=添加组图片
iPos=插入位置
hImage=图片句柄

[func (a *AdapterListView) Group_AddItemImageEx(pName string, hImage int, iPos int) int {]
ff=添加组图片EX
iPos=插入位置
hImage=图片句柄
pName=字段称

[func (a *AdapterListView) Group_SetText(iGroup int, iColumn int, pValue string) bool {]
ff=组设置文本
pValue=值
iColumn=列索引
iGroup=组索引

[func (a *AdapterListView) Group_SetTextEx(iGroup int, pName string, pValue string) bool {]
ff=组设置文本EX
pValue=值
pName=字段名
iGroup=组索引

[func (a *AdapterListView) Group_SetImage(iGroup int, iColumn int, hImage int) bool {]
ff=组设置图片
hImage=图片句柄
iColumn=列索引
iGroup=组索引

[func (a *AdapterListView) Group_SetImageEx(iGroup int, pName string, hImage int) bool {]
ff=组设置图片EX
hImage=图片句柄
pName=字段名
iGroup=组索引

[func (a *AdapterListView) Item_AddColumn(pName string) int {]
ff=项添加列
pName=字段称

[func (a *AdapterListView) Group_GetCount() int {]
ff=取组数量

[func (a *AdapterListView) Item_GetCount(iGroup int) int {]
ff=取组中项数量
iGroup=组索引

[func (a *AdapterListView) Item_AddItemText(iGroup int, pValue string, iPos int) int {]
ff=添加项文本
iPos=插入位置
pValue=值
iGroup=组索引

[func (a *AdapterListView) Item_AddItemTextEx(iGroup int, pName string, pValue string, iPos int) int {]
ff=添加项文本EX
iPos=插入位置
pValue=值
pName=字段称
iGroup=组索引

[func (a *AdapterListView) Item_AddItemImage(iGroup int, hImage int, iPos int) int {]
ff=添加项图片
iPos=插入位置
hImage=图片句柄
iGroup=组索引

[func (a *AdapterListView) Item_AddItemImageEx(iGroup int, pName string, hImage int, iPos int) int {]
ff=添加项图片EX
iPos=插入位置
hImage=图片句柄
pName=字段称
iGroup=组索引

[func (a *AdapterListView) Group_DeleteItem(iGroup int) bool {]
ff=组删除项
iGroup=组索引

[func (a *AdapterListView) Group_DeleteAllChildItem(iGroup int) int {]
ff=删除全部子项
iGroup=组索引

[func (a *AdapterListView) Item_DeleteItem(iGroup int, iItem int) bool {]
ff=删除项
iItem=项索引
iGroup=组索引

[func (a *AdapterListView) DeleteAll() int {]
ff=删除全部

[func (a *AdapterListView) DeleteAllGroup() int {]
ff=删除全部组

[func (a *AdapterListView) DeleteAllItem() int {]
ff=删除全部项

[func (a *AdapterListView) DeleteColumnGroup(iColumn int) int {]
ff=组删除列
iColumn=列索引

[func (a *AdapterListView) DeleteColumnItem(iColumn int) int {]
ff=项删除列
iColumn=列索引

[func (a *AdapterListView) Item_GetTextEx(iGroup int, iItem int, pName string) string {]
ff=项获取文本EX
pName=字段称
iItem=项索引
iGroup=组索引

[func (a *AdapterListView) Item_GetImageEx(iGroup int, iItem int, pName string) int {]
ff=项获取图片EX
pName=字段称
iItem=项索引
iGroup=组索引

[func (a *AdapterListView) Item_SetText(iGroup int, iItem int, iColumn int, pValue string) bool {]
ff=项置文本
pValue=值
iColumn=列索引
iItem=项索引
iGroup=组索引

[func (a *AdapterListView) Item_SetTextEx(iGroup int, iItem int, pName string, pValue string) bool {]
ff=项置文本EX
pValue=值
pName=字段称
iItem=项索引
iGroup=组索引

[func (a *AdapterListView) Item_SetImage(iGroup int, iItem int, iColumn int, hImage int) bool {]
ff=项置图片
hImage=图片句柄
iColumn=列索引
iItem=项索引
iGroup=组索引

[func (a *AdapterListView) Item_SetImageEx(iGroup int, iItem int, pName string, hImage int) bool {]
ff=项置图片EX
hImage=图片句柄
pName=字段称
iItem=项索引
iGroup=组索引

[func (a *AdapterListView) Group_GetText(iGroup int, iColumn int) string {]
ff=组取文本
iColumn=列索引
iGroup=组索引

[func (a *AdapterListView) Group_GetTextEx(iGroup int, pName string) string {]
ff=组取文本EX
pName=字段名称
iGroup=组索引

[func (a *AdapterListView) Group_GetImage(iGroup int, iColumn int) int {]
ff=组取图片
iColumn=列索引
iGroup=组索引

[func (a *AdapterListView) Group_GetImageEx(iGroup int, pName string) int {]
ff=组取图片EX
pName=字段名称
iGroup=组索引

[func (a *AdapterListView) Item_GetText(iGroup int, iItem int, iColumn int) string {]
ff=项取文本
iColumn=列索引
iItem=项索引
iGroup=组索引

[func (a *AdapterListView) Item_GetImage(iGroup int, iItem int, iColumn int) int {]
ff=项取图片
iColumn=列索引
iItem=项索引
iGroup=组索引
