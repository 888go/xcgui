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

[func XSvg_LoadFile(pFileName string) int {]
ff=SVG_加载从文件
pFileName=文件名

[func XSvg_LoadString(pString string) int {]
ff=SVG_加载从字符串
pString=字符串

[func XSvg_LoadZip(pZipFileName, pFileName, pPassword string) int {]
ff=SVG_加载从ZIP
pPassword=zip密码
pFileName=svg文件名
pZipFileName=zip文件名

[func XSvg_LoadZipRes(id int32, pFileName, pPassword string, hModule uintptr) int {]
ff=SVG_加载从资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=svg文件名
id=资源ID

[func XSvg_LoadRes(id int32, pType string, hModule uintptr) int {]
ff=SVG_加载从资源
hModule=从指定模块加载
pType=资源类型
id=资源ID

[func XSvg_SetSize(hSvg int, nWidth, nHeight int32) {]
ff=SVG_置大小
nHeight=高度
nWidth=宽度
hSvg=SVG句柄

[func XSvg_GetSize(hSvg int, pWidth, pHeight *int32) {]
ff=SVG_取大小
pHeight=接收返回高度
pWidth=接收返回宽度
hSvg=SVG句柄

[func XSvg_GetWidth(hSvg int) int32 {]
ff=SVG_取宽度
hSvg=SVG句柄

[func XSvg_GetHeight(hSvg int) int32 {]
ff=SVG_取高度
hSvg=SVG句柄

[func XSvg_SetPosition(hSvg int, x, y int32) {]
ff=SVG_置偏移
y=y轴偏移
x=x轴偏移
hSvg=SVG句柄

[func XSvg_GetPosition(hSvg int, pX, pY *int32) {]
ff=SVG_取偏移
pY=y轴偏移
pX=x轴偏移
hSvg=SVG句柄

[func XSvg_SetPositionF(hSvg int, x, y float32) {]
ff=SVG_置偏移F
y=y轴偏移
x=x轴偏移
hSvg=SVG句柄

[func XSvg_GetPositionF(hSvg int, pX, pY *float32) {]
ff=SVG_取偏移F
pY=y轴偏移
pX=x轴偏移
hSvg=SVG句柄

[func XSvg_GetViewBox(hSvg int, pViewBox *RECT) {]
ff=SVG_取视图框
pViewBox=接收返回视图框
hSvg=SVG句柄

[func XSvg_EnableAutoDestroy(hSvg int, bEnable bool) {]
ff=SVG_启用自动销毁
bEnable=是否自动销毁
hSvg=SVG句柄

[func XSvg_AddRef(hSvg int) {]
ff=SVG_增加引用计数
hSvg=SVG句柄

[func XSvg_Release(hSvg int) {]
ff=SVG_释放引用计数
hSvg=SVG句柄

[func XSvg_GetRefCount(hSvg int) int32 {]
ff=SVG_取引用计数
hSvg=SVG句柄

[func XSvg_Destroy(hSvg int) {]
ff=SVG_销毁
hSvg=SVG句柄

[func XSvg_SetAlpha(hSvg int, alpha byte) {]
ff=SVG_置透明度
alpha=透明度
hSvg=SVG句柄

[func XSvg_GetAlpha(hSvg int) byte {]
ff=SVG_取透明度
hSvg=SVG句柄

[func XSvg_SetUserFillColor(hSvg int, color int, bEnable bool) {]
ff=SVG_置用户填充颜色
bEnable=是否有效
color=AGBR颜色
hSvg=SVG句柄

[func XSvg_SetUserStrokeColor(hSvg int, color int, strokeWidth float32, bEnable bool) {]
ff=SVG_置用户笔触颜色
bEnable=是否有效
strokeWidth=笔触宽度
color=AGBR颜色
hSvg=SVG句柄

[func XSvg_GetUserFillColor(hSvg int, pColor *int) bool {]
ff=SVG_取用户填充颜色
pColor=返回颜色值
hSvg=SVG句柄

[func XSvg_GetUserStrokeColor(hSvg int, pColor *int, pStrokeWidth *float32) bool {]
ff=SVG_取用户笔触颜色
pColor=返回颜色值
hSvg=SVG句柄

[func XSvg_SetRotateAngle(hSvg int, angle float32) {]
ff=SVG_置旋转角度
angle=转角度
hSvg=SVG句柄

[func XSvg_GetRotateAngle(hSvg int) float32 {]
ff=SVG_取旋转角度
hSvg=SVG句柄

[func XSvg_SetRotate(hSvg int, angle float32, x float32, y float32, bOffset bool) {]
ff=SVG_置旋转
bOffset=TRUE
y=旋转中心点Y
x=旋转中心点X
angle=角度
hSvg=SVG句柄

[func XSvg_GetRotate(hSvg int, pAngle *float32, pX *float32, pY *float32, pbOffset *bool) {]
ff=SVG_取旋转
pbOffset=返回TRUE
pY=返回旋转中心点Y
pX=返回旋转中心点X
pAngle=返回角度
hSvg=SVG句柄

[func XSvg_Show(hSvg int, bShow bool) {]
ff=SVG_显示
bShow=是否显示
hSvg=SVG句柄

[func XSvg_LoadStringW(pString string) int {]
ff=SVG_加载从字符串W
pString=字符串

[func XSvg_LoadStringUtf8(pString string) int {]
ff=SVG_加载从字符串UTF8
pString=字符串

[func XSvg_LoadZipMem(data #左中括号##右中括号#byte, pFileName, pPassword string) int {]
ff=SVG_加载从内存ZIP
pPassword=zip密码
pFileName=svg文件名
data=zip数据
