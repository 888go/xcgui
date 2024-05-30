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

[func NewButton(x, y, cx, cy int32, pName string, hParent int) *Button {]
ff=创建按钮
hParent=父窗口句柄
pName=标题
cy=高度
cx=宽度
y=按钮y坐标
x=按钮x坐标

[func NewButtonByHandle(handle int) *Button {]
ff=创建按钮并按句柄
handle=句柄

[func NewButtonByName(name string) *Button {]
ff=创建按钮并按名称
name=名称

[func NewButtonByUID(nUID int) *Button {]
ff=创建按钮并按UID

[func NewButtonByUIDName(name string) *Button {]
ff=创建按钮并按UID名称
name=UID名称

[func (b *Button) IsCheck() bool {]
ff=判断选中

[func (b *Button) SetCheck(bCheck bool) bool {]
ff=置选中
bCheck=是否选中

[func (b *Button) SetState(nState xcc.Common_State3_) int {]
ff=置状态
nState=按钮状态

[func (b *Button) GetState() xcc.Common_State3_ {]
ff=取状态

[func (b *Button) GetStateEx() xcc.Button_State_ {]
ff=取状态EX

[func (b *Button) SetTypeEx(nType xcc.XC_OBJECT_TYPE_EX) int {]
ff=置类型EX
nType=按钮类型

[func (b *Button) SetGroupID(nID int32) int {]
ff=置组ID
nID=组ID

[func (b *Button) GetGroupID() int32 {]
ff=取组ID

[func (b *Button) SetBindEle(hBindEle int) int {]
ff=置绑定元素
hBindEle=将要绑定的元素

[func (b *Button) GetBindEle() int {]
ff=取绑定元素

[func (b *Button) SetTextAlign(nFlags xcc.TextFormatFlag_) int {]
ff=置文本对齐
nFlags=对齐方式

[func (b *Button) GetTextAlign() xcc.TextFormatFlag_ {]
ff=取文本对齐方式

[func (b *Button) SetIconAlign(align xcc.Button_Icon_Align_) int {]
ff=置图标对齐
align=对齐方式

[func (b *Button) SetOffset(x int, y int) int {]
ff=置偏移
y=偏移量y
x=偏移量x

[func (b *Button) SetOffsetIcon(x int, y int) int {]
ff=置图标偏移
y=偏移量y
x=偏移量x

[func (b *Button) SetIconSpace(size int) int {]
ff=置图标间隔
size=间隔大小

[func (b *Button) SetText(pName string) int {]
ff=置文本
pName=文本内容

[func (b *Button) GetText() string {]
ff=取文本

[func (b *Button) SetIcon(hImage int) int {]
ff=置图标
hImage=图片句柄

[func (b *Button) SetIconDisable(hImage int) int {]
ff=置禁用图标
hImage=图片句柄

[func (b *Button) GetIcon(nType int) int {]
ff=取图标
nType=图标类型

[func (b *Button) AddAnimationFrame(hImage int, uElapse int) int {]
ff=添加动画帧
uElapse=图片帧延时
hImage=图片句柄

[func (b *Button) EnableAnimation(bEnable bool, bLoopPlay bool) int {]
ff=启用动画
bLoopPlay=是否循环播放
bEnable=启用

[func (b *Button) Event_BnClick(pFun XE_BNCLICK) bool {]
ff=事件_被单击

[func (b *Button) Event_BnClick1(pFun XE_BNCLICK1) bool {]
ff=事件_被单击1

[func (b *Button) Event_BUTTON_CHECK(pFun XE_BUTTON_CHECK) bool {]
ff=事件_选中

[func (b *Button) Event_BUTTON_CHECK1(pFun XE_BUTTON_CHECK1) bool {]
ff=事件_选中1
