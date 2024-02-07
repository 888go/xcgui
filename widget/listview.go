package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 列表视图.
type ListView struct {
	ScrollView
}

// 列表视_创建.
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
func X创建列表视(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *ListView {
	p := &ListView{}
	p.X设置句柄(炫彩基类.X列表视_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 列表视_创建Ex. 创建列表视图元素, 使用内置项模板, 自动创建数据适配器.
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
func X创建列表视Ex(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄, 列数量 int32) *ListView {
	p := &ListView{}
	p.X设置句柄(炫彩基类.X列表视_创建Ex(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄, 列数量))
	return p
}

// 从句柄创建对象.
func X创建列表视并按句柄(handle int) *ListView {
	p := &ListView{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建列表视并按名称(name string) *ListView {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ListView{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建列表视并按UID(nUID int) *ListView {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ListView{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建列表视并按UID名称(name string) *ListView {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ListView{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 列表视_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *ListView) X创建数据适配器() int {
	return 炫彩基类.X列表视_创建数据适配器(l.Handle)
}

// 列表视_绑定数据适配器.
//
// hAdapter: 数据适配器XAdListView.
func (l *ListView) X绑定数据适配器(数据适配器 int) int {
	return 炫彩基类.X列表视_绑定数据适配器(l.Handle, 数据适配器)
}

// 列表视_取数据适配器, 返回数据适配器句柄.
func (l *ListView) X取数据适配器() int {
	return 炫彩基类.X列表视_取数据适配器(l.Handle)
}

// 列表视_置项模板文件.
//
// pXmlFile: 文件名.
func (l *ListView) X置项模板文件(文件名 string) bool {
	return 炫彩基类.X列表视_置项模板文件(l.Handle, 文件名)
}

// 列表视_置项模板从字符串.
//
// pStringXML: 字符串.
func (l *ListView) X置项模板从字符串(字符串 string) bool {
	return 炫彩基类.X列表视_置项模板从字符串(l.Handle, 字符串)
}

// 列表视_置项模板, 置列表项模板.
//
// hTemp: 模板句柄.
func (l *ListView) X置项模板(模板句柄 int) bool {
	return 炫彩基类.X列表视_置项模板(l.Handle, 模板句柄)
}

// 列表视_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nTempItemID: 模板项ID.
func (l *ListView) X取模板对象(组索引 int, 项索引 int, 模板项ID int) int {
	return 炫彩基类.X列表视_取模板对象(l.Handle, 组索引, 项索引, 模板项ID)
}

// 列表视_取模板对象组, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iGroup: 组索引.
//
// nTempItemID: 模板项ID.
func (l *ListView) X取模板对象组(组索引 int, 模板项ID int) int {
	return 炫彩基类.X列表视_取模板对象组(l.Handle, 组索引, 模板项ID)
}

// 列表视_取对象所在项, 获取当前对象所在模板实例, 属于列表视中哪一个项.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func (l *ListView) X取对象所在项(对象句柄 int, 接收组索引 *int32, 接收项索引 *int32) bool {
	return 炫彩基类.X列表视_取对象所在项(l.Handle, 对象句柄, 接收组索引, 接收项索引)
}

// 列表视_测试点击项, 检查坐标点所在项.
//
// pPt: 坐标点.
//
// pOutGroup: 接收组索引.
//
// pOutItem: 接收项索引.
func (l *ListView) X测试点击项(坐标点 *炫彩基类.POINT, 接收组索引 *int32, 接收项索引 *int32) bool {
	return 炫彩基类.X列表视_测试点击项(l.Handle, 坐标点, 接收组索引, 接收项索引)
}

// 列表视_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// pPt: 坐标点.
//
// pOutGroup: 接收做索引.
//
// pOutItem: 接收项索引.
func (l *ListView) X测试点击项EX(坐标点 *炫彩基类.POINT, 接收做索引 *int32, 接收项索引 *int32) bool {
	return 炫彩基类.X列表视_测试点击项EX(l.Handle, 坐标点, 接收做索引, 接收项索引)
}

// 列表视_启用多选.
//
// bEnable: 是否启用.
func (l *ListView) X启用多选(是否启用 bool) int {
	return 炫彩基类.X列表视_启用多选(l.Handle, 是否启用)
}

// 列表视_启用模板复用.
//
// bEnable: 是否启用.
func (l *ListView) X启用模板复用(是否启用 bool) int {
	return 炫彩基类.X列表视_启用模板复用(l.Handle, 是否启用)
}

// 列表视_启用虚表.
//
// bEnable: 是否启用.
func (l *ListView) X启用虚表(是否启用 bool) int {
	return 炫彩基类.X列表视_启用虚表(l.Handle, 是否启用)
}

// 列表视_置虚表项数量.
//
// iGroup: 组索引.
//
// nCount: 项数量.
func (l *ListView) X置虚表项数量(组索引 int, 项数量 int) bool {
	return 炫彩基类.X列表视_置虚表项数量(l.Handle, 组索引, 项数量)
}

// 列表视_置项背景绘制标志, 置是否绘制指定状态下项的背景.
//
// nFlags: 标志位: List_DrawItemBk_Flag_.
func (l *ListView) X置项背景绘制标志(标志位 炫彩常量类.List_DrawItemBk_Flag_) int {
	return 炫彩基类.X列表视_置项背景绘制标志(l.Handle, 标志位)
}

// 列表视_置选择项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) X置选择项(组索引 int, 项索引 int) bool {
	return 炫彩基类.X列表视_置选择项(l.Handle, 组索引, 项索引)
}

// 列表视_取选择项.
//
// piGroup: 接收组索引.
//
// piItem: 接收项索引.
func (l *ListView) X取选择项(接收组索引 *int32, 接收项索引 *int32) bool {
	return 炫彩基类.X列表视_取选择项(l.Handle, 接收组索引, 接收项索引)
}

// 列表视_添加选择项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) X添加选择项(组索引 int, 项索引 int) bool {
	return 炫彩基类.X列表视_添加选择项(l.Handle, 组索引, 项索引)
}

// 列表视_显示指定项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) X显示指定项(组索引 int, 项索引 int) int {
	return 炫彩基类.X列表视_显示指定项(l.Handle, 组索引, 项索引)
}

// 列表视_取可视项范围, 获取当前可见项范围.
//
// piGroup1: 可视开始组.
//
// piGroup2: 可视结束组.
//
// piStartGroup: 可视开始组.
//
// piStartItem: 可视开始项.
//
// piEndGroup: 可视结束组.
//
// piEndItem: 可视结束项.
func (l *ListView) X取可视项范围(组1 *int32, 组2 *int32, 可视开始组 *int32, 可视开始项 *int32, 可视结束组 *int32, 可视结束项 *int32) int {
	return 炫彩基类.X列表视_取可视项范围(l.Handle, 组1, 组2, 可视开始组, 可视开始项, 可视结束组, 可视结束项)
}

// 列表视_取选择项数量.
func (l *ListView) X取选择项数量() int {
	return 炫彩基类.X列表视_取选择项数量(l.Handle)
}

// 列表视_取选择项全部, 获取选择的项ID, 返回接收项数量.
//
// pArray: 数组.
//
// nArraySize: 数组大小.
func (l *ListView) X取选择项全部(数组 *[]炫彩基类.ListView_Item_Id_, 数组大小 int) int {
	return 炫彩基类.X列表视_取选择项全部(l.Handle, 数组, 数组大小)
}

// 列表视_置选择项全部, 选择所有的项.
func (l *ListView) X置选择项全部() int {
	return 炫彩基类.X列表视_置选择项全部(l.Handle)
}

// 列表视_取消选择项全部, 取消选择所有项.
func (l *ListView) X取消选择项全部() int {
	return 炫彩基类.X列表视_取消选择项全部(l.Handle)
}

// 列表视_置列间隔, 置列间隔大小.
//
// space: 间隔大小.
func (l *ListView) X置列间隔(间隔大小 int) int {
	return 炫彩基类.X列表视_置列间隔(l.Handle, 间隔大小)
}

// 列表视_置行间隔, 置行间隔大小.
//
// space: 间隔大小.
func (l *ListView) X置行间隔(间隔大小 int) int {
	return 炫彩基类.X列表视_置行间隔(l.Handle, 间隔大小)
}

// 列表视_置项大小.
//
// width: 宽度.
//
// height: 高度.
func (l *ListView) X置项大小(宽度 int, 高度 int) int {
	return 炫彩基类.X列表视_置项大小(l.Handle, 宽度, 高度)
}

// 列表视_取项大小.
//
// pSize: 接收返回大小.
func (l *ListView) X取项大小(接收返回大小 *炫彩基类.SIZE) int {
	return 炫彩基类.X列表视_取项大小(l.Handle, 接收返回大小)
}

// 列表视_置组高度.
//
// height: 高度.
func (l *ListView) X置组高度(高度 int) int {
	return 炫彩基类.X列表视_置组高度(l.Handle, 高度)
}

// 列表视_取组高度.
func (l *ListView) X取组高度() int {
	return 炫彩基类.X列表视_取组高度(l.Handle)
}

// 列表视_置组用户数据.
//
// iGroup: 组索引.
//
// nData: 数据.
func (l *ListView) X置组用户数据(组索引 int, 数据 int) int {
	return 炫彩基类.X列表视_置组用户数据(l.Handle, 组索引, 数据)
}

// 列表视_置项用户数据.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// nData: 数据.
func (l *ListView) X置项用户数据(组索引 int, 项索引 int, 数据 int) int {
	return 炫彩基类.X列表视_置项用户数据(l.Handle, 组索引, 项索引, 数据)
}

// 列表视_取组用户数据.
//
// iGroup: 组索引.
func (l *ListView) X取组用户数据(组索引 int) int {
	return 炫彩基类.X列表视_取组用户数据(l.Handle, 组索引)
}

// 列表视_取项用户数据.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) X取项用户数据(组索引 int, 项索引 int) int {
	return 炫彩基类.X列表视_取项用户数据(l.Handle, 组索引, 项索引)
}

// 列表视_刷新项数据.
func (l *ListView) X刷新项数据() int {
	return 炫彩基类.X列表视_刷新项数据(l.Handle)
}

// 列表视_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// iGroup: 组索引.
//
// iItem: 项索引, 如果为-1, 代表为组.
func (l *ListView) X刷新指定项(组索引 int, 项索引 int) int {
	return 炫彩基类.X列表视_刷新指定项(l.Handle, 组索引, 项索引)
}

// 列表视_展开组, 成功返回TRUE否则返回FALSE, 如果状态没有改变返回FALSE.
//
// iGroup: 组索引.
//
// bExpand: 是否展开.
func (l *ListView) X展开组(组索引 int, 是否展开 bool) bool {
	return 炫彩基类.X列表视_展开组(l.Handle, 组索引, 是否展开)
}

// 列表视_组添加列, 返回列索引.
//
// pName: 字段称.
func (l *ListView) X组添加列(字段称 string) int {
	return 炫彩基类.X列表视_组添加列(l.Handle, 字段称)
}

// 列表视_组添加项文本, 返回组索引.
//
// pValue: 值.
//
// iPos: 插入位置.
func (l *ListView) X组添加项文本(值 string, 插入位置 int) int {
	return 炫彩基类.X列表视_组添加项文本(l.Handle, 值, 插入位置)
}

// 列表视_组添加项文本扩展, 返回组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置.
func (l *ListView) X组添加项文本EX(字段称 string, 值 string, 插入位置 int) int {
	return 炫彩基类.X列表视_组添加项文本EX(l.Handle, 字段称, 值, 插入位置)
}

// 列表视_组添加项图片, 返回组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func (l *ListView) X组添加项图片(图片句柄 int, 插入位置 int) int {
	return 炫彩基类.X列表视_组添加项图片(l.Handle, 图片句柄, 插入位置)
}

// 列表视_组添加项图片扩展, 返回组索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置.
func (l *ListView) X组添加项图片EX(字段称 string, 图片句柄 int, 插入位置 int) int {
	return 炫彩基类.X列表视_组添加项图片EX(l.Handle, 字段称, 图片句柄, 插入位置)
}

// 列表视_组置文本.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (l *ListView) X组置文本(组索引 int, 列索引 int, 值 string) bool {
	return 炫彩基类.X列表视_组置文本(l.Handle, 组索引, 列索引, 值)
}

// 列表视_组置文本扩展.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func (l *ListView) X组置文本EX(组索引 int, 字段名 string, 值 string) bool {
	return 炫彩基类.X列表视_组置文本EX(l.Handle, 组索引, 字段名, 值)
}

// 列表视_组置图片.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (l *ListView) X组置图片(组索引 int, 列索引 int, 图片句柄 int) bool {
	return 炫彩基类.X列表视_组置图片(l.Handle, 组索引, 列索引, 图片句柄)
}

// 列表视_组置图片扩展.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (l *ListView) X组置图片EX(组索引 int, 字段名 string, 图片句柄 int) bool {
	return 炫彩基类.X列表视_组置图片EX(l.Handle, 组索引, 字段名, 图片句柄)
}

