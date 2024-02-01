package xcc

// XC_NAME

const (
	XC_NAME1 = "name1"
	XC_NAME2 = "name2"
	XC_NAME3 = "name3"
	XC_NAME4 = "name4"
	XC_NAME5 = "name5"
	XC_NAME6 = "name6"
)

// 特殊ID

const (
	I常量_根节点    = 0  // 根节点
	I常量_错误     = -1 // ID错误
	I常量_插入开始位置 = -2 // 插入开始位置
	I常量_插入末尾位置 = -3 // 插入末尾位置
)

// I常量_窗口形式_ 是显示窗口的形式.
type I常量_窗口形式_ int32

const (
	I常量_窗口形式_隐藏            I常量_窗口形式_ = iota // 窗口_隐藏 ，活动状态给另一个窗口
	I常量_窗口形式_显示且激活_原尺寸位置                    // 窗口_显示且激活 用原来的大小和位置显示一个窗口，并将其激活
	I常量_窗口形式_最小化并激活                         // 窗口_最小化并激活 ，并将其激活
	I常量_窗口形式_最大化并激活                         // 窗口_最大化并激活 ，并将其激活
	I常量_窗口形式_显示且不激活                         // 窗口_显示且不激活 用最近的大小和位置显示一个窗口，同时不激活该窗口
	I常量_窗口形式_显示并激活                          // 窗口_显示并激活 用当前的大小和位置显示一个窗口，并将其激活
	I常量_窗口形式_最小化窗口                          // 窗口_最小化窗口，活动状态给另一个窗口
	I常量_窗口形式_最小化且不改变活动                      // 窗口_最小化且不改变活动 ，同时不改变活动窗口
	I常量_窗口形式_显示并按原状态                        // 窗口_显示并按原状态  原来的状态显示
	I常量_窗口形式_激活并还原显示窗口                      // 窗口_激活并还原显示窗口
	I常量_窗口形式_按默认状态显示                        // 窗口_按默认状态显示
	I常量_窗口形式_FORCEMINIMIZE                  // 最小化窗口，即使拥有该窗口的线程被挂起也会最小化。在从其他线程最小化窗口时才使用这个参数
)

// AdjustLayout_ 是调整布局标识位.
type AdjustLayout_ int

const (
	AdjustLayout_No   AdjustLayout_ = iota // 不调整布局
	AdjustLayout_All                       // 强制调整自身和子对象布局
	AdjustLayout_Self                      // 只调整自身布局, 不调整子对象布局
)

// I常量_窗口透明_ 是炫彩窗口透明标识.
type I常量_窗口透明_ int

const (
	I常量_窗口透明_默认     I常量_窗口透明_ = iota // I常量_窗口透明_默认, 不透明
	I常量_窗口透明_透明                      // I常量_窗口透明_透明, 带透明通道, 异型
	I常量_窗口透明_阴影                      // I常量_窗口透明_阴影, 带透明通道, 边框阴影, 窗口透明或半透明
	I常量_窗口透明_简单                      // I常量_窗口透明_透明, 不带透明通道, 指定半透明度, 指定透明色
	I常量_窗口透明_WIN7玻璃                  // I常量_窗口透明_WIN7玻璃, 需要WIN7开启特效, 当前未启用
)

// I常量_布局轴对齐_ 是布局轴对齐.
type I常量_布局轴对齐_ int

const (
	I常量_布局轴对齐_自动            I常量_布局轴对齐_ = iota // 无
	Layout_Align_Axis_Start                   // 水平布局(顶部), 垂直布局(左侧)
	I常量_布局轴对齐_居中                              // 居中
	Layout_Align_Axis_End                     // 水平布局(底部), 垂直布局(右侧)
)

// I常量_布局大小_ 是布局大小类型.
type I常量_布局大小_ int

const (
	I常量_布局大小_固定   I常量_布局大小_ = iota // I常量_布局大小_固定
	I常量_布局大小_填充父                   // I常量_布局大小_填充父
	I常量_布局大小_自动大小                  // I常量_布局大小_自动大小, 根据内容计算大小
	I常量_布局大小_按照比例                  // I常量_布局大小_按照比例分配剩余空间
	I常量_布局大小_百分比                   // I常量_布局大小_百分比
	I常量_布局大小_不使用                   // I常量_布局大小_不使用
)

// I常量_布局对齐_ 是布局_对齐.
type I常量_布局对齐_ int

const (
	I常量_布局对齐_左侧 I常量_布局对齐_ = iota // I常量_布局对齐_左侧
	I常量_布局对齐_顶部                  // I常量_布局对齐_顶部
	I常量_布局对齐_右侧                  // I常量_布局对齐_右侧
	I常量_布局对齐_底部                  // I常量_布局对齐_底部
	I常量_布局对齐_居中                  // I常量_布局对齐_居中
	I常量_布局对齐_等距                  // I常量_布局对齐_等距
)

// XC_OBJECT_STYLE 对象样式(用于区分外观).
type XC_OBJECT_STYLE int

const (
	Button_Style_Default XC_OBJECT_STYLE = iota // 默认风格
	Button_Style_Radio                          // 单选按钮
	Button_Style_Check                          // 复选按钮
	Button_Style_Icon                           // 图标按钮
	Button_Style_Expand                         // 展开收缩按钮

	Button_Style_Close // 关闭按钮
	Button_Style_Max   // 最大化按钮
	Button_Style_Min   // 最小化按钮

	Button_Style_Scrollbar_Left     // 水平滚动条-左按钮
	Button_Style_Scrollbar_Right    // 水平滚动条-右按钮
	Button_Style_Scrollbar_Up       // 垂直滚动条-上按钮
	Button_Style_Scrollbar_Down     // 垂直滚动条-下按钮
	Button_Style_Scrollbar_Slider_H // 水平滚动条-滑块
	Button_Style_Scrollbar_Slider_V // 垂直滚动条-滑块

	Button_Style_TabBar // TabBar-按钮
	Button_Style_Slider // 滑动条-滑块

	Button_Style_ToolBar       // 工具条-按钮
	Button_Style_ToolBar_Left  // 工具条-左滚动按钮
	Button_Style_ToolBar_Right // 工具条-右滚动按钮

	Button_Style_Pane_Close // 窗格-关闭按钮
	Button_Style_Pane_Lock  // 窗格-锁定按钮
	Button_Style_Pane_Menu  // 窗格-下拉菜单按钮

	Button_Style_Pane_Dock_Left   // 窗格-码头按钮左
	Button_Style_Pane_Dock_Top    // 窗格-码头按钮上
	Button_Style_Pane_Dock_Right  // 窗格-码头按钮右
	Button_Style_Pane_Dock_Bottom // 窗格-码头按钮下

	Element_Style_FrameWnd_Dock_Left   // 框架窗口-停靠码头左
	Element_Style_FrameWnd_Dock_Top    // 框架窗口-停靠码头上
	Element_Style_FrameWnd_Dock_Right  // 框架窗口-停靠码头右
	Element_Style_FrameWnd_Dock_Bottom // 框架窗口-停靠码头下

	Element_Style_ToolBar_Separator // 工具条-分割线
	ListBox_Style_ComboBox          // 组合框-下拉列表框, 下拉组合框弹出的ListBox
)

