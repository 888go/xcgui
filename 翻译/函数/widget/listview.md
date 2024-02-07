
# <翻译开始>
func NewListView(x int, y int, cx int, cy int, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewListView(x int, y int, cx int, cy
高度
# <翻译结束>

# <翻译开始>
func NewListView(x int, y int, cx
宽度
# <翻译结束>

# <翻译开始>
func NewListView(x int, y
元素y坐标
# <翻译结束>

# <翻译开始>
func NewListView(x
元素x坐标
# <翻译结束>

# <翻译开始>
func NewListView
X创建列表视
# <翻译结束>


# <翻译开始>
func NewListViewEx(x, y, cx, cy int32, hParent, col_extend_count
列数量
# <翻译结束>

# <翻译开始>
func NewListViewEx(x, y, cx, cy int32, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewListViewEx(x, y, cx, cy
高度
# <翻译结束>

# <翻译开始>
func NewListViewEx(x, y, cx
宽度
# <翻译结束>

# <翻译开始>
func NewListViewEx(x, y
元素y坐标
# <翻译结束>

# <翻译开始>
func NewListViewEx(x
元素x坐标
# <翻译结束>

# <翻译开始>
func NewListViewEx
X创建列表视Ex
# <翻译结束>


# <翻译开始>
func NewListViewByHandle
X创建列表视并按句柄
# <翻译结束>


# <翻译开始>
func NewListViewByName
X创建列表视并按名称
# <翻译结束>


# <翻译开始>
func NewListViewByUID
X创建列表视并按UID
# <翻译结束>


# <翻译开始>
func NewListViewByUIDName
X创建列表视并按UID名称
# <翻译结束>


# <翻译开始>
func (l *ListView) CreateAdapter
X创建数据适配器
# <翻译结束>


# <翻译开始>
func (l *ListView) BindAdapter(hAdapter
数据适配器
# <翻译结束>

# <翻译开始>
func (l *ListView) BindAdapter
X绑定数据适配器
# <翻译结束>


# <翻译开始>
func (l *ListView) GetAdapter
X取数据适配器
# <翻译结束>


# <翻译开始>
func (l *ListView) SetItemTemplateXML(pXmlFile
文件名
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemTemplateXML
X置项模板文件
# <翻译结束>


# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromString(pStringXML
字符串
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromString
X置项模板从字符串
# <翻译结束>


# <翻译开始>
func (l *ListView) SetItemTemplate(hTemp
模板句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemTemplate
X置项模板
# <翻译结束>


# <翻译开始>
func (l *ListView) GetTemplateObject(iGroup int, iItem int, nTempItemID
模板项ID
# <翻译结束>

# <翻译开始>
func (l *ListView) GetTemplateObject(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetTemplateObject(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetTemplateObject
X取模板对象
# <翻译结束>


# <翻译开始>
func (l *ListView) GetTemplateObjectGroup(iGroup int, nTempItemID
模板项ID
# <翻译结束>

# <翻译开始>
func (l *ListView) GetTemplateObjectGroup(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetTemplateObjectGroup
X取模板对象组
# <翻译结束>


# <翻译开始>
func (l *ListView) GetItemIDFromHXCGUI(hXCGUI int, piGroup *int32, piItem
接收项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetItemIDFromHXCGUI(hXCGUI int, piGroup
接收组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetItemIDFromHXCGUI(hXCGUI
对象句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) GetItemIDFromHXCGUI
X取对象所在项
# <翻译结束>


# <翻译开始>
func (l *ListView) HitTest(pPt *xc.POINT, pOutGroup *int32, pOutItem
接收项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) HitTest(pPt *xc.POINT, pOutGroup
接收组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) HitTest(pPt
坐标点
# <翻译结束>

# <翻译开始>
func (l *ListView) HitTest
X测试点击项
# <翻译结束>


# <翻译开始>
func (l *ListView) HitTestOffset(pPt *xc.POINT, pOutGroup *int32, pOutItem
接收项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) HitTestOffset(pPt *xc.POINT, pOutGroup
接收做索引
# <翻译结束>

# <翻译开始>
func (l *ListView) HitTestOffset(pPt
坐标点
# <翻译结束>

# <翻译开始>
func (l *ListView) HitTestOffset
X测试点击项EX
# <翻译结束>


# <翻译开始>
func (l *ListView) EnableMultiSel(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (l *ListView) EnableMultiSel
X启用多选
# <翻译结束>


# <翻译开始>
func (l *ListView) EnableTemplateReuse(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (l *ListView) EnableTemplateReuse
X启用模板复用
# <翻译结束>


# <翻译开始>
func (l *ListView) EnableVirtualTable(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (l *ListView) EnableVirtualTable
X启用虚表
# <翻译结束>


# <翻译开始>
func (l *ListView) SetVirtualItemCount(iGroup int, nCount
项数量
# <翻译结束>

# <翻译开始>
func (l *ListView) SetVirtualItemCount(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) SetVirtualItemCount
X置虚表项数量
# <翻译结束>


# <翻译开始>
func (l *ListView) SetDrawItemBkFlags(nFlags
标志位
# <翻译结束>

# <翻译开始>
func (l *ListView) SetDrawItemBkFlags
X置项背景绘制标志
# <翻译结束>


# <翻译开始>
func (l *ListView) SetSelectItem(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) SetSelectItem(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) SetSelectItem
X置选择项
# <翻译结束>


# <翻译开始>
func (l *ListView) GetSelectItem(piGroup *int32, piItem
接收项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetSelectItem(piGroup
接收组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetSelectItem
X取选择项
# <翻译结束>


# <翻译开始>
func (l *ListView) AddSelectItem(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) AddSelectItem(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) AddSelectItem
X添加选择项
# <翻译结束>


# <翻译开始>
func (l *ListView) VisibleItem(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) VisibleItem(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) VisibleItem
X显示指定项
# <翻译结束>


# <翻译开始>
func (l *ListView) GetVisibleItemRange(piGroup1 *int32, piGroup2 *int32, piStartGroup *int32, piStartItem *int32, piEndGroup *int32, piEndItem
可视结束项
# <翻译结束>

# <翻译开始>
func (l *ListView) GetVisibleItemRange(piGroup1 *int32, piGroup2 *int32, piStartGroup *int32, piStartItem *int32, piEndGroup
可视结束组
# <翻译结束>

# <翻译开始>
func (l *ListView) GetVisibleItemRange(piGroup1 *int32, piGroup2 *int32, piStartGroup *int32, piStartItem
可视开始项
# <翻译结束>

# <翻译开始>
func (l *ListView) GetVisibleItemRange(piGroup1 *int32, piGroup2 *int32, piStartGroup
可视开始组
# <翻译结束>

# <翻译开始>
func (l *ListView) GetVisibleItemRange(piGroup1 *int32, piGroup2
组2
# <翻译结束>

# <翻译开始>
func (l *ListView) GetVisibleItemRange(piGroup1
组1
# <翻译结束>

# <翻译开始>
func (l *ListView) GetVisibleItemRange
X取可视项范围
# <翻译结束>


# <翻译开始>
func (l *ListView) GetSelectItemCount
X取选择项数量
# <翻译结束>


# <翻译开始>
func (l *ListView) GetSelectAll(pArray *[]xc.ListView_Item_Id_, nArraySize
数组大小
# <翻译结束>

# <翻译开始>
func (l *ListView) GetSelectAll(pArray
数组
# <翻译结束>

# <翻译开始>
func (l *ListView) GetSelectAll
X取选择项全部
# <翻译结束>


# <翻译开始>
func (l *ListView) SetSelectAll
X置选择项全部
# <翻译结束>


# <翻译开始>
func (l *ListView) CancelSelectAll
X取消选择项全部
# <翻译结束>


# <翻译开始>
func (l *ListView) SetColumnSpace(space
间隔大小
# <翻译结束>

# <翻译开始>
func (l *ListView) SetColumnSpace
X置列间隔
# <翻译结束>


# <翻译开始>
func (l *ListView) SetRowSpace(space
间隔大小
# <翻译结束>

# <翻译开始>
func (l *ListView) SetRowSpace
X置行间隔
# <翻译结束>


# <翻译开始>
func (l *ListView) SetItemSize(width int, height
高度
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemSize(width
宽度
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemSize
X置项大小
# <翻译结束>


# <翻译开始>
func (l *ListView) GetItemSize(pSize
接收返回大小
# <翻译结束>

# <翻译开始>
func (l *ListView) GetItemSize
X取项大小
# <翻译结束>


# <翻译开始>
func (l *ListView) SetGroupHeight(height
高度
# <翻译结束>

# <翻译开始>
func (l *ListView) SetGroupHeight
X置组高度
# <翻译结束>


# <翻译开始>
func (l *ListView) GetGroupHeight
X取组高度
# <翻译结束>


# <翻译开始>
func (l *ListView) SetGroupUserData(iGroup int, nData
数据
# <翻译结束>

# <翻译开始>
func (l *ListView) SetGroupUserData(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) SetGroupUserData
X置组用户数据
# <翻译结束>


# <翻译开始>
func (l *ListView) SetItemUserData(iGroup int, iItem int, nData
数据
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemUserData(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemUserData(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemUserData
X置项用户数据
# <翻译结束>


# <翻译开始>
func (l *ListView) GetGroupUserData(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetGroupUserData
X取组用户数据
# <翻译结束>


# <翻译开始>
func (l *ListView) GetItemUserData(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetItemUserData(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) GetItemUserData
X取项用户数据
# <翻译结束>


# <翻译开始>
func (l *ListView) RefreshData
X刷新项数据
# <翻译结束>


# <翻译开始>
func (l *ListView) RefreshItem(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) RefreshItem(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) RefreshItem
X刷新指定项
# <翻译结束>


# <翻译开始>
func (l *ListView) ExpandGroup(iGroup int, bExpand
是否展开
# <翻译结束>

# <翻译开始>
func (l *ListView) ExpandGroup(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) ExpandGroup
X展开组
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_AddColumn(pName
字段称
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddColumn
X组添加列
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_AddItemText(pValue string, iPos
插入位置
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemText(pValue
值
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemText
X组添加项文本
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_AddItemTextEx(pName string, pValue string, iPos
插入位置
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemTextEx(pName string, pValue
值
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemTextEx(pName
字段称
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemTextEx
X组添加项文本EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_AddItemImage(hImage int, iPos
插入位置
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemImage(hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemImage
X组添加项图片
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_AddItemImageEx(pName string, hImage int, iPos
插入位置
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemImageEx(pName string, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemImageEx(pName
字段称
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_AddItemImageEx
X组添加项图片EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_SetText(iGroup int, iColumn int, pValue
值
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetText(iGroup int, iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetText(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetText
X组置文本
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_SetTextEx(iGroup int, pName string, pValue
值
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetTextEx(iGroup int, pName
字段名
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetTextEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetTextEx
X组置文本EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_SetImage(iGroup int, iColumn int, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetImage(iGroup int, iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetImage(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetImage
X组置图片
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_SetImageEx(iGroup int, pName string, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetImageEx(iGroup int, pName
字段名
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetImageEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_SetImageEx
X组置图片EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_GetCount
X组获取数量
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_GetCount(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetCount
X项获取数量
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_AddColumn(pName
字段名
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddColumn
X项添加列
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_AddItemText(iGroup int, pValue string, iPos
插入位置
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemText(iGroup int, pValue
值
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemText(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemText
X项添加文本
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_AddItemTextEx(iGroup int, pName string, pValue string, iPos
插入位置
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemTextEx(iGroup int, pName string, pValue
值
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemTextEx(iGroup int, pName
字段名
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemTextEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemTextEx
X项添加文本EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_AddItemImage(iGroup int, hImage int, iPos
插入位置
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemImage(iGroup int, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemImage(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemImage
X项添加图片
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_AddItemImageEx(iGroup int, pName string, hImage int, iPos
插入位置
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemImageEx(iGroup int, pName string, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemImageEx(iGroup int, pName
字段名
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemImageEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_AddItemImageEx
X项添加图片EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_SetText(iGroup int, iItem int, iColumn int, pValue
值
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetText(iGroup int, iItem int, iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetText(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetText(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetText
X项置文本
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_SetTextEx(iGroup int, iItem int, pName string, pValue
值
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetTextEx(iGroup int, iItem int, pName
字段名
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetTextEx(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetTextEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetTextEx
X项置文本EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_SetImage(iGroup int, iItem int, iColumn int, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetImage(iGroup int, iItem int, iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetImage(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetImage(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetImage
X项置图片
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_SetImageEx(iGroup int, iItem int, pName string, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetImageEx(iGroup int, iItem int, pName
列名称
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetImageEx(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetImageEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_SetImageEx
X项置图片EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_DeleteItem(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_DeleteItem
X组删除项
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_DeleteAllChildItem(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_DeleteAllChildItem
X组删除全部子项
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_DeleteItem(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_DeleteItem(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_DeleteItem
X项删除
# <翻译结束>


# <翻译开始>
func (l *ListView) DeleteAll
X删除全部
# <翻译结束>


# <翻译开始>
func (l *ListView) DeleteAllGroup
X删除全部组
# <翻译结束>


# <翻译开始>
func (l *ListView) DeleteAllItem
X删除全部项
# <翻译结束>


# <翻译开始>
func (l *ListView) DeleteColumnGroup(iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) DeleteColumnGroup
X组删除列
# <翻译结束>


# <翻译开始>
func (l *ListView) DeleteColumnItem(iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) DeleteColumnItem
X项删除列
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_GetTextEx(iGroup int, iItem int, pName
字段称
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetTextEx(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetTextEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetTextEx
X项获取文本
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_GetImageEx(iGroup int, iItem int, pName
字段称
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetImageEx(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetImageEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetImageEx
X项获取图片EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_GetText(iGroup int, iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_GetText(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_GetText
X组取文本
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_GetTextEx(iGroup int, pName
字段名称
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_GetTextEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_GetTextEx
X组取文本EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_GetImage(iGroup int, iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_GetImage(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_GetImage
X组取图片
# <翻译结束>


# <翻译开始>
func (l *ListView) Group_GetImageEx(iGroup int, pName
字段名称
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_GetImageEx(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Group_GetImageEx
X组取图片EX
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_GetText(iGroup int, iItem int, iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetText(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetText(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetText
X项取文本
# <翻译结束>


# <翻译开始>
func (l *ListView) Item_GetImage(iGroup int, iItem int, iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetImage(iGroup int, iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetImage(iGroup
组索引
# <翻译结束>

# <翻译开始>
func (l *ListView) Item_GetImage
X项取图片
# <翻译结束>


# <翻译开始>
func (l *ListView) SetDragRectColor(color int, width
线宽度
# <翻译结束>

# <翻译开始>
func (l *ListView) SetDragRectColor(color
ABGR颜色
# <翻译结束>

# <翻译开始>
func (l *ListView) SetDragRectColor
X置拖动矩形颜色
# <翻译结束>


# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromMem(data
模板数据
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromMem
X置项模板并按内存
# <翻译结束>


# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule
模块句柄
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromZipRes(id int32, pFileName
文件名
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromZipRes(id
RC资源ID
# <翻译结束>

# <翻译开始>
func (l *ListView) SetItemTemplateXMLFromZipRes
X置项模板并按资源ZIP
# <翻译结束>


# <翻译开始>
func (l *ListView) GetItemTemplate
X取项模板
# <翻译结束>


# <翻译开始>
func (l *ListView) GetItemTemplateGroup
X取项模板组
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_TEMP_CREATE
X事件_项模板创建
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_TEMP_CREATE1
X事件_项模板创建1
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_TEMP_CREATE_END
X事件_项模板创建完成
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_TEMP_CREATE_END1
X事件_项模板创建完成1
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_TEMP_DESTROY
X事件_项模板销毁
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_TEMP_DESTROY1
X事件_项模板销毁1
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_TEMP_ADJUST_COORDINATE
X事件_项模板调整坐标
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_TEMP_ADJUST_COORDINATE1
X事件_项模板调整坐标1
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_DRAWITEM
X事件_自绘项
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_DRAWITEM1
X事件_自绘项1
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_SELECT
X事件_项选择
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_SELECT1
X事件_项选择1
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_EXPAND
X事件_组展开收缩
# <翻译结束>


# <翻译开始>
func (l *ListView) Event_LISTVIEW_EXPAND1
X事件_组展开收缩1
# <翻译结束>

