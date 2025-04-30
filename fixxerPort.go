//go:build windows

package main

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func doInjection(dllPath string) {
	// convert dll path to UTF16
	dllPathPtr, err := windows.UTF16PtrFromString(dllPath)
	if err != nil {
		panic(err)
	}
	// open the process we suspended with all access
	process, err := windows.OpenProcess(PROCESS_ALL_ACCESS, false, uint32(gProcessInfo.ProcessId))
	if err != nil {
		panic(err)
	}

	// get address of LoadLibraryA
	LoadLibAddy, err := GetAddressLoadLibraryW()
	if err != nil {
		panic(err)
	}

	// Allocate memory within Daybreak's virtual address space, this is where the DLL path will be written.
	remoteMem, err := VirtualAllocEx(process, 0, uintptr(len(dllPath)+1), windows.MEM_RESERVE|windows.MEM_COMMIT, windows.PAGE_READWRITE)
	if err != nil {
		panic(err)
	}

	// writing the dll path size to the allocated memory
	if err = WriteProcessMemory(process, remoteMem, dllPathPtr, len(dllPath)+1, 0); err != nil {
		panic(err)
	}
	// Inject.
	remoteThread, err := CreateRemoteThread(process, 0, 0, LoadLibAddy, remoteMem, 0, 0)
	if err != nil {
		panic(err)
	}
	// Wait for injection to complete then close the thread handle.
	windows.WaitForSingleObject(remoteThread, windows.INFINITE)
	windows.CloseHandle(remoteThread)
	// free the memory we allocated in the target process
	if err = VirtualFreeEx(process, remoteMem, 0, windows.MEM_RELEASE); err != nil {
		panic(err)
	}
	// resume process
	returned, err := windows.ResumeThread(gProcessInfo.Thread)
	if err != nil && err != syscall.Errno(0) {
		panic(err)
	}
	if returned == 0 {
		panic("Failed to resume thread")
	}
	print("Injected\n")

}

var gProcessInfo *windows.ProcessInformation

func fixxer() {
	gProcessInfo = &windows.ProcessInformation{}
	exePath := "./DaybreakDX.exe"
	procName, err := windows.UTF16PtrFromString(exePath)
	if err != nil {
		panic(err)
	}
	// create a suspended process
	if err = windows.CreateProcess(procName, nil, nil, nil, false, windows.CREATE_SUSPENDED, nil, nil, &windows.StartupInfo{}, gProcessInfo); err != nil {
		panic(err)
	}

	// set affinity to a single core
	if err = SetProcessAffinityMask(gProcessInfo, 1); err != nil {
		panic(err)
	}
	// sets priority to high
	if err = SetPriorityClass(gProcessInfo, windows.HIGH_PRIORITY_CLASS); err != nil {
		panic(err)
	}

	myDll := "./DaybreakFixer.dll"

	// inject dll
	doInjection(myDll)

	windows.CloseHandle(gProcessInfo.Process)
	windows.CloseHandle(gProcessInfo.Thread)
}
