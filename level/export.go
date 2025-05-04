package level

import (
	"os"
	"strconv"
)

func ExportLvl() {
	f, err := os.Create(Filename + ".golvl")
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(strconv.Itoa(LevelWidth) + "\n")
	f.WriteString(strconv.Itoa(LevelHeight) + "\n")

	for _, row := range Tilemap {
		for _, tile := range row {
			f.WriteString(strconv.Itoa(tile))
		}
		f.WriteString("\n")
	}
}
