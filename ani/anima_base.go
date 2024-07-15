package 炫彩动画类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
)

// 动画特效基类.
type animaBase struct {
	炫彩对象基类.ObjectBase
}

// 动画_运行, 并且增加引用计数.
//
// hRedrawObjectUI: 当更新UI时重绘的UI层. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.

// ff:运行
// hRedrawObjectUI:当更新UI时重绘的UI层
func (a *animaBase) X运行(当更新UI时重绘的UI层 int) int {
	return 炫彩基类.X动画_运行(a.Handle, 当更新UI时重绘的UI层)
}

// 动画_释放, 停止动画, 并释放引用, 当引用计数为0时自动销毁.
//
// bEnd: 是否立即执行到终点.

// ff:释放
// bEnd:是否立即执行到终点
func (a *animaBase) X释放(是否立即执行到终点 bool) bool {
	return 炫彩基类.X动画_释放(a.Handle, 是否立即执行到终点)
}

// 动画_释放扩展, 释放与指定UI对象关联的所有动画, 返回释放动画数量.
//
// bEnd: 是否立即执行到终点.

// ff:释放EX
// bEnd:是否立即执行到终点
func (a *animaBase) X释放EX(是否立即执行到终点 bool) int {
	return 炫彩基类.X动画_释放EX(a.Handle, 是否立即执行到终点)
}

// 动画_取UI对象, 获取动画关联的UI对象, 返回UI对象句柄.

// ff:取UI对象
func (a *animaBase) X取UI对象() int {
	return 炫彩基类.X动画_取UI对象(a.Handle)
}

// 动画_启用自动销毁, TRUE: 当引用计数为0时自动销毁, FALSE: 手动销毁.
//
// bEnable: 是否启用.

// ff:启用自动销毁
// bEnable:是否启用
func (a *animaBase) X启用自动销毁(是否启用 bool) int {
	return 炫彩基类.X动画_启用自动销毁(a.Handle, 是否启用)
}

// 动画_置回调.
//
// callback: 回调函数.

// ff:置回调
// callback:回调函数
func (a *animaBase) X置回调(回调函数 interface{}) int {
	return 炫彩基类.X动画_置回调(a.Handle, 回调函数)
}

// 动画_置用户数据.
//
// nUserData: 用户数据.

// ff:置用户数据
// nUserData:用户数据
func (a *animaBase) X置用户数据(用户数据 int) int {
	return 炫彩基类.X动画_置用户数据(a.Handle, 用户数据)
}

// 动画_取用户数据, 返回用户数据.

// ff:取用户数据
func (a *animaBase) X取用户数据() int {
	return 炫彩基类.X动画_取用户数据(a.Handle)
}

// 动画_停止.

// ff:停止
func (a *animaBase) X停止() bool {
	return 炫彩基类.X动画_停止(a.Handle)
}

// 动画_启动.

// ff:启动
func (a *animaBase) X启动() bool {
	return 炫彩基类.X动画_启动(a.Handle)
}

// 动画_暂停.

// ff:暂停
func (a *animaBase) X暂停() bool {
	return 炫彩基类.X动画_暂停(a.Handle)
}
