package undbg

import (
	"syscall"
	"log"
)

func (dbg *undbg) revStep(n int) int {
	if n != 1 {
		log.Fatal("RevStep is not implemented for multiple instructions!")
	}

	dbg.idx--
	currentState := dbg.state()
	syscall.PtraceSetRegs(dbg.pid, &currentState.regs)
	return 1
}
