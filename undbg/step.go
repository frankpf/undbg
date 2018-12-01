package undbg

import (
	"log"
	"syscall"
	"fmt"
)

type StepFunc func (int, *syscall.WaitStatus) int

func RevStep(n int) StepFunc {
	if (n != 1) {
		log.Fatal("RevStep is not implemented for multiple instructions!")
		return nil
	}

	return func(pid int, ws *syscall.WaitStatus) int {
		fmt.Printf("RevStep is not implemented!\n")
		return 0
	}
}

func Step(n int) StepFunc {
	return func(pid int, ws *syscall.WaitStatus) int {
		var i = 0
		for ; i < n; i++ {
			StepOnce(pid, ws)
		}

		var msg = "Stepped %d instruction"
		if n > 1 {
			msg += "s"
		}
		msg += "\n"

		fmt.Printf(msg, n)

		return i
	}
}

func StepOnce(pid int, ws *syscall.WaitStatus) {
	if err := syscall.PtraceSingleStep(pid); err != nil {
		log.Fatal(err)
	}
	syscall.Wait4(pid, ws, syscall.WALL, nil)
}
