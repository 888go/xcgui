
# <翻译开始>
func NewList(x int, y int, cx int, cy int, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewList(x int, y int, cx int, cy
高度
# <翻译结束>

# <翻译开始>
func NewList(x int, y int, cx
宽度
# <翻译结束>

# <翻译开始>
func NewList(x int, y
元素y坐标
# <翻译结束>

# <翻译开始>
func NewList(x
元素x坐标
# <翻译结束>

# <翻译开始>
func NewList
X创建列表
# <翻译结束>


# <翻译开始>
func NewListEx(x, y, cx, cy int32, hParent, col_extend_count
列数量
# <翻译结束>

# <翻译开始>
func NewListEx(x, y, cx, cy int32, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewListEx(x, y, cx, cy
高度
# <翻译结束>

# <翻译开始>
func NewListEx(x, y, cx
宽度
# <翻译结束>

# <翻译开始>
func NewListEx(x, y
元素y坐标
# <翻译结束>

# <翻译开始>
func NewListEx(x
元素x坐标
# <翻译结束>

# <翻译开始>
func NewListEx
X创建列表Ex
# <翻译结束>


# <翻译开始>
func NewListByHandle
X创建列表并按句柄
# <翻译结束>


# <翻译开始>
func NewListByName
X创建列表并按名称
# <翻译结束>


# <翻译开始>
func NewListByUID
X创建列表并按UID
# <翻译结束>


# <翻译开始>
func NewListByUIDName
X创建列表并按UID名称
# <翻译结束>


# <翻译开始>
func (l *List) AddColumn(width
列宽度
# <翻译结束>

# <翻译开始>
func (l *List) AddColumn
X增加列
# <翻译结束>


# <翻译开始>
func (l *List) InsertColumn(width int, iItem
插入位置索引
# <翻译结束>

# <翻译开始>
func (l *List) InsertColumn(width
列宽度
# <翻译结束>

# <翻译开始>
func (l *List) InsertColumn
X插入列
# <翻译结束>


# <翻译开始>
func (l *List) EnableMultiSel(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (l *List) EnableMultiSel
X启用多选
# <翻译结束>


# <翻译开始>
func (l *List) EnableDragChangeColumnWidth(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (l *List) EnableDragChangeColumnWidth
X启用拖动更改列宽
# <翻译结束>


# <翻译开始>
func (l *List) EnableVScrollBarTop(bTop
是否启用
# <翻译结束>

# <翻译开始>
func (l *List) EnableVScrollBarTop
X启用垂直滚动条顶部对齐
# <翻译结束>


# <翻译开始>
func (l *List) EnableItemBkFullRow(bFull
是否启用
# <翻译结束>

# <翻译开始>
func (l *List) EnableItemBkFullRow
X启用项背景全行模式
# <翻译结束>


# <翻译开始>
func (l *List) EnableFixedRowHeight(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (l *List) EnableFixedRowHeight
X启用固定行高
# <翻译结束>


# <翻译开始>
func (l *List) EnableTemplateReuse(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (l *List) EnableTemplateReuse
X启用模板复用
# <翻译结束>


# <翻译开始>
func (l *List) EnableVirtualTable(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (l *List) EnableVirtualTable
X启用虚表
# <翻译结束>


# <翻译开始>
func (l *List) SetVirtualRowCount(nRowCount
行数
# <翻译结束>

# <翻译开始>
func (l *List) SetVirtualRowCount
X置虚表行数
# <翻译结束>


# <翻译开始>
func (l *List) SetSort(iColumn int, iColumnAdapter int, bEnable
是否启用排序
# <翻译结束>

# <翻译开始>
func (l *List) SetSort(iColumn int, iColumnAdapter
数据适配器中列索引
# <翻译结束>

# <翻译开始>
func (l *List) SetSort(iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *List) SetSort
X置排序
# <翻译结束>


# <翻译开始>
func (l *List) SetDrawItemBkFlags(nFlags
标志位
# <翻译结束>

# <翻译开始>
func (l *List) SetDrawItemBkFlags
X置绘制项背景标志
# <翻译结束>


# <翻译开始>
func (l *List) SetColumnWidth(iItem int, width
宽度
# <翻译结束>

# <翻译开始>
func (l *List) SetColumnWidth(iItem
列索引
# <翻译结束>

# <翻译开始>
func (l *List) SetColumnWidth
X置列宽
# <翻译结束>


# <翻译开始>
func (l *List) SetColumnMinWidth(iItem int, width
宽度
# <翻译结束>

# <翻译开始>
func (l *List) SetColumnMinWidth(iItem
列索引
# <翻译结束>

# <翻译开始>
func (l *List) SetColumnMinWidth
X置列最小宽度
# <翻译结束>


# <翻译开始>
func (l *List) SetColumnWidthFixed(iColumn int, bFixed
是否固定宽度
# <翻译结束>

# <翻译开始>
func (l *List) SetColumnWidthFixed(iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *List) SetColumnWidthFixed
X置列宽度固定
# <翻译结束>


# <翻译开始>
func (l *List) GetColumnWidth(iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *List) GetColumnWidth
X取列宽度
# <翻译结束>


# <翻译开始>
func (l *List) GetColumnCount
X取列数量
# <翻译结束>


# <翻译开始>
func (l *List) SetItemData(iItem int, iSubItem int, data
用户数据
# <翻译结束>

# <翻译开始>
func (l *List) SetItemData(iItem int, iSubItem
子项索引
# <翻译结束>

# <翻译开始>
func (l *List) SetItemData(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) SetItemData
X置项数据
# <翻译结束>


# <翻译开始>
func (l *List) GetItemData(iItem int, iSubItem
子项索引
# <翻译结束>

# <翻译开始>
func (l *List) GetItemData(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) GetItemData
X取项数据
# <翻译结束>


# <翻译开始>
func (l *List) SetSelectItem(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) SetSelectItem
X置选择项
# <翻译结束>


# <翻译开始>
func (l *List) GetSelectItem
X取选择项
# <翻译结束>


# <翻译开始>
func (l *List) GetSelectItemCount
X取选择项数量
# <翻译结束>


# <翻译开始>
func (l *List) AddSelectItem(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) AddSelectItem
X添加选择项
# <翻译结束>


# <翻译开始>
func (l *List) SetSelectAll
X置选择全部
# <翻译结束>


# <翻译开始>
func (l *List) GetSelectAll(pArray *[]int32, nArraySize
数组大小
# <翻译结束>

# <翻译开始>
func (l *List) GetSelectAll(pArray
接收行索引数组
# <翻译结束>

# <翻译开始>
func (l *List) GetSelectAll
X取全部选择
# <翻译结束>


# <翻译开始>
func (l *List) VisibleItem(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) VisibleItem
X显示指定项
# <翻译结束>


# <翻译开始>
func (l *List) CancelSelectItem(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) CancelSelectItem
X取消选择项
# <翻译结束>


# <翻译开始>
func (l *List) CancelSelectAll
X取消全部选择项
# <翻译结束>


# <翻译开始>
func (l *List) GetHeaderHELE
X取列表头
# <翻译结束>


# <翻译开始>
func (l *List) DeleteColumn(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) DeleteColumn
X删除列
# <翻译结束>


# <翻译开始>
func (l *List) DeleteColumnAll
X删除全部列
# <翻译结束>


# <翻译开始>
func (l *List) BindAdapter(hAdapter
数据适配器句柄
# <翻译结束>

# <翻译开始>
func (l *List) BindAdapter
X绑定数据适配器
# <翻译结束>


# <翻译开始>
func (l *List) BindAdapterHeader(hAdapter
数据适配器句柄
# <翻译结束>

# <翻译开始>
func (l *List) BindAdapterHeader
X列表头绑定数据适配器
# <翻译结束>


# <翻译开始>
func (l *List) CreateAdapter(colExtend_count
	预计列表总列数
# <翻译结束>

# <翻译开始>
func (l *List) CreateAdapter
X创建数据适配器
# <翻译结束>


# <翻译开始>
func (l *List) CreateAdapterHeader
X创建列表头数据适配器
# <翻译结束>


# <翻译开始>
func (l *List) GetAdapter
X取数据适配器
# <翻译结束>


# <翻译开始>
func (l *List) GetAdapterHeader
X取列表头数据适配器
# <翻译结束>


# <翻译开始>
func (l *List) SetItemTemplateXML(pXmlFile
文件名
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTemplateXML
X置项模板文件
# <翻译结束>


# <翻译开始>
func (l *List) SetItemTemplateXMLFromString(pStringXML
字符串
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTemplateXMLFromString
X置项模板从字符串
# <翻译结束>


# <翻译开始>
func (l *List) SetItemTemplate(hTemp
模板句柄
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTemplate
X置项模板
# <翻译结束>


# <翻译开始>
func (l *List) GetTemplateObject(iItem int, iSubItem int, nTempItemID
模板项itemID
# <翻译结束>

# <翻译开始>
func (l *List) GetTemplateObject(iItem int, iSubItem
子项索引
# <翻译结束>

# <翻译开始>
func (l *List) GetTemplateObject(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) GetTemplateObject
X取项模板对象
# <翻译结束>


# <翻译开始>
func (l *List) GetItemIndexFromHXCGUI(hXCGUI
对象句柄
# <翻译结束>

# <翻译开始>
func (l *List) GetItemIndexFromHXCGUI
X取对象所在行
# <翻译结束>


# <翻译开始>
func (l *List) GetHeaderTemplateObject(iItem int, nTempItemID
模板项ID
# <翻译结束>

# <翻译开始>
func (l *List) GetHeaderTemplateObject(iItem
列表头项ID
# <翻译结束>

# <翻译开始>
func (l *List) GetHeaderTemplateObject
X取列表头模板对象
# <翻译结束>


# <翻译开始>
func (l *List) GetHeaderItemIndexFromHXCGUI(hXCGUI
对象句柄
# <翻译结束>

# <翻译开始>
func (l *List) GetHeaderItemIndexFromHXCGUI
X取列表头对象所在行
# <翻译结束>


# <翻译开始>
func (l *List) SetHeaderHeight(height
高度
# <翻译结束>

# <翻译开始>
func (l *List) SetHeaderHeight
X置列表头高度
# <翻译结束>


# <翻译开始>
func (l *List) GetHeaderHeight
X取列表头高度
# <翻译结束>


# <翻译开始>
func (l *List) GetVisibleRowRange(piStart *int32, piEnd
结束行索引
# <翻译结束>

# <翻译开始>
func (l *List) GetVisibleRowRange(piStart
开始行索引
# <翻译结束>

# <翻译开始>
func (l *List) GetVisibleRowRange
X取可视行范围
# <翻译结束>


# <翻译开始>
func (l *List) SetItemHeightDefault(nHeight int32, nSelHeight
选中时高度
# <翻译结束>

# <翻译开始>
func (l *List) SetItemHeightDefault(nHeight
高度
# <翻译结束>

# <翻译开始>
func (l *List) SetItemHeightDefault
X置项默认高度
# <翻译结束>


# <翻译开始>
func (l *List) GetItemHeightDefault(pHeight *int32, pSelHeight
选中时高度
# <翻译结束>

# <翻译开始>
func (l *List) GetItemHeightDefault(pHeight
高度
# <翻译结束>

# <翻译开始>
func (l *List) GetItemHeightDefault
X取项默认高度
# <翻译结束>


# <翻译开始>
func (l *List) SetRowSpace(nSpace
行间距大小
# <翻译结束>

# <翻译开始>
func (l *List) SetRowSpace
X置行间距
# <翻译结束>


# <翻译开始>
func (l *List) GetRowSpace
X取行间距
# <翻译结束>


# <翻译开始>
func (l *List) SetLockColumnLeft(iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *List) SetLockColumnLeft
X置锁定列左侧
# <翻译结束>


# <翻译开始>
func (l *List) SetLockColumnRight(iColumn
列索引
# <翻译结束>

# <翻译开始>
func (l *List) SetLockColumnRight
X置锁定列右侧
# <翻译结束>


# <翻译开始>
func (l *List) SetLockRowBottom(bLock
是否锁定
# <翻译结束>

# <翻译开始>
func (l *List) SetLockRowBottom
X置锁定行底部
# <翻译结束>


# <翻译开始>
func (l *List) SetLockRowBottomOverlap(bOverlap
是否重叠
# <翻译结束>

# <翻译开始>
func (l *List) SetLockRowBottomOverlap
X置锁定行底部重叠
# <翻译结束>


# <翻译开始>
func (l *List) HitTest(pPt *xc.POINT, piItem *int32, piSubItem
子项索引
# <翻译结束>

# <翻译开始>
func (l *List) HitTest(pPt *xc.POINT, piItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) HitTest(pPt
坐标点
# <翻译结束>

# <翻译开始>
func (l *List) HitTest
X测试点击项
# <翻译结束>


# <翻译开始>
func (l *List) HitTestOffset(pPt *xc.POINT, piItem *int32, piSubItem
子项索引
# <翻译结束>

# <翻译开始>
func (l *List) HitTestOffset(pPt *xc.POINT, piItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) HitTestOffset(pPt
坐标点
# <翻译结束>

# <翻译开始>
func (l *List) HitTestOffset
X测试点击项EX
# <翻译结束>


# <翻译开始>
func (l *List) RefreshData
X刷新项数据
# <翻译结束>


# <翻译开始>
func (l *List) RefreshItem(iItem
项索引
# <翻译结束>

# <翻译开始>
func (l *List) RefreshItem
X刷新指定项
# <翻译结束>


# <翻译开始>
func (l *List) AddColumnText(nWidth int, pName string, pText
文本
# <翻译结束>

# <翻译开始>
func (l *List) AddColumnText(nWidth int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) AddColumnText(nWidth
列宽
# <翻译结束>

# <翻译开始>
func (l *List) AddColumnText
X添加列文本
# <翻译结束>


# <翻译开始>
func (l *List) AddColumnImage(nWidth int, pName string, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (l *List) AddColumnImage(nWidth int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) AddColumnImage(nWidth
列宽
# <翻译结束>

# <翻译开始>
func (l *List) AddColumnImage
X添加列图片
# <翻译结束>


# <翻译开始>
func (l *List) AddItemText(pText
文本
# <翻译结束>

# <翻译开始>
func (l *List) AddItemText
X添加项文本
# <翻译结束>


# <翻译开始>
func (l *List) AddItemTextEx(pName string, pText
文本
# <翻译结束>

# <翻译开始>
func (l *List) AddItemTextEx(pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) AddItemTextEx
X添加项文本EX
# <翻译结束>


# <翻译开始>
func (l *List) AddItemImage(hImage
图片
# <翻译结束>

# <翻译开始>
func (l *List) AddItemImage
X添加项图片
# <翻译结束>


# <翻译开始>
func (l *List) AddItemImageEx(pName string, hImage
图片
# <翻译结束>

# <翻译开始>
func (l *List) AddItemImageEx(pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) AddItemImageEx
X添加项图片EX
# <翻译结束>


# <翻译开始>
func (l *List) InsertItemText(iItem int, pValue
文本
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemText(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemText
X插入项文本
# <翻译结束>


# <翻译开始>
func (l *List) InsertItemTextEx(iItem int, pName string, pValue
文本
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemTextEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemTextEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemTextEx
X插入项文本EX
# <翻译结束>


# <翻译开始>
func (l *List) InsertItemImage(iItem int, hImage
图片
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemImage(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemImage
X插入项图片
# <翻译结束>


# <翻译开始>
func (l *List) InsertItemImageEx(iItem int, pName string, hImage
图片
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemImageEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemImageEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) InsertItemImageEx
X插入项图片EX
# <翻译结束>


# <翻译开始>
func (l *List) SetItemText(iItem int, iColumn int, pText
文本
# <翻译结束>

# <翻译开始>
func (l *List) SetItemText(iItem int, iColumn
列
# <翻译结束>

# <翻译开始>
func (l *List) SetItemText(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) SetItemText
X置项文本
# <翻译结束>


# <翻译开始>
func (l *List) SetItemTextEx(iItem int, pName string, pText
文本
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTextEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTextEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTextEx
X置项文本EX
# <翻译结束>


# <翻译开始>
func (l *List) SetItemImage(iItem int, iColumn int, hImage
图片
# <翻译结束>

# <翻译开始>
func (l *List) SetItemImage(iItem int, iColumn
列
# <翻译结束>

# <翻译开始>
func (l *List) SetItemImage(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) SetItemImage
X置项图片
# <翻译结束>


# <翻译开始>
func (l *List) SetItemImageEx(iItem int, pName string, hImage
图片
# <翻译结束>

# <翻译开始>
func (l *List) SetItemImageEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) SetItemImageEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) SetItemImageEx
X置项图片EX
# <翻译结束>


# <翻译开始>
func (l *List) SetItemInt(iItem int, iColumn int, nValue
值
# <翻译结束>

# <翻译开始>
func (l *List) SetItemInt(iItem int, iColumn
列
# <翻译结束>

# <翻译开始>
func (l *List) SetItemInt(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) SetItemInt
X置项指数值
# <翻译结束>


# <翻译开始>
func (l *List) SetItemIntEx(iItem int, pName string, nValue
值
# <翻译结束>

# <翻译开始>
func (l *List) SetItemIntEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) SetItemIntEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) SetItemIntEx
X置项整数值EX
# <翻译结束>


# <翻译开始>
func (l *List) SetItemFloat(iItem int, iColumn int, fFloat
值
# <翻译结束>

# <翻译开始>
func (l *List) SetItemFloat(iItem int, iColumn
列
# <翻译结束>

# <翻译开始>
func (l *List) SetItemFloat(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) SetItemFloat
X置项浮点值
# <翻译结束>


# <翻译开始>
func (l *List) SetItemFloatEx(iItem int, pName string, fFloat
值
# <翻译结束>

# <翻译开始>
func (l *List) SetItemFloatEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) SetItemFloatEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) SetItemFloatEx
X置项浮点值EX
# <翻译结束>


# <翻译开始>
func (l *List) GetItemText(iItem int, iColumn
列
# <翻译结束>

# <翻译开始>
func (l *List) GetItemText(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) GetItemText
X取项文本
# <翻译结束>


# <翻译开始>
func (l *List) GetItemTextEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) GetItemTextEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) GetItemTextEx
X取项文本EX
# <翻译结束>


# <翻译开始>
func (l *List) GetItemImage(iItem int, iColumn
列
# <翻译结束>

# <翻译开始>
func (l *List) GetItemImage(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) GetItemImage
X取项图片
# <翻译结束>


# <翻译开始>
func (l *List) GetItemImageEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) GetItemImageEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) GetItemImageEx
X取项图片EX
# <翻译结束>


# <翻译开始>
func (l *List) GetItemInt(iItem int, iColumn int, pOutValue
值指针
# <翻译结束>

# <翻译开始>
func (l *List) GetItemInt(iItem int, iColumn
列
# <翻译结束>

# <翻译开始>
func (l *List) GetItemInt(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) GetItemInt
X取项整数值
# <翻译结束>


# <翻译开始>
func (l *List) GetItemIntEx(iItem int, pName string, pOutValue
值指针
# <翻译结束>

# <翻译开始>
func (l *List) GetItemIntEx(iItem int, pName
文本
# <翻译结束>

# <翻译开始>
func (l *List) GetItemIntEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) GetItemIntEx
X取项整数值EX
# <翻译结束>


# <翻译开始>
func (l *List) GetItemFloat(iItem int, iColumn int, pOutValue
值指针
# <翻译结束>

# <翻译开始>
func (l *List) GetItemFloat(iItem int, iColumn
列
# <翻译结束>

# <翻译开始>
func (l *List) GetItemFloat(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) GetItemFloat
X取项浮点值
# <翻译结束>


# <翻译开始>
func (l *List) GetItemFloatEx(iItem int, pName string, pOutValue
值指针
# <翻译结束>

# <翻译开始>
func (l *List) GetItemFloatEx(iItem int, pName
名称
# <翻译结束>

# <翻译开始>
func (l *List) GetItemFloatEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) GetItemFloatEx
X取项浮点值EX
# <翻译结束>


# <翻译开始>
func (l *List) DeleteItem(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) DeleteItem
X删除项
# <翻译结束>


# <翻译开始>
func (l *List) DeleteItemEx(iItem int, nCount
计数
# <翻译结束>

# <翻译开始>
func (l *List) DeleteItemEx(iItem
项
# <翻译结束>

# <翻译开始>
func (l *List) DeleteItemEx
X删除项EX
# <翻译结束>


# <翻译开始>
func (l *List) DeleteItemAll
X删除项全部
# <翻译结束>


# <翻译开始>
func (l *List) DeleteColumnAll_AD
X删除列全部AD
# <翻译结束>


# <翻译开始>
func (l *List) GetCount_AD
X取项数量AD
# <翻译结束>


# <翻译开始>
func (l *List) GetCountColumn_AD
X取列数量AD
# <翻译结束>


# <翻译开始>
func (l *List) SetSplitLineColor(color
ABGR颜色值
# <翻译结束>

# <翻译开始>
func (l *List) SetSplitLineColor
X置分割线颜色
# <翻译结束>


# <翻译开始>
func (l *List) SetItemHeight(iRow int, nHeight, nSelHeight
选中时高度
# <翻译结束>

# <翻译开始>
func (l *List) SetItemHeight(iRow int, nHeight
高度
# <翻译结束>

# <翻译开始>
func (l *List) SetItemHeight(iRow
行索引
# <翻译结束>

# <翻译开始>
func (l *List) SetItemHeight
X置项高度
# <翻译结束>


# <翻译开始>
func (l *List) GetItemHeight(iRow int, pHeight, pSelHeight
选中时高度
# <翻译结束>

# <翻译开始>
func (l *List) GetItemHeight(iRow int, pHeight
高度
# <翻译结束>

# <翻译开始>
func (l *List) GetItemHeight(iRow
行索引
# <翻译结束>

# <翻译开始>
func (l *List) GetItemHeight
X取项高度
# <翻译结束>


# <翻译开始>
func (l *List) SetDragRectColor(color, width
线宽度
# <翻译结束>

# <翻译开始>
func (l *List) SetDragRectColor(color
ABGR颜色值
# <翻译结束>

# <翻译开始>
func (l *List) SetDragRectColor
X置拖动矩形颜色
# <翻译结束>


# <翻译开始>
func (l *List) GetItemTemplate
X取项模板
# <翻译结束>


# <翻译开始>
func (l *List) GetItemTemplateHeader
X取项模板列表头
# <翻译结束>


# <翻译开始>
func (l *List) RefreshDataHeader
X刷新项数据列表头
# <翻译结束>


# <翻译开始>
func (l *List) SetItemTemplateXMLFromMem(data
模板数据
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTemplateXMLFromMem
X置项模板从内存
# <翻译结束>


# <翻译开始>
func (l *List) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule
模块句柄
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword
zip密码
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTemplateXMLFromZipRes(id int32, pFileName
项模板文件名
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTemplateXMLFromZipRes(id
RC资源ID
# <翻译结束>

# <翻译开始>
func (l *List) SetItemTemplateXMLFromZipRes
X置项模板并按资源ZIP
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_TEMP_CREATE
X事件_项模板创建
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_TEMP_CREATE1
X事件_项模板创建1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_TEMP_CREATE_END
X事件_项模板创建完成
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_TEMP_CREATE_END1
X事件_项模板创建完成1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_TEMP_DESTROY
X事件_项模板销毁
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_TEMP_DESTROY1
X事件_项模板销毁1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_TEMP_ADJUST_COORDINATE
X停用_项模板调整坐标
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_TEMP_ADJUST_COORDINATE1
X停用_项模板调整坐标1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_DRAWITEM
X事件_绘制项
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_DRAWITEM1
X事件_绘制项1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_SELECT
X事件_项选择
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_SELECT1
X事件_项选择1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_DRAWITEM
X事件_列表元素绘制列表头项
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_DRAWITEM1
X事件_列表元素绘制列表头项1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_CLICK
X事件_列表头项点击
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_CLICK1
X事件_列表头项点击1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_WIDTH_CHANGE
X事件_列表头项宽度改变
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_WIDTH_CHANGE1
X事件_列表头项宽度改变1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_TEMP_CREATE
X事件_列表头项模板创建
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_TEMP_CREATE1
X事件_列表头项模板创建1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_TEMP_CREATE_END
X事件_列表头项模板创建完成
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_TEMP_CREATE_END1
X事件_列表头项模板创建完成1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_TEMP_DESTROY
X事件_列表头项模板销毁
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_TEMP_DESTROY1
X事件_列表头项模板销毁1
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_TEMP_ADJUST_COORDINATE
X停用_列表头项模板调整坐标
# <翻译结束>


# <翻译开始>
func (l *List) Event_LIST_HEADER_TEMP_ADJUST_COORDINATE1
X停用_列表头项模板调整坐标1
# <翻译结束>

