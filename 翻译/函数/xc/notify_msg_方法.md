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

[func XNotifyMsg_Popup(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {]
ff=通知消息_弹出
position=位置

[func XNotifyMsg_PopupEx(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {]
ff=通知消息_弹出EX
position=位置

[func XNotifyMsg_WindowPopup(hWindow int, position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {]
ff=通知消息_窗口中弹出
position=位置
hWindow=窗口句柄

[func XNotifyMsg_WindowPopupEx(hWindow int, position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {]
ff=通知消息_窗口中弹出EX
position=位置
hWindow=窗口句柄

[func XNotifyMsg_SetDuration(hWindow, duration int) int {]
ff=通知消息_置持续时间
duration=持续时间
hWindow=窗口句柄

[func XNotifyMsg_SetParentMargin(hWindow, left, top, right, bottom int) int {]
ff=通知消息_置父边距
bottom=底部间隔
right=右侧间隔
top=顶部间隔
left=左侧间隔
hWindow=窗口句柄

[func XNotifyMsg_SetCaptionHeight(hWindow, nHeight int) int {]
ff=通知消息_置标题高度
nHeight=高度
hWindow=窗口句柄

[func XNotifyMsg_SetWidth(hWindow, nWidth int) int {]
ff=通知消息_置宽度
nWidth=宽度
hWindow=窗口句柄

[func XNotifyMsg_SetSpace(hWindow, nSpace int) int {]
ff=通知消息_置间隔
nSpace=间隔大小
hWindow=窗口句柄

[func XNotifyMsg_SetBorderSize(hWindow, left, top, right, bottom int) int {]
ff=通知消息_置边大小
bottom=底边
right=右边
top=顶边
left=左边
hWindow=窗口句柄
