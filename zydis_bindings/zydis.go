package zydis_bindings

// #cgo CFLAGS: -I ${SRCDIR}/../dependencies/zydis6/include/
// #cgo CFLAGS: -I ${SRCDIR}/../dependencies/zydis6/dependencies/zycore/include/
// #cgo CFLAGS: -I ${SRCDIR}/../dependencies/zydis6/build/
// #cgo CFLAGS: -I ${SRCDIR}/../dependencies/zydis6/build/dependencies/zycore/
// #cgo LDFLAGS: ${SRCDIR}/../dependencies/zydis6/build/libZydis.a
// #include "test.h"
import "C"

func PrintBytes(buf []byte) {
	ptr := C.CBytes(buf)
	length := len(buf)
	C.printme((*C.uchar)(ptr), (C.ulong)(length), 0)
}

func main() {
	buf := make([]byte, 5)
	buf = append(buf, 0xb8)
	buf = append(buf, 0x3c)
	buf = append(buf, 0x00)
	buf = append(buf, 0x00)
	buf = append(buf, 0x00)

	data := C.CBytes(buf)
	C.printme((*C.uchar)(data), 5, 0)
}
