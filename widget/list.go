package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 列表.
type List struct {
	ScrollView
}

// 列表_创建, 创建列表元素.
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
func X创建列表(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *List {
	p := &List{}
	p.X设置句柄(炫彩基类.X列表_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 列表_创建Ex, 创建列表元素, 使用内置项模板, 自动创建数据适配器.
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
func X创建列表Ex(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄, 列数量 int32) *List {
	p := &List{}
	p.X设置句柄(炫彩基类.X列表_创建Ex(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄, 列数量))
	return p
}

// 从句柄创建对象.
func X创建列表并按句柄(handle int) *List {
	p := &List{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建列表并按名称(name string) *List {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &List{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建列表并按UID(nUID int) *List {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &List{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建列表并按UID名称(name string) *List {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &List{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 列表_增加列, 返回位置索引.
//
// width: 列宽度.
func (l *List) X增加列(列宽度 int) int {
	return 炫彩基类.X列表_增加列(l.Handle, 列宽度)
}

// 列表_插入列, 返回插入位置索引.
//
// width: 列宽度.
//
// iItem: 插入位置索引.
func (l *List) X插入列(列宽度 int, 插入位置索引 int) int {
	return 炫彩基类.X列表_插入列(l.Handle, 列宽度, 插入位置索引)
}

// 列表_启用多选, 启用或关闭多选功能.
//
// bEnable: 是否启用.
func (l *List) X启用多选(是否启用 bool) int {
	return 炫彩基类.X列表_启用多选(l.Handle, 是否启用)
}

// 列表_启用拖动更改列宽, 启用拖动改变列宽度.
//
// bEnable: 是否启用.
func (l *List) X启用拖动更改列宽(是否启用 bool) int {
	return 炫彩基类.X列表_启用拖动更改列宽(l.Handle, 是否启用)
}

// 列表_启用垂直滚动条顶部对齐.
//
// bTop: 是否启用.
func (l *List) X启用垂直滚动条顶部对齐(是否启用 bool) int {
	return 炫彩基类.X列表_启用垂直滚动条顶部对齐(l.Handle, 是否启用)
}

// 列表_启用项背景全行模式, 启用项背景全行填充模式.
//
// bFull: 是否启用.
func (l *List) X启用项背景全行模式(是否启用 bool) int {
	return 炫彩基类.X列表_启用项背景全行模式(l.Handle, 是否启用)
}

// 列表_启用固定行高.
//
// bEnable: 是否启用.
func (l *List) X启用固定行高(是否启用 bool) int {
	return 炫彩基类.X列表_启用固定行高(l.Handle, 是否启用)
}

// 列表_启用模板复用.
//
// bEnable: 是否启用.
func (l *List) X启用模板复用(是否启用 bool) int {
	return 炫彩基类.X列表_启用模板复用(l.Handle, 是否启用)
}

// 列表_启用虚表.
//
// bEnable: 是否启用.
func (l *List) X启用虚表(是否启用 bool) int {
	return 炫彩基类.X列表_启用虚表(l.Handle, 是否启用)
}

// 列表_置虚表行数.
//
// nRowCount: 行数.
func (l *List) X置虚表行数(行数 int) int {
	return 炫彩基类.X列表_置虚表行数(l.Handle, 行数)
}

// 列表_置排序, 设置排序属性.
//
// iColumn: 列索引.
//
// iColumnAdapter: 需要排序的数据在数据适配器中的列索引.
//
// bEnable: 是否启用排序功能.
func (l *List) X置排序(列索引 int, 数据适配器中列索引 int, 是否启用排序 bool) int {
	return 炫彩基类.X列表_置排序(l.Handle, 列索引, 数据适配器中列索引, 是否启用排序)
}

// 列表_置绘制项背景标志, 设置是否绘制指定状态下项的背景.
//
// nFlags: 标志位, List_DrawItemBk_Flag_.
func (l *List) X置绘制项背景标志(标志位 炫彩常量类.List_DrawItemBk_Flag_) int {
	return 炫彩基类.X列表_置绘制项背景标志(l.Handle, 标志位)
}

// 列表_置列宽.
//
// iItem: 列索引.
//
// width: 宽度.
func (l *List) X置列宽(列索引 int, 宽度 int) int {
	return 炫彩基类.X列表_置列宽(l.Handle, 列索引, 宽度)
}

// 列表_置列最小宽度.
//
// iItem: 列索引.
//
// width: 宽度.
func (l *List) X置列最小宽度(列索引 int, 宽度 int) int {
	return 炫彩基类.X列表_置列最小宽度(l.Handle, 列索引, 宽度)
}

// 列表_置列宽度固定.
//
// iColumn: 列索引.
//
// bFixed: 是否固定宽度.
func (l *List) X置列宽度固定(列索引 int, 是否固定宽度 bool) int {
	return 炫彩基类.X列表_置列宽度固定(l.Handle, 列索引, 是否固定宽度)
}

// 列表_取列宽度.
//
// iColumn: 列索引.
func (l *List) X取列宽度(列索引 int) int {
	return 炫彩基类.X列表_取列宽度(l.Handle, 列索引)
}

// 列表_取列数量.
func (l *List) X取列数量() int {
	return 炫彩基类.X列表_取列数量(l.Handle)
}

// 列表_置项数据, 设置项用户数据.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// data: 用户数据.
func (l *List) X置项数据(项索引 int, 子项索引 int, 用户数据 int) bool {
	return 炫彩基类.X列表_置项数据(l.Handle, 项索引, 子项索引, 用户数据)
}

// 列表_取项数据, 获取项用户数据.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
func (l *List) X取项数据(项索引 int, 子项索引 int) int {
	return 炫彩基类.X列表_取项数据(l.Handle, 项索引, 子项索引)
}

// 列表_置选择项.
//
// iItem: 项索引.
func (l *List) X置选择项(项索引 int) bool {
	return 炫彩基类.X列表_置选择项(l.Handle, 项索引)
}

// 列表_取选择项, 返回项索引.
func (l *List) X取选择项() int {
	return 炫彩基类.X列表_取选择项(l.Handle)
}

// 列表_取选择项数量, 获取选择项数量.
func (l *List) X取选择项数量() int {
	return 炫彩基类.X列表_取选择项数量(l.Handle)
}

// 列表_添加选择项.
//
// iItem: 项索引.
func (l *List) X添加选择项(项索引 int) bool {
	return 炫彩基类.X列表_添加选择项(l.Handle, 项索引)
}

// 列表_置选择全部, 选择全部行.
func (l *List) X置选择全部() int {
	return 炫彩基类.X列表_置选择全部(l.Handle)
}

// 列表_取全部选择, 获取全部选择的行, 返回行数量.
//
// pArray: 接收行索引切片.
//
// nArraySize: 切片大小.
func (l *List) X取全部选择(接收行索引切片 *[]int32, 切片大小 int) int {
	return 炫彩基类.X列表_取全部选择(l.Handle, 接收行索引切片, 切片大小)
}

// 列表_显示指定项, 滚动视图让指定项可见.
//
// iItem: 项索引.
func (l *List) X显示指定项(项索引 int) int {
	return 炫彩基类.X列表_显示指定项(l.Handle, 项索引)
}

// 列表_取消选择项, 取消选择指定项(这里的项可以理解为行).
//
// iItem: 项索引.
func (l *List) X取消选择项(项索引 int) bool {
	return 炫彩基类.X列表_取消选择项(l.Handle, 项索引)
}

// 列表_取消全部选择项, 取消选择所有项(这里的项可以理解为行).
func (l *List) X取消全部选择项() int {
	return 炫彩基类.X列表_取消全部选择项(l.Handle)
}

// 列表_取列表头, 获取列表头元素, 返回列表头元素句柄.
func (l *List) X取列表头() int {
	return 炫彩基类.X列表_取列表头(l.Handle)
}

// 列表_删除列.
//
// iItem: 项索引.
func (l *List) X删除列(项索引 int) bool {
	return 炫彩基类.X列表_删除列(l.Handle, 项索引)
}

// 列表_删除全部列, 删除所有的列, 仅删除List的, 数据适配器的列不变.
func (l *List) X删除全部列() int {
	return 炫彩基类.X列表_删除全部列(l.Handle)
}

// 列表_绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdTable.
func (l *List) X绑定数据适配器(数据适配器句柄 int) int {
	return 炫彩基类.X列表_绑定数据适配器(l.Handle, 数据适配器句柄)
}

// 列表_列表头绑定数据适配器.
//
// hAdapter: 数据适配器句柄 XAdMap.
func (l *List) X列表头绑定数据适配器(数据适配器句柄 int) int {
	return 炫彩基类.X列表_列表头绑定数据适配器(l.Handle, 数据适配器句柄)
}

// 列表_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
//
// colExtend_count:	列延伸-预计列表总列数, 默认值0; 限制最大延伸范围, 避免超出范围, 增加不必要的字段.
func (l *List) X创建数据适配器(预计列表总列数 int) int {
	return 炫彩基类.X列表_创建数据适配器(l.Handle, 预计列表总列数)
}

// 列表_列表头创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *List) X创建列表头数据适配器() int {
	return 炫彩基类.X列表_列表头创建数据适配器(l.Handle)
}

// 列表_取数据适配器, 返回数据适配器句柄.
func (l *List) X取数据适配器() int {
	return 炫彩基类.X列表_取数据适配器(l.Handle)
}

// 列表_列表头获取数据适配器, 获取列表头数据适配器句柄.
func (l *List) X取列表头数据适配器() int {
	return 炫彩基类.X列表_列表头获取数据适配器(l.Handle)
}

// 列表_置项模板文件, 设置项布局模板文件.
//
// pXmlFile: 文件名.
func (l *List) X置项模板文件(文件名 string) bool {
	return 炫彩基类.X列表_置项模板文件(l.Handle, 文件名)
}

// 列表_置项模板从字符串, 设置项布局模板文件.
//
// pStringXML: 字符串.
func (l *List) X置项模板从字符串(字符串 string) bool {
	return 炫彩基类.X列表_置项模板从字符串(l.Handle, 字符串)
}

// 列表_置项模板, 设置列表项模板.
//
// hTemp: 模板句柄.
func (l *List) X置项模板(模板句柄 int) bool {
	return 炫彩基类.X列表_置项模板(l.Handle, 模板句柄)
}

// 列表_取项模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iItem: 项索引.
//
// iSubItem: 子项索引.
//
// nTempItemID: 模板项itemID.
func (l *List) X取项模板对象(项索引 int, 子项索引 int, 模板项itemID int) int {
	return 炫彩基类.X列表_取项模板对象(l.Handle, 项索引, 子项索引, 模板项itemID)
}

// 列表_取对象所在行, 获取当前对象所在模板实例, 属于列表中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄.
func (l *List) X取对象所在行(对象句柄 int) int {
	return 炫彩基类.X列表_取对象所在行(l.Handle, 对象句柄)
}

// 列表_取列表头模板对象, 列表头, 通过模板项ID, 获取实例化模板项ID对应的对象句柄.
//
// iItem: 列表头项ID.
//
// nTempItemID: 模板项ID.
func (l *List) X取列表头模板对象(列表头项ID int, 模板项ID int) int {
	return 炫彩基类.X列表_取列表头模板对象(l.Handle, 列表头项ID, 模板项ID)
}

// 列表_取列表头对象所在行, 列表头, 获取当前对象所在模板实例, 属于列表头中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
//
// hXCGUI: 对象句柄.
func (l *List) X取列表头对象所在行(对象句柄 int) int {
	return 炫彩基类.X列表_取列表头对象所在行(l.Handle, 对象句柄)
}

// 列表_置列表头高度.
//
// height: 高度.
func (l *List) X置列表头高度(高度 int) int {
	return 炫彩基类.X列表_置列表头高度(l.Handle, 高度)
}

// 列表_取列表头高度.
func (l *List) X取列表头高度() int {
	return 炫彩基类.X列表_取列表头高度(l.Handle)
}

// 列表_取可视行范围.
//
// piStart: 开始行索引.
//
// piEnd: 结束行索引.
func (l *List) X取可视行范围(开始行索引 *int32, 结束行索引 *int32) int {
	return 炫彩基类.X列表_取可视行范围(l.Handle, 开始行索引, 结束行索引)
}

// 列表_置项默认高度.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (l *List) X置项默认高度(高度 int32, 选中时高度 int32) int {
	return 炫彩基类.X列表_置项默认高度(l.Handle, 高度, 选中时高度)
}

// 列表_取项默认高度.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *List) X取项默认高度(高度 *int32, 选中时高度 *int32) int {
	return 炫彩基类.X列表_取项默认高度(l.Handle, 高度, 选中时高度)
}

// 列表_置行间距.
//
// nSpace: 行间距大小.
func (l *List) X置行间距(行间距大小 int) int {
	return 炫彩基类.X列表_置行间距(l.Handle, 行间距大小)
}

// 列表_取行间距.
func (l *List) X取行间距() int {
	return 炫彩基类.X列表_取行间距(l.Handle)
}

// 列表_置锁定列左侧, 锁定列, 设置左侧锁定列分界列索引.
//
// iColumn: 列索引, -1代表不锁定.
func (l *List) X置锁定列左侧(列索引 int) int {
	return 炫彩基类.X列表_置锁定列左侧(l.Handle, 列索引)
}

// 列表_置锁定列右侧.
//
// iColumn: 列索引, -1代表不锁定. 暂时只支持锁定末尾列.
func (l *List) X置锁定列右侧(列索引 int) int {
	return 炫彩基类.X列表_置锁定列右侧(l.Handle, 列索引)
}

// 列表_置锁定行底部, 设置是否锁定末尾行.
//
// bLock: 是否锁定.
func (l *List) X置锁定行底部(是否锁定 bool) int {
	return 炫彩基类.X列表_置锁定行底部(l.Handle, 是否锁定)
}

// 列表_置锁定行底部重叠.
//
// bOverlap: 是否重叠.
func (l *List) X置锁定行底部重叠(是否重叠 bool) int {
	return 炫彩基类.X列表_置锁定行底部重叠(l.Handle, 是否重叠)
}

// 列表_测试点击项, 检测坐标点所在项.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func (l *List) X测试点击项(坐标点 *炫彩基类.POINT, 项索引 *int32, 子项索引 *int32) bool {
	return 炫彩基类.X列表_测试点击项(l.Handle, 坐标点, 项索引, 子项索引)
}

// 列表_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量.
//
// pPt: 坐标点.
//
// piItem: 项索引.
//
// piSubItem: 子项索引.
func (l *List) X测试点击项EX(坐标点 *炫彩基类.POINT, 项索引 *int32, 子项索引 *int32) bool {
	return 炫彩基类.X列表_测试点击项EX(l.Handle, 坐标点, 项索引, 子项索引)
}

// 列表_刷新项数据.
func (l *List) X刷新项数据() int {
	return 炫彩基类.X列表_刷新项数据(l.Handle)
}

// 列表_刷新指定项, 刷新指定项模板, 以便更新UI.
//
// iItem: 项索引.
func (l *List) X刷新指定项(项索引 int) int {
	return 炫彩基类.X列表_刷新指定项(l.Handle, 项索引)
}

// 列表_添加列文本.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen. 你可以理解为创建表头数据适配器后, 内部有一个Map来存储每一列的表头名(列名), 这个name名就是Map的Key, 这个函数就相当于给每一列的Key赋值, 然后List会根据这个name名从Map读取Value来显示表头到界面.
//
// pText: 文本.
func (l *List) X添加列文本(列宽 int, 名称 string, 文本 string) int {
	return 炫彩基类.X列表_添加列文本(l.Handle, 列宽, 名称, 文本)
}

// 列表_添加列图片.
//
// nWidth: 列宽.
//
// pName: 模板里绑定的name名. 在List内部存在有默认模板, name名是从name1到namen. 你可以理解为创建表头数据适配器后, 内部有一个Map来存储每一列的表头名(列名), 这个name名就是Map的Key, 这个函数就相当于给每一列的Key赋值, 然后List会根据这个name名从Map读取Value来显示表头到界面.
//
// hImage: 图片句柄.
func (l *List) X添加列图片(列宽 int, 名称 string, 图片句柄 int) int {
	return 炫彩基类.X列表_添加列图片(l.Handle, 列宽, 名称, 图片句柄)
}

// 列表_添加项文本.
//
// pText:.
func (l *List) X添加项文本(文本 string) int {
	return 炫彩基类.X列表_添加项文本(l.Handle, 文本)
}

// 列表_添加项文本扩展.
//
// pName:.
//
// pText:.
func (l *List) X添加项文本EX(名称 string, 文本 string) int {
	return 炫彩基类.X列表_添加项文本EX(l.Handle, 名称, 文本)
}

// 列表_添加项图片.
//
// hImage:.
func (l *List) X添加项图片(图片 int) int {
	return 炫彩基类.X列表_添加项图片(l.Handle, 图片)
}

// 列表_添加项图片扩展.
//
// pName:.
//
// hImage:.
func (l *List) X添加项图片EX(名称 string, 图片 int) int {
	return 炫彩基类.X列表_添加项图片EX(l.Handle, 名称, 图片)
}

// 列表_插入项文本.
//
// iItem:.
//
// pValue:.
func (l *List) X插入项文本(项 int, 文本 string) int {
	return 炫彩基类.X列表_插入项文本(l.Handle, 项, 文本)
}

// 列表_插入项文本扩展.
//
// iItem:.
//
// pName:.
//
// pValue:.
func (l *List) X插入项文本EX(项 int, 名称 string, 文本 string) int {
	return 炫彩基类.X列表_插入项文本EX(l.Handle, 项, 名称, 文本)
}

// 列表_插入项图片.
//
// iItem:.
//
// hImage:.
func (l *List) X插入项图片(项 int, 图片 int) int {
	return 炫彩基类.X列表_插入项图片(l.Handle, 项, 图片)
}

// 列表_插入项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *List) X插入项图片EX(项 int, 名称 string, 图片 int) int {
	return 炫彩基类.X列表_插入项图片EX(l.Handle, 项, 名称, 图片)
}

// 列表_置项文本.
//
// iItem:.
//
// iColumn:.
//
// pText:.
func (l *List) X置项文本(项 int, 列 int, 文本 string) bool {
	return 炫彩基类.X列表_置项文本(l.Handle, 项, 列, 文本)
}

// 列表_置项文本扩展.
//
// iItem:.
//
// pName:.
//
// pText:.
func (l *List) X置项文本EX(项 int, 名称 string, 文本 string) bool {
	return 炫彩基类.X列表_置项文本EX(l.Handle, 项, 名称, 文本)
}

// 列表_置项图片.
//
// iItem:.
//
// iColumn:.
//
// hImage:.
func (l *List) X置项图片(项 int, 列 int, 图片 int) bool {
	return 炫彩基类.X列表_置项图片(l.Handle, 项, 列, 图片)
}

// 列表_置项图片扩展.
//
// iItem:.
//
// pName:.
//
// hImage:.
func (l *List) X置项图片EX(项 int, 名称 string, 图片 int) bool {
	return 炫彩基类.X列表_置项图片EX(l.Handle, 项, 名称, 图片)
}

// 列表_置项指数值.
//
// iItem:.
//
// iColumn:.
//
// nValue:.
func (l *List) X置项指数值(项 int, 列 int, 值 int) bool {
	return 炫彩基类.X列表_置项指数值(l.Handle, 项, 列, 值)
}

// 列表_置项整数值扩展.
//
// iItem:.
//
// pName:.
//
// nValue:.
func (l *List) X置项整数值EX(项 int, 名称 string, 值 int) bool {
	return 炫彩基类.X列表_置项整数值EX(l.Handle, 项, 名称, 值)
}

// 列表_置项浮点值.
//
// iItem:.
//
// iColumn:.
//
// fFloat:.
func (l *List) X置项浮点值(项 int, 列 int, 值 float32) bool {
	return 炫彩基类.X列表_置项浮点值(l.Handle, 项, 列, 值)
}

// 列表_置项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// fFloat:.
func (l *List) X置项浮点值EX(项 int, 名称 string, 值 float32) bool {
	return 炫彩基类.X列表_置项浮点值EX(l.Handle, 项, 名称, 值)
}

// 列表_取项文本.
//
// iItem:.
//
// iColumn:.
func (l *List) X取项文本(项 int, 列 int) string {
	return 炫彩基类.X列表_取项文本(l.Handle, 项, 列)
}

// 列表_取项文本扩展.
//
// iItem:.
//
// pName:.
func (l *List) X取项文本EX(项 int, 名称 string) string {
	return 炫彩基类.X列表_取项文本EX(l.Handle, 项, 名称)
}

// 列表_取项图片.
//
// iItem:.
//
// iColumn:.
func (l *List) X取项图片(项 int, 列 int) int {
	return 炫彩基类.X列表_取项图片(l.Handle, 项, 列)
}

// 列表_取项图片扩展.
//
// iItem:.
//
// pName:.
func (l *List) X取项图片EX(项 int, 名称 string) int {
	return 炫彩基类.X列表_取项图片EX(l.Handle, 项, 名称)
}

// 列表_取项整数值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *List) X取项整数值(项 int, 列 int, 值指针 *int32) bool {
	return 炫彩基类.X列表_取项整数值(l.Handle, 项, 列, 值指针)
}

// 列表_取项整数值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *List) X取项整数值EX(项 int, 文本 string, 值指针 *int32) bool {
	return 炫彩基类.X列表_取项整数值EX(l.Handle, 项, 文本, 值指针)
}

// 列表_取项浮点值.
//
// iItem:.
//
// iColumn:.
//
// pOutValue:.
func (l *List) X取项浮点值(项 int, 列 int, 值指针 *float32) bool {
	return 炫彩基类.X列表_取项浮点值(l.Handle, 项, 列, 值指针)
}

// 列表_取项浮点值扩展.
//
// iItem:.
//
// pName:.
//
// pOutValue:.
func (l *List) X取项浮点值EX(项 int, 名称 string, 值指针 *float32) bool {
	return 炫彩基类.X列表_取项浮点值EX(l.Handle, 项, 名称, 值指针)
}

// 列表_删除项.
//
// iItem:.
func (l *List) X删除项(项 int) bool {
	return 炫彩基类.X列表_删除项(l.Handle, 项)
}

// 列表_删除项扩展.
//
// iItem:.
//
// nCount:.
func (l *List) X删除项EX(项 int, 计数 int) bool {
	return 炫彩基类.X列表_删除项EX(l.Handle, 项, 计数)
}

// 列表_删除项全部.
func (l *List) X删除项全部() int {
	return 炫彩基类.X列表_删除项全部(l.Handle)
}

// 列表_删除列全部AD.
func (l *List) X删除列全部AD() int {
	return 炫彩基类.X列表_删除列全部AD(l.Handle)
}

// 列表_取项数量AD.
func (l *List) X取项数量AD() int {
	return 炫彩基类.X列表_取项数量AD(l.Handle)
}

// 列表_取列数量AD.
func (l *List) X取列数量AD() int {
	return 炫彩基类.X列表_取列数量AD(l.Handle)
}

// 列表_置分割线颜色.
//
// color: ABGR 颜色值.
func (l *List) X置分割线颜色(ABGR颜色值 int) int {
	return 炫彩基类.X列表_置分割线颜色(l.Handle, ABGR颜色值)
}

// 列表_置项高度.
//
// iRow: 行索引.
//
// nHeight: 高度.
//
// nSelHeight: 选中时高度.
func (l *List) X置项高度(行索引 int, 高度, 选中时高度 int32) int {
	return 炫彩基类.X列表_置项高度(l.Handle, 行索引, 高度, 选中时高度)
}

// 列表_取项高度.
//
// iRow: 行索引.
//
// pHeight: 高度.
//
// pSelHeight: 选中时高度.
func (l *List) X取项高度(行索引 int, 高度, 选中时高度 *int32) int {
	return 炫彩基类.X列表_取项高度(l.Handle, 行索引, 高度, 选中时高度)
}

// 列表_置拖动矩形颜色.
//
// color: ABGR 颜色值.
//
// width: 线宽度.
func (l *List) X置拖动矩形颜色(ABGR颜色值, 线宽度 int) int {
	return 炫彩基类.X列表_置拖动矩形颜色(l.Handle, ABGR颜色值, 线宽度)
}

// 列表_取项模板. 返回列表项模板句柄.
func (l *List) X取项模板() int {
	return 炫彩基类.X列表_取项模板(l.Handle)
}

// 列表_取项模板列表头. 返回列表头项模板句柄.
func (l *List) X取项模板列表头() int {
	return 炫彩基类.X列表_取项模板列表头(l.Handle)
}

// 列表_刷新项数据列表头.
func (l *List) X刷新项数据列表头() int {
	return 炫彩基类.X列表_刷新项数据列表头(l.Handle)
}

// 列表_置项模板从内存.
//
// data: 模板数据.
func (l *List) X置项模板从内存(模板数据 []byte) bool {
	return 炫彩基类.X列表_置项模板从内存(l.Handle, 模板数据)
}

// 列表_置项模板从资源ZIP. 从RC资源ZIP加载.
//
// id: RC资源ID.
//
// pFileName: 项模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func (l *List) X置项模板并按资源ZIP(RC资源ID int32, 项模板文件名 string, zip密码 string, 模块句柄 uintptr) bool {
	return 炫彩基类.X列表_置项模板从资源ZIP(l.Handle, RC资源ID, 项模板文件名, zip密码, 模块句柄)
}

/*
以下都是事件
*/

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE func(pItem *炫彩基类.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE1 func(hEle int, pItem *炫彩基类.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE_END func(pItem *炫彩基类.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
type XE_LIST_TEMP_CREATE_END1 func(hEle int, pItem *炫彩基类.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LIST_TEMP_DESTROY func(pItem *炫彩基类.List_Item_, nFlag int32, pbHandled *bool) int

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
type XE_LIST_TEMP_DESTROY1 func(hEle int, pItem *炫彩基类.List_Item_, nFlag int32, pbHandled *bool) int
type XE_LIST_TEMP_ADJUST_COORDINATE func(pItem *炫彩基类.List_Item_, pbHandled *bool) int                          // 列表元素,项模板调整坐标. 已停用.
type XE_LIST_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *炫彩基类.List_Item_, pbHandled *bool) int               // 列表元素,项模板调整坐标. 已停用.
type XE_LIST_DRAWITEM func(hDraw int, pItem *炫彩基类.List_Item_, pbHandled *bool) int                             // 列表元素,绘制项.
type XE_LIST_DRAWITEM1 func(hEle int, hDraw int, pItem *炫彩基类.List_Item_, pbHandled *bool) int                  // 列表元素,绘制项.
type XE_LIST_SELECT func(iItem int32, pbHandled *bool) int                                                     // 列表元素,项选择事件.
type XE_LIST_SELECT1 func(hEle int, iItem int32, pbHandled *bool) int                                          // 列表元素,项选择事件.
type XE_LIST_HEADER_DRAWITEM func(hDraw int, pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int               // 列表元素绘制列表头项.
type XE_LIST_HEADER_DRAWITEM1 func(hEle int, hDraw int, pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int    // 列表元素绘制列表头项.
type XE_LIST_HEADER_CLICK func(iItem int32, pbHandled *bool) int                                               // 列表元素,列表头项点击事件.
type XE_LIST_HEADER_CLICK1 func(hEle int, iItem int32, pbHandled *bool) int                                    // 列表元素,列表头项点击事件.
type XE_LIST_HEADER_WIDTH_CHANGE func(iItem int32, nWidth int32, pbHandled *bool) int                          // 列表元素,列表头项宽度改变事件.
type XE_LIST_HEADER_WIDTH_CHANGE1 func(hEle int, iItem int32, nWidth int32, pbHandled *bool) int               // 列表元素,列表头项宽度改变事件.
type XE_LIST_HEADER_TEMP_CREATE func(pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int                       // 列表元素,列表头项模板创建.
type XE_LIST_HEADER_TEMP_CREATE1 func(hEle int, pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int            // 列表元素,列表头项模板创建.
type XE_LIST_HEADER_TEMP_CREATE_END func(pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int                   // 列表元素,列表头项模板创建完成事件.
type XE_LIST_HEADER_TEMP_CREATE_END1 func(hEle int, pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int        // 列表元素,列表头项模板创建完成事件.
type XE_LIST_HEADER_TEMP_DESTROY func(pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int                      // 列表元素,列表头项模板销毁.
type XE_LIST_HEADER_TEMP_DESTROY1 func(hEle int, pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int           // 列表元素,列表头项模板销毁.
type XE_LIST_HEADER_TEMP_ADJUST_COORDINATE func(pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int            // 列表元素,列表头项模板调整坐标. 已停用.
type XE_LIST_HEADER_TEMP_ADJUST_COORDINATE1 func(hEle int, pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int // 列表元素,列表头项模板调整坐标. 已停用.

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (l *List) X事件_项模板创建(pFun XE_LIST_TEMP_CREATE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_TEMP_CREATE, pFun)
}

// 列表元素-项模板创建事件,模板复用机制需先启用;替换模板无效判断nFlag,因为内部会检查模板是否改变,不用担心重复.
//
// nFlag  0:状态改变; 1:新模板实例; 2:旧模板复用
func (l *List) X事件_项模板创建1(pFun XE_LIST_TEMP_CREATE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_TEMP_CREATE, pFun)
}

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (l *List) X事件_项模板创建完成(pFun XE_LIST_TEMP_CREATE_END) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_TEMP_CREATE_END, pFun)
}

// 列表元素-项模板创建完成事件,模板复用机制需先启用;不管是新建还是复用,都需要更新数据, 当为复用时不要注册事件以免重复注册.
//
// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
func (l *List) X事件_项模板创建完成1(pFun XE_LIST_TEMP_CREATE_END1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_TEMP_CREATE_END, pFun)
}

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (l *List) X事件_项模板销毁(pFun XE_LIST_TEMP_DESTROY) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_TEMP_DESTROY, pFun)
}

// 列表元素,项模板销毁.
//
// nFlag   0:正常销毁;  1:移动到缓存(不会被销毁,临时缓存备用,当需要时被复用)
func (l *List) X事件_项模板销毁1(pFun XE_LIST_TEMP_DESTROY1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_TEMP_DESTROY, pFun)
}

// 列表元素,项模板调整坐标. 已停用.
func (l *List) X停用_项模板调整坐标(pFun XE_LIST_TEMP_ADJUST_COORDINATE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表元素,项模板调整坐标. 已停用.
func (l *List) X停用_项模板调整坐标1(pFun XE_LIST_TEMP_ADJUST_COORDINATE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表元素,绘制项.
func (l *List) X事件_绘制项(pFun XE_LIST_DRAWITEM) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_DRAWITEM, pFun)
}

// 列表元素,绘制项.
func (l *List) X事件_绘制项1(pFun XE_LIST_DRAWITEM1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_DRAWITEM, pFun)
}

// 列表元素,项选择事件.
func (l *List) X事件_项选择(pFun XE_LIST_SELECT) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_SELECT, pFun)
}

// 列表元素,项选择事件.
func (l *List) X事件_项选择1(pFun XE_LIST_SELECT1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_SELECT, pFun)
}

// 列表元素绘制列表头项.
func (l *List) X事件_列表元素绘制列表头项(pFun XE_LIST_HEADER_DRAWITEM) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_HEADER_DRAWITEM, pFun)
}

// 列表元素绘制列表头项.
func (l *List) X事件_列表元素绘制列表头项1(pFun XE_LIST_HEADER_DRAWITEM1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_HEADER_DRAWITEM, pFun)
}

// 列表元素,列表头项点击事件.
func (l *List) X事件_列表头项点击(pFun XE_LIST_HEADER_CLICK) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_HEADER_CLICK, pFun)
}

