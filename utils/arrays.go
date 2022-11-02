package utils

import "math"

func ChunkBy[T comparable](quantityOfItemsPerChunk int, items []T) [][]T {
	chunks := [][]T{}
	quantityOfChunks := math.Ceil(float64(len(items)) / float64(quantityOfItemsPerChunk))
	for v := 0; v < int(quantityOfChunks); v++ {
		chunks = append(chunks, []T{})
	}

	arrIndex := 0
	discriminant := 0
	for _, item := range items {
		if discriminant >= quantityOfItemsPerChunk {
			discriminant = 0
			arrIndex += 1
		}

		chunks[arrIndex] = append(chunks[arrIndex], item)
		discriminant += 1
	}

	return chunks
}
