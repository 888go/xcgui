package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 编辑框(常规, 富文本, 聊天气泡).
type Edit struct {
	ScrollView
}

// 编辑框_创建.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func X创建编辑框(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *Edit {
	p := &Edit{}
	p.X设置句柄(炫彩基类.X编辑框_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 编辑框_创建扩展.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// nType: 类型, Edit_Type_.
//
// hParent: 父为窗口句柄或元素句柄.
func X创建编辑框EX(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 类型 炫彩常量类.Edit_Type_, 父窗口句柄或元素句柄 int) *Edit {
	p := &Edit{}
	p.Handle = 炫彩基类.X编辑框_创建EX(元素x坐标, 元素y坐标, 宽度, 高度, 类型, 父窗口句柄或元素句柄)
	return p
}

// 从句柄创建对象.
func X创建编辑框并按句柄(句柄 int) *Edit {
	p := &Edit{}
	p.X设置句柄(句柄)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建编辑框并按名称(名称 string) *Edit {
	handle := 炫彩基类.X取对象从名称(名称)
	if handle > 0 {
		p := &Edit{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建编辑框并按UID(nUID int) *Edit {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Edit{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建编辑框并按UID名称(UID名称 string) *Edit {
	handle := 炫彩基类.X取对象从UID名称(UID名称)
	if handle > 0 {
		p := &Edit{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 编辑框_启用自动换行.
//
// bEnable: 是否启用.
func (e *Edit) X启用自动换行(是否启用 bool) int {
	return 炫彩基类.X编辑框_启用自动换行(e.Handle, 是否启用)
}

// 编辑框_启用只读.
//
// bEnable: 是否启用.
func (e *Edit) X启用只读(是否启用 bool) int {
	return 炫彩基类.X编辑框_启用只读(e.Handle, 是否启用)
}

// 编辑框_启用多行.
//
// bEnable:.
func (e *Edit) X启用多行(是否启用 bool) int {
	return 炫彩基类.X编辑框_启用多行(e.Handle, 是否启用)
}

// 编辑框_启用密码, 启用密码模式(只支持默认类型编辑框).
//
// bEnable: 是否启用.
func (e *Edit) X启用密码(是否启用 bool) int {
	return 炫彩基类.X编辑框_启用密码(e.Handle, 是否启用)
}

// 编辑框_启用自动选择, 当获得焦点时,自动选择所有内容.
//
// bEnable: 是否启用.
func (e *Edit) X启用自动选择(是否启用 bool) int {
	return 炫彩基类.X编辑框_启用自动选择(e.Handle, 是否启用)
}

// 编辑框_启用自动取消选择, 当失去焦点时自动取消选择.
//
// bEnable: 是否启用.
func (e *Edit) X启用自动取消选择(是否启用 bool) int {
	return 炫彩基类.X编辑框_启用自动取消选择(e.Handle, 是否启用)
}

// 编辑框_是否只读.
func (e *Edit) X是否只读() bool {
	return 炫彩基类.X编辑框_是否只读(e.Handle)
}

// 编辑框_是否多行.
func (e *Edit) X是否多行() bool {
	return 炫彩基类.X编辑框_是否多行(e.Handle)
}

// 编辑框_是否密码.
func (e *Edit) X是否密码() bool {
	return 炫彩基类.X编辑框_是否密码(e.Handle)
}

// 编辑框_是否自动换行.
func (e *Edit) X是否自动换行() bool {
	return 炫彩基类.X编辑框_是否自动换行(e.Handle)
}

// 编辑框_判断为空.
func (e *Edit) X判断为空() bool {
	return 炫彩基类.X编辑框_判断为空(e.Handle)
}

// 编辑框_是否在选择区域.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) X是否在选择区域(行索引 int, 列索引 int) bool {
	return 炫彩基类.X编辑框_是否在选择区域(e.Handle, 行索引, 列索引)
}

// 编辑框_取总行数.
func (e *Edit) X取总行数() int {
	return 炫彩基类.X编辑框_取总行数(e.Handle)
}

// 编辑框_取数据, 包含文本或非文本内容.
func (e *Edit) X取数据() 炫彩基类.Edit_Data_Copy_ {
	return 炫彩基类.X编辑框_取数据(e.Handle)
}

// 编辑框_添加数据.
//
// pData: 数据结构.
//
// styleTable: 样式表.
//
// nStyleCount: 样式数量.
func (e *Edit) X添加数据(数据结构 *炫彩基类.Edit_Data_Copy_, 样式表 []uint16, 样式数量 int) int {
	return 炫彩基类.X编辑框_添加数据(e.Handle, 数据结构, 样式表, 样式数量)
}

// 编辑框_释放数据.
func (e *Edit) X释放数据(pData *炫彩基类.Edit_Data_Copy_) int {
	return 炫彩基类.X编辑框_释放数据(pData)
}

// 编辑框_置默认文本, 当内容为空时, 显示默认文本.
//
// pString: 文本内容.
func (e *Edit) X置默认文本(文本内容 string) int {
	return 炫彩基类.X编辑框_置默认文本(e.Handle, 文本内容)
}

// 编辑框_置默认文本颜色.
//
// color: ABGR 颜色值.
func (e *Edit) X置默认文本颜色(ABGR颜色值 int) int {
	return 炫彩基类.X编辑框_置默认文本颜色(e.Handle, ABGR颜色值)
}

// 编辑框_置密码字符.
//
// ch: 字符.
func (e *Edit) X置密码字符(字符 int) int {
	return 炫彩基类.X编辑框_置密码字符(e.Handle, 字符)
}

// 编辑框_置文本对齐, 单行模式下有效.
//
// align: 对齐方式, Edit_TextAlign_Flag_.
func (e *Edit) X置文本对齐(对齐方式 炫彩常量类.Edit_TextAlign_Flag_) int {
	return 炫彩基类.X编辑框_置文本对齐(e.Handle, 对齐方式)
}

// 编辑框_置TAB空格.
//
// nSpace: 空格数量.
func (e *Edit) X置TAB空格(空格数量 int) int {
	return 炫彩基类.X编辑框_置TAB空格(e.Handle, 空格数量)
}

// 编辑框_置文本.
//
// pString: 字符串.
func (e *Edit) X置文本(字符串 string) int {
	return 炫彩基类.X编辑框_置文本(e.Handle, 字符串)
}

// 编辑框_置文本整数.
//
// nValue: 整数值.
func (e *Edit) X置文本整数(整数值 int) int {
	return 炫彩基类.X编辑框_置文本整数(e.Handle, 整数值)
}

// 编辑框_取文本, 不包含非文本内容, 返回实际接收文本长度.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 内存大小. 例: GetLength()+1 .
func (e *Edit) X取文本(接收文本指针 *string, 内存大小 int) int {
	return 炫彩基类.X编辑框_取文本(e.Handle, 接收文本指针, 内存大小)
}

// 编辑框_取文本Ex, 不包含非文本内容, 返回文本.
func (e *Edit) X取文本Ex() string {
	var s string
	炫彩基类.X编辑框_取文本(e.Handle, &s, 炫彩基类.X编辑框_取内容长度(e.Handle)+1)
	return s
}

// 编辑框_取选择文本Ex, 不包括非文本内容, 返回文本.
func (e *Edit) X取选择文本Ex() string {
	var s string
	炫彩基类.X编辑框_取选择文本(e.Handle, &s, 炫彩基类.X编辑框_取选择文本长度(e.Handle)+1)
	return s
}

// 编辑框_取文本行.
//
// iRow: 行索引.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 接收文本内存块长度. 例: GetLengthRow()+1 .
func (e *Edit) X取文本行(行索引 int, 接收文本指针 *string, 接收文本内存块长度 int) int {
	return 炫彩基类.X编辑框_取文本行(e.Handle, 行索引, 接收文本指针, 接收文本内存块长度)
}

// 编辑框_取文本行Ex, 返回文本.
//
// iRow: 行索引.
func (e *Edit) X取文本行Ex(行索引 int) string {
	var s string
	炫彩基类.X编辑框_取文本行(e.Handle, 行索引, &s, 炫彩基类.X编辑框_取内容长度行(e.Handle, 行索引)+1)
	return s
}

// 编辑框_取内容长度, 包含非文本内容.
func (e *Edit) X取内容长度() int {
	return 炫彩基类.X编辑框_取内容长度(e.Handle)
}

// 编辑框_取内容长度行, 包含非文本内容.
//
// iRow: 行索引.
func (e *Edit) X取内容长度行(行索引 int) int {
	return 炫彩基类.X编辑框_取内容长度行(e.Handle, 行索引)
}

// 编辑框_取字符, 返回指定位置字符.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) X取字符(行索引 int, 列索引 int) int {
	return 炫彩基类.X编辑框_取字符(e.Handle, 行索引, 列索引)
}

// 编辑框_插入文本.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
func (e *Edit) X插入文本(行索引 int, 列索引 int, 字符串 string) int {
	return 炫彩基类.X编辑框_插入文本(e.Handle, 行索引, 列索引, 字符串)
}

// AddTextUser 编辑框_插入文本模拟用户操作, 自动刷新UI, 支持撤销/恢复.
//
//	@param pString 字符串.
//	@return int
func (e *Edit) X插入文本模拟用户操作(字符串 string) int {
	return 炫彩基类.X编辑框_插入文本模拟用户操作(e.Handle, 字符串)
}

// 编辑框_添加文本.
//
// pString: 字符串.
func (e *Edit) X添加文本(字符串 string) int {
	return 炫彩基类.X编辑框_添加文本(e.Handle, 字符串)
}

// 编辑框_添加文本扩展.
//
// pString: 字符串.
//
// iStyle: 样式索引.
func (e *Edit) X添加文本EX(字符串 string, 样式索引 int) int {
	return 炫彩基类.X编辑框_添加文本EX(e.Handle, 字符串, 样式索引)
}

// 编辑框_添加对象, 例如: 字体, 图片, UI对象, 返回样式索引.
//
// hObj: 对象句柄.
func (e *Edit) X添加对象(对象句柄 int) int {
	return 炫彩基类.X编辑框_添加对象(e.Handle, 对象句柄)
}

// 编辑框_添加对象从样式, 当样式为图片时有效.
//
// iStyle: 样式索引.
func (e *Edit) X添加对象并按样式(样式索引 int) int {
	return 炫彩基类.X编辑框_添加对象从样式(e.Handle, 样式索引)
}

// 编辑框_添加样式, 返回样式索引.
//
// hFont_image_Obj: 字体.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) X添加样式(字体 int, 颜色 int, 是否使用颜色 bool) int {
	return 炫彩基类.X编辑框_添加样式(e.Handle, 字体, 颜色, 是否使用颜色)
}

// 编辑框_添加样式扩展, 返回样式索引.
//
// fontName: 字体名称.
//
// fontSize: 字体大小.
//
// fontStyle: 字体样式, FontStyle_.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) X添加样式EX(字体名称 string, 字体大小 int, 字体样式 炫彩常量类.FontStyle_, 颜色 int, 是否使用颜色 bool) int {
	return 炫彩基类.X编辑框_添加样式EX(e.Handle, 字体名称, 字体大小, 字体样式, 颜色, 是否使用颜色)
}

// 编辑框_取样式信息.
//
// iStyle: 样式索引.
//
// info: 返回样式信息.
func (e *Edit) X取样式信息(样式索引 int, 返回样式信息 *炫彩基类.Edit_Style_Info_) bool {
	return 炫彩基类.X编辑框_取样式信息(e.Handle, 样式索引, 返回样式信息)
}

// 编辑框_置当前样式.
//
// iStyle: 样式索引.
func (e *Edit) X置当前样式(样式索引 int) int {
	return 炫彩基类.X编辑框_置当前样式(e.Handle, 样式索引)
}

// 编辑框_置插入符颜色.
//
// color: 颜色.
func (e *Edit) X置插入符颜色(颜色 int) int {
	return 炫彩基类.X编辑框_置插入符颜色(e.Handle, 颜色)
}

// 编辑框_置插入符宽度.
//
// nWidth: 宽度.
func (e *Edit) X置插入符宽度(宽度 int) int {
	return 炫彩基类.X编辑框_置插入符宽度(e.Handle, 宽度)
}

// 编辑框_置选择背景颜色.
//
// color: ABGR 颜色.
func (e *Edit) X置选择背景颜色(ABGR颜色 int) int {
	return 炫彩基类.X编辑框_置选择背景颜色(e.Handle, ABGR颜色)
}

// 编辑框_置默认行高.
//
// nHeight: 行高.
func (e *Edit) X置默认行高(行高 int) int {
	return 炫彩基类.X编辑框_置默认行高(e.Handle, 行高)
}

// 编辑框_置指定行高度, 类型为 Edit_Type_Richedit 支持指定不同行高.
//
// iRow: 行索引.
//
// nHeight: 高度.
func (e *Edit) X置指定行高度(行索引 int, 高度 int) int {
	return 炫彩基类.X编辑框_置指定行高度(e.Handle, 行索引, 高度)
}

// SetCurPos 编辑框_置当前位置.
//
//	@param iRow 行索引.
//	@return int
func (e *Edit) X置当前位置(行索引 int) int {
	return 炫彩基类.X编辑框_置当前位置(e.Handle, 行索引)
}

// 编辑框_取当前位置点, 返回范围位置点.
func (e *Edit) X取当前位置点() int {
	return 炫彩基类.X编辑框_取当前位置点(e.Handle)
}

// 编辑框_取当前行, 返回行索引.
func (e *Edit) X取当前行() int {
	return 炫彩基类.X编辑框_取当前行(e.Handle)
}

// 编辑框_取当前列, 返回列索引.
func (e *Edit) X取当前列() int {
	return 炫彩基类.X编辑框_取当前列(e.Handle)
}

// 编辑框_取坐标点.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pOut: 接收返回坐标点.
func (e *Edit) X取坐标点(行索引 int, 列索引 int, 接收返回坐标点 *炫彩基类.POINT) int {
	return 炫彩基类.X编辑框_取坐标点(e.Handle, 行索引, 列索引, 接收返回坐标点)
}

// 编辑框_自动滚动, 视图自动滚动到当前插入符位置.
func (e *Edit) X自动滚动() bool {
	return 炫彩基类.X编辑框_自动滚动(e.Handle)
}

// 编辑框_自动滚动扩展, 视图自动滚动到指定位置.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) X自动滚动EX(行索引 int, 列索引 int) bool {
	return 炫彩基类.X编辑框_自动滚动EX(e.Handle, 行索引, 列索引)
}

// PosToRowCol 编辑框_转换位置, 转换位置点到行列.
//
//	@param iPos 位置点.
//	@param pInfo 行列.
//	@return int
func (e *Edit) X转换位置(位置点 int, 行列 *炫彩基类.Position_) int {
	return 炫彩基类.X编辑框_转换位置(e.Handle, 位置点, 行列)
}

// 编辑框_选择全部.
func (e *Edit) X选择全部() bool {
	return 炫彩基类.X编辑框_选择全部(e.Handle)
}

// 编辑框_取消选择.
func (e *Edit) X取消选择() bool {
	return 炫彩基类.X编辑框_取消选择(e.Handle)
}

// 编辑框_删除选择内容.
func (e *Edit) X删除选择内容() bool {
	return 炫彩基类.X编辑框_删除选择内容(e.Handle)
}

// 编辑框_置选择.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func (e *Edit) X置选择(起始行索引 int, 起始行列索引 int, 结束行索引 int, 结束行列索引 int) bool {
	return 炫彩基类.X编辑框_置选择(e.Handle, 起始行索引, 起始行列索引, 结束行索引, 结束行列索引)
}

// 编辑框_取选择文本, 不包括非文本内容, 返回接收文本内容实际长度.
//
// pOut: 接收返回文本内容.
//
// nOutLen: 接收内存大小. GetSelectTextLength()+1 .
func (e *Edit) X取选择文本(接收返回文本 *string, 接收内存大小 int) int {
	return 炫彩基类.X编辑框_取选择文本(e.Handle, 接收返回文本, 接收内存大小)
}

// 编辑框_取选择内容范围.
//
// pBegin: 起始位置.
//
// pEnd: 结束位置.
func (e *Edit) X取选择内容范围(起始位置 *炫彩基类.Position_, 结束位置 *炫彩基类.Position_) bool {
	return 炫彩基类.X编辑框_取选择内容范围(e.Handle, 起始位置, 结束位置)
}

// 编辑框_取可视行范围.
//
// piStart: 起始行索引.
//
// piEnd: 结束行索引.
func (e *Edit) X取可视行范围(起始行索引 *int32, 结束行索引 *int32) int {
	return 炫彩基类.X编辑框_取可视行范围(e.Handle, 起始行索引, 结束行索引)
}

// 编辑框_删除, 删除指定范围内容.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func (e *Edit) X删除(起始行索引 int, 起始行列索引 int, 结束行索引 int, 结束行列索引 int) bool {
	return 炫彩基类.X编辑框_删除(e.Handle, 起始行索引, 起始行列索引, 结束行索引, 结束行列索引)
}

// 编辑框_删除行.
//
// iRow: 行索引.
func (e *Edit) X删除行(行索引 int) bool {
	return 炫彩基类.X编辑框_删除行(e.Handle, 行索引)
}

// 编辑框_剪贴板剪切.
func (e *Edit) X剪贴板剪切() bool {
	return 炫彩基类.X编辑框_剪贴板剪切(e.Handle)
}

// 编辑框_剪贴板复制.
func (e *Edit) X剪贴板复制() bool {
	return 炫彩基类.X编辑框_剪贴板复制(e.Handle)
}

// 编辑框_剪贴板粘贴.
func (e *Edit) X剪贴板粘贴() bool {
	return 炫彩基类.X编辑框_剪贴板粘贴(e.Handle)
}

// 编辑框_撤销.
func (e *Edit) X撤销() bool {
	return 炫彩基类.X编辑框_撤销(e.Handle)
}

// 编辑框_恢复.
func (e *Edit) X恢复() bool {
	return 炫彩基类.X编辑框_恢复(e.Handle)
}

// 编辑框_添加气泡开始, 当前行开始.
//
// hImageAvatar: 头像.
//
// hImageBubble: 气泡背景.
//
// nFlag: 标志, Chat_Flag_.
func (e *Edit) X添加气泡开始(头像 int, 气泡背景 int, 标志 炫彩常量类.Chat_Flag_) int {
	return 炫彩基类.X编辑框_添加气泡开始(e.Handle, 头像, 气泡背景, 标志)
}

// 编辑框_添加气泡结束, 当前行结束.
func (e *Edit) X添加气泡结束() int {
	return 炫彩基类.X编辑框_添加气泡结束(e.Handle)
}

// 编辑框_置气泡缩进, 设置聊天气泡内容缩进.
//
// nIndentation: 缩进值.
func (e *Edit) X置气泡缩进(缩进值 int) int {
	return 炫彩基类.X编辑框_置气泡缩进(e.Handle, 缩进值)
}

// 编辑框_置行间隔, 设置行间隔大小, 多行模式有效.
//
// nSpace: 行间隔大小.
func (e *Edit) X置行间隔(行间隔大小 int) int {
	return 炫彩基类.X编辑框_置行间隔(e.Handle, 行间隔大小)
}

// 编辑框_置当前位置扩展.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) X置当前位置EX(行索引, 列索引 int32) int {
	return 炫彩基类.X编辑框_置当前位置EX(e.Handle, 行索引, 列索引)
}

// 编辑框_取当前位置扩展.
//
// iRow: 返回行索引.
//
// iCol: 返回列索引.
func (e *Edit) X取当前位置EX(返回行索引, 返回列索引 *int32) int {
	return 炫彩基类.X编辑框_取当前位置EX(e.Handle, 返回行索引, 返回列索引)
}

// 编辑框_移动到末尾.
func (e *Edit) X移动到末尾() int {
	return 炫彩基类.X编辑框_移动到末尾(e.Handle)
}

// 编辑框_行列到位置, 返回位置点.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) X行列到位置(行索引 int, 列索引 int) int {
	return 炫彩基类.X编辑框_行列到位置(e.Handle, 行索引, 列索引)
}

// 编辑框_置后备字体, 置中文字体. 如果已设置, 当遇到中文字符时使用后备字体, 解决不支持中文的字体的问题
//
// hFont: 字体.
func (e *Edit) X置后备字体(字体 int) int {
	return 炫彩基类.X编辑框_置后备字体(e.Handle, 字体)
}

// 编辑框_释放样式.
//
// iStyle: 样式.
func (e *Edit) X释放样式(样式 int) bool {
	return 炫彩基类.X编辑框_释放样式(e.Handle, 样式)
}

// 编辑框_修改样式.
//
// iStyle: 样式索引.
//
// hFont: 字体句柄.
//
// color: ABGR 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) X修改样式(样式索引 int, 字体句柄 int, ABGR颜色 int, 是否使用颜色 bool) bool {
	return 炫彩基类.X编辑框_修改样式(e.Handle, 样式索引, 字体句柄, ABGR颜色, 是否使用颜色)
}

