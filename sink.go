package main

// #cgo CFLAGS: -I/opt/halon/include
// #cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
// #include <HalonMTA.h>
// #include <stdlib.h>
import "C"
import (
	"time"
	"unsafe"
)

func main() {}

//export Halon_version
func Halon_version() C.int {
	return C.HALONMTA_PLUGIN_VERSION
}

func deliverTask(hdc *C.HalonDeliverContext, delay float64) {
	if delay != 0 {
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	code := C.int(250)
	reason := C.CString("Ok")
	reason_up := unsafe.Pointer(reason)
	defer C.free(reason_up)
	C.HalonMTA_deliver_setinfo(hdc, C.HALONMTA_RESULT_CODE, unsafe.Pointer(&code), 0)
	C.HalonMTA_deliver_setinfo(hdc, C.HALONMTA_RESULT_REASON, reason_up, 0)
	C.HalonMTA_deliver_done(hdc)
}

//export Halon_deliver
func Halon_deliver(hdc *C.HalonDeliverContext) {
	var delay_float64 float64 = 0
	var arguments *C.HalonHSLValue
	if C.HalonMTA_deliver_getinfo(hdc, C.HALONMTA_INFO_ARGUMENTS, nil, 0, unsafe.Pointer(&arguments), nil) {
		delay_cs := C.CString("delay")
		delay_cs_up := unsafe.Pointer(delay_cs)
		defer C.free(delay_cs_up)
		delay_hv := C.HalonMTA_hsl_value_array_find(arguments, delay_cs)
		if delay_hv != nil {
			var delay_cd C.double
			if C.HalonMTA_hsl_value_get(delay_hv, C.HALONMTA_HSL_TYPE_NUMBER, unsafe.Pointer(&delay_cd), nil) {
				delay_float64 = float64(delay_cd)
			}
		}
	}
	go deliverTask(hdc, delay_float64)
}
