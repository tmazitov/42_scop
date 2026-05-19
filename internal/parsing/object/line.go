package objectParsing

import (
	"strings"
)

type objLineType int8

const (
	objVertex objLineType = iota
	objInit
	objVertexTexture
	objVertexNormal
	objFace
	objNewMaterial
	objUseMaterial
	objSmooth
	objComment
	objNone
)

var objLineDictionary = map[string]objLineType{
	"o":      objInit,
	"#":      objComment,
	"g":      objComment, // groups are organizational, silently skip
	"l":      objComment, // line elements, silently skip
	"mg":     objComment, // merging groups, silently skip
	"p":      objComment, // point elements, silently skip
	"v":      objVertex,
	"vt":     objVertexTexture,
	"vn":     objVertexNormal,
	"f":      objFace,
	"mtllib": objNewMaterial,
	"usemtl": objUseMaterial,
	"s":      objSmooth,
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