// 编辑框_置空格大小.
//
// size: 空格大小.
func (e *Edit) X置空格大小(空格大小 int) int {
	return 炫彩基类.X编辑框_置空格大小(e.Handle, 空格大小)
}

// 编辑框_置字符间距.
//
// size: 英文字符间距大小.
//
// sizeZh: 中文字符间距大小.
func (e *Edit) X置字符间距(英文字符间距 int, 中文字符间距 int) int {
	return 炫彩基类.X编辑框_置字符间距(e.Handle, 英文字符间距, 中文字符间距)
}

// 编辑框_取选择文本长度, 不包括非文本内容, 返回文本内容长度.
func (e *Edit) X取选择文本长度() int {
	return 炫彩基类.X编辑框_取选择文本长度(e.Handle)
}

// 编辑框_置选择文本样式.
//
// iStyle: 样式索引.
func (e *Edit) X置选择文本样式(样式索引 int) int {
	return 炫彩基类.X编辑框_置选择文本样式(e.Handle, 样式索引)
}

// 编辑框_取文本_临时, 不包含非文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
func (e *Edit) X取文本Tmp() string {
	return 炫彩基类.X编辑框_取文本Tmp(e.Handle)
}

// 编辑框_取文本行_临时, 获取指定行文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
//
// iRow: 行索引.
func (e *Edit) X取行文本Tmp(行索引 int) string {
	return 炫彩基类.X编辑框_取文本行Tmp(e.Handle, 行索引)
}

