package undbg

import (
	"syscall"
	"log"
)

func (dbg *undbg) getCurrentRegs() syscall.PtraceRegs {
	var regs = syscall.PtraceRegs{}
	err := syscall.PtraceGetRegs(dbg.pid, &regs)
	if err != nil {
		log.Fatal(err)
	}
	return regs
}
