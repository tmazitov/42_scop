package materialParsing

import (
	"strings"
)

type mtlLineType int8

const (
    mtlNewMaterial mtlLineType = iota
    mtlShininess
    mtlAmbientColor
    mtlDiffuseColor
    mtlSpecularColor
    mtlOpticalDensity
    mtlDissolve
    mtlIlluminationModel
	mtlDiffuseTextureMap
    mtlNone
)

var mtlLineDictionary = map[string]mtlLineType {
	"newmtl" : mtlNewMaterial,
	"Ns"	 : mtlShininess,
	"Ka"	 : mtlAmbientColor,
	"Kd"	 : mtlDiffuseColor,
	"Ks"	 : mtlSpecularColor,
	"Ni"	 : mtlOpticalDensity,
	"d"		 : mtlDissolve,
	"illum"	 : mtlIlluminationModel,
	"map_Kd" : mtlDiffuseTextureMap,
}

func filterMtlFileLine(line string) (mtlLineType, []string) {
	if len(line) == 0 {
		return mtlNone, nil
	}

	args := strings.Fields(line)
	if len(args) == 0 {
		return mtlNone, nil
	}

	lineType, ok := mtlLineDictionary[args[0]]
	if !ok {
		return mtlNone, nil
	}
	return lineType, args
}