// 编辑框_取选择文本, 不包含非文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
func (e *Edit) X取选择文本Tmp() string {
	return 炫彩基类.X编辑框_取选择文本tmp(e.Handle)
}

// 编辑框_插入气泡开始, 当前行开始.
//
// hImageAvatar: 头像图片句柄.
//
// hImageBubble: 气泡背景图片句柄.
//
// nFlag: 聊天气泡对齐方式: xcc.Chat_Flag_ .
func (e *Edit) X插入气泡开始(头像图片句柄 int, 背景图片句柄 int, 对齐方式 炫彩常量类.Chat_Flag_) int {
	return 炫彩基类.X编辑框_插入气泡开始(e.Handle, 头像图片句柄, 背景图片句柄, 对齐方式)
}

// 编辑框_取指定行气泡标识. 返回行标识: xcc.Chat_Flag_
//
// iRow: 行索引.
func (e *Edit) X取指定行气泡标识(行索引 int) 炫彩常量类.Chat_Flag_ {
	return 炫彩基类.X编辑框_取指定行气泡标识(e.Handle, 行索引)
}

// 编辑框_插入文本扩展.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
//
// iStyle: 样式.
func (e *Edit) X插入文本EX(行索引 int, 列索引 int, 字符串 string, 样式 int) int {
	return 炫彩基类.X编辑框_插入文本EX(e.Handle, 行索引, 列索引, 字符串, 样式)
}

