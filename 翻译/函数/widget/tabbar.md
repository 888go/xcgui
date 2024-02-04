
# <翻译开始>
func NewTabBar(x int, y int, cx int, cy int, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewTabBar(x int, y int, cx int, cy
高度
# <翻译结束>

# <翻译开始>
func NewTabBar(x int, y int, cx
宽度
# <翻译结束>

# <翻译开始>
func NewTabBar(x int, y
元素y坐标
# <翻译结束>

# <翻译开始>
func NewTabBar(x
元素x坐标
# <翻译结束>

# <翻译开始>
func NewTabBar
X创建TAB条
# <翻译结束>


# <翻译开始>
func NewTabBarByHandle
X创建TAB条并按句柄
# <翻译结束>


# <翻译开始>
func NewTabBarByName
X创建TAB条并按名称
# <翻译结束>


# <翻译开始>
func NewTabBarByUID
X创建TAB条并按UID
# <翻译结束>


# <翻译开始>
func NewTabBarByUIDName
X创建TAB条并按UID名称
# <翻译结束>


# <翻译开始>
func (t *TabBar) AddLabel(pName
标签文本
# <翻译结束>

# <翻译开始>
func (t *TabBar) AddLabel
X添加标签
# <翻译结束>


# <翻译开始>
func (t *TabBar) InsertLabel(index int, pName
标签文本
# <翻译结束>

# <翻译开始>
func (t *TabBar) InsertLabel(index
插入位置
# <翻译结束>

# <翻译开始>
func (t *TabBar) InsertLabel
XTAB条插入_标签
# <翻译结束>


# <翻译开始>
func (t *TabBar) MoveLabel(iSrc int, iDest
目标位置索引
# <翻译结束>

# <翻译开始>
func (t *TabBar) MoveLabel(iSrc
源位置索引
# <翻译结束>

# <翻译开始>
func (t *TabBar) MoveLabel
X移动标签
# <翻译结束>


# <翻译开始>
func (t *TabBar) DeleteLabel(index
位置索引
# <翻译结束>

# <翻译开始>
func (t *TabBar) DeleteLabel
X删除标签
# <翻译结束>


# <翻译开始>
func (t *TabBar) DeleteLabelAll
X删除全部
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetLabel(index
位置索引
# <翻译结束>

# <翻译开始>
func (t *TabBar) GetLabel
X取标签
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetLabelClose(index
位置索引
# <翻译结束>

# <翻译开始>
func (t *TabBar) GetLabelClose
X取标签上的关闭按钮
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetButtonLeft
X取左滚动按钮
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetButtonRight
X取右滚动按钮
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetButtonDropMenu
X取下拉菜单按钮句柄
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetSelect
X取当前选择
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetLabelSpacing
X取间隔
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetLabelCount
X取标签数量
# <翻译结束>


# <翻译开始>
func (t *TabBar) GetindexByEle(hLabel
标签按钮句柄
# <翻译结束>

# <翻译开始>
func (t *TabBar) GetindexByEle
X取标签位置索引
# <翻译结束>


# <翻译开始>
func (t *TabBar) SetLabelSpacing(spacing
标签间隔大小
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetLabelSpacing
X置间隔
# <翻译结束>


# <翻译开始>
func (t *TabBar) SetPadding(left int, top int, right int, bottom
下
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetPadding(left int, top int, right
右
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetPadding(left int, top
上
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetPadding(left
左
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetPadding
X置边距
# <翻译结束>


# <翻译开始>
func (t *TabBar) SetSelect(index
索引
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetSelect
X置选择
# <翻译结束>


# <翻译开始>
func (t *TabBar) SetUp
X左滚动
# <翻译结束>


# <翻译开始>
func (t *TabBar) SetDown
X右滚动
# <翻译结束>


# <翻译开始>
func (t *TabBar) EnableTile(bTile
是否启用
# <翻译结束>

# <翻译开始>
func (t *TabBar) EnableTile
X启用平铺
# <翻译结束>


# <翻译开始>
func (t *TabBar) EnableDropMenu(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (t *TabBar) EnableDropMenu
X启用下拉菜单按钮
# <翻译结束>


# <翻译开始>
func (t *TabBar) EnableClose(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (t *TabBar) EnableClose
X启用标签带关闭按钮
# <翻译结束>


# <翻译开始>
func (t *TabBar) SetCloseSize(pSize
大小值
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetCloseSize
X置关闭按钮大小
# <翻译结束>


# <翻译开始>
func (t *TabBar) SetTurnButtonSize(pSize
大小值
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetTurnButtonSize
X置滚动按钮大小
# <翻译结束>


# <翻译开始>
func (t *TabBar) SetLabelWidth(index int, nWidth
宽度
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetLabelWidth(index
索引
# <翻译结束>

# <翻译开始>
func (t *TabBar) SetLabelWidth
X置指定标签固定宽度
# <翻译结束>


# <翻译开始>
func (t *TabBar) ShowLabel(index int, bShow
是否显示
# <翻译结束>

# <翻译开始>
func (t *TabBar) ShowLabel(index
标签索引
# <翻译结束>

# <翻译开始>
func (t *TabBar) ShowLabel
X显示标签
# <翻译结束>


# <翻译开始>
func (t *TabBar) Event_TABBAR_SELECT
X事件_选择改变
# <翻译结束>


# <翻译开始>
func (t *TabBar) Event_TABBAR_SELECT1
X事件_选择改变1
# <翻译结束>


# <翻译开始>
func (t *TabBar) Event_TABBAR_DELETE
X事件_删除
# <翻译结束>


# <翻译开始>
func (t *TabBar) Event_TABBAR_DELETE1
X事件_删除1
# <翻译结束>

