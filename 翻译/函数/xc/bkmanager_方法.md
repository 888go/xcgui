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

[func XBkM_Create() int {]
ff=背景_创建

[func XBkM_Destroy(hBkInfoM int) int {]
ff=背景_销毁
hBkInfoM=背景管理器句柄

[func XBkM_SetBkInfo(hBkInfoM int, pText string) int {]
ff=废弃_XBkM_SetBkInfo
pText=背景内容字符串
hBkInfoM=背景管理器句柄

[func XBkM_AddInfo(hBkInfoM int, pText string) int {]
ff=背景_添加内容
pText=背景内容字符串
hBkInfoM=背景管理器句柄

[func XBkM_AddBorder(hBkInfoM int, nState xcc.CombinedState, color, width, id int) int {]
ff=背景_添加边框
nState=组合状态
hBkInfoM=背景管理器句柄

[func XBkM_AddFill(hBkInfoM int, nState xcc.CombinedState, color, id int) int {]
ff=背景_添加填充
nState=组合状态
hBkInfoM=背景管理器句柄

[func XBkM_AddImage(hBkInfoM int, nState xcc.CombinedState, hImage, id int) int {]
ff=背景_添加图片
nState=组合状态
hBkInfoM=背景管理器句柄

[func XBkM_GetCount(hBkInfoM int) int {]
ff=背景_取数量
hBkInfoM=背景管理器句柄

[func XBkM_Clear(hBkInfoM int) int {]
ff=背景_清空
hBkInfoM=背景管理器句柄

[func XBkM_Draw(hBkInfoM int, nState xcc.CombinedState, hDraw int, pRect *RECT) bool {]
ff=背景_绘制
nState=组合状态
hBkInfoM=背景管理器句柄

[func XBkM_DrawEx(hBkInfoM int, nState xcc.CombinedState, hDraw int, pRect *RECT, nStateEx xcc.CombinedState) bool {]
ff=背景_绘制EX
nState=组合状态
hBkInfoM=背景管理器句柄

[func XBkM_EnableAutoDestroy(hBkInfoM int, bEnable bool) int {]
ff=背景_启用自动销毁
bEnable=是否启用
hBkInfoM=背景管理器句柄

[func XBkM_AddRef(hBkInfoM int) int {]
ff=背景_增加引用计数
hBkInfoM=背景管理器句柄

[func XBkM_Release(hBkInfoM int) int {]
ff=背景_释放引用计数
hBkInfoM=背景管理器句柄

[func XBkM_GetRefCount(hBkInfoM int) int {]
ff=背景_取引用计数
hBkInfoM=背景管理器句柄

[func XBkM_SetInfo(hBkInfoM int, pText string) int {]
ff=背景_设置内容
pText=背景内容字符串
hBkInfoM=背景管理器句柄

[func XBkM_GetStateTextColor(hBkInfoM int, nState xcc.CombinedState, color *int) bool {]
ff=背景_取指定状态文本颜色
nState=组合状态
hBkInfoM=背景管理器句柄

[func XBkM_GetObject(hBkInfoM int, id int) int {]
ff=背景_取背景对象
id=背景对象ID
hBkInfoM=背景管理器句柄
