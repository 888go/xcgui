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

[func XPane_Create(pName string, nWidth int, nHeight int, hFrameWnd int) int {]
ff=窗格_创建
hFrameWnd=框架窗口
nHeight=高度
nWidth=宽度
pName=窗格标题

[func XPane_SetView(hEle int, hView int) int {]
ff=窗格_置视图
hView=绑定视图元素
hEle=元素句柄

[func XPane_SetTitle(hEle int, pTitle string) int {]
ff=窗格_置标题
pTitle=文本内容
hEle=元素句柄

[func XPane_GetTitle(hEle int) string {]
ff=窗格_取标题
hEle=元素句柄

[func XPane_SetCaptionHeight(hEle int, nHeight int) int {]
ff=窗格_置标题栏高度
nHeight=高度
hEle=元素句柄

[func XPane_GetCaptionHeight(hEle int) int {]
ff=窗格_取标题栏高度
hEle=元素句柄

[func XPane_IsShowPane(hEle int) bool {]
ff=窗格_判断显示
hEle=元素句柄

[func XPane_SetSize(hEle int, nWidth int, nHeight int) int {]
ff=窗格_置大小
nHeight=高度
nWidth=宽度
hEle=元素句柄

[func XPane_GetState(hEle int) xcc.Pane_State_ {]
ff=窗格_取状态
hEle=元素句柄

[func XPane_GetViewRect(hEle int, pRect *RECT) int {]
ff=窗格_取视图坐标
pRect=接收返回坐标
hEle=元素句柄

[func XPane_HidePane(hEle int, bGroupActivate bool) int {]
ff=窗格_隐藏
bGroupActivate=延迟组成员激活
hEle=元素句柄

[func XPane_ShowPane(hEle int, bGroupActivate bool) int {]
ff=窗格_显示
bGroupActivate=延迟组成员激活
hEle=元素句柄

[func XPane_DockPane(hEle int) int {]
ff=窗格_停靠
hEle=元素句柄

[func XPane_LockPane(hEle int) int {]
ff=窗格_锁定
hEle=元素句柄

[func XPane_FloatPane(hEle int) int {]
ff=窗格_浮动
hEle=元素句柄

[func XPane_DrawPane(hEle int, hDraw int) int {]
ff=窗格_绘制
hDraw=图形绘制句柄
hEle=元素句柄

[func XPane_SetSelect(hEle int) bool {]
ff=窗口_置选中
hEle=元素句柄

[func XPane_IsGroupActivate(hEle int) bool {]
ff=窗格_是否激活
hEle=元素句柄
