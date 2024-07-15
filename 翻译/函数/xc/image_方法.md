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

[func XImage_LoadSrc(hImageSrc int) int {]
ff=图片_加载从图片源
hImageSrc=图片源句柄

[func XImage_LoadFile(pFileName string) int {]
ff=图片_加载从文件
pFileName=图片文件

[func XImage_LoadFileAdaptive(pFileName string, leftSize, topSize, rightSize, bottomSize int32) int {]
ff=图片_加载从文件自适应
bottomSize=下
rightSize=右
topSize=上
leftSize=左
pFileName=图片文件

[func XImage_LoadFileRect(pFileName string, x, y, cx, cy int32) int {]
ff=图片_加载从文件指定区域
cy=高度
cx=宽度
y=右
x=左
pFileName=图片文件

[func XImage_LoadResAdaptive(id int32, pType string, leftSize, topSize, rightSize, bottomSize int32, hModule uintptr) int {]
ff=图片_加载从资源自适应
hModule=从指定模块加载
bottomSize=下
rightSize=右
topSize=上
leftSize=左
pType=资源类型
id=资源ID

[func XImage_LoadRes(id int32, pType string, bStretch bool, hModule uintptr) int {]
ff=图片_加载从资源
hModule=从指定模块加载
bStretch=是否拉伸图片
pType=资源类型
id=资源ID

[func XImage_LoadZip(pZipFileName string, pFileName string, pPassword string) int {]
ff=图片_加载从ZIP
pPassword=ZIP压缩包密码
pFileName=图片文件名
pZipFileName=ZIP压缩包文件名

[func XImage_LoadZipRes(id int32, pFileName string, pPassword string, hModule uintptr) int {]
ff=图片_加载从资源ZIP
hModule=模块句柄
pPassword=ZIP压缩包密码
pFileName=图片文件名
id=RC资源ID

[func XImage_LoadZipAdaptive(pZipFileName string, pFileName string, pPassword string, x1, x2, y1, y2 int32) int {]
ff=图片_加载从ZIP自适应
pPassword=ZIP压缩包密码
pFileName=图片文件名
pZipFileName=ZIP压缩包文件名

[func XImage_LoadZipRect(pZipFileName string, pFileName string, pPassword string, x, y, cx, cy int32) int {]
ff=图片_加载从ZIP指定区域
cy=高度
cx=宽度
y=右
x=左
pPassword=密码
pFileName=图片名称
pZipFileName=ZIP文件

[func XImage_LoadZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string) int {]
ff=图片_加载从内存ZIP
pPassword=zip压缩包密码
pFileName=图片名称
data=图片数据

[func XImage_LoadMemory(pBuffer #左中括号##右中括号#byte) int {]
ff=图片_加载从内存
pBuffer=图片数据

[func XImage_LoadMemoryRect(pBuffer #左中括号##右中括号#byte, x, y, cx, cy int32) int {]
ff=图片_加载从内存指定区域
cy=高度
cx=宽度
y=右
x=左
pBuffer=图片数据

[func XImage_LoadMemoryAdaptive(pBuffer #左中括号##右中括号#byte, leftSize, topSize, rightSize, bottomSize int32) int {]
ff=图片_加载从内存自适应
bottomSize=下
rightSize=右
topSize=上
leftSize=左
pBuffer=图片数据

[func XImage_LoadFromImage(pImage uintptr) int {]
ff=图片_加载从Image
pImage=GDI图片对象指针

[func XImage_LoadFromExtractIcon(pFileName string) int {]
ff=图片_加载文件图标
pFileName=文件名

[func XImage_LoadFromHICON(hIcon uintptr) int {]
ff=图片_加载从HICON
hIcon=图标句柄

[func XImage_LoadFromHBITMAP(hBitmap uintptr) int {]
ff=图片_加载从HBITMAP
hBitmap=位图句柄

[func XImage_IsStretch(hImage int) bool {]
ff=图片_判断缩放
hImage=图片句柄

[func XImage_IsAdaptive(hImage int) bool {]
ff=图片_判断自适应
hImage=图片句柄

[func XImage_IsTile(hImage int) bool {]
ff=图片_判断平铺
hImage=图片句柄

[func XImage_SetDrawType(hImage int, nType xcc.Image_Draw_Type_) bool {]
ff=图片_置绘制类型
nType=图片绘制类型
hImage=图片句柄

[func XImage_SetDrawTypeAdaptive(hImage int, leftSize, topSize, rightSize, bottomSize int32) bool {]
ff=图片_置绘制类型自适应
bottomSize=下
rightSize=右
topSize=上
leftSize=左
hImage=图片句柄

[func XImage_SetTranColor(hImage int, color int) {]
ff=图片_置透明色
color=ABGR颜色
hImage=图片句柄

[func XImage_SetTranColorEx(hImage int, color int, tranColor byte) {]
ff=图片_置透明色EX
tranColor=透明色的透明度
color=ABGR颜色
hImage=图片句柄

[func XImage_SetRotateAngle(hImage int, fAngle float32) float32 {]
ff=图片_置旋转角度
fAngle=选择角度
hImage=图片句柄

[func XImage_SetSplitEqual(hImage int, nCount int32, iIndex int32) {]
ff=图片_置等分
iIndex=索引
nCount=等分数量
hImage=图片句柄

[func XImage_EnableTranColor(hImage int, bEnable bool) {]
ff=图片_启用透明色
bEnable=启用TRUE
hImage=图片句柄

[func XImage_EnableAutoDestroy(hImage int, bEnable bool) {]
ff=图片_启用自动销毁
bEnable=启用自动销毁TRUE
hImage=图片句柄

[func XImage_EnableCenter(hImage int, bCenter bool) {]
ff=图片_启用居中
bCenter=是否居中显示
hImage=图片句柄

[func XImage_IsCenter(hImage int) bool {]
ff=图片_判断居中
hImage=图片句柄

[func XImage_GetDrawType(hImage int) xcc.Image_Draw_Type_ {]
ff=图片_取绘制类型
hImage=图片句柄

[func XImage_GetWidth(hImage int) int32 {]
ff=图片_取宽度
hImage=图片句柄

[func XImage_GetHeight(hImage int) int32 {]
ff=图片_取高度
hImage=图片句柄

[func XImage_GetImageSrc(hImage int) int {]
ff=图片_取图片源
hImage=图片句柄

[func XImage_AddRef(hImage int) {]
ff=图片_增加引用计数
hImage=图片句柄

[func XImage_Release(hImage int) {]
ff=图片_释放引用计数
hImage=图片句柄

[func XImage_GetRefCount(hImage int) int32 {]
ff=图片_取引用计数
hImage=图片句柄

[func XImage_Destroy(hImage int) {]
ff=图片_销毁
hImage=图片句柄

[func XImage_LoadSvg(hSvg int) int {]
ff=图片_加载从SVG
hSvg=SVG句柄

[func XImage_LoadSvgFile(pFileName string) int {]
ff=图片_加载从SVG文件
pFileName=文件名

[func XImage_LoadSvgString(pString string) int {]
ff=图片_加载从SVG字符串
pString=字符串

[func XImage_GetSvg(hImage int) int {]
ff=图片_取SVG
hImage=图片句柄

[func XImage_LoadSvgStringW(pString string) int {]
ff=图片_加载从SVG字符串W
pString=字符串

[func XImage_LoadSvgStringUtf8(pString string) int {]
ff=图片_加载从SVG字符串UTF8
pString=字符串

[func XImage_SetScaleSize(hImage int, width, height int32) {]
ff=图片_置缩放大小
height=高度
width=宽度
hImage=图片句柄
