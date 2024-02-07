package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 列表树元素.
type Tree struct {
	ScrollView
}

// 列表树_创建, 创建树元素.
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
func X创建列表树(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄 int) *Tree {
	p := &Tree{}
	p.X设置句柄(炫彩基类.X列表树_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 列表树_创建Ex, 创建树元素, 使用内置项模板, 自动创建数据适配器.
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
func X创建列表树Ex(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄, 列数量 int32) *Tree {
	p := &Tree{}
	p.X设置句柄(炫彩基类.X列表树_创建Ex(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄, 列数量))
	return p
}

// 从句柄创建对象.
func X创建列表树并按句柄(handle int) *Tree {
	p := &Tree{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建列表树并按名称(name string) *Tree {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &Tree{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建列表树并按UID(nUID int) *Tree {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Tree{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建列表树并按UID名称(name string) *Tree {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &Tree{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 列表树_启用拖动项, 启用拖动项功能.
//
// bEnable: 是否启用.
func (t *Tree) X启用拖动项(是否启用 bool) {
	炫彩基类.X列表树_启用拖动项(t.Handle, 是否启用)
}

// 列表树_启用连接线, 启用或禁用显示项的连接线.
//
// bEnable: 是否启用.
//
// bSolid: 实线或虚线; TRUE: 实线, FALSE: 虚线.
func (t *Tree) X启用连接线(是否启用 bool, 实线或虚线 bool) {
	炫彩基类.X列表树_启用连接线(t.Handle, 是否启用, 实线或虚线)
}

// 列表树_启用展开, 启动或关闭默认展开功能, 如果开启新插入的项将自动展开.
//
// bEnable: 是否启用.
func (t *Tree) X启用展开(是否启用 bool) {
	炫彩基类.X列表树_启用展开(t.Handle, 是否启用)
}

// 列表树_启用模板复用.
//
// bEnable: 是否启用.
func (t *Tree) X启用模板复用(是否启用 bool) {
	炫彩基类.X列表树_启用模板复用(t.Handle, 是否启用)
}

// 列表树_置连接线颜色.
//
// color: ABGR 颜色.
func (t *Tree) X置连接线颜色(ABGR颜色 int) {
	炫彩基类.X列表树_置连接线颜色(t.Handle, ABGR颜色)
}

// 列表树_置展开按钮大小, 设置展开按钮占用空间大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (t *Tree) X置展开按钮大小(宽度, 高度 int32) {
	炫彩基类.X列表树_置展开按钮大小(t.Handle, 宽度, 高度)
}

// 列表树_置连接线长度, 设置连线绘制长度, 展开按钮与项内容之间的连线.
//
// nLength: 连线绘制长度.
func (t *Tree) X置连接线长度(连线绘制长度 int32) {
	炫彩基类.X列表树_置连接线长度(t.Handle, 连线绘制长度)
}

// 列表树_置拖动项插入位置颜色, 设置拖动项插入位置颜色提示.
//
// color: ABGR 颜色.
func (t *Tree) X置拖动项插入位置颜色(ABGR颜色 int) {
	炫彩基类.X列表树_置拖动项插入位置颜色(t.Handle, ABGR颜色)
}

// 列表树_置项模板文件.
//
// pXmlFile: 文件名.
func (t *Tree) X置项模板文件(文件名 string) bool {
	return 炫彩基类.X列表树_置项模板文件(t.Handle, 文件名)
}

// 列表树_置选择项模板文件, 设置项模板文件, 项选中状态.
//
// pXmlFile: 文件名.
func (t *Tree) X置选择项模板文件(文件名 string) bool {
	return 炫彩基类.X列表树_置选择项模板文件(t.Handle, 文件名)
}

// 列表树_置项模板.
//
// hTemp: 模板句柄.
func (t *Tree) X置项模板(模板句柄 int) bool {
	return 炫彩基类.X列表树_置项模板(t.Handle, 模板句柄)
}

// 列表树_置选择项模板, 设置列表项模板, 项选中状态.
//
// hTemp: 模板句柄.
func (t *Tree) X置选择项模板(模板句柄 int) bool {
	return 炫彩基类.X列表树_置选择项模板(t.Handle, 模板句柄)
}

// 列表树_置项模板从字符串, 设置项模板文件.
//
// pStringXML: 字符串.
func (t *Tree) X置项模板并按字符串(字符串 string) bool {
	return 炫彩基类.X列表树_置项模板从字符串(t.Handle, 字符串)
}

// 列表树_置选择项模板从字符串, 设置项模板文件, 项选中状态.
//
// pStringXML: 字符串.
func (t *Tree) X置选择项模板并按字符串(字符串 string) bool {
	return 炫彩基类.X列表树_置选择项模板从字符串(t.Handle, 字符串)
}

// 列表树_置项背景绘制标志, 设置是否绘制指定状态下项的背景.
//
// nFlags: 标志位: xcc.List_DrawItemBk_Flag_.
func (t *Tree) X置项背景绘制标志(标志位 炫彩常量类.List_DrawItemBk_Flag_) {
	炫彩基类.X列表树_置项背景绘制标志(t.Handle, 标志位)
}

// 列表树_置项数据, 设置项用户数据.
//
// nID: 项ID.
//
// nUserData: 用户数据.
func (t *Tree) X置项数据(项ID int32, 用户数据 int) bool {
	return 炫彩基类.X列表树_置项数据(t.Handle, 项ID, 用户数据)
}

// 列表树_取项数据, 获取项用户数据.
//
// nID: 项ID.
func (t *Tree) X取项数据(项ID int32) int {
	return 炫彩基类.X列表树_取项数据(t.Handle, 项ID)
}

// 列表树_置选择项.
//
// nID: 项ID.
func (t *Tree) X置选择项(项ID int32) bool {
	return 炫彩基类.X列表树_置选择项(t.Handle, 项ID)
}

// 列表树_取选择项, 返回项ID.
func (t *Tree) X取选择项() int32 {
	return 炫彩基类.X列表树_取选择项(t.Handle)
}

// 列表树_可视指定项, 滚动视图让指定项可见.
//
// nID: 项索引.
func (t *Tree) X可视指定项(项索引 int32) {
	炫彩基类.X列表树_可视指定项(t.Handle, 项索引)
}

// 列表树_判断展开.
//
// nID: 项ID.
func (t *Tree) X判断展开(项ID int32) bool {
	return 炫彩基类.X列表树_判断展开(t.Handle, 项ID)
}

// 列表树_展开项, 判断项是否展开.
//
// nID: 项ID.
//
// bExpand: 是否展开.
func (t *Tree) X展开项(项ID int32, 是否展开 bool) bool {
	return 炫彩基类.X列表树_展开项(t.Handle, 项ID, 是否展开)
}

// 列表树_展开全部子项, 展开所有的子项.
//
// nID: 项ID.
//
// bExpand: 是否展开.
func (t *Tree) X展开全部子项(项ID int32, 是否展开 bool) bool {
	return 炫彩基类.X列表树_展开全部子项(t.Handle, 项ID, 是否展开)
}

// 列表树_测试点击项, 检测坐标点所在项, 返回项ID.
//
// pPt: 坐标点.
func (t *Tree) X测试点击项(坐标点 *炫彩基类.POINT) int32 {
	return 炫彩基类.X列表树_测试点击项(t.Handle, 坐标点)
}

// 列表树_测试点击项扩展, 检测坐标点所在项, 自动添加滚动视图偏移坐标, 返回项ID.
//
// pPt: 坐标点.
func (t *Tree) X测试点击项EX(坐标点 *炫彩基类.POINT) int32 {
	return 炫彩基类.X列表树_测试点击项EX(t.Handle, 坐标点)
}

// 列表树_取第一个子项, 获取第一个子项. 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) X取第一个子项(项ID int32) int32 {
	return 炫彩基类.X列表树_取第一个子项(t.Handle, 项ID)
}

// 列表树_取末尾子项, 获取末尾子项. 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) X取末尾子项(项ID int32) int32 {
	return 炫彩基类.X列表树_取末尾子项(t.Handle, 项ID)
}

// 列表树_取上一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) X取上一个兄弟项(项ID int32) int32 {
	return 炫彩基类.X列表树_取上一个兄弟项(t.Handle, 项ID)
}

// 列表树_取下一个兄弟项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) X取下一个兄弟项(项ID int32) int32 {
	return 炫彩基类.X列表树_取下一个兄弟项(t.Handle, 项ID)
}

// 列表树_取父项, 成功返回项ID, 失败返回XC_ID_ERROR.
//
// nID: 项ID.
func (t *Tree) X取父项(项ID int32) int32 {
	return 炫彩基类.X列表树_取父项(t.Handle, 项ID)
}

// 列表树_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (t *Tree) X创建数据适配器() int {
	return 炫彩基类.X列表树_创建数据适配器(t.Handle)
}

// 列表树_绑定数据适配器.
//
// hAdapter: 数据适配器句柄, XAdTree.
func (t *Tree) X绑定数据适配器(数据适配器句柄 int) {
	炫彩基类.X列表树_绑定数据适配器(t.Handle, 数据适配器句柄)
}

// 列表树_取数据适配器, 返回数据适配器句柄.
func (t *Tree) X取数据适配器() int {
	return 炫彩基类.X列表树_取数据适配器(t.Handle)
}

// 列表树_刷新数据, 刷新所有项模板, 以便更新UI.
func (t *Tree) X刷新数据() {
	炫彩基类.X列表树_刷新数据(t.Handle)
}

// 列表树_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// nID: 项ID.
func (t *Tree) X刷新指定项(项ID int32) {
	炫彩基类.X列表树_刷新指定项(t.Handle, 项ID)
}

// 列表树_置缩进, 设置缩进大小.
//
// nWidth: 缩进宽度.
func (t *Tree) X置缩进(缩进宽度 int32) {
	炫彩基类.X列表树_置缩进(t.Handle, 缩进宽度)
}

// 列表树_取缩进, 返回缩进值大小.
func (t *Tree) X取缩进() int32 {
	return 炫彩基类.X列表树_取缩进(t.Handle)
}

// 列表树_置项默认高度.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (t *Tree) X置项默认高度(高度, 选中时高度 int32) {
	炫彩基类.X列表树_置项默认高度(t.Handle, 高度, 选中时高度)
}

// 列表树_取项默认高度.
//
// pHeight: 接收返回高度.
//
// pSelHeight: 接收返回值, 当项选中时的高度.
func (t *Tree) X取项默认高度(接收返回高度, 接收返回值 *int32) {
	炫彩基类.X列表树_取项默认高度(t.Handle, 接收返回高度, 接收返回值)
}

// 列表树_置项高度.
//
// nID: 项ID.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (t *Tree) X置项高度(项ID, 高度, 选中时高度 int32) {
	炫彩基类.X列表树_置项高度(t.Handle, 项ID, 高度, 选中时高度)
}

// 列表树_取项高度.
//
// nID: 项ID.
//
// pHeight: 接收返回高度.
//
// pSelHeight: 接收返回值, 当项选中时的高度.
func (t *Tree) X取项高度(项ID int32, 接收返回高度, 接收返回值 *int32) {
	炫彩基类.X列表树_取项高度(t.Handle, 项ID, 接收返回高度, 接收返回值)
}

// 列表树_置行间距.
//
// nSpace: 行间隔大小.
func (t *Tree) X置行间距(行间隔大小 int32) {
	炫彩基类.X列表树_置行间距(t.Handle, 行间隔大小)
}

// 列表树_取行间距.
func (t *Tree) X取行间距() int32 {
	return 炫彩基类.X列表树_取行间距(t.Handle)
}

// 列表树_移动项, 移动项的位置.
//
// nMoveItem: 要移动的项ID.
//
// nDestItem: 目标项ID, 参照位置.
//
// nFlag: 0:目标前面, 1:目标后面, 2:目标子项首, 3:目标子项尾.
func (t *Tree) X移动项(要移动的项ID, 目标项ID, 位置 int32) bool {
	return 炫彩基类.X列表树_移动项(t.Handle, 要移动的项ID, 目标项ID, 位置)
}

// 列表树_取模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// nID: 树项ID.
//
// nTempItemID: 模板项ID.
func (t *Tree) X取模板对象(项ID, 模板项ID int32) int {
	return 炫彩基类.X列表树_取模板对象(t.Handle, 项ID, 模板项ID)
}

// 列表树_取对象所在项, 获取当前对象所在模板实例, 属于列表树中哪一个项. 成功返回项ID, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄.
func (t *Tree) X取对象所在项(对象句柄 int) int32 {
	return 炫彩基类.X列表树_取对象所在项(t.Handle, 对象句柄)
}

// 列表树_插入项文本.
//
// pValue:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) X插入项文本(值 string, 父id, 插入id int32) int32 {
	return 炫彩基类.X列表树_插入项文本(t.Handle, 值, 父id, 插入id)
}

// 列表树_插入项文本扩展.
//
// pName:.
//
// pValue:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) X插入项文本EX(名称 string, 值 string, 父id, 插入id int32) int32 {
	return 炫彩基类.X列表树_插入项文本EX(t.Handle, 名称, 值, 父id, 插入id)
}

