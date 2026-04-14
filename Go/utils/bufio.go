package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func LerString(msg string) string {
	fmt.Print(msg)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}
