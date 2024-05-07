package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 列表框.
type ListBox struct {
	ScrollView
}

// 列表框_创建, 创建列表框元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func X创建列表框(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *ListBox {
	p := &ListBox{}
	p.X设置句柄(炫彩基类.X列表框_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 列表框_创建Ex, 创建列表框元素, 使用内置项模板, 自动创建数据适配器.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
//
// col_extend_count: 列数量. 例如: 内置模板是1列, 如果数据有5列, 那么此参数填5.
func X创建列表框Ex(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄, 列数量 int32) *ListBox {
	p := &ListBox{}
	p.X设置句柄(炫彩基类.X列表框_创建Ex(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄, 列数量))
	return p
}

// 从句柄创建对象.
func X创建列表框并按句柄(handle int) *ListBox {
	p := &ListBox{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建列表框并按名称(name string) *ListBox {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ListBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建列表框并按UID(nUID int) *ListBox {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ListBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建列表框并按UID名称(name string) *ListBox {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ListBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 列表框_启用固定行高.
//
// bEnable: 是否启用.
func (l *ListBox) X启用固定行高(是否启用 bool) int {
	return 炫彩基类.X列表框_启用固定行高(l.Handle, 是否启用)
}

// 列表框_启用模板复用.
//
// bEnable: 是否启用.
func (l *ListBox) X启用模板复用(是否启用 bool) int {
	return 炫彩基类.X列表框_启用模板复用(l.Handle, 是否启用)
}

// 列表框_启用虚表.
//
// bEnable: 是否启用.
func (l *ListBox) X启用虚表(是否启用 bool) int {
	return 炫彩基类.X列表框_启用虚表(l.Handle, 是否启用)
}

// 列表框_置虚表行数.
//
// nRowCount: 行数.
func (l *ListBox) X置虚表行数(行数 int) int {
	return 炫彩基类.X列表框_置虚表行数(l.Handle, 行数)
}

// 列表框_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func (l *ListBox) X置绘制项背景标志(标志位 炫彩常量类.List_DrawItemBk_Flag_) int {
	return 炫彩基类.X列表框_置绘制项背景标志(l.Handle, 标志位)
}

// 列表框_置项数据, 设置项用户数据.
//
// iItem: 想索引.
//
// nUserData: 用户数据.
func (l *ListBox) X置项数据(想索引 int, 用户数据 int) bool {
	return 炫彩基类.X列表框_置项数据(l.Handle, 想索引, 用户数据)
}

// 列表框_取项数据, 获取项用户数据.
//
// iItem: 项索引.
func (l *ListBox) X取项数据(项索引 int) int {
	return 炫彩基类.X列表框_取项数据(l.Handle, 项索引)
}

// 列表框_置项信息.
//
// iItem: 项索引.
//
// pItem: 项信息.
func (l *ListBox) X置项信息(项索引 int, 项信息 *炫彩基类.ListBox_Item_Info_) bool {
	return 炫彩基类.X列表框_置项信息(l.Handle, 项索引, 项信息)
}

// 列表框_取项背景信息, 获取项信息.
//
// iItem: 项索引.
//
// pItem: 项信息.
func (l *ListBox) X取项背景信息(项索引 int, 项信息 *炫彩基类.ListBox_Item_Info_) bool {
	return 炫彩基类.X列表框_取项背景信息(l.Handle, 项索引, 项信息)
}

// 列表框_置选择项, 设置选择选.
//
// iItem: 项索引.
func (l *ListBox) X置选择项(项索引 int) bool {
	return 炫彩基类.X列表框_置选择项(l.Handle, 项索引)
}

// 列表框_取选择项, 返回项索引.
func (l *ListBox) X取选择项() int {
	return 炫彩基类.X列表框_取选择项(l.Handle)
}

// 列表框_添加选择项.
//
// iItem: 项索引.
func (l *ListBox) X添加选择项(项索引 int) bool {
	return 炫彩基类.X列表框_添加选择项(l.Handle, 项索引)
}

// 列表框_取消选择项.
//
// iItem: 项索引.
func (l *ListBox) X取消选择项(项索引 int) bool {
	return 炫彩基类.X列表框_取消选择项(l.Handle, 项索引)
}

// 列表框_取消选择全部, 如果之前有选择状态的项返回TRUE, 此时可以更新UI, 否则返回FALSE.
func (l *ListBox) X取消选择全部() bool {
	return 炫彩基类.X列表框_取消选择全部(l.Handle)
}

// 列表框_取全部选择, 获取所有选择项, 返回接收数量.
//
// pArray: 切片缓冲区.
//
// nArraySize: 切片大小.
func (l *ListBox) X取全部选择(切片缓冲区 *[]int32, 切片大小 int) int {
	return 炫彩基类.X列表框_取全部选择(l.Handle, 切片缓冲区, 切片大小)
}

// 列表框_取选择项数量, 获取选择项数量.
func (l *ListBox) X取选择项数量() int {
	return 炫彩基类.X列表框_取选择项数量(l.Handle)
}

// 列表框_取鼠标停留项, 返回鼠标所在项.
func (l *ListBox) X取鼠标停留项() int {
	return 炫彩基类.X列表框_取鼠标停留项(l.Handle)
}

// 列表框_选择全部项.
func (l *ListBox) X选择全部项() bool {
	return 炫彩基类.X列表框_选择全部项(l.Handle)
}

// 列表框_显示指定项, 滚动视图让指定项可见.
//
// iItem: 项索引.
func (l *ListBox) X显示指定项(项索引 int) int {
	return 炫彩基类.X列表框_显示指定项(l.Handle, 项索引)
}

// 列表框_取可视行范围, 获取当前可见行范围.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func (l *ListBox) X取可视行范围(开始行索引 *int32, 结束行索引 *int32) int {
	return 炫彩基类.X列表框_取可视行范围(l.Handle, 开始行索引, 结束行索引)
}

// 列表框_置项默认高度.
//
// nHeight: 项高度.
//
// nSelHeight: 选中项高度.
func (l *ListBox) X置项默认高度(项高度 int32, 选中项高度 int32) int {
	return 炫彩基类.X列表框_置项默认高度(l.Handle, 项高度, 选中项高度)
}

// 列表框_取项默认高度.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *ListBox) X取项默认高度(高度 *int32, 选中时高度 *int32) int {
	return 炫彩基类.X列表框_取项默认高度(l.Handle, 高度, 选中时高度)
}

// 列表框_取所在行索引, 获取当前对象所在模板实例, 属于列表中哪一个项(行). 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func (l *ListBox) X取所在行索引(对象句柄 int) int {
	return 炫彩基类.X列表框_取所在行索引(l.Handle, 对象句柄)
}

// 列表框_置行间距.
//
// nSpace: 间距大小.
func (l *ListBox) X置行间距(间距大小 int) int {
	return 炫彩基类.X列表框_置行间距(l.Handle, 间距大小)
}

// 列表框_取行间距.
func (l *ListBox) X取行间距() int {
	return 炫彩基类.X列表框_取行间距(l.Handle)
}

// 列表框_测试点击项, 检测坐标点所在项, 返回项索引.
//
// pPt: 坐标点.
func (l *ListBox) X测试点击项(坐标点 *炫彩基类.POINT) int {
	return 炫彩基类.X列表框_测试点击项(l.Handle, 坐标点)
}

// 列表框_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移量, 返回项索引.
//
// pPt: 坐标点.
func (l *ListBox) X测试点击项EX(坐标点 *炫彩基类.POINT) int {
	return 炫彩基类.X列表框_测试点击项EX(l.Handle, 坐标点)
}

// 列表框_置项模板文件, 设置列表项模板文件.
//
// pXmlFile: 文件名.
func (l *ListBox) X置项模板文件(文件名 string) bool {
	return 炫彩基类.X列表框_置项模板文件(l.Handle, 文件名)
}

// 列表框_置项模板, 设置列表项模板.
//
// hTemp: 模板句柄.
func (l *ListBox) X置项模板(模板句柄 int) bool {
	return 炫彩基类.X列表框_置项模板(l.Handle, 模板句柄)
}

// 列表框_置项模板从字符串, 设置项布局模板文件.
//
// pStringXML: 字符串.
func (l *ListBox) X置项模板并按字符串(字符串 string) bool {
	return 炫彩基类.X列表框_置项模板从字符串(l.Handle, 字符串)
}

// 列表框_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄, 成功返回对象句柄, 否则返回NULL.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func (l *ListBox) X取模板对象(项索引 int, 模板项ID int) int {
	return 炫彩基类.X列表框_取模板对象(l.Handle, 项索引, 模板项ID)
}

// 列表框_启用多选, 是否启用多行选择功能.
//
// bEnable: 是否启用.
func (l *ListBox) X启用多选(是否启用 bool) int {
	return 炫彩基类.X列表框_启用多选(l.Handle, 是否启用)
}

// 列表框_创建数据适配器, 创建数据适配器并绑定, 根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *ListBox) X创建数据适配器() int {
	return 炫彩基类.X列表框_创建数据适配器(l.Handle)
}

