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

[func DragQueryFileW(hDrop uintptr, iFile uint32, lpszFile *string, cch uint32) int {]
ff=拖放文件取路径

[func DragQueryPoint(hDrop uintptr, ppt *xc.POINT) bool {]
ff=拖放文件取鼠标位置
ppt=接收鼠标指针的坐标
hDrop=句柄

[func ShellExecuteW(hwnd uintptr, lpOperation, lpFile, lpParameters, lpDirectory string, nShowCmd xcc.SW_) int {]
ff=对指定文件执行操作
nShowCmd=显示Cmd
lpDirectory=默认完整路径
lpParameters=参数
lpFile=文件名
lpOperation=操作类型
hwnd=父窗口句柄

[func SHBrowseForFolderW(browseInfo *BrowseInfoW) uintptr {]
ff=对话框打开文件夹
browseInfo=对话框选项

[func SHGetPathFromIDListW(pidl uintptr, pszPath *string) bool {]
ff=文件夹指针取实际路径
pszPath=返回文件路径
pidl=文件夹指针
