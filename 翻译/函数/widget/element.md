
# <翻译开始>
func NewElement(x, y, cx, cy int32, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewElement(x, y, cx, cy
高度
# <翻译结束>

# <翻译开始>
func NewElement(x, y, cx
宽度
# <翻译结束>

# <翻译开始>
func NewElement(x, y
y坐标
# <翻译结束>

# <翻译开始>
func NewElement(x
x坐标
# <翻译结束>

# <翻译开始>
func NewElement
X创建基础元素
# <翻译结束>


# <翻译开始>
func NewElementByHandle
X创建基础元素并按句柄
# <翻译结束>


# <翻译开始>
func NewElementByName
X创建基础元素并按名称
# <翻译结束>


# <翻译开始>
func NewElementByUID
X创建基础元素并按UID
# <翻译结束>


# <翻译开始>
func NewElementByUIDName
X创建基础元素并按UID名称
# <翻译结束>


# <翻译开始>
func (e *Element) RegEventC(nEvent xcc.XE_, pFun
事件函数
# <翻译结束>

# <翻译开始>
func (e *Element) RegEventC(nEvent
事件类型
# <翻译结束>

# <翻译开始>
func (e *Element) RegEventC
X注册事件C
# <翻译结束>


# <翻译开始>
func (e *Element) RegEventC1(nEvent xcc.XE_, pFun
事件函数
# <翻译结束>

# <翻译开始>
func (e *Element) RegEventC1(nEvent
事件类型
# <翻译结束>

# <翻译开始>
func (e *Element) RegEventC1
X注册事件C1
# <翻译结束>


# <翻译开始>
func (e *Element) RemoveEventC(nEvent xcc.XE_, pFun
事件函数
# <翻译结束>

# <翻译开始>
func (e *Element) RemoveEventC(nEvent
事件类型
# <翻译结束>

# <翻译开始>
func (e *Element) RemoveEventC
X移除事件C
# <翻译结束>


# <翻译开始>
func (e *Element) RegEventCEx(nEvent xcc.XE_, pFun
事件函数指针
# <翻译结束>

# <翻译开始>
func (e *Element) RegEventCEx(nEvent
事件类型
# <翻译结束>

# <翻译开始>
func (e *Element) RegEventCEx
X注册事件CEx
# <翻译结束>


# <翻译开始>
func (e *Element) RegEventC1Ex(nEvent xcc.XE_, pFun
事件函数指针
# <翻译结束>

# <翻译开始>
func (e *Element) RegEventC1Ex(nEvent
事件类型
# <翻译结束>

# <翻译开始>
func (e *Element) RegEventC1Ex
X注册事件C1Ex
# <翻译结束>


# <翻译开始>
func (e *Element) RemoveEventCEx(nEvent xcc.XE_, pFun
事件函数指针
# <翻译结束>

# <翻译开始>
func (e *Element) RemoveEventCEx(nEvent
事件类型
# <翻译结束>

# <翻译开始>
func (e *Element) RemoveEventCEx
X移除事件CEx
# <翻译结束>


# <翻译开始>
func (e *Element) SendEvent(nEvent xcc.XE_, wParam, lParam
参数2
# <翻译结束>

# <翻译开始>
func (e *Element) SendEvent(nEvent xcc.XE_, wParam
参数1
# <翻译结束>

# <翻译开始>
func (e *Element) SendEvent(nEvent
事件类型
# <翻译结束>

# <翻译开始>
func (e *Element) SendEvent
X发送事件
# <翻译结束>


# <翻译开始>
func (e *Element) PostEvent(nEvent xcc.XE_, wParam, lParam
参数2
# <翻译结束>

# <翻译开始>
func (e *Element) PostEvent(nEvent xcc.XE_, wParam
参数1
# <翻译结束>

# <翻译开始>
func (e *Element) PostEvent(nEvent
事件类型
# <翻译结束>

# <翻译开始>
func (e *Element) PostEvent
X投递事件
# <翻译结束>


# <翻译开始>
func (e *Element) GetRect(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) GetRect
X取坐标
# <翻译结束>


# <翻译开始>
func (e *Element) GetRectLogic(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) GetRectLogic
X取逻辑坐标
# <翻译结束>


# <翻译开始>
func (e *Element) GetClientRect(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) GetClientRect
X取客户区坐标
# <翻译结束>


# <翻译开始>
func (e *Element) SetWidth(nWidth
宽度
# <翻译结束>

# <翻译开始>
func (e *Element) SetWidth
X置宽度
# <翻译结束>


# <翻译开始>
func (e *Element) SetHeight(nHeight
高度
# <翻译结束>

# <翻译开始>
func (e *Element) SetHeight
X置高度
# <翻译结束>


# <翻译开始>
func (e *Element) GetWidth
X取宽度
# <翻译结束>


# <翻译开始>
func (e *Element) GetHeight
X取高度
# <翻译结束>


# <翻译开始>
func (e *Element) RectWndClientToEleClient(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) RectWndClientToEleClient
X窗口客户区坐标到元素客户区坐标
# <翻译结束>


# <翻译开始>
func (e *Element) PointWndClientToEleClient(pPt
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) PointWndClientToEleClient
X窗口客户区点到元素客户区
# <翻译结束>


# <翻译开始>
func (e *Element) RectClientToWndClient(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) RectClientToWndClient
X客户区坐标到窗口客户区
# <翻译结束>


# <翻译开始>
func (e *Element) PointClientToWndClient(pPt
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) PointClientToWndClient
X客户区点到窗口客户区
# <翻译结束>


# <翻译开始>
func (e *Element) GetWndClientRect(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) GetWndClientRect
X取窗口客户区坐标
# <翻译结束>


# <翻译开始>
func (e *Element) GetCursor
X取光标
# <翻译结束>


# <翻译开始>
func (e *Element) SetCursor(hCursor
光标句柄
# <翻译结束>

# <翻译开始>
func (e *Element) SetCursor
X置光标
# <翻译结束>


# <翻译开始>
func (e *Element) AddChild(hChild
子元素句柄或形状对象句柄
# <翻译结束>

# <翻译开始>
func (e *Element) AddChild
X添加子对象
# <翻译结束>


# <翻译开始>
func (e *Element) InsertChild(hChild int, index
插入位置索引
# <翻译结束>

# <翻译开始>
func (e *Element) InsertChild(hChild
元素句柄或形状对象句柄
# <翻译结束>

# <翻译开始>
func (e *Element) InsertChild
X插入子对象
# <翻译结束>


# <翻译开始>
func (e *Element) SetRect(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo
流水号
# <翻译结束>

# <翻译开始>
func (e *Element) SetRect(pRect *xc.RECT, bRedraw bool, nFlags
识位
# <翻译结束>

# <翻译开始>
func (e *Element) SetRect(pRect *xc.RECT, bRedraw
是否重绘
# <翻译结束>

# <翻译开始>
func (e *Element) SetRect(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetRect
X置坐标
# <翻译结束>


# <翻译开始>
func (e *Element) SetRectEx(x int, y int, cx int, cy int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo
流水号
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectEx(x int, y int, cx int, cy int, bRedraw bool, nFlags
标识
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectEx(x int, y int, cx int, cy int, bRedraw
是否重绘
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectEx(x int, y int, cx int, cy
高度
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectEx(x int, y int, cx
宽度
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectEx(x int, y
Y坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectEx(x
X坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectEx
X置坐标EX
# <翻译结束>


# <翻译开始>
func (e *Element) SetRectLogic(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo
流水号
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectLogic(pRect *xc.RECT, bRedraw bool, nFlags
识位
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectLogic(pRect *xc.RECT, bRedraw
是否重绘
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectLogic(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetRectLogic
X置逻辑坐标
# <翻译结束>


# <翻译开始>
func (e *Element) SetPosition(x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo
流水号
# <翻译结束>

# <翻译开始>
func (e *Element) SetPosition(x, y int32, bRedraw bool, nFlags
识位
# <翻译结束>

# <翻译开始>
func (e *Element) SetPosition(x, y int32, bRedraw
是否重绘
# <翻译结束>

# <翻译开始>
func (e *Element) SetPosition(x, y
Y坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetPosition(x
X坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetPosition
X移动
# <翻译结束>


# <翻译开始>
func (e *Element) SetPositionLogic(x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo
流水号
# <翻译结束>

# <翻译开始>
func (e *Element) SetPositionLogic(x, y int32, bRedraw bool, nFlags
识位
# <翻译结束>

# <翻译开始>
func (e *Element) SetPositionLogic(x, y int32, bRedraw
是否重绘
# <翻译结束>

# <翻译开始>
func (e *Element) SetPositionLogic(x, y
Y坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetPositionLogic(x
X坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetPositionLogic
X移动逻辑坐标
# <翻译结束>


# <翻译开始>
func (e *Element) IsDrawFocus
X判断绘制焦点
# <翻译结束>


# <翻译开始>
func (e *Element) IsEnable
X判断启用
# <翻译结束>


# <翻译开始>
func (e *Element) IsEnableFocus
X判断启用焦点
# <翻译结束>


# <翻译开始>
func (e *Element) IsMouseThrough
X判断鼠标穿透
# <翻译结束>


# <翻译开始>
func (e *Element) HitChildEle(pPt
坐标点
# <翻译结束>

# <翻译开始>
func (e *Element) HitChildEle
X测试点击元素
# <翻译结束>


# <翻译开始>
func (e *Element) IsBkTransparent
X判断背景透明
# <翻译结束>


# <翻译开始>
func (e *Element) IsEnableEvent_XE_PAINT_END
X判断启用事件_XE_PAINT_END
# <翻译结束>


# <翻译开始>
func (e *Element) IsKeyTab
X判断接受TAB
# <翻译结束>


# <翻译开始>
func (e *Element) IsSwitchFocus
X判断接受切换焦点
# <翻译结束>


# <翻译开始>
func (e *Element) IsEnable_XE_MOUSEWHEEL
X判断启用_XE_MOUSEWHEEL
# <翻译结束>


# <翻译开始>
func (e *Element) IsChildEle(hChildEle
子元素句柄
# <翻译结束>

# <翻译开始>
func (e *Element) IsChildEle
X判断为子元素
# <翻译结束>


# <翻译开始>
func (e *Element) IsEnableCanvas
X判断启用画布
# <翻译结束>


# <翻译开始>
func (e *Element) IsFocus
X判断焦点
# <翻译结束>


# <翻译开始>
func (e *Element) IsFocusEx
X判断焦点EX
# <翻译结束>


# <翻译开始>
func (e *Element) Enable(bEnable
启用或禁用
# <翻译结束>

# <翻译开始>
func (e *Element) Enable
X启用
# <翻译结束>


# <翻译开始>
func (e *Element) EnableFocus(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableFocus
X启用焦点
# <翻译结束>


# <翻译开始>
func (e *Element) EnableDrawFocus(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableDrawFocus
X启用绘制焦点
# <翻译结束>


# <翻译开始>
func (e *Element) EnableDrawBorder(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableDrawBorder
X启用绘制边框
# <翻译结束>


# <翻译开始>
func (e *Element) EnableCanvas(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableCanvas
X启用画布
# <翻译结束>


# <翻译开始>
func (e *Element) EnableEvent_XE_PAINT_END(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableEvent_XE_PAINT_END
X启用事件_XE_PAINT_END
# <翻译结束>


# <翻译开始>
func (e *Element) EnableBkTransparent(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableBkTransparent
X启用背景透明
# <翻译结束>


# <翻译开始>
func (e *Element) EnableMouseThrough(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableMouseThrough
X启用鼠标穿透
# <翻译结束>


# <翻译开始>
func (e *Element) EnableKeyTab(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableKeyTab
X启用接收TAB
# <翻译结束>


# <翻译开始>
func (e *Element) EnableSwitchFocus(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableSwitchFocus
X启用切换焦点
# <翻译结束>


# <翻译开始>
func (e *Element) EnableEvent_XE_MOUSEWHEEL(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Element) EnableEvent_XE_MOUSEWHEEL
X启用事件_XE_MOUSEWHEEL
# <翻译结束>


# <翻译开始>
func (e *Element) Remove
X移除
# <翻译结束>


# <翻译开始>
func (e *Element) SetZOrder(index
位置索引
# <翻译结束>

# <翻译开始>
func (e *Element) SetZOrder
X置Z序
# <翻译结束>


# <翻译开始>
func (e *Element) SetZOrderEx(hDestEle int, nType
类型
# <翻译结束>

# <翻译开始>
func (e *Element) SetZOrderEx(hDestEle
目标元素
# <翻译结束>

# <翻译开始>
func (e *Element) SetZOrderEx
X置Z序EX
# <翻译结束>


# <翻译开始>
func (e *Element) GetZOrder
X取Z序
# <翻译结束>


# <翻译开始>
func (e *Element) EnableTopmost(bTopmost
是否置顶显示
# <翻译结束>

# <翻译开始>
func (e *Element) EnableTopmost
X启用置顶
# <翻译结束>


# <翻译开始>
func (e *Element) Redraw(bImmediate
是否立即重绘
# <翻译结束>

# <翻译开始>
func (e *Element) Redraw
X重绘
# <翻译结束>


# <翻译开始>
func (e *Element) RedrawRect(pRect *xc.RECT, bImmediate
是否立即重绘
# <翻译结束>

# <翻译开始>
func (e *Element) RedrawRect(pRect
坐标
# <翻译结束>

# <翻译开始>
func (e *Element) RedrawRect
X重绘指定区域
# <翻译结束>


# <翻译开始>
func (e *Element) GetChildCount
X取子对象数量
# <翻译结束>


# <翻译开始>
func (e *Element) GetChildByIndex(index
索引
# <翻译结束>

# <翻译开始>
func (e *Element) GetChildByIndex
X取子对象并按索引
# <翻译结束>


# <翻译开始>
func (e *Element) GetChildByID(nID
元素ID
# <翻译结束>

# <翻译开始>
func (e *Element) GetChildByID
X取子对象并按ID
# <翻译结束>


# <翻译开始>
func (e *Element) SetBorderSize(left int, top int, right int, bottom
下边
# <翻译结束>

# <翻译开始>
func (e *Element) SetBorderSize(left int, top int, right
右边
# <翻译结束>

# <翻译开始>
func (e *Element) SetBorderSize(left int, top
上边
# <翻译结束>

# <翻译开始>
func (e *Element) SetBorderSize(left
左边
# <翻译结束>

# <翻译开始>
func (e *Element) SetBorderSize
X置边框大小
# <翻译结束>


# <翻译开始>
func (e *Element) GetBorderSize(pBorder
大小
# <翻译结束>

# <翻译开始>
func (e *Element) GetBorderSize
X取边框大小
# <翻译结束>


# <翻译开始>
func (e *Element) SetPadding(left int, top int, right int, bottom
下边
# <翻译结束>

# <翻译开始>
func (e *Element) SetPadding(left int, top int, right
右边
# <翻译结束>

# <翻译开始>
func (e *Element) SetPadding(left int, top
上边
# <翻译结束>

# <翻译开始>
func (e *Element) SetPadding(left
左边
# <翻译结束>

# <翻译开始>
func (e *Element) SetPadding
X置内填充大小
# <翻译结束>


# <翻译开始>
func (e *Element) GetPadding(pPadding
大小
# <翻译结束>

# <翻译开始>
func (e *Element) GetPadding
X取内填充大小
# <翻译结束>


# <翻译开始>
func (e *Element) SetDragBorder(nFlags
边框位置组合
# <翻译结束>

# <翻译开始>
func (e *Element) SetDragBorder
X置拖动边框
# <翻译结束>


# <翻译开始>
func (e *Element) SetDragBorderBindEle(nFlags xcc.Element_Position_, hBindEle int, nSpace
元素间隔大小
# <翻译结束>

# <翻译开始>
func (e *Element) SetDragBorderBindEle(nFlags xcc.Element_Position_, hBindEle
绑定元素
# <翻译结束>

# <翻译开始>
func (e *Element) SetDragBorderBindEle(nFlags
边框位置标识
# <翻译结束>

# <翻译开始>
func (e *Element) SetDragBorderBindEle
X置拖动边框绑定元素
# <翻译结束>


# <翻译开始>
func (e *Element) SetMinSize(nWidth int, nHeight
最小高度
# <翻译结束>

# <翻译开始>
func (e *Element) SetMinSize(nWidth
最小宽度
# <翻译结束>

# <翻译开始>
func (e *Element) SetMinSize
X置最小大小
# <翻译结束>


# <翻译开始>
func (e *Element) SetMaxSize(nWidth int, nHeight
最大高度
# <翻译结束>

# <翻译开始>
func (e *Element) SetMaxSize(nWidth
最大宽度
# <翻译结束>

# <翻译开始>
func (e *Element) SetMaxSize
X置最大大小
# <翻译结束>


# <翻译开始>
func (e *Element) SetLockScroll(bHorizon bool, bVertical
是否锁定垂直
# <翻译结束>

# <翻译开始>
func (e *Element) SetLockScroll(bHorizon
是否锁定水平
# <翻译结束>

# <翻译开始>
func (e *Element) SetLockScroll
X置锁定滚动
# <翻译结束>


# <翻译开始>
func (e *Element) SetTextColor(color
ABGR颜色值
# <翻译结束>

# <翻译开始>
func (e *Element) SetTextColor
X置文本颜色
# <翻译结束>


# <翻译开始>
func (e *Element) GetTextColor
X取文本颜色
# <翻译结束>


# <翻译开始>
func (e *Element) GetTextColorEx
X取文本颜色EX
# <翻译结束>


# <翻译开始>
func (e *Element) SetFocusBorderColor(color
ABGR颜色值
# <翻译结束>

# <翻译开始>
func (e *Element) SetFocusBorderColor
X置焦点边框颜色
# <翻译结束>


# <翻译开始>
func (e *Element) GetFocusBorderColor
X取焦点边框颜色
# <翻译结束>


# <翻译开始>
func (e *Element) SetFont(hFontx
炫彩字体
# <翻译结束>

# <翻译开始>
func (e *Element) SetFont
X置字体
# <翻译结束>


# <翻译开始>
func (e *Element) GetFont
X取字体
# <翻译结束>


# <翻译开始>
func (e *Element) GetFontEx
X取字体EX
# <翻译结束>


# <翻译开始>
func (e *Element) SetAlpha
X置透明度
# <翻译结束>


# <翻译开始>
func (e *Element) Destroy
X销毁
# <翻译结束>


# <翻译开始>
func (e *Element) AddBkBorder(nState xcc.CombinedState, color int, width
线宽
# <翻译结束>

# <翻译开始>
func (e *Element) AddBkBorder(nState xcc.CombinedState, color
ABGR颜色
# <翻译结束>

# <翻译开始>
func (e *Element) AddBkBorder(nState
组合状态
# <翻译结束>

# <翻译开始>
func (e *Element) AddBkBorder
X添加背景边框
# <翻译结束>


# <翻译开始>
func (e *Element) AddBkFill(nState xcc.CombinedState, color
ABGR颜色
# <翻译结束>

# <翻译开始>
func (e *Element) AddBkFill(nState
组合状态
# <翻译结束>

# <翻译开始>
func (e *Element) AddBkFill
X添加背景填充
# <翻译结束>


# <翻译开始>
func (e *Element) AddBkImage(nState xcc.CombinedState, hImage
图片句柄
# <翻译结束>

# <翻译开始>
func (e *Element) AddBkImage(nState
组合状态
# <翻译结束>

# <翻译开始>
func (e *Element) AddBkImage
X添加背景图片
# <翻译结束>


# <翻译开始>
func (e *Element) GetBkInfoCount
X取背景对象数量
# <翻译结束>


# <翻译开始>
func (e *Element) ClearBkInfo
X清空背景对象
# <翻译结束>


# <翻译开始>
func (e *Element) GetBkManager
X取背景管理器
# <翻译结束>


# <翻译开始>
func (e *Element) GetBkManagerEx
X取背景管理器EX
# <翻译结束>


# <翻译开始>
func (e *Element) SetBkManager(hBkInfoM
背景管理器
# <翻译结束>

# <翻译开始>
func (e *Element) SetBkManager
X置背景管理器
# <翻译结束>


# <翻译开始>
func (e *Element) GetStateFlags
X取状态
# <翻译结束>


# <翻译开始>
func (e *Element) DrawFocus(hDraw int, pRect
区域坐标
# <翻译结束>

# <翻译开始>
func (e *Element) DrawFocus(hDraw
图形绘制句柄
# <翻译结束>

# <翻译开始>
func (e *Element) DrawFocus
X绘制焦点
# <翻译结束>


# <翻译开始>
func (e *Element) DrawEle(hDraw
图形绘制句柄
# <翻译结束>

# <翻译开始>
func (e *Element) DrawEle
X绘制
# <翻译结束>


# <翻译开始>
func (e *Element) SetUserData(nData
用户数据
# <翻译结束>

# <翻译开始>
func (e *Element) SetUserData
X置用户数据
# <翻译结束>


# <翻译开始>
func (e *Element) GetUserData
X取用户数据
# <翻译结束>


# <翻译开始>
func (e *Element) GetContentSize(bHorizon bool, cx int, cy int, pSize
返回大小
# <翻译结束>

# <翻译开始>
func (e *Element) GetContentSize(bHorizon bool, cx int, cy
高度
# <翻译结束>

# <翻译开始>
func (e *Element) GetContentSize(bHorizon bool, cx
宽度
# <翻译结束>

# <翻译开始>
func (e *Element) GetContentSize(bHorizon
水平或垂直
# <翻译结束>

# <翻译开始>
func (e *Element) GetContentSize
X取内容大小
# <翻译结束>


# <翻译开始>
func (e *Element) SetCapture(b b
TRUE设置
# <翻译结束>

# <翻译开始>
func (e *Element) SetCapture
X置鼠标捕获
# <翻译结束>


# <翻译开始>
func (e *Element) EnableTransparentChannel(bEnable
启用或关闭
# <翻译结束>

# <翻译开始>
func (e *Element) EnableTransparentChannel
X启用透明通道
# <翻译结束>


# <翻译开始>
func (e *Element) SetXCTimer(nIDEvent int, uElapse
延时毫秒
# <翻译结束>

# <翻译开始>
func (e *Element) SetXCTimer(nIDEvent
事件ID
# <翻译结束>

# <翻译开始>
func (e *Element) SetXCTimer
X置炫彩定时器
# <翻译结束>


# <翻译开始>
func (e *Element) KillXCTimer(nIDEvent
事件ID
# <翻译结束>

# <翻译开始>
func (e *Element) KillXCTimer
X关闭炫彩定时器
# <翻译结束>


# <翻译开始>
func (e *Element) SetToolTip(pText
内容
# <翻译结束>

# <翻译开始>
func (e *Element) SetToolTip
X置工具提示
# <翻译结束>


# <翻译开始>
func (e *Element) SetToolTipEx(pText string, nTextAlign
对齐方式
# <翻译结束>

# <翻译开始>
func (e *Element) SetToolTipEx(pText
内容
# <翻译结束>

# <翻译开始>
func (e *Element) SetToolTipEx
X置工具提示EX
# <翻译结束>


# <翻译开始>
func (e *Element) GetToolTip
X取工具提示
# <翻译结束>


# <翻译开始>
func (e *Element) PopupToolTip(x int, y
Y坐标
# <翻译结束>

# <翻译开始>
func (e *Element) PopupToolTip(x
X坐标
# <翻译结束>

# <翻译开始>
func (e *Element) PopupToolTip
X弹出工具提示
# <翻译结束>


# <翻译开始>
func (e *Element) AdjustLayout(nAdjustNo
流水号
# <翻译结束>

# <翻译开始>
func (e *Element) AdjustLayout
X调整布局
# <翻译结束>


# <翻译开始>
func (e *Element) AdjustLayoutEx(nFlags xcc.AdjustLayout_, nAdjustNo
流水号
# <翻译结束>

# <翻译开始>
func (e *Element) AdjustLayoutEx(nFlags
识位
# <翻译结束>

# <翻译开始>
func (e *Element) AdjustLayoutEx
X调整布局EX
# <翻译结束>


# <翻译开始>
func (e *Element) GetAlpha
X取透明度
# <翻译结束>


# <翻译开始>
func (e *Element) GetPosition(pOutX, pOutY
返回Y坐标
# <翻译结束>

# <翻译开始>
func (e *Element) GetPosition(pOutX
返回X坐标
# <翻译结束>

# <翻译开始>
func (e *Element) GetPosition
X取位置
# <翻译结束>


# <翻译开始>
func (e *Element) SetSize(nWidth, nHeight int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo
流水号
# <翻译结束>

# <翻译开始>
func (e *Element) SetSize(nWidth, nHeight int32, bRedraw bool, nFlags
识位
# <翻译结束>

# <翻译开始>
func (e *Element) SetSize(nWidth, nHeight int32, bRedraw
是否重绘
# <翻译结束>

# <翻译开始>
func (e *Element) SetSize(nWidth, nHeight
高度
# <翻译结束>

# <翻译开始>
func (e *Element) SetSize(nWidth
宽度
# <翻译结束>

# <翻译开始>
func (e *Element) SetSize
X置大小
# <翻译结束>


# <翻译开始>
func (e *Element) GetSize(pOutWidth, pOutHeight
返回高度
# <翻译结束>

# <翻译开始>
func (e *Element) GetSize(pOutWidth
返回宽度
# <翻译结束>

# <翻译开始>
func (e *Element) GetSize
X取大小
# <翻译结束>


# <翻译开始>
func (e *Element) SetBkInfo(pText
字符串
# <翻译结束>

# <翻译开始>
func (e *Element) SetBkInfo
X置背景
# <翻译结束>


# <翻译开始>
func (e *Element) GetWndClientRectDPI(pRect
接收返回坐标
# <翻译结束>

# <翻译开始>
func (e *Element) GetWndClientRectDPI
X取窗口客户区坐标DPI
# <翻译结束>


# <翻译开始>
func (e *Element) PointClientToWndClientDPI(pPt
接收返回坐标点
# <翻译结束>

# <翻译开始>
func (e *Element) PointClientToWndClientDPI
X取窗口客户区坐标DPI
# <翻译结束>


# <翻译开始>
func (e *Element) RectClientToWndClientDPI(pRect
接收返回坐标
# <翻译结束>

# <翻译开始>
func (e *Element) RectClientToWndClientDPI
X客户区坐标到窗口客户区DPI
# <翻译结束>


# <翻译开始>
func (e *Element) SetFocus
X置焦点
# <翻译结束>


# <翻译开始>
func (e *Element) GetLeft
X取左边
# <翻译结束>


# <翻译开始>
func (e *Element) GetTop
X取顶边
# <翻译结束>


# <翻译开始>
func (e *Element) GetRight
X取右边
# <翻译结束>


# <翻译开始>
func (e *Element) GetBottom
X取底边
# <翻译结束>


# <翻译开始>
func (e *Element) SetLeft(x int32, bRedraw
是否重绘
# <翻译结束>

# <翻译开始>
func (e *Element) SetLeft(x
左边x坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetLeft
X置左边
# <翻译结束>


# <翻译开始>
func (e *Element) SetTop(y int32, bRedraw
是否重绘
# <翻译结束>

# <翻译开始>
func (e *Element) SetTop(y
顶边y坐标
# <翻译结束>

# <翻译开始>
func (e *Element) SetTop
X置顶边
# <翻译结束>


# <翻译开始>
func (e *Element) Event_ELEPROCE
X事件_处理过程
# <翻译结束>


# <翻译开始>
func (e *Element) Event_ELEPROCE1
X事件_处理过程1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_PAINT
X事件_绘制事件
# <翻译结束>


# <翻译开始>
func (e *Element) Event_PAINT1
X事件_绘制事件1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_PAINT_END
X事件_绘制完成
# <翻译结束>


# <翻译开始>
func (e *Element) Event_PAINT_END1
X事件_绘制完成1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_PAINT_SCROLLVIEW
X事件_滚动视图
# <翻译结束>


# <翻译开始>
func (e *Element) Event_PAINT_SCROLLVIEW1
X事件_滚动视图1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSEMOVE
X事件_鼠标移动
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSEMOVE1
X事件_鼠标移动1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSESTAY
X事件_鼠标进入
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSESTAY1
X事件_鼠标进入1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSEHOVER
X事件_鼠标悬停
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSEHOVER1
X事件_鼠标悬停1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSELEAVE
X事件_鼠标离开
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSELEAVE1
X事件_鼠标离开1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSEWHEEL
X事件_鼠标滚轮滚动
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MOUSEWHEEL1
X事件_鼠标滚轮滚动1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_LBUTTONDOWN
X事件_左键按下
# <翻译结束>


# <翻译开始>
func (e *Element) Event_LBUTTONDOWN1
X事件_左键按下1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_LBUTTONUP
X事件_左键弹起
# <翻译结束>


# <翻译开始>
func (e *Element) Event_LBUTTONUP1
X事件_左键弹起1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_RBUTTONDOWN
X事件_右键按下
# <翻译结束>


# <翻译开始>
func (e *Element) Event_RBUTTONDOWN1
X事件_右键按下1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_RBUTTONUP
X事件_右键弹起
# <翻译结束>


# <翻译开始>
func (e *Element) Event_RBUTTONUP1
X事件_右键弹起1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_LBUTTONDBCLICK
X事件_左键双击
# <翻译结束>


# <翻译开始>
func (e *Element) Event_LBUTTONDBCLICK1
X事件_左键双击1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_XC_TIMER
X事件_炫彩定时器
# <翻译结束>


# <翻译开始>
func (e *Element) Event_XC_TIMER1
X事件_炫彩定时器1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_ADJUSTLAYOUT
X事件_调整布局
# <翻译结束>


# <翻译开始>
func (e *Element) Event_ADJUSTLAYOUT1
X事件_调整布局1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_ADJUSTLAYOUT_END
X事件_调整布局完成
# <翻译结束>


# <翻译开始>
func (e *Element) Event_ADJUSTLAYOUT_END1
X事件_调整布局完成
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SETFOCUS
X事件_获得焦点
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SETFOCUS1
X事件_获得焦点1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_KILLFOCUS
X事件_失去焦点
# <翻译结束>


# <翻译开始>
func (e *Element) Event_KILLFOCUS1
X事件_失去焦点1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_DESTROY
X事件_即将销毁
# <翻译结束>


# <翻译开始>
func (e *Element) Event_DESTROY1
X事件_即将销毁1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_DESTROY_END
X事件_销毁完成
# <翻译结束>


# <翻译开始>
func (e *Element) Event_DESTROY_END1
X事件_销毁完成1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SIZE
X事件_大小改变事件
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SIZE1
X事件_大小改变事件1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SHOW
X事件_显示隐藏事件
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SHOW1
X事件_显示隐藏事件1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SETFONT
X事件_设置字体
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SETFONT1
X事件_设置字体1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_KEYDOWN
X事件_按键按下
# <翻译结束>


# <翻译开始>
func (e *Element) Event_KEYDOWN1
X事件_按键按下1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_KEYUP
X事件_按键弹起
# <翻译结束>


# <翻译开始>
func (e *Element) Event_KEYUP1
X事件_按键弹起1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_CHAR
X事件_CHAR
# <翻译结束>


# <翻译开始>
func (e *Element) Event_CHAR1
X事件_CHAR1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SETCAPTURE
X事件_设置鼠标捕获
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SETCAPTURE1
X事件_设置鼠标捕获1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_KILLCAPTURE
X事件_失去鼠标捕获
# <翻译结束>


# <翻译开始>
func (e *Element) Event_KILLCAPTURE1
X事件_失去鼠标捕获1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SETCURSOR
X事件_设置鼠标光标
# <翻译结束>


# <翻译开始>
func (e *Element) Event_SETCURSOR1
X事件_设置鼠标光标1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_DROPFILES
X事件_文件拖放
# <翻译结束>


# <翻译开始>
func (e *Element) Event_DROPFILES1
X事件_文件拖放事件1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_TOOLTIP_POPUP
X事件_工具提示弹出
# <翻译结束>


# <翻译开始>
func (e *Element) Event_TOOLTIP_POPUP1
X事件_工具提示弹出1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_SELECT
X事件_弹出菜单项被选择
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_POPUP
X事件_菜单弹出
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_EXIT
X事件_菜单退出
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_POPUP_WND
X事件_菜单弹出窗口
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_DRAW_BACKGROUND
X事件_绘制菜单背景
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_DRAWITEM
X事件_绘制菜单项事件
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_SELECT1
X事件_弹出菜单项被选择1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_POPUP1
X事件_菜单弹出1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_EXIT1
X事件_菜单退出1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_POPUP_WND1
X菜单弹出窗口1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_DRAW_BACKGROUND1
X事件_绘制菜单背景1
# <翻译结束>


# <翻译开始>
func (e *Element) Event_MENU_DRAWITEM1
X事件_绘制菜单项1
# <翻译结束>

