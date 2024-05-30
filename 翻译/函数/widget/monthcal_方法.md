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

[func NewMonthCal(x int, y int, cx int, cy int, hParent int) *MonthCal {]
ff=创建月历
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func NewMonthCalByHandle(handle int) *MonthCal {]
ff=创建月历并按句柄

[func NewMonthCalByName(name string) *MonthCal {]
ff=创建月历并按名称

[func NewMonthCalByUID(nUID int) *MonthCal {]
ff=创建月历并按UID

[func NewMonthCalByUIDName(name string) *MonthCal {]
ff=创建月历并按UID名称

[func (m *MonthCal) GetButton(nType int) int {]
ff=取内部按钮
nType=按钮类型

[func (m *MonthCal) SetToday(nYear int32, nMonth int32, nDay int32) int {]
ff=置当前日期
nDay=日
nMonth=月
nYear=年

[func (m *MonthCal) GetToday(pnYear *int32, pnMonth *int32, pnDay *int32) int {]
ff=取当前日期
pnDay=日指针
pnMonth=月指针
pnYear=年指针

[func (m *MonthCal) GetSelDate(pnYear *int32, pnMonth *int32, pnDay *int32) int {]
ff=取选择日期
pnDay=日指针
pnMonth=月指针
pnYear=年指针

[func (m *MonthCal) SetTextColor(nFlag int32, color int) int {]
ff=置文本颜色
color=ABGR颜色值
nFlag=类型

[func (m *MonthCal) Event_MONTHCAL_CHANGE(pFun XE_MONTHCAL_CHANGE) bool {]
ff=事件_日期改变

[func (m *MonthCal) Event_MONTHCAL_CHANGE1(pFun XE_MONTHCAL_CHANGE1) bool {]
ff=事件_日期改变1