// 列表树_插入项图片.
//
// hImage:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) X插入项图片(图片 int, 父id, 插入id int32) int32 {
	return 炫彩基类.X列表树_插入项图片(t.Handle, 图片, 父id, 插入id)
}

// 列表树_插入项图片扩展.
//
// pName:.
//
// hImage:.
//
// nParentID:.
//
// insertID:.
func (t *Tree) X插入项图片EX(名称 string, 图片 int, 父id, 插入id int32) int32 {
	return 炫彩基类.X列表树_插入项图片EX(t.Handle, 名称, 图片, 父id, 插入id)
}

// 列表树_取项数量.
func (t *Tree) X取项数量() int32 {
	return 炫彩基类.X列表树_取项数量(t.Handle)
}

// 列表树_取列数量.
func (t *Tree) X取列数量() int32 {
	return 炫彩基类.X列表树_取列数量(t.Handle)
}

// 列表树_置项文本.
//
// nID:.
//
// iColumn:.
//
// pValue:.
func (t *Tree) X置项文本(项ID, 列 int32, 值 string) bool {
	return 炫彩基类.X列表树_置项文本(t.Handle, 项ID, 列, 值)
}

// 列表树_置项文本扩展.
//
// nID:.
//
// pName:.
//
// pValue:.
func (t *Tree) X置项文本EX(项ID int32, 名称 string, 值 string) bool {
	return 炫彩基类.X列表树_置项文本EX(t.Handle, 项ID, 名称, 值)
}