// 列表框_绑定数据适配器, 绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdTable.
func (l *ListBox) X绑定数据适配器(数据适配器句柄 int) int {
	return 炫彩基类.X列表框_绑定数据适配器(l.Handle, 数据适配器句柄)
}

// 列表框_取数据适配器, 获取绑定的数据适配器, 返回数据适配器句柄.
func (l *ListBox) X取数据适配器() int {
	return 炫彩基类.X列表框_取数据适配器(l.Handle)
}

// 列表框_排序.
//
// iColumnAdapter: 需要排序的数据在数据适配器中所属列索引.
//
// bAscending: 升序(TRUE)或降序(FALSE).
func (l *ListBox) X排序(列索引 int, 升序 bool) int {
	return 炫彩基类.X列表框_排序(l.Handle, 列索引, 升序)
}

// 列表框_刷新数据.
func (l *ListBox) X刷新数据() int {
	return 炫彩基类.X列表框_刷新数据(l.Handle)
}

// 列表框_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// iItem: 项索引.
func (l *ListBox) X刷新指定项(项索引 int) int {
	return 炫彩基类.X列表框_刷新指定项(l.Handle, 项索引)
}

// 列表框_添加项文本, XAdTable_AddItemText, 返回项索引.
//
// pText:.
func (l *ListBox) X添加项文本(文本 string) int {
	return 炫彩基类.X列表框_添加项文本(l.Handle, 文本)
}

