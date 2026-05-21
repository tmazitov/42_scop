package appx

type State struct {
	IsVertexOnly      bool
	IsRotationEnabled bool
	IsTextureEnabled  bool
	TextureBlend      float32
	SelectedObjectIdx int
}

func NewState() *State {
	return &State{
		IsVertexOnly:      false,
		IsRotationEnabled: false,
		IsTextureEnabled:  false,
		TextureBlend:      0,
		SelectedObjectIdx: 0,
	}
}