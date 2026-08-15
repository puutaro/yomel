package arg_list

import (
	"os"
)

func Gen() []string {
	return os.Args[1:]
}