// 列表框_添加项文本扩展, XAdTable_AddItemTextEx.
//
// pName:.
//
// pText:.
func (l *ListBox) X添加项文本EX(名称 string, 文本 string) int {
	return 炫彩基类.X列表框_添加项文本EX(l.Handle, 名称, 文本)
}

// 列表框_添加项图片, XAdTable_AddItemImage.
//
// hImage:.
func (l *ListBox) X添加项图片(图片 int) int {
	return 炫彩基类.X列表框_添加项图片(l.Handle, 图片)
}

// 列表框_添加项图片扩展, XAdTable_AddItemImageEx.
//
// pName:.
//
// hImage:.
func (l *ListBox) X添加项图片EX(名称 string, 图片 int) int {
	return 炫彩基类.X列表框_添加项图片EX(l.Handle, 名称, 图片)
}

// 列表框_插入项文本.
//
// iItem:.
//
// pValue:.
func (l *ListBox) X插入项文本(项 int, 文本 string) int {
	return 炫彩基类.X列表框_插入项文本(l.Handle, 项, 文本)
}

// 列表框_插入项文本扩展.
//
// iItem:.
//
// pName:.
//
// pValue:.
func (l *ListBox) X插入项文本EX(项 int, 名称 string, 文本 string) int {
	return 炫彩基类.X列表框_插入项文本EX(l.Handle, 项, 名称, 文本)
}

// 列表框_插入项图片.
//
// iItem:.
//
// hImage:.
func (l *ListBox) X插入项图片(项 int, 图片 int) int {
	return 炫彩基类.X列表框_插入项图片(l.Handle, 项, 图片)
}

// 列表框_插入项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *ListBox) X插入项图片EX(项 int, 名称 string, 图片 int) int {
	return 炫彩基类.X列表框_插入项图片EX(l.Handle, 项, 名称, 图片)
}

// 列表框_置项文本.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func (l *ListBox) X置项文本(项 int, 列 int, 文本 string) bool {
	return 炫彩基类.X列表框_置项文本(l.Handle, 项, 列, 文本)
}

// 列表框_置项文本扩展.
//
// iItem:.
//
// pName:.
//
// pText:.
func (l *ListBox) X置项文本EX(项 int, 名称 string, 文本 string) bool {
	return 炫彩基类.X列表框_置项文本EX(l.Handle, 项, 名称, 文本)
}

// 列表框_置项图片.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func (l *ListBox) X置项图片(项 int, 列 int, 图片 int) bool {
	return 炫彩基类.X列表框_置项图片(l.Handle, 项, 列, 图片)
}

// 列表框_置项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *ListBox) X置项图片EX(项 int, 名称 string, 图片 int) bool {
	return 炫彩基类.X列表框_置项图片EX(l.Handle, 项, 名称, 图片)
}

// 列表框_置项整数值.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func (l *ListBox) X置项整数值(项 int, 列 int, 文本 int) bool {
	return 炫彩基类.X列表框_置项整数值(l.Handle, 项, 列, 文本)
}

