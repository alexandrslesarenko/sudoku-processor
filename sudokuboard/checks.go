package sudokuboard

import (
	"github.com/alexandrslesarenko/sudoku-processor/errors"
	"strconv"
)

func (b *Board) checkElementByItsRow(x int, y int) (bool, error) {
	element, err := b.getElementByCoordsAndCheckNonEmpty(x, y)
	if err != nil {
		return false, err
	}
	for i := 0; i < b.Dim*b.Dim; i++ {
		if i != x {
			currElement, err := b.getElementByCoords(i, y)
			if err != nil {
				return false, err
			}
			if currElement == element {
				return false, nil
			}
		}
	}
	return true, nil
}

func (b *Board) checkElementByItsCol(x int, y int) (bool, error) {
	element, err := b.getElementByCoordsAndCheckNonEmpty(x, y)
	if err != nil {
		return false, err
	}
	for i := 0; i < b.Dim*b.Dim; i++ {
		if i != y {
			currElement, err := b.getElementByCoords(x, i)
			if err != nil {
				return false, err
			}
			if currElement == element {
				return false, nil
			}
		}
	}
	return true, nil
}

func (b *Board) checkElementByItsChunk(x int, y int) (bool, error) {
	element, err := b.getElementByCoordsAndCheckNonEmpty(x, y)
	if err != nil {
		return false, err
	}
	numOfChunkInRow := x / b.Dim
	numOfChunkInCol := y / b.Dim
	for iy := 0; iy < b.Dim; iy++ {
		for ix := 0; ix < b.Dim; ix++ {
			if !(x%b.Dim == ix && y%b.Dim == iy) {
				currElement := b.Chunks[numOfChunkInRow][numOfChunkInCol].Elements[ix][iy]
				if currElement == element {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

func (b *Board) CheckElement(x int, y int) (bool, error) {
	checkFlag, err := b.checkElementByItsCol(x, y)
	if err == nil {
		if checkFlag {
			checkFlag, err = b.checkElementByItsRow(x, y)
			if err == nil {
				if checkFlag {
					checkFlag, err = b.checkElementByItsChunk(x, y)
					if err == nil {
						return checkFlag, nil
					}
				}
			}
		}
	}
	return false, err
}

func (b *Board) getElementByCoordsAndCheckNonEmpty(x int, y int) (int, error) {
	element, err := b.getElementByCoords(x, y)
	if err != nil {
		return -1, err
	} else {
		if element == 0 {
			return -1, errors.CommonError("Element value must be defined")
		}
	}
	return element, nil
}

func (b *Board) checkCoordErrors(x int, y int) error {
	if x < 0 || x >= b.Dim*b.Dim {
		msg := "X coordinate is out of range: Dim=" + strconv.Itoa(b.Dim) + "; X=" + strconv.Itoa(x)
		return errors.CommonError(msg)
	}
	if y < 0 || y >= b.Dim*b.Dim {
		msg := "Y coordinate is out of range: Dim=" + strconv.Itoa(b.Dim) + "; Y=" + strconv.Itoa(y)
		return errors.CommonError(msg)
	}
	return nil
}
