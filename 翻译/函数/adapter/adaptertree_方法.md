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

[func NewAdapterTree() *AdapterTree {]
ff=创建适配器树

[func NewAdapterTreeByHandle(handle int) *AdapterTree {]
ff=创建适配器树并按句柄

[func (a *AdapterTree) AddColumn(pName string) int {]
ff=添加列
pName=字段称

[func (a *AdapterTree) SetColumn(pColName string) int {]
ff=置列
pColName=列名

[func (a *AdapterTree) InsertItemText(pValue string, nParentID int, insertID int) int {]
ff=插入项文本
insertID=插入位置ID
nParentID=父ID
pValue=值

[func (a *AdapterTree) InsertItemTextEx(pName string, pValue string, nParentID int, insertID int) int {]
ff=插入项文本EX
insertID=插入位置ID
nParentID=父ID
pValue=值
pName=字段称

[func (a *AdapterTree) InsertItemImage(hImage int, nParentID int, insertID int) int {]
ff=插入项图片
insertID=插入位置ID
nParentID=父ID
hImage=图片句柄

[func (a *AdapterTree) InsertItemImageEx(pName string, hImage int, nParentID int, insertID int) int {]
ff=插入项图片EX
insertID=插入位置ID
nParentID=父ID
hImage=图片句柄
pName=字段称

[func (a *AdapterTree) GetCount() int {]
ff=取项数量

[func (a *AdapterTree) GetCountColumn() int {]
ff=取列数量

[func (a *AdapterTree) SetItemText(nID int, iColumn int, pValue string) bool {]
ff=置项文本
pValue=值
iColumn=列索引
nID=项ID

[func (a *AdapterTree) SetItemTextEx(nID int, pName string, pValue string) bool {]
ff=置项文本EX
pValue=值
pName=字段称
nID=项ID

[func (a *AdapterTree) SetItemImage(nID int, iColumn int, hImage int) bool {]
ff=置项图片
hImage=图片句柄
iColumn=列索引
nID=项ID

[func (a *AdapterTree) SetItemImageEx(nID int, pName string, hImage int) bool {]
ff=置项图片EX
hImage=图片句柄
pName=字段称
nID=项ID

[func (a *AdapterTree) GetItemText(nID int, iColumn int) string {]
ff=取项文本
iColumn=列索引
nID=项ID

[func (a *AdapterTree) GetItemTextEx(nID int, pName string) string {]
ff=取项文本EX
pName=字段称
nID=项ID

[func (a *AdapterTree) GetItemImage(nID int, iColumn int) int {]
ff=取项图片
iColumn=列索引
nID=项ID

[func (a *AdapterTree) GetItemImageEx(nID int, pName string) int {]
ff=取项图片EX
pName=字段称
nID=项ID

[func (a *AdapterTree) DeleteItem(nID int) bool {]
ff=删除项
nID=项ID

[func (a *AdapterTree) DeleteItemAll() int {]
ff=删除项全部

[func (a *AdapterTree) DeleteColumnAll() int {]
ff=删除列全部
