package xc

import (
	"sync"
	"syscall"
)

var (
	uiThreadCallBackFunc func(data int) int
	uiThreadCallBackPtr  = syscall.NewCallback(uiThreadCallBack)
	rwm                  sync.RWMutex
)

// UI线程回调函数
func uiThreadCallBack(data int) int {
	return uiThreadCallBackFunc(data)
}

// 炫彩_调用界面线程Ex, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//	@param f: 回调函数.
func XC_CallUiThreadEx(f func(data int) int, data int) int {
	rwm.Lock()
	uiThreadCallBackFunc = f
	r, _, _ : xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(data))
	rwm.Unlock()
	return int(r)
}

// 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//	@param f: 回调函数, 没有参数也没有返回值, 可以直接使用匿名函数.
func XC_CallUT(f func()) {
	rwm.Lock()
	uiThreadCallBackFunc = func(data int) int {
		f()
		return 0
	}
	xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(0))
	rwm.Unlock()
}

// 用于在UI线程操作UI.
type UiThreader interface {
	UiThreadCallBack(data int) int // 回调函数, 把操作UI的代码写在这里面.
}

// 炫彩_调用界面线程1, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
func XC_CallUiThreader(u UiThreader, data int) int {
	rwm.Lock()
	uiThreadCallBackFunc = u.UiThreadCallBack
	r, _, _ : xC_CallUiThread.Call(uiThreadCallBackPtr, uintptr(data))
	rwm.Unlock()
	return int(r)
}
