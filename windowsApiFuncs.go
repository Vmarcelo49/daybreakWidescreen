//go:build windows

package main

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const PROCESS_ALL_ACCESS = 0x1F0FFF

var kernel32 *windows.LazyDLL

func VirtualAllocEx(hProcess windows.Handle, lpAddress, dwSize, flAllocationType, flProtect uintptr) (uintptr, error) {
	if kernel32 == nil {
		kernel32 = windows.NewLazyDLL("kernel32.dll")
	}
	remoteMem, _, err := kernel32.NewProc("VirtualAllocEx").Call(uintptr(hProcess), lpAddress, dwSize, flAllocationType, flProtect)
	if err != nil && err != syscall.Errno(0) {
		return 0, errors.New("VirtualAllocEx failed: " + err.Error())
	}
	return remoteMem, nil
}

func SetProcessAffinityMask(processInfo *windows.ProcessInformation, processAffinityMask uint32) error {
	if kernel32 == nil {
		kernel32 = windows.NewLazyDLL("kernel32.dll")
	}
	_, _, err := kernel32.NewProc("SetProcessAffinityMask").Call(uintptr(processInfo.Process), uintptr(processAffinityMask))
	if err != nil && err != syscall.Errno(0) {
		return errors.New("SetProcessAffinityMask failed: " + err.Error())
	}
	return nil
}

func SetPriorityClass(processInfo *windows.ProcessInformation, priorityClass uint32) error {
	if kernel32 == nil {
		kernel32 = windows.NewLazyDLL("kernel32.dll")
	}
	_, _, err := kernel32.NewProc("SetPriorityClass").Call(uintptr(processInfo.Process), uintptr(priorityClass))
	if err != nil && err != syscall.Errno(0) {
		return errors.New("SetPriorityClass failed: " + err.Error())
	}
	return nil
}

func GetAddressLoadLibraryW() (uintptr, error) {
	if kernel32 == nil {
		kernel32 = windows.NewLazyDLL("kernel32.dll")
	}
	loadLibraryProc := kernel32.NewProc("LoadLibraryW")
	err := loadLibraryProc.Find()
	if err != nil {
		return 0, errors.New("GetAddressLoadLibraryW failed: " + err.Error())
	}
	return loadLibraryProc.Addr(), nil
}

func WriteProcessMemory(hProcess windows.Handle, baseAddress uintptr, buffer *uint16, nBytesToBeWritten int, bytesWritten uintptr) error {
	if kernel32 == nil {
		kernel32 = windows.NewLazyDLL("kernel32.dll")
	}
	_, _, err := kernel32.NewProc("WriteProcessMemory").Call(uintptr(hProcess), baseAddress, uintptr(unsafe.Pointer(buffer)), uintptr(nBytesToBeWritten), bytesWritten)
	if err != nil && err != syscall.Errno(0) {
		return errors.New("WriteProcessMemory failed: " + err.Error())
	}
	return nil
}

func CreateRemoteThread(hProcess windows.Handle, lpThreadAttributes, dwStackSize, lpStartAddress, lpParameter, dwCreationFlags, lpThreadId uintptr) (windows.Handle, error) {
	if kernel32 == nil {
		kernel32 = windows.NewLazyDLL("kernel32.dll")
	}
	remoteThread, _, err := kernel32.NewProc("CreateRemoteThread").Call(uintptr(hProcess), lpThreadAttributes, dwStackSize, lpStartAddress, lpParameter, dwCreationFlags, lpThreadId)
	if err != nil && err != syscall.Errno(0) {
		return 0, errors.New("CreateRemoteThread failed: " + err.Error())
	}
	return windows.Handle(remoteThread), nil
}

func VirtualFreeEx(hProcess windows.Handle, lpAddress, dwSize, dwFreeType uintptr) error {
	if kernel32 == nil {
		kernel32 = windows.NewLazyDLL("kernel32.dll")
	}
	VirtualFreeEx := kernel32.NewProc("VirtualFreeEx")
	_, _, err := VirtualFreeEx.Call(uintptr(hProcess), lpAddress, dwSize, windows.MEM_RELEASE)
	if err != nil && err != syscall.Errno(0) {
		return errors.New("VirtualFreeEx failed: " + err.Error())
	}
	return nil
}
