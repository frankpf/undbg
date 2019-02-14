package undbg

import (
	"fmt"
	"github.com/frankpf/undbg/utils"
	"log"
	"os"
	"os/exec"
	"syscall"
)

type state struct {
	regs     syscall.PtraceRegs
	memWrite bool
	memDst   uintptr
	memOldValue uint64
}

type undbg struct {
	pid      int
	ws       syscall.WaitStatus
	states   []state
	idx      int
	icounter int
}

func Start(input string) {
	cmd := startTarget(input)
	startDebugger(cmd.Process.Pid)
}

func newDebugger(pid int) *undbg {
	dbg := &undbg{
		pid:      pid,
		ws:       0,
		states:   make([]state, 0),
		idx:      -1,
		icounter: 0,
	}

	return dbg
}

func startDebugger(pid int) {
	dbg := newDebugger(pid)

	/* Wait for target to stop on its first instruction */
	dbg.wait()

	initialState := createStateSnapshot(dbg.getCurrentRegs(), false, 0, 0)
	dbg.states = append(dbg.states, initialState)
	dbg.idx++

	var cmd string
	for dbg.stopped() {
		cmd = utils.ReadLine("> ")
		dbg.runCommand(cmd)
	}

	fmt.Printf("Total instructions executed = %d\n", dbg.icounter)
}

func startTarget(name string) *exec.Cmd {
	log.Println("Starting target " + name)

	prog := exec.Command(name)
	prog.Stdout = os.Stdout
	prog.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err := prog.Start()
	if err != nil {
		log.Fatal(err)
	}
	return prog
}