// 编辑框_插入对象.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hObj: 对象句柄.
func (e *Edit) X插入对象(行索引 int, 列索引 int, 对象句柄 int) int {
	return 炫彩基类.X编辑框_插入对象(e.Handle, 行索引, 列索引, 对象句柄)
}

// 编辑框_置气泡最大宽度. 当值为0时代表不限制宽度.
//
// nWidth: 最大宽度.
func (e *Edit) X置气泡最大宽度(最大宽度 int32) {
	炫彩基类.X编辑框_置气泡最大宽度(e.Handle, 最大宽度)
}

/*
以下都是事件
*/

type XE_EDIT_SET func(pbHandled *bool) int                                       // 元素事件_编辑框设置.
type XE_EDIT_SET1 func(hEle int, pbHandled *bool) int                            // 元素事件_编辑框设置.
type XE_EDIT_DRAWROW func(hDraw int, iRow int32, pbHandled *bool) int            // 暂未使用.
type XE_EDIT_DRAWROW1 func(hEle int, hDraw int, iRow int32, pbHandled *bool) int // 暂未使用.
type XE_EDIT_CHANGED func(pbHandled *bool) int                                   // 编辑框_内容被改变.
type XE_EDIT_CHANGED1 func(hEle int, pbHandled *bool) int                        // 编辑框_内容被改变.
type XE_EDIT_POS_CHANGED func(iPos int32, pbHandled *bool) int                   // 编辑框_光标位置_被改变.
type XE_EDIT_POS_CHANGED1 func(hEle int, iPos int32, pbHandled *bool) int        // 编辑框_光标位置_被改变.
type XE_EDIT_STYLE_CHANGED func(iStyle int32, pbHandled *bool) int               // 编辑框_样式_被改变.
type XE_EDIT_STYLE_CHANGED1 func(hEle int, iStyle int32, pbHandled *bool) int    // 编辑框_样式_被改变.
type XE_EDIT_ENTER_GET_TABALIGN func(pbHandled *bool) int                        // 回车TAB对齐,返回需要TAB数量.
type XE_EDIT_ENTER_GET_TABALIGN1 func(hEle int, pbHandled *bool) int             // 回车TAB对齐,返回需要TAB数量.
// 编辑框_行_被改变.
//
// iRow: 更改行开始位置索引,  if(nChangeRows>0) iEnd= iRow + nChangeRows
//
// nChangeRows: 改变行数, 正数添加行, 负数删除行
type XE_EDIT_ROW_CHANGED func(iRow int32, nChangeRows int32, pbHandled *bool) int

