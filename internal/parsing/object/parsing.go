package objectParsing


type objParsingFunc func(object *objectParsingProcess, args []string) error

var (
	objParsingActionsDictionary = map[objLineType]objParsingFunc{
		objVertex 			: vertexHandler,
		objVertexTexture	: vertexTextureHandler,
		objVertexNormal		: vertexNormalHandler,
		objFace   			: faceHandler,
		objNewMaterial 		: materialHandler,
	}
)