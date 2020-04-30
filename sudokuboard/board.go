package sudokuboard

type Board struct {
	Chunks [][]Chunk
	Dim    int
	x      int
	y      int
	solved bool
	valid  bool
}

func New(dim int) (Board, error) {
	newBoard := Board{
		Dim:    dim,
		Chunks: makeChunkArray(dim),
	}
	return newBoard, nil
}

func (b *Board) Clone() (*Board, error) {
	newBoard, err := New(b.Dim)
	if err == nil {
		for y := 0; y < b.Dim*b.Dim; y++ {
			for x := 0; x < b.Dim*b.Dim; x++ {
				curElement, err := b.getElementByCoords(x, y)
				if err == nil {
					err = newBoard.setElementByCoords(x, y, curElement)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			}
		}
	} else {
		return nil, err
	}
	return &newBoard, nil
}
