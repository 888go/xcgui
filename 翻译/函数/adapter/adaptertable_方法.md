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

[func NewAdapterTable() *AdapterTable {]
ff=创建适配器表

[func NewAdapterTableByHandle(handle int) *AdapterTable {]
ff=创建适配器表并按句柄

[func (a *AdapterTable) Sort(iColumn int, bAscending bool) int {]
ff=排序
bAscending=升序排序
iColumn=要排序的列索引

[func (a *AdapterTable) GetItemDataType(iItem int, iColumn int) xcc.Adapter_Date_Type_ {]
ff=取项数据类型
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) GetItemDataTypeEx(iItem int, pName string) xcc.Adapter_Date_Type_ {]
ff=取项数据类型EX
pName=字段称
iItem=项索引

[func (a *AdapterTable) AddColumn(pName string) int {]
ff=添加列
pName=字段称

[func (a *AdapterTable) SetColumn(pColName string) int {]
ff=置列
pColName=列名

[func (a *AdapterTable) AddItemText(pValue string) int {]
ff=添加项文本
pValue=值

[func (a *AdapterTable) AddItemTextEx(pName string, pValue string) int {]
ff=添加项文本EX
pValue=值
pName=字段称

[func (a *AdapterTable) AddItemImage(hImage int) int {]
ff=添加项图片
hImage=图片句柄

[func (a *AdapterTable) AddItemImageEx(pName string, hImage int) int {]
ff=添加项图片EX
hImage=图片句柄
pName=字段称

[func (a *AdapterTable) InsertItemText(iItem int, pValue string) int {]
ff=插入项文本
pValue=值
iItem=插入位置索引

[func (a *AdapterTable) InsertItemTextEx(iItem int, pName string, pValue string) int {]
ff=插入项文本EX
pValue=值
pName=字段称
iItem=插入位置索引

[func (a *AdapterTable) InsertItemImage(iItem int, hImage int) int {]
ff=插入项图片
hImage=图片句柄
iItem=插入位置索引

[func (a *AdapterTable) InsertItemImageEx(iItem int, pName string, hImage int) int {]
ff=插入项图片EX
hImage=图片句柄
pName=字段称
iItem=插入位置索引

[func (a *AdapterTable) SetItemText(iItem int, iColumn int, pValue string) bool {]
ff=置项文本
pValue=值
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) SetItemTextEx(iItem int, pName string, pValue string) bool {]
ff=置项文本EX
pValue=值
pName=字段称
iItem=项索引

[func (a *AdapterTable) SetItemInt(iItem int, iColumn int, nValue int32) bool {]
ff=置项整数值
nValue=值
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) SetItemIntEx(iItem int, pName string, nValue int32) bool {]
ff=置项整数值EX
nValue=值
pName=字段称
iItem=项索引

[func (a *AdapterTable) SetItemFloat(iItem int, iColumn int, fValue float32) bool {]
ff=置项浮点值
fValue=值
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) SetItemFloatEx(iItem int, pName string, fValue float32) bool {]
ff=置项浮点值EX
fValue=值
pName=字段称
iItem=项索引

[func (a *AdapterTable) SetItemImage(iItem int, iColumn int, hImage int) bool {]
ff=置项图片
hImage=图片句柄
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) SetItemImageEx(iItem int, pName string, hImage int) bool {]
ff=置项图片EX
hImage=图片句柄
pName=字段称
iItem=项索引

[func (a *AdapterTable) DeleteItem(iItem int) bool {]
ff=删除项
iItem=项索引

[func (a *AdapterTable) DeleteItemEx(iItem int, nCount int) bool {]
ff=删除项EX
nCount=删除项数量
iItem=项起始索引

[func (a *AdapterTable) DeleteItemAll() int {]
ff=删除项全部

[func (a *AdapterTable) DeleteColumnAll() int {]
ff=删除列全部

[func (a *AdapterTable) GetCount() int {]
ff=取项数量

[func (a *AdapterTable) GetCountColumn() int {]
ff=取列数量

[func (a *AdapterTable) GetItemText(iItem int, iColumn int) string {]
ff=取项文本
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) GetItemTextEx(iItem int, pName string) string {]
ff=取项文本EX
pName=字段称
iItem=项索引

[func (a *AdapterTable) GetItemImage(iItem int, iColumn int) int {]
ff=取项图片
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) GetItemImageEx(iItem int, pName string) int {]
ff=取项图片EX
pName=字段称
iItem=项索引

[func (a *AdapterTable) GetItemInt(iItem int, iColumn int, pOutValue *int32) bool {]
ff=取项整数值
pOutValue=接收返还值
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) GetItemIntEx(iItem int, pName string, pOutValue *int32) bool {]
ff=取项整数值EX
pOutValue=接收返还值
pName=字段称
iItem=项索引

[func (a *AdapterTable) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {]
ff=取项浮点值
pOutValue=接收返还值
iColumn=列索引
iItem=项索引

[func (a *AdapterTable) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {]
ff=取项浮点值EX
pOutValue=接收返还值
pName=字段称
iItem=项索引
