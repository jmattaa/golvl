package screen

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jmattaa/golvl/level"
)

var panSpeed float32 = 10

var lastTileX, lastTileY = -1, -1

func HandleEditor(cam *rl.Camera2D) {
	if rl.IsKeyPressed(rl.KeyE) {
		level.ExportLvl()
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		cam.Offset.X -= panSpeed
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		cam.Offset.X += panSpeed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		cam.Offset.Y -= panSpeed
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		cam.Offset.Y += panSpeed
	}

	mouseWheel := rl.GetMouseWheelMove()

	if mouseWheel > 0 && cam.Zoom < 10 {
		cam.Zoom += .1
	}
	if mouseWheel < 0 && cam.Zoom > .01 {
		cam.Zoom -= .1
	}

	mouse := rl.GetMousePosition()
	worldMouse := rl.GetScreenToWorld2D(mouse, *cam)

	rl.BeginDrawing()
	rl.BeginMode2D(*cam)
	rl.ClearBackground(rl.RayWhite)

	tileSize := int32(32)
	width := level.LevelWidth
	height := level.LevelHeight

	if len(level.Tilemap) == 0 {
		level.Tilemap = make([][]int, height)
		for y := range level.Tilemap {
			level.Tilemap[y] = make([]int, width)
		}
	}

	mouseTileX := -1
	mouseTileY := -1

	for y := range height {
		for x := range width {
			xPos := int32(x) * tileSize
			yPos := int32(y) * tileSize
			rect := rl.NewRectangle(float32(xPos), float32(yPos), float32(tileSize), float32(tileSize))

			color := rl.LightGray
			switch level.Tilemap[y][x] {
			case 1:
				color = rl.Blue
			case 2:
				color = rl.Green
			case 3:
				color = rl.Red
			}

			rl.DrawRectangleRec(rect, color)
			rl.DrawRectangleLines(xPos, yPos, tileSize, tileSize, rl.DarkGray)

			if rl.CheckCollisionPointRec(worldMouse, rect) {
				mouseTileX = x
				mouseTileY = y
			}
		}
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) && mouseTileX != -1 && mouseTileY != -1 {
		if mouseTileX != lastTileX || mouseTileY != lastTileY {
			level.Tilemap[mouseTileY][mouseTileX] = (level.Tilemap[mouseTileY][mouseTileX] + 1) % level.NumTileTypes
			lastTileX = mouseTileX
			lastTileY = mouseTileY
		}
	} else if rl.IsMouseButtonDown(rl.MouseRightButton) &&
		mouseTileX != -1 &&
		mouseTileY != -1 {
		level.Tilemap[mouseTileY][mouseTileX] = 0
		lastTileX = mouseTileX
		lastTileY = mouseTileY
	} else {
		lastTileX = -1
		lastTileY = -1
	}

	rl.EndMode2D()
	rl.EndDrawing()
}