// 列表视_组获取数量, 返回组数量.
func (l *ListView) X组获取数量() int {
	return 炫彩基类.X列表视_组获取数量(l.Handle)
}

// 列表视_项获取数量, 成功返回项数量, 否则返回 XC_ID_ERROR.
//
// iGroup: 组索引.
func (l *ListView) X项获取数量(组索引 int) int {
	return 炫彩基类.X列表视_项获取数量(l.Handle, 组索引)
}

// 列表视_项添加列, 返回列索引.
//
// pName: 字段名.
func (l *ListView) X项添加列(字段名 string) int {
	return 炫彩基类.X列表视_项添加列(l.Handle, 字段名)
}

// 列表视_项添加文本, 返回项索引.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) X项添加文本(组索引 int, 值 string, 插入位置 int) int {
	return 炫彩基类.X列表视_项添加文本(l.Handle, 组索引, 值, 插入位置)
}

// 列表视_项添加文本扩展, 返回项索引.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) X项添加文本EX(组索引 int, 字段名 string, 值 string, 插入位置 int) int {
	return 炫彩基类.X列表视_项添加文本EX(l.Handle, 组索引, 字段名, 值, 插入位置)
}

// 列表视_项添加图片, 返回项索引.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) X项添加图片(组索引 int, 图片句柄 int, 插入位置 int) int {
	return 炫彩基类.X列表视_项添加图片(l.Handle, 组索引, 图片句柄, 插入位置)
}

