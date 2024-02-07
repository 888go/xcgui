package 炫彩模板类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 节点.
type Node struct {
	Handle int
}

// 模板_创建节点.
//
// nType: 对象类型: xcc.XC_OBJECT_TYPE : xcc.XC_.
func X创建节点(对象类型 炫彩常量类.XC_OBJECT_TYPE) *Node {
	p := &Node{
		Handle: 炫彩基类.X模板_创建节点(对象类型),
	}
	return p
}

// 给本类的Handle赋值.
func (n *Node) X设置句柄(句柄 int) {
	n.Handle = 句柄
}

// 模板_取节点, 获取节点, 根据itemID. 返回节点对象.
//
// pNode: 节点指针.
//
// itemID: ID.
func (n *Node) X取节点(ID int32) *Node {
	p := &Node{
		Handle: 炫彩基类.X模板_取节点(n.Handle, ID),
	}
	return p
}

// 模板_克隆节点, 获取列表项模板类型, 返回节点对象.
func (n *Node) X取副本() *Node {
	p := &Node{
		Handle: 炫彩基类.X模板_克隆节点(n.Handle),
	}
	return p
}

// 模板_添加子节点.
//
// pNode: 节点指针.
func (n *Node) X添加子节点(节点指针 int) bool {
	return 炫彩基类.X模板_添加子节点(n.Handle, 节点指针)
}

// 模板_置节点属性.
//
// pName: 属性名.
//
// pAttr: 属性值.
func (n *Node) X置节点属性(属性名 string, 属性值 string) bool {
	return 炫彩基类.X模板_置节点属性(n.Handle, 属性名, 属性值)
}

// 模板_置节点属性扩展.
//
// itemID: 模板项ID.
//
// pName: 属性名.
//
// pAttr: 属性值.
func (n *Node) X置节点属性EX(模板项ID int32, 属性名 string, 属性值 string) bool {
	return 炫彩基类.X模板_置节点属性EX(n.Handle, 模板项ID, 属性名, 属性值)
}
