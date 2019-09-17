package log

import (
	"strings"
)

func DepL(s string) {
	i := len(s)

	a := (120 - i) / 2

	print(strings.Repeat("=", a))
	print(" ", s, " ")
	print(strings.Repeat("=", 120-i-a))
	println()
}
