package undbg

import (
	"strings"
	"strconv"
	"fmt"
)

func parseCommand(cmd string, pid int) stepFunc {
	if (cmd == "step" || cmd == "s") {
		return step(1)
	} else if (strings.HasPrefix(cmd, "step ") || strings.HasPrefix(cmd, "s ")) {
		countStr := strings.Split(cmd, " ")[1]
		count, err := strconv.Atoi(countStr)
		if (err != nil) {
			fmt.Println(countStr + " is not a valid argument to step")
			return nil
		}

		return step(count)
	} else if (cmd == "rev-step" || cmd == "rs") {
		return revStep(1)
	} else {
		return nil
	}
}
