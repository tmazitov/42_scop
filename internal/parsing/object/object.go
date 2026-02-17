package objectParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"bufio"
	"fmt"
	"log"
	"path/filepath"
	"os"
)


func ParseObj(filePath string) (*rende.Object, error) {

	var (
		object *rende.Object
		objectParseProcess = newObjectParsingProcess(filePath)
	)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	object = rende.NewObject(filepath.Base(filePath))

    // Create a new scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Loop through the file and read each line
	var (
		counter = -1
		lineArgs []string
		lineType objLineType
	)

    for scanner.Scan() {
		counter++
        line := scanner.Text() // Get the line as a string
		lineType, lineArgs = filterObjFileLine(line) 
		if lineType == objNone{
			log.Printf("obj parsing warn : unsupported line '%s'\n", line)
			continue
		}

		lineHandler, ok := objParsingActionsDictionary[lineType]
		if !ok {
			return nil, fmt.Errorf("obj parsing line %d error : unsupported line type", counter)
		}

		err = lineHandler(objectParseProcess, lineArgs)
		if err != nil {
			return nil, fmt.Errorf("obj parsing line %d error : %w", counter, err)
		}
    }

    if err := scanner.Err(); err != nil {
		return nil, err
    }

	if err := objectParseProcess.Prepare(); err != nil {
		return nil, err
	}

	object.
		SetShape(objectParseProcess.vertices).
		SetIndices(objectParseProcess.indices).
		SetMaterials(objectParseProcess.materials)

	return object, nil
}