// 列表框_置项整数值扩展.
//
// iItem:.
//
// pName:.
//
// nValue:.
func (l *ListBox) X置项整数值EX(项 int, 名称 string, 文本 int) bool {
	return 炫彩基类.X列表框_置项整数值EX(l.Handle, 项, 名称, 文本)
}

// 列表框_置项浮点值.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func (l *ListBox) X置项浮点值(项 int, 列 int, 小数值 float32) bool {
	return 炫彩基类.X列表框_置项浮点值(l.Handle, 项, 列, 小数值)
}

// 列表框_置项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func (l *ListBox) X置项浮点值EX(项 int, 名称 string, 小数值 float32) bool {
	return 炫彩基类.X列表框_置项浮点值EX(l.Handle, 项, 名称, 小数值)
}

// 列表框_取项文本.
//
// iItem:.
//
// iColumn:.
func (l *ListBox) X取项文本(项 int, 列 int) string {
	return 炫彩基类.X列表框_取项文本(l.Handle, 项, 列)
}

// 列表框_取项文本扩展.
//
// iItem:.
//
// pName:.
func (l *ListBox) X取项文本EX(项 int, 名称 string) string {
	return 炫彩基类.X列表框_取项文本EX(l.Handle, 项, 名称)
}

// 列表框_取项图片.
//
// iItem:.
//
// iColumn:.
func (l *ListBox) X取项图片(项 int, 列 int) int {
	return 炫彩基类.X列表框_取项图片(l.Handle, 项, 列)
}

// 列表框_取项图片扩展.
//
// iItem:.
//
// pName:.
func (l *ListBox) X取项图片EX(项 int, 名称 string) int {
	return 炫彩基类.X列表框_取项图片EX(l.Handle, 项, 名称)
}

// 列表框_取项整数值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *ListBox) X取项整数值(项 int, 列 int, 值指针 *int32) bool {
	return 炫彩基类.X列表框_取项整数值(l.Handle, 项, 列, 值指针)
}

// 列表框_取项整数值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *ListBox) X取项整数值EX(项 int, 名称 string, 值指针 *int32) bool {
	return 炫彩基类.X列表框_取项整数值EX(l.Handle, 项, 名称, 值指针)
}

// 列表框_取项浮点值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *ListBox) X取项浮点值(项 int, 列 int, 值指针 *float32) bool {
	return 炫彩基类.X列表框_取项浮点值(l.Handle, 项, 列, 值指针)
}

// 列表框_取项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *ListBox) X取项浮点值EX(项 int, 名称 string, 值指针 *float32) bool {
	return 炫彩基类.X列表框_取项浮点值EX(l.Handle, 项, 名称, 值指针)
}

// 列表框_删除项.
//
// iItem:.
func (l *ListBox) X删除项(项 int) bool {
	return 炫彩基类.X列表框_删除项(l.Handle, 项)
}

// 列表框_删除项扩展.
//
// iItem:.
//
// nCount:.
func (l *ListBox) X删除项EX(项 int, 计数 int) bool {
	return 炫彩基类.X列表框_删除项EX(l.Handle, 项, 计数)
}

// 列表框_删除项全部.
func (l *ListBox) X删除项全部() int {
	return 炫彩基类.X列表框_删除项全部(l.Handle)
}

// 列表框_删除列全部.
func (l *ListBox) X删除列全部() int {
	return 炫彩基类.X列表框_删除列全部(l.Handle)
}

// 列表框_取项数量AD.
func (l *ListBox) X取项数量AD() int {
	return 炫彩基类.X列表框_取项数量AD(l.Handle)
}

// 列表框_取列数量AD.
func (l *ListBox) X取列数量AD() int {
	return 炫彩基类.X列表框_取列数量AD(l.Handle)
}

// 列表框_置分割线颜色.
//
// color: ABGR 颜色值.
func (l *ListBox) X置分割线颜色(ABGR颜色值 int) int {
	return 炫彩基类.X列表框_置分割线颜色(l.Handle, ABGR颜色值)
}

// 列表框_置拖动矩形颜色.
//
// color: ABGR 颜色值.
//
// width: 线宽度.
func (l *ListBox) X置拖动矩形颜色(ABGR颜色值, 线宽度 int) int {
	return 炫彩基类.X列表框_置拖动矩形颜色(l.Handle, ABGR颜色值, 线宽度)
}

// 列表框_置项模板从内存. 设置项模板文件.
//
// data: 模板数据.
func (l *ListBox) X置项模板从内存(模板数据 []byte) bool {
	return 炫彩基类.X列表框_置项模板从内存(l.Handle, 模板数据)
}

