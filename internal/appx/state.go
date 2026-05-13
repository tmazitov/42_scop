package appx

type State struct {
	IsVertexOnly		bool
	IsRotationEnabled	bool
}

func NewState() *State {
	return &State{
		IsVertexOnly: false,
		IsRotationEnabled: false,
	}
}