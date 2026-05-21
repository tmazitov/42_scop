package appx

type State struct {
	IsVertexOnly      bool
	IsRotationEnabled bool
	IsTextureEnabled  bool
	TextureBlend      float32
}

func NewState() *State {
	return &State{
		IsVertexOnly:      false,
		IsRotationEnabled: false,
		IsTextureEnabled:  false,
		TextureBlend:      0,
	}
}