// 编辑框_行_被改变.
//
// iRow: 更改行开始位置索引,  if(nChangeRows>0) iEnd= iRow + nChangeRows
//
// nChangeRows: 改变行数, 正数添加行, 负数删除行
type XE_EDIT_ROW_CHANGED1 func(hEle int, iRow int32, nChangeRows int32, pbHandled *bool) int
type XE_EDIT_SWAPROW func(iRow, bArrowUp int32, pbHandled *bool) int            // 元素事件_交换行
type XE_EDIT_SWAPROW1 func(hEle int, iRow, bArrowUp int32, pbHandled *bool) int // 元素事件_交换行
type XE_EDIT_COLOR_CHANGE func(color int, pbHandled *bool) int                  // 编辑框_颜色被改变
type XE_EDIT_COLOR_CHANGE1 func(hEle int, color int, pbHandled *bool) int       // 编辑框_颜色被改变

// 编辑框_颜色被改变.
func (e *Edit) X事件_颜色被改变(pFun XE_EDIT_COLOR_CHANGE) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_COLOR_CHANGE, pFun)
}

// 编辑框_颜色被改变.
func (e *Edit) X事件_颜色被改变1(pFun XE_EDIT_COLOR_CHANGE1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_COLOR_CHANGE, pFun)
}

