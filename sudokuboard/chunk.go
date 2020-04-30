package sudokuboard

type Chunk struct {
	// set of elements
	Elements [][]int
}

func initElementsForChunk(dim int) [][]int {
	result := make([][]int, dim)
	for i := range result {
		shard := make([]int, dim)
		result[i] = shard
	}
	return result
}

func NewChunk(dim int) Chunk {
	chunk := Chunk{
		Elements: initElementsForChunk(dim),
	}
	return chunk
}

func makeChunkArray(dim int) [][]Chunk {
	result := make([][]Chunk, dim)
	for i := range result {
		shard := make([]Chunk, dim)
		for j := range shard {
			shard[j] = NewChunk(dim)
		}
		result[i] = shard
	}
	return result
}