// 列表视_项添加图片扩展, 返回项索引.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, -1添加到末尾.
func (l *ListView) X项添加图片EX(组索引 int, 字段名 string, 图片句柄 int, 插入位置 int) int {
	return 炫彩基类.X列表视_项添加图片EX(l.Handle, 组索引, 字段名, 图片句柄, 插入位置)
}

// 列表视_项置文本.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (l *ListView) X项置文本(组索引 int, 项索引 int, 列索引 int, 值 string) bool {
	return 炫彩基类.X列表视_项置文本(l.Handle, 组索引, 项索引, 列索引, 值)
}

// 列表视_项置文本扩展.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段名.
//
// pValue: 值.
func (l *ListView) X项置文本EX(组索引 int, 项索引 int, 字段名 string, 值 string) bool {
	return 炫彩基类.X列表视_项置文本EX(l.Handle, 组索引, 项索引, 字段名, 值)
}

// 列表视_项置图片.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (l *ListView) X项置图片(组索引 int, 项索引 int, 列索引 int, 图片句柄 int) bool {
	return 炫彩基类.X列表视_项置图片(l.Handle, 组索引, 项索引, 列索引, 图片句柄)
}

// 列表视_项置图片扩展.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 列名称.
//
// hImage: 图片句柄.
func (l *ListView) X项置图片EX(组索引 int, 项索引 int, 列名称 string, 图片句柄 int) bool {
	return 炫彩基类.X列表视_项置图片EX(l.Handle, 组索引, 项索引, 列名称, 图片句柄)
}

