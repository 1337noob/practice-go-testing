package invertmap

func InvertMap(m map[string]int) map[int]string {
	inverted := make(map[int]string)
	for k, v := range m {
		inverted[v] = k
	}

	return inverted
}
