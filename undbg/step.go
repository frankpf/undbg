package undbg

import (
	"fmt"
	"golang.org/x/arch/x86/x86asm"
	"log"
	"syscall"
)

func (dbg *undbg) printRegs() {
	regs := dbg.getCurrentRegs()
	fmt.Printf("%+v\n", regs)
}

func (dbg *undbg) getCurrentRegs() syscall.PtraceRegs {
	var regs = syscall.PtraceRegs{}
	err := syscall.PtraceGetRegs(dbg.pid, &regs)
	if err != nil {
		log.Fatal(err)
	}
	return regs
}

func createStateSnapshot(regs syscall.PtraceRegs, memWrite bool, memDst uintptr, memValue uint64) state {
	s := state{
		regs:     regs,
		memWrite: memWrite,
		memDst:   memDst,
		memValue: memValue,
	}

	return s
}

// FIXME: This is really hacky. We probably should be using
// structs and methods to manage access to the program state.
var globalState []syscall.PtraceRegs
var currentReg = 0

func (dbg *undbg) pc() uintptr {
	currentState := dbg.state()
	return uintptr(currentState.regs.Rip)
}

func (dbg *undbg) state() state {
	currentState := dbg.states[dbg.idx]
	return currentState
}

func (dbg *undbg) printCurrentInstruction() {
	text := make([]byte, 15)
	_, err := syscall.PtracePeekText(dbg.pid, dbg.pc(), text)
	if err != nil {
		log.Fatal(err)
	}

	inst, _ := x86asm.Decode(text, 64)
	fmt.Println("Inst " + x86asm.IntelSyntax(inst, 0, nil))
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

func (dbg *undbg) revStep(n int) int {
	if n != 1 {
		log.Fatal("RevStep is not implemented for multiple instructions!")
	}

	dbg.idx--
	currentState := dbg.state()
	syscall.PtraceSetRegs(dbg.pid, &currentState.regs)
	return 1
}

func (dbg *undbg) step(n int) int {
	var i = 0
	for ; i < n; i++ {
		dbg.stepOnce()
	}
	return i
}