// 列表视_组删除项.
//
// iGroup: 组索引.
func (l *ListView) X组删除项(组索引 int) bool {
	return 炫彩基类.X列表视_组删除项(l.Handle, 组索引)
}

// 列表视_组删除全部子项.
//
// iGroup: 组索引.
func (l *ListView) X组删除全部子项(组索引 int) int {
	return 炫彩基类.X列表视_组删除全部子项(l.Handle, 组索引)
}

// 列表视_项删除.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (l *ListView) X项删除(组索引 int, 项索引 int) bool {
	return 炫彩基类.X列表视_项删除(l.Handle, 组索引, 项索引)
}

// 列表视_删除全部.
func (l *ListView) X删除全部() int {
	return 炫彩基类.X列表视_删除全部(l.Handle)
}

// 列表视_删除全部组.
func (l *ListView) X删除全部组() int {
	return 炫彩基类.X列表视_删除全部组(l.Handle)
}

// 列表视_删除全部项.
func (l *ListView) X删除全部项() int {
	return 炫彩基类.X列表视_删除全部项(l.Handle)
}

// 列表视_组删除列.
//
// iColumn: 列索引.
func (l *ListView) X组删除列(列索引 int) int {
	return 炫彩基类.X列表视_组删除列(l.Handle, 列索引)
}

