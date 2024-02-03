
# <翻译开始>
func New(nType
模板类型
# <翻译结束>

# <翻译开始>
func New
X创建
# <翻译结束>


# <翻译开始>
func NewByHandle
X创建并按句柄
# <翻译结束>


# <翻译开始>
func NewByXML(nType xcc.ListItemTemp_Type_, pFileName
文件名
# <翻译结束>

# <翻译开始>
func NewByXML(nType
模板类型
# <翻译结束>

# <翻译开始>
func NewByXML
X创建并按文件
# <翻译结束>


# <翻译开始>
func NewByZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func NewByZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName
文件名
# <翻译结束>

# <翻译开始>
func NewByZip(nType xcc.ListItemTemp_Type_, pZipFile
zip文件
# <翻译结束>

# <翻译开始>
func NewByZip(nType
模板类型
# <翻译结束>

# <翻译开始>
func NewByZip
X创建并按ZIP
# <翻译结束>


# <翻译开始>
func NewByZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule
模块句柄
# <翻译结束>

# <翻译开始>
func NewByZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func NewByZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName
模板文件名
# <翻译结束>

# <翻译开始>
func NewByZipRes(nType xcc.ListItemTemp_Type_, id
RC资源ID
# <翻译结束>

# <翻译开始>
func NewByZipRes(nType
模板类型
# <翻译结束>

# <翻译开始>
func NewByZipRes
X创建并按ZIP
# <翻译结束>


# <翻译开始>
func NewByZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func NewByZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName
文件名
# <翻译结束>

# <翻译开始>
func NewByZipMem(nType xcc.ListItemTemp_Type_, data
模板文件数据
# <翻译结束>

# <翻译开始>
func NewByZipMem(nType
模板类型
# <翻译结束>

# <翻译开始>
func NewByZipMem
X创建并按内存ZIP
# <翻译结束>


# <翻译开始>
func NewByXmlMem(nType xcc.ListItemTemp_Type_, data
模板文件数据
# <翻译结束>

# <翻译开始>
func NewByXmlMem(nType
模板类型
# <翻译结束>

# <翻译开始>
func NewByXmlMem
X创建并按内存
# <翻译结束>


# <翻译开始>
func NewByString(nType xcc.ListItemTemp_Type_, pStringXML
字符串
# <翻译结束>

# <翻译开始>
func NewByString(nType
模板类型
# <翻译结束>

# <翻译开始>
func NewByString
X创建并按字符串
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) Clone
X取副本
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) GetType
X取类型
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) Destroy
X销毁
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) AddNodeRoot(pNode
节点指针
# <翻译结束>

# <翻译开始>
func (l *ListItemTemplate) AddNodeRoot
X添加根节点
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) List_GetNode(index
节点位置索引
# <翻译结束>

# <翻译开始>
func (l *ListItemTemplate) List_GetNode
X取列表中节点
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) List_InsertNode(index int32, pNode
节点指针
# <翻译结束>

# <翻译开始>
func (l *ListItemTemplate) List_InsertNode(index
插入位置索引
# <翻译结束>

# <翻译开始>
func (l *ListItemTemplate) List_InsertNode
X列表插入节点
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) List_DeleteNode(index
位置索引
# <翻译结束>

# <翻译开始>
func (l *ListItemTemplate) List_DeleteNode
X列表删除节点
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) List_GetCount
X列表取数量
# <翻译结束>


# <翻译开始>
func (l *ListItemTemplate) List_MoveColumn(iColSrc, iColDest
目标列索引
# <翻译结束>

# <翻译开始>
func (l *ListItemTemplate) List_MoveColumn(iColSrc
源列索引
# <翻译结束>

# <翻译开始>
func (l *ListItemTemplate) List_MoveColumn
X列表移动列
# <翻译结束>


# <翻译开始>
func LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2
返回模板句柄2
# <翻译结束>

# <翻译开始>
func LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1
返回模板句柄1
# <翻译结束>

# <翻译开始>
func LoadEx(nType xcc.ListItemTemp_Type_, pFileName
文件名
# <翻译结束>

# <翻译开始>
func LoadEx(nType
模板类型
# <翻译结束>

# <翻译开始>
func LoadEx
X加载并按文件扩展
# <翻译结束>


# <翻译开始>
func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2
返回模板句柄2
# <翻译结束>

# <翻译开始>
func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1
返回模板句柄1
# <翻译结束>

# <翻译开始>
func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName
文件名
# <翻译结束>

# <翻译开始>
func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile
zip文件
# <翻译结束>

# <翻译开始>
func LoadZipEx(nType
模板类型
# <翻译结束>

# <翻译开始>
func LoadZipEx
X加载并按ZIP扩展
# <翻译结束>


# <翻译开始>
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2
返回模板句柄2
# <翻译结束>

# <翻译开始>
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1
返回模板句柄1
# <翻译结束>

# <翻译开始>
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName
文件名
# <翻译结束>

# <翻译开始>
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data
模板文件数据
# <翻译结束>

# <翻译开始>
func LoadZipMemEx(nType
模板类型
# <翻译结束>

# <翻译开始>
func LoadZipMemEx
X加载并按内存ZIP扩展
# <翻译结束>


# <翻译开始>
func LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2
返回模板句柄2
# <翻译结束>

# <翻译开始>
func LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1
返回模板句柄1
# <翻译结束>

# <翻译开始>
func LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML
字符串内容
# <翻译结束>

# <翻译开始>
func LoadFromStringEx(nType
模板类型
# <翻译结束>

# <翻译开始>
func LoadFromStringEx
X加载并按字符串扩展
# <翻译结束>


# <翻译开始>
func LoadFromMemEx(nType xcc.ListItemTemp_Type_, data []byte, pOutTemp1 *int, pOutTemp2
返回模板句柄2
# <翻译结束>

# <翻译开始>
func LoadFromMemEx(nType xcc.ListItemTemp_Type_, data []byte, pOutTemp1
返回模板句柄1
# <翻译结束>

# <翻译开始>
func LoadFromMemEx(nType xcc.ListItemTemp_Type_, data
模板文件数据
# <翻译结束>

# <翻译开始>
func LoadFromMemEx(nType
模板类型
# <翻译结束>

# <翻译开始>
func LoadFromMemEx
X加载并按内存扩展
# <翻译结束>


# <翻译开始>
func LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int, hModule
模块句柄
# <翻译结束>

# <翻译开始>
func LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2
返回模板句柄2
# <翻译结束>

# <翻译开始>
func LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1
返回模板句柄1
# <翻译结束>

# <翻译开始>
func LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName
模板文件名
# <翻译结束>

# <翻译开始>
func LoadZipResEx(nType xcc.ListItemTemp_Type_, id
RC资源ID
# <翻译结束>

# <翻译开始>
func LoadZipResEx(nType
模板类型
# <翻译结束>

# <翻译开始>
func LoadZipResEx
X加载并按资源ZIP扩展
# <翻译结束>