// I常量_弹出消息框_ 弹出消息框.
type I常量_弹出消息框_ int

const (
	I常量_弹出消息框_其他   I常量_弹出消息框_ = 0    // I常量_弹出消息框_其他
	I常量_弹出消息框_确定按钮 I常量_弹出消息框_ = 0x01 // I常量_弹出消息框_确定按钮
	I常量_弹出消息框_取消按钮 I常量_弹出消息框_ = 0x02 // I常量_弹出消息框_取消按钮

	I常量_弹出消息框_应用程序 I常量_弹出消息框_ = 0x01000 // 图标 I常量_弹出消息框_应用程序 IDI_APPLICATION
	I常量_弹出消息框_信息   I常量_弹出消息框_ = 0x02000 // 图标 I常量_弹出消息框_信息 IDI_ASTERISK
	I常量_弹出消息框_问询   I常量_弹出消息框_ = 0x04000 // 图标 I常量_弹出消息框_问询  /帮助/提问 IDI_QUESTION
	I常量_弹出消息框_错误   I常量_弹出消息框_ = 0x08000 // 图标 I常量_弹出消息框_错误   /拒绝/禁止 IDI_ERROR
	I常量_弹出消息框_警告   I常量_弹出消息框_ = 0x10000 // 图标 I常量_弹出消息框_警告 IDI_WARNING
	I常量_弹出消息框_盾牌   I常量_弹出消息框_ = 0x20000 // 图标 I常量_弹出消息框_盾牌 IDI_SHIELD
)

// Common_State3_ 普通三种状态.
type Common_State3_ int32

const (
	Common_State3_Leave Common_State3_ = iota // 离开
	Common_State3_Stay                        // 停留
	Common_State3_Down                        // 按下
)

// I常量_按钮状态_ 按钮状态.
type I常量_按钮状态_ int

const (
	I常量_按钮状态_离开 I常量_按钮状态_ = iota // I常量_按钮状态_离开
	I常量_按钮状态_停留                  // I常量_按钮状态_停留
	I常量_按钮状态_按下                  // I常量_按钮状态_按下
	I常量_按钮状态_选中                  // I常量_按钮状态_选中
	I常量_按钮状态_禁用                  // I常量_按钮状态_禁用
)

// I常量_文本对齐_ 文本对齐.
type I常量_文本对齐_ int

const (
	I常量_文本对齐_左    I常量_文本对齐_ = 0      // I常量_文本对齐_左
	I常量_文本对齐_垂直顶  I常量_文本对齐_ = 0      // I常量_文本对齐_垂直顶
	I常量_文本对齐_内部保留 I常量_文本对齐_ = 0x4000 // I常量_文本对齐_内部保留
	I常量_文本对齐_水平居中 I常量_文本对齐_ = 0x1    // I常量_文本对齐_水平居中
	I常量_文本对齐_右    I常量_文本对齐_ = 0x2    // I常量_文本对齐_右
	I常量_文本对齐_垂直居中 I常量_文本对齐_ = 0x4    // I常量_文本对齐_垂直居中
	I常量_文本对齐_垂直底  I常量_文本对齐_ = 0x8    // I常量_文本对齐_垂直底

	TextFormatFlag_DirectionRightToLeft  I常量_文本对齐_ = 0x10   // 从右向左 顺序显示文本
	TextFormatFlag_NoWrap                I常量_文本对齐_ = 0x20   // 禁止换行
	TextFormatFlag_DirectionVertical     I常量_文本对齐_ = 0x40   // 垂直显示  文本
	TextFormatFlag_NoFitBlackBox         I常量_文本对齐_ = 0x80   // 允许部分字符延伸该字符串的布局矩形。默认情况下，将重新定位字符以避免任何延伸
	TextFormatFlag_DisplayFormatControl  I常量_文本对齐_ = 0x100  // 控制字符（如从左到右标记）随具有代表性的标志符号一起显示在输出中。
	TextFormatFlag_NoFontFallback        I常量_文本对齐_ = 0x200  // 对于请求的字体中不支持的字符，禁用回退到可选字体。缺失的任何字符都用缺失标志符号的字体显示，通常是一个空的方块
	TextFormatFlag_MeasureTrailingSpaces I常量_文本对齐_ = 0x400  // 包括每一行结尾处的尾随空格。在默认情况下，MeasureString 方法返回的边框都将排除每一行结尾处的空格。设置此标记以便在测定时将空格包括进去
	TextFormatFlag_LineLimit             I常量_文本对齐_ = 0x800  // 如果内容显示高度不够一行,那么不显示
	TextFormatFlag_NoClip                I常量_文本对齐_ = 0x1000 // 允许显示标志符号的伸出部分和延伸到边框外的未换行文本。在默认情况下，延伸到边框外侧的所有文本和标志符号部分都被剪裁

	TextTrimming_None              I常量_文本对齐_ = 0       // 不使用去尾
	TextTrimming_Character         I常量_文本对齐_ = 0x40000 // 以字符为单位去尾
	TextTrimming_Word              I常量_文本对齐_ = 0x80000 // 以单词为单位去尾
	TextTrimming_EllipsisCharacter I常量_文本对齐_ = 0x8000  // 以字符为单位去尾,省略部分使用且略号表示
	TextTrimming_EllipsisWord      I常量_文本对齐_ = 0x10000 // 以单词为单位去尾,
	TextTrimming_EllipsisPath      I常量_文本对齐_ = 0x20000 // 略去字符串中间部分，保证字符的首尾都能够显示
)

// I常量_按钮图标对齐_ 按钮图标对齐方式.
type I常量_按钮图标对齐_ int

