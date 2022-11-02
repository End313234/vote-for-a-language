package utils

func MergeMaps[K comparable, V any](ma map[K]V, mb map[K]V) map[K]V {
	for k, v := range ma {
		mb[k] = v
	}

	return mb
}
