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

[func NewByFile(pFileName string) *Svg {]
ff=创建并按文件
pFileName=文件名

[func NewByString(pString string) *Svg {]
ff=创建并按字符串
pString=字符串

[func NewByStringW(pString string) *Svg {]
ff=创建并按字符串W
pString=字符串

[func NewByStringUtf8(pString string) *Svg {]
ff=创建并按字符串UTF8
pString=字符串

[func NewByZip(pZipFileName, pFileName, pPassword string) *Svg {]
ff=创建并按ZIP
pPassword=zip密码
pFileName=svg文件名
pZipFileName=zip文件名

[func NewByZipRes(id int32, pFileName, pPassword string, hModule uintptr) *Svg {]
ff=创建并按资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=svg文件名
id=资源ID

[func NewByZipMem(data #左中括号##右中括号#byte, pFileName, pPassword string) *Svg {]
ff=创建并按内存ZIP
pPassword=zip密码
pFileName=svg文件名
data=zip数据

[func NewByRes(id int32, pType string, hModule uintptr) *Svg {]
ff=创建并按资源
hModule=从指定模块加载
pType=资源类型
id=资源ID

[func (s *Svg) SetSize(nWidth, nHeight int32) {]
ff=置大小
nHeight=高度
nWidth=宽度

[func (s *Svg) GetSize(pWidth, pHeight *int32) {]
ff=取大小
pHeight=接收返回高度
pWidth=接收返回宽度

[func (s *Svg) GetWidth() int32 {]
ff=取宽度

[func (s *Svg) GetHeight() int32 {]
ff=取高度

[func (s *Svg) SetPosition(x, y int32) {]
ff=置偏移
y=y轴偏移
x=x轴偏移

[func (s *Svg) GetPosition(pX, pY *int32) {]
ff=取偏移
pY=y轴偏移
pX=x轴偏移

[func (s *Svg) SetPositionF(x, y float32) {]
ff=置偏移F
y=y轴偏移
x=x轴偏移

[func (s *Svg) GetPositionF(pX, pY *float32) {]
ff=取偏移F
pY=y轴偏移
pX=x轴偏移

[func (s *Svg) GetViewBox(pViewBox *xc.RECT) {]
ff=取视图框
pViewBox=接收返回视图框

[func (s *Svg) EnableAutoDestroy(bEnable bool) {]
ff=启用自动销毁
bEnable=是否自动销毁

[func (s *Svg) AddRef() {]
ff=增加引用计数

[func (s *Svg) Release() {]
ff=释放引用计数

[func (s *Svg) GetRefCount() int32 {]
ff=取引用计数

[func (s *Svg) Destroy() {]
ff=销毁

[func (s *Svg) SetAlpha(alpha byte) {]
ff=置透明度
alpha=透明度

[func (s *Svg) GetAlpha() byte {]
ff=取透明度

[func (s *Svg) SetUserFillColor(color int, bEnable bool) {]
ff=置用户填充颜色
bEnable=是否有效
color=颜色

[func (s *Svg) SetUserStrokeColor(color int, strokeWidth float32, bEnable bool) {]
ff=置用户笔触颜色
bEnable=是否有效
strokeWidth=笔触宽度
color=颜色

[func (s *Svg) GetUserFillColor(pColor *int) bool {]
ff=取用户填充颜色
pColor=返回颜色值

[func (s *Svg) GetUserStrokeColor(pColor *int, pStrokeWidth *float32) bool {]
ff=取用户笔触颜色
pColor=返回颜色值

[func (s *Svg) SetRotateAngle(angle float32) {]
ff=置旋转角度
angle=转角度

[func (s *Svg) GetRotateAngle() float32 {]
ff=取旋转角度

[func (s *Svg) SetRotate(angle float32, x float32, y float32, bOffset bool) {]
ff=置旋转
bOffset=偏移方式
y=旋转中心点Y
x=旋转中心点X
angle=角度

[func (s *Svg) GetRotate(pAngle *float32, pX *float32, pY *float32, pbOffset *bool) {]
ff=取旋转
pbOffset=返回偏移方式
pY=返回y
pX=返回x
pAngle=返回角度

[func (s *Svg) Show(bShow bool) {]
ff=显示
bShow=是否显示
