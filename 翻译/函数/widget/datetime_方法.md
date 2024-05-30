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

[func NewDateTime(x int, y int, cx int, cy int, hParent int) *DateTime {]
ff=创建日期
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func NewDateTimeByHandle(handle int) *DateTime {]
ff=创建日期并按句柄
handle=句柄

[func NewDateTimeByName(name string) *DateTime {]
ff=创建日期并按名称
name=名称

[func NewDateTimeByUID(nUID int) *DateTime {]
ff=创建日期并按UID

[func NewDateTimeByUIDName(name string) *DateTime {]
ff=创建日期并按UID名称
name=UID名称

[func (d *DateTime) SetStyle(nStyle int) int {]
ff=置样式
nStyle=样式

[func (d *DateTime) GetStyle() int {]
ff=取样式

[func (d *DateTime) EnableSplitSlash(bSlash bool) int {]
ff=启用分割栏为斜线
bSlash=斜线

[func (d *DateTime) GetButton(nType int) int {]
ff=取内部按钮
nType=按钮类型

[func (d *DateTime) GetSelBkColor() int {]
ff=取选择日期背景颜色

[func (d *DateTime) SetSelBkColor(crSelectBk int) int {]
ff=置选择日期背景颜色
crSelectBk=背景色

[func (d *DateTime) GetDate(pnYear *int32, pnMonth *int32, pnDay *int32) int {]
ff=取当前日期
pnDay=日指针
pnMonth=月指针
pnYear=年指针

[func (d *DateTime) SetDate(nYear int32, nMonth int32, nDay int32) int {]
ff=置当前日期
nDay=日
nMonth=月
nYear=年

[func (d *DateTime) GetTime(pnHour *int32, pnMinute *int32, pnSecond *int32) int {]
ff=取当前时间
pnSecond=秒指针
pnMinute=分指针
pnHour=时指针

[func (d *DateTime) SetTime(nHour int32, nMinute int32, nSecond int32) int {]
ff=设置当前时间
nSecond=秒
nMinute=分
nHour=时

[func (d *DateTime) Popup() int {]
ff=弹出月历卡片

[func (d *DateTime) Event_DATETIME_CHANGE(pFun XE_DATETIME_CHANGE) bool {]
ff=事件_内容改变

[func (d *DateTime) Event_DATETIME_CHANGE1(pFun XE_DATETIME_CHANGE1) bool {]
ff=事件_内容改变1

[func (d *DateTime) Event_DATETIME_POPUP_MONTHCAL(pFun XE_DATETIME_POPUP_MONTHCAL) bool {]
ff=事件_弹出月历卡片

[func (d *DateTime) Event_DATETIME_POPUP_MONTHCAL1(pFun XE_DATETIME_POPUP_MONTHCAL1) bool {]
ff=事件_弹出月历卡片1

[func (d *DateTime) Event_DATETIME_EXIT_MONTHCAL(pFun XE_DATETIME_EXIT_MONTHCAL) bool {]
ff=事件_月历卡片退出

[func (d *DateTime) Event_DATETIME_EXIT_MONTHCAL1(pFun XE_DATETIME_EXIT_MONTHCAL1) bool {]
ff=事件_月历卡片退出1
