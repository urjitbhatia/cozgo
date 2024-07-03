package cozgo

//#cgo LDFLAGS: -ldl
//#include <cozgo.h>
//#include <stdlib.h>
import "C"
import "unsafe"

//Begin tells Coz-Profiler to begin a profiled transaction
func Begin(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.cozBegin(cname)
}

//End tells Coz-Profiler to end the profiled transaction
func End(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.cozEnd(cname)
}

//NamedProgress marks a named progress checkpoint in the span of a profiled transaction
func NamedProgress(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.cozProgressNamed(cname)
}

//Progress marks an anonymous progress checkpoint in the span of a profiled transaction.
//Coz will try to pick up the line numbers and filename in this case
func Progress() {
	C.cozProgress()
}
