package appx

type State struct {
	IsVertexOnly bool
}

func NewState() *State {
	return &State{
		IsVertexOnly: false,
	}
}