package screen

type ScreenType uint8

const (
	SCRMenu ScreenType = iota
	SCREditor
)

type Screen struct {
	Type ScreenType
}

var Scr Screen = Init()

func Init() Screen {
	return Screen{Type: SCRMenu}
}

func (s *Screen) Handle() {
	switch s.Type {
	case SCRMenu:
		HandleMenu()
	case SCREditor:
		HandleEditor()
	}
}
