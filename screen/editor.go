package screen

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jmattaa/golvl/level"
)

var cameraOffset = rl.NewVector2(0, 0)
var panSpeed float32 = 10

var lastTileX, lastTileY = -1, -1

func HandleEditor() {
	if rl.IsKeyPressed(rl.KeyE) {
		level.ExportLvl()
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	tileSize := int32(level.TileSize)
	width := level.LevelWidth
	height := level.LevelHeight

	if len(level.Tilemap) == 0 {
		level.Tilemap = make([][]int, height)
		for y := range level.Tilemap {
			level.Tilemap[y] = make([]int, width)
		}
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		cameraOffset.X -= panSpeed
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		cameraOffset.X += panSpeed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		cameraOffset.Y -= panSpeed
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		cameraOffset.Y += panSpeed
	}

	mouse := rl.GetMousePosition()

	mouseTileX := -1
	mouseTileY := -1

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			xPos := int32(x)*tileSize + int32(cameraOffset.X)
			yPos := int32(y)*tileSize + int32(cameraOffset.Y)
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

			if rl.CheckCollisionPointRec(mouse, rect) {
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
	} else {
		lastTileX = -1
		lastTileY = -1
	}

	rl.EndDrawing()
}
