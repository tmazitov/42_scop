package objectParsing

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tmazitov/42_scop/internal/rende"
)

func ParseObj(filePath string) ([]*rende.Object, error) {

	var (
		objects            []*rende.Object
		materials          *materialStorage      = newMaterialStorage(filePath)
		sharedVertices     *vertexStorage        = newVertexStorage()
		objectParseProcess *objectParsingProcess = newObjectParsingProcess("Default", filePath, materials, sharedVertices)
	)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through the file and read each line
	var (
		counter  = -1
		lineArgs []string
		lineType objLineType
	)

main:
	for scanner.Scan() {
		counter++
		line := scanner.Text() // Get the line as a string
		lineType, lineArgs = filterObjFileLine(line)

		switch lineType {
		case objNone:
			log.Printf("obj parsing warn : unsupported line '%s'\n", line)
			continue main
		case objInit:

			if len(lineArgs) != 2 {
				return nil, ErrInvalidInitObjectLine
			}

			// Complete previous process
			if objectParseProcess != nil && !objectParseProcess.IsEmpty() {
				o, err := objectParseProcess.ToObject()
				if err != nil {
					return nil, err
				}
				objects = append(objects, o)
			}

			// Init new object (shares global vertex/texture/normal arrays)
			name := lineArgs[1]
			objectParseProcess = newObjectParsingProcess(name, filePath, materials, sharedVertices)

			continue main
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

	o, err := objectParseProcess.ToObject()
	if err != nil {
		return nil, err
	}

	objects = append(objects, o)

	return objects, nil
}
