package objectParsing

import (
	"strings"
)

type objLineType int8

const (
    objVertex objLineType = iota
	objVertexTexture
	objVertexNormal
	objFace
	objNewMaterial
	objNone
)

var objLineDictionary = map[string]objLineType {
	"v" 	: objVertex,
	"vt"	: objVertexTexture,
	"vn"	: objVertexNormal,
	"f"		: objFace,
	"mtllib": objNewMaterial,
}
	
func filterObjFileLine(line string) (objLineType, []string) {
	if len(line) == 0 {
		return objNone, nil
	}

	args := strings.Fields(line)
	if len(args) == 0 {
		return objNone, nil
	}

	lineType, ok := objLineDictionary[args[0]]
	if !ok {
		return objNone, nil
	}
	return lineType, args 
}