// 元素事件_交换行.
func (e *Edit) X事件_交换行(pFun XE_EDIT_SWAPROW) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_SWAPROW, pFun)
}

// 元素事件_交换行.
func (e *Edit) X事件_交换行1(pFun XE_EDIT_SWAPROW1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_SWAPROW, pFun)
}

// 元素事件_编辑框设置.
func (e *Edit) X事件_编辑框设置(pFun XE_EDIT_SET) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_SET, pFun)
}

// 元素事件_编辑框设置.
func (e *Edit) X事件_编辑框设置1(pFun XE_EDIT_SET1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_SET, pFun)
}

// 暂未使用.
func (e *Edit) X暂未使用Event_EDIT_DRAWROW(pFun XE_EDIT_DRAWROW) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_DRAWROW, pFun)
}

// 暂未使用.
func (e *Edit) X暂未使用Event_EDIT_DRAWROW1(pFun XE_EDIT_DRAWROW1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_DRAWROW, pFun)
}

// 编辑框_内容被改变.
func (e *Edit) X事件_内容被改变(pFun XE_EDIT_CHANGED) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_CHANGED, pFun)
}

// 编辑框_内容被改变.
func (e *Edit) X事件_内容被改变1(pFun XE_EDIT_CHANGED1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_CHANGED, pFun)
}

