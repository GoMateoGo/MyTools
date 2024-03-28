package utils

import (
	"syscall"
	"unsafe"
)

func RunGame(appName, path string) {
	dll, err := syscall.LoadDLL("rungame.dll")
	if err != nil {
		SendMessageBox("错误", "未找到rungame.dll:"+err.Error())
		//fmt.Println("dll未找到:", err.Error())
		return
	}

	//找C函数
	proc, err := dll.FindProc("RunGame") //C函数
	if err != nil {
		SendMessageBox("错误", "rungame.dll加载错误:"+err.Error())
		//fmt.Println("没有找到对应函数:", err)
		return
	}

	p := strToWidePtr(path)
	p2 := strToWidePtr(appName)

	//执行C函数
	call, _, err := proc.Call(p2, p)
	if err != nil {
		//SendMessageBox("错误", "RuntimeCall错误:"+err.Error())
		//fmt.Println("RuntimeCall错误:", err.Error())
	}
	if call > 0 {
		//fmt.Println("call结果", int64(call))
	}
}

// 字符串转指针,实际结果为*uint16
func strToPtr(s string) uintptr {
	b, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(b))
}

// 字符串转换宽字符<支持中文>
func strToWidePtr(s string) uintptr {
	b, _ := syscall.UTF16FromString(s)
	return uintptr(unsafe.Pointer(&b[0]))
}

// 字符串转char*
func strToCharPtr(s string) uintptr {
	b, _ := syscall.BytePtrFromString(s)
	p := unsafe.Pointer(b)
	return uintptr(p)
}

// 方法名字起的不是很好,实际上是转换为uintptr类型
func IntPtr(i int) uintptr {
	return uintptr(i)
}

// 封装的winAPi方法,具体使用可以去msdn查询
// 此处有个疑惑,有一些dll的调用需要释放,还没研究.
func SendMessageBox(title, msg string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	msgBox := user32.NewProc("MessageBoxW")
	msgBox.Call(IntPtr(0), strToWidePtr(msg), strToWidePtr(title), IntPtr(0))
}
