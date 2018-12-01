package main

import (
	"strconv"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"strings"
)

func ReadLine(prompt string) string {
	fmt.Printf(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func StartTarget(name string) *exec.Cmd {
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

type StepFunc func (int, *syscall.WaitStatus) int


func ParseCommand(cmd string, pid int) StepFunc {
	if (cmd == "step" || cmd == "s") {
		return Step(1)
	} else if (strings.HasPrefix(cmd, "step ") || strings.HasPrefix(cmd, "s ")) {
		countStr := strings.Split(cmd, " ")[1]
		count, err := strconv.Atoi(countStr)
		if (err != nil) {
			fmt.Println(countStr + " is not a valid argument to step")
			return nil
		}

		return Step(count)
	} else if (cmd == "rev-step" || cmd == "rs") {
		return RevStep(1)
	} else {
		return nil
	}
}

func StartDebugger(pid int) {
	var ws syscall.WaitStatus
	var icounter = 0

	/* Wait for target to stop on its first instruction */
	syscall.Wait4(pid, &ws, syscall.WALL, nil)

	var cmd string
	for ws.Stopped() {
		cmd = ReadLine("> ")
		fn := ParseCommand(cmd, pid)
		if (fn == nil) {
			fmt.Println("Invalid command \"" + cmd + "\"")
		} else {
			stepped := fn(pid, &ws)
			icounter += stepped
		}
	}

	fmt.Printf("Instructions = %d\n", icounter)
}

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

func main() {
	flag.Parse()
	input := flag.Arg(0)

	if input == "" {
		fmt.Println("Usage: undbg <target>")
		os.Exit(1)
	}

	cmd := StartTarget(input)

	StartDebugger(cmd.Process.Pid)
}
