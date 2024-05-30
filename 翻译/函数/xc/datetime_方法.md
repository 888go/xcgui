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

[func XDateTime_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=日期_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func XDateTime_SetStyle(hEle int, nStyle int) int {]
ff=日期_置样式
nStyle=样式
hEle=元素句柄

[func XDateTime_GetStyle(hEle int) int {]
ff=日期_取样式
hEle=元素句柄

[func XDateTime_EnableSplitSlash(hEle int, bSlash bool) int {]
ff=日期_启用分割栏为斜线
bSlash=TRUE
hEle=元素句柄

[func XDateTime_GetButton(hEle int, nType int) int {]
ff=日期_取内部按钮
nType=按钮类型
hEle=元素句柄

[func XDateTime_GetSelBkColor(hEle int) int {]
ff=日期_取选择日期背景颜色
hEle=元素句柄

[func XDateTime_SetSelBkColor(hEle int, crSelectBk int) int {]
ff=日期_置选择日期背景颜色
crSelectBk=文字被选中背景色
hEle=元素句柄

[func XDateTime_GetDate(hEle int, pnYear *int32, pnMonth *int32, pnDay *int32) int {]
ff=日期_取当前日期
pnDay=日指针
pnMonth=月指针
pnYear=年指针
hEle=元素句柄

[func XDateTime_SetDate(hEle int, nYear int32, nMonth int32, nDay int32) int {]
ff=日期_置当前日期
nDay=日
nMonth=月
nYear=年
hEle=元素句柄

[func XDateTime_GetTime(hEle int, pnHour *int32, pnMinute *int32, pnSecond *int32) int {]
ff=日期_取当前时间
pnSecond=秒指针
pnMinute=分指针
pnHour=时指针
hEle=元素句柄

[func XDateTime_SetTime(hEle int, nHour int32, nMinute int32, nSecond int32) int {]
ff=日期_社区当前时间
nSecond=秒
nMinute=分
nHour=时
hEle=元素句柄

[func XDateTime_Popup(hEle int) int {]
ff=日期_弹出
hEle=元素句柄
