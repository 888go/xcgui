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

[func XTrayIcon_Reset() {]
ff=托盘图标_重置

[func XTrayIcon_Del() bool {]
ff=托盘图标_删除

[func XTrayIcon_Modify() bool {]
ff=托盘图标_修改

[func XTrayIcon_SetFocus() bool {]
ff=托盘图标_置焦点

[func XTrayIcon_Add(hWindow int, id int32) bool {]
ff=托盘图标_添加
id=托盘图标唯一标识符
hWindow=关联窗口句柄

[func XTrayIcon_SetIcon(hIcon uintptr) {]
ff=托盘图标_置图标
hIcon=图标句柄

[func XTrayIcon_SetTips(pTips string) {]
ff=托盘图标_置提示文本
pTips=提示文本内容

[func XTrayIcon_SetCallbackMessage(user_message uint32) {]
ff=托盘图标_置回调消息
user_message=用户自定义消息

[func XTrayIcon_SetPopupBalloon(pTitle, pText string, hBalloonIcon uintptr, flags xcc.TrayIcon_Flag_) {]
ff=托盘图标_置弹出气泡
flags=标识
hBalloonIcon=气泡图标
pText=弹出气泡内容
pTitle=弹出气泡标题