const (
	I常量_按钮图标对齐_左 I常量_按钮图标对齐_ = iota // I常量_按钮图标对齐_左
	I常量_按钮图标对齐_顶                    // I常量_按钮图标对齐_顶
	I常量_按钮图标对齐_右                    // I常量_按钮图标对齐_右
	I常量_按钮图标对齐_底                    // I常量_按钮图标对齐_底
)

// I常量_窗口样式_ 窗口样式.
type I常量_窗口样式_ int

const (
	I常量_窗口样式_无     I常量_窗口样式_ = 0x0000 // I常量_窗口样式_无 什么也没有
	I常量_窗口样式_标题栏   I常量_窗口样式_ = 0x0001 // I常量_窗口样式_标题栏
	I常量_窗口样式_边框    I常量_窗口样式_ = 0x0002 // I常量_窗口样式_边框  ,如果没有指定,那么边框大小为0
	I常量_窗口样式_居中    I常量_窗口样式_ = 0x0004 // I常量_窗口样式_居中
	I常量_窗口样式_拖动边框  I常量_窗口样式_ = 0x0008 // I常量_窗口样式_拖动边框
	I常量_窗口样式_拖动窗口  I常量_窗口样式_ = 0x0010 // I常量_窗口样式_拖动窗口
	I常量_窗口样式_允许最大化 I常量_窗口样式_ = 0x0020 // I常量_窗口样式_允许最大化

	I常量_窗口样式_图标     I常量_窗口样式_ = 0x0040 // I常量_窗口样式_图标
	I常量_窗口样式_标题     I常量_窗口样式_ = 0x0080 // I常量_窗口样式_标题
	I常量_窗口样式_按钮_最小化 I常量_窗口样式_ = 0x0100 // I常量_窗口样式_按钮_最小化
	I常量_窗口样式_按钮_最大化 I常量_窗口样式_ = 0x0200 // I常量_窗口样式_按钮_最大化
	I常量_窗口样式_按钮_关闭  I常量_窗口样式_ = 0x0400 // I常量_窗口样式_按钮_关闭

	//I常量_窗口样式_默认 窗口样式-控制按钮: 居中 图标, 标题, 关闭按钮, 最大化按钮, 最小化按钮
	I常量_窗口样式_默认 = I常量_窗口样式_标题栏 | I常量_窗口样式_边框 | I常量_窗口样式_居中 | I常量_窗口样式_拖动边框 | I常量_窗口样式_允许最大化 | I常量_窗口样式_图标 | I常量_窗口样式_标题 | I常量_窗口样式_按钮_最小化 | I常量_窗口样式_按钮_最大化 | I常量_窗口样式_按钮_关闭 // 窗口样式-控制按钮: 居中 图标, 标题, 关闭按钮, 最大化按钮, 最小化按钮

	//I常量_窗口样式_简单
	I常量_窗口样式_简单 = I常量_窗口样式_标题栏 | I常量_窗口样式_边框 | I常量_窗口样式_居中 | I常量_窗口样式_拖动边框 | I常量_窗口样式_允许最大化 // 窗口样式-简单: 居中

	//I常量_窗口样式_弹出
	I常量_窗口样式_弹出 = I常量_窗口样式_标题栏 | I常量_窗口样式_边框 | I常量_窗口样式_居中 | I常量_窗口样式_拖动边框 | I常量_窗口样式_允许最大化 | I常量_窗口样式_图标 |
		I常量_窗口样式_标题 // 窗口样式-弹出窗口: 图标, 标题, 关闭按钮

	I常量_窗口样式_模态 = I常量_窗口样式_标题栏 | I常量_窗口样式_边框 | I常量_窗口样式_居中 | I常量_窗口样式_图标 | I常量_窗口样式_标题 | I常量_窗口样式_按钮_关闭 // 模态窗口样式-控制按钮: 居中, 图标, 标题, 关闭按钮

	I常量_窗口样式_模态_简单 = I常量_窗口样式_标题栏 | I常量_窗口样式_边框 | I常量_窗口样式_居中 // 模态窗口样式-简单: 居中
)

// I常量_对象句柄类型_ 对象句柄类型.
type I常量_对象句柄类型_ int

