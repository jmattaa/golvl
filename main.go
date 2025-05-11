package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jmattaa/golvl/level"
	"github.com/jmattaa/golvl/screen"
	"github.com/jmattaa/golvl/utils"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(utils.WIN_W, utils.WIN_H, "golvl - v"+utils.VERSION)
	defer rl.CloseWindow()

	rl.SetExitKey(0) // we don want no esc quit shi
	rl.SetTargetFPS(60)

	cam := rl.NewCamera2D(
		rl.NewVector2(float32(utils.WIN_W)/2, float32(utils.WIN_H)/2),
		rl.NewVector2(0, 0),
		0,
		1,
	)

	if len(os.Args) > 1 {
		fload := os.Args[1]
		level.Load(fload)
		screen.Scr.Type = screen.SCREditor // if we lod then it be editing
	}

	for !rl.WindowShouldClose() {
		// global handling for all screens
		if rl.IsWindowResized() {
			utils.WIN_W = int32(rl.GetScreenWidth())
			utils.WIN_H = int32(rl.GetScreenHeight())
		}
		// handle current screen
		screen.Scr.Handle(&cam)
	}
}
