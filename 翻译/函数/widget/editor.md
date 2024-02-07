
# <翻译开始>
func NewEditor(x int, y int, cx int, cy int, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewEditor(x int, y int, cx int, cy
高度
# <翻译结束>

# <翻译开始>
func NewEditor(x int, y int, cx
宽度
# <翻译结束>

# <翻译开始>
func NewEditor(x int, y
元素y坐标
# <翻译结束>

# <翻译开始>
func NewEditor(x
元素x坐标
# <翻译结束>

# <翻译开始>
func NewEditor
X创建代码编辑框
# <翻译结束>


# <翻译开始>
func NewEditorByHandle
X创建代码编辑框并按句柄
# <翻译结束>


# <翻译开始>
func NewEditorByName
X创建代码编辑框并按名称
# <翻译结束>


# <翻译开始>
func NewEditorByUID
X创建代码编辑框并按UID
# <翻译结束>


# <翻译开始>
func NewEditorByUIDName
X创建代码编辑框并按UID名称
# <翻译结束>


# <翻译开始>
func (e *Editor) EnableAutoMatchSpaseSelect(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Editor) EnableAutoMatchSpaseSelect
X启用空格选择自动匹配项
# <翻译结束>


# <翻译开始>
func (e *Editor) IsBreakpoint(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) IsBreakpoint
X判断断点
# <翻译结束>


# <翻译开始>
func (e *Editor) SetBreakpoint(iRow int, bActivate
是否激活
# <翻译结束>

# <翻译开始>
func (e *Editor) SetBreakpoint(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) SetBreakpoint
X置断点
# <翻译结束>


# <翻译开始>
func (e *Editor) RemoveBreakpoint(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) RemoveBreakpoint
X移除断点
# <翻译结束>


# <翻译开始>
func (e *Editor) ClearBreakpoint
X清空断点
# <翻译结束>


# <翻译开始>
func (e *Editor) SetRunRow(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) SetRunRow
X置当前运行
# <翻译结束>


# <翻译开始>
func (e *Editor) GetColor(pInfo
颜色信息结构体指针
# <翻译结束>

# <翻译开始>
func (e *Editor) GetColor
X取颜色信息
# <翻译结束>


# <翻译开始>
func (e *Editor) SetColor(pInfo
颜色信息结构体指针
# <翻译结束>

# <翻译开始>
func (e *Editor) SetColor
X置颜色
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleKeyword(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleKeyword
X置常量样式
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleFunction(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleFunction
X置函数样式
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleVar(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleVar
X置变量样式
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleDataType(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleDataType
X置数据类型样式
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleClass(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleClass
X置类样式
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleMacro(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleMacro
X置宏样式
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleString(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleString
X置字符串样式
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleComment(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleComment
X置注释样式
# <翻译结束>


# <翻译开始>
func (e *Editor) SetStyleNumber(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetStyleNumber
X置数字样式
# <翻译结束>


# <翻译开始>
func (e *Editor) GetBreakpointCount
X取断点数量
# <翻译结束>


# <翻译开始>
func (e *Editor) GetBreakpoints(aPoints *[]int32, nCount
数组大小
# <翻译结束>

# <翻译开始>
func (e *Editor) GetBreakpoints(aPoints
接收断点数组
# <翻译结束>

# <翻译开始>
func (e *Editor) GetBreakpoints
X取全部断点
# <翻译结束>


# <翻译开始>
func (e *Editor) SetCurRow(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) SetCurRow
X设置当前行
# <翻译结束>


# <翻译开始>
func (e *Editor) GetDepth(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) GetDepth
X获取深度
# <翻译结束>


# <翻译开始>
func (e *Editor) ToExpandRow(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) ToExpandRow
X转换到展开行
# <翻译结束>


# <翻译开始>
func (e *Editor) ExpandEx(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) ExpandEx
X展开EX
# <翻译结束>


# <翻译开始>
func (e *Editor) ExpandAll(bExpand
是否展开
# <翻译结束>

# <翻译开始>
func (e *Editor) ExpandAll
X展开全部
# <翻译结束>


# <翻译开始>
func (e *Editor) Expand(iRow int, bExpand
是否展开
# <翻译结束>

# <翻译开始>
func (e *Editor) Expand(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Editor) Expand
X展开指定行
# <翻译结束>


# <翻译开始>
func (e *Editor) AddKeyword(pKey string, iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Editor) AddKeyword(pKey
字符串
# <翻译结束>

# <翻译开始>
func (e *Editor) AddKeyword
X添加关键字
# <翻译结束>


# <翻译开始>
func (e *Editor) AddConst(pKey
字符串
# <翻译结束>

# <翻译开始>
func (e *Editor) AddConst
X添加自动匹配字符串
# <翻译结束>


# <翻译开始>
func (e *Editor) AddFunction(pKey
字符串
# <翻译结束>

# <翻译开始>
func (e *Editor) AddFunction
X添加自动匹配函数
# <翻译结束>


# <翻译开始>
func (e *Editor) AddExcludeDefVarKeyword(pKeyword
字符串
# <翻译结束>

# <翻译开始>
func (e *Editor) AddExcludeDefVarKeyword
X添加排除定义变量关键字
# <翻译结束>


# <翻译开始>
func (e *Editor) GetExpandState
X获取折叠状态
# <翻译结束>


# <翻译开始>
func (e *Editor) SetExpandState
X设置折叠状态
# <翻译结束>


# <翻译开始>
func (e *Editor) GetIndentation(iRow
行
# <翻译结束>

# <翻译开始>
func (e *Editor) GetIndentation
X获取缩进
# <翻译结束>


# <翻译开始>
func (e *Editor) IsEmptyRow(iRow
行
# <翻译结束>

# <翻译开始>
func (e *Editor) IsEmptyRow
X是否为空行
# <翻译结束>


# <翻译开始>
func (e *Editor) SetAutoMatchMode(mode
模式
# <翻译结束>

# <翻译开始>
func (e *Editor) SetAutoMatchMode
X置自动匹配结果显示模式
# <翻译结束>


# <翻译开始>
func (e *Editor) Event_EDITOR_MODIFY_ROWS
X事件_多行内容改变
# <翻译结束>


# <翻译开始>
func (e *Editor) Event_EDITOR_MODIFY_ROWS1
X事件_多行内容改变1
# <翻译结束>


# <翻译开始>
func (e *Editor) Event_EDITOR_SETBREAKPOINT
X事件_设置断点
# <翻译结束>



# <翻译开始>
func (e *Editor) Event_EDITOR_SETBREAKPOINT1
X事件_设置断点1
# <翻译结束>


# <翻译开始>
func (e *Editor) Event_EDITOR_REMOVEBREAKPOINT
X事件_移除断点
# <翻译结束>


# <翻译开始>
func (e *Editor) Event_EDITOR_REMOVEBREAKPOINT1
X事件_移除断点1
# <翻译结束>


# <翻译开始>
func (e *Editor) Event_EDITOR_AUTOMATCH_SELECT
X事件_自动匹配选择
# <翻译结束>


# <翻译开始>
func (e *Editor) Event_EDITOR_AUTOMATCH_SELECT1
X事件_自动匹配选择1
# <翻译结束>

