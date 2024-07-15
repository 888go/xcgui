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

[func NewBySrc(hImageSrc int) *Image {]
ff=创建并按图片源句柄
hImageSrc=图片源句柄

[func NewByFile(pFileName string) *Image {]
ff=创建并按文件
pFileName=图片文件

[func NewByFileAdaptive(pFileName string, leftSize, topSize, rightSize, bottomSize int32) *Image {]
ff=创建并按文件且自适应
bottomSize=下
rightSize=右
topSize=上
leftSize=左
pFileName=图片文件

[func NewByFileRect(pFileName string, x, y, cx, cy int32) *Image {]
ff=创建并按文件且指定区域
cy=高度
cx=宽度
y=坐标y
x=坐标x
pFileName=图片文件

[func NewByResAdaptive(id int32, pType string, leftSize, topSize, rightSize, bottomSize int32, hModule uintptr) *Image {]
ff=创建并按资源且自适应
hModule=从指定模块加载
bottomSize=下
rightSize=右
topSize=上
leftSize=左
pType=资源类型
id=资源ID

[func NewByRes(id int32, pType string, bStretch bool, hModule uintptr) *Image {]
ff=创建并按资源
hModule=从指定模块加载
bStretch=是否拉伸图片
pType=资源类型
id=资源ID

[func NewByZip(pZipFileName string, pFileName string, pPassword string) *Image {]
ff=创建并按ZIP
pPassword=ZIP压缩包密码
pFileName=图片文件名
pZipFileName=ZIP压缩包文件名

[func NewByZipRes(id int32, pFileName string, pPassword string, hModule uintptr) *Image {]
ff=创建并按资源ZIP
hModule=模块句柄
pPassword=ZIP压缩包密码
pFileName=图片文件名
id=RC资源ID

[func NewByZipAdaptive(pZipFileName string, pFileName string, pPassword string, x1, x2, y1, y2 int32) *Image {]
ff=创建并按ZIP且自适应
y2=坐标y2
y1=坐标y1
x2=坐标x2
x1=坐标x1
pPassword=ZIP压缩包密码
pFileName=图片文件名
pZipFileName=ZIP压缩包文件名

[func NewByZipRect(pZipFileName string, pFileName string, pPassword string, x, y, cx, cy int32) *Image {]
ff=创建并按ZIP且指定区域
cy=高度
cx=宽度
y=坐标y
x=坐标x
pPassword=密码
pFileName=图片名称
pZipFileName=ZIP文件

[func NewByZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string) *Image {]
ff=创建并按内存ZIP
pPassword=zip压缩包密码
pFileName=图片名称
data=图片数据

[func NewByMem(pBuffer #左中括号##右中括号#byte) *Image {]
ff=创建并按内存
pBuffer=图片数据

[func NewByMemRect(pBuffer #左中括号##右中括号#byte, x, y, cx, cy int32) *Image {]
ff=创建并按内存且指定区域
cy=高度
cx=宽度
y=坐标y
x=坐标x
pBuffer=图片数据

[func NewByMemAdaptive(pBuffer #左中括号##右中括号#byte, leftSize, topSize, rightSize, bottomSize int32) *Image {]
ff=创建并按内存且自适应
bottomSize=坐标下
rightSize=坐标右
topSize=坐标上
leftSize=坐标左
pBuffer=图片数据

[func NewByImage(pImage uintptr) *Image {]
ff=创建并按GDI图片对象指针
pImage=GDI图片对象指针

[func NewByExtractIcon(pFileName string) *Image {]
ff=创建并按图标文件
pFileName=文件名

[func NewByHICON(hIcon uintptr) *Image {]
ff=创建并按图标句柄
hIcon=图标句柄

[func NewByHBITMAP(hBitmap uintptr) *Image {]
ff=创建并按位图句柄
hBitmap=位图句柄

[func NewBySvg(hSvg int) *Image {]
ff=创建并按SVG
hSvg=SVG句柄

[func NewBySvgFile(pFileName string) *Image {]
ff=创建并按SVG文件
pFileName=文件名

[func NewBySvgString(pString string) *Image {]
ff=创建并按SVG字符串
pString=字符串

[func NewBySvgStringW(pString string) *Image {]
ff=创建并按SVG字符串W
pString=字符串

[func NewBySvgStringUtf8(pString string) *Image {]
ff=创建并按SVG字符串UTF8
pString=字符串

[func NewByHandle(handle int) *Image {]
ff=创建并按句柄
handle=句柄

[func NewByName(name string) *Image {]
ff=创建并按资源名称
name=资源名称

[func NewByNameEx(fileName, name string) *Image {]
ff=创建并按资源文件名
name=资源名称
fileName=资源文件名

[func (i *Image) GetSvg() int {]
ff=取SVG

[func (i *Image) IsStretch() bool {]
ff=判断缩放

[func (i *Image) IsAdaptive() bool {]
ff=判断自适应

[func (i *Image) IsTile() bool {]
ff=判断平铺

[func (i *Image) SetDrawType(nType xcc.Image_Draw_Type_) bool {]
ff=置绘制类型
nType=图片绘制类型

[func (i *Image) SetDrawTypeAdaptive(leftSize, topSize, rightSize, bottomSize int32) bool {]
ff=置绘制类型自适应
bottomSize=坐标下
rightSize=坐标右
topSize=坐标上
leftSize=坐标左

[func (i *Image) SetTranColor(color int) {]
ff=置透明色
color=ABGR颜色

[func (i *Image) SetTranColorEx(color int, tranColor byte) {]
ff=置透明色EX
tranColor=透明度
color=ABGR颜色

[func (i *Image) SetRotateAngle(fAngle float32) float32 {]
ff=置旋转角度
fAngle=选择角度

[func (i *Image) SetSplitEqual(nCount, iIndex int32) {]
ff=置等分
iIndex=索引
nCount=等分数量

[func (i *Image) EnableTranColor(bEnable bool) {]
ff=启用透明色
bEnable=启用TRUE

[func (i *Image) EnableAutoDestroy(bEnable bool) {]
ff=启用自动销毁
bEnable=启用自动销毁TRUE

[func (i *Image) EnableCenter(bCenter bool) {]
ff=启用居中
bCenter=是否居中显示

[func (i *Image) IsCenter() bool {]
ff=判断居中

[func (i *Image) GetDrawType() xcc.Image_Draw_Type_ {]
ff=取绘制类型

[func (i *Image) GetWidth() int32 {]
ff=取宽度

[func (i *Image) GetHeight() int32 {]
ff=取高度

[func (i *Image) GetImageSrc() int {]
ff=取图片源

[func (i *Image) AddRef() {]
ff=增加引用计数

[func (i *Image) Release() {]
ff=释放引用计数

[func (i *Image) GetRefCount() int32 {]
ff=取引用计数

[func (i *Image) Destroy() {]
ff=销毁

[func (i *Image) SetScaleSize(width, height int32) {]
ff=置缩放大小
height=高度
width=宽度
