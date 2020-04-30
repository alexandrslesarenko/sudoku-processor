package sudokuboard

import (
	"fmt"
	"os"
)

func (b *Board) printLine() {
	for i := 0; i < b.Dim*b.Dim*3+b.Dim; i++ {
		fmt.Print("_")
	}
	fmt.Println("_")
}

func (b *Board) Print() {
	for y := 0; y < b.Dim*b.Dim; y++ {
		if y%b.Dim == 0 {
			b.printLine()
		}
		for x := 0; x < b.Dim*b.Dim; x++ {
			if x%b.Dim == 0 {
				fmt.Print("|")
			}
			element, err := b.getElementByCoords(x, y)
			if err == nil {
				fmt.Printf("%3d", element)
			} else {
				fmt.Println(err)
				os.Exit(-1)
			}
		}
		fmt.Println("|")
	}
	b.printLine()
}
