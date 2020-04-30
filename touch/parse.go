package touch

import (
	"strings"
)

func Split(filename string) (name, suffix string) {
	res := strings.Split(filename, ".")
	name = res[0]
	suffix = res[1]
	return
}
