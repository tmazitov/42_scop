package parsing

import (
	"strings"
)

func filterObjFileLine(line string) ([]string) {
	if len(line) == 0 {
		return nil
	}

	args := strings.Split(line, " ")
	if len(args) == 0 {
		return nil
	}


	switch (args[0]) {
	case "v", "f", "mtllib":
		return args
	default:
		return nil
	} 
}