// 列表元素,列表头项点击事件.
func (l *List) X事件_列表头项点击1(pFun XE_LIST_HEADER_CLICK1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_HEADER_CLICK, pFun)
}

// 列表元素,列表头项宽度改变事件.
func (l *List) X事件_列表头项宽度改变(pFun XE_LIST_HEADER_WIDTH_CHANGE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_HEADER_WIDTH_CHANGE, pFun)
}

// 列表元素,列表头项宽度改变事件.
func (l *List) X事件_列表头项宽度改变1(pFun XE_LIST_HEADER_WIDTH_CHANGE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_HEADER_WIDTH_CHANGE, pFun)
}

// 列表元素,列表头项模板创建.
func (l *List) X事件_列表头项模板创建(pFun XE_LIST_HEADER_TEMP_CREATE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_HEADER_TEMP_CREATE, pFun)
}

// 列表元素,列表头项模板创建.
func (l *List) X事件_列表头项模板创建1(pFun XE_LIST_HEADER_TEMP_CREATE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_HEADER_TEMP_CREATE, pFun)
}

// 列表元素,列表头项模板创建完成事件.
func (l *List) X事件_列表头项模板创建完成(pFun XE_LIST_HEADER_TEMP_CREATE_END) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_HEADER_TEMP_CREATE_END, pFun)
}

