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
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
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
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func XTable_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=表格_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=按钮y坐标
x=按钮x坐标

[func XTable_Reset(hShape int, nRow int, nCol int) int {]
ff=表格_重置
nCol=列数量
nRow=行数量
hShape=形状对象句柄

[func XTable_ComboRow(hShape int, iRow int, iCol int, count int) int {]
ff=表格_组合行
count=数量
iCol=列索引
iRow=行索引
hShape=形状对象句柄

[func XTable_ComboCol(hShape int, iRow int, iCol int, count int) int {]
ff=表格_组合列
count=数量
iCol=列索引
iRow=行索引
hShape=形状对象句柄

[func XTable_SetColWidth(hShape int, iCol int, width int) int {]
ff=表格_置列宽
width=宽度
iCol=列索引
hShape=形状对象句柄

[func XTable_SetRowHeight(hShape int, iRow int, height int) int {]
ff=表格_置行高
height=高度
iRow=行索引
hShape=形状对象句柄

[func XTable_SetBorderColor(hShape int, color int) int {]
ff=表格_置边颜色
color=颜色
hShape=形状对象句柄

[func XTable_SetTextColor(hShape int, color int) int {]
ff=表格_置文本颜色
color=颜色
hShape=形状对象句柄

[func XTable_SetFont(hShape int, hFont int) int {]
ff=表格_置字体
hFont=字体句柄
hShape=形状对象句柄

[func XTable_SetItemPadding(hShape int, leftSize int, topSize int, rightSize int, bottomSize int) int {]
ff=表格_置项内填充
bottomSize=下
rightSize=右
topSize=上
leftSize=左
hShape=形状对象句柄

[func XTable_SetItemText(hShape int, iRow int, iCol int, pText string) int {]
ff=表格_置项文本
pText=文本
iCol=列索引
iRow=行索引
hShape=形状对象句柄

[func XTable_SetItemFont(hShape int, iRow int, iCol int, hFont int) int {]
ff=表格_置项字体
hFont=字体句柄
iCol=列索引
iRow=行索引
hShape=形状对象句柄

[func XTable_SetItemTextAlign(hShape int, iRow int, iCol int, nAlign xcc.TextFormatFlag_) int {]
ff=表格_置项文本对齐
nAlign=对齐方式
iCol=列索引
iRow=行索引
hShape=形状对象句柄

[func XTable_SetItemTextColor(hShape int, iRow int, iCol int, color int, bColor bool) int {]
ff=表格_置项文本色
bColor=是否使用
color=颜色
iCol=列索引
iRow=行索引
hShape=形状对象句柄

[func XTable_SetItemBkColor(hShape int, iRow int, iCol int, color int, bColor bool) int {]
ff=表格_置项背景色
bColor=是否使用
color=颜色
iCol=列索引
iRow=行索引
hShape=形状对象句柄

[func XTable_SetItemLine(hShape int, iRow1 int, iCol1 int, iRow2 int, iCol2 int, nFlag int, color int) int {]
ff=表格_置项线
color=颜色
nFlag=标识
iCol2=列索引2
iRow2=行索引2
iCol1=列索引1
iRow1=行索引1
hShape=形状对象句柄

[func XTable_SetItemFlag(hShape int, iRow int, iCol int, flag xcc.Table_Flag_) int {]
ff=表格_置项标识
flag=标识
iCol=列索引
iRow=行索引
hShape=形状对象句柄

[func XTable_GetItemRect(hShape int, iRow int, iCol int, pRect *RECT) bool {]
ff=表格_取项坐标
pRect=接收返回坐标
iCol=列索引
iRow=行索引
hShape=形状对象句柄
