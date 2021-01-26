package main

// #cgo LDFLAGS: -L${SRCDIR}/lib -ladd
// #include "glue.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Print("Hello ")
	cstrFormat := C.CString("%s\n")
	defer C.free(unsafe.Pointer(cstrFormat))
	cstrArg := C.CString("Crustaceans")
	defer C.free(unsafe.Pointer(cstrArg))
	C.myprint(cstrFormat, cstrArg)
	fmt.Println("42: ", C.add(40, 2))
}
