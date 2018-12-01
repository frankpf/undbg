package undbg

import (
	"fmt"
	"strconv"
	"strings"
)

func (dbg *undbg) runCommand(cmd string) {
	if cmd == "regs" {
		dbg.printRegs()
		return
	}

	if cmd == "step" || cmd == "s" {
		dbg.step(1)
		return
	}

	if strings.HasPrefix(cmd, "step ") || strings.HasPrefix(cmd, "s ") {
		countStr := strings.Split(cmd, " ")[1]
		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Println(countStr + " is not a valid argument to step")
		} else {
			dbg.step(count)
		}
		return
	}

	if cmd == "rev-step" || cmd == "rs" {
		dbg.revStep(1)
		return
	}

	if cmd == "print" || cmd == "p" {
		dbg.printCurrentInstruction()
		return
	}

	fmt.Println("Invalid command \"" + cmd + "\"")

}