// 列表树_置项图片.
//
// nID:.
//
// iColumn:.
//
// hImage:.
func (t *Tree) X置项图片(项ID, 列 int32, 图片 int) bool {
	return 炫彩基类.X列表树_置项图片(t.Handle, 项ID, 列, 图片)
}

// 列表树_置项图片扩展.
//
// nID:.
//
// pName:.
//
// hImage:.
func (t *Tree) X置项图片EX(项ID int32, 名称 string, 图片 int) bool {
	return 炫彩基类.X列表树_置项图片EX(t.Handle, 项ID, 名称, 图片)
}

// 列表树_取项文本.
//
// nID:.
//
// iColumn:.
func (t *Tree) X取项文本(项ID, 列 int32) string {
	return 炫彩基类.X列表树_取项文本(t.Handle, 项ID, 列)
}

// 列表树_取项文本扩展.
//
// nID:.
//
// pName:.
func (t *Tree) X取项文本EX(项ID int32, 名称 string) string {
	return 炫彩基类.X列表树_取项文本EX(t.Handle, 项ID, 名称)
}

// 列表树_取项图片.
//
// nID:.
//
// iColumn:.
func (t *Tree) X取项图片(项ID, 列 int32) int {
	return 炫彩基类.X列表树_取项图片(t.Handle, 项ID, 列)
}

