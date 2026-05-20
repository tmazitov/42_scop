package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"bufio"
	"log"
	"os"
)

func ParseMtl(filePath string) ([]*rende.Material, error) {

	var (
		proc *mtlParsingProcess = newMtlParsingProcess(filePath)
	)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

    // Create a new scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

    // Loop through the file and read each line
	var (
		counter = -1
		lineType mtlLineType
		lineArgs []string
		lineHandler mtlParsingFunc
		ok bool
	)


    for scanner.Scan() {
		counter++
        line := scanner.Text() // Get the line as a string
		if len(line) == 0 {
			continue
		}

		lineType, lineArgs = filterMtlFileLine(line)
		if lineType == mtlComment {
			continue
		}
		if lineType == mtlNone {
			log.Printf("mtl parsing warn : unsupported line '%s'\n", line)
			continue
		}

		if lineHandler, ok = mtlParsingActionsDictionary[lineType]; !ok {
			log.Printf("mtl parsing warn : unsupported line type '%s'\n", lineArgs[0])
			continue
		}

		if err := lineHandler(proc, lineArgs); err != nil {
			return nil, err
		}
    }

    if err := scanner.Err(); err != nil {
		return nil, err
    }

	return proc.materials, nil
}