const (
	I常量_对象句柄类型_错误类型    I常量_对象句柄类型_ = -1 // I常量_对象句柄类型_错误类型
	XC_NOTHING         I常量_对象句柄类型_ = 0
	I常量_对象句柄类型_窗口      I常量_对象句柄类型_ = 1  // I常量_对象句柄类型_窗口
	I常量_对象句柄类型_模态窗口    I常量_对象句柄类型_ = 2  // I常量_对象句柄类型_模态窗口
	I常量_对象句柄类型_框架窗口    I常量_对象句柄类型_ = 3  // I常量_对象句柄类型_框架窗口
	I常量_对象句柄类型_浮动窗口    I常量_对象句柄类型_ = 4  // I常量_对象句柄类型_浮动窗口
	I常量_对象句柄类型_组合框     I常量_对象句柄类型_ = 11 // I常量_对象句柄类型_组合框 comboBoxWindow_ 组合框弹出下拉列表窗口
	I常量_对象句柄类型_弹出菜单主窗口 I常量_对象句柄类型_ = 12 // I常量_对象句柄类型_弹出菜单主窗口 popupMenuWindow_ 弹出菜单主窗口
	I常量_对象句柄类型_弹出菜单子窗口 I常量_对象句柄类型_ = 13 // I常量_对象句柄类型_弹出菜单子窗口 popupMenuChildWindow_ 弹出菜单子窗口
	I常量_对象句柄类型_可视对象    I常量_对象句柄类型_ = 19 // I常量_对象句柄类型_可视对象
	I常量_对象句柄类型_窗口组件    I常量_对象句柄类型_ = 20 // I常量_对象句柄类型_窗口组件
	I常量_对象句柄类型_基础元素    I常量_对象句柄类型_ = 21 // I常量_对象句柄类型_基础元素
	I常量_对象句柄类型_布局元素    I常量_对象句柄类型_ = 53 // I常量_对象句柄类型_布局元素
	I常量_对象句柄类型_流式布局    I常量_对象句柄类型_ = 54 // I常量_对象句柄类型_流式布局
	I常量_对象句柄类型_按钮      I常量_对象句柄类型_ = 22 // I常量_对象句柄类型_按钮
	I常量_对象句柄类型_编辑框     I常量_对象句柄类型_ = 45 // I常量_对象句柄类型_编辑框
	I常量_对象句柄类型_代码编辑框   I常量_对象句柄类型_ = 46 // I常量_对象句柄类型_代码编辑框

	I常量_对象句柄类型_富文本编辑框 I常量_对象句柄类型_ = 23 // I常量_对象句柄类型_富文本编辑框
	I常量_对象句柄类型_下拉组合框  I常量_对象句柄类型_ = 24 // I常量_对象句柄类型_下拉组合框
	I常量_对象句柄类型_滚动条    I常量_对象句柄类型_ = 25 // I常量_对象句柄类型_滚动条
	I常量_对象句柄类型_滚动视图   I常量_对象句柄类型_ = 26 // I常量_对象句柄类型_滚动视图
	I常量_对象句柄类型_列表     I常量_对象句柄类型_ = 27 // I常量_对象句柄类型_列表
	I常量_对象句柄类型_列表框    I常量_对象句柄类型_ = 28 // I常量_对象句柄类型_列表框
	I常量_对象句柄类型_列表视图   I常量_对象句柄类型_ = 29 // I常量_对象句柄类型_列表视图   ,大图标
	I常量_对象句柄类型_列表树    I常量_对象句柄类型_ = 30 // I常量_对象句柄类型_列表树
	I常量_对象句柄类型_菜单条    I常量_对象句柄类型_ = 31 // I常量_对象句柄类型_菜单条
	I常量_对象句柄类型_滑动条    I常量_对象句柄类型_ = 32 // I常量_对象句柄类型_滑动条
	I常量_对象句柄类型_进度条    I常量_对象句柄类型_ = 33 // I常量_对象句柄类型_进度条
	I常量_对象句柄类型_工具条    I常量_对象句柄类型_ = 34 // I常量_对象句柄类型_工具条
	I常量_对象句柄类型_月历卡片   I常量_对象句柄类型_ = 35 // I常量_对象句柄类型_月历卡片
	I常量_对象句柄类型_日期时间   I常量_对象句柄类型_ = 36 // I常量_对象句柄类型_日期时间
	I常量_对象句柄类型_属性网格   I常量_对象句柄类型_ = 37 // I常量_对象句柄类型_属性网格
	I常量_对象句柄类型_颜色选择框  I常量_对象句柄类型_ = 38 // I常量_对象句柄类型_颜色选择框
	I常量_对象句柄类型_设置编辑框  I常量_对象句柄类型_ = 39 // I常量_对象句柄类型_设置编辑框
	I常量_对象句柄类型_tab条   I常量_对象句柄类型_ = 40 // I常量_对象句柄类型_tab条
	I常量_对象句柄类型_文本链接按钮 I常量_对象句柄类型_ = 41 // I常量_对象句柄类型_文本链接按钮

	I常量_对象句柄类型_窗格       I常量_对象句柄类型_ = 42 // I常量_对象句柄类型_窗格
	I常量_对象句柄类型_窗格拖动分割条  I常量_对象句柄类型_ = 43 // I常量_对象句柄类型_窗格拖动分割条
	I常量_对象句柄类型_菜单条上的按钮  I常量_对象句柄类型_ = 44 // I常量_对象句柄类型_菜单条上的按钮
	I常量_对象句柄类型_文件选择编辑框  I常量_对象句柄类型_ = 50 // EditFile I常量_对象句柄类型_文件选择编辑框
	I常量_对象句柄类型_文件夹选择编辑框 I常量_对象句柄类型_ = 51 // EditFolder I常量_对象句柄类型_文件夹选择编辑框
	I常量_对象句柄类型_列表头元素    I常量_对象句柄类型_ = 52 // I常量_对象句柄类型_列表头元素

	I常量_对象句柄类型_形状对象     I常量_对象句柄类型_ = 61 // I常量_对象句柄类型_形状对象
	I常量_对象句柄类型_形状对象_文本  I常量_对象句柄类型_ = 62 // I常量_对象句柄类型_形状对象_文本
	I常量_对象句柄类型_形状对象_图片  I常量_对象句柄类型_ = 63 // I常量_对象句柄类型_形状对象_图片
	I常量_对象句柄类型_形状对象_矩形  I常量_对象句柄类型_ = 64 // I常量_对象句柄类型_形状对象_矩形
	I常量_对象句柄类型_形状对象_圆   I常量_对象句柄类型_ = 65 // I常量_对象句柄类型_形状对象_圆
	I常量_对象句柄类型_形状对象_直线  I常量_对象句柄类型_ = 66 // I常量_对象句柄类型_形状对象_直线
	I常量_对象句柄类型_形状对象_组框  I常量_对象句柄类型_ = 67 // I常量_对象句柄类型_形状对象_组框
	I常量_对象句柄类型_形状对象_GIF I常量_对象句柄类型_ = 68 // I常量_对象句柄类型_形状对象_GIF
	I常量_对象句柄类型_形状对象_表格  I常量_对象句柄类型_ = 69 // I常量_对象句柄类型_形状对象_表格

	I常量_对象句柄类型_弹出菜单    I常量_对象句柄类型_ = 81            // I常量_对象句柄类型_弹出菜单
	I常量_对象句柄类型_图片      I常量_对象句柄类型_ = 82            // I常量_对象句柄类型_图片
	XC_IMAGE_TEXTURE               = I常量_对象句柄类型_图片 // 图片纹理, 图片源, 图片素材
	I常量_对象句柄类型_绘图操作    I常量_对象句柄类型_ = 83            // I常量_对象句柄类型_绘图操作
	I常量_对象句柄类型_炫彩字体    I常量_对象句柄类型_ = 84            // I常量_对象句柄类型_炫彩字体
	I常量_对象句柄类型_flash   I常量_对象句柄类型_ = 85            // I常量_对象句柄类型_flash
	XC_PANE_CELL       I常量_对象句柄类型_ = 86            // ...
	XC_WEB             I常量_对象句柄类型_ = 87            // ...
	I常量_对象句柄类型_图片帧     I常量_对象句柄类型_ = 88            // I常量_对象句柄类型_图片帧, 指定图片的渲染属性
	I常量_对象句柄类型_SVG矢量图形 I常量_对象句柄类型_ = 89            // I常量_对象句柄类型_SVG矢量图形

	I常量_对象句柄类型_布局对象                 I常量_对象句柄类型_ = 101 // I常量_对象句柄类型_布局对象 LayoutObject
	XC_ADAPTER                      I常量_对象句柄类型_ = 102 // ...
	I常量_对象句柄类型_数据适配器AdapterTable    I常量_对象句柄类型_ = 103 // I常量_对象句柄类型_数据适配器AdapterTable
	I常量_对象句柄类型_数据适配器AdapterTree     I常量_对象句柄类型_ = 104 // I常量_对象句柄类型_数据适配器AdapterTree
	I常量_对象句柄类型_数据适配器AdapterListView I常量_对象句柄类型_ = 105 // I常量_对象句柄类型_数据适配器AdapterListView
	I常量_对象句柄类型_数据适配器AdapterMap      I常量_对象句柄类型_ = 106 // I常量_对象句柄类型_数据适配器AdapterMap

	I常量_对象句柄类型_背景管理器 I常量_对象句柄类型_ = 116 // I常量_对象句柄类型_背景管理器

	// 无实体对象, 只是用来判断布局

	XC_LAYOUT_LISTVIEW     I常量_对象句柄类型_ = 111
	XC_LAYOUT_LIST         I常量_对象句柄类型_ = 112
	XC_LAYOUT_OBJECT_GROUP I常量_对象句柄类型_ = 113
	XC_LAYOUT_OBJECT_ITEM  I常量_对象句柄类型_ = 114
	XC_LAYOUT_PANEL        I常量_对象句柄类型_ = 115

	// 无实体对象, 只是用来判断类型

	XC_LIST_ITEM      I常量_对象句柄类型_ = 121 // 列表项模板 list_item
	XC_LISTVIEW_GROUP I常量_对象句柄类型_ = 122
	XC_LISTVIEW_ITEM  I常量_对象句柄类型_ = 123
	XC_LAYOUT_BOX     I常量_对象句柄类型_ = 124

	XC_ANIMATION_SEQUENCE I常量_对象句柄类型_ = 131 // 动画序列
	XC_ANIMATION_GROUP    I常量_对象句柄类型_ = 132 // 动画同步组
	XC_ANIMATION_ITEM     I常量_对象句柄类型_ = 133 // 动画项
)

