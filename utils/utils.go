package utils

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func ReadLine(prompt string) string {
	fmt.Printf(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}
