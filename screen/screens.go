package screen

import rl "github.com/gen2brain/raylib-go/raylib"

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

func (s *Screen) Handle(cam *rl.Camera2D) {
	switch s.Type {
	case SCRMenu:
		HandleMenu()
	case SCREditor:
		HandleEditor(cam)
	}
}
