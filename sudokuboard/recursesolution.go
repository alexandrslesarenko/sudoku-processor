package sudokuboard

import "fmt"

func (b *Board) Solved() (bool, error) {
	for y := 0; y < b.Dim*b.Dim; y++ {
		for x := 0; x < b.Dim*b.Dim; x++ {
			element, err := b.getElementByCoords(x, y)
			if err != nil || element == 0 {
				return false, err
			}
		}
	}
	return true, nil
}

func (b *Board) GetRecurseSolution(showProgress bool) (*Board, bool, error) {
	solved, err := b.Solved()
	if err != nil {
		return nil, false, err
	} else if solved {
		return b, true, nil
	}
	for y := 0; y < b.Dim*b.Dim; y++ {
		for x := 0; x < b.Dim*b.Dim; x++ {
			element, err := b.getElementByCoords(x, y)
			if err != nil {
				return nil, false, err
			} else {
				if element == 0 {
					newBoard, err := b.Clone()
					if err == nil {
						for i := 1; i <= b.Dim*b.Dim; i++ {
							err = newBoard.setElementByCoords(x, y, i)
							if showProgress {
								b.Print()
								fmt.Printf("%c[%d;%df", 0x1B, 0, 0)
							}
							if err != nil {
								return nil, false, err
							} else {
								elementChecked, err := newBoard.CheckElement(x, y)
								if err != nil {
									return nil, false, err
								} else {
									if elementChecked {
										resultBoard, solved, err := newBoard.GetRecurseSolution(showProgress)
										if solved {
											return resultBoard, solved, err
										}
									}
								}
							}
						}
						// Board has no solutions because no one value is right for current element
						return nil, false, nil
					} else {
						return nil, false, err
					}
				}
			}
		}
	}
	// all variants was fail, no solutions found
	return nil, false, nil
}
