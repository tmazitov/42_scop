package parsing

var (
	vLine = "v"
	fLine = "f"
)

func filterLine(line string) string {
	if len(line) == 0 {
		return "x"
	}

	switch (line[0]) {
	case 'v':
		return "v"
	case 'f':
		return "f"
	default:
		return "x"
	} 
}