// I常量_对象扩展类型_ 对象扩展类型(功能扩展).
type I常量_对象扩展类型_ int

const (
	I常量_对象扩展类型_默认        I常量_对象扩展类型_ = iota // I常量_对象扩展类型_默认
	I常量_对象扩展类型_单选按钮                         // I常量_对象扩展类型_单选按钮
	I常量_对象扩展类型_复选按钮                         // I常量_对象扩展类型_复选按钮
	I常量_对象扩展类型_窗口关闭按钮                       // I常量_对象扩展类型_窗口关闭按钮
	I常量_对象扩展类型_窗口最小化按钮                      // I常量_对象扩展类型_窗口最小化按钮
	I常量_对象扩展类型_窗口最大化还原按钮                    // I常量_对象扩展类型_窗口最大化还原按钮

	I常量_对象扩展类型_布局元素 // I常量_对象扩展类型_布局元素, 启用布局功能的元素

	I常量_对象扩展类型_错误类型 I常量_对象扩展类型_ = -1 // I常量_对象扩展类型_错误类型
)

// Element_Position_ UI元素位置.
type Element_Position_ int

const (
	Element_Position_No     Element_Position_ = 0x00 // 无效
	Element_Position_Left   Element_Position_ = 0x01 // 左边
	Element_Position_Top    Element_Position_ = 0x02 // 上边
	Element_Position_Right  Element_Position_ = 0x04 // 右边
	Element_Position_Bottom Element_Position_ = 0x08 // 下边
)

// I常量_位置标识_ 位置标识.
type I常量_位置标识_ int

const (
	I常量_位置标识_左   I常量_位置标识_ = iota // I常量_位置标识_左
	I常量_位置标识_上                    // I常量_位置标识_上
	I常量_位置标识_右                    // I常量_位置标识_右
	I常量_位置标识_下                    // I常量_位置标识_下
	I常量_位置标识_左上角                  // I常量_位置标识_左上角
	I常量_位置标识_左下角                  // I常量_位置标识_左下角
	I常量_位置标识_右上角                  // I常量_位置标识_右上角
	I常量_位置标识_右下角                  // I常量_位置标识_右下角
	I常量_位置标识_中心                   // I常量_位置标识_中心
)

// I常量_框架窗口对齐_ 窗格_对齐.
type I常量_框架窗口对齐_ int

const (
	I常量_框架窗口对齐_左  I常量_框架窗口对齐_ = iota // I常量_框架窗口对齐_左
	I常量_框架窗口对齐_顶                     // I常量_框架窗口对齐_顶
	I常量_框架窗口对齐_右                     // I常量_框架窗口对齐_右
	I常量_框架窗口对齐_底                     // I常量_框架窗口对齐_底
	I常量_框架窗口对齐_居中                    // I常量_框架窗口对齐_居中
	I常量_框架窗口对齐_错误 = -1               // I常量_框架窗口对齐_错误
)

// I常量_菜单_标识_ 菜单项标识.
type I常量_菜单_标识_ int32

const (
	I常量_菜单标识_正常           I常量_菜单_标识_ = 0x00 //I常量_菜单标识_正常
	Menu_Item_Flag_Select I常量_菜单_标识_ = 0x01 // I常量_菜单标识_选择或鼠标停留
	I常量_菜单标识_选择或鼠标停留      I常量_菜单_标识_ = 0x01 // I常量_菜单标识_选择或鼠标停留
	I常量_菜单标识_勾选           I常量_菜单_标识_ = 0x02 // I常量_菜单标识_勾选
	I常量_菜单标识_弹出           I常量_菜单_标识_ = 0x04 // I常量_菜单标识_弹出
	I常量_菜单标识_分隔栏          I常量_菜单_标识_ = 0x08 // I常量_菜单标识_分隔栏   ID号任意, ID号被忽略
	I常量_菜单标识_禁用           I常量_菜单_标识_ = 0x10 // I常量_菜单标识_禁用
)

// I常量_菜单弹出方向_ 弹出菜单方向.
type I常量_菜单弹出方向_ int

