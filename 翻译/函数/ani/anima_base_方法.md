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

[func (a *animaBase) Run(hRedrawObjectUI int) int {]
ff=运行
hRedrawObjectUI=当更新UI时重绘的UI层

[func (a *animaBase) Release(bEnd bool) bool {]
ff=释放
bEnd=是否立即执行到终点

[func (a *animaBase) ReleaseEx(bEnd bool) int {]
ff=释放EX
bEnd=是否立即执行到终点

[func (a *animaBase) GetObjectUI() int {]
ff=取UI对象

[func (a *animaBase) EnableAutoDestroy(bEnable bool) int {]
ff=启用自动销毁
bEnable=是否启用

[func (a *animaBase) SetCallBack(callback interface{}) int {]
ff=置回调
callback=回调函数

[func (a *animaBase) SetUserData(nUserData int) int {]
ff=置用户数据
nUserData=用户数据

[func (a *animaBase) GetUserData() int {]
ff=取用户数据

[func (a *animaBase) Stop() bool {]
ff=停止

[func (a *animaBase) Start() bool {]
ff=启动
[func (a *animaBase) Pause() bool {]
ff=暂停
