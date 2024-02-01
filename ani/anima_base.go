package ani

import (
	"e.coding.net/gogit/go/xcgui/objectbase"
	"e.coding.net/gogit/go/xcgui/xc"
)

// 动画特效基类.
type animaBase struct {
	objectbase.ObjectBase
}

// 动画_运行, 并且增加引用计数.
//
// hRedrawObjectUI: 当更新UI时重绘的UI层. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.
func (a *animaBase) Run(hRedrawObjectUI int) int {
	return xc.XAnima_Run(a.I句柄, hRedrawObjectUI)
}

// 动画_释放, 停止动画, 并释放引用, 当引用计数为0时自动销毁.
//
// bEnd: 是否立即执行到终点.
func (a *animaBase) Release(bEnd bool) bool {
	return xc.XAnima_Release(a.I句柄, bEnd)
}

// 动画_释放扩展, 释放与指定UI对象关联的所有动画, 返回释放动画数量.
//
// bEnd: 是否立即执行到终点.
func (a *animaBase) ReleaseEx(bEnd bool) int {
	return xc.XAnima_ReleaseEx(a.I句柄, bEnd)
}

// 动画_取UI对象, 获取动画关联的UI对象, 返回UI对象句柄.
func (a *animaBase) GetObjectUI() int {
	return xc.XAnima_GetObjectUI(a.I句柄)
}

// 动画_启用自动销毁, TRUE: 当引用计数为0时自动销毁, FALSE: 手动销毁.
//
// bEnable: 是否启用.
func (a *animaBase) EnableAutoDestroy(bEnable bool) int {
	return xc.XAnima_EnableAutoDestroy(a.I句柄, bEnable)
}

// 动画_置回调.
//
// callback: 回调函数.
func (a *animaBase) SetCallBack(callback interface{}) int {
	return xc.XAnima_SetCallBack(a.I句柄, callback)
}

// 动画_置用户数据.
//
// nUserData: 用户数据.
func (a *animaBase) SetUserData(nUserData int) int {
	return xc.XAnima_SetUserData(a.I句柄, nUserData)
}

// 动画_取用户数据, 返回用户数据.
func (a *animaBase) GetUserData() int {
	return xc.XAnima_GetUserData(a.I句柄)
}

// 动画_停止.
func (a *animaBase) Stop() bool {
	return xc.XAnima_Stop(a.I句柄)
}

// 动画_启动.
func (a *animaBase) Start() bool {
	return xc.XAnima_Start(a.I句柄)
}

// 动画_暂停.
func (a *animaBase) Pause() bool {
	return xc.XAnima_Pause(a.I句柄)
}
