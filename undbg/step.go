package undbg

import (
	"log"
	"syscall"
	"fmt"
	"golang.org/x/arch/x86/x86asm"
)

type stepFunc func (int, *syscall.WaitStatus) int

// FIXME: This is really hacky. We probably should be using
// structs and methods to manage access to the program state.
var globalState []syscall.PtraceRegs
var currentReg = 0

func printCurrentInstruction(pid int, pc uint64) {
	text := make([]byte, 15)
	_, err := syscall.PtracePeekText(pid, uintptr(pc), text)
	if err != nil {
		log.Fatal(err)
	}

	inst, _ := x86asm.Decode(text, 64)
	fmt.Println("Inst " + x86asm.IntelSyntax(inst, 0, nil))
}

func stepOnce(pid int, ws *syscall.WaitStatus) {

	var regs = syscall.PtraceRegs{}
	err := syscall.PtraceGetRegs(pid, &regs)
	if err != nil {
		log.Fatal(err)
	}
	printCurrentInstruction(pid, regs.Rip)

	currentReg++
	if currentReg > len(globalState) {
		globalState = append(globalState, regs)
	}


	err = syscall.PtraceSingleStep(pid)
	if err != nil {
		log.Fatal(err)
	}


	syscall.Wait4(pid, ws, syscall.WALL, nil)
}

func revStep(n int) stepFunc {
	if (n != 1) {
		log.Fatal("RevStep is not implemented for multiple instructions!")
		return nil
	}

	return func(pid int, ws *syscall.WaitStatus) int {
		currentReg--
		syscall.PtraceSetRegs(pid, &globalState[currentReg])
		return 1
	}
}

func step(n int) stepFunc {
	return func(pid int, ws *syscall.WaitStatus) int {
		var i = 0
		for ; i < n; i++ {
			stepOnce(pid, ws)
		}

		return i
	}
}

