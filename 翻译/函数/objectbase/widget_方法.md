# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func (w *Widget) IsShow() bool {]
ff=判断显示

[func (w *Widget) Show(bShow bool) int {]
ff=显示
bShow=是否显示

[func (w *Widget) EnableLayoutControl(bEnable bool) int {]
ff=启用布局控制
bEnable=是否启用

[func (w *Widget) IsLayoutControl() bool {]
ff=是否布局控制

[func (w *Widget) GetParentEle() int {]
ff=取父元素

[func (w *Widget) GetParent() int {]
ff=取父对象

[func (w *Widget) GetHWND() uintptr {]
ff=取HWND

[func (w *Widget) GetHWINDOW() int {]
ff=取HWINDOW

[func (w *Widget) LayoutItem_EnableWrap(bWrap bool) int {]
ff=布局项启用换行
bWrap=是否换行

[func (w *Widget) LayoutItem_EnableSwap(bEnable bool) int {]
ff=布局项启用交换
bEnable=是否启用

[func (w *Widget) LayoutItem_EnableFloat(bFloat bool) int {]
ff=布局项启用浮动
bFloat=是否浮动

[func (w *Widget) LayoutItem_SetWidth(nType xcc.Layout_Size_, nWidth int32) int {]
ff=布局项置宽度
nType=类型

[func (w *Widget) LayoutItem_SetHeight(nType xcc.Layout_Size_, nHeight int32) int {]
ff=布局项置高度
nType=类型

[func (w *Widget) LayoutItem_GetWidth(pType *xcc.Layout_Size_, pWidth *int32) int {]
ff=布局项取宽度
pType=返回类型

[func (w *Widget) LayoutItem_GetHeight(pType *xcc.Layout_Size_, pHeight *int32) int {]
ff=布局项取高度
pType=返回类型

[func (w *Widget) LayoutItem_SetAlign(nAlign xcc.Layout_Align_Axis_) int {]
ff=布局项置对齐
nAlign=对齐方式

[func (w *Widget) LayoutItem_SetMargin(left, top, right, bottom int) int {]
ff=布局项置外间距

[func (w *Widget) LayoutItem_GetMargin(pMargin *xc.RECT) int {]
ff=布局项取外间距
pMargin=接收返回

[func (w *Widget) LayoutItem_SetMinSize(width, height int) int {]
ff=布局项置最小大小
height=最小高度
width=最小宽度

[func (w *Widget) LayoutItem_SetPosition(left, top, right, bottom int) int {]
ff=布局项置位置
bottom=下边距离
right=右边距离
top=上边距离
left=左边距离

[func (w *Widget) SetID(nID int32) int {]
ff=置ID
nID=ID值

[func (w *Widget) SetUID(nUID int32) int {]
ff=置UID
nUID=UID值

[func (w *Widget) SetName(pName string) int {]
ff=置名称
pName=name值

[func (w *Widget) GetID() int32 {]
ff=取ID

[func (w *Widget) GetUID() int32 {]
ff=取UID

[func (w *Widget) GetName() string {]
ff=取名称
