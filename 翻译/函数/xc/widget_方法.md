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

[func XWidget_IsShow(hXCGUI int) bool {]
ff=窗口组件_判断显示
hXCGUI=对象句柄

[func XWidget_Show(hXCGUI int, bShow bool) int {]
ff=窗口组件_显示
bShow=是否显示
hXCGUI=对象句柄

[func XWidget_EnableLayoutControl(hXCGUI int, bEnable bool) int {]
ff=窗口组件_启用布局控制
bEnable=是否启用
hXCGUI=对象句柄

[func XWidget_IsLayoutControl(hXCGUI int) bool {]
ff=窗口组件_是否布局控制
hXCGUI=对象句柄

[func XWidget_GetParentEle(hXCGUI int) int {]
ff=窗口组件_取父元素
hXCGUI=对象句柄

[func XWidget_GetParent(hXCGUI int) int {]
ff=窗口组件_取父对象
hXCGUI=对象句柄

[func XWidget_GetHWND(hXCGUI int) uintptr {]
ff=窗口组件_取HWND
hXCGUI=对象句柄

[func XWidget_GetHWINDOW(hXCGUI int) int {]
ff=窗口组件_取HWINDOW
hXCGUI=对象句柄

[func XWidget_LayoutItem_EnableWrap(hXCGUI int, bWrap bool) int {]
ff=窗口组件_布局项_启用换行
bWrap=是否换行
hXCGUI=对象句柄

[func XWidget_LayoutItem_EnableSwap(hXCGUI int, bEnable bool) int {]
ff=窗口组件_布局项_启用交换
bEnable=是否启用
hXCGUI=对象句柄

[func XWidget_LayoutItem_EnableFloat(hXCGUI int, bFloat bool) int {]
ff=窗口组件_布局项_启用浮动
bFloat=是否浮动
hXCGUI=对象句柄

[func XWidget_LayoutItem_SetWidth(hXCGUI int, nType xcc.Layout_Size_, nWidth int32) int {]
ff=窗口组件_布局项_置宽度
nType=类型
hXCGUI=对象句柄

[func XWidget_LayoutItem_SetHeight(hXCGUI int, nType xcc.Layout_Size_, nHeight int32) int {]
ff=窗口组件_布局项_置高度
nType=类型
hXCGUI=对象句柄

[func XWidget_LayoutItem_GetWidth(hXCGUI int, pType *xcc.Layout_Size_, pWidth *int32) int {]
ff=窗口组件_布局项_取宽度
pType=返回类型
hXCGUI=对象句柄

[func XWidget_LayoutItem_GetHeight(hXCGUI int, pType *xcc.Layout_Size_, pHeight *int32) int {]
ff=窗口组件_布局项_取高度
pType=返回类型
hXCGUI=对象句柄

[func XWidget_LayoutItem_SetAlign(hXCGUI int, nAlign xcc.Layout_Align_Axis_) int {]
ff=窗口组件_布局项_置对齐
nAlign=对齐方式
hXCGUI=对象句柄

[func XWidget_LayoutItem_SetMargin(hXCGUI, left, top, right, bottom int) int {]
ff=窗口组件_布局项_置外间距
hXCGUI=对象句柄

[func XWidget_LayoutItem_GetMargin(hXCGUI int, pMargin *RECT) int {]
ff=窗口组件_布局项_取外间距
pMargin=接收返回
hXCGUI=对象句柄

[func XWidget_LayoutItem_SetMinSize(hXCGUI, width, height int) int {]
ff=窗口组件_布局项_置最小大小
height=最小高度
width=最小宽度
hXCGUI=对象句柄

[func XWidget_LayoutItem_SetPosition(hXCGUI, left, top, right, bottom int) int {]
ff=窗口组件_布局项_置位置
bottom=下边距离
right=右边距离
top=上边距离
left=左边距离
hXCGUI=对象句柄

[func XWidget_SetID(hXCGUI int, nID int32) int {]
ff=窗口组件_置ID
nID=ID值
hXCGUI=对象句柄

[func XWidget_SetUID(hXCGUI int, nUID int32) int {]
ff=窗口组件_置UID
nUID=UID值
hXCGUI=对象句柄

[func XWidget_SetName(hXCGUI int, pName string) int {]
ff=窗口组件_置名称
pName=名称
hXCGUI=对象句柄

[func XWidget_GetID(hXCGUI int) int32 {]
ff=窗口组件_取ID
hXCGUI=对象句柄

[func XWidget_GetUID(hXCGUI int) int32 {]
ff=窗口组件_取UID
hXCGUI=对象句柄

[func XWidget_GetName(hXCGUI int) string {]
ff=窗口组件_取名称
hXCGUI=对象句柄