// 列表树_取项图片扩展.
//
// nID:.
//
// pName:.
func (t *Tree) X取项图片EX(项ID int32, 名称 string) int {
	return 炫彩基类.X列表树_取项图片EX(t.Handle, 项ID, 名称)
}

// 列表树_删除项.
//
// nID:.
func (t *Tree) X删除项(nID int32) bool {
	return 炫彩基类.X列表树_删除项(t.Handle, nID)
}

// 列表树_删除全部项.
func (t *Tree) X删除全部项() {
	炫彩基类.X列表树_删除全部项(t.Handle)
}

// 列表树_删除列全部.
func (t *Tree) X删除列全部() {
	炫彩基类.X列表树_删除列全部(t.Handle)
}

// 列表树_置分割线颜色.
//
// color: ABGR 颜色值.
func (t *Tree) X置分割线颜色(ABGR颜色值 int) {
	炫彩基类.X列表树_置分割线颜色(t.Handle, ABGR颜色值)
}

// 列表树_置项模板从内存.
//
// data: 模板数据.
func (t *Tree) X置项模板从内存(模板数据 []byte) bool {
	return 炫彩基类.X列表树_置项模板从内存(t.Handle, 模板数据)
}

// 列表树_置项模板从资源ZIP.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (t *Tree) X置项模板并按资源ZIP(RC资源ID int32, 文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	return 炫彩基类.X列表树_置项模板从资源ZIP(t.Handle, RC资源ID, 文件名, zip密码, 模块句柄)
}