// 列表视_项删除列.
//
// iColumn: 列索引.
func (l *ListView) X项删除列(列索引 int) int {
	return 炫彩基类.X列表视_项删除列(l.Handle, 列索引)
}

// 列表视_项获取文本.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (l *ListView) X项获取文本(组索引 int, 项索引 int, 字段称 string) string {
	return 炫彩基类.X列表视_项获取文本(l.Handle, 组索引, 项索引, 字段称)
}

// 列表视_项获取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (l *ListView) X项获取图片EX(组索引 int, 项索引 int, 字段称 string) int {
	return 炫彩基类.X列表视_项获取图片EX(l.Handle, 组索引, 项索引, 字段称)
}

// 列表视_组取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (l *ListView) X组取文本(组索引 int, 列索引 int) string {
	return 炫彩基类.X列表视_组取文本(l.Handle, 组索引, 列索引)
}

// 列表视_组取文本扩展, 返回文本内容.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (l *ListView) X组取文本EX(组索引 int, 字段名称 string) string {
	return 炫彩基类.X列表视_组取文本EX(l.Handle, 组索引, 字段名称)
}

// 列表视_组取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (l *ListView) X组取图片(组索引 int, 列索引 int) int {
	return 炫彩基类.X列表视_组取图片(l.Handle, 组索引, 列索引)
}

// 列表视_组取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (l *ListView) X组取图片EX(组索引 int, 字段名称 string) int {
	return 炫彩基类.X列表视_组取图片EX(l.Handle, 组索引, 字段名称)
}

// 列表视_项取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (l *ListView) X项取文本(组索引 int, 项索引 int, 列索引 int) string {
	return 炫彩基类.X列表视_项取文本(l.Handle, 组索引, 项索引, 列索引)
}

// 列表视_项取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (l *ListView) X项取图片(组索引 int, 项索引 int, 列索引 int) int {
	return 炫彩基类.X列表视_项取图片(l.Handle, 组索引, 项索引, 列索引)
}

// 列表视_置拖动矩形颜色.
//
// color: ABGR 颜色.
//
// width: 线宽度.
func (l *ListView) X置拖动矩形颜色(ABGR颜色 int, 线宽度 int) int {
	return 炫彩基类.X列表视_置拖动矩形颜色(l.Handle, ABGR颜色, 线宽度)
}

// 列表视_置项模板从内存.
//
// data: 模板数据.
func (l *ListView) X置项模板并按内存(模板数据 []byte) bool {
	return 炫彩基类.X列表视_置项模板从内存(l.Handle, 模板数据)
}

// 列表视_置项模板从资源ZIP.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (l *ListView) X置项模板并按资源ZIP(RC资源ID int32, 文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	return 炫彩基类.X列表视_置项模板从资源ZIP(l.Handle, RC资源ID, 文件名, zip密码, 模块句柄)
}

// 列表视_取项模板, 返回项模板句柄.
func (l *ListView) X取项模板() int {
	return 炫彩基类.X列表视_取项模板(l.Handle)
}

// 列表视_取项模板组, 返回项模板组句柄.
func (l *ListView) X取项模板组() int {
	return 炫彩基类.X列表视_取项模板组(l.Handle)
}

