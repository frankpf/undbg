package undbg

import (
	"syscall"
	"golang.org/x/arch/x86/x86asm"
	"fmt"
	"log"
)

func (dbg *undbg) printRegs() {
	regs := dbg.getCurrentRegs()
	fmt.Printf("%+v\n", regs)
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
