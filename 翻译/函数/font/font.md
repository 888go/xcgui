
# <翻译开始>
func New(size
字体大小
# <翻译结束>

# <翻译开始>
func New
X创建
# <翻译结束>


# <翻译开始>
func NewEX(pName string, size int32, style
字体样式
# <翻译结束>

# <翻译开始>
func NewEX(pName string, size
字体大小
# <翻译结束>

# <翻译开始>
func NewEX(pName
字体名称
# <翻译结束>

# <翻译开始>
func NewEX
X创建扩展
# <翻译结束>


# <翻译开始>
func NewLOGFONTW(pFontInfo
字体选项
# <翻译结束>

# <翻译开始>
func NewLOGFONTW
X创建并按选项
# <翻译结束>


# <翻译开始>
func NewByHFONT(hFont
字体句柄
# <翻译结束>

# <翻译开始>
func NewByHFONT
X创建并按字体句柄
# <翻译结束>


# <翻译开始>
func NewByFont(pFont
GDI字体指针
# <翻译结束>

# <翻译开始>
func NewByFont
X创建并按GDI字体指针
# <翻译结束>


# <翻译开始>
func NewByFile(pFontFile string, size int32, style
字体样式
# <翻译结束>

# <翻译开始>
func NewByFile(pFontFile string, size
字体大小
# <翻译结束>

# <翻译开始>
func NewByFile(pFontFile
字体文件名
# <翻译结束>

# <翻译开始>
func NewByFile
X创建并按文件
# <翻译结束>


# <翻译开始>
func NewByZip(pZipFileName, pFileName, pPassword string, fontSize int32, style
字体样式
# <翻译结束>

# <翻译开始>
func NewByZip(pZipFileName, pFileName, pPassword string, fontSize
字体大小
# <翻译结束>

# <翻译开始>
func NewByZip(pZipFileName, pFileName, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func NewByZip(pZipFileName, pFileName
zip文件名
# <翻译结束>

# <翻译开始>
func NewByZip(pZipFileName
zip文件名
# <翻译结束>

# <翻译开始>
func NewByZip
X创建并按ZIP
# <翻译结束>


# <翻译开始>
func NewByZipMem(data []byte, pFileName, pPassword string, fontSize int32, style
字体样式
# <翻译结束>

# <翻译开始>
func NewByZipMem(data []byte, pFileName, pPassword string, fontSize
字体大小
# <翻译结束>

# <翻译开始>
func NewByZipMem(data []byte, pFileName, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func NewByZipMem(data []byte, pFileName
字体文件名
# <翻译结束>

# <翻译开始>
func NewByZipMem(data
zip数据
# <翻译结束>

# <翻译开始>
func NewByZipMem
X创建并按内存ZIP
# <翻译结束>


# <翻译开始>
func NewByMem(data []byte, fontSize int32, style
字体样式
# <翻译结束>

# <翻译开始>
func NewByMem(data []byte, fontSize
字体大小
# <翻译结束>

# <翻译开始>
func NewByMem(data
字体文件数据
# <翻译结束>

# <翻译开始>
func NewByMem
X创建并按内存
# <翻译结束>


# <翻译开始>
func NewByRes(id int32, pType string, fontSize int32, style xcc.FontStyle_, hModule
模块
# <翻译结束>

# <翻译开始>
func NewByRes(id int32, pType string, fontSize int32, style
字体样式
# <翻译结束>

# <翻译开始>
func NewByRes(id int32, pType string, fontSize
字体大小
# <翻译结束>

# <翻译开始>
func NewByRes(id int32, pType
类型
# <翻译结束>

# <翻译开始>
func NewByRes
X创建并按资源
# <翻译结束>


# <翻译开始>
func NewByHandle(handle
句柄
# <翻译结束>

# <翻译开始>
func NewByHandle
X创建并按句柄
# <翻译结束>


# <翻译开始>
func NewByName(name
名称
# <翻译结束>

# <翻译开始>
func NewByName
X创建并按名称
# <翻译结束>


# <翻译开始>
func (f *Font) EnableAutoDestroy(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (f *Font) EnableAutoDestroy
X启用自动销毁
# <翻译结束>


# <翻译开始>
func (f *Font) GetFont
X取字体指针
# <翻译结束>


# <翻译开始>
func (f *Font) GetFontInfo(pInfo
返回字体信息
# <翻译结束>

# <翻译开始>
func (f *Font) GetFontInfo
X取信息
# <翻译结束>


# <翻译开始>
func (f *Font) GetLOGFONTW(hdc uintptr, pOut
返回字体属性
# <翻译结束>

# <翻译开始>
func (f *Font) GetLOGFONTW(hdc
hdc句柄
# <翻译结束>

# <翻译开始>
func (f *Font) GetLOGFONTW
X取属性
# <翻译结束>


# <翻译开始>
func (f *Font) Destroy
X销毁
# <翻译结束>


# <翻译开始>
func (f *Font) AddRef
X增加引用计数
# <翻译结束>


# <翻译开始>
func (f *Font) GetRefCount
X取引用计数
# <翻译结束>


# <翻译开始>
func (f *Font) Release
X释放引用计数
# <翻译结束>

