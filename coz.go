package cozgo

//#cgo LDFLAGS: -ldl
//#include <cozgo.h>
import "C"

//Begin tells Coz-Profiler to begin a profiled transaction
func Begin(name string) {
	C.cozBegin(C.CString(name))
}

//End tells Coz-Profiler to end the profiled transaction
func End(name string) {
	C.cozEnd(C.CString(name))
}

//NamedProgress marks a named progress checkpoint in the span of a profiled transaction
func NamedProgress(name string) {
	C.cozProgressNamed(C.CString(name))
}

//Progress marks an anonymous progress checkpoint in the span of a profiled transaction.
//Coz will try to pick up the line numbers and filename in this case
func Progress() {
	C.cozProgress()
}