/*
以下都是事件
*/

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
type XE_LISTVIEW_TEMP_CREATE func(pItem *炫彩基类.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
type XE_LISTVIEW_TEMP_CREATE1 func(hEle int, pItem *炫彩基类.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
type XE_LISTVIEW_TEMP_CREATE_END func(pItem *炫彩基类.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
type XE_LISTVIEW_TEMP_CREATE_END1 func(hEle int, pItem *炫彩基类.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
type XE_LISTVIEW_TEMP_DESTROY func(pItem *炫彩基类.ListView_Item_, nFlag int32, pbHandled *bool) int

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
type XE_LISTVIEW_TEMP_DESTROY1 func(hEle int, pItem *炫彩基类.ListView_Item_, nFlag int32, pbHandled *bool) int
type XE_LISTVIEW_TEMP_ADJUST_COORDINATE func(pItem *炫彩基类.ListView_Item_, pbHandled *bool) int            // 列表视元素,项模板调整坐标.已停用.
type XE_LISTVIEW_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *炫彩基类.ListView_Item_, pbHandled *bool) int // 列表视元素,项模板调整坐标.已停用.
type XE_LISTVIEW_DRAWITEM func(hDraw int, pItem *炫彩基类.ListView_Item_, pbHandled *bool) int               // 列表视元素,自绘项.
type XE_LISTVIEW_DRAWITEM1 func(hEle int, hDraw int, pItem *炫彩基类.ListView_Item_, pbHandled *bool) int    // 列表视元素,自绘项.
type XE_LISTVIEW_SELECT func(iGroup int32, iItem int32, pbHandled *bool) int                           // 列表视元素,项选择事件.
type XE_LISTVIEW_SELECT1 func(hEle int, iGroup int32, iItem int32, pbHandled *bool) int                // 列表视元素,项选择事件.
type XE_LISTVIEW_EXPAND func(iGroup int32, bExpand bool, pbHandled *bool) int                          // 列表视元素,组展开收缩事件.
type XE_LISTVIEW_EXPAND1 func(hEle int, iGroup int32, bExpand bool, pbHandled *bool) int               // 列表视元素,组展开收缩事件.

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
func (l *ListView) X事件_项模板创建(pFun XE_LISTVIEW_TEMP_CREATE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTVIEW_TEMP_CREATE, pFun)
}

// 列表视元素-项模板创建事件,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变(当前未使用); 1新模板实例; 2旧模板复用
func (l *ListView) X事件_项模板创建1(pFun XE_LISTVIEW_TEMP_CREATE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTVIEW_TEMP_CREATE, pFun)
}

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
func (l *ListView) X事件_项模板创建完成(pFun XE_LISTVIEW_TEMP_CREATE_END) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTVIEW_TEMP_CREATE_END, pFun)
}

// 列表视元素-项模板创建完成事件,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用,当前未使用); 1:新模板实例; 2:旧模板复用
func (l *ListView) X事件_项模板创建完成1(pFun XE_LISTVIEW_TEMP_CREATE_END1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTVIEW_TEMP_CREATE_END, pFun)
}

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
func (l *ListView) X事件_项模板销毁(pFun XE_LISTVIEW_TEMP_DESTROY) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTVIEW_TEMP_DESTROY, pFun)
}

// 列表视元素-项模板销毁, 模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存列表(不会被销毁, 临时缓存备用, 当需要时被复用)
func (l *ListView) X事件_项模板销毁1(pFun XE_LISTVIEW_TEMP_DESTROY1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTVIEW_TEMP_DESTROY, pFun)
}

// 列表视元素,项模板调整坐标.已停用.
func (l *ListView) X事件_项模板调整坐标(pFun XE_LISTVIEW_TEMP_ADJUST_COORDINATE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTVIEW_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表视元素,项模板调整坐标.已停用.
func (l *ListView) X事件_项模板调整坐标1(pFun XE_LISTVIEW_TEMP_ADJUST_COORDINATE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTVIEW_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表视元素,自绘项.
func (l *ListView) X事件_自绘项(pFun XE_LISTVIEW_DRAWITEM) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTVIEW_DRAWITEM, pFun)
}

// 列表视元素,自绘项.
func (l *ListView) X事件_自绘项1(pFun XE_LISTVIEW_DRAWITEM1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTVIEW_DRAWITEM, pFun)
}

// 列表视元素,项选择事件.
func (l *ListView) X事件_项选择(pFun XE_LISTVIEW_SELECT) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTVIEW_SELECT, pFun)
}

// 列表视元素,项选择事件.
func (l *ListView) X事件_项选择1(pFun XE_LISTVIEW_SELECT1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTVIEW_SELECT, pFun)
}

// 列表视元素,组展开收缩事件.
func (l *ListView) X事件_组展开收缩(pFun XE_LISTVIEW_EXPAND) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LISTVIEW_EXPAND, pFun)
}

// 列表视元素,组展开收缩事件.
func (l *ListView) X事件_组展开收缩1(pFun XE_LISTVIEW_EXPAND1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LISTVIEW_EXPAND, pFun)
}
