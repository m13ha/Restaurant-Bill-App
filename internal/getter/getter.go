package getter

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(query string, r *bufio.Reader) (string, error) {
	fmt.Print(query)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
