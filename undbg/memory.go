package undbg

import (
	"syscall"
	"encoding/binary"
	"fmt"
	"log"
)

func (dbg *undbg) getValueAtAddr(addr uintptr) uint64 {
	buf := make([]byte, 8)
	count, err := syscall.PtracePeekData(dbg.pid, addr, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[memory.getValue] count = %d\n", count)
	value := binary.LittleEndian.Uint64(buf)
	return value
}

func (dbg *undbg) setValueAtAddr(addr uintptr, value uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, value)
	count, err := syscall.PtracePokeData(dbg.pid, addr, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[memory.setValue] count = %d\n", count)
}
