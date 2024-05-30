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

[func New(nType xcc.ListItemTemp_Type_) *ListItemTemplate {]
ff=创建
nType=模板类型

[func NewByHandle(handle int) *ListItemTemplate {]
ff=创建并按句柄
handle=句柄

[func NewByXML(nType xcc.ListItemTemp_Type_, pFileName string) *ListItemTemplate {]
ff=创建并按文件
nType=模板类型

[func NewByZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) *ListItemTemplate {]
ff=创建并按ZIP
nType=模板类型

[func NewByZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) *ListItemTemplate {]
ff=创建并按RC资源ZIP
nType=模板类型

[func NewByZipMem(nType xcc.ListItemTemp_Type_, data #左中括号##右中括号#byte, pFileName string, pPassword string) *ListItemTemplate {]
ff=创建并按内存ZIP
nType=模板类型

[func NewByXmlMem(nType xcc.ListItemTemp_Type_, data #左中括号##右中括号#byte) *ListItemTemplate {]
ff=创建并按内存
nType=模板类型

[func NewByString(nType xcc.ListItemTemp_Type_, pStringXML string) *ListItemTemplate {]
ff=创建并按字符串
nType=模板类型

[func (l *ListItemTemplate) Clone() *ListItemTemplate {]
ff=取副本

[func (l *ListItemTemplate) GetType() xcc.ListItemTemp_Type_ {]
ff=取类型

[func (l *ListItemTemplate) Destroy() bool {]
ff=销毁

[func (l *ListItemTemplate) AddNodeRoot(pNode int) bool {]
ff=添加根节点
pNode=节点指针

[func (l *ListItemTemplate) List_GetNode(index int32) int {]
ff=取列表中节点
index=节点位置索引

[func (l *ListItemTemplate) List_InsertNode(index int32, pNode int) bool {]
ff=列表插入节点
pNode=节点指针
index=插入位置索引

[func (l *ListItemTemplate) List_DeleteNode(index int32) bool {]
ff=列表删除节点
index=位置索引

[func (l *ListItemTemplate) List_GetCount() int32 {]
ff=列表取数量

[func (l *ListItemTemplate) List_MoveColumn(iColSrc, iColDest int32) bool {]
ff=列表移动列
iColDest=目标列索引
iColSrc=源列索引

[func LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=加载并按文件EX
nType=模板类型

[func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=加载并按ZIPEX
nType=模板类型

[func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data #左中括号##右中括号#byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=加载并按内存ZIPEX
nType=模板类型

[func LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=加载并按字符串EX
nType=模板类型

[func LoadFromMemEx(nType xcc.ListItemTemp_Type_, data #左中括号##右中括号#byte, pOutTemp1 *int, pOutTemp2 *int) bool {]
ff=加载并按内存EX
nType=模板类型

[func LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int, hModule uintptr) int {]
ff=加载并按资源ZIPEX
nType=模板类型
