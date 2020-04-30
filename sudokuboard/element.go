package sudokuboard

func (b *Board) getAbsCoords(x int, y int) (int, int, int, int, error) {
	err := b.checkCoordErrors(x, y)
	if err != nil {
		return -1, -1, -1, -1, err
	}
	numOfChunkInRow := x / b.Dim
	numOfElemInRow := x % b.Dim
	numOfChunkInCol := y / b.Dim
	numOfElemInCol := y % b.Dim
	return numOfChunkInRow, numOfElemInRow, numOfChunkInCol, numOfElemInCol, nil
}

func (b *Board) getElementByCoords(x int, y int) (int, error) {
	numOfChunkInRow, numOfElemInRow, numOfChunkInCol, numOfElemInCol, err := b.getAbsCoords(x, y)
	if err != nil {
		return -1, err
	}
	return b.Chunks[numOfChunkInRow][numOfChunkInCol].Elements[numOfElemInRow][numOfElemInCol], nil
}

func (b *Board) setElementByCoords(x int, y int, value int) error {
	numOfChunkInRow, numOfElemInRow, numOfChunkInCol, numOfElemInCol, err := b.getAbsCoords(x, y)
	if err == nil {
		b.Chunks[numOfChunkInRow][numOfChunkInCol].Elements[numOfElemInRow][numOfElemInCol] = value
		return err
	}
	return nil
}