// 列表树_取项模板, 返回项模板句柄.
func (t *Tree) X取项模板() int {
	return 炫彩基类.X列表树_取项模板(t.Handle)
}

/*
以下都是事件
*/

// 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_TREE_TEMP_CREATE func(pItem *炫彩基类.Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_TREE_TEMP_CREATE1 func(hEle int, pItem *炫彩基类.Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_TREE_TEMP_CREATE_END func(pItem *炫彩基类.Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_TREE_TEMP_CREATE_END1 func(hEle int, pItem *炫彩基类.Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板销毁,模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_TREE_TEMP_DESTROY func(pItem *炫彩基类.Tree_Item_, nFlag int32, pbHandled *bool) int

// 列表树元素-项模板销毁,模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_TREE_TEMP_DESTROY1 func(hEle int, pItem *炫彩基类.Tree_Item_, nFlag int32, pbHandled *bool) int
type XE_TREE_TEMP_ADJUST_COORDINATE func(pItem *炫彩基类.Tree_Item_, pbHandled *bool) int            // 树元素,项模板,调整项坐标. 已停用.
type XE_TREE_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *炫彩基类.Tree_Item_, pbHandled *bool) int // 树元素,项模板,调整项坐标. 已停用.
type XE_TREE_DRAWITEM func(hDraw int, pItem *炫彩基类.Tree_Item_, pbHandled *bool) int               // 树元素,绘制项.
type XE_TREE_DRAWITEM1 func(hEle int, hDraw int, pItem *炫彩基类.Tree_Item_, pbHandled *bool) int    // 树元素,绘制项.
type XE_TREE_SELECT func(nItemID int32, pbHandled *bool) int                                   // 树元素,项选择事件.
type XE_TREE_SELECT1 func(hEle int, nItemID int32, pbHandled *bool) int                        // 树元素,项选择事件.
type XE_TREE_EXPAND func(id int32, bExpand bool, pbHandled *bool) int                          // 树元素,项展开收缩事件.
type XE_TREE_EXPAND1 func(hEle int, id int32, bExpand bool, pbHandled *bool) int               // 树元素,项展开收缩事件.
type XE_TREE_DRAG_ITEM_ING func(pInfo *炫彩基类.Tree_Drag_Item_, pbHandled *bool) int                // 树元素,用户正在拖动项, 可对参数值修改.
type XE_TREE_DRAG_ITEM_ING1 func(hEle int, pInfo *炫彩基类.Tree_Drag_Item_, pbHandled *bool) int     // 树元素,用户正在拖动项, 可对参数值修改.
type XE_TREE_DRAG_ITEM func(pInfo *炫彩基类.Tree_Drag_Item_, pbHandled *bool) int                    // 树元素,拖动项事件.
type XE_TREE_DRAG_ITEM1 func(hEle int, pInfo *炫彩基类.Tree_Drag_Item_, pbHandled *bool) int         // 树元素,拖动项事件.

