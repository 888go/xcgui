package ani

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// 动画特效基类.
type animaBase struct {
	objectbase.ObjectBase
}

// 动画_运行, 并且增加引用计数.
//. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.
// ff:运行
// hRedrawObjectUI:当更新UI时重绘的UI层
func (a *animaBase) Run(hRedrawObjectUI int) int {
	return xc.XAnima_Run(a.Handle, hRedrawObjectUI)
}

// 动画_释放, 停止动画, 并释放引用, 当引用计数为0时自动销毁.
//.
// ff:释放
// bEnd:是否立即执行到终点
func (a *animaBase) Release(bEnd bool) bool {
	return xc.XAnima_Release(a.Handle, bEnd)
}

// 动画_释放扩展, 释放与指定UI对象关联的所有动画, 返回释放动画数量.
//.
// ff:释放EX
// bEnd:是否立即执行到终点
func (a *animaBase) ReleaseEx(bEnd bool) int {
	return xc.XAnima_ReleaseEx(a.Handle, bEnd)
}

// 动画_取UI对象, 获取动画关联的UI对象, 返回UI对象句柄.
// ff:取UI对象
func (a *animaBase) GetObjectUI() int {
	return xc.XAnima_GetObjectUI(a.Handle)
}

// 动画_启用自动销毁, TRUE: 当引用计数为0时自动销毁, FALSE: 手动销毁.
//.
// ff:启用自动销毁
// bEnable:是否启用
func (a *animaBase) EnableAutoDestroy(bEnable bool) int {
	return xc.XAnima_EnableAutoDestroy(a.Handle, bEnable)
}

// 动画_置回调.
//.
// ff:置回调
// callback:回调函数
func (a *animaBase) SetCallBack(callback interface{}) int {
	return xc.XAnima_SetCallBack(a.Handle, callback)
}

// 动画_置用户数据.
//.
// ff:置用户数据
// nUserData:用户数据
func (a *animaBase) SetUserData(nUserData int) int {
	return xc.XAnima_SetUserData(a.Handle, nUserData)
}

// 动画_取用户数据, 返回用户数据.
// ff:取用户数据
func (a *animaBase) GetUserData() int {
	return xc.XAnima_GetUserData(a.Handle)
}

// 动画_停止.
// ff:停止
func (a *animaBase) Stop() bool {
	return xc.XAnima_Stop(a.Handle)
}

// 动画_启动.
// ff:启动
func (a *animaBase) Start() bool {
	return xc.XAnima_Start(a.Handle)
}

// 动画_暂停.
// ff:暂停
func (a *animaBase) Pause() bool {
	return xc.XAnima_Pause(a.Handle)
}
