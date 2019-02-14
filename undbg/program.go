package undbg

import (
	"syscall"
)

func (dbg *undbg) pc() uintptr {
	currentState := dbg.state()
	return uintptr(currentState.regs.Rip)
}

func (dbg *undbg) state() state {
	currentState := dbg.states[dbg.idx]
	return currentState
}

func (dbg *undbg) wait() {
	syscall.Wait4(dbg.pid, &dbg.ws, syscall.WALL, nil)
}

func (dbg *undbg) stopped() bool {
	return dbg.ws.Stopped()
}

func createStateSnapshot(regs syscall.PtraceRegs, memWrite bool, memDst uintptr, memOldValue uint64) state {
	s := state{
		regs:     regs,
		memWrite: memWrite,
		memDst:   memDst,
		memOldValue: memOldValue,
	}

	return s
}
