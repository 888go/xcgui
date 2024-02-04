
# <翻译开始>
func NewMenu
X创建菜单
# <翻译结束>


# <翻译开始>
func NewMenuByHandle
X创建菜单并按句柄
# <翻译结束>


# <翻译开始>
func NewMenuByName
X创建菜单并按名称
# <翻译结束>


# <翻译开始>
func NewMenuByUID
X创建菜单并按UID
# <翻译结束>


# <翻译开始>
func NewMenuByUIDName
X创建菜单并按UID名称
# <翻译结束>


# <翻译开始>
func (m *Menu) AddItem(nID int32, pText string, nParentID int32, nFlags
标识
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItem(nID int32, pText string, nParentID
父项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItem(nID int32, pText
文本内容
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItem(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItem
X添加项
# <翻译结束>


# <翻译开始>
func (m *Menu) AddItemIcon(nID int32, pText string, nParentID int32, hIcon int, nFlags
标识
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItemIcon(nID int32, pText string, nParentID int32, hIcon
图标句柄
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItemIcon(nID int32, pText string, nParentID
父项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItemIcon(nID int32, pText
文本内容
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItemIcon(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) AddItemIcon
X添加项图标
# <翻译结束>


# <翻译开始>
func (m *Menu) InsertItem(nID int32, pText string, nFlags xcc.Menu_Item_Flag_, insertID
插入位置ID
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItem(nID int32, pText string, nFlags
标识
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItem(nID int32, pText
文本内容
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItem(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItem
X插入项
# <翻译结束>


# <翻译开始>
func (m *Menu) InsertItemIcon(nID int32, pText string, hIcon int, nFlags xcc.Menu_Item_Flag_, insertID
插入位置ID
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItemIcon(nID int32, pText string, hIcon int, nFlags
标识
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItemIcon(nID int32, pText string, hIcon
图标句柄
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItemIcon(nID int32, pText
文本内容
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItemIcon(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) InsertItemIcon
X插入项图标
# <翻译结束>


# <翻译开始>
func (m *Menu) GetFirstChildItem(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) GetFirstChildItem
X取第一个子项
# <翻译结束>


# <翻译开始>
func (m *Menu) GetEndChildItem(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) GetEndChildItem
X取末尾子项
# <翻译结束>


# <翻译开始>
func (m *Menu) GetPrevSiblingItem(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) GetPrevSiblingItem
X取上一个兄弟项
# <翻译结束>


# <翻译开始>
func (m *Menu) GetNextSiblingItem(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) GetNextSiblingItem
X取下一个兄弟项
# <翻译结束>


# <翻译开始>
func (m *Menu) GetParentItem(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) GetParentItem
X取父项
# <翻译结束>


# <翻译开始>
func (m *Menu) SetAutoDestroy(bAuto
是否自动销毁
# <翻译结束>

# <翻译开始>
func (m *Menu) SetAutoDestroy
X置自动销毁
# <翻译结束>


# <翻译开始>
func (m *Menu) EnableDrawBackground(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (m *Menu) EnableDrawBackground
X启用用户绘制背景
# <翻译结束>


# <翻译开始>
func (m *Menu) EnableDrawItem(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (m *Menu) EnableDrawItem
X启用用户绘制项
# <翻译结束>


# <翻译开始>
func (m *Menu) Popup(hParentWnd uintptr, x, y int32, hParentEle int, nPosition
弹出位置
# <翻译结束>

# <翻译开始>
func (m *Menu) Popup(hParentWnd uintptr, x, y int32, hParentEle
父元素句柄
# <翻译结束>

# <翻译开始>
func (m *Menu) Popup(hParentWnd uintptr, x, y
y坐标
# <翻译结束>

# <翻译开始>
func (m *Menu) Popup(hParentWnd uintptr, x
x坐标
# <翻译结束>

# <翻译开始>
func (m *Menu) Popup(hParentWnd
父窗口句柄
# <翻译结束>

# <翻译开始>
func (m *Menu) Popup
X弹出
# <翻译结束>


# <翻译开始>
func (m *Menu) DestroyMenu
X销毁
# <翻译结束>


# <翻译开始>
func (m *Menu) CloseMenu
X关闭
# <翻译结束>


# <翻译开始>
func (m *Menu) SetBkImage(hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (m *Menu) SetBkImage
X置背景图片
# <翻译结束>


# <翻译开始>
func (m *Menu) SetItemText(nID int32, pText
文本内容
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemText(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemText
X置项文本
# <翻译结束>


# <翻译开始>
func (m *Menu) GetItemText(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) GetItemText
X取项文本
# <翻译结束>


# <翻译开始>
func (m *Menu) GetItemTextLength(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) GetItemTextLength
X取项文本长度
# <翻译结束>


# <翻译开始>
func (m *Menu) SetItemIcon(nID int32, hIcon
菜单项图标句柄
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemIcon(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemIcon
X置项图标
# <翻译结束>


# <翻译开始>
func (m *Menu) SetItemFlags(nID int32, uFlags
标识
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemFlags(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemFlags
X置项标志
# <翻译结束>


# <翻译开始>
func (m *Menu) SetItemHeight(height
高度
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemHeight
X置项高度
# <翻译结束>


# <翻译开始>
func (m *Menu) GetItemHeight
X取项高度
# <翻译结束>


# <翻译开始>
func (m *Menu) SetBorderColor(crColor
ABGR颜色
# <翻译结束>

# <翻译开始>
func (m *Menu) SetBorderColor
X置边框颜色
# <翻译结束>


# <翻译开始>
func (m *Menu) SetBorderSize(nLeft, nTop, nRight, nBottom
下
# <翻译结束>

# <翻译开始>
func (m *Menu) SetBorderSize(nLeft, nTop, nRight
右
# <翻译结束>

# <翻译开始>
func (m *Menu) SetBorderSize(nLeft, nTop
上
# <翻译结束>

# <翻译开始>
func (m *Menu) SetBorderSize(nLeft
左
# <翻译结束>

# <翻译开始>
func (m *Menu) SetBorderSize
X置边框大小
# <翻译结束>


# <翻译开始>
func (m *Menu) GetLeftWidth
X取左侧宽度
# <翻译结束>


# <翻译开始>
func (m *Menu) GetLeftSpaceText
X取左侧文本间隔
# <翻译结束>


# <翻译开始>
func (m *Menu) GetItemCount
X取项数量
# <翻译结束>


# <翻译开始>
func (m *Menu) SetItemCheck(nID int32, bCheck
勾选TRUE
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemCheck(nID
菜单项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemCheck
X置项勾选
# <翻译结束>


# <翻译开始>
func (m *Menu) IsItemCheck(nID
菜单项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) IsItemCheck
X判断项勾选
# <翻译结束>


# <翻译开始>
func (m *Menu) SetItemWidth(nID, nWidth
宽度
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemWidth(nID
项ID
# <翻译结束>

# <翻译开始>
func (m *Menu) SetItemWidth
X置项宽度
# <翻译结束>


# <翻译开始>
func (m *Menu) GetMenuBar
X取菜单条
# <翻译结束>