// 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (t *Tree) X事件_项模板创建(pFun XE_TREE_TEMP_CREATE) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_TEMP_CREATE, pFun)
}

// 列表树元素-项模板创建,模板复用机制需先启用; 替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (t *Tree) X事件_项模板创建1(pFun XE_TREE_TEMP_CREATE1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_TEMP_CREATE, pFun)
}

// 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (t *Tree) X事件_项模板创建完成(pFun XE_TREE_TEMP_CREATE_END) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_TEMP_CREATE_END, pFun)
}

// 列表树元素-项模板创建完成,模板复用机制需先启用; 不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (t *Tree) X事件_项模板创建完成1(pFun XE_TREE_TEMP_CREATE_END1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_TEMP_CREATE_END, pFun)
}

// 列表树元素-项模板销毁,模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (t *Tree) X事件_项模板销毁(pFun XE_TREE_TEMP_DESTROY) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_TEMP_DESTROY, pFun)
}

// 列表树元素-项模板销毁,模板复用机制需先启用;
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (t *Tree) X事件_项模板销毁1(pFun XE_TREE_TEMP_DESTROY1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_TEMP_DESTROY, pFun)
}

// 树元素,项模板,调整项坐标. 已停用.
func (t *Tree) X停用_项模板(pFun XE_TREE_TEMP_ADJUST_COORDINATE) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_TEMP_ADJUST_COORDINATE, pFun)
}

// 树元素,项模板,调整项坐标. 已停用.
func (t *Tree) X停用_项模板1(pFun XE_TREE_TEMP_ADJUST_COORDINATE1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_TEMP_ADJUST_COORDINATE, pFun)
}

// 树元素,绘制项.
func (t *Tree) X事件_绘制项(pFun XE_TREE_DRAWITEM) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_DRAWITEM, pFun)
}

// 树元素,绘制项.
func (t *Tree) X事件_绘制项1(pFun XE_TREE_DRAWITEM1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_DRAWITEM, pFun)
}

// 树元素,项选择事件.
func (t *Tree) X事件_项选择(pFun XE_TREE_SELECT) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_SELECT, pFun)
}

// 树元素,项选择事件.
func (t *Tree) X事件_项选择1(pFun XE_TREE_SELECT1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_SELECT, pFun)
}

// 树元素,项展开收缩事件.
func (t *Tree) X事件_项展开收缩(pFun XE_TREE_EXPAND) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_EXPAND, pFun)
}

// 树元素,项展开收缩事件.
func (t *Tree) X事件_项展开收缩1(pFun XE_TREE_EXPAND1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_EXPAND, pFun)
}

// 树元素,用户正在拖动项, 可对参数值修改.
func (t *Tree) X事件_用户正在拖动项(pFun XE_TREE_DRAG_ITEM_ING) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_DRAG_ITEM_ING, pFun)
}

// 树元素,用户正在拖动项, 可对参数值修改.
func (t *Tree) X事件_用户正在拖动项1(pFun XE_TREE_DRAG_ITEM_ING1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_DRAG_ITEM_ING, pFun)
}

// 树元素,拖动项事件.
func (t *Tree) X事件_拖动项(pFun XE_TREE_DRAG_ITEM) bool {
	return 炫彩基类.X元素_注册事件C(t.Handle, 炫彩常量类.XE_TREE_DRAG_ITEM, pFun)
}

// 树元素,拖动项事件.
func (t *Tree) X事件_拖动项1(pFun XE_TREE_DRAG_ITEM1) bool {
	return 炫彩基类.X元素_注册事件C1(t.Handle, 炫彩常量类.XE_TREE_DRAG_ITEM, pFun)
}
