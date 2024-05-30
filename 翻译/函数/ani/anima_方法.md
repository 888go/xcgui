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

[func NewAnima(hObjectUI int, nLoopCount int) *Anima {]
ff=创建动画序列
nLoopCount=动画循环次数
hObjectUI=绑定的UI对象

[func (a *Anima) Move(duration int, x float32, y float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=移动
ease_flag=缓动标识
nLoopCount=动画循环次数
y=终点位置Y
x=终点位置X
duration=持续时间

[func (a *Anima) MoveEx(duration int, from_x float32, from_y float32, to_x float32, to_y float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=移动EX
ease_flag=缓动标识
nLoopCount=动画循环次数
to_y=终点位置Y
to_x=终点位置X
from_y=起点位置Y
from_x=起点位置X
duration=持续时间

[func (a *Anima) Rotate(duration int, angle float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaRotate {]
ff=旋转
ease_flag=缓动标识
nLoopCount=动画循环次数
angle=角度
duration=持续时间

[func (a *Anima) RotateEx(duration int, from float32, to float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaRotate {]
ff=旋转EX
ease_flag=缓动标识
nLoopCount=动画循环次数
to=终点角度
from=起点角度
duration=持续时间

[func (a *Anima) Scale(duration int, scaleX float32, scaleY float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaScale {]
ff=缩放
ease_flag=缓动标识
nLoopCount=动画循环次数
scaleY=Y轴缩放比例
scaleX=轴缩放比例
duration=持续时间

[func (a *Anima) ScaleSize(duration int, width float32, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaScale {]
ff=缩放大小
ease_flag=缓动标识
nLoopCount=动画循环次数
height=高度
width=宽度
duration=持续时间

[func (a *Anima) Alpha(duration int, alpha uint8, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=透明度
ease_flag=缓动标识
nLoopCount=动画循环次数
alpha=透明度
duration=持续时间

[func (a *Anima) AlphaEx(duration int, from_alpha uint8, to_alpha uint8, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=透明度EX
ease_flag=缓动标识
nLoopCount=动画循环次数
to_alpha=终止透明度
from_alpha=起始透明度
duration=持续时间

[func (a *Anima) Color(duration int, color int, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=颜色
ease_flag=缓动标识
nLoopCount=动画循环次数
color=ABGR
duration=持续时间

[func (a *Anima) ColorEx(duration int, from int, to int, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=颜色EX
ease_flag=缓动标识
nLoopCount=动画循环次数
to=终点颜色
from=起点颜色
duration=持续时间

[func (a *Anima) LayoutWidth(duration int, nType xcc.Layout_Size_, width float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=布局宽度
nType=布局宽度类型
duration=持续时间

[func (a *Anima) LayoutHeight(duration int, nType xcc.Layout_Size_, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=布局高度
nType=布局高度类型
duration=持续时间

[func (a *Anima) LayoutSize(duration int, nWidthType xcc.Layout_Size_, width float32, nHeightType xcc.Layout_Size_, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {]
ff=布局大小
nWidthType=布局宽度类型
duration=持续时间

[func (a *Anima) Delay(duration float32) *AnimaItem {]
ff=延迟
duration=持续时间

[func (a *Anima) Show(duration float32, bShow bool) *AnimaItem {]
ff=显示
bShow=显示或隐藏
duration=持续时间

[func (a *Anima) DestroyObjectUI(duration float32) *AnimaItem {]
ff=销毁UI对象
duration=持续时间

[func (a *Anima) DelayEx(duration float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=延迟EX
ease_flag=缓动标识
nLoopCount=动画循环次数
duration=持续时间
