package undbg

import (
	"strings"
	"strconv"
	"fmt"
)

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
