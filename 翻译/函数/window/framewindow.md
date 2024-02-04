
# <翻译开始>
func NewFrameWindow(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle
GUI库窗口样式
# <翻译结束>

# <翻译开始>
func NewFrameWindow(x, y, cx, cy int32, pTitle string, hWndParent
父窗口句柄
# <翻译结束>

# <翻译开始>
func NewFrameWindow(x, y, cx, cy int32, pTitle
标题
# <翻译结束>

# <翻译开始>
func NewFrameWindow(x, y, cx, cy
高度
# <翻译结束>

# <翻译开始>
func NewFrameWindow(x, y, cx
宽度
# <翻译结束>

# <翻译开始>
func NewFrameWindow(x, y
y坐标
# <翻译结束>

# <翻译开始>
func NewFrameWindow(x
x坐标
# <翻译结束>

# <翻译开始>
func NewFrameWindow
X创建框架窗口
# <翻译结束>


# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle
GUI库窗口样式
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent
父窗口
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle
窗口名
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy
高度
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx
宽度
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y
y坐标
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x
x坐标
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName
类名
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle int, dwStyle
样式
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx(dwExStyle
扩展样式
# <翻译结束>

# <翻译开始>
func NewFrameWindowEx
X创建框架窗口EX
# <翻译结束>


# <翻译开始>
func NewFrameWindowByHandle
X创建框架窗口并按句柄
# <翻译结束>


# <翻译开始>
func NewFrameWindowByName
X创建框架窗口并按名称
# <翻译结束>


# <翻译开始>
func NewFrameWindowByUID
X创建框架窗口并按UID
# <翻译结束>


# <翻译开始>
func NewFrameWindowByUIDName
X创建框架窗口并按UID名称
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) GetLayoutAreaRect(pRect
返回坐标
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) GetLayoutAreaRect
X取布局区域坐标
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) SetView(hEle
元素句柄
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) SetView
X置视图
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) SetPaneSplitBarColor(color
ABGR颜色值
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) SetPaneSplitBarColor
X置窗格分隔条颜色
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) SetTabBarHeight(nHeight
高度
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) SetTabBarHeight
X置TabBar条高度
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) SaveLayoutToFile(pFileName
文件名
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) SaveLayoutToFile
X保存布局到文件
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) LoadLayoutFile(aPaneList []int, nPaneCount int32, pFileName
文件名
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) LoadLayoutFile(aPaneList []int, nPaneCount
窗格数量
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) LoadLayoutFile(aPaneList
窗格句柄数组
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) LoadLayoutFile
X加载布局信息文件
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) AddPane(hPaneDest int, hPaneNew int, align
对齐方式
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) AddPane(hPaneDest int, hPaneNew
当前窗格
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) AddPane(hPaneDest
目标窗格
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) AddPane
X添加窗格
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) MergePane(hPaneDest int, hPaneNew
当前窗格
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) MergePane(hPaneDest
目标窗格
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) MergePane
X合并窗格
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) GetDragFloatWndTopFlag
X取拖动浮动窗格停留位置标识
# <翻译结束>


# <翻译开始>
func FrameWnd_Attach(hWnd uintptr, XCStyle
炫彩窗口样式
# <翻译结束>

# <翻译开始>
func FrameWnd_Attach(hWnd
外部窗口句柄
# <翻译结束>

# <翻译开始>
func FrameWnd_Attach
X附加窗口
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) GetViewRect(pRect
返回坐标
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) GetViewRect
X取主视图坐标
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) SetPaneSplitBarWidth(nWidth
宽度
# <翻译结束>

# <翻译开始>
func (fw *FrameWindow) SetPaneSplitBarWidth
X置窗格分隔条宽度
# <翻译结束>


# <翻译开始>
func (fw *FrameWindow) GetPaneSplitBarWidth
X取窗格分隔条宽度
# <翻译结束>