const (
	I常量_菜单弹出方向_左上角 I常量_菜单弹出方向_ = iota // I常量_菜单弹出方向_左上角
	I常量_菜单弹出方向_左下角                    // I常量_菜单弹出方向_左下角
	I常量_菜单弹出方向_右上角                    // I常量_菜单弹出方向_右上角
	I常量_菜单弹出方向_右下角                    // I常量_菜单弹出方向_右下角
	I常量_菜单弹出方向_左居中                    // I常量_菜单弹出方向_左居中
	I常量_菜单弹出方向_上居中                    // I常量_菜单弹出方向_上居中
	I常量_菜单弹出方向_右居中                    // I常量_菜单弹出方向_右居中
	I常量_菜单弹出方向_下居中                    // I常量_菜单弹出方向_下居中
)

// I常量_组合框状态_ ComboBox状态.
type I常量_组合框状态_ int

const (
	I常量_组合框状态_鼠标离开 I常量_组合框状态_ = iota // I常量_组合框状态_鼠标离开
	I常量_组合框状态_鼠标停留                   // I常量_组合框状态_鼠标停留
	I常量_组合框状态_按下                     // I常量_组合框状态_按下
)

// Adapter_Date_Type_ 数据适配器数据类型.
type Adapter_Date_Type_ int

const (
	Adapter_Date_Type_Int    Adapter_Date_Type_ = iota // 整形
	Adapter_Date_Type_Float                            // 浮点型
	Adapter_Date_Type_String                           // 字符串
	Adapter_Date_Type_Image                            // 图片
	Adapter_Date_Type_Error  Adapter_Date_Type_ = -1
)

// Chat_Flag_ Edit 聊天气泡对齐方式.
type Chat_Flag_ int

const (
	Chat_Flag_Left            Chat_Flag_ = 0x1 // 左侧
	Chat_Flag_Right           Chat_Flag_ = 0x2 // 右侧
	Chat_Flag_Center          Chat_Flag_ = 0x4 // 中间
	Chat_Flag_Next_Row_Bubble Chat_Flag_ = 0x8 // 下一行显示气泡
)

// I常量_编辑框类型_ 编辑框类型.
type I常量_编辑框类型_ int

const (
	I常量_编辑框类型_普通   I常量_编辑框类型_ = iota // I常量_编辑框类型_普通  编辑框, 每行的高度相同
	I常量_编辑框类型_代码编辑                   // I常量_编辑框类型_代码编辑
	I常量_编辑框类型_富文本                    // I常量_编辑框类型_富文本  编辑框, 每行的高度可能不同
	I常量_编辑框类型_聊天气泡                   // I常量_编辑框类型_聊天气泡, 每行的高度可能不同
	I常量_编辑框类型_代码表格                   // I常量_编辑框类型_代码表格, 内部使用, 每行的高度相同
)

// I常量_编辑框风格类型_ 编辑框风格类型.
type I常量_编辑框风格类型_ int32

const (
	I常量_编辑框风格类型_字体   I常量_编辑框风格类型_ = iota + 1 // I常量_编辑框风格类型_字体
	I常量_编辑框风格类型_图片                           // I常量_编辑框风格类型_图片
	I常量_编辑框风格类型_UI对象                         // I常量_编辑框风格类型_UI对象
)

// I常量_编辑框文本对齐_ 编辑框文本对齐标志.
type I常量_编辑框文本对齐_ int

const (
	I常量_编辑框文本对齐_左    I常量_编辑框文本对齐_ = 0x0 // I常量_编辑框文本对齐_左
	I常量_编辑框文本对齐_右    I常量_编辑框文本对齐_ = 0x1 // I常量_编辑框文本对齐_右
	I常量_编辑框文本对齐_水平居中 I常量_编辑框文本对齐_ = 0x2 // I常量_编辑框文本对齐_水平居中

	I常量_编辑框文本对齐_顶    I常量_编辑框文本对齐_ = 0x0 // I常量_编辑框文本对齐_顶
	I常量_编辑框文本对齐_底    I常量_编辑框文本对齐_ = 0x4 // I常量_编辑框文本对齐_底
	I常量_编辑框文本对齐_垂直居中 I常量_编辑框文本对齐_ = 0x8 // I常量_编辑框文本对齐_垂直居中
)

// I常量_表格标识_ Table 标识.
type I常量_表格标识_ int

const (
	I常量_表格标识_铺满组合单元格 I常量_表格标识_ = iota // I常量_表格标识_铺满组合单元格
	I常量_表格标识_正常最小单元格                  // I常量_表格标识_正常最小单元格
)

// GRADIENT_FILL_ 渐变填充模式.
type GRADIENT_FILL_ int

const (
	GRADIENT_FILL_RECT_H   GRADIENT_FILL_ = iota // 水平填充
	GRADIENT_FILL_RECT_V                         // 垂直填充
	GRADIENT_FILL_TRIANGLE                       // 三角形
)

// I常量_缓动类型_ 缓动类型.
type I常量_缓动类型_ int

const (
	I常量_缓动类型_慢到快    I常量_缓动类型_ = iota // I常量_缓动类型_慢到快
	I常量_缓动类型_快到慢                     // I常量_缓动类型_快到慢
	I常量_缓动类型_慢到快再到慢                  // I常量_缓动类型_慢到快再到慢
)

// I常量_缓动标识_ 缓动标识.
type I常量_缓动标识_ int

const (
	I常量_缓动标识_线性    I常量_缓动标识_ = iota // I常量_缓动标识_线性, 直线
	I常量_缓动标识_二次方曲线                  // I常量_缓动标识_二次方曲线
	I常量_缓动标识_三次方曲线                  // I常量_缓动标识_三次方曲线, 圆弧
	I常量_缓动标识_四次方曲线                  // I常量_缓动标识_四次方曲线
	I常量_缓动标识_五次方曲线                  // I常量_缓动标识_五次方曲线

	I常量_缓动标识_正弦   // I常量_缓动标识_正弦, 在末端变化
	I常量_缓动标识_突击   // I常量_缓动标识_突击, 突然一下
	I常量_缓动标识_圆环   // I常量_缓动标识_圆环, 好比绕过一个圆环
	I常量_缓动标识_强力回弹 // I常量_缓动标识_强力回弹
	I常量_缓动标识_回弹   // I常量_缓动标识_回弹, 比较缓慢
	I常量_缓动标识_弹跳   // I常量_缓动标识_弹跳, 模拟小球落地弹跳

	I常量_缓动标识_慢到快    I常量_缓动标识_ = 0x010000 // I常量_缓动标识_慢到快
	I常量_缓动标识_快到慢    I常量_缓动标识_ = 0x020000 // I常量_缓动标识_快到慢
	I常量_缓动标识_慢到快再到慢 I常量_缓动标识_ = 0x030000 // I常量_缓动标识_慢到快再到慢
)

