package arglist

import "os"

func Gen() []string {
	return os.Args[1:]
}
