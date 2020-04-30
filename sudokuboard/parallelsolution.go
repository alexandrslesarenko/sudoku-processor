package sudokuboard

import "fmt"

func boardValidatorWorker(job <-chan *Board, result chan<- *Board) {
	for b := range job {
		b.valid, _ = b.CheckElement(b.x, b.y)
		b.solved = true
		for y := 0; y < b.Dim*b.Dim; y++ {
			for x := 0; x < b.Dim*b.Dim; x++ {
				v, _ := b.getElementByCoords(x, y)
				if v == 0 {
					b.solved = false
					break
				}
				if !b.solved {
					break
				}
			}
		}
		result <- b
	}
}

func (b *Board) getFirstEmptyElementCoords() (bool, int, int) {
	for y := 0; y < b.Dim*b.Dim; y++ {
		for x := 0; x < b.Dim*b.Dim; x++ {
			v, _ := b.getElementByCoords(x, y)
			if v == 0 {
				return false, x, y
			}
		}
	}
	return true, -1, -1
}

func (b *Board) GetParallelSolution(showProgress bool) (*Board, bool, error) {
	jobs := make(chan *Board)
	results := make(chan *Board, b.Dim*b.Dim)
	var resultList []*Board
	for w := 1; w <= b.Dim*b.Dim; w++ {
		go boardValidatorWorker(jobs, results)
	}

	// getting first empty cell if it exists
	solved, x, y := b.getFirstEmptyElementCoords()
	if solved {
		return b, true, nil
	}

	// lets get valid variants of first empty cell
	for i := 1; i <= b.Dim*b.Dim; i++ {
		newBoard, err := b.Clone()
		if err == nil {
			newBoard.solved = false
			newBoard.valid = false
			newBoard.x = x
			newBoard.y = y
			err = newBoard.setElementByCoords(x, y, i)
			if err == nil {
				jobs <- newBoard
				validatedBoard := <-results
				if validatedBoard.solved {
					return validatedBoard, true, nil
				}
				if validatedBoard.valid {
					resultList = append(resultList, validatedBoard)
				}
			}
		} else {
			return nil, false, err
		}
	}
	// if board has no solutions
	if resultList == nil {
		b.valid = false
		return b, false, nil
	}

	for i := range resultList {
		if resultList[i].solved {
			return resultList[i], true, nil
		}
		if showProgress {
			resultList[i].Print()
			fmt.Printf("%c[%d;%df", 0x1B, 0, 0)
		}
		x, solved, err := resultList[i].GetParallelSolution(showProgress)
		if err == nil {
			if solved {
				return x, true, nil
			}
		}

	}
	return nil, false, nil
}