// 编辑框_光标位置_被改变.
func (e *Edit) X事件_光标位置被改变(pFun XE_EDIT_POS_CHANGED) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_POS_CHANGED, pFun)
}

// 编辑框_光标位置_被改变.
func (e *Edit) X事件_光标位置被改变1(pFun XE_EDIT_POS_CHANGED1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_POS_CHANGED, pFun)
}

// 编辑框_样式_被改变.
func (e *Edit) X事件_样式被改变(pFun XE_EDIT_STYLE_CHANGED) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_STYLE_CHANGED, pFun)
}

// 编辑框_样式_被改变.
func (e *Edit) X事件_样式被改变1(pFun XE_EDIT_STYLE_CHANGED1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_STYLE_CHANGED, pFun)
}

// 回车TAB对齐,返回需要TAB数量.
func (e *Edit) X事件_回车TAB对齐(pFun XE_EDIT_ENTER_GET_TABALIGN) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_ENTER_GET_TABALIGN, pFun)
}

// 回车TAB对齐,返回需要TAB数量.
func (e *Edit) X事件_回车TAB对齐1(pFun XE_EDIT_ENTER_GET_TABALIGN1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_ENTER_GET_TABALIGN, pFun)
}

// 编辑框_行_被改变.
//
// iRow: 更改行开始位置索引,  if(nChangeRows>0) iEnd= iRow + nChangeRows
//
// nChangeRows: 改变行数, 正数添加行, 负数删除行
func (e *Edit) X事件_行被改变(pFun XE_EDIT_ROW_CHANGED) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDIT_ROW_CHANGED, pFun)
}

// 编辑框_行_被改变.
//
// iRow: 更改行开始位置索引,  if(nChangeRows>0) iEnd= iRow + nChangeRows
//
// nChangeRows: 改变行数, 正数添加行, 负数删除行
func (e *Edit) X事件_行被改变1(pFun XE_EDIT_ROW_CHANGED1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDIT_ROW_CHANGED, pFun)
}
