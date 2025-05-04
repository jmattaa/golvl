package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jmattaa/golvl/screen"
	"github.com/jmattaa/golvl/utils"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(utils.WIN_W, utils.WIN_H, "golvl - v"+utils.VERSION)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// global handling for all screens
		if rl.IsWindowResized() {
			utils.WIN_W = int32(rl.GetScreenWidth())
			utils.WIN_H = int32(rl.GetScreenHeight())
		}

		// handle current screen
		screen.Scr.Handle()
	}
}
