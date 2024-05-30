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

[func XAnima_Run(hAnimation int, hRedrawObjectUI int) int {]
ff=动画_运行
hRedrawObjectUI=重绘的UI层
hAnimation=动画序列或动画组句柄

[func XAnima_Release(hAnimation int, bEnd bool) bool {]
ff=动画_释放
bEnd=是否立即执行到终点
hAnimation=动画序列或动画组句柄

[func XAnima_ReleaseEx(hObjectUI int, bEnd bool) int {]
ff=动画_释放EX
bEnd=是否立即执行到终点
hObjectUI=指定UI对象句柄

[func XAnima_Create(hObjectUI int, nLoopCount int) int {]
ff=动画_创建动画序列
nLoopCount=动画循环次数
hObjectUI=绑定的UI对象

[func XAnima_Move(hSequence int, duration int, x float32, y float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_移动
ease_flag=缓动标识
nLoopCount=动画循环次数
y=终点位置Y
x=终点位置X
duration=持续时间
hSequence=动画序列句柄

[func XAnima_MoveEx(hSequence int, duration int, from_x float32, from_y float32, to_x float32, to_y float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_移动EX
ease_flag=缓动标识
nLoopCount=动画循环次数
to_y=终点位置Y
to_x=终点位置X
from_y=起点位置Y
from_x=起点位置X
duration=持续时间
hSequence=动画序列句柄

[func XAnima_Rotate(hSequence int, duration int, angle float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_旋转
ease_flag=缓动标识
nLoopCount=动画循环次数
angle=角度
duration=持续时间
hSequence=动画序列句柄

[func XAnima_RotateEx(hSequence int, duration int, from float32, to float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_旋转EX
ease_flag=缓动标识
nLoopCount=动画循环次数
to=终点角度
from=起点角度
duration=持续时间
hSequence=动画序列句柄

[func XAnima_Scale(hSequence int, duration int, scaleX float32, scaleY float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_缩放
ease_flag=缓动标识
nLoopCount=动画循环次数
scaleY=Y轴缩放比例
scaleX=轴缩放比例
duration=持续时间
hSequence=动画序列句柄

[func XAnima_ScaleSize(hSequence int, duration int, width float32, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_缩放大小
ease_flag=缓动标识
nLoopCount=动画循环次数
height=高度
width=宽度
duration=持续时间
hSequence=动画序列句柄

[func XAnima_Alpha(hSequence int, duration int, alpha uint8, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_透明度
ease_flag=缓动标识
nLoopCount=动画循环次数
alpha=透明度
duration=持续时间
hSequence=动画序列句柄

[func XAnima_AlphaEx(hSequence int, duration int, from_alpha uint8, to_alpha uint8, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_透明度EX
ease_flag=缓动标识
nLoopCount=动画循环次数
to_alpha=终止透明度
from_alpha=起始透明度
duration=持续时间
hSequence=动画序列句柄

[func XAnima_Color(hSequence int, duration int, color int, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_颜色
ease_flag=缓动标识
nLoopCount=动画循环次数
color=ABGR
duration=持续时间
hSequence=动画序列句柄

[func XAnima_ColorEx(hSequence int, duration int, from int, to int, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_颜色EX
ease_flag=缓动标识
nLoopCount=动画循环次数
to=终点颜色
from=起点颜色
duration=持续时间
hSequence=动画序列句柄

[func XAnima_LayoutWidth(hSequence int, duration int, nType xcc.Layout_Size_, width float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_布局宽度
nType=布局宽度类型
duration=持续时间
hSequence=动画序列句柄

[func XAnima_LayoutHeight(hSequence int, duration int, nType xcc.Layout_Size_, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_布局高度
nType=布局高度类型
duration=持续时间
hSequence=动画序列句柄

[func XAnima_LayoutSize(hSequence int, duration int, nWidthType xcc.Layout_Size_, width float32, nHeightType xcc.Layout_Size_, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_布局大小
nWidthType=布局宽度类型
duration=持续时间
hSequence=动画序列句柄

[func XAnima_Delay(hSequence int, duration float32) int {]
ff=动画_延迟
duration=持续时间
hSequence=动画序列句柄

[func XAnima_Show(hSequence int, duration float32, bShow bool) int {]
ff=动画_显示
bShow=显示或隐藏
duration=持续时间
hSequence=动画序列句柄

[func XAnimaGroup_Create(nLoopCount int) int {]
ff=动画组_创建
nLoopCount=循环次数

[func XAnimaGroup_AddItem(hGroup int, hSequence int) int {]
ff=动画组_添加项
hSequence=动画序列句柄
hGroup=动画组句柄

[func XAnimaRotate_SetCenter(hAnimationRotate int, x float32, y float32, bOffset bool) bool {]
ff=动画旋转_置中心
bOffset=TRUE
y=坐标Y
x=坐标X
hAnimationRotate=动画旋转项句柄

[func XAnimaScale_SetPosition(hAnimationScale int, position xcc.Position_Flag_) bool {]
ff=动画缩放_置延伸位置
position=位置
hAnimationScale=动画缩放项句柄

[func XAnima_GetObjectUI(hAnimation int) int {]
ff=动画_取UI对象
hAnimation=动画序列或动画组句柄

[func XAnimaItem_EnableCompleteRelease(hAnimationItem int, bEnable bool) int {]
ff=动画项_启用完成释放
bEnable=是否启用
hAnimationItem=动画项句柄

[func XAnima_EnableAutoDestroy(hAnimation int, bEnable bool) int {]
ff=动画_启用自动销毁
bEnable=是否启用
hAnimation=动画项或动画序列或动画组句柄

[func XAnima_DestroyObjectUI(hSequence int, duration float32) int {]
ff=动画_销毁UI对象
duration=持续时间
hSequence=动画序列句柄

[func XAnima_SetCallBack(hAnimationEx int, callback interface{}) int {]
ff=动画_置回调
callback=回调函数
hAnimationEx=动画序列或动画组句柄

[func XAnima_SetUserData(hAnimationEx, nUserData int) int {]
ff=动画_置用户数据
nUserData=用户数据
hAnimationEx=动画序列或动画组句柄

[func XAnima_GetUserData(hAnimationEx int) int {]
ff=动画_取用户数据
hAnimationEx=动画序列或动画组句柄

[func XAnima_Stop(hAnimationEx int) bool {]
ff=动画_停止
hAnimationEx=动画序列或动画组句柄

[func XAnima_Start(hAnimationEx int) bool {]
ff=动画_启动
hAnimationEx=动画序列或动画组句柄

[func XAnima_Pause(hAnimationEx int) bool {]
ff=动画_暂停
hAnimationEx=动画序列或动画组句柄

[func XAnimaItem_SetCallback(hAnimationItem int, callback interface{}) int {]
ff=动画项_置回调
callback=回调函数
hAnimationItem=动画项句柄

[func XAnimaItem_SetUserData(hAnimationItem, nUserData int) int {]
ff=动画项_置用户数据
nUserData=用户数据
hAnimationItem=动画项句柄

[func XAnimaItem_GetUserData(hAnimationItem int) int {]
ff=动画项_取用户数据
hAnimationItem=动画项句柄

[func XAnimaItem_EnableAutoDestroy(hAnimationItem int, bEnable bool) int {]
ff=动画项_启用自动销毁
bEnable=是否启用
hAnimationItem=动画项句柄

[func XAnima_DelayEx(hSequence int, duration float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {]
ff=动画_延迟EX
ease_flag=缓动标识
nLoopCount=动画循环次数
duration=持续时间
hSequence=动画序列句柄

[func XAnimaMove_SetFlag(hAnimationMove int, flags xcc.Animation_Move_) int {]
ff=动画移动_置标识
flags=动画移动标识
hAnimationMove=动画移动项句柄
