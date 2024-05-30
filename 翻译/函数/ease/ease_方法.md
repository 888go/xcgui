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

[func Linear(p float32) float32 {]
ff=缓动线性
p=位置

[func Quad(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动二次方曲线
flag=缓动类型
p=位置

[func Cubic(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动三次方曲线
flag=缓动类型
p=位置

[func Quart(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动四次方曲线
flag=缓动类型
p=位置

[func Quint(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动五次方曲线
flag=缓动类型
p=位置

[func Sine(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动正弦曲线
flag=缓动类型
p=位置

[func Expo(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动突击曲线
flag=缓动类型
p=位置

[func Circ(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动圆环
flag=缓动类型
p=位置

[func Elastic(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动强力回弹
flag=缓动类型
p=位置

[func Back(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动回弹
flag=缓动类型
p=位置

[func Bounce(p float32, flag xcc.Ease_Type_) float32 {]
ff=缓动弹跳
flag=缓动类型
p=位置

[func Ex(pos float32, flag xcc.Ease_Flag_) float32 {]
ff=缓动EX
flag=缓动标识
pos=位置
