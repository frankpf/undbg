package undbg

import (
	"syscall"
	"log"
)

func (dbg *undbg) step(n int) int {
	var i = 0
	for ; i < n; i++ {
		dbg.stepOnce()
	}
	return i
}

func (dbg *undbg) stepOnce() {
	err := syscall.PtraceSingleStep(dbg.pid)
	if err != nil {
		log.Fatal(err)
	}
	dbg.icounter++
	dbg.wait()

	regs := dbg.getCurrentRegs()
	dbg.idx++
	if dbg.idx == len(dbg.states) {
		snapshot := createStateSnapshot(regs, false, 0, 0)
		dbg.states = append(dbg.states, snapshot)
	}
}
