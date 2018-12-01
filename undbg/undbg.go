package undbg

import (
	"os"
	"os/exec"
	"syscall"
	"fmt"
	"github.com/frankpf/undbg/utils"
	"log"
)

func Start(input string) {
	cmd := startTarget(input)
	startDebugger(cmd.Process.Pid)
}

func startDebugger(pid int) {
	var ws syscall.WaitStatus
	var icounter = 0

	/* Wait for target to stop on its first instruction */
	syscall.Wait4(pid, &ws, syscall.WALL, nil)

	var cmd string
	for ws.Stopped() {
		cmd = utils.ReadLine("> ")
		fn := parseCommand(cmd, pid)
		if (fn == nil) {
			fmt.Println("Invalid command \"" + cmd + "\"")
		} else {
			stepped := fn(pid, &ws)
			icounter += stepped
		}
	}

	fmt.Printf("Total instructions executed = %d\n", icounter)
}

func startTarget(name string) *exec.Cmd {
	log.Println("Starting target " + name)

	cmd := exec.Command(name)
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}

