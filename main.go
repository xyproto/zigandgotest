package main

// #cgo LDFLAGS: -L${SRCDIR}/lib -ladd
// #include "glue.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello from Go")
	cstrFormat := C.CString("%s\n")
	defer C.free(unsafe.Pointer(cstrFormat))
	cstrArg := C.CString("Hello from C")
	defer C.free(unsafe.Pointer(cstrArg))
	C.myprint(cstrFormat, cstrArg)
	fmt.Println("40 + 2 == ", C.add(40, 2))
}
