package 炫彩WinApi工具类//bm:炫彩WinApi工具类

import (
	"errors"
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/wapi"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

// waitOpenClipboard 循环打开剪贴板，最多等待一秒钟.
//
//	@return bool
func waitOpenClipboard() bool {
	started := time.Now()
	limit := started.Add(time.Second)
	for time.Now().Before(limit) {
		if 炫彩WinApi类.X剪辑版打开(0) {
			return true
		}
		time.Sleep(time.Millisecond)
	}
	return false
}

// GetClipboardText 获取剪贴板中的文本.
//
//	@return string
//	@return error

// ff:剪贴板取文本
func X剪贴板取文本() (string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// 确定剪贴板是否包含指定格式的数据
	if !炫彩WinApi类.X剪贴板内容格式判断(炫彩WinApi类.CF_UNICODETEXT) {
		return "", nil // 剪贴板中不包含Unicode文本数据
	}

	// 打开剪贴板
	if !waitOpenClipboard() {
		return "", errors.New("打开剪贴板失败")
	}
	defer 炫彩WinApi类.X剪辑版关闭() // 关闭剪贴板

	// 取剪贴板数据句柄
	hMem := 炫彩WinApi类.X剪贴板取指定格式内容(炫彩WinApi类.CF_UNICODETEXT)
	if hMem == 0 {
		return "", errors.New("未获取到剪贴板数据句柄")
	}

	// 锁定
	lpData := 炫彩WinApi类.X全局内存对象锁定(hMem)
	if lpData == 0 {
		return "", errors.New("锁定剪贴板数据内存失败")
	}
	defer 炫彩WinApi类.X全局内存对象解锁(hMem) // 解锁

	// 获取数据大小
	nSize := 炫彩WinApi类.X全局内存对象取大小(hMem)
	if nSize == 0 {
		return "", errors.New("获取到剪贴板文本数据尺寸为0")
	}

	buf := make([]uint16, nSize)
	炫彩WinApi类.X内存复制(炫彩工具类.Uint16SliceDataPtr(&buf), lpData, nSize)
	return syscall.UTF16ToString(buf), nil
}

// SetClipboardText 将文本置入剪贴板.
//
//	@param text 要置入的文本.
//	@return error

// ff:剪贴板置文本
// text:文本
func X剪贴板置文本(文本 string) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// 打开剪贴板
	if !waitOpenClipboard() {
		return errors.New("打开剪贴板失败")
	}
	defer 炫彩WinApi类.X剪辑版关闭() // 关闭剪贴板

	// 清空剪贴板
	if !炫彩WinApi类.X剪辑版清空() {
		return errors.New("清空剪贴板失败")
	}

	// 转换
	bytes, _ := syscall.UTF16FromString(文本)
	dwLength := len(bytes) * int(unsafe.Sizeof(bytes[0]))

	// 分配全局内存
	hGlobalMemory := 炫彩WinApi类.X全局内存分配(炫彩WinApi类.GMEM_Moveable, uint(dwLength))
	if hGlobalMemory == 0 {
		return errors.New("分配全局内存失败")
	}
	defer 炫彩WinApi类.X全局内存对象关闭(hGlobalMemory) // 释放全局内存

	// 锁住内存区
	lpGlobalMemory := 炫彩WinApi类.X全局内存对象锁定(hGlobalMemory)
	if lpGlobalMemory == 0 {
		return errors.New("锁住内存区失败")
	}
	defer 炫彩WinApi类.X全局内存对象解锁(hGlobalMemory) // 解锁内存区

	// 内存复制
	h := 炫彩WinApi类.X复制到缓冲区(lpGlobalMemory, uintptr(unsafe.Pointer(&bytes[0])))
	if h == 0 {
		return errors.New("内存复制失败")
	}

	// 设置剪贴板数据
	if 炫彩WinApi类.X剪贴板设置数据(炫彩WinApi类.CF_UNICODETEXT, hGlobalMemory) == 0 {
		return errors.New("设置剪贴板数据失败")
	}

	return nil
}
