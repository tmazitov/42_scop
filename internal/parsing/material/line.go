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
    mtlEmissiveColor
    mtlOpticalDensity
    mtlDissolve
    mtlTransparency
    mtlTransmissionFilter
    mtlIlluminationModel
	mtlDiffuseTextureMap
	mtlAmbientTextureMap
	mtlSpecularTextureMap
	mtlShininessTextureMap
	mtlDissolveTextureMap
	mtlComment
    mtlNone
)

var mtlLineDictionary = map[string]mtlLineType {
	"newmtl"  : mtlNewMaterial,
	"Ns"      : mtlShininess,
	"Ka"      : mtlAmbientColor,
	"Kd"      : mtlDiffuseColor,
	"Ks"      : mtlSpecularColor,
	"Ke"      : mtlEmissiveColor,
	"Ni"      : mtlOpticalDensity,
	"d"       : mtlDissolve,
	"Tr"      : mtlTransparency,
	"Tf"      : mtlTransmissionFilter,
	"illum"   : mtlIlluminationModel,
	"map_Kd"  : mtlDiffuseTextureMap,
	"map_Ka"  : mtlAmbientTextureMap,
	"map_Ks"  : mtlSpecularTextureMap,
	"map_Ns"  : mtlShininessTextureMap,
	"map_d"   : mtlDissolveTextureMap,
	// PBR extensions — not applicable in fixed-function GL, silently ignored
	"Pr"      : mtlComment,
	"Pm"      : mtlComment,
	"Ps"      : mtlComment,
	"Pc"      : mtlComment,
	"Pcr"     : mtlComment,
	"Pds"     : mtlComment,
	"Pl"      : mtlComment,
	"aniso"   : mtlComment,
	"anisor"  : mtlComment,
	"map_Pr"  : mtlComment,
	"map_Pm"  : mtlComment,
	"map_Ps"  : mtlComment,
	"map_Ke"  : mtlComment,
	"norm"    : mtlComment,
	"#"       : mtlComment,
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