// I常量_字体样式_ 字体样式.
type I常量_字体样式_ int32

const (
	I常量_字体样式_正常  I常量_字体样式_ = 0 // I常量_字体样式_正常
	I常量_字体样式_粗体  I常量_字体样式_ = 1 // I常量_字体样式_粗体
	I常量_字体样式_斜体  I常量_字体样式_ = 2 // I常量_字体样式_斜体
	I常量_字体样式_粗斜体 I常量_字体样式_ = 3 // I常量_字体样式_粗斜体
	I常量_字体样式_下划线 I常量_字体样式_ = 4 // I常量_字体样式_下划线
	I常量_字体样式_删除线 I常量_字体样式_ = 8 // I常量_字体样式_删除线
)

// Image_Draw_Type_ 图片绘制类型.
type Image_Draw_Type_ int

const (
	Image_Draw_Type_Default         Image_Draw_Type_ = iota // 默认
	Image_Draw_Type_Stretch                                 // 拉伸
	Image_Draw_Type_Adaptive                                // 自适应,九宫格
	Image_Draw_Type_Tile                                    // 平铺
	Image_Draw_Type_Fixed_Ratio                             // 固定比例,当图片超出显示范围时,按照原始比例压缩显示图片
	Image_Draw_Type_Adaptive_Border                         // 九宫格不绘制中间区域
)

// ListItemTemp_Type_ 列表项模板类型.
type ListItemTemp_Type_ int

const (
	ListItemTemp_Type_Tree           ListItemTemp_Type_ = 0x01                                                               // tree
	ListItemTemp_Type_ListBox        ListItemTemp_Type_ = 0x02                                                               // listBox
	ListItemTemp_Type_List_Head      ListItemTemp_Type_ = 0x04                                                               // list 列表头
	ListItemTemp_Type_List_Item      ListItemTemp_Type_ = 0x08                                                               // list 列表项
	ListItemTemp_Type_ListView_Group ListItemTemp_Type_ = 0x10                                                               // listView 列表视组
	ListItemTemp_Type_ListView_Item  ListItemTemp_Type_ = 0x20                                                               // listView 列表视项
	ListItemTemp_Type_List                              = ListItemTemp_Type_List_Head | ListItemTemp_Type_List_Item          // list (列表头)与(列表项)组合
	ListItemTemp_Type_ListView                          = ListItemTemp_Type_ListView_Group | ListItemTemp_Type_ListView_Item // listView (列表视组)与(列表视项)组合
)

// Window_Position_ 窗口位置.
type Window_Position_ int

const (
	Window_Position_Top    Window_Position_ = iota // 上
	Window_Position_Bottom                         // 下
	Window_Position_Left                           // 左
	Window_Position_Right                          // 右
	Window_Position_Body                           // body
	Window_Position_Window                         // window 整个窗口
	Window_Position_Error  Window_Position_ = -1   // 错误
)

// 窗口位置

const (
	WINDOW_TOP         = iota + 1 // I窗口位置_上
	WINDOW_BOTTOM                 // I窗口位置_下
	WINDOW_LEFT                   // I窗口位置_左
	WINDOW_RIGHT                  // I窗口位置_右
	WINDOW_TOPLEFT                // I窗口位置_左上角
	WINDOW_TOPRIGHT               // I窗口位置_右上角
	WINDOW_BOTTOMLEFT             // I窗口位置_左下角
	WINDOW_BOTTOMRIGHT            // I窗口位置_右下角
	WINDOW_CAPTION                // I窗口位置_标题栏移动窗口区域
	WINDOW_BODY
)

// List_Item_State_ List项状态.
type List_Item_State_ int32

const (
	List_Item_State_Leave  List_Item_State_ = iota // 项鼠标离开状态
	List_Item_State_Stay                           // 项鼠标停留状态
	List_Item_State_Select                         // 项选择状态
	List_Item_State_Cache                          // 缓存的项
)

// Tree_Item_State_ Tree项状态.
type Tree_Item_State_ int32

const (
	Tree_Item_State_Leave  Tree_Item_State_ = iota // 项鼠标离开状态
	Tree_Item_State_Stay                           // 项鼠标停留状态
	Tree_Item_State_Select                         // 项选择状态
)

// List_DrawItemBk_Flag_ 项背景绘制标志位(List, ListBox, ListView, Tree).
type List_DrawItemBk_Flag_ int

const (
	List_DrawItemBk_Flag_Leave       List_DrawItemBk_Flag_ = 1 << iota // 绘制鼠标离开状态项背景
	List_DrawItemBk_Flag_Stay                                          // 绘制鼠标选择状态项背景
	List_DrawItemBk_Flag_Select                                        // 绘制鼠标停留状态项项背景
	List_DrawItemBk_Flag_Group_Leave                                   // 绘制鼠标离开状态组背景, 当项为组时
	List_DrawItemBk_Flag_Group_Stay                                    // 绘制鼠标停留状态组背景, 当项为组时

	List_DrawItemBk_Flag_Line  // 列表绘制水平分割线
	List_DrawItemBk_Flag_LineV // 列表绘制垂直分割线

	List_DrawItemBk_Flag_Nothing List_DrawItemBk_Flag_ = 0 // 不绘制
)

// PropertyGrid_Item_Type_ 属性网格项类型.
type PropertyGrid_Item_Type_ int

const (
	PropertyGrid_Item_Type_Text       PropertyGrid_Item_Type_ = iota // 默认,字符串类型
	PropertyGrid_Item_Type_Edit                                      // 编辑框
	PropertyGrid_Item_Type_Edit_Color                                // 颜色选择元素
	PropertyGrid_Item_Type_Edit_File                                 // 文件选择编辑框
	PropertyGrid_Item_Type_Edit_Set                                  // 设置
	PropertyGrid_Item_Type_ComboBox                                  // 组合框
	PropertyGrid_Item_Type_Group                                     // 组
	PropertyGrid_Item_Type_Panel                                     // 面板
)

// Zorder_ Z序位置.
type Zorder_ int

const (
	Zorder_Top    Zorder_ = iota // 最上面
	Zorder_Bottom                // 最下面
	Zorder_Before                // 指定目标下面
	Zorder_After                 // 指定目标上面
)

// I常量_窗格状态_ 窗格状态.
type I常量_窗格状态_ int

