package screen

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jmattaa/golvl/level"
)

var (
	filenameBuf    = []rune{}
	tileSizeBuf    = []rune{}
	levelWidthBuf  = []rune{}
	levelHeightBuf = []rune{}
	selectedInput  int
	inputError     bool
)

const maxInputLength = 20

func HandleMenu() {
	screenW := float32(rl.GetScreenWidth())
	screenH := float32(rl.GetScreenHeight())
	xStart := screenW * 0.1
	yStart := screenH * 0.2
	lineSpacing := float32(100)

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	drawInputBox("Filename", &filenameBuf, int32(xStart), int32(yStart+lineSpacing*0), 0)
	drawInputBox(
		"Tile Size(dosent change shi it's basically the zoom of the editor, like 32 or sum is nice)",
		&tileSizeBuf, int32(xStart),
		int32(yStart+lineSpacing*1),
		1,
	)
	drawInputBox("Level Width", &levelWidthBuf, int32(xStart), int32(yStart+lineSpacing*2), 2)
	drawInputBox("Level Height", &levelHeightBuf, int32(xStart), int32(yStart+lineSpacing*3), 3)

	if inputError {
		rl.DrawText("Invalid number input!", int32(xStart), int32(yStart+lineSpacing*4), 20, rl.Red)
	}

	buttonW := float32(120)
	buttonH := float32(40)
	buttonX := xStart + 120
	buttonY := yStart + lineSpacing*5

	drawDoneButton(buttonX, buttonY, buttonW, buttonH, handleDone)

	// info
	rl.DrawText(
		"When done editing in the next screen press 'e' to export",
		int32(xStart),
		int32(yStart+lineSpacing*6),
		20,
		rl.Black,
	)

	rl.EndDrawing()

	handleInput()
}

func handleInput() {
	if rl.IsKeyPressed(rl.KeyDown) {
		selectedInput = (selectedInput + 1) % 4
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		selectedInput = (selectedInput + 3) % 4
	}

	char := rl.GetCharPressed()
	if char > 0 {
		r := rune(char)
		switch selectedInput {
		case 0:
			if len(filenameBuf) < maxInputLength {
				filenameBuf = append(filenameBuf, r)
			}
		case 1, 2, 3:
			if r >= '0' && r <= '9' {
				buf := getBuffer(selectedInput)
				if len(*buf) < maxInputLength {
					*buf = append(*buf, r)
				}
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyBackspace) {
		buf := getBuffer(selectedInput)
		if len(*buf) > 0 {
			*buf = (*buf)[:len(*buf)-1]
		}
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		handleDone()
	}
}

func handleDone() {
	inpsFilled := len(filenameBuf) > 0 && len(tileSizeBuf) > 0 && len(levelWidthBuf) > 0 && len(levelHeightBuf) > 0
	if !inpsFilled {
		inputError = true
		return
	}

	tileSize, err1 := strconv.Atoi(string(tileSizeBuf))
	width, err2 := strconv.Atoi(string(levelWidthBuf))
	height, err3 := strconv.Atoi(string(levelHeightBuf))

	if err1 == nil && err2 == nil && err3 == nil {
		level.Filename = string(filenameBuf)
		level.TileSize = tileSize
		level.LevelWidth = width
		level.LevelHeight = height
		inputError = false
		Scr.Type = SCREditor
	} else {
		inputError = true
	}
}

func drawInputBox(label string, buffer *[]rune, x, y int32, index int) {
	boxWidth := int32(300)
	boxHeight := int32(30)

	rect := rl.NewRectangle(float32(x), float32(y+boxHeight), float32(boxWidth), float32(boxHeight))
	rl.DrawText(label, x, y, 20, rl.Black)

	color := rl.Gray
	if selectedInput == index {
		color = rl.LightGray
	}
	rl.DrawRectangleRec(rect, color)
	rl.DrawRectangleLines(int32(rect.X), int32(rect.Y), int32(rect.Width), int32(rect.Height), rl.DarkGray)

	rl.DrawText(string(*buffer), int32(rect.X)+5, int32(rect.Y)+5, 20, rl.Black)
}

func drawDoneButton(x, y, w, h float32, onClick func()) {
	mouseX := float32(rl.GetMouseX())
	mouseY := float32(rl.GetMouseY())
	mouseRect := rl.NewRectangle(mouseX, mouseY, 1, 1)
	buttonRect := rl.NewRectangle(x, y, w, h)

	buttonColor := rl.LightGray
	if rl.CheckCollisionRecs(buttonRect, mouseRect) {
		buttonColor = rl.Gray
	}

	rl.DrawRectangleRec(buttonRect, buttonColor)
	rl.DrawRectangleLines(int32(x), int32(y), int32(w), int32(h), rl.DarkGray)
	rl.DrawText("Done", int32(x+30), int32(y+10), 20, rl.Black)

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionRecs(buttonRect, mouseRect) {
		onClick()
	}
}

func getBuffer(index int) *[]rune {
	switch index {
	case 0:
		return &filenameBuf
	case 1:
		return &tileSizeBuf
	case 2:
		return &levelWidthBuf
	case 3:
		return &levelHeightBuf
	}
	return nil
}