// 列表元素,列表头项模板创建完成事件.
func (l *List) X事件_列表头项模板创建完成1(pFun XE_LIST_HEADER_TEMP_CREATE_END1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_HEADER_TEMP_CREATE_END, pFun)
}

// 列表元素,列表头项模板销毁.
func (l *List) X事件_列表头项模板销毁(pFun XE_LIST_HEADER_TEMP_DESTROY) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_HEADER_TEMP_DESTROY, pFun)
}

// 列表元素,列表头项模板销毁.
func (l *List) X事件_列表头项模板销毁1(pFun XE_LIST_HEADER_TEMP_DESTROY1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_HEADER_TEMP_DESTROY, pFun)
}

// 列表元素,列表头项模板调整坐标. 已停用.
func (l *List) X停用_列表头项模板调整坐标(pFun XE_LIST_HEADER_TEMP_ADJUST_COORDINATE) bool {
	return 炫彩基类.X元素_注册事件C(l.Handle, 炫彩常量类.XE_LIST_HEADER_TEMP_ADJUST_COORDINATE, pFun)
}

// 列表元素,列表头项模板调整坐标. 已停用.
func (l *List) X停用_列表头项模板调整坐标1(pFun XE_LIST_HEADER_TEMP_ADJUST_COORDINATE1) bool {
	return 炫彩基类.X元素_注册事件C1(l.Handle, 炫彩常量类.XE_LIST_HEADER_TEMP_ADJUST_COORDINATE, pFun)
}