const (
	I常量_窗格状态_Any  I常量_窗格状态_ = iota //I常量_窗格状态_Any
	I常量_窗格状态_锁定                    // I常量_窗格状态_锁定
	I常量_窗格状态_停靠码头                  // I常量_窗格状态_停靠码头
	I常量_窗格状态_浮动窗格                  // I常量_窗格状态_浮动窗格
	I常量_窗格状态_错误   I常量_窗格状态_ = -1   //I常量_窗格状态_错误
)

// XC_DWRITE_RENDERING_MODE_ D2D文本渲染模式.
type XC_DWRITE_RENDERING_MODE_ int

const (
	XC_DWRITE_RENDERING_MODE_DEFAULT                     XC_DWRITE_RENDERING_MODE_ = iota // 指定根据字体和大小自动确定呈现模式。
	XC_DWRITE_RENDERING_MODE_ALIASED                                                      // 指定不执行抗锯齿。 每个像素要么设置为文本的前景色，要么保留背景的颜色。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_GDI_CLASSIC                                        // 使用与别名文本相同的度量指定 ClearType 呈现。 字形只能定位在整个像素的边界上。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_GDI_NATURAL                                        // 使用使用 CLEARTYPE_NATURAL_QUALITY 创建的字体，使用与使用 GDI 的文本呈现相同的指标指定 ClearType 呈现。 与使用别名文本相比，字形度量更接近其理想值，但字形仍然位于整个像素的边界上。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_NATURAL                                            // 仅在水平维度中指定具有抗锯齿功能的 ClearType 渲染。 这通常用于中小字体大小（最多 16 ppem）。
	XC_DWRITE_RENDERING_MODE_CLEARTYPE_NATURAL_SYMMETRIC                                  // 指定在水平和垂直维度上具有抗锯齿的 ClearType 渲染。这通常用于较大的尺寸，以使曲线和对角线看起来更平滑，但会牺牲一些柔和度。
	XC_DWRITE_RENDERING_MODE_OUTLINE                                                      // 指定渲染应绕过光栅化器并直接使用轮廓。 这通常用于非常大的尺寸。
)

// MonthCal_Button_Type_ 月历元素上的按钮类型.
type MonthCal_Button_Type_ int

const (
	MonthCal_Button_Type_Today      MonthCal_Button_Type_ = iota // 今天按钮
	MonthCal_Button_Type_Last_Year                               // 上一年
	MonthCal_Button_Type_Next_Year                               // 下一年
	MonthCal_Button_Type_Last_Month                              // 上一月
	MonthCal_Button_Type_Next_Month                              // 下一月
)

// 窗格菜单 当前未使用

const (
	IDM_LOCK  = 1000000006 // 锁定
	IDM_DOCK  = 1000000007 // 停靠
	IDM_FLOAT = 1000000008 // 浮动
	IDM_HIDE  = 1000000009 // 隐藏

	Edit_Style_No = 0xFFFF // 无效样式
)

// 菜单ID, 当前未使用

const (
	IDM_CLIP      = 1000000000 // 剪切
	IDM_COPY      = 1000000001 // 复制
	IDM_PASTE     = 1000000002 // 粘贴
	IDM_DELETE    = 1000000003 // 删除
	IDM_SELECTALL = 1000000004 // 全选
	IDM_DELETEALL = 1000000005 // 清空
)

// I常量_通知消息外观_ 通知消息外观.
type I常量_通知消息外观_ int

const (
	I常量_通知消息外观_默认 I常量_通知消息外观_ = iota // I常量_通知消息外观_默认
	I常量_通知消息外观_成功                    // I常量_通知消息外观_成功
	I常量_通知消息外观_警告                    // I常量_通知消息外观_警告
	I常量_通知消息外观_消息                    // I常量_通知消息外观_消息
	I常量_通知消息外观_错误                    // I常量_通知消息外观_错误
)

// I常量_动画移动标识_ 动画移动标识.
type I常量_动画移动标识_ int

const (
	I常量_动画移动标识_X I常量_动画移动标识_ = 0x01 // I常量_动画移动标识_X 轴移动.
	I常量_动画移动标识_Y I常量_动画移动标识_ = 0x02 // I常量_动画移动标识_Y 轴移动.
)

// BkObject_Align_Flag_ 背景对象对齐方式.
type BkObject_Align_Flag_ int

const (
	BkObject_Align_Flag_Left     BkObject_Align_Flag_ = 1 << iota // 左对齐, 当设置此标识时, 外间距(margin.left)代表左侧间距; 当right未设置时,那么外间距(margin.right)代表宽度;
	BkObject_Align_Flag_Top                                       // 顶对齐, 当设置此标识时, 外间距(margin.top)代表顶部间距; 当bottom未设置时,那么外间距(margin.bottom)代表高度;
	BkObject_Align_Flag_Right                                     // 右对齐, 当设置此标识时, 外间距(margin.right)代表右侧间距; 当left未设置时,那么外间距(margin.left)代表宽度;
	BkObject_Align_Flag_Bottom                                    // 底对齐, 当设置此标识时, 外间距(margin.bottom)代表底部间距; 当top未设置时,那么外间距(margin.top)代表高度;
	BkObject_Align_Flag_Center                                    // 水平居中, 当设置此标识时, 外间距(margin.left)代表宽度;
	BkObject_Align_Flag_Center_v                                  // 垂直居中, 当设置此标识时, 外间距(margin.top)代表高度;
	BkObject_Align_Flag_No       BkObject_Align_Flag_ = 0         // 无
)

// I常量_框架窗口单元格类型_ 框架窗口单元格类型
type I常量_框架窗口单元格类型_ int

const (
	I常量_框架窗口单元格类型_无    I常量_框架窗口单元格类型_ = 0 // I常量_框架窗口单元格类型_无
	I常量_框架窗口单元格类型_窗格   I常量_框架窗口单元格类型_ = 1 // I常量_框架窗口单元格类型_窗格
	I常量_框架窗口单元格类型_窗格组  I常量_框架窗口单元格类型_ = 2 // I常量_框架窗口单元格类型_窗格组
	I常量_框架窗口单元格类型_主视图区 I常量_框架窗口单元格类型_ = 3 // I常量_框架窗口单元格类型_主视图区
	I常量_框架窗口单元格类型_上下布局 I常量_框架窗口单元格类型_ = 4 // I常量_框架窗口单元格类型_上下布局
	I常量_框架窗口单元格类型_左右布局 I常量_框架窗口单元格类型_ = 5 // I常量_框架窗口单元格类型_左右布局
)
