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

[func XBtn_Create(x, y, cx, cy int32, pName string, hParent int) int {]
ff=按钮_创建
hParent=父窗口句柄或元素句柄
pName=标题
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func XBtn_IsCheck(hEle int) bool {]
ff=按钮_判断选中
hEle=元素句柄

[func XBtn_SetCheck(hEle int, bCheck bool) bool {]
ff=按钮_置选中
bCheck=是否选中
hEle=元素句柄

[func XBtn_SetState(hEle int, nState xcc.Common_State3_) int {]
ff=按钮_置状态
nState=按钮状态
hEle=元素句柄

[func XBtn_GetState(hEle int) xcc.Common_State3_ {]
ff=按钮_取状态
hEle=元素句柄

[func XBtn_GetStateEx(hEle int) xcc.Button_State_ {]
ff=按钮_取状态EX
hEle=元素句柄

[func XBtn_SetTypeEx(hEle int, nType xcc.XC_OBJECT_TYPE_EX) int {]
ff=按钮_置类型EX
nType=按钮类型
hEle=元素句柄

[func XBtn_SetGroupID(hEle int, nID int32) int {]
ff=按钮_置组ID
nID=组ID
hEle=元素句柄

[func XBtn_GetGroupID(hEle int) int32 {]
ff=按钮_取组ID
hEle=元素句柄

[func XBtn_SetBindEle(hEle int, hBindEle int) int {]
ff=按钮_置绑定元素
hBindEle=将要绑定的元素
hEle=元素句柄

[func XBtn_GetBindEle(hEle int) int {]
ff=按钮_取绑定元素
hEle=元素句柄

[func XBtn_SetTextAlign(hEle int, nFlags xcc.TextFormatFlag_) int {]
ff=按钮_置文本对齐
nFlags=对齐方式
hEle=元素句柄

[func XBtn_GetTextAlign(hEle int) xcc.TextFormatFlag_ {]
ff=按钮_取文本对齐方式
hEle=元素句柄

[func XBtn_SetIconAlign(hEle int, align xcc.Button_Icon_Align_) int {]
ff=按钮_置图标对齐
align=对齐方式
hEle=元素句柄

[func XBtn_SetOffset(hEle int, x int, y int) int {]
ff=按钮_置偏移
y=偏移量y
x=偏移量x
hEle=元素句柄

[func XBtn_SetOffsetIcon(hEle int, x int, y int) int {]
ff=按钮_置图标偏移
y=偏移量y
x=偏移量x
hEle=元素句柄

[func XBtn_SetIconSpace(hEle int, size int) int {]
ff=按钮_置图标间隔
size=间隔大小
hEle=元素句柄

[func XBtn_SetText(hEle int, pName string) int {]
ff=按钮_置文本
pName=文本内容
hEle=元素句柄

[func XBtn_GetText(hEle int) string {]
ff=按钮_取文本
hEle=元素句柄

[func XBtn_SetIcon(hEle int, hImage int) int {]
ff=按钮_置图标
hImage=图片句柄
hEle=元素句柄

[func XBtn_SetIconDisable(hEle int, hImage int) int {]
ff=按钮_置禁用图标
hImage=图片句柄
hEle=元素句柄

[func XBtn_GetIcon(hEle int, nType int) int {]
ff=按钮_取图标
nType=图标类型
hEle=元素句柄

[func XBtn_AddAnimationFrame(hEle int, hImage int, uElapse int) int {]
ff=按钮_添加动画帧
uElapse=图片帧延时
hImage=图片句柄
hEle=元素句柄

[func XBtn_EnableAnimation(hEle int, bEnable bool, bLoopPlay bool) int {]
ff=按钮_启用动画
bLoopPlay=是否循环播放
bEnable=开始播放动画TRUE
hEle=元素句柄
