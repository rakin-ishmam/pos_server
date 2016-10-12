package data

// Spliter split string
func Spliter(fields ...string) []string {
	sps := []string{}

	for _, field := range fields {
		for i := 0; i < len(field); i++ {
			for j := i + 1; j <= len(field); j++ {
				sps = append(sps, field[i:j])
			}
		}
	}

	return sps
}