// 列表框_置项模板从资源ZIP. 设置项模板文件.
//
// id: RC资源ID.
//
// pFileName: 项模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (l *ListBox) X置项模板并按资源ZIP(RC资源ID int, 项模板文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	return 炫彩基类.X列表框_置项模板从资源ZIP(l.Handle, RC资源ID, 项模板文件名, zip密码, 模块句柄)
}

// 列表框_取项模板. 获取列表项模板, 返回项模板句柄.
func (l *ListBox) X取项模板() int {
	return 炫彩基类.X列表框_取项模板(l.Handle)
}

/*
以下都是事件
*/

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE func(pItem *炫彩基类.ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE1 func(hEle int, pItem *炫彩基类.ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE_END func(pItem *炫彩基类.ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LISTBOX_TEMP_CREATE_END1 func(hEle int, pItem *炫彩基类.ListBox_Item_, nFlag int32, pbHandled *bool) int

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LISTBOX_TEMP_DESTROY func(pItem *炫彩基类.ListBox_Item_, nFlag int, pbHandled *bool) int

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LISTBOX_TEMP_DESTROY1 func(hEle int, pItem *炫彩基类.ListBox_Item_, nFlag int, pbHandled *bool) int
type XE_LISTBOX_TEMP_ADJUST_COORDINATE func(pItem *炫彩基类.ListBox_Item_, pbHandled *bool) int            // 列表框元素,项模板调整坐标. 已停用.
type XE_LISTBOX_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *炫彩基类.ListBox_Item_, pbHandled *bool) int // 列表框元素,项模板调整坐标. 已停用.
type XE_LISTBOX_DRAWITEM func(hDraw int, pItem *炫彩基类.ListBox_Item_, pbHandled *bool) int               // 列表框元素,项绘制事件.
type XE_LISTBOX_DRAWITEM1 func(hEle int, hDraw int, pItem *炫彩基类.ListBox_Item_, pbHandled *bool) int    // 列表框元素,项绘制事件.
type XE_LISTBOX_SELECT func(iItem int32, pbHandled *bool) int                                          // 列表框元素,项选择事件.
type XE_LISTBOX_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int                               // 列表框元素,项选择事件.

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (l *ListBox) X事件_项模板创建(pFun XE_LISTBOX_TEMP_CREATE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTBOX_TEMP_CREATE, pFun)
}

// 列表框元素-项模板创建事件, 模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (l *ListBox) X事件_项模板创建1(pFun XE_LISTBOX_TEMP_CREATE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTBOX_TEMP_CREATE, pFun)
}

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (l *ListBox) X事件_项模板创建完成(pFun XE_LISTBOX_TEMP_CREATE_END) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTBOX_TEMP_CREATE_END, pFun)
}

// 列表框元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (l *ListBox) X事件_项模板创建完成1(pFun XE_LISTBOX_TEMP_CREATE_END1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTBOX_TEMP_CREATE_END, pFun)
}

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (l *ListBox) X事件_项模板销毁(pFun XE_LISTBOX_TEMP_DESTROY) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTBOX_TEMP_DESTROY, pFun)
}

// 列表框元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (l *ListBox) X事件_项模板销毁1(pFun XE_LISTBOX_TEMP_DESTROY1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTBOX_TEMP_DESTROY, pFun)
}

// 列表框元素,项模板调整坐标. 已停用.
func (l *ListBox) X停用_项模板调整坐标(pFun XE_LISTBOX_TEMP_ADJUST_COORDINATE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表框元素,项模板调整坐标. 已停用.
func (l *ListBox) X停用_项模板调整坐标1(pFun XE_LISTBOX_TEMP_ADJUST_COORDINATE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTBOX_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表框元素,项绘制事件.
func (l *ListBox) X事件_项绘制事件(pFun XE_LISTBOX_DRAWITEM) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTBOX_DRAWITEM, pFun)
}

// 列表框元素,项绘制事件.
func (l *ListBox) X事件_项绘制事件1(pFun XE_LISTBOX_DRAWITEM1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTBOX_DRAWITEM, pFun)
}

// 列表框元素,项选择事件.
func (l *ListBox) X事件_项选择事件(pFun XE_LISTBOX_SELECT) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTBOX_SELECT, pFun)
}

// 列表框元素,项选择事件.
func (l *ListBox) X事件_项选择事件1(pFun XE_LISTBOX_SELECT1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTBOX_SELECT, pFun)
}
