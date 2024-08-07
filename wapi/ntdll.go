package 炫彩WinApi类

import "syscall"

var (
	// Library.
	ntdll = syscall.NewLazyDLL("ntdll.dll")

	// Functions.
	rtlMoveMemory = ntdll.NewProc("RtlMoveMemory")
)

// RtlMoveMemory 将源内存块的内容复制到目标内存块，并支持重叠的源内存块和目标内存块.
//
//	@Description 详情: https://docs.microsoft.com/zh-cn/windows/win32/devnotes/rtlmovememory.
//	@param Destination 指向要复制字节的目标内存块的指针.
//	@param Source 指向要复制字节的源内存块的指针.
//	@param Length 从源复制到目标中的字节数.

// ff:内存复制
// Length:字节数
// Source:源内存块指针
// Destination:目标内存块指针
func X内存复制(目标内存块指针 uintptr, 源内存块指针 uintptr, 字节数 uint) {
	rtlMoveMemory.Call(目标内存块指针, 源内存块指针, uintptr(字节数))
}
