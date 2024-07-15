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

[func New() *BkManager {]
ff=创建

[func NewByHandle(handle int) *BkManager {]
ff=创建并按句柄

[func NewByName(name string) *BkManager {]
ff=创建并按名称

[func (b *BkManager) SetBkInfo(pText string) int {]
ff=作废SetBkInfo
pText=背景内容

[func (b *BkManager) Destroy() int {]
ff=销毁

[func (b *BkManager) AddInfo(pText string) int {]
ff=添加内容
pText=背景内容

[func (b *BkManager) AddBorder(nState xcc.CombinedState, color, width, id int) int {]
ff=添加边框
nState=组合状态

[func (b *BkManager) AddFill(nState xcc.CombinedState, color, id int) int {]
ff=添加填充
nState=组合状态

[func (b *BkManager) AddImage(nState xcc.CombinedState, hImage, id int) int {]
ff=添加图片
nState=组合状态

[func (b *BkManager) GetCount() int {]
ff=取数量

[func (b *BkManager) Clear() int {]
ff=清空

[func (b *BkManager) Draw(nState xcc.CombinedState, hDraw int, pRect *xc.RECT) bool {]
ff=绘制
nState=组合状态

[func (b *BkManager) DrawEx(nState xcc.CombinedState, hDraw int, pRect *xc.RECT, nStateEx xcc.CombinedState) bool {]
ff=绘制EX
nState=组合状态

[func (b *BkManager) EnableAutoDestroy(bEnable bool) int {]
ff=启用自动销毁
bEnable=是否启用

[func (b *BkManager) AddRef() int {]
ff=增加引用计数

[func (b *BkManager) Release() int {]
ff=释放引用计数

[func (b *BkManager) GetRefCount() int {]
ff=取引用计数

[func (b *BkManager) SetInfo(pText string) int {]
ff=设置背景内容
pText=背景内容

[func (b *BkManager) GetStateTextColor(nState xcc.CombinedState, color *int) bool {]
ff=取指定状态文本颜色
nState=组合状态

[func (b *BkManager) GetObject(id int) int {]
ff=取背景对象
id=背景对象ID
