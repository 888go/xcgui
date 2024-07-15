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

[func NewProgressBar(x int, y int, cx int, cy int, hParent int) *ProgressBar {]
ff=创建进度条
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewProgressBarByHandle(handle int) *ProgressBar {]
ff=创建进度条并按句柄

[func NewProgressBarByName(name string) *ProgressBar {]
ff=创建进度条并按名称

[func NewProgressBarByUID(nUID int) *ProgressBar {]
ff=创建进度条并按UID

[func NewProgressBarByUIDName(name string) *ProgressBar {]
ff=创建进度条并按UID名称

[func (p *ProgressBar) SetRange(range_ int) int {]
ff=置范围
range_=范围

[func (p *ProgressBar) GetRange() int {]
ff=取范围

[func (p *ProgressBar) SetImageLoad(hImage int) int {]
ff=置进度图片
hImage=图片句柄

[func (p *ProgressBar) SetPos(pos int) int {]
ff=置进度
pos=位置点

[func (p *ProgressBar) GetPos() int {]
ff=取进度

[func (p *ProgressBar) EnableHorizon(bHorizon bool) int {]
ff=置水平
bHorizon=水平或垂直

[func (p *ProgressBar) EnableStretch(bStretch bool) bool {]
ff=启用缩放
bStretch=缩放

[func (p *ProgressBar) EnableShowText(bShow bool) bool {]
ff=启用进度文本
bShow=是否启用

[func (p *ProgressBar) SetColorLoad(color int) bool {]
ff=置进度颜色
color=ABGR颜色

[func (p *ProgressBar) Event_PROGRESSBAR_CHANGE(pFun XE_PROGRESSBAR_CHANGE) bool {]
ff=事件_进度改变事件

[func (p *ProgressBar) Event_PROGRESSBAR_CHANGE1(pFun XE_PROGRESSBAR_CHANGE1) bool {]
ff=事件_进度改变1
