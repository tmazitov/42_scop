package objectParsing

import "strconv"

func smoothGroupHandler(object *objectParsingProcess, args []string) error {
	if len(args) < 2 {
		return nil
	}

	arg := args[1]
	if arg == "off" {
		object.smoothGroup = 0
		return nil
	}

	n, err := strconv.Atoi(arg)
	if err != nil || n < 0 {
		return nil // ignore invalid values
	}
	object.smoothGroup = n
	return nil
}
