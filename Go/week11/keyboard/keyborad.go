package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetInteger() (int, error) {
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a = strings.TrimSpace(a) // python strip
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, nil
}
