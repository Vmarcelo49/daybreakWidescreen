package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var gProcessInfo *windows.ProcessInformation

func doInjection(dllPath string, kernel32 *windows.LazyDLL) {
	// convert dll path to utf16
	dllPathPtr, err := windows.UTF16PtrFromString(dllPath)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}
	// open process
	process, err := windows.OpenProcess(
		0xffff, // PROCESS_ALL_ACCESS
		false,
		uint32(gProcessInfo.ProcessId),
	)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}

	// get address of LoadLibraryA
	LoadLibAddy, err := syscall.GetProcAddress(syscall.Handle(kernel32.Handle()), "LoadLibraryA")
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}

	// Allocate memory within Daybreak's virtual address space.
	// This is where the DLL path will be written.
	remoteMem, _, err := kernel32.NewProc("VirtualAllocEx").Call(uintptr(process), 0, uintptr(len(dllPath)+1), windows.MEM_RESERVE|windows.MEM_COMMIT, windows.PAGE_READWRITE)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}
	// Write path for LoadLibraryA in previously allocated memory.
	_, _, err = kernel32.NewProc("WriteProcessMemory").Call(uintptr(process), remoteMem, uintptr(unsafe.Pointer(dllPathPtr)), uintptr(len(dllPath)+1), 0)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}
	// Inject.
	remoteThread, _, err := kernel32.NewProc("CreateRemoteThread").Call(uintptr(process), 0, 0, LoadLibAddy, remoteMem, 0, 0)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}
	// Wait for injection to complete.
	windows.WaitForSingleObject(windows.Handle(remoteThread), windows.INFINITE)
	// Close handle.
	windows.CloseHandle(windows.Handle(remoteThread))
	VirtualFreeEx := kernel32.NewProc("VirtualFreeEx")
	_, _, err = VirtualFreeEx.Call(uintptr(process), remoteMem, 0, windows.MEM_RELEASE)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}
	// resume process
	returned, err := windows.ResumeThread(windows.Handle(gProcessInfo.Thread))
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}
	if returned == 0 {
		panic("Failed to resume thread")
	}
	print("Injected\n")

}

func fixxer() {
	gProcessInfo = &windows.ProcessInformation{}
	kernel32 := windows.NewLazyDLL("kernel32.dll")

	exePath := "./DaybreakDX.exe"
	procName, err := windows.UTF16PtrFromString(exePath)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}

	err = windows.CreateProcess(
		procName,
		nil,
		nil,
		nil,
		false,
		windows.CREATE_SUSPENDED,
		nil,
		nil,
		&windows.StartupInfo{},
		gProcessInfo,
	)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}

	// set affinity to cpu 0
	_, _, err = kernel32.NewProc("SetProcessAffinityMask").Call(uintptr(gProcessInfo.Process), uintptr(0x1))
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}
	// sets priority to high
	_, _, err = kernel32.NewProc("SetPriorityClass").Call(uintptr(gProcessInfo.Process), uintptr(0x80))
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}

	myDll := "./DaybreakFixer.dll"

	// inject dll
	doInjection(myDll, kernel32)

	windows.CloseHandle(gProcessInfo.Process)
	windows.CloseHandle(gProcessInfo.Thread)
}
