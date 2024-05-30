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

[func NewComboBox(x int, y int, cx int, cy int, hParent int) *ComboBox {]
ff=创建组合框
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewComboBoxByHandle(handle int) *ComboBox {]
ff=创建组合框并按句柄
handle=句柄

[func NewComboBoxByName(name string) *ComboBox {]
ff=创建组合框并按名称
name=名称

[func NewComboBoxByUID(nUID int) *ComboBox {]
ff=创建组合框并按UID

[func NewComboBoxByUIDName(name string) *ComboBox {]
ff=创建组合框并按UID名称
name=UID名称

[func (c *ComboBox) SetSelItem(iIndex int) bool {]
ff=置选择项
iIndex=项索引

[func (c *ComboBox) CreateAdapter() int {]
ff=创建数据适配器

[func (c *ComboBox) BindAdapter(hAdapter int) int {]
ff=绑定数据适配器
hAdapter=适配器句柄

[func (c *ComboBox) GetAdapter() int {]
ff=取数据适配器

[func (c *ComboBox) SetBindName(pName string) int {]
ff=置绑定名称
pName=字段名

[func (c *ComboBox) GetButtonRect(pRect *xc.RECT) int {]
ff=取下拉按钮坐标
pRect=坐标

[func (c *ComboBox) SetButtonSize(size int) int {]
ff=置下拉按钮大小
size=大小

[func (c *ComboBox) SetDropHeight(height int) int {]
ff=置下拉列表高度
height=高度

[func (c *ComboBox) GetDropHeight() int {]
ff=取下拉列表高度

[func (c *ComboBox) SetItemTemplateXML(pXmlFile string) int {]
ff=置项模板
pXmlFile=项模板文件

[func (c *ComboBox) SetItemTemplateXMLFromString(pStringXML string) int {]
ff=置项模板并按字符串
pStringXML=字符串

[func (c *ComboBox) EnableDrawButton(bEnable bool) int {]
ff=启用绘制下拉按钮
bEnable=是否绘制

[func (c *ComboBox) EnableEdit(bEdit bool) int {]
ff=启用编辑
bEdit=TRUE可编辑

[func (c *ComboBox) EnableDropHeightFixed(bEnable bool) int {]
ff=启用下拉列表高度固定大小
bEnable=是否启用

[func (c *ComboBox) GetSelItem() int {]
ff=取选择项

[func (c *ComboBox) GetState() xcc.ComboBox_State_ {]
ff=取状态

[func (c *ComboBox) AddItemText(pText string) int {]
ff=添加项文本
pText=文本

[func (c *ComboBox) AddItemTextEx(pName string, pText string) int {]
ff=添加项文本EX
pText=文本
pName=字段名

[func (c *ComboBox) AddItemImage(hImage int) int {]
ff=添加项图片
hImage=图片句柄

[func (c *ComboBox) AddItemImageEx(pName string, hImage int) int {]
ff=添加项图片EX
hImage=图片句柄
pName=字段名

[func (c *ComboBox) InsertItemText(iItem int, pText string) int {]
ff=插入项文本
pText=文本
iItem=项索引

[func (c *ComboBox) InsertItemTextEx(iItem int, pName string, pText string) int {]
ff=插入项文本EX
pText=文本
pName=字段名
iItem=项索引

[func (c *ComboBox) InsertItemImage(iItem int, hImage int) int {]
ff=插入项图片
hImage=图片句柄
iItem=项索引

[func (c *ComboBox) InsertItemImageEx(iItem int, pName string, hImage int) int {]
ff=插入项图片EX
hImage=图片句柄
pName=字段名
iItem=项索引

[func (c *ComboBox) SetItemText(iItem int, iColumn int, pText string) bool {]
ff=置项文本
pText=文本
iColumn=列索引
iItem=项索引

[func (c *ComboBox) SetItemTextEx(iItem int, pName string, pText string) bool {]
ff=置项文本EX
pText=文本
pName=字段名
iItem=项索引

[func (c *ComboBox) SetItemImage(iItem int, iColumn int, hImage int) bool {]
ff=置项图片
hImage=图片句柄
iColumn=列索引
iItem=项索引

[func (c *ComboBox) SetItemImageEx(iItem int, pName string, hImage int) bool {]
ff=置项图片EX
hImage=图片句柄
pName=字段名
iItem=项索引

[func (c *ComboBox) SetItemInt(iItem int, iColumn int, nValue int32) bool {]
ff=置项整数值
nValue=整数值
iColumn=列索引
iItem=项索引

[func (c *ComboBox) SetItemIntEx(iItem int, pName string, nValue int32) bool {]
ff=置项指数值EX
nValue=整数值
pName=字段名
iItem=项索引

[func (c *ComboBox) SetItemFloat(iItem int, iColumn int, fFloat float32) bool {]
ff=置项浮点值
fFloat=浮点数
iColumn=列索引
iItem=项索引

[func (c *ComboBox) SetItemFloatEx(iItem int, pName string, fFloat float32) bool {]
ff=置项浮点值EX
fFloat=浮点数
pName=字段名
iItem=项索引

[func (c *ComboBox) GetItemText(iItem int32, iColumn int32) string {]
ff=取项文本
iColumn=列索引
iItem=项索引

[func (c *ComboBox) GetItemTextEx(iItem int, pName string) string {]
ff=取项文本EX
pName=字段名
iItem=项索引

[func (c *ComboBox) GetItemImage(iItem int, iColumn int) int {]
ff=取项图片
iColumn=列索引
iItem=项索引

[func (c *ComboBox) GetItemImageEx(iItem int, pName string) int {]
ff=取项图片EX
pName=字段名
iItem=项索引

[func (c *ComboBox) GetItemInt(iItem int, iColumn int, pOutValue *int32) bool {]
ff=取项整数值
pOutValue=接收返回整数值
iColumn=列索引
iItem=项索引

[func (c *ComboBox) GetItemIntEx(iItem int, pName string, pOutValue *int32) bool {]
ff=取项整数值EX
pOutValue=接收返回整数值
pName=字段名
iItem=项索引

[func (c *ComboBox) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {]
ff=取项浮点值
pOutValue=接收返回浮点值
iColumn=列索引
iItem=项索引

[func (c *ComboBox) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {]
ff=取项浮点值EX
pOutValue=接收返回浮点值
pName=字段名
iItem=项索引

[func (c *ComboBox) DeleteItem(iItem int) bool {]
ff=删除项
iItem=项索引

[func (c *ComboBox) DeleteItemEx(iItem int, nCount int) bool {]
ff=删除项EX
nCount=删除数量
iItem=项索引

[func (c *ComboBox) DeleteItemAll() int {]
ff=删除项全部

[func (c *ComboBox) DeleteColumnAll() int {]
ff=删除列全部

[func (c *ComboBox) GetCount() int {]
ff=取项数量

[func (c *ComboBox) GetCountColumn() int {]
ff=取列数量

[func (c *ComboBox) PopupDropList() int {]
ff=弹出下拉列表

[func (c *ComboBox) SetItemTemplate(hTemp int) int {]
ff=设置项模板
hTemp=模板句柄

[func (c *ComboBox) SetItemTemplateXMLFromMem(data #左中括号##右中括号#byte) bool {]
ff=置项模板并按内存
data=模板数据

[func (c *ComboBox) SetItemTemplateXMLFromZipRes(id int32, pFileName string, pPassword string, hModule uintptr) bool {]
ff=置项模板并按资源ZIP
hModule=模块句柄
pPassword=zip密码
pFileName=文件名
id=RC资源ID

[func (c *ComboBox) GetItemTemplate() int {]
ff=取项模板

[func (c *ComboBox) Event_ComboBox_Select_End(pFun XE_COMBOBOX_SELECT_END) bool {]
ff=事件_下拉列表项选择完成

[func (c *ComboBox) Event_ComboBox_Select_End1(pFun XE_COMBOBOX_SELECT_END1) bool {]
ff=事件_下拉列表项选择完成1

[func (c *ComboBox) Event_COMBOBOX_SELECT(pFun XE_COMBOBOX_SELECT) bool {]
ff=事件_下拉列表项选择

[func (c *ComboBox) Event_COMBOBOX_SELECT1(pFun XE_COMBOBOX_SELECT1) bool {]
ff=事件_下拉列表项选择1

[func (c *ComboBox) Event_COMBOBOX_POPUP_LIST(pFun XE_COMBOBOX_POPUP_LIST) bool {]
ff=事件_下拉列表弹出

[func (c *ComboBox) Event_COMBOBOX_POPUP_LIST1(pFun XE_COMBOBOX_POPUP_LIST1) bool {]
ff=事件_下拉列表弹出1

[func (c *ComboBox) Event_COMBOBOX_EXIT_LIST(pFun XE_COMBOBOX_EXIT_LIST) bool {]
ff=事件_下拉列表退出

[func (c *ComboBox) Event_COMBOBOX_EXIT_LIST1(pFun XE_COMBOBOX_EXIT_LIST1) bool {]
ff=事件_下拉列表退出1
