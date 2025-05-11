package level

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmattaa/golvl/utils"
)

func Export() {
	f, err := os.Create(Filename)
	if err != nil {
		println("error creating file", err.Error())
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

func Load(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		println("error opening file", err.Error())
		return
	}
	defer f.Close()
	Filename = fname // esta muy importante 🤔

	fmt.Fscanf(f, "%d\n", &LevelWidth)
	fmt.Fscanf(f, "%d\n", &LevelHeight)

	if LevelWidth <= 0 || LevelHeight <= 0 {
		println("invalid level size", LevelWidth, LevelHeight)
		os.Exit(1)
	}

	println("loading level", LevelWidth, "x", LevelHeight)

	var ch rune

	Tilemap = make([][]int, LevelHeight)
	for i := range LevelHeight {
		Tilemap[i] = make([]int, LevelWidth)
		for j := range LevelWidth {
			fmt.Fscanf(f, "%c", &ch)
			ch = ch - '0'
			if ch < 0 && ch > 9 {
				fmt.Printf("cannot load file, invalid char %c\n", ch)
				os.Exit(1)
			}
			Tilemap[i][j] = int(ch)
		}
		fmt.Fscanf(f, "%c", &ch)
		if ch != '\n' && ch != utils.EOF {
			fmt.Printf("cannot load file, invalid char %c\n", ch)
			os.Exit(1)
		}
	}
}
