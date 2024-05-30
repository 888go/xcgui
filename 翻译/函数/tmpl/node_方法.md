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

[func NewNode(nType xcc.XC_OBJECT_TYPE) *Node {]
ff=创建节点
nType=对象类型

[func (n *Node) SetHandle(handle int) {]
ff=设置句柄
handle=句柄

[func (n *Node) GetNode(itemID int32) *Node {]
ff=取节点
itemID=ID

[func (n *Node) CloneNode() *Node {]
ff=取副本

[func (n *Node) AddNode(pNode int) bool {]
ff=添加子节点
pNode=节点指针

[func (n *Node) SetNodeAttribute(pName string, pAttr string) bool {]
ff=置节点属性
pAttr=属性值
pName=属性名

[func (n *Node) SetNodeAttributeEx(itemID int32, pName string, pAttr string) bool {]
ff=置节点属性EX
pAttr=属性值
pName=属性名
itemID=模板项ID
