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

[func GetDropFiles(hDropInfo uintptr) #左中括号##右中括号#string {]
ff=拖放文件取路径
hDropInfo=拖放信息句柄

[func OpenDir(hParent int) string {]
ff=对话框打开文件夹
hParent=炫彩窗口句柄

[func OpenFile(hParent int, filters #左中括号##右中括号#string, defaultDir string) string {]
ff=对话框打开单个文件
defaultDir=初始目录
filters=过滤器切片
hParent=炫彩窗口句柄

[func OpenFiles(hParent int, filters #左中括号##右中括号#string, defaultDir string) #左中括号##右中括号#string {]
ff=对话框打开多个文件
defaultDir=初始目录
filters=过滤器切片
hParent=炫彩窗口句柄

[func SaveFile(hParent int, filters #左中括号##右中括号#string, defaultDir, defaultFileName string) string {]
ff=对话框保存文件
defaultFileName=默认文件名
defaultDir=初始目录
filters=过滤器切片
hParent=炫彩窗口句柄

[func ChooseColor(hParent int) int {]
ff=对话框选择颜色
hParent=炫彩窗口句柄
