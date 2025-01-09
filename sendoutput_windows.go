package main

import (
	"syscall"
	"unsafe"
)

var output string = ""
var g_callback uintptr

func sendOutput() {

	b := []byte(output)
	dataPtr := uintptr(unsafe.Pointer(&b[0]))
	dataSize := uintptr(uint32(len(b)))

	_, _, errNo := syscall.SyscallN(g_callback, dataPtr, dataSize)
	if errNo != 0 {
		println("Got error: %s", errNo.Error())
	}
	